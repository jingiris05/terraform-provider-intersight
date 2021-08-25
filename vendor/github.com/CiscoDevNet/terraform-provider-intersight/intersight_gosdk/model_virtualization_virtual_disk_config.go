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

// VirtualizationVirtualDiskConfig Virtual disk file information.
type VirtualizationVirtualDiskConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Disk capacity to be provided with units example - 10Gi.
	Capacity *string `json:"Capacity,omitempty"`
	// File mode of the disk, example - Filesystem, Block. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
	Mode *string `json:"Mode,omitempty"`
	// Base64 encoded CA certificates of the https source to check against.
	SourceCerts *string `json:"SourceCerts,omitempty"`
	// Source disk name from where the clone is done.
	SourceDiskToClone *string `json:"SourceDiskToClone,omitempty"`
	// Disk image source for the virtual machine.
	SourceFilePath       *string `json:"SourceFilePath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVirtualDiskConfig VirtualizationVirtualDiskConfig

// NewVirtualizationVirtualDiskConfig instantiates a new VirtualizationVirtualDiskConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVirtualDiskConfig(classId string, objectType string) *VirtualizationVirtualDiskConfig {
	this := VirtualizationVirtualDiskConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// NewVirtualizationVirtualDiskConfigWithDefaults instantiates a new VirtualizationVirtualDiskConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVirtualDiskConfigWithDefaults() *VirtualizationVirtualDiskConfig {
	this := VirtualizationVirtualDiskConfig{}
	var classId string = "virtualization.VirtualDiskConfig"
	this.ClassId = classId
	var objectType string = "virtualization.VirtualDiskConfig"
	this.ObjectType = objectType
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVirtualDiskConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVirtualDiskConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVirtualDiskConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVirtualDiskConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskConfig) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfig) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskConfig) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *VirtualizationVirtualDiskConfig) SetCapacity(v string) {
	o.Capacity = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskConfig) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfig) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskConfig) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VirtualizationVirtualDiskConfig) SetMode(v string) {
	o.Mode = &v
}

// GetSourceCerts returns the SourceCerts field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskConfig) GetSourceCerts() string {
	if o == nil || o.SourceCerts == nil {
		var ret string
		return ret
	}
	return *o.SourceCerts
}

// GetSourceCertsOk returns a tuple with the SourceCerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfig) GetSourceCertsOk() (*string, bool) {
	if o == nil || o.SourceCerts == nil {
		return nil, false
	}
	return o.SourceCerts, true
}

// HasSourceCerts returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskConfig) HasSourceCerts() bool {
	if o != nil && o.SourceCerts != nil {
		return true
	}

	return false
}

// SetSourceCerts gets a reference to the given string and assigns it to the SourceCerts field.
func (o *VirtualizationVirtualDiskConfig) SetSourceCerts(v string) {
	o.SourceCerts = &v
}

// GetSourceDiskToClone returns the SourceDiskToClone field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskConfig) GetSourceDiskToClone() string {
	if o == nil || o.SourceDiskToClone == nil {
		var ret string
		return ret
	}
	return *o.SourceDiskToClone
}

// GetSourceDiskToCloneOk returns a tuple with the SourceDiskToClone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfig) GetSourceDiskToCloneOk() (*string, bool) {
	if o == nil || o.SourceDiskToClone == nil {
		return nil, false
	}
	return o.SourceDiskToClone, true
}

// HasSourceDiskToClone returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskConfig) HasSourceDiskToClone() bool {
	if o != nil && o.SourceDiskToClone != nil {
		return true
	}

	return false
}

// SetSourceDiskToClone gets a reference to the given string and assigns it to the SourceDiskToClone field.
func (o *VirtualizationVirtualDiskConfig) SetSourceDiskToClone(v string) {
	o.SourceDiskToClone = &v
}

// GetSourceFilePath returns the SourceFilePath field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskConfig) GetSourceFilePath() string {
	if o == nil || o.SourceFilePath == nil {
		var ret string
		return ret
	}
	return *o.SourceFilePath
}

// GetSourceFilePathOk returns a tuple with the SourceFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskConfig) GetSourceFilePathOk() (*string, bool) {
	if o == nil || o.SourceFilePath == nil {
		return nil, false
	}
	return o.SourceFilePath, true
}

// HasSourceFilePath returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskConfig) HasSourceFilePath() bool {
	if o != nil && o.SourceFilePath != nil {
		return true
	}

	return false
}

// SetSourceFilePath gets a reference to the given string and assigns it to the SourceFilePath field.
func (o *VirtualizationVirtualDiskConfig) SetSourceFilePath(v string) {
	o.SourceFilePath = &v
}

func (o VirtualizationVirtualDiskConfig) MarshalJSON() ([]byte, error) {
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
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.SourceCerts != nil {
		toSerialize["SourceCerts"] = o.SourceCerts
	}
	if o.SourceDiskToClone != nil {
		toSerialize["SourceDiskToClone"] = o.SourceDiskToClone
	}
	if o.SourceFilePath != nil {
		toSerialize["SourceFilePath"] = o.SourceFilePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVirtualDiskConfig) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVirtualDiskConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Disk capacity to be provided with units example - 10Gi.
		Capacity *string `json:"Capacity,omitempty"`
		// File mode of the disk, example - Filesystem, Block. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
		Mode *string `json:"Mode,omitempty"`
		// Base64 encoded CA certificates of the https source to check against.
		SourceCerts *string `json:"SourceCerts,omitempty"`
		// Source disk name from where the clone is done.
		SourceDiskToClone *string `json:"SourceDiskToClone,omitempty"`
		// Disk image source for the virtual machine.
		SourceFilePath *string `json:"SourceFilePath,omitempty"`
	}

	varVirtualizationVirtualDiskConfigWithoutEmbeddedStruct := VirtualizationVirtualDiskConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVirtualDiskConfigWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVirtualDiskConfig := _VirtualizationVirtualDiskConfig{}
		varVirtualizationVirtualDiskConfig.ClassId = varVirtualizationVirtualDiskConfigWithoutEmbeddedStruct.ClassId
		varVirtualizationVirtualDiskConfig.ObjectType = varVirtualizationVirtualDiskConfigWithoutEmbeddedStruct.ObjectType
		varVirtualizationVirtualDiskConfig.Capacity = varVirtualizationVirtualDiskConfigWithoutEmbeddedStruct.Capacity
		varVirtualizationVirtualDiskConfig.Mode = varVirtualizationVirtualDiskConfigWithoutEmbeddedStruct.Mode
		varVirtualizationVirtualDiskConfig.SourceCerts = varVirtualizationVirtualDiskConfigWithoutEmbeddedStruct.SourceCerts
		varVirtualizationVirtualDiskConfig.SourceDiskToClone = varVirtualizationVirtualDiskConfigWithoutEmbeddedStruct.SourceDiskToClone
		varVirtualizationVirtualDiskConfig.SourceFilePath = varVirtualizationVirtualDiskConfigWithoutEmbeddedStruct.SourceFilePath
		*o = VirtualizationVirtualDiskConfig(varVirtualizationVirtualDiskConfig)
	} else {
		return err
	}

	varVirtualizationVirtualDiskConfig := _VirtualizationVirtualDiskConfig{}

	err = json.Unmarshal(bytes, &varVirtualizationVirtualDiskConfig)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVirtualDiskConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "SourceCerts")
		delete(additionalProperties, "SourceDiskToClone")
		delete(additionalProperties, "SourceFilePath")

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

type NullableVirtualizationVirtualDiskConfig struct {
	value *VirtualizationVirtualDiskConfig
	isSet bool
}

func (v NullableVirtualizationVirtualDiskConfig) Get() *VirtualizationVirtualDiskConfig {
	return v.value
}

func (v *NullableVirtualizationVirtualDiskConfig) Set(val *VirtualizationVirtualDiskConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVirtualDiskConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVirtualDiskConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVirtualDiskConfig(val *VirtualizationVirtualDiskConfig) *NullableVirtualizationVirtualDiskConfig {
	return &NullableVirtualizationVirtualDiskConfig{value: val, isSet: true}
}

func (v NullableVirtualizationVirtualDiskConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVirtualDiskConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
