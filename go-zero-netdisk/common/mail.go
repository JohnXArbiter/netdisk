package common

import "gopkg.in/gomail.v2"

type Email struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	SenderName string `json:"senderName"`
}

func (email *Email) SendCode(to, code string) error {
	m := gomail.NewMessage(
		gomail.SetEncoding(gomail.Base64),
	)
	m.SetHeader("From", m.FormatAddress(email.Username, email.SenderName))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "咪咪网盘注册验证码")
	m.SetBody("text/html",
		`<h2>咪咪网盘</h2>
			  <span>将在5分钟内过期</span>
			  <span style="font-size: 30px; 
					padding: 10px;
					background-color:wheat">`+
			code+
			`</span>`)

	d := gomail.NewDialer(email.Host, email.Port, email.Username, email.Password)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d.DialAndSend(m)
}
