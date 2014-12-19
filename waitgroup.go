package fp

type WaitGroup struct {
	v VectorChan
}

// sets initial capacity
func NewWaitGroup(capacity int) *WaitGroup {
	wg := new(WaitGroup)
	wg.v.q = makepromise(capacity)
	return wg
}

func (this *WaitGroup) Add(qs ...*Promise) {
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
	for {
		if 0 == this.v.Len() {
			break
		}
		if x, ok := this.v.Recv(); ok {
			flushchan(x)
		}
	}

}

func Flush(p *Promise) {
	if nil == p {
		return
	}
	for _ = range p.q {
	}
}

func flushchan(a AnyVal) {
	if c, b := a.(*Promise); b {
		Flush(c)
	}
}
