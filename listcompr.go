package fp

import (
	"reflect"
)

//!!! not tested
// list comprehension
func listCompr(f FuncAny1, alist interface{}, async bool, predicates ...FuncBool1) (p *Promise) {
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
		p.close()
	}()

	return

}

func ListCompr(f FuncAny1, alist interface{}, predicates ...FuncBool1) (p *Promise) {
	return listCompr(f, alist, false, predicates...)
}

func ListComprAsync(f FuncAny1, alist interface{}, predicates ...FuncBool1) (p *Promise) {
	return listCompr(f, alist, true, predicates...)
}

//!!! not tested
// list comprehension
func listComprGen(f FuncAny1, in *Promise, async bool, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()

	go func() {
		var wg WaitGroup
		for a := range in.q {
			p1 := test_predicates1(f, a.a, p, async, predicates...)
			if async && nil != p1 {
				wg.Add(p1)
			}
		}

		if async {
			wg.Wait()
		}
		p.close()
	}()

	return

}

func ListComprGen(f FuncAny1, in *Promise, predicates ...FuncBool1) (p *Promise) {
	return listComprGen(f, in, false, predicates...)
}

func ListComprGenAsync(f FuncAny1, in *Promise, predicates ...FuncBool1) (p *Promise) {
	return listComprGen(f, in, true, predicates...)
}

func test_predicates1(f FuncAny1, a interface{}, outchan *Promise, async bool, predicates ...FuncBool1) (p *Promise) {
	if are_all_true1(a, predicates...) {
		if async {
			p = Async(func() {
				msg := new(qMsg)
				msg.a = f(a)
				outchan.send(msg)
			})
		} else {
			msg := new(qMsg)
			msg.a = f(a)
			outchan.send(msg)
		}
	}

	return
}

//!!! not tested
// list comprehension, 2 lists
func listCompr2(f FuncAny2, alist, blist interface{}, async bool, predicates ...FuncBool2) (p *Promise) {

	p = makepromise()
	go func() {
		q1 := Zip2(alist, blist)
		test_predicates2(f, q1, p, async, predicates...)
		p.close()
	}()

	return

}

func ListCompr2(f FuncAny2, alist, blist interface{}, predicates ...FuncBool2) (p *Promise) {
	return listCompr2(f, alist, blist, false, predicates...)
}

func ListComprAsync2(f FuncAny2, alist, blist interface{}, predicates ...FuncBool2) (p *Promise) {
	return listCompr2(f, alist, blist, true, predicates...)
}

//!!! not tested
// list comprehension, 2 lists
func listComprGen2(f FuncAny2, a, b *Promise, async bool, predicates ...FuncBool2) (p *Promise) {
	p = makepromise()
	go func() {
		q1 := ZipGen2(a, b)
		test_predicates2(f, q1, p, async, predicates...)
		p.close()
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

	for data := range promTuple.q {
		tuple = data.a.(*Tuple2)
		if are_all_true2(tuple.A, tuple.B, predicates...) {
			if async {
				wg.Add(AsyncAnyN(func(n ...interface{}) interface{} {
					msg := new(qMsg)
					msg.a = f(n[0], n[1])
					outchan.send(msg)
					return true
				}, tuple.A, tuple.B))
			} else {
				msg := new(qMsg)
				msg.a = f(tuple.A, tuple.B)
				outchan.send(msg)
			}
		}
	}

	if async {
		wg.Wait()
	}
}
