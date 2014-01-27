package gapp

import (
	. "launchpad.net/gocheck"
	"testing"
	"reflect"
)

func Test(t *testing.T) { TestingT(t) }

type ConfigSuite struct{}
var _ = Suite(&ConfigSuite{})

func (s *ConfigSuite) TearDownTest(c *C) {
	Config.clear()
}

func (s *ConfigSuite) TestEmpty(c *C) {
	c.Assert(len(Config.keys()), Equals, 0)
}

func (s *ConfigSuite) TestAddDefaults(c *C) {
	resp, err := Config.add("name", "n", "user name", "Philip J. Fry", reflect.String, false, false)
	c.Assert(err, IsNil)
	c.Assert(resp.Long, Equals, "name")
	c.Assert(resp.Short, Equals, "n")
	c.Assert(resp.Description, Equals, "user name")
	c.Assert(resp.Value, Equals, "Philip J. Fry")
	c.Assert(resp.Type, Equals, reflect.String)
	c.Assert(resp.Required, Equals, false)
	c.Assert(resp.CLI, Equals, false)
}

func (s *ConfigSuite) TestAddInteger(c *C) {
	resp, err := Config.add("power", "p", "power level", 3000, reflect.Int, false, false)
	c.Assert(err, IsNil)
	c.Assert(resp.Value, Equals, 3000)
	c.Assert(resp.Type, Equals, reflect.Int)
}

func (s *ConfigSuite) TestTypeEnforcement(c *C) {
	_, err := Config.add("power", "p", "power level", "3000", reflect.Int, false, false)
	c.Assert(err, ErrorMatches, "Config Entry `power` is not of type `int`.*")
}

func (s *ConfigSuite) TestGetValue(c *C) {
	_, err := Config.add("power", "p", "power level", 3000, reflect.Int, false, false)
	c.Assert(err, IsNil)
	resp, ok := Config.get("power")
	c.Assert(ok, Equals, true)
	c.Assert(resp, Equals, 3000)
}

func (s *ConfigSuite) TestInvalidGetValue(c *C) {
	_, err := Config.add("power", "p", "power level", 3000, reflect.Int, false, false)
	c.Assert(err, IsNil)
	resp, ok := Config.get("pow")
	c.Assert(ok, Equals, false)
	c.Assert(resp, IsNil)
}

func (s *ConfigSuite) TestModifyInteger(c *C) {
	_, err := Config.add("power", "p", "power level", 3000, reflect.Int, false, false)
	c.Assert(err, Equals, nil)
	resp2, err2 := Config.set("power", 9000)
	c.Assert(err2, Equals, nil)
	c.Assert(resp2.Value, Equals, 9000)
	c.Assert(resp2.Type, Equals, reflect.Int)
}
