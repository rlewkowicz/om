// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type BoshFormService struct {
	GetFormStub        func(path string) (api.Form, error)
	getFormMutex       sync.RWMutex
	getFormArgsForCall []struct {
		path string
	}
	getFormReturns struct {
		result1 api.Form
		result2 error
	}
	PostFormStub        func(api.PostFormInput) error
	postFormMutex       sync.RWMutex
	postFormArgsForCall []struct {
		arg1 api.PostFormInput
	}
	postFormReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BoshFormService) GetForm(path string) (api.Form, error) {
	fake.getFormMutex.Lock()
	fake.getFormArgsForCall = append(fake.getFormArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("GetForm", []interface{}{path})
	fake.getFormMutex.Unlock()
	if fake.GetFormStub != nil {
		return fake.GetFormStub(path)
	} else {
		return fake.getFormReturns.result1, fake.getFormReturns.result2
	}
}

func (fake *BoshFormService) GetFormCallCount() int {
	fake.getFormMutex.RLock()
	defer fake.getFormMutex.RUnlock()
	return len(fake.getFormArgsForCall)
}

func (fake *BoshFormService) GetFormArgsForCall(i int) string {
	fake.getFormMutex.RLock()
	defer fake.getFormMutex.RUnlock()
	return fake.getFormArgsForCall[i].path
}

func (fake *BoshFormService) GetFormReturns(result1 api.Form, result2 error) {
	fake.GetFormStub = nil
	fake.getFormReturns = struct {
		result1 api.Form
		result2 error
	}{result1, result2}
}

func (fake *BoshFormService) PostForm(arg1 api.PostFormInput) error {
	fake.postFormMutex.Lock()
	fake.postFormArgsForCall = append(fake.postFormArgsForCall, struct {
		arg1 api.PostFormInput
	}{arg1})
	fake.recordInvocation("PostForm", []interface{}{arg1})
	fake.postFormMutex.Unlock()
	if fake.PostFormStub != nil {
		return fake.PostFormStub(arg1)
	} else {
		return fake.postFormReturns.result1
	}
}

func (fake *BoshFormService) PostFormCallCount() int {
	fake.postFormMutex.RLock()
	defer fake.postFormMutex.RUnlock()
	return len(fake.postFormArgsForCall)
}

func (fake *BoshFormService) PostFormArgsForCall(i int) api.PostFormInput {
	fake.postFormMutex.RLock()
	defer fake.postFormMutex.RUnlock()
	return fake.postFormArgsForCall[i].arg1
}

func (fake *BoshFormService) PostFormReturns(result1 error) {
	fake.PostFormStub = nil
	fake.postFormReturns = struct {
		result1 error
	}{result1}
}

func (fake *BoshFormService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getFormMutex.RLock()
	defer fake.getFormMutex.RUnlock()
	fake.postFormMutex.RLock()
	defer fake.postFormMutex.RUnlock()
	return fake.invocations
}

func (fake *BoshFormService) recordInvocation(key string, args []interface{}) {
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