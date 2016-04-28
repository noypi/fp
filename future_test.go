package fp

import (
	"fmt"
	"reflect"
	"time"

	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestFuture_x01(c *C) {
	p := Future(func() (ret interface{}, err error) {
		ret = 1
		return
	})

	Flush(p.Then(func(a interface{}) (interface{}, error) {
		if !reflect.DeepEqual(a, 1) {
			panic("not equal")
		}
		return nil, nil
	}))
}

func (suite *MySuite) TestFuture_x03(c *C) {

	didCall := false
	p := Future(func() (ret interface{}, err error) {
		fmt.Print("adrian guwapo")
		didCall = true
		return
	})
	Flush(p)

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

}

func (suite *MySuite) TestPromise_close(c *C) {
	p := makepromise()
	p.close()

	Flush(p.Then(func(a interface{}) (interface{}, error) {
		if !reflect.DeepEqual(a, nil) {
			panic("not equal")
		}
		return nil, nil
	}))

}
