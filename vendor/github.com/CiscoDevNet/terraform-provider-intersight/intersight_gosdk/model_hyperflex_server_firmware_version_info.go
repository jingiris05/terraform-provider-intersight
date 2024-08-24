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

// checks if the HyperflexServerFirmwareVersionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexServerFirmwareVersionInfo{}

// HyperflexServerFirmwareVersionInfo The firmware version details for UCS servers.
type HyperflexServerFirmwareVersionInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The platform type for UCS server. * `M5` - M5 generation of UCS server. * `M3` - M3 generation of UCS server. * `M4` - M4 generation of UCS server. * `M6` - M6 generation of UCS server.
	ServerPlatform *string `json:"ServerPlatform,omitempty"`
	// The server firmware bundle version.
	Version              *string `json:"Version,omitempty" validate:"regexp=(^3\\\\.[1-9]\\\\([1-9][a-z]\\\\)$|^[4-9]\\\\.[0-9]\\\\([1-9][a-z]\\\\)$|^[4-9]\\\\.[0-9]\\\\([0-9](\\\\.[0-9]+|[a-z]).*\\\\)$)"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexServerFirmwareVersionInfo HyperflexServerFirmwareVersionInfo

// NewHyperflexServerFirmwareVersionInfo instantiates a new HyperflexServerFirmwareVersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexServerFirmwareVersionInfo(classId string, objectType string) *HyperflexServerFirmwareVersionInfo {
	this := HyperflexServerFirmwareVersionInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var serverPlatform string = "M5"
	this.ServerPlatform = &serverPlatform
	return &this
}

// NewHyperflexServerFirmwareVersionInfoWithDefaults instantiates a new HyperflexServerFirmwareVersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexServerFirmwareVersionInfoWithDefaults() *HyperflexServerFirmwareVersionInfo {
	this := HyperflexServerFirmwareVersionInfo{}
	var classId string = "hyperflex.ServerFirmwareVersionInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.ServerFirmwareVersionInfo"
	this.ObjectType = objectType
	var serverPlatform string = "M5"
	this.ServerPlatform = &serverPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexServerFirmwareVersionInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexServerFirmwareVersionInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ServerFirmwareVersionInfo" of the ClassId field.
func (o *HyperflexServerFirmwareVersionInfo) GetDefaultClassId() interface{} {
	return "hyperflex.ServerFirmwareVersionInfo"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexServerFirmwareVersionInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexServerFirmwareVersionInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ServerFirmwareVersionInfo" of the ObjectType field.
func (o *HyperflexServerFirmwareVersionInfo) GetDefaultObjectType() interface{} {
	return "hyperflex.ServerFirmwareVersionInfo"
}

// GetServerPlatform returns the ServerPlatform field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionInfo) GetServerPlatform() string {
	if o == nil || IsNil(o.ServerPlatform) {
		var ret string
		return ret
	}
	return *o.ServerPlatform
}

// GetServerPlatformOk returns a tuple with the ServerPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfo) GetServerPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.ServerPlatform) {
		return nil, false
	}
	return o.ServerPlatform, true
}

// HasServerPlatform returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionInfo) HasServerPlatform() bool {
	if o != nil && !IsNil(o.ServerPlatform) {
		return true
	}

	return false
}

// SetServerPlatform gets a reference to the given string and assigns it to the ServerPlatform field.
func (o *HyperflexServerFirmwareVersionInfo) SetServerPlatform(v string) {
	o.ServerPlatform = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexServerFirmwareVersionInfo) SetVersion(v string) {
	o.Version = &v
}

func (o HyperflexServerFirmwareVersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexServerFirmwareVersionInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ServerPlatform) {
		toSerialize["ServerPlatform"] = o.ServerPlatform
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexServerFirmwareVersionInfo) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The platform type for UCS server. * `M5` - M5 generation of UCS server. * `M3` - M3 generation of UCS server. * `M4` - M4 generation of UCS server. * `M6` - M6 generation of UCS server.
		ServerPlatform *string `json:"ServerPlatform,omitempty"`
		// The server firmware bundle version.
		Version *string `json:"Version,omitempty" validate:"regexp=(^3\\\\.[1-9]\\\\([1-9][a-z]\\\\)$|^[4-9]\\\\.[0-9]\\\\([1-9][a-z]\\\\)$|^[4-9]\\\\.[0-9]\\\\([0-9](\\\\.[0-9]+|[a-z]).*\\\\)$)"`
	}

	varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct := HyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexServerFirmwareVersionInfo := _HyperflexServerFirmwareVersionInfo{}
		varHyperflexServerFirmwareVersionInfo.ClassId = varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct.ClassId
		varHyperflexServerFirmwareVersionInfo.ObjectType = varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct.ObjectType
		varHyperflexServerFirmwareVersionInfo.ServerPlatform = varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct.ServerPlatform
		varHyperflexServerFirmwareVersionInfo.Version = varHyperflexServerFirmwareVersionInfoWithoutEmbeddedStruct.Version
		*o = HyperflexServerFirmwareVersionInfo(varHyperflexServerFirmwareVersionInfo)
	} else {
		return err
	}

	varHyperflexServerFirmwareVersionInfo := _HyperflexServerFirmwareVersionInfo{}

	err = json.Unmarshal(data, &varHyperflexServerFirmwareVersionInfo)
	if err == nil {
		o.MoBaseComplexType = varHyperflexServerFirmwareVersionInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ServerPlatform")
		delete(additionalProperties, "Version")

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

type NullableHyperflexServerFirmwareVersionInfo struct {
	value *HyperflexServerFirmwareVersionInfo
	isSet bool
}

func (v NullableHyperflexServerFirmwareVersionInfo) Get() *HyperflexServerFirmwareVersionInfo {
	return v.value
}

func (v *NullableHyperflexServerFirmwareVersionInfo) Set(val *HyperflexServerFirmwareVersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerFirmwareVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerFirmwareVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerFirmwareVersionInfo(val *HyperflexServerFirmwareVersionInfo) *NullableHyperflexServerFirmwareVersionInfo {
	return &NullableHyperflexServerFirmwareVersionInfo{value: val, isSet: true}
}

func (v NullableHyperflexServerFirmwareVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerFirmwareVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
