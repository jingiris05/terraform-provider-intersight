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
	"time"
)

// checks if the ApplianceMetadataManifestVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceMetadataManifestVersion{}

// ApplianceMetadataManifestVersion BucketVersion is a 3 tuple (bucketName, timestamp of when the metamanifest was touched, checksum). Uniquely identifies a single metadata update to a bucket.
type ApplianceMetadataManifestVersion struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Metamanifest file checksum.
	FileSha *string `json:"FileSha,omitempty"`
	// The timestamp when the metamanifest was touched.
	FileTime *time.Time `json:"FileTime,omitempty"`
	// Name of the bucket that is being monitored. * `hcl-meta` - Hcl bucket, metadata update will be automatically enabled. * `advisories` - Advisory bucket, metadata update will be automatically enabled. * `onprem-images` - Onprem images bucket, metadata update will be automatically enableds.
	MetadataType         *string `json:"MetadataType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceMetadataManifestVersion ApplianceMetadataManifestVersion

// NewApplianceMetadataManifestVersion instantiates a new ApplianceMetadataManifestVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceMetadataManifestVersion(classId string, objectType string) *ApplianceMetadataManifestVersion {
	this := ApplianceMetadataManifestVersion{}
	this.ClassId = classId
	this.ObjectType = objectType
	var metadataType string = "hcl-meta"
	this.MetadataType = &metadataType
	return &this
}

// NewApplianceMetadataManifestVersionWithDefaults instantiates a new ApplianceMetadataManifestVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceMetadataManifestVersionWithDefaults() *ApplianceMetadataManifestVersion {
	this := ApplianceMetadataManifestVersion{}
	var classId string = "appliance.MetadataManifestVersion"
	this.ClassId = classId
	var objectType string = "appliance.MetadataManifestVersion"
	this.ObjectType = objectType
	var metadataType string = "hcl-meta"
	this.MetadataType = &metadataType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceMetadataManifestVersion) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceMetadataManifestVersion) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceMetadataManifestVersion) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.MetadataManifestVersion" of the ClassId field.
func (o *ApplianceMetadataManifestVersion) GetDefaultClassId() interface{} {
	return "appliance.MetadataManifestVersion"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceMetadataManifestVersion) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceMetadataManifestVersion) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceMetadataManifestVersion) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.MetadataManifestVersion" of the ObjectType field.
func (o *ApplianceMetadataManifestVersion) GetDefaultObjectType() interface{} {
	return "appliance.MetadataManifestVersion"
}

// GetFileSha returns the FileSha field value if set, zero value otherwise.
func (o *ApplianceMetadataManifestVersion) GetFileSha() string {
	if o == nil || IsNil(o.FileSha) {
		var ret string
		return ret
	}
	return *o.FileSha
}

// GetFileShaOk returns a tuple with the FileSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetadataManifestVersion) GetFileShaOk() (*string, bool) {
	if o == nil || IsNil(o.FileSha) {
		return nil, false
	}
	return o.FileSha, true
}

// HasFileSha returns a boolean if a field has been set.
func (o *ApplianceMetadataManifestVersion) HasFileSha() bool {
	if o != nil && !IsNil(o.FileSha) {
		return true
	}

	return false
}

// SetFileSha gets a reference to the given string and assigns it to the FileSha field.
func (o *ApplianceMetadataManifestVersion) SetFileSha(v string) {
	o.FileSha = &v
}

// GetFileTime returns the FileTime field value if set, zero value otherwise.
func (o *ApplianceMetadataManifestVersion) GetFileTime() time.Time {
	if o == nil || IsNil(o.FileTime) {
		var ret time.Time
		return ret
	}
	return *o.FileTime
}

// GetFileTimeOk returns a tuple with the FileTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetadataManifestVersion) GetFileTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FileTime) {
		return nil, false
	}
	return o.FileTime, true
}

// HasFileTime returns a boolean if a field has been set.
func (o *ApplianceMetadataManifestVersion) HasFileTime() bool {
	if o != nil && !IsNil(o.FileTime) {
		return true
	}

	return false
}

// SetFileTime gets a reference to the given time.Time and assigns it to the FileTime field.
func (o *ApplianceMetadataManifestVersion) SetFileTime(v time.Time) {
	o.FileTime = &v
}

// GetMetadataType returns the MetadataType field value if set, zero value otherwise.
func (o *ApplianceMetadataManifestVersion) GetMetadataType() string {
	if o == nil || IsNil(o.MetadataType) {
		var ret string
		return ret
	}
	return *o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetadataManifestVersion) GetMetadataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataType) {
		return nil, false
	}
	return o.MetadataType, true
}

// HasMetadataType returns a boolean if a field has been set.
func (o *ApplianceMetadataManifestVersion) HasMetadataType() bool {
	if o != nil && !IsNil(o.MetadataType) {
		return true
	}

	return false
}

// SetMetadataType gets a reference to the given string and assigns it to the MetadataType field.
func (o *ApplianceMetadataManifestVersion) SetMetadataType(v string) {
	o.MetadataType = &v
}

func (o ApplianceMetadataManifestVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceMetadataManifestVersion) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FileSha) {
		toSerialize["FileSha"] = o.FileSha
	}
	if !IsNil(o.FileTime) {
		toSerialize["FileTime"] = o.FileTime
	}
	if !IsNil(o.MetadataType) {
		toSerialize["MetadataType"] = o.MetadataType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceMetadataManifestVersion) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceMetadataManifestVersionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Metamanifest file checksum.
		FileSha *string `json:"FileSha,omitempty"`
		// The timestamp when the metamanifest was touched.
		FileTime *time.Time `json:"FileTime,omitempty"`
		// Name of the bucket that is being monitored. * `hcl-meta` - Hcl bucket, metadata update will be automatically enabled. * `advisories` - Advisory bucket, metadata update will be automatically enabled. * `onprem-images` - Onprem images bucket, metadata update will be automatically enableds.
		MetadataType *string `json:"MetadataType,omitempty"`
	}

	varApplianceMetadataManifestVersionWithoutEmbeddedStruct := ApplianceMetadataManifestVersionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceMetadataManifestVersionWithoutEmbeddedStruct)
	if err == nil {
		varApplianceMetadataManifestVersion := _ApplianceMetadataManifestVersion{}
		varApplianceMetadataManifestVersion.ClassId = varApplianceMetadataManifestVersionWithoutEmbeddedStruct.ClassId
		varApplianceMetadataManifestVersion.ObjectType = varApplianceMetadataManifestVersionWithoutEmbeddedStruct.ObjectType
		varApplianceMetadataManifestVersion.FileSha = varApplianceMetadataManifestVersionWithoutEmbeddedStruct.FileSha
		varApplianceMetadataManifestVersion.FileTime = varApplianceMetadataManifestVersionWithoutEmbeddedStruct.FileTime
		varApplianceMetadataManifestVersion.MetadataType = varApplianceMetadataManifestVersionWithoutEmbeddedStruct.MetadataType
		*o = ApplianceMetadataManifestVersion(varApplianceMetadataManifestVersion)
	} else {
		return err
	}

	varApplianceMetadataManifestVersion := _ApplianceMetadataManifestVersion{}

	err = json.Unmarshal(data, &varApplianceMetadataManifestVersion)
	if err == nil {
		o.MoBaseComplexType = varApplianceMetadataManifestVersion.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileSha")
		delete(additionalProperties, "FileTime")
		delete(additionalProperties, "MetadataType")

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

type NullableApplianceMetadataManifestVersion struct {
	value *ApplianceMetadataManifestVersion
	isSet bool
}

func (v NullableApplianceMetadataManifestVersion) Get() *ApplianceMetadataManifestVersion {
	return v.value
}

func (v *NullableApplianceMetadataManifestVersion) Set(val *ApplianceMetadataManifestVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceMetadataManifestVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceMetadataManifestVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceMetadataManifestVersion(val *ApplianceMetadataManifestVersion) *NullableApplianceMetadataManifestVersion {
	return &NullableApplianceMetadataManifestVersion{value: val, isSet: true}
}

func (v NullableApplianceMetadataManifestVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceMetadataManifestVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
