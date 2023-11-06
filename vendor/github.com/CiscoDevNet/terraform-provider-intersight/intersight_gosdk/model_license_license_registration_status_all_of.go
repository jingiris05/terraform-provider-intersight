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

// LicenseLicenseRegistrationStatusAllOf Definition of the list of properties defined in 'license.LicenseRegistrationStatus', excluding properties defined in parent classes.
type LicenseLicenseRegistrationStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Stores information on whether account has gone through the registration wizard. True if has not.
	AccountCreationState *bool `json:"AccountCreationState,omitempty"`
	// Stores information on whether account is new. This data is used for UI theme upgrade, where existing users will be shown a slightly different screen. True if new.
	IsNewAccount *bool `json:"IsNewAccount,omitempty"`
	// Stores information on the current flow of license registration. * `RegistrationNotStarted` - The license registration state to chose between trial and registration. * `RegistrationStarted` - The license registration state during set up flow. * `RegistrationComplete` - The license registration state after completion.
	LicenseRegistrationState *string `json:"LicenseRegistrationState,omitempty"`
	// Stores information on whether trial flow has been completed. True if trial registration finish.
	TrialRegistrationComplete *bool                                  `json:"TrialRegistrationComplete,omitempty"`
	AccountLicenseData        *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _LicenseLicenseRegistrationStatusAllOf LicenseLicenseRegistrationStatusAllOf

// NewLicenseLicenseRegistrationStatusAllOf instantiates a new LicenseLicenseRegistrationStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseLicenseRegistrationStatusAllOf(classId string, objectType string) *LicenseLicenseRegistrationStatusAllOf {
	this := LicenseLicenseRegistrationStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var licenseRegistrationState string = "RegistrationNotStarted"
	this.LicenseRegistrationState = &licenseRegistrationState
	return &this
}

// NewLicenseLicenseRegistrationStatusAllOfWithDefaults instantiates a new LicenseLicenseRegistrationStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseLicenseRegistrationStatusAllOfWithDefaults() *LicenseLicenseRegistrationStatusAllOf {
	this := LicenseLicenseRegistrationStatusAllOf{}
	var classId string = "license.LicenseRegistrationStatus"
	this.ClassId = classId
	var objectType string = "license.LicenseRegistrationStatus"
	this.ObjectType = objectType
	var licenseRegistrationState string = "RegistrationNotStarted"
	this.LicenseRegistrationState = &licenseRegistrationState
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseLicenseRegistrationStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseLicenseRegistrationStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LicenseLicenseRegistrationStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseLicenseRegistrationStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccountCreationState returns the AccountCreationState field value if set, zero value otherwise.
func (o *LicenseLicenseRegistrationStatusAllOf) GetAccountCreationState() bool {
	if o == nil || o.AccountCreationState == nil {
		var ret bool
		return ret
	}
	return *o.AccountCreationState
}

// GetAccountCreationStateOk returns a tuple with the AccountCreationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) GetAccountCreationStateOk() (*bool, bool) {
	if o == nil || o.AccountCreationState == nil {
		return nil, false
	}
	return o.AccountCreationState, true
}

// HasAccountCreationState returns a boolean if a field has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) HasAccountCreationState() bool {
	if o != nil && o.AccountCreationState != nil {
		return true
	}

	return false
}

// SetAccountCreationState gets a reference to the given bool and assigns it to the AccountCreationState field.
func (o *LicenseLicenseRegistrationStatusAllOf) SetAccountCreationState(v bool) {
	o.AccountCreationState = &v
}

// GetIsNewAccount returns the IsNewAccount field value if set, zero value otherwise.
func (o *LicenseLicenseRegistrationStatusAllOf) GetIsNewAccount() bool {
	if o == nil || o.IsNewAccount == nil {
		var ret bool
		return ret
	}
	return *o.IsNewAccount
}

// GetIsNewAccountOk returns a tuple with the IsNewAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) GetIsNewAccountOk() (*bool, bool) {
	if o == nil || o.IsNewAccount == nil {
		return nil, false
	}
	return o.IsNewAccount, true
}

// HasIsNewAccount returns a boolean if a field has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) HasIsNewAccount() bool {
	if o != nil && o.IsNewAccount != nil {
		return true
	}

	return false
}

// SetIsNewAccount gets a reference to the given bool and assigns it to the IsNewAccount field.
func (o *LicenseLicenseRegistrationStatusAllOf) SetIsNewAccount(v bool) {
	o.IsNewAccount = &v
}

// GetLicenseRegistrationState returns the LicenseRegistrationState field value if set, zero value otherwise.
func (o *LicenseLicenseRegistrationStatusAllOf) GetLicenseRegistrationState() string {
	if o == nil || o.LicenseRegistrationState == nil {
		var ret string
		return ret
	}
	return *o.LicenseRegistrationState
}

