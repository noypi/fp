package fp

type WaitGroup struct {
	v VectorChan
}

// sets initial capacity
func NewWaitGroup(capacity int) *WaitGroup {
	wg := new(WaitGroup)
	wg.v.q = make(chan interface{}, capacity)
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
		<-this.v.getchan()
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
		x := <-this.v.getchan()
		flushchan(x)
	}

}

func Flush(p *Promise) {
	if nil == p {
		return
	}
	for _ = range p.q {
	}
}

func flushchan(a interface{}) {
	if c, b := a.(*Promise); b {
		Flush(c)
	}
}
