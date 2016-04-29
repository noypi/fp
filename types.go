package fp

import (
	"time"
)

type qChan chan *qMsg
type Promise struct {
	q      qChan
	err    error
	closed bool
}
type LazyFn func() *Promise
type LazyFn1 func(a interface{}) *Promise
type LazyFnN func(a ...interface{}) *Promise

type qMsg struct {
	a   interface{}
	err error
}

type Tuple2 struct {
	A interface{}
	B interface{}
}

type FuncQ func() (ares, aerr, anotify interface{})
type FuncMuteQ func(ares, aerr, anotify interface{})
type FuncChainQ func(ares, aerr, anotify interface{}) (bresult, berror, bnotify interface{})

//-
type Func0 func() (interface{}, error)
type Func1 func(a interface{}) (interface{}, error)
type Func2 func(a, b interface{}) (interface{}, error)
type FuncN func(n ...interface{}) (interface{}, error)

//-
type FuncVoid0 func()
type FuncVoid1 func(a interface{})
type FuncVoid2 func(a, b interface{})
type FuncVoidN func(n ...interface{})

//-
type FuncBool0 func() bool
type FuncBool1 func(a interface{}) bool
type FuncBool2 func(a, b interface{}) bool
type FuncBoolN func(a ...interface{}) bool

//-
type FuncAny0 func() interface{}
type FuncAny1 func(a interface{}) interface{}
type FuncAny2 func(a, b interface{}) interface{}
type FuncAnyN func(n ...interface{}) interface{}

//-
type FuncTickAny func(now, previous time.Time) interface{}
type FuncTickBool func(now, previous time.Time) bool
type FuncTickVoid func(now, previous time.Time)

//-
type Ranger func(Func2, interface{}, ...int) *Promise
