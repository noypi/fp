package fp

import (
	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestRangeList(c *C) {
	list := []int{1, 2, 3, 4, 5}

	p := RangeList(func(a, i AnyVal) (ret AnyVal, skip bool) {
		c.Log("i=", i)
		ret = a.(int) * 3
		return
	}, list)

	for i := range p {
		c.Log(i.(int))
	}
}

func (suite *MySuite) TestRangeDict(c *C) {
	list := map[string]int{
		"david": 4,
		"mama":  22,
		"papa":  31,
	}

	p := RangeDict(func(v, k AnyVal) (ret AnyVal, skip bool) {
		c.Logf("v=%v, k=%v", v, k)
		ret = v
		if v.(int) != 4 {
			skip = true
		}
		return
	}, list)

	hasItem := false
	for i := range p {
		hasItem = true
		c.Assert(i.(int) == 4, Equals, true)
	}

	c.Assert(hasItem, Equals, true)
}

func (suite *MySuite) TestParallelLoop(c *C) {
	q := ParallelLoop(func(a, i AnyVal) (ret AnyVal, skip bool) {
		ret = a
		return
	}, func(a AnyVal) (ret AnyVal, skip bool) {
		ret = a.(int) * 2
		return
	}, []int{10, 31, 53})

	c.Assert((<-q).(int), Equals, 20)
	c.Assert((<-q).(int), Equals, 62)
	c.Assert((<-q).(int), Equals, 106)
	c.Assert(<-q, Equals, nil)

}
