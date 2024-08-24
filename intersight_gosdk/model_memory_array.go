/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18012
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the MemoryArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemoryArray{}

// MemoryArray Holder housing multiple memory units.
type MemoryArray struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The instance number of the memory array.
	ArrayId *int64 `json:"ArrayId,omitempty"`
	// ID of the CPU that access this memory array.
	CpuId *int64 `json:"CpuId,omitempty"`
	// Current capacity of all the memory units on a server.
	CurrentCapacity *string `json:"CurrentCapacity,omitempty"`
	// The primary hardware error detection or correction method supported by the memory array.
	ErrorCorrection *string `json:"ErrorCorrection,omitempty"`
	// Maximum capacity of all the memory units on a server.
	MaxCapacity *string `json:"MaxCapacity,omitempty"`
	// The maximum number of slots or sockets available for memory devices in the memory array.
	MaxDevices *string `json:"MaxDevices,omitempty"`
	// The power state indicator of the memory array.
	OperPowerState      *string                                 `json:"OperPowerState,omitempty"`
	ComputeBlade        NullableComputeBladeRelationship        `json:"ComputeBlade,omitempty"`
	ComputeBoard        NullableComputeBoardRelationship        `json:"ComputeBoard,omitempty"`
	ComputeRackUnit     NullableComputeRackUnitRelationship     `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo NullableInventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to memoryPersistentMemoryUnit resources.
	PersistentMemoryUnits []MemoryPersistentMemoryUnitRelationship    `json:"PersistentMemoryUnits,omitempty"`
	RegisteredDevice      NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to memoryUnit resources.
	Units                []MemoryUnitRelationship `json:"Units,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryArray MemoryArray

// NewMemoryArray instantiates a new MemoryArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryArray(classId string, objectType string) *MemoryArray {
	this := MemoryArray{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryArrayWithDefaults instantiates a new MemoryArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryArrayWithDefaults() *MemoryArray {
	this := MemoryArray{}
	var classId string = "memory.Array"
	this.ClassId = classId
	var objectType string = "memory.Array"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryArray) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryArray) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryArray) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "memory.Array" of the ClassId field.
func (o *MemoryArray) GetDefaultClassId() interface{} {
	return "memory.Array"
}

// GetObjectType returns the ObjectType field value
func (o *MemoryArray) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryArray) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryArray) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "memory.Array" of the ObjectType field.
func (o *MemoryArray) GetDefaultObjectType() interface{} {
	return "memory.Array"
}

// GetArrayId returns the ArrayId field value if set, zero value otherwise.
func (o *MemoryArray) GetArrayId() int64 {
	if o == nil || IsNil(o.ArrayId) {
		var ret int64
		return ret
	}
	return *o.ArrayId
}

// GetArrayIdOk returns a tuple with the ArrayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArray) GetArrayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ArrayId) {
		return nil, false
	}
	return o.ArrayId, true
}

// HasArrayId returns a boolean if a field has been set.
func (o *MemoryArray) HasArrayId() bool {
	if o != nil && !IsNil(o.ArrayId) {
		return true
	}

	return false
}

// SetArrayId gets a reference to the given int64 and assigns it to the ArrayId field.
func (o *MemoryArray) SetArrayId(v int64) {
	o.ArrayId = &v
}

// GetCpuId returns the CpuId field value if set, zero value otherwise.
func (o *MemoryArray) GetCpuId() int64 {
	if o == nil || IsNil(o.CpuId) {
		var ret int64
		return ret
	}
	return *o.CpuId
}

// GetCpuIdOk returns a tuple with the CpuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArray) GetCpuIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CpuId) {
		return nil, false
	}
	return o.CpuId, true
}

// HasCpuId returns a boolean if a field has been set.
func (o *MemoryArray) HasCpuId() bool {
	if o != nil && !IsNil(o.CpuId) {
		return true
	}

	return false
}

// SetCpuId gets a reference to the given int64 and assigns it to the CpuId field.
func (o *MemoryArray) SetCpuId(v int64) {
	o.CpuId = &v
}

// GetCurrentCapacity returns the CurrentCapacity field value if set, zero value otherwise.
func (o *MemoryArray) GetCurrentCapacity() string {
	if o == nil || IsNil(o.CurrentCapacity) {
		var ret string
		return ret
	}
	return *o.CurrentCapacity
}

// GetCurrentCapacityOk returns a tuple with the CurrentCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArray) GetCurrentCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentCapacity) {
		return nil, false
	}
	return o.CurrentCapacity, true
}

