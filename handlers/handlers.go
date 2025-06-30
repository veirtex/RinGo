package handlers

import (
	"fmt"
	"ringo/errs"
	"ringo/models"
	"strconv"
	"strings"
	"time"
)

type Handler interface {
	Handle(args []string, r *models.RinGoObject) error
}

type SetHandler struct{}

type SSetHandler struct{}

type HSetHandler struct{}

type DeleteHandler struct{}

func processTime(sndStr string) (time.Duration, error) {
	secondsInt, err := strconv.Atoi(sndStr)
	if err != nil {
		return 0, err
	} else {
		return time.Duration(secondsInt) * time.Second, nil
	}
}

func (h SetHandler) Handle(args []string, r *models.RinGoObject) error {
	if len(args) < 4 {
		return errs.ErrArgs
	}
	key := args[2]
	value := args[3]
	var exp time.Duration
	if len(args) > 5 {
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
}

func (h SSetHandler) Handle(args []string, r *models.RinGoObject) error {
	key := args[2]
	values := []string{}
	var exp time.Duration
	for i := 3; i < len(args); i++ {
		if strings.ToLower(args[i]) == "exp" {
			if i+1 < len(args) {
				dur, err := processTime(args[i+1])
				if err != nil {
					return err
				}
				exp = dur
			}
			break
		}
		values = append(values, args[i])
	}
	if err := r.Store(key, values, exp); err != nil {
		return err
	} else {
		return nil
	}
}

func (h HSetHandler) Handle(args []string, r *models.RinGoObject) error {
	key := args[2]
	values := make(map[string]string)
	var exp time.Duration
	for i := 3; i < len(args); i += 2 {
		if strings.ToLower(args[i]) == "exp" {
			if i+1 < len(args) {
				dur, err := processTime(args[i+1])
				if err != nil {
					return err
				}
				exp = dur
			}
			break
		}
		if i+1 >= len(args) {
			return fmt.Errorf("expected value after key '%s'", args[i])
		}
		hsetKey := args[i]
		hsetVal := args[i+1]
		values[hsetKey] = hsetVal
	}
	if err := r.Store(key, values, exp); err != nil {
		return err
	} else {
		return nil
	}
}

func GetHandle(args []string, r *models.RinGoObject) (interface{}, error) {
	if len(args) < 3 {
		return nil, errs.ErrArgs
	}
	key := args[2]
	if key == "" {
		return nil, errs.ErrNilKey
	}
	if val, err := r.Get(key); err != nil {
		return nil, err
	} else {
		return val, nil
	}
}

func (h DeleteHandler) Handle(args []string, r *models.RinGoObject) error {
	if len(args) < 3 {
		return errs.ErrArgs
	}
	key := args[2]
	if key == "" {
		return errs.ErrNilKey
	}
	if err := r.Delete(key); err != nil {
		return err
	}
	return nil
}
