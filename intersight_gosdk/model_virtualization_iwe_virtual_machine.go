/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VirtualizationIweVirtualMachine The Virtual machine that runs on a Intersight Workload Engine compute host.
type VirtualizationIweVirtualMachine struct {
	VirtualizationBaseVirtualMachine
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string          `json:"ObjectType"`
	AffinitySelectors []InfraMetaData `json:"AffinitySelectors,omitempty"`
	// Denotes age or life time of the VM in nano seconds.
	Age                   *string                `json:"Age,omitempty"`
	AntiAffinitySelectors []InfraMetaData        `json:"AntiAffinitySelectors,omitempty"`
	Disks                 []VirtualizationVmDisk `json:"Disks,omitempty"`
	// Reason of the failure when VM is in failed state.
	FailureReason *string                     `json:"FailureReason,omitempty"`
	Interfaces    []VirtualizationVmInterface `json:"Interfaces,omitempty"`
	Labels        []InfraMetaData             `json:"Labels,omitempty"`
	// Number network interfaces associated with the virtual machine.
	NetworkCount *int64 `json:"NetworkCount,omitempty"`
	// Denotes the VM start timestamp.
	StartTime *string `json:"StartTime,omitempty"`
	// Status of virtual machine. * `Unknown` - Virtual machine state is not available. * `Running` - Virtual machine is running normally. * `Stopped` - Virtual machine has been stopped. * `WaitForLaunch` - Virtual machine is wating to be launched. * `Paused` - Virtual machine is currently paused. * `ImportInProgress` - Virtual machine image is being imported into the platform. * `ImportFailed` - Virtual machine image import operation failed. * `DiskCloneInProgress` - Disk clone operation for the virtual machine is in progress. * `DiskCloneFailed` - Disk clone operation for the virtual machine failed. * `DiskCreateInProgress` - Disk create operation for the virtual machine is in progress. * `DiskCreateFailed` - Disk create operation for the virtual machine failed. * `Processing` - Virtual machine is being created. * `UnSchedulable` - Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications. * `Failed` - Some virtual machine operation has failed. More information is available as part of the results of the operation. * `Warning` - CPU/Memory utilisation for the virtual machine has crossed the threshold value. * `` - Virtual machine status is not available.
	Status               *string                               `json:"Status,omitempty"`
	Cluster              *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
	Host                 *VirtualizationIweHostRelationship    `json:"Host,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationIweVirtualMachine VirtualizationIweVirtualMachine

// NewVirtualizationIweVirtualMachine instantiates a new VirtualizationIweVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIweVirtualMachine(classId string, objectType string) *VirtualizationIweVirtualMachine {
	this := VirtualizationIweVirtualMachine{}
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
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewVirtualizationIweVirtualMachineWithDefaults instantiates a new VirtualizationIweVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIweVirtualMachineWithDefaults() *VirtualizationIweVirtualMachine {
	this := VirtualizationIweVirtualMachine{}
	var classId string = "virtualization.IweVirtualMachine"
	this.ClassId = classId
	var objectType string = "virtualization.IweVirtualMachine"
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIweVirtualMachine) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachine) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIweVirtualMachine) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIweVirtualMachine) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachine) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIweVirtualMachine) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAffinitySelectors returns the AffinitySelectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweVirtualMachine) GetAffinitySelectors() []InfraMetaData {
	if o == nil {
		var ret []InfraMetaData
		return ret
	}
	return o.AffinitySelectors
}

// GetAffinitySelectorsOk returns a tuple with the AffinitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweVirtualMachine) GetAffinitySelectorsOk() (*[]InfraMetaData, bool) {
	if o == nil || o.AffinitySelectors == nil {
		return nil, false
	}
	return &o.AffinitySelectors, true
}

// HasAffinitySelectors returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasAffinitySelectors() bool {
	if o != nil && o.AffinitySelectors != nil {
		return true
	}

	return false
}

// SetAffinitySelectors gets a reference to the given []InfraMetaData and assigns it to the AffinitySelectors field.
func (o *VirtualizationIweVirtualMachine) SetAffinitySelectors(v []InfraMetaData) {
	o.AffinitySelectors = v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachine) GetAge() string {
	if o == nil || o.Age == nil {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachine) GetAgeOk() (*string, bool) {
	if o == nil || o.Age == nil {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasAge() bool {
	if o != nil && o.Age != nil {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *VirtualizationIweVirtualMachine) SetAge(v string) {
	o.Age = &v
}

// GetAntiAffinitySelectors returns the AntiAffinitySelectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweVirtualMachine) GetAntiAffinitySelectors() []InfraMetaData {
	if o == nil {
		var ret []InfraMetaData
		return ret
	}
	return o.AntiAffinitySelectors
}

// GetAntiAffinitySelectorsOk returns a tuple with the AntiAffinitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweVirtualMachine) GetAntiAffinitySelectorsOk() (*[]InfraMetaData, bool) {
	if o == nil || o.AntiAffinitySelectors == nil {
		return nil, false
	}
	return &o.AntiAffinitySelectors, true
}

// HasAntiAffinitySelectors returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasAntiAffinitySelectors() bool {
	if o != nil && o.AntiAffinitySelectors != nil {
		return true
	}

	return false
}

// SetAntiAffinitySelectors gets a reference to the given []InfraMetaData and assigns it to the AntiAffinitySelectors field.
func (o *VirtualizationIweVirtualMachine) SetAntiAffinitySelectors(v []InfraMetaData) {
	o.AntiAffinitySelectors = v
}

// GetDisks returns the Disks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweVirtualMachine) GetDisks() []VirtualizationVmDisk {
	if o == nil {
		var ret []VirtualizationVmDisk
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweVirtualMachine) GetDisksOk() (*[]VirtualizationVmDisk, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return &o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []VirtualizationVmDisk and assigns it to the Disks field.
func (o *VirtualizationIweVirtualMachine) SetDisks(v []VirtualizationVmDisk) {
	o.Disks = v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachine) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachine) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *VirtualizationIweVirtualMachine) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweVirtualMachine) GetInterfaces() []VirtualizationVmInterface {
	if o == nil {
		var ret []VirtualizationVmInterface
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweVirtualMachine) GetInterfacesOk() (*[]VirtualizationVmInterface, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return &o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []VirtualizationVmInterface and assigns it to the Interfaces field.
func (o *VirtualizationIweVirtualMachine) SetInterfaces(v []VirtualizationVmInterface) {
	o.Interfaces = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweVirtualMachine) GetLabels() []InfraMetaData {
	if o == nil {
		var ret []InfraMetaData
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweVirtualMachine) GetLabelsOk() (*[]InfraMetaData, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []InfraMetaData and assigns it to the Labels field.
func (o *VirtualizationIweVirtualMachine) SetLabels(v []InfraMetaData) {
	o.Labels = v
}

// GetNetworkCount returns the NetworkCount field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachine) GetNetworkCount() int64 {
	if o == nil || o.NetworkCount == nil {
		var ret int64
		return ret
	}
	return *o.NetworkCount
}

// GetNetworkCountOk returns a tuple with the NetworkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachine) GetNetworkCountOk() (*int64, bool) {
	if o == nil || o.NetworkCount == nil {
		return nil, false
	}
	return o.NetworkCount, true
}

// HasNetworkCount returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasNetworkCount() bool {
	if o != nil && o.NetworkCount != nil {
		return true
	}

	return false
}

// SetNetworkCount gets a reference to the given int64 and assigns it to the NetworkCount field.
func (o *VirtualizationIweVirtualMachine) SetNetworkCount(v int64) {
	o.NetworkCount = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachine) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachine) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *VirtualizationIweVirtualMachine) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachine) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachine) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VirtualizationIweVirtualMachine) SetStatus(v string) {
	o.Status = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachine) GetCluster() VirtualizationIweClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationIweClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachine) GetClusterOk() (*VirtualizationIweClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationIweClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationIweVirtualMachine) SetCluster(v VirtualizationIweClusterRelationship) {
	o.Cluster = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachine) GetHost() VirtualizationIweHostRelationship {
	if o == nil || o.Host == nil {
		var ret VirtualizationIweHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachine) GetHostOk() (*VirtualizationIweHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachine) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given VirtualizationIweHostRelationship and assigns it to the Host field.
func (o *VirtualizationIweVirtualMachine) SetHost(v VirtualizationIweHostRelationship) {
	o.Host = &v
}

func (o VirtualizationIweVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualMachine, errVirtualizationBaseVirtualMachine := json.Marshal(o.VirtualizationBaseVirtualMachine)
	if errVirtualizationBaseVirtualMachine != nil {
		return []byte{}, errVirtualizationBaseVirtualMachine
	}
	errVirtualizationBaseVirtualMachine = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualMachine), &toSerialize)
	if errVirtualizationBaseVirtualMachine != nil {
		return []byte{}, errVirtualizationBaseVirtualMachine
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AffinitySelectors != nil {
		toSerialize["AffinitySelectors"] = o.AffinitySelectors
	}
	if o.Age != nil {
		toSerialize["Age"] = o.Age
	}
	if o.AntiAffinitySelectors != nil {
		toSerialize["AntiAffinitySelectors"] = o.AntiAffinitySelectors
	}
	if o.Disks != nil {
		toSerialize["Disks"] = o.Disks
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.Interfaces != nil {
		toSerialize["Interfaces"] = o.Interfaces
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.NetworkCount != nil {
		toSerialize["NetworkCount"] = o.NetworkCount
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIweVirtualMachine) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationIweVirtualMachineWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string          `json:"ObjectType"`
		AffinitySelectors []InfraMetaData `json:"AffinitySelectors,omitempty"`
		// Denotes age or life time of the VM in nano seconds.
		Age                   *string                `json:"Age,omitempty"`
		AntiAffinitySelectors []InfraMetaData        `json:"AntiAffinitySelectors,omitempty"`
		Disks                 []VirtualizationVmDisk `json:"Disks,omitempty"`
		// Reason of the failure when VM is in failed state.
		FailureReason *string                     `json:"FailureReason,omitempty"`
		Interfaces    []VirtualizationVmInterface `json:"Interfaces,omitempty"`
		Labels        []InfraMetaData             `json:"Labels,omitempty"`
		// Number network interfaces associated with the virtual machine.
		NetworkCount *int64 `json:"NetworkCount,omitempty"`
		// Denotes the VM start timestamp.
		StartTime *string `json:"StartTime,omitempty"`
		// Status of virtual machine. * `Unknown` - Virtual machine state is not available. * `Running` - Virtual machine is running normally. * `Stopped` - Virtual machine has been stopped. * `WaitForLaunch` - Virtual machine is wating to be launched. * `Paused` - Virtual machine is currently paused. * `ImportInProgress` - Virtual machine image is being imported into the platform. * `ImportFailed` - Virtual machine image import operation failed. * `DiskCloneInProgress` - Disk clone operation for the virtual machine is in progress. * `DiskCloneFailed` - Disk clone operation for the virtual machine failed. * `DiskCreateInProgress` - Disk create operation for the virtual machine is in progress. * `DiskCreateFailed` - Disk create operation for the virtual machine failed. * `Processing` - Virtual machine is being created. * `UnSchedulable` - Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications. * `Failed` - Some virtual machine operation has failed. More information is available as part of the results of the operation. * `Warning` - CPU/Memory utilisation for the virtual machine has crossed the threshold value. * `` - Virtual machine status is not available.
		Status  *string                               `json:"Status,omitempty"`
		Cluster *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
		Host    *VirtualizationIweHostRelationship    `json:"Host,omitempty"`
	}

	varVirtualizationIweVirtualMachineWithoutEmbeddedStruct := VirtualizationIweVirtualMachineWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationIweVirtualMachineWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationIweVirtualMachine := _VirtualizationIweVirtualMachine{}
		varVirtualizationIweVirtualMachine.ClassId = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.ClassId
		varVirtualizationIweVirtualMachine.ObjectType = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.ObjectType
		varVirtualizationIweVirtualMachine.AffinitySelectors = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.AffinitySelectors
		varVirtualizationIweVirtualMachine.Age = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.Age
		varVirtualizationIweVirtualMachine.AntiAffinitySelectors = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.AntiAffinitySelectors
		varVirtualizationIweVirtualMachine.Disks = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.Disks
		varVirtualizationIweVirtualMachine.FailureReason = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.FailureReason
		varVirtualizationIweVirtualMachine.Interfaces = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.Interfaces
		varVirtualizationIweVirtualMachine.Labels = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.Labels
		varVirtualizationIweVirtualMachine.NetworkCount = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.NetworkCount
		varVirtualizationIweVirtualMachine.StartTime = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.StartTime
		varVirtualizationIweVirtualMachine.Status = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.Status
		varVirtualizationIweVirtualMachine.Cluster = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.Cluster
		varVirtualizationIweVirtualMachine.Host = varVirtualizationIweVirtualMachineWithoutEmbeddedStruct.Host
		*o = VirtualizationIweVirtualMachine(varVirtualizationIweVirtualMachine)
	} else {
		return err
	}

	varVirtualizationIweVirtualMachine := _VirtualizationIweVirtualMachine{}

	err = json.Unmarshal(bytes, &varVirtualizationIweVirtualMachine)
	if err == nil {
		o.VirtualizationBaseVirtualMachine = varVirtualizationIweVirtualMachine.VirtualizationBaseVirtualMachine
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AffinitySelectors")
		delete(additionalProperties, "Age")
		delete(additionalProperties, "AntiAffinitySelectors")
		delete(additionalProperties, "Disks")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "Interfaces")
		delete(additionalProperties, "Labels")
		delete(additionalProperties, "NetworkCount")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Host")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualMachine := reflect.ValueOf(o.VirtualizationBaseVirtualMachine)
		for i := 0; i < reflectVirtualizationBaseVirtualMachine.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualMachine.Type().Field(i)

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

type NullableVirtualizationIweVirtualMachine struct {
	value *VirtualizationIweVirtualMachine
	isSet bool
}

func (v NullableVirtualizationIweVirtualMachine) Get() *VirtualizationIweVirtualMachine {
	return v.value
}

func (v *NullableVirtualizationIweVirtualMachine) Set(val *VirtualizationIweVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIweVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIweVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIweVirtualMachine(val *VirtualizationIweVirtualMachine) *NullableVirtualizationIweVirtualMachine {
	return &NullableVirtualizationIweVirtualMachine{value: val, isSet: true}
}

func (v NullableVirtualizationIweVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIweVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