// HasCurrentCapacity returns a boolean if a field has been set.
func (o *MemoryArray) HasCurrentCapacity() bool {
	if o != nil && !IsNil(o.CurrentCapacity) {
		return true
	}

	return false
}

// SetCurrentCapacity gets a reference to the given string and assigns it to the CurrentCapacity field.
func (o *MemoryArray) SetCurrentCapacity(v string) {
	o.CurrentCapacity = &v
}

// GetErrorCorrection returns the ErrorCorrection field value if set, zero value otherwise.
func (o *MemoryArray) GetErrorCorrection() string {
	if o == nil || IsNil(o.ErrorCorrection) {
		var ret string
		return ret
	}
	return *o.ErrorCorrection
}

// GetErrorCorrectionOk returns a tuple with the ErrorCorrection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArray) GetErrorCorrectionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCorrection) {
		return nil, false
	}
	return o.ErrorCorrection, true
}

// HasErrorCorrection returns a boolean if a field has been set.
func (o *MemoryArray) HasErrorCorrection() bool {
	if o != nil && !IsNil(o.ErrorCorrection) {
		return true
	}

	return false
}

// SetErrorCorrection gets a reference to the given string and assigns it to the ErrorCorrection field.
func (o *MemoryArray) SetErrorCorrection(v string) {
	o.ErrorCorrection = &v
}

// GetMaxCapacity returns the MaxCapacity field value if set, zero value otherwise.
func (o *MemoryArray) GetMaxCapacity() string {
	if o == nil || IsNil(o.MaxCapacity) {
		var ret string
		return ret
	}
	return *o.MaxCapacity
}

// GetMaxCapacityOk returns a tuple with the MaxCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArray) GetMaxCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCapacity) {
		return nil, false
	}
	return o.MaxCapacity, true
}

// HasMaxCapacity returns a boolean if a field has been set.
func (o *MemoryArray) HasMaxCapacity() bool {
	if o != nil && !IsNil(o.MaxCapacity) {
		return true
	}

	return false
}

// SetMaxCapacity gets a reference to the given string and assigns it to the MaxCapacity field.
func (o *MemoryArray) SetMaxCapacity(v string) {
	o.MaxCapacity = &v
}

// GetMaxDevices returns the MaxDevices field value if set, zero value otherwise.
func (o *MemoryArray) GetMaxDevices() string {
	if o == nil || IsNil(o.MaxDevices) {
		var ret string
		return ret
	}
	return *o.MaxDevices
}

// GetMaxDevicesOk returns a tuple with the MaxDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArray) GetMaxDevicesOk() (*string, bool) {
	if o == nil || IsNil(o.MaxDevices) {
		return nil, false
	}
	return o.MaxDevices, true
}

// HasMaxDevices returns a boolean if a field has been set.
func (o *MemoryArray) HasMaxDevices() bool {
	if o != nil && !IsNil(o.MaxDevices) {
		return true
	}

	return false
}

// SetMaxDevices gets a reference to the given string and assigns it to the MaxDevices field.
func (o *MemoryArray) SetMaxDevices(v string) {
	o.MaxDevices = &v
}

// GetOperPowerState returns the OperPowerState field value if set, zero value otherwise.
func (o *MemoryArray) GetOperPowerState() string {
	if o == nil || IsNil(o.OperPowerState) {
		var ret string
		return ret
	}
	return *o.OperPowerState
}

// GetOperPowerStateOk returns a tuple with the OperPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArray) GetOperPowerStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperPowerState) {
		return nil, false
	}
	return o.OperPowerState, true
}

// HasOperPowerState returns a boolean if a field has been set.
func (o *MemoryArray) HasOperPowerState() bool {
	if o != nil && !IsNil(o.OperPowerState) {
		return true
	}

	return false
}

// SetOperPowerState gets a reference to the given string and assigns it to the OperPowerState field.
func (o *MemoryArray) SetOperPowerState(v string) {
	o.OperPowerState = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryArray) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || IsNil(o.ComputeBlade.Get()) {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade.Get()
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryArray) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBlade.Get(), o.ComputeBlade.IsSet()
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *MemoryArray) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade.IsSet() {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given NullableComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *MemoryArray) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade.Set(&v)
}

// SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil
func (o *MemoryArray) SetComputeBladeNil() {
	o.ComputeBlade.Set(nil)
}

// UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
func (o *MemoryArray) UnsetComputeBlade() {
	o.ComputeBlade.Unset()
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryArray) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || IsNil(o.ComputeBoard.Get()) {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard.Get()
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryArray) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBoard.Get(), o.ComputeBoard.IsSet()
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *MemoryArray) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard.IsSet() {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given NullableComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *MemoryArray) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard.Set(&v)
}

// SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil
func (o *MemoryArray) SetComputeBoardNil() {
	o.ComputeBoard.Set(nil)
}

// UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
func (o *MemoryArray) UnsetComputeBoard() {
	o.ComputeBoard.Unset()
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryArray) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || IsNil(o.ComputeRackUnit.Get()) {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit.Get()
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryArray) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeRackUnit.Get(), o.ComputeRackUnit.IsSet()
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *MemoryArray) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit.IsSet() {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given NullableComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *MemoryArray) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit.Set(&v)
}

// SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil
func (o *MemoryArray) SetComputeRackUnitNil() {
	o.ComputeRackUnit.Set(nil)
}

// UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
func (o *MemoryArray) UnsetComputeRackUnit() {
	o.ComputeRackUnit.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryArray) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryArray) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryArray) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryArray) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *MemoryArray) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *MemoryArray) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetPersistentMemoryUnits returns the PersistentMemoryUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryArray) GetPersistentMemoryUnits() []MemoryPersistentMemoryUnitRelationship {
	if o == nil {
		var ret []MemoryPersistentMemoryUnitRelationship
		return ret
	}
	return o.PersistentMemoryUnits
}

// GetPersistentMemoryUnitsOk returns a tuple with the PersistentMemoryUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryArray) GetPersistentMemoryUnitsOk() ([]MemoryPersistentMemoryUnitRelationship, bool) {
	if o == nil || IsNil(o.PersistentMemoryUnits) {
		return nil, false
	}
	return o.PersistentMemoryUnits, true
}

// HasPersistentMemoryUnits returns a boolean if a field has been set.
func (o *MemoryArray) HasPersistentMemoryUnits() bool {
	if o != nil && !IsNil(o.PersistentMemoryUnits) {
		return true
	}

	return false
}

// SetPersistentMemoryUnits gets a reference to the given []MemoryPersistentMemoryUnitRelationship and assigns it to the PersistentMemoryUnits field.
func (o *MemoryArray) SetPersistentMemoryUnits(v []MemoryPersistentMemoryUnitRelationship) {
	o.PersistentMemoryUnits = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryArray) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryArray) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryArray) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryArray) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *MemoryArray) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *MemoryArray) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetUnits returns the Units field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryArray) GetUnits() []MemoryUnitRelationship {
	if o == nil {
		var ret []MemoryUnitRelationship
		return ret
	}
	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryArray) GetUnitsOk() ([]MemoryUnitRelationship, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *MemoryArray) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given []MemoryUnitRelationship and assigns it to the Units field.
func (o *MemoryArray) SetUnits(v []MemoryUnitRelationship) {
	o.Units = v
}

