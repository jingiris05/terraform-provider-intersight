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
)

// ApplianceBackupMonitorAllOf Definition of the list of properties defined in 'appliance.BackupMonitor', excluding properties defined in parent classes.
type ApplianceBackupMonitorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the most recent Intersight Appliance backup. * `BackupFound` - Backup is successful and complete. * `BackupFailed` - The current Backup failed. * `BackupOutdated` - Backup is old and outdated.
	LastBackupStatus     *string                 `json:"LastBackupStatus,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceBackupMonitorAllOf ApplianceBackupMonitorAllOf

// NewApplianceBackupMonitorAllOf instantiates a new ApplianceBackupMonitorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceBackupMonitorAllOf(classId string, objectType string) *ApplianceBackupMonitorAllOf {
	this := ApplianceBackupMonitorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceBackupMonitorAllOfWithDefaults instantiates a new ApplianceBackupMonitorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceBackupMonitorAllOfWithDefaults() *ApplianceBackupMonitorAllOf {
	this := ApplianceBackupMonitorAllOf{}
	var classId string = "appliance.BackupMonitor"
	this.ClassId = classId
	var objectType string = "appliance.BackupMonitor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceBackupMonitorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupMonitorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceBackupMonitorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceBackupMonitorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupMonitorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceBackupMonitorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLastBackupStatus returns the LastBackupStatus field value if set, zero value otherwise.
func (o *ApplianceBackupMonitorAllOf) GetLastBackupStatus() string {
	if o == nil || o.LastBackupStatus == nil {
		var ret string
		return ret
	}
	return *o.LastBackupStatus
}

// GetLastBackupStatusOk returns a tuple with the LastBackupStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupMonitorAllOf) GetLastBackupStatusOk() (*string, bool) {
	if o == nil || o.LastBackupStatus == nil {
		return nil, false
	}
	return o.LastBackupStatus, true
}

// HasLastBackupStatus returns a boolean if a field has been set.
func (o *ApplianceBackupMonitorAllOf) HasLastBackupStatus() bool {
	if o != nil && o.LastBackupStatus != nil {
		return true
	}

	return false
}

// SetLastBackupStatus gets a reference to the given string and assigns it to the LastBackupStatus field.
func (o *ApplianceBackupMonitorAllOf) SetLastBackupStatus(v string) {
	o.LastBackupStatus = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceBackupMonitorAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupMonitorAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceBackupMonitorAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceBackupMonitorAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceBackupMonitorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LastBackupStatus != nil {
		toSerialize["LastBackupStatus"] = o.LastBackupStatus
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceBackupMonitorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceBackupMonitorAllOf := _ApplianceBackupMonitorAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceBackupMonitorAllOf); err == nil {
		*o = ApplianceBackupMonitorAllOf(varApplianceBackupMonitorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LastBackupStatus")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceBackupMonitorAllOf struct {
	value *ApplianceBackupMonitorAllOf
	isSet bool
}

func (v NullableApplianceBackupMonitorAllOf) Get() *ApplianceBackupMonitorAllOf {
	return v.value
}

func (v *NullableApplianceBackupMonitorAllOf) Set(val *ApplianceBackupMonitorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceBackupMonitorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceBackupMonitorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceBackupMonitorAllOf(val *ApplianceBackupMonitorAllOf) *NullableApplianceBackupMonitorAllOf {
	return &NullableApplianceBackupMonitorAllOf{value: val, isSet: true}
}

func (v NullableApplianceBackupMonitorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceBackupMonitorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
