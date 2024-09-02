package server

import (
	"log"
	"net/http"

	rabbitmq "github.com/frost/pkg/rabbitMQ"
	"github.com/frost/pkg/server/handlers"
	//amqp "github.com/rabbitmq/amqp091-go"

	"github.com/frost/pkg/database"
	"github.com/frost/utils"
)


type APISERVER struct {
	addr string
}


func NewAPISERVER(addr string) *APISERVER  {

	return &APISERVER{
		addr: addr,
	}
	
}


func (s *APISERVER) Run() error {


	// connect to database
	db, err := database.Database()

	utils.LogError(err, "Cannot connect to the database")

	// connect to rabbitMQ
	conn, err := rabbitmq.RabbitMQ()

	utils.LogError(err, "Cannot connect to rabbitMQ")

	db.MustExec(database.GenerationTable)

	router := http.NewServeMux()

	router.HandleFunc("POST /generate", handlers.GenerateTask(db, conn))
	
	go utils.Receivetask(conn)

	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}


	log.Printf("Server is running on port %s", s.addr)

	return server.ListenAndServe()

} 

