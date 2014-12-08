package fp_test

func ExampleLazy() {
	qLazy := Lazy(func() AnyVal {
		seq = append(seq, "lazy")
		return 10 + 2
	})

	// ...
	res := <-<-qLazy // use double arrows
}

func ExampleQ() {

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

	// waits for notify's goroutine to finish running
	wg.Wait()

}

func ExampleRangeList() {

	// populate inputs
	inputs := []int{}
	for i := 0; i < 10; i++ {
		inputs = append(inputs, 15+int(rand.Int31n(15)))
	}

	// concurrently executes each
	q := RangeList(func(x, index AnyVal) (ret AnyVal, skip bool) {
		// assign result to be sent to promise
		ret = expensive_run_with_res(x.(int))
		// ignore some elements
		skip = (0 == (index.(int) % 2))
		// -- can also ignore base on ret?
		return
	}, inputs)

	// receive inputs
	for a := range q {
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
	for a := range q {
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
	for a := range q {
		log.Println("result a=", a)
	}

}

func ExampleAsync() {
	q := Async(func() {
		// do something
	})

	// wait
	<-q
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

func ExamplePromise() {
	p := Promise(func() (ret AnyVal, skip bool) {
		// do something...
		ret = 1 // some value

		// default is false
		// if true, the result will not be sent to the promise channel
		skip = false
		return
	})

	for res := range p {
		// will not be here if skipped
	}
}

func ExampleChainQ_bind() {

	cq := Q(func() (ares, aerr, anotify AnyVal) {
		// do something...
		return
	})

	// chain 1
	cq1 := cq.Bind(func(ares, aerr, anot AnyVal) (bres, berr, bnot AnyVal) {
		if nil != anot {
			bnot = fmt.Sprintf("from chain1 anot=%v", anot)
		} else {
			bres = fmt.Sprintf("from chain1 ares=%v", ares)
		}
		return
	})

	// waits for a result or an error
	select {
	case res := <-cq1.Qresult:
		log.Println("WrapsAProgressBar result=", res)
	case err := <-cq1.Qerror:
		log.Println("WrapsAProgressBar error=", err)
	}
}

func ExampleLazyIn1() {

	qLazy := make(LazyInChan, 1)

	q1 := LazyIn1(func(x AnyVal) (ret AnyVal, skip bool) {
		ret = x.(int) * 2
		return
	}, qLazy)

	var wg WaitGroup
	wg.Add(Async(func() {
		// sends input
		qLazy <- 10
		qLazy <- 31
		qLazy <- 53
		close(qLazy)
	}))

	as := []int{}
	// needs to close qLazy to break from the loop below
	for a := range q1 {
		as = append(as, a.(int))
	}

	wg.Wait()

}
