// This file was generated by counterfeiter

package iioutil

import (
	"io"
	"os"
	"sync"
)

//IoutilFake ...
type IoutilFake struct {
	ReadAllStub        func(io.Reader) ([]byte, error)
	readAllMutex       sync.RWMutex
	readAllArgsForCall []struct {
		arg1 io.Reader
	}
	readAllReturns struct {
		result1 []byte
		result2 error
	}
	ReadDirStub        func(string) ([]os.FileInfo, error)
	readDirMutex       sync.RWMutex
	readDirArgsForCall []struct {
		arg1 string
	}
	readDirReturns struct {
		result1 []os.FileInfo
		result2 error
	}
	ReadFileStub        func(string) ([]byte, error)
	readFileMutex       sync.RWMutex
	readFileArgsForCall []struct {
		arg1 string
	}
	readFileReturns struct {
		result1 []byte
		result2 error
	}
	TempDirStub        func(string, string) (string, error)
	tempDirMutex       sync.RWMutex
	tempDirArgsForCall []struct {
		arg1 string
		arg2 string
	}
	tempDirReturns struct {
		result1 string
		result2 error
	}
	TempFileStub        func(string) (*os.File, error)
	tempFileMutex       sync.RWMutex
	tempFileArgsForCall []struct {
		arg1 string
	}
	tempFileReturns struct {
		result1 *os.File
		result2 error
	}
	WriteFileStub        func(string, []byte, os.FileMode) error
	writeFileMutex       sync.RWMutex
	writeFileArgsForCall []struct {
		arg1 string
		arg2 []byte
		arg3 os.FileMode
	}
	writeFileReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

//ReadAll ...
func (fake *IoutilFake) ReadAll(arg1 io.Reader) ([]byte, error) {
	fake.readAllMutex.Lock()
	fake.readAllArgsForCall = append(fake.readAllArgsForCall, struct {
		arg1 io.Reader
	}{arg1})
	fake.recordInvocation("ReadAll", []interface{}{arg1})
	fake.readAllMutex.Unlock()
	if fake.ReadAllStub != nil {
		return fake.ReadAllStub(arg1)
	}
	return fake.readAllReturns.result1, fake.readAllReturns.result2
}

//ReadAllCallCount ...
func (fake *IoutilFake) ReadAllCallCount() int {
	fake.readAllMutex.RLock()
	defer fake.readAllMutex.RUnlock()
	return len(fake.readAllArgsForCall)
}

//ReadAllArgsForCall ...
func (fake *IoutilFake) ReadAllArgsForCall(i int) io.Reader {
	fake.readAllMutex.RLock()
	defer fake.readAllMutex.RUnlock()
	return fake.readAllArgsForCall[i].arg1
}

//ReadAllReturns ...
func (fake *IoutilFake) ReadAllReturns(result1 []byte, result2 error) {
	fake.ReadAllStub = nil
	fake.readAllReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

//ReadDir ...
func (fake *IoutilFake) ReadDir(arg1 string) ([]os.FileInfo, error) {
	fake.readDirMutex.Lock()
	fake.readDirArgsForCall = append(fake.readDirArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReadDir", []interface{}{arg1})
	fake.readDirMutex.Unlock()
	if fake.ReadDirStub != nil {
		return fake.ReadDirStub(arg1)
	}
	return fake.readDirReturns.result1, fake.readDirReturns.result2
}

//ReadDirCallCount ...
func (fake *IoutilFake) ReadDirCallCount() int {
	fake.readDirMutex.RLock()
	defer fake.readDirMutex.RUnlock()
	return len(fake.readDirArgsForCall)
}

//ReadDirArgsForCall ...
func (fake *IoutilFake) ReadDirArgsForCall(i int) string {
	fake.readDirMutex.RLock()
	defer fake.readDirMutex.RUnlock()
	return fake.readDirArgsForCall[i].arg1
}

//ReadDirReturns ...
func (fake *IoutilFake) ReadDirReturns(result1 []os.FileInfo, result2 error) {
	fake.ReadDirStub = nil
	fake.readDirReturns = struct {
		result1 []os.FileInfo
		result2 error
	}{result1, result2}
}

//ReadFile ...
func (fake *IoutilFake) ReadFile(arg1 string) ([]byte, error) {
	fake.readFileMutex.Lock()
	fake.readFileArgsForCall = append(fake.readFileArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReadFile", []interface{}{arg1})
	fake.readFileMutex.Unlock()
	if fake.ReadFileStub != nil {
		return fake.ReadFileStub(arg1)
	}
	return fake.readFileReturns.result1, fake.readFileReturns.result2
}

//ReadFileCallCount ...
func (fake *IoutilFake) ReadFileCallCount() int {
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	return len(fake.readFileArgsForCall)
}

//ReadFileArgsForCall ...
func (fake *IoutilFake) ReadFileArgsForCall(i int) string {
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	return fake.readFileArgsForCall[i].arg1
}

//ReadFileReturns ...
func (fake *IoutilFake) ReadFileReturns(result1 []byte, result2 error) {
	fake.ReadFileStub = nil
	fake.readFileReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

//TempDir ...
func (fake *IoutilFake) TempDir(arg1 string, arg2 string) (string, error) {
	fake.tempDirMutex.Lock()
	fake.tempDirArgsForCall = append(fake.tempDirArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("TempDir", []interface{}{arg1, arg2})
	fake.tempDirMutex.Unlock()
	if fake.TempDirStub != nil {
		return fake.TempDirStub(arg1, arg2)
	}
	return fake.tempDirReturns.result1, fake.tempDirReturns.result2
}

//TempDirCallCount ...
func (fake *IoutilFake) TempDirCallCount() int {
	fake.tempDirMutex.RLock()
	defer fake.tempDirMutex.RUnlock()
	return len(fake.tempDirArgsForCall)
}

//TempDirArgsForCall ...
func (fake *IoutilFake) TempDirArgsForCall(i int) (string, string) {
	fake.tempDirMutex.RLock()
	defer fake.tempDirMutex.RUnlock()
	return fake.tempDirArgsForCall[i].arg1, fake.tempDirArgsForCall[i].arg2
}

//TempDirReturns ...
func (fake *IoutilFake) TempDirReturns(result1 string, result2 error) {
	fake.TempDirStub = nil
	fake.tempDirReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

//TempFile ...
func (fake *IoutilFake) TempFile(arg1 string) (*os.File, error) {
	fake.tempFileMutex.Lock()
	fake.tempFileArgsForCall = append(fake.tempFileArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("TempFile", []interface{}{arg1})
	fake.tempFileMutex.Unlock()
	if fake.TempFileStub != nil {
		return fake.TempFileStub(arg1)
	}
	return fake.tempFileReturns.result1, fake.tempFileReturns.result2
}

//TempFileCallCount ...
func (fake *IoutilFake) TempFileCallCount() int {
	fake.tempFileMutex.RLock()
	defer fake.tempFileMutex.RUnlock()
	return len(fake.tempFileArgsForCall)
}

//TempFileArgsForCall ...
func (fake *IoutilFake) TempFileArgsForCall(i int) string {
	fake.tempFileMutex.RLock()
	defer fake.tempFileMutex.RUnlock()
	return fake.tempFileArgsForCall[i].arg1
}

//TempFileReturns ...
func (fake *IoutilFake) TempFileReturns(result1 *os.File, result2 error) {
	fake.TempFileStub = nil
	fake.tempFileReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

//WriteFile ...
func (fake *IoutilFake) WriteFile(arg1 string, arg2 []byte, arg3 os.FileMode) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeFileMutex.Lock()
	fake.writeFileArgsForCall = append(fake.writeFileArgsForCall, struct {
		arg1 string
		arg2 []byte
		arg3 os.FileMode
	}{arg1, arg2Copy, arg3})
	fake.recordInvocation("WriteFile", []interface{}{arg1, arg2Copy, arg3})
	fake.writeFileMutex.Unlock()
	if fake.WriteFileStub != nil {
		return fake.WriteFileStub(arg1, arg2, arg3)
	}
	return fake.writeFileReturns.result1
}

//WriteFileCallCount ...
func (fake *IoutilFake) WriteFileCallCount() int {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return len(fake.writeFileArgsForCall)
}

//WriteFileArgsForCall ...
func (fake *IoutilFake) WriteFileArgsForCall(i int) (string, []byte, os.FileMode) {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return fake.writeFileArgsForCall[i].arg1, fake.writeFileArgsForCall[i].arg2, fake.writeFileArgsForCall[i].arg3
}

//WriteFileReturns ...
func (fake *IoutilFake) WriteFileReturns(result1 error) {
	fake.WriteFileStub = nil
	fake.writeFileReturns = struct {
		result1 error
	}{result1}
}

//Invocations ...
func (fake *IoutilFake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readAllMutex.RLock()
	defer fake.readAllMutex.RUnlock()
	fake.readDirMutex.RLock()
	defer fake.readDirMutex.RUnlock()
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	fake.tempDirMutex.RLock()
	defer fake.tempDirMutex.RUnlock()
	fake.tempFileMutex.RLock()
	defer fake.tempFileMutex.RUnlock()
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return fake.invocations
}

func (fake *IoutilFake) recordInvocation(key string, args []interface{}) {
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

var _ Ioutil = new(IoutilFake)
