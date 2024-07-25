// Code generated by counterfeiter. DO NOT EDIT.
package metricsfakes

import (
	"sync"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
)

type FakeSubStatter struct {
	DecStub        func(string, int64, float32, ...statsd.Tag) error
	decMutex       sync.RWMutex
	decArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}
	decReturns struct {
		result1 error
	}
	decReturnsOnCall map[int]struct {
		result1 error
	}
	GaugeStub        func(string, int64, float32, ...statsd.Tag) error
	gaugeMutex       sync.RWMutex
	gaugeArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}
	gaugeReturns struct {
		result1 error
	}
	gaugeReturnsOnCall map[int]struct {
		result1 error
	}
	GaugeDeltaStub        func(string, int64, float32, ...statsd.Tag) error
	gaugeDeltaMutex       sync.RWMutex
	gaugeDeltaArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}
	gaugeDeltaReturns struct {
		result1 error
	}
	gaugeDeltaReturnsOnCall map[int]struct {
		result1 error
	}
	IncStub        func(string, int64, float32, ...statsd.Tag) error
	incMutex       sync.RWMutex
	incArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}
	incReturns struct {
		result1 error
	}
	incReturnsOnCall map[int]struct {
		result1 error
	}
	NewSubStatterStub        func(string) statsd.SubStatter
	newSubStatterMutex       sync.RWMutex
	newSubStatterArgsForCall []struct {
		arg1 string
	}
	newSubStatterReturns struct {
		result1 statsd.SubStatter
	}
	newSubStatterReturnsOnCall map[int]struct {
		result1 statsd.SubStatter
	}
	RawStub        func(string, string, float32, ...statsd.Tag) error
	rawMutex       sync.RWMutex
	rawArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 float32
		arg4 []statsd.Tag
	}
	rawReturns struct {
		result1 error
	}
	rawReturnsOnCall map[int]struct {
		result1 error
	}
	SetStub        func(string, string, float32, ...statsd.Tag) error
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 float32
		arg4 []statsd.Tag
	}
	setReturns struct {
		result1 error
	}
	setReturnsOnCall map[int]struct {
		result1 error
	}
	SetIntStub        func(string, int64, float32, ...statsd.Tag) error
	setIntMutex       sync.RWMutex
	setIntArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}
	setIntReturns struct {
		result1 error
	}
	setIntReturnsOnCall map[int]struct {
		result1 error
	}
	SetSamplerFuncStub        func(statsd.SamplerFunc)
	setSamplerFuncMutex       sync.RWMutex
	setSamplerFuncArgsForCall []struct {
		arg1 statsd.SamplerFunc
	}
	TimingStub        func(string, int64, float32, ...statsd.Tag) error
	timingMutex       sync.RWMutex
	timingArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}
	timingReturns struct {
		result1 error
	}
	timingReturnsOnCall map[int]struct {
		result1 error
	}
	TimingDurationStub        func(string, time.Duration, float32, ...statsd.Tag) error
	timingDurationMutex       sync.RWMutex
	timingDurationArgsForCall []struct {
		arg1 string
		arg2 time.Duration
		arg3 float32
		arg4 []statsd.Tag
	}
	timingDurationReturns struct {
		result1 error
	}
	timingDurationReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSubStatter) Dec(arg1 string, arg2 int64, arg3 float32, arg4 ...statsd.Tag) error {
	fake.decMutex.Lock()
	ret, specificReturn := fake.decReturnsOnCall[len(fake.decArgsForCall)]
	fake.decArgsForCall = append(fake.decArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}{arg1, arg2, arg3, arg4})
	stub := fake.DecStub
	fakeReturns := fake.decReturns
	fake.recordInvocation("Dec", []interface{}{arg1, arg2, arg3, arg4})
	fake.decMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubStatter) DecCallCount() int {
	fake.decMutex.RLock()
	defer fake.decMutex.RUnlock()
	return len(fake.decArgsForCall)
}

func (fake *FakeSubStatter) DecCalls(stub func(string, int64, float32, ...statsd.Tag) error) {
	fake.decMutex.Lock()
	defer fake.decMutex.Unlock()
	fake.DecStub = stub
}

