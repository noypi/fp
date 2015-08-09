package fp

type ChainQ struct {
	Qerror  *Promise
	Qresult *Promise
	Qnotify *Promise
	f       FuncQ
	fcurry  FuncChainQ
	fmute   FuncMuteQ
	nexts   []*ChainQ
}

func NewChain(f FuncQ) (cq *ChainQ) {
	cq = &ChainQ{
		Qerror:  makepromise(),
		Qresult: makepromise(),
		Qnotify: makepromise(),
	}
	cq.f = f

	return
}

func (this *ChainQ) Run() {
	go func() {
		for {
			if this.call() {
				break
			}
		}
		this.close()
	}()

}

func (this *ChainQ) Bind(f FuncChainQ) (cqNew *ChainQ) {
	cqNew = NewChain(nil)
	cqNew.fcurry = func(ares, aerr, anot AnyVal) (bres, berr, bnot AnyVal) {
		bres, berr, bnot = f(ares, aerr, anot)
		cqNew.send(bres, berr, bnot)
		return
	}
	this.nexts = append(this.nexts, cqNew)
	return
}

func (this *ChainQ) BindMute(f FuncMuteQ) (cqNew *ChainQ) {
	cqNew = NewChain(nil)
	cqNew.fmute = f
	this.nexts = append(this.nexts, cqNew)
	return
}

// !!!  not tested
func (this *ChainQ) Intercept(f FuncChainQ) {
	prevCurry := this.fcurry
	this.fcurry = func(ares, aerr, anot AnyVal) (bres, berr, bnot AnyVal) {
		bres, berr, bnot = f(ares, aerr, anot)
		prevCurry(bres, berr, bnot)
		return
	}
	return
}

func (this ChainQ) call() (done bool) {
	if nil == this.f {
		done = true
		return
	}

	ares, aerr, anot := this.f()
	done = this.send(ares, aerr, anot)

	return
}

func (this *ChainQ) close() {
	this.Qnotify.close()
	this.Qerror.close()
	this.Qresult.close()

	for _, next := range this.nexts {
		next.close()
	}
}

func (this *ChainQ) send(ares, aerr, anot AnyVal) (done bool) {
	done = true

	q := Async(func() {
		for _, next := range this.nexts {
			if nil != next.fmute {
				next.fmute(ares, aerr, anot)
			}
			if nil != next.fcurry {
				next.fcurry(ares, aerr, anot)
			}
		}
	})

	if nil != aerr {
		this.Qerror.send(aerr)

	} else if nil != anot {
		this.Qnotify.send(anot)
		done = false

	} else {
		// results can receive nil values
		this.Qresult.send(ares)

	}

	<-q.Q()

	return
}
