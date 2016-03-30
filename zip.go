package fp

import (
	"reflect"
)

// !!! not tested
func Zip2(alist, blist AnyVal) (p *Promise) {

	p = makepromise()

	av := reflect.ValueOf(alist)
	bv := reflect.ValueOf(blist)

	go func() {
		for i := 0; i < av.Len(); i++ {
			if i < bv.Len() {
				tuple := new(Tuple2)
				tuple.A = av.Index(i).Interface()
				tuple.B = bv.Index(i).Interface()
				msg := new(qMsg)
				msg.a = tuple
				p.send(msg)
			} else {
				break
			}
		}
		p.close()
	}()

	return
}

// !!! not tested
func ZipGen2(a, b *Promise) (p *Promise) {

	p = makepromise()

	go func() {

		for {
			var x, y AnyVal
			var ok bool

			chosen, xyi, ok := reflect.Select([]reflect.SelectCase{
				reflect.SelectCase{
					Dir:  reflect.SelectRecv,
					Chan: a.vq,
				},
				reflect.SelectCase{
					Dir:  reflect.SelectRecv,
					Chan: b.vq,
				},
			})

			if !ok {
				break
			}

			if 0 == chosen {
				x = xyi.Interface()
				if ok {
					y = (<-b.q).a
				}
			} else {
				y = xyi.Interface()
				if ok {
					x = (<-a.q).a
				}
			}

			if !ok {
				break
			}

			tuple := new(Tuple2)
			tuple.A = x
			tuple.B = y
			msg := new(qMsg)
			msg.a = tuple

			p.send(msg)

		}
		p.close()
	}()

	return
}

//!!! not yet tested
func ZipWith2(f FuncAny2, alist, blist AnyVal) (p *Promise) {
	q := Zip2(alist, blist)
	p = zipwith(f, q)
	return
}

//!!! not yet tested
func ZipGenWith2(f FuncAny2, q1, q2 *Promise) (p *Promise) {
	q := ZipGen2(q1, q2)
	p = zipwith(f, q)
	return
}

func zipwith(f FuncAny2, in *Promise) (p *Promise) {
	p = makepromise()
	for xy := range in.q {
		tuple := xy.a.(*Tuple2)
		msg := new(qMsg)
		msg.a = f(tuple.A, tuple.B)
		p.send(msg)
	}
	return
}