func (fake *FakeSubStatter) DecArgsForCall(i int) (string, int64, float32, []statsd.Tag) {
	fake.decMutex.RLock()
	defer fake.decMutex.RUnlock()
	argsForCall := fake.decArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSubStatter) DecReturns(result1 error) {
	fake.decMutex.Lock()
	defer fake.decMutex.Unlock()
	fake.DecStub = nil
	fake.decReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) DecReturnsOnCall(i int, result1 error) {
	fake.decMutex.Lock()
	defer fake.decMutex.Unlock()
	fake.DecStub = nil
	if fake.decReturnsOnCall == nil {
		fake.decReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.decReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) Gauge(arg1 string, arg2 int64, arg3 float32, arg4 ...statsd.Tag) error {
	fake.gaugeMutex.Lock()
	ret, specificReturn := fake.gaugeReturnsOnCall[len(fake.gaugeArgsForCall)]
	fake.gaugeArgsForCall = append(fake.gaugeArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}{arg1, arg2, arg3, arg4})
	stub := fake.GaugeStub
	fakeReturns := fake.gaugeReturns
	fake.recordInvocation("Gauge", []interface{}{arg1, arg2, arg3, arg4})
	fake.gaugeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubStatter) GaugeCallCount() int {
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	return len(fake.gaugeArgsForCall)
}

func (fake *FakeSubStatter) GaugeCalls(stub func(string, int64, float32, ...statsd.Tag) error) {
	fake.gaugeMutex.Lock()
	defer fake.gaugeMutex.Unlock()
	fake.GaugeStub = stub
}

func (fake *FakeSubStatter) GaugeArgsForCall(i int) (string, int64, float32, []statsd.Tag) {
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	argsForCall := fake.gaugeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSubStatter) GaugeReturns(result1 error) {
	fake.gaugeMutex.Lock()
	defer fake.gaugeMutex.Unlock()
	fake.GaugeStub = nil
	fake.gaugeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) GaugeReturnsOnCall(i int, result1 error) {
	fake.gaugeMutex.Lock()
	defer fake.gaugeMutex.Unlock()
	fake.GaugeStub = nil
	if fake.gaugeReturnsOnCall == nil {
		fake.gaugeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.gaugeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) GaugeDelta(arg1 string, arg2 int64, arg3 float32, arg4 ...statsd.Tag) error {
	fake.gaugeDeltaMutex.Lock()
	ret, specificReturn := fake.gaugeDeltaReturnsOnCall[len(fake.gaugeDeltaArgsForCall)]
	fake.gaugeDeltaArgsForCall = append(fake.gaugeDeltaArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}{arg1, arg2, arg3, arg4})
	stub := fake.GaugeDeltaStub
	fakeReturns := fake.gaugeDeltaReturns
	fake.recordInvocation("GaugeDelta", []interface{}{arg1, arg2, arg3, arg4})
	fake.gaugeDeltaMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubStatter) GaugeDeltaCallCount() int {
	fake.gaugeDeltaMutex.RLock()
	defer fake.gaugeDeltaMutex.RUnlock()
	return len(fake.gaugeDeltaArgsForCall)
}

func (fake *FakeSubStatter) GaugeDeltaCalls(stub func(string, int64, float32, ...statsd.Tag) error) {
	fake.gaugeDeltaMutex.Lock()
	defer fake.gaugeDeltaMutex.Unlock()
	fake.GaugeDeltaStub = stub
}

func (fake *FakeSubStatter) GaugeDeltaArgsForCall(i int) (string, int64, float32, []statsd.Tag) {
	fake.gaugeDeltaMutex.RLock()
	defer fake.gaugeDeltaMutex.RUnlock()
	argsForCall := fake.gaugeDeltaArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSubStatter) GaugeDeltaReturns(result1 error) {
	fake.gaugeDeltaMutex.Lock()
	defer fake.gaugeDeltaMutex.Unlock()
	fake.GaugeDeltaStub = nil
	fake.gaugeDeltaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) GaugeDeltaReturnsOnCall(i int, result1 error) {
	fake.gaugeDeltaMutex.Lock()
	defer fake.gaugeDeltaMutex.Unlock()
	fake.GaugeDeltaStub = nil
	if fake.gaugeDeltaReturnsOnCall == nil {
		fake.gaugeDeltaReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.gaugeDeltaReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) Inc(arg1 string, arg2 int64, arg3 float32, arg4 ...statsd.Tag) error {
	fake.incMutex.Lock()
	ret, specificReturn := fake.incReturnsOnCall[len(fake.incArgsForCall)]
	fake.incArgsForCall = append(fake.incArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}{arg1, arg2, arg3, arg4})
	stub := fake.IncStub
	fakeReturns := fake.incReturns
	fake.recordInvocation("Inc", []interface{}{arg1, arg2, arg3, arg4})
	fake.incMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubStatter) IncCallCount() int {
	fake.incMutex.RLock()
	defer fake.incMutex.RUnlock()
	return len(fake.incArgsForCall)
}

