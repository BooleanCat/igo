package iexec

import (
	"bytes"
	"io/ioutil"
	"os/exec"

	"github.com/BooleanCat/igo/ios"
)

//Exec is an interface around os/exec
type Exec interface {
	Command(string, ...string) Cmd
}

/*
PureFake returns a struct containing fake Exec with nested initialised fake
members.

The following Fakes are available:
- Cmd: a FakeCmd returned by Exec.Command()
- Process: a FakeProcess returned by Cmd.GetProcess()
*/
type PureFake struct {
	Exec    *Fake
	Cmd     *CmdFake
	Process *ios.ProcessFake
}

//NewPureFake returns a fake Exec with nested new fakes within
func NewPureFake() *PureFake {
	execFake := NewFake()
	processFake := new(ios.ProcessFake)
	cmdFake := newPureCmdFake(processFake)
	execFake.CommandReturns(cmdFake)
	return &PureFake{
		Exec:    execFake,
		Cmd:     cmdFake,
		Process: processFake,
	}
}

func newPureCmdFake(process ios.Process) *CmdFake {
	fake := NewCmdFake()
	fake.GetProcessReturns(process)
	fake.StdoutPipeReturns(ioutil.NopCloser(new(bytes.Buffer)), nil)
	fake.StderrPipeReturns(ioutil.NopCloser(new(bytes.Buffer)), nil)
	return fake
}

//Real is a wrapper around exec that implements iexec.Exec
type Real struct{}

//New creates a struct that behaves like the exec package
func New() *Real {
	return new(Real)
}

//Command is a wrapper around exec.Command()
func (*Real) Command(name string, args ...string) Cmd {
	return &CmdReal{cmd: exec.Command(name, args...)}
}
