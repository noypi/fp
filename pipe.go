package fp

// caller is response in closing 'a'
func PipeChan(a chan AnyVal) (p *Promise) {
	p = makepromise()
	go func() {
		for b := range a {
			msg := new(qMsg)
			p.send(msg)
		}
		p.close()
	}()

	return
}

func PipeChan(a *Promise) (p *Promise) {
	p = makepromise()
	go func() {
		for b := range a.q {
			p.send(b)
		}
		p.close()
	}()

	return
}
