### Golang futures, lazy, list comprehension...etc...

things that help me simplify my code.

#### Examples

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
