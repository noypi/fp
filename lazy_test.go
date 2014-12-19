package fp

import (
	. "gopkg.in/check.v1"
	//"time"
)

func (suite *MySuite) TestLazy(c *C) {
	/* TODO: fix this
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
	*/
}

func (suite *MySuite) TestLazyInAsync1(c *C) {
	q := makepromise()

	q1 := LazyInAsync1(func(x AnyVal) (ret AnyVal, skip bool) {
		ret = x.(int) * 2
		return
	}, q)

	var wg WaitGroup
	wg.Add(Async(func() {
		q.send(10)
		q.send(31)
		q.send(53)
		q.close()
	}))

	as := []int{}
	for {
		if a, ok := q1.Recv(); ok {
			as = append(as, a.(int))
		} else {
			break
		}
	}

	c.Assert(len(as), Equals, 3)
	c.Assert(as[0], Equals, 20)
	c.Assert(as[1], Equals, 62)
	c.Assert(as[2], Equals, 106)

	wg.Wait()

}
