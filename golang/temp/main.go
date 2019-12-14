package main

import (
   "github.com/gin-gonic/gin"
   "net/http/httputil"
   "net/http"
   "fmt"
)


const Host  = "127.0.0.1:8080"

var simpleHostProxy = httputil.ReverseProxy{
   Director: func(req *http.Request) {
      req.URL.Scheme = "http"
      req.URL.Host = Host
      req.Host = Host
   },
}

func main() {
   engine := gin.New()
   vi := engine.Group("")
   vi.Any("/*action", WithHeader)
   err := engine.Run(":80")
   if err != nil {
      fmt.Println(err)
   }
}

func WithHeader(ctx *gin.Context) {

   ctx.Request.Header.Add("requester-uid", "id")
   simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
}
