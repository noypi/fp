package fp

import (
	"sync"
)

func DistributeWork(src *Promise, worker func(AnyVal), nProcessor uint) (q *Promise) {
	q = makepromise()
	go func() {
		var wg sync.WaitGroup
		var i uint
		for a := range q.q {
			go func() { worker(a); wg.Done() }()
			i++
			if 0 == (nProcessor - i) {
				wg.Wait()
				i = 0

			}
		}
	}()
	return
}
