igo
===

igo is a collection of interfaces around Golang's standard packages - making it
easy to mock and make assertions about calls to functions within these packages.

A fake (generated using counterfeiter) is also provided for each interface for
along with helper functions that makes more complicated faking something you
don't need to think about.

Packages that expose public attributes have been given corresponding get and set
methods.

This is by no means a complete collection of interfaces and fakes and the focus
is on the most commonly used packages.

Installation
============

`go get github.com/BooleanCat/igo`

Usage
=====

iexec (interfacing os/exec)
---------------------------

When writing unit tests for a package that calls a specific system command, it
can be helpful to assert that the call was made without actually calling the
target binary (it may be log running or have side effects that are undesirable
to have to deal with in a unit test). `iexec.Exec.Command` is a wrapper around `exec.Command` that makes this possible by using a provider pattern.

```
func foo(command iexec.CmdProvider) {
    cmd := command("my-binary")
    cmd.Run()
}
```

Since the `CmdProvider` has been used, we can substitute a fake instead of
the real thing that can be used to make assertions. This is demonstrated using the `ginkgo` testing framework below.

```
...
Describe("foo", func(), {
    var (
        execFake    *iexec.ExecFake
        commandName string
    )

    BeforeEach(func() {
        execFake = new(iexec.ExecFake)
        foo(execFake.Command)
        commandName, _ = execFake.CommandArgsForCall(0)
    })

    It("calls out to my-binary", func() {
        Expect(commandName).To(Equal('my-binary'))
    })
})
...
```

When using `foo` in a non unit test situation, just use the provided
interface: `foo(iexec.Exec.Command)`.

In the example above, we haven't covered testing around whether or not our function called `cmd.Run()`. This can be easily achieved by using the helper `iexec.NewPureFake()`:

```
...
Describe("foo", func(), {
    var execFakes *iexec.PureFakeExec

    BeforeEach(func() {
        execFakes = iexec.NewPureFake()
        foo(execFakes.Exec.Command)
    })

    It("runs my-binary", func() {
        Expect(execFakes.Cmd.RunCallCount()).To(Equal(1))
    })
})
...
```

`iexec.NewPureFake()` will even set up a fake `ios.Process` (interfacing
`os.Process`). In the below example, the killing of a process can be asserted.

```
func bar(command iexec.CmdProvider) {
    cmd := command("my-binary")
    cmd.Start()
    cmd.GetProcess().Kill()
}

...
Describe("foo", func(), {
    var execFakes *iexec.PureFakeExec

    BeforeEach(func() {
        execFakes = iexec.NewPureFake()
        bar(execFakes.Exec.Command)
    })

    It("kills my-binary", func() {
        Expect(execFakes.Process.KillCallCount()).To(Equal(1))
    })
})
...
```
