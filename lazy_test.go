package fp

import (
	. "gopkg.in/check.v1"
	"time"
)

func (suite *MySuite) TestLazy(c *C) {

	seq := []string{}

	qLazy := Lazy(func() AnyVal {
		seq = append(seq, "lazy")
		return 10 + 2
	})

	seq = append(seq, "begin")
	time.Sleep(300 * time.Millisecond)

	seq = append(seq, "after sleep1")
	time.Sleep(300 * time.Millisecond)

	seq = append(seq, "after sleep2")
	time.Sleep(300 * time.Millisecond)

	res := <-<-qLazy // use double arrows
	seq = append(seq, "got res")

	c.Assert(len(seq), Equals, 5)
	c.Assert(seq[0], Equals, "begin")
	c.Assert(seq[1], Equals, "after sleep1")
	c.Assert(seq[2], Equals, "after sleep2")
	c.Assert(seq[3], Equals, "lazy")
	c.Assert(seq[4], Equals, "got res")
	c.Assert(res, Equals, 12)
}

func (suite *MySuite) TestLazyIn(c *C) {
	qLazy := make(LazyInChan, 1)

	q1 := LazyIn1(func(x AnyVal) (ret AnyVal, skip bool) {
		ret = x.(int) * 2
		return
	}, qLazy)

	var wg WaitGroup
	wg.Add(Async(func() {
		// sends input
		qLazy <- 10
		qLazy <- 31
		qLazy <- 53
		close(qLazy)
	}))

	c.Assert((<-q1).(int), Equals, 20)
	c.Assert((<-q1).(int), Equals, 62)
	c.Assert((<-q1).(int), Equals, 106)
	c.Assert(<-q1, Equals, nil)

	wg.Wait()
}

func (suite *MySuite) TestLazyIn_x02(c *C) {
	qLazy := make(LazyInChan, 1)

	q1 := LazyIn1(func(x AnyVal) (ret AnyVal, skip bool) {
		ret = x.(int) * 2
		return
	}, qLazy)

	var wg WaitGroup
	wg.Add(Async(func() {
		// sends input
		qLazy <- 10
		qLazy <- 31
		qLazy <- 53
		close(qLazy)
	}))

	as := []int{}
	for a := range q1 {
		as = append(as, a.(int))
	}

	c.Assert(len(as), Equals, 3)
	c.Assert(as[0], Equals, 20)
	c.Assert(as[1], Equals, 62)
	c.Assert(as[2], Equals, 106)

	wg.Wait()

}

func (suite *MySuite) TestLazyInAsync1(c *C) {
	qLazy := make(LazyInChan, 1)

	q1 := LazyInAsync1(func(x AnyVal) (ret AnyVal, skip bool) {
		ret = x.(int) * 2
		return
	}, qLazy)

	var wg WaitGroup
	wg.Add(Async(func() {
		// sends input
		qLazy <- 10
		qLazy <- 31
		qLazy <- 53
		close(qLazy)
	}))

	as := []int{}
	for a := range q1 {
		as = append(as, a.(int))
	}

	c.Assert(len(as), Equals, 3)
	c.Assert(as[0], Equals, 20)
	c.Assert(as[1], Equals, 62)
	c.Assert(as[2], Equals, 106)

	wg.Wait()

}
