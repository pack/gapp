package gapp

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type _ConfigEntry struct {
	Long        string
	Short       string
	Description string
	Value       interface{}
	Type        reflect.Kind
	Required    bool
	CLI         bool
	Listeners   []chan interface{}
}

type _Config struct {
	sync.RWMutex
	entries map[string]_ConfigEntry
}

var Config *_Config

func DefaultEntry() _ConfigEntry {
	return _ConfigEntry{
		Long:        "",
		Short:       "",
		Description: "",
		Value:       nil,
		Type:        reflect.String,
		Required:    false,
		CLI:         false,
		Listeners:   make([]chan interface{}, 0),
	}
}

// Initializes configuration to an empty state
func _Load() {
	Config = &_Config{entries: make(map[string]_ConfigEntry)}
}

func init() {
	_Load()
}

// Returns the splice of entry keys
func (c *_Config) keys() []string {
	klist := []string{}
	for k, _ := range c.entries {
		klist = append(klist, k)
	}
	return klist
}

// Remove all configuration parameters
func (c *_Config) clear() {
	_Load()
}

// Add up a new configuration parameter
func (c *_Config) add(long, short, description string, value interface{}, tpe reflect.Kind, required, cli bool) (_ConfigEntry, error) {
	entry := DefaultEntry()
	entry.Long = long
	entry.Short = short
	entry.Description = description
	entry.Value = value
	entry.Type = tpe
	entry.Required = required
	entry.CLI = cli
	c.Lock()
	defer c.Unlock()
	err := enforce_type(value, entry)
	if err == nil {
		c.entries[long] = entry
	}
	return entry, err
}

// Type enforce a configuration variable
func enforce_type(value interface{}, cfg _ConfigEntry) error {
	val := reflect.ValueOf(value)
	tpe := val.Kind()
	if tpe != cfg.Type {
		errmsg := fmt.Sprintf("Config Entry `%s` is not of type `%s`: (%s [%s])", cfg.Long, cfg.Type, val, tpe)
		return errors.New(errmsg)
	}
	return nil
}

// Retrieve a configuation entry
func (c *_Config) get_entry(key string) (_ConfigEntry, bool) {
	c.RLock()
	defer c.RUnlock()
	cfg, ok := c.entries[key]
	return cfg, ok
}

// Retrieve a configuration value
func (c *_Config) get(key string) (interface{}, bool) {
	cfg, ok := c.get_entry(key)
	if ok == false {
		return nil, ok
	}
	return cfg.Value, ok
}

//Sets a configuration parameter
func (c *_Config) set(key string, value interface{}) (_ConfigEntry, error) {
	entry, ok := c.get_entry(key)
	if ok == true {
		if reflect.ValueOf(value) == reflect.ValueOf(entry.Value) {
			return entry, nil
		}
	}

	c.Lock()
	defer c.Unlock()
	err := enforce_type(value, entry)
	if err == nil {
		entry.Value = value
		c.entries[key] = entry
	}
	return entry, err
}
