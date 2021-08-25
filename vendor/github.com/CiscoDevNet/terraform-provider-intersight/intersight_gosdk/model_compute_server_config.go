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

// ComputeServerConfig Configurable properties of servers common to both a server and a server profile. The user will apply these properties on a server directly when the server is not associated with a server profile and through a server profile when the server is attached with a server profile.
type ComputeServerConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// User defined asset tag of the server.
	AssetTag *string `json:"AssetTag,omitempty"`
	// User defined description of the server.
	UserLabel            *string `json:"UserLabel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeServerConfig ComputeServerConfig

// NewComputeServerConfig instantiates a new ComputeServerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeServerConfig(classId string, objectType string) *ComputeServerConfig {
	this := ComputeServerConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeServerConfigWithDefaults instantiates a new ComputeServerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeServerConfigWithDefaults() *ComputeServerConfig {
	this := ComputeServerConfig{}
	var classId string = "compute.ServerConfig"
	this.ClassId = classId
	var objectType string = "compute.ServerConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeServerConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeServerConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeServerConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeServerConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeServerConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeServerConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise.
func (o *ComputeServerConfig) GetAssetTag() string {
	if o == nil || o.AssetTag == nil {
		var ret string
		return ret
	}
	return *o.AssetTag
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerConfig) GetAssetTagOk() (*string, bool) {
	if o == nil || o.AssetTag == nil {
		return nil, false
	}
	return o.AssetTag, true
}

// HasAssetTag returns a boolean if a field has been set.
func (o *ComputeServerConfig) HasAssetTag() bool {
	if o != nil && o.AssetTag != nil {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given string and assigns it to the AssetTag field.
func (o *ComputeServerConfig) SetAssetTag(v string) {
	o.AssetTag = &v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ComputeServerConfig) GetUserLabel() string {
	if o == nil || o.UserLabel == nil {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerConfig) GetUserLabelOk() (*string, bool) {
	if o == nil || o.UserLabel == nil {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ComputeServerConfig) HasUserLabel() bool {
	if o != nil && o.UserLabel != nil {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ComputeServerConfig) SetUserLabel(v string) {
	o.UserLabel = &v
}

func (o ComputeServerConfig) MarshalJSON() ([]byte, error) {
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
	if o.AssetTag != nil {
		toSerialize["AssetTag"] = o.AssetTag
	}
	if o.UserLabel != nil {
		toSerialize["UserLabel"] = o.UserLabel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeServerConfig) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeServerConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// User defined asset tag of the server.
		AssetTag *string `json:"AssetTag,omitempty"`
		// User defined description of the server.
		UserLabel *string `json:"UserLabel,omitempty"`
	}

	varComputeServerConfigWithoutEmbeddedStruct := ComputeServerConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeServerConfigWithoutEmbeddedStruct)
	if err == nil {
		varComputeServerConfig := _ComputeServerConfig{}
		varComputeServerConfig.ClassId = varComputeServerConfigWithoutEmbeddedStruct.ClassId
		varComputeServerConfig.ObjectType = varComputeServerConfigWithoutEmbeddedStruct.ObjectType
		varComputeServerConfig.AssetTag = varComputeServerConfigWithoutEmbeddedStruct.AssetTag
		varComputeServerConfig.UserLabel = varComputeServerConfigWithoutEmbeddedStruct.UserLabel
		*o = ComputeServerConfig(varComputeServerConfig)
	} else {
		return err
	}

	varComputeServerConfig := _ComputeServerConfig{}

	err = json.Unmarshal(bytes, &varComputeServerConfig)
	if err == nil {
		o.MoBaseComplexType = varComputeServerConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssetTag")
		delete(additionalProperties, "UserLabel")

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

type NullableComputeServerConfig struct {
	value *ComputeServerConfig
	isSet bool
}

func (v NullableComputeServerConfig) Get() *ComputeServerConfig {
	return v.value
}

func (v *NullableComputeServerConfig) Set(val *ComputeServerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeServerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeServerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeServerConfig(val *ComputeServerConfig) *NullableComputeServerConfig {
	return &NullableComputeServerConfig{value: val, isSet: true}
}

func (v NullableComputeServerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeServerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
