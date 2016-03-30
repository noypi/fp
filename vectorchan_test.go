package fp

import (
	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestVectorChan(c *C) {
	var v VectorChan
	ch := makepromise(2)
	v.Add(ch)
	c.Assert(v.Len(), Equals, 1)
	c.Assert(v.Cap(), Equals, 4)

	ch.send(&qMsg{a: 12})
	ch.send(&qMsg{a: 14})
	c.Assert(v.Len(), Equals, 1)

	// full
	for i := 0; i < 7; i++ {
		v.Add(makepromise())
	}
	c.Assert(v.Len(), Equals, 8)
	c.Assert(v.Cap(), Equals, 8)

	// doubled
	v.Add(makepromise())
	c.Assert(v.Len(), Equals, 9)
	c.Assert(v.Cap(), Equals, 16)

	ch1 := <-v.getchan()

	a := <-ch1.(*Promise).q
	c.Assert(a.a, Equals, 12)
	a = <-ch1.(*Promise).q
	c.Assert(a.a, Equals, 14)
	c.Assert(v.Len(), Equals, 8)
	c.Assert(v.Cap(), Equals, 16)

	// var params
	qs := []AnyVal{}
	for i := 0; i < 9; i++ {
		qs = append(qs, makepromise())
	}
	v.Add(qs...)
	c.Assert(v.Len(), Equals, 17)
	c.Assert(v.Cap(), Equals, 32)

}
