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
	"reflect"
	"strings"
)

// ConnectorStartStreamFromDevice Start a stream from the device connector. Message type sent from the device connector to Intersight to trigger a stream open.
type ConnectorStartStreamFromDevice struct {
	ConnectorStreamMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The asset.ClusterMember member identity that is opening this stream.
	MemberId *string `json:"MemberId,omitempty"`
	// The stream is to be started against the cluster member.
	MemberStream *bool `json:"MemberStream,omitempty"`
	// Any extra configuration needed to open/identify a stream.
	StreamConfig interface{} `json:"StreamConfig,omitempty"`
	// Identifies the type of stream to open to the device. The Intersight service will validate that the device should open a stream of this type and if so build a stream configuration and send it down to the device. The streamType should identify a unique stream to open to a device, that is if the device sends a stream open message and a stream of that type is already open in the cloud the existing stream should be re-used.
	StreamType *string `json:"StreamType,omitempty"`
	// The topic the device should send the stream open message to.
	Topic                *string `json:"Topic,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorStartStreamFromDevice ConnectorStartStreamFromDevice

// NewConnectorStartStreamFromDevice instantiates a new ConnectorStartStreamFromDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorStartStreamFromDevice(classId string, objectType string) *ConnectorStartStreamFromDevice {
	this := ConnectorStartStreamFromDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorStartStreamFromDeviceWithDefaults instantiates a new ConnectorStartStreamFromDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorStartStreamFromDeviceWithDefaults() *ConnectorStartStreamFromDevice {
	this := ConnectorStartStreamFromDevice{}
	var classId string = "connector.StartStreamFromDevice"
	this.ClassId = classId
	var objectType string = "connector.StartStreamFromDevice"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorStartStreamFromDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamFromDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorStartStreamFromDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorStartStreamFromDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamFromDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorStartStreamFromDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *ConnectorStartStreamFromDevice) GetMemberId() string {
	if o == nil || o.MemberId == nil {
		var ret string
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamFromDevice) GetMemberIdOk() (*string, bool) {
	if o == nil || o.MemberId == nil {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *ConnectorStartStreamFromDevice) HasMemberId() bool {
	if o != nil && o.MemberId != nil {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given string and assigns it to the MemberId field.
func (o *ConnectorStartStreamFromDevice) SetMemberId(v string) {
	o.MemberId = &v
}

// GetMemberStream returns the MemberStream field value if set, zero value otherwise.
func (o *ConnectorStartStreamFromDevice) GetMemberStream() bool {
	if o == nil || o.MemberStream == nil {
		var ret bool
		return ret
	}
	return *o.MemberStream
}

// GetMemberStreamOk returns a tuple with the MemberStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamFromDevice) GetMemberStreamOk() (*bool, bool) {
	if o == nil || o.MemberStream == nil {
		return nil, false
	}
	return o.MemberStream, true
}

// HasMemberStream returns a boolean if a field has been set.
func (o *ConnectorStartStreamFromDevice) HasMemberStream() bool {
	if o != nil && o.MemberStream != nil {
		return true
	}

	return false
}

// SetMemberStream gets a reference to the given bool and assigns it to the MemberStream field.
func (o *ConnectorStartStreamFromDevice) SetMemberStream(v bool) {
	o.MemberStream = &v
}

// GetStreamConfig returns the StreamConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorStartStreamFromDevice) GetStreamConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StreamConfig
}

// GetStreamConfigOk returns a tuple with the StreamConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorStartStreamFromDevice) GetStreamConfigOk() (*interface{}, bool) {
	if o == nil || o.StreamConfig == nil {
		return nil, false
	}
	return &o.StreamConfig, true
}

// HasStreamConfig returns a boolean if a field has been set.
func (o *ConnectorStartStreamFromDevice) HasStreamConfig() bool {
	if o != nil && o.StreamConfig != nil {
		return true
	}

	return false
}

// SetStreamConfig gets a reference to the given interface{} and assigns it to the StreamConfig field.
func (o *ConnectorStartStreamFromDevice) SetStreamConfig(v interface{}) {
	o.StreamConfig = v
}

// GetStreamType returns the StreamType field value if set, zero value otherwise.
func (o *ConnectorStartStreamFromDevice) GetStreamType() string {
	if o == nil || o.StreamType == nil {
		var ret string
		return ret
	}
	return *o.StreamType
}

// GetStreamTypeOk returns a tuple with the StreamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamFromDevice) GetStreamTypeOk() (*string, bool) {
	if o == nil || o.StreamType == nil {
		return nil, false
	}
	return o.StreamType, true
}

// HasStreamType returns a boolean if a field has been set.
func (o *ConnectorStartStreamFromDevice) HasStreamType() bool {
	if o != nil && o.StreamType != nil {
		return true
	}

	return false
}

// SetStreamType gets a reference to the given string and assigns it to the StreamType field.
func (o *ConnectorStartStreamFromDevice) SetStreamType(v string) {
	o.StreamType = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *ConnectorStartStreamFromDevice) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStartStreamFromDevice) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *ConnectorStartStreamFromDevice) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *ConnectorStartStreamFromDevice) SetTopic(v string) {
	o.Topic = &v
}

func (o ConnectorStartStreamFromDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorStreamMessage, errConnectorStreamMessage := json.Marshal(o.ConnectorStreamMessage)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	errConnectorStreamMessage = json.Unmarshal([]byte(serializedConnectorStreamMessage), &toSerialize)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MemberId != nil {
		toSerialize["MemberId"] = o.MemberId
	}
	if o.MemberStream != nil {
		toSerialize["MemberStream"] = o.MemberStream
	}
	if o.StreamConfig != nil {
		toSerialize["StreamConfig"] = o.StreamConfig
	}
	if o.StreamType != nil {
		toSerialize["StreamType"] = o.StreamType
	}
	if o.Topic != nil {
		toSerialize["Topic"] = o.Topic
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorStartStreamFromDevice) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorStartStreamFromDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The asset.ClusterMember member identity that is opening this stream.
		MemberId *string `json:"MemberId,omitempty"`
		// The stream is to be started against the cluster member.
		MemberStream *bool `json:"MemberStream,omitempty"`
		// Any extra configuration needed to open/identify a stream.
		StreamConfig interface{} `json:"StreamConfig,omitempty"`
		// Identifies the type of stream to open to the device. The Intersight service will validate that the device should open a stream of this type and if so build a stream configuration and send it down to the device. The streamType should identify a unique stream to open to a device, that is if the device sends a stream open message and a stream of that type is already open in the cloud the existing stream should be re-used.
		StreamType *string `json:"StreamType,omitempty"`
		// The topic the device should send the stream open message to.
		Topic *string `json:"Topic,omitempty"`
	}

	varConnectorStartStreamFromDeviceWithoutEmbeddedStruct := ConnectorStartStreamFromDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorStartStreamFromDeviceWithoutEmbeddedStruct)
	if err == nil {
		varConnectorStartStreamFromDevice := _ConnectorStartStreamFromDevice{}
		varConnectorStartStreamFromDevice.ClassId = varConnectorStartStreamFromDeviceWithoutEmbeddedStruct.ClassId
		varConnectorStartStreamFromDevice.ObjectType = varConnectorStartStreamFromDeviceWithoutEmbeddedStruct.ObjectType
		varConnectorStartStreamFromDevice.MemberId = varConnectorStartStreamFromDeviceWithoutEmbeddedStruct.MemberId
		varConnectorStartStreamFromDevice.MemberStream = varConnectorStartStreamFromDeviceWithoutEmbeddedStruct.MemberStream
		varConnectorStartStreamFromDevice.StreamConfig = varConnectorStartStreamFromDeviceWithoutEmbeddedStruct.StreamConfig
		varConnectorStartStreamFromDevice.StreamType = varConnectorStartStreamFromDeviceWithoutEmbeddedStruct.StreamType
		varConnectorStartStreamFromDevice.Topic = varConnectorStartStreamFromDeviceWithoutEmbeddedStruct.Topic
		*o = ConnectorStartStreamFromDevice(varConnectorStartStreamFromDevice)
	} else {
		return err
	}

	varConnectorStartStreamFromDevice := _ConnectorStartStreamFromDevice{}

	err = json.Unmarshal(bytes, &varConnectorStartStreamFromDevice)
	if err == nil {
		o.ConnectorStreamMessage = varConnectorStartStreamFromDevice.ConnectorStreamMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MemberId")
		delete(additionalProperties, "MemberStream")
		delete(additionalProperties, "StreamConfig")
		delete(additionalProperties, "StreamType")
		delete(additionalProperties, "Topic")

		// remove fields from embedded structs
		reflectConnectorStreamMessage := reflect.ValueOf(o.ConnectorStreamMessage)
		for i := 0; i < reflectConnectorStreamMessage.Type().NumField(); i++ {
			t := reflectConnectorStreamMessage.Type().Field(i)

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

type NullableConnectorStartStreamFromDevice struct {
	value *ConnectorStartStreamFromDevice
	isSet bool
}

func (v NullableConnectorStartStreamFromDevice) Get() *ConnectorStartStreamFromDevice {
	return v.value
}

func (v *NullableConnectorStartStreamFromDevice) Set(val *ConnectorStartStreamFromDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorStartStreamFromDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorStartStreamFromDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorStartStreamFromDevice(val *ConnectorStartStreamFromDevice) *NullableConnectorStartStreamFromDevice {
	return &NullableConnectorStartStreamFromDevice{value: val, isSet: true}
}

func (v NullableConnectorStartStreamFromDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorStartStreamFromDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
