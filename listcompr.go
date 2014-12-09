package fp

import (
	"reflect"
)

//!!! not tested
// list comprehension
func ListCompr(f FuncAny1, alist AnyVal, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()
	va := reflect.ValueOf(alist)

	go func() {
		var wg WaitGroup
		for i := 0; i < va.Len(); i++ {
			a := va.Index(i).Interface()
			if p1 := test_predicates1(f, a, p, predicates...); nil != p1 {
				wg.Add(p1)
			}
		}

		wg.Wait()
		p.Close()
	}()

	return

}

//!!! not tested
// list comprehension
func ListComprGen(f FuncAny1, in *Promise, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()

	go func() {
		var wg WaitGroup
		for a := range in.Q() {
			if p1 := test_predicates1(f, a, p, predicates...); nil != p1 {
				wg.Add(p1)
			}
		}

		wg.Wait()
		p.Close()
	}()

	return

}

func test_predicates1(f FuncAny1, a AnyVal, outchan *Promise, predicates ...FuncBool1) (p *Promise) {
	trueCnt := 0
	for _, pred := range predicates {
		if pred(a) {
			trueCnt++
		} else {
			break
		}
	}
	if len(predicates) == trueCnt {
		p = Async(func() {
			outchan.send(f(a))
		})
	}

	return
}

//!!! not tested
// list comprehension, 2 lists
func ListCompr2(f FuncAny2, alist, blist AnyVal, predicates ...FuncBool2) (p *Promise) {

	p = makepromise()
	go func() {
		q1 := Zip2(alist, blist)
		test_predicates2(f, q1, p, predicates...)
		p.Close()
	}()

	return

}

//!!! not tested
// list comprehension, 2 lists
func ListComprGen2(f FuncAny2, a, b *Promise, predicates ...FuncBool2) (p *Promise) {

	p = makepromise()
	go func() {
		q1 := ZipGen2(a, b)
		test_predicates2(f, q1, p, predicates...)
		p.Close()
	}()

	return

}

func test_predicates2(f FuncAny2, promTuple *Promise, outchan *Promise, predicates ...FuncBool2) {
	var tuple *Tuple2
	var wg WaitGroup

	for data := range promTuple.Q() {
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
			wg.Add(AsyncAnyN(func(n ...AnyVal) AnyVal {
				outchan.send(f(n[0], n[1]))
				return true
			}, tuple.A, tuple.B))
		}
	}

	wg.Wait()
}
