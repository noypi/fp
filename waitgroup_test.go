package fp

import (
	. "gopkg.in/check.v1"
	"time"
)

func (suite *MySuite) TestWaitGroup(c *C) {
	var wg WaitGroup
	wg.Wait() // test if it won't panic, waiting for empty

	ch := make(PromiseChan, 1)
	wg.Add(ch)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch <- 12
		close(ch)
	}()

	ch1 := make(PromiseChan, 1)
	wg.Add(ch1)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- 12
		close(ch1)
	}()

	wg.Wait()
}

func (suite *MySuite) TestWaitGroup1(c *C) {
	var wg WaitGroup
	ch := make(PromiseChan, 1)
	wg.Add(ch)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch <- 12
		close(ch)
	}()

	ch1 := make(PromiseChan, 1)
	wg.Add(ch1)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- 12
		close(ch1)
	}()

	wg.WaitN(2)

}

func (suite *MySuite) TestWaitGroup2(c *C) {
	var wg WaitGroup

	cnt := 0
	f := func() {
		cnt++
	}
	for i := 0; i < 100; i++ {
		wg.Add(Async(f), Async(f))
	}

	wg.Wait()
	c.Assert(cnt, Equals, 200)
}
