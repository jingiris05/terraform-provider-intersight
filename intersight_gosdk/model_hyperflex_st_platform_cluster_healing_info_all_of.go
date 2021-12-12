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
)

// HyperflexStPlatformClusterHealingInfoAllOf Definition of the list of properties defined in 'hyperflex.StPlatformClusterHealingInfo', excluding properties defined in parent classes.
type HyperflexStPlatformClusterHealingInfoAllOf struct {
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

type _HyperflexStPlatformClusterHealingInfoAllOf HyperflexStPlatformClusterHealingInfoAllOf

// NewHyperflexStPlatformClusterHealingInfoAllOf instantiates a new HyperflexStPlatformClusterHealingInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexStPlatformClusterHealingInfoAllOf(classId string, objectType string) *HyperflexStPlatformClusterHealingInfoAllOf {
	this := HyperflexStPlatformClusterHealingInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexStPlatformClusterHealingInfoAllOfWithDefaults instantiates a new HyperflexStPlatformClusterHealingInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexStPlatformClusterHealingInfoAllOfWithDefaults() *HyperflexStPlatformClusterHealingInfoAllOf {
	this := HyperflexStPlatformClusterHealingInfoAllOf{}
	var classId string = "hyperflex.StPlatformClusterHealingInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.StPlatformClusterHealingInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEstimatedCompletionTimeInSeconds returns the EstimatedCompletionTimeInSeconds field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetEstimatedCompletionTimeInSeconds() int64 {
	if o == nil || o.EstimatedCompletionTimeInSeconds == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedCompletionTimeInSeconds
}

// GetEstimatedCompletionTimeInSecondsOk returns a tuple with the EstimatedCompletionTimeInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetEstimatedCompletionTimeInSecondsOk() (*int64, bool) {
	if o == nil || o.EstimatedCompletionTimeInSeconds == nil {
		return nil, false
	}
	return o.EstimatedCompletionTimeInSeconds, true
}

// HasEstimatedCompletionTimeInSeconds returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasEstimatedCompletionTimeInSeconds() bool {
	if o != nil && o.EstimatedCompletionTimeInSeconds != nil {
		return true
	}

	return false
}

// SetEstimatedCompletionTimeInSeconds gets a reference to the given int64 and assigns it to the EstimatedCompletionTimeInSeconds field.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetEstimatedCompletionTimeInSeconds(v int64) {
	o.EstimatedCompletionTimeInSeconds = &v
}

// GetInProgress returns the InProgress field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetInProgress() bool {
	if o == nil || o.InProgress == nil {
		var ret bool
		return ret
	}
	return *o.InProgress
}

// GetInProgressOk returns a tuple with the InProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetInProgressOk() (*bool, bool) {
	if o == nil || o.InProgress == nil {
		return nil, false
	}
	return o.InProgress, true
}

// HasInProgress returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasInProgress() bool {
	if o != nil && o.InProgress != nil {
		return true
	}

	return false
}

// SetInProgress gets a reference to the given bool and assigns it to the InProgress field.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetInProgress(v bool) {
	o.InProgress = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return &o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetMessages(v []string) {
	o.Messages = v
}

// GetMessagesIterator returns the MessagesIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessagesIterator() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MessagesIterator
}

// GetMessagesIteratorOk returns a tuple with the MessagesIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessagesIteratorOk() (*interface{}, bool) {
	if o == nil || o.MessagesIterator == nil {
		return nil, false
	}
	return &o.MessagesIterator, true
}

// HasMessagesIterator returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasMessagesIterator() bool {
	if o != nil && o.MessagesIterator != nil {
		return true
	}

	return false
}

// SetMessagesIterator gets a reference to the given interface{} and assigns it to the MessagesIterator field.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetMessagesIterator(v interface{}) {
	o.MessagesIterator = v
}

// GetMessagesSize returns the MessagesSize field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessagesSize() int64 {
	if o == nil || o.MessagesSize == nil {
		var ret int64
		return ret
	}
	return *o.MessagesSize
}

// GetMessagesSizeOk returns a tuple with the MessagesSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetMessagesSizeOk() (*int64, bool) {
	if o == nil || o.MessagesSize == nil {
		return nil, false
	}
	return o.MessagesSize, true
}

// HasMessagesSize returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasMessagesSize() bool {
	if o != nil && o.MessagesSize != nil {
		return true
	}

	return false
}

// SetMessagesSize gets a reference to the given int64 and assigns it to the MessagesSize field.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetMessagesSize(v int64) {
	o.MessagesSize = &v
}

// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetPercentComplete() int64 {
	if o == nil || o.PercentComplete == nil {
		var ret int64
		return ret
	}
	return *o.PercentComplete
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) GetPercentCompleteOk() (*int64, bool) {
	if o == nil || o.PercentComplete == nil {
		return nil, false
	}
	return o.PercentComplete, true
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) HasPercentComplete() bool {
	if o != nil && o.PercentComplete != nil {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given int64 and assigns it to the PercentComplete field.
func (o *HyperflexStPlatformClusterHealingInfoAllOf) SetPercentComplete(v int64) {
	o.PercentComplete = &v
}

func (o HyperflexStPlatformClusterHealingInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EstimatedCompletionTimeInSeconds != nil {
		toSerialize["EstimatedCompletionTimeInSeconds"] = o.EstimatedCompletionTimeInSeconds
	}
	if o.InProgress != nil {
		toSerialize["InProgress"] = o.InProgress
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.MessagesIterator != nil {
		toSerialize["MessagesIterator"] = o.MessagesIterator
	}
	if o.MessagesSize != nil {
		toSerialize["MessagesSize"] = o.MessagesSize
	}
	if o.PercentComplete != nil {
		toSerialize["PercentComplete"] = o.PercentComplete
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexStPlatformClusterHealingInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexStPlatformClusterHealingInfoAllOf := _HyperflexStPlatformClusterHealingInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexStPlatformClusterHealingInfoAllOf); err == nil {
		*o = HyperflexStPlatformClusterHealingInfoAllOf(varHyperflexStPlatformClusterHealingInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EstimatedCompletionTimeInSeconds")
		delete(additionalProperties, "InProgress")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "MessagesIterator")
		delete(additionalProperties, "MessagesSize")
		delete(additionalProperties, "PercentComplete")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexStPlatformClusterHealingInfoAllOf struct {
	value *HyperflexStPlatformClusterHealingInfoAllOf
	isSet bool
}

func (v NullableHyperflexStPlatformClusterHealingInfoAllOf) Get() *HyperflexStPlatformClusterHealingInfoAllOf {
	return v.value
}

func (v *NullableHyperflexStPlatformClusterHealingInfoAllOf) Set(val *HyperflexStPlatformClusterHealingInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexStPlatformClusterHealingInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexStPlatformClusterHealingInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexStPlatformClusterHealingInfoAllOf(val *HyperflexStPlatformClusterHealingInfoAllOf) *NullableHyperflexStPlatformClusterHealingInfoAllOf {
	return &NullableHyperflexStPlatformClusterHealingInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexStPlatformClusterHealingInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexStPlatformClusterHealingInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
