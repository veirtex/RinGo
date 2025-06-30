package main

import (
	"fmt"
	"log"
	"os"
	"ringo/errs"
	"ringo/models"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		log.Panic(errs.ErrArgs)
	}
	handlers := map[string]Handler{
		"set":    SetHandler{},
		"sset":   SSetHandler{},
		"hset":   HSetHandler{},
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
