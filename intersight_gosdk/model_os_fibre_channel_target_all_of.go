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

// OsFibreChannelTargetAllOf Definition of the list of properties defined in 'os.FibreChannelTarget', excluding properties defined in parent classes.
type OsFibreChannelTargetAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The WWPN Address of the underlying fibre channel interface at initator used for SAN boot. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 20:00:D4:C9:3C:35:02:01.
	InitiatorWwpn *string `json:"InitiatorWwpn,omitempty"`
	// Logical Unit Number (LUN) of the install target.
	LunId *int64 `json:"LunId,omitempty"`
	// The WWPN Address of the underlying fibre channel interface at target used by the SAN boot device. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.  For example, 51:4F:0C:50:14:1F:AF:01.
	TargetWwpn           *string `json:"TargetWwpn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsFibreChannelTargetAllOf OsFibreChannelTargetAllOf

// NewOsFibreChannelTargetAllOf instantiates a new OsFibreChannelTargetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsFibreChannelTargetAllOf(classId string, objectType string) *OsFibreChannelTargetAllOf {
	this := OsFibreChannelTargetAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var lunId int64 = 0
	this.LunId = &lunId
	return &this
}

// NewOsFibreChannelTargetAllOfWithDefaults instantiates a new OsFibreChannelTargetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsFibreChannelTargetAllOfWithDefaults() *OsFibreChannelTargetAllOf {
	this := OsFibreChannelTargetAllOf{}
	var classId string = "os.FibreChannelTarget"
	this.ClassId = classId
	var objectType string = "os.FibreChannelTarget"
	this.ObjectType = objectType
	var lunId int64 = 0
	this.LunId = &lunId
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsFibreChannelTargetAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsFibreChannelTargetAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsFibreChannelTargetAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsFibreChannelTargetAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsFibreChannelTargetAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsFibreChannelTargetAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInitiatorWwpn returns the InitiatorWwpn field value if set, zero value otherwise.
func (o *OsFibreChannelTargetAllOf) GetInitiatorWwpn() string {
	if o == nil || o.InitiatorWwpn == nil {
		var ret string
		return ret
	}
	return *o.InitiatorWwpn
}

// GetInitiatorWwpnOk returns a tuple with the InitiatorWwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsFibreChannelTargetAllOf) GetInitiatorWwpnOk() (*string, bool) {
	if o == nil || o.InitiatorWwpn == nil {
		return nil, false
	}
	return o.InitiatorWwpn, true
}

// HasInitiatorWwpn returns a boolean if a field has been set.
func (o *OsFibreChannelTargetAllOf) HasInitiatorWwpn() bool {
	if o != nil && o.InitiatorWwpn != nil {
		return true
	}

	return false
}

// SetInitiatorWwpn gets a reference to the given string and assigns it to the InitiatorWwpn field.
func (o *OsFibreChannelTargetAllOf) SetInitiatorWwpn(v string) {
	o.InitiatorWwpn = &v
}

// GetLunId returns the LunId field value if set, zero value otherwise.
func (o *OsFibreChannelTargetAllOf) GetLunId() int64 {
	if o == nil || o.LunId == nil {
		var ret int64
		return ret
	}
	return *o.LunId
}

// GetLunIdOk returns a tuple with the LunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsFibreChannelTargetAllOf) GetLunIdOk() (*int64, bool) {
	if o == nil || o.LunId == nil {
		return nil, false
	}
	return o.LunId, true
}

// HasLunId returns a boolean if a field has been set.
func (o *OsFibreChannelTargetAllOf) HasLunId() bool {
	if o != nil && o.LunId != nil {
		return true
	}

	return false
}

// SetLunId gets a reference to the given int64 and assigns it to the LunId field.
func (o *OsFibreChannelTargetAllOf) SetLunId(v int64) {
	o.LunId = &v
}

// GetTargetWwpn returns the TargetWwpn field value if set, zero value otherwise.
func (o *OsFibreChannelTargetAllOf) GetTargetWwpn() string {
	if o == nil || o.TargetWwpn == nil {
		var ret string
		return ret
	}
	return *o.TargetWwpn
}

// GetTargetWwpnOk returns a tuple with the TargetWwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsFibreChannelTargetAllOf) GetTargetWwpnOk() (*string, bool) {
	if o == nil || o.TargetWwpn == nil {
		return nil, false
	}
	return o.TargetWwpn, true
}

// HasTargetWwpn returns a boolean if a field has been set.
func (o *OsFibreChannelTargetAllOf) HasTargetWwpn() bool {
	if o != nil && o.TargetWwpn != nil {
		return true
	}

	return false
}

// SetTargetWwpn gets a reference to the given string and assigns it to the TargetWwpn field.
func (o *OsFibreChannelTargetAllOf) SetTargetWwpn(v string) {
	o.TargetWwpn = &v
}

func (o OsFibreChannelTargetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InitiatorWwpn != nil {
		toSerialize["InitiatorWwpn"] = o.InitiatorWwpn
	}
	if o.LunId != nil {
		toSerialize["LunId"] = o.LunId
	}
	if o.TargetWwpn != nil {
		toSerialize["TargetWwpn"] = o.TargetWwpn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsFibreChannelTargetAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsFibreChannelTargetAllOf := _OsFibreChannelTargetAllOf{}

	if err = json.Unmarshal(bytes, &varOsFibreChannelTargetAllOf); err == nil {
		*o = OsFibreChannelTargetAllOf(varOsFibreChannelTargetAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InitiatorWwpn")
		delete(additionalProperties, "LunId")
		delete(additionalProperties, "TargetWwpn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsFibreChannelTargetAllOf struct {
	value *OsFibreChannelTargetAllOf
	isSet bool
}

func (v NullableOsFibreChannelTargetAllOf) Get() *OsFibreChannelTargetAllOf {
	return v.value
}

func (v *NullableOsFibreChannelTargetAllOf) Set(val *OsFibreChannelTargetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsFibreChannelTargetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsFibreChannelTargetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsFibreChannelTargetAllOf(val *OsFibreChannelTargetAllOf) *NullableOsFibreChannelTargetAllOf {
	return &NullableOsFibreChannelTargetAllOf{value: val, isSet: true}
}

func (v NullableOsFibreChannelTargetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsFibreChannelTargetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
