/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// AssetWorkloadOptimizerDynatraceOptions Captures configuration specific to the Dynatrace target for the Workload Optimizer service.
type AssetWorkloadOptimizerDynatraceOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Each environment monitored with Dynatrace is identified with a unique character string—the environment ID. The Dynatrace API relies heavily on environment IDs to ensure that it pulls monitoring data from and pushes relevant external events to the correct Dynatrace environments.
	EnvironmentId        *string `json:"EnvironmentId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerDynatraceOptions AssetWorkloadOptimizerDynatraceOptions

// NewAssetWorkloadOptimizerDynatraceOptions instantiates a new AssetWorkloadOptimizerDynatraceOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerDynatraceOptions(classId string, objectType string) *AssetWorkloadOptimizerDynatraceOptions {
	this := AssetWorkloadOptimizerDynatraceOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerDynatraceOptionsWithDefaults instantiates a new AssetWorkloadOptimizerDynatraceOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerDynatraceOptionsWithDefaults() *AssetWorkloadOptimizerDynatraceOptions {
	this := AssetWorkloadOptimizerDynatraceOptions{}
	var classId string = "asset.WorkloadOptimizerDynatraceOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerDynatraceOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerDynatraceOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerDynatraceOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerDynatraceOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerDynatraceOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerDynatraceOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerDynatraceOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerDynatraceOptions) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerDynatraceOptions) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerDynatraceOptions) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId != nil {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *AssetWorkloadOptimizerDynatraceOptions) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

func (o AssetWorkloadOptimizerDynatraceOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnvironmentId != nil {
		toSerialize["EnvironmentId"] = o.EnvironmentId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerDynatraceOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetWorkloadOptimizerDynatraceOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Each environment monitored with Dynatrace is identified with a unique character string—the environment ID. The Dynatrace API relies heavily on environment IDs to ensure that it pulls monitoring data from and pushes relevant external events to the correct Dynatrace environments.
		EnvironmentId *string `json:"EnvironmentId,omitempty"`
	}

	varAssetWorkloadOptimizerDynatraceOptionsWithoutEmbeddedStruct := AssetWorkloadOptimizerDynatraceOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerDynatraceOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetWorkloadOptimizerDynatraceOptions := _AssetWorkloadOptimizerDynatraceOptions{}
		varAssetWorkloadOptimizerDynatraceOptions.ClassId = varAssetWorkloadOptimizerDynatraceOptionsWithoutEmbeddedStruct.ClassId
		varAssetWorkloadOptimizerDynatraceOptions.ObjectType = varAssetWorkloadOptimizerDynatraceOptionsWithoutEmbeddedStruct.ObjectType
		varAssetWorkloadOptimizerDynatraceOptions.EnvironmentId = varAssetWorkloadOptimizerDynatraceOptionsWithoutEmbeddedStruct.EnvironmentId
		*o = AssetWorkloadOptimizerDynatraceOptions(varAssetWorkloadOptimizerDynatraceOptions)
	} else {
		return err
	}

	varAssetWorkloadOptimizerDynatraceOptions := _AssetWorkloadOptimizerDynatraceOptions{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerDynatraceOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetWorkloadOptimizerDynatraceOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnvironmentId")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetWorkloadOptimizerDynatraceOptions struct {
	value *AssetWorkloadOptimizerDynatraceOptions
	isSet bool
}

func (v NullableAssetWorkloadOptimizerDynatraceOptions) Get() *AssetWorkloadOptimizerDynatraceOptions {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerDynatraceOptions) Set(val *AssetWorkloadOptimizerDynatraceOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerDynatraceOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerDynatraceOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerDynatraceOptions(val *AssetWorkloadOptimizerDynatraceOptions) *NullableAssetWorkloadOptimizerDynatraceOptions {
	return &NullableAssetWorkloadOptimizerDynatraceOptions{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerDynatraceOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerDynatraceOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
