package main

import (
	"os"
	"log"

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

	content , err := os.ReadFile(filepath)

	if err != nil {
		log.Println("Cannot read file")
	}

	log.Println(string(content))

	return nil
	
}

