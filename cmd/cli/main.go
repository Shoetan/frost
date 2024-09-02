package main

import (
	"log"
	"net/http"
	"os"
	"bytes"

	"github.com/spf13/cobra"
)
	


func main()  {

	//parent command
	rootCMD := &cobra.Command{Use: "frost" }

	generateImage := &cobra.Command{
		Use: "generate [file]",
		Short: "Simple print function",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return generateImage(args[0])
		},
	}

	rootCMD.AddCommand(generateImage)

	if err := rootCMD.Execute(); err != nil {
		log.Fatal(err)
	}
}

// this function will take the path to the file path read it and send the contents as body to and endpoint 
func generateImage(filepath string)  error {

	content , err := os.ReadFile(filepath) // Read contents from the file specified in the command

	if err != nil {
		log.Println("Cannot read file")
	}

	// create a new http request with apiURL 
	req, err := http.NewRequest("POST","http://localhost:9000/generate" , bytes.NewReader(content))

	if err != nil {
		log.Printf("Request cannot be created %s", err.Error())
	}

	//Create a client to send request over http
	client := &http.Client{}

	resp, err := client.Do(req) //send the request over http

	if err != nil {
		log.Printf("Cannot send request over http %s", err.Error())
	}

	log.Print(resp.StatusCode)
	
	//defer resp.Body.Close()

	return nil
	
}

/* 
		Next steps 
		1. Collect text from response body in a readable human form: Done
		2. Queue the generation process because It is not immediate: Done
		3. Once queued return an Id or something reference Alexis or claude: Done
		4. Design the creation process of puting text on canvas with text highlighting for different languages
 */

