/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-24T21:18:00Z.
 *
 * API version: 1.0.9-4404
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// AssetDeviceTransactionAllOf Definition of the list of properties defined in 'asset.DeviceTransaction', excluding properties defined in parent classes.
type AssetDeviceTransactionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The action taken by Cisco Install Base on the device.
	Action *string `json:"Action,omitempty"`
	// Description of status of Cisco device reported by Cisco Install Base.
	StatusDescription *string `json:"StatusDescription,omitempty"`
	// Status ID of device reported by Cisco Install Base. * `0` - A default value to catch cases where device status is not correctly detected. * `10000` - Device is installed successfully. * `1010041` - Device is currently in Return Material Authorization process. * `10005` - Device is replaced successfully with another device. * `10007` - Device is returned succcessfuly. * `10009` - Device is terminated at customer's end.
	StatusId *int32 `json:"StatusId,omitempty"`
	// Timestamp field reported by Cisco Install Base.
	Timestamp *string `json:"Timestamp,omitempty"`
	// Transaction batch ID reported by Cisco Install Base.
	TransactionBatchId *int64 `json:"TransactionBatchId,omitempty"`
	// Transaction Date reported by Cisco Install Base.
	TransactionDate *string `json:"TransactionDate,omitempty"`
	// Transaction sequence reported by Cisco Install Base.
	TransactionSequence  *int64 `json:"TransactionSequence,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDeviceTransactionAllOf AssetDeviceTransactionAllOf

// NewAssetDeviceTransactionAllOf instantiates a new AssetDeviceTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceTransactionAllOf(classId string, objectType string) *AssetDeviceTransactionAllOf {
	this := AssetDeviceTransactionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var statusId int32 = 0
	this.StatusId = &statusId
	return &this
}

// NewAssetDeviceTransactionAllOfWithDefaults instantiates a new AssetDeviceTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceTransactionAllOfWithDefaults() *AssetDeviceTransactionAllOf {
	this := AssetDeviceTransactionAllOf{}
	var classId string = "asset.DeviceTransaction"
	this.ClassId = classId
	var objectType string = "asset.DeviceTransaction"
	this.ObjectType = objectType
	var statusId int32 = 0
	this.StatusId = &statusId
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceTransactionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceTransactionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceTransactionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceTransactionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceTransactionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceTransactionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AssetDeviceTransactionAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceTransactionAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AssetDeviceTransactionAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AssetDeviceTransactionAllOf) SetAction(v string) {
	o.Action = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *AssetDeviceTransactionAllOf) GetStatusDescription() string {
	if o == nil || o.StatusDescription == nil {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceTransactionAllOf) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || o.StatusDescription == nil {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *AssetDeviceTransactionAllOf) HasStatusDescription() bool {
	if o != nil && o.StatusDescription != nil {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *AssetDeviceTransactionAllOf) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *AssetDeviceTransactionAllOf) GetStatusId() int32 {
	if o == nil || o.StatusId == nil {
		var ret int32
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceTransactionAllOf) GetStatusIdOk() (*int32, bool) {
	if o == nil || o.StatusId == nil {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *AssetDeviceTransactionAllOf) HasStatusId() bool {
	if o != nil && o.StatusId != nil {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given int32 and assigns it to the StatusId field.
func (o *AssetDeviceTransactionAllOf) SetStatusId(v int32) {
	o.StatusId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AssetDeviceTransactionAllOf) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceTransactionAllOf) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AssetDeviceTransactionAllOf) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *AssetDeviceTransactionAllOf) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetTransactionBatchId returns the TransactionBatchId field value if set, zero value otherwise.
func (o *AssetDeviceTransactionAllOf) GetTransactionBatchId() int64 {
	if o == nil || o.TransactionBatchId == nil {
		var ret int64
		return ret
	}
	return *o.TransactionBatchId
}

// GetTransactionBatchIdOk returns a tuple with the TransactionBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceTransactionAllOf) GetTransactionBatchIdOk() (*int64, bool) {
	if o == nil || o.TransactionBatchId == nil {
		return nil, false
	}
	return o.TransactionBatchId, true
}

// HasTransactionBatchId returns a boolean if a field has been set.
func (o *AssetDeviceTransactionAllOf) HasTransactionBatchId() bool {
	if o != nil && o.TransactionBatchId != nil {
		return true
	}

	return false
}

// SetTransactionBatchId gets a reference to the given int64 and assigns it to the TransactionBatchId field.
func (o *AssetDeviceTransactionAllOf) SetTransactionBatchId(v int64) {
	o.TransactionBatchId = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *AssetDeviceTransactionAllOf) GetTransactionDate() string {
	if o == nil || o.TransactionDate == nil {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceTransactionAllOf) GetTransactionDateOk() (*string, bool) {
	if o == nil || o.TransactionDate == nil {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *AssetDeviceTransactionAllOf) HasTransactionDate() bool {
	if o != nil && o.TransactionDate != nil {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *AssetDeviceTransactionAllOf) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetTransactionSequence returns the TransactionSequence field value if set, zero value otherwise.
func (o *AssetDeviceTransactionAllOf) GetTransactionSequence() int64 {
	if o == nil || o.TransactionSequence == nil {
		var ret int64
		return ret
	}
	return *o.TransactionSequence
}

// GetTransactionSequenceOk returns a tuple with the TransactionSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceTransactionAllOf) GetTransactionSequenceOk() (*int64, bool) {
	if o == nil || o.TransactionSequence == nil {
		return nil, false
	}
	return o.TransactionSequence, true
}

// HasTransactionSequence returns a boolean if a field has been set.
func (o *AssetDeviceTransactionAllOf) HasTransactionSequence() bool {
	if o != nil && o.TransactionSequence != nil {
		return true
	}

	return false
}

// SetTransactionSequence gets a reference to the given int64 and assigns it to the TransactionSequence field.
func (o *AssetDeviceTransactionAllOf) SetTransactionSequence(v int64) {
	o.TransactionSequence = &v
}

func (o AssetDeviceTransactionAllOf) MarshalJSON() ([]byte, error) {
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
	if o.StatusDescription != nil {
		toSerialize["StatusDescription"] = o.StatusDescription
	}
	if o.StatusId != nil {
		toSerialize["StatusId"] = o.StatusId
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.TransactionBatchId != nil {
		toSerialize["TransactionBatchId"] = o.TransactionBatchId
	}
	if o.TransactionDate != nil {
		toSerialize["TransactionDate"] = o.TransactionDate
	}
	if o.TransactionSequence != nil {
		toSerialize["TransactionSequence"] = o.TransactionSequence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetDeviceTransactionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetDeviceTransactionAllOf := _AssetDeviceTransactionAllOf{}

	if err = json.Unmarshal(bytes, &varAssetDeviceTransactionAllOf); err == nil {
		*o = AssetDeviceTransactionAllOf(varAssetDeviceTransactionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "StatusDescription")
		delete(additionalProperties, "StatusId")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "TransactionBatchId")
		delete(additionalProperties, "TransactionDate")
		delete(additionalProperties, "TransactionSequence")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDeviceTransactionAllOf struct {
	value *AssetDeviceTransactionAllOf
	isSet bool
}

func (v NullableAssetDeviceTransactionAllOf) Get() *AssetDeviceTransactionAllOf {
	return v.value
}

func (v *NullableAssetDeviceTransactionAllOf) Set(val *AssetDeviceTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceTransactionAllOf(val *AssetDeviceTransactionAllOf) *NullableAssetDeviceTransactionAllOf {
	return &NullableAssetDeviceTransactionAllOf{value: val, isSet: true}
}

func (v NullableAssetDeviceTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
