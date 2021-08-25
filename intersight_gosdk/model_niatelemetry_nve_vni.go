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

// NiatelemetryNveVni Stores nve vni count information per switch.
type NiatelemetryNveVni struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return value of cp vni count.
	CpVniCount *int64 `json:"CpVniCount,omitempty"`
	// Return value of cp vni down count.
	CpVniDown *int64 `json:"CpVniDown,omitempty"`
	// Return value of cp vni up count.
	CpVniUp *int64 `json:"CpVniUp,omitempty"`
	// Return value of dp vni count.
	DpVniCount *int64 `json:"DpVniCount,omitempty"`
	// Return value of cp vni down count.
	DpVniDown *int64 `json:"DpVniDown,omitempty"`
	// Return value of cp vni up count.
	DpVniUp              *int64 `json:"DpVniUp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNveVni NiatelemetryNveVni

// NewNiatelemetryNveVni instantiates a new NiatelemetryNveVni object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNveVni(classId string, objectType string) *NiatelemetryNveVni {
	this := NiatelemetryNveVni{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNveVniWithDefaults instantiates a new NiatelemetryNveVni object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNveVniWithDefaults() *NiatelemetryNveVni {
	this := NiatelemetryNveVni{}
	var classId string = "niatelemetry.NveVni"
	this.ClassId = classId
	var objectType string = "niatelemetry.NveVni"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNveVni) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVni) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNveVni) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNveVni) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVni) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNveVni) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpVniCount returns the CpVniCount field value if set, zero value otherwise.
func (o *NiatelemetryNveVni) GetCpVniCount() int64 {
	if o == nil || o.CpVniCount == nil {
		var ret int64
		return ret
	}
	return *o.CpVniCount
}

// GetCpVniCountOk returns a tuple with the CpVniCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVni) GetCpVniCountOk() (*int64, bool) {
	if o == nil || o.CpVniCount == nil {
		return nil, false
	}
	return o.CpVniCount, true
}

// HasCpVniCount returns a boolean if a field has been set.
func (o *NiatelemetryNveVni) HasCpVniCount() bool {
	if o != nil && o.CpVniCount != nil {
		return true
	}

	return false
}

// SetCpVniCount gets a reference to the given int64 and assigns it to the CpVniCount field.
func (o *NiatelemetryNveVni) SetCpVniCount(v int64) {
	o.CpVniCount = &v
}

// GetCpVniDown returns the CpVniDown field value if set, zero value otherwise.
func (o *NiatelemetryNveVni) GetCpVniDown() int64 {
	if o == nil || o.CpVniDown == nil {
		var ret int64
		return ret
	}
	return *o.CpVniDown
}

// GetCpVniDownOk returns a tuple with the CpVniDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVni) GetCpVniDownOk() (*int64, bool) {
	if o == nil || o.CpVniDown == nil {
		return nil, false
	}
	return o.CpVniDown, true
}

// HasCpVniDown returns a boolean if a field has been set.
func (o *NiatelemetryNveVni) HasCpVniDown() bool {
	if o != nil && o.CpVniDown != nil {
		return true
	}

	return false
}

// SetCpVniDown gets a reference to the given int64 and assigns it to the CpVniDown field.
func (o *NiatelemetryNveVni) SetCpVniDown(v int64) {
	o.CpVniDown = &v
}

// GetCpVniUp returns the CpVniUp field value if set, zero value otherwise.
func (o *NiatelemetryNveVni) GetCpVniUp() int64 {
	if o == nil || o.CpVniUp == nil {
		var ret int64
		return ret
	}
	return *o.CpVniUp
}

// GetCpVniUpOk returns a tuple with the CpVniUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVni) GetCpVniUpOk() (*int64, bool) {
	if o == nil || o.CpVniUp == nil {
		return nil, false
	}
	return o.CpVniUp, true
}

// HasCpVniUp returns a boolean if a field has been set.
func (o *NiatelemetryNveVni) HasCpVniUp() bool {
	if o != nil && o.CpVniUp != nil {
		return true
	}

	return false
}

// SetCpVniUp gets a reference to the given int64 and assigns it to the CpVniUp field.
func (o *NiatelemetryNveVni) SetCpVniUp(v int64) {
	o.CpVniUp = &v
}

// GetDpVniCount returns the DpVniCount field value if set, zero value otherwise.
func (o *NiatelemetryNveVni) GetDpVniCount() int64 {
	if o == nil || o.DpVniCount == nil {
		var ret int64
		return ret
	}
	return *o.DpVniCount
}

// GetDpVniCountOk returns a tuple with the DpVniCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVni) GetDpVniCountOk() (*int64, bool) {
	if o == nil || o.DpVniCount == nil {
		return nil, false
	}
	return o.DpVniCount, true
}

// HasDpVniCount returns a boolean if a field has been set.
func (o *NiatelemetryNveVni) HasDpVniCount() bool {
	if o != nil && o.DpVniCount != nil {
		return true
	}

	return false
}

// SetDpVniCount gets a reference to the given int64 and assigns it to the DpVniCount field.
func (o *NiatelemetryNveVni) SetDpVniCount(v int64) {
	o.DpVniCount = &v
}

// GetDpVniDown returns the DpVniDown field value if set, zero value otherwise.
func (o *NiatelemetryNveVni) GetDpVniDown() int64 {
	if o == nil || o.DpVniDown == nil {
		var ret int64
		return ret
	}
	return *o.DpVniDown
}

// GetDpVniDownOk returns a tuple with the DpVniDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVni) GetDpVniDownOk() (*int64, bool) {
	if o == nil || o.DpVniDown == nil {
		return nil, false
	}
	return o.DpVniDown, true
}

// HasDpVniDown returns a boolean if a field has been set.
func (o *NiatelemetryNveVni) HasDpVniDown() bool {
	if o != nil && o.DpVniDown != nil {
		return true
	}

	return false
}

// SetDpVniDown gets a reference to the given int64 and assigns it to the DpVniDown field.
func (o *NiatelemetryNveVni) SetDpVniDown(v int64) {
	o.DpVniDown = &v
}

// GetDpVniUp returns the DpVniUp field value if set, zero value otherwise.
func (o *NiatelemetryNveVni) GetDpVniUp() int64 {
	if o == nil || o.DpVniUp == nil {
		var ret int64
		return ret
	}
	return *o.DpVniUp
}

// GetDpVniUpOk returns a tuple with the DpVniUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVni) GetDpVniUpOk() (*int64, bool) {
	if o == nil || o.DpVniUp == nil {
		return nil, false
	}
	return o.DpVniUp, true
}

// HasDpVniUp returns a boolean if a field has been set.
func (o *NiatelemetryNveVni) HasDpVniUp() bool {
	if o != nil && o.DpVniUp != nil {
		return true
	}

	return false
}

// SetDpVniUp gets a reference to the given int64 and assigns it to the DpVniUp field.
func (o *NiatelemetryNveVni) SetDpVniUp(v int64) {
	o.DpVniUp = &v
}

func (o NiatelemetryNveVni) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CpVniCount != nil {
		toSerialize["CpVniCount"] = o.CpVniCount
	}
	if o.CpVniDown != nil {
		toSerialize["CpVniDown"] = o.CpVniDown
	}
	if o.CpVniUp != nil {
		toSerialize["CpVniUp"] = o.CpVniUp
	}
	if o.DpVniCount != nil {
		toSerialize["DpVniCount"] = o.DpVniCount
	}
	if o.DpVniDown != nil {
		toSerialize["DpVniDown"] = o.DpVniDown
	}
	if o.DpVniUp != nil {
		toSerialize["DpVniUp"] = o.DpVniUp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNveVni) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryNveVniWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Return value of cp vni count.
		CpVniCount *int64 `json:"CpVniCount,omitempty"`
		// Return value of cp vni down count.
		CpVniDown *int64 `json:"CpVniDown,omitempty"`
		// Return value of cp vni up count.
		CpVniUp *int64 `json:"CpVniUp,omitempty"`
		// Return value of dp vni count.
		DpVniCount *int64 `json:"DpVniCount,omitempty"`
		// Return value of cp vni down count.
		DpVniDown *int64 `json:"DpVniDown,omitempty"`
		// Return value of cp vni up count.
		DpVniUp *int64 `json:"DpVniUp,omitempty"`
	}

	varNiatelemetryNveVniWithoutEmbeddedStruct := NiatelemetryNveVniWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryNveVniWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNveVni := _NiatelemetryNveVni{}
		varNiatelemetryNveVni.ClassId = varNiatelemetryNveVniWithoutEmbeddedStruct.ClassId
		varNiatelemetryNveVni.ObjectType = varNiatelemetryNveVniWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNveVni.CpVniCount = varNiatelemetryNveVniWithoutEmbeddedStruct.CpVniCount
		varNiatelemetryNveVni.CpVniDown = varNiatelemetryNveVniWithoutEmbeddedStruct.CpVniDown
		varNiatelemetryNveVni.CpVniUp = varNiatelemetryNveVniWithoutEmbeddedStruct.CpVniUp
		varNiatelemetryNveVni.DpVniCount = varNiatelemetryNveVniWithoutEmbeddedStruct.DpVniCount
		varNiatelemetryNveVni.DpVniDown = varNiatelemetryNveVniWithoutEmbeddedStruct.DpVniDown
		varNiatelemetryNveVni.DpVniUp = varNiatelemetryNveVniWithoutEmbeddedStruct.DpVniUp
		*o = NiatelemetryNveVni(varNiatelemetryNveVni)
	} else {
		return err
	}

	varNiatelemetryNveVni := _NiatelemetryNveVni{}

	err = json.Unmarshal(bytes, &varNiatelemetryNveVni)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryNveVni.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CpVniCount")
		delete(additionalProperties, "CpVniDown")
		delete(additionalProperties, "CpVniUp")
		delete(additionalProperties, "DpVniCount")
		delete(additionalProperties, "DpVniDown")
		delete(additionalProperties, "DpVniUp")

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

type NullableNiatelemetryNveVni struct {
	value *NiatelemetryNveVni
	isSet bool
}

func (v NullableNiatelemetryNveVni) Get() *NiatelemetryNveVni {
	return v.value
}

func (v *NullableNiatelemetryNveVni) Set(val *NiatelemetryNveVni) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNveVni) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNveVni) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNveVni(val *NiatelemetryNveVni) *NullableNiatelemetryNveVni {
	return &NullableNiatelemetryNveVni{value: val, isSet: true}
}

func (v NullableNiatelemetryNveVni) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNveVni) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
