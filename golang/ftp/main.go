package main

import (
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	c, err := ftp.Dial("123.57.36.41:222", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("FTP_LiuHR", "LiuHRDF93jdm359d1@#dn2")
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the FTP conn

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
