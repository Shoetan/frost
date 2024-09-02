package utils

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
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



func createImage(code string, outputPath string) error {

	width , height := 600, 600

	dc := gg.NewContext(width, height)

	//Set background color 
	dc.SetColor(color.RGBA{40,44,52,255})

	dc.Clear()

	//load a font

	font, err := truetype.Parse(goregular.TTF)
	LogError(err, "Could not load fonts")

	//set font face
	face := truetype.NewFace(font, &truetype.Options{Size: 16})
	dc.SetFontFace(face)

	//set text color
	dc.SetColor(color.White)

		// Draw the code
		margin := 40.0
		lineHeight := 1.5
		y := margin
		for _, line := range strings.Split(code, "\n") {
			dc.DrawString(line, margin, y)
			y += dc.FontHeight() * lineHeight
		}
	
		// Save the image
		return dc.SavePNG(outputPath)

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
		false,   // auto-ack
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
			LogError(err, "Could not unmarshal task")

			log.Printf("Received a task: GenerationID: %s, Text: %s", task.GenerationId, task.Text)


			//add the generation image code here 
			err = createImage(string(task.Text), "worki.png")

			LogError(err, "Image creation failed")

			d.Ack(false)

		}
	}()


}