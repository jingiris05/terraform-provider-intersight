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

// RecommendationHardwareExpansionRequestAllOf Definition of the list of properties defined in 'recommendation.HardwareExpansionRequest', excluding properties defined in parent classes.
type RecommendationHardwareExpansionRequestAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action to be triggered for computing recommendation. Default value is None. * `None` - The Enum value None represents that no action is triggered on the forecast Instance managed object. * `Evaluate` - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered.
	Action *string `json:"Action,omitempty"`
	// User visible error message for the Hardware Expansion request.
	Message *string `json:"Message,omitempty"`
	// Status of the Hardware Expansion request. * `None` - The Enum value None represents the default status value before any API call is made. * `Success` - The Enum value Success represents that the API call returned with success. * `Fail` - The Enum value Fail represents that the API call returned with a failure.
	Status *string `json:"Status,omitempty"`
	// An array of relationships to recommendationHardwareExpansionRequestItem resources.
	HwExpansionRequestItems []RecommendationHardwareExpansionRequestItemRelationship `json:"HwExpansionRequestItems,omitempty"`
	RegisteredDevice        *AssetDeviceRegistrationRelationship                     `json:"RegisteredDevice,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _RecommendationHardwareExpansionRequestAllOf RecommendationHardwareExpansionRequestAllOf

// NewRecommendationHardwareExpansionRequestAllOf instantiates a new RecommendationHardwareExpansionRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationHardwareExpansionRequestAllOf(classId string, objectType string) *RecommendationHardwareExpansionRequestAllOf {
	this := RecommendationHardwareExpansionRequestAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// NewRecommendationHardwareExpansionRequestAllOfWithDefaults instantiates a new RecommendationHardwareExpansionRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationHardwareExpansionRequestAllOfWithDefaults() *RecommendationHardwareExpansionRequestAllOf {
	this := RecommendationHardwareExpansionRequestAllOf{}
	var classId string = "recommendation.HardwareExpansionRequest"
	this.ClassId = classId
	var objectType string = "recommendation.HardwareExpansionRequest"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationHardwareExpansionRequestAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationHardwareExpansionRequestAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationHardwareExpansionRequestAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationHardwareExpansionRequestAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RecommendationHardwareExpansionRequestAllOf) SetAction(v string) {
	o.Action = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RecommendationHardwareExpansionRequestAllOf) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RecommendationHardwareExpansionRequestAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetHwExpansionRequestItems returns the HwExpansionRequestItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationHardwareExpansionRequestAllOf) GetHwExpansionRequestItems() []RecommendationHardwareExpansionRequestItemRelationship {
	if o == nil {
		var ret []RecommendationHardwareExpansionRequestItemRelationship
		return ret
	}
	return o.HwExpansionRequestItems
}

// GetHwExpansionRequestItemsOk returns a tuple with the HwExpansionRequestItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationHardwareExpansionRequestAllOf) GetHwExpansionRequestItemsOk() ([]RecommendationHardwareExpansionRequestItemRelationship, bool) {
	if o == nil || o.HwExpansionRequestItems == nil {
		return nil, false
	}
	return o.HwExpansionRequestItems, true
}

// HasHwExpansionRequestItems returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) HasHwExpansionRequestItems() bool {
	if o != nil && o.HwExpansionRequestItems != nil {
		return true
	}

	return false
}

// SetHwExpansionRequestItems gets a reference to the given []RecommendationHardwareExpansionRequestItemRelationship and assigns it to the HwExpansionRequestItems field.
func (o *RecommendationHardwareExpansionRequestAllOf) SetHwExpansionRequestItems(v []RecommendationHardwareExpansionRequestItemRelationship) {
	o.HwExpansionRequestItems = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *RecommendationHardwareExpansionRequestAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o RecommendationHardwareExpansionRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.HwExpansionRequestItems != nil {
		toSerialize["HwExpansionRequestItems"] = o.HwExpansionRequestItems
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationHardwareExpansionRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecommendationHardwareExpansionRequestAllOf := _RecommendationHardwareExpansionRequestAllOf{}

	if err = json.Unmarshal(bytes, &varRecommendationHardwareExpansionRequestAllOf); err == nil {
		*o = RecommendationHardwareExpansionRequestAllOf(varRecommendationHardwareExpansionRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "HwExpansionRequestItems")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommendationHardwareExpansionRequestAllOf struct {
	value *RecommendationHardwareExpansionRequestAllOf
	isSet bool
}

func (v NullableRecommendationHardwareExpansionRequestAllOf) Get() *RecommendationHardwareExpansionRequestAllOf {
	return v.value
}

func (v *NullableRecommendationHardwareExpansionRequestAllOf) Set(val *RecommendationHardwareExpansionRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationHardwareExpansionRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationHardwareExpansionRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationHardwareExpansionRequestAllOf(val *RecommendationHardwareExpansionRequestAllOf) *NullableRecommendationHardwareExpansionRequestAllOf {
	return &NullableRecommendationHardwareExpansionRequestAllOf{value: val, isSet: true}
}

func (v NullableRecommendationHardwareExpansionRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationHardwareExpansionRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
