package fp

import (
	"sync"
)

type _Q struct {
	success []LazyFn1
	failure []LazyFn1
	notify  []LazyFn1
	done    []LazyFn1
	syncq   sync.Mutex
}

//Q(fn, successCb, failureCb, doneCb, notifyCb)
func Q(fns ...FuncAny1) (q *_Q) {
	q = new(_Q)
	if 1 <= len(fns) {
		q.success = append(q.success, wrapcb(fns[0]))
	}
	if 2 <= len(fns) {
		q.failure = append(q.success, wrapcb(fns[1]))
	}
	if 3 <= len(fns) {
		q.done = append(q.success, wrapcb(fns[2]))
	}
	if 4 <= len(fns) {
		q.notify = append(q.success, wrapcb(fns[3]))
	}

	return
}

func AllSettled(as ...*Promise) *Promise {
	bs := make([]AnyVal, len(as))
	return Future(func() (AnyVal, error) {
		for i := 0; i < len(bs); i++ {
			bs[i] = <-as[i].Q()
		}
		return bs, nil
	})
}

func (this *_Q) Call(fn func(s QSignal)) (qp *Promise, qsig QSignal) {
	if nil == fn {
		return
	}
	s := &qsignal{q: this}
	qsig = s
	qp = s.qdone()
	go fn(s)
	return
}

func (this *_Q) CallAll(fns ...func(s QSignal)) (qps []*Promise) {
	if 0 == len(fns) {
		return
	}

	qps = make([]*Promise, len(fns))
	for i, fn := range fns {
		s := &qsignal{q: this}
		qps[i] = s.qdone()
		go fn(s)
	}

	return
}

func (this *_Q) Then(successCb, failureCb, doneCb, notifyCb func(AnyVal) AnyVal) (q *_Q) {
	this.syncq.Lock()
	defer this.syncq.Unlock()

	this.success = append(this.success, wrapcb(successCb))
	q.failure = append(this.failure, wrapcb(failureCb))
	q.done = append(this.done, wrapcb(doneCb))
	q.notify = append(this.notify, wrapcb(notifyCb))

	return
}

func (this *_Q) OnSuccess(successCb func(AnyVal) AnyVal) (q *_Q) {
	this.syncq.Lock()
	this.success = append(this.success, wrapcb(successCb))
	this.syncq.Unlock()
	return
}

func (this *_Q) OnFailure(failureCb func(AnyVal) AnyVal) (q *_Q) {
	this.syncq.Lock()
	this.failure = append(this.failure, wrapcb(failureCb))
	this.syncq.Unlock()
	return
}

func (this *_Q) OnDone(doneCb func(AnyVal) AnyVal) (q *_Q) {
	this.syncq.Lock()
	this.done = append(this.done, wrapcb(doneCb))
	this.syncq.Unlock()
	return
}

func (this *_Q) OnNotify(notifyCb func(AnyVal) AnyVal) (q *_Q) {
	this.syncq.Lock()
	this.notify = append(this.notify, wrapcb(notifyCb))
	this.syncq.Unlock()
	return
}

func wrapcb(fn FuncAny1) LazyFn1 {
	return Lazy1(func(in AnyVal) AnyVal {
		switch v := in.(type) {
		case *Promise:
			res := <-v.Q()
			return fn(res)
		default:
			return fn(in)
		}
	})
}
