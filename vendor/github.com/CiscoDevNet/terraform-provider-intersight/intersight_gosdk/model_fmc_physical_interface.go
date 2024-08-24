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

// checks if the FmcPhysicalInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FmcPhysicalInterface{}

// FmcPhysicalInterface FMC Devices Physical Interfaces.
type FmcPhysicalInterface struct {
	FmcInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Represents Device Id associated with Physical Interface.
	DeviceId *string `json:"DeviceId,omitempty"`
	// Represents Domain Id associated with Physical Interface.
	DomainId *string `json:"DomainId,omitempty"`
	// Represents Physical Interface Name.
	Name *string `json:"Name,omitempty"`
	// Represents Physical Interface Id.
	PhysicalInterfaceId  *string `json:"PhysicalInterfaceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FmcPhysicalInterface FmcPhysicalInterface

// NewFmcPhysicalInterface instantiates a new FmcPhysicalInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFmcPhysicalInterface(classId string, objectType string) *FmcPhysicalInterface {
	this := FmcPhysicalInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFmcPhysicalInterfaceWithDefaults instantiates a new FmcPhysicalInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFmcPhysicalInterfaceWithDefaults() *FmcPhysicalInterface {
	this := FmcPhysicalInterface{}
	var classId string = "fmc.PhysicalInterface"
	this.ClassId = classId
	var objectType string = "fmc.PhysicalInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FmcPhysicalInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FmcPhysicalInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FmcPhysicalInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fmc.PhysicalInterface" of the ClassId field.
func (o *FmcPhysicalInterface) GetDefaultClassId() interface{} {
	return "fmc.PhysicalInterface"
}

// GetObjectType returns the ObjectType field value
func (o *FmcPhysicalInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FmcPhysicalInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FmcPhysicalInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fmc.PhysicalInterface" of the ObjectType field.
func (o *FmcPhysicalInterface) GetDefaultObjectType() interface{} {
	return "fmc.PhysicalInterface"
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *FmcPhysicalInterface) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FmcPhysicalInterface) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *FmcPhysicalInterface) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *FmcPhysicalInterface) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *FmcPhysicalInterface) GetDomainId() string {
	if o == nil || IsNil(o.DomainId) {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FmcPhysicalInterface) GetDomainIdOk() (*string, bool) {
	if o == nil || IsNil(o.DomainId) {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *FmcPhysicalInterface) HasDomainId() bool {
	if o != nil && !IsNil(o.DomainId) {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *FmcPhysicalInterface) SetDomainId(v string) {
	o.DomainId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FmcPhysicalInterface) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FmcPhysicalInterface) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FmcPhysicalInterface) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FmcPhysicalInterface) SetName(v string) {
	o.Name = &v
}

// GetPhysicalInterfaceId returns the PhysicalInterfaceId field value if set, zero value otherwise.
func (o *FmcPhysicalInterface) GetPhysicalInterfaceId() string {
	if o == nil || IsNil(o.PhysicalInterfaceId) {
		var ret string
		return ret
	}
	return *o.PhysicalInterfaceId
}

// GetPhysicalInterfaceIdOk returns a tuple with the PhysicalInterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FmcPhysicalInterface) GetPhysicalInterfaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalInterfaceId) {
		return nil, false
	}
	return o.PhysicalInterfaceId, true
}

// HasPhysicalInterfaceId returns a boolean if a field has been set.
func (o *FmcPhysicalInterface) HasPhysicalInterfaceId() bool {
	if o != nil && !IsNil(o.PhysicalInterfaceId) {
		return true
	}

	return false
}

// SetPhysicalInterfaceId gets a reference to the given string and assigns it to the PhysicalInterfaceId field.
func (o *FmcPhysicalInterface) SetPhysicalInterfaceId(v string) {
	o.PhysicalInterfaceId = &v
}

func (o FmcPhysicalInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FmcPhysicalInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFmcInventoryEntity, errFmcInventoryEntity := json.Marshal(o.FmcInventoryEntity)
	if errFmcInventoryEntity != nil {
		return map[string]interface{}{}, errFmcInventoryEntity
	}
	errFmcInventoryEntity = json.Unmarshal([]byte(serializedFmcInventoryEntity), &toSerialize)
	if errFmcInventoryEntity != nil {
		return map[string]interface{}{}, errFmcInventoryEntity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DeviceId) {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if !IsNil(o.DomainId) {
		toSerialize["DomainId"] = o.DomainId
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PhysicalInterfaceId) {
		toSerialize["PhysicalInterfaceId"] = o.PhysicalInterfaceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FmcPhysicalInterface) UnmarshalJSON(data []byte) (err error) {
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
	type FmcPhysicalInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Represents Device Id associated with Physical Interface.
		DeviceId *string `json:"DeviceId,omitempty"`
		// Represents Domain Id associated with Physical Interface.
		DomainId *string `json:"DomainId,omitempty"`
		// Represents Physical Interface Name.
		Name *string `json:"Name,omitempty"`
		// Represents Physical Interface Id.
		PhysicalInterfaceId *string `json:"PhysicalInterfaceId,omitempty"`
	}

	varFmcPhysicalInterfaceWithoutEmbeddedStruct := FmcPhysicalInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFmcPhysicalInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varFmcPhysicalInterface := _FmcPhysicalInterface{}
		varFmcPhysicalInterface.ClassId = varFmcPhysicalInterfaceWithoutEmbeddedStruct.ClassId
		varFmcPhysicalInterface.ObjectType = varFmcPhysicalInterfaceWithoutEmbeddedStruct.ObjectType
		varFmcPhysicalInterface.DeviceId = varFmcPhysicalInterfaceWithoutEmbeddedStruct.DeviceId
		varFmcPhysicalInterface.DomainId = varFmcPhysicalInterfaceWithoutEmbeddedStruct.DomainId
		varFmcPhysicalInterface.Name = varFmcPhysicalInterfaceWithoutEmbeddedStruct.Name
		varFmcPhysicalInterface.PhysicalInterfaceId = varFmcPhysicalInterfaceWithoutEmbeddedStruct.PhysicalInterfaceId
		*o = FmcPhysicalInterface(varFmcPhysicalInterface)
	} else {
		return err
	}

	varFmcPhysicalInterface := _FmcPhysicalInterface{}

	err = json.Unmarshal(data, &varFmcPhysicalInterface)
	if err == nil {
		o.FmcInventoryEntity = varFmcPhysicalInterface.FmcInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "DomainId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PhysicalInterfaceId")

		// remove fields from embedded structs
		reflectFmcInventoryEntity := reflect.ValueOf(o.FmcInventoryEntity)
		for i := 0; i < reflectFmcInventoryEntity.Type().NumField(); i++ {
			t := reflectFmcInventoryEntity.Type().Field(i)

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

type NullableFmcPhysicalInterface struct {
	value *FmcPhysicalInterface
	isSet bool
}

func (v NullableFmcPhysicalInterface) Get() *FmcPhysicalInterface {
	return v.value
}

func (v *NullableFmcPhysicalInterface) Set(val *FmcPhysicalInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableFmcPhysicalInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableFmcPhysicalInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFmcPhysicalInterface(val *FmcPhysicalInterface) *NullableFmcPhysicalInterface {
	return &NullableFmcPhysicalInterface{value: val, isSet: true}
}

func (v NullableFmcPhysicalInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFmcPhysicalInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
