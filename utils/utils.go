package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Task struct{
	GenerationId string `json:"generation_id"`
	Text []byte `json:"file_content"`
}

func GetEnvVariables(keys ...string) map[string]string {
	err := godotenv.Load()
	if err != nil {
			log.Printf("Error loading .env: %s", err.Error())
	}

	envVars := make(map[string]string)
	for _, key := range keys {
			envVars[key] = os.Getenv(key)
	}

	return envVars
}


func LogError( err error, message string)  {
		if err != nil  {
			log.Printf( "%s : %v", message, err.Error())
		}
}


func PublishTask(conn *amqp.Connection, randomId string, text []byte) {
	//create a channel 
	channel, err := conn.Channel()
	LogError(err, "Could not create a channel")

	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"task", //name of queue
		false,
		false,
		false,
		false,
		nil,
	)

	LogError(err, "Could not declare a queue")

	task := Task{
		GenerationId: randomId,
		Text: text,
	}

	//convert the task into json
	body, err := json.Marshal(task)

	LogError(err, "unable to unmarshall")

	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})

	LogError(err, "Could not publish message")

}

func Receivetask(conn *amqp.Connection){
		//create a channel 
		channel, err := conn.Channel()
		LogError(err, "Could not create a channel")
	
		defer channel.Close()
	
		queue, err := channel.QueueDeclare(
			"task", //name of queue
			false,
			false,
			false,
			false,
			nil,
		)
	
		LogError(err, "Could not declare a queue")

		// Receive messages from the queue
		msgs, err := channel.Consume(
		queue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	LogError(err, "Failed to register a consumer")

	//process 

	go func() {
		for d := range msgs{
			var task Task

			err := json.Unmarshal(d.Body, &task)

			LogError(err, "Could not unmarshall task")

			log.Printf("Received a task: GenerationID=%s, Text=%s", task.GenerationId, task.Text)

		}
	}()


}