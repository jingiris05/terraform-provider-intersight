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

// KubernetesOvsBond An OpenVSwitch bonded network interface.
type KubernetesOvsBond struct {
	KubernetesNetworkInterface
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string   `json:"ObjectType"`
	Interfaces []string `json:"Interfaces,omitempty"`
	// Native VLAN for to use for the bond.
	Vlan                 *int64 `json:"Vlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesOvsBond KubernetesOvsBond

// NewKubernetesOvsBond instantiates a new KubernetesOvsBond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesOvsBond(classId string, objectType string) *KubernetesOvsBond {
	this := KubernetesOvsBond{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesOvsBondWithDefaults instantiates a new KubernetesOvsBond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesOvsBondWithDefaults() *KubernetesOvsBond {
	this := KubernetesOvsBond{}
	var classId string = "kubernetes.OvsBond"
	this.ClassId = classId
	var objectType string = "kubernetes.OvsBond"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesOvsBond) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesOvsBond) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesOvsBond) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesOvsBond) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesOvsBond) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesOvsBond) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesOvsBond) GetInterfaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesOvsBond) GetInterfacesOk() ([]string, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *KubernetesOvsBond) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []string and assigns it to the Interfaces field.
func (o *KubernetesOvsBond) SetInterfaces(v []string) {
	o.Interfaces = v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *KubernetesOvsBond) GetVlan() int64 {
	if o == nil || o.Vlan == nil {
		var ret int64
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesOvsBond) GetVlanOk() (*int64, bool) {
	if o == nil || o.Vlan == nil {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *KubernetesOvsBond) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int64 and assigns it to the Vlan field.
func (o *KubernetesOvsBond) SetVlan(v int64) {
	o.Vlan = &v
}

func (o KubernetesOvsBond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesNetworkInterface, errKubernetesNetworkInterface := json.Marshal(o.KubernetesNetworkInterface)
	if errKubernetesNetworkInterface != nil {
		return []byte{}, errKubernetesNetworkInterface
	}
	errKubernetesNetworkInterface = json.Unmarshal([]byte(serializedKubernetesNetworkInterface), &toSerialize)
	if errKubernetesNetworkInterface != nil {
		return []byte{}, errKubernetesNetworkInterface
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Interfaces != nil {
		toSerialize["Interfaces"] = o.Interfaces
	}
	if o.Vlan != nil {
		toSerialize["Vlan"] = o.Vlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesOvsBond) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesOvsBondWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string   `json:"ObjectType"`
		Interfaces []string `json:"Interfaces,omitempty"`
		// Native VLAN for to use for the bond.
		Vlan *int64 `json:"Vlan,omitempty"`
	}

	varKubernetesOvsBondWithoutEmbeddedStruct := KubernetesOvsBondWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesOvsBondWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesOvsBond := _KubernetesOvsBond{}
		varKubernetesOvsBond.ClassId = varKubernetesOvsBondWithoutEmbeddedStruct.ClassId
		varKubernetesOvsBond.ObjectType = varKubernetesOvsBondWithoutEmbeddedStruct.ObjectType
		varKubernetesOvsBond.Interfaces = varKubernetesOvsBondWithoutEmbeddedStruct.Interfaces
		varKubernetesOvsBond.Vlan = varKubernetesOvsBondWithoutEmbeddedStruct.Vlan
		*o = KubernetesOvsBond(varKubernetesOvsBond)
	} else {
		return err
	}

	varKubernetesOvsBond := _KubernetesOvsBond{}

	err = json.Unmarshal(bytes, &varKubernetesOvsBond)
	if err == nil {
		o.KubernetesNetworkInterface = varKubernetesOvsBond.KubernetesNetworkInterface
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Interfaces")
		delete(additionalProperties, "Vlan")

		// remove fields from embedded structs
		reflectKubernetesNetworkInterface := reflect.ValueOf(o.KubernetesNetworkInterface)
		for i := 0; i < reflectKubernetesNetworkInterface.Type().NumField(); i++ {
			t := reflectKubernetesNetworkInterface.Type().Field(i)

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

type NullableKubernetesOvsBond struct {
	value *KubernetesOvsBond
	isSet bool
}

func (v NullableKubernetesOvsBond) Get() *KubernetesOvsBond {
	return v.value
}

func (v *NullableKubernetesOvsBond) Set(val *KubernetesOvsBond) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesOvsBond) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesOvsBond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesOvsBond(val *KubernetesOvsBond) *NullableKubernetesOvsBond {
	return &NullableKubernetesOvsBond{value: val, isSet: true}
}

func (v NullableKubernetesOvsBond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesOvsBond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
