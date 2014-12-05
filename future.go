package fp

import (
	"reflect"
	"sync"
)

type AnyVal interface{}
type PromiseChan chan AnyVal
type LazyChan chan AnyVal
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
type FuncAny0 func() (AnyVal)
type FuncAny1 func(a AnyVal) (AnyVal)
type FuncAny2 func(a, b AnyVal) (AnyVal)
//-
type Ranger func(Func2, AnyVal, ...int) PromiseChan

func makepromise(chanlen ...int) (p PromiseChan) {
	if 0 < len(chanlen) {
		p = make(PromiseChan, chanlen[0])
	} else {
		p = make(PromiseChan)
	}
	return
}

func Promise(f Func0, chanlen ...int) (p PromiseChan) {

	p = makepromise(chanlen...)
	go func() {
		defer close(p)
		if ret, skip := f(); !skip {
			p <- ret
		}
	}()
	return p
}

func RangeList(f Func2, list AnyVal, chanlen ...int) (p PromiseChan) {
	v, ok := list.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(list)
	}

	n := v.Len()
	p = makepromise(chanlen...)

	wg := new(sync.WaitGroup)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(index int, ch PromiseChan) {
			defer wg.Done()
			if ret, skip := f(v.Index(index).Interface(), index); !skip {
				ch <- ret
			}

		}(i, p)
	}

	go func() { wg.Wait(); close(p) }()

	return p
}

// calls Func2 as func(value, key)
func RangeDict(f Func2, dict AnyVal, chanlen ...int) (p PromiseChan) {
	v, ok := dict.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(dict)
	}

	n := v.Len()
	p = makepromise(chanlen...)
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for _, vk := range v.MapKeys() {
		go func(vk reflect.Value, ch PromiseChan) {
			defer wg.Done()
			if ret, skip := f(v.MapIndex(vk).Interface(), vk.Interface()); !skip {
				ch <- ret
			}

		}(vk, p)
	}

	go func() { wg.Wait(); close(p) }()

	return p
}

func Range(f Func2, listOrMap AnyVal, chanlen ...int) (p PromiseChan) {
	typ := reflect.TypeOf(listOrMap)
	var ranger Ranger
	switch typ.Kind() {
	case reflect.Slice:
		ranger = RangeList
	case reflect.Map:
		ranger = RangeDict
	default:
		return
	}

	p = ranger(f, listOrMap, chanlen...)
	return
}

//!!! not tested
// list comprehension
func ListCompr(f FuncAny2, listOrMap AnyVal, predicates ...FuncBool1) (p PromiseChan) {

	p = Range(func(a, b AnyVal) (ret AnyVal, skip bool){
		trueCnt := 0
		for _, pred := range predicates {
			if pred(a) {
				trueCnt++
			} else {
				break
			}
		}
		if len(predicates) == trueCnt {
			ret = f(a, b)
		} else {
			skip = true
		}
		return
	}, listOrMap)
	
	return
	
}

// !!! not yet tested
// caller closes l and p
func Lazy(f Func1) (p PromiseChan, l LazyChan) {
	p = make(PromiseChan)
	l = make(LazyChan)
	go func() {
		for input := range l {
			ret, _ := f(input)
			p <- ret
		}
	}()

	return
}

// !!! not yet tested
func Chain(ch LazyChan, a ...Func1) (p PromiseChan) {
	if 0 >= len(a) {
		return
	}
	p = make(PromiseChan)
	go func() {
		v, _ := a[0](<-ch)
		for i := 1; i < len(a); i++ {
			v, _ = a[i](v)
		}
		p <- v
		close(p)
	}()
	return p
}

// !!! not yet tested
func ParallelLoop(af Func2, bf Func1, aListOrMap AnyVal, chanlen ...int) (p PromiseChan) {

	p = makepromise(chanlen...)
	q1 := Range(af, aListOrMap, chanlen...)
	go func() {
		for qres := range q1 {
			if ret, skip := bf(qres); !skip {
				p <- ret
			}
		}
		close(p)
	}()

	return

}
