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

// NiaapiReleaseRecommend This contains the recommend version based on the hardware.
type NiaapiReleaseRecommend struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Current long-lived release.
	Cll *string `json:"Cll,omitempty"`
	// Customer recommended releases.
	Crr *string `json:"Crr,omitempty"`
	// Hardware model identificator.
	Pid *string `json:"Pid,omitempty"`
	// Upcoming long-lived release.
	Ull                  *string `json:"Ull,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiReleaseRecommend NiaapiReleaseRecommend

// NewNiaapiReleaseRecommend instantiates a new NiaapiReleaseRecommend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiReleaseRecommend(classId string, objectType string) *NiaapiReleaseRecommend {
	this := NiaapiReleaseRecommend{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiReleaseRecommendWithDefaults instantiates a new NiaapiReleaseRecommend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiReleaseRecommendWithDefaults() *NiaapiReleaseRecommend {
	this := NiaapiReleaseRecommend{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiReleaseRecommend) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiReleaseRecommend) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiReleaseRecommend) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiReleaseRecommend) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiReleaseRecommend) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiReleaseRecommend) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCll returns the Cll field value if set, zero value otherwise.
func (o *NiaapiReleaseRecommend) GetCll() string {
	if o == nil || o.Cll == nil {
		var ret string
		return ret
	}
	return *o.Cll
}

// GetCllOk returns a tuple with the Cll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiReleaseRecommend) GetCllOk() (*string, bool) {
	if o == nil || o.Cll == nil {
		return nil, false
	}
	return o.Cll, true
}

// HasCll returns a boolean if a field has been set.
func (o *NiaapiReleaseRecommend) HasCll() bool {
	if o != nil && o.Cll != nil {
		return true
	}

	return false
}

// SetCll gets a reference to the given string and assigns it to the Cll field.
func (o *NiaapiReleaseRecommend) SetCll(v string) {
	o.Cll = &v
}

// GetCrr returns the Crr field value if set, zero value otherwise.
func (o *NiaapiReleaseRecommend) GetCrr() string {
	if o == nil || o.Crr == nil {
		var ret string
		return ret
	}
	return *o.Crr
}

// GetCrrOk returns a tuple with the Crr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiReleaseRecommend) GetCrrOk() (*string, bool) {
	if o == nil || o.Crr == nil {
		return nil, false
	}
	return o.Crr, true
}

// HasCrr returns a boolean if a field has been set.
func (o *NiaapiReleaseRecommend) HasCrr() bool {
	if o != nil && o.Crr != nil {
		return true
	}

	return false
}

// SetCrr gets a reference to the given string and assigns it to the Crr field.
func (o *NiaapiReleaseRecommend) SetCrr(v string) {
	o.Crr = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *NiaapiReleaseRecommend) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiReleaseRecommend) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *NiaapiReleaseRecommend) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *NiaapiReleaseRecommend) SetPid(v string) {
	o.Pid = &v
}

// GetUll returns the Ull field value if set, zero value otherwise.
func (o *NiaapiReleaseRecommend) GetUll() string {
	if o == nil || o.Ull == nil {
		var ret string
		return ret
	}
	return *o.Ull
}

// GetUllOk returns a tuple with the Ull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiReleaseRecommend) GetUllOk() (*string, bool) {
	if o == nil || o.Ull == nil {
		return nil, false
	}
	return o.Ull, true
}

// HasUll returns a boolean if a field has been set.
func (o *NiaapiReleaseRecommend) HasUll() bool {
	if o != nil && o.Ull != nil {
		return true
	}

	return false
}

// SetUll gets a reference to the given string and assigns it to the Ull field.
func (o *NiaapiReleaseRecommend) SetUll(v string) {
	o.Ull = &v
}

func (o NiaapiReleaseRecommend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cll != nil {
		toSerialize["Cll"] = o.Cll
	}
	if o.Crr != nil {
		toSerialize["Crr"] = o.Crr
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.Ull != nil {
		toSerialize["Ull"] = o.Ull
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiReleaseRecommend) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiReleaseRecommendWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Current long-lived release.
		Cll *string `json:"Cll,omitempty"`
		// Customer recommended releases.
		Crr *string `json:"Crr,omitempty"`
		// Hardware model identificator.
		Pid *string `json:"Pid,omitempty"`
		// Upcoming long-lived release.
		Ull *string `json:"Ull,omitempty"`
	}

	varNiaapiReleaseRecommendWithoutEmbeddedStruct := NiaapiReleaseRecommendWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiReleaseRecommendWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiReleaseRecommend := _NiaapiReleaseRecommend{}
		varNiaapiReleaseRecommend.ClassId = varNiaapiReleaseRecommendWithoutEmbeddedStruct.ClassId
		varNiaapiReleaseRecommend.ObjectType = varNiaapiReleaseRecommendWithoutEmbeddedStruct.ObjectType
		varNiaapiReleaseRecommend.Cll = varNiaapiReleaseRecommendWithoutEmbeddedStruct.Cll
		varNiaapiReleaseRecommend.Crr = varNiaapiReleaseRecommendWithoutEmbeddedStruct.Crr
		varNiaapiReleaseRecommend.Pid = varNiaapiReleaseRecommendWithoutEmbeddedStruct.Pid
		varNiaapiReleaseRecommend.Ull = varNiaapiReleaseRecommendWithoutEmbeddedStruct.Ull
		*o = NiaapiReleaseRecommend(varNiaapiReleaseRecommend)
	} else {
		return err
	}

	varNiaapiReleaseRecommend := _NiaapiReleaseRecommend{}

	err = json.Unmarshal(bytes, &varNiaapiReleaseRecommend)
	if err == nil {
		o.MoBaseMo = varNiaapiReleaseRecommend.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cll")
		delete(additionalProperties, "Crr")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "Ull")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableNiaapiReleaseRecommend struct {
	value *NiaapiReleaseRecommend
	isSet bool
}

func (v NullableNiaapiReleaseRecommend) Get() *NiaapiReleaseRecommend {
	return v.value
}

func (v *NullableNiaapiReleaseRecommend) Set(val *NiaapiReleaseRecommend) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiReleaseRecommend) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiReleaseRecommend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiReleaseRecommend(val *NiaapiReleaseRecommend) *NullableNiaapiReleaseRecommend {
	return &NullableNiaapiReleaseRecommend{value: val, isSet: true}
}

func (v NullableNiaapiReleaseRecommend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiReleaseRecommend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
