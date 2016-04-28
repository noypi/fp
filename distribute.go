package fp

import (
	"sync"
)

func DistributeWork(src *Promise, worker func(interface{}) (interface{}, error), nProcessor uint) (q *Promise) {
	panic("todo")
	return
}

func DistributeWorkCh(src chan interface{}, worker func(interface{}) (interface{}, error), nProcessor uint) (q *Promise) {
	q = makepromise()
	go func() {
		var wg sync.WaitGroup
		var i uint
		for a := range src {
			wg.Add(1)
			go func(a1 interface{}) {
				var msg = new(qMsg)
				msg.a, msg.err = worker(a1)
				q.q <- msg
				wg.Done()
			}(a)
			i++
			if 0 >= (nProcessor - i) {
				wg.Wait()
				i = 0
			}
		}
		wg.Wait()
		q.close()
	}()
	return
}
