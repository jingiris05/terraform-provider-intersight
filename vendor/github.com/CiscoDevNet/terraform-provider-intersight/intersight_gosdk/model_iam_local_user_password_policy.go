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

// IamLocalUserPasswordPolicy Configuration for LocalUserPasswordPolicy.
type IamLocalUserPasswordPolicy struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Minimum number of characters different from previous password.
	MinCharDifference *int64 `json:"MinCharDifference,omitempty"`
	// Minimum Days allowed between password change.
	MinDaysBetweenPasswordChange *int64 `json:"MinDaysBetweenPasswordChange,omitempty"`
	// Minimum length of password.
	MinLengthPassword *int64 `json:"MinLengthPassword,omitempty"`
	// Minimum number of required lower case characters.
	MinLowerCase *int64 `json:"MinLowerCase,omitempty"`
	// Minimum number of required numeric characters.
	MinNumeric *int64 `json:"MinNumeric,omitempty"`
	// Minimum number of required special characters.
	MinSpecialChar *int64 `json:"MinSpecialChar,omitempty"`
	// Minimum number of required upper case characters.
	MinUpperCase *int64 `json:"MinUpperCase,omitempty"`
	// Number of previous passwords disallowed.
	NumPreviousPasswordsDisallowed *int64                  `json:"NumPreviousPasswordsDisallowed,omitempty"`
	Account                        *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties           map[string]interface{}
}

type _IamLocalUserPasswordPolicy IamLocalUserPasswordPolicy

// NewIamLocalUserPasswordPolicy instantiates a new IamLocalUserPasswordPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamLocalUserPasswordPolicy(classId string, objectType string) *IamLocalUserPasswordPolicy {
	this := IamLocalUserPasswordPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var minCharDifference int64 = 0
	this.MinCharDifference = &minCharDifference
	var minDaysBetweenPasswordChange int64 = 0
	this.MinDaysBetweenPasswordChange = &minDaysBetweenPasswordChange
	var minLengthPassword int64 = 8
	this.MinLengthPassword = &minLengthPassword
	var minLowerCase int64 = 1
	this.MinLowerCase = &minLowerCase
	var minNumeric int64 = 1
	this.MinNumeric = &minNumeric
	var minSpecialChar int64 = 0
	this.MinSpecialChar = &minSpecialChar
	var minUpperCase int64 = 1
	this.MinUpperCase = &minUpperCase
	var numPreviousPasswordsDisallowed int64 = 0
	this.NumPreviousPasswordsDisallowed = &numPreviousPasswordsDisallowed
	return &this
}

// NewIamLocalUserPasswordPolicyWithDefaults instantiates a new IamLocalUserPasswordPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamLocalUserPasswordPolicyWithDefaults() *IamLocalUserPasswordPolicy {
	this := IamLocalUserPasswordPolicy{}
	var classId string = "iam.LocalUserPasswordPolicy"
	this.ClassId = classId
	var objectType string = "iam.LocalUserPasswordPolicy"
	this.ObjectType = objectType
	var minCharDifference int64 = 0
	this.MinCharDifference = &minCharDifference
	var minDaysBetweenPasswordChange int64 = 0
	this.MinDaysBetweenPasswordChange = &minDaysBetweenPasswordChange
	var minLengthPassword int64 = 8
	this.MinLengthPassword = &minLengthPassword
	var minLowerCase int64 = 1
	this.MinLowerCase = &minLowerCase
	var minNumeric int64 = 1
	this.MinNumeric = &minNumeric
	var minSpecialChar int64 = 0
	this.MinSpecialChar = &minSpecialChar
	var minUpperCase int64 = 1
	this.MinUpperCase = &minUpperCase
	var numPreviousPasswordsDisallowed int64 = 0
	this.NumPreviousPasswordsDisallowed = &numPreviousPasswordsDisallowed
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamLocalUserPasswordPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamLocalUserPasswordPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamLocalUserPasswordPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamLocalUserPasswordPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMinCharDifference returns the MinCharDifference field value if set, zero value otherwise.
func (o *IamLocalUserPasswordPolicy) GetMinCharDifference() int64 {
	if o == nil || o.MinCharDifference == nil {
		var ret int64
		return ret
	}
	return *o.MinCharDifference
}

// GetMinCharDifferenceOk returns a tuple with the MinCharDifference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetMinCharDifferenceOk() (*int64, bool) {
	if o == nil || o.MinCharDifference == nil {
		return nil, false
	}
	return o.MinCharDifference, true
}

