package fp_test

import (
	"errors"
	"fmt"
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

		func (this Resource) Get(n ...AnyVal) (p *Promise) {
			return r.fget(n...)
		}

		func (this Resource) Put(n ...AnyVal) (p *Promise) {
			return r.fput(n...)
		}
	*/

}

func ExampleQ() {

	q := Future(func() (AnyVal, error) {
		// do some work
		return "resolved", errors.New("failed")
	})

	q = q.Then(func(a AnyVal) (AnyVal, error) {
		// on resolved
		return "resolved", errors.New("failed")
	}, func(a AnyVal) (AnyVal, error) {
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

	q := RangeList(func(x, index AnyVal) (ret AnyVal, err error) {
		// assign result to be sent to promise
		ret = expensive_run_with_res(x.(int))
		// ignore some elements
		if 0 == (index.(int) % 2) {
			err = errors.New("some error")
		}
		// -- can also ignore base on ret?
		return
	}, inputs)

	// print results
	for a := range q.Q() {
		log.Println("result a=", a)
	}
}

func ExampleListCompr() {
	list := []int{1, 2, 3, 4, 5, 6}

	q := ListCompr(func(a AnyVal) (ret AnyVal) {
		return a.(int) * 3
	}, list, func(a AnyVal) bool {
		return 0 == (a.(int) % 2)
	})

	// receive inputs
	for a := range q.Q() {
		log.Println("result a=", a)
	}

}

func ExampleListCompr2() {
	alist := []int{1, 2, 3, 4, 5, 6}
	blist := []int{2, 3, 4, 5, 6, 7}

	q := ListCompr2(func(a, b AnyVal) (ret AnyVal) {
		return a.(int) + b.(int)
	}, alist, blist, func(a, b AnyVal) bool {
		return 0 == (a.(int) % 2)
	})

	// receive inputs
	for a := range q.Q() {
		log.Println("result a=", a)
	}

}

func ExampleAsync() {
	q := Async(func() {
		// do something
	})

	// wait
	q.Recv()
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
	p := Future(func() (ret AnyVal, err error) {
		// do something...
		ret = 1 // some value

		// default is false
		// if true, the result will not be sent to the promise channel
		err = nil
		return
	})

	val, ok := p.Recv()
	fmt.Println(val, ok)
}
