package models

import (
	"errors"
	"fmt"
	"ringo/errs"
	"time"
)

var ()

func (r *RinGoObject) Store(key string, value interface{}, duration time.Duration) error {
	if key == "" {
		return fmt.Errorf("key: %s, %w", key, errs.ErrNilKey)
	}
	var exp time.Time
	if duration != 0 {
		exp = time.Now().Add(duration)
	}
	switch v := value.(type) {
	case string:
		r.Values[key] = GlobalObject{Value: v, ExpirationTime: exp}
	case []string:
		if err := r.handleValues(v, key, exp); err != nil {
			return err
		}
	case map[string]string:
		if err := r.handleValues(v, key, exp); err != nil {
			return err
		}
	default:
		return errs.ErrDatatype
	}
	return nil
}

func (r *RinGoObject) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, fmt.Errorf("key: %s, %w", key, errs.ErrNilKey)
	}
	val, ok := r.Values[key]
	if !ok {
		return nil, fmt.Errorf("key: %s, %w", key, errs.ErrNotFound)
	}

	if !val.ExpirationTime.IsZero() && time.Now().After(val.ExpirationTime) {
		delete(r.Values, key)
		return nil, fmt.Errorf("key: %s, %w", key, errs.ErrKeyDeleted)
	}

	switch v := val.Value.(type) {
	case string:
		return v, nil
	case []string:
		return v, nil
	case map[string]string:
		return v, nil
	default:
		return nil, errors.New("failed to define the datatype of the value")
	}
}

func (r *RinGoObject) Delete(key string) error {
	if key == "" {
		return fmt.Errorf("key: %s, %w", key, errs.ErrNilKey)
	}
	delete(r.Values, key)
	return nil
}

func (r *RinGoObject) getExpirationDate(key string) (time.Time, error) {
	if key == "" {
		return time.Time{}, fmt.Errorf("key: %s, %w", key, errs.ErrNilKey)
	}
	val, ok := r.Values[key]
	if !ok {
		return time.Time{}, fmt.Errorf("key: %s, %w", key, errs.ErrNotFound)
	}
	return val.ExpirationTime, nil
}

func (r *RinGoObject) handleValues(newVal interface{}, key string, exp time.Time) error {
	val, err := r.Get(key)
	if err != nil {
		r.Values[key] = GlobalObject{Value: newVal, ExpirationTime: exp}
		return nil
	}
	oldExp, errExp := r.getExpirationDate(key)
	if errExp != nil {
		oldExp = exp
	}
	merged, err := mergeFunc(val, newVal)
	if err != nil {
		return err
	}
	r.Values[key] = GlobalObject{Value: merged, ExpirationTime: oldExp}
	return nil
}

func mergeFunc(val interface{}, newVal interface{}) (interface{}, error) {
	switch v := val.(type) {
	case []string:
		v2, ok := newVal.([]string)
		if !ok {
			return nil, errs.ErrDatatype
		}
		v = append(v, v2...)
		return v, nil
	case map[string]string:
		v2, ok := newVal.(map[string]string)
		if !ok {
			return nil, errs.ErrDatatype
		}
		for k, value := range v2 {
			v[k] = value
		}
		return v, nil
	default:
		return nil, errs.ErrDatatype
	}
}

func NewRinGoObject() *RinGoObject {
	return &RinGoObject{
		Values: make(map[string]GlobalObject),
	}
}
