package inet_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/BooleanCat/igo/inet"
)

var _ = Describe("inet", func() {
	Describe("Net", func() {
		It("is implemented by Real", func() {
			var _ inet.Net = inet.New()
		})

		It("is implemented by Fake", func() {
			var _ inet.Net = inet.NewFake()
		})
	})
})
