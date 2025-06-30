package main

import (
	"errors"
	"ringo/errs"
	"ringo/models"
	"testing"
	"time"
)

// NOTE TO SELF
// TESTS ARE WRITTEN BY GPT I NEED TO REWRITE EVERYTHING BY MYSELF
func TestStoreHandler_Set(t *testing.T) {
	r := &models.RinGoObject{Values: make(map[string]models.GlobalObject)}
	h := SetHandler{}

	args := []string{"cmd", "set", "mykey", "myvalue"}
	err := h.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	val, err := r.Get("mykey")
	if err != nil {
		t.Fatalf("expected key to exist, got error %v", err)
	}
	if val != "myvalue" {
		t.Fatalf("expected value 'myvalue', got %v", val)
	}
}

func TestStoreHandler_SSet(t *testing.T) {
	r := &models.RinGoObject{Values: make(map[string]models.GlobalObject)}
	h := SSetHandler{}

	args := []string{"cmd", "sset", "listkey", "val1", "val2", "val3"}
	err := h.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	val, err := r.Get("listkey")
	if err != nil {
		t.Fatalf("expected listkey to exist, got error %v", err)
	}

	slice, ok := val.([]string)
	if !ok {
		t.Fatalf("expected []string, got %T", val)
	}
	expected := []string{"val1", "val2", "val3"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("expected %v at index %d, got %v", v, i, slice[i])
		}
	}

	args = []string{"cmd", "sset", "listkey", "val4", "Val5", "val6"}
	err = h.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error with expiration, got %v", err)
	}

	val, err = r.Get("listkey")
	if err != nil {
		t.Fatalf("expected listkey to exist, got error %v", err)
	}
	slice, ok = val.([]string)
	if !ok {
		t.Fatalf("expected []string, got %T", val)
	}
	expected = []string{"val1", "val2", "val3", "val4", "Val5", "val6"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("expected %v at index %d, got %v", v, i, slice[i])
		}
	}
}

func TestStoreHandler_HSet(t *testing.T) {
	r := &models.RinGoObject{Values: make(map[string]models.GlobalObject)}
	h := HSetHandler{}
	args := []string{"cmd", "hset", "listkey", "key1", "val1", "key2", "val2", "key3", "val3"}
	err := h.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	val, err := r.Get("listkey")
	if err != nil {
		t.Fatalf("expected listkey to exist, got error %v", err)
	}

	dict, ok := val.(map[string]string)
	if !ok {
		t.Fatalf("expected map[string]string, got %T", val)
	}

	expected := map[string]string{"key1": "val1", "key2": "val2", "key3": "val3"}
	for k, v := range expected {
		if dict[k] != v {
			t.Errorf("expected %v for key %s, got %v", v, k, dict[k])
		}
	}

	args = []string{"cmd", "hset", "listkey", "key4", "val4", "key5", "val5"}
	err = h.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	val, err = r.Get("listkey")
	if err != nil {
		t.Fatalf("expected listkey to exist, got error %v", err)
	}

	dict, ok = val.(map[string]string)
	if !ok {
		t.Fatalf("expected map[string]string, got %T", val)
	}

	expected = map[string]string{"key1": "val1", "key2": "val2", "key3": "val3", "key4": "val4", "key5": "val5"}
	for k, v := range expected {
		if dict[k] != v {
			t.Errorf("expected %v for key %s, got %v", v, k, dict[k])
		}
	}
}

func TestExpirationTime(t *testing.T) {
	r := &models.RinGoObject{Values: make(map[string]models.GlobalObject)}
	sH := SetHandler{}
	args := []string{"cmd", "set", "key", "val", "exp", "1"}
	err := sH.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	time.Sleep(2 * time.Second)

	_, err = r.Get("key")
	if !errors.Is(err, errs.ErrKeyDeleted) {
		t.Errorf("expected error %v, got %v", errs.ErrKeyDeleted, err)
	}

	sSH := SSetHandler{}

	args = []string{"cmd", "sset", "key", "val", "val", "exp", "1"}
	err = sSH.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	time.Sleep(2 * time.Second)

	_, err = r.Get("key")
	if !errors.Is(err, errs.ErrKeyDeleted) {
		t.Errorf("expected error %v, got %v", errs.ErrKeyDeleted, err)
	}

	hSH := HSetHandler{}

	args = []string{"cmd", "hset", "key", "key1", "val1", "exp", "1"}
	err = hSH.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	time.Sleep(2 * time.Second)

	_, err = r.Get("key")
	if !errors.Is(err, errs.ErrKeyDeleted) {
		t.Errorf("expected error %v, got %v", errs.ErrKeyDeleted, err)
	}
}

func TestDeleting(t *testing.T) {
	r := &models.RinGoObject{Values: make(map[string]models.GlobalObject)}
	sH := SetHandler{}
	dH := DeleteHandler{}
	args := []string{"cmd", "set", "key", "val"}
	err := sH.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	args = []string{"cmd", "delete", "key"}
	err = dH.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	_, err = r.Get("key")
	if !errors.Is(err, errs.ErrNotFound) {
		t.Errorf("expected error %v, got %v", errs.ErrNotFound, err)
	}
}
