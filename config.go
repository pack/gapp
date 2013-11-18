package gapp

import (
	"sync"
	"reflect"
	"fmt"
)

var Config *_Config

type ConfigEntry struct {
	Value interface{}
	Type reflect.Type
}

type _Config struct {
	sync.RWMutex
	entries map[string]ConfigEntry
	// listeners []chan *Msg
}

// Returns the splice of entry keys
func (c *_Config) keys() []string {
	klist := []string{}
	for k, _ := range c.entries {
		klist = append(klist, k)
	}
	return klist
}

// Sets a configuration parameter
func (c *_Config) set(v, t reflect.Type) {
	ce := &ConfigEntry{Value: v, Type: t}
	fmt.Println(c.entries)
}

func _Load() {
	Config = &_Config{entries: make(map[string]ConfigEntry)}
}

func init() {
	_Load()
}
