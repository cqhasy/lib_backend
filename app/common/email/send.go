package email

import (
	"AILN/app/common"
	"AILN/app/common/tool"
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"net/textproto"
)

const (
	Register = iota
	SetPassword
	SetEmail
)

type SenderConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
}

type Msg struct {
	Email string `json:"email"`
	Type  int    `json:"type"`
}

func (message *Msg) send() error {
	Email := message.Email
	Type := message.Type
	if Type != Register && Type != SetPassword && Type != SetEmail {
		return fmt.Errorf("invalid email usage")
	}
	//生成一个验证码
	randCode := tool.RandStringBytes(6)

	subject := fmt.Sprintf("护工 - %v", Type)

	emailInfo := common.CONFIG.StringMap("email")
	//发送
	e := &email.Email{
		To:      []string{Email},
		From:    emailInfo["username"],
		Subject: subject,
		HTML:    []byte(randCode),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send(emailInfo["host"]+":"+emailInfo["port"], smtp.PlainAuth("", emailInfo["username"], emailInfo["password"], emailInfo["host"]))
	if err != nil {
		return err
	}
	//存到redis

	return err
}

var queue chan *Msg

func Push(msg *Msg) error {
	select {
	case queue <- msg:
		return nil
	default:
		return errors.New("邮件发送失败")
	}
}

func Sender() {
	queue = make(chan *Msg, 100)
	var msg *Msg
	var err error
	for {
		msg = <-queue
		err = msg.send()
		if err != nil {
			log.Print("emailSender: send failed")
		}
	}
}
