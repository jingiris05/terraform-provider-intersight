/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-11-20T05:29:54Z.
 *
 * API version: 1.0.9-2713
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IaasLicenseUtilizationInfo License list with the Utilization info for UCSD.
type IaasLicenseUtilizationInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of licenses actually used for this feature.
	ActualUsed *int64 `json:"ActualUsed,omitempty"`
	// License label of UCSD license.
	Label *string `json:"Label,omitempty"`
	// License limit for this license feature.
	LicensedLimit *string `json:"LicensedLimit,omitempty"`
	// SKU for the UCSD license.
	Sku                  *string `json:"Sku,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasLicenseUtilizationInfo IaasLicenseUtilizationInfo

// NewIaasLicenseUtilizationInfo instantiates a new IaasLicenseUtilizationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasLicenseUtilizationInfo(classId string, objectType string) *IaasLicenseUtilizationInfo {
	this := IaasLicenseUtilizationInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasLicenseUtilizationInfoWithDefaults instantiates a new IaasLicenseUtilizationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasLicenseUtilizationInfoWithDefaults() *IaasLicenseUtilizationInfo {
	this := IaasLicenseUtilizationInfo{}
	var classId string = "iaas.LicenseUtilizationInfo"
	this.ClassId = classId
	var objectType string = "iaas.LicenseUtilizationInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasLicenseUtilizationInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasLicenseUtilizationInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasLicenseUtilizationInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasLicenseUtilizationInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasLicenseUtilizationInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasLicenseUtilizationInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActualUsed returns the ActualUsed field value if set, zero value otherwise.
func (o *IaasLicenseUtilizationInfo) GetActualUsed() int64 {
	if o == nil || o.ActualUsed == nil {
		var ret int64
		return ret
	}
	return *o.ActualUsed
}

// GetActualUsedOk returns a tuple with the ActualUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseUtilizationInfo) GetActualUsedOk() (*int64, bool) {
	if o == nil || o.ActualUsed == nil {
		return nil, false
	}
	return o.ActualUsed, true
}

// HasActualUsed returns a boolean if a field has been set.
func (o *IaasLicenseUtilizationInfo) HasActualUsed() bool {
	if o != nil && o.ActualUsed != nil {
		return true
	}

	return false
}

// SetActualUsed gets a reference to the given int64 and assigns it to the ActualUsed field.
func (o *IaasLicenseUtilizationInfo) SetActualUsed(v int64) {
	o.ActualUsed = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IaasLicenseUtilizationInfo) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseUtilizationInfo) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IaasLicenseUtilizationInfo) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *IaasLicenseUtilizationInfo) SetLabel(v string) {
	o.Label = &v
}

// GetLicensedLimit returns the LicensedLimit field value if set, zero value otherwise.
func (o *IaasLicenseUtilizationInfo) GetLicensedLimit() string {
	if o == nil || o.LicensedLimit == nil {
		var ret string
		return ret
	}
	return *o.LicensedLimit
}

// GetLicensedLimitOk returns a tuple with the LicensedLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseUtilizationInfo) GetLicensedLimitOk() (*string, bool) {
	if o == nil || o.LicensedLimit == nil {
		return nil, false
	}
	return o.LicensedLimit, true
}

// HasLicensedLimit returns a boolean if a field has been set.
func (o *IaasLicenseUtilizationInfo) HasLicensedLimit() bool {
	if o != nil && o.LicensedLimit != nil {
		return true
	}

	return false
}

// SetLicensedLimit gets a reference to the given string and assigns it to the LicensedLimit field.
func (o *IaasLicenseUtilizationInfo) SetLicensedLimit(v string) {
	o.LicensedLimit = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *IaasLicenseUtilizationInfo) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseUtilizationInfo) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *IaasLicenseUtilizationInfo) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *IaasLicenseUtilizationInfo) SetSku(v string) {
	o.Sku = &v
}

func (o IaasLicenseUtilizationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActualUsed != nil {
		toSerialize["ActualUsed"] = o.ActualUsed
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.LicensedLimit != nil {
		toSerialize["LicensedLimit"] = o.LicensedLimit
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasLicenseUtilizationInfo) UnmarshalJSON(bytes []byte) (err error) {
	type IaasLicenseUtilizationInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Number of licenses actually used for this feature.
		ActualUsed *int64 `json:"ActualUsed,omitempty"`
		// License label of UCSD license.
		Label *string `json:"Label,omitempty"`
		// License limit for this license feature.
		LicensedLimit *string `json:"LicensedLimit,omitempty"`
		// SKU for the UCSD license.
		Sku *string `json:"Sku,omitempty"`
	}

	varIaasLicenseUtilizationInfoWithoutEmbeddedStruct := IaasLicenseUtilizationInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIaasLicenseUtilizationInfoWithoutEmbeddedStruct)
	if err == nil {
		varIaasLicenseUtilizationInfo := _IaasLicenseUtilizationInfo{}
		varIaasLicenseUtilizationInfo.ClassId = varIaasLicenseUtilizationInfoWithoutEmbeddedStruct.ClassId
		varIaasLicenseUtilizationInfo.ObjectType = varIaasLicenseUtilizationInfoWithoutEmbeddedStruct.ObjectType
		varIaasLicenseUtilizationInfo.ActualUsed = varIaasLicenseUtilizationInfoWithoutEmbeddedStruct.ActualUsed
		varIaasLicenseUtilizationInfo.Label = varIaasLicenseUtilizationInfoWithoutEmbeddedStruct.Label
		varIaasLicenseUtilizationInfo.LicensedLimit = varIaasLicenseUtilizationInfoWithoutEmbeddedStruct.LicensedLimit
		varIaasLicenseUtilizationInfo.Sku = varIaasLicenseUtilizationInfoWithoutEmbeddedStruct.Sku
		*o = IaasLicenseUtilizationInfo(varIaasLicenseUtilizationInfo)
	} else {
		return err
	}

	varIaasLicenseUtilizationInfo := _IaasLicenseUtilizationInfo{}

	err = json.Unmarshal(bytes, &varIaasLicenseUtilizationInfo)
	if err == nil {
		o.MoBaseComplexType = varIaasLicenseUtilizationInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActualUsed")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "LicensedLimit")
		delete(additionalProperties, "Sku")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableIaasLicenseUtilizationInfo struct {
	value *IaasLicenseUtilizationInfo
	isSet bool
}

func (v NullableIaasLicenseUtilizationInfo) Get() *IaasLicenseUtilizationInfo {
	return v.value
}

func (v *NullableIaasLicenseUtilizationInfo) Set(val *IaasLicenseUtilizationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasLicenseUtilizationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasLicenseUtilizationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasLicenseUtilizationInfo(val *IaasLicenseUtilizationInfo) *NullableIaasLicenseUtilizationInfo {
	return &NullableIaasLicenseUtilizationInfo{value: val, isSet: true}
}

func (v NullableIaasLicenseUtilizationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasLicenseUtilizationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
