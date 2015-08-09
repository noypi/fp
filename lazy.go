package fp

// !!! not yet tested
func Lazy(f FuncAny0) LazyFn {
	return func() *Promise {
		return Future(func() (ret AnyVal, skip bool) {
			ret = f()
			return
		})
	}
}

// !!! not yet tested
func Lazy1(f FuncAny1) LazyFn1 {
	return func(a AnyVal) *Promise {
		return Future(func() (ret AnyVal, skip bool) {
			ret = f(a)
			return
		})
	}
}

// !!! not yet tested
func LazyN(f FuncAnyN) LazyFnN {
	return func(a ...AnyVal) *Promise {
		return Future(func() (ret AnyVal, skip bool) {
			ret = f(a...)
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
		p.close()
	}()

	return
}

// !!! not yet tested
func lazyInParams(f FuncN, qL *Promise, mute bool, n ...AnyVal) (p *Promise) {
	p = makepromise()
	go func() {
		for a := range qL.Q() {
			var ret AnyVal
			var skip bool

			if 0 < len(n) {
				params := append([]AnyVal{a}, n...)
				ret, skip = f(params...)
			} else {
				ret, skip = f(a)
			}

			if !skip {
				if !mute {
					p.send(ret)
				}
			}
		}
		p.close()
	}()

	return
}

func LazyInParams(f FuncN, qL *Promise, n ...AnyVal) (p *Promise) {
	return lazyInParams(f, qL, false, n...)
}

func LazyInParamsMute(f FuncVoidN, qL *Promise, n ...AnyVal) (p *Promise) {
	return lazyInParams(func(n ...AnyVal) (ret AnyVal, skip bool) {
		f(n...)
		return
	}, qL, true, n...)
}

// !!! not yet tested
func lazyIn2(f Func2, qL1, qL2 *Promise, mute bool) (p *Promise) {
	p = makepromise()
	go func() {
		q := ZipGen2(qL1, qL2)
		for a := range q.Q() {
			tuple := a.(*Tuple2)
			if ret, skip := f(tuple.A, tuple.B); !skip {
				if !mute {
					p.send(ret)
				}
			}
		}
		p.close()
	}()

	return
}

func LazyIn2(f Func2, qL1, qL2 *Promise) (p *Promise) {
	return lazyIn2(f, qL1, qL2, false)
}

func LazyInMute2(f Func2, qL1, qL2 *Promise) (p *Promise) {
	return lazyIn2(f, qL1, qL2, true)
}
