package redispubhandler

import (
	"gopkg.in/redis.v5"
)

type Context struct {
	Error   error
	Message string
}

type Request interface {
	Response(*Context)
}

func Handle(r *redis.Client, sub string, req Request) error {
	subscriber, err := r.Subscribe(sub)
	if err != nil {
		return err
	}
	go func() {
		for {

			if msg, err := subscriber.ReceiveMessage(); err != nil {
				req.Response(&Context{
					Error: err,
				})
			} else {
				req.Response(&Context{Message: msg.Payload})
			}

		}
	}()
	return nil
}
