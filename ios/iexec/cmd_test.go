package iexec_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/BooleanCat/igo/ios/iexec"
)

var _ = Describe("iexec", func() {
	Describe("Cmd", func() {
		It("is implemented by CmdReal", func() {
			var _ iexec.Cmd = iexec.NewCmd()
		})

		It("is implemented by CmdFake", func() {
			var _ iexec.Cmd = iexec.NewCmdFake()
		})
	})
})
