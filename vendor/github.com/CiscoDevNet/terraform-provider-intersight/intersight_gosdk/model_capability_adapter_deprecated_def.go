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

// checks if the CapabilityAdapterDeprecatedDef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityAdapterDeprecatedDef{}

// CapabilityAdapterDeprecatedDef Object to represent an unsupported/deprecated adapter. Meant to be used under server descriptor object.
type CapabilityAdapterDeprecatedDef struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Model of the unsupported adapter.
	Model *string `json:"Model,omitempty"`
	// Vendor of the unsupported adapter.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityAdapterDeprecatedDef CapabilityAdapterDeprecatedDef

// NewCapabilityAdapterDeprecatedDef instantiates a new CapabilityAdapterDeprecatedDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityAdapterDeprecatedDef(classId string, objectType string) *CapabilityAdapterDeprecatedDef {
	this := CapabilityAdapterDeprecatedDef{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityAdapterDeprecatedDefWithDefaults instantiates a new CapabilityAdapterDeprecatedDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityAdapterDeprecatedDefWithDefaults() *CapabilityAdapterDeprecatedDef {
	this := CapabilityAdapterDeprecatedDef{}
	var classId string = "capability.AdapterDeprecatedDef"
	this.ClassId = classId
	var objectType string = "capability.AdapterDeprecatedDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityAdapterDeprecatedDef) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterDeprecatedDef) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityAdapterDeprecatedDef) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.AdapterDeprecatedDef" of the ClassId field.
func (o *CapabilityAdapterDeprecatedDef) GetDefaultClassId() interface{} {
	return "capability.AdapterDeprecatedDef"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityAdapterDeprecatedDef) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterDeprecatedDef) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityAdapterDeprecatedDef) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.AdapterDeprecatedDef" of the ObjectType field.
func (o *CapabilityAdapterDeprecatedDef) GetDefaultObjectType() interface{} {
	return "capability.AdapterDeprecatedDef"
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CapabilityAdapterDeprecatedDef) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterDeprecatedDef) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CapabilityAdapterDeprecatedDef) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CapabilityAdapterDeprecatedDef) SetModel(v string) {
	o.Model = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *CapabilityAdapterDeprecatedDef) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterDeprecatedDef) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *CapabilityAdapterDeprecatedDef) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *CapabilityAdapterDeprecatedDef) SetVendor(v string) {
	o.Vendor = &v
}

func (o CapabilityAdapterDeprecatedDef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityAdapterDeprecatedDef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	if !IsNil(o.Vendor) {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityAdapterDeprecatedDef) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityAdapterDeprecatedDefWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Model of the unsupported adapter.
		Model *string `json:"Model,omitempty"`
		// Vendor of the unsupported adapter.
		Vendor *string `json:"Vendor,omitempty"`
	}

	varCapabilityAdapterDeprecatedDefWithoutEmbeddedStruct := CapabilityAdapterDeprecatedDefWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityAdapterDeprecatedDefWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityAdapterDeprecatedDef := _CapabilityAdapterDeprecatedDef{}
		varCapabilityAdapterDeprecatedDef.ClassId = varCapabilityAdapterDeprecatedDefWithoutEmbeddedStruct.ClassId
		varCapabilityAdapterDeprecatedDef.ObjectType = varCapabilityAdapterDeprecatedDefWithoutEmbeddedStruct.ObjectType
		varCapabilityAdapterDeprecatedDef.Model = varCapabilityAdapterDeprecatedDefWithoutEmbeddedStruct.Model
		varCapabilityAdapterDeprecatedDef.Vendor = varCapabilityAdapterDeprecatedDefWithoutEmbeddedStruct.Vendor
		*o = CapabilityAdapterDeprecatedDef(varCapabilityAdapterDeprecatedDef)
	} else {
		return err
	}

	varCapabilityAdapterDeprecatedDef := _CapabilityAdapterDeprecatedDef{}

	err = json.Unmarshal(data, &varCapabilityAdapterDeprecatedDef)
	if err == nil {
		o.CapabilityCapability = varCapabilityAdapterDeprecatedDef.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Vendor")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableCapabilityAdapterDeprecatedDef struct {
	value *CapabilityAdapterDeprecatedDef
	isSet bool
}

func (v NullableCapabilityAdapterDeprecatedDef) Get() *CapabilityAdapterDeprecatedDef {
	return v.value
}

func (v *NullableCapabilityAdapterDeprecatedDef) Set(val *CapabilityAdapterDeprecatedDef) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityAdapterDeprecatedDef) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityAdapterDeprecatedDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityAdapterDeprecatedDef(val *CapabilityAdapterDeprecatedDef) *NullableCapabilityAdapterDeprecatedDef {
	return &NullableCapabilityAdapterDeprecatedDef{value: val, isSet: true}
}

func (v NullableCapabilityAdapterDeprecatedDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityAdapterDeprecatedDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
