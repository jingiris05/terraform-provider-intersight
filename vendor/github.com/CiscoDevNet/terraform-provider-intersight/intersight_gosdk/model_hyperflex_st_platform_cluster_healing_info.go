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

// checks if the HyperflexStPlatformClusterHealingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexStPlatformClusterHealingInfo{}

// HyperflexStPlatformClusterHealingInfo The detailed status of the cluster healing process.
type HyperflexStPlatformClusterHealingInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The estimated time in seconds it will take to complete the auto-healing process.
	EstimatedCompletionTimeInSeconds *int64 `json:"EstimatedCompletionTimeInSeconds,omitempty"`
	// The status of the cluster's auto-healing process. If 'true', auto-healing is in progress for the cluster.
	InProgress *bool    `json:"InProgress,omitempty"`
	Messages   []string `json:"Messages,omitempty"`
	// The current message describing the auto-healing process of the cluster.
	MessagesIterator interface{} `json:"MessagesIterator,omitempty"`
	// The number of elements in the messages collection.
	MessagesSize *int64 `json:"MessagesSize,omitempty"`
	// The progress of the auto-healing process as a percentage.
	PercentComplete      *int64 `json:"PercentComplete,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexStPlatformClusterHealingInfo HyperflexStPlatformClusterHealingInfo

// NewHyperflexStPlatformClusterHealingInfo instantiates a new HyperflexStPlatformClusterHealingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexStPlatformClusterHealingInfo(classId string, objectType string) *HyperflexStPlatformClusterHealingInfo {
	this := HyperflexStPlatformClusterHealingInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexStPlatformClusterHealingInfoWithDefaults instantiates a new HyperflexStPlatformClusterHealingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexStPlatformClusterHealingInfoWithDefaults() *HyperflexStPlatformClusterHealingInfo {
	this := HyperflexStPlatformClusterHealingInfo{}
	var classId string = "hyperflex.StPlatformClusterHealingInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.StPlatformClusterHealingInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexStPlatformClusterHealingInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexStPlatformClusterHealingInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.StPlatformClusterHealingInfo" of the ClassId field.
func (o *HyperflexStPlatformClusterHealingInfo) GetDefaultClassId() interface{} {
	return "hyperflex.StPlatformClusterHealingInfo"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexStPlatformClusterHealingInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexStPlatformClusterHealingInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.StPlatformClusterHealingInfo" of the ObjectType field.
func (o *HyperflexStPlatformClusterHealingInfo) GetDefaultObjectType() interface{} {
	return "hyperflex.StPlatformClusterHealingInfo"
}

// GetEstimatedCompletionTimeInSeconds returns the EstimatedCompletionTimeInSeconds field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterHealingInfo) GetEstimatedCompletionTimeInSeconds() int64 {
	if o == nil || IsNil(o.EstimatedCompletionTimeInSeconds) {
		var ret int64
		return ret
	}
	return *o.EstimatedCompletionTimeInSeconds
}

// GetEstimatedCompletionTimeInSecondsOk returns a tuple with the EstimatedCompletionTimeInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfo) GetEstimatedCompletionTimeInSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.EstimatedCompletionTimeInSeconds) {
		return nil, false
	}
	return o.EstimatedCompletionTimeInSeconds, true
}

// HasEstimatedCompletionTimeInSeconds returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfo) HasEstimatedCompletionTimeInSeconds() bool {
	if o != nil && !IsNil(o.EstimatedCompletionTimeInSeconds) {
		return true
	}

	return false
}

// SetEstimatedCompletionTimeInSeconds gets a reference to the given int64 and assigns it to the EstimatedCompletionTimeInSeconds field.
func (o *HyperflexStPlatformClusterHealingInfo) SetEstimatedCompletionTimeInSeconds(v int64) {
	o.EstimatedCompletionTimeInSeconds = &v
}

// GetInProgress returns the InProgress field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterHealingInfo) GetInProgress() bool {
	if o == nil || IsNil(o.InProgress) {
		var ret bool
		return ret
	}
	return *o.InProgress
}

// GetInProgressOk returns a tuple with the InProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfo) GetInProgressOk() (*bool, bool) {
	if o == nil || IsNil(o.InProgress) {
		return nil, false
	}
	return o.InProgress, true
}

// HasInProgress returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfo) HasInProgress() bool {
	if o != nil && !IsNil(o.InProgress) {
		return true
	}

	return false
}

// SetInProgress gets a reference to the given bool and assigns it to the InProgress field.
func (o *HyperflexStPlatformClusterHealingInfo) SetInProgress(v bool) {
	o.InProgress = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexStPlatformClusterHealingInfo) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexStPlatformClusterHealingInfo) GetMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfo) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *HyperflexStPlatformClusterHealingInfo) SetMessages(v []string) {
	o.Messages = v
}

// GetMessagesIterator returns the MessagesIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexStPlatformClusterHealingInfo) GetMessagesIterator() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MessagesIterator
}

// GetMessagesIteratorOk returns a tuple with the MessagesIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexStPlatformClusterHealingInfo) GetMessagesIteratorOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MessagesIterator) {
		return nil, false
	}
	return &o.MessagesIterator, true
}

// HasMessagesIterator returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfo) HasMessagesIterator() bool {
	if o != nil && !IsNil(o.MessagesIterator) {
		return true
	}

	return false
}

// SetMessagesIterator gets a reference to the given interface{} and assigns it to the MessagesIterator field.
func (o *HyperflexStPlatformClusterHealingInfo) SetMessagesIterator(v interface{}) {
	o.MessagesIterator = v
}

// GetMessagesSize returns the MessagesSize field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterHealingInfo) GetMessagesSize() int64 {
	if o == nil || IsNil(o.MessagesSize) {
		var ret int64
		return ret
	}
	return *o.MessagesSize
}

// GetMessagesSizeOk returns a tuple with the MessagesSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfo) GetMessagesSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MessagesSize) {
		return nil, false
	}
	return o.MessagesSize, true
}

// HasMessagesSize returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfo) HasMessagesSize() bool {
	if o != nil && !IsNil(o.MessagesSize) {
		return true
	}

	return false
}

// SetMessagesSize gets a reference to the given int64 and assigns it to the MessagesSize field.
func (o *HyperflexStPlatformClusterHealingInfo) SetMessagesSize(v int64) {
	o.MessagesSize = &v
}

// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterHealingInfo) GetPercentComplete() int64 {
	if o == nil || IsNil(o.PercentComplete) {
		var ret int64
		return ret
	}
	return *o.PercentComplete
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfo) GetPercentCompleteOk() (*int64, bool) {
	if o == nil || IsNil(o.PercentComplete) {
		return nil, false
	}
	return o.PercentComplete, true
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfo) HasPercentComplete() bool {
	if o != nil && !IsNil(o.PercentComplete) {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given int64 and assigns it to the PercentComplete field.
func (o *HyperflexStPlatformClusterHealingInfo) SetPercentComplete(v int64) {
	o.PercentComplete = &v
}

func (o HyperflexStPlatformClusterHealingInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexStPlatformClusterHealingInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EstimatedCompletionTimeInSeconds) {
		toSerialize["EstimatedCompletionTimeInSeconds"] = o.EstimatedCompletionTimeInSeconds
	}
	if !IsNil(o.InProgress) {
		toSerialize["InProgress"] = o.InProgress
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.MessagesIterator != nil {
		toSerialize["MessagesIterator"] = o.MessagesIterator
	}
	if !IsNil(o.MessagesSize) {
		toSerialize["MessagesSize"] = o.MessagesSize
	}
	if !IsNil(o.PercentComplete) {
		toSerialize["PercentComplete"] = o.PercentComplete
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexStPlatformClusterHealingInfo) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The estimated time in seconds it will take to complete the auto-healing process.
		EstimatedCompletionTimeInSeconds *int64 `json:"EstimatedCompletionTimeInSeconds,omitempty"`
		// The status of the cluster's auto-healing process. If 'true', auto-healing is in progress for the cluster.
		InProgress *bool    `json:"InProgress,omitempty"`
		Messages   []string `json:"Messages,omitempty"`
		// The current message describing the auto-healing process of the cluster.
		MessagesIterator interface{} `json:"MessagesIterator,omitempty"`
		// The number of elements in the messages collection.
		MessagesSize *int64 `json:"MessagesSize,omitempty"`
		// The progress of the auto-healing process as a percentage.
		PercentComplete *int64 `json:"PercentComplete,omitempty"`
	}

	varHyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct := HyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexStPlatformClusterHealingInfo := _HyperflexStPlatformClusterHealingInfo{}
		varHyperflexStPlatformClusterHealingInfo.ClassId = varHyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct.ClassId
		varHyperflexStPlatformClusterHealingInfo.ObjectType = varHyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct.ObjectType
		varHyperflexStPlatformClusterHealingInfo.EstimatedCompletionTimeInSeconds = varHyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct.EstimatedCompletionTimeInSeconds
		varHyperflexStPlatformClusterHealingInfo.InProgress = varHyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct.InProgress
		varHyperflexStPlatformClusterHealingInfo.Messages = varHyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct.Messages
		varHyperflexStPlatformClusterHealingInfo.MessagesIterator = varHyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct.MessagesIterator
		varHyperflexStPlatformClusterHealingInfo.MessagesSize = varHyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct.MessagesSize
		varHyperflexStPlatformClusterHealingInfo.PercentComplete = varHyperflexStPlatformClusterHealingInfoWithoutEmbeddedStruct.PercentComplete
		*o = HyperflexStPlatformClusterHealingInfo(varHyperflexStPlatformClusterHealingInfo)
	} else {
		return err
	}

	varHyperflexStPlatformClusterHealingInfo := _HyperflexStPlatformClusterHealingInfo{}

	err = json.Unmarshal(data, &varHyperflexStPlatformClusterHealingInfo)
	if err == nil {
		o.MoBaseComplexType = varHyperflexStPlatformClusterHealingInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EstimatedCompletionTimeInSeconds")
		delete(additionalProperties, "InProgress")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "MessagesIterator")
		delete(additionalProperties, "MessagesSize")
		delete(additionalProperties, "PercentComplete")

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

type NullableHyperflexStPlatformClusterHealingInfo struct {
	value *HyperflexStPlatformClusterHealingInfo
	isSet bool
}

func (v NullableHyperflexStPlatformClusterHealingInfo) Get() *HyperflexStPlatformClusterHealingInfo {
	return v.value
}

func (v *NullableHyperflexStPlatformClusterHealingInfo) Set(val *HyperflexStPlatformClusterHealingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexStPlatformClusterHealingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexStPlatformClusterHealingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexStPlatformClusterHealingInfo(val *HyperflexStPlatformClusterHealingInfo) *NullableHyperflexStPlatformClusterHealingInfo {
	return &NullableHyperflexStPlatformClusterHealingInfo{value: val, isSet: true}
}

func (v NullableHyperflexStPlatformClusterHealingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexStPlatformClusterHealingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
