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

// checks if the EquipmentExpanderModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentExpanderModule{}

// EquipmentExpanderModule Expander module inside the chassis.
type EquipmentExpanderModule struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Module identifier for the expander module.
	ModuleId   *int64   `json:"ModuleId,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	// Operational state of expander module.
	OperState *string `json:"OperState,omitempty"`
	// Part number identifier for the expander module.
	PartNumber       *string                              `json:"PartNumber,omitempty"`
	EquipmentChassis NullableEquipmentChassisRelationship `json:"EquipmentChassis,omitempty"`
	// An array of relationships to equipmentFanModule resources.
	FanModules           []EquipmentFanModuleRelationship            `json:"FanModules,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentExpanderModule EquipmentExpanderModule

// NewEquipmentExpanderModule instantiates a new EquipmentExpanderModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentExpanderModule(classId string, objectType string) *EquipmentExpanderModule {
	this := EquipmentExpanderModule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentExpanderModuleWithDefaults instantiates a new EquipmentExpanderModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentExpanderModuleWithDefaults() *EquipmentExpanderModule {
	this := EquipmentExpanderModule{}
	var classId string = "equipment.ExpanderModule"
	this.ClassId = classId
	var objectType string = "equipment.ExpanderModule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentExpanderModule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentExpanderModule) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "equipment.ExpanderModule" of the ClassId field.
func (o *EquipmentExpanderModule) GetDefaultClassId() interface{} {
	return "equipment.ExpanderModule"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentExpanderModule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentExpanderModule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "equipment.ExpanderModule" of the ObjectType field.
func (o *EquipmentExpanderModule) GetDefaultObjectType() interface{} {
	return "equipment.ExpanderModule"
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentExpanderModule) GetModuleId() int64 {
	if o == nil || IsNil(o.ModuleId) {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetModuleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ModuleId) {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasModuleId() bool {
	if o != nil && !IsNil(o.ModuleId) {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentExpanderModule) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentExpanderModule) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentExpanderModule) GetOperReasonOk() ([]string, bool) {
	if o == nil || IsNil(o.OperReason) {
		return nil, false
	}
	return o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasOperReason() bool {
	if o != nil && !IsNil(o.OperReason) {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *EquipmentExpanderModule) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentExpanderModule) GetOperState() string {
	if o == nil || IsNil(o.OperState) {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetOperStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperState) {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasOperState() bool {
	if o != nil && !IsNil(o.OperState) {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentExpanderModule) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentExpanderModule) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentExpanderModule) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentExpanderModule) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || IsNil(o.EquipmentChassis.Get()) {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis.Get()
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentExpanderModule) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.EquipmentChassis.Get(), o.EquipmentChassis.IsSet()
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis.IsSet() {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given NullableEquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentExpanderModule) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis.Set(&v)
}

// SetEquipmentChassisNil sets the value for EquipmentChassis to be an explicit nil
func (o *EquipmentExpanderModule) SetEquipmentChassisNil() {
	o.EquipmentChassis.Set(nil)
}

// UnsetEquipmentChassis ensures that no value is present for EquipmentChassis, not even an explicit nil
func (o *EquipmentExpanderModule) UnsetEquipmentChassis() {
	o.EquipmentChassis.Unset()
}

// GetFanModules returns the FanModules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentExpanderModule) GetFanModules() []EquipmentFanModuleRelationship {
	if o == nil {
		var ret []EquipmentFanModuleRelationship
		return ret
	}
	return o.FanModules
}

// GetFanModulesOk returns a tuple with the FanModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentExpanderModule) GetFanModulesOk() ([]EquipmentFanModuleRelationship, bool) {
	if o == nil || IsNil(o.FanModules) {
		return nil, false
	}
	return o.FanModules, true
}

// HasFanModules returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasFanModules() bool {
	if o != nil && !IsNil(o.FanModules) {
		return true
	}

	return false
}

// SetFanModules gets a reference to the given []EquipmentFanModuleRelationship and assigns it to the FanModules field.
func (o *EquipmentExpanderModule) SetFanModules(v []EquipmentFanModuleRelationship) {
	o.FanModules = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentExpanderModule) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentExpanderModule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentExpanderModule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *EquipmentExpanderModule) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *EquipmentExpanderModule) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o EquipmentExpanderModule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentExpanderModule) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ModuleId) {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if !IsNil(o.OperState) {
		toSerialize["OperState"] = o.OperState
	}
	if !IsNil(o.PartNumber) {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.EquipmentChassis.IsSet() {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis.Get()
	}
	if o.FanModules != nil {
		toSerialize["FanModules"] = o.FanModules
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentExpanderModule) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentExpanderModuleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Module identifier for the expander module.
		ModuleId   *int64   `json:"ModuleId,omitempty"`
		OperReason []string `json:"OperReason,omitempty"`
		// Operational state of expander module.
		OperState *string `json:"OperState,omitempty"`
		// Part number identifier for the expander module.
		PartNumber       *string                              `json:"PartNumber,omitempty"`
		EquipmentChassis NullableEquipmentChassisRelationship `json:"EquipmentChassis,omitempty"`
		// An array of relationships to equipmentFanModule resources.
		FanModules       []EquipmentFanModuleRelationship            `json:"FanModules,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentExpanderModuleWithoutEmbeddedStruct := EquipmentExpanderModuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentExpanderModuleWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentExpanderModule := _EquipmentExpanderModule{}
		varEquipmentExpanderModule.ClassId = varEquipmentExpanderModuleWithoutEmbeddedStruct.ClassId
		varEquipmentExpanderModule.ObjectType = varEquipmentExpanderModuleWithoutEmbeddedStruct.ObjectType
		varEquipmentExpanderModule.ModuleId = varEquipmentExpanderModuleWithoutEmbeddedStruct.ModuleId
		varEquipmentExpanderModule.OperReason = varEquipmentExpanderModuleWithoutEmbeddedStruct.OperReason
		varEquipmentExpanderModule.OperState = varEquipmentExpanderModuleWithoutEmbeddedStruct.OperState
		varEquipmentExpanderModule.PartNumber = varEquipmentExpanderModuleWithoutEmbeddedStruct.PartNumber
		varEquipmentExpanderModule.EquipmentChassis = varEquipmentExpanderModuleWithoutEmbeddedStruct.EquipmentChassis
		varEquipmentExpanderModule.FanModules = varEquipmentExpanderModuleWithoutEmbeddedStruct.FanModules
		varEquipmentExpanderModule.RegisteredDevice = varEquipmentExpanderModuleWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentExpanderModule(varEquipmentExpanderModule)
	} else {
		return err
	}

	varEquipmentExpanderModule := _EquipmentExpanderModule{}

	err = json.Unmarshal(data, &varEquipmentExpanderModule)
	if err == nil {
		o.EquipmentBase = varEquipmentExpanderModule.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "FanModules")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableEquipmentExpanderModule struct {
	value *EquipmentExpanderModule
	isSet bool
}

func (v NullableEquipmentExpanderModule) Get() *EquipmentExpanderModule {
	return v.value
}

func (v *NullableEquipmentExpanderModule) Set(val *EquipmentExpanderModule) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentExpanderModule) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentExpanderModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentExpanderModule(val *EquipmentExpanderModule) *NullableEquipmentExpanderModule {
	return &NullableEquipmentExpanderModule{value: val, isSet: true}
}

func (v NullableEquipmentExpanderModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentExpanderModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
