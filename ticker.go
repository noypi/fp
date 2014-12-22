package fp

import (
	"time"
)

func TickUntil(f FuncTickAny, d time.Duration, predicates ...FuncTickBool) (p *Promise) {
	return tickUntil(f, d, false, predicates...)
}

func TickUntilMute(f FuncTickVoid, d time.Duration, predicates ...FuncTickBool) (p *Promise) {
	predicates = append([]FuncTickBool{func(a, b time.Time) bool {
		f(a, b)
		return true
	}}, predicates...)
	return tickUntil(nil, d, true, predicates...)
}

func tickUntil(f FuncTickAny, d time.Duration, mute bool, predicates ...FuncTickBool) (p *Promise) {
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
				if !mute {
					p.send(a)
				}
			}
			prevtime = now
		}
		p.close()
	}()

	return
}

func TickWhile(f FuncTickAny, d time.Duration, predicates ...FuncBool1) (p *Promise) {
	return tickWhile(f, d, false, predicates...)
}

func TickWhileMute(f FuncTickVoid, d time.Duration, predicates ...FuncBool1) (p *Promise) {
	return tickWhile(func(a, b time.Time) (ret AnyVal) {
		f(a, b)
		return
	}, d, true, predicates...)
}

func tickWhile(f FuncTickAny, d time.Duration, mute bool, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()

	tickr := time.Tick(d)
	if nil == tickr {
		p.close()
		return
	}

	var prevtime time.Time
	return produceWhile(func() (ret AnyVal) {
		now := <-tickr
		if nil != f {
			ret = f(now, prevtime)
		}
		prevtime = now
		return
	}, mute, predicates...)

}
