package iuser_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/BooleanCat/igo/ios/iuser"
)

var _ = Describe("iuser", func() {
	It("is implemented by Real", func() {
		var _ iuser.User = iuser.New()
	})

	It("is implemented by Fake", func() {
		var _ iuser.User = iuser.NewFake()
	})
})
