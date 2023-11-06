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

// NiatelemetryProcSysMem15minAllOf Definition of the list of properties defined in 'niatelemetry.ProcSysMem15min', excluding properties defined in parent classes.
type NiatelemetryProcSysMem15minAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the free average value.
	FreeAvg *string `json:"FreeAvg,omitempty"`
	// Returns the total average value.
	TotalAvg             *string `json:"TotalAvg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryProcSysMem15minAllOf NiatelemetryProcSysMem15minAllOf

// NewNiatelemetryProcSysMem15minAllOf instantiates a new NiatelemetryProcSysMem15minAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryProcSysMem15minAllOf(classId string, objectType string) *NiatelemetryProcSysMem15minAllOf {
	this := NiatelemetryProcSysMem15minAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryProcSysMem15minAllOfWithDefaults instantiates a new NiatelemetryProcSysMem15minAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryProcSysMem15minAllOfWithDefaults() *NiatelemetryProcSysMem15minAllOf {
	this := NiatelemetryProcSysMem15minAllOf{}
	var classId string = "niatelemetry.ProcSysMem15min"
	this.ClassId = classId
	var objectType string = "niatelemetry.ProcSysMem15min"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryProcSysMem15minAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysMem15minAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryProcSysMem15minAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryProcSysMem15minAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysMem15minAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryProcSysMem15minAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFreeAvg returns the FreeAvg field value if set, zero value otherwise.
func (o *NiatelemetryProcSysMem15minAllOf) GetFreeAvg() string {
	if o == nil || o.FreeAvg == nil {
		var ret string
		return ret
	}
	return *o.FreeAvg
}

// GetFreeAvgOk returns a tuple with the FreeAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysMem15minAllOf) GetFreeAvgOk() (*string, bool) {
	if o == nil || o.FreeAvg == nil {
		return nil, false
	}
	return o.FreeAvg, true
}

// HasFreeAvg returns a boolean if a field has been set.
func (o *NiatelemetryProcSysMem15minAllOf) HasFreeAvg() bool {
	if o != nil && o.FreeAvg != nil {
		return true
	}

	return false
}

// SetFreeAvg gets a reference to the given string and assigns it to the FreeAvg field.
func (o *NiatelemetryProcSysMem15minAllOf) SetFreeAvg(v string) {
	o.FreeAvg = &v
}

// GetTotalAvg returns the TotalAvg field value if set, zero value otherwise.
func (o *NiatelemetryProcSysMem15minAllOf) GetTotalAvg() string {
	if o == nil || o.TotalAvg == nil {
		var ret string
		return ret
	}
	return *o.TotalAvg
}

// GetTotalAvgOk returns a tuple with the TotalAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysMem15minAllOf) GetTotalAvgOk() (*string, bool) {
	if o == nil || o.TotalAvg == nil {
		return nil, false
	}
	return o.TotalAvg, true
}

// HasTotalAvg returns a boolean if a field has been set.
func (o *NiatelemetryProcSysMem15minAllOf) HasTotalAvg() bool {
	if o != nil && o.TotalAvg != nil {
		return true
	}

	return false
}

// SetTotalAvg gets a reference to the given string and assigns it to the TotalAvg field.
func (o *NiatelemetryProcSysMem15minAllOf) SetTotalAvg(v string) {
	o.TotalAvg = &v
}

func (o NiatelemetryProcSysMem15minAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FreeAvg != nil {
		toSerialize["FreeAvg"] = o.FreeAvg
	}
	if o.TotalAvg != nil {
		toSerialize["TotalAvg"] = o.TotalAvg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryProcSysMem15minAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryProcSysMem15minAllOf := _NiatelemetryProcSysMem15minAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryProcSysMem15minAllOf); err == nil {
		*o = NiatelemetryProcSysMem15minAllOf(varNiatelemetryProcSysMem15minAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FreeAvg")
		delete(additionalProperties, "TotalAvg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryProcSysMem15minAllOf struct {
	value *NiatelemetryProcSysMem15minAllOf
	isSet bool
}

func (v NullableNiatelemetryProcSysMem15minAllOf) Get() *NiatelemetryProcSysMem15minAllOf {
	return v.value
}

func (v *NullableNiatelemetryProcSysMem15minAllOf) Set(val *NiatelemetryProcSysMem15minAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryProcSysMem15minAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryProcSysMem15minAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryProcSysMem15minAllOf(val *NiatelemetryProcSysMem15minAllOf) *NullableNiatelemetryProcSysMem15minAllOf {
	return &NullableNiatelemetryProcSysMem15minAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryProcSysMem15minAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryProcSysMem15minAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
