package fp

func Future(f Func0) (p *Promise) {
	p = makepromise()
	go func() {
		ret, err := f()
		p.err = err
		if nil == p.err {
			p.send(ret)
		}
		p.close()
	}()
	return
}

func Future1(f Func1, param AnyVal) (p *Promise) {
	p = makepromise()
	go func() {
		ret, err := f(param)
		p.err = err
		if nil == p.err {
			p.send(ret)
		}
		p.close()
	}()
	return
}

func FutureN(f FuncN, params ...AnyVal) (p *Promise) {
	p = makepromise()
	go func() {
		ret, err := f(params...)
		p.err = err
		if nil == p.err {
			p.send(ret)
		}
		p.close()
	}()
	return
}

func fnFutureQ(f Func0, pin *Promise) (p *Promise) {
	p = pin
	go func() {
		ret, err := f()
		p.err = err
		if nil == p.err {
			p.send(ret)
		}
		p.close()
	}()
	return
}
