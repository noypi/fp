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
		for x := range qL.q {
			wg.Add(Async1(func(a AnyVal) (ret AnyVal) {
				msg := new(qMsg)
				msg.a, msg.err = f(a)
				p.send(msg)
				return
			}, x))
		}
		wg.Wait()
		p.close()
	}()

	return
}

// !!! not yet tested
func lazyInParams(f FuncN, qL *Promise, n ...AnyVal) (p *Promise) {
	p = makepromise()
	go func() {
		for a := range qL.q {
			msg := new(qMsg)

			if 0 < len(n) {
				params := append([]AnyVal{a}, n...)
				msg.a, msg.err = f(params...)
			} else {
				msg.a, msg.err = f(a)
			}
			p.send(msg)
		}
		p.close()
	}()

	return
}

func LazyInParams(f FuncN, qL *Promise, n ...AnyVal) (p *Promise) {
	return lazyInParams(f, qL, n...)
}

// !!! not yet tested
func lazyIn2(f Func2, qL1, qL2 *Promise) (p *Promise) {
	p = makepromise()
	go func() {
		q := ZipGen2(qL1, qL2)
		for a := range q.q {
			msg := new(qMsg)
			tuple := a.a.(*Tuple2)
			ret, err := f(tuple.A, tuple.B)
			msg.a = ret
			msg.err = err
			p.send(msg)
		}
		p.close()
	}()

	return
}

func LazyIn2(f Func2, qL1, qL2 *Promise) (p *Promise) {
	return lazyIn2(f, qL1, qL2)
}
