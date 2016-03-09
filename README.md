[![Build Status](https://travis-ci.org/noypi/fp.svg?branch=master)](https://travis-ci.org/noypi/fp)
[![GoDoc](https://godoc.org/github.com/noypi/fp?status.png)](http://godoc.org/github.com/noypi/fp)

### Golang futures, lazy, list comprehension...etc...

things that help me simplify my code.

### Other projects
- github.com/fanliao/go-promise

#### Examples

##### implementing q

- inspired by angularjs $q
- chainable, and async
- the difference is that a Q is not a promise, use Future or Lazy instead

```go
	q := Q(func(a AnyVal) AnyVal {
		return "from success1"
	})

	q.OnSuccess(func(a AnyVal) AnyVal {
		return "from success 2"
	})

	q.OnDone(func(a AnyVal) AnyVal {
		return "from done 1"
	})

	qres, qsig := q.Call(func(s QSignal) {
		s.Resolve("resolved!")
	})

	fmt.Println("result=", <-qres.Q())
	fmt.Println(qsig.HaveSucceeded())
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
