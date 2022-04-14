package utils

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "xxx@qq.com",
		"pass": "xyz",
		"host": "smtp.qq.com",
		"port": "465",
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", "Green Studio"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                                 //发送给多个用户
	m.SetHeader("Subject", subject)                              //设置邮件主题
	m.SetBody("text/html", body)                                 //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}

func Mail(reciever string, res bool) {
	//定义收件人
	mailTo := []string{
		reciever,
	}

	//邮件主题为
	subject := "阳光晨跑结果通知"
	// 邮件正文
	var body string

	if res {
		//跑步成功
		body = "跑步成功"
		SendMail(mailTo, subject, body)
	} else {
		//跑步失败
		body = "跑步失败"
		SendMail(mailTo, subject, body)
	}

}
