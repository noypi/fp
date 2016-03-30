package fp

func Future(f Func0) (p *Promise) {
	p = makepromise()
	go func() {
		msg := new(qMsg)
		a, err := f()
		msg.a = a
		msg.err = err
		p.send(msg)
		p.close()
	}()
	return
}

func FutureWithNotify(fn func(chan AnyVal) (AnyVal, error)) (q *Promise, notify *Promise) {
	ch := make(chan AnyVal, 1)
	notify = PipeChan(ch)
	q = Future(func() (ret AnyVal, err error) {
		ret, err = fn(ch)
		close(ch)
		return
	})

	return
}

func Future1(f Func1, param AnyVal) (p *Promise) {
	p = makepromise()
	go func() {
		msg := new(qMsg)
		ret, err := f(param)
		msg.a = ret
		msg.err = err
		p.send(msg)
		p.close()
	}()
	return
}

func FutureN(f FuncN, params ...AnyVal) (p *Promise) {
	p = makepromise()
	go func() {
		ret, err := f(params...)
		msg := new(qMsg)
		msg.a = ret
		msg.err = err
		p.send(msg)
		p.close()
	}()
	return
}
