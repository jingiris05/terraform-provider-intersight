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

// checks if the StoragePurePort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoragePurePort{}

// StoragePurePort Port entity in Pure FlashArray.
type StoragePurePort struct {
	StorageBasePhysicalPort
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the port to which this port has failed over.
	Failover *string `json:"Failover,omitempty"`
	// The NVMe Qualified Name (NQN) associated with the host for ethernet port.
	Nqn *string `json:"Nqn,omitempty"`
	// Ip address of iSCSI portal configured on the port.
	Portal               *string                                     `json:"Portal,omitempty"`
	Array                NullableStoragePureArrayRelationship        `json:"Array,omitempty"`
	Controller           NullableStoragePureControllerRelationship   `json:"Controller,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePurePort StoragePurePort

// NewStoragePurePort instantiates a new StoragePurePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePurePort(classId string, objectType string) *StoragePurePort {
	this := StoragePurePort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePurePortWithDefaults instantiates a new StoragePurePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePurePortWithDefaults() *StoragePurePort {
	this := StoragePurePort{}
	var classId string = "storage.PurePort"
	this.ClassId = classId
	var objectType string = "storage.PurePort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePurePort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePurePort) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.PurePort" of the ClassId field.
func (o *StoragePurePort) GetDefaultClassId() interface{} {
	return "storage.PurePort"
}

// GetObjectType returns the ObjectType field value
func (o *StoragePurePort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePurePort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.PurePort" of the ObjectType field.
func (o *StoragePurePort) GetDefaultObjectType() interface{} {
	return "storage.PurePort"
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *StoragePurePort) GetFailover() string {
	if o == nil || IsNil(o.Failover) {
		var ret string
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetFailoverOk() (*string, bool) {
	if o == nil || IsNil(o.Failover) {
		return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *StoragePurePort) HasFailover() bool {
	if o != nil && !IsNil(o.Failover) {
		return true
	}

	return false
}

// SetFailover gets a reference to the given string and assigns it to the Failover field.
func (o *StoragePurePort) SetFailover(v string) {
	o.Failover = &v
}

// GetNqn returns the Nqn field value if set, zero value otherwise.
func (o *StoragePurePort) GetNqn() string {
	if o == nil || IsNil(o.Nqn) {
		var ret string
		return ret
	}
	return *o.Nqn
}

// GetNqnOk returns a tuple with the Nqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetNqnOk() (*string, bool) {
	if o == nil || IsNil(o.Nqn) {
		return nil, false
	}
	return o.Nqn, true
}

// HasNqn returns a boolean if a field has been set.
func (o *StoragePurePort) HasNqn() bool {
	if o != nil && !IsNil(o.Nqn) {
		return true
	}

	return false
}

// SetNqn gets a reference to the given string and assigns it to the Nqn field.
func (o *StoragePurePort) SetNqn(v string) {
	o.Nqn = &v
}

// GetPortal returns the Portal field value if set, zero value otherwise.
func (o *StoragePurePort) GetPortal() string {
	if o == nil || IsNil(o.Portal) {
		var ret string
		return ret
	}
	return *o.Portal
}

// GetPortalOk returns a tuple with the Portal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePort) GetPortalOk() (*string, bool) {
	if o == nil || IsNil(o.Portal) {
		return nil, false
	}
	return o.Portal, true
}

// HasPortal returns a boolean if a field has been set.
func (o *StoragePurePort) HasPortal() bool {
	if o != nil && !IsNil(o.Portal) {
		return true
	}

	return false
}

// SetPortal gets a reference to the given string and assigns it to the Portal field.
func (o *StoragePurePort) SetPortal(v string) {
	o.Portal = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePurePort) GetArray() StoragePureArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePurePort) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePurePort) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePurePort) SetArray(v StoragePureArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StoragePurePort) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StoragePurePort) UnsetArray() {
	o.Array.Unset()
}

// GetController returns the Controller field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePurePort) GetController() StoragePureControllerRelationship {
	if o == nil || IsNil(o.Controller.Get()) {
		var ret StoragePureControllerRelationship
		return ret
	}
	return *o.Controller.Get()
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePurePort) GetControllerOk() (*StoragePureControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Controller.Get(), o.Controller.IsSet()
}

// HasController returns a boolean if a field has been set.
func (o *StoragePurePort) HasController() bool {
	if o != nil && o.Controller.IsSet() {
		return true
	}

	return false
}

// SetController gets a reference to the given NullableStoragePureControllerRelationship and assigns it to the Controller field.
func (o *StoragePurePort) SetController(v StoragePureControllerRelationship) {
	o.Controller.Set(&v)
}

// SetControllerNil sets the value for Controller to be an explicit nil
func (o *StoragePurePort) SetControllerNil() {
	o.Controller.Set(nil)
}

// UnsetController ensures that no value is present for Controller, not even an explicit nil
func (o *StoragePurePort) UnsetController() {
	o.Controller.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePurePort) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePurePort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePurePort) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePurePort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StoragePurePort) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StoragePurePort) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StoragePurePort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoragePurePort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBasePhysicalPort, errStorageBasePhysicalPort := json.Marshal(o.StorageBasePhysicalPort)
	if errStorageBasePhysicalPort != nil {
		return map[string]interface{}{}, errStorageBasePhysicalPort
	}
	errStorageBasePhysicalPort = json.Unmarshal([]byte(serializedStorageBasePhysicalPort), &toSerialize)
	if errStorageBasePhysicalPort != nil {
		return map[string]interface{}{}, errStorageBasePhysicalPort
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Failover) {
		toSerialize["Failover"] = o.Failover
	}
	if !IsNil(o.Nqn) {
		toSerialize["Nqn"] = o.Nqn
	}
	if !IsNil(o.Portal) {
		toSerialize["Portal"] = o.Portal
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.Controller.IsSet() {
		toSerialize["Controller"] = o.Controller.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoragePurePort) UnmarshalJSON(data []byte) (err error) {
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
	type StoragePurePortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the port to which this port has failed over.
		Failover *string `json:"Failover,omitempty"`
		// The NVMe Qualified Name (NQN) associated with the host for ethernet port.
		Nqn *string `json:"Nqn,omitempty"`
		// Ip address of iSCSI portal configured on the port.
		Portal           *string                                     `json:"Portal,omitempty"`
		Array            NullableStoragePureArrayRelationship        `json:"Array,omitempty"`
		Controller       NullableStoragePureControllerRelationship   `json:"Controller,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStoragePurePortWithoutEmbeddedStruct := StoragePurePortWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStoragePurePortWithoutEmbeddedStruct)
	if err == nil {
		varStoragePurePort := _StoragePurePort{}
		varStoragePurePort.ClassId = varStoragePurePortWithoutEmbeddedStruct.ClassId
		varStoragePurePort.ObjectType = varStoragePurePortWithoutEmbeddedStruct.ObjectType
		varStoragePurePort.Failover = varStoragePurePortWithoutEmbeddedStruct.Failover
		varStoragePurePort.Nqn = varStoragePurePortWithoutEmbeddedStruct.Nqn
		varStoragePurePort.Portal = varStoragePurePortWithoutEmbeddedStruct.Portal
		varStoragePurePort.Array = varStoragePurePortWithoutEmbeddedStruct.Array
		varStoragePurePort.Controller = varStoragePurePortWithoutEmbeddedStruct.Controller
		varStoragePurePort.RegisteredDevice = varStoragePurePortWithoutEmbeddedStruct.RegisteredDevice
		*o = StoragePurePort(varStoragePurePort)
	} else {
		return err
	}

	varStoragePurePort := _StoragePurePort{}

	err = json.Unmarshal(data, &varStoragePurePort)
	if err == nil {
		o.StorageBasePhysicalPort = varStoragePurePort.StorageBasePhysicalPort
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Failover")
		delete(additionalProperties, "Nqn")
		delete(additionalProperties, "Portal")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Controller")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBasePhysicalPort := reflect.ValueOf(o.StorageBasePhysicalPort)
		for i := 0; i < reflectStorageBasePhysicalPort.Type().NumField(); i++ {
			t := reflectStorageBasePhysicalPort.Type().Field(i)

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

type NullableStoragePurePort struct {
	value *StoragePurePort
	isSet bool
}

func (v NullableStoragePurePort) Get() *StoragePurePort {
	return v.value
}

func (v *NullableStoragePurePort) Set(val *StoragePurePort) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePurePort) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePurePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePurePort(val *StoragePurePort) *NullableStoragePurePort {
	return &NullableStoragePurePort{value: val, isSet: true}
}

func (v NullableStoragePurePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePurePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