func (fake *FakeSubStatter) IncCalls(stub func(string, int64, float32, ...statsd.Tag) error) {
	fake.incMutex.Lock()
	defer fake.incMutex.Unlock()
	fake.IncStub = stub
}

func (fake *FakeSubStatter) IncArgsForCall(i int) (string, int64, float32, []statsd.Tag) {
	fake.incMutex.RLock()
	defer fake.incMutex.RUnlock()
	argsForCall := fake.incArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSubStatter) IncReturns(result1 error) {
	fake.incMutex.Lock()
	defer fake.incMutex.Unlock()
	fake.IncStub = nil
	fake.incReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) IncReturnsOnCall(i int, result1 error) {
	fake.incMutex.Lock()
	defer fake.incMutex.Unlock()
	fake.IncStub = nil
	if fake.incReturnsOnCall == nil {
		fake.incReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.incReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) NewSubStatter(arg1 string) statsd.SubStatter {
	fake.newSubStatterMutex.Lock()
	ret, specificReturn := fake.newSubStatterReturnsOnCall[len(fake.newSubStatterArgsForCall)]
	fake.newSubStatterArgsForCall = append(fake.newSubStatterArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.NewSubStatterStub
	fakeReturns := fake.newSubStatterReturns
	fake.recordInvocation("NewSubStatter", []interface{}{arg1})
	fake.newSubStatterMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubStatter) NewSubStatterCallCount() int {
	fake.newSubStatterMutex.RLock()
	defer fake.newSubStatterMutex.RUnlock()
	return len(fake.newSubStatterArgsForCall)
}

func (fake *FakeSubStatter) NewSubStatterCalls(stub func(string) statsd.SubStatter) {
	fake.newSubStatterMutex.Lock()
	defer fake.newSubStatterMutex.Unlock()
	fake.NewSubStatterStub = stub
}

func (fake *FakeSubStatter) NewSubStatterArgsForCall(i int) string {
	fake.newSubStatterMutex.RLock()
	defer fake.newSubStatterMutex.RUnlock()
	argsForCall := fake.newSubStatterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubStatter) NewSubStatterReturns(result1 statsd.SubStatter) {
	fake.newSubStatterMutex.Lock()
	defer fake.newSubStatterMutex.Unlock()
	fake.NewSubStatterStub = nil
	fake.newSubStatterReturns = struct {
		result1 statsd.SubStatter
	}{result1}
}

func (fake *FakeSubStatter) NewSubStatterReturnsOnCall(i int, result1 statsd.SubStatter) {
	fake.newSubStatterMutex.Lock()
	defer fake.newSubStatterMutex.Unlock()
	fake.NewSubStatterStub = nil
	if fake.newSubStatterReturnsOnCall == nil {
		fake.newSubStatterReturnsOnCall = make(map[int]struct {
			result1 statsd.SubStatter
		})
	}
	fake.newSubStatterReturnsOnCall[i] = struct {
		result1 statsd.SubStatter
	}{result1}
}

func (fake *FakeSubStatter) Raw(arg1 string, arg2 string, arg3 float32, arg4 ...statsd.Tag) error {
	fake.rawMutex.Lock()
	ret, specificReturn := fake.rawReturnsOnCall[len(fake.rawArgsForCall)]
	fake.rawArgsForCall = append(fake.rawArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 float32
		arg4 []statsd.Tag
	}{arg1, arg2, arg3, arg4})
	stub := fake.RawStub
	fakeReturns := fake.rawReturns
	fake.recordInvocation("Raw", []interface{}{arg1, arg2, arg3, arg4})
	fake.rawMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubStatter) RawCallCount() int {
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	return len(fake.rawArgsForCall)
}

func (fake *FakeSubStatter) RawCalls(stub func(string, string, float32, ...statsd.Tag) error) {
	fake.rawMutex.Lock()
	defer fake.rawMutex.Unlock()
	fake.RawStub = stub
}

func (fake *FakeSubStatter) RawArgsForCall(i int) (string, string, float32, []statsd.Tag) {
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	argsForCall := fake.rawArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSubStatter) RawReturns(result1 error) {
	fake.rawMutex.Lock()
	defer fake.rawMutex.Unlock()
	fake.RawStub = nil
	fake.rawReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) RawReturnsOnCall(i int, result1 error) {
	fake.rawMutex.Lock()
	defer fake.rawMutex.Unlock()
	fake.RawStub = nil
	if fake.rawReturnsOnCall == nil {
		fake.rawReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rawReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) Set(arg1 string, arg2 string, arg3 float32, arg4 ...statsd.Tag) error {
	fake.setMutex.Lock()
	ret, specificReturn := fake.setReturnsOnCall[len(fake.setArgsForCall)]
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 float32
		arg4 []statsd.Tag
	}{arg1, arg2, arg3, arg4})
	stub := fake.SetStub
	fakeReturns := fake.setReturns
	fake.recordInvocation("Set", []interface{}{arg1, arg2, arg3, arg4})
	fake.setMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubStatter) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeSubStatter) SetCalls(stub func(string, string, float32, ...statsd.Tag) error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = stub
}

