package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	srv "micro/rpc/srv/proto/srv"
)

type Srv struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Srv) Call(ctx context.Context, req *srv.Request, rsp *srv.Response) error {
	log.Log("Received Srv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Srv) Stream(ctx context.Context, req *srv.StreamingRequest, stream srv.Srv_StreamStream) error {
	log.Logf("Received Srv.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&srv.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Srv) PingPong(ctx context.Context, stream srv.Srv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&srv.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
