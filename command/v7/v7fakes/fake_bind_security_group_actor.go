// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	v7 "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/resources"
)

type FakeBindSecurityGroupActor struct {
	BindSecurityGroupToSpaceStub        func(string, string, constant.SecurityGroupLifecycle) (v7action.Warnings, error)
	bindSecurityGroupToSpaceMutex       sync.RWMutex
	bindSecurityGroupToSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 constant.SecurityGroupLifecycle
	}
	bindSecurityGroupToSpaceReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	bindSecurityGroupToSpaceReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetOrganizationByNameStub        func(string) (v7action.Organization, v7action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		arg1 string
	}
	getOrganizationByNameReturns struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	GetOrganizationSpacesStub        func(string) ([]v7action.Space, v7action.Warnings, error)
	getOrganizationSpacesMutex       sync.RWMutex
	getOrganizationSpacesArgsForCall []struct {
		arg1 string
	}
	getOrganizationSpacesReturns struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationSpacesReturnsOnCall map[int]struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	GetSecurityGroupStub        func(string) (resources.SecurityGroup, v7action.Warnings, error)
	getSecurityGroupMutex       sync.RWMutex
	getSecurityGroupArgsForCall []struct {
		arg1 string
	}
	getSecurityGroupReturns struct {
		result1 resources.SecurityGroup
		result2 v7action.Warnings
		result3 error
	}
	getSecurityGroupReturnsOnCall map[int]struct {
		result1 resources.SecurityGroup
		result2 v7action.Warnings
		result3 error
	}
	GetSpaceByNameAndOrganizationStub        func(string, string) (v7action.Space, v7action.Warnings, error)
	getSpaceByNameAndOrganizationMutex       sync.RWMutex
	getSpaceByNameAndOrganizationArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceByNameAndOrganizationReturns struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	getSpaceByNameAndOrganizationReturnsOnCall map[int]struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBindSecurityGroupActor) BindSecurityGroupToSpace(arg1 string, arg2 string, arg3 constant.SecurityGroupLifecycle) (v7action.Warnings, error) {
	fake.bindSecurityGroupToSpaceMutex.Lock()
	ret, specificReturn := fake.bindSecurityGroupToSpaceReturnsOnCall[len(fake.bindSecurityGroupToSpaceArgsForCall)]
	fake.bindSecurityGroupToSpaceArgsForCall = append(fake.bindSecurityGroupToSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 constant.SecurityGroupLifecycle
	}{arg1, arg2, arg3})
	fake.recordInvocation("BindSecurityGroupToSpace", []interface{}{arg1, arg2, arg3})
	fake.bindSecurityGroupToSpaceMutex.Unlock()
	if fake.BindSecurityGroupToSpaceStub != nil {
		return fake.BindSecurityGroupToSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.bindSecurityGroupToSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBindSecurityGroupActor) BindSecurityGroupToSpaceCallCount() int {
	fake.bindSecurityGroupToSpaceMutex.RLock()
	defer fake.bindSecurityGroupToSpaceMutex.RUnlock()
	return len(fake.bindSecurityGroupToSpaceArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) BindSecurityGroupToSpaceCalls(stub func(string, string, constant.SecurityGroupLifecycle) (v7action.Warnings, error)) {
	fake.bindSecurityGroupToSpaceMutex.Lock()
	defer fake.bindSecurityGroupToSpaceMutex.Unlock()
	fake.BindSecurityGroupToSpaceStub = stub
}

func (fake *FakeBindSecurityGroupActor) BindSecurityGroupToSpaceArgsForCall(i int) (string, string, constant.SecurityGroupLifecycle) {
	fake.bindSecurityGroupToSpaceMutex.RLock()
	defer fake.bindSecurityGroupToSpaceMutex.RUnlock()
	argsForCall := fake.bindSecurityGroupToSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBindSecurityGroupActor) BindSecurityGroupToSpaceReturns(result1 v7action.Warnings, result2 error) {
	fake.bindSecurityGroupToSpaceMutex.Lock()
	defer fake.bindSecurityGroupToSpaceMutex.Unlock()
	fake.BindSecurityGroupToSpaceStub = nil
	fake.bindSecurityGroupToSpaceReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeBindSecurityGroupActor) BindSecurityGroupToSpaceReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.bindSecurityGroupToSpaceMutex.Lock()
	defer fake.bindSecurityGroupToSpaceMutex.Unlock()
	fake.BindSecurityGroupToSpaceStub = nil
	if fake.bindSecurityGroupToSpaceReturnsOnCall == nil {
		fake.bindSecurityGroupToSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.bindSecurityGroupToSpaceReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeBindSecurityGroupActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeBindSecurityGroupActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeBindSecurityGroupActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBindSecurityGroupActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationByName(arg1 string) (v7action.Organization, v7action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationByName", []interface{}{arg1})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationByNameCalls(stub func(string) (v7action.Organization, v7action.Warnings, error)) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = stub
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	argsForCall := fake.getOrganizationByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationByNameReturns(result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationByNameReturnsOnCall(i int, result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Organization
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationSpaces(arg1 string) ([]v7action.Space, v7action.Warnings, error) {
	fake.getOrganizationSpacesMutex.Lock()
	ret, specificReturn := fake.getOrganizationSpacesReturnsOnCall[len(fake.getOrganizationSpacesArgsForCall)]
	fake.getOrganizationSpacesArgsForCall = append(fake.getOrganizationSpacesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationSpaces", []interface{}{arg1})
	fake.getOrganizationSpacesMutex.Unlock()
	if fake.GetOrganizationSpacesStub != nil {
		return fake.GetOrganizationSpacesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationSpacesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationSpacesCallCount() int {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	return len(fake.getOrganizationSpacesArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationSpacesCalls(stub func(string) ([]v7action.Space, v7action.Warnings, error)) {
	fake.getOrganizationSpacesMutex.Lock()
	defer fake.getOrganizationSpacesMutex.Unlock()
	fake.GetOrganizationSpacesStub = stub
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationSpacesArgsForCall(i int) string {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	argsForCall := fake.getOrganizationSpacesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationSpacesReturns(result1 []v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationSpacesMutex.Lock()
	defer fake.getOrganizationSpacesMutex.Unlock()
	fake.GetOrganizationSpacesStub = nil
	fake.getOrganizationSpacesReturns = struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) GetOrganizationSpacesReturnsOnCall(i int, result1 []v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationSpacesMutex.Lock()
	defer fake.getOrganizationSpacesMutex.Unlock()
	fake.GetOrganizationSpacesStub = nil
	if fake.getOrganizationSpacesReturnsOnCall == nil {
		fake.getOrganizationSpacesReturnsOnCall = make(map[int]struct {
			result1 []v7action.Space
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationSpacesReturnsOnCall[i] = struct {
		result1 []v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) GetSecurityGroup(arg1 string) (resources.SecurityGroup, v7action.Warnings, error) {
	fake.getSecurityGroupMutex.Lock()
	ret, specificReturn := fake.getSecurityGroupReturnsOnCall[len(fake.getSecurityGroupArgsForCall)]
	fake.getSecurityGroupArgsForCall = append(fake.getSecurityGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetSecurityGroup", []interface{}{arg1})
	fake.getSecurityGroupMutex.Unlock()
	if fake.GetSecurityGroupStub != nil {
		return fake.GetSecurityGroupStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSecurityGroupReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBindSecurityGroupActor) GetSecurityGroupCallCount() int {
	fake.getSecurityGroupMutex.RLock()
	defer fake.getSecurityGroupMutex.RUnlock()
	return len(fake.getSecurityGroupArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) GetSecurityGroupCalls(stub func(string) (resources.SecurityGroup, v7action.Warnings, error)) {
	fake.getSecurityGroupMutex.Lock()
	defer fake.getSecurityGroupMutex.Unlock()
	fake.GetSecurityGroupStub = stub
}

func (fake *FakeBindSecurityGroupActor) GetSecurityGroupArgsForCall(i int) string {
	fake.getSecurityGroupMutex.RLock()
	defer fake.getSecurityGroupMutex.RUnlock()
	argsForCall := fake.getSecurityGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindSecurityGroupActor) GetSecurityGroupReturns(result1 resources.SecurityGroup, result2 v7action.Warnings, result3 error) {
	fake.getSecurityGroupMutex.Lock()
	defer fake.getSecurityGroupMutex.Unlock()
	fake.GetSecurityGroupStub = nil
	fake.getSecurityGroupReturns = struct {
		result1 resources.SecurityGroup
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) GetSecurityGroupReturnsOnCall(i int, result1 resources.SecurityGroup, result2 v7action.Warnings, result3 error) {
	fake.getSecurityGroupMutex.Lock()
	defer fake.getSecurityGroupMutex.Unlock()
	fake.GetSecurityGroupStub = nil
	if fake.getSecurityGroupReturnsOnCall == nil {
		fake.getSecurityGroupReturnsOnCall = make(map[int]struct {
			result1 resources.SecurityGroup
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSecurityGroupReturnsOnCall[i] = struct {
		result1 resources.SecurityGroup
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) GetSpaceByNameAndOrganization(arg1 string, arg2 string) (v7action.Space, v7action.Warnings, error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	ret, specificReturn := fake.getSpaceByNameAndOrganizationReturnsOnCall[len(fake.getSpaceByNameAndOrganizationArgsForCall)]
	fake.getSpaceByNameAndOrganizationArgsForCall = append(fake.getSpaceByNameAndOrganizationArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSpaceByNameAndOrganization", []interface{}{arg1, arg2})
	fake.getSpaceByNameAndOrganizationMutex.Unlock()
	if fake.GetSpaceByNameAndOrganizationStub != nil {
		return fake.GetSpaceByNameAndOrganizationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceByNameAndOrganizationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBindSecurityGroupActor) GetSpaceByNameAndOrganizationCallCount() int {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	return len(fake.getSpaceByNameAndOrganizationArgsForCall)
}

func (fake *FakeBindSecurityGroupActor) GetSpaceByNameAndOrganizationCalls(stub func(string, string) (v7action.Space, v7action.Warnings, error)) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = stub
}

func (fake *FakeBindSecurityGroupActor) GetSpaceByNameAndOrganizationArgsForCall(i int) (string, string) {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	argsForCall := fake.getSpaceByNameAndOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBindSecurityGroupActor) GetSpaceByNameAndOrganizationReturns(result1 v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	fake.getSpaceByNameAndOrganizationReturns = struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) GetSpaceByNameAndOrganizationReturnsOnCall(i int, result1 v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	if fake.getSpaceByNameAndOrganizationReturnsOnCall == nil {
		fake.getSpaceByNameAndOrganizationReturnsOnCall = make(map[int]struct {
			result1 v7action.Space
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSpaceByNameAndOrganizationReturnsOnCall[i] = struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBindSecurityGroupActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bindSecurityGroupToSpaceMutex.RLock()
	defer fake.bindSecurityGroupToSpaceMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	fake.getSecurityGroupMutex.RLock()
	defer fake.getSecurityGroupMutex.RUnlock()
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBindSecurityGroupActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.BindSecurityGroupActor = new(FakeBindSecurityGroupActor)