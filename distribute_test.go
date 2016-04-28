package fp

import (
	. "gopkg.in/check.v1"
	"runtime"
	"sync/atomic"
)

func (suite *MySuite) TestDistributeWork(c *C) {
	src := make(chan interface{}, 100)
	go func() {
		for i := 1; i <= 100; i++ {
			src <- i
		}
		close(src)
	}()

	var total int32
	work := func(a interface{}) (out interface{}, err error) {
		atomic.AddInt32(&total, int32(a.(int)))
		return
	}

	q := DistributeWorkCh(src, work, uint(runtime.NumCPU()))

	// wait
	Flush(q)
	c.Assert(total, Equals, int32(5050))
}
