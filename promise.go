package fp

import (
	"runtime"
)

func (this *Promise) Flush() {
	Flush(this)
}

func (this Promise) IsEmpty() bool {
	return 0 == len(this.q)
}

func (this *Promise) close() {
	close(this.q)
	this.closed = true
}

func (this *Promise) Then(fns ...Func1) (p *Promise) {
	p = makepromise()
	go func() {
		var res2 interface{}
		var err2 error
		for res := range this.q {
			if nil != res.err {
				if 1 < len(fns) {
					res2, err2 = fns[1](res.err)
				}
			} else {
				if 0 < len(fns) {
					res2, err2 = fns[0](res.a)
				}
			}

			msg := new(qMsg)
			msg.a = res2
			msg.err = err2
			p.send(msg)
		}
		p.close()
	}()

	return
}

func (this *Promise) send(a *qMsg) {
	if nil == a.a && nil == a.err {
		// skip
	} else {
		this.q <- a
	}
}

func Fcall(v interface{}) *Promise {
	p := makepromise()
	msg := new(qMsg)
	msg.a = v
	p.send(msg)
	close(p.q)
	return p
}

func makepromise(chanlen ...int) (p *Promise) {
	p = new(Promise)
	if 0 < len(chanlen) {
		if 0 == chanlen[0] {
			panic("does not support chanlen=0. could break promise.")
		}
		p.q = make(qChan, chanlen[0])
	} else {
		p.q = make(qChan, runtime.NumCPU())
	}
	return
}
