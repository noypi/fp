package fp

// !!! not yet tested
func Lazy(f FuncAny0) LazyFn {
	return func() (*Promise) {
		return Future(func()(ret AnyVal, skip bool) {
			ret = f()
			return
		})
	}
}

// !!! not yet tested
func Lazy1(f FuncAny1) (LazyFn1) {
	return func(a AnyVal) (*Promise) {
		return Future(func()(ret AnyVal, skip bool) {
			ret = f(a)
			return
		})
	}
}

// !!! not yet tested
// caller closes l to stop goroutine
func LazyInAsync1(f Func1, qL *Promise, chanlen ...int) (p *Promise) {

	var wg WaitGroup
	p = makepromise()
	go func() {
		for x := range qL.Q() {
			wg.Add(Async1(func(a AnyVal) (ret AnyVal) {
				var skip bool
				if ret, skip = f(a); !skip {
					p.send(ret)
				}
				return
			}, x))
		}
		wg.Wait()
		p.Close()
	}()

	return
}

// !!! not yet tested
func LazyIn2(f Func2, qL1, qL2 *Promise) (p *Promise) {
	// was purposely set to 0, so, results are received in the expected order
	p = makepromise()
	go func() {
		q := ZipGen2(qL1, qL2)
		for a := range q.Q() {
			tuple := a.(*Tuple2)
			if ret, skip := f(tuple.A, tuple.B); !skip {
				p.send(ret)
			}
		}
		p.Close()
	}()

	return
}