func (o MemoryArray) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemoryArray) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ArrayId) {
		toSerialize["ArrayId"] = o.ArrayId
	}
	if !IsNil(o.CpuId) {
		toSerialize["CpuId"] = o.CpuId
	}
	if !IsNil(o.CurrentCapacity) {
		toSerialize["CurrentCapacity"] = o.CurrentCapacity
	}
	if !IsNil(o.ErrorCorrection) {
		toSerialize["ErrorCorrection"] = o.ErrorCorrection
	}
	if !IsNil(o.MaxCapacity) {
		toSerialize["MaxCapacity"] = o.MaxCapacity
	}
	if !IsNil(o.MaxDevices) {
		toSerialize["MaxDevices"] = o.MaxDevices
	}
	if !IsNil(o.OperPowerState) {
		toSerialize["OperPowerState"] = o.OperPowerState
	}
	if o.ComputeBlade.IsSet() {
		toSerialize["ComputeBlade"] = o.ComputeBlade.Get()
	}
	if o.ComputeBoard.IsSet() {
		toSerialize["ComputeBoard"] = o.ComputeBoard.Get()
	}
	if o.ComputeRackUnit.IsSet() {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.PersistentMemoryUnits != nil {
		toSerialize["PersistentMemoryUnits"] = o.PersistentMemoryUnits
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.Units != nil {
		toSerialize["Units"] = o.Units
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemoryArray) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type MemoryArrayWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The instance number of the memory array.
		ArrayId *int64 `json:"ArrayId,omitempty"`
		// ID of the CPU that access this memory array.
		CpuId *int64 `json:"CpuId,omitempty"`
		// Current capacity of all the memory units on a server.
		CurrentCapacity *string `json:"CurrentCapacity,omitempty"`
		// The primary hardware error detection or correction method supported by the memory array.
		ErrorCorrection *string `json:"ErrorCorrection,omitempty"`
		// Maximum capacity of all the memory units on a server.
		MaxCapacity *string `json:"MaxCapacity,omitempty"`
		// The maximum number of slots or sockets available for memory devices in the memory array.
		MaxDevices *string `json:"MaxDevices,omitempty"`
		// The power state indicator of the memory array.
		OperPowerState      *string                                 `json:"OperPowerState,omitempty"`
		ComputeBlade        NullableComputeBladeRelationship        `json:"ComputeBlade,omitempty"`
		ComputeBoard        NullableComputeBoardRelationship        `json:"ComputeBoard,omitempty"`
		ComputeRackUnit     NullableComputeRackUnitRelationship     `json:"ComputeRackUnit,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		// An array of relationships to memoryPersistentMemoryUnit resources.
		PersistentMemoryUnits []MemoryPersistentMemoryUnitRelationship    `json:"PersistentMemoryUnits,omitempty"`
		RegisteredDevice      NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to memoryUnit resources.
		Units []MemoryUnitRelationship `json:"Units,omitempty"`
	}

	varMemoryArrayWithoutEmbeddedStruct := MemoryArrayWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMemoryArrayWithoutEmbeddedStruct)
	if err == nil {
		varMemoryArray := _MemoryArray{}
		varMemoryArray.ClassId = varMemoryArrayWithoutEmbeddedStruct.ClassId
		varMemoryArray.ObjectType = varMemoryArrayWithoutEmbeddedStruct.ObjectType
		varMemoryArray.ArrayId = varMemoryArrayWithoutEmbeddedStruct.ArrayId
		varMemoryArray.CpuId = varMemoryArrayWithoutEmbeddedStruct.CpuId
		varMemoryArray.CurrentCapacity = varMemoryArrayWithoutEmbeddedStruct.CurrentCapacity
		varMemoryArray.ErrorCorrection = varMemoryArrayWithoutEmbeddedStruct.ErrorCorrection
		varMemoryArray.MaxCapacity = varMemoryArrayWithoutEmbeddedStruct.MaxCapacity
		varMemoryArray.MaxDevices = varMemoryArrayWithoutEmbeddedStruct.MaxDevices
		varMemoryArray.OperPowerState = varMemoryArrayWithoutEmbeddedStruct.OperPowerState
		varMemoryArray.ComputeBlade = varMemoryArrayWithoutEmbeddedStruct.ComputeBlade
		varMemoryArray.ComputeBoard = varMemoryArrayWithoutEmbeddedStruct.ComputeBoard
		varMemoryArray.ComputeRackUnit = varMemoryArrayWithoutEmbeddedStruct.ComputeRackUnit
		varMemoryArray.InventoryDeviceInfo = varMemoryArrayWithoutEmbeddedStruct.InventoryDeviceInfo
		varMemoryArray.PersistentMemoryUnits = varMemoryArrayWithoutEmbeddedStruct.PersistentMemoryUnits
		varMemoryArray.RegisteredDevice = varMemoryArrayWithoutEmbeddedStruct.RegisteredDevice
		varMemoryArray.Units = varMemoryArrayWithoutEmbeddedStruct.Units
		*o = MemoryArray(varMemoryArray)
	} else {
		return err
	}

	varMemoryArray := _MemoryArray{}

	err = json.Unmarshal(data, &varMemoryArray)
	if err == nil {
		o.EquipmentBase = varMemoryArray.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArrayId")
		delete(additionalProperties, "CpuId")
		delete(additionalProperties, "CurrentCapacity")
		delete(additionalProperties, "ErrorCorrection")
		delete(additionalProperties, "MaxCapacity")
		delete(additionalProperties, "MaxDevices")
		delete(additionalProperties, "OperPowerState")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PersistentMemoryUnits")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Units")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableMemoryArray struct {
	value *MemoryArray
	isSet bool
}

func (v NullableMemoryArray) Get() *MemoryArray {
	return v.value
}

func (v *NullableMemoryArray) Set(val *MemoryArray) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryArray) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryArray(val *MemoryArray) *NullableMemoryArray {
	return &NullableMemoryArray{value: val, isSet: true}
}

func (v NullableMemoryArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
