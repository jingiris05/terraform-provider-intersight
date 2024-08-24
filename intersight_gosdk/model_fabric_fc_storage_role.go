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

// checks if the FabricFcStorageRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricFcStorageRole{}

// FabricFcStorageRole Configuration object sent by user to create a fc uplink port.
type FabricFcStorageRole struct {
	FabricPortRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured speed for the port. * `16Gbps` - Admin configurable speed 16Gbps. * `8Gbps` - Admin configurable speed 8Gbps. * `32Gbps` - Admin configurable speed 32Gbps. * `Auto` - Admin configurable speed AUTO ( default ).
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Virtual San Identifier associated to the FC port.
	VsanId               *int64 `json:"VsanId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcStorageRole FabricFcStorageRole

// NewFabricFcStorageRole instantiates a new FabricFcStorageRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcStorageRole(classId string, objectType string) *FabricFcStorageRole {
	this := FabricFcStorageRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "16Gbps"
	this.AdminSpeed = &adminSpeed
	return &this
}

// NewFabricFcStorageRoleWithDefaults instantiates a new FabricFcStorageRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcStorageRoleWithDefaults() *FabricFcStorageRole {
	this := FabricFcStorageRole{}
	var classId string = "fabric.FcStorageRole"
	this.ClassId = classId
	var objectType string = "fabric.FcStorageRole"
	this.ObjectType = objectType
	var adminSpeed string = "16Gbps"
	this.AdminSpeed = &adminSpeed
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricFcStorageRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricFcStorageRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricFcStorageRole) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.FcStorageRole" of the ClassId field.
func (o *FabricFcStorageRole) GetDefaultClassId() interface{} {
	return "fabric.FcStorageRole"
}

// GetObjectType returns the ObjectType field value
func (o *FabricFcStorageRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricFcStorageRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricFcStorageRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.FcStorageRole" of the ObjectType field.
func (o *FabricFcStorageRole) GetDefaultObjectType() interface{} {
	return "fabric.FcStorageRole"
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricFcStorageRole) GetAdminSpeed() string {
	if o == nil || IsNil(o.AdminSpeed) {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcStorageRole) GetAdminSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.AdminSpeed) {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricFcStorageRole) HasAdminSpeed() bool {
	if o != nil && !IsNil(o.AdminSpeed) {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricFcStorageRole) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetVsanId returns the VsanId field value if set, zero value otherwise.
func (o *FabricFcStorageRole) GetVsanId() int64 {
	if o == nil || IsNil(o.VsanId) {
		var ret int64
		return ret
	}
	return *o.VsanId
}

// GetVsanIdOk returns a tuple with the VsanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcStorageRole) GetVsanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VsanId) {
		return nil, false
	}
	return o.VsanId, true
}

// HasVsanId returns a boolean if a field has been set.
func (o *FabricFcStorageRole) HasVsanId() bool {
	if o != nil && !IsNil(o.VsanId) {
		return true
	}

	return false
}

// SetVsanId gets a reference to the given int64 and assigns it to the VsanId field.
func (o *FabricFcStorageRole) SetVsanId(v int64) {
	o.VsanId = &v
}

func (o FabricFcStorageRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricFcStorageRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortRole, errFabricPortRole := json.Marshal(o.FabricPortRole)
	if errFabricPortRole != nil {
		return map[string]interface{}{}, errFabricPortRole
	}
	errFabricPortRole = json.Unmarshal([]byte(serializedFabricPortRole), &toSerialize)
	if errFabricPortRole != nil {
		return map[string]interface{}{}, errFabricPortRole
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdminSpeed) {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if !IsNil(o.VsanId) {
		toSerialize["VsanId"] = o.VsanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricFcStorageRole) UnmarshalJSON(data []byte) (err error) {
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
	type FabricFcStorageRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin configured speed for the port. * `16Gbps` - Admin configurable speed 16Gbps. * `8Gbps` - Admin configurable speed 8Gbps. * `32Gbps` - Admin configurable speed 32Gbps. * `Auto` - Admin configurable speed AUTO ( default ).
		AdminSpeed *string `json:"AdminSpeed,omitempty"`
		// Virtual San Identifier associated to the FC port.
		VsanId *int64 `json:"VsanId,omitempty"`
	}

	varFabricFcStorageRoleWithoutEmbeddedStruct := FabricFcStorageRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricFcStorageRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricFcStorageRole := _FabricFcStorageRole{}
		varFabricFcStorageRole.ClassId = varFabricFcStorageRoleWithoutEmbeddedStruct.ClassId
		varFabricFcStorageRole.ObjectType = varFabricFcStorageRoleWithoutEmbeddedStruct.ObjectType
		varFabricFcStorageRole.AdminSpeed = varFabricFcStorageRoleWithoutEmbeddedStruct.AdminSpeed
		varFabricFcStorageRole.VsanId = varFabricFcStorageRoleWithoutEmbeddedStruct.VsanId
		*o = FabricFcStorageRole(varFabricFcStorageRole)
	} else {
		return err
	}

	varFabricFcStorageRole := _FabricFcStorageRole{}

	err = json.Unmarshal(data, &varFabricFcStorageRole)
	if err == nil {
		o.FabricPortRole = varFabricFcStorageRole.FabricPortRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "VsanId")

		// remove fields from embedded structs
		reflectFabricPortRole := reflect.ValueOf(o.FabricPortRole)
		for i := 0; i < reflectFabricPortRole.Type().NumField(); i++ {
			t := reflectFabricPortRole.Type().Field(i)

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

type NullableFabricFcStorageRole struct {
	value *FabricFcStorageRole
	isSet bool
}

func (v NullableFabricFcStorageRole) Get() *FabricFcStorageRole {
	return v.value
}

func (v *NullableFabricFcStorageRole) Set(val *FabricFcStorageRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcStorageRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcStorageRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcStorageRole(val *FabricFcStorageRole) *NullableFabricFcStorageRole {
	return &NullableFabricFcStorageRole{value: val, isSet: true}
}

func (v NullableFabricFcStorageRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcStorageRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
