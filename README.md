[![Build Status](https://travis-ci.org/noypi/fp.svg?branch=master)](https://travis-ci.org/noypi/fp)
[![GoDoc](https://godoc.org/github.com/noypi/fp?status.png)](http://godoc.org/github.com/noypi/fp)

### Golang promise, futures, lazy, list comprehension...etc...

#### Examples

##### implementing q

- inspired by angularjs $q
- chainable, and async

```go
	q := Future(func() (AnyVal, error) {
		// do some work
		return "resolved", errors.New("failed")
	})

	q = q.Then(func(a AnyVal) (AnyVal, error) {
		// on resolved
		return "resolved", errors.New("failed")
	}, func(a AnyVal) (AnyVal, error) {
		// on error
		return "resolved", errors.New("failed")
	})

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
	qLazy := Range(func(a, i AnyVal) (ret AnyVal, err error) {
		ret = &Tuple2{
			A: a,
			B: i,
		}
		return
	}, as)

	q1 := LazyInAsync1(func(x AnyVal) (ret AnyVal, err error) {
		ret = fb(x.(*Tuple2).A.(int))
		return
	}, qLazy)


```

See examples for more.



### shoutout

Thanks minux for reviewing... =)

#### Other similar projects
- github.com/fanliao/go-promise

