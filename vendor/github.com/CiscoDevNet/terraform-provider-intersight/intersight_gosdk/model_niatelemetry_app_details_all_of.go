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

// NiatelemetryAppDetailsAllOf Definition of the list of properties defined in 'niatelemetry.AppDetails', excluding properties defined in parent classes.
type NiatelemetryAppDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Names of apps running on ND.
	AppName *string `json:"AppName,omitempty"`
	// Status of apps running on ND.
	AppStatus *string `json:"AppStatus,omitempty"`
	// Versions of apps running on ND.
	AppVersion           *string                              `json:"AppVersion,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryAppDetailsAllOf NiatelemetryAppDetailsAllOf

// NewNiatelemetryAppDetailsAllOf instantiates a new NiatelemetryAppDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryAppDetailsAllOf(classId string, objectType string) *NiatelemetryAppDetailsAllOf {
	this := NiatelemetryAppDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryAppDetailsAllOfWithDefaults instantiates a new NiatelemetryAppDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryAppDetailsAllOfWithDefaults() *NiatelemetryAppDetailsAllOf {
	this := NiatelemetryAppDetailsAllOf{}
	var classId string = "niatelemetry.AppDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.AppDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryAppDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryAppDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryAppDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryAppDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *NiatelemetryAppDetailsAllOf) GetAppName() string {
	if o == nil || o.AppName == nil {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetailsAllOf) GetAppNameOk() (*string, bool) {
	if o == nil || o.AppName == nil {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *NiatelemetryAppDetailsAllOf) HasAppName() bool {
	if o != nil && o.AppName != nil {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *NiatelemetryAppDetailsAllOf) SetAppName(v string) {
	o.AppName = &v
}

// GetAppStatus returns the AppStatus field value if set, zero value otherwise.
func (o *NiatelemetryAppDetailsAllOf) GetAppStatus() string {
	if o == nil || o.AppStatus == nil {
		var ret string
		return ret
	}
	return *o.AppStatus
}

// GetAppStatusOk returns a tuple with the AppStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetailsAllOf) GetAppStatusOk() (*string, bool) {
	if o == nil || o.AppStatus == nil {
		return nil, false
	}
	return o.AppStatus, true
}

// HasAppStatus returns a boolean if a field has been set.
func (o *NiatelemetryAppDetailsAllOf) HasAppStatus() bool {
	if o != nil && o.AppStatus != nil {
		return true
	}

	return false
}

// SetAppStatus gets a reference to the given string and assigns it to the AppStatus field.
func (o *NiatelemetryAppDetailsAllOf) SetAppStatus(v string) {
	o.AppStatus = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *NiatelemetryAppDetailsAllOf) GetAppVersion() string {
	if o == nil || o.AppVersion == nil {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetailsAllOf) GetAppVersionOk() (*string, bool) {
	if o == nil || o.AppVersion == nil {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *NiatelemetryAppDetailsAllOf) HasAppVersion() bool {
	if o != nil && o.AppVersion != nil {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *NiatelemetryAppDetailsAllOf) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryAppDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAppDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryAppDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryAppDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryAppDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AppName != nil {
		toSerialize["AppName"] = o.AppName
	}
	if o.AppStatus != nil {
		toSerialize["AppStatus"] = o.AppStatus
	}
	if o.AppVersion != nil {
		toSerialize["AppVersion"] = o.AppVersion
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryAppDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryAppDetailsAllOf := _NiatelemetryAppDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryAppDetailsAllOf); err == nil {
		*o = NiatelemetryAppDetailsAllOf(varNiatelemetryAppDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AppName")
		delete(additionalProperties, "AppStatus")
		delete(additionalProperties, "AppVersion")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryAppDetailsAllOf struct {
	value *NiatelemetryAppDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryAppDetailsAllOf) Get() *NiatelemetryAppDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryAppDetailsAllOf) Set(val *NiatelemetryAppDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryAppDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryAppDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryAppDetailsAllOf(val *NiatelemetryAppDetailsAllOf) *NullableNiatelemetryAppDetailsAllOf {
	return &NullableNiatelemetryAppDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryAppDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryAppDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
