package fp

// caller is response in closing 'a'
func PipeChan(a chan interface{}) (p *Promise) {
	p = makepromise()
	go func() {
		for b := range a {
			msg := new(qMsg)
			msg.a = b
			p.send(msg)
		}
		p.close()
	}()

	return
}
