package fp

func ChainQ_old(f FuncQ) (cq *ChainQ) {
	cq = NewChain(f)
	cq.Run()
	return
}

func Future(f Func0) (p *Promise) {
	p = makepromise()
	go func() {
		if ret, skip := f(); !skip {
			p.send(ret)
		}
		p.close()
	}()
	return
}

func Future1(f Func1, param AnyVal) (p *Promise) {
	p = makepromise()
	go func() {
		if ret, skip := f(param); !skip {
			p.send(ret)
		}
		p.close()
	}()
	return
}

func FutureN(f FuncN, params ...AnyVal) (p *Promise) {
	p = makepromise()
	go func() {
		if ret, skip := f(params...); !skip {
			p.send(ret)
		}
		p.close()
	}()
	return
}

func fnFutureQ(f Func0, pin *Promise) (p *Promise) {
	p = pin
	go func() {
		if ret, skip := f(); !skip {
			p.send(ret)
		}
		p.close()
	}()
	return
}
