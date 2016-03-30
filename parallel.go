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

	p = LazyInAsync1(func(a AnyVal) (AnyVal, error) {
		return bf(a)
	}, q1, chanlen...)

	return

}
func RangeList(f Func2, list AnyVal, chanlen ...int) (p *Promise) {
	return rangeList(f, list, false, chanlen...)
}

func RangeListAsync(f Func2, list AnyVal, chanlen ...int) (p *Promise) {
	return rangeList(f, list, true, chanlen...)
}

func rangeList(f Func2, list AnyVal, async bool, chanlen ...int) (p *Promise) {
	v, ok := list.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(list)
	}

	n := v.Len()
	var wg sync.WaitGroup
	wg.Add(n)
	fnexec := func(index int, ch *Promise) {
		ret, err := f(v.Index(index).Interface(), index)
		msg := new(qMsg)
		msg.a = ret
		msg.err = err
		ch.send(msg)
		wg.Done()
	}

	p = makepromise(chanlen...)
	go func() {
		for i := 0; i < n; i++ {
			if async {
				fnexec(i, p)
			} else {
				fnexec(i, p)
			}
		}
	}()

	go func() {
		wg.Wait()
		p.close()
	}()

	return p
}

// calls Func2 as func(value, key)
func RangeDict(f Func2, dict AnyVal, chanlen ...int) (p *Promise) {
	return rangeDict(f, dict, false, chanlen...)
}
func RangeDictAsync(f Func2, dict AnyVal, chanlen ...int) (p *Promise) {
	return rangeDict(f, dict, true, chanlen...)
}
func rangeDict(f Func2, dict AnyVal, async bool, chanlen ...int) (p *Promise) {
	v, ok := dict.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(dict)
	}

	n := v.Len()
	p = makepromise(chanlen...)
	var wg sync.WaitGroup
	wg.Add(n)

	fnexec := func(vk reflect.Value, ch *Promise) {
		ret, err := f(v.MapIndex(vk).Interface(), vk.Interface())
		msg := new(qMsg)
		msg.a = ret
		msg.err = err
		ch.send(msg)
		wg.Done()
	}

	go func() {
		for _, vk := range v.MapKeys() {
			if async {
				go fnexec(vk, p)
			} else {
				fnexec(vk, p)
			}
		}
	}()

	go func() { wg.Wait(); p.close() }()

	return p
}
