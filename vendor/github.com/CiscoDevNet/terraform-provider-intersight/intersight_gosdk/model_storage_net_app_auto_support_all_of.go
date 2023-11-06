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

// StorageNetAppAutoSupportAllOf Definition of the list of properties defined in 'storage.NetAppAutoSupport', excluding properties defined in parent classes.
type StorageNetAppAutoSupportAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies whether the AutoSupport daemon is enabled. When this setting is disabled, delivery of all AutoSupport messages is turned off.
	Enabled *bool `json:"Enabled,omitempty"`
	// The e-mail address from which the AutoSupport messages are sent.
	From *string `json:"From,omitempty"`
	// Proxy server for AutoSupport message delivery via HTTP/S.
	ProxyUrl *string `json:"ProxyUrl,omitempty"`
	// The name of the transport protocol used to deliver AutoSupport messages. * `none` - Default value of none when nothing is returned. * `smtp` - Simple Mail Transfer Protocol. * `http` - Hypertext Transfer Protocol. * `https` - Hypertext Transfer Protocol Secure.
	Transport            *string `json:"Transport,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppAutoSupportAllOf StorageNetAppAutoSupportAllOf

// NewStorageNetAppAutoSupportAllOf instantiates a new StorageNetAppAutoSupportAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppAutoSupportAllOf(classId string, objectType string) *StorageNetAppAutoSupportAllOf {
	this := StorageNetAppAutoSupportAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppAutoSupportAllOfWithDefaults instantiates a new StorageNetAppAutoSupportAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppAutoSupportAllOfWithDefaults() *StorageNetAppAutoSupportAllOf {
	this := StorageNetAppAutoSupportAllOf{}
	var classId string = "storage.NetAppAutoSupport"
	this.ClassId = classId
	var objectType string = "storage.NetAppAutoSupport"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppAutoSupportAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppAutoSupportAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppAutoSupportAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppAutoSupportAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppAutoSupportAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppAutoSupportAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppAutoSupportAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAutoSupportAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppAutoSupportAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *StorageNetAppAutoSupportAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *StorageNetAppAutoSupportAllOf) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAutoSupportAllOf) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *StorageNetAppAutoSupportAllOf) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *StorageNetAppAutoSupportAllOf) SetFrom(v string) {
	o.From = &v
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise.
func (o *StorageNetAppAutoSupportAllOf) GetProxyUrl() string {
	if o == nil || o.ProxyUrl == nil {
		var ret string
		return ret
	}
	return *o.ProxyUrl
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAutoSupportAllOf) GetProxyUrlOk() (*string, bool) {
	if o == nil || o.ProxyUrl == nil {
		return nil, false
	}
	return o.ProxyUrl, true
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *StorageNetAppAutoSupportAllOf) HasProxyUrl() bool {
	if o != nil && o.ProxyUrl != nil {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given string and assigns it to the ProxyUrl field.
func (o *StorageNetAppAutoSupportAllOf) SetProxyUrl(v string) {
	o.ProxyUrl = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *StorageNetAppAutoSupportAllOf) GetTransport() string {
	if o == nil || o.Transport == nil {
		var ret string
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAutoSupportAllOf) GetTransportOk() (*string, bool) {
	if o == nil || o.Transport == nil {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *StorageNetAppAutoSupportAllOf) HasTransport() bool {
	if o != nil && o.Transport != nil {
		return true
	}

	return false
}

// SetTransport gets a reference to the given string and assigns it to the Transport field.
func (o *StorageNetAppAutoSupportAllOf) SetTransport(v string) {
	o.Transport = &v
}

func (o StorageNetAppAutoSupportAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.From != nil {
		toSerialize["From"] = o.From
	}
	if o.ProxyUrl != nil {
		toSerialize["ProxyUrl"] = o.ProxyUrl
	}
	if o.Transport != nil {
		toSerialize["Transport"] = o.Transport
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppAutoSupportAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppAutoSupportAllOf := _StorageNetAppAutoSupportAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppAutoSupportAllOf); err == nil {
		*o = StorageNetAppAutoSupportAllOf(varStorageNetAppAutoSupportAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "From")
		delete(additionalProperties, "ProxyUrl")
		delete(additionalProperties, "Transport")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppAutoSupportAllOf struct {
	value *StorageNetAppAutoSupportAllOf
	isSet bool
}

func (v NullableStorageNetAppAutoSupportAllOf) Get() *StorageNetAppAutoSupportAllOf {
	return v.value
}

func (v *NullableStorageNetAppAutoSupportAllOf) Set(val *StorageNetAppAutoSupportAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppAutoSupportAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppAutoSupportAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppAutoSupportAllOf(val *StorageNetAppAutoSupportAllOf) *NullableStorageNetAppAutoSupportAllOf {
	return &NullableStorageNetAppAutoSupportAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppAutoSupportAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppAutoSupportAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
