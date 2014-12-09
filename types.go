package fp

import (
	"reflect"
	"sync"
)

type AnyVal interface{}
type ChanAny chan AnyVal
type Promise struct {
	q      ChanAny
	m      sync.Mutex
	vq     reflect.Value
	closed bool
}
type LazyFn func() *Promise
type LazyFn1 func(a AnyVal) *Promise
type LazyFnN func(a ...AnyVal) *Promise

type Tuple2 struct {
	A AnyVal
	B AnyVal
}

type FuncQ func() (ares, aerr, anotify AnyVal)
type FuncChainQ func(ares, aerr, anotify AnyVal) (bresult, berror, bnotify AnyVal)

//-
type Func0 func() (AnyVal, bool)
type Func1 func(a AnyVal) (AnyVal, bool)
type Func2 func(a, b AnyVal) (AnyVal, bool)

//-
type FuncVoid0 func()
type FuncVoid1 func(a AnyVal)
type FuncVoid2 func(a, b AnyVal)

//-
type FuncBool0 func() bool
type FuncBool1 func(a AnyVal) bool
type FuncBool2 func(a, b AnyVal) bool

//-
type FuncAny0 func() AnyVal
type FuncAny1 func(a AnyVal) AnyVal
type FuncAny2 func(a, b AnyVal) AnyVal
type FuncAnyN func(n ...AnyVal) AnyVal

//-
type Ranger func(Func2, AnyVal, ...int) *Promise

// if ok is false, the promise is closed without receiving any.
func (this Promise) Recv() (a AnyVal, ok bool) {
	av, ok := this.vq.Recv()
	a = av.Interface()
	return
}

func (this Promise) IsEmpty() bool {
	return 0 == len(this.q)
}

func (this *Promise) Close() {
	this.m.Lock()
	defer this.m.Unlock()
	close(this.q)
	this.closed = true
}

func (this *Promise) Q() ChanAny {
	return this.q
}

func (this *Promise) send(a AnyVal) {
	this.q <- a
}

func makepromise(chanlen ...int) (p *Promise) {
	p = new(Promise)
	if 0 < len(chanlen) {
		if 0 == chanlen[0] {
			panic("does not support chanlen=0. could break promise.")
		}
		p.q = make(ChanAny, chanlen[0])
	} else {
		p.q = make(ChanAny, 1)
	}
	p.vq = reflect.ValueOf(p.q)
	return
}
