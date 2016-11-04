package iuser

import "os/user"

//User is an interface around os/user
type User interface {
	Lookup(string) (*user.User, error)
	LookupGroup(string) (*user.Group, error)
}

//UserWrap is a wrapper around user that implements iuser.User
type UserWrap struct{}

//Lookup is a wrapper around user.Lookup()
func (userWrap *UserWrap) Lookup(username string) (*user.User, error) {
	return user.Lookup(username)
}

//LookupGroup is a wrapper around user.LookupGroup()
func (userWrap *UserWrap) LookupGroup(name string) (*user.Group, error) {
	return user.LookupGroup(name)
}
