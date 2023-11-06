/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14237
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

// VirtualizationBaseVirtualMachine Common attributes of a virtual machine managed by a  hypervisor. Serves as a base class for all concrete virtual machine types. A virtual machine (VM) is what a user and applications interact with. A VM usually runs a guest OS and applications run on the guest OS.
type VirtualizationBaseVirtualMachine struct {
	VirtualizationBaseSourceDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Time when this VM booted up.
	BootTime *time.Time                `json:"BootTime,omitempty"`
	Capacity NullableInfraHardwareInfo `json:"Capacity,omitempty"`
	// Average CPU utilization percentage derived as a ratio of CPU used to CPU allocated. The value is calculated whenever inventory is performed.
	CpuUtilization *float32                        `json:"CpuUtilization,omitempty"`
	GuestInfo      NullableVirtualizationGuestInfo `json:"GuestInfo,omitempty"`
	// Type of hypervisor where the virtual machine is hosted for example ESXi. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference).
	Identity       *string                              `json:"Identity,omitempty"`
	IpAddress      []string                             `json:"IpAddress,omitempty"`
	MemoryCapacity NullableVirtualizationMemoryCapacity `json:"MemoryCapacity,omitempty"`
	// Average memory utilization percentage derived as a ratio of memory used to available memory. The value is calculated whenever inventory is performed.
	MemoryUtilization *float32 `json:"MemoryUtilization,omitempty"`
	// User-provided name to identify the virtual machine.
	Name *string `json:"Name,omitempty"`
	// Power state of the virtual machine. * `Unknown` - The entity's power state is unknown. * `PoweringOn` - The entity is powering on. * `PoweredOn` - The entity is powered on. * `PoweringOff` - The entity is powering off. * `PoweredOff` - The entity is powered down. * `StandBy` - The entity is in standby mode. * `Paused` - The entity is in pause state. * `Rebooting` - The entity reboot is in progress. * `` - The entity's power state is not available.
	PowerState        *string                               `json:"PowerState,omitempty"`
	ProcessorCapacity NullableVirtualizationComputeCapacity `json:"ProcessorCapacity,omitempty"`
	// Cloud platform, where the virtual machine is launched. * `Unknown` - Cloud provider is not known. * `VMwarevSphere` - Cloud provider named VMware vSphere. * `AmazonWebServices` - Cloud provider named Amazon Web Services. * `MicrosoftAzure` - Cloud provider named Microsoft Azure. * `GoogleCloudPlatform` - Cloud provider named Google Cloud Platform.
	Provider *string `json:"Provider,omitempty"`
	// The current state of the virtual machine. For example, starting, stopped, etc. * `None` - A place holder for the default value. * `Creating` - Virtual machine creation is in progress. * `Pending` - The virtual machine is preparing to enter the started state. * `Starting` - The virtual machine is starting. * `Started` - The virtual machine is running and ready for use. * `Stopping` - The virtual machine is preparing to be stopped. * `Stopped` - The virtual machine is shut down and cannot be used. The virtual machine can be started again at any time. * `Pausing` - The virtual machine is preparing to be paused. * `Paused` - The virtual machine enters into paused state due to low free disk space. * `Suspending` - The virtual machine is preparing to be suspended. * `Suspended` - Virtual machine is in sleep mode.When a virtual machine is suspended, the current state of theoperating system, and applications is saved, and the virtual machine put into a suspended mode. * `Deleting` - The virtual machine is preparing to be terminated. * `Terminated` - The virtual machine has been permanently deleted and cannot be started. * `Rebooting` - The virtual machine reboot is in progress. * `Error` - The deployment of virtual machine is failed. * `Warning` - The virtual machine is in warning state.
	State *string `json:"State,omitempty"`
	// The uuid of this virtual machine. The uuid is internally generated and not user specified.
	Uuid *string `json:"Uuid,omitempty"`
	// Time when this virtualmachine is created.
	VmCreationTime       *time.Time `json:"VmCreationTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseVirtualMachine VirtualizationBaseVirtualMachine

// NewVirtualizationBaseVirtualMachine instantiates a new VirtualizationBaseVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVirtualMachine(classId string, objectType string) *VirtualizationBaseVirtualMachine {
	this := VirtualizationBaseVirtualMachine{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var powerState string = "Unknown"
	this.PowerState = &powerState
	var provider string = "Unknown"
	this.Provider = &provider
	var state string = "None"
	this.State = &state
	return &this
}

// NewVirtualizationBaseVirtualMachineWithDefaults instantiates a new VirtualizationBaseVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVirtualMachineWithDefaults() *VirtualizationBaseVirtualMachine {
	this := VirtualizationBaseVirtualMachine{}
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var powerState string = "Unknown"
	this.PowerState = &powerState
	var provider string = "Unknown"
	this.Provider = &provider
	var state string = "None"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseVirtualMachine) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseVirtualMachine) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseVirtualMachine) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseVirtualMachine) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootTime returns the BootTime field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetBootTime() time.Time {
	if o == nil || o.BootTime == nil {
		var ret time.Time
		return ret
	}
	return *o.BootTime
}

// GetBootTimeOk returns a tuple with the BootTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetBootTimeOk() (*time.Time, bool) {
	if o == nil || o.BootTime == nil {
		return nil, false
	}
	return o.BootTime, true
}

// HasBootTime returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasBootTime() bool {
	if o != nil && o.BootTime != nil {
		return true
	}

	return false
}

// SetBootTime gets a reference to the given time.Time and assigns it to the BootTime field.
func (o *VirtualizationBaseVirtualMachine) SetBootTime(v time.Time) {
	o.BootTime = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseVirtualMachine) GetCapacity() InfraHardwareInfo {
	if o == nil || o.Capacity.Get() == nil {
		var ret InfraHardwareInfo
		return ret
	}
	return *o.Capacity.Get()
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseVirtualMachine) GetCapacityOk() (*InfraHardwareInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capacity.Get(), o.Capacity.IsSet()
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasCapacity() bool {
	if o != nil && o.Capacity.IsSet() {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given NullableInfraHardwareInfo and assigns it to the Capacity field.
func (o *VirtualizationBaseVirtualMachine) SetCapacity(v InfraHardwareInfo) {
	o.Capacity.Set(&v)
}

// SetCapacityNil sets the value for Capacity to be an explicit nil
func (o *VirtualizationBaseVirtualMachine) SetCapacityNil() {
	o.Capacity.Set(nil)
}

// UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
func (o *VirtualizationBaseVirtualMachine) UnsetCapacity() {
	o.Capacity.Unset()
}

// GetCpuUtilization returns the CpuUtilization field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetCpuUtilization() float32 {
	if o == nil || o.CpuUtilization == nil {
		var ret float32
		return ret
	}
	return *o.CpuUtilization
}

// GetCpuUtilizationOk returns a tuple with the CpuUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetCpuUtilizationOk() (*float32, bool) {
	if o == nil || o.CpuUtilization == nil {
		return nil, false
	}
	return o.CpuUtilization, true
}

// HasCpuUtilization returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasCpuUtilization() bool {
	if o != nil && o.CpuUtilization != nil {
		return true
	}

	return false
}

// SetCpuUtilization gets a reference to the given float32 and assigns it to the CpuUtilization field.
func (o *VirtualizationBaseVirtualMachine) SetCpuUtilization(v float32) {
	o.CpuUtilization = &v
}

// GetGuestInfo returns the GuestInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseVirtualMachine) GetGuestInfo() VirtualizationGuestInfo {
	if o == nil || o.GuestInfo.Get() == nil {
		var ret VirtualizationGuestInfo
		return ret
	}
	return *o.GuestInfo.Get()
}

// GetGuestInfoOk returns a tuple with the GuestInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseVirtualMachine) GetGuestInfoOk() (*VirtualizationGuestInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuestInfo.Get(), o.GuestInfo.IsSet()
}

// HasGuestInfo returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasGuestInfo() bool {
	if o != nil && o.GuestInfo.IsSet() {
		return true
	}

	return false
}

// SetGuestInfo gets a reference to the given NullableVirtualizationGuestInfo and assigns it to the GuestInfo field.
func (o *VirtualizationBaseVirtualMachine) SetGuestInfo(v VirtualizationGuestInfo) {
	o.GuestInfo.Set(&v)
}

// SetGuestInfoNil sets the value for GuestInfo to be an explicit nil
func (o *VirtualizationBaseVirtualMachine) SetGuestInfoNil() {
	o.GuestInfo.Set(nil)
}

// UnsetGuestInfo ensures that no value is present for GuestInfo, not even an explicit nil
func (o *VirtualizationBaseVirtualMachine) UnsetGuestInfo() {
	o.GuestInfo.Unset()
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *VirtualizationBaseVirtualMachine) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseVirtualMachine) SetIdentity(v string) {
	o.Identity = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseVirtualMachine) GetIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseVirtualMachine) GetIpAddressOk() ([]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *VirtualizationBaseVirtualMachine) SetIpAddress(v []string) {
	o.IpAddress = v
}

// GetMemoryCapacity returns the MemoryCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseVirtualMachine) GetMemoryCapacity() VirtualizationMemoryCapacity {
	if o == nil || o.MemoryCapacity.Get() == nil {
		var ret VirtualizationMemoryCapacity
		return ret
	}
	return *o.MemoryCapacity.Get()
}

// GetMemoryCapacityOk returns a tuple with the MemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseVirtualMachine) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryCapacity.Get(), o.MemoryCapacity.IsSet()
}

// HasMemoryCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasMemoryCapacity() bool {
	if o != nil && o.MemoryCapacity.IsSet() {
		return true
	}

	return false
}

// SetMemoryCapacity gets a reference to the given NullableVirtualizationMemoryCapacity and assigns it to the MemoryCapacity field.
func (o *VirtualizationBaseVirtualMachine) SetMemoryCapacity(v VirtualizationMemoryCapacity) {
	o.MemoryCapacity.Set(&v)
}

// SetMemoryCapacityNil sets the value for MemoryCapacity to be an explicit nil
func (o *VirtualizationBaseVirtualMachine) SetMemoryCapacityNil() {
	o.MemoryCapacity.Set(nil)
}

// UnsetMemoryCapacity ensures that no value is present for MemoryCapacity, not even an explicit nil
func (o *VirtualizationBaseVirtualMachine) UnsetMemoryCapacity() {
	o.MemoryCapacity.Unset()
}

// GetMemoryUtilization returns the MemoryUtilization field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetMemoryUtilization() float32 {
	if o == nil || o.MemoryUtilization == nil {
		var ret float32
		return ret
	}
	return *o.MemoryUtilization
}

// GetMemoryUtilizationOk returns a tuple with the MemoryUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetMemoryUtilizationOk() (*float32, bool) {
	if o == nil || o.MemoryUtilization == nil {
		return nil, false
	}
	return o.MemoryUtilization, true
}

// HasMemoryUtilization returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasMemoryUtilization() bool {
	if o != nil && o.MemoryUtilization != nil {
		return true
	}

	return false
}

// SetMemoryUtilization gets a reference to the given float32 and assigns it to the MemoryUtilization field.
func (o *VirtualizationBaseVirtualMachine) SetMemoryUtilization(v float32) {
	o.MemoryUtilization = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseVirtualMachine) SetName(v string) {
	o.Name = &v
}

// GetPowerState returns the PowerState field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetPowerState() string {
	if o == nil || o.PowerState == nil {
		var ret string
		return ret
	}
	return *o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetPowerStateOk() (*string, bool) {
	if o == nil || o.PowerState == nil {
		return nil, false
	}
	return o.PowerState, true
}

// HasPowerState returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasPowerState() bool {
	if o != nil && o.PowerState != nil {
		return true
	}

	return false
}

// SetPowerState gets a reference to the given string and assigns it to the PowerState field.
func (o *VirtualizationBaseVirtualMachine) SetPowerState(v string) {
	o.PowerState = &v
}

// GetProcessorCapacity returns the ProcessorCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseVirtualMachine) GetProcessorCapacity() VirtualizationComputeCapacity {
	if o == nil || o.ProcessorCapacity.Get() == nil {
		var ret VirtualizationComputeCapacity
		return ret
	}
	return *o.ProcessorCapacity.Get()
}

// GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseVirtualMachine) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessorCapacity.Get(), o.ProcessorCapacity.IsSet()
}

// HasProcessorCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasProcessorCapacity() bool {
	if o != nil && o.ProcessorCapacity.IsSet() {
		return true
	}

	return false
}

// SetProcessorCapacity gets a reference to the given NullableVirtualizationComputeCapacity and assigns it to the ProcessorCapacity field.
func (o *VirtualizationBaseVirtualMachine) SetProcessorCapacity(v VirtualizationComputeCapacity) {
	o.ProcessorCapacity.Set(&v)
}

// SetProcessorCapacityNil sets the value for ProcessorCapacity to be an explicit nil
func (o *VirtualizationBaseVirtualMachine) SetProcessorCapacityNil() {
	o.ProcessorCapacity.Set(nil)
}

// UnsetProcessorCapacity ensures that no value is present for ProcessorCapacity, not even an explicit nil
func (o *VirtualizationBaseVirtualMachine) UnsetProcessorCapacity() {
	o.ProcessorCapacity.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *VirtualizationBaseVirtualMachine) SetProvider(v string) {
	o.Provider = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VirtualizationBaseVirtualMachine) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VirtualizationBaseVirtualMachine) SetUuid(v string) {
	o.Uuid = &v
}

// GetVmCreationTime returns the VmCreationTime field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachine) GetVmCreationTime() time.Time {
	if o == nil || o.VmCreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.VmCreationTime
}

// GetVmCreationTimeOk returns a tuple with the VmCreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachine) GetVmCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.VmCreationTime == nil {
		return nil, false
	}
	return o.VmCreationTime, true
}

// HasVmCreationTime returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachine) HasVmCreationTime() bool {
	if o != nil && o.VmCreationTime != nil {
		return true
	}

	return false
}

// SetVmCreationTime gets a reference to the given time.Time and assigns it to the VmCreationTime field.
func (o *VirtualizationBaseVirtualMachine) SetVmCreationTime(v time.Time) {
	o.VmCreationTime = &v
}

func (o VirtualizationBaseVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseSourceDevice, errVirtualizationBaseSourceDevice := json.Marshal(o.VirtualizationBaseSourceDevice)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	errVirtualizationBaseSourceDevice = json.Unmarshal([]byte(serializedVirtualizationBaseSourceDevice), &toSerialize)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BootTime != nil {
		toSerialize["BootTime"] = o.BootTime
	}
	if o.Capacity.IsSet() {
		toSerialize["Capacity"] = o.Capacity.Get()
	}
	if o.CpuUtilization != nil {
		toSerialize["CpuUtilization"] = o.CpuUtilization
	}
	if o.GuestInfo.IsSet() {
		toSerialize["GuestInfo"] = o.GuestInfo.Get()
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.MemoryCapacity.IsSet() {
		toSerialize["MemoryCapacity"] = o.MemoryCapacity.Get()
	}
	if o.MemoryUtilization != nil {
		toSerialize["MemoryUtilization"] = o.MemoryUtilization
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PowerState != nil {
		toSerialize["PowerState"] = o.PowerState
	}
	if o.ProcessorCapacity.IsSet() {
		toSerialize["ProcessorCapacity"] = o.ProcessorCapacity.Get()
	}
	if o.Provider != nil {
		toSerialize["Provider"] = o.Provider
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.VmCreationTime != nil {
		toSerialize["VmCreationTime"] = o.VmCreationTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseVirtualMachine) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationBaseVirtualMachineWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Time when this VM booted up.
		BootTime *time.Time                `json:"BootTime,omitempty"`
		Capacity NullableInfraHardwareInfo `json:"Capacity,omitempty"`
		// Average CPU utilization percentage derived as a ratio of CPU used to CPU allocated. The value is calculated whenever inventory is performed.
		CpuUtilization *float32                        `json:"CpuUtilization,omitempty"`
		GuestInfo      NullableVirtualizationGuestInfo `json:"GuestInfo,omitempty"`
		// Type of hypervisor where the virtual machine is hosted for example ESXi. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
		HypervisorType *string `json:"HypervisorType,omitempty"`
		// The internally generated identity of this VM. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference).
		Identity       *string                              `json:"Identity,omitempty"`
		IpAddress      []string                             `json:"IpAddress,omitempty"`
		MemoryCapacity NullableVirtualizationMemoryCapacity `json:"MemoryCapacity,omitempty"`
		// Average memory utilization percentage derived as a ratio of memory used to available memory. The value is calculated whenever inventory is performed.
		MemoryUtilization *float32 `json:"MemoryUtilization,omitempty"`
		// User-provided name to identify the virtual machine.
		Name *string `json:"Name,omitempty"`
		// Power state of the virtual machine. * `Unknown` - The entity's power state is unknown. * `PoweringOn` - The entity is powering on. * `PoweredOn` - The entity is powered on. * `PoweringOff` - The entity is powering off. * `PoweredOff` - The entity is powered down. * `StandBy` - The entity is in standby mode. * `Paused` - The entity is in pause state. * `Rebooting` - The entity reboot is in progress. * `` - The entity's power state is not available.
		PowerState        *string                               `json:"PowerState,omitempty"`
		ProcessorCapacity NullableVirtualizationComputeCapacity `json:"ProcessorCapacity,omitempty"`
		// Cloud platform, where the virtual machine is launched. * `Unknown` - Cloud provider is not known. * `VMwarevSphere` - Cloud provider named VMware vSphere. * `AmazonWebServices` - Cloud provider named Amazon Web Services. * `MicrosoftAzure` - Cloud provider named Microsoft Azure. * `GoogleCloudPlatform` - Cloud provider named Google Cloud Platform.
		Provider *string `json:"Provider,omitempty"`
		// The current state of the virtual machine. For example, starting, stopped, etc. * `None` - A place holder for the default value. * `Creating` - Virtual machine creation is in progress. * `Pending` - The virtual machine is preparing to enter the started state. * `Starting` - The virtual machine is starting. * `Started` - The virtual machine is running and ready for use. * `Stopping` - The virtual machine is preparing to be stopped. * `Stopped` - The virtual machine is shut down and cannot be used. The virtual machine can be started again at any time. * `Pausing` - The virtual machine is preparing to be paused. * `Paused` - The virtual machine enters into paused state due to low free disk space. * `Suspending` - The virtual machine is preparing to be suspended. * `Suspended` - Virtual machine is in sleep mode.When a virtual machine is suspended, the current state of theoperating system, and applications is saved, and the virtual machine put into a suspended mode. * `Deleting` - The virtual machine is preparing to be terminated. * `Terminated` - The virtual machine has been permanently deleted and cannot be started. * `Rebooting` - The virtual machine reboot is in progress. * `Error` - The deployment of virtual machine is failed. * `Warning` - The virtual machine is in warning state.
		State *string `json:"State,omitempty"`
		// The uuid of this virtual machine. The uuid is internally generated and not user specified.
		Uuid *string `json:"Uuid,omitempty"`
		// Time when this virtualmachine is created.
		VmCreationTime *time.Time `json:"VmCreationTime,omitempty"`
	}

	varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct := VirtualizationBaseVirtualMachineWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBaseVirtualMachine := _VirtualizationBaseVirtualMachine{}
		varVirtualizationBaseVirtualMachine.ClassId = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.ClassId
		varVirtualizationBaseVirtualMachine.ObjectType = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.ObjectType
		varVirtualizationBaseVirtualMachine.BootTime = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.BootTime
		varVirtualizationBaseVirtualMachine.Capacity = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.Capacity
		varVirtualizationBaseVirtualMachine.CpuUtilization = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.CpuUtilization
		varVirtualizationBaseVirtualMachine.GuestInfo = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.GuestInfo
		varVirtualizationBaseVirtualMachine.HypervisorType = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.HypervisorType
		varVirtualizationBaseVirtualMachine.Identity = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.Identity
		varVirtualizationBaseVirtualMachine.IpAddress = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.IpAddress
		varVirtualizationBaseVirtualMachine.MemoryCapacity = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.MemoryCapacity
		varVirtualizationBaseVirtualMachine.MemoryUtilization = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.MemoryUtilization
		varVirtualizationBaseVirtualMachine.Name = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.Name
		varVirtualizationBaseVirtualMachine.PowerState = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.PowerState
		varVirtualizationBaseVirtualMachine.ProcessorCapacity = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.ProcessorCapacity
		varVirtualizationBaseVirtualMachine.Provider = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.Provider
		varVirtualizationBaseVirtualMachine.State = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.State
		varVirtualizationBaseVirtualMachine.Uuid = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.Uuid
		varVirtualizationBaseVirtualMachine.VmCreationTime = varVirtualizationBaseVirtualMachineWithoutEmbeddedStruct.VmCreationTime
		*o = VirtualizationBaseVirtualMachine(varVirtualizationBaseVirtualMachine)
	} else {
		return err
	}

	varVirtualizationBaseVirtualMachine := _VirtualizationBaseVirtualMachine{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualMachine)
	if err == nil {
		o.VirtualizationBaseSourceDevice = varVirtualizationBaseVirtualMachine.VirtualizationBaseSourceDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BootTime")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "CpuUtilization")
		delete(additionalProperties, "GuestInfo")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "MemoryCapacity")
		delete(additionalProperties, "MemoryUtilization")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PowerState")
		delete(additionalProperties, "ProcessorCapacity")
		delete(additionalProperties, "Provider")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VmCreationTime")

		// remove fields from embedded structs
		reflectVirtualizationBaseSourceDevice := reflect.ValueOf(o.VirtualizationBaseSourceDevice)
		for i := 0; i < reflectVirtualizationBaseSourceDevice.Type().NumField(); i++ {
			t := reflectVirtualizationBaseSourceDevice.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationBaseVirtualMachine struct {
	value *VirtualizationBaseVirtualMachine
	isSet bool
}

func (v NullableVirtualizationBaseVirtualMachine) Get() *VirtualizationBaseVirtualMachine {
	return v.value
}

func (v *NullableVirtualizationBaseVirtualMachine) Set(val *VirtualizationBaseVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVirtualMachine(val *VirtualizationBaseVirtualMachine) *NullableVirtualizationBaseVirtualMachine {
	return &NullableVirtualizationBaseVirtualMachine{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
