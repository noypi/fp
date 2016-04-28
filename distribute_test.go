package fp

import (
	"fmt"
	. "gopkg.in/check.v1"
	"runtime"
)

func (suite *MySuite) TestDistributeWork(c *C) {
	src := make(chan interface{}, 100)
	go func() {
		for i := 0; i < 100; i++ {
			src <- i
		}
		close(src)
	}()

	work := func(a interface{}) (out interface{}, err error) {
		out = a.(int) * 10
		return
	}

	q := DistributeWorkCh(src, work, uint(runtime.NumCPU()))

	q.Then(func(a interface{}) (out interface{}, err error) {
		fmt.Println(a)
		return
	})

	// wait
	Flush(q)
}
