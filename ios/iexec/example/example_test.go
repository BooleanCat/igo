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
	var execFakes *iexec.PureFake

	BeforeEach(func() {
		execFakes = iexec.NewPureFake()
	})

	Describe("foo", func() {
		var (
			commandName string
			commandArgs []string
			joinedArgs  string
		)

		BeforeEach(func() {
			foo(execFakes.Exec.Command)
			commandName, commandArgs = execFakes.Exec.CommandArgsForCall(0)
			joinedArgs = strings.Join(commandArgs, " ")
		})

		It("calls my-binary", func() {
			Expect(commandName).To(Equal("my-binary"))
		})

		It("uses the correct flags", func() {
			Expect(joinedArgs).To(ContainSubstring("-p 4"))
		})

		It("calls cmd.Start()", func() {
			Expect(execFakes.Cmd.StartCallCount()).To(Equal(1))
		})

		It("calls Process.Kill()", func() {
			Expect(execFakes.Process.KillCallCount()).To(Equal(1))
		})
	})
})
