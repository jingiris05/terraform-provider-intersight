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

// checks if the ApplianceKeyValuePair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceKeyValuePair{}

// ApplianceKeyValuePair An arbitrary key and value pair that can be used to for having pair of key and values in application such as Appliance Capabilities.
type ApplianceKeyValuePair struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The string representation of a tag key.
	Key *string `json:"Key,omitempty"`
	// The string representation of a tag value.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceKeyValuePair ApplianceKeyValuePair

// NewApplianceKeyValuePair instantiates a new ApplianceKeyValuePair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceKeyValuePair(classId string, objectType string) *ApplianceKeyValuePair {
	this := ApplianceKeyValuePair{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceKeyValuePairWithDefaults instantiates a new ApplianceKeyValuePair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceKeyValuePairWithDefaults() *ApplianceKeyValuePair {
	this := ApplianceKeyValuePair{}
	var classId string = "appliance.KeyValuePair"
	this.ClassId = classId
	var objectType string = "appliance.KeyValuePair"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceKeyValuePair) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceKeyValuePair) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceKeyValuePair) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.KeyValuePair" of the ClassId field.
func (o *ApplianceKeyValuePair) GetDefaultClassId() interface{} {
	return "appliance.KeyValuePair"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceKeyValuePair) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceKeyValuePair) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceKeyValuePair) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.KeyValuePair" of the ObjectType field.
func (o *ApplianceKeyValuePair) GetDefaultObjectType() interface{} {
	return "appliance.KeyValuePair"
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ApplianceKeyValuePair) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceKeyValuePair) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ApplianceKeyValuePair) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ApplianceKeyValuePair) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApplianceKeyValuePair) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceKeyValuePair) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApplianceKeyValuePair) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApplianceKeyValuePair) SetValue(v string) {
	o.Value = &v
}

func (o ApplianceKeyValuePair) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceKeyValuePair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceKeyValuePair) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceKeyValuePairWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The string representation of a tag key.
		Key *string `json:"Key,omitempty"`
		// The string representation of a tag value.
		Value *string `json:"Value,omitempty"`
	}

	varApplianceKeyValuePairWithoutEmbeddedStruct := ApplianceKeyValuePairWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceKeyValuePairWithoutEmbeddedStruct)
	if err == nil {
		varApplianceKeyValuePair := _ApplianceKeyValuePair{}
		varApplianceKeyValuePair.ClassId = varApplianceKeyValuePairWithoutEmbeddedStruct.ClassId
		varApplianceKeyValuePair.ObjectType = varApplianceKeyValuePairWithoutEmbeddedStruct.ObjectType
		varApplianceKeyValuePair.Key = varApplianceKeyValuePairWithoutEmbeddedStruct.Key
		varApplianceKeyValuePair.Value = varApplianceKeyValuePairWithoutEmbeddedStruct.Value
		*o = ApplianceKeyValuePair(varApplianceKeyValuePair)
	} else {
		return err
	}

	varApplianceKeyValuePair := _ApplianceKeyValuePair{}

	err = json.Unmarshal(data, &varApplianceKeyValuePair)
	if err == nil {
		o.MoBaseComplexType = varApplianceKeyValuePair.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Value")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableApplianceKeyValuePair struct {
	value *ApplianceKeyValuePair
	isSet bool
}

func (v NullableApplianceKeyValuePair) Get() *ApplianceKeyValuePair {
	return v.value
}

func (v *NullableApplianceKeyValuePair) Set(val *ApplianceKeyValuePair) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceKeyValuePair) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceKeyValuePair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceKeyValuePair(val *ApplianceKeyValuePair) *NullableApplianceKeyValuePair {
	return &NullableApplianceKeyValuePair{value: val, isSet: true}
}

func (v NullableApplianceKeyValuePair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceKeyValuePair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
