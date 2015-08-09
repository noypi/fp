package fp

import (
	"reflect"
	"sync"
)

func Range(f Func2, listOrMap AnyVal, chanlen ...int) (p *Promise) {
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

// !!! not yet tested
// execute af() for each element; and for each valid element run bf() in parallel
// p receives the result of bf()
func ParallelLoop(af Func2, bf Func1, aListOrMap AnyVal, chanlen ...int) (p *Promise) {

	q1 := Range(af, aListOrMap)

	p = LazyInAsync1(func(a AnyVal) (ret AnyVal, skip bool) {
		return bf(a)
	}, q1, chanlen...)

	return

}

func RangeList(f Func2, list AnyVal, chanlen ...int) (p *Promise) {
	v, ok := list.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(list)
	}

	n := v.Len()
	p = makepromise(chanlen...)
	qtmp := make(chan bool, cap(p.Q()))

	wg := new(sync.WaitGroup)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(index int, ch *Promise) {
			qtmp <- true
			if ret, skip := f(v.Index(index).Interface(), index); !skip {
				ch.send(ret)
			}
			wg.Done()
			<-qtmp
		}(i, p)
	}

	go func() { wg.Wait(); close(qtmp); p.close() }()

	return p
}

// calls Func2 as func(value, key)
func RangeDict(f Func2, dict AnyVal, chanlen ...int) (p *Promise) {
	v, ok := dict.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(dict)
	}

	n := v.Len()
	p = makepromise(chanlen...)
	qtmp := make(chan bool, cap(p.Q()))
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for _, vk := range v.MapKeys() {
		go func(vk reflect.Value, ch *Promise) {
			qtmp <- true
			if ret, skip := f(v.MapIndex(vk).Interface(), vk.Interface()); !skip {
				ch.send(ret)
			}
			wg.Done()
			<-qtmp
		}(vk, p)
	}

	go func() { wg.Wait(); p.close() }()

	return p
}
