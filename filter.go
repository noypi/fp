package fp

//!! not yet tested
func FilterGen(q *Promise, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()
	go func() {
		for a := range q.q {
			if are_all_true1(a, predicates...) {
				q.send(a)
			}
		}
		p.close()
	}()
	return
}
