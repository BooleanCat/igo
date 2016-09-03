// This file was generated by counterfeiter

package ios

import (
	"os"
	"sync"
)

//ProcessFake ...
type ProcessFake struct {
	KillStub        func() error
	killMutex       sync.RWMutex
	killArgsForCall []struct{}
	killReturns     struct {
		result1 error
	}
	ReleaseStub        func() error
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct{}
	releaseReturns     struct {
		result1 error
	}
	SignalStub        func(os.Signal) error
	signalMutex       sync.RWMutex
	signalArgsForCall []struct {
		arg1 os.Signal
	}
	signalReturns struct {
		result1 error
	}
	WaitStub        func() (*os.ProcessState, error)
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns     struct {
		result1 *os.ProcessState
		result2 error
	}
	GetPidStub        func() int
	getPidMutex       sync.RWMutex
	getPidArgsForCall []struct{}
	getPidReturns     struct {
		result1 int
	}
	SetPidStub        func(int)
	setPidMutex       sync.RWMutex
	setPidArgsForCall []struct {
		arg1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

//Kill ...
func (fake *ProcessFake) Kill() error {
	fake.killMutex.Lock()
	fake.killArgsForCall = append(fake.killArgsForCall, struct{}{})
	fake.recordInvocation("Kill", []interface{}{})
	fake.killMutex.Unlock()
	if fake.KillStub != nil {
		return fake.KillStub()
	}
	return fake.killReturns.result1
}

//KillCallCount ...
func (fake *ProcessFake) KillCallCount() int {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return len(fake.killArgsForCall)
}

//KillReturns ...
func (fake *ProcessFake) KillReturns(result1 error) {
	fake.KillStub = nil
	fake.killReturns = struct {
		result1 error
	}{result1}
}

//Release ...
func (fake *ProcessFake) Release() error {
	fake.releaseMutex.Lock()
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct{}{})
	fake.recordInvocation("Release", []interface{}{})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		return fake.ReleaseStub()
	}
	return fake.releaseReturns.result1
}

//ReleaseCallCount ...
func (fake *ProcessFake) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

//ReleaseReturns ...
func (fake *ProcessFake) ReleaseReturns(result1 error) {
	fake.ReleaseStub = nil
	fake.releaseReturns = struct {
		result1 error
	}{result1}
}

//Signal ...
func (fake *ProcessFake) Signal(arg1 os.Signal) error {
	fake.signalMutex.Lock()
	fake.signalArgsForCall = append(fake.signalArgsForCall, struct {
		arg1 os.Signal
	}{arg1})
	fake.recordInvocation("Signal", []interface{}{arg1})
	fake.signalMutex.Unlock()
	if fake.SignalStub != nil {
		return fake.SignalStub(arg1)
	}
	return fake.signalReturns.result1
}

//SignalCallCount ...
func (fake *ProcessFake) SignalCallCount() int {
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	return len(fake.signalArgsForCall)
}

//SignalArgsForCall ...
func (fake *ProcessFake) SignalArgsForCall(i int) os.Signal {
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	return fake.signalArgsForCall[i].arg1
}

//SignalReturns ...
func (fake *ProcessFake) SignalReturns(result1 error) {
	fake.SignalStub = nil
	fake.signalReturns = struct {
		result1 error
	}{result1}
}

//Wait ...
func (fake *ProcessFake) Wait() (*os.ProcessState, error) {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	}
	return fake.waitReturns.result1, fake.waitReturns.result2
}

//WaitCallCount ...
func (fake *ProcessFake) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

//WaitReturns ...
func (fake *ProcessFake) WaitReturns(result1 *os.ProcessState, result2 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 *os.ProcessState
		result2 error
	}{result1, result2}
}

//GetPid ...
func (fake *ProcessFake) GetPid() int {
	fake.getPidMutex.Lock()
	fake.getPidArgsForCall = append(fake.getPidArgsForCall, struct{}{})
	fake.recordInvocation("GetPid", []interface{}{})
	fake.getPidMutex.Unlock()
	if fake.GetPidStub != nil {
		return fake.GetPidStub()
	}
	return fake.getPidReturns.result1
}

//GetPidCallCount ...
func (fake *ProcessFake) GetPidCallCount() int {
	fake.getPidMutex.RLock()
	defer fake.getPidMutex.RUnlock()
	return len(fake.getPidArgsForCall)
}

//GetPidReturns ...
func (fake *ProcessFake) GetPidReturns(result1 int) {
	fake.GetPidStub = nil
	fake.getPidReturns = struct {
		result1 int
	}{result1}
}

//SetPid ...
func (fake *ProcessFake) SetPid(arg1 int) {
	fake.setPidMutex.Lock()
	fake.setPidArgsForCall = append(fake.setPidArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("SetPid", []interface{}{arg1})
	fake.setPidMutex.Unlock()
	if fake.SetPidStub != nil {
		fake.SetPidStub(arg1)
	}
}

//SetPidCallCount ...
func (fake *ProcessFake) SetPidCallCount() int {
	fake.setPidMutex.RLock()
	defer fake.setPidMutex.RUnlock()
	return len(fake.setPidArgsForCall)
}

//SetPidArgsForCall ...
func (fake *ProcessFake) SetPidArgsForCall(i int) int {
	fake.setPidMutex.RLock()
	defer fake.setPidMutex.RUnlock()
	return fake.setPidArgsForCall[i].arg1
}

//Invocations ...
func (fake *ProcessFake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	fake.getPidMutex.RLock()
	defer fake.getPidMutex.RUnlock()
	fake.setPidMutex.RLock()
	defer fake.setPidMutex.RUnlock()
	return fake.invocations
}

func (fake *ProcessFake) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ Process = new(ProcessFake)
