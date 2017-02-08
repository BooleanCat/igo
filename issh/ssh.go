package issh

import "golang.org/x/crypto/ssh"

//SSH is an interface around ssh
type SSH interface {
	Dial(string, string, *ssh.ClientConfig) (*ssh.Client, error)
}

//Real is a wrapper around ssh that implements issh.SSH
type Real struct{}

//New creates a struct that behaves like the ssh package
func New() *Real {
	return new(Real)
}

//Dial is a wrapper around ssh.Dial()
func (*Real) Dial(protocol, address string, config *ssh.ClientConfig) (*ssh.Client, error) {
	return ssh.Dial(protocol, address, config)
}
