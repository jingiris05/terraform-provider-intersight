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

// checks if the PoolAbstractIdPoolMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolAbstractIdPoolMember{}

// PoolAbstractIdPoolMember A single member of an identity pool such as IP, MAC, UUID etc.
type PoolAbstractIdPoolMember struct {
	PoolAbstractPoolMember
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Boolean to represent whether the ID is used either statically or by another pool.
	AssignedByAnother *bool `json:"AssignedByAnother,omitempty"`
	// Boolean to represent whether the ID is reserved.
	Reserved             *bool `json:"Reserved,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractIdPoolMember PoolAbstractIdPoolMember

// NewPoolAbstractIdPoolMember instantiates a new PoolAbstractIdPoolMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractIdPoolMember(classId string, objectType string) *PoolAbstractIdPoolMember {
	this := PoolAbstractIdPoolMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assigned bool = false
	this.Assigned = &assigned
	var assignedByAnother bool = false
	this.AssignedByAnother = &assignedByAnother
	var reserved bool = false
	this.Reserved = &reserved
	return &this
}

// NewPoolAbstractIdPoolMemberWithDefaults instantiates a new PoolAbstractIdPoolMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractIdPoolMemberWithDefaults() *PoolAbstractIdPoolMember {
	this := PoolAbstractIdPoolMember{}
	var assignedByAnother bool = false
	this.AssignedByAnother = &assignedByAnother
	var reserved bool = false
	this.Reserved = &reserved
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractIdPoolMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractIdPoolMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractIdPoolMember) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractIdPoolMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractIdPoolMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractIdPoolMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssignedByAnother returns the AssignedByAnother field value if set, zero value otherwise.
func (o *PoolAbstractIdPoolMember) GetAssignedByAnother() bool {
	if o == nil || IsNil(o.AssignedByAnother) {
		var ret bool
		return ret
	}
	return *o.AssignedByAnother
}

// GetAssignedByAnotherOk returns a tuple with the AssignedByAnother field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractIdPoolMember) GetAssignedByAnotherOk() (*bool, bool) {
	if o == nil || IsNil(o.AssignedByAnother) {
		return nil, false
	}
	return o.AssignedByAnother, true
}

// HasAssignedByAnother returns a boolean if a field has been set.
func (o *PoolAbstractIdPoolMember) HasAssignedByAnother() bool {
	if o != nil && !IsNil(o.AssignedByAnother) {
		return true
	}

	return false
}

// SetAssignedByAnother gets a reference to the given bool and assigns it to the AssignedByAnother field.
func (o *PoolAbstractIdPoolMember) SetAssignedByAnother(v bool) {
	o.AssignedByAnother = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *PoolAbstractIdPoolMember) GetReserved() bool {
	if o == nil || IsNil(o.Reserved) {
		var ret bool
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractIdPoolMember) GetReservedOk() (*bool, bool) {
	if o == nil || IsNil(o.Reserved) {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *PoolAbstractIdPoolMember) HasReserved() bool {
	if o != nil && !IsNil(o.Reserved) {
		return true
	}

	return false
}

// SetReserved gets a reference to the given bool and assigns it to the Reserved field.
func (o *PoolAbstractIdPoolMember) SetReserved(v bool) {
	o.Reserved = &v
}

func (o PoolAbstractIdPoolMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolAbstractIdPoolMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPoolMember, errPoolAbstractPoolMember := json.Marshal(o.PoolAbstractPoolMember)
	if errPoolAbstractPoolMember != nil {
		return map[string]interface{}{}, errPoolAbstractPoolMember
	}
	errPoolAbstractPoolMember = json.Unmarshal([]byte(serializedPoolAbstractPoolMember), &toSerialize)
	if errPoolAbstractPoolMember != nil {
		return map[string]interface{}{}, errPoolAbstractPoolMember
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AssignedByAnother) {
		toSerialize["AssignedByAnother"] = o.AssignedByAnother
	}
	if !IsNil(o.Reserved) {
		toSerialize["Reserved"] = o.Reserved
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PoolAbstractIdPoolMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type PoolAbstractIdPoolMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Boolean to represent whether the ID is used either statically or by another pool.
		AssignedByAnother *bool `json:"AssignedByAnother,omitempty"`
		// Boolean to represent whether the ID is reserved.
		Reserved *bool `json:"Reserved,omitempty"`
	}

	varPoolAbstractIdPoolMemberWithoutEmbeddedStruct := PoolAbstractIdPoolMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPoolAbstractIdPoolMemberWithoutEmbeddedStruct)
	if err == nil {
		varPoolAbstractIdPoolMember := _PoolAbstractIdPoolMember{}
		varPoolAbstractIdPoolMember.ClassId = varPoolAbstractIdPoolMemberWithoutEmbeddedStruct.ClassId
		varPoolAbstractIdPoolMember.ObjectType = varPoolAbstractIdPoolMemberWithoutEmbeddedStruct.ObjectType
		varPoolAbstractIdPoolMember.AssignedByAnother = varPoolAbstractIdPoolMemberWithoutEmbeddedStruct.AssignedByAnother
		varPoolAbstractIdPoolMember.Reserved = varPoolAbstractIdPoolMemberWithoutEmbeddedStruct.Reserved
		*o = PoolAbstractIdPoolMember(varPoolAbstractIdPoolMember)
	} else {
		return err
	}

	varPoolAbstractIdPoolMember := _PoolAbstractIdPoolMember{}

	err = json.Unmarshal(data, &varPoolAbstractIdPoolMember)
	if err == nil {
		o.PoolAbstractPoolMember = varPoolAbstractIdPoolMember.PoolAbstractPoolMember
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssignedByAnother")
		delete(additionalProperties, "Reserved")

		// remove fields from embedded structs
		reflectPoolAbstractPoolMember := reflect.ValueOf(o.PoolAbstractPoolMember)
		for i := 0; i < reflectPoolAbstractPoolMember.Type().NumField(); i++ {
			t := reflectPoolAbstractPoolMember.Type().Field(i)

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

type NullablePoolAbstractIdPoolMember struct {
	value *PoolAbstractIdPoolMember
	isSet bool
}

func (v NullablePoolAbstractIdPoolMember) Get() *PoolAbstractIdPoolMember {
	return v.value
}

func (v *NullablePoolAbstractIdPoolMember) Set(val *PoolAbstractIdPoolMember) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractIdPoolMember) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractIdPoolMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractIdPoolMember(val *PoolAbstractIdPoolMember) *NullablePoolAbstractIdPoolMember {
	return &NullablePoolAbstractIdPoolMember{value: val, isSet: true}
}

func (v NullablePoolAbstractIdPoolMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractIdPoolMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