// HasMinCharDifference returns a boolean if a field has been set.
func (o *IamLocalUserPasswordPolicy) HasMinCharDifference() bool {
	if o != nil && o.MinCharDifference != nil {
		return true
	}

	return false
}

// SetMinCharDifference gets a reference to the given int64 and assigns it to the MinCharDifference field.
func (o *IamLocalUserPasswordPolicy) SetMinCharDifference(v int64) {
	o.MinCharDifference = &v
}

// GetMinDaysBetweenPasswordChange returns the MinDaysBetweenPasswordChange field value if set, zero value otherwise.
func (o *IamLocalUserPasswordPolicy) GetMinDaysBetweenPasswordChange() int64 {
	if o == nil || o.MinDaysBetweenPasswordChange == nil {
		var ret int64
		return ret
	}
	return *o.MinDaysBetweenPasswordChange
}

// GetMinDaysBetweenPasswordChangeOk returns a tuple with the MinDaysBetweenPasswordChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetMinDaysBetweenPasswordChangeOk() (*int64, bool) {
	if o == nil || o.MinDaysBetweenPasswordChange == nil {
		return nil, false
	}
	return o.MinDaysBetweenPasswordChange, true
}

// HasMinDaysBetweenPasswordChange returns a boolean if a field has been set.
func (o *IamLocalUserPasswordPolicy) HasMinDaysBetweenPasswordChange() bool {
	if o != nil && o.MinDaysBetweenPasswordChange != nil {
		return true
	}

	return false
}

// SetMinDaysBetweenPasswordChange gets a reference to the given int64 and assigns it to the MinDaysBetweenPasswordChange field.
func (o *IamLocalUserPasswordPolicy) SetMinDaysBetweenPasswordChange(v int64) {
	o.MinDaysBetweenPasswordChange = &v
}

// GetMinLengthPassword returns the MinLengthPassword field value if set, zero value otherwise.
func (o *IamLocalUserPasswordPolicy) GetMinLengthPassword() int64 {
	if o == nil || o.MinLengthPassword == nil {
		var ret int64
		return ret
	}
	return *o.MinLengthPassword
}

// GetMinLengthPasswordOk returns a tuple with the MinLengthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetMinLengthPasswordOk() (*int64, bool) {
	if o == nil || o.MinLengthPassword == nil {
		return nil, false
	}
	return o.MinLengthPassword, true
}

// HasMinLengthPassword returns a boolean if a field has been set.
func (o *IamLocalUserPasswordPolicy) HasMinLengthPassword() bool {
	if o != nil && o.MinLengthPassword != nil {
		return true
	}

	return false
}

// SetMinLengthPassword gets a reference to the given int64 and assigns it to the MinLengthPassword field.
func (o *IamLocalUserPasswordPolicy) SetMinLengthPassword(v int64) {
	o.MinLengthPassword = &v
}

// GetMinLowerCase returns the MinLowerCase field value if set, zero value otherwise.
func (o *IamLocalUserPasswordPolicy) GetMinLowerCase() int64 {
	if o == nil || o.MinLowerCase == nil {
		var ret int64
		return ret
	}
	return *o.MinLowerCase
}

// GetMinLowerCaseOk returns a tuple with the MinLowerCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetMinLowerCaseOk() (*int64, bool) {
	if o == nil || o.MinLowerCase == nil {
		return nil, false
	}
	return o.MinLowerCase, true
}

// HasMinLowerCase returns a boolean if a field has been set.
func (o *IamLocalUserPasswordPolicy) HasMinLowerCase() bool {
	if o != nil && o.MinLowerCase != nil {
		return true
	}

	return false
}

// SetMinLowerCase gets a reference to the given int64 and assigns it to the MinLowerCase field.
func (o *IamLocalUserPasswordPolicy) SetMinLowerCase(v int64) {
	o.MinLowerCase = &v
}

// GetMinNumeric returns the MinNumeric field value if set, zero value otherwise.
func (o *IamLocalUserPasswordPolicy) GetMinNumeric() int64 {
	if o == nil || o.MinNumeric == nil {
		var ret int64
		return ret
	}
	return *o.MinNumeric
}

// GetMinNumericOk returns a tuple with the MinNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetMinNumericOk() (*int64, bool) {
	if o == nil || o.MinNumeric == nil {
		return nil, false
	}
	return o.MinNumeric, true
}

// HasMinNumeric returns a boolean if a field has been set.
func (o *IamLocalUserPasswordPolicy) HasMinNumeric() bool {
	if o != nil && o.MinNumeric != nil {
		return true
	}

	return false
}

