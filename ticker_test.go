package fp

import (
	. "gopkg.in/check.v1"
	"time"
)

func (suite *MySuite) TestTickWhile(c *C) {

	var cnt int
	q := TickWhile(func(now, previous time.Time) (ret AnyVal) {
		cnt++
		return previous
	}, 100*time.Millisecond, func(a AnyVal) bool {
		return cnt < 5
	})

	var prevPrev time.Time
	q.Recv()
	for prev := range q.Q() {
		c.Log("prev=", prev, "prevPrev=", prevPrev)
		c.Assert(prevPrev.Before(prev.(time.Time)), Equals, true)
		prevPrev = prev.(time.Time)
	}

}
