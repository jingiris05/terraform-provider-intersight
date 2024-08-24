/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18012
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the NiatelemetryVpcDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryVpcDetails{}

// NiatelemetryVpcDetails Stores information related to VPC pairs on every switch.
type NiatelemetryVpcDetails struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns boolean if VPC is configured on switch or not.
	IsVpcConfigured *bool `json:"IsVpcConfigured,omitempty"`
	// Returns peer switch id if VPC configured.
	PeerSwitchDbId *int64 `json:"PeerSwitchDbId,omitempty"`
	// Returns the switch id of the switch.
	SwitchDbId           *int64 `json:"SwitchDbId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryVpcDetails NiatelemetryVpcDetails

// NewNiatelemetryVpcDetails instantiates a new NiatelemetryVpcDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryVpcDetails(classId string, objectType string) *NiatelemetryVpcDetails {
	this := NiatelemetryVpcDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryVpcDetailsWithDefaults instantiates a new NiatelemetryVpcDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryVpcDetailsWithDefaults() *NiatelemetryVpcDetails {
	this := NiatelemetryVpcDetails{}
	var classId string = "niatelemetry.VpcDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.VpcDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryVpcDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryVpcDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryVpcDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.VpcDetails" of the ClassId field.
func (o *NiatelemetryVpcDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.VpcDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryVpcDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryVpcDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryVpcDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.VpcDetails" of the ObjectType field.
func (o *NiatelemetryVpcDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.VpcDetails"
}

// GetIsVpcConfigured returns the IsVpcConfigured field value if set, zero value otherwise.
func (o *NiatelemetryVpcDetails) GetIsVpcConfigured() bool {
	if o == nil || IsNil(o.IsVpcConfigured) {
		var ret bool
		return ret
	}
	return *o.IsVpcConfigured
}

// GetIsVpcConfiguredOk returns a tuple with the IsVpcConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryVpcDetails) GetIsVpcConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVpcConfigured) {
		return nil, false
	}
	return o.IsVpcConfigured, true
}

// HasIsVpcConfigured returns a boolean if a field has been set.
func (o *NiatelemetryVpcDetails) HasIsVpcConfigured() bool {
	if o != nil && !IsNil(o.IsVpcConfigured) {
		return true
	}

	return false
}

// SetIsVpcConfigured gets a reference to the given bool and assigns it to the IsVpcConfigured field.
func (o *NiatelemetryVpcDetails) SetIsVpcConfigured(v bool) {
	o.IsVpcConfigured = &v
}

// GetPeerSwitchDbId returns the PeerSwitchDbId field value if set, zero value otherwise.
func (o *NiatelemetryVpcDetails) GetPeerSwitchDbId() int64 {
	if o == nil || IsNil(o.PeerSwitchDbId) {
		var ret int64
		return ret
	}
	return *o.PeerSwitchDbId
}

// GetPeerSwitchDbIdOk returns a tuple with the PeerSwitchDbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryVpcDetails) GetPeerSwitchDbIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PeerSwitchDbId) {
		return nil, false
	}
	return o.PeerSwitchDbId, true
}

// HasPeerSwitchDbId returns a boolean if a field has been set.
func (o *NiatelemetryVpcDetails) HasPeerSwitchDbId() bool {
	if o != nil && !IsNil(o.PeerSwitchDbId) {
		return true
	}

	return false
}

// SetPeerSwitchDbId gets a reference to the given int64 and assigns it to the PeerSwitchDbId field.
func (o *NiatelemetryVpcDetails) SetPeerSwitchDbId(v int64) {
	o.PeerSwitchDbId = &v
}

// GetSwitchDbId returns the SwitchDbId field value if set, zero value otherwise.
func (o *NiatelemetryVpcDetails) GetSwitchDbId() int64 {
	if o == nil || IsNil(o.SwitchDbId) {
		var ret int64
		return ret
	}
	return *o.SwitchDbId
}

// GetSwitchDbIdOk returns a tuple with the SwitchDbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryVpcDetails) GetSwitchDbIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SwitchDbId) {
		return nil, false
	}
	return o.SwitchDbId, true
}

// HasSwitchDbId returns a boolean if a field has been set.
func (o *NiatelemetryVpcDetails) HasSwitchDbId() bool {
	if o != nil && !IsNil(o.SwitchDbId) {
		return true
	}

	return false
}

// SetSwitchDbId gets a reference to the given int64 and assigns it to the SwitchDbId field.
func (o *NiatelemetryVpcDetails) SetSwitchDbId(v int64) {
	o.SwitchDbId = &v
}

func (o NiatelemetryVpcDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryVpcDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.IsVpcConfigured) {
		toSerialize["IsVpcConfigured"] = o.IsVpcConfigured
	}
	if !IsNil(o.PeerSwitchDbId) {
		toSerialize["PeerSwitchDbId"] = o.PeerSwitchDbId
	}
	if !IsNil(o.SwitchDbId) {
		toSerialize["SwitchDbId"] = o.SwitchDbId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryVpcDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type NiatelemetryVpcDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Returns boolean if VPC is configured on switch or not.
		IsVpcConfigured *bool `json:"IsVpcConfigured,omitempty"`
		// Returns peer switch id if VPC configured.
		PeerSwitchDbId *int64 `json:"PeerSwitchDbId,omitempty"`
		// Returns the switch id of the switch.
		SwitchDbId *int64 `json:"SwitchDbId,omitempty"`
	}

	varNiatelemetryVpcDetailsWithoutEmbeddedStruct := NiatelemetryVpcDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryVpcDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryVpcDetails := _NiatelemetryVpcDetails{}
		varNiatelemetryVpcDetails.ClassId = varNiatelemetryVpcDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryVpcDetails.ObjectType = varNiatelemetryVpcDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryVpcDetails.IsVpcConfigured = varNiatelemetryVpcDetailsWithoutEmbeddedStruct.IsVpcConfigured
		varNiatelemetryVpcDetails.PeerSwitchDbId = varNiatelemetryVpcDetailsWithoutEmbeddedStruct.PeerSwitchDbId
		varNiatelemetryVpcDetails.SwitchDbId = varNiatelemetryVpcDetailsWithoutEmbeddedStruct.SwitchDbId
		*o = NiatelemetryVpcDetails(varNiatelemetryVpcDetails)
	} else {
		return err
	}

	varNiatelemetryVpcDetails := _NiatelemetryVpcDetails{}

	err = json.Unmarshal(data, &varNiatelemetryVpcDetails)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryVpcDetails.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsVpcConfigured")
		delete(additionalProperties, "PeerSwitchDbId")
		delete(additionalProperties, "SwitchDbId")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableNiatelemetryVpcDetails struct {
	value *NiatelemetryVpcDetails
	isSet bool
}

func (v NullableNiatelemetryVpcDetails) Get() *NiatelemetryVpcDetails {
	return v.value
}

func (v *NullableNiatelemetryVpcDetails) Set(val *NiatelemetryVpcDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryVpcDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryVpcDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryVpcDetails(val *NiatelemetryVpcDetails) *NullableNiatelemetryVpcDetails {
	return &NullableNiatelemetryVpcDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryVpcDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryVpcDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
