package models

import "time"

type GlobalObject struct {
	ExpirationTime time.Time
	Value          interface{}
}

type RinGoObject struct {
	Values map[string]GlobalObject
}
