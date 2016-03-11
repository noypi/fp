package fp

import (
	"reflect"
	"sync"
	"time"
)

type AnyVal interface{}
type ChanAny chan AnyVal
type Promise struct {
	q      ChanAny
	m      sync.Mutex
	vq     reflect.Value
	err    error
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
type FuncMuteQ func(ares, aerr, anotify AnyVal)
type FuncChainQ func(ares, aerr, anotify AnyVal) (bresult, berror, bnotify AnyVal)

//-
type Func0 func() (AnyVal, error)
type Func1 func(a AnyVal) (AnyVal, error)
type Func2 func(a, b AnyVal) (AnyVal, error)
type FuncN func(n ...AnyVal) (AnyVal, error)

//-
type FuncVoid0 func()
type FuncVoid1 func(a AnyVal)
type FuncVoid2 func(a, b AnyVal)
type FuncVoidN func(n ...AnyVal)

//-
type FuncBool0 func() bool
type FuncBool1 func(a AnyVal) bool
type FuncBool2 func(a, b AnyVal) bool
type FuncBoolN func(a ...AnyVal) bool

//-
type FuncAny0 func() AnyVal
type FuncAny1 func(a AnyVal) AnyVal
type FuncAny2 func(a, b AnyVal) AnyVal
type FuncAnyN func(n ...AnyVal) AnyVal

//-
type FuncTickAny func(now, previous time.Time) AnyVal
type FuncTickBool func(now, previous time.Time) bool
type FuncTickVoid func(now, previous time.Time)

//-
type Ranger func(Func2, AnyVal, ...int) *Promise
