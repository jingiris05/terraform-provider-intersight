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

// NiatelemetryNiccAllOf Definition of the list of properties defined in 'niatelemetry.Nicc', excluding properties defined in parent classes.
type NiatelemetryNiccAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Configuration issues depicts the failures for NICC managed package upgrade on APIC.
	ConfigIssues *string `json:"ConfigIssues,omitempty"`
	// NICC state. NiccState checks the current operational state of NICC app on APIC.
	NiccState *string `json:"NiccState,omitempty"`
	// NICC state last updated timestamp. It indicates the last updated timestamp for operational state of NICC app.
	NiccStateLastUpdateTs *string `json:"NiccStateLastUpdateTs,omitempty"`
	// NICC version. NiccVersion is used to check compatibility with Nexus Cloud features.
	NiccVersion          *string                              `json:"NiccVersion,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNiccAllOf NiatelemetryNiccAllOf

// NewNiatelemetryNiccAllOf instantiates a new NiatelemetryNiccAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNiccAllOf(classId string, objectType string) *NiatelemetryNiccAllOf {
	this := NiatelemetryNiccAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNiccAllOfWithDefaults instantiates a new NiatelemetryNiccAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNiccAllOfWithDefaults() *NiatelemetryNiccAllOf {
	this := NiatelemetryNiccAllOf{}
	var classId string = "niatelemetry.Nicc"
	this.ClassId = classId
	var objectType string = "niatelemetry.Nicc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNiccAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiccAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNiccAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNiccAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiccAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNiccAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigIssues returns the ConfigIssues field value if set, zero value otherwise.
func (o *NiatelemetryNiccAllOf) GetConfigIssues() string {
	if o == nil || o.ConfigIssues == nil {
		var ret string
		return ret
	}
	return *o.ConfigIssues
}

// GetConfigIssuesOk returns a tuple with the ConfigIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiccAllOf) GetConfigIssuesOk() (*string, bool) {
	if o == nil || o.ConfigIssues == nil {
		return nil, false
	}
	return o.ConfigIssues, true
}

// HasConfigIssues returns a boolean if a field has been set.
func (o *NiatelemetryNiccAllOf) HasConfigIssues() bool {
	if o != nil && o.ConfigIssues != nil {
		return true
	}

	return false
}

// SetConfigIssues gets a reference to the given string and assigns it to the ConfigIssues field.
func (o *NiatelemetryNiccAllOf) SetConfigIssues(v string) {
	o.ConfigIssues = &v
}

// GetNiccState returns the NiccState field value if set, zero value otherwise.
func (o *NiatelemetryNiccAllOf) GetNiccState() string {
	if o == nil || o.NiccState == nil {
		var ret string
		return ret
	}
	return *o.NiccState
}

// GetNiccStateOk returns a tuple with the NiccState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiccAllOf) GetNiccStateOk() (*string, bool) {
	if o == nil || o.NiccState == nil {
		return nil, false
	}
	return o.NiccState, true
}

// HasNiccState returns a boolean if a field has been set.
func (o *NiatelemetryNiccAllOf) HasNiccState() bool {
	if o != nil && o.NiccState != nil {
		return true
	}

	return false
}

// SetNiccState gets a reference to the given string and assigns it to the NiccState field.
func (o *NiatelemetryNiccAllOf) SetNiccState(v string) {
	o.NiccState = &v
}

// GetNiccStateLastUpdateTs returns the NiccStateLastUpdateTs field value if set, zero value otherwise.
func (o *NiatelemetryNiccAllOf) GetNiccStateLastUpdateTs() string {
	if o == nil || o.NiccStateLastUpdateTs == nil {
		var ret string
		return ret
	}
	return *o.NiccStateLastUpdateTs
}

// GetNiccStateLastUpdateTsOk returns a tuple with the NiccStateLastUpdateTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiccAllOf) GetNiccStateLastUpdateTsOk() (*string, bool) {
	if o == nil || o.NiccStateLastUpdateTs == nil {
		return nil, false
	}
	return o.NiccStateLastUpdateTs, true
}

// HasNiccStateLastUpdateTs returns a boolean if a field has been set.
func (o *NiatelemetryNiccAllOf) HasNiccStateLastUpdateTs() bool {
	if o != nil && o.NiccStateLastUpdateTs != nil {
		return true
	}

	return false
}

// SetNiccStateLastUpdateTs gets a reference to the given string and assigns it to the NiccStateLastUpdateTs field.
func (o *NiatelemetryNiccAllOf) SetNiccStateLastUpdateTs(v string) {
	o.NiccStateLastUpdateTs = &v
}

// GetNiccVersion returns the NiccVersion field value if set, zero value otherwise.
func (o *NiatelemetryNiccAllOf) GetNiccVersion() string {
	if o == nil || o.NiccVersion == nil {
		var ret string
		return ret
	}
	return *o.NiccVersion
}

// GetNiccVersionOk returns a tuple with the NiccVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiccAllOf) GetNiccVersionOk() (*string, bool) {
	if o == nil || o.NiccVersion == nil {
		return nil, false
	}
	return o.NiccVersion, true
}

// HasNiccVersion returns a boolean if a field has been set.
func (o *NiatelemetryNiccAllOf) HasNiccVersion() bool {
	if o != nil && o.NiccVersion != nil {
		return true
	}

	return false
}

// SetNiccVersion gets a reference to the given string and assigns it to the NiccVersion field.
func (o *NiatelemetryNiccAllOf) SetNiccVersion(v string) {
	o.NiccVersion = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNiccAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiccAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNiccAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNiccAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNiccAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigIssues != nil {
		toSerialize["ConfigIssues"] = o.ConfigIssues
	}
	if o.NiccState != nil {
		toSerialize["NiccState"] = o.NiccState
	}
	if o.NiccStateLastUpdateTs != nil {
		toSerialize["NiccStateLastUpdateTs"] = o.NiccStateLastUpdateTs
	}
	if o.NiccVersion != nil {
		toSerialize["NiccVersion"] = o.NiccVersion
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNiccAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNiccAllOf := _NiatelemetryNiccAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNiccAllOf); err == nil {
		*o = NiatelemetryNiccAllOf(varNiatelemetryNiccAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigIssues")
		delete(additionalProperties, "NiccState")
		delete(additionalProperties, "NiccStateLastUpdateTs")
		delete(additionalProperties, "NiccVersion")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNiccAllOf struct {
	value *NiatelemetryNiccAllOf
	isSet bool
}

func (v NullableNiatelemetryNiccAllOf) Get() *NiatelemetryNiccAllOf {
	return v.value
}

func (v *NullableNiatelemetryNiccAllOf) Set(val *NiatelemetryNiccAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiccAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiccAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiccAllOf(val *NiatelemetryNiccAllOf) *NullableNiatelemetryNiccAllOf {
	return &NullableNiatelemetryNiccAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNiccAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiccAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
