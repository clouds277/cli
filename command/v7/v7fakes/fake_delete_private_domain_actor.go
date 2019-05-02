// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeDeletePrivateDomainActor struct {
	CheckSharedDomainStub        func(string) (v7action.Warnings, error)
	checkSharedDomainMutex       sync.RWMutex
	checkSharedDomainArgsForCall []struct {
		arg1 string
	}
	checkSharedDomainReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	checkSharedDomainReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	DeletePrivateDomainStub        func(string) (v7action.Warnings, error)
	deletePrivateDomainMutex       sync.RWMutex
	deletePrivateDomainArgsForCall []struct {
		arg1 string
	}
	deletePrivateDomainReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	deletePrivateDomainReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeletePrivateDomainActor) CheckSharedDomain(arg1 string) (v7action.Warnings, error) {
	fake.checkSharedDomainMutex.Lock()
	ret, specificReturn := fake.checkSharedDomainReturnsOnCall[len(fake.checkSharedDomainArgsForCall)]
	fake.checkSharedDomainArgsForCall = append(fake.checkSharedDomainArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("CheckSharedDomain", []interface{}{arg1})
	fake.checkSharedDomainMutex.Unlock()
	if fake.CheckSharedDomainStub != nil {
		return fake.CheckSharedDomainStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.checkSharedDomainReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeletePrivateDomainActor) CheckSharedDomainCallCount() int {
	fake.checkSharedDomainMutex.RLock()
	defer fake.checkSharedDomainMutex.RUnlock()
	return len(fake.checkSharedDomainArgsForCall)
}

func (fake *FakeDeletePrivateDomainActor) CheckSharedDomainCalls(stub func(string) (v7action.Warnings, error)) {
	fake.checkSharedDomainMutex.Lock()
	defer fake.checkSharedDomainMutex.Unlock()
	fake.CheckSharedDomainStub = stub
}

func (fake *FakeDeletePrivateDomainActor) CheckSharedDomainArgsForCall(i int) string {
	fake.checkSharedDomainMutex.RLock()
	defer fake.checkSharedDomainMutex.RUnlock()
	argsForCall := fake.checkSharedDomainArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDeletePrivateDomainActor) CheckSharedDomainReturns(result1 v7action.Warnings, result2 error) {
	fake.checkSharedDomainMutex.Lock()
	defer fake.checkSharedDomainMutex.Unlock()
	fake.CheckSharedDomainStub = nil
	fake.checkSharedDomainReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeletePrivateDomainActor) CheckSharedDomainReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.checkSharedDomainMutex.Lock()
	defer fake.checkSharedDomainMutex.Unlock()
	fake.CheckSharedDomainStub = nil
	if fake.checkSharedDomainReturnsOnCall == nil {
		fake.checkSharedDomainReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.checkSharedDomainReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeletePrivateDomainActor) DeletePrivateDomain(arg1 string) (v7action.Warnings, error) {
	fake.deletePrivateDomainMutex.Lock()
	ret, specificReturn := fake.deletePrivateDomainReturnsOnCall[len(fake.deletePrivateDomainArgsForCall)]
	fake.deletePrivateDomainArgsForCall = append(fake.deletePrivateDomainArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeletePrivateDomain", []interface{}{arg1})
	fake.deletePrivateDomainMutex.Unlock()
	if fake.DeletePrivateDomainStub != nil {
		return fake.DeletePrivateDomainStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deletePrivateDomainReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeletePrivateDomainActor) DeletePrivateDomainCallCount() int {
	fake.deletePrivateDomainMutex.RLock()
	defer fake.deletePrivateDomainMutex.RUnlock()
	return len(fake.deletePrivateDomainArgsForCall)
}

func (fake *FakeDeletePrivateDomainActor) DeletePrivateDomainCalls(stub func(string) (v7action.Warnings, error)) {
	fake.deletePrivateDomainMutex.Lock()
	defer fake.deletePrivateDomainMutex.Unlock()
	fake.DeletePrivateDomainStub = stub
}

func (fake *FakeDeletePrivateDomainActor) DeletePrivateDomainArgsForCall(i int) string {
	fake.deletePrivateDomainMutex.RLock()
	defer fake.deletePrivateDomainMutex.RUnlock()
	argsForCall := fake.deletePrivateDomainArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDeletePrivateDomainActor) DeletePrivateDomainReturns(result1 v7action.Warnings, result2 error) {
	fake.deletePrivateDomainMutex.Lock()
	defer fake.deletePrivateDomainMutex.Unlock()
	fake.DeletePrivateDomainStub = nil
	fake.deletePrivateDomainReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeletePrivateDomainActor) DeletePrivateDomainReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.deletePrivateDomainMutex.Lock()
	defer fake.deletePrivateDomainMutex.Unlock()
	fake.DeletePrivateDomainStub = nil
	if fake.deletePrivateDomainReturnsOnCall == nil {
		fake.deletePrivateDomainReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.deletePrivateDomainReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeletePrivateDomainActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkSharedDomainMutex.RLock()
	defer fake.checkSharedDomainMutex.RUnlock()
	fake.deletePrivateDomainMutex.RLock()
	defer fake.deletePrivateDomainMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeletePrivateDomainActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.DeletePrivateDomainActor = new(FakeDeletePrivateDomainActor)
