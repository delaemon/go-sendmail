package main

import (
	"log"
	"net/smtp"
	"bytes"
	"flag"
	"github.com/BurntSushi/toml"
	"fmt"
)

type Smtp struct {
	User string
	Pwd  string
	Host string
	Port string
}

type Config struct {
	Smtp    Smtp
	From    string
	To      string
	Subject string
	Body    string
	Mode    string
}

var (
	config 	  Config
	smtp_user string
	smtp_pwd  string
	smtp_host string
	smtp_port string
	from      string
	to		  string
	subject   string
	body 	  string
	mode	  string
	help 	  bool
)

func stream() {
	c, err := smtp.Dial(smtp_host + ":" + smtp_port)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Mail(from)
	c.Rcpt(to)
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString(body)
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}

func sendmail() {
	auth := smtp.PlainAuth("", smtp_user, smtp_pwd, smtp_host)
	msg := []byte(
	"To: " + to + "\r\n" +
	"Subject: " + subject + "\r\n" +
	"\r\n" +
	body + "\r\n")
	err := smtp.SendMail(smtp_host + ":" + smtp_port, auth, from, []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	out := `
usage:
	go run src/sendmail.go [option] (default-setting: config/default.toml)

option:
	-u, --user 			smtp login user
	-p, --password 		smtp login password
	-h, --host			smtp server host
	-P, --Port 			stmp server port
	-f, --from			email sender
	-t, --to 			email recipient
	-s, --subject 		email subject
	-b, --body 			email body
	-m, --mode 			send mode(sendmail|stream|config)
	--help			 	view usage

example:
	go run src/sendmail.go \
		-u account@gmail.com \
		-p password \
		-h smtp.gmail.com \
		-P 587 \
		-f sender@example.org \
		-t recipient@example.net \
		-s "What's happening?" \
		-b "Read a book." \
		-m sendmail
	`
	fmt.Println(out)
}

func showConfig() {
	fmt.Println(
		"[smtp]\n"  +
		"user: " 	+ smtp_user + "\n" +
		"pwd: "  	+ smtp_pwd  + "\n" +
		"host: " 	+ smtp_host + "\n" +
		"port: " 	+ smtp_port + "\n" +
		"[mail]\n" 	+
		"from: " 	+ from	    + "\n" +
		"to: " 		+ to 	    + "\n" +
		"subject: " + subject	+ "\n" +
		"body: " 	+ body 	+ "\n")
}

func setFlag() {
	flag.StringVar(&smtp_user,	"u", 		config.Smtp.User, "smtp login user")
	flag.StringVar(&smtp_user,	"user", 	config.Smtp.User, "smtp login user")
	flag.StringVar(&smtp_pwd ,	"p", 		config.Smtp.Pwd,  "smtp login password")
	flag.StringVar(&smtp_pwd ,	"password", config.Smtp.Pwd,  "smtp login password")
	flag.StringVar(&smtp_host,	"h", 		config.Smtp.Host, "smtp server host")
	flag.StringVar(&smtp_host,	"host", 	config.Smtp.Host, "smtp server host")
	flag.StringVar(&smtp_port,	"P", 		config.Smtp.Port, "stmp server port")
	flag.StringVar(&smtp_port,	"Port", 	config.Smtp.Port, "stmp server port")
	flag.StringVar(&from     ,	"f", 		config.From,      "email sender")
	flag.StringVar(&from     ,	"from", 	config.From,      "email sender")
	flag.StringVar(&to       ,	"t", 		config.To,        "email recipient")
	flag.StringVar(&to       ,	"to", 		config.To,        "email recipient")
	flag.StringVar(&subject  ,	"s", 		config.Subject,   "email subject")
	flag.StringVar(&subject  ,	"subject", 	config.Subject,   "email subject")
	flag.StringVar(&body     ,	"b", 		config.Body,      "email body")
	flag.StringVar(&body     ,	"body", 	config.Body,      "email body")
	flag.StringVar(&mode     ,	"m", 		config.Mode,      "send mode(sendmail|stream|config)")
	flag.StringVar(&mode     ,	"mode", 	config.Mode,      "send mode(sendmail|stream|config)")
	flag.BoolVar  (&help     ,	"help",		false, 		   	  "View usage")
}

func main() {
	if _, err := toml.DecodeFile("config/default.toml", &config); err != nil {
		log.Fatal(err)
		return
	}
	setFlag()
	flag.Parse()

	if help {
		usage()
		return
	}

	switch (mode){
	case "strem":
		stream()
	case "sendmail":
		sendmail()
	case "config":
		showConfig()
	default:
		usage()
	}
}
