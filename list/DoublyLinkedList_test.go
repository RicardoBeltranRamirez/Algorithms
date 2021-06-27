package list

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDoublyLinkedList_Add(t *testing.T) {
	c.Convey("Subject: Add", t, func() {

		c.Convey("When the list is empty", func() {
			list := NewDoublyLinkedList()
			list.Add("element1")
			c.So(list.Size(), c.ShouldEqual, 1)
		})

		c.Convey("When adding more that one element", func() {
			list := NewDoublyLinkedList()
			list.Add("element1")
			list.Add("element2")
			c.So(list.Size(), c.ShouldEqual, 2)
		})
	})
}
func TestDoubleLinkedList_Clear(t *testing.T) {
	c.Convey("Subject: Clear", t, func() {
		list := NewDoublyLinkedList()
		list.Add("element1")
		list.Add("element2")

		list.Clear()

		c.So(list.IsEmpty(), c.ShouldEqual, true)
	})
}

func TestDoublyLinkedList_PeekFirst(t *testing.T) {
	c.Convey("Subject: PeekFirst", t, func() {

		c.Convey("When the list is empty", func() {
			list := NewDoublyLinkedList()
			value, err := list.PeekFirst()

			c.So(value, c.ShouldEqual, nil)
			c.So(err.Error(), c.ShouldEqual, "linked list is empty")

		})

		c.Convey("When the list is not empty", func() {

			list := NewDoublyLinkedList()
			list.Add("element1")
			list.Add("element2")
			value, err := list.PeekFirst()

			c.So(value, c.ShouldEqual, "element1")
			c.So(err, c.ShouldEqual, nil)
		})
	})
}

func TestDoublyLinkedList_PeekLast(t *testing.T) {
	c.Convey("Subject: PeekLast", t, func() {

		c.Convey("When the list is empty", func() {
			list := NewDoublyLinkedList()
			value, err := list.PeekLast()

			c.So(value, c.ShouldEqual, nil)
			c.So(err.Error(), c.ShouldEqual, "linked list is empty")

		})

		c.Convey("When the list is not empty", func() {
			list := NewDoublyLinkedList()
			list.Add("element1")
			list.Add("element2")
			value, err := list.PeekLast()

			c.So(value, c.ShouldEqual, "element2")
			c.So(err, c.ShouldEqual, nil)
		})
	})
}

func TestDoublyLinkedList_RemoveFirst(t *testing.T) {
	c.Convey("Subject: RemoveFirst", t, func() {

		c.Convey("When the list is empty", func() {
			list := NewDoublyLinkedList()
			value, err := list.RemoveFirst()

			c.So(value, c.ShouldEqual, nil)
			c.So(err.Error(), c.ShouldEqual, "linked list is empty")
		})

		c.Convey("When the list is not empty", func() {
			list := NewDoublyLinkedList()
			list.Add("element1")
			list.Add("element2")

			value, err := list.RemoveFirst()

			c.So(value, c.ShouldEqual, "element1")
			c.So(err, c.ShouldEqual, nil)

			first, err := list.PeekFirst()

			c.So(err, c.ShouldEqual, nil)
			c.So(first, c.ShouldEqual, "element2")

		})
	})
}

func TestDoublyLinkedList_RemoveLast(t *testing.T) {
	c.Convey("Subject: RemoveLast", t, func() {

		c.Convey("When the list is empty", func() {
			list := NewDoublyLinkedList()
			value, err := list.RemoveLast()

			c.So(value, c.ShouldEqual, nil)
			c.So(err.Error(), c.ShouldEqual, "linked list is empty")
		})

		c.Convey("When the list is not empty", func() {
			list := NewDoublyLinkedList()
			list.Add("element1")
			list.Add("element2")

			value, err := list.RemoveLast()

			c.So(value, c.ShouldEqual, "element2")
			c.So(err, c.ShouldEqual, nil)

			first, err := list.PeekLast()

			c.So(err, c.ShouldEqual, nil)
			c.So(first, c.ShouldEqual, "element1")

		})
	})
}

func TestDoublyLinkedList_RemoveAt(t *testing.T) {

	c.Convey("Subject: RemoveAt", t, func() {

		c.Convey("When the list is empty", func() {
			list := NewDoublyLinkedList()
			value, err := list.RemoveAt(0)

			c.So(value, c.ShouldEqual, nil)
			c.So(err.Error(), c.ShouldEqual, "illegal index")
		})

		c.Convey("When the list is not empty, but trying to access to an invalid index", func() {
			list := NewDoublyLinkedList()
			list.Add("element1")
			value, err := list.RemoveAt(1)

			c.So(value, c.ShouldEqual, nil)
			c.So(err.Error(), c.ShouldEqual, "illegal index")
		})

		c.Convey("When the list is not empty, and trying to access to a valid index", func() {
			list := NewDoublyLinkedList()
			list.Add("element1")
			list.Add("element2")

			value, err := list.RemoveAt(1)

			c.So(err, c.ShouldEqual, nil)
			c.So(value, c.ShouldEqual, "element2")

			first, err := list.PeekLast()
			c.So(err, c.ShouldEqual, nil)
			c.So(first, c.ShouldEqual, "element1")
		})

	})
}

func TestDoublyLinkedList_Remove(t *testing.T) {

	c.Convey("Subject: Remove", t, func() {

		c.Convey("When the list is empty", func() {
			list := NewDoublyLinkedList()
			removed, err := list.Remove("element1")

			c.So(removed, c.ShouldEqual, false)
			c.So(err, c.ShouldEqual, nil)
		})

		c.Convey("When the list is not empty, and trying to remove an element that doesn't exists", func() {
			list := NewDoublyLinkedList()
			list.Add("element1")
			removed, err := list.Remove("element2")

			c.So(removed, c.ShouldEqual, false)
			c.So(err, c.ShouldEqual, nil)
		})

		c.Convey("When the list is not empty, and trying to remove an element that exists", func() {
			list := NewDoublyLinkedList()
			list.Add("element1")
			removed, err := list.Remove("element1")

			c.So(removed, c.ShouldEqual, true)
			c.So(err, c.ShouldEqual, nil)
		})

	})
}

func TestDoublyLinkedList_Contains(t *testing.T) {
	c.Convey("Subject: Contains", t, func() {

		c.Convey("When the list is empty", func() {
			list := NewDoublyLinkedList()
			contains := list.Contains("element1")

			c.So(contains, c.ShouldEqual, false)
		})

		c.Convey("When the list is not empty", func() {
			list := NewDoublyLinkedList()
			list.Add("element1")
			list.Add("element2")
			list.Add("element3")
			list.Add("element4")
			contains := list.Contains("element2")

			c.So(contains, c.ShouldEqual, true)
		})
	})
}
