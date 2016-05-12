package fp

import (
	"sync"
)

func distributework(srcq qChan, srcqAny chan interface{}, q *Promise, worker func(interface{}) (interface{}, error), nProcessor uint) {
	var wg sync.WaitGroup
	var wCh = make(chan struct{}, nProcessor)
	var blank = struct{}{}

	doa := func(a interface{}) {
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

	if nil == srcq {
		for a := range srcqAny {
			doa(a)
		}
	} else {
		for a := range srcq {
			doa(a)
		}
	}

	close(wCh)
	wg.Wait()
	q.close()
}

func DistributeWork(src *Promise, worker func(interface{}) (interface{}, error), nProcessor uint) (q *Promise) {
	q = makepromise()
	go distributework(src.q, nil, q, worker, nProcessor)
	return
}

func DistributeWorkCh(src chan interface{}, worker func(interface{}) (interface{}, error), nProcessor uint) (q *Promise) {
	q = makepromise()
	go distributework(nil, src, q, worker, nProcessor)
	return q
}
