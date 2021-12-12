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
)

// BootSdCardAllOf Definition of the list of properties defined in 'boot.SdCard', excluding properties defined in parent classes.
type BootSdCardAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                 `json:"ObjectType"`
	Bootloader NullableBootBootloader `json:"Bootloader,omitempty"`
	// The Logical Unit Number (LUN) of the device.
	Lun *int64 `json:"Lun,omitempty"`
	// The subtype for the selected device type. * `None` - No sub type for SD card boot device. * `flex-util` - Use of FlexUtil (microSD) card as sub type for SD card boot device. * `flex-flash` - Use of FlexFlash (SD) card as sub type for SD card boot device. * `SDCARD` - Use of SD card as sub type for the SD Card boot device.
	Subtype              *string `json:"Subtype,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootSdCardAllOf BootSdCardAllOf

// NewBootSdCardAllOf instantiates a new BootSdCardAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootSdCardAllOf(classId string, objectType string) *BootSdCardAllOf {
	this := BootSdCardAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var lun int64 = 0
	this.Lun = &lun
	var subtype string = "None"
	this.Subtype = &subtype
	return &this
}

// NewBootSdCardAllOfWithDefaults instantiates a new BootSdCardAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootSdCardAllOfWithDefaults() *BootSdCardAllOf {
	this := BootSdCardAllOf{}
	var classId string = "boot.SdCard"
	this.ClassId = classId
	var objectType string = "boot.SdCard"
	this.ObjectType = objectType
	var lun int64 = 0
	this.Lun = &lun
	var subtype string = "None"
	this.Subtype = &subtype
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootSdCardAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootSdCardAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootSdCardAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootSdCardAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootSdCardAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootSdCardAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootloader returns the Bootloader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootSdCardAllOf) GetBootloader() BootBootloader {
	if o == nil || o.Bootloader.Get() == nil {
		var ret BootBootloader
		return ret
	}
	return *o.Bootloader.Get()
}

// GetBootloaderOk returns a tuple with the Bootloader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootSdCardAllOf) GetBootloaderOk() (*BootBootloader, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bootloader.Get(), o.Bootloader.IsSet()
}

// HasBootloader returns a boolean if a field has been set.
func (o *BootSdCardAllOf) HasBootloader() bool {
	if o != nil && o.Bootloader.IsSet() {
		return true
	}

	return false
}

// SetBootloader gets a reference to the given NullableBootBootloader and assigns it to the Bootloader field.
func (o *BootSdCardAllOf) SetBootloader(v BootBootloader) {
	o.Bootloader.Set(&v)
}

// SetBootloaderNil sets the value for Bootloader to be an explicit nil
func (o *BootSdCardAllOf) SetBootloaderNil() {
	o.Bootloader.Set(nil)
}

// UnsetBootloader ensures that no value is present for Bootloader, not even an explicit nil
func (o *BootSdCardAllOf) UnsetBootloader() {
	o.Bootloader.Unset()
}

// GetLun returns the Lun field value if set, zero value otherwise.
func (o *BootSdCardAllOf) GetLun() int64 {
	if o == nil || o.Lun == nil {
		var ret int64
		return ret
	}
	return *o.Lun
}

// GetLunOk returns a tuple with the Lun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootSdCardAllOf) GetLunOk() (*int64, bool) {
	if o == nil || o.Lun == nil {
		return nil, false
	}
	return o.Lun, true
}

// HasLun returns a boolean if a field has been set.
func (o *BootSdCardAllOf) HasLun() bool {
	if o != nil && o.Lun != nil {
		return true
	}

	return false
}

// SetLun gets a reference to the given int64 and assigns it to the Lun field.
func (o *BootSdCardAllOf) SetLun(v int64) {
	o.Lun = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *BootSdCardAllOf) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootSdCardAllOf) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *BootSdCardAllOf) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *BootSdCardAllOf) SetSubtype(v string) {
	o.Subtype = &v
}

func (o BootSdCardAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootloader.IsSet() {
		toSerialize["Bootloader"] = o.Bootloader.Get()
	}
	if o.Lun != nil {
		toSerialize["Lun"] = o.Lun
	}
	if o.Subtype != nil {
		toSerialize["Subtype"] = o.Subtype
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootSdCardAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootSdCardAllOf := _BootSdCardAllOf{}

	if err = json.Unmarshal(bytes, &varBootSdCardAllOf); err == nil {
		*o = BootSdCardAllOf(varBootSdCardAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootloader")
		delete(additionalProperties, "Lun")
		delete(additionalProperties, "Subtype")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootSdCardAllOf struct {
	value *BootSdCardAllOf
	isSet bool
}

func (v NullableBootSdCardAllOf) Get() *BootSdCardAllOf {
	return v.value
}

func (v *NullableBootSdCardAllOf) Set(val *BootSdCardAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootSdCardAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootSdCardAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootSdCardAllOf(val *BootSdCardAllOf) *NullableBootSdCardAllOf {
	return &NullableBootSdCardAllOf{value: val, isSet: true}
}

func (v NullableBootSdCardAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootSdCardAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
