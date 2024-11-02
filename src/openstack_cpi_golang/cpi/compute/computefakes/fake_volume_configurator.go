// Code generated by counterfeiter. DO NOT EDIT.
package computefakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-openstack-cpi-release/src/openstack_cpi_golang/cpi/compute"
	"github.com/cloudfoundry/bosh-openstack-cpi-release/src/openstack_cpi_golang/cpi/config"
	"github.com/cloudfoundry/bosh-openstack-cpi-release/src/openstack_cpi_golang/cpi/properties"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
)

type FakeVolumeConfigurator struct {
	ConfigureVolumesStub        func(string, config.OpenstackConfig, properties.CreateVM, flavors.Flavor) ([]bootfromvolume.BlockDevice, error)
	configureVolumesMutex       sync.RWMutex
	configureVolumesArgsForCall []struct {
		arg1 string
		arg2 config.OpenstackConfig
		arg3 properties.CreateVM
		arg4 flavors.Flavor
	}
	configureVolumesReturns struct {
		result1 []bootfromvolume.BlockDevice
		result2 error
	}
	configureVolumesReturnsOnCall map[int]struct {
		result1 []bootfromvolume.BlockDevice
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeConfigurator) ConfigureVolumes(arg1 string, arg2 config.OpenstackConfig, arg3 properties.CreateVM, arg4 flavors.Flavor) ([]bootfromvolume.BlockDevice, error) {
	fake.configureVolumesMutex.Lock()
	ret, specificReturn := fake.configureVolumesReturnsOnCall[len(fake.configureVolumesArgsForCall)]
	fake.configureVolumesArgsForCall = append(fake.configureVolumesArgsForCall, struct {
		arg1 string
		arg2 config.OpenstackConfig
		arg3 properties.CreateVM
		arg4 flavors.Flavor
	}{arg1, arg2, arg3, arg4})
	stub := fake.ConfigureVolumesStub
	fakeReturns := fake.configureVolumesReturns
	fake.recordInvocation("ConfigureVolumes", []interface{}{arg1, arg2, arg3, arg4})
	fake.configureVolumesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolumeConfigurator) ConfigureVolumesCallCount() int {
	fake.configureVolumesMutex.RLock()
	defer fake.configureVolumesMutex.RUnlock()
	return len(fake.configureVolumesArgsForCall)
}

func (fake *FakeVolumeConfigurator) ConfigureVolumesCalls(stub func(string, config.OpenstackConfig, properties.CreateVM, flavors.Flavor) ([]bootfromvolume.BlockDevice, error)) {
	fake.configureVolumesMutex.Lock()
	defer fake.configureVolumesMutex.Unlock()
	fake.ConfigureVolumesStub = stub
}

func (fake *FakeVolumeConfigurator) ConfigureVolumesArgsForCall(i int) (string, config.OpenstackConfig, properties.CreateVM, flavors.Flavor) {
	fake.configureVolumesMutex.RLock()
	defer fake.configureVolumesMutex.RUnlock()
	argsForCall := fake.configureVolumesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeVolumeConfigurator) ConfigureVolumesReturns(result1 []bootfromvolume.BlockDevice, result2 error) {
	fake.configureVolumesMutex.Lock()
	defer fake.configureVolumesMutex.Unlock()
	fake.ConfigureVolumesStub = nil
	fake.configureVolumesReturns = struct {
		result1 []bootfromvolume.BlockDevice
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeConfigurator) ConfigureVolumesReturnsOnCall(i int, result1 []bootfromvolume.BlockDevice, result2 error) {
	fake.configureVolumesMutex.Lock()
	defer fake.configureVolumesMutex.Unlock()
	fake.ConfigureVolumesStub = nil
	if fake.configureVolumesReturnsOnCall == nil {
		fake.configureVolumesReturnsOnCall = make(map[int]struct {
			result1 []bootfromvolume.BlockDevice
			result2 error
		})
	}
	fake.configureVolumesReturnsOnCall[i] = struct {
		result1 []bootfromvolume.BlockDevice
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeConfigurator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configureVolumesMutex.RLock()
	defer fake.configureVolumesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolumeConfigurator) recordInvocation(key string, args []interface{}) {
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

var _ compute.VolumeConfigurator = new(FakeVolumeConfigurator)