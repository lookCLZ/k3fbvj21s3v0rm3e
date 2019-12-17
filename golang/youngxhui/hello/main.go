package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router:=gin.Default()
	router.GET("/",func(ctx *gin.Context){
		context.String(http.StatusOK,"hello gin")
	})
	fmt.Println("启动")
	router.Run()
}