package gapp

import (
	//"errors"
	//"fmt"
	"reflect"
	"sync"
)

type _ConfigEntry struct {
	Long string
	Short string
	Description string
	Value interface{}
	Type reflect.Kind
	CLI bool
	Required bool
	Listeners []chan interface{}
}

type _Config struct {
	sync.RWMutex
	entries map[string]_ConfigEntry
}

var Config *_Config

func DefaultEntry() _ConfigEntry {
	return _ConfigEntry{
		Long: "",
		Short: "",
		Description: "",
		Value: nil,
		Type: reflect.String,
		CLI: false,
		Required: false,
		Listeners: make([]chan interface{}, 0),
	}
}

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

// Add up a new configuration parameter
func (c *_Config) add(option, short, description string, dfault interface{}, dtype reflect.Kind, cli, required bool) (_ConfigEntry, error) {
	entry := DefaultEntry()
	entry.Long = option
	entry.Short = short
	entry.Description = description
	entry.Value = dfault
	entry.Type = dtype
	entry.CLI = cli
	entry.Required = required
	c.Lock()
	defer c.Unlock()
	c.entries[option] = entry
	return entry, nil
}

func validate_input(input interface{}, cfg _ConfigEntry) (error) {
	reflect.ValueOf(input).Kind()
	//error_string := fmt.Sprintf("Config Error: %s is not a kind of %s", v, t)
	return nil
	//return errors.New(error_string)
}

// // Sets a configuration parameter
// func (c *_Config) set(k string, v interface{}) (_ConfigEntry, error) {
// 	entry := ConfigEntry{Value: v, Type: t}
// 	if t == reflect.ValueOf(v).Kind() {
// 		c.Lock()
// 		defer c.Unlock()
// 		c.entries[k] = entry
// 		return entry, nil
// 	}
// 	error_string := fmt.Sprintf("gapp: %s is not a kind of %s", v, t)
// 	return ConfigEntry{}, errors.New(error_string)
// }

// // Gets information for a configuration parameter
// func (c *_Config) info(k string) (ConfigEntry) {
// 	c.RLock()
// 	defer c.RUnlock()
// 	return c.entries[k]
// }

// // Gets a configuration value
// func (c *_Config) get(k string) (interface{}) {
// 	return c.info(k).Value
// }
