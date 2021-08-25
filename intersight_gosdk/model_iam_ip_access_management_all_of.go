/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-24T21:18:00Z.
 *
 * API version: 1.0.9-4404
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// IamIpAccessManagementAllOf Definition of the list of properties defined in 'iam.IpAccessManagement', excluding properties defined in parent classes.
type IamIpAccessManagementAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Flag stores the state of IP address based access management. Access management is enabled when it's true.
	Enable *bool `json:"Enable,omitempty"`
	// The access to account gets locked out if wrong IP addresses are configured. Account Administrators have privilege to unblock the account. It stores the time when the account was last recovered from lock out.
	LastRecoveryTime *time.Time                     `json:"LastRecoveryTime,omitempty"`
	Holder           *IamSecurityHolderRelationship `json:"Holder,omitempty"`
	// An array of relationships to iamIpAddress resources.
	IpAddresses          []IamIpAddressRelationship `json:"IpAddresses,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamIpAccessManagementAllOf IamIpAccessManagementAllOf

// NewIamIpAccessManagementAllOf instantiates a new IamIpAccessManagementAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamIpAccessManagementAllOf(classId string, objectType string) *IamIpAccessManagementAllOf {
	this := IamIpAccessManagementAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamIpAccessManagementAllOfWithDefaults instantiates a new IamIpAccessManagementAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamIpAccessManagementAllOfWithDefaults() *IamIpAccessManagementAllOf {
	this := IamIpAccessManagementAllOf{}
	var classId string = "iam.IpAccessManagement"
	this.ClassId = classId
	var objectType string = "iam.IpAccessManagement"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamIpAccessManagementAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamIpAccessManagementAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamIpAccessManagementAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamIpAccessManagementAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamIpAccessManagementAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamIpAccessManagementAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *IamIpAccessManagementAllOf) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAccessManagementAllOf) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *IamIpAccessManagementAllOf) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *IamIpAccessManagementAllOf) SetEnable(v bool) {
	o.Enable = &v
}

// GetLastRecoveryTime returns the LastRecoveryTime field value if set, zero value otherwise.
func (o *IamIpAccessManagementAllOf) GetLastRecoveryTime() time.Time {
	if o == nil || o.LastRecoveryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRecoveryTime
}

// GetLastRecoveryTimeOk returns a tuple with the LastRecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAccessManagementAllOf) GetLastRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || o.LastRecoveryTime == nil {
		return nil, false
	}
	return o.LastRecoveryTime, true
}

// HasLastRecoveryTime returns a boolean if a field has been set.
func (o *IamIpAccessManagementAllOf) HasLastRecoveryTime() bool {
	if o != nil && o.LastRecoveryTime != nil {
		return true
	}

	return false
}

// SetLastRecoveryTime gets a reference to the given time.Time and assigns it to the LastRecoveryTime field.
func (o *IamIpAccessManagementAllOf) SetLastRecoveryTime(v time.Time) {
	o.LastRecoveryTime = &v
}

// GetHolder returns the Holder field value if set, zero value otherwise.
func (o *IamIpAccessManagementAllOf) GetHolder() IamSecurityHolderRelationship {
	if o == nil || o.Holder == nil {
		var ret IamSecurityHolderRelationship
		return ret
	}
	return *o.Holder
}

// GetHolderOk returns a tuple with the Holder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAccessManagementAllOf) GetHolderOk() (*IamSecurityHolderRelationship, bool) {
	if o == nil || o.Holder == nil {
		return nil, false
	}
	return o.Holder, true
}

// HasHolder returns a boolean if a field has been set.
func (o *IamIpAccessManagementAllOf) HasHolder() bool {
	if o != nil && o.Holder != nil {
		return true
	}

	return false
}

// SetHolder gets a reference to the given IamSecurityHolderRelationship and assigns it to the Holder field.
func (o *IamIpAccessManagementAllOf) SetHolder(v IamSecurityHolderRelationship) {
	o.Holder = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIpAccessManagementAllOf) GetIpAddresses() []IamIpAddressRelationship {
	if o == nil {
		var ret []IamIpAddressRelationship
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIpAccessManagementAllOf) GetIpAddressesOk() (*[]IamIpAddressRelationship, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *IamIpAccessManagementAllOf) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []IamIpAddressRelationship and assigns it to the IpAddresses field.
func (o *IamIpAccessManagementAllOf) SetIpAddresses(v []IamIpAddressRelationship) {
	o.IpAddresses = v
}

func (o IamIpAccessManagementAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enable != nil {
		toSerialize["Enable"] = o.Enable
	}
	if o.LastRecoveryTime != nil {
		toSerialize["LastRecoveryTime"] = o.LastRecoveryTime
	}
	if o.Holder != nil {
		toSerialize["Holder"] = o.Holder
	}
	if o.IpAddresses != nil {
		toSerialize["IpAddresses"] = o.IpAddresses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamIpAccessManagementAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamIpAccessManagementAllOf := _IamIpAccessManagementAllOf{}

	if err = json.Unmarshal(bytes, &varIamIpAccessManagementAllOf); err == nil {
		*o = IamIpAccessManagementAllOf(varIamIpAccessManagementAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enable")
		delete(additionalProperties, "LastRecoveryTime")
		delete(additionalProperties, "Holder")
		delete(additionalProperties, "IpAddresses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamIpAccessManagementAllOf struct {
	value *IamIpAccessManagementAllOf
	isSet bool
}

func (v NullableIamIpAccessManagementAllOf) Get() *IamIpAccessManagementAllOf {
	return v.value
}

func (v *NullableIamIpAccessManagementAllOf) Set(val *IamIpAccessManagementAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamIpAccessManagementAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamIpAccessManagementAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamIpAccessManagementAllOf(val *IamIpAccessManagementAllOf) *NullableIamIpAccessManagementAllOf {
	return &NullableIamIpAccessManagementAllOf{value: val, isSet: true}
}

func (v NullableIamIpAccessManagementAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamIpAccessManagementAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