// SetMinNumeric gets a reference to the given int64 and assigns it to the MinNumeric field.
func (o *IamLocalUserPasswordPolicy) SetMinNumeric(v int64) {
	o.MinNumeric = &v
}

// GetMinSpecialChar returns the MinSpecialChar field value if set, zero value otherwise.
func (o *IamLocalUserPasswordPolicy) GetMinSpecialChar() int64 {
	if o == nil || o.MinSpecialChar == nil {
		var ret int64
		return ret
	}
	return *o.MinSpecialChar
}

// GetMinSpecialCharOk returns a tuple with the MinSpecialChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetMinSpecialCharOk() (*int64, bool) {
	if o == nil || o.MinSpecialChar == nil {
		return nil, false
	}
	return o.MinSpecialChar, true
}

// HasMinSpecialChar returns a boolean if a field has been set.
func (o *IamLocalUserPasswordPolicy) HasMinSpecialChar() bool {
	if o != nil && o.MinSpecialChar != nil {
		return true
	}

	return false
}

// SetMinSpecialChar gets a reference to the given int64 and assigns it to the MinSpecialChar field.
func (o *IamLocalUserPasswordPolicy) SetMinSpecialChar(v int64) {
	o.MinSpecialChar = &v
}

// GetMinUpperCase returns the MinUpperCase field value if set, zero value otherwise.
func (o *IamLocalUserPasswordPolicy) GetMinUpperCase() int64 {
	if o == nil || o.MinUpperCase == nil {
		var ret int64
		return ret
	}
	return *o.MinUpperCase
}

// GetMinUpperCaseOk returns a tuple with the MinUpperCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetMinUpperCaseOk() (*int64, bool) {
	if o == nil || o.MinUpperCase == nil {
		return nil, false
	}
	return o.MinUpperCase, true
}

// HasMinUpperCase returns a boolean if a field has been set.
func (o *IamLocalUserPasswordPolicy) HasMinUpperCase() bool {
	if o != nil && o.MinUpperCase != nil {
		return true
	}

	return false
}

// SetMinUpperCase gets a reference to the given int64 and assigns it to the MinUpperCase field.
func (o *IamLocalUserPasswordPolicy) SetMinUpperCase(v int64) {
	o.MinUpperCase = &v
}

// GetNumPreviousPasswordsDisallowed returns the NumPreviousPasswordsDisallowed field value if set, zero value otherwise.
func (o *IamLocalUserPasswordPolicy) GetNumPreviousPasswordsDisallowed() int64 {
	if o == nil || o.NumPreviousPasswordsDisallowed == nil {
		var ret int64
		return ret
	}
	return *o.NumPreviousPasswordsDisallowed
}

// GetNumPreviousPasswordsDisallowedOk returns a tuple with the NumPreviousPasswordsDisallowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetNumPreviousPasswordsDisallowedOk() (*int64, bool) {
	if o == nil || o.NumPreviousPasswordsDisallowed == nil {
		return nil, false
	}
	return o.NumPreviousPasswordsDisallowed, true
}

// HasNumPreviousPasswordsDisallowed returns a boolean if a field has been set.
func (o *IamLocalUserPasswordPolicy) HasNumPreviousPasswordsDisallowed() bool {
	if o != nil && o.NumPreviousPasswordsDisallowed != nil {
		return true
	}

	return false
}

