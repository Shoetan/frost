package utils


import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

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