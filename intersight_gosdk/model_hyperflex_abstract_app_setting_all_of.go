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
)

// HyperflexAbstractAppSettingAllOf Definition of the list of properties defined in 'hyperflex.AbstractAppSetting', excluding properties defined in parent classes.
type HyperflexAbstractAppSettingAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The application setting identifier.
	Name *string `json:"Name,omitempty"`
	// The application setting value.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexAbstractAppSettingAllOf HyperflexAbstractAppSettingAllOf

// NewHyperflexAbstractAppSettingAllOf instantiates a new HyperflexAbstractAppSettingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAbstractAppSettingAllOf(classId string, objectType string) *HyperflexAbstractAppSettingAllOf {
	this := HyperflexAbstractAppSettingAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexAbstractAppSettingAllOfWithDefaults instantiates a new HyperflexAbstractAppSettingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAbstractAppSettingAllOfWithDefaults() *HyperflexAbstractAppSettingAllOf {
	this := HyperflexAbstractAppSettingAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexAbstractAppSettingAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexAbstractAppSettingAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexAbstractAppSettingAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexAbstractAppSettingAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexAbstractAppSettingAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexAbstractAppSettingAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexAbstractAppSettingAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAbstractAppSettingAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexAbstractAppSettingAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexAbstractAppSettingAllOf) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HyperflexAbstractAppSettingAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAbstractAppSettingAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HyperflexAbstractAppSettingAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HyperflexAbstractAppSettingAllOf) SetValue(v string) {
	o.Value = &v
}

func (o HyperflexAbstractAppSettingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexAbstractAppSettingAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexAbstractAppSettingAllOf := _HyperflexAbstractAppSettingAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexAbstractAppSettingAllOf); err == nil {
		*o = HyperflexAbstractAppSettingAllOf(varHyperflexAbstractAppSettingAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexAbstractAppSettingAllOf struct {
	value *HyperflexAbstractAppSettingAllOf
	isSet bool
}

func (v NullableHyperflexAbstractAppSettingAllOf) Get() *HyperflexAbstractAppSettingAllOf {
	return v.value
}

func (v *NullableHyperflexAbstractAppSettingAllOf) Set(val *HyperflexAbstractAppSettingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAbstractAppSettingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAbstractAppSettingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAbstractAppSettingAllOf(val *HyperflexAbstractAppSettingAllOf) *NullableHyperflexAbstractAppSettingAllOf {
	return &NullableHyperflexAbstractAppSettingAllOf{value: val, isSet: true}
}

func (v NullableHyperflexAbstractAppSettingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAbstractAppSettingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
