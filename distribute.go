package fp

import (
	"sync"
)

type _distributesrc struct {
	srcq    qChan
	srcqAny chan interface{}
	srcl0   LazyFn
}

func distributework(src _distributesrc, q *Promise, worker func(interface{}) (interface{}, error), nProcessor uint) {
	var wg sync.WaitGroup
	var wCh = make(chan struct{}, nProcessor)
	var blank = struct{}{}

	doa := func(a1 interface{}) {
		switch v := a1.(type) {
		case *qMsg:
			if nil == v.err {
				v.a, v.err = worker(v.a)
			}
			<-wCh
			q.send(v)

		default:
			var msg = new(qMsg)
			msg.a, msg.err = worker(a1)
			<-wCh
			q.send(msg)

		}
		wg.Done()
	}

	if nil != src.srcqAny {
		for a := range src.srcqAny {
			wg.Add(1)
			wCh <- blank
			go doa(a)
		}
	} else if nil != src.srcq {
		for a := range src.srcq {
			wg.Add(1)
			wCh <- blank
			go doa(a)
		}
	} else if nil != src.srcl0 {
		for {
			a := src.srcl0()
			a1 := <-a.q
			if nil == a1 {
				break
			} else {
				wg.Add(1)
				wCh <- blank
				go doa(a1)
			}
		}
	}

	close(wCh)
	wg.Wait()
	q.close()
}

func DistributeWork(src *Promise, worker func(interface{}) (interface{}, error), nProcessor uint) (q *Promise) {
	q = makepromise()
	go distributework(_distributesrc{srcq: src.q}, q, worker, nProcessor)
	return
}

func DistributeWorkCh(src chan interface{}, worker func(interface{}) (interface{}, error), nProcessor uint) (q *Promise) {
	q = makepromise()
	go distributework(_distributesrc{srcqAny: src}, q, worker, nProcessor)
	return q
}

func DistributeWorkL0(ql LazyFn, worker func(interface{}) (interface{}, error), nProcessor uint) (q *Promise) {
	q = makepromise()
	go distributework(_distributesrc{srcl0: ql}, q, worker, nProcessor)
	return
}
