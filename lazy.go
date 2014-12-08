package fp

// !!! needs to be tested, Lazies should return closures,...
// use double arrows, example, <-<-l
func Lazy(f FuncAny0) LazyFn {
	return func() (PromiseChan) {
		return Promise(func()(ret AnyVal, skip bool) {
			ret = f()
			return
		})
	}	
}

// !!! not yet tested
// caller closes l to stop goroutine
func LazyIn1(f Func1, l LazyInChan) (p PromiseChan) {
	// was purposely set to 0, so, results are received in the expected order
	p = make(PromiseChan, 0)
	go func() {
		for x := range l {
			if ret, skip := f(x); !skip {
				p <- ret
			}
		}
		close(p)
	}()

	return
}

// !!! not yet tested
// caller closes l to stop goroutine
func LazyInAsync1(f Func1, qL LazyInChan, chanlen ...int) (p PromiseChan) {

	var wg WaitGroup
	p = makepromise(chanlen...)
	go func() {
		for x := range qL {
			wg.Add(Async1(func(a AnyVal) (ret AnyVal) {
				var skip bool
				if ret, skip = f(a); !skip {
					p <- ret
				}
				return
			}, x))
		}
		wg.Wait()
		close(p)
	}()

	return
}

// !!! not yet tested
func LazyIn2(f Func2, qL1, qL2 LazyInChan) (p PromiseChan) {
	// was purposely set to 0, so, results are received in the expected order
	p = make(PromiseChan, 0)
	go func() {
		q := ZipGen2(qL1, qL2)
		for a := range q {
			tuple := a.(*Tuple2)
			if ret, skip := f(tuple.A, tuple.B); !skip {
				p <- ret
			}
		}
		close(p)
	}()

	return
}
