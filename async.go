package fp

// useful because this returns a channel, which receives notifications when done
func Async(f FuncVoid0) (p *Promise) {
	p = makepromise()
	go func() {
		f()
		p.send(true)
		p.Close()
	}()
	return
}

func Async1(f FuncAny1, param AnyVal) (p *Promise) {
	p = makepromise()
	go func() {
		p.send(f(param))
		p.Close()
	}()
	return
}
