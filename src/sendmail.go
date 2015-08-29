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
	config Config
	smtp_user *string
	smtp_pwd  *string
	smtp_host *string
	smtp_port *string
	from      *string
	to		  *string
	subject   *string
	body 	  *string
	mode	  *string
)

func stream() {
	c, err := smtp.Dial(*smtp_host + ":" + *smtp_port)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Mail(*from)
	c.Rcpt(*to)
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString(*body)
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}

func sendmail() {
	auth := smtp.PlainAuth("", *smtp_user, *smtp_pwd, *smtp_host)
	msg := []byte(
	"To: " + *to + "\r\n" +
	"Subject: " + *subject + "\r\n" +
	"\r\n" +
	*body + "\r\n")
	err := smtp.SendMail(*smtp_host + ":" + *smtp_port, auth, *from, []string{*to}, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	out := `
usage:
	go run src/sendmail.go [option] (default-setting: config/default.toml)

option:
	-u 	smtp login user
	-p 	smtp login password
	-h 	smtp server host
	-P 	stmp server port
	-f 	email sender
	-t 	email recipient
	-s 	email subject
	-b 	email body
	-m 	send mode(sendmail|stream|config)

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
		"user: " 	+ *smtp_user + "\n" +
		"pwd: "  	+ *smtp_pwd  + "\n" +
		"host: " 	+ *smtp_host + "\n" +
		"port: " 	+ *smtp_port + "\n" +
		"[mail]\n" 	+
		"from: " 	+ *from	    + "\n" +
		"to: " 		+ *to 	    + "\n" +
		"subject: " + *subject	+ "\n" +
		"body: " 	+ *body 	+ "\n")
}

func setFlag() {
	smtp_user = flag.String("u", config.Smtp.User, "smtp login user")
	smtp_pwd  = flag.String("p", config.Smtp.Pwd,  "smtp login password")
	smtp_host = flag.String("h", config.Smtp.Host, "smtp server host")
	smtp_port = flag.String("P", config.Smtp.Port, "stmp server port")
	from      = flag.String("f", config.From,      "email sender")
	to        = flag.String("t", config.To,        "email recipient")
	subject   = flag.String("s", config.Subject,   "email subject")
	body      = flag.String("b", config.Body,      "email body")
	mode      = flag.String("m", config.Mode,      "send mode(sendmail|stream|config)")
}

func main() {
	if _, err := toml.DecodeFile("config/default.toml", &config); err != nil {
		log.Fatal(err)
		return
	}
	setFlag()
	flag.Parse()

	switch (*mode){
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
