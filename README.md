# go-sendmail

```
usage:
    go run src/sendmail.go [option] (config/default.toml)

option:
    -u  smtp login user
    -p  smtp login password
    -h  smtp server host
    -P  stmp server port
    -f  email sender
    -t  email recipient
    -s  email subject
    -b  email body
    -m  send mode(sendmail|stream|config)

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
```
