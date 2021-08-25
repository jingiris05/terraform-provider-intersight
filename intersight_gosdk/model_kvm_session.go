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

// KvmSession Virtual KVM Session that provides Single Sign-On access to the vKVM console of the server. The vKVM access can be direct or can be tunneled by specifying the tunnel to be used for the access.
type KvmSession struct {
	SessionAbstractSubSession
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Temporary one-time password for vKVM access.
	OneTimePassword *string `json:"OneTimePassword,omitempty"`
	// Indicates if vKVM SSO is supported on the server.
	SsoSupported *bool `json:"SsoSupported,omitempty"`
	// Username used for vKVM access.
	Username             *string                              `json:"Username,omitempty"`
	Device               *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	Server               *ComputePhysicalRelationship         `json:"Server,omitempty"`
	Tunnel               *KvmTunnelRelationship               `json:"Tunnel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmSession KvmSession

// NewKvmSession instantiates a new KvmSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmSession(classId string, objectType string) *KvmSession {
	this := KvmSession{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKvmSessionWithDefaults instantiates a new KvmSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmSessionWithDefaults() *KvmSession {
	this := KvmSession{}
	var classId string = "kvm.Session"
	this.ClassId = classId
	var objectType string = "kvm.Session"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmSession) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmSession) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmSession) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmSession) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmSession) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmSession) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOneTimePassword returns the OneTimePassword field value if set, zero value otherwise.
func (o *KvmSession) GetOneTimePassword() string {
	if o == nil || o.OneTimePassword == nil {
		var ret string
		return ret
	}
	return *o.OneTimePassword
}

// GetOneTimePasswordOk returns a tuple with the OneTimePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetOneTimePasswordOk() (*string, bool) {
	if o == nil || o.OneTimePassword == nil {
		return nil, false
	}
	return o.OneTimePassword, true
}

// HasOneTimePassword returns a boolean if a field has been set.
func (o *KvmSession) HasOneTimePassword() bool {
	if o != nil && o.OneTimePassword != nil {
		return true
	}

	return false
}

// SetOneTimePassword gets a reference to the given string and assigns it to the OneTimePassword field.
func (o *KvmSession) SetOneTimePassword(v string) {
	o.OneTimePassword = &v
}

// GetSsoSupported returns the SsoSupported field value if set, zero value otherwise.
func (o *KvmSession) GetSsoSupported() bool {
	if o == nil || o.SsoSupported == nil {
		var ret bool
		return ret
	}
	return *o.SsoSupported
}

// GetSsoSupportedOk returns a tuple with the SsoSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetSsoSupportedOk() (*bool, bool) {
	if o == nil || o.SsoSupported == nil {
		return nil, false
	}
	return o.SsoSupported, true
}

// HasSsoSupported returns a boolean if a field has been set.
func (o *KvmSession) HasSsoSupported() bool {
	if o != nil && o.SsoSupported != nil {
		return true
	}

	return false
}

// SetSsoSupported gets a reference to the given bool and assigns it to the SsoSupported field.
func (o *KvmSession) SetSsoSupported(v bool) {
	o.SsoSupported = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *KvmSession) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *KvmSession) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *KvmSession) SetUsername(v string) {
	o.Username = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *KvmSession) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *KvmSession) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *KvmSession) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *KvmSession) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *KvmSession) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *KvmSession) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *KvmSession) GetTunnel() KvmTunnelRelationship {
	if o == nil || o.Tunnel == nil {
		var ret KvmTunnelRelationship
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmSession) GetTunnelOk() (*KvmTunnelRelationship, bool) {
	if o == nil || o.Tunnel == nil {
		return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *KvmSession) HasTunnel() bool {
	if o != nil && o.Tunnel != nil {
		return true
	}

	return false
}

// SetTunnel gets a reference to the given KvmTunnelRelationship and assigns it to the Tunnel field.
func (o *KvmSession) SetTunnel(v KvmTunnelRelationship) {
	o.Tunnel = &v
}

func (o KvmSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSessionAbstractSubSession, errSessionAbstractSubSession := json.Marshal(o.SessionAbstractSubSession)
	if errSessionAbstractSubSession != nil {
		return []byte{}, errSessionAbstractSubSession
	}
	errSessionAbstractSubSession = json.Unmarshal([]byte(serializedSessionAbstractSubSession), &toSerialize)
	if errSessionAbstractSubSession != nil {
		return []byte{}, errSessionAbstractSubSession
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OneTimePassword != nil {
		toSerialize["OneTimePassword"] = o.OneTimePassword
	}
	if o.SsoSupported != nil {
		toSerialize["SsoSupported"] = o.SsoSupported
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.Tunnel != nil {
		toSerialize["Tunnel"] = o.Tunnel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KvmSession) UnmarshalJSON(bytes []byte) (err error) {
	type KvmSessionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Temporary one-time password for vKVM access.
		OneTimePassword *string `json:"OneTimePassword,omitempty"`
		// Indicates if vKVM SSO is supported on the server.
		SsoSupported *bool `json:"SsoSupported,omitempty"`
		// Username used for vKVM access.
		Username *string                              `json:"Username,omitempty"`
		Device   *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
		Server   *ComputePhysicalRelationship         `json:"Server,omitempty"`
		Tunnel   *KvmTunnelRelationship               `json:"Tunnel,omitempty"`
	}

	varKvmSessionWithoutEmbeddedStruct := KvmSessionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKvmSessionWithoutEmbeddedStruct)
	if err == nil {
		varKvmSession := _KvmSession{}
		varKvmSession.ClassId = varKvmSessionWithoutEmbeddedStruct.ClassId
		varKvmSession.ObjectType = varKvmSessionWithoutEmbeddedStruct.ObjectType
		varKvmSession.OneTimePassword = varKvmSessionWithoutEmbeddedStruct.OneTimePassword
		varKvmSession.SsoSupported = varKvmSessionWithoutEmbeddedStruct.SsoSupported
		varKvmSession.Username = varKvmSessionWithoutEmbeddedStruct.Username
		varKvmSession.Device = varKvmSessionWithoutEmbeddedStruct.Device
		varKvmSession.Server = varKvmSessionWithoutEmbeddedStruct.Server
		varKvmSession.Tunnel = varKvmSessionWithoutEmbeddedStruct.Tunnel
		*o = KvmSession(varKvmSession)
	} else {
		return err
	}

	varKvmSession := _KvmSession{}

	err = json.Unmarshal(bytes, &varKvmSession)
	if err == nil {
		o.SessionAbstractSubSession = varKvmSession.SessionAbstractSubSession
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OneTimePassword")
		delete(additionalProperties, "SsoSupported")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "Tunnel")

		// remove fields from embedded structs
		reflectSessionAbstractSubSession := reflect.ValueOf(o.SessionAbstractSubSession)
		for i := 0; i < reflectSessionAbstractSubSession.Type().NumField(); i++ {
			t := reflectSessionAbstractSubSession.Type().Field(i)

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

type NullableKvmSession struct {
	value *KvmSession
	isSet bool
}

func (v NullableKvmSession) Get() *KvmSession {
	return v.value
}

func (v *NullableKvmSession) Set(val *KvmSession) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmSession) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmSession(val *KvmSession) *NullableKvmSession {
	return &NullableKvmSession{value: val, isSet: true}
}

func (v NullableKvmSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
