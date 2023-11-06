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
	"reflect"
	"strings"
)

// NiatelemetryProcSysCpu5min Aci node performance info in last 5 mintutes.
type NiatelemetryProcSysCpu5min struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the kernal average value.
	KernalAvg *string `json:"KernalAvg,omitempty"`
	// Returns the user average value.
	UserAvg              *string `json:"UserAvg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryProcSysCpu5min NiatelemetryProcSysCpu5min

// NewNiatelemetryProcSysCpu5min instantiates a new NiatelemetryProcSysCpu5min object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryProcSysCpu5min(classId string, objectType string) *NiatelemetryProcSysCpu5min {
	this := NiatelemetryProcSysCpu5min{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryProcSysCpu5minWithDefaults instantiates a new NiatelemetryProcSysCpu5min object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryProcSysCpu5minWithDefaults() *NiatelemetryProcSysCpu5min {
	this := NiatelemetryProcSysCpu5min{}
	var classId string = "niatelemetry.ProcSysCpu5min"
	this.ClassId = classId
	var objectType string = "niatelemetry.ProcSysCpu5min"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryProcSysCpu5min) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysCpu5min) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryProcSysCpu5min) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryProcSysCpu5min) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysCpu5min) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryProcSysCpu5min) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKernalAvg returns the KernalAvg field value if set, zero value otherwise.
func (o *NiatelemetryProcSysCpu5min) GetKernalAvg() string {
	if o == nil || o.KernalAvg == nil {
		var ret string
		return ret
	}
	return *o.KernalAvg
}

// GetKernalAvgOk returns a tuple with the KernalAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysCpu5min) GetKernalAvgOk() (*string, bool) {
	if o == nil || o.KernalAvg == nil {
		return nil, false
	}
	return o.KernalAvg, true
}

// HasKernalAvg returns a boolean if a field has been set.
func (o *NiatelemetryProcSysCpu5min) HasKernalAvg() bool {
	if o != nil && o.KernalAvg != nil {
		return true
	}

	return false
}

// SetKernalAvg gets a reference to the given string and assigns it to the KernalAvg field.
func (o *NiatelemetryProcSysCpu5min) SetKernalAvg(v string) {
	o.KernalAvg = &v
}

// GetUserAvg returns the UserAvg field value if set, zero value otherwise.
func (o *NiatelemetryProcSysCpu5min) GetUserAvg() string {
	if o == nil || o.UserAvg == nil {
		var ret string
		return ret
	}
	return *o.UserAvg
}

// GetUserAvgOk returns a tuple with the UserAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryProcSysCpu5min) GetUserAvgOk() (*string, bool) {
	if o == nil || o.UserAvg == nil {
		return nil, false
	}
	return o.UserAvg, true
}

// HasUserAvg returns a boolean if a field has been set.
func (o *NiatelemetryProcSysCpu5min) HasUserAvg() bool {
	if o != nil && o.UserAvg != nil {
		return true
	}

	return false
}

// SetUserAvg gets a reference to the given string and assigns it to the UserAvg field.
func (o *NiatelemetryProcSysCpu5min) SetUserAvg(v string) {
	o.UserAvg = &v
}

func (o NiatelemetryProcSysCpu5min) MarshalJSON() ([]byte, error) {
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
	if o.KernalAvg != nil {
		toSerialize["KernalAvg"] = o.KernalAvg
	}
	if o.UserAvg != nil {
		toSerialize["UserAvg"] = o.UserAvg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryProcSysCpu5min) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryProcSysCpu5minWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Returns the kernal average value.
		KernalAvg *string `json:"KernalAvg,omitempty"`
		// Returns the user average value.
		UserAvg *string `json:"UserAvg,omitempty"`
	}

	varNiatelemetryProcSysCpu5minWithoutEmbeddedStruct := NiatelemetryProcSysCpu5minWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryProcSysCpu5minWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryProcSysCpu5min := _NiatelemetryProcSysCpu5min{}
		varNiatelemetryProcSysCpu5min.ClassId = varNiatelemetryProcSysCpu5minWithoutEmbeddedStruct.ClassId
		varNiatelemetryProcSysCpu5min.ObjectType = varNiatelemetryProcSysCpu5minWithoutEmbeddedStruct.ObjectType
		varNiatelemetryProcSysCpu5min.KernalAvg = varNiatelemetryProcSysCpu5minWithoutEmbeddedStruct.KernalAvg
		varNiatelemetryProcSysCpu5min.UserAvg = varNiatelemetryProcSysCpu5minWithoutEmbeddedStruct.UserAvg
		*o = NiatelemetryProcSysCpu5min(varNiatelemetryProcSysCpu5min)
	} else {
		return err
	}

	varNiatelemetryProcSysCpu5min := _NiatelemetryProcSysCpu5min{}

	err = json.Unmarshal(bytes, &varNiatelemetryProcSysCpu5min)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryProcSysCpu5min.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KernalAvg")
		delete(additionalProperties, "UserAvg")

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

type NullableNiatelemetryProcSysCpu5min struct {
	value *NiatelemetryProcSysCpu5min
	isSet bool
}

func (v NullableNiatelemetryProcSysCpu5min) Get() *NiatelemetryProcSysCpu5min {
	return v.value
}

func (v *NullableNiatelemetryProcSysCpu5min) Set(val *NiatelemetryProcSysCpu5min) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryProcSysCpu5min) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryProcSysCpu5min) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryProcSysCpu5min(val *NiatelemetryProcSysCpu5min) *NullableNiatelemetryProcSysCpu5min {
	return &NullableNiatelemetryProcSysCpu5min{value: val, isSet: true}
}

func (v NullableNiatelemetryProcSysCpu5min) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryProcSysCpu5min) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
