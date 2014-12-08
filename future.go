package fp

func Q(f FuncQ) (cq *ChainQ) {
	cq = NewChain(f)
	cq.Run()
	return
}

func Promise(f Func0) (p PromiseChan) {
	p = make(PromiseChan, 0)
	go func() {
		if ret, skip := f(); !skip {
			p <- ret
		}
		close(p)
	}()
	return
}
