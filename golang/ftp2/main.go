package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"gopkg.in/dutchcoders/goftp.v1"
)

func main() {

	var err error
	var ftp *goftp.FTP
	if ftp, err = goftp.Connect("123.57.36.41:222"); err != nil {
		panic(err)
	}

	defer ftp.Close()
	fmt.Println("Successfully connected to", ftp)

	// Username/password authentication

	if err = ftp.Login("FTP_LiuHR", "LiuHRDF93jdm359d1@#dn2"); err != nil {

		panic(err)

	}

	if err = ftp.Cwd("/"); err != nil {

		panic(err)

	}

	var curpath string

	if curpath, err = ftp.Pwd(); err != nil {

		panic(err)

	}

	fmt.Printf("Current path: %s", curpath)

	// Upload a file

	var file *os.File

	if file, err = os.Open("/Users/liuhongrui/Desktop/xxx.txt"); err != nil {
		panic(err)
	}
// Data connection already open; Transfer starting
	if err := ftp.Stor("/FTP_LiuHR/test233.txt", file); err != nil {
		panic(err)
	}

	// Download each file into local memory, and calculate it's sha256 hash

	err = ftp.Walk("/", func(path string, info os.FileMode, err error) error {

		_, err = ftp.Retr(path, func(r io.Reader) error {

			var hasher = sha256.New()

			if _, err = io.Copy(hasher, r); err != nil {

				return err

			}

			hash := fmt.Sprintf("%s %x", path, hex.EncodeToString(hasher.Sum(nil)))

			fmt.Println(hash)

			return err

		})

		return nil

	})

}
