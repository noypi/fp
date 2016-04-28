package fp

// useful because this returns a channel, which receives notifications when done
func Async(f FuncVoid0) (p *Promise) {
	p = makepromise()
	go func() {
		f()
		p.close()
	}()
	return
}

func Async1(f FuncAny1, param interface{}) (p *Promise) {
	p = makepromise()
	go func() {
		msg := new(qMsg)
		msg.a = f(param)
		p.send(msg)
		p.close()
	}()
	return
}

func AsyncAnyN(f FuncAnyN, param ...interface{}) (p *Promise) {
	p = makepromise()
	go func() {
		msg := new(qMsg)
		msg.a = f(param...)
		p.send(msg)
		p.close()
	}()
	return
}
