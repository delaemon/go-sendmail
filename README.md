# go-sendmail

##Description
Simple sendmail command.
Settings can be specified from the write to a config/default.toml or command line.

###Usage
```
usage:
    bin/OS_ARC/sendmail [option] (config/default.toml)

option:
	-u, --user 			smtp login user
	-p, --password 		smtp login password
	-h, --host			smtp server host
	-P, --port 			stmp server port
	-f, --from			email sender
	-t, --to 			email recipient
	-s, --subject 		email subject
	-a, --attach        email attach files (file1,file2,...)
	-c, --content-type	email body content-type (default: text/html)
	-b, --body 			email body (message body or require file path)
	--show				view config
	--help			 	view usage
```

###Example
####apply config/default.toml only
```
./bin/${OS_ARC}/sendmail
```
####overwrite config/default.toml
```
./bin/${OS_ARC}/sendmail \
    -u account@gmail.com \
    -p password \
    -h smtp.gmail.com \
    -P 587 \
    -f sender@example.org \
    -t recipient@example.net \
    -s "Default Html Mail" \
    -a "/image1.png,image2.png" \
    -b "<html> ~ </html>"
```    
####text mail
```
./bin/${OS_ARC}/sendmail \
    -u account@gmail.com \
    -p password \
    -h smtp.gmail.com \
    -P 587 \
    -f sender@example.org \
    -t recipient@example.net \
    -s "It's Text Mail" \
    -c "text/plain"
    -b "Hello"
```
####note
"./bin/${OS_ARC}/sendmail" possible is replaced with "go run sendamail.go"
And if to describe any of the config to a default.toml set on the command line is not required.
