package main

import (
	"fmt"
	"log"
	"os"
	"ringo/errs"
	"ringo/handlers"
	"ringo/models"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Printf("Error: %s", errs.ErrArgs)
		return
	}
	handlrs := map[string]handlers.Handler{
		"set":    handlers.SetHandler{},
		"sset":   handlers.SSetHandler{},
		"hset":   handlers.HSetHandler{},
		"delete": handlers.DeleteHandler{},
	}
	var command string
	if len(args) > 1 {
		command = args[1]
	} else {
		log.Printf("Error: %s", errs.ErrUnknownCommand)
		return
	}

	var r models.RinGoObject
	r.Values = make(map[string]models.GlobalObject)
	if handler, ok := handlrs[command]; ok {
		err := handler.Handle(args, &r)
		if err != nil {
			log.Printf("Error: %s", err)
		} else {
			fmt.Println("Done")
		}
	} else {
		if command == "get" {
			if val, err := handlers.GetHandle(args, &r); err != nil {
				log.Printf("Error: %s", err)
			} else {
				fmt.Println(val)
			}
		} else {
			log.Printf("Error: %s", errs.ErrUnknownCommand)
		}
	}
}
