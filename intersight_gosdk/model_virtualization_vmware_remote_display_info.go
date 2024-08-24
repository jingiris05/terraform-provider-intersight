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

// checks if the VirtualizationVmwareRemoteDisplayInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationVmwareRemoteDisplayInfo{}

// VirtualizationVmwareRemoteDisplayInfo Details of the remote display settings.
type VirtualizationVmwareRemoteDisplayInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The password used for remote access. It is stored base64 encoded.
	RemoteDisplayPassword *string `json:"RemoteDisplayPassword,omitempty"`
	// The access key for the remote display, potentially a long string.
	RemoteDisplayVncKey *string `json:"RemoteDisplayVncKey,omitempty"`
	// Applies only when remoteDisplayvnc is enabled.
	RemoteDisplayVncPort *int64 `json:"RemoteDisplayVncPort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareRemoteDisplayInfo VirtualizationVmwareRemoteDisplayInfo

// NewVirtualizationVmwareRemoteDisplayInfo instantiates a new VirtualizationVmwareRemoteDisplayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareRemoteDisplayInfo(classId string, objectType string) *VirtualizationVmwareRemoteDisplayInfo {
	this := VirtualizationVmwareRemoteDisplayInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareRemoteDisplayInfoWithDefaults instantiates a new VirtualizationVmwareRemoteDisplayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareRemoteDisplayInfoWithDefaults() *VirtualizationVmwareRemoteDisplayInfo {
	this := VirtualizationVmwareRemoteDisplayInfo{}
	var classId string = "virtualization.VmwareRemoteDisplayInfo"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareRemoteDisplayInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareRemoteDisplayInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareRemoteDisplayInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareRemoteDisplayInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareRemoteDisplayInfo" of the ClassId field.
func (o *VirtualizationVmwareRemoteDisplayInfo) GetDefaultClassId() interface{} {
	return "virtualization.VmwareRemoteDisplayInfo"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareRemoteDisplayInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareRemoteDisplayInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareRemoteDisplayInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareRemoteDisplayInfo" of the ObjectType field.
func (o *VirtualizationVmwareRemoteDisplayInfo) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareRemoteDisplayInfo"
}

// GetRemoteDisplayPassword returns the RemoteDisplayPassword field value if set, zero value otherwise.
func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayPassword() string {
	if o == nil || IsNil(o.RemoteDisplayPassword) {
		var ret string
		return ret
	}
	return *o.RemoteDisplayPassword
}

// GetRemoteDisplayPasswordOk returns a tuple with the RemoteDisplayPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteDisplayPassword) {
		return nil, false
	}
	return o.RemoteDisplayPassword, true
}

// HasRemoteDisplayPassword returns a boolean if a field has been set.
func (o *VirtualizationVmwareRemoteDisplayInfo) HasRemoteDisplayPassword() bool {
	if o != nil && !IsNil(o.RemoteDisplayPassword) {
		return true
	}

	return false
}

// SetRemoteDisplayPassword gets a reference to the given string and assigns it to the RemoteDisplayPassword field.
func (o *VirtualizationVmwareRemoteDisplayInfo) SetRemoteDisplayPassword(v string) {
	o.RemoteDisplayPassword = &v
}

// GetRemoteDisplayVncKey returns the RemoteDisplayVncKey field value if set, zero value otherwise.
func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayVncKey() string {
	if o == nil || IsNil(o.RemoteDisplayVncKey) {
		var ret string
		return ret
	}
	return *o.RemoteDisplayVncKey
}

// GetRemoteDisplayVncKeyOk returns a tuple with the RemoteDisplayVncKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayVncKeyOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteDisplayVncKey) {
		return nil, false
	}
	return o.RemoteDisplayVncKey, true
}

// HasRemoteDisplayVncKey returns a boolean if a field has been set.
func (o *VirtualizationVmwareRemoteDisplayInfo) HasRemoteDisplayVncKey() bool {
	if o != nil && !IsNil(o.RemoteDisplayVncKey) {
		return true
	}

	return false
}

// SetRemoteDisplayVncKey gets a reference to the given string and assigns it to the RemoteDisplayVncKey field.
func (o *VirtualizationVmwareRemoteDisplayInfo) SetRemoteDisplayVncKey(v string) {
	o.RemoteDisplayVncKey = &v
}

// GetRemoteDisplayVncPort returns the RemoteDisplayVncPort field value if set, zero value otherwise.
func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayVncPort() int64 {
	if o == nil || IsNil(o.RemoteDisplayVncPort) {
		var ret int64
		return ret
	}
	return *o.RemoteDisplayVncPort
}

// GetRemoteDisplayVncPortOk returns a tuple with the RemoteDisplayVncPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareRemoteDisplayInfo) GetRemoteDisplayVncPortOk() (*int64, bool) {
	if o == nil || IsNil(o.RemoteDisplayVncPort) {
		return nil, false
	}
	return o.RemoteDisplayVncPort, true
}

// HasRemoteDisplayVncPort returns a boolean if a field has been set.
func (o *VirtualizationVmwareRemoteDisplayInfo) HasRemoteDisplayVncPort() bool {
	if o != nil && !IsNil(o.RemoteDisplayVncPort) {
		return true
	}

	return false
}

// SetRemoteDisplayVncPort gets a reference to the given int64 and assigns it to the RemoteDisplayVncPort field.
func (o *VirtualizationVmwareRemoteDisplayInfo) SetRemoteDisplayVncPort(v int64) {
	o.RemoteDisplayVncPort = &v
}

func (o VirtualizationVmwareRemoteDisplayInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationVmwareRemoteDisplayInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RemoteDisplayPassword) {
		toSerialize["RemoteDisplayPassword"] = o.RemoteDisplayPassword
	}
	if !IsNil(o.RemoteDisplayVncKey) {
		toSerialize["RemoteDisplayVncKey"] = o.RemoteDisplayVncKey
	}
	if !IsNil(o.RemoteDisplayVncPort) {
		toSerialize["RemoteDisplayVncPort"] = o.RemoteDisplayVncPort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationVmwareRemoteDisplayInfo) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationVmwareRemoteDisplayInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The password used for remote access. It is stored base64 encoded.
		RemoteDisplayPassword *string `json:"RemoteDisplayPassword,omitempty"`
		// The access key for the remote display, potentially a long string.
		RemoteDisplayVncKey *string `json:"RemoteDisplayVncKey,omitempty"`
		// Applies only when remoteDisplayvnc is enabled.
		RemoteDisplayVncPort *int64 `json:"RemoteDisplayVncPort,omitempty"`
	}

	varVirtualizationVmwareRemoteDisplayInfoWithoutEmbeddedStruct := VirtualizationVmwareRemoteDisplayInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationVmwareRemoteDisplayInfoWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareRemoteDisplayInfo := _VirtualizationVmwareRemoteDisplayInfo{}
		varVirtualizationVmwareRemoteDisplayInfo.ClassId = varVirtualizationVmwareRemoteDisplayInfoWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareRemoteDisplayInfo.ObjectType = varVirtualizationVmwareRemoteDisplayInfoWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareRemoteDisplayInfo.RemoteDisplayPassword = varVirtualizationVmwareRemoteDisplayInfoWithoutEmbeddedStruct.RemoteDisplayPassword
		varVirtualizationVmwareRemoteDisplayInfo.RemoteDisplayVncKey = varVirtualizationVmwareRemoteDisplayInfoWithoutEmbeddedStruct.RemoteDisplayVncKey
		varVirtualizationVmwareRemoteDisplayInfo.RemoteDisplayVncPort = varVirtualizationVmwareRemoteDisplayInfoWithoutEmbeddedStruct.RemoteDisplayVncPort
		*o = VirtualizationVmwareRemoteDisplayInfo(varVirtualizationVmwareRemoteDisplayInfo)
	} else {
		return err
	}

	varVirtualizationVmwareRemoteDisplayInfo := _VirtualizationVmwareRemoteDisplayInfo{}

	err = json.Unmarshal(data, &varVirtualizationVmwareRemoteDisplayInfo)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVmwareRemoteDisplayInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RemoteDisplayPassword")
		delete(additionalProperties, "RemoteDisplayVncKey")
		delete(additionalProperties, "RemoteDisplayVncPort")

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

type NullableVirtualizationVmwareRemoteDisplayInfo struct {
	value *VirtualizationVmwareRemoteDisplayInfo
	isSet bool
}

func (v NullableVirtualizationVmwareRemoteDisplayInfo) Get() *VirtualizationVmwareRemoteDisplayInfo {
	return v.value
}

func (v *NullableVirtualizationVmwareRemoteDisplayInfo) Set(val *VirtualizationVmwareRemoteDisplayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareRemoteDisplayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareRemoteDisplayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareRemoteDisplayInfo(val *VirtualizationVmwareRemoteDisplayInfo) *NullableVirtualizationVmwareRemoteDisplayInfo {
	return &NullableVirtualizationVmwareRemoteDisplayInfo{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareRemoteDisplayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareRemoteDisplayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
