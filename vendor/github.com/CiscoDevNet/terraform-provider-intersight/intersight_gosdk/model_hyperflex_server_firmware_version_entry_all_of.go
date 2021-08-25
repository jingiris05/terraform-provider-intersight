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

// HyperflexServerFirmwareVersionEntryAllOf Definition of the list of properties defined in 'hyperflex.ServerFirmwareVersionEntry', excluding properties defined in parent classes.
type HyperflexServerFirmwareVersionEntryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                                `json:"ObjectType"`
	Constraint NullableHyperflexAppSettingConstraint `json:"Constraint,omitempty"`
	// The server platform type that is applicable for the server firmware bundle version. * `M5` - M5 generation of UCS server. * `M3` - M3 generation of UCS server. * `M4` - M4 generation of UCS server.
	ServerPlatform *string `json:"ServerPlatform,omitempty"`
	// The server firmware bundle version.
	Version               *string                                     `json:"Version,omitempty"`
	ServerFirmwareVersion *HyperflexServerFirmwareVersionRelationship `json:"ServerFirmwareVersion,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _HyperflexServerFirmwareVersionEntryAllOf HyperflexServerFirmwareVersionEntryAllOf

// NewHyperflexServerFirmwareVersionEntryAllOf instantiates a new HyperflexServerFirmwareVersionEntryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexServerFirmwareVersionEntryAllOf(classId string, objectType string) *HyperflexServerFirmwareVersionEntryAllOf {
	this := HyperflexServerFirmwareVersionEntryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var serverPlatform string = "M5"
	this.ServerPlatform = &serverPlatform
	return &this
}

// NewHyperflexServerFirmwareVersionEntryAllOfWithDefaults instantiates a new HyperflexServerFirmwareVersionEntryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexServerFirmwareVersionEntryAllOfWithDefaults() *HyperflexServerFirmwareVersionEntryAllOf {
	this := HyperflexServerFirmwareVersionEntryAllOf{}
	var classId string = "hyperflex.ServerFirmwareVersionEntry"
	this.ClassId = classId
	var objectType string = "hyperflex.ServerFirmwareVersionEntry"
	this.ObjectType = objectType
	var serverPlatform string = "M5"
	this.ServerPlatform = &serverPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexServerFirmwareVersionEntryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexServerFirmwareVersionEntryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConstraint returns the Constraint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetConstraint() HyperflexAppSettingConstraint {
	if o == nil || o.Constraint.Get() == nil {
		var ret HyperflexAppSettingConstraint
		return ret
	}
	return *o.Constraint.Get()
}

// GetConstraintOk returns a tuple with the Constraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetConstraintOk() (*HyperflexAppSettingConstraint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraint.Get(), o.Constraint.IsSet()
}

// HasConstraint returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionEntryAllOf) HasConstraint() bool {
	if o != nil && o.Constraint.IsSet() {
		return true
	}

	return false
}

// SetConstraint gets a reference to the given NullableHyperflexAppSettingConstraint and assigns it to the Constraint field.
func (o *HyperflexServerFirmwareVersionEntryAllOf) SetConstraint(v HyperflexAppSettingConstraint) {
	o.Constraint.Set(&v)
}

// SetConstraintNil sets the value for Constraint to be an explicit nil
func (o *HyperflexServerFirmwareVersionEntryAllOf) SetConstraintNil() {
	o.Constraint.Set(nil)
}

// UnsetConstraint ensures that no value is present for Constraint, not even an explicit nil
func (o *HyperflexServerFirmwareVersionEntryAllOf) UnsetConstraint() {
	o.Constraint.Unset()
}

// GetServerPlatform returns the ServerPlatform field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetServerPlatform() string {
	if o == nil || o.ServerPlatform == nil {
		var ret string
		return ret
	}
	return *o.ServerPlatform
}

// GetServerPlatformOk returns a tuple with the ServerPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetServerPlatformOk() (*string, bool) {
	if o == nil || o.ServerPlatform == nil {
		return nil, false
	}
	return o.ServerPlatform, true
}

// HasServerPlatform returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionEntryAllOf) HasServerPlatform() bool {
	if o != nil && o.ServerPlatform != nil {
		return true
	}

	return false
}

// SetServerPlatform gets a reference to the given string and assigns it to the ServerPlatform field.
func (o *HyperflexServerFirmwareVersionEntryAllOf) SetServerPlatform(v string) {
	o.ServerPlatform = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionEntryAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexServerFirmwareVersionEntryAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetServerFirmwareVersion returns the ServerFirmwareVersion field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetServerFirmwareVersion() HyperflexServerFirmwareVersionRelationship {
	if o == nil || o.ServerFirmwareVersion == nil {
		var ret HyperflexServerFirmwareVersionRelationship
		return ret
	}
	return *o.ServerFirmwareVersion
}

// GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionEntryAllOf) GetServerFirmwareVersionOk() (*HyperflexServerFirmwareVersionRelationship, bool) {
	if o == nil || o.ServerFirmwareVersion == nil {
		return nil, false
	}
	return o.ServerFirmwareVersion, true
}

// HasServerFirmwareVersion returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionEntryAllOf) HasServerFirmwareVersion() bool {
	if o != nil && o.ServerFirmwareVersion != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersion gets a reference to the given HyperflexServerFirmwareVersionRelationship and assigns it to the ServerFirmwareVersion field.
func (o *HyperflexServerFirmwareVersionEntryAllOf) SetServerFirmwareVersion(v HyperflexServerFirmwareVersionRelationship) {
	o.ServerFirmwareVersion = &v
}

func (o HyperflexServerFirmwareVersionEntryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Constraint.IsSet() {
		toSerialize["Constraint"] = o.Constraint.Get()
	}
	if o.ServerPlatform != nil {
		toSerialize["ServerPlatform"] = o.ServerPlatform
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.ServerFirmwareVersion != nil {
		toSerialize["ServerFirmwareVersion"] = o.ServerFirmwareVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexServerFirmwareVersionEntryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexServerFirmwareVersionEntryAllOf := _HyperflexServerFirmwareVersionEntryAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexServerFirmwareVersionEntryAllOf); err == nil {
		*o = HyperflexServerFirmwareVersionEntryAllOf(varHyperflexServerFirmwareVersionEntryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Constraint")
		delete(additionalProperties, "ServerPlatform")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "ServerFirmwareVersion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexServerFirmwareVersionEntryAllOf struct {
	value *HyperflexServerFirmwareVersionEntryAllOf
	isSet bool
}

func (v NullableHyperflexServerFirmwareVersionEntryAllOf) Get() *HyperflexServerFirmwareVersionEntryAllOf {
	return v.value
}

func (v *NullableHyperflexServerFirmwareVersionEntryAllOf) Set(val *HyperflexServerFirmwareVersionEntryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerFirmwareVersionEntryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerFirmwareVersionEntryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerFirmwareVersionEntryAllOf(val *HyperflexServerFirmwareVersionEntryAllOf) *NullableHyperflexServerFirmwareVersionEntryAllOf {
	return &NullableHyperflexServerFirmwareVersionEntryAllOf{value: val, isSet: true}
}

func (v NullableHyperflexServerFirmwareVersionEntryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerFirmwareVersionEntryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
