package gapp

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("Empty configuration should have no entries", t, func() {
		So(len(Config.keys()), ShouldEqual, 0)
	})

	Convey("Should be able to set defaults", t, func() {
		resp, err := Config.add("name", "n", "name of the user", "Leeroy Jenkins", reflect.String, true, true)
		So(err, ShouldEqual, nil)
		So(resp.Long, ShouldEqual, "name")
		So(resp.Short, ShouldEqual, "n")
		So(resp.CLI, ShouldEqual, true)
		So(resp.Required, ShouldEqual, true)
	})

	 Convey("After setting an entry it should be retrievable", t, func() {
		 resp, err := Config.set("power", "p", "power level", 9000, reflect.Int)
	 	So(err, ShouldEqual, nil)
	 	So(resp.Value, ShouldEqual, 9000)
	 })

	// Convey("Adding an invalid entry should throw an error", t, func() {
	// 	_, err := Config.set("power", 9000, reflect.String)
	// 	So(err, ShouldNotEqual, nil)
	// })
}
