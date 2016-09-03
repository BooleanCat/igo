Installation
============

`go get github.com/BooleanCat/igo/os/exec`

Usage
=====

iexec (interfacing os/exec)
---------------------------

When writing unit tests for a package that calls a specific system command, it can be helpful to assert that the call was made without actually calling the target binary (it may be long running or have side effects that are undesirable to have to deal with in a unit test). `iexec.Exec.Command` is a wrapper around `exec.Command` that makes this possible by using a provider pattern such as:

```go
func foo(command iexec.CmdProvider) {
    cmd := command("my-binary")
    cmd.Run()
}
```

Since the `CmdProvider` has been used, we can substitute a fake instead of the real thing that can be used to make assertions. This is demonstrated using the [ginkgo](https://onsi.github.io/ginkgo/) testing framework below.

```go
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
        Expect(commandName).To(Equal("my-binary"))
    })
})
...
```

When using `foo` in a non unit test situation, just use `foo(new(iexec.Exec).Command)`.

In the example above, we haven't covered testing around whether or not our function called `cmd.Run()`. This can be easily achieved by using the helper `iexec.NewPureFake()`:

```go
...
Describe("foo", func(), {
    var execFakes *iexec.PureFake

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

`iexec.NewPureFake()` will even set up a fake `ios.Process` (interfacing `os.Process`). In the below example, the killing of a process can be asserted.

```go
func bar(command iexec.CmdProvider) {
    cmd := command("my-binary")
    cmd.Start()
    cmd.GetProcess().Kill()
}

...
Describe("foo", func(), {
    var execFakes *iexec.PureFake

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
