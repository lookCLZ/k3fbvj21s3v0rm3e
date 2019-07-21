package main

import (
	"fmt"
	"os"
	"net/http"
	"net/url"
)

func OpenSendFile(fName string,w http.ResponseWriter) {
	pathFileName:="/Users/liuhongrui/Downloads"+fName
	f,_:=os.Open(pathFileName)
	defer f.Close()

	buf:=make([]byte,4096)
	for {
		n,_:=f.Read(buf)
		if n == 0 {
			return 
		}
		w.Write(buf[:n])
	}
}

func myHandler(w http.ResponseWriter,r *http.Request) {
	decodedUrl, _ := url.QueryUnescape(r.URL.String())
	fmt.Println("客户端请求，",decodedUrl)
	OpenSendFile(decodedUrl,w)
}

func main(){
	http.HandleFunc("/",myHandler)
	http.ListenAndServe("127.0.0.1:8000",nil)

}