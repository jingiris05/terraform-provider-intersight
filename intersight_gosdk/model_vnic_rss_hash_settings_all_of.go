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

// VnicRssHashSettingsAllOf Definition of the list of properties defined in 'vnic.RssHashSettings', excluding properties defined in parent classes.
type VnicRssHashSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When enabled, the IPv4 address is used for traffic distribution.
	Ipv4Hash *bool `json:"Ipv4Hash,omitempty"`
	// When enabled, the IPv6 extensions are used for traffic distribution.
	Ipv6ExtHash *bool `json:"Ipv6ExtHash,omitempty"`
	// When enabled, the IPv6 address is used for traffic distribution.
	Ipv6Hash *bool `json:"Ipv6Hash,omitempty"`
	// When enabled, both the IPv4 address and TCP port number are used for traffic distribution.
	TcpIpv4Hash *bool `json:"TcpIpv4Hash,omitempty"`
	// When enabled, both the IPv6 extensions and TCP port number are used for traffic distribution.
	TcpIpv6ExtHash *bool `json:"TcpIpv6ExtHash,omitempty"`
	// When enabled, both the IPv6 address and TCP port number are used for traffic distribution.
	TcpIpv6Hash *bool `json:"TcpIpv6Hash,omitempty"`
	// When enabled, both the IPv4 address and UDP port number are used for traffic distribution.
	UdpIpv4Hash *bool `json:"UdpIpv4Hash,omitempty"`
	// When enabled, both the IPv6 address and UDP port number are used for traffic distribution.
	UdpIpv6Hash          *bool `json:"UdpIpv6Hash,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicRssHashSettingsAllOf VnicRssHashSettingsAllOf

// NewVnicRssHashSettingsAllOf instantiates a new VnicRssHashSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicRssHashSettingsAllOf(classId string, objectType string) *VnicRssHashSettingsAllOf {
	this := VnicRssHashSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ipv4Hash bool = true
	this.Ipv4Hash = &ipv4Hash
	var ipv6ExtHash bool = false
	this.Ipv6ExtHash = &ipv6ExtHash
	var ipv6Hash bool = true
	this.Ipv6Hash = &ipv6Hash
	var tcpIpv4Hash bool = true
	this.TcpIpv4Hash = &tcpIpv4Hash
	var tcpIpv6ExtHash bool = false
	this.TcpIpv6ExtHash = &tcpIpv6ExtHash
	var tcpIpv6Hash bool = true
	this.TcpIpv6Hash = &tcpIpv6Hash
	var udpIpv4Hash bool = false
	this.UdpIpv4Hash = &udpIpv4Hash
	var udpIpv6Hash bool = false
	this.UdpIpv6Hash = &udpIpv6Hash
	return &this
}

// NewVnicRssHashSettingsAllOfWithDefaults instantiates a new VnicRssHashSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicRssHashSettingsAllOfWithDefaults() *VnicRssHashSettingsAllOf {
	this := VnicRssHashSettingsAllOf{}
	var classId string = "vnic.RssHashSettings"
	this.ClassId = classId
	var objectType string = "vnic.RssHashSettings"
	this.ObjectType = objectType
	var ipv4Hash bool = true
	this.Ipv4Hash = &ipv4Hash
	var ipv6ExtHash bool = false
	this.Ipv6ExtHash = &ipv6ExtHash
	var ipv6Hash bool = true
	this.Ipv6Hash = &ipv6Hash
	var tcpIpv4Hash bool = true
	this.TcpIpv4Hash = &tcpIpv4Hash
	var tcpIpv6ExtHash bool = false
	this.TcpIpv6ExtHash = &tcpIpv6ExtHash
	var tcpIpv6Hash bool = true
	this.TcpIpv6Hash = &tcpIpv6Hash
	var udpIpv4Hash bool = false
	this.UdpIpv4Hash = &udpIpv4Hash
	var udpIpv6Hash bool = false
	this.UdpIpv6Hash = &udpIpv6Hash
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicRssHashSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicRssHashSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicRssHashSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicRssHashSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpv4Hash returns the Ipv4Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettingsAllOf) GetIpv4Hash() bool {
	if o == nil || o.Ipv4Hash == nil {
		var ret bool
		return ret
	}
	return *o.Ipv4Hash
}

// GetIpv4HashOk returns a tuple with the Ipv4Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettingsAllOf) GetIpv4HashOk() (*bool, bool) {
	if o == nil || o.Ipv4Hash == nil {
		return nil, false
	}
	return o.Ipv4Hash, true
}

// HasIpv4Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettingsAllOf) HasIpv4Hash() bool {
	if o != nil && o.Ipv4Hash != nil {
		return true
	}

	return false
}

// SetIpv4Hash gets a reference to the given bool and assigns it to the Ipv4Hash field.
func (o *VnicRssHashSettingsAllOf) SetIpv4Hash(v bool) {
	o.Ipv4Hash = &v
}

// GetIpv6ExtHash returns the Ipv6ExtHash field value if set, zero value otherwise.
func (o *VnicRssHashSettingsAllOf) GetIpv6ExtHash() bool {
	if o == nil || o.Ipv6ExtHash == nil {
		var ret bool
		return ret
	}
	return *o.Ipv6ExtHash
}

// GetIpv6ExtHashOk returns a tuple with the Ipv6ExtHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettingsAllOf) GetIpv6ExtHashOk() (*bool, bool) {
	if o == nil || o.Ipv6ExtHash == nil {
		return nil, false
	}
	return o.Ipv6ExtHash, true
}

// HasIpv6ExtHash returns a boolean if a field has been set.
func (o *VnicRssHashSettingsAllOf) HasIpv6ExtHash() bool {
	if o != nil && o.Ipv6ExtHash != nil {
		return true
	}

	return false
}

// SetIpv6ExtHash gets a reference to the given bool and assigns it to the Ipv6ExtHash field.
func (o *VnicRssHashSettingsAllOf) SetIpv6ExtHash(v bool) {
	o.Ipv6ExtHash = &v
}

// GetIpv6Hash returns the Ipv6Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettingsAllOf) GetIpv6Hash() bool {
	if o == nil || o.Ipv6Hash == nil {
		var ret bool
		return ret
	}
	return *o.Ipv6Hash
}

// GetIpv6HashOk returns a tuple with the Ipv6Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettingsAllOf) GetIpv6HashOk() (*bool, bool) {
	if o == nil || o.Ipv6Hash == nil {
		return nil, false
	}
	return o.Ipv6Hash, true
}

// HasIpv6Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettingsAllOf) HasIpv6Hash() bool {
	if o != nil && o.Ipv6Hash != nil {
		return true
	}

	return false
}

// SetIpv6Hash gets a reference to the given bool and assigns it to the Ipv6Hash field.
func (o *VnicRssHashSettingsAllOf) SetIpv6Hash(v bool) {
	o.Ipv6Hash = &v
}

// GetTcpIpv4Hash returns the TcpIpv4Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettingsAllOf) GetTcpIpv4Hash() bool {
	if o == nil || o.TcpIpv4Hash == nil {
		var ret bool
		return ret
	}
	return *o.TcpIpv4Hash
}

// GetTcpIpv4HashOk returns a tuple with the TcpIpv4Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettingsAllOf) GetTcpIpv4HashOk() (*bool, bool) {
	if o == nil || o.TcpIpv4Hash == nil {
		return nil, false
	}
	return o.TcpIpv4Hash, true
}

// HasTcpIpv4Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettingsAllOf) HasTcpIpv4Hash() bool {
	if o != nil && o.TcpIpv4Hash != nil {
		return true
	}

	return false
}

// SetTcpIpv4Hash gets a reference to the given bool and assigns it to the TcpIpv4Hash field.
func (o *VnicRssHashSettingsAllOf) SetTcpIpv4Hash(v bool) {
	o.TcpIpv4Hash = &v
}

// GetTcpIpv6ExtHash returns the TcpIpv6ExtHash field value if set, zero value otherwise.
func (o *VnicRssHashSettingsAllOf) GetTcpIpv6ExtHash() bool {
	if o == nil || o.TcpIpv6ExtHash == nil {
		var ret bool
		return ret
	}
	return *o.TcpIpv6ExtHash
}

// GetTcpIpv6ExtHashOk returns a tuple with the TcpIpv6ExtHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettingsAllOf) GetTcpIpv6ExtHashOk() (*bool, bool) {
	if o == nil || o.TcpIpv6ExtHash == nil {
		return nil, false
	}
	return o.TcpIpv6ExtHash, true
}

// HasTcpIpv6ExtHash returns a boolean if a field has been set.
func (o *VnicRssHashSettingsAllOf) HasTcpIpv6ExtHash() bool {
	if o != nil && o.TcpIpv6ExtHash != nil {
		return true
	}

	return false
}

// SetTcpIpv6ExtHash gets a reference to the given bool and assigns it to the TcpIpv6ExtHash field.
func (o *VnicRssHashSettingsAllOf) SetTcpIpv6ExtHash(v bool) {
	o.TcpIpv6ExtHash = &v
}

// GetTcpIpv6Hash returns the TcpIpv6Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettingsAllOf) GetTcpIpv6Hash() bool {
	if o == nil || o.TcpIpv6Hash == nil {
		var ret bool
		return ret
	}
	return *o.TcpIpv6Hash
}

// GetTcpIpv6HashOk returns a tuple with the TcpIpv6Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettingsAllOf) GetTcpIpv6HashOk() (*bool, bool) {
	if o == nil || o.TcpIpv6Hash == nil {
		return nil, false
	}
	return o.TcpIpv6Hash, true
}

// HasTcpIpv6Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettingsAllOf) HasTcpIpv6Hash() bool {
	if o != nil && o.TcpIpv6Hash != nil {
		return true
	}

	return false
}

// SetTcpIpv6Hash gets a reference to the given bool and assigns it to the TcpIpv6Hash field.
func (o *VnicRssHashSettingsAllOf) SetTcpIpv6Hash(v bool) {
	o.TcpIpv6Hash = &v
}

// GetUdpIpv4Hash returns the UdpIpv4Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettingsAllOf) GetUdpIpv4Hash() bool {
	if o == nil || o.UdpIpv4Hash == nil {
		var ret bool
		return ret
	}
	return *o.UdpIpv4Hash
}

// GetUdpIpv4HashOk returns a tuple with the UdpIpv4Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettingsAllOf) GetUdpIpv4HashOk() (*bool, bool) {
	if o == nil || o.UdpIpv4Hash == nil {
		return nil, false
	}
	return o.UdpIpv4Hash, true
}

// HasUdpIpv4Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettingsAllOf) HasUdpIpv4Hash() bool {
	if o != nil && o.UdpIpv4Hash != nil {
		return true
	}

	return false
}

// SetUdpIpv4Hash gets a reference to the given bool and assigns it to the UdpIpv4Hash field.
func (o *VnicRssHashSettingsAllOf) SetUdpIpv4Hash(v bool) {
	o.UdpIpv4Hash = &v
}

// GetUdpIpv6Hash returns the UdpIpv6Hash field value if set, zero value otherwise.
func (o *VnicRssHashSettingsAllOf) GetUdpIpv6Hash() bool {
	if o == nil || o.UdpIpv6Hash == nil {
		var ret bool
		return ret
	}
	return *o.UdpIpv6Hash
}

// GetUdpIpv6HashOk returns a tuple with the UdpIpv6Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRssHashSettingsAllOf) GetUdpIpv6HashOk() (*bool, bool) {
	if o == nil || o.UdpIpv6Hash == nil {
		return nil, false
	}
	return o.UdpIpv6Hash, true
}

// HasUdpIpv6Hash returns a boolean if a field has been set.
func (o *VnicRssHashSettingsAllOf) HasUdpIpv6Hash() bool {
	if o != nil && o.UdpIpv6Hash != nil {
		return true
	}

	return false
}

// SetUdpIpv6Hash gets a reference to the given bool and assigns it to the UdpIpv6Hash field.
func (o *VnicRssHashSettingsAllOf) SetUdpIpv6Hash(v bool) {
	o.UdpIpv6Hash = &v
}

func (o VnicRssHashSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Ipv4Hash != nil {
		toSerialize["Ipv4Hash"] = o.Ipv4Hash
	}
	if o.Ipv6ExtHash != nil {
		toSerialize["Ipv6ExtHash"] = o.Ipv6ExtHash
	}
	if o.Ipv6Hash != nil {
		toSerialize["Ipv6Hash"] = o.Ipv6Hash
	}
	if o.TcpIpv4Hash != nil {
		toSerialize["TcpIpv4Hash"] = o.TcpIpv4Hash
	}
	if o.TcpIpv6ExtHash != nil {
		toSerialize["TcpIpv6ExtHash"] = o.TcpIpv6ExtHash
	}
	if o.TcpIpv6Hash != nil {
		toSerialize["TcpIpv6Hash"] = o.TcpIpv6Hash
	}
	if o.UdpIpv4Hash != nil {
		toSerialize["UdpIpv4Hash"] = o.UdpIpv4Hash
	}
	if o.UdpIpv6Hash != nil {
		toSerialize["UdpIpv6Hash"] = o.UdpIpv6Hash
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicRssHashSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicRssHashSettingsAllOf := _VnicRssHashSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varVnicRssHashSettingsAllOf); err == nil {
		*o = VnicRssHashSettingsAllOf(varVnicRssHashSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ipv4Hash")
		delete(additionalProperties, "Ipv6ExtHash")
		delete(additionalProperties, "Ipv6Hash")
		delete(additionalProperties, "TcpIpv4Hash")
		delete(additionalProperties, "TcpIpv6ExtHash")
		delete(additionalProperties, "TcpIpv6Hash")
		delete(additionalProperties, "UdpIpv4Hash")
		delete(additionalProperties, "UdpIpv6Hash")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicRssHashSettingsAllOf struct {
	value *VnicRssHashSettingsAllOf
	isSet bool
}

func (v NullableVnicRssHashSettingsAllOf) Get() *VnicRssHashSettingsAllOf {
	return v.value
}

func (v *NullableVnicRssHashSettingsAllOf) Set(val *VnicRssHashSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicRssHashSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicRssHashSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicRssHashSettingsAllOf(val *VnicRssHashSettingsAllOf) *NullableVnicRssHashSettingsAllOf {
	return &NullableVnicRssHashSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicRssHashSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicRssHashSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
