package main

import (
	"fmt"
	"os"
	"ringo/models"
)

func main() {
	args := os.Args
	handlers := map[string]Handler{
		"store":  StoreHandler{},
		"get":    GetHandler{},
		"delete": DeleteHandler{},
	}
	var command string
	if len(args) > 1 {
		command = args[1]
	}

	var r models.RinGoObject
	r.Values = make(map[string]models.GlobalObject)
	if handler, ok := handlers[command]; ok {
		err := handler.Handle(args, &r)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Unknown command:", command)
	}
}
