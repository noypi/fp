package fp

import (
	"reflect"
)

// !!! not tested
func Zip2(alist, blist AnyVal) (p PromiseChan) {

	p = make(PromiseChan, 1)

	av := reflect.ValueOf(alist)
	bv := reflect.ValueOf(blist)

	go func() {
		for i := 0; i < av.Len(); i++ {
			if i < bv.Len() {
				p <- &Tuple2{
					A: av.Index(i).Interface(),
					B: bv.Index(i).Interface(),
				}
			} else {
				break
			}
		}
		close(p)
	}()

	return
}

// !!! not tested
// this Zip does not accept nil as a value to tuples... need to improve?
// closing a and b stops the goroutine
func ZipGen2(a, b LazyInChan) (p PromiseChan) {

	p = make(PromiseChan, 0)

	go func() {

		for {
			var x, y AnyVal
			select {
			case x = <-a:
				for y = range b {
					// breaks when b is closed
				}
			case y = <-b:
				for x = range a {
					// breaks when a is closed
				}
			}

			if nil == x || nil == y {
				break
			}

			p <- &Tuple2{
				A: x,
				B: y,
			}

		}
		close(p)
	}()

	return
}
