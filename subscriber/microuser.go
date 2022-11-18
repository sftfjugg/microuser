package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	microuser "github.com/sftfjugg/microuser/proto/microuser"
)

type Microuser struct{}

func (e *Microuser) Handle(ctx context.Context, msg *microuser.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *microuser.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
