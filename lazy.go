package fp

// !!! not yet tested
func Lazy(f FuncAny0) LazyFn {
	return func() *Promise {
		return Future(func() (AnyVal, error) {
			return f(), nil
		})
	}
}

// !!! not yet tested
func Lazy1(f FuncAny1) LazyFn1 {
	return func(a AnyVal) *Promise {
		return Future(func() (AnyVal, error) {
			return f(a), nil
		})
	}
}

// !!! not yet tested
func LazyN(f FuncAnyN) LazyFnN {
	return func(a ...AnyVal) *Promise {
		return Future(func() (AnyVal, error) {
			return f(a...), nil
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
				if ret, p.err = f(a); nil == p.err {
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
			var err error

			if 0 < len(n) {
				params := append([]AnyVal{a}, n...)
				ret, err = f(params...)
			} else {
				ret, err = f(a)
			}

			p.err = err
			if nil != err {
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
	return lazyInParams(func(n ...AnyVal) (AnyVal, error) {
		f(n...)
		return nil, nil
	}, qL, true, n...)
}

// !!! not yet tested
func lazyIn2(f Func2, qL1, qL2 *Promise, mute bool) (p *Promise) {
	p = makepromise()
	go func() {
		q := ZipGen2(qL1, qL2)
		for a := range q.Q() {
			tuple := a.(*Tuple2)
			ret, err := f(tuple.A, tuple.B)
			p.err = err
			if nil == p.err {
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
