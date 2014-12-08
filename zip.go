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
				p.send(&Tuple2{
					A: av.Index(i).Interface(),
					B: bv.Index(i).Interface(),
				})
			} else {
				break
			}
		}
		p.Close()
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
					Dir: reflect.SelectRecv,
					Chan: a.vq,
				},
				reflect.SelectCase{
					Dir: reflect.SelectRecv,
					Chan: b.vq,
				},
			})
			
			if !ok {
				break
			}
			
			if 0==chosen {
				x = xyi.Interface()
				if ok {
					y, ok = b.Recv()
				}
			} else {
				y = xyi.Interface()
				if ok {
					x, ok = a.Recv()
				}
			}
			
			if !ok {
				break
			}

			p.send(&Tuple2{
				A: x,
				B: y,
			})

		}
		p.Close()
	}()

	return
}
