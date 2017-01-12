package iexec_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/BooleanCat/igo/ios/iexec"
)

var _ = Describe("iexec", func() {
	Describe("Exec", func() {
		It("is implemented by Real", func() {
			var _ iexec.Exec = iexec.New()
		})

		It("is implemented by Fake", func() {
			var _ iexec.Exec = iexec.NewFake()
		})
	})

	Describe("Cmd", func() {
		It("is implemented by CmdReal", func() {
			var _ iexec.Cmd = iexec.NewCmd()
		})

		It("is implemented by CmdFake", func() {
			var _ iexec.Cmd = iexec.NewCmdFake()
		})
	})
})
