package emailsender

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

type EmailSend struct {
	addr    string
	port    string
	myMail  string
	subject string
	auth    smtp.Auth
}

func NewEmailSend(HostAddress, HostPort, Username, Password, Subject, TestMail string) (*EmailSend, error) {
	emailSend := &EmailSend{
		addr:    HostAddress,
		port:    HostPort,
		auth:    smtp.PlainAuth("", Username, Password, HostAddress),
		myMail:  Username,
		subject: Subject,
	}
	// fmt.Println("test email send")
	// if err := emailSend.Send(TestMail, "test email"); err != nil {
	// 	fmt.Println("test email send failed")
	// 	return nil, err
	// }
	// fmt.Println("test email send success")
	return emailSend, nil
}

func (e *EmailSend) Send(email, emailCode string) error {
	fmt.Println("仅测试显示：", emailCode)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.qq.com",
	}
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", e.addr, e.port), tlsconfig)
	if err != nil {
		return fmt.Errorf("tls.Dial error: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, e.addr)
	if err != nil {
		return fmt.Errorf("smtp.NewClient error: %v", err)
	}

	// auth 认证
	if err = client.Auth(e.auth); err != nil {
		return fmt.Errorf("client.Auth error: %v", err)
	}

	// 设置发件人与收件人
	if err = client.Mail(e.myMail); err != nil {
		return fmt.Errorf("client.Mail error: %v", err)
	}
	if err = client.Rcpt(email); err != nil {
		return fmt.Errorf("client.Rcpt error: %v", err)
	}

	// 写入邮件内容
	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("client.Data error: %v", err)
	}
	defer wc.Close()
	_, err = wc.Write([]byte(fmt.Sprintf("你的验证码为： %s", emailCode)))
	if err != nil {
		return fmt.Errorf("wc.Write error: %v", err)
	}
	return nil
}
