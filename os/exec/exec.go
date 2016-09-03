package iexec

import (
	"bytes"
	"io/ioutil"
	"os/exec"

	ios "github.com/BooleanCat/igo/os"
)

//Exec is an interface around os/exec
type Exec interface {
	Command(string, ...string) Cmd
}

//NewFake returns a fake Exec with nested new fakes within
func NewFake() (*ExecFake, *CmdFake, *ios.ProcessFake) {
	fake := new(ExecFake)
	cmdFake := new(CmdFake)
	processFake := new(ios.ProcessFake)
	cmdFake.GetProcessReturns(processFake)
	cmdFake.StdoutPipeReturns(ioutil.NopCloser(new(bytes.Buffer)), nil)
	cmdFake.StderrPipeReturns(ioutil.NopCloser(new(bytes.Buffer)), nil)
	fake.CommandReturns(cmdFake)
	return fake, cmdFake, processFake
}

//ExecWrap is a wrapper around exec that implements iexec.Exec
type ExecWrap struct{}

//Command is a wrapper around exec.Command()
func (e *ExecWrap) Command(name string, args ...string) Cmd {
	return &CmdWrap{cmd: exec.Command(name, args...)}
}
