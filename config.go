package gapp

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var Config *_Config

type ConfigEntry struct {
	Value interface{}
	Type  reflect.Kind
}

type _Config struct {
	sync.RWMutex
	entries map[string]ConfigEntry
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
func (c *_Config) set(k string, v interface{}, t reflect.Kind) (ConfigEntry, error) {
	entry := ConfigEntry{Value: v, Type: t}
	if t == reflect.ValueOf(v).Kind() {
		c.Lock()
		defer c.Unlock()
		c.entries[k] = entry
		return entry, nil
	}
	error_string := fmt.Sprintf("gapp: %s is not of kind %s", v, t)
	return ConfigEntry{}, errors.New(error_string)
}

// Gets a configuration parameter
func (c *_Config) get(k string) (ConfigEntry) {
	c.RLock()
	defer c.RUnlock()
	return c.entries[k]
}

func _Load() {
	Config = &_Config{entries: make(map[string]ConfigEntry)}
}

func init() {
	_Load()
}
