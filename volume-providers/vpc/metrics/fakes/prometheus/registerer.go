// Code generated by counterfeiter. DO NOT EDIT.
package prometheus

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type Registerer struct {
	RegisterStub        func(prometheus.Collector) error
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		arg1 prometheus.Collector
	}
	registerReturns struct {
		result1 error
	}
	registerReturnsOnCall map[int]struct {
		result1 error
	}
	MustRegisterStub        func(...prometheus.Collector)
	mustRegisterMutex       sync.RWMutex
	mustRegisterArgsForCall []struct {
		arg1 []prometheus.Collector
	}
	UnregisterStub        func(prometheus.Collector) bool
	unregisterMutex       sync.RWMutex
	unregisterArgsForCall []struct {
		arg1 prometheus.Collector
	}
	unregisterReturns struct {
		result1 bool
	}
	unregisterReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Registerer) Register(arg1 prometheus.Collector) error {
	fake.registerMutex.Lock()
	ret, specificReturn := fake.registerReturnsOnCall[len(fake.registerArgsForCall)]
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		arg1 prometheus.Collector
	}{arg1})
	fake.recordInvocation("Register", []interface{}{arg1})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		return fake.RegisterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.registerReturns.result1
}

func (fake *Registerer) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *Registerer) RegisterArgsForCall(i int) prometheus.Collector {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return fake.registerArgsForCall[i].arg1
}

func (fake *Registerer) RegisterReturns(result1 error) {
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1 error
	}{result1}
}

func (fake *Registerer) RegisterReturnsOnCall(i int, result1 error) {
	fake.RegisterStub = nil
	if fake.registerReturnsOnCall == nil {
		fake.registerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Registerer) MustRegister(arg1 ...prometheus.Collector) {
	fake.mustRegisterMutex.Lock()
	fake.mustRegisterArgsForCall = append(fake.mustRegisterArgsForCall, struct {
		arg1 []prometheus.Collector
	}{arg1})
	fake.recordInvocation("MustRegister", []interface{}{arg1})
	fake.mustRegisterMutex.Unlock()
	if fake.MustRegisterStub != nil {
		fake.MustRegisterStub(arg1...)
	}
}

func (fake *Registerer) MustRegisterCallCount() int {
	fake.mustRegisterMutex.RLock()
	defer fake.mustRegisterMutex.RUnlock()
	return len(fake.mustRegisterArgsForCall)
}

func (fake *Registerer) MustRegisterArgsForCall(i int) []prometheus.Collector {
	fake.mustRegisterMutex.RLock()
	defer fake.mustRegisterMutex.RUnlock()
	return fake.mustRegisterArgsForCall[i].arg1
}

func (fake *Registerer) Unregister(arg1 prometheus.Collector) bool {
	fake.unregisterMutex.Lock()
	ret, specificReturn := fake.unregisterReturnsOnCall[len(fake.unregisterArgsForCall)]
	fake.unregisterArgsForCall = append(fake.unregisterArgsForCall, struct {
		arg1 prometheus.Collector
	}{arg1})
	fake.recordInvocation("Unregister", []interface{}{arg1})
	fake.unregisterMutex.Unlock()
	if fake.UnregisterStub != nil {
		return fake.UnregisterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unregisterReturns.result1
}

func (fake *Registerer) UnregisterCallCount() int {
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	return len(fake.unregisterArgsForCall)
}

func (fake *Registerer) UnregisterArgsForCall(i int) prometheus.Collector {
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	return fake.unregisterArgsForCall[i].arg1
}

func (fake *Registerer) UnregisterReturns(result1 bool) {
	fake.UnregisterStub = nil
	fake.unregisterReturns = struct {
		result1 bool
	}{result1}
}

func (fake *Registerer) UnregisterReturnsOnCall(i int, result1 bool) {
	fake.UnregisterStub = nil
	if fake.unregisterReturnsOnCall == nil {
		fake.unregisterReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.unregisterReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *Registerer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.mustRegisterMutex.RLock()
	defer fake.mustRegisterMutex.RUnlock()
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Registerer) recordInvocation(key string, args []interface{}) {
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