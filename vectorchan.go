package fp

import (
	"sync"
)

const cDefaultCapacity = 4

// a channel that can grow in size when needed
type VectorChan struct {
	q      PromiseChan
	mutex  sync.Mutex
	closed bool
}

func NewVectorChan(capacity int) *VectorChan {
	v := new(VectorChan)
	v.q = make(PromiseChan, capacity)
	return v
}

func (this *VectorChan) Add(in ...PromiseChan) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for _, a := range in {
		this.send(a)
	}

}

func (this *VectorChan) Send(as ...AnyVal) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for _, a := range as {
		this.send(a)
	}
}

func (this *VectorChan) Close() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	close(this.q)
	this.closed = true
}

func (this *VectorChan) send(x AnyVal) {
	if this.closed {
		panic("VectorChan already closed")
	}

	if nil == this.q {
		this.q = make(PromiseChan, cDefaultCapacity)
	}

	if 0 == (cap(this.q) - len(this.q)) {
		q1 := make(PromiseChan, cap(this.q)<<1)
		close(this.q)

		for a := range this.q {
			q1 <- a
		}

		this.q = q1
	}

	this.q <- x
}

func (this *VectorChan) Recv() (a AnyVal, ok bool) {
	q := this.getchan()
	ok = (0 < len(q))
	if ok {
		a = <-q
	}
	return
}

func (this VectorChan) Len() int {
	return len(this.getchan())
}

func (this VectorChan) Cap() int {
	return cap(this.getchan())
}

func (this *VectorChan) getchan() PromiseChan {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	// can still attempt to read after close, but not send on it
	return this.q
	// !!! what happens if this was returned then vector grows after?
}
