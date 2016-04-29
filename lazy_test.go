package fp

import (
	. "gopkg.in/check.v1"
	//"time"
)

func (suite *MySuite) TestLazy(c *C) {
	/* TODO: fix this
	seq := []string{}

	qLazy := Lazy(func() interface{} {
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
