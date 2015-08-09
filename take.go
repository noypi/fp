package fp

//see ProduceWhile
func TakeWhileGen(q *Promise, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()
	go func() {
		for a := range q.q {
			if !are_all_true1(a, predicates...) {
				break
			} else {
				p.send(a)
			}
		}
		p.close()
		Flush(q) // you might like ProduceWhile
	}()

	return
}
