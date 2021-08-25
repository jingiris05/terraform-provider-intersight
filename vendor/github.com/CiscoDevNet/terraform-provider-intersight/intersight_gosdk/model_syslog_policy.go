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
	"reflect"
	"strings"
)

// SyslogPolicy The syslog policy configure the syslog server to receive CIMC log entries.
type SyslogPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                                `json:"ObjectType"`
	LocalClients  []SyslogLocalClientBase               `json:"LocalClients,omitempty"`
	RemoteClients []SyslogRemoteClientBase              `json:"RemoteClients,omitempty"`
	Organization  *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SyslogPolicy SyslogPolicy

// NewSyslogPolicy instantiates a new SyslogPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogPolicy(classId string, objectType string) *SyslogPolicy {
	this := SyslogPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSyslogPolicyWithDefaults instantiates a new SyslogPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogPolicyWithDefaults() *SyslogPolicy {
	this := SyslogPolicy{}
	var classId string = "syslog.Policy"
	this.ClassId = classId
	var objectType string = "syslog.Policy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SyslogPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SyslogPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SyslogPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SyslogPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SyslogPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SyslogPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLocalClients returns the LocalClients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogPolicy) GetLocalClients() []SyslogLocalClientBase {
	if o == nil {
		var ret []SyslogLocalClientBase
		return ret
	}
	return o.LocalClients
}

// GetLocalClientsOk returns a tuple with the LocalClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogPolicy) GetLocalClientsOk() (*[]SyslogLocalClientBase, bool) {
	if o == nil || o.LocalClients == nil {
		return nil, false
	}
	return &o.LocalClients, true
}

// HasLocalClients returns a boolean if a field has been set.
func (o *SyslogPolicy) HasLocalClients() bool {
	if o != nil && o.LocalClients != nil {
		return true
	}

	return false
}

// SetLocalClients gets a reference to the given []SyslogLocalClientBase and assigns it to the LocalClients field.
func (o *SyslogPolicy) SetLocalClients(v []SyslogLocalClientBase) {
	o.LocalClients = v
}

// GetRemoteClients returns the RemoteClients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogPolicy) GetRemoteClients() []SyslogRemoteClientBase {
	if o == nil {
		var ret []SyslogRemoteClientBase
		return ret
	}
	return o.RemoteClients
}

// GetRemoteClientsOk returns a tuple with the RemoteClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogPolicy) GetRemoteClientsOk() (*[]SyslogRemoteClientBase, bool) {
	if o == nil || o.RemoteClients == nil {
		return nil, false
	}
	return &o.RemoteClients, true
}

// HasRemoteClients returns a boolean if a field has been set.
func (o *SyslogPolicy) HasRemoteClients() bool {
	if o != nil && o.RemoteClients != nil {
		return true
	}

	return false
}

// SetRemoteClients gets a reference to the given []SyslogRemoteClientBase and assigns it to the RemoteClients field.
func (o *SyslogPolicy) SetRemoteClients(v []SyslogRemoteClientBase) {
	o.RemoteClients = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SyslogPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SyslogPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SyslogPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SyslogPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *SyslogPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o SyslogPolicy) MarshalJSON() ([]byte, error) {
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
	if o.LocalClients != nil {
		toSerialize["LocalClients"] = o.LocalClients
	}
	if o.RemoteClients != nil {
		toSerialize["RemoteClients"] = o.RemoteClients
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SyslogPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type SyslogPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                                `json:"ObjectType"`
		LocalClients  []SyslogLocalClientBase               `json:"LocalClients,omitempty"`
		RemoteClients []SyslogRemoteClientBase              `json:"RemoteClients,omitempty"`
		Organization  *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varSyslogPolicyWithoutEmbeddedStruct := SyslogPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSyslogPolicyWithoutEmbeddedStruct)
	if err == nil {
		varSyslogPolicy := _SyslogPolicy{}
		varSyslogPolicy.ClassId = varSyslogPolicyWithoutEmbeddedStruct.ClassId
		varSyslogPolicy.ObjectType = varSyslogPolicyWithoutEmbeddedStruct.ObjectType
		varSyslogPolicy.LocalClients = varSyslogPolicyWithoutEmbeddedStruct.LocalClients
		varSyslogPolicy.RemoteClients = varSyslogPolicyWithoutEmbeddedStruct.RemoteClients
		varSyslogPolicy.Organization = varSyslogPolicyWithoutEmbeddedStruct.Organization
		varSyslogPolicy.Profiles = varSyslogPolicyWithoutEmbeddedStruct.Profiles
		*o = SyslogPolicy(varSyslogPolicy)
	} else {
		return err
	}

	varSyslogPolicy := _SyslogPolicy{}

	err = json.Unmarshal(bytes, &varSyslogPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varSyslogPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LocalClients")
		delete(additionalProperties, "RemoteClients")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

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

type NullableSyslogPolicy struct {
	value *SyslogPolicy
	isSet bool
}

func (v NullableSyslogPolicy) Get() *SyslogPolicy {
	return v.value
}

func (v *NullableSyslogPolicy) Set(val *SyslogPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogPolicy(val *SyslogPolicy) *NullableSyslogPolicy {
	return &NullableSyslogPolicy{value: val, isSet: true}
}

func (v NullableSyslogPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
