package fp

import (
	. "gopkg.in/check.v1"
)

func doTestSuccess(c *C) {

	results := []string{}

	q := Q(func(a AnyVal) AnyVal {
		c.Assert(a, Equals, "resolved!")
		results = append(results, a.(string))
		return "from success1"
	})

	q.OnSuccess(func(a AnyVal) AnyVal {
		c.Assert(a, Equals, "from success1")
		results = append(results, a.(string))
		return "from success 2"
	})

	q.OnDone(func(a AnyVal) AnyVal {
		c.Assert(a, Equals, "from success 2")
		results = append(results, a.(string))
		return "from done 1"
	})

	qres, qsig := q.Call(func(s QSignal) {
		s.Resolve("resolved!")
	})

	c.Assert(<-qres.Q(), Equals, "from done 1")
	c.Assert(len(results), Equals, 3)
	c.Assert(qsig.IsRejected(), Equals, false)
	c.Assert(qsig.HaveSucceeded(), Equals, true)

}

func (suite *MySuite) TestSuccess(c *C) {
	for i := 0; i < 1000; i++ {
		go doTestSuccess(c)
	}
}
