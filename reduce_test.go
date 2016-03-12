package fp

import (
	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestReduceParams(c *C) {
	add := func(a, b int) int {
		return a + b
	}

	res := ReduceParams(add, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	c.Assert(<-res.Q(), Equals, 45)

}

func (suite *MySuite) TestReduceFuncs(c *C) {
	double := func(a int) int {
		return a * 2
	}
	triple := func(a int) int {
		return a * 3
	}
	add5 := func(a int) int {
		return a + 5
	}

	res := ReduceFuncs(3, []func(int) int{double, triple, add5})
	c.Assert(<-res.Q(), Equals, 23)

}

func (suite *MySuite) TestReduceFuncsErr(c *C) {
	double := func(a, b int) int {
		return a * 2
	}

	res := ReduceFuncs(3, []func(int, int) int{double})
	c.Assert(<-res.Q(), IsNil)
	c.Assert(res.Error(), NotNil)
}
