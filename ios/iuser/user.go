package iuser

import "os/user"

//User is an interface around os/user
type User interface {
	Lookup(string) (*user.User, error)
	LookupGroup(string) (*user.Group, error)
	Current() (*user.User, error)
}

//Real is a wrapper around user that implements iuser.User
type Real struct{}

//New creates a struct that behaves like the user package
func New() *Real {
	return new(Real)
}

//Lookup is a wrapper around user.Lookup()
func (*Real) Lookup(username string) (*user.User, error) {
	return user.Lookup(username)
}

//LookupGroup is a wrapper around user.LookupGroup()
func (*Real) LookupGroup(name string) (*user.Group, error) {
	return user.LookupGroup(name)
}

//Current is a wrapper around user.Current()
func (*Real) Current() (*user.User, error) {
	return user.Current()
}