func (fake *FakeSubStatter) SetArgsForCall(i int) (string, string, float32, []statsd.Tag) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	argsForCall := fake.setArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSubStatter) SetReturns(result1 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	fake.setReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) SetReturnsOnCall(i int, result1 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	if fake.setReturnsOnCall == nil {
		fake.setReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) SetInt(arg1 string, arg2 int64, arg3 float32, arg4 ...statsd.Tag) error {
	fake.setIntMutex.Lock()
	ret, specificReturn := fake.setIntReturnsOnCall[len(fake.setIntArgsForCall)]
	fake.setIntArgsForCall = append(fake.setIntArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}{arg1, arg2, arg3, arg4})
	stub := fake.SetIntStub
	fakeReturns := fake.setIntReturns
	fake.recordInvocation("SetInt", []interface{}{arg1, arg2, arg3, arg4})
	fake.setIntMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubStatter) SetIntCallCount() int {
	fake.setIntMutex.RLock()
	defer fake.setIntMutex.RUnlock()
	return len(fake.setIntArgsForCall)
}

func (fake *FakeSubStatter) SetIntCalls(stub func(string, int64, float32, ...statsd.Tag) error) {
	fake.setIntMutex.Lock()
	defer fake.setIntMutex.Unlock()
	fake.SetIntStub = stub
}

func (fake *FakeSubStatter) SetIntArgsForCall(i int) (string, int64, float32, []statsd.Tag) {
	fake.setIntMutex.RLock()
	defer fake.setIntMutex.RUnlock()
	argsForCall := fake.setIntArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSubStatter) SetIntReturns(result1 error) {
	fake.setIntMutex.Lock()
	defer fake.setIntMutex.Unlock()
	fake.SetIntStub = nil
	fake.setIntReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) SetIntReturnsOnCall(i int, result1 error) {
	fake.setIntMutex.Lock()
	defer fake.setIntMutex.Unlock()
	fake.SetIntStub = nil
	if fake.setIntReturnsOnCall == nil {
		fake.setIntReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setIntReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) SetSamplerFunc(arg1 statsd.SamplerFunc) {
	fake.setSamplerFuncMutex.Lock()
	fake.setSamplerFuncArgsForCall = append(fake.setSamplerFuncArgsForCall, struct {
		arg1 statsd.SamplerFunc
	}{arg1})
	stub := fake.SetSamplerFuncStub
	fake.recordInvocation("SetSamplerFunc", []interface{}{arg1})
	fake.setSamplerFuncMutex.Unlock()
	if stub != nil {
		fake.SetSamplerFuncStub(arg1)
	}
}

func (fake *FakeSubStatter) SetSamplerFuncCallCount() int {
	fake.setSamplerFuncMutex.RLock()
	defer fake.setSamplerFuncMutex.RUnlock()
	return len(fake.setSamplerFuncArgsForCall)
}

func (fake *FakeSubStatter) SetSamplerFuncCalls(stub func(statsd.SamplerFunc)) {
	fake.setSamplerFuncMutex.Lock()
	defer fake.setSamplerFuncMutex.Unlock()
	fake.SetSamplerFuncStub = stub
}

func (fake *FakeSubStatter) SetSamplerFuncArgsForCall(i int) statsd.SamplerFunc {
	fake.setSamplerFuncMutex.RLock()
	defer fake.setSamplerFuncMutex.RUnlock()
	argsForCall := fake.setSamplerFuncArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubStatter) Timing(arg1 string, arg2 int64, arg3 float32, arg4 ...statsd.Tag) error {
	fake.timingMutex.Lock()
	ret, specificReturn := fake.timingReturnsOnCall[len(fake.timingArgsForCall)]
	fake.timingArgsForCall = append(fake.timingArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
		arg4 []statsd.Tag
	}{arg1, arg2, arg3, arg4})
	stub := fake.TimingStub
	fakeReturns := fake.timingReturns
	fake.recordInvocation("Timing", []interface{}{arg1, arg2, arg3, arg4})
	fake.timingMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubStatter) TimingCallCount() int {
	fake.timingMutex.RLock()
	defer fake.timingMutex.RUnlock()
	return len(fake.timingArgsForCall)
}

