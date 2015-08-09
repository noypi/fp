### Golang futures, lazy, list comprehension...etc...

things that help me simplify my code.

#### Examples

##### implementing a message loop

```go
	func MessageLoop() (ares, aerr, anotify AnyVal) {
		// read from channel from C's callback'
		// or wait for incoming message from a websocket =)
		anotify = some.Pool().(*SomeEvent) // set anotify to receive notification messages

		// set aerr to stop the loop and notifies Qerror
		aerr = "some error"

		// if anotify is nil and aerr is nil, the loop stops with a notification to Qresult
		ares = "any value including nil"

		return
	}

	func OnNotify(val *SomeEvent) {
		// do something
	}

	func setupLoop() (wg *WaitGroup){
		chain := Q(MessageLoop)

		wg := new(WaitGroup)

		// receive notification messages
		q := Async(LazyInParamsMute(n ...AnyVal){

			OnNotify(n[0].(*SomeEvent))

		}, chain.Qnotify)// can add more parameters, example sender

		wg.Add(q)

		// can also chain handlers
		chain1 := chain.Bind(...)
		//-- or
		chain1b := chain.BindMute(...) // if Qnotify, Qerror, and Qresult is not used
		
		return
	}

	func main() {
		wg := setupLoop()
		wg.Wait()
	}
```

##### example of implementing a resource using LazyN

```go
	type Resource struct {
		fget FuncAnyN
		fput FuncAnyN
	}
	func NewResource(fget, fput FuncAnyN) (r *Resource) {
		r = new(Resource)
		r.fget = LazyN(fget)
		r.fput = LazyN(fput)
		return
	}
	func (this Resource) Get(n ...AnyVal) (p *Promise) {
		return r.fget(n...)
	}
	func (this Resource) Put(n ...AnyVal) (p *Promise) {
		return r.fput(n...)
	}
```

##### other examples

```go
	as := []int{26, 27, 29, 0, 1, 2, 26, 27, 29, 0, 1, 2, 26, 27, 29, 0, 1, 2}
	// range will concurrently execute each
	qLazy := Range(func(a, i AnyVal) (ret AnyVal, skip bool) {
		ret = &Tuple2{
			A: a,
			B: i,
		}
		return
	}, as)

	q1 := LazyInAsync1(func(x AnyVal) (ret AnyVal, skip bool) {
		ret = fb(x.(*Tuple2).A.(int))
		return
	}, qLazy)

	// print results
	for a := range q1.Q() {
		log.Printf("ret=%d\n", a)
	}
```

See examples for more.



### shoutout

Thanks minux for reviewing... =)
