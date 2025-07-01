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
	Handle(args []string, r *models.RinGoObject) (interface{}, error)
}

type SetHandler struct{}

type SSetHandler struct{}

type HSetHandler struct{}

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

func (h SetHandler) Handle(args []string, r *models.RinGoObject) (interface{}, error) {
	if len(args) < 3 {
		return false, errs.ErrArgs
	}
	key := args[1]
	value := args[2]
	var exp time.Duration
	if len(args) > 4 {
		if strings.ToLower(args[3]) == "exp" {
			if dur, err := processTime(args[4]); err != nil {
				return false, err
			} else {
				exp = dur
			}
		}
	}
	if err := r.Store(key, value, exp); err != nil {
		return false, err
	}
	return true, nil
}

func (h SSetHandler) Handle(args []string, r *models.RinGoObject) (interface{}, error) {
	if len(args) < 3 {
		return false, errs.ErrArgs
	}
	key := args[1]
	values := []string{}
	var exp time.Duration
	for i := 2; i < len(args); i++ {
		if strings.ToLower(args[i]) == "exp" {
			if i+1 < len(args) {
				dur, err := processTime(args[i+1])
				if err != nil {
					return false, err
				}
				exp = dur
			}
			break
		}
		values = append(values, args[i])
	}
	if err := r.Store(key, values, exp); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (h HSetHandler) Handle(args []string, r *models.RinGoObject) (interface{}, error) {
	if len(args) < 4 {
		return false, errs.ErrArgs
	}
	key := args[1]
	values := make(map[string]string)
	var exp time.Duration
	for i := 2; i < len(args); i += 2 {
		if strings.ToLower(args[i]) == "exp" {
			if i+1 < len(args) {
				dur, err := processTime(args[i+1])
				if err != nil {
					return false, err
				}
				exp = dur
			}
			break
		}
		if i+1 >= len(args) {
			return false, fmt.Errorf("expected value after key '%s'", args[i])
		}
		hsetKey := args[i]
		hsetVal := args[i+1]
		values[hsetKey] = hsetVal
	}
	if err := r.Store(key, values, exp); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (h GetHandler) Handle(args []string, r *models.RinGoObject) (interface{}, error) {
	if len(args) < 2 {
		return nil, errs.ErrArgs
	}
	key := args[1]
	if key == "" {
		return nil, errs.ErrNilKey
	}
	if val, err := r.Get(key); err != nil {
		return nil, err
	} else {
		return val, nil
	}
}

func (h DeleteHandler) Handle(args []string, r *models.RinGoObject) (interface{}, error) {
	if len(args) < 2 {
		return false, errs.ErrArgs
	}
	key := args[1]
	if key == "" {
		return false, errs.ErrNilKey
	}
	if err := r.Delete(key); err != nil {
		return false, err
	}
	return true, nil
}