// GetLicenseRegistrationStateOk returns a tuple with the LicenseRegistrationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) GetLicenseRegistrationStateOk() (*string, bool) {
	if o == nil || o.LicenseRegistrationState == nil {
		return nil, false
	}
	return o.LicenseRegistrationState, true
}

// HasLicenseRegistrationState returns a boolean if a field has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) HasLicenseRegistrationState() bool {
	if o != nil && o.LicenseRegistrationState != nil {
		return true
	}

	return false
}

// SetLicenseRegistrationState gets a reference to the given string and assigns it to the LicenseRegistrationState field.
func (o *LicenseLicenseRegistrationStatusAllOf) SetLicenseRegistrationState(v string) {
	o.LicenseRegistrationState = &v
}

// GetTrialRegistrationComplete returns the TrialRegistrationComplete field value if set, zero value otherwise.
func (o *LicenseLicenseRegistrationStatusAllOf) GetTrialRegistrationComplete() bool {
	if o == nil || o.TrialRegistrationComplete == nil {
		var ret bool
		return ret
	}
	return *o.TrialRegistrationComplete
}

// GetTrialRegistrationCompleteOk returns a tuple with the TrialRegistrationComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) GetTrialRegistrationCompleteOk() (*bool, bool) {
	if o == nil || o.TrialRegistrationComplete == nil {
		return nil, false
	}
	return o.TrialRegistrationComplete, true
}

// HasTrialRegistrationComplete returns a boolean if a field has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) HasTrialRegistrationComplete() bool {
	if o != nil && o.TrialRegistrationComplete != nil {
		return true
	}

	return false
}

// SetTrialRegistrationComplete gets a reference to the given bool and assigns it to the TrialRegistrationComplete field.
func (o *LicenseLicenseRegistrationStatusAllOf) SetTrialRegistrationComplete(v bool) {
	o.TrialRegistrationComplete = &v
}

// GetAccountLicenseData returns the AccountLicenseData field value if set, zero value otherwise.
func (o *LicenseLicenseRegistrationStatusAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship {
	if o == nil || o.AccountLicenseData == nil {
		var ret LicenseAccountLicenseDataRelationship
		return ret
	}
	return *o.AccountLicenseData
}

// GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool) {
	if o == nil || o.AccountLicenseData == nil {
		return nil, false
	}
	return o.AccountLicenseData, true
}

// HasAccountLicenseData returns a boolean if a field has been set.
func (o *LicenseLicenseRegistrationStatusAllOf) HasAccountLicenseData() bool {
	if o != nil && o.AccountLicenseData != nil {
		return true
	}

	return false
}

// SetAccountLicenseData gets a reference to the given LicenseAccountLicenseDataRelationship and assigns it to the AccountLicenseData field.
func (o *LicenseLicenseRegistrationStatusAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship) {
	o.AccountLicenseData = &v
}

func (o LicenseLicenseRegistrationStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccountCreationState != nil {
		toSerialize["AccountCreationState"] = o.AccountCreationState
	}
	if o.IsNewAccount != nil {
		toSerialize["IsNewAccount"] = o.IsNewAccount
	}
	if o.LicenseRegistrationState != nil {
		toSerialize["LicenseRegistrationState"] = o.LicenseRegistrationState
	}
	if o.TrialRegistrationComplete != nil {
		toSerialize["TrialRegistrationComplete"] = o.TrialRegistrationComplete
	}
	if o.AccountLicenseData != nil {
		toSerialize["AccountLicenseData"] = o.AccountLicenseData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseLicenseRegistrationStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLicenseLicenseRegistrationStatusAllOf := _LicenseLicenseRegistrationStatusAllOf{}

	if err = json.Unmarshal(bytes, &varLicenseLicenseRegistrationStatusAllOf); err == nil {
		*o = LicenseLicenseRegistrationStatusAllOf(varLicenseLicenseRegistrationStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccountCreationState")
		delete(additionalProperties, "IsNewAccount")
		delete(additionalProperties, "LicenseRegistrationState")
		delete(additionalProperties, "TrialRegistrationComplete")
		delete(additionalProperties, "AccountLicenseData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseLicenseRegistrationStatusAllOf struct {
	value *LicenseLicenseRegistrationStatusAllOf
	isSet bool
}

func (v NullableLicenseLicenseRegistrationStatusAllOf) Get() *LicenseLicenseRegistrationStatusAllOf {
	return v.value
}

func (v *NullableLicenseLicenseRegistrationStatusAllOf) Set(val *LicenseLicenseRegistrationStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseLicenseRegistrationStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseLicenseRegistrationStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseLicenseRegistrationStatusAllOf(val *LicenseLicenseRegistrationStatusAllOf) *NullableLicenseLicenseRegistrationStatusAllOf {
	return &NullableLicenseLicenseRegistrationStatusAllOf{value: val, isSet: true}
}

func (v NullableLicenseLicenseRegistrationStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseLicenseRegistrationStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
