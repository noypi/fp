### Golang futures, lazy, list comprehension...etc...

things that help me simplify my code.

#### Examples
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
