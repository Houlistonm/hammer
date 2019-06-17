// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/pivotal/pcf-cli/commands"
	"github.com/pivotal/pcf-cli/environment"
)

type FakeEnvReader struct {
	ReadStub        func(string) (environment.Config, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 string
	}
	readReturns struct {
		result1 environment.Config
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 environment.Config
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnvReader) Read(arg1 string) (environment.Config, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Read", []interface{}{arg1})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.readReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEnvReader) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeEnvReader) ReadCalls(stub func(string) (environment.Config, error)) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeEnvReader) ReadArgsForCall(i int) string {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	argsForCall := fake.readArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEnvReader) ReadReturns(result1 environment.Config, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 environment.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvReader) ReadReturnsOnCall(i int, result1 environment.Config, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 environment.Config
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 environment.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEnvReader) recordInvocation(key string, args []interface{}) {
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

var _ commands.EnvReader = new(FakeEnvReader)
