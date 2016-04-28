package fp

import (
	"fmt"
	"reflect"
	"time"

	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestTickWhile(c *C) {

	var cnt int
	q := TickWhile(func(now, previous time.Time) (ret interface{}) {
		cnt++
		return previous
	}, 100*time.Millisecond, func(a interface{}) bool {
		return cnt < 5
	})

	var prevPrev time.Time

	i := 0
	Flush(q.Then(func(prev interface{}) (res interface{}, err error) {
		fmt.Println("i=", i, "prev=", prev, "prevPrev=", prevPrev)
		if 0 < i {
			if !reflect.DeepEqual(prevPrev.Before(prev.(time.Time)), true) {
				panic("not equal")
			}
		}
		prevPrev = prev.(time.Time)
		i++

		return nil, nil

	}))

	c.Assert(i, Equals, 4)

}
