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
)

// ConvergedinfraBasePodSummary Summary information for the base pod. This inclues properties like number of nodes, stroage capacity/utilization, aggregated alarms for all the components of the pod etc.
type ConvergedinfraBasePodSummary struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType        string                                  `json:"ObjectType"`
	AlarmSummary      NullableConvergedinfraAlarmSummary      `json:"AlarmSummary,omitempty"`
	ComplianceSummary NullableConvergedinfraComplianceSummary `json:"ComplianceSummary,omitempty"`
	// Number of nodes associated with the pod.
	NodeCount *int64 `json:"NodeCount,omitempty"`
	// The available storage capacity for this pod.
	StorageAvailable *int64 `json:"StorageAvailable,omitempty"`
	// The total storage capacity for this pod.
	StorageCapacity *int64 `json:"StorageCapacity,omitempty"`
	// The percentage storage utilization for this pod.
	StorageUtilization   *float32 `json:"StorageUtilization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraBasePodSummary ConvergedinfraBasePodSummary

// NewConvergedinfraBasePodSummary instantiates a new ConvergedinfraBasePodSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraBasePodSummary(classId string, objectType string) *ConvergedinfraBasePodSummary {
	this := ConvergedinfraBasePodSummary{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraBasePodSummaryWithDefaults instantiates a new ConvergedinfraBasePodSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraBasePodSummaryWithDefaults() *ConvergedinfraBasePodSummary {
	this := ConvergedinfraBasePodSummary{}
	var classId string = "convergedinfra.PodSummary"
	this.ClassId = classId
	var objectType string = "convergedinfra.PodSummary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraBasePodSummary) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePodSummary) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraBasePodSummary) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraBasePodSummary) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePodSummary) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraBasePodSummary) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlarmSummary returns the AlarmSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraBasePodSummary) GetAlarmSummary() ConvergedinfraAlarmSummary {
	if o == nil || o.AlarmSummary.Get() == nil {
		var ret ConvergedinfraAlarmSummary
		return ret
	}
	return *o.AlarmSummary.Get()
}

// GetAlarmSummaryOk returns a tuple with the AlarmSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraBasePodSummary) GetAlarmSummaryOk() (*ConvergedinfraAlarmSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlarmSummary.Get(), o.AlarmSummary.IsSet()
}

// HasAlarmSummary returns a boolean if a field has been set.
func (o *ConvergedinfraBasePodSummary) HasAlarmSummary() bool {
	if o != nil && o.AlarmSummary.IsSet() {
		return true
	}

	return false
}

// SetAlarmSummary gets a reference to the given NullableConvergedinfraAlarmSummary and assigns it to the AlarmSummary field.
func (o *ConvergedinfraBasePodSummary) SetAlarmSummary(v ConvergedinfraAlarmSummary) {
	o.AlarmSummary.Set(&v)
}

// SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil
func (o *ConvergedinfraBasePodSummary) SetAlarmSummaryNil() {
	o.AlarmSummary.Set(nil)
}

// UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
func (o *ConvergedinfraBasePodSummary) UnsetAlarmSummary() {
	o.AlarmSummary.Unset()
}

// GetComplianceSummary returns the ComplianceSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraBasePodSummary) GetComplianceSummary() ConvergedinfraComplianceSummary {
	if o == nil || o.ComplianceSummary.Get() == nil {
		var ret ConvergedinfraComplianceSummary
		return ret
	}
	return *o.ComplianceSummary.Get()
}

// GetComplianceSummaryOk returns a tuple with the ComplianceSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraBasePodSummary) GetComplianceSummaryOk() (*ConvergedinfraComplianceSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComplianceSummary.Get(), o.ComplianceSummary.IsSet()
}

// HasComplianceSummary returns a boolean if a field has been set.
func (o *ConvergedinfraBasePodSummary) HasComplianceSummary() bool {
	if o != nil && o.ComplianceSummary.IsSet() {
		return true
	}

	return false
}

// SetComplianceSummary gets a reference to the given NullableConvergedinfraComplianceSummary and assigns it to the ComplianceSummary field.
func (o *ConvergedinfraBasePodSummary) SetComplianceSummary(v ConvergedinfraComplianceSummary) {
	o.ComplianceSummary.Set(&v)
}

// SetComplianceSummaryNil sets the value for ComplianceSummary to be an explicit nil
func (o *ConvergedinfraBasePodSummary) SetComplianceSummaryNil() {
	o.ComplianceSummary.Set(nil)
}

// UnsetComplianceSummary ensures that no value is present for ComplianceSummary, not even an explicit nil
func (o *ConvergedinfraBasePodSummary) UnsetComplianceSummary() {
	o.ComplianceSummary.Unset()
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *ConvergedinfraBasePodSummary) GetNodeCount() int64 {
	if o == nil || o.NodeCount == nil {
		var ret int64
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePodSummary) GetNodeCountOk() (*int64, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *ConvergedinfraBasePodSummary) HasNodeCount() bool {
	if o != nil && o.NodeCount != nil {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int64 and assigns it to the NodeCount field.
func (o *ConvergedinfraBasePodSummary) SetNodeCount(v int64) {
	o.NodeCount = &v
}

// GetStorageAvailable returns the StorageAvailable field value if set, zero value otherwise.
func (o *ConvergedinfraBasePodSummary) GetStorageAvailable() int64 {
	if o == nil || o.StorageAvailable == nil {
		var ret int64
		return ret
	}
	return *o.StorageAvailable
}

// GetStorageAvailableOk returns a tuple with the StorageAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePodSummary) GetStorageAvailableOk() (*int64, bool) {
	if o == nil || o.StorageAvailable == nil {
		return nil, false
	}
	return o.StorageAvailable, true
}

// HasStorageAvailable returns a boolean if a field has been set.
func (o *ConvergedinfraBasePodSummary) HasStorageAvailable() bool {
	if o != nil && o.StorageAvailable != nil {
		return true
	}

	return false
}

// SetStorageAvailable gets a reference to the given int64 and assigns it to the StorageAvailable field.
func (o *ConvergedinfraBasePodSummary) SetStorageAvailable(v int64) {
	o.StorageAvailable = &v
}

// GetStorageCapacity returns the StorageCapacity field value if set, zero value otherwise.
func (o *ConvergedinfraBasePodSummary) GetStorageCapacity() int64 {
	if o == nil || o.StorageCapacity == nil {
		var ret int64
		return ret
	}
	return *o.StorageCapacity
}

// GetStorageCapacityOk returns a tuple with the StorageCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePodSummary) GetStorageCapacityOk() (*int64, bool) {
	if o == nil || o.StorageCapacity == nil {
		return nil, false
	}
	return o.StorageCapacity, true
}

// HasStorageCapacity returns a boolean if a field has been set.
func (o *ConvergedinfraBasePodSummary) HasStorageCapacity() bool {
	if o != nil && o.StorageCapacity != nil {
		return true
	}

	return false
}

// SetStorageCapacity gets a reference to the given int64 and assigns it to the StorageCapacity field.
func (o *ConvergedinfraBasePodSummary) SetStorageCapacity(v int64) {
	o.StorageCapacity = &v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise.
func (o *ConvergedinfraBasePodSummary) GetStorageUtilization() float32 {
	if o == nil || o.StorageUtilization == nil {
		var ret float32
		return ret
	}
	return *o.StorageUtilization
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBasePodSummary) GetStorageUtilizationOk() (*float32, bool) {
	if o == nil || o.StorageUtilization == nil {
		return nil, false
	}
	return o.StorageUtilization, true
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *ConvergedinfraBasePodSummary) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization != nil {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given float32 and assigns it to the StorageUtilization field.
func (o *ConvergedinfraBasePodSummary) SetStorageUtilization(v float32) {
	o.StorageUtilization = &v
}

func (o ConvergedinfraBasePodSummary) MarshalJSON() ([]byte, error) {
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
	if o.AlarmSummary.IsSet() {
		toSerialize["AlarmSummary"] = o.AlarmSummary.Get()
	}
	if o.ComplianceSummary.IsSet() {
		toSerialize["ComplianceSummary"] = o.ComplianceSummary.Get()
	}
	if o.NodeCount != nil {
		toSerialize["NodeCount"] = o.NodeCount
	}
	if o.StorageAvailable != nil {
		toSerialize["StorageAvailable"] = o.StorageAvailable
	}
	if o.StorageCapacity != nil {
		toSerialize["StorageCapacity"] = o.StorageCapacity
	}
	if o.StorageUtilization != nil {
		toSerialize["StorageUtilization"] = o.StorageUtilization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraBasePodSummary) UnmarshalJSON(bytes []byte) (err error) {
	type ConvergedinfraBasePodSummaryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType        string                                  `json:"ObjectType"`
		AlarmSummary      NullableConvergedinfraAlarmSummary      `json:"AlarmSummary,omitempty"`
		ComplianceSummary NullableConvergedinfraComplianceSummary `json:"ComplianceSummary,omitempty"`
		// Number of nodes associated with the pod.
		NodeCount *int64 `json:"NodeCount,omitempty"`
		// The available storage capacity for this pod.
		StorageAvailable *int64 `json:"StorageAvailable,omitempty"`
		// The total storage capacity for this pod.
		StorageCapacity *int64 `json:"StorageCapacity,omitempty"`
		// The percentage storage utilization for this pod.
		StorageUtilization *float32 `json:"StorageUtilization,omitempty"`
	}

	varConvergedinfraBasePodSummaryWithoutEmbeddedStruct := ConvergedinfraBasePodSummaryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConvergedinfraBasePodSummaryWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraBasePodSummary := _ConvergedinfraBasePodSummary{}
		varConvergedinfraBasePodSummary.ClassId = varConvergedinfraBasePodSummaryWithoutEmbeddedStruct.ClassId
		varConvergedinfraBasePodSummary.ObjectType = varConvergedinfraBasePodSummaryWithoutEmbeddedStruct.ObjectType
		varConvergedinfraBasePodSummary.AlarmSummary = varConvergedinfraBasePodSummaryWithoutEmbeddedStruct.AlarmSummary
		varConvergedinfraBasePodSummary.ComplianceSummary = varConvergedinfraBasePodSummaryWithoutEmbeddedStruct.ComplianceSummary
		varConvergedinfraBasePodSummary.NodeCount = varConvergedinfraBasePodSummaryWithoutEmbeddedStruct.NodeCount
		varConvergedinfraBasePodSummary.StorageAvailable = varConvergedinfraBasePodSummaryWithoutEmbeddedStruct.StorageAvailable
		varConvergedinfraBasePodSummary.StorageCapacity = varConvergedinfraBasePodSummaryWithoutEmbeddedStruct.StorageCapacity
		varConvergedinfraBasePodSummary.StorageUtilization = varConvergedinfraBasePodSummaryWithoutEmbeddedStruct.StorageUtilization
		*o = ConvergedinfraBasePodSummary(varConvergedinfraBasePodSummary)
	} else {
		return err
	}

	varConvergedinfraBasePodSummary := _ConvergedinfraBasePodSummary{}

	err = json.Unmarshal(bytes, &varConvergedinfraBasePodSummary)
	if err == nil {
		o.MoBaseComplexType = varConvergedinfraBasePodSummary.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlarmSummary")
		delete(additionalProperties, "ComplianceSummary")
		delete(additionalProperties, "NodeCount")
		delete(additionalProperties, "StorageAvailable")
		delete(additionalProperties, "StorageCapacity")
		delete(additionalProperties, "StorageUtilization")

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

type NullableConvergedinfraBasePodSummary struct {
	value *ConvergedinfraBasePodSummary
	isSet bool
}

func (v NullableConvergedinfraBasePodSummary) Get() *ConvergedinfraBasePodSummary {
	return v.value
}

func (v *NullableConvergedinfraBasePodSummary) Set(val *ConvergedinfraBasePodSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraBasePodSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraBasePodSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraBasePodSummary(val *ConvergedinfraBasePodSummary) *NullableConvergedinfraBasePodSummary {
	return &NullableConvergedinfraBasePodSummary{value: val, isSet: true}
}

func (v NullableConvergedinfraBasePodSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraBasePodSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
