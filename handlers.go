package main

import (
	"fmt"
	"ringo/models"
	"strconv"
	"strings"
	"time"
)

type Handler interface {
	Handle(args []string, r *models.RinGoObject) error
}

type StoreHandler struct{}

type GetHandler struct{}

type DeleteHandler struct{}

func processTime(sndStr string) (time.Duration, error) {
	secondsInt, err := strconv.Atoi(sndStr)
	if err != nil {
		return 0, err
	} else {
		return time.Duration(secondsInt) * time.Second, nil
	}
}

func (h StoreHandler) Handle(args []string, r *models.RinGoObject) error {
	if len(args) < 4 {
		return fmt.Errorf(`args are not enough try "command" -h `)
	}
	command := args[1]
	switch strings.ToLower(command) {
	case "set":
		key := args[2]
		value := args[3]
		var exp time.Duration
		if len(args) > 6 {
			if strings.ToLower(args[4]) == "exp" {
				if dur, err := processTime(args[5]); err != nil {
					return err
				} else {
					exp = dur
				}
			}
		}
		if err := r.Store(key, value, exp); err != nil {
			return err
		}
		return nil
	case "sset":
		key := args[2]
		values := []string{}
		var exp time.Duration
		for ind, arg := range args[3:] {
			if strings.ToLower(arg) == "exp" {
				if len(args[3:]) >= ind+1 {
					if dur, err := processTime(args[3:][ind+1]); err != nil {
						return err
					} else {
						exp = dur
					}
				}
				break
			}
			values = append(values, arg)
		}
		if err := r.Store(key, values, exp); err != nil {
			return err
		} else {
			return nil
		}
	case "hset":
		// TODO
		return nil
	default:
		// TODO
		return nil
	}
}

func (h GetHandler) Handle(args []string, r *models.RinGoObject) error {
	// TODO
	return nil
}

func (h DeleteHandler) Handle(args []string, r *models.RinGoObject) error {
	// TODO
	return nil
}
