package fp

import (
	"reflect"

	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestListCompr(c *C) {

	list := []int{1, 2, 3, 4, 5, 6}

	q := ListCompr(func(a interface{}) (ret interface{}) {
		return a.(int) * 3
	}, list, func(a interface{}) bool {
		return 0 == (a.(int) % 2)
	})

	res := []interface{}{
		6,
		12,
		18,
		nil,
	}

	i := 0
	Flush(q.Then(func(a interface{}) (interface{}, error) {
		if !reflect.DeepEqual(a, res[i]) {
			panic("not equal")
		}
		i++
		return nil, nil
	}))

}

func (suite *MySuite) TestListCompr2(c *C) {

	alist := []int{1, 2, 3, 4, 5, 6}
	blist := []int{2, 3, 4, 5, 6, 7}

	q := ListCompr2(func(a, b interface{}) (ret interface{}) {
		return a.(int) + b.(int)

	}, alist, blist, func(a, b interface{}) bool {
		return 0 == (a.(int) % 2)
	})

	type _res struct {
		v  interface{}
		ok bool
	}
	res := []_res{
		{5, true},
		{9, true},
		{13, true},
		{nil, false},
	}
	i := 0
	Flush(q.Then(func(a interface{}) (interface{}, error) {
		if !reflect.DeepEqual(a, res[i].v) {
			panic("not equal")
		}
		i++
		return nil, nil
	}))
}
