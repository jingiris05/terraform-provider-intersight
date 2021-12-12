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
	"time"
)

// ApplianceBackupAllOf Definition of the list of properties defined in 'appliance.Backup', excluding properties defined in parent classes.
type ApplianceBackupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Elapsed time in seconds since the backup process has started.
	ElapsedTime *int64 `json:"ElapsedTime,omitempty"`
	// End date and time of the backup process.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool    `json:"IsPasswordSet,omitempty"`
	Messages      []string `json:"Messages,omitempty"`
	// Password to authenticate the fileserver.
	Password *string `json:"Password,omitempty"`
	// Start date and time of the backup process.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Status of the backup managed object. * `Started` - Backup or restore process has started. * `Created` - Backup or restore is in created state. * `Failed` - Backup or restore process has failed. * `Completed` - Backup or restore process has completed. * `Copied` - Backup file has been copied.
	Status               *string                 `json:"Status,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceBackupAllOf ApplianceBackupAllOf

// NewApplianceBackupAllOf instantiates a new ApplianceBackupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceBackupAllOf(classId string, objectType string) *ApplianceBackupAllOf {
	this := ApplianceBackupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceBackupAllOfWithDefaults instantiates a new ApplianceBackupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceBackupAllOfWithDefaults() *ApplianceBackupAllOf {
	this := ApplianceBackupAllOf{}
	var classId string = "appliance.Backup"
	this.ClassId = classId
	var objectType string = "appliance.Backup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceBackupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceBackupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceBackupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceBackupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *ApplianceBackupAllOf) GetElapsedTime() int64 {
	if o == nil || o.ElapsedTime == nil {
		var ret int64
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupAllOf) GetElapsedTimeOk() (*int64, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *ApplianceBackupAllOf) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int64 and assigns it to the ElapsedTime field.
func (o *ApplianceBackupAllOf) SetElapsedTime(v int64) {
	o.ElapsedTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplianceBackupAllOf) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupAllOf) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplianceBackupAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplianceBackupAllOf) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *ApplianceBackupAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *ApplianceBackupAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *ApplianceBackupAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceBackupAllOf) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceBackupAllOf) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return &o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ApplianceBackupAllOf) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *ApplianceBackupAllOf) SetMessages(v []string) {
	o.Messages = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApplianceBackupAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApplianceBackupAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApplianceBackupAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplianceBackupAllOf) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplianceBackupAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplianceBackupAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceBackupAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceBackupAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceBackupAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceBackupAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceBackupAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceBackupAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceBackupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ElapsedTime != nil {
		toSerialize["ElapsedTime"] = o.ElapsedTime
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceBackupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceBackupAllOf := _ApplianceBackupAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceBackupAllOf); err == nil {
		*o = ApplianceBackupAllOf(varApplianceBackupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ElapsedTime")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceBackupAllOf struct {
	value *ApplianceBackupAllOf
	isSet bool
}

func (v NullableApplianceBackupAllOf) Get() *ApplianceBackupAllOf {
	return v.value
}

func (v *NullableApplianceBackupAllOf) Set(val *ApplianceBackupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceBackupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceBackupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceBackupAllOf(val *ApplianceBackupAllOf) *NullableApplianceBackupAllOf {
	return &NullableApplianceBackupAllOf{value: val, isSet: true}
}

func (v NullableApplianceBackupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceBackupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
