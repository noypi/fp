package fp

import (
	"time"
)

func TickUntil(f FuncTick2, d time.Duration, predicates ...FuncTickBool2) (p *Promise) {
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
			p.send(f(now, prevtime))
			prevtime = now
		}
		p.close()
	}()

	return
}

func TickWhile(f FuncTick2, d time.Duration, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()

	tickr := time.Tick(d)
	if nil == tickr {
		p.close()
		return
	}

	var prevtime time.Time
	return ProduceWhile(func() (ret AnyVal) {
		now := <-tickr
		ret = f(now, prevtime)
		prevtime = now
		return
	}, predicates...)

}
