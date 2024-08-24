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

// checks if the HyperflexProtectionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexProtectionInfo{}

// HyperflexProtectionInfo ProtectionInfo for this snapshot.
type HyperflexProtectionInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType                     string                                  `json:"ObjectType"`
	VmCurrentProtectionInfo        NullableHyperflexSnapshotInfoBrief      `json:"VmCurrentProtectionInfo,omitempty"`
	VmLastSuccessfulProtectionInfo NullableHyperflexSnapshotInfoBrief      `json:"VmLastSuccessfulProtectionInfo,omitempty"`
	VmSpaceUsage                   NullableHyperflexVmProtectionSpaceUsage `json:"VmSpaceUsage,omitempty"`
	AdditionalProperties           map[string]interface{}
}

type _HyperflexProtectionInfo HyperflexProtectionInfo

// NewHyperflexProtectionInfo instantiates a new HyperflexProtectionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexProtectionInfo(classId string, objectType string) *HyperflexProtectionInfo {
	this := HyperflexProtectionInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexProtectionInfoWithDefaults instantiates a new HyperflexProtectionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexProtectionInfoWithDefaults() *HyperflexProtectionInfo {
	this := HyperflexProtectionInfo{}
	var classId string = "hyperflex.ProtectionInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.ProtectionInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexProtectionInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexProtectionInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexProtectionInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ProtectionInfo" of the ClassId field.
func (o *HyperflexProtectionInfo) GetDefaultClassId() interface{} {
	return "hyperflex.ProtectionInfo"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexProtectionInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexProtectionInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexProtectionInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ProtectionInfo" of the ObjectType field.
func (o *HyperflexProtectionInfo) GetDefaultObjectType() interface{} {
	return "hyperflex.ProtectionInfo"
}

// GetVmCurrentProtectionInfo returns the VmCurrentProtectionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProtectionInfo) GetVmCurrentProtectionInfo() HyperflexSnapshotInfoBrief {
	if o == nil || IsNil(o.VmCurrentProtectionInfo.Get()) {
		var ret HyperflexSnapshotInfoBrief
		return ret
	}
	return *o.VmCurrentProtectionInfo.Get()
}

// GetVmCurrentProtectionInfoOk returns a tuple with the VmCurrentProtectionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProtectionInfo) GetVmCurrentProtectionInfoOk() (*HyperflexSnapshotInfoBrief, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmCurrentProtectionInfo.Get(), o.VmCurrentProtectionInfo.IsSet()
}

// HasVmCurrentProtectionInfo returns a boolean if a field has been set.
func (o *HyperflexProtectionInfo) HasVmCurrentProtectionInfo() bool {
	if o != nil && o.VmCurrentProtectionInfo.IsSet() {
		return true
	}

	return false
}

// SetVmCurrentProtectionInfo gets a reference to the given NullableHyperflexSnapshotInfoBrief and assigns it to the VmCurrentProtectionInfo field.
func (o *HyperflexProtectionInfo) SetVmCurrentProtectionInfo(v HyperflexSnapshotInfoBrief) {
	o.VmCurrentProtectionInfo.Set(&v)
}

// SetVmCurrentProtectionInfoNil sets the value for VmCurrentProtectionInfo to be an explicit nil
func (o *HyperflexProtectionInfo) SetVmCurrentProtectionInfoNil() {
	o.VmCurrentProtectionInfo.Set(nil)
}

// UnsetVmCurrentProtectionInfo ensures that no value is present for VmCurrentProtectionInfo, not even an explicit nil
func (o *HyperflexProtectionInfo) UnsetVmCurrentProtectionInfo() {
	o.VmCurrentProtectionInfo.Unset()
}

// GetVmLastSuccessfulProtectionInfo returns the VmLastSuccessfulProtectionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProtectionInfo) GetVmLastSuccessfulProtectionInfo() HyperflexSnapshotInfoBrief {
	if o == nil || IsNil(o.VmLastSuccessfulProtectionInfo.Get()) {
		var ret HyperflexSnapshotInfoBrief
		return ret
	}
	return *o.VmLastSuccessfulProtectionInfo.Get()
}

// GetVmLastSuccessfulProtectionInfoOk returns a tuple with the VmLastSuccessfulProtectionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProtectionInfo) GetVmLastSuccessfulProtectionInfoOk() (*HyperflexSnapshotInfoBrief, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmLastSuccessfulProtectionInfo.Get(), o.VmLastSuccessfulProtectionInfo.IsSet()
}

// HasVmLastSuccessfulProtectionInfo returns a boolean if a field has been set.
func (o *HyperflexProtectionInfo) HasVmLastSuccessfulProtectionInfo() bool {
	if o != nil && o.VmLastSuccessfulProtectionInfo.IsSet() {
		return true
	}

	return false
}

// SetVmLastSuccessfulProtectionInfo gets a reference to the given NullableHyperflexSnapshotInfoBrief and assigns it to the VmLastSuccessfulProtectionInfo field.
func (o *HyperflexProtectionInfo) SetVmLastSuccessfulProtectionInfo(v HyperflexSnapshotInfoBrief) {
	o.VmLastSuccessfulProtectionInfo.Set(&v)
}

// SetVmLastSuccessfulProtectionInfoNil sets the value for VmLastSuccessfulProtectionInfo to be an explicit nil
func (o *HyperflexProtectionInfo) SetVmLastSuccessfulProtectionInfoNil() {
	o.VmLastSuccessfulProtectionInfo.Set(nil)
}

