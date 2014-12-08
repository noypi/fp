package fp

import (
	"fmt"
	. "gopkg.in/check.v1"
	"time"
)

func (suite *MySuite) TestFuture_x01(c *C) {
	p := Future(func() (ret AnyVal, skip bool) {
		ret = 1
		return
	})

	a, ok := p.Recv()
	c.Assert(a, Equals, 1)
	c.Assert(ok, Equals, true)

	a, ok = p.Recv()
	c.Assert(a, Equals, nil)
	c.Assert(ok, Equals, false)

}

func (suite *MySuite) TestFuture_x03(c *C) {

	didCall := false
	p := Future(func() (ret AnyVal, skip bool) {
		fmt.Print("adrian guwapo")
		didCall = true
		return
	})

	wasTested := false
	for i := 0; i < 50; i++ {
		fmt.Print("*")
		if 20 == i {
			time.Sleep(300 * time.Millisecond)
		} else if 21 == i {
			c.Assert(didCall, Equals, true)
			wasTested = true
		}
	}
	c.Assert(wasTested, Equals, true)
	c.Log("adrian guwapo 2")
	p.Recv()
}

func (suite *MySuite) TestPromise_close(c *C) {
	p := makepromise()
	p.Close()

	a, ok := p.Recv()
	c.Assert(a, Equals, nil)
	c.Assert(ok, Equals, false)

}
