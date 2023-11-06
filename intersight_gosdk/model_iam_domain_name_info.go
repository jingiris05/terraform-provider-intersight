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
	"time"
)

// IamDomainNameInfo The organisation's domain name such as cisco.com that has been used to log in to Intersight.
type IamDomainNameInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Regenerate TXT record and validate TXT record. * `generate` - Generate TXT record for domain name ownership validation. * `verify` - Verify TXT record for domain name ownership validation.
	Action *string `json:"Action,omitempty"`
	// Email domain name. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
	DomainName     *string                   `json:"DomainName,omitempty"`
	FailureDetails NullableIamFailureDetails `json:"FailureDetails,omitempty"`
	// Expiration time of TXT Record.
	RecordExpiryTime *time.Time `json:"RecordExpiryTime,omitempty"`
	// Status of Domain Ownership Verification. * `Pending` - Domain verification is pending. * `Failed` - Domain verification failed. Re-generate token and verify. * `Verified` - Domain verification succeeded. * `Expired` - TXT Record for Domain verification expired.
	Status *string `json:"Status,omitempty"`
	// Resource record used to verify Domain Ownership. Users need to verify the domain by adding the TXT Record in their DNS Host.
	TxtRecord            *string                 `json:"TxtRecord,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamDomainNameInfo IamDomainNameInfo

// NewIamDomainNameInfo instantiates a new IamDomainNameInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamDomainNameInfo(classId string, objectType string) *IamDomainNameInfo {
	this := IamDomainNameInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "generate"
	this.Action = &action
	return &this
}

// NewIamDomainNameInfoWithDefaults instantiates a new IamDomainNameInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamDomainNameInfoWithDefaults() *IamDomainNameInfo {
	this := IamDomainNameInfo{}
	var classId string = "iam.DomainNameInfo"
	this.ClassId = classId
	var objectType string = "iam.DomainNameInfo"
	this.ObjectType = objectType
	var action string = "generate"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamDomainNameInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamDomainNameInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamDomainNameInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamDomainNameInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *IamDomainNameInfo) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfo) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *IamDomainNameInfo) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *IamDomainNameInfo) SetAction(v string) {
	o.Action = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *IamDomainNameInfo) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfo) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *IamDomainNameInfo) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *IamDomainNameInfo) SetDomainName(v string) {
	o.DomainName = &v
}

// GetFailureDetails returns the FailureDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamDomainNameInfo) GetFailureDetails() IamFailureDetails {
	if o == nil || o.FailureDetails.Get() == nil {
		var ret IamFailureDetails
		return ret
	}
	return *o.FailureDetails.Get()
}

// GetFailureDetailsOk returns a tuple with the FailureDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamDomainNameInfo) GetFailureDetailsOk() (*IamFailureDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureDetails.Get(), o.FailureDetails.IsSet()
}

// HasFailureDetails returns a boolean if a field has been set.
func (o *IamDomainNameInfo) HasFailureDetails() bool {
	if o != nil && o.FailureDetails.IsSet() {
		return true
	}

	return false
}

// SetFailureDetails gets a reference to the given NullableIamFailureDetails and assigns it to the FailureDetails field.
func (o *IamDomainNameInfo) SetFailureDetails(v IamFailureDetails) {
	o.FailureDetails.Set(&v)
}

// SetFailureDetailsNil sets the value for FailureDetails to be an explicit nil
func (o *IamDomainNameInfo) SetFailureDetailsNil() {
	o.FailureDetails.Set(nil)
}

// UnsetFailureDetails ensures that no value is present for FailureDetails, not even an explicit nil
func (o *IamDomainNameInfo) UnsetFailureDetails() {
	o.FailureDetails.Unset()
}

// GetRecordExpiryTime returns the RecordExpiryTime field value if set, zero value otherwise.
func (o *IamDomainNameInfo) GetRecordExpiryTime() time.Time {
	if o == nil || o.RecordExpiryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.RecordExpiryTime
}

// GetRecordExpiryTimeOk returns a tuple with the RecordExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfo) GetRecordExpiryTimeOk() (*time.Time, bool) {
	if o == nil || o.RecordExpiryTime == nil {
		return nil, false
	}
	return o.RecordExpiryTime, true
}

// HasRecordExpiryTime returns a boolean if a field has been set.
func (o *IamDomainNameInfo) HasRecordExpiryTime() bool {
	if o != nil && o.RecordExpiryTime != nil {
		return true
	}

	return false
}

// SetRecordExpiryTime gets a reference to the given time.Time and assigns it to the RecordExpiryTime field.
func (o *IamDomainNameInfo) SetRecordExpiryTime(v time.Time) {
	o.RecordExpiryTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IamDomainNameInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IamDomainNameInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IamDomainNameInfo) SetStatus(v string) {
	o.Status = &v
}

// GetTxtRecord returns the TxtRecord field value if set, zero value otherwise.
func (o *IamDomainNameInfo) GetTxtRecord() string {
	if o == nil || o.TxtRecord == nil {
		var ret string
		return ret
	}
	return *o.TxtRecord
}

// GetTxtRecordOk returns a tuple with the TxtRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfo) GetTxtRecordOk() (*string, bool) {
	if o == nil || o.TxtRecord == nil {
		return nil, false
	}
	return o.TxtRecord, true
}

// HasTxtRecord returns a boolean if a field has been set.
func (o *IamDomainNameInfo) HasTxtRecord() bool {
	if o != nil && o.TxtRecord != nil {
		return true
	}

	return false
}

// SetTxtRecord gets a reference to the given string and assigns it to the TxtRecord field.
func (o *IamDomainNameInfo) SetTxtRecord(v string) {
	o.TxtRecord = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamDomainNameInfo) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfo) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamDomainNameInfo) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamDomainNameInfo) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o IamDomainNameInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.DomainName != nil {
		toSerialize["DomainName"] = o.DomainName
	}
	if o.FailureDetails.IsSet() {
		toSerialize["FailureDetails"] = o.FailureDetails.Get()
	}
	if o.RecordExpiryTime != nil {
		toSerialize["RecordExpiryTime"] = o.RecordExpiryTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TxtRecord != nil {
		toSerialize["TxtRecord"] = o.TxtRecord
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamDomainNameInfo) UnmarshalJSON(bytes []byte) (err error) {
	type IamDomainNameInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Regenerate TXT record and validate TXT record. * `generate` - Generate TXT record for domain name ownership validation. * `verify` - Verify TXT record for domain name ownership validation.
		Action *string `json:"Action,omitempty"`
		// Email domain name. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
		DomainName     *string                   `json:"DomainName,omitempty"`
		FailureDetails NullableIamFailureDetails `json:"FailureDetails,omitempty"`
		// Expiration time of TXT Record.
		RecordExpiryTime *time.Time `json:"RecordExpiryTime,omitempty"`
		// Status of Domain Ownership Verification. * `Pending` - Domain verification is pending. * `Failed` - Domain verification failed. Re-generate token and verify. * `Verified` - Domain verification succeeded. * `Expired` - TXT Record for Domain verification expired.
		Status *string `json:"Status,omitempty"`
		// Resource record used to verify Domain Ownership. Users need to verify the domain by adding the TXT Record in their DNS Host.
		TxtRecord *string                 `json:"TxtRecord,omitempty"`
		Account   *IamAccountRelationship `json:"Account,omitempty"`
	}

	varIamDomainNameInfoWithoutEmbeddedStruct := IamDomainNameInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamDomainNameInfoWithoutEmbeddedStruct)
	if err == nil {
		varIamDomainNameInfo := _IamDomainNameInfo{}
		varIamDomainNameInfo.ClassId = varIamDomainNameInfoWithoutEmbeddedStruct.ClassId
		varIamDomainNameInfo.ObjectType = varIamDomainNameInfoWithoutEmbeddedStruct.ObjectType
		varIamDomainNameInfo.Action = varIamDomainNameInfoWithoutEmbeddedStruct.Action
		varIamDomainNameInfo.DomainName = varIamDomainNameInfoWithoutEmbeddedStruct.DomainName
		varIamDomainNameInfo.FailureDetails = varIamDomainNameInfoWithoutEmbeddedStruct.FailureDetails
		varIamDomainNameInfo.RecordExpiryTime = varIamDomainNameInfoWithoutEmbeddedStruct.RecordExpiryTime
		varIamDomainNameInfo.Status = varIamDomainNameInfoWithoutEmbeddedStruct.Status
		varIamDomainNameInfo.TxtRecord = varIamDomainNameInfoWithoutEmbeddedStruct.TxtRecord
		varIamDomainNameInfo.Account = varIamDomainNameInfoWithoutEmbeddedStruct.Account
		*o = IamDomainNameInfo(varIamDomainNameInfo)
	} else {
		return err
	}

	varIamDomainNameInfo := _IamDomainNameInfo{}

	err = json.Unmarshal(bytes, &varIamDomainNameInfo)
	if err == nil {
		o.MoBaseMo = varIamDomainNameInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "DomainName")
		delete(additionalProperties, "FailureDetails")
		delete(additionalProperties, "RecordExpiryTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TxtRecord")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIamDomainNameInfo struct {
	value *IamDomainNameInfo
	isSet bool
}

func (v NullableIamDomainNameInfo) Get() *IamDomainNameInfo {
	return v.value
}

func (v *NullableIamDomainNameInfo) Set(val *IamDomainNameInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIamDomainNameInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIamDomainNameInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamDomainNameInfo(val *IamDomainNameInfo) *NullableIamDomainNameInfo {
	return &NullableIamDomainNameInfo{value: val, isSet: true}
}

func (v NullableIamDomainNameInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamDomainNameInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
