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
	"reflect"
	"strings"
)

// AaaRetentionPolicy An account level policy specifying the period for the audit log retention.
type AaaRetentionPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time period in months for audit log retention. Audit logs beyond this period will be automatically deleted.
	RetentionPeriod      *int64                  `json:"RetentionPeriod,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AaaRetentionPolicy AaaRetentionPolicy

// NewAaaRetentionPolicy instantiates a new AaaRetentionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAaaRetentionPolicy(classId string, objectType string) *AaaRetentionPolicy {
	this := AaaRetentionPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAaaRetentionPolicyWithDefaults instantiates a new AaaRetentionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAaaRetentionPolicyWithDefaults() *AaaRetentionPolicy {
	this := AaaRetentionPolicy{}
	var classId string = "aaa.RetentionPolicy"
	this.ClassId = classId
	var objectType string = "aaa.RetentionPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AaaRetentionPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AaaRetentionPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AaaRetentionPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AaaRetentionPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AaaRetentionPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AaaRetentionPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *AaaRetentionPolicy) GetRetentionPeriod() int64 {
	if o == nil || o.RetentionPeriod == nil {
		var ret int64
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaRetentionPolicy) GetRetentionPeriodOk() (*int64, bool) {
	if o == nil || o.RetentionPeriod == nil {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *AaaRetentionPolicy) HasRetentionPeriod() bool {
	if o != nil && o.RetentionPeriod != nil {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given int64 and assigns it to the RetentionPeriod field.
func (o *AaaRetentionPolicy) SetRetentionPeriod(v int64) {
	o.RetentionPeriod = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AaaRetentionPolicy) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaRetentionPolicy) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AaaRetentionPolicy) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *AaaRetentionPolicy) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o AaaRetentionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.RetentionPeriod != nil {
		toSerialize["RetentionPeriod"] = o.RetentionPeriod
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AaaRetentionPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type AaaRetentionPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The time period in months for audit log retention. Audit logs beyond this period will be automatically deleted.
		RetentionPeriod *int64                  `json:"RetentionPeriod,omitempty"`
		Account         *IamAccountRelationship `json:"Account,omitempty"`
	}

	varAaaRetentionPolicyWithoutEmbeddedStruct := AaaRetentionPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAaaRetentionPolicyWithoutEmbeddedStruct)
	if err == nil {
		varAaaRetentionPolicy := _AaaRetentionPolicy{}
		varAaaRetentionPolicy.ClassId = varAaaRetentionPolicyWithoutEmbeddedStruct.ClassId
		varAaaRetentionPolicy.ObjectType = varAaaRetentionPolicyWithoutEmbeddedStruct.ObjectType
		varAaaRetentionPolicy.RetentionPeriod = varAaaRetentionPolicyWithoutEmbeddedStruct.RetentionPeriod
		varAaaRetentionPolicy.Account = varAaaRetentionPolicyWithoutEmbeddedStruct.Account
		*o = AaaRetentionPolicy(varAaaRetentionPolicy)
	} else {
		return err
	}

	varAaaRetentionPolicy := _AaaRetentionPolicy{}

	err = json.Unmarshal(bytes, &varAaaRetentionPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varAaaRetentionPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RetentionPeriod")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableAaaRetentionPolicy struct {
	value *AaaRetentionPolicy
	isSet bool
}

func (v NullableAaaRetentionPolicy) Get() *AaaRetentionPolicy {
	return v.value
}

func (v *NullableAaaRetentionPolicy) Set(val *AaaRetentionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAaaRetentionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAaaRetentionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAaaRetentionPolicy(val *AaaRetentionPolicy) *NullableAaaRetentionPolicy {
	return &NullableAaaRetentionPolicy{value: val, isSet: true}
}

func (v NullableAaaRetentionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAaaRetentionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
