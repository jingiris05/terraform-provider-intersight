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

// NiatelemetryApicSysLogSrc Object to capture the SysLogGroup details in APIC.
type NiatelemetryApicSysLogSrc struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the SysLogSrc in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Minimum sevirity on SysLogSrc MO in APIC.
	MinSev *string `json:"MinSev,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicSysLogSrc NiatelemetryApicSysLogSrc

// NewNiatelemetryApicSysLogSrc instantiates a new NiatelemetryApicSysLogSrc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicSysLogSrc(classId string, objectType string) *NiatelemetryApicSysLogSrc {
	this := NiatelemetryApicSysLogSrc{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicSysLogSrcWithDefaults instantiates a new NiatelemetryApicSysLogSrc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicSysLogSrcWithDefaults() *NiatelemetryApicSysLogSrc {
	this := NiatelemetryApicSysLogSrc{}
	var classId string = "niatelemetry.ApicSysLogSrc"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicSysLogSrc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicSysLogSrc) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrc) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicSysLogSrc) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicSysLogSrc) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrc) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicSysLogSrc) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrc) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrc) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrc) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicSysLogSrc) SetDn(v string) {
	o.Dn = &v
}

// GetMinSev returns the MinSev field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrc) GetMinSev() string {
	if o == nil || o.MinSev == nil {
		var ret string
		return ret
	}
	return *o.MinSev
}

// GetMinSevOk returns a tuple with the MinSev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrc) GetMinSevOk() (*string, bool) {
	if o == nil || o.MinSev == nil {
		return nil, false
	}
	return o.MinSev, true
}

// HasMinSev returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrc) HasMinSev() bool {
	if o != nil && o.MinSev != nil {
		return true
	}

	return false
}

// SetMinSev gets a reference to the given string and assigns it to the MinSev field.
func (o *NiatelemetryApicSysLogSrc) SetMinSev(v string) {
	o.MinSev = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrc) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrc) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrc) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicSysLogSrc) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrc) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrc) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrc) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicSysLogSrc) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrc) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrc) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrc) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicSysLogSrc) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrc) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrc) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrc) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicSysLogSrc) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicSysLogSrc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.MinSev != nil {
		toSerialize["MinSev"] = o.MinSev
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicSysLogSrc) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryApicSysLogSrcWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn of the SysLogSrc in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Minimum sevirity on SysLogSrc MO in APIC.
		MinSev *string `json:"MinSev,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicSysLogSrcWithoutEmbeddedStruct := NiatelemetryApicSysLogSrcWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicSysLogSrcWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicSysLogSrc := _NiatelemetryApicSysLogSrc{}
		varNiatelemetryApicSysLogSrc.ClassId = varNiatelemetryApicSysLogSrcWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicSysLogSrc.ObjectType = varNiatelemetryApicSysLogSrcWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicSysLogSrc.Dn = varNiatelemetryApicSysLogSrcWithoutEmbeddedStruct.Dn
		varNiatelemetryApicSysLogSrc.MinSev = varNiatelemetryApicSysLogSrcWithoutEmbeddedStruct.MinSev
		varNiatelemetryApicSysLogSrc.RecordType = varNiatelemetryApicSysLogSrcWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicSysLogSrc.RecordVersion = varNiatelemetryApicSysLogSrcWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicSysLogSrc.SiteName = varNiatelemetryApicSysLogSrcWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicSysLogSrc.RegisteredDevice = varNiatelemetryApicSysLogSrcWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicSysLogSrc(varNiatelemetryApicSysLogSrc)
	} else {
		return err
	}

	varNiatelemetryApicSysLogSrc := _NiatelemetryApicSysLogSrc{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicSysLogSrc)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicSysLogSrc.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "MinSev")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableNiatelemetryApicSysLogSrc struct {
	value *NiatelemetryApicSysLogSrc
	isSet bool
}

func (v NullableNiatelemetryApicSysLogSrc) Get() *NiatelemetryApicSysLogSrc {
	return v.value
}

func (v *NullableNiatelemetryApicSysLogSrc) Set(val *NiatelemetryApicSysLogSrc) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicSysLogSrc) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicSysLogSrc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicSysLogSrc(val *NiatelemetryApicSysLogSrc) *NullableNiatelemetryApicSysLogSrc {
	return &NullableNiatelemetryApicSysLogSrc{value: val, isSet: true}
}

func (v NullableNiatelemetryApicSysLogSrc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicSysLogSrc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
