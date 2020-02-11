package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	srv "micro/rpc/srv/proto/srv"
)

type Srv struct{}

func (e *Srv) Handle(ctx context.Context, msg *srv.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *srv.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
