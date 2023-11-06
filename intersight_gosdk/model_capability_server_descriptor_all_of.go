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

// CapabilityServerDescriptorAllOf Definition of the list of properties defined in 'capability.ServerDescriptor', excluding properties defined in parent classes.
type CapabilityServerDescriptorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the CIMC to VIC side-band interface is enabled on the server.
	IsNcsiEnabled *bool `json:"IsNcsiEnabled,omitempty"`
	// The form factor (blade/rack/etc) of the server. * `unknown` - The form factor of the server is unknown. * `blade` - Blade server form factor. * `rack` - Rack unit server form factor.
	ServerFormFactor     *string `json:"ServerFormFactor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityServerDescriptorAllOf CapabilityServerDescriptorAllOf

// NewCapabilityServerDescriptorAllOf instantiates a new CapabilityServerDescriptorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityServerDescriptorAllOf(classId string, objectType string) *CapabilityServerDescriptorAllOf {
	this := CapabilityServerDescriptorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityServerDescriptorAllOfWithDefaults instantiates a new CapabilityServerDescriptorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityServerDescriptorAllOfWithDefaults() *CapabilityServerDescriptorAllOf {
	this := CapabilityServerDescriptorAllOf{}
	var classId string = "capability.ServerDescriptor"
	this.ClassId = classId
	var objectType string = "capability.ServerDescriptor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityServerDescriptorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityServerDescriptorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityServerDescriptorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityServerDescriptorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityServerDescriptorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityServerDescriptorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsNcsiEnabled returns the IsNcsiEnabled field value if set, zero value otherwise.
func (o *CapabilityServerDescriptorAllOf) GetIsNcsiEnabled() bool {
	if o == nil || o.IsNcsiEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsNcsiEnabled
}

// GetIsNcsiEnabledOk returns a tuple with the IsNcsiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerDescriptorAllOf) GetIsNcsiEnabledOk() (*bool, bool) {
	if o == nil || o.IsNcsiEnabled == nil {
		return nil, false
	}
	return o.IsNcsiEnabled, true
}

// HasIsNcsiEnabled returns a boolean if a field has been set.
func (o *CapabilityServerDescriptorAllOf) HasIsNcsiEnabled() bool {
	if o != nil && o.IsNcsiEnabled != nil {
		return true
	}

	return false
}

// SetIsNcsiEnabled gets a reference to the given bool and assigns it to the IsNcsiEnabled field.
func (o *CapabilityServerDescriptorAllOf) SetIsNcsiEnabled(v bool) {
	o.IsNcsiEnabled = &v
}

// GetServerFormFactor returns the ServerFormFactor field value if set, zero value otherwise.
func (o *CapabilityServerDescriptorAllOf) GetServerFormFactor() string {
	if o == nil || o.ServerFormFactor == nil {
		var ret string
		return ret
	}
	return *o.ServerFormFactor
}

// GetServerFormFactorOk returns a tuple with the ServerFormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerDescriptorAllOf) GetServerFormFactorOk() (*string, bool) {
	if o == nil || o.ServerFormFactor == nil {
		return nil, false
	}
	return o.ServerFormFactor, true
}

// HasServerFormFactor returns a boolean if a field has been set.
func (o *CapabilityServerDescriptorAllOf) HasServerFormFactor() bool {
	if o != nil && o.ServerFormFactor != nil {
		return true
	}

	return false
}

// SetServerFormFactor gets a reference to the given string and assigns it to the ServerFormFactor field.
func (o *CapabilityServerDescriptorAllOf) SetServerFormFactor(v string) {
	o.ServerFormFactor = &v
}

func (o CapabilityServerDescriptorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsNcsiEnabled != nil {
		toSerialize["IsNcsiEnabled"] = o.IsNcsiEnabled
	}
	if o.ServerFormFactor != nil {
		toSerialize["ServerFormFactor"] = o.ServerFormFactor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityServerDescriptorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityServerDescriptorAllOf := _CapabilityServerDescriptorAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityServerDescriptorAllOf); err == nil {
		*o = CapabilityServerDescriptorAllOf(varCapabilityServerDescriptorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsNcsiEnabled")
		delete(additionalProperties, "ServerFormFactor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityServerDescriptorAllOf struct {
	value *CapabilityServerDescriptorAllOf
	isSet bool
}

func (v NullableCapabilityServerDescriptorAllOf) Get() *CapabilityServerDescriptorAllOf {
	return v.value
}

func (v *NullableCapabilityServerDescriptorAllOf) Set(val *CapabilityServerDescriptorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityServerDescriptorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityServerDescriptorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityServerDescriptorAllOf(val *CapabilityServerDescriptorAllOf) *NullableCapabilityServerDescriptorAllOf {
	return &NullableCapabilityServerDescriptorAllOf{value: val, isSet: true}
}

func (v NullableCapabilityServerDescriptorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityServerDescriptorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
