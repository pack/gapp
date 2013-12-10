package gapp

import (
	"testing"
	"reflect"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("Empty configuration should have no entries", t, func() {
		So(len(Config.keys()), ShouldEqual, 0)
	})

	Convey("After setting an entry it should be retrievable", t, func() {
		resp, err := Config.set("power", 9000, reflect.Int)
		So(err, ShouldEqual, nil)
		So(resp.Value, ShouldEqual, 9000)
	})

	Convey("Adding an invalid entry should throw an error", t, func() {
		_, err := Config.set("power", 9000, reflect.String)
		So(err, ShouldNotEqual, nil)
	})
}