// UnsetVmLastSuccessfulProtectionInfo ensures that no value is present for VmLastSuccessfulProtectionInfo, not even an explicit nil
func (o *HyperflexProtectionInfo) UnsetVmLastSuccessfulProtectionInfo() {
	o.VmLastSuccessfulProtectionInfo.Unset()
}

// GetVmSpaceUsage returns the VmSpaceUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProtectionInfo) GetVmSpaceUsage() HyperflexVmProtectionSpaceUsage {
	if o == nil || IsNil(o.VmSpaceUsage.Get()) {
		var ret HyperflexVmProtectionSpaceUsage
		return ret
	}
	return *o.VmSpaceUsage.Get()
}

// GetVmSpaceUsageOk returns a tuple with the VmSpaceUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProtectionInfo) GetVmSpaceUsageOk() (*HyperflexVmProtectionSpaceUsage, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmSpaceUsage.Get(), o.VmSpaceUsage.IsSet()
}

// HasVmSpaceUsage returns a boolean if a field has been set.
func (o *HyperflexProtectionInfo) HasVmSpaceUsage() bool {
	if o != nil && o.VmSpaceUsage.IsSet() {
		return true
	}

	return false
}

// SetVmSpaceUsage gets a reference to the given NullableHyperflexVmProtectionSpaceUsage and assigns it to the VmSpaceUsage field.
func (o *HyperflexProtectionInfo) SetVmSpaceUsage(v HyperflexVmProtectionSpaceUsage) {
	o.VmSpaceUsage.Set(&v)
}

// SetVmSpaceUsageNil sets the value for VmSpaceUsage to be an explicit nil
func (o *HyperflexProtectionInfo) SetVmSpaceUsageNil() {
	o.VmSpaceUsage.Set(nil)
}

// UnsetVmSpaceUsage ensures that no value is present for VmSpaceUsage, not even an explicit nil
func (o *HyperflexProtectionInfo) UnsetVmSpaceUsage() {
	o.VmSpaceUsage.Unset()
}

func (o HyperflexProtectionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexProtectionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.VmCurrentProtectionInfo.IsSet() {
		toSerialize["VmCurrentProtectionInfo"] = o.VmCurrentProtectionInfo.Get()
	}
	if o.VmLastSuccessfulProtectionInfo.IsSet() {
		toSerialize["VmLastSuccessfulProtectionInfo"] = o.VmLastSuccessfulProtectionInfo.Get()
	}
	if o.VmSpaceUsage.IsSet() {
		toSerialize["VmSpaceUsage"] = o.VmSpaceUsage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexProtectionInfo) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexProtectionInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType                     string                                  `json:"ObjectType"`
		VmCurrentProtectionInfo        NullableHyperflexSnapshotInfoBrief      `json:"VmCurrentProtectionInfo,omitempty"`
		VmLastSuccessfulProtectionInfo NullableHyperflexSnapshotInfoBrief      `json:"VmLastSuccessfulProtectionInfo,omitempty"`
		VmSpaceUsage                   NullableHyperflexVmProtectionSpaceUsage `json:"VmSpaceUsage,omitempty"`
	}

	varHyperflexProtectionInfoWithoutEmbeddedStruct := HyperflexProtectionInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexProtectionInfoWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexProtectionInfo := _HyperflexProtectionInfo{}
		varHyperflexProtectionInfo.ClassId = varHyperflexProtectionInfoWithoutEmbeddedStruct.ClassId
		varHyperflexProtectionInfo.ObjectType = varHyperflexProtectionInfoWithoutEmbeddedStruct.ObjectType
		varHyperflexProtectionInfo.VmCurrentProtectionInfo = varHyperflexProtectionInfoWithoutEmbeddedStruct.VmCurrentProtectionInfo
		varHyperflexProtectionInfo.VmLastSuccessfulProtectionInfo = varHyperflexProtectionInfoWithoutEmbeddedStruct.VmLastSuccessfulProtectionInfo
		varHyperflexProtectionInfo.VmSpaceUsage = varHyperflexProtectionInfoWithoutEmbeddedStruct.VmSpaceUsage
		*o = HyperflexProtectionInfo(varHyperflexProtectionInfo)
	} else {
		return err
	}

	varHyperflexProtectionInfo := _HyperflexProtectionInfo{}

	err = json.Unmarshal(data, &varHyperflexProtectionInfo)
	if err == nil {
		o.MoBaseComplexType = varHyperflexProtectionInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VmCurrentProtectionInfo")
		delete(additionalProperties, "VmLastSuccessfulProtectionInfo")
		delete(additionalProperties, "VmSpaceUsage")

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

type NullableHyperflexProtectionInfo struct {
	value *HyperflexProtectionInfo
	isSet bool
}

func (v NullableHyperflexProtectionInfo) Get() *HyperflexProtectionInfo {
	return v.value
}

func (v *NullableHyperflexProtectionInfo) Set(val *HyperflexProtectionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexProtectionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexProtectionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexProtectionInfo(val *HyperflexProtectionInfo) *NullableHyperflexProtectionInfo {
	return &NullableHyperflexProtectionInfo{value: val, isSet: true}
}

func (v NullableHyperflexProtectionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexProtectionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
