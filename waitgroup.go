package fp

type WaitGroup struct {
	v VectorChan
}

// sets initial capacity
func NewWaitGroup(capacity int) *WaitGroup {
	wg := new(WaitGroup)
	wg.v.q = make(PromiseChan, capacity)
	return wg
}

func (this *WaitGroup) Add(qs ...PromiseChan) {
	for _, a := range qs {
		this.v.Add(a)
	}
}

func (this WaitGroup) WaitN(n int) {
	var cnt int
	for {
		_, _ = this.v.Recv()
		// not flushing x
		// just counting received
		cnt++
		if n == cnt {
			break
		}
	}
}

func (this WaitGroup) Wait() {
	if 0 == this.v.Len() {
		return
	}
	for {
		if x, ok := this.v.Recv(); ok {
			flushchan(x)
		} else {
			break
		}
	}

}

func flushchan(a AnyVal) {
	if c, b := a.(PromiseChan); b {
		for _ = range c {
			// close c to break
		}
	}
}
