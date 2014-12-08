package fp

import (
	"reflect"
)

//!!! not tested
// list comprehension
func ListCompr(f FuncAny1, alist AnyVal, predicates ...FuncBool1) (p PromiseChan) {
	p = make(PromiseChan, 1)
	va := reflect.ValueOf(alist)

	go func() {
		for i := 0; i < va.Len(); i++ {
			a := va.Index(i).Interface()
			test_predicates1(f, a, p, predicates...)
		}
		close(p)
	}()

	return

}

//!!! not tested
// list comprehension
func ListComprGen(f FuncAny1, in LazyInChan, predicates ...FuncBool1) (p PromiseChan) {
	p = make(PromiseChan, 1)

	go func() {
		for a := range in {
			test_predicates1(f, a, p, predicates...)
		}
		close(p)
	}()

	return

}

func test_predicates1(f FuncAny1, a AnyVal, outchan PromiseChan, predicates ...FuncBool1) {
	trueCnt := 0
	for _, pred := range predicates {
		if pred(a) {
			trueCnt++
		} else {
			break
		}
	}
	if len(predicates) == trueCnt {
		outchan <- f(a)
	}
}

//!!! not tested
// list comprehension, 2 lists
func ListCompr2(f FuncAny2, alist, blist AnyVal, predicates ...FuncBool2) (p PromiseChan) {

	p = make(PromiseChan, 1)
	go func() {
		q1 := Zip2(alist, blist)
		test_predicates2(f, LazyInChan(q1), p, predicates...)
		close(p)
	}()

	return

}

//!!! not tested
// list comprehension, 2 lists
func ListComprGen2(f FuncAny2, a, b LazyInChan, predicates ...FuncBool2) (p PromiseChan) {

	p = make(PromiseChan, 1)
	go func() {
		q1 := ZipGen2(a, b)
		test_predicates2(f, LazyInChan(q1), p, predicates...)
		close(p)
	}()

	return

}

func test_predicates2(f FuncAny2, tupleChan LazyInChan, outchan PromiseChan, predicates ...FuncBool2) {
	var tuple *Tuple2
	for data := range tupleChan {
		tuple = data.(*Tuple2)
		trueCnt := 0
		for _, pred := range predicates {
			if pred(tuple.A, tuple.B) {
				trueCnt++
			} else {
				break
			}
		}
		if len(predicates) == trueCnt {
			outchan <- f(tuple.A, tuple.B)
		}
	}
}
