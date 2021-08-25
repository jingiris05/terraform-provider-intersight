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
)

// CommAbstractHttpProxyPolicyAllOf Definition of the list of properties defined in 'comm.AbstractHttpProxyPolicy', excluding properties defined in parent classes.
type CommAbstractHttpProxyPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// HTTP Proxy server FQDN or IP.
	Hostname *string `json:"Hostname,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// The password for the HTTP Proxy.
	Password *string `json:"Password,omitempty"`
	// The HTTP Proxy port number. The port number of the HTTP proxy must be between 1 and 65535, inclusive.
	Port *int64 `json:"Port,omitempty"`
	// The username for the HTTP Proxy.
	Username             *string                               `json:"Username,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommAbstractHttpProxyPolicyAllOf CommAbstractHttpProxyPolicyAllOf

// NewCommAbstractHttpProxyPolicyAllOf instantiates a new CommAbstractHttpProxyPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommAbstractHttpProxyPolicyAllOf(classId string, objectType string) *CommAbstractHttpProxyPolicyAllOf {
	this := CommAbstractHttpProxyPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// NewCommAbstractHttpProxyPolicyAllOfWithDefaults instantiates a new CommAbstractHttpProxyPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommAbstractHttpProxyPolicyAllOfWithDefaults() *CommAbstractHttpProxyPolicyAllOf {
	this := CommAbstractHttpProxyPolicyAllOf{}
	var classId string = "comm.HttpProxyPolicy"
	this.ClassId = classId
	var objectType string = "comm.HttpProxyPolicy"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *CommAbstractHttpProxyPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CommAbstractHttpProxyPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CommAbstractHttpProxyPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CommAbstractHttpProxyPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicyAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *CommAbstractHttpProxyPolicyAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicyAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *CommAbstractHttpProxyPolicyAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicyAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CommAbstractHttpProxyPolicyAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicyAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *CommAbstractHttpProxyPolicyAllOf) SetPort(v int64) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicyAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CommAbstractHttpProxyPolicyAllOf) SetUsername(v string) {
	o.Username = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CommAbstractHttpProxyPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CommAbstractHttpProxyPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *CommAbstractHttpProxyPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o CommAbstractHttpProxyPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommAbstractHttpProxyPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCommAbstractHttpProxyPolicyAllOf := _CommAbstractHttpProxyPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varCommAbstractHttpProxyPolicyAllOf); err == nil {
		*o = CommAbstractHttpProxyPolicyAllOf(varCommAbstractHttpProxyPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommAbstractHttpProxyPolicyAllOf struct {
	value *CommAbstractHttpProxyPolicyAllOf
	isSet bool
}

func (v NullableCommAbstractHttpProxyPolicyAllOf) Get() *CommAbstractHttpProxyPolicyAllOf {
	return v.value
}

func (v *NullableCommAbstractHttpProxyPolicyAllOf) Set(val *CommAbstractHttpProxyPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommAbstractHttpProxyPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommAbstractHttpProxyPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommAbstractHttpProxyPolicyAllOf(val *CommAbstractHttpProxyPolicyAllOf) *NullableCommAbstractHttpProxyPolicyAllOf {
	return &NullableCommAbstractHttpProxyPolicyAllOf{value: val, isSet: true}
}

func (v NullableCommAbstractHttpProxyPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommAbstractHttpProxyPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
