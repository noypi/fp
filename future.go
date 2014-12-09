package fp

func Q(f FuncQ) (cq *ChainQ) {
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
		p.Close()
	}()
	return
}

func fnFutureQ(f Func0, pin *Promise) (p *Promise) {
	p = pin
	go func() {
		if ret, skip := f(); !skip {
			p.send(ret)
		}
		p.Close()
	}()
	return
}
