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

// checks if the CapabilityAdapterFirmwareRequirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityAdapterFirmwareRequirement{}

// CapabilityAdapterFirmwareRequirement Firmware requirements for enabling Intersight based management for an adaptor.
type CapabilityAdapterFirmwareRequirement struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Series of adapter. Example - Cruz, Bodega.
	AdapterSeries *string `json:"AdapterSeries,omitempty"`
	// Do not update if the current version is reported as empty.
	IgnoreEmptyCurrentVersion *bool `json:"IgnoreEmptyCurrentVersion,omitempty"`
	// The minimum adapter version that supports this adapter upgrade.
	MinimumAdapterVersion *string `json:"MinimumAdapterVersion,omitempty"`
	// The minimum BMC version that supports this adapter upgrade.
	MinimumBmcVersion *string `json:"MinimumBmcVersion,omitempty"`
	// The recommended BMC version that supports this adapter upgrade.
	RecommendedBmcVersion *string  `json:"RecommendedBmcVersion,omitempty"`
	SupportedModels       []string `json:"SupportedModels,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _CapabilityAdapterFirmwareRequirement CapabilityAdapterFirmwareRequirement

// NewCapabilityAdapterFirmwareRequirement instantiates a new CapabilityAdapterFirmwareRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityAdapterFirmwareRequirement(classId string, objectType string) *CapabilityAdapterFirmwareRequirement {
	this := CapabilityAdapterFirmwareRequirement{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityAdapterFirmwareRequirementWithDefaults instantiates a new CapabilityAdapterFirmwareRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityAdapterFirmwareRequirementWithDefaults() *CapabilityAdapterFirmwareRequirement {
	this := CapabilityAdapterFirmwareRequirement{}
	var classId string = "capability.AdapterFirmwareRequirement"
	this.ClassId = classId
	var objectType string = "capability.AdapterFirmwareRequirement"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityAdapterFirmwareRequirement) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterFirmwareRequirement) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityAdapterFirmwareRequirement) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.AdapterFirmwareRequirement" of the ClassId field.
func (o *CapabilityAdapterFirmwareRequirement) GetDefaultClassId() interface{} {
	return "capability.AdapterFirmwareRequirement"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityAdapterFirmwareRequirement) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterFirmwareRequirement) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityAdapterFirmwareRequirement) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.AdapterFirmwareRequirement" of the ObjectType field.
func (o *CapabilityAdapterFirmwareRequirement) GetDefaultObjectType() interface{} {
	return "capability.AdapterFirmwareRequirement"
}

// GetAdapterSeries returns the AdapterSeries field value if set, zero value otherwise.
func (o *CapabilityAdapterFirmwareRequirement) GetAdapterSeries() string {
	if o == nil || IsNil(o.AdapterSeries) {
		var ret string
		return ret
	}
	return *o.AdapterSeries
}

// GetAdapterSeriesOk returns a tuple with the AdapterSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterFirmwareRequirement) GetAdapterSeriesOk() (*string, bool) {
	if o == nil || IsNil(o.AdapterSeries) {
		return nil, false
	}
	return o.AdapterSeries, true
}

// HasAdapterSeries returns a boolean if a field has been set.
func (o *CapabilityAdapterFirmwareRequirement) HasAdapterSeries() bool {
	if o != nil && !IsNil(o.AdapterSeries) {
		return true
	}

	return false
}

// SetAdapterSeries gets a reference to the given string and assigns it to the AdapterSeries field.
func (o *CapabilityAdapterFirmwareRequirement) SetAdapterSeries(v string) {
	o.AdapterSeries = &v
}

// GetIgnoreEmptyCurrentVersion returns the IgnoreEmptyCurrentVersion field value if set, zero value otherwise.
func (o *CapabilityAdapterFirmwareRequirement) GetIgnoreEmptyCurrentVersion() bool {
	if o == nil || IsNil(o.IgnoreEmptyCurrentVersion) {
		var ret bool
		return ret
	}
	return *o.IgnoreEmptyCurrentVersion
}

// GetIgnoreEmptyCurrentVersionOk returns a tuple with the IgnoreEmptyCurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterFirmwareRequirement) GetIgnoreEmptyCurrentVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreEmptyCurrentVersion) {
		return nil, false
	}
	return o.IgnoreEmptyCurrentVersion, true
}

// HasIgnoreEmptyCurrentVersion returns a boolean if a field has been set.
func (o *CapabilityAdapterFirmwareRequirement) HasIgnoreEmptyCurrentVersion() bool {
	if o != nil && !IsNil(o.IgnoreEmptyCurrentVersion) {
		return true
	}

	return false
}

// SetIgnoreEmptyCurrentVersion gets a reference to the given bool and assigns it to the IgnoreEmptyCurrentVersion field.
func (o *CapabilityAdapterFirmwareRequirement) SetIgnoreEmptyCurrentVersion(v bool) {
	o.IgnoreEmptyCurrentVersion = &v
}

// GetMinimumAdapterVersion returns the MinimumAdapterVersion field value if set, zero value otherwise.
func (o *CapabilityAdapterFirmwareRequirement) GetMinimumAdapterVersion() string {
	if o == nil || IsNil(o.MinimumAdapterVersion) {
		var ret string
		return ret
	}
	return *o.MinimumAdapterVersion
}

// GetMinimumAdapterVersionOk returns a tuple with the MinimumAdapterVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterFirmwareRequirement) GetMinimumAdapterVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumAdapterVersion) {
		return nil, false
	}
	return o.MinimumAdapterVersion, true
}

// HasMinimumAdapterVersion returns a boolean if a field has been set.
func (o *CapabilityAdapterFirmwareRequirement) HasMinimumAdapterVersion() bool {
	if o != nil && !IsNil(o.MinimumAdapterVersion) {
		return true
	}

	return false
}

// SetMinimumAdapterVersion gets a reference to the given string and assigns it to the MinimumAdapterVersion field.
func (o *CapabilityAdapterFirmwareRequirement) SetMinimumAdapterVersion(v string) {
	o.MinimumAdapterVersion = &v
}

// GetMinimumBmcVersion returns the MinimumBmcVersion field value if set, zero value otherwise.
func (o *CapabilityAdapterFirmwareRequirement) GetMinimumBmcVersion() string {
	if o == nil || IsNil(o.MinimumBmcVersion) {
		var ret string
		return ret
	}
	return *o.MinimumBmcVersion
}

// GetMinimumBmcVersionOk returns a tuple with the MinimumBmcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterFirmwareRequirement) GetMinimumBmcVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumBmcVersion) {
		return nil, false
	}
	return o.MinimumBmcVersion, true
}

// HasMinimumBmcVersion returns a boolean if a field has been set.
func (o *CapabilityAdapterFirmwareRequirement) HasMinimumBmcVersion() bool {
	if o != nil && !IsNil(o.MinimumBmcVersion) {
		return true
	}

	return false
}

// SetMinimumBmcVersion gets a reference to the given string and assigns it to the MinimumBmcVersion field.
func (o *CapabilityAdapterFirmwareRequirement) SetMinimumBmcVersion(v string) {
	o.MinimumBmcVersion = &v
}

// GetRecommendedBmcVersion returns the RecommendedBmcVersion field value if set, zero value otherwise.
func (o *CapabilityAdapterFirmwareRequirement) GetRecommendedBmcVersion() string {
	if o == nil || IsNil(o.RecommendedBmcVersion) {
		var ret string
		return ret
	}
	return *o.RecommendedBmcVersion
}

// GetRecommendedBmcVersionOk returns a tuple with the RecommendedBmcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterFirmwareRequirement) GetRecommendedBmcVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendedBmcVersion) {
		return nil, false
	}
	return o.RecommendedBmcVersion, true
}

// HasRecommendedBmcVersion returns a boolean if a field has been set.
func (o *CapabilityAdapterFirmwareRequirement) HasRecommendedBmcVersion() bool {
	if o != nil && !IsNil(o.RecommendedBmcVersion) {
		return true
	}

	return false
}

// SetRecommendedBmcVersion gets a reference to the given string and assigns it to the RecommendedBmcVersion field.
func (o *CapabilityAdapterFirmwareRequirement) SetRecommendedBmcVersion(v string) {
	o.RecommendedBmcVersion = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityAdapterFirmwareRequirement) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityAdapterFirmwareRequirement) GetSupportedModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedModels) {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *CapabilityAdapterFirmwareRequirement) HasSupportedModels() bool {
	if o != nil && !IsNil(o.SupportedModels) {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *CapabilityAdapterFirmwareRequirement) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

func (o CapabilityAdapterFirmwareRequirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityAdapterFirmwareRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return map[string]interface{}{}, errCapabilityCapability
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdapterSeries) {
		toSerialize["AdapterSeries"] = o.AdapterSeries
	}
	if !IsNil(o.IgnoreEmptyCurrentVersion) {
		toSerialize["IgnoreEmptyCurrentVersion"] = o.IgnoreEmptyCurrentVersion
	}
	if !IsNil(o.MinimumAdapterVersion) {
		toSerialize["MinimumAdapterVersion"] = o.MinimumAdapterVersion
	}
	if !IsNil(o.MinimumBmcVersion) {
		toSerialize["MinimumBmcVersion"] = o.MinimumBmcVersion
	}
	if !IsNil(o.RecommendedBmcVersion) {
		toSerialize["RecommendedBmcVersion"] = o.RecommendedBmcVersion
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityAdapterFirmwareRequirement) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Series of adapter. Example - Cruz, Bodega.
		AdapterSeries *string `json:"AdapterSeries,omitempty"`
		// Do not update if the current version is reported as empty.
		IgnoreEmptyCurrentVersion *bool `json:"IgnoreEmptyCurrentVersion,omitempty"`
		// The minimum adapter version that supports this adapter upgrade.
		MinimumAdapterVersion *string `json:"MinimumAdapterVersion,omitempty"`
		// The minimum BMC version that supports this adapter upgrade.
		MinimumBmcVersion *string `json:"MinimumBmcVersion,omitempty"`
		// The recommended BMC version that supports this adapter upgrade.
		RecommendedBmcVersion *string  `json:"RecommendedBmcVersion,omitempty"`
		SupportedModels       []string `json:"SupportedModels,omitempty"`
	}

	varCapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct := CapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityAdapterFirmwareRequirement := _CapabilityAdapterFirmwareRequirement{}
		varCapabilityAdapterFirmwareRequirement.ClassId = varCapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct.ClassId
		varCapabilityAdapterFirmwareRequirement.ObjectType = varCapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct.ObjectType
		varCapabilityAdapterFirmwareRequirement.AdapterSeries = varCapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct.AdapterSeries
		varCapabilityAdapterFirmwareRequirement.IgnoreEmptyCurrentVersion = varCapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct.IgnoreEmptyCurrentVersion
		varCapabilityAdapterFirmwareRequirement.MinimumAdapterVersion = varCapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct.MinimumAdapterVersion
		varCapabilityAdapterFirmwareRequirement.MinimumBmcVersion = varCapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct.MinimumBmcVersion
		varCapabilityAdapterFirmwareRequirement.RecommendedBmcVersion = varCapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct.RecommendedBmcVersion
		varCapabilityAdapterFirmwareRequirement.SupportedModels = varCapabilityAdapterFirmwareRequirementWithoutEmbeddedStruct.SupportedModels
		*o = CapabilityAdapterFirmwareRequirement(varCapabilityAdapterFirmwareRequirement)
	} else {
		return err
	}

	varCapabilityAdapterFirmwareRequirement := _CapabilityAdapterFirmwareRequirement{}

	err = json.Unmarshal(data, &varCapabilityAdapterFirmwareRequirement)
	if err == nil {
		o.CapabilityCapability = varCapabilityAdapterFirmwareRequirement.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterSeries")
		delete(additionalProperties, "IgnoreEmptyCurrentVersion")
		delete(additionalProperties, "MinimumAdapterVersion")
		delete(additionalProperties, "MinimumBmcVersion")
		delete(additionalProperties, "RecommendedBmcVersion")
		delete(additionalProperties, "SupportedModels")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableCapabilityAdapterFirmwareRequirement struct {
	value *CapabilityAdapterFirmwareRequirement
	isSet bool
}

func (v NullableCapabilityAdapterFirmwareRequirement) Get() *CapabilityAdapterFirmwareRequirement {
	return v.value
}

func (v *NullableCapabilityAdapterFirmwareRequirement) Set(val *CapabilityAdapterFirmwareRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityAdapterFirmwareRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityAdapterFirmwareRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityAdapterFirmwareRequirement(val *CapabilityAdapterFirmwareRequirement) *NullableCapabilityAdapterFirmwareRequirement {
	return &NullableCapabilityAdapterFirmwareRequirement{value: val, isSet: true}
}

func (v NullableCapabilityAdapterFirmwareRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityAdapterFirmwareRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
