// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf-experimental/routing-api/authentication"
)

type FakeCloudControllerApi struct {
	GetInfoStub        func() *authentication.Info
	getInfoMutex       sync.RWMutex
	getInfoArgsForCall []struct{}
	getInfoReturns struct {
		result1 *authentication.Info
	}
}

func (fake *FakeCloudControllerApi) GetInfo() *authentication.Info {
	fake.getInfoMutex.Lock()
	fake.getInfoArgsForCall = append(fake.getInfoArgsForCall, struct{}{})
	fake.getInfoMutex.Unlock()
	if fake.GetInfoStub != nil {
		return fake.GetInfoStub()
	} else {
		return fake.getInfoReturns.result1
	}
}

func (fake *FakeCloudControllerApi) GetInfoCallCount() int {
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return len(fake.getInfoArgsForCall)
}

func (fake *FakeCloudControllerApi) GetInfoReturns(result1 *authentication.Info) {
	fake.GetInfoStub = nil
	fake.getInfoReturns = struct {
		result1 *authentication.Info
	}{result1}
}

var _ authentication.CloudControllerApi = new(FakeCloudControllerApi)