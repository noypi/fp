package fp_test

import (
	"errors"
	"log"
	"math/rand"

	. "github.com/noypi/fp"
)

func fb(_ int) {
	// some slow func
}

func ExampleLazyN() {
	/*
		type Resource struct {
			fget FuncAnyN
			fput FuncAnyN
		}

		func NewResource(fget, fput FuncAnyN) (r *Resource) {
			r = new(Resource)
			r.fget = LazyN(fget)
			r.fput = LazyN(fput)
			return
		}

		func (this Resource) Get(n ...interface{}) (p *Promise) {
			return r.fget(n...)
		}

		func (this Resource) Put(n ...interface{}) (p *Promise) {
			return r.fput(n...)
		}
	*/

}

func ExampleQ() {

	q := Future(func() (interface{}, error) {
		// do some work
		return "resolved", errors.New("failed")
	})

	q = q.Then(func(a interface{}) (interface{}, error) {
		// on resolved
		return "resolved", errors.New("failed")
	}, func(a interface{}) (interface{}, error) {
		// on error
		return "resolved", errors.New("failed")
	})

}

func ExampleRangeList() {

	expensive_run_with_res := func(a int) int { return 0 }

	// populate inputs
	inputs := []int{}
	for i := 0; i < 10; i++ {
		inputs = append(inputs, 15+int(rand.Int31n(15)))
	}

	q := RangeList(func(x, index interface{}) (ret interface{}, err error) {
		// assign result to be sent to promise
		ret = expensive_run_with_res(x.(int))
		// ignore some elements
		if 0 == (index.(int) % 2) {
			err = errors.New("some error")
		}
		// -- can also ignore base on ret?
		return
	}, inputs)

	var wg WaitGroup
	wg.Add(q.Then(func(a interface{}) (interface{}, error) {
		log.Printf("ret=%d\n", a)
		return "resolved", nil
	}))
	wg.Wait()
}

func ExampleListCompr() {
	list := []int{1, 2, 3, 4, 5, 6}

	q := ListCompr(func(a interface{}) (ret interface{}) {
		return a.(int) * 3
	}, list, func(a interface{}) bool {
		return 0 == (a.(int) % 2)
	})

	// receive inputs
	Flush(q.Then(func(a interface{}) (interface{}, error) {
		log.Println("result a=", a)
		return "resolved", nil
	}))

}

func ExampleListCompr2() {
	alist := []int{1, 2, 3, 4, 5, 6}
	blist := []int{2, 3, 4, 5, 6, 7}

	q := ListCompr2(func(a, b interface{}) (ret interface{}) {
		return a.(int) + b.(int)
	}, alist, blist, func(a, b interface{}) bool {
		return 0 == (a.(int) % 2)
	})

	// receive inputs
	Flush(q.Then(func(a interface{}) (interface{}, error) {
		log.Println("result a=", a)
		return "resolved", nil
	}))

}

func ExampleAsync() {
	q := Async(func() {
		// do something
	})

	// wait
	Flush(q)
}

func ExampleWaitGroup() {
	var wg WaitGroup

	wg.Add(Async(func() {
		// do something
	}), Async(func() {
		// do something
	}), Async(func() {
		// do something
	}))

	// wait for all go routines to finish
	wg.Wait()
}

func ExampleFuture() {
	p := Future(func() (ret interface{}, err error) {
		// do something...
		ret = 1 // some value

		// default is false
		// if true, the result will not be sent to the promise channel
		err = nil
		return
	})

	Flush(p.Then(func(a interface{}) (interface{}, error) {
		log.Println("result a=", a)
		return "resolved", nil
	}))
}