// SetNumPreviousPasswordsDisallowed gets a reference to the given int64 and assigns it to the NumPreviousPasswordsDisallowed field.
func (o *IamLocalUserPasswordPolicy) SetNumPreviousPasswordsDisallowed(v int64) {
	o.NumPreviousPasswordsDisallowed = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamLocalUserPasswordPolicy) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLocalUserPasswordPolicy) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamLocalUserPasswordPolicy) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamLocalUserPasswordPolicy) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o IamLocalUserPasswordPolicy) MarshalJSON() ([]byte, error) {
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
	if o.MinCharDifference != nil {
		toSerialize["MinCharDifference"] = o.MinCharDifference
	}
	if o.MinDaysBetweenPasswordChange != nil {
		toSerialize["MinDaysBetweenPasswordChange"] = o.MinDaysBetweenPasswordChange
	}
	if o.MinLengthPassword != nil {
		toSerialize["MinLengthPassword"] = o.MinLengthPassword
	}
	if o.MinLowerCase != nil {
		toSerialize["MinLowerCase"] = o.MinLowerCase
	}
	if o.MinNumeric != nil {
		toSerialize["MinNumeric"] = o.MinNumeric
	}
	if o.MinSpecialChar != nil {
		toSerialize["MinSpecialChar"] = o.MinSpecialChar
	}
	if o.MinUpperCase != nil {
		toSerialize["MinUpperCase"] = o.MinUpperCase
	}
	if o.NumPreviousPasswordsDisallowed != nil {
		toSerialize["NumPreviousPasswordsDisallowed"] = o.NumPreviousPasswordsDisallowed
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamLocalUserPasswordPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type IamLocalUserPasswordPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Minimum number of characters different from previous password.
		MinCharDifference *int64 `json:"MinCharDifference,omitempty"`
		// Minimum Days allowed between password change.
		MinDaysBetweenPasswordChange *int64 `json:"MinDaysBetweenPasswordChange,omitempty"`
		// Minimum length of password.
		MinLengthPassword *int64 `json:"MinLengthPassword,omitempty"`
		// Minimum number of required lower case characters.
		MinLowerCase *int64 `json:"MinLowerCase,omitempty"`
		// Minimum number of required numeric characters.
		MinNumeric *int64 `json:"MinNumeric,omitempty"`
		// Minimum number of required special characters.
		MinSpecialChar *int64 `json:"MinSpecialChar,omitempty"`
		// Minimum number of required upper case characters.
		MinUpperCase *int64 `json:"MinUpperCase,omitempty"`
		// Number of previous passwords disallowed.
		NumPreviousPasswordsDisallowed *int64                  `json:"NumPreviousPasswordsDisallowed,omitempty"`
		Account                        *IamAccountRelationship `json:"Account,omitempty"`
	}

	varIamLocalUserPasswordPolicyWithoutEmbeddedStruct := IamLocalUserPasswordPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamLocalUserPasswordPolicyWithoutEmbeddedStruct)
	if err == nil {
		varIamLocalUserPasswordPolicy := _IamLocalUserPasswordPolicy{}
		varIamLocalUserPasswordPolicy.ClassId = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.ClassId
		varIamLocalUserPasswordPolicy.ObjectType = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.ObjectType
		varIamLocalUserPasswordPolicy.MinCharDifference = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.MinCharDifference
		varIamLocalUserPasswordPolicy.MinDaysBetweenPasswordChange = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.MinDaysBetweenPasswordChange
		varIamLocalUserPasswordPolicy.MinLengthPassword = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.MinLengthPassword
		varIamLocalUserPasswordPolicy.MinLowerCase = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.MinLowerCase
		varIamLocalUserPasswordPolicy.MinNumeric = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.MinNumeric
		varIamLocalUserPasswordPolicy.MinSpecialChar = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.MinSpecialChar
		varIamLocalUserPasswordPolicy.MinUpperCase = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.MinUpperCase
		varIamLocalUserPasswordPolicy.NumPreviousPasswordsDisallowed = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.NumPreviousPasswordsDisallowed
		varIamLocalUserPasswordPolicy.Account = varIamLocalUserPasswordPolicyWithoutEmbeddedStruct.Account
		*o = IamLocalUserPasswordPolicy(varIamLocalUserPasswordPolicy)
	} else {
		return err
	}

	varIamLocalUserPasswordPolicy := _IamLocalUserPasswordPolicy{}

	err = json.Unmarshal(bytes, &varIamLocalUserPasswordPolicy)
	if err == nil {
		o.MoBaseMo = varIamLocalUserPasswordPolicy.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MinCharDifference")
		delete(additionalProperties, "MinDaysBetweenPasswordChange")
		delete(additionalProperties, "MinLengthPassword")
		delete(additionalProperties, "MinLowerCase")
		delete(additionalProperties, "MinNumeric")
		delete(additionalProperties, "MinSpecialChar")
		delete(additionalProperties, "MinUpperCase")
		delete(additionalProperties, "NumPreviousPasswordsDisallowed")
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

type NullableIamLocalUserPasswordPolicy struct {
	value *IamLocalUserPasswordPolicy
	isSet bool
}

func (v NullableIamLocalUserPasswordPolicy) Get() *IamLocalUserPasswordPolicy {
	return v.value
}

func (v *NullableIamLocalUserPasswordPolicy) Set(val *IamLocalUserPasswordPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIamLocalUserPasswordPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIamLocalUserPasswordPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamLocalUserPasswordPolicy(val *IamLocalUserPasswordPolicy) *NullableIamLocalUserPasswordPolicy {
	return &NullableIamLocalUserPasswordPolicy{value: val, isSet: true}
}

func (v NullableIamLocalUserPasswordPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamLocalUserPasswordPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
