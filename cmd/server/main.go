package main

import (
	"github.com/frost/pkg/server"
)


func main()  {
	
	s := server.NewAPISERVER(":9000")
	
	s.Run()
	
}
