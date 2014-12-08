package fp

import (
	"fmt"
	. "gopkg.in/check.v1"
	"time"
)

func (suite *MySuite) TestPromise(c *C) {

	didCall := false
	p := Promise(func() (ret AnyVal, skip bool) {
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
	<-p
}
