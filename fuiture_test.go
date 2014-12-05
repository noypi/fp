package fp

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func Test(t *testing.T) { TestingT(t) }

func (suite *MySuite) TestPromise(c *C) {

	didCall := false
	p := Promise(func() (ret AnyVal, skip bool) {
		fmt.Print("adrian guwapo")
		didCall = true
		return
	})

	wasTested := false
	for i := 0; i < 50; i++ {
		fmt.Print("*")
		if 20 == i {
			time.Sleep(300 * time.Millisecond)
		} else if 21 == i {
			c.Assert(didCall, Equals, true)
			wasTested = true
		}
	}
	c.Assert(wasTested, Equals, true)
	c.Log("adrian guwapo 2")
	<-p
}

func (suite *MySuite) TestListCompr(c *C) {

	list := []int{1, 2, 3, 4, 5, 6}

	q := ListCompr(func(a, b AnyVal) (ret AnyVal) {
		return a.(int) * 3
	}, list, func(a AnyVal) bool {
		return 0 == (a.(int) % 2)
	})

	c.Assert(<-q, Equals, 6)
	c.Assert(<-q, Equals, 12)
	c.Assert(<-q, Equals, 18)
	c.Assert(<-q, Equals, nil)
}

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
