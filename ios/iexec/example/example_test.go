package example_test

import (
	"strings"

	"github.com/BooleanCat/igo/ios/iexec"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func foo(command iexec.CmdProvider) {
	cmd := command("my-binary", "-p", "4")
	cmd.Start()
	cmd.GetProcess().Kill()
}

var _ = Describe("iexec", func() {
	var fakes *iexec.NestedCommandFake

	BeforeEach(func() {
		fakes = iexec.NewNestedCommandFake()
	})

	Describe("foo", func() {
		var (
			commandName string
			commandArgs []string
			joinedArgs  string
		)

		BeforeEach(func() {
			foo(fakes.Exec.Command)
			commandName, commandArgs = fakes.Exec.CommandArgsForCall(0)
			joinedArgs = strings.Join(commandArgs, " ")
		})

		It("calls my-binary", func() {
			Expect(commandName).To(Equal("my-binary"))
		})

		It("uses the correct flags", func() {
			Expect(joinedArgs).To(ContainSubstring("-p 4"))
		})

		It("calls cmd.Start()", func() {
			Expect(fakes.Cmd.StartCallCount()).To(Equal(1))
		})

		It("calls Process.Kill()", func() {
			Expect(fakes.Process.KillCallCount()).To(Equal(1))
		})
	})
})
