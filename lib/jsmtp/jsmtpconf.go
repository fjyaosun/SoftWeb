package jsmtp

var SmtpConf struct {
	Password string
	Host     string
	User     string
	Default  struct {
		To      string
		Subject string
	}
}

func ConfigInit() {
	SmtpConf.Host = "smtp.126.com:25"
	SmtpConf.Password = "iyjgwykedbhsqvob"
	SmtpConf.User = "fjyaosun@126.com"
	SmtpConf.Default.To = "fjyaosun@163.com,584285135@qq.com"
	SmtpConf.Default.Subject = "Golang"
}
