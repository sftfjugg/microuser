package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	microuser "github.com/sftfjugg/microuser/proto/microuser"
)

type Microuser struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Microuser) Call(ctx context.Context, req *microuser.Request, rsp *microuser.Response) error {
	log.Info("Received Microuser.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Microuser) Stream(ctx context.Context, req *microuser.StreamingRequest, stream microuser.Microuser_StreamStream) error {
	log.Infof("Received Microuser.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&microuser.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Microuser) PingPong(ctx context.Context, stream microuser.Microuser_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&microuser.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
