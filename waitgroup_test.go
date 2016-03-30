package fp

import (
	"time"

	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestWaitGroup(c *C) {
	var wg WaitGroup
	wg.Wait() // test if it won't panic, waiting for empty

	ch := makepromise()
	wg.Add(ch)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch.send(&qMsg{a: 12})
		ch.close()
	}()

	ch1 := makepromise()
	wg.Add(ch1)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1.send(&qMsg{a: 12})
		ch1.close()
	}()

	wg.Wait()
}

func (suite *MySuite) TestWaitGroup1(c *C) {
	var wg WaitGroup
	ch := makepromise()
	wg.Add(ch)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch.send(&qMsg{a: 12})
		ch.close()
	}()

	ch1 := makepromise()
	wg.Add(ch1)

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch.send(&qMsg{a: 12})
		ch.close()
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
