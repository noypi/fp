package fp

type AnyVal interface{}
type PromiseChan chan AnyVal
type LazyChan chan PromiseChan
type LazyInChan PromiseChan
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

//-
type Ranger func(Func2, AnyVal, ...int) PromiseChan

func makepromise(chanlen ...int) (p PromiseChan) {
	if 0 < len(chanlen) {
		p = make(PromiseChan, chanlen[0])
	} else {
		p = make(PromiseChan, 1)
	}
	return
}
