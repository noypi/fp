package main

import (
	"fmt"
	. "github.com/noypi/fp"
	"log"
	"math/rand"
	"time"
)

func main() {
	// example of using lazy
	WrapALazyFunctionSample()

	// example of using async
	WrapExpensiveProcessing()

	// example of using range
	WrapExpensiveProcessing_WithResult()

	// example of using Q
	WrapsAProgressBar()

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

	qLazy := make(LazyInChan, 1)
	defer close(qLazy)

	q1 := LazyIn1(func(x AnyVal) (ret AnyVal, skip bool) {
		ret = fb(x.(int))
		return
	}, qLazy)

	// inputs
	as := []int{26, 27, 29, 0, 1, 2, 26, 27, 29, 0, 1, 2, 26, 27, 29, 0, 1, 2}

	// range will concurrently execute each
	q2 := Range(func(a, i AnyVal) (ret AnyVal, skip bool) {
		qLazy <- a
		ret = &Tuple2{
			A: <-q1,
			B: i,
		}
		return
	}, as)

	// print results
	for a := range q2 {
		tuple := a.(*Tuple2)
		log.Printf("a=%d, i=%d\n", tuple.A.(int), tuple.B.(int))
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

	for a := range q {
		log.Println("result a=", a)
	}
}

func WrapsAProgressBar() {
	log.Println("+WrapExpensiveProcessing_WithResult()")
	defer log.Println("-WrapExpensiveProcessing_WithResult()")

	percent := 0

	cq := Q(func() (ares, aerr, anotify AnyVal) {

		// do something here
		fb(10)
		//---

		// update progressbar
		if 100 <= percent {
			ares = "It works! Successfully loaded Q! =)"
		} else {
			time.Sleep(300 * time.Millisecond)
			percent += int(rand.Int31n(10))
			anotify = percent
		}

		return
	})

	// chain 1
	cq1 := cq.Bind(func(ares, aerr, anot AnyVal) (bres, berr, bnot AnyVal) {
		log.Printf("\t chain 1 anot=%v, ares=%v\n", anot, ares)
		if nil != anot {
			bnot = fmt.Sprintf("from chain1 anot=%v", anot)
		} else {
			bres = fmt.Sprintf("from chain1 ares=%v", ares)
		}
		return
	})

	// chain 1b
	cq1b := cq.Bind(func(ares, aerr, anot AnyVal) (bres, berr, bnot AnyVal) {
		log.Printf("\t chain 1b anot=%v, ares=%v\n", anot, ares)
		if nil != anot {
			bnot = "from chain1b"
		} else {
			bres = fmt.Sprintf("from chain1b ares=%v", ares)
		}
		return
	})

	// chain 2
	cq2 := cq1.Bind(func(ares, aerr, anot AnyVal) (bres, berr, bnot AnyVal) {
		log.Printf("\tchain 2 anot=%v, ares=%v", anot, ares)

		if nil != anot {
			bnot = fmt.Sprintf("from chain2 anot=%v", anot)
		} else {
			bres = fmt.Sprintf("from chain2 ares=%v", ares)
		}

		return
	})

	// run progressbar updates asynchronously
	var wg WaitGroup

	wg.Add(Async(func() {
		for notify_message := range cq.Qnotify {
			log.Printf("%d%%...", notify_message.(int))
		}

	}), Async(func() {
		for not1 := range cq1.Qnotify {
			log.Printf("cq1 notify =%v\n", not1)
		}

	}), Async(func() {
		for not1b := range cq1b.Qnotify {
			log.Printf("cq1b notify =%v\n", not1b)
		}

	}), Async(func() {
		for not2 := range cq2.Qnotify {
			log.Printf("\t\tcq2 notify =%v\n", not2)
		}

	}))

	// waits for a result or an error
	select {
	case res := <-cq.Qresult:
		log.Println("WrapsAProgressBar result=", res)
	case err := <-cq.Qerror:
		log.Println("WrapsAProgressBar error=", err)
	}

	// waits for a result or an error
	for res := range cq1.Qresult {
		log.Println("WrapsAProgressBar cq1 result=", res)
	}

	for err := range cq1.Qerror {
		log.Println("WrapsAProgressBar cq1 error=", err)
	}

	// waits for notify's goroutine to finish running
	wg.Wait()

}
