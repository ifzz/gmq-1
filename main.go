package main

import (
	"go-mq/mq"
	"go-mq/service"
)

func main() {
	// new dispatcher
	dispatcher := mq.NewDispatcher()
	go dispatcher.Run()

	// new service
	serv := service.NewService(dispatcher)
	serv.Target.Run()
}