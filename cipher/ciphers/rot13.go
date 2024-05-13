package ciphers

type Rot13Cipher struct{}

type Rot13Actions interface {
}

func (r *Rot13Cipher) Handle(s string, e int) {}
