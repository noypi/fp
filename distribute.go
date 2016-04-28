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
			go func() {
				var msg = new(qMsg)
				msg.a, msg.err = worker(a)
				q.q <- msg
				wg.Done()
			}()
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
