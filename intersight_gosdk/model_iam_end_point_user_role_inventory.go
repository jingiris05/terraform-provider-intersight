/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18012
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the IamEndPointUserRoleInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamEndPointUserRoleInventory{}

// IamEndPointUserRoleInventory Mapping of endpoint user to endpoint roles.
type IamEndPointUserRoleInventory struct {
	PolicyAbstractInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Denotes whether password has changed.
	ChangePassword *bool `json:"ChangePassword,omitempty"`
	// Enables the user account on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// The password must have a minimum of 8 and a maximum of 127 characters. For servers with IPMI user role enabled, the maximum length is limited to 20 characters. When strong password is enabled, must satisfy below requirements: A. The password must not contain the User's Name. B. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (! , @, #, $, %, ^, &, *, -, _, +, =).
	Password *string `json:"Password,omitempty" validate:"regexp=^[a-zA-Z0-9!@#$%^&\\\\*+-_=]+$"`
	// An array of relationships to iamEndPointRole resources.
	EndPointRole         []IamEndPointRoleRelationship                      `json:"EndPointRole,omitempty"`
	EndPointUser         NullableIamEndPointUserInventoryRelationship       `json:"EndPointUser,omitempty"`
	EndPointUserPolicy   NullableIamEndPointUserPolicyInventoryRelationship `json:"EndPointUserPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointUserRoleInventory IamEndPointUserRoleInventory

// NewIamEndPointUserRoleInventory instantiates a new IamEndPointUserRoleInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointUserRoleInventory(classId string, objectType string) *IamEndPointUserRoleInventory {
	this := IamEndPointUserRoleInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamEndPointUserRoleInventoryWithDefaults instantiates a new IamEndPointUserRoleInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointUserRoleInventoryWithDefaults() *IamEndPointUserRoleInventory {
	this := IamEndPointUserRoleInventory{}
	var classId string = "iam.EndPointUserRoleInventory"
	this.ClassId = classId
	var objectType string = "iam.EndPointUserRoleInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamEndPointUserRoleInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamEndPointUserRoleInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.EndPointUserRoleInventory" of the ClassId field.
func (o *IamEndPointUserRoleInventory) GetDefaultClassId() interface{} {
	return "iam.EndPointUserRoleInventory"
}

// GetObjectType returns the ObjectType field value
func (o *IamEndPointUserRoleInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamEndPointUserRoleInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.EndPointUserRoleInventory" of the ObjectType field.
func (o *IamEndPointUserRoleInventory) GetDefaultObjectType() interface{} {
	return "iam.EndPointUserRoleInventory"
}

// GetChangePassword returns the ChangePassword field value if set, zero value otherwise.
func (o *IamEndPointUserRoleInventory) GetChangePassword() bool {
	if o == nil || IsNil(o.ChangePassword) {
		var ret bool
		return ret
	}
	return *o.ChangePassword
}

// GetChangePasswordOk returns a tuple with the ChangePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventory) GetChangePasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.ChangePassword) {
		return nil, false
	}
	return o.ChangePassword, true
}

// HasChangePassword returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventory) HasChangePassword() bool {
	if o != nil && !IsNil(o.ChangePassword) {
		return true
	}

	return false
}

// SetChangePassword gets a reference to the given bool and assigns it to the ChangePassword field.
func (o *IamEndPointUserRoleInventory) SetChangePassword(v bool) {
	o.ChangePassword = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IamEndPointUserRoleInventory) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventory) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventory) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IamEndPointUserRoleInventory) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *IamEndPointUserRoleInventory) GetIsPasswordSet() bool {
	if o == nil || IsNil(o.IsPasswordSet) {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventory) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPasswordSet) {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventory) HasIsPasswordSet() bool {
	if o != nil && !IsNil(o.IsPasswordSet) {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *IamEndPointUserRoleInventory) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IamEndPointUserRoleInventory) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventory) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventory) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IamEndPointUserRoleInventory) SetPassword(v string) {
	o.Password = &v
}

// GetEndPointRole returns the EndPointRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserRoleInventory) GetEndPointRole() []IamEndPointRoleRelationship {
	if o == nil {
		var ret []IamEndPointRoleRelationship
		return ret
	}
	return o.EndPointRole
}

// GetEndPointRoleOk returns a tuple with the EndPointRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserRoleInventory) GetEndPointRoleOk() ([]IamEndPointRoleRelationship, bool) {
	if o == nil || IsNil(o.EndPointRole) {
		return nil, false
	}
	return o.EndPointRole, true
}

// HasEndPointRole returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventory) HasEndPointRole() bool {
	if o != nil && !IsNil(o.EndPointRole) {
		return true
	}

	return false
}

// SetEndPointRole gets a reference to the given []IamEndPointRoleRelationship and assigns it to the EndPointRole field.
func (o *IamEndPointUserRoleInventory) SetEndPointRole(v []IamEndPointRoleRelationship) {
	o.EndPointRole = v
}

// GetEndPointUser returns the EndPointUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserRoleInventory) GetEndPointUser() IamEndPointUserInventoryRelationship {
	if o == nil || IsNil(o.EndPointUser.Get()) {
		var ret IamEndPointUserInventoryRelationship
		return ret
	}
	return *o.EndPointUser.Get()
}

// GetEndPointUserOk returns a tuple with the EndPointUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserRoleInventory) GetEndPointUserOk() (*IamEndPointUserInventoryRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndPointUser.Get(), o.EndPointUser.IsSet()
}

// HasEndPointUser returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventory) HasEndPointUser() bool {
	if o != nil && o.EndPointUser.IsSet() {
		return true
	}

	return false
}

// SetEndPointUser gets a reference to the given NullableIamEndPointUserInventoryRelationship and assigns it to the EndPointUser field.
func (o *IamEndPointUserRoleInventory) SetEndPointUser(v IamEndPointUserInventoryRelationship) {
	o.EndPointUser.Set(&v)
}

// SetEndPointUserNil sets the value for EndPointUser to be an explicit nil
func (o *IamEndPointUserRoleInventory) SetEndPointUserNil() {
	o.EndPointUser.Set(nil)
}

// UnsetEndPointUser ensures that no value is present for EndPointUser, not even an explicit nil
func (o *IamEndPointUserRoleInventory) UnsetEndPointUser() {
	o.EndPointUser.Unset()
}

// GetEndPointUserPolicy returns the EndPointUserPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserRoleInventory) GetEndPointUserPolicy() IamEndPointUserPolicyInventoryRelationship {
	if o == nil || IsNil(o.EndPointUserPolicy.Get()) {
		var ret IamEndPointUserPolicyInventoryRelationship
		return ret
	}
	return *o.EndPointUserPolicy.Get()
}

// GetEndPointUserPolicyOk returns a tuple with the EndPointUserPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserRoleInventory) GetEndPointUserPolicyOk() (*IamEndPointUserPolicyInventoryRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndPointUserPolicy.Get(), o.EndPointUserPolicy.IsSet()
}

// HasEndPointUserPolicy returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventory) HasEndPointUserPolicy() bool {
	if o != nil && o.EndPointUserPolicy.IsSet() {
		return true
	}

	return false
}

// SetEndPointUserPolicy gets a reference to the given NullableIamEndPointUserPolicyInventoryRelationship and assigns it to the EndPointUserPolicy field.
func (o *IamEndPointUserRoleInventory) SetEndPointUserPolicy(v IamEndPointUserPolicyInventoryRelationship) {
	o.EndPointUserPolicy.Set(&v)
}

// SetEndPointUserPolicyNil sets the value for EndPointUserPolicy to be an explicit nil
func (o *IamEndPointUserRoleInventory) SetEndPointUserPolicyNil() {
	o.EndPointUserPolicy.Set(nil)
}

// UnsetEndPointUserPolicy ensures that no value is present for EndPointUserPolicy, not even an explicit nil
func (o *IamEndPointUserRoleInventory) UnsetEndPointUserPolicy() {
	o.EndPointUserPolicy.Unset()
}

func (o IamEndPointUserRoleInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamEndPointUserRoleInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractInventory, errPolicyAbstractInventory := json.Marshal(o.PolicyAbstractInventory)
	if errPolicyAbstractInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractInventory
	}
	errPolicyAbstractInventory = json.Unmarshal([]byte(serializedPolicyAbstractInventory), &toSerialize)
	if errPolicyAbstractInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractInventory
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ChangePassword) {
		toSerialize["ChangePassword"] = o.ChangePassword
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.IsPasswordSet) {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if !IsNil(o.Password) {
		toSerialize["Password"] = o.Password
	}
	if o.EndPointRole != nil {
		toSerialize["EndPointRole"] = o.EndPointRole
	}
	if o.EndPointUser.IsSet() {
		toSerialize["EndPointUser"] = o.EndPointUser.Get()
	}
	if o.EndPointUserPolicy.IsSet() {
		toSerialize["EndPointUserPolicy"] = o.EndPointUserPolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamEndPointUserRoleInventory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type IamEndPointUserRoleInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Denotes whether password has changed.
		ChangePassword *bool `json:"ChangePassword,omitempty"`
		// Enables the user account on the endpoint.
		Enabled *bool `json:"Enabled,omitempty"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// The password must have a minimum of 8 and a maximum of 127 characters. For servers with IPMI user role enabled, the maximum length is limited to 20 characters. When strong password is enabled, must satisfy below requirements: A. The password must not contain the User's Name. B. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (! , @, #, $, %, ^, &, *, -, _, +, =).
		Password *string `json:"Password,omitempty" validate:"regexp=^[a-zA-Z0-9!@#$%^&\\\\*+-_=]+$"`
		// An array of relationships to iamEndPointRole resources.
		EndPointRole       []IamEndPointRoleRelationship                      `json:"EndPointRole,omitempty"`
		EndPointUser       NullableIamEndPointUserInventoryRelationship       `json:"EndPointUser,omitempty"`
		EndPointUserPolicy NullableIamEndPointUserPolicyInventoryRelationship `json:"EndPointUserPolicy,omitempty"`
	}

	varIamEndPointUserRoleInventoryWithoutEmbeddedStruct := IamEndPointUserRoleInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamEndPointUserRoleInventoryWithoutEmbeddedStruct)
	if err == nil {
		varIamEndPointUserRoleInventory := _IamEndPointUserRoleInventory{}
		varIamEndPointUserRoleInventory.ClassId = varIamEndPointUserRoleInventoryWithoutEmbeddedStruct.ClassId
		varIamEndPointUserRoleInventory.ObjectType = varIamEndPointUserRoleInventoryWithoutEmbeddedStruct.ObjectType
		varIamEndPointUserRoleInventory.ChangePassword = varIamEndPointUserRoleInventoryWithoutEmbeddedStruct.ChangePassword
		varIamEndPointUserRoleInventory.Enabled = varIamEndPointUserRoleInventoryWithoutEmbeddedStruct.Enabled
		varIamEndPointUserRoleInventory.IsPasswordSet = varIamEndPointUserRoleInventoryWithoutEmbeddedStruct.IsPasswordSet
		varIamEndPointUserRoleInventory.Password = varIamEndPointUserRoleInventoryWithoutEmbeddedStruct.Password
		varIamEndPointUserRoleInventory.EndPointRole = varIamEndPointUserRoleInventoryWithoutEmbeddedStruct.EndPointRole
		varIamEndPointUserRoleInventory.EndPointUser = varIamEndPointUserRoleInventoryWithoutEmbeddedStruct.EndPointUser
		varIamEndPointUserRoleInventory.EndPointUserPolicy = varIamEndPointUserRoleInventoryWithoutEmbeddedStruct.EndPointUserPolicy
		*o = IamEndPointUserRoleInventory(varIamEndPointUserRoleInventory)
	} else {
		return err
	}

	varIamEndPointUserRoleInventory := _IamEndPointUserRoleInventory{}

	err = json.Unmarshal(data, &varIamEndPointUserRoleInventory)
	if err == nil {
		o.PolicyAbstractInventory = varIamEndPointUserRoleInventory.PolicyAbstractInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChangePassword")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "EndPointRole")
		delete(additionalProperties, "EndPointUser")
		delete(additionalProperties, "EndPointUserPolicy")

		// remove fields from embedded structs
		reflectPolicyAbstractInventory := reflect.ValueOf(o.PolicyAbstractInventory)
		for i := 0; i < reflectPolicyAbstractInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractInventory.Type().Field(i)

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

type NullableIamEndPointUserRoleInventory struct {
	value *IamEndPointUserRoleInventory
	isSet bool
}

func (v NullableIamEndPointUserRoleInventory) Get() *IamEndPointUserRoleInventory {
	return v.value
}

func (v *NullableIamEndPointUserRoleInventory) Set(val *IamEndPointUserRoleInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserRoleInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserRoleInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserRoleInventory(val *IamEndPointUserRoleInventory) *NullableIamEndPointUserRoleInventory {
	return &NullableIamEndPointUserRoleInventory{value: val, isSet: true}
}

func (v NullableIamEndPointUserRoleInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserRoleInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
