/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14237
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IppoolReservationReference The reference to the reservation object.
type IppoolReservationReference struct {
	PoolReservationReference
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The consumer name for which the reserved IP would be used.
	ConsumerName *string `json:"ConsumerName,omitempty"`
	// The consumer type for which the reserved IP would be used. * `OutofbandIpv4-Access` - IP reservation would be used for out of band management. * `InbandIpv4-Access` - IP reservation would be used for inband management. * `InbandIpv6-Access` - IP reservation would be used for inband management. * `ISCSI` - IP reservation would be used for ISCSI management.
	ConsumerType         *string `json:"ConsumerType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolReservationReference IppoolReservationReference

// NewIppoolReservationReference instantiates a new IppoolReservationReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolReservationReference(classId string, objectType string) *IppoolReservationReference {
	this := IppoolReservationReference{}
	this.ClassId = classId
	this.ObjectType = objectType
	var consumerType string = "OutofbandIpv4-Access"
	this.ConsumerType = &consumerType
	return &this
}

// NewIppoolReservationReferenceWithDefaults instantiates a new IppoolReservationReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolReservationReferenceWithDefaults() *IppoolReservationReference {
	this := IppoolReservationReference{}
	var classId string = "ippool.ReservationReference"
	this.ClassId = classId
	var objectType string = "ippool.ReservationReference"
	this.ObjectType = objectType
	var consumerType string = "OutofbandIpv4-Access"
	this.ConsumerType = &consumerType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolReservationReference) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolReservationReference) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolReservationReference) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IppoolReservationReference) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolReservationReference) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolReservationReference) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConsumerName returns the ConsumerName field value if set, zero value otherwise.
func (o *IppoolReservationReference) GetConsumerName() string {
	if o == nil || o.ConsumerName == nil {
		var ret string
		return ret
	}
	return *o.ConsumerName
}

// GetConsumerNameOk returns a tuple with the ConsumerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationReference) GetConsumerNameOk() (*string, bool) {
	if o == nil || o.ConsumerName == nil {
		return nil, false
	}
	return o.ConsumerName, true
}

// HasConsumerName returns a boolean if a field has been set.
func (o *IppoolReservationReference) HasConsumerName() bool {
	if o != nil && o.ConsumerName != nil {
		return true
	}

	return false
}

// SetConsumerName gets a reference to the given string and assigns it to the ConsumerName field.
func (o *IppoolReservationReference) SetConsumerName(v string) {
	o.ConsumerName = &v
}

// GetConsumerType returns the ConsumerType field value if set, zero value otherwise.
func (o *IppoolReservationReference) GetConsumerType() string {
	if o == nil || o.ConsumerType == nil {
		var ret string
		return ret
	}
	return *o.ConsumerType
}

// GetConsumerTypeOk returns a tuple with the ConsumerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationReference) GetConsumerTypeOk() (*string, bool) {
	if o == nil || o.ConsumerType == nil {
		return nil, false
	}
	return o.ConsumerType, true
}

// HasConsumerType returns a boolean if a field has been set.
func (o *IppoolReservationReference) HasConsumerType() bool {
	if o != nil && o.ConsumerType != nil {
		return true
	}

	return false
}

// SetConsumerType gets a reference to the given string and assigns it to the ConsumerType field.
func (o *IppoolReservationReference) SetConsumerType(v string) {
	o.ConsumerType = &v
}

func (o IppoolReservationReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolReservationReference, errPoolReservationReference := json.Marshal(o.PoolReservationReference)
	if errPoolReservationReference != nil {
		return []byte{}, errPoolReservationReference
	}
	errPoolReservationReference = json.Unmarshal([]byte(serializedPoolReservationReference), &toSerialize)
	if errPoolReservationReference != nil {
		return []byte{}, errPoolReservationReference
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConsumerName != nil {
		toSerialize["ConsumerName"] = o.ConsumerName
	}
	if o.ConsumerType != nil {
		toSerialize["ConsumerType"] = o.ConsumerType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolReservationReference) UnmarshalJSON(bytes []byte) (err error) {
	type IppoolReservationReferenceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The consumer name for which the reserved IP would be used.
		ConsumerName *string `json:"ConsumerName,omitempty"`
		// The consumer type for which the reserved IP would be used. * `OutofbandIpv4-Access` - IP reservation would be used for out of band management. * `InbandIpv4-Access` - IP reservation would be used for inband management. * `InbandIpv6-Access` - IP reservation would be used for inband management. * `ISCSI` - IP reservation would be used for ISCSI management.
		ConsumerType *string `json:"ConsumerType,omitempty"`
	}

	varIppoolReservationReferenceWithoutEmbeddedStruct := IppoolReservationReferenceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIppoolReservationReferenceWithoutEmbeddedStruct)
	if err == nil {
		varIppoolReservationReference := _IppoolReservationReference{}
		varIppoolReservationReference.ClassId = varIppoolReservationReferenceWithoutEmbeddedStruct.ClassId
		varIppoolReservationReference.ObjectType = varIppoolReservationReferenceWithoutEmbeddedStruct.ObjectType
		varIppoolReservationReference.ConsumerName = varIppoolReservationReferenceWithoutEmbeddedStruct.ConsumerName
		varIppoolReservationReference.ConsumerType = varIppoolReservationReferenceWithoutEmbeddedStruct.ConsumerType
		*o = IppoolReservationReference(varIppoolReservationReference)
	} else {
		return err
	}

	varIppoolReservationReference := _IppoolReservationReference{}

	err = json.Unmarshal(bytes, &varIppoolReservationReference)
	if err == nil {
		o.PoolReservationReference = varIppoolReservationReference.PoolReservationReference
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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

type NullableIppoolReservationReference struct {
	value *IppoolReservationReference
	isSet bool
}

func (v NullableIppoolReservationReference) Get() *IppoolReservationReference {
	return v.value
}

func (v *NullableIppoolReservationReference) Set(val *IppoolReservationReference) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolReservationReference) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolReservationReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolReservationReference(val *IppoolReservationReference) *NullableIppoolReservationReference {
	return &NullableIppoolReservationReference{value: val, isSet: true}
}

func (v NullableIppoolReservationReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolReservationReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
