package gapp

import (
	"testing"
	"reflect"
)

func TestEmptyKeys(t *testing.T) {
	if len(Config.keys()) != 0 {
		t.Error("Config entries are not empty")
	}
}

func TestAddingEntry(t *testing.T) {
	resp, err := Config.set("power", 9000, reflect.Int)
	if err != nil {
		t.Error(err)
		return
	}
	if resp.Value != 9000 {
		t.Error("Added configuration entry value did not match")
	}
}

func TestAddingInvalidEntry(t *testing.T) {
	_, err := Config.set("power", 9000, reflect.String)
	if err == nil {
		t.Error("Failed to properly handle adding an invalid entry")
	}
}
