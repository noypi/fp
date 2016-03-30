package fp

import (
	"errors"
	"fmt"
	"reflect"

	. "gopkg.in/check.v1"
)

func (suite *MySuite) TestRangeList(c *C) {
	list := []int{1, 2, 3, 4, 5}

	p := RangeList(func(a, i AnyVal) (ret AnyVal, err error) {
		fmt.Println("a=", a)
		ret = a.(int) * 3
		return
	}, list, 3)

	res := []AnyVal{
		3, 6, 9, 12, 15,
	}
	i := 0
	q3 := p.Then(func(a AnyVal) (AnyVal, error) {
		fmt.Println("test rangelist a=", a, ", res[i]=", res[i])
		if !reflect.DeepEqual(a, res[i]) {
			panic("not equal")
		}
		i++
		return "resolved", nil
	})

	Flush(q3)

}

func (suite *MySuite) TestRangeDict(c *C) {
	list := map[string]int{
		"david": 4,
		"mama":  22,
		"papa":  31,
	}

	p := RangeDict(func(v, k AnyVal) (ret AnyVal, err error) {
		fmt.Println("TestRangeDict v=%v, k=%v", v, k)
		ret = v
		if v.(int) != 4 {
			err = errors.New(k.(string))
		}
		return
	}, list)

	hasItem := false
	hasPapa := false
	hasMama := false

	Flush(p.Then(func(a AnyVal) (AnyVal, error) {
		fmt.Println("TestRangeDict Then resolved a=", a)
		hasItem = true
		if !reflect.DeepEqual(a, 4) {
			panic("not equal")
		}
		return nil, nil
	}, func(a AnyVal) (AnyVal, error) {
		fmt.Println("TestRangeDict Then failed a=", a)
		if a.(error).Error() == "mama" {
			hasMama = true
		} else if a.(error).Error() == "papa" {
			hasPapa = true
		}

		return nil, nil

	}))

	c.Assert(hasItem, Equals, true)
	c.Assert(hasMama, Equals, true)
	c.Assert(hasPapa, Equals, true)
}

func (suite *MySuite) Disable_TestParallelLoop(c *C) {
	q := ParallelLoop(func(a, i AnyVal) (ret AnyVal, err error) {
		ret = a
		return
	}, func(a AnyVal) (ret AnyVal, err error) {
		ret = a.(int) * 2
		return
	}, []int{10, 31, 53})

	type _res struct {
		v  AnyVal
		ok bool
	}
	res := []_res{
		{20, true},
		{62, true},
		{106, true},
		{nil, false},
	}
	i := 0
	Flush(q.Then(func(a AnyVal) (AnyVal, error) {
		if !reflect.DeepEqual(a, res[i].v) {
			panic("not equal")
		}
		i++
		return nil, nil
	}))

}
