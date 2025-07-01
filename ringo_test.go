package main

import (
	"errors"
	"ringo/errs"
	"ringo/handlers"
	"ringo/models"
	"testing"
	"time"
)

// NOTE TO SELF
// TESTS ARE WRITTEN BY GPT I NEED TO REWRITE EVERYTHING BY MYSELF
// AYT NVM I FIXED THE MISSING THINGS NOW I NEED TO WRITE MORE TEST I THINK TO COVER MORE THINGS
func TestStoreHandler_Set(t *testing.T) {
	r := &models.RinGoObject{Values: make(map[string]models.GlobalObject)}
	h := handlers.SetHandler{}

	args := []string{"set", "mykey", "myvalue"}
	_, err := h.Handle(args, r)
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
	h := handlers.SSetHandler{}

	args := []string{"sset", "listkey", "val1", "val2", "val3"}
	_, err := h.Handle(args, r)
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

	args = []string{"sset", "listkey", "val4", "Val5", "val6"}
	_, err = h.Handle(args, r)
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
	h := handlers.HSetHandler{}
	args := []string{"hset", "listkey", "key1", "val1", "key2", "val2", "key3", "val3"}
	_, err := h.Handle(args, r)
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

	args = []string{"hset", "listkey", "key4", "val4", "key5", "val5"}
	_, err = h.Handle(args, r)
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
	sH := handlers.SetHandler{}
	args := []string{"set", "key", "val", "exp", "1"}
	_, err := sH.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	time.Sleep(2 * time.Second)

	_, err = r.Get("key")
	if !errors.Is(err, errs.ErrKeyDeleted) {
		t.Errorf("expected error %v, got %v", errs.ErrKeyDeleted, err)
	}

	sSH := handlers.SSetHandler{}

	args = []string{"sset", "key", "val", "val", "exp", "1"}
	_, err = sSH.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	time.Sleep(2 * time.Second)

	_, err = r.Get("key")
	if !errors.Is(err, errs.ErrKeyDeleted) {
		t.Errorf("expected error %v, got %v", errs.ErrKeyDeleted, err)
	}

	hSH := handlers.HSetHandler{}

	args = []string{"hset", "key", "key1", "val1", "exp", "1"}
	_, err = hSH.Handle(args, r)
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
	sH := handlers.SetHandler{}
	dH := handlers.DeleteHandler{}
	args := []string{"set", "key", "val"}
	_, err := sH.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	args = []string{"delete", "key"}
	_, err = dH.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	_, err = r.Get("key")
	if !errors.Is(err, errs.ErrNotFound) {
		t.Errorf("expected error %v, got %v", errs.ErrNotFound, err)
	}
}
