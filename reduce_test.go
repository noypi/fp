package fp

import (
	"reflect"

	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestReduceParams(c *C) {
	add := func(a, b int) int {
		return a + b
	}

	res := ReduceParams(add, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	bCalled := false
	Flush(res.Then(func(a interface{}) (interface{}, error) {
		if !reflect.DeepEqual(a, 45) {
			panic("not equal")
		}
		bCalled = true
		return nil, nil
	}))
	c.Assert(bCalled, Equals, true)

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
	bCalled := false
	Flush(res.Then(func(a interface{}) (interface{}, error) {
		if !reflect.DeepEqual(a, 23) {
			panic("not equal")
		}
		bCalled = true
		return nil, nil
	}))
	c.Assert(bCalled, Equals, true)

}

func (suite *MySuite) TestReduceFuncsErr(c *C) {
	double := func(a, b int) int {
		return a * 2
	}

	res := ReduceFuncs(3, []func(int, int) int{double})

	bCalledResolved := false
	bCalledFailed := false
	Flush(res.Then(func(a interface{}) (interface{}, error) {
		bCalledResolved = true
		return nil, nil
	}, func(a interface{}) (interface{}, error) {
		bCalledFailed = true
		return nil, nil
	}))

	c.Assert(bCalledResolved, Equals, false)
	c.Assert(bCalledFailed, Equals, true)

}
