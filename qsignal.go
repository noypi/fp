package fp

type QSignal interface {
	Reject(AnyVal) *Promise
	Resolve(AnyVal) *Promise
	Notify(AnyVal) *Promise
	HaveSucceeded() bool
	IsRejected() bool
}

type qsignal struct {
	q         *_Q
	qdonechan ChanAny
	succeeded bool
	rejected  bool
}

func (this qsignal) HaveSucceeded() bool {
	return this.succeeded
}

func (this qsignal) IsRejected() bool {
	return this.rejected
}

func (this *qsignal) qdone() *Promise {
	this.qdonechan = make(ChanAny, 1)
	return Future(func() (AnyVal, error) {
		return <-this.qdonechan, nil
	})
}

func (this *qsignal) execEach(as []LazyFn1, param AnyVal) (out *Promise) {
	if this.succeeded != this.rejected {
		// Done already called
		return
	}
	if 0 == len(as) {
		return
	}

	var bs []LazyFn1
	{ //lock
		this.q.syncq.Lock()
		bs = as[0:]
		this.q.syncq.Unlock()
	}

	for _, fn := range bs {
		param = fn(param)
	}

	if 0 < len(as) {
		out = param.(*Promise)
	}
	return
}

func (this *qsignal) Reject(result AnyVal) (out *Promise) {
	out = this.execEach(this.q.failure, result)
	var param AnyVal = out
	if nil == out {
		param = result
	}
	this.Done(param, false)
	return nil
}

func (this *qsignal) Resolve(result AnyVal) (out *Promise) {
	out = this.execEach(this.q.success, result)
	var param AnyVal = out
	if nil == out {
		param = result
	}
	this.Done(param, true)
	return nil
}

func (this *qsignal) Notify(result AnyVal) (out *Promise) {
	out = this.execEach(this.q.notify, result)
	return nil
}

func (this *qsignal) Done(result AnyVal, succeed bool) (out *Promise) {
	out = this.execEach(append(this.q.done, wrapcb(this.lastdone(succeed))), result)
	return nil
}

func (this *qsignal) lastdone(succeed bool) func(AnyVal) AnyVal {
	return func(param AnyVal) AnyVal {
		if nil != this.qdonechan {
			this.qdonechan <- param
			close(this.qdonechan)
		}

		this.succeeded = succeed
		this.rejected = !succeed

		return param
	}
}
