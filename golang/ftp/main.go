package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/jlaffaye/ftp"
)

func ConnectAndLogin() *ftp.ServerConn {
	conn, err := ftp.Dial("123.57.36.41" + ":" + "222")
	if err != nil {
		log.Println(err)
		return nil
	}
	if err := conn.Login("FTP_LiuHR", "LiuHRDF93jdm359d1@#dn2"); err != nil {
		log.Println(err)
		return nil
	}
	return conn
}

func main() {
	c := ConnectAndLogin()

	data := bytes.NewBufferString("Hello World")
	if err := c.Stor("/FTP_LiuHR/test2.txt", data); err != nil {
		fmt.Printf("%+v\n", err)
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}

// use of closed network connection
// 230 User FTP_LiuHR logged in.
// Remote system type is Windows_NT.
