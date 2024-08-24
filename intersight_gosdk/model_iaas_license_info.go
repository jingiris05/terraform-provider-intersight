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

// checks if the IaasLicenseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IaasLicenseInfo{}

// IaasLicenseInfo Describes about license info currently available in UCSD.
type IaasLicenseInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UCS Director license expiration date.
	LicenseExpirationDate *string               `json:"LicenseExpirationDate,omitempty"`
	LicenseKeysInfo       []IaasLicenseKeysInfo `json:"LicenseKeysInfo,omitempty"`
	// License type of UCSD whether it is EVAL/Permanent/Subscription..
	LicenseType            *string                          `json:"LicenseType,omitempty"`
	LicenseUtilizationInfo []IaasLicenseUtilizationInfo     `json:"LicenseUtilizationInfo,omitempty"`
	Guid                   NullableIaasUcsdInfoRelationship `json:"Guid,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _IaasLicenseInfo IaasLicenseInfo

// NewIaasLicenseInfo instantiates a new IaasLicenseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasLicenseInfo(classId string, objectType string) *IaasLicenseInfo {
	this := IaasLicenseInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasLicenseInfoWithDefaults instantiates a new IaasLicenseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasLicenseInfoWithDefaults() *IaasLicenseInfo {
	this := IaasLicenseInfo{}
	var classId string = "iaas.LicenseInfo"
	this.ClassId = classId
	var objectType string = "iaas.LicenseInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasLicenseInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasLicenseInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasLicenseInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iaas.LicenseInfo" of the ClassId field.
func (o *IaasLicenseInfo) GetDefaultClassId() interface{} {
	return "iaas.LicenseInfo"
}

// GetObjectType returns the ObjectType field value
func (o *IaasLicenseInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasLicenseInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasLicenseInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iaas.LicenseInfo" of the ObjectType field.
func (o *IaasLicenseInfo) GetDefaultObjectType() interface{} {
	return "iaas.LicenseInfo"
}

// GetLicenseExpirationDate returns the LicenseExpirationDate field value if set, zero value otherwise.
func (o *IaasLicenseInfo) GetLicenseExpirationDate() string {
	if o == nil || IsNil(o.LicenseExpirationDate) {
		var ret string
		return ret
	}
	return *o.LicenseExpirationDate
}

// GetLicenseExpirationDateOk returns a tuple with the LicenseExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseInfo) GetLicenseExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseExpirationDate) {
		return nil, false
	}
	return o.LicenseExpirationDate, true
}

// HasLicenseExpirationDate returns a boolean if a field has been set.
func (o *IaasLicenseInfo) HasLicenseExpirationDate() bool {
	if o != nil && !IsNil(o.LicenseExpirationDate) {
		return true
	}

	return false
}

// SetLicenseExpirationDate gets a reference to the given string and assigns it to the LicenseExpirationDate field.
func (o *IaasLicenseInfo) SetLicenseExpirationDate(v string) {
	o.LicenseExpirationDate = &v
}

// GetLicenseKeysInfo returns the LicenseKeysInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasLicenseInfo) GetLicenseKeysInfo() []IaasLicenseKeysInfo {
	if o == nil {
		var ret []IaasLicenseKeysInfo
		return ret
	}
	return o.LicenseKeysInfo
}

// GetLicenseKeysInfoOk returns a tuple with the LicenseKeysInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasLicenseInfo) GetLicenseKeysInfoOk() ([]IaasLicenseKeysInfo, bool) {
	if o == nil || IsNil(o.LicenseKeysInfo) {
		return nil, false
	}
	return o.LicenseKeysInfo, true
}

// HasLicenseKeysInfo returns a boolean if a field has been set.
func (o *IaasLicenseInfo) HasLicenseKeysInfo() bool {
	if o != nil && !IsNil(o.LicenseKeysInfo) {
		return true
	}

	return false
}

// SetLicenseKeysInfo gets a reference to the given []IaasLicenseKeysInfo and assigns it to the LicenseKeysInfo field.
func (o *IaasLicenseInfo) SetLicenseKeysInfo(v []IaasLicenseKeysInfo) {
	o.LicenseKeysInfo = v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *IaasLicenseInfo) GetLicenseType() string {
	if o == nil || IsNil(o.LicenseType) {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseInfo) GetLicenseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseType) {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *IaasLicenseInfo) HasLicenseType() bool {
	if o != nil && !IsNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *IaasLicenseInfo) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetLicenseUtilizationInfo returns the LicenseUtilizationInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasLicenseInfo) GetLicenseUtilizationInfo() []IaasLicenseUtilizationInfo {
	if o == nil {
		var ret []IaasLicenseUtilizationInfo
		return ret
	}
	return o.LicenseUtilizationInfo
}

// GetLicenseUtilizationInfoOk returns a tuple with the LicenseUtilizationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasLicenseInfo) GetLicenseUtilizationInfoOk() ([]IaasLicenseUtilizationInfo, bool) {
	if o == nil || IsNil(o.LicenseUtilizationInfo) {
		return nil, false
	}
	return o.LicenseUtilizationInfo, true
}

// HasLicenseUtilizationInfo returns a boolean if a field has been set.
func (o *IaasLicenseInfo) HasLicenseUtilizationInfo() bool {
	if o != nil && !IsNil(o.LicenseUtilizationInfo) {
		return true
	}

	return false
}

// SetLicenseUtilizationInfo gets a reference to the given []IaasLicenseUtilizationInfo and assigns it to the LicenseUtilizationInfo field.
func (o *IaasLicenseInfo) SetLicenseUtilizationInfo(v []IaasLicenseUtilizationInfo) {
	o.LicenseUtilizationInfo = v
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasLicenseInfo) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasLicenseInfo) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasLicenseInfo) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableIaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasLicenseInfo) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid.Set(&v)
}

// SetGuidNil sets the value for Guid to be an explicit nil
func (o *IaasLicenseInfo) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *IaasLicenseInfo) UnsetGuid() {
	o.Guid.Unset()
}

func (o IaasLicenseInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IaasLicenseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.LicenseExpirationDate) {
		toSerialize["LicenseExpirationDate"] = o.LicenseExpirationDate
	}
	if o.LicenseKeysInfo != nil {
		toSerialize["LicenseKeysInfo"] = o.LicenseKeysInfo
	}
	if !IsNil(o.LicenseType) {
		toSerialize["LicenseType"] = o.LicenseType
	}
	if o.LicenseUtilizationInfo != nil {
		toSerialize["LicenseUtilizationInfo"] = o.LicenseUtilizationInfo
	}
	if o.Guid.IsSet() {
		toSerialize["Guid"] = o.Guid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IaasLicenseInfo) UnmarshalJSON(data []byte) (err error) {
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
	type IaasLicenseInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// UCS Director license expiration date.
		LicenseExpirationDate *string               `json:"LicenseExpirationDate,omitempty"`
		LicenseKeysInfo       []IaasLicenseKeysInfo `json:"LicenseKeysInfo,omitempty"`
		// License type of UCSD whether it is EVAL/Permanent/Subscription..
		LicenseType            *string                          `json:"LicenseType,omitempty"`
		LicenseUtilizationInfo []IaasLicenseUtilizationInfo     `json:"LicenseUtilizationInfo,omitempty"`
		Guid                   NullableIaasUcsdInfoRelationship `json:"Guid,omitempty"`
	}

	varIaasLicenseInfoWithoutEmbeddedStruct := IaasLicenseInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIaasLicenseInfoWithoutEmbeddedStruct)
	if err == nil {
		varIaasLicenseInfo := _IaasLicenseInfo{}
		varIaasLicenseInfo.ClassId = varIaasLicenseInfoWithoutEmbeddedStruct.ClassId
		varIaasLicenseInfo.ObjectType = varIaasLicenseInfoWithoutEmbeddedStruct.ObjectType
		varIaasLicenseInfo.LicenseExpirationDate = varIaasLicenseInfoWithoutEmbeddedStruct.LicenseExpirationDate
		varIaasLicenseInfo.LicenseKeysInfo = varIaasLicenseInfoWithoutEmbeddedStruct.LicenseKeysInfo
		varIaasLicenseInfo.LicenseType = varIaasLicenseInfoWithoutEmbeddedStruct.LicenseType
		varIaasLicenseInfo.LicenseUtilizationInfo = varIaasLicenseInfoWithoutEmbeddedStruct.LicenseUtilizationInfo
		varIaasLicenseInfo.Guid = varIaasLicenseInfoWithoutEmbeddedStruct.Guid
		*o = IaasLicenseInfo(varIaasLicenseInfo)
	} else {
		return err
	}

	varIaasLicenseInfo := _IaasLicenseInfo{}

	err = json.Unmarshal(data, &varIaasLicenseInfo)
	if err == nil {
		o.MoBaseMo = varIaasLicenseInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LicenseExpirationDate")
		delete(additionalProperties, "LicenseKeysInfo")
		delete(additionalProperties, "LicenseType")
		delete(additionalProperties, "LicenseUtilizationInfo")
		delete(additionalProperties, "Guid")

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

type NullableIaasLicenseInfo struct {
	value *IaasLicenseInfo
	isSet bool
}

func (v NullableIaasLicenseInfo) Get() *IaasLicenseInfo {
	return v.value
}

func (v *NullableIaasLicenseInfo) Set(val *IaasLicenseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasLicenseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasLicenseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasLicenseInfo(val *IaasLicenseInfo) *NullableIaasLicenseInfo {
	return &NullableIaasLicenseInfo{value: val, isSet: true}
}

func (v NullableIaasLicenseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasLicenseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
