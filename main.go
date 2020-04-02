package main

import (
	"context"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"flag"
	"fmt"
	"os"
)

const (
	nagiosOK	int = 0
	nagiosWarning	int = 1
	nagiosCritical	int = 2
	nagiosUnknown	int = 3
)
var nagStr = []string {"OK","WARNING","CRITICAL","UNKNOWN"}
type param struct {
	hostname	string
	secret		string
	username 	string
	password 	string
}

var cfg param

func init() {
	flag.StringVar(&cfg.hostname, "hostname", "127.0.0.1:1812", "Radius server and port")
	flag.StringVar(&cfg.secret,   "secret",   "secret123",      "Radius secret")
	flag.StringVar(&cfg.username, "username", "john.doe",  "Username")
	flag.StringVar(&cfg.password, "password", "12345",     "Password")
	flag.Parse()
}
func main() {
	ret:=nagiosUnknown
	msg:=""
	packet := radius.New(radius.CodeAccessRequest, []byte(cfg.secret))
	rfc2865.UserName_SetString(packet, cfg.username)
	rfc2865.UserPassword_SetString(packet, cfg.password)
	response, err := radius.Exchange(context.Background(), packet, cfg.hostname)
	if err != nil {
		ret=nagiosCritical
		msg=fmt.Sprintf("%s",err)
	} else {
		if response.Code==radius.CodeAccessAccept {
			ret=nagiosOK
			msg="Access-Accept"
		} else if response.Code==radius.CodeAccessReject {
			ret=nagiosCritical
			msg="Access-Reject"
		}
	}
	fmt.Printf("%s|%s\n",nagStr[ret],msg)
	os.Exit(ret)
}
