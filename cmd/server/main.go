package main

import (
	rabbitmq "github.com/frost/pkg/rabbitMQ"
	"github.com/frost/pkg/server"
	"github.com/frost/utils"
)


func main()  {

	conn, err := rabbitmq.RabbitMQ()

	utils.LogError(err, "Could not connect to rabbitMQ")

	
	s := server.NewAPISERVER(":9000")
	
	s.Run(conn)
	
	utils.Receivetask(conn)
}
