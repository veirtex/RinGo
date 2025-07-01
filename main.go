package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"ringo/errs"
	"ringo/handlers"
	"ringo/models"
	"strings"
)

func main() {
	isRunning := true
	handlrs := map[string]handlers.Handler{
		"set":    handlers.SetHandler{},
		"sset":   handlers.SSetHandler{},
		"hset":   handlers.HSetHandler{},
		"get":    handlers.GetHandler{},
		"delete": handlers.DeleteHandler{},
	}
	var command string
	var r models.RinGoObject
	r.Values = make(map[string]models.GlobalObject)
	for isRunning {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command = scanner.Text()
		args := handleCommand(command)
		if args == nil {
			log.Printf("%s", errs.ErrArgs)
			continue
		}
		if args[0] == "exit" {
			isRunning = false
			break
		}
		if handler, ok := handlrs[args[0]]; ok {
			res, err := handler.Handle(args, &r)
			if err != nil {
				log.Printf("%s", err)
			} else {
				printing(res)
			}
		} else {
			log.Printf("%s", errs.ErrUnknownCommand)
		}
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func handleCommand(command string) []string {
	fields := strings.Fields(command)
	return fields
}

func printing(val interface{}) {
	switch v := val.(type) {
	case bool:
		fmt.Printf("%d\n", boolToInt(v))
	case string:
		fmt.Println(v)
	case []string:
		fmt.Println(strings.Join(v, ", "))
	case map[string]string:
		for k, value := range v {
			fmt.Printf("%s: %s ", k, value)
		}
		fmt.Println()
	default:
		log.Println(errs.ErrDatatype)
	}
}
