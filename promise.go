package fp

import (
	"reflect"
)

func (this Promise) IsEmpty() bool {
	return 0 == len(this.q)
}

func (this *Promise) SetContinueOnError() {
	panic("soon")
}

func (this *Promise) close() {
	this.m.Lock()
	close(this.q)
	this.closed = true
	this.m.Unlock()
}

func (this *Promise) Then(fns ...Func1) (p *Promise) {
	p = makepromise()
	go func() {
		var res2 AnyVal
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
	this.q <- a
}

func Fcall(v AnyVal) *Promise {
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
		p.q = make(qChan, 1)
	}
	p.vq = reflect.ValueOf(p.q)
	return
}
