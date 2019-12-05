package main

import (
	"fmt"
	"net/http"
	"os"
)

const port = "8000"

func main() {
	msg := "Starting main"
	a := 22

	fmt.Println(msg)
	fmt.Println(a)
}

func hi(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "HostName: %s", hostName)
}
