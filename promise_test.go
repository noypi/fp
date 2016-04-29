package fp

import (
	. "gopkg.in/check.v1"
	"runtime"
	"sync/atomic"
)

func (suite *MySuite) TestPromiseThen(c *C) {
	src := make(chan interface{}, 100)
	go func() {
		for i := 0; i < 100; i++ {
			src <- i
		}
		close(src)
	}()

	work := func(a interface{}) (out interface{}, err error) {
		out = a.(int) + 1
		return
	}

	q := DistributeWorkCh(src, work, uint(runtime.NumCPU()))

	var total int32
	q = q.Then(func(a interface{}) (out interface{}, err error) {
		atomic.AddInt32(&total, int32(a.(int)))
		return
	})

	// wait
	Flush(q)
	c.Assert(total, Equals, int32(5050))
}
