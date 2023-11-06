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

// RecoveryScheduleConfigPolicy Base Schedule config which contains all the required inputs to do schedule on a local or remote server.
type RecoveryScheduleConfigPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                         `json:"ObjectType"`
	Schedule   NullableRecoveryBackupSchedule `json:"Schedule,omitempty"`
	// An array of relationships to recoveryBackupProfile resources.
	BackupProfiles       []RecoveryBackupProfileRelationship   `json:"BackupProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryScheduleConfigPolicy RecoveryScheduleConfigPolicy

// NewRecoveryScheduleConfigPolicy instantiates a new RecoveryScheduleConfigPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryScheduleConfigPolicy(classId string, objectType string) *RecoveryScheduleConfigPolicy {
	this := RecoveryScheduleConfigPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecoveryScheduleConfigPolicyWithDefaults instantiates a new RecoveryScheduleConfigPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryScheduleConfigPolicyWithDefaults() *RecoveryScheduleConfigPolicy {
	this := RecoveryScheduleConfigPolicy{}
	var classId string = "recovery.ScheduleConfigPolicy"
	this.ClassId = classId
	var objectType string = "recovery.ScheduleConfigPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryScheduleConfigPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryScheduleConfigPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryScheduleConfigPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryScheduleConfigPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryScheduleConfigPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryScheduleConfigPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryScheduleConfigPolicy) GetSchedule() RecoveryBackupSchedule {
	if o == nil || o.Schedule.Get() == nil {
		var ret RecoveryBackupSchedule
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryScheduleConfigPolicy) GetScheduleOk() (*RecoveryBackupSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *RecoveryScheduleConfigPolicy) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableRecoveryBackupSchedule and assigns it to the Schedule field.
func (o *RecoveryScheduleConfigPolicy) SetSchedule(v RecoveryBackupSchedule) {
	o.Schedule.Set(&v)
}

// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *RecoveryScheduleConfigPolicy) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *RecoveryScheduleConfigPolicy) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetBackupProfiles returns the BackupProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryScheduleConfigPolicy) GetBackupProfiles() []RecoveryBackupProfileRelationship {
	if o == nil {
		var ret []RecoveryBackupProfileRelationship
		return ret
	}
	return o.BackupProfiles
}

// GetBackupProfilesOk returns a tuple with the BackupProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryScheduleConfigPolicy) GetBackupProfilesOk() ([]RecoveryBackupProfileRelationship, bool) {
	if o == nil || o.BackupProfiles == nil {
		return nil, false
	}
	return o.BackupProfiles, true
}

// HasBackupProfiles returns a boolean if a field has been set.
func (o *RecoveryScheduleConfigPolicy) HasBackupProfiles() bool {
	if o != nil && o.BackupProfiles != nil {
		return true
	}

	return false
}

// SetBackupProfiles gets a reference to the given []RecoveryBackupProfileRelationship and assigns it to the BackupProfiles field.
func (o *RecoveryScheduleConfigPolicy) SetBackupProfiles(v []RecoveryBackupProfileRelationship) {
	o.BackupProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RecoveryScheduleConfigPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryScheduleConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RecoveryScheduleConfigPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *RecoveryScheduleConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o RecoveryScheduleConfigPolicy) MarshalJSON() ([]byte, error) {
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
	if o.Schedule.IsSet() {
		toSerialize["Schedule"] = o.Schedule.Get()
	}
	if o.BackupProfiles != nil {
		toSerialize["BackupProfiles"] = o.BackupProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryScheduleConfigPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type RecoveryScheduleConfigPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                         `json:"ObjectType"`
		Schedule   NullableRecoveryBackupSchedule `json:"Schedule,omitempty"`
		// An array of relationships to recoveryBackupProfile resources.
		BackupProfiles []RecoveryBackupProfileRelationship   `json:"BackupProfiles,omitempty"`
		Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varRecoveryScheduleConfigPolicyWithoutEmbeddedStruct := RecoveryScheduleConfigPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecoveryScheduleConfigPolicyWithoutEmbeddedStruct)
	if err == nil {
		varRecoveryScheduleConfigPolicy := _RecoveryScheduleConfigPolicy{}
		varRecoveryScheduleConfigPolicy.ClassId = varRecoveryScheduleConfigPolicyWithoutEmbeddedStruct.ClassId
		varRecoveryScheduleConfigPolicy.ObjectType = varRecoveryScheduleConfigPolicyWithoutEmbeddedStruct.ObjectType
		varRecoveryScheduleConfigPolicy.Schedule = varRecoveryScheduleConfigPolicyWithoutEmbeddedStruct.Schedule
		varRecoveryScheduleConfigPolicy.BackupProfiles = varRecoveryScheduleConfigPolicyWithoutEmbeddedStruct.BackupProfiles
		varRecoveryScheduleConfigPolicy.Organization = varRecoveryScheduleConfigPolicyWithoutEmbeddedStruct.Organization
		*o = RecoveryScheduleConfigPolicy(varRecoveryScheduleConfigPolicy)
	} else {
		return err
	}

	varRecoveryScheduleConfigPolicy := _RecoveryScheduleConfigPolicy{}

	err = json.Unmarshal(bytes, &varRecoveryScheduleConfigPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varRecoveryScheduleConfigPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Schedule")
		delete(additionalProperties, "BackupProfiles")
		delete(additionalProperties, "Organization")

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

type NullableRecoveryScheduleConfigPolicy struct {
	value *RecoveryScheduleConfigPolicy
	isSet bool
}

func (v NullableRecoveryScheduleConfigPolicy) Get() *RecoveryScheduleConfigPolicy {
	return v.value
}

func (v *NullableRecoveryScheduleConfigPolicy) Set(val *RecoveryScheduleConfigPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryScheduleConfigPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryScheduleConfigPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryScheduleConfigPolicy(val *RecoveryScheduleConfigPolicy) *NullableRecoveryScheduleConfigPolicy {
	return &NullableRecoveryScheduleConfigPolicy{value: val, isSet: true}
}

func (v NullableRecoveryScheduleConfigPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryScheduleConfigPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
