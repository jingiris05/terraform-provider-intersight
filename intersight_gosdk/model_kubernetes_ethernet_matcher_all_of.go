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

// KubernetesEthernetMatcherAllOf Definition of the list of properties defined in 'kubernetes.EthernetMatcher', excluding properties defined in parent classes.
type KubernetesEthernetMatcherAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Which property we should use to find the ethernet interface. * `Name` - A network interface name, e.g. eth0, eno9. * `MacAddress` - A network interface Mac Address.
	Type *string `json:"Type,omitempty"`
	// The value to match for the property specified by type.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesEthernetMatcherAllOf KubernetesEthernetMatcherAllOf

// NewKubernetesEthernetMatcherAllOf instantiates a new KubernetesEthernetMatcherAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesEthernetMatcherAllOf(classId string, objectType string) *KubernetesEthernetMatcherAllOf {
	this := KubernetesEthernetMatcherAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "Name"
	this.Type = &type_
	return &this
}

// NewKubernetesEthernetMatcherAllOfWithDefaults instantiates a new KubernetesEthernetMatcherAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesEthernetMatcherAllOfWithDefaults() *KubernetesEthernetMatcherAllOf {
	this := KubernetesEthernetMatcherAllOf{}
	var classId string = "kubernetes.EthernetMatcher"
	this.ClassId = classId
	var objectType string = "kubernetes.EthernetMatcher"
	this.ObjectType = objectType
	var type_ string = "Name"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesEthernetMatcherAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesEthernetMatcherAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesEthernetMatcherAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesEthernetMatcherAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesEthernetMatcherAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesEthernetMatcherAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KubernetesEthernetMatcherAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEthernetMatcherAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KubernetesEthernetMatcherAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KubernetesEthernetMatcherAllOf) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *KubernetesEthernetMatcherAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEthernetMatcherAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *KubernetesEthernetMatcherAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *KubernetesEthernetMatcherAllOf) SetValue(v string) {
	o.Value = &v
}

func (o KubernetesEthernetMatcherAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesEthernetMatcherAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesEthernetMatcherAllOf := _KubernetesEthernetMatcherAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesEthernetMatcherAllOf); err == nil {
		*o = KubernetesEthernetMatcherAllOf(varKubernetesEthernetMatcherAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesEthernetMatcherAllOf struct {
	value *KubernetesEthernetMatcherAllOf
	isSet bool
}

func (v NullableKubernetesEthernetMatcherAllOf) Get() *KubernetesEthernetMatcherAllOf {
	return v.value
}

func (v *NullableKubernetesEthernetMatcherAllOf) Set(val *KubernetesEthernetMatcherAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesEthernetMatcherAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesEthernetMatcherAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesEthernetMatcherAllOf(val *KubernetesEthernetMatcherAllOf) *NullableKubernetesEthernetMatcherAllOf {
	return &NullableKubernetesEthernetMatcherAllOf{value: val, isSet: true}
}

func (v NullableKubernetesEthernetMatcherAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesEthernetMatcherAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
