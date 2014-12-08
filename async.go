package fp

// useful because this returns a channel, which receives notifications when done
func Async(f FuncVoid0) (p PromiseChan) {
	p = make(PromiseChan, 1)
	go func() {
		f()
		p <- true
		close(p)
	}()
	return
}

func Async1(f FuncAny1, param AnyVal) (p PromiseChan) {
	p = make(PromiseChan, 1)
	go func() {
		p <- f(param)
		close(p)
	}()
	return
}
