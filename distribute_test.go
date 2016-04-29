package fp

import (
	. "gopkg.in/check.v1"
	"runtime"
	"sync/atomic"
)

func (suite *MySuite) TestDistributeWork(c *C) {
	ns := make([]int, 100)
	for i := 0; i < 100; i++ {
		ns[i] = i
	}
	qList := RangeList(func(a, i interface{}) (out interface{}, err error) {
		out = a.(int) + 1
		return
	}, ns)

	var total int32
	work := func(a interface{}) (out interface{}, err error) {
		atomic.AddInt32(&total, int32(a.(int)))
		return
	}

	q := DistributeWork(qList, work, uint(runtime.NumCPU()))

	// wait
	Flush(q)
	c.Assert(total, Equals, int32(5050))
}

func (suite *MySuite) TestDistributeWorkCh(c *C) {
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
