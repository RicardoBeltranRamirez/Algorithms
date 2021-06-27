package stack

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStack_Push(t *testing.T) {
	c.Convey("Subject: Push", t, func() {
		stk := NewStack()
		stk.Push("element1")

		c.So(stk.IsEmpty(), c.ShouldEqual, false)
		c.So(stk.Size(), c.ShouldEqual, 1)
	})
}

func TestStack_Peek(t *testing.T) {
	c.Convey("Subject: Peek", t, func() {
		stk := NewStack()
		stk.Push("element1")

		value, err := stk.Peek()

		c.So(err, c.ShouldEqual, nil)
		c.So(value, c.ShouldEqual, "element1")
		c.So(stk.IsEmpty(), c.ShouldEqual, false)
		c.So(stk.Size(), c.ShouldEqual, 1)
	})
}

func TestStack_Pop(t *testing.T) {
	c.Convey("Subject: Pop", t, func() {
		stk := NewStack()
		stk.Push("element1")

		value, err := stk.Pop()

		c.So(err, c.ShouldEqual, nil)
		c.So(value, c.ShouldEqual, "element1")
		c.So(stk.IsEmpty(), c.ShouldEqual, true)
		c.So(stk.Size(), c.ShouldEqual, 0)
	})
}
