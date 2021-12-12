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

// IamEndPointUserPolicy Enables creation of local users on endpoints.
type IamEndPointUserPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType         string                                `json:"ObjectType"`
	PasswordProperties NullableIamEndPointPasswordProperties `json:"PasswordProperties,omitempty"`
	// An array of relationships to iamEndPointUserRole resources.
	EndPointUserRoles []IamEndPointUserRoleRelationship     `json:"EndPointUserRoles,omitempty"`
	Organization      *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointUserPolicy IamEndPointUserPolicy

// NewIamEndPointUserPolicy instantiates a new IamEndPointUserPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointUserPolicy(classId string, objectType string) *IamEndPointUserPolicy {
	this := IamEndPointUserPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamEndPointUserPolicyWithDefaults instantiates a new IamEndPointUserPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointUserPolicyWithDefaults() *IamEndPointUserPolicy {
	this := IamEndPointUserPolicy{}
	var classId string = "iam.EndPointUserPolicy"
	this.ClassId = classId
	var objectType string = "iam.EndPointUserPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamEndPointUserPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamEndPointUserPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamEndPointUserPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamEndPointUserPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPasswordProperties returns the PasswordProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserPolicy) GetPasswordProperties() IamEndPointPasswordProperties {
	if o == nil || o.PasswordProperties.Get() == nil {
		var ret IamEndPointPasswordProperties
		return ret
	}
	return *o.PasswordProperties.Get()
}

// GetPasswordPropertiesOk returns a tuple with the PasswordProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserPolicy) GetPasswordPropertiesOk() (*IamEndPointPasswordProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordProperties.Get(), o.PasswordProperties.IsSet()
}

// HasPasswordProperties returns a boolean if a field has been set.
func (o *IamEndPointUserPolicy) HasPasswordProperties() bool {
	if o != nil && o.PasswordProperties.IsSet() {
		return true
	}

	return false
}

// SetPasswordProperties gets a reference to the given NullableIamEndPointPasswordProperties and assigns it to the PasswordProperties field.
func (o *IamEndPointUserPolicy) SetPasswordProperties(v IamEndPointPasswordProperties) {
	o.PasswordProperties.Set(&v)
}

// SetPasswordPropertiesNil sets the value for PasswordProperties to be an explicit nil
func (o *IamEndPointUserPolicy) SetPasswordPropertiesNil() {
	o.PasswordProperties.Set(nil)
}

// UnsetPasswordProperties ensures that no value is present for PasswordProperties, not even an explicit nil
func (o *IamEndPointUserPolicy) UnsetPasswordProperties() {
	o.PasswordProperties.Unset()
}

// GetEndPointUserRoles returns the EndPointUserRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserPolicy) GetEndPointUserRoles() []IamEndPointUserRoleRelationship {
	if o == nil {
		var ret []IamEndPointUserRoleRelationship
		return ret
	}
	return o.EndPointUserRoles
}

// GetEndPointUserRolesOk returns a tuple with the EndPointUserRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserPolicy) GetEndPointUserRolesOk() (*[]IamEndPointUserRoleRelationship, bool) {
	if o == nil || o.EndPointUserRoles == nil {
		return nil, false
	}
	return &o.EndPointUserRoles, true
}

// HasEndPointUserRoles returns a boolean if a field has been set.
func (o *IamEndPointUserPolicy) HasEndPointUserRoles() bool {
	if o != nil && o.EndPointUserRoles != nil {
		return true
	}

	return false
}

// SetEndPointUserRoles gets a reference to the given []IamEndPointUserRoleRelationship and assigns it to the EndPointUserRoles field.
func (o *IamEndPointUserPolicy) SetEndPointUserRoles(v []IamEndPointUserRoleRelationship) {
	o.EndPointUserRoles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *IamEndPointUserPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *IamEndPointUserPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IamEndPointUserPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *IamEndPointUserPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *IamEndPointUserPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o IamEndPointUserPolicy) MarshalJSON() ([]byte, error) {
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
	if o.PasswordProperties.IsSet() {
		toSerialize["PasswordProperties"] = o.PasswordProperties.Get()
	}
	if o.EndPointUserRoles != nil {
		toSerialize["EndPointUserRoles"] = o.EndPointUserRoles
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

func (o *IamEndPointUserPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type IamEndPointUserPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType         string                                `json:"ObjectType"`
		PasswordProperties NullableIamEndPointPasswordProperties `json:"PasswordProperties,omitempty"`
		// An array of relationships to iamEndPointUserRole resources.
		EndPointUserRoles []IamEndPointUserRoleRelationship     `json:"EndPointUserRoles,omitempty"`
		Organization      *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varIamEndPointUserPolicyWithoutEmbeddedStruct := IamEndPointUserPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamEndPointUserPolicyWithoutEmbeddedStruct)
	if err == nil {
		varIamEndPointUserPolicy := _IamEndPointUserPolicy{}
		varIamEndPointUserPolicy.ClassId = varIamEndPointUserPolicyWithoutEmbeddedStruct.ClassId
		varIamEndPointUserPolicy.ObjectType = varIamEndPointUserPolicyWithoutEmbeddedStruct.ObjectType
		varIamEndPointUserPolicy.PasswordProperties = varIamEndPointUserPolicyWithoutEmbeddedStruct.PasswordProperties
		varIamEndPointUserPolicy.EndPointUserRoles = varIamEndPointUserPolicyWithoutEmbeddedStruct.EndPointUserRoles
		varIamEndPointUserPolicy.Organization = varIamEndPointUserPolicyWithoutEmbeddedStruct.Organization
		varIamEndPointUserPolicy.Profiles = varIamEndPointUserPolicyWithoutEmbeddedStruct.Profiles
		*o = IamEndPointUserPolicy(varIamEndPointUserPolicy)
	} else {
		return err
	}

	varIamEndPointUserPolicy := _IamEndPointUserPolicy{}

	err = json.Unmarshal(bytes, &varIamEndPointUserPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varIamEndPointUserPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PasswordProperties")
		delete(additionalProperties, "EndPointUserRoles")
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

type NullableIamEndPointUserPolicy struct {
	value *IamEndPointUserPolicy
	isSet bool
}

func (v NullableIamEndPointUserPolicy) Get() *IamEndPointUserPolicy {
	return v.value
}

func (v *NullableIamEndPointUserPolicy) Set(val *IamEndPointUserPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserPolicy(val *IamEndPointUserPolicy) *NullableIamEndPointUserPolicy {
	return &NullableIamEndPointUserPolicy{value: val, isSet: true}
}

func (v NullableIamEndPointUserPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
