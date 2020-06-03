package goawsses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var DefaultCharSet = "UTF-8"

type Configs struct {
	AccessIdKey, AccessIdSecret, Region, FromAddress string
}

type Service struct {
	c Configs
	s *ses.SES
}

// Send 发送邮件，返回 msg id
func (s *Service) Send(to, subject, body string) (string, error) {
	// Assemble the to.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{
			},
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(DefaultCharSet),
					Data:    aws.String(body),
				},
				//Text: &ses.Content{
				//	Charset: aws.String(CharSet),
				//	Data:    aws.String(TextBody),
				//},
			},
			Subject: &ses.Content{
				Charset: aws.String(DefaultCharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(s.c.FromAddress),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the to.
	result, err := s.s.SendEmail(input)
	if err != nil {
		return "", err
	}

	return *result.MessageId, nil

	// Display error messages if they occur.
	//if err != nil {
	//	if aerr, ok := err.(awserr.Error); ok {
	//		switch aerr.Code() {
	//		case ses.ErrCodeMessageRejected:
	//			fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
	//		case ses.ErrCodeMailFromDomainNotVerifiedException:
	//			fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
	//		case ses.ErrCodeConfigurationSetDoesNotExistException:
	//			fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
	//		default:
	//			fmt.Println(aerr.Error())
	//		}
	//	} else {
	//		// Print the error, cast err to awserr.Error to get the Code and
	//		// Message from an error.
	//		fmt.Println(err.Error())
	//	}
	//
	//}
}

func NewService(key, secret, region, from string) (*Service, error) {
	c := Configs{
		AccessIdKey:    key,
		AccessIdSecret: secret,
		Region:         region,
		FromAddress:    from,
	}
	s, err := newSES(c.AccessIdKey, c.AccessIdSecret, c.Region)
	if err != nil {
		return nil, err
	}

	return &Service{
		c: c,
		s: s,
	}, nil
}

func newSES(key, secret, region string) (*ses.SES, error) {
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	})
	if err != nil {
		return nil, err
	}

	return ses.New(s), nil
}
