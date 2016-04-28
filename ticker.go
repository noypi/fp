package fp

import (
	"time"
)

func TickUntil(f FuncTickAny, d time.Duration, predicates ...FuncTickBool) (p *Promise) {
	return tickUntil(f, d, predicates...)
}

func tickUntil(f FuncTickAny, d time.Duration, predicates ...FuncTickBool) (p *Promise) {
	p = makepromise()

	tickr := time.Tick(d)
	if nil == tickr {
		p.close()
		return
	}

	go func() {
		var prevtime time.Time
		for now := range tickr {
			if !are_all_tick_true2(now, prevtime, predicates...) {
				break
			}
			if nil != f {
				a := f(now, prevtime)
				msg := new(qMsg)
				msg.a = a
				p.send(msg)
			}
			prevtime = now
		}
		p.close()
	}()

	return
}

func TickWhile(f FuncTickAny, d time.Duration, predicates ...FuncBool1) (p *Promise) {
	return tickWhile(f, d, predicates...)
}

func tickWhile(f FuncTickAny, d time.Duration, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()

	tickr := time.Tick(d)
	if nil == tickr {
		p.close()
		return
	}

	var prevtime time.Time
	return produceWhile(func() (ret interface{}) {
		now := <-tickr
		if nil != f {
			ret = f(now, prevtime)
		}
		prevtime = now
		return
	}, predicates...)

}
