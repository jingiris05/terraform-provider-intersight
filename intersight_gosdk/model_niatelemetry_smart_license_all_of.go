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

// NiatelemetrySmartLicenseAllOf Definition of the list of properties defined in 'niatelemetry.SmartLicense', excluding properties defined in parent classes.
type NiatelemetrySmartLicenseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicate the mode smart license is curerntly running.
	ActiveMode *string `json:"ActiveMode,omitempty"`
	// Authorization status of the smart license.
	AuthStatus *string `json:"AuthStatus,omitempty"`
	// License Udi of the smart license.
	LicenseUdi *string `json:"LicenseUdi,omitempty"`
	// Smart licensing account name in CSSM and is retrieved from CSSM after regsitration.
	SmartAccount         *string `json:"SmartAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySmartLicenseAllOf NiatelemetrySmartLicenseAllOf

// NewNiatelemetrySmartLicenseAllOf instantiates a new NiatelemetrySmartLicenseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySmartLicenseAllOf(classId string, objectType string) *NiatelemetrySmartLicenseAllOf {
	this := NiatelemetrySmartLicenseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySmartLicenseAllOfWithDefaults instantiates a new NiatelemetrySmartLicenseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySmartLicenseAllOfWithDefaults() *NiatelemetrySmartLicenseAllOf {
	this := NiatelemetrySmartLicenseAllOf{}
	var classId string = "niatelemetry.SmartLicense"
	this.ClassId = classId
	var objectType string = "niatelemetry.SmartLicense"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySmartLicenseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicenseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySmartLicenseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySmartLicenseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicenseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySmartLicenseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveMode returns the ActiveMode field value if set, zero value otherwise.
func (o *NiatelemetrySmartLicenseAllOf) GetActiveMode() string {
	if o == nil || o.ActiveMode == nil {
		var ret string
		return ret
	}
	return *o.ActiveMode
}

// GetActiveModeOk returns a tuple with the ActiveMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicenseAllOf) GetActiveModeOk() (*string, bool) {
	if o == nil || o.ActiveMode == nil {
		return nil, false
	}
	return o.ActiveMode, true
}

// HasActiveMode returns a boolean if a field has been set.
func (o *NiatelemetrySmartLicenseAllOf) HasActiveMode() bool {
	if o != nil && o.ActiveMode != nil {
		return true
	}

	return false
}

// SetActiveMode gets a reference to the given string and assigns it to the ActiveMode field.
func (o *NiatelemetrySmartLicenseAllOf) SetActiveMode(v string) {
	o.ActiveMode = &v
}

// GetAuthStatus returns the AuthStatus field value if set, zero value otherwise.
func (o *NiatelemetrySmartLicenseAllOf) GetAuthStatus() string {
	if o == nil || o.AuthStatus == nil {
		var ret string
		return ret
	}
	return *o.AuthStatus
}

// GetAuthStatusOk returns a tuple with the AuthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicenseAllOf) GetAuthStatusOk() (*string, bool) {
	if o == nil || o.AuthStatus == nil {
		return nil, false
	}
	return o.AuthStatus, true
}

// HasAuthStatus returns a boolean if a field has been set.
func (o *NiatelemetrySmartLicenseAllOf) HasAuthStatus() bool {
	if o != nil && o.AuthStatus != nil {
		return true
	}

	return false
}

// SetAuthStatus gets a reference to the given string and assigns it to the AuthStatus field.
func (o *NiatelemetrySmartLicenseAllOf) SetAuthStatus(v string) {
	o.AuthStatus = &v
}

// GetLicenseUdi returns the LicenseUdi field value if set, zero value otherwise.
func (o *NiatelemetrySmartLicenseAllOf) GetLicenseUdi() string {
	if o == nil || o.LicenseUdi == nil {
		var ret string
		return ret
	}
	return *o.LicenseUdi
}

// GetLicenseUdiOk returns a tuple with the LicenseUdi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicenseAllOf) GetLicenseUdiOk() (*string, bool) {
	if o == nil || o.LicenseUdi == nil {
		return nil, false
	}
	return o.LicenseUdi, true
}

// HasLicenseUdi returns a boolean if a field has been set.
func (o *NiatelemetrySmartLicenseAllOf) HasLicenseUdi() bool {
	if o != nil && o.LicenseUdi != nil {
		return true
	}

	return false
}

// SetLicenseUdi gets a reference to the given string and assigns it to the LicenseUdi field.
func (o *NiatelemetrySmartLicenseAllOf) SetLicenseUdi(v string) {
	o.LicenseUdi = &v
}

// GetSmartAccount returns the SmartAccount field value if set, zero value otherwise.
func (o *NiatelemetrySmartLicenseAllOf) GetSmartAccount() string {
	if o == nil || o.SmartAccount == nil {
		var ret string
		return ret
	}
	return *o.SmartAccount
}

// GetSmartAccountOk returns a tuple with the SmartAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySmartLicenseAllOf) GetSmartAccountOk() (*string, bool) {
	if o == nil || o.SmartAccount == nil {
		return nil, false
	}
	return o.SmartAccount, true
}

// HasSmartAccount returns a boolean if a field has been set.
func (o *NiatelemetrySmartLicenseAllOf) HasSmartAccount() bool {
	if o != nil && o.SmartAccount != nil {
		return true
	}

	return false
}

// SetSmartAccount gets a reference to the given string and assigns it to the SmartAccount field.
func (o *NiatelemetrySmartLicenseAllOf) SetSmartAccount(v string) {
	o.SmartAccount = &v
}

func (o NiatelemetrySmartLicenseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActiveMode != nil {
		toSerialize["ActiveMode"] = o.ActiveMode
	}
	if o.AuthStatus != nil {
		toSerialize["AuthStatus"] = o.AuthStatus
	}
	if o.LicenseUdi != nil {
		toSerialize["LicenseUdi"] = o.LicenseUdi
	}
	if o.SmartAccount != nil {
		toSerialize["SmartAccount"] = o.SmartAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetrySmartLicenseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetrySmartLicenseAllOf := _NiatelemetrySmartLicenseAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetrySmartLicenseAllOf); err == nil {
		*o = NiatelemetrySmartLicenseAllOf(varNiatelemetrySmartLicenseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveMode")
		delete(additionalProperties, "AuthStatus")
		delete(additionalProperties, "LicenseUdi")
		delete(additionalProperties, "SmartAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetrySmartLicenseAllOf struct {
	value *NiatelemetrySmartLicenseAllOf
	isSet bool
}

func (v NullableNiatelemetrySmartLicenseAllOf) Get() *NiatelemetrySmartLicenseAllOf {
	return v.value
}

func (v *NullableNiatelemetrySmartLicenseAllOf) Set(val *NiatelemetrySmartLicenseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySmartLicenseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySmartLicenseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySmartLicenseAllOf(val *NiatelemetrySmartLicenseAllOf) *NullableNiatelemetrySmartLicenseAllOf {
	return &NullableNiatelemetrySmartLicenseAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetrySmartLicenseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySmartLicenseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
