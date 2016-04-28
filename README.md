[![Build Status](https://travis-ci.org/noypi/fp.svg?branch=master)](https://travis-ci.org/noypi/fp)
[![GoDoc](https://godoc.org/github.com/noypi/fp?status.png)](http://godoc.org/github.com/noypi/fp)
[![codebeat badge](https://codebeat.co/badges/a7fbc443-04aa-47c9-8b78-20b1395a9dc1)](https://codebeat.co/projects/github-com-noypi-fp)

### Golang promise, futures, lazy, list comprehension...etc...

#### Examples

##### implementing q

- inspired by angularjs $q
- chainable, and async

```go
	q := Future(func() (interface{}, error) {
		// do some work
		return "resolved", errors.New("failed")
	})

	q = q.Then(func(a interface{}) (interface{}, error) {
		// on resolved
		return "resolved", errors.New("failed")
	}, func(a interface{}) (interface{}, error) {
		// on error
		return "resolved", errors.New("failed")
	})

```

##### distribute work (TODO: still in progress)
```go
	src := make(chan interface{}, 100)
	go func() {
		for i := 0; i < 100; i++ {
			src <- i
		}
		close(src)
	}()

	work := func(a interface{}) (out interface{}, err error) {
		out = a.(int) + 1
		return
	}

	q := DistributeWorkCh(src, work, uint(runtime.NumCPU()))

	var total int32
	q = q.Then(func(a interface{}) (out interface{}, err error) {
		atomic.AddInt32(&total, int32(a.(int)))
		return
	})

	// wait
	Flush(q)
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
	func (this Resource) Get(n ...interface{}) (p *Promise) {
		return r.fget(n...)
	}
	func (this Resource) Put(n ...interface{}) (p *Promise) {
		return r.fput(n...)
	}
```

##### other examples

```go
	as := []int{26, 27, 29, 0, 1, 2, 26, 27, 29, 0, 1, 2, 26, 27, 29, 0, 1, 2}
	// range will concurrently execute each
	qLazy := Range(func(a, i interface{}) (ret interface{}, err error) {
		ret = &Tuple2{
			A: a,
			B: i,
		}
		return
	}, as)

	q1 := LazyInAsync1(func(x interface{}) (ret interface{}, err error) {
		ret = fb(x.(*Tuple2).A.(int))
		return
	}, qLazy)


```

See examples for more.



### shoutout

Thanks minux for reviewing... =)

#### Other similar projects
- github.com/fanliao/go-promise

