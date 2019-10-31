package email

import (
	"gopkg.in/gomail.v2"
	"strconv"
)


//最后一个参数 merchant 商户号 不知道的话 找陈伟民
func Send(title, content, touser, merchant string) error {
	//TODO 暂不做调用
	return nil;
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "xxx@vip.qq.com",
		"pass": "qjqngvjrkrscbhhe",
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", merchant+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", touser)                            //发送给多个用户
	m.SetHeader("Subject", title)                         //设置邮件主题
	m.SetBody("text/html", content)                            //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}
