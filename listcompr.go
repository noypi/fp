package fp

import (
	"reflect"
)

//!!! not tested
// list comprehension
func listCompr(f FuncAny1, alist AnyVal, async bool, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()
	va := reflect.ValueOf(alist)

	go func() {
		var wg WaitGroup
		for i := 0; i < va.Len(); i++ {
			a := va.Index(i).Interface()
			p1 := test_predicates1(f, a, p, async, predicates...)
			if async && nil != p1 {
				wg.Add(p1)
			}
		}

		if async {
			wg.Wait()
		}
		p.Close()
	}()

	return

}

func ListCompr(f FuncAny1, alist AnyVal, predicates ...FuncBool1) (p *Promise) {
	return listCompr(f, alist, false, predicates...)
}

func ListComprAsync(f FuncAny1, alist AnyVal, predicates ...FuncBool1) (p *Promise) {
	return listCompr(f, alist, true, predicates...)
}

//!!! not tested
// list comprehension
func listComprGen(f FuncAny1, in *Promise, async bool, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()

	go func() {
		var wg WaitGroup
		for a := range in.Q() {
			p1 := test_predicates1(f, a, p, async, predicates...)
			if async && nil != p1 {
				wg.Add(p1)
			}
		}

		if async {
			wg.Wait()
		}
		p.Close()
	}()

	return

}

func ListComprGen(f FuncAny1, in *Promise, predicates ...FuncBool1) (p *Promise) {
	return listComprGen(f, in, false, predicates...)
}

func ListComprGenAsync(f FuncAny1, in *Promise, predicates ...FuncBool1) (p *Promise) {
	return listComprGen(f, in, true, predicates...)
}

func test_predicates1(f FuncAny1, a AnyVal, outchan *Promise, async bool, predicates ...FuncBool1) (p *Promise) {
	trueCnt := 0
	for _, pred := range predicates {
		if pred(a) {
			trueCnt++
		} else {
			break
		}
	}
	if len(predicates) == trueCnt {

		if async {
			p = Async(func() {
				outchan.send(f(a))
			})
		} else {
			outchan.send(f(a))
		}
	}

	return
}

//!!! not tested
// list comprehension, 2 lists
func listCompr2(f FuncAny2, alist, blist AnyVal, async bool, predicates ...FuncBool2) (p *Promise) {

	p = makepromise()
	go func() {
		q1 := Zip2(alist, blist)
		test_predicates2(f, q1, p, async, predicates...)
		p.Close()
	}()

	return

}

func ListCompr2(f FuncAny2, alist, blist AnyVal, predicates ...FuncBool2) (p *Promise) {
	return listCompr2(f, alist, blist, false, predicates...)
}

func ListComprAsync2(f FuncAny2, alist, blist AnyVal, predicates ...FuncBool2) (p *Promise) {
	return listCompr2(f, alist, blist, true, predicates...)
}

//!!! not tested
// list comprehension, 2 lists
func listComprGen2(f FuncAny2, a, b *Promise, async bool, predicates ...FuncBool2) (p *Promise) {
	p = makepromise()
	go func() {
		q1 := ZipGen2(a, b)
		test_predicates2(f, q1, p, async, predicates...)
		p.Close()
	}()

	return
}

func ListComprGen2(f FuncAny2, a, b *Promise, predicates ...FuncBool2) (p *Promise) {
	return listComprGen2(f, a, b, false, predicates...)
}

func ListComprGenAsync2(f FuncAny2, a, b *Promise, predicates ...FuncBool2) (p *Promise) {
	return listComprGen2(f, a, b, true, predicates...)
}

func test_predicates2(f FuncAny2, promTuple *Promise, outchan *Promise, async bool, predicates ...FuncBool2) {
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
			if async {
				wg.Add(AsyncAnyN(func(n ...AnyVal) AnyVal {
					outchan.send(f(n[0], n[1]))
					return true
				}, tuple.A, tuple.B))
			} else {
				outchan.send(f(tuple.A, tuple.B))
			}
		}
	}

	if async {
		wg.Wait()
	}
}
