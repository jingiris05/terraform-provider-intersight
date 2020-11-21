/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// IamSecurityHolderAllOf Definition of the list of properties defined in 'iam.SecurityHolder', excluding properties defined in parent classes.
type IamSecurityHolderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                             `json:"ObjectType"`
	Account              *IamAccountRelationship            `json:"Account,omitempty"`
	IpRulesConfiguration *IamIpAccessManagementRelationship `json:"IpRulesConfiguration,omitempty"`
	// An array of relationships to iamResourcePermission resources.
	ResourcePermissions  []IamResourcePermissionRelationship `json:"ResourcePermissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamSecurityHolderAllOf IamSecurityHolderAllOf

// NewIamSecurityHolderAllOf instantiates a new IamSecurityHolderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSecurityHolderAllOf(classId string, objectType string) *IamSecurityHolderAllOf {
	this := IamSecurityHolderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamSecurityHolderAllOfWithDefaults instantiates a new IamSecurityHolderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSecurityHolderAllOfWithDefaults() *IamSecurityHolderAllOf {
	this := IamSecurityHolderAllOf{}
	var classId string = "iam.SecurityHolder"
	this.ClassId = classId
	var objectType string = "iam.SecurityHolder"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamSecurityHolderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamSecurityHolderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamSecurityHolderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamSecurityHolderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamSecurityHolderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamSecurityHolderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamSecurityHolderAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSecurityHolderAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamSecurityHolderAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamSecurityHolderAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetIpRulesConfiguration returns the IpRulesConfiguration field value if set, zero value otherwise.
func (o *IamSecurityHolderAllOf) GetIpRulesConfiguration() IamIpAccessManagementRelationship {
	if o == nil || o.IpRulesConfiguration == nil {
		var ret IamIpAccessManagementRelationship
		return ret
	}
	return *o.IpRulesConfiguration
}

// GetIpRulesConfigurationOk returns a tuple with the IpRulesConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSecurityHolderAllOf) GetIpRulesConfigurationOk() (*IamIpAccessManagementRelationship, bool) {
	if o == nil || o.IpRulesConfiguration == nil {
		return nil, false
	}
	return o.IpRulesConfiguration, true
}

// HasIpRulesConfiguration returns a boolean if a field has been set.
func (o *IamSecurityHolderAllOf) HasIpRulesConfiguration() bool {
	if o != nil && o.IpRulesConfiguration != nil {
		return true
	}

	return false
}

// SetIpRulesConfiguration gets a reference to the given IamIpAccessManagementRelationship and assigns it to the IpRulesConfiguration field.
func (o *IamSecurityHolderAllOf) SetIpRulesConfiguration(v IamIpAccessManagementRelationship) {
	o.IpRulesConfiguration = &v
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamSecurityHolderAllOf) GetResourcePermissions() []IamResourcePermissionRelationship {
	if o == nil {
		var ret []IamResourcePermissionRelationship
		return ret
	}
	return o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamSecurityHolderAllOf) GetResourcePermissionsOk() (*[]IamResourcePermissionRelationship, bool) {
	if o == nil || o.ResourcePermissions == nil {
		return nil, false
	}
	return &o.ResourcePermissions, true
}

// HasResourcePermissions returns a boolean if a field has been set.
func (o *IamSecurityHolderAllOf) HasResourcePermissions() bool {
	if o != nil && o.ResourcePermissions != nil {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given []IamResourcePermissionRelationship and assigns it to the ResourcePermissions field.
func (o *IamSecurityHolderAllOf) SetResourcePermissions(v []IamResourcePermissionRelationship) {
	o.ResourcePermissions = v
}

func (o IamSecurityHolderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.IpRulesConfiguration != nil {
		toSerialize["IpRulesConfiguration"] = o.IpRulesConfiguration
	}
	if o.ResourcePermissions != nil {
		toSerialize["ResourcePermissions"] = o.ResourcePermissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamSecurityHolderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamSecurityHolderAllOf := _IamSecurityHolderAllOf{}

	if err = json.Unmarshal(bytes, &varIamSecurityHolderAllOf); err == nil {
		*o = IamSecurityHolderAllOf(varIamSecurityHolderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "IpRulesConfiguration")
		delete(additionalProperties, "ResourcePermissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamSecurityHolderAllOf struct {
	value *IamSecurityHolderAllOf
	isSet bool
}

func (v NullableIamSecurityHolderAllOf) Get() *IamSecurityHolderAllOf {
	return v.value
}

func (v *NullableIamSecurityHolderAllOf) Set(val *IamSecurityHolderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSecurityHolderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSecurityHolderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSecurityHolderAllOf(val *IamSecurityHolderAllOf) *NullableIamSecurityHolderAllOf {
	return &NullableIamSecurityHolderAllOf{value: val, isSet: true}
}

func (v NullableIamSecurityHolderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSecurityHolderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
