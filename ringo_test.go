package main

import (
	"ringo/models"
	"testing"
)

// NOTE TO SELF
// TESTS ARE WRITTEN BY GPT I NEED TO REWRITE EVERYTHING BY MYSELF
func TestStoreHandler_Set(t *testing.T) {
	r := &models.RinGoObject{Values: make(map[string]models.GlobalObject)}
	h := StoreHandler{}

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

	args = []string{"cmd", "set", "key2", "value2", "exp", "5"}
	err = h.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error with expiration, got %v", err)
	}

	val, err = r.Get("key2")
	if err != nil {
		t.Fatalf("expected key2 to exist, got error %v", err)
	}
	if val != "value2" {
		t.Fatalf("expected value 'value2', got %v", val)
	}
}

func TestStoreHandler_SSet(t *testing.T) {
	r := &models.RinGoObject{Values: make(map[string]models.GlobalObject)}
	h := StoreHandler{}

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

	args = []string{"cmd", "sset", "listkey2", "valA", "valB", "exp", "2"}
	err = h.Handle(args, r)
	if err != nil {
		t.Fatalf("expected no error with expiration, got %v", err)
	}

	val, err = r.Get("listkey2")
	if err != nil {
		t.Fatalf("expected listkey2 to exist, got error %v", err)
	}
	slice, ok = val.([]string)
	if !ok {
		t.Fatalf("expected []string, got %T", val)
	}
	expected = []string{"valA", "valB"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("expected %v at index %d, got %v", v, i, slice[i])
		}
	}
}

func TestStoreHandler_HSet(t *testing.T) {
	// TODO

}

func TestExpirationTime(t *testing.T) {
	// TODO
}

func TestGet(t *testing.T) {
	// TODO
}

func TestDeleting(t *testing.T) {
	// TODO
}
