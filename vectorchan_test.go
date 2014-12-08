package fp

import (
	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestVectorChan(c *C) {
	var v VectorChan
	ch := make(PromiseChan, 2)
	v.Add(ch)
	c.Assert(v.Len(), Equals, 1)
	c.Assert(v.Cap(), Equals, 4)

	ch <- 12
	ch <- 14
	c.Assert(v.Len(), Equals, 1)

	// full
	for i := 0; i < 7; i++ {
		v.Add(make(PromiseChan))
	}
	c.Assert(v.Len(), Equals, 8)
	c.Assert(v.Cap(), Equals, 8)

	// doubled
	v.Add(make(PromiseChan))
	c.Assert(v.Len(), Equals, 9)
	c.Assert(v.Cap(), Equals, 16)

	ch1, _ := v.Recv()
	c.Assert(<-ch1.(PromiseChan), Equals, 12)
	c.Assert(<-ch1.(PromiseChan), Equals, 14)
	c.Assert(v.Len(), Equals, 8)
	c.Assert(v.Cap(), Equals, 16)

	// var params
	qs := []PromiseChan{}
	for i := 0; i < 9; i++ {
		qs = append(qs, make(PromiseChan))
	}
	v.Add(qs...)
	c.Assert(v.Len(), Equals, 17)
	c.Assert(v.Cap(), Equals, 32)

}
