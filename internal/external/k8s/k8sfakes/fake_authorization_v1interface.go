// Code generated by counterfeiter. DO NOT EDIT.
package k8sfakes

import (
	"sync"

	"github.com/observatorium/opa-openshift/internal/external/k8s"
	v1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/client-go/rest"
)

type FakeAuthorizationV1Interface struct {
	LocalSubjectAccessReviewsStub        func(string) v1.LocalSubjectAccessReviewInterface
	localSubjectAccessReviewsMutex       sync.RWMutex
	localSubjectAccessReviewsArgsForCall []struct {
		arg1 string
	}
	localSubjectAccessReviewsReturns struct {
		result1 v1.LocalSubjectAccessReviewInterface
	}
	localSubjectAccessReviewsReturnsOnCall map[int]struct {
		result1 v1.LocalSubjectAccessReviewInterface
	}
	RESTClientStub        func() rest.Interface
	rESTClientMutex       sync.RWMutex
	rESTClientArgsForCall []struct {
	}
	rESTClientReturns struct {
		result1 rest.Interface
	}
	rESTClientReturnsOnCall map[int]struct {
		result1 rest.Interface
	}
	SelfSubjectAccessReviewsStub        func() v1.SelfSubjectAccessReviewInterface
	selfSubjectAccessReviewsMutex       sync.RWMutex
	selfSubjectAccessReviewsArgsForCall []struct {
	}
	selfSubjectAccessReviewsReturns struct {
		result1 v1.SelfSubjectAccessReviewInterface
	}
	selfSubjectAccessReviewsReturnsOnCall map[int]struct {
		result1 v1.SelfSubjectAccessReviewInterface
	}
	SelfSubjectRulesReviewsStub        func() v1.SelfSubjectRulesReviewInterface
	selfSubjectRulesReviewsMutex       sync.RWMutex
	selfSubjectRulesReviewsArgsForCall []struct {
	}
	selfSubjectRulesReviewsReturns struct {
		result1 v1.SelfSubjectRulesReviewInterface
	}
	selfSubjectRulesReviewsReturnsOnCall map[int]struct {
		result1 v1.SelfSubjectRulesReviewInterface
	}
	SubjectAccessReviewsStub        func() v1.SubjectAccessReviewInterface
	subjectAccessReviewsMutex       sync.RWMutex
	subjectAccessReviewsArgsForCall []struct {
	}
	subjectAccessReviewsReturns struct {
		result1 v1.SubjectAccessReviewInterface
	}
	subjectAccessReviewsReturnsOnCall map[int]struct {
		result1 v1.SubjectAccessReviewInterface
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthorizationV1Interface) LocalSubjectAccessReviews(arg1 string) v1.LocalSubjectAccessReviewInterface {
	fake.localSubjectAccessReviewsMutex.Lock()
	ret, specificReturn := fake.localSubjectAccessReviewsReturnsOnCall[len(fake.localSubjectAccessReviewsArgsForCall)]
	fake.localSubjectAccessReviewsArgsForCall = append(fake.localSubjectAccessReviewsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LocalSubjectAccessReviewsStub
	fakeReturns := fake.localSubjectAccessReviewsReturns
	fake.recordInvocation("LocalSubjectAccessReviews", []interface{}{arg1})
	fake.localSubjectAccessReviewsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthorizationV1Interface) LocalSubjectAccessReviewsCallCount() int {
	fake.localSubjectAccessReviewsMutex.RLock()
	defer fake.localSubjectAccessReviewsMutex.RUnlock()
	return len(fake.localSubjectAccessReviewsArgsForCall)
}

func (fake *FakeAuthorizationV1Interface) LocalSubjectAccessReviewsCalls(stub func(string) v1.LocalSubjectAccessReviewInterface) {
	fake.localSubjectAccessReviewsMutex.Lock()
	defer fake.localSubjectAccessReviewsMutex.Unlock()
	fake.LocalSubjectAccessReviewsStub = stub
}

func (fake *FakeAuthorizationV1Interface) LocalSubjectAccessReviewsArgsForCall(i int) string {
	fake.localSubjectAccessReviewsMutex.RLock()
	defer fake.localSubjectAccessReviewsMutex.RUnlock()
	argsForCall := fake.localSubjectAccessReviewsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAuthorizationV1Interface) LocalSubjectAccessReviewsReturns(result1 v1.LocalSubjectAccessReviewInterface) {
	fake.localSubjectAccessReviewsMutex.Lock()
	defer fake.localSubjectAccessReviewsMutex.Unlock()
	fake.LocalSubjectAccessReviewsStub = nil
	fake.localSubjectAccessReviewsReturns = struct {
		result1 v1.LocalSubjectAccessReviewInterface
	}{result1}
}

func (fake *FakeAuthorizationV1Interface) LocalSubjectAccessReviewsReturnsOnCall(i int, result1 v1.LocalSubjectAccessReviewInterface) {
	fake.localSubjectAccessReviewsMutex.Lock()
	defer fake.localSubjectAccessReviewsMutex.Unlock()
	fake.LocalSubjectAccessReviewsStub = nil
	if fake.localSubjectAccessReviewsReturnsOnCall == nil {
		fake.localSubjectAccessReviewsReturnsOnCall = make(map[int]struct {
			result1 v1.LocalSubjectAccessReviewInterface
		})
	}
	fake.localSubjectAccessReviewsReturnsOnCall[i] = struct {
		result1 v1.LocalSubjectAccessReviewInterface
	}{result1}
}

func (fake *FakeAuthorizationV1Interface) RESTClient() rest.Interface {
	fake.rESTClientMutex.Lock()
	ret, specificReturn := fake.rESTClientReturnsOnCall[len(fake.rESTClientArgsForCall)]
	fake.rESTClientArgsForCall = append(fake.rESTClientArgsForCall, struct {
	}{})
	stub := fake.RESTClientStub
	fakeReturns := fake.rESTClientReturns
	fake.recordInvocation("RESTClient", []interface{}{})
	fake.rESTClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthorizationV1Interface) RESTClientCallCount() int {
	fake.rESTClientMutex.RLock()
	defer fake.rESTClientMutex.RUnlock()
	return len(fake.rESTClientArgsForCall)
}

func (fake *FakeAuthorizationV1Interface) RESTClientCalls(stub func() rest.Interface) {
	fake.rESTClientMutex.Lock()
	defer fake.rESTClientMutex.Unlock()
	fake.RESTClientStub = stub
}

func (fake *FakeAuthorizationV1Interface) RESTClientReturns(result1 rest.Interface) {
	fake.rESTClientMutex.Lock()
	defer fake.rESTClientMutex.Unlock()
	fake.RESTClientStub = nil
	fake.rESTClientReturns = struct {
		result1 rest.Interface
	}{result1}
}

func (fake *FakeAuthorizationV1Interface) RESTClientReturnsOnCall(i int, result1 rest.Interface) {
	fake.rESTClientMutex.Lock()
	defer fake.rESTClientMutex.Unlock()
	fake.RESTClientStub = nil
	if fake.rESTClientReturnsOnCall == nil {
		fake.rESTClientReturnsOnCall = make(map[int]struct {
			result1 rest.Interface
		})
	}
	fake.rESTClientReturnsOnCall[i] = struct {
		result1 rest.Interface
	}{result1}
}

func (fake *FakeAuthorizationV1Interface) SelfSubjectAccessReviews() v1.SelfSubjectAccessReviewInterface {
	fake.selfSubjectAccessReviewsMutex.Lock()
	ret, specificReturn := fake.selfSubjectAccessReviewsReturnsOnCall[len(fake.selfSubjectAccessReviewsArgsForCall)]
	fake.selfSubjectAccessReviewsArgsForCall = append(fake.selfSubjectAccessReviewsArgsForCall, struct {
	}{})
	stub := fake.SelfSubjectAccessReviewsStub
	fakeReturns := fake.selfSubjectAccessReviewsReturns
	fake.recordInvocation("SelfSubjectAccessReviews", []interface{}{})
	fake.selfSubjectAccessReviewsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthorizationV1Interface) SelfSubjectAccessReviewsCallCount() int {
	fake.selfSubjectAccessReviewsMutex.RLock()
	defer fake.selfSubjectAccessReviewsMutex.RUnlock()
	return len(fake.selfSubjectAccessReviewsArgsForCall)
}

func (fake *FakeAuthorizationV1Interface) SelfSubjectAccessReviewsCalls(stub func() v1.SelfSubjectAccessReviewInterface) {
	fake.selfSubjectAccessReviewsMutex.Lock()
	defer fake.selfSubjectAccessReviewsMutex.Unlock()
	fake.SelfSubjectAccessReviewsStub = stub
}

func (fake *FakeAuthorizationV1Interface) SelfSubjectAccessReviewsReturns(result1 v1.SelfSubjectAccessReviewInterface) {
	fake.selfSubjectAccessReviewsMutex.Lock()
	defer fake.selfSubjectAccessReviewsMutex.Unlock()
	fake.SelfSubjectAccessReviewsStub = nil
	fake.selfSubjectAccessReviewsReturns = struct {
		result1 v1.SelfSubjectAccessReviewInterface
	}{result1}
}

func (fake *FakeAuthorizationV1Interface) SelfSubjectAccessReviewsReturnsOnCall(i int, result1 v1.SelfSubjectAccessReviewInterface) {
	fake.selfSubjectAccessReviewsMutex.Lock()
	defer fake.selfSubjectAccessReviewsMutex.Unlock()
	fake.SelfSubjectAccessReviewsStub = nil
	if fake.selfSubjectAccessReviewsReturnsOnCall == nil {
		fake.selfSubjectAccessReviewsReturnsOnCall = make(map[int]struct {
			result1 v1.SelfSubjectAccessReviewInterface
		})
	}
	fake.selfSubjectAccessReviewsReturnsOnCall[i] = struct {
		result1 v1.SelfSubjectAccessReviewInterface
	}{result1}
}

func (fake *FakeAuthorizationV1Interface) SelfSubjectRulesReviews() v1.SelfSubjectRulesReviewInterface {
	fake.selfSubjectRulesReviewsMutex.Lock()
	ret, specificReturn := fake.selfSubjectRulesReviewsReturnsOnCall[len(fake.selfSubjectRulesReviewsArgsForCall)]
	fake.selfSubjectRulesReviewsArgsForCall = append(fake.selfSubjectRulesReviewsArgsForCall, struct {
	}{})
	stub := fake.SelfSubjectRulesReviewsStub
	fakeReturns := fake.selfSubjectRulesReviewsReturns
	fake.recordInvocation("SelfSubjectRulesReviews", []interface{}{})
	fake.selfSubjectRulesReviewsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthorizationV1Interface) SelfSubjectRulesReviewsCallCount() int {
	fake.selfSubjectRulesReviewsMutex.RLock()
	defer fake.selfSubjectRulesReviewsMutex.RUnlock()
	return len(fake.selfSubjectRulesReviewsArgsForCall)
}

func (fake *FakeAuthorizationV1Interface) SelfSubjectRulesReviewsCalls(stub func() v1.SelfSubjectRulesReviewInterface) {
	fake.selfSubjectRulesReviewsMutex.Lock()
	defer fake.selfSubjectRulesReviewsMutex.Unlock()
	fake.SelfSubjectRulesReviewsStub = stub
}

func (fake *FakeAuthorizationV1Interface) SelfSubjectRulesReviewsReturns(result1 v1.SelfSubjectRulesReviewInterface) {
	fake.selfSubjectRulesReviewsMutex.Lock()
	defer fake.selfSubjectRulesReviewsMutex.Unlock()
	fake.SelfSubjectRulesReviewsStub = nil
	fake.selfSubjectRulesReviewsReturns = struct {
		result1 v1.SelfSubjectRulesReviewInterface
	}{result1}
}

func (fake *FakeAuthorizationV1Interface) SelfSubjectRulesReviewsReturnsOnCall(i int, result1 v1.SelfSubjectRulesReviewInterface) {
	fake.selfSubjectRulesReviewsMutex.Lock()
	defer fake.selfSubjectRulesReviewsMutex.Unlock()
	fake.SelfSubjectRulesReviewsStub = nil
	if fake.selfSubjectRulesReviewsReturnsOnCall == nil {
		fake.selfSubjectRulesReviewsReturnsOnCall = make(map[int]struct {
			result1 v1.SelfSubjectRulesReviewInterface
		})
	}
	fake.selfSubjectRulesReviewsReturnsOnCall[i] = struct {
		result1 v1.SelfSubjectRulesReviewInterface
	}{result1}
}

func (fake *FakeAuthorizationV1Interface) SubjectAccessReviews() v1.SubjectAccessReviewInterface {
	fake.subjectAccessReviewsMutex.Lock()
	ret, specificReturn := fake.subjectAccessReviewsReturnsOnCall[len(fake.subjectAccessReviewsArgsForCall)]
	fake.subjectAccessReviewsArgsForCall = append(fake.subjectAccessReviewsArgsForCall, struct {
	}{})
	stub := fake.SubjectAccessReviewsStub
	fakeReturns := fake.subjectAccessReviewsReturns
	fake.recordInvocation("SubjectAccessReviews", []interface{}{})
	fake.subjectAccessReviewsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthorizationV1Interface) SubjectAccessReviewsCallCount() int {
	fake.subjectAccessReviewsMutex.RLock()
	defer fake.subjectAccessReviewsMutex.RUnlock()
	return len(fake.subjectAccessReviewsArgsForCall)
}

func (fake *FakeAuthorizationV1Interface) SubjectAccessReviewsCalls(stub func() v1.SubjectAccessReviewInterface) {
	fake.subjectAccessReviewsMutex.Lock()
	defer fake.subjectAccessReviewsMutex.Unlock()
	fake.SubjectAccessReviewsStub = stub
}

func (fake *FakeAuthorizationV1Interface) SubjectAccessReviewsReturns(result1 v1.SubjectAccessReviewInterface) {
	fake.subjectAccessReviewsMutex.Lock()
	defer fake.subjectAccessReviewsMutex.Unlock()
	fake.SubjectAccessReviewsStub = nil
	fake.subjectAccessReviewsReturns = struct {
		result1 v1.SubjectAccessReviewInterface
	}{result1}
}

func (fake *FakeAuthorizationV1Interface) SubjectAccessReviewsReturnsOnCall(i int, result1 v1.SubjectAccessReviewInterface) {
	fake.subjectAccessReviewsMutex.Lock()
	defer fake.subjectAccessReviewsMutex.Unlock()
	fake.SubjectAccessReviewsStub = nil
	if fake.subjectAccessReviewsReturnsOnCall == nil {
		fake.subjectAccessReviewsReturnsOnCall = make(map[int]struct {
			result1 v1.SubjectAccessReviewInterface
		})
	}
	fake.subjectAccessReviewsReturnsOnCall[i] = struct {
		result1 v1.SubjectAccessReviewInterface
	}{result1}
}

func (fake *FakeAuthorizationV1Interface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.localSubjectAccessReviewsMutex.RLock()
	defer fake.localSubjectAccessReviewsMutex.RUnlock()
	fake.rESTClientMutex.RLock()
	defer fake.rESTClientMutex.RUnlock()
	fake.selfSubjectAccessReviewsMutex.RLock()
	defer fake.selfSubjectAccessReviewsMutex.RUnlock()
	fake.selfSubjectRulesReviewsMutex.RLock()
	defer fake.selfSubjectRulesReviewsMutex.RUnlock()
	fake.subjectAccessReviewsMutex.RLock()
	defer fake.subjectAccessReviewsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthorizationV1Interface) recordInvocation(key string, args []interface{}) {
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

var _ k8s.AuthorizationV1Interface = new(FakeAuthorizationV1Interface)
