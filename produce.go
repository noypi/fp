package fp

//!! not yet tested
func ProduceWhile(f FuncAny0, predicates ...FuncBool1) (p *Promise) {
	return produceWhile(f, predicates...)
}

func produceWhile(f FuncAny0, predicates ...FuncBool1) (p *Promise) {
	p = makepromise()
	go func() {
		var a AnyVal
		for {
			a = f()
			if !are_all_true1(a, predicates...) {
				break
			}

			msg := new(qMsg)
			msg.a = a
			p.send(msg)

		}
		p.close()

	}()

	return
}
