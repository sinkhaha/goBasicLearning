package bdd

import (
	"testing"
	// 这里导入有个点，表示当前导入的包的方法都是在当前的名字空间，直接使用即可，如Convey()
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given 2 even numbers", t, func() {
		a := 3
		b := 4

		Convey("When add the two numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
