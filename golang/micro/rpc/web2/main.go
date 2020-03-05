package main

import (
        "github.com/micro/go-micro/v2/util/log"
	"net/http"

        "github.com/micro/go-micro/v2/web"
        "micro/rpc/web2/handler"
)

func main() {
	// create new web service
        service := web.NewService(
                web.Name("go.micro.web.web2"),
                web.Version("latest"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/web2/call", handler.Web2Call)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
