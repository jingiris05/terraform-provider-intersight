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

// checks if the EtherHostPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EtherHostPort{}

// EtherHostPort Host Interface ports available on the I/O module or the Fabric Extender that facilitate connectivity between the Fabric Interconnect and the Cisco UCS B/C/X-Series servers.
type EtherHostPort struct {
	EtherPhysicalPortBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Breakout port member in the fabric extender.
	AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
	// Fabric extender identifier for this port.
	ModuleId *int64 `json:"ModuleId,omitempty"`
	// Host Port Speed of IO card or fabric extender.
	Speed                *string                                     `json:"Speed,omitempty"`
	EquipmentIoCardBase  NullableEquipmentIoCardBaseRelationship     `json:"EquipmentIoCardBase,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EtherHostPort EtherHostPort

// NewEtherHostPort instantiates a new EtherHostPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtherHostPort(classId string, objectType string) *EtherHostPort {
	this := EtherHostPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEtherHostPortWithDefaults instantiates a new EtherHostPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtherHostPortWithDefaults() *EtherHostPort {
	this := EtherHostPort{}
	var classId string = "ether.HostPort"
	this.ClassId = classId
	var objectType string = "ether.HostPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EtherHostPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EtherHostPort) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "ether.HostPort" of the ClassId field.
func (o *EtherHostPort) GetDefaultClassId() interface{} {
	return "ether.HostPort"
}

// GetObjectType returns the ObjectType field value
func (o *EtherHostPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EtherHostPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "ether.HostPort" of the ObjectType field.
func (o *EtherHostPort) GetDefaultObjectType() interface{} {
	return "ether.HostPort"
}

// GetAggregatePortId returns the AggregatePortId field value if set, zero value otherwise.
func (o *EtherHostPort) GetAggregatePortId() int64 {
	if o == nil || IsNil(o.AggregatePortId) {
		var ret int64
		return ret
	}
	return *o.AggregatePortId
}

// GetAggregatePortIdOk returns a tuple with the AggregatePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetAggregatePortIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AggregatePortId) {
		return nil, false
	}
	return o.AggregatePortId, true
}

// HasAggregatePortId returns a boolean if a field has been set.
func (o *EtherHostPort) HasAggregatePortId() bool {
	if o != nil && !IsNil(o.AggregatePortId) {
		return true
	}

	return false
}

// SetAggregatePortId gets a reference to the given int64 and assigns it to the AggregatePortId field.
func (o *EtherHostPort) SetAggregatePortId(v int64) {
	o.AggregatePortId = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EtherHostPort) GetModuleId() int64 {
	if o == nil || IsNil(o.ModuleId) {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetModuleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ModuleId) {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EtherHostPort) HasModuleId() bool {
	if o != nil && !IsNil(o.ModuleId) {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EtherHostPort) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *EtherHostPort) GetSpeed() string {
	if o == nil || IsNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherHostPort) GetSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *EtherHostPort) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *EtherHostPort) SetSpeed(v string) {
	o.Speed = &v
}

// GetEquipmentIoCardBase returns the EquipmentIoCardBase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EtherHostPort) GetEquipmentIoCardBase() EquipmentIoCardBaseRelationship {
	if o == nil || IsNil(o.EquipmentIoCardBase.Get()) {
		var ret EquipmentIoCardBaseRelationship
		return ret
	}
	return *o.EquipmentIoCardBase.Get()
}

// GetEquipmentIoCardBaseOk returns a tuple with the EquipmentIoCardBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EtherHostPort) GetEquipmentIoCardBaseOk() (*EquipmentIoCardBaseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.EquipmentIoCardBase.Get(), o.EquipmentIoCardBase.IsSet()
}

// HasEquipmentIoCardBase returns a boolean if a field has been set.
func (o *EtherHostPort) HasEquipmentIoCardBase() bool {
	if o != nil && o.EquipmentIoCardBase.IsSet() {
		return true
	}

	return false
}

// SetEquipmentIoCardBase gets a reference to the given NullableEquipmentIoCardBaseRelationship and assigns it to the EquipmentIoCardBase field.
func (o *EtherHostPort) SetEquipmentIoCardBase(v EquipmentIoCardBaseRelationship) {
	o.EquipmentIoCardBase.Set(&v)
}

// SetEquipmentIoCardBaseNil sets the value for EquipmentIoCardBase to be an explicit nil
func (o *EtherHostPort) SetEquipmentIoCardBaseNil() {
	o.EquipmentIoCardBase.Set(nil)
}

// UnsetEquipmentIoCardBase ensures that no value is present for EquipmentIoCardBase, not even an explicit nil
func (o *EtherHostPort) UnsetEquipmentIoCardBase() {
	o.EquipmentIoCardBase.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EtherHostPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EtherHostPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EtherHostPort) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EtherHostPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *EtherHostPort) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *EtherHostPort) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o EtherHostPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EtherHostPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEtherPhysicalPortBase, errEtherPhysicalPortBase := json.Marshal(o.EtherPhysicalPortBase)
	if errEtherPhysicalPortBase != nil {
		return map[string]interface{}{}, errEtherPhysicalPortBase
	}
	errEtherPhysicalPortBase = json.Unmarshal([]byte(serializedEtherPhysicalPortBase), &toSerialize)
	if errEtherPhysicalPortBase != nil {
		return map[string]interface{}{}, errEtherPhysicalPortBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AggregatePortId) {
		toSerialize["AggregatePortId"] = o.AggregatePortId
	}
	if !IsNil(o.ModuleId) {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if !IsNil(o.Speed) {
		toSerialize["Speed"] = o.Speed
	}
	if o.EquipmentIoCardBase.IsSet() {
		toSerialize["EquipmentIoCardBase"] = o.EquipmentIoCardBase.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EtherHostPort) UnmarshalJSON(data []byte) (err error) {
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
	type EtherHostPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Breakout port member in the fabric extender.
		AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
		// Fabric extender identifier for this port.
		ModuleId *int64 `json:"ModuleId,omitempty"`
		// Host Port Speed of IO card or fabric extender.
		Speed               *string                                     `json:"Speed,omitempty"`
		EquipmentIoCardBase NullableEquipmentIoCardBaseRelationship     `json:"EquipmentIoCardBase,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEtherHostPortWithoutEmbeddedStruct := EtherHostPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEtherHostPortWithoutEmbeddedStruct)
	if err == nil {
		varEtherHostPort := _EtherHostPort{}
		varEtherHostPort.ClassId = varEtherHostPortWithoutEmbeddedStruct.ClassId
		varEtherHostPort.ObjectType = varEtherHostPortWithoutEmbeddedStruct.ObjectType
		varEtherHostPort.AggregatePortId = varEtherHostPortWithoutEmbeddedStruct.AggregatePortId
		varEtherHostPort.ModuleId = varEtherHostPortWithoutEmbeddedStruct.ModuleId
		varEtherHostPort.Speed = varEtherHostPortWithoutEmbeddedStruct.Speed
		varEtherHostPort.EquipmentIoCardBase = varEtherHostPortWithoutEmbeddedStruct.EquipmentIoCardBase
		varEtherHostPort.RegisteredDevice = varEtherHostPortWithoutEmbeddedStruct.RegisteredDevice
		*o = EtherHostPort(varEtherHostPort)
	} else {
		return err
	}

	varEtherHostPort := _EtherHostPort{}

	err = json.Unmarshal(data, &varEtherHostPort)
	if err == nil {
		o.EtherPhysicalPortBase = varEtherHostPort.EtherPhysicalPortBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggregatePortId")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "EquipmentIoCardBase")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEtherPhysicalPortBase := reflect.ValueOf(o.EtherPhysicalPortBase)
		for i := 0; i < reflectEtherPhysicalPortBase.Type().NumField(); i++ {
			t := reflectEtherPhysicalPortBase.Type().Field(i)

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

type NullableEtherHostPort struct {
	value *EtherHostPort
	isSet bool
}

func (v NullableEtherHostPort) Get() *EtherHostPort {
	return v.value
}

func (v *NullableEtherHostPort) Set(val *EtherHostPort) {
	v.value = val
	v.isSet = true
}

func (v NullableEtherHostPort) IsSet() bool {
	return v.isSet
}

func (v *NullableEtherHostPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtherHostPort(val *EtherHostPort) *NullableEtherHostPort {
	return &NullableEtherHostPort{value: val, isSet: true}
}

func (v NullableEtherHostPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtherHostPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
