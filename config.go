package gapp

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// The generic configuration entry
type _ConfigEntry struct {
	Long        string
	Short       string
	Description string
	Value       interface{}
	Type        reflect.Kind
	CLI         bool
	Listeners   []chan interface{}
}

// Config wrapper struct with RWMutex
type _Config struct {
	sync.RWMutex
	alias   map[string]string
	entries map[string]_ConfigEntry
}

var Config *_Config

// Empty configuration entry
func DefaultEntry() _ConfigEntry {
	return _ConfigEntry{
		Long:        "",
		Short:       "",
		Description: "",
		Value:       nil,
		Type:        reflect.String,
		CLI:         false,
		Listeners:   make([]chan interface{}, 0),
	}
}

// Initializes configuration to an empty state
func _Load() {
	Config = &_Config{
		alias:   make(map[string]string),
		entries: make(map[string]_ConfigEntry),
	}
}

func init() {
	_Load()
}

// Returns the splice of entry keys
func (c *_Config) Keys() []string {
	klist := []string{}
	for k, _ := range c.entries {
		klist = append(klist, k)
	}
	return klist
}

// Remove all configuration parameters
func (c *_Config) Clear() {
	_Load()
}

// Add up a new configuration parameter
func (c *_Config) Add(long, short, description string, value interface{}, tpe reflect.Kind, cli bool) (_ConfigEntry, error) {
	entry := DefaultEntry()
	entry.Long = long
	entry.Short = short
	entry.Description = description
	entry.Value = value
	entry.Type = tpe
	entry.CLI = cli
	c.Lock()
	defer c.Unlock()
	err := enforce_type(value, entry)
	if err == nil {
		c.alias[short] = long
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
func (c *_Config) Get_entry(key string) (_ConfigEntry, bool) {
	c.RLock()
	defer c.RUnlock()
	entry := key
	if alias, ok := c.alias[key]; ok {
		entry = alias
	}
	cfg, ok2 := c.entries[entry]
	return cfg, ok2
}

// Retrieve a configuration value
func (c *_Config) Get(key string) (interface{}, bool) {
	cfg, ok := c.Get_entry(key)
	if !ok {
		return nil, ok
	}
	return cfg.Value, ok
}

//Sets a configuration parameter
func (c *_Config) Set(key string, value interface{}) (_ConfigEntry, error) {
	entry, ok := c.Get_entry(key)
	if ok {
		if reflect.ValueOf(value) == reflect.ValueOf(entry.Value) {
			return entry, nil
		}
	}

	c.Lock()
	defer c.Unlock()
	err := enforce_type(value, entry)
	if err == nil {
		entry.Value = value
		c.alias[entry.Short] = key
		c.entries[key] = entry
		go c._notify_subscribers(key, value)
	}
	return entry, err
}

// Helper function for listener notifications
func (c *_Config) _notify_subscribers(key string, value interface{}) {
	entry, _ := c.Get_entry(key)
	for _, ch := range entry.Listeners {
		ch <- value
	}
}

// Subscribe to a given configuration value
func (c *_Config) Subscribe_to(key string) (chan interface{}, error) {
	ch := make(chan interface{})
	entry, ok := c.Get_entry(key)
	if ok {
		c.Lock()
		defer c.Unlock()
		entry.Listeners = append(entry.Listeners, ch)
		c.entries[key] = entry
		return ch, nil
	}
	return ch, errors.New(fmt.Sprintf("Unable to subscribe to config: `%s`", key))
}
