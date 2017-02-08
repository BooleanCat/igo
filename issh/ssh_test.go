package issh_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/BooleanCat/igo/issh"
)

var _ = Describe("issh", func() {
	Describe("SSH", func() {
		It("is implemented by Real", func() {
			var _ issh.SSH = issh.New()
		})

		It("is implemented by Fake", func() {
			var _ issh.SSH = issh.NewFake()
		})
	})
})
