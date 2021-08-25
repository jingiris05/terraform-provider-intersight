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

// FirmwareDistributable An image distributed by Cisco.
type FirmwareDistributable struct {
	FirmwareBaseDistributable
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The file location of the distributable.
	FileLocation *string `json:"FileLocation,omitempty"`
	// The category into which the distributable falls into according to the supported platform series. For e.g.; C-Series/B-Series/Infrastructure.
	ImageCategory *string `json:"ImageCategory,omitempty"`
	// The source of the distributable. If it has been created by the user or system. * `System` - The distributable has been created by the System. * `User` - The distributable has been created by the User.
	Origin               *string                                `json:"Origin,omitempty"`
	Catalog              *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareDistributable FirmwareDistributable

// NewFirmwareDistributable instantiates a new FirmwareDistributable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareDistributable(classId string, objectType string) *FirmwareDistributable {
	this := FirmwareDistributable{}
	this.ClassId = classId
	this.ObjectType = objectType
	var origin string = "System"
	this.Origin = &origin
	return &this
}

// NewFirmwareDistributableWithDefaults instantiates a new FirmwareDistributable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareDistributableWithDefaults() *FirmwareDistributable {
	this := FirmwareDistributable{}
	var classId string = "firmware.Distributable"
	this.ClassId = classId
	var objectType string = "firmware.Distributable"
	this.ObjectType = objectType
	var origin string = "System"
	this.Origin = &origin
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareDistributable) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareDistributable) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareDistributable) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareDistributable) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareDistributable) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareDistributable) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *FirmwareDistributable) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributable) GetFileLocationOk() (*string, bool) {
	if o == nil || o.FileLocation == nil {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *FirmwareDistributable) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *FirmwareDistributable) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetImageCategory returns the ImageCategory field value if set, zero value otherwise.
func (o *FirmwareDistributable) GetImageCategory() string {
	if o == nil || o.ImageCategory == nil {
		var ret string
		return ret
	}
	return *o.ImageCategory
}

// GetImageCategoryOk returns a tuple with the ImageCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributable) GetImageCategoryOk() (*string, bool) {
	if o == nil || o.ImageCategory == nil {
		return nil, false
	}
	return o.ImageCategory, true
}

// HasImageCategory returns a boolean if a field has been set.
func (o *FirmwareDistributable) HasImageCategory() bool {
	if o != nil && o.ImageCategory != nil {
		return true
	}

	return false
}

// SetImageCategory gets a reference to the given string and assigns it to the ImageCategory field.
func (o *FirmwareDistributable) SetImageCategory(v string) {
	o.ImageCategory = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *FirmwareDistributable) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributable) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *FirmwareDistributable) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *FirmwareDistributable) SetOrigin(v string) {
	o.Origin = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *FirmwareDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *FirmwareDistributable) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *FirmwareDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o FirmwareDistributable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseDistributable, errFirmwareBaseDistributable := json.Marshal(o.FirmwareBaseDistributable)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	errFirmwareBaseDistributable = json.Unmarshal([]byte(serializedFirmwareBaseDistributable), &toSerialize)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileLocation != nil {
		toSerialize["FileLocation"] = o.FileLocation
	}
	if o.ImageCategory != nil {
		toSerialize["ImageCategory"] = o.ImageCategory
	}
	if o.Origin != nil {
		toSerialize["Origin"] = o.Origin
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareDistributable) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareDistributableWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The file location of the distributable.
		FileLocation *string `json:"FileLocation,omitempty"`
		// The category into which the distributable falls into according to the supported platform series. For e.g.; C-Series/B-Series/Infrastructure.
		ImageCategory *string `json:"ImageCategory,omitempty"`
		// The source of the distributable. If it has been created by the user or system. * `System` - The distributable has been created by the System. * `User` - The distributable has been created by the User.
		Origin  *string                                `json:"Origin,omitempty"`
		Catalog *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	}

	varFirmwareDistributableWithoutEmbeddedStruct := FirmwareDistributableWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareDistributableWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareDistributable := _FirmwareDistributable{}
		varFirmwareDistributable.ClassId = varFirmwareDistributableWithoutEmbeddedStruct.ClassId
		varFirmwareDistributable.ObjectType = varFirmwareDistributableWithoutEmbeddedStruct.ObjectType
		varFirmwareDistributable.FileLocation = varFirmwareDistributableWithoutEmbeddedStruct.FileLocation
		varFirmwareDistributable.ImageCategory = varFirmwareDistributableWithoutEmbeddedStruct.ImageCategory
		varFirmwareDistributable.Origin = varFirmwareDistributableWithoutEmbeddedStruct.Origin
		varFirmwareDistributable.Catalog = varFirmwareDistributableWithoutEmbeddedStruct.Catalog
		*o = FirmwareDistributable(varFirmwareDistributable)
	} else {
		return err
	}

	varFirmwareDistributable := _FirmwareDistributable{}

	err = json.Unmarshal(bytes, &varFirmwareDistributable)
	if err == nil {
		o.FirmwareBaseDistributable = varFirmwareDistributable.FirmwareBaseDistributable
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileLocation")
		delete(additionalProperties, "ImageCategory")
		delete(additionalProperties, "Origin")
		delete(additionalProperties, "Catalog")

		// remove fields from embedded structs
		reflectFirmwareBaseDistributable := reflect.ValueOf(o.FirmwareBaseDistributable)
		for i := 0; i < reflectFirmwareBaseDistributable.Type().NumField(); i++ {
			t := reflectFirmwareBaseDistributable.Type().Field(i)

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

type NullableFirmwareDistributable struct {
	value *FirmwareDistributable
	isSet bool
}

func (v NullableFirmwareDistributable) Get() *FirmwareDistributable {
	return v.value
}

func (v *NullableFirmwareDistributable) Set(val *FirmwareDistributable) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareDistributable) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareDistributable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareDistributable(val *FirmwareDistributable) *NullableFirmwareDistributable {
	return &NullableFirmwareDistributable{value: val, isSet: true}
}

func (v NullableFirmwareDistributable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareDistributable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
