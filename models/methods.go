package models

import (
	"fmt"
	"time"
)

func (r *RinGoObject) Store(key string, value interface{}, duration time.Duration) error {
	if key == "" || value == nil {
		return fmt.Errorf("the key or value is nil")
	}
	var exp time.Time
	if duration != 0 {
		exp = time.Now().Add(duration)
	}
	switch v := value.(type) {
	case string:
		r.Values[key] = GlobalObject{Value: v, ExpirationTime: exp}
	case []string:
		r.Values[key] = GlobalObject{Value: v, ExpirationTime: exp}
	case map[string]string:
		r.Values[key] = GlobalObject{Value: v, ExpirationTime: exp}
	default:
		return fmt.Errorf("failed to define the datatype of the value")
	}
	return nil
}

func (r *RinGoObject) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, fmt.Errorf("key is nil")
	}
	val, ok := r.Values[key]
	if !ok {
		return nil, fmt.Errorf("value wasn't found")
	}

	if !val.ExpirationTime.IsZero() && time.Now().After(val.ExpirationTime) {
		delete(r.Values, key)
		return nil, fmt.Errorf("value was deleted")
	}

	switch v := val.Value.(type) {
	case string:
		return v, nil
	case []string:
		return v, nil
	case map[string]string:
		return v, nil
	default:
		return nil, fmt.Errorf("failed to define the datatype of the value")
	}
}

func (r *RinGoObject) Delete(key string) error {
	if key == "" {
		return fmt.Errorf("key is nil")
	}
	delete(r.Values, key)
	return nil
}

func NewRinGoObject() *RinGoObject {
	return &RinGoObject{
		Values: make(map[string]GlobalObject),
	}
}
