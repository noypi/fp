package fp

import (
	"reflect"
)

func ReduceParams(fn interface{}, as interface{}) *Promise {
	return Future(func() (res interface{}, err error) {
		defer func() {
			err = recover2err(recover(), ErrInvalidParam)
		}()

		vfn := reflect.ValueOf(fn)
		params := make([]reflect.Value, 2)
		vas := reflect.ValueOf(as)
		prev := vas.Index(0)
		for i := 1; i < vas.Len(); i++ {
			params[0] = prev
			params[1] = vas.Index(i)
			prev = (vfn.Call(params))[0]
		}
		res = prev.Interface()
		return
	})
}

func ReduceFuncs(initparam interface{}, fns interface{}) *Promise {
	return Future(func() (res interface{}, err error) {
		defer func() {
			err = recover2err(recover(), ErrInvalidParam)
		}()
		vfns := reflect.ValueOf(fns)
		prevfn := vfns.Index(0)
		params := make([]reflect.Value, 1)
		params[0] = reflect.ValueOf(initparam)
		params[0] = prevfn.Call(params)[0]

		for i := 1; i < vfns.Len(); i++ {
			params[0] = vfns.Index(i).Call(params)[0]
		}

		res = params[0].Interface()
		return
	})

}
