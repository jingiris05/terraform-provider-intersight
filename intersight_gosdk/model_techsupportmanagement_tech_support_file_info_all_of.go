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

// TechsupportmanagementTechSupportFileInfoAllOf Definition of the list of properties defined in 'techsupportmanagement.TechSupportFileInfo', excluding properties defined in parent classes.
type TechsupportmanagementTechSupportFileInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the techsupport file.
	FileName *string `json:"FileName,omitempty"`
	// Techsupport file size in bytes.
	FileSize *int64 `json:"FileSize,omitempty"`
	// The Url to download the techsupport file.
	TechsupportDownloadUrl *string `json:"TechsupportDownloadUrl,omitempty"`
	// The upload status of the techsupport file.
	UploadStatus         *string `json:"UploadStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TechsupportmanagementTechSupportFileInfoAllOf TechsupportmanagementTechSupportFileInfoAllOf

// NewTechsupportmanagementTechSupportFileInfoAllOf instantiates a new TechsupportmanagementTechSupportFileInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementTechSupportFileInfoAllOf(classId string, objectType string) *TechsupportmanagementTechSupportFileInfoAllOf {
	this := TechsupportmanagementTechSupportFileInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTechsupportmanagementTechSupportFileInfoAllOfWithDefaults instantiates a new TechsupportmanagementTechSupportFileInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementTechSupportFileInfoAllOfWithDefaults() *TechsupportmanagementTechSupportFileInfoAllOf {
	this := TechsupportmanagementTechSupportFileInfoAllOf{}
	var classId string = "techsupportmanagement.TechSupportFileInfo"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.TechSupportFileInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementTechSupportFileInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementTechSupportFileInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetFileSize() int64 {
	if o == nil || o.FileSize == nil {
		var ret int64
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetFileSizeOk() (*int64, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int64 and assigns it to the FileSize field.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) SetFileSize(v int64) {
	o.FileSize = &v
}

// GetTechsupportDownloadUrl returns the TechsupportDownloadUrl field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetTechsupportDownloadUrl() string {
	if o == nil || o.TechsupportDownloadUrl == nil {
		var ret string
		return ret
	}
	return *o.TechsupportDownloadUrl
}

// GetTechsupportDownloadUrlOk returns a tuple with the TechsupportDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetTechsupportDownloadUrlOk() (*string, bool) {
	if o == nil || o.TechsupportDownloadUrl == nil {
		return nil, false
	}
	return o.TechsupportDownloadUrl, true
}

// HasTechsupportDownloadUrl returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) HasTechsupportDownloadUrl() bool {
	if o != nil && o.TechsupportDownloadUrl != nil {
		return true
	}

	return false
}

// SetTechsupportDownloadUrl gets a reference to the given string and assigns it to the TechsupportDownloadUrl field.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) SetTechsupportDownloadUrl(v string) {
	o.TechsupportDownloadUrl = &v
}

// GetUploadStatus returns the UploadStatus field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetUploadStatus() string {
	if o == nil || o.UploadStatus == nil {
		var ret string
		return ret
	}
	return *o.UploadStatus
}

// GetUploadStatusOk returns a tuple with the UploadStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) GetUploadStatusOk() (*string, bool) {
	if o == nil || o.UploadStatus == nil {
		return nil, false
	}
	return o.UploadStatus, true
}

// HasUploadStatus returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) HasUploadStatus() bool {
	if o != nil && o.UploadStatus != nil {
		return true
	}

	return false
}

// SetUploadStatus gets a reference to the given string and assigns it to the UploadStatus field.
func (o *TechsupportmanagementTechSupportFileInfoAllOf) SetUploadStatus(v string) {
	o.UploadStatus = &v
}

func (o TechsupportmanagementTechSupportFileInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.FileSize != nil {
		toSerialize["FileSize"] = o.FileSize
	}
	if o.TechsupportDownloadUrl != nil {
		toSerialize["TechsupportDownloadUrl"] = o.TechsupportDownloadUrl
	}
	if o.UploadStatus != nil {
		toSerialize["UploadStatus"] = o.UploadStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TechsupportmanagementTechSupportFileInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTechsupportmanagementTechSupportFileInfoAllOf := _TechsupportmanagementTechSupportFileInfoAllOf{}

	if err = json.Unmarshal(bytes, &varTechsupportmanagementTechSupportFileInfoAllOf); err == nil {
		*o = TechsupportmanagementTechSupportFileInfoAllOf(varTechsupportmanagementTechSupportFileInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileName")
		delete(additionalProperties, "FileSize")
		delete(additionalProperties, "TechsupportDownloadUrl")
		delete(additionalProperties, "UploadStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTechsupportmanagementTechSupportFileInfoAllOf struct {
	value *TechsupportmanagementTechSupportFileInfoAllOf
	isSet bool
}

func (v NullableTechsupportmanagementTechSupportFileInfoAllOf) Get() *TechsupportmanagementTechSupportFileInfoAllOf {
	return v.value
}

func (v *NullableTechsupportmanagementTechSupportFileInfoAllOf) Set(val *TechsupportmanagementTechSupportFileInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementTechSupportFileInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementTechSupportFileInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementTechSupportFileInfoAllOf(val *TechsupportmanagementTechSupportFileInfoAllOf) *NullableTechsupportmanagementTechSupportFileInfoAllOf {
	return &NullableTechsupportmanagementTechSupportFileInfoAllOf{value: val, isSet: true}
}

func (v NullableTechsupportmanagementTechSupportFileInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementTechSupportFileInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