func (fake *FakeSubStatter) TimingCalls(stub func(string, int64, float32, ...statsd.Tag) error) {
	fake.timingMutex.Lock()
	defer fake.timingMutex.Unlock()
	fake.TimingStub = stub
}

func (fake *FakeSubStatter) TimingArgsForCall(i int) (string, int64, float32, []statsd.Tag) {
	fake.timingMutex.RLock()
	defer fake.timingMutex.RUnlock()
	argsForCall := fake.timingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSubStatter) TimingReturns(result1 error) {
	fake.timingMutex.Lock()
	defer fake.timingMutex.Unlock()
	fake.TimingStub = nil
	fake.timingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) TimingReturnsOnCall(i int, result1 error) {
	fake.timingMutex.Lock()
	defer fake.timingMutex.Unlock()
	fake.TimingStub = nil
	if fake.timingReturnsOnCall == nil {
		fake.timingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.timingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) TimingDuration(arg1 string, arg2 time.Duration, arg3 float32, arg4 ...statsd.Tag) error {
	fake.timingDurationMutex.Lock()
	ret, specificReturn := fake.timingDurationReturnsOnCall[len(fake.timingDurationArgsForCall)]
	fake.timingDurationArgsForCall = append(fake.timingDurationArgsForCall, struct {
		arg1 string
		arg2 time.Duration
		arg3 float32
		arg4 []statsd.Tag
	}{arg1, arg2, arg3, arg4})
	stub := fake.TimingDurationStub
	fakeReturns := fake.timingDurationReturns
	fake.recordInvocation("TimingDuration", []interface{}{arg1, arg2, arg3, arg4})
	fake.timingDurationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubStatter) TimingDurationCallCount() int {
	fake.timingDurationMutex.RLock()
	defer fake.timingDurationMutex.RUnlock()
	return len(fake.timingDurationArgsForCall)
}

func (fake *FakeSubStatter) TimingDurationCalls(stub func(string, time.Duration, float32, ...statsd.Tag) error) {
	fake.timingDurationMutex.Lock()
	defer fake.timingDurationMutex.Unlock()
	fake.TimingDurationStub = stub
}

func (fake *FakeSubStatter) TimingDurationArgsForCall(i int) (string, time.Duration, float32, []statsd.Tag) {
	fake.timingDurationMutex.RLock()
	defer fake.timingDurationMutex.RUnlock()
	argsForCall := fake.timingDurationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSubStatter) TimingDurationReturns(result1 error) {
	fake.timingDurationMutex.Lock()
	defer fake.timingDurationMutex.Unlock()
	fake.TimingDurationStub = nil
	fake.timingDurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) TimingDurationReturnsOnCall(i int, result1 error) {
	fake.timingDurationMutex.Lock()
	defer fake.timingDurationMutex.Unlock()
	fake.TimingDurationStub = nil
	if fake.timingDurationReturnsOnCall == nil {
		fake.timingDurationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.timingDurationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSubStatter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.decMutex.RLock()
	defer fake.decMutex.RUnlock()
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	fake.gaugeDeltaMutex.RLock()
	defer fake.gaugeDeltaMutex.RUnlock()
	fake.incMutex.RLock()
	defer fake.incMutex.RUnlock()
	fake.newSubStatterMutex.RLock()
	defer fake.newSubStatterMutex.RUnlock()
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	fake.setIntMutex.RLock()
	defer fake.setIntMutex.RUnlock()
	fake.setSamplerFuncMutex.RLock()
	defer fake.setSamplerFuncMutex.RUnlock()
	fake.timingMutex.RLock()
	defer fake.timingMutex.RUnlock()
	fake.timingDurationMutex.RLock()
	defer fake.timingDurationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSubStatter) recordInvocation(key string, args []interface{}) {
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

var _ statsd.SubStatter = new(FakeSubStatter)
