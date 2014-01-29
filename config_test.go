package gapp

import (
	. "launchpad.net/gocheck"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) { TestingT(t) }

type EmptyConfigSuite struct{}
type PowerConfigSuite struct{ resp _ConfigEntry }

var _ = Suite(&EmptyConfigSuite{})
var _ = Suite(&PowerConfigSuite{})

func (s *EmptyConfigSuite) TearDownTest(c *C) {
	Config.Clear()
}

func (s *PowerConfigSuite) TearDownTest(c *C) {
	Config.Clear()
}

func (s *PowerConfigSuite) SetUpTest(c *C) {
	resp, err := Config.Add("power", "p", "power level", 3000, reflect.Int, false)
	c.Assert(err, IsNil)
	s.resp = resp
}

func (s *EmptyConfigSuite) TestEmpty(c *C) {
	c.Assert(len(Config.Keys()), Equals, 0)
}

func (s *EmptyConfigSuite) TestAddDefaults(c *C) {
	resp, err := Config.Add("name", "n", "user name", "Philip J. Fry", reflect.String, false)
	c.Assert(err, IsNil)
	c.Assert(resp.Long, Equals, "name")
	c.Assert(resp.Short, Equals, "n")
	c.Assert(resp.Description, Equals, "user name")
	c.Assert(resp.Value, Equals, "Philip J. Fry")
	c.Assert(resp.Type, Equals, reflect.String)
	c.Assert(resp.CLI, Equals, false)
}

func (s *PowerConfigSuite) TestAddInteger(c *C) {
	c.Assert(s.resp.Value, Equals, 3000)
	c.Assert(s.resp.Type, Equals, reflect.Int)
}

func (s *PowerConfigSuite) TestTypeEnforcement(c *C) {
	_, err := Config.Add("power", "p", "power level", "3000", reflect.Int, false)
	c.Assert(err, ErrorMatches, "Config Entry `power` is not of type `int`.*")
}

func (s *PowerConfigSuite) TestGetValue(c *C) {
	resp, ok := Config.Get("power")
	c.Assert(ok, Equals, true)
	c.Assert(resp, Equals, 3000)
}

func (s *PowerConfigSuite) TestGetAlias(c *C) {
	resp, ok := Config.Get("p")
	c.Assert(ok, Equals, true)
	c.Assert(resp, Equals, 3000)
}

func (s *PowerConfigSuite) TestInvalidGetValue(c *C) {
	resp, ok := Config.Get("youshouldfail")
	c.Assert(ok, Equals, false)
	c.Assert(resp, IsNil)
}

func (s *PowerConfigSuite) TestModifyInteger(c *C) {
	resp, err := Config.Set("power", 9000)
	c.Assert(err, Equals, nil)
	c.Assert(resp.Value, Equals, 9000)
	c.Assert(resp.Type, Equals, reflect.Int)
}

func (s *PowerConfigSuite) TestSubscription(c *C) {
	ch, err := Config.Subscribe_to("power")
	c.Assert(err, Equals, nil)
	_, err2 := Config.Set("power", 9000)
	c.Assert(err2, Equals, nil)
	value := <-ch
	c.Assert(value, Equals, 9000)
}
