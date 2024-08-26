package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/frost/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)


type GenerateResponse struct {
	GenerationID string `json:"generation_Id"`
	Status       string `json:"status"`
}

func GenerateTask(db *sqlx.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//Get contents from body 
		_, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Could not read contents from body", http.StatusBadRequest)
			return
		}

		//generate a new random Id for the generation task 
		generationID := uuid.New().String()


		//Insert into the database 

		_, err = db.Exec("INSERT INTO generation_tasks(generation_Id, status) VALUES($1, $2)", generationID, "Queued")

		utils.LogError(err, "Could not insert into Database")


		response := GenerateResponse{
			GenerationID: generationID,
			Status: "queued",
		}

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(response)
	
	
		
	}
	
}