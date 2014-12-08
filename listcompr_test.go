package fp

import (
	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestListCompr(c *C) {

	list := []int{1, 2, 3, 4, 5, 6}

	q := ListCompr(func(a AnyVal) (ret AnyVal) {
		return a.(int) * 3
	}, list, func(a AnyVal) bool {
		return 0 == (a.(int) % 2)
	})

	c.Assert(<-q, Equals, 6)
	c.Assert(<-q, Equals, 12)
	c.Assert(<-q, Equals, 18)
	c.Assert(<-q, Equals, nil)
}

func (suite *MySuite) TestListCompr2(c *C) {

	alist := []int{1, 2, 3, 4, 5, 6}
	blist := []int{2, 3, 4, 5, 6, 7}

	q := ListCompr2(func(a, b AnyVal) (ret AnyVal) {
		return a.(int) + b.(int)

	}, alist, blist, func(a, b AnyVal) bool {
		return 0 == (a.(int) % 2)
	})

	c.Assert(<-q, Equals, 5)
	c.Assert(<-q, Equals, 9)
	c.Assert(<-q, Equals, 13)
	c.Assert(<-q, Equals, nil)
}
