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

// ConnectorStreamAcknowledgeAllOf Definition of the list of properties defined in 'connector.StreamAcknowledge', excluding properties defined in parent classes.
type ConnectorStreamAcknowledgeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The latest message sequence processed in the cloud. Device connector will drop all messages up to this sequence from its cache.
	AckSequence          *int64 `json:"AckSequence,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorStreamAcknowledgeAllOf ConnectorStreamAcknowledgeAllOf

// NewConnectorStreamAcknowledgeAllOf instantiates a new ConnectorStreamAcknowledgeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorStreamAcknowledgeAllOf(classId string, objectType string) *ConnectorStreamAcknowledgeAllOf {
	this := ConnectorStreamAcknowledgeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorStreamAcknowledgeAllOfWithDefaults instantiates a new ConnectorStreamAcknowledgeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorStreamAcknowledgeAllOfWithDefaults() *ConnectorStreamAcknowledgeAllOf {
	this := ConnectorStreamAcknowledgeAllOf{}
	var classId string = "connector.StreamAcknowledge"
	this.ClassId = classId
	var objectType string = "connector.StreamAcknowledge"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorStreamAcknowledgeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorStreamAcknowledgeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorStreamAcknowledgeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorStreamAcknowledgeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorStreamAcknowledgeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorStreamAcknowledgeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAckSequence returns the AckSequence field value if set, zero value otherwise.
func (o *ConnectorStreamAcknowledgeAllOf) GetAckSequence() int64 {
	if o == nil || o.AckSequence == nil {
		var ret int64
		return ret
	}
	return *o.AckSequence
}

// GetAckSequenceOk returns a tuple with the AckSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStreamAcknowledgeAllOf) GetAckSequenceOk() (*int64, bool) {
	if o == nil || o.AckSequence == nil {
		return nil, false
	}
	return o.AckSequence, true
}

// HasAckSequence returns a boolean if a field has been set.
func (o *ConnectorStreamAcknowledgeAllOf) HasAckSequence() bool {
	if o != nil && o.AckSequence != nil {
		return true
	}

	return false
}

// SetAckSequence gets a reference to the given int64 and assigns it to the AckSequence field.
func (o *ConnectorStreamAcknowledgeAllOf) SetAckSequence(v int64) {
	o.AckSequence = &v
}

func (o ConnectorStreamAcknowledgeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AckSequence != nil {
		toSerialize["AckSequence"] = o.AckSequence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorStreamAcknowledgeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorStreamAcknowledgeAllOf := _ConnectorStreamAcknowledgeAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorStreamAcknowledgeAllOf); err == nil {
		*o = ConnectorStreamAcknowledgeAllOf(varConnectorStreamAcknowledgeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AckSequence")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorStreamAcknowledgeAllOf struct {
	value *ConnectorStreamAcknowledgeAllOf
	isSet bool
}

func (v NullableConnectorStreamAcknowledgeAllOf) Get() *ConnectorStreamAcknowledgeAllOf {
	return v.value
}

func (v *NullableConnectorStreamAcknowledgeAllOf) Set(val *ConnectorStreamAcknowledgeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorStreamAcknowledgeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorStreamAcknowledgeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorStreamAcknowledgeAllOf(val *ConnectorStreamAcknowledgeAllOf) *NullableConnectorStreamAcknowledgeAllOf {
	return &NullableConnectorStreamAcknowledgeAllOf{value: val, isSet: true}
}

func (v NullableConnectorStreamAcknowledgeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorStreamAcknowledgeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
