package fp

import (
	. "gopkg.in/check.v1"
	"time"
)

func (suite *MySuite) TestAsync(c *C) {

	ns := []int{}
	executed := false
	q := Async(func() {
		ns = append(ns, 1)
		executed = true
		time.Sleep(300 * time.Millisecond)
	})

	ns = append(ns, 0)
	<-q
	ns = append(ns, 2)
	c.Assert(executed, Equals, true)
	c.Assert(len(ns), Equals, 3)
	c.Assert(ns[0], Equals, 0)
	c.Assert(ns[1], Equals, 1)
	c.Assert(ns[2], Equals, 2)

}
