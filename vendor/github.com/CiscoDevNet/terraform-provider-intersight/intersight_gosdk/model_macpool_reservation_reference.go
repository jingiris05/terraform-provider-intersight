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

// checks if the MacpoolReservationReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MacpoolReservationReference{}

// MacpoolReservationReference The reference to the reservation object.
type MacpoolReservationReference struct {
	PoolReservationReference
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The consumer name for which the reserved MAC would be used.
	ConsumerName *string `json:"ConsumerName,omitempty"`
	// The consumer type for which the reserved MAC would be used. * `Vnic` - MAC reservation would be used by VNIC.
	ConsumerType         *string `json:"ConsumerType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MacpoolReservationReference MacpoolReservationReference

// NewMacpoolReservationReference instantiates a new MacpoolReservationReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacpoolReservationReference(classId string, objectType string) *MacpoolReservationReference {
	this := MacpoolReservationReference{}
	this.ClassId = classId
	this.ObjectType = objectType
	var consumerType string = "Vnic"
	this.ConsumerType = &consumerType
	return &this
}

// NewMacpoolReservationReferenceWithDefaults instantiates a new MacpoolReservationReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacpoolReservationReferenceWithDefaults() *MacpoolReservationReference {
	this := MacpoolReservationReference{}
	var classId string = "macpool.ReservationReference"
	this.ClassId = classId
	var objectType string = "macpool.ReservationReference"
	this.ObjectType = objectType
	var consumerType string = "Vnic"
	this.ConsumerType = &consumerType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MacpoolReservationReference) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MacpoolReservationReference) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MacpoolReservationReference) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "macpool.ReservationReference" of the ClassId field.
func (o *MacpoolReservationReference) GetDefaultClassId() interface{} {
	return "macpool.ReservationReference"
}

// GetObjectType returns the ObjectType field value
func (o *MacpoolReservationReference) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MacpoolReservationReference) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MacpoolReservationReference) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "macpool.ReservationReference" of the ObjectType field.
func (o *MacpoolReservationReference) GetDefaultObjectType() interface{} {
	return "macpool.ReservationReference"
}

// GetConsumerName returns the ConsumerName field value if set, zero value otherwise.
func (o *MacpoolReservationReference) GetConsumerName() string {
	if o == nil || IsNil(o.ConsumerName) {
		var ret string
		return ret
	}
	return *o.ConsumerName
}

// GetConsumerNameOk returns a tuple with the ConsumerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolReservationReference) GetConsumerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerName) {
		return nil, false
	}
	return o.ConsumerName, true
}

// HasConsumerName returns a boolean if a field has been set.
func (o *MacpoolReservationReference) HasConsumerName() bool {
	if o != nil && !IsNil(o.ConsumerName) {
		return true
	}

	return false
}

// SetConsumerName gets a reference to the given string and assigns it to the ConsumerName field.
func (o *MacpoolReservationReference) SetConsumerName(v string) {
	o.ConsumerName = &v
}

// GetConsumerType returns the ConsumerType field value if set, zero value otherwise.
func (o *MacpoolReservationReference) GetConsumerType() string {
	if o == nil || IsNil(o.ConsumerType) {
		var ret string
		return ret
	}
	return *o.ConsumerType
}

// GetConsumerTypeOk returns a tuple with the ConsumerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolReservationReference) GetConsumerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerType) {
		return nil, false
	}
	return o.ConsumerType, true
}

// HasConsumerType returns a boolean if a field has been set.
func (o *MacpoolReservationReference) HasConsumerType() bool {
	if o != nil && !IsNil(o.ConsumerType) {
		return true
	}

	return false
}

// SetConsumerType gets a reference to the given string and assigns it to the ConsumerType field.
func (o *MacpoolReservationReference) SetConsumerType(v string) {
	o.ConsumerType = &v
}

func (o MacpoolReservationReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MacpoolReservationReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolReservationReference, errPoolReservationReference := json.Marshal(o.PoolReservationReference)
	if errPoolReservationReference != nil {
		return map[string]interface{}{}, errPoolReservationReference
	}
	errPoolReservationReference = json.Unmarshal([]byte(serializedPoolReservationReference), &toSerialize)
	if errPoolReservationReference != nil {
		return map[string]interface{}{}, errPoolReservationReference
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ConsumerName) {
		toSerialize["ConsumerName"] = o.ConsumerName
	}
	if !IsNil(o.ConsumerType) {
		toSerialize["ConsumerType"] = o.ConsumerType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MacpoolReservationReference) UnmarshalJSON(data []byte) (err error) {
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
	type MacpoolReservationReferenceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The consumer name for which the reserved MAC would be used.
		ConsumerName *string `json:"ConsumerName,omitempty"`
		// The consumer type for which the reserved MAC would be used. * `Vnic` - MAC reservation would be used by VNIC.
		ConsumerType *string `json:"ConsumerType,omitempty"`
	}

	varMacpoolReservationReferenceWithoutEmbeddedStruct := MacpoolReservationReferenceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMacpoolReservationReferenceWithoutEmbeddedStruct)
	if err == nil {
		varMacpoolReservationReference := _MacpoolReservationReference{}
		varMacpoolReservationReference.ClassId = varMacpoolReservationReferenceWithoutEmbeddedStruct.ClassId
		varMacpoolReservationReference.ObjectType = varMacpoolReservationReferenceWithoutEmbeddedStruct.ObjectType
		varMacpoolReservationReference.ConsumerName = varMacpoolReservationReferenceWithoutEmbeddedStruct.ConsumerName
		varMacpoolReservationReference.ConsumerType = varMacpoolReservationReferenceWithoutEmbeddedStruct.ConsumerType
		*o = MacpoolReservationReference(varMacpoolReservationReference)
	} else {
		return err
	}

	varMacpoolReservationReference := _MacpoolReservationReference{}

	err = json.Unmarshal(data, &varMacpoolReservationReference)
	if err == nil {
		o.PoolReservationReference = varMacpoolReservationReference.PoolReservationReference
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConsumerName")
		delete(additionalProperties, "ConsumerType")

		// remove fields from embedded structs
		reflectPoolReservationReference := reflect.ValueOf(o.PoolReservationReference)
		for i := 0; i < reflectPoolReservationReference.Type().NumField(); i++ {
			t := reflectPoolReservationReference.Type().Field(i)

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

type NullableMacpoolReservationReference struct {
	value *MacpoolReservationReference
	isSet bool
}

func (v NullableMacpoolReservationReference) Get() *MacpoolReservationReference {
	return v.value
}

func (v *NullableMacpoolReservationReference) Set(val *MacpoolReservationReference) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolReservationReference) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolReservationReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolReservationReference(val *MacpoolReservationReference) *NullableMacpoolReservationReference {
	return &NullableMacpoolReservationReference{value: val, isSet: true}
}

func (v NullableMacpoolReservationReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolReservationReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
