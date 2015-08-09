package fp

import (
	"sync"
)

const cDefaultCapacity = 4

// a channel that can grow in size when needed
type VectorChan struct {
	q      *Promise
	mutex  sync.Mutex
	closed bool
}

func NewVectorChan(capacity int) *VectorChan {
	v := new(VectorChan)
	v.q = makepromise(capacity)
	return v
}

func (this *VectorChan) Add(in ...*Promise) {
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
	this.q.close()
	this.closed = true
}

func (this *VectorChan) send(x AnyVal) {
	if this.closed {
		panic("VectorChan already closed")
	}

	if nil == this.q {
		this.q = makepromise(cDefaultCapacity)
	}

	if 0 == (cap(this.q.q) - len(this.q.q)) {
		q1 := makepromise(cap(this.q.q) << 1)
		close(this.q.q)

		for a := range this.q.q {
			q1.send(a)
		}

		this.q = q1
	}

	this.q.send(x)
}

func (this *VectorChan) Recv() (a AnyVal, ok bool) {
	return this.getchan().Recv()
}

func (this VectorChan) Len() int {
	return len(this.getchan().q)
}

func (this VectorChan) Cap() int {
	return cap(this.getchan().q)
}

func (this *VectorChan) getchan() *Promise {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if nil == this.q {
		this.q = makepromise(cDefaultCapacity)
	}
	// can still attempt to read after close, but not send on it
	return this.q
	// !!! what happens if this was returned then vector grows after?
}
