package fp

import (
	"sync"
)

func DistributeWork(src *Promise, worker func(interface{}) (interface{}, error), nProcessor uint) (q *Promise) {
	q = makepromise()
	go func() {
		var wg sync.WaitGroup
		var wCh = make(chan struct{}, nProcessor)
		var blank = struct{}{}
		for a := range src.q {
			wg.Add(1)
			wCh <- blank
			go func(a1 interface{}) {
				switch v := a1.(type) {
				case *qMsg:
					if nil == v.err {
						v.a, v.err = worker(v.a)
					}
					q.q <- v
				default:
					var msg = new(qMsg)
					msg.a, msg.err = worker(a1)
					q.q <- msg

				}
				wg.Done()
				<-wCh
			}(a)
		}
		close(wCh)
		wg.Wait()
		q.close()
		for _ = range wCh {

		}
	}()
	return
}

func DistributeWorkCh(src chan interface{}, worker func(interface{}) (interface{}, error), nProcessor uint) (q *Promise) {
	return DistributeWork(PipeChan(src), worker, nProcessor)
}
