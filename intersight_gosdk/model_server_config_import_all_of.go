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

// ServerConfigImportAllOf Definition of the list of properties defined in 'server.ConfigImport', excluding properties defined in parent classes.
type ServerConfigImportAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the imported profile.
	Description *string `json:"Description,omitempty"`
	// Policy prefix for the policies of the imported server profile.
	PolicyPrefix *string  `json:"PolicyPrefix,omitempty"`
	PolicyTypes  []string `json:"PolicyTypes,omitempty"`
	// Profile name for the imported server profile.
	ProfileName          *string                               `json:"ProfileName,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Server               *ComputeRackUnitRelationship          `json:"Server,omitempty"`
	ServerProfile        *ServerProfileRelationship            `json:"ServerProfile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerConfigImportAllOf ServerConfigImportAllOf

// NewServerConfigImportAllOf instantiates a new ServerConfigImportAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigImportAllOf(classId string, objectType string) *ServerConfigImportAllOf {
	this := ServerConfigImportAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServerConfigImportAllOfWithDefaults instantiates a new ServerConfigImportAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigImportAllOfWithDefaults() *ServerConfigImportAllOf {
	this := ServerConfigImportAllOf{}
	var classId string = "server.ConfigImport"
	this.ClassId = classId
	var objectType string = "server.ConfigImport"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerConfigImportAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerConfigImportAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerConfigImportAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServerConfigImportAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerConfigImportAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerConfigImportAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServerConfigImportAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigImportAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServerConfigImportAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServerConfigImportAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyPrefix returns the PolicyPrefix field value if set, zero value otherwise.
func (o *ServerConfigImportAllOf) GetPolicyPrefix() string {
	if o == nil || o.PolicyPrefix == nil {
		var ret string
		return ret
	}
	return *o.PolicyPrefix
}

// GetPolicyPrefixOk returns a tuple with the PolicyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigImportAllOf) GetPolicyPrefixOk() (*string, bool) {
	if o == nil || o.PolicyPrefix == nil {
		return nil, false
	}
	return o.PolicyPrefix, true
}

// HasPolicyPrefix returns a boolean if a field has been set.
func (o *ServerConfigImportAllOf) HasPolicyPrefix() bool {
	if o != nil && o.PolicyPrefix != nil {
		return true
	}

	return false
}

// SetPolicyPrefix gets a reference to the given string and assigns it to the PolicyPrefix field.
func (o *ServerConfigImportAllOf) SetPolicyPrefix(v string) {
	o.PolicyPrefix = &v
}

// GetPolicyTypes returns the PolicyTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerConfigImportAllOf) GetPolicyTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PolicyTypes
}

// GetPolicyTypesOk returns a tuple with the PolicyTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerConfigImportAllOf) GetPolicyTypesOk() (*[]string, bool) {
	if o == nil || o.PolicyTypes == nil {
		return nil, false
	}
	return &o.PolicyTypes, true
}

// HasPolicyTypes returns a boolean if a field has been set.
func (o *ServerConfigImportAllOf) HasPolicyTypes() bool {
	if o != nil && o.PolicyTypes != nil {
		return true
	}

	return false
}

// SetPolicyTypes gets a reference to the given []string and assigns it to the PolicyTypes field.
func (o *ServerConfigImportAllOf) SetPolicyTypes(v []string) {
	o.PolicyTypes = v
}

// GetProfileName returns the ProfileName field value if set, zero value otherwise.
func (o *ServerConfigImportAllOf) GetProfileName() string {
	if o == nil || o.ProfileName == nil {
		var ret string
		return ret
	}
	return *o.ProfileName
}

// GetProfileNameOk returns a tuple with the ProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigImportAllOf) GetProfileNameOk() (*string, bool) {
	if o == nil || o.ProfileName == nil {
		return nil, false
	}
	return o.ProfileName, true
}

// HasProfileName returns a boolean if a field has been set.
func (o *ServerConfigImportAllOf) HasProfileName() bool {
	if o != nil && o.ProfileName != nil {
		return true
	}

	return false
}

// SetProfileName gets a reference to the given string and assigns it to the ProfileName field.
func (o *ServerConfigImportAllOf) SetProfileName(v string) {
	o.ProfileName = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ServerConfigImportAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigImportAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ServerConfigImportAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ServerConfigImportAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ServerConfigImportAllOf) GetServer() ComputeRackUnitRelationship {
	if o == nil || o.Server == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigImportAllOf) GetServerOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ServerConfigImportAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputeRackUnitRelationship and assigns it to the Server field.
func (o *ServerConfigImportAllOf) SetServer(v ComputeRackUnitRelationship) {
	o.Server = &v
}

// GetServerProfile returns the ServerProfile field value if set, zero value otherwise.
func (o *ServerConfigImportAllOf) GetServerProfile() ServerProfileRelationship {
	if o == nil || o.ServerProfile == nil {
		var ret ServerProfileRelationship
		return ret
	}
	return *o.ServerProfile
}

// GetServerProfileOk returns a tuple with the ServerProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigImportAllOf) GetServerProfileOk() (*ServerProfileRelationship, bool) {
	if o == nil || o.ServerProfile == nil {
		return nil, false
	}
	return o.ServerProfile, true
}

// HasServerProfile returns a boolean if a field has been set.
func (o *ServerConfigImportAllOf) HasServerProfile() bool {
	if o != nil && o.ServerProfile != nil {
		return true
	}

	return false
}

// SetServerProfile gets a reference to the given ServerProfileRelationship and assigns it to the ServerProfile field.
func (o *ServerConfigImportAllOf) SetServerProfile(v ServerProfileRelationship) {
	o.ServerProfile = &v
}

func (o ServerConfigImportAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.PolicyPrefix != nil {
		toSerialize["PolicyPrefix"] = o.PolicyPrefix
	}
	if o.PolicyTypes != nil {
		toSerialize["PolicyTypes"] = o.PolicyTypes
	}
	if o.ProfileName != nil {
		toSerialize["ProfileName"] = o.ProfileName
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.ServerProfile != nil {
		toSerialize["ServerProfile"] = o.ServerProfile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerConfigImportAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServerConfigImportAllOf := _ServerConfigImportAllOf{}

	if err = json.Unmarshal(bytes, &varServerConfigImportAllOf); err == nil {
		*o = ServerConfigImportAllOf(varServerConfigImportAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "PolicyPrefix")
		delete(additionalProperties, "PolicyTypes")
		delete(additionalProperties, "ProfileName")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "ServerProfile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerConfigImportAllOf struct {
	value *ServerConfigImportAllOf
	isSet bool
}

func (v NullableServerConfigImportAllOf) Get() *ServerConfigImportAllOf {
	return v.value
}

func (v *NullableServerConfigImportAllOf) Set(val *ServerConfigImportAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigImportAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigImportAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigImportAllOf(val *ServerConfigImportAllOf) *NullableServerConfigImportAllOf {
	return &NullableServerConfigImportAllOf{value: val, isSet: true}
}

func (v NullableServerConfigImportAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigImportAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
