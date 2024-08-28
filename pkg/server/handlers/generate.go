package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/frost/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	amqp "github.com/rabbitmq/amqp091-go"
)


type GenerateResponse struct {
	GenerationID string `json:"generation_Id"`
	Status       string `json:"status"`
}

func GenerateTask(db *sqlx.DB, conn *amqp.Connection) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//Get contents from body 
		content, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Could not read contents from body", http.StatusBadRequest)
			return
		}

		//generate a new random Id for the generation task to just 8 digits
		generationID := uuid.New().String()[:8]


		//Insert into the database 

		_, err = db.Exec("INSERT INTO generation_tasks(generation_Id, status, file_content) VALUES($1, $2, $3)", generationID, "Processing", content)

		utils.LogError(err, "Could not insert into Database")


		response := GenerateResponse{
			GenerationID: generationID,
			Status: "Queued & processing",
		}

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(response)


		//Queue task into  redis or rabbitMQ

		utils.PublishTask(conn,generationID, content)
	
	
		
	}
	
}