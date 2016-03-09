package main

import (
	. "github.com/noypi/fp"
	"log"
	"math/rand"
)

func main() {
	// example of using lazy
	WrapALazyFunctionSample()

	// example of using async
	WrapExpensiveProcessing()

	// example of using range
	WrapExpensiveProcessing_WithResult()

}

func fb(x int) int {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fb(x-1) + fb(x-2)
	}
	return 0
}

// wraps a function and is executed later with input is available
func WrapALazyFunctionSample() {
	log.Println("+WrapAFunctionSample()")
	defer log.Println("-WrapAFunctionSample()")

	as := []int{26, 27, 29, 0, 1, 2, 26, 27, 29, 0, 1, 2, 26, 27, 29, 0, 1, 2}
	// range will concurrently execute each
	qLazy := Range(func(a, i AnyVal) (ret AnyVal, skip bool) {
		ret = &Tuple2{
			A: a,
			B: i,
		}
		return
	}, as)

	q1 := LazyInAsync1(func(x AnyVal) (ret AnyVal, skip bool) {
		ret = fb(x.(*Tuple2).A.(int))
		return
	}, qLazy)

	// print results
	for a := range q1.Q() {
		log.Printf("ret=%d\n", a)
	}

}

//
// wraps expensive processing and execute in parallel
//
func expensive_run(x int) {
	log.Println("execute something... x=", x)
}
func WrapExpensiveProcessing() {
	log.Println("+WrapExpensiveProcessing()")
	defer log.Println("-WrapExpensiveProcessing()")
	var wg WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(Async(func() {
			// i here is not determined, maybe 10 at all times
			expensive_run(i)
		}))

		func(index int) {
			// pass to index, so the expected parameter is passed
			wg.Add(Async(func() {
				expensive_run(index)
			}))
		}(i) // i here is passed as parameter to the anonymous function
	}

	// waits until all async is done
	wg.Wait()
}

//
// wraps expensive processing and execute in parallel,
//
func expensive_run_with_res(x int) int {
	log.Println("execute something... x=", x)
	fb(x)

	return x
}
func WrapExpensiveProcessing_WithResult() {
	log.Println("+WrapExpensiveProcessing_WithResult()")
	defer log.Println("-WrapExpensiveProcessing_WithResult()")

	inputs := []int{}
	for i := 0; i < 10; i++ {
		inputs = append(inputs, 15+int(rand.Int31n(15)))
	}

	q := RangeList(func(x, index AnyVal) (ret AnyVal, skip bool) {
		// assign result to be sent to promise
		ret = expensive_run_with_res(x.(int))
		// ignore some elements
		skip = (0 == (index.(int) % 2))
		// -- can also ignore base on ret?
		return
	}, inputs)

	// print results
	for a := range q.Q() {
		log.Println("result a=", a)
	}
}
