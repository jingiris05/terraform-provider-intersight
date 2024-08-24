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

// checks if the ResourceSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSelector{}

// ResourceSelector Complex type representing the filter condition to extract the resources.
type ResourceSelector struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// ODATA filter to select resources. The group selector may include URLs of individual resource, or OData query with filters that match multiple queries. The URLs must be relative (i.e. do not include the host).
	Selector             *string `json:"Selector,omitempty" validate:"regexp=^$|\\/api\\/v1\\/.*"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSelector ResourceSelector

// NewResourceSelector instantiates a new ResourceSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSelector(classId string, objectType string) *ResourceSelector {
	this := ResourceSelector{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourceSelectorWithDefaults instantiates a new ResourceSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSelectorWithDefaults() *ResourceSelector {
	this := ResourceSelector{}
	var classId string = "resource.Selector"
	this.ClassId = classId
	var objectType string = "resource.Selector"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceSelector) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceSelector) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceSelector) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "resource.Selector" of the ClassId field.
func (o *ResourceSelector) GetDefaultClassId() interface{} {
	return "resource.Selector"
}

// GetObjectType returns the ObjectType field value
func (o *ResourceSelector) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceSelector) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceSelector) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "resource.Selector" of the ObjectType field.
func (o *ResourceSelector) GetDefaultObjectType() interface{} {
	return "resource.Selector"
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *ResourceSelector) GetSelector() string {
	if o == nil || IsNil(o.Selector) {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelector) GetSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *ResourceSelector) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *ResourceSelector) SetSelector(v string) {
	o.Selector = &v
}

func (o ResourceSelector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSelector) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Selector) {
		toSerialize["Selector"] = o.Selector
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSelector) UnmarshalJSON(data []byte) (err error) {
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
	type ResourceSelectorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// ODATA filter to select resources. The group selector may include URLs of individual resource, or OData query with filters that match multiple queries. The URLs must be relative (i.e. do not include the host).
		Selector *string `json:"Selector,omitempty" validate:"regexp=^$|\\/api\\/v1\\/.*"`
	}

	varResourceSelectorWithoutEmbeddedStruct := ResourceSelectorWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varResourceSelectorWithoutEmbeddedStruct)
	if err == nil {
		varResourceSelector := _ResourceSelector{}
		varResourceSelector.ClassId = varResourceSelectorWithoutEmbeddedStruct.ClassId
		varResourceSelector.ObjectType = varResourceSelectorWithoutEmbeddedStruct.ObjectType
		varResourceSelector.Selector = varResourceSelectorWithoutEmbeddedStruct.Selector
		*o = ResourceSelector(varResourceSelector)
	} else {
		return err
	}

	varResourceSelector := _ResourceSelector{}

	err = json.Unmarshal(data, &varResourceSelector)
	if err == nil {
		o.MoBaseComplexType = varResourceSelector.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Selector")

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

type NullableResourceSelector struct {
	value *ResourceSelector
	isSet bool
}

func (v NullableResourceSelector) Get() *ResourceSelector {
	return v.value
}

func (v *NullableResourceSelector) Set(val *ResourceSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSelector(val *ResourceSelector) *NullableResourceSelector {
	return &NullableResourceSelector{value: val, isSet: true}
}

func (v NullableResourceSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
