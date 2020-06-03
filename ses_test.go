package goawsses

import "testing"

func TestService_Send(t *testing.T) {
	key := "xxx"
	secret := "xxx"
	region := "us-west-2"
	from := "" // 需要已验证，否则会报错：Email address is not verified

	s, err := NewService(key, secret, region, from)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	to := "xxx@163.com"
	subject := "测试 Test"
	body := "<h1>亲爱的：</h1><br><p>这是测试邮件，test email，您的验证码是：584125</p>"

	msgId, err := s.Send(to, subject, body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(msgId)
}
