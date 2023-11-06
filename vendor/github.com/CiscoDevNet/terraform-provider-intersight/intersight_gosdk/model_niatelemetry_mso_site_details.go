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
	"reflect"
	"strings"
)

// NiatelemetryMsoSiteDetails Details of sites in Multi-Site Orchestrator.
type NiatelemetryMsoSiteDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of cloudSec on Multi-Site Orchestrator site.
	IsCloudSecEnabled *string `json:"IsCloudSecEnabled,omitempty"`
	// Number of leafs per site in Multi-Site Orchestrator.
	NumberOfLeafsPerSiteInMso *int64 `json:"NumberOfLeafsPerSiteInMso,omitempty"`
	// Number of pods per site in Multi-Site Orchestrator.
	NumberOfPodsPerSiteInMso *int64 `json:"NumberOfPodsPerSiteInMso,omitempty"`
	// Number of spines per site in Multi-Site Orchestrator.
	NumberOfSpinesPerSiteInMso *int64 `json:"NumberOfSpinesPerSiteInMso,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// ID of site in Multi-Site Orchestrator.
	SiteId *string `json:"SiteId,omitempty"`
	// Name of the site in Multi-Site Orchestrator.
	SiteName *string `json:"SiteName,omitempty"`
	// Version of the controller in the site.
	SiteVersion          *string                              `json:"SiteVersion,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryMsoSiteDetails NiatelemetryMsoSiteDetails

// NewNiatelemetryMsoSiteDetails instantiates a new NiatelemetryMsoSiteDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryMsoSiteDetails(classId string, objectType string) *NiatelemetryMsoSiteDetails {
	this := NiatelemetryMsoSiteDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryMsoSiteDetailsWithDefaults instantiates a new NiatelemetryMsoSiteDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryMsoSiteDetailsWithDefaults() *NiatelemetryMsoSiteDetails {
	this := NiatelemetryMsoSiteDetails{}
	var classId string = "niatelemetry.MsoSiteDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.MsoSiteDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryMsoSiteDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryMsoSiteDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryMsoSiteDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryMsoSiteDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsCloudSecEnabled returns the IsCloudSecEnabled field value if set, zero value otherwise.
func (o *NiatelemetryMsoSiteDetails) GetIsCloudSecEnabled() string {
	if o == nil || o.IsCloudSecEnabled == nil {
		var ret string
		return ret
	}
	return *o.IsCloudSecEnabled
}

// GetIsCloudSecEnabledOk returns a tuple with the IsCloudSecEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetIsCloudSecEnabledOk() (*string, bool) {
	if o == nil || o.IsCloudSecEnabled == nil {
		return nil, false
	}
	return o.IsCloudSecEnabled, true
}

// HasIsCloudSecEnabled returns a boolean if a field has been set.
func (o *NiatelemetryMsoSiteDetails) HasIsCloudSecEnabled() bool {
	if o != nil && o.IsCloudSecEnabled != nil {
		return true
	}

	return false
}

// SetIsCloudSecEnabled gets a reference to the given string and assigns it to the IsCloudSecEnabled field.
func (o *NiatelemetryMsoSiteDetails) SetIsCloudSecEnabled(v string) {
	o.IsCloudSecEnabled = &v
}

// GetNumberOfLeafsPerSiteInMso returns the NumberOfLeafsPerSiteInMso field value if set, zero value otherwise.
func (o *NiatelemetryMsoSiteDetails) GetNumberOfLeafsPerSiteInMso() int64 {
	if o == nil || o.NumberOfLeafsPerSiteInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfLeafsPerSiteInMso
}

// GetNumberOfLeafsPerSiteInMsoOk returns a tuple with the NumberOfLeafsPerSiteInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetNumberOfLeafsPerSiteInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfLeafsPerSiteInMso == nil {
		return nil, false
	}
	return o.NumberOfLeafsPerSiteInMso, true
}

// HasNumberOfLeafsPerSiteInMso returns a boolean if a field has been set.
func (o *NiatelemetryMsoSiteDetails) HasNumberOfLeafsPerSiteInMso() bool {
	if o != nil && o.NumberOfLeafsPerSiteInMso != nil {
		return true
	}

	return false
}

// SetNumberOfLeafsPerSiteInMso gets a reference to the given int64 and assigns it to the NumberOfLeafsPerSiteInMso field.
func (o *NiatelemetryMsoSiteDetails) SetNumberOfLeafsPerSiteInMso(v int64) {
	o.NumberOfLeafsPerSiteInMso = &v
}

// GetNumberOfPodsPerSiteInMso returns the NumberOfPodsPerSiteInMso field value if set, zero value otherwise.
func (o *NiatelemetryMsoSiteDetails) GetNumberOfPodsPerSiteInMso() int64 {
	if o == nil || o.NumberOfPodsPerSiteInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfPodsPerSiteInMso
}

// GetNumberOfPodsPerSiteInMsoOk returns a tuple with the NumberOfPodsPerSiteInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetNumberOfPodsPerSiteInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfPodsPerSiteInMso == nil {
		return nil, false
	}
	return o.NumberOfPodsPerSiteInMso, true
}

// HasNumberOfPodsPerSiteInMso returns a boolean if a field has been set.
func (o *NiatelemetryMsoSiteDetails) HasNumberOfPodsPerSiteInMso() bool {
	if o != nil && o.NumberOfPodsPerSiteInMso != nil {
		return true
	}

	return false
}

// SetNumberOfPodsPerSiteInMso gets a reference to the given int64 and assigns it to the NumberOfPodsPerSiteInMso field.
func (o *NiatelemetryMsoSiteDetails) SetNumberOfPodsPerSiteInMso(v int64) {
	o.NumberOfPodsPerSiteInMso = &v
}

// GetNumberOfSpinesPerSiteInMso returns the NumberOfSpinesPerSiteInMso field value if set, zero value otherwise.
func (o *NiatelemetryMsoSiteDetails) GetNumberOfSpinesPerSiteInMso() int64 {
	if o == nil || o.NumberOfSpinesPerSiteInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSpinesPerSiteInMso
}

// GetNumberOfSpinesPerSiteInMsoOk returns a tuple with the NumberOfSpinesPerSiteInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetNumberOfSpinesPerSiteInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfSpinesPerSiteInMso == nil {
		return nil, false
	}
	return o.NumberOfSpinesPerSiteInMso, true
}

// HasNumberOfSpinesPerSiteInMso returns a boolean if a field has been set.
func (o *NiatelemetryMsoSiteDetails) HasNumberOfSpinesPerSiteInMso() bool {
	if o != nil && o.NumberOfSpinesPerSiteInMso != nil {
		return true
	}

	return false
}

// SetNumberOfSpinesPerSiteInMso gets a reference to the given int64 and assigns it to the NumberOfSpinesPerSiteInMso field.
func (o *NiatelemetryMsoSiteDetails) SetNumberOfSpinesPerSiteInMso(v int64) {
	o.NumberOfSpinesPerSiteInMso = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryMsoSiteDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryMsoSiteDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryMsoSiteDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *NiatelemetryMsoSiteDetails) GetSiteId() string {
	if o == nil || o.SiteId == nil {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetSiteIdOk() (*string, bool) {
	if o == nil || o.SiteId == nil {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *NiatelemetryMsoSiteDetails) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *NiatelemetryMsoSiteDetails) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryMsoSiteDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryMsoSiteDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryMsoSiteDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSiteVersion returns the SiteVersion field value if set, zero value otherwise.
func (o *NiatelemetryMsoSiteDetails) GetSiteVersion() string {
	if o == nil || o.SiteVersion == nil {
		var ret string
		return ret
	}
	return *o.SiteVersion
}

// GetSiteVersionOk returns a tuple with the SiteVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetSiteVersionOk() (*string, bool) {
	if o == nil || o.SiteVersion == nil {
		return nil, false
	}
	return o.SiteVersion, true
}

// HasSiteVersion returns a boolean if a field has been set.
func (o *NiatelemetryMsoSiteDetails) HasSiteVersion() bool {
	if o != nil && o.SiteVersion != nil {
		return true
	}

	return false
}

// SetSiteVersion gets a reference to the given string and assigns it to the SiteVersion field.
func (o *NiatelemetryMsoSiteDetails) SetSiteVersion(v string) {
	o.SiteVersion = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryMsoSiteDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSiteDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryMsoSiteDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryMsoSiteDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryMsoSiteDetails) MarshalJSON() ([]byte, error) {
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
	if o.IsCloudSecEnabled != nil {
		toSerialize["IsCloudSecEnabled"] = o.IsCloudSecEnabled
	}
	if o.NumberOfLeafsPerSiteInMso != nil {
		toSerialize["NumberOfLeafsPerSiteInMso"] = o.NumberOfLeafsPerSiteInMso
	}
	if o.NumberOfPodsPerSiteInMso != nil {
		toSerialize["NumberOfPodsPerSiteInMso"] = o.NumberOfPodsPerSiteInMso
	}
	if o.NumberOfSpinesPerSiteInMso != nil {
		toSerialize["NumberOfSpinesPerSiteInMso"] = o.NumberOfSpinesPerSiteInMso
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.SiteId != nil {
		toSerialize["SiteId"] = o.SiteId
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.SiteVersion != nil {
		toSerialize["SiteVersion"] = o.SiteVersion
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryMsoSiteDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryMsoSiteDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of cloudSec on Multi-Site Orchestrator site.
		IsCloudSecEnabled *string `json:"IsCloudSecEnabled,omitempty"`
		// Number of leafs per site in Multi-Site Orchestrator.
		NumberOfLeafsPerSiteInMso *int64 `json:"NumberOfLeafsPerSiteInMso,omitempty"`
		// Number of pods per site in Multi-Site Orchestrator.
		NumberOfPodsPerSiteInMso *int64 `json:"NumberOfPodsPerSiteInMso,omitempty"`
		// Number of spines per site in Multi-Site Orchestrator.
		NumberOfSpinesPerSiteInMso *int64 `json:"NumberOfSpinesPerSiteInMso,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// ID of site in Multi-Site Orchestrator.
		SiteId *string `json:"SiteId,omitempty"`
		// Name of the site in Multi-Site Orchestrator.
		SiteName *string `json:"SiteName,omitempty"`
		// Version of the controller in the site.
		SiteVersion      *string                              `json:"SiteVersion,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct := NiatelemetryMsoSiteDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryMsoSiteDetails := _NiatelemetryMsoSiteDetails{}
		varNiatelemetryMsoSiteDetails.ClassId = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryMsoSiteDetails.ObjectType = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryMsoSiteDetails.IsCloudSecEnabled = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.IsCloudSecEnabled
		varNiatelemetryMsoSiteDetails.NumberOfLeafsPerSiteInMso = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.NumberOfLeafsPerSiteInMso
		varNiatelemetryMsoSiteDetails.NumberOfPodsPerSiteInMso = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.NumberOfPodsPerSiteInMso
		varNiatelemetryMsoSiteDetails.NumberOfSpinesPerSiteInMso = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.NumberOfSpinesPerSiteInMso
		varNiatelemetryMsoSiteDetails.RecordType = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryMsoSiteDetails.SiteId = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.SiteId
		varNiatelemetryMsoSiteDetails.SiteName = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryMsoSiteDetails.SiteVersion = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.SiteVersion
		varNiatelemetryMsoSiteDetails.RegisteredDevice = varNiatelemetryMsoSiteDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryMsoSiteDetails(varNiatelemetryMsoSiteDetails)
	} else {
		return err
	}

	varNiatelemetryMsoSiteDetails := _NiatelemetryMsoSiteDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryMsoSiteDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryMsoSiteDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsCloudSecEnabled")
		delete(additionalProperties, "NumberOfLeafsPerSiteInMso")
		delete(additionalProperties, "NumberOfPodsPerSiteInMso")
		delete(additionalProperties, "NumberOfSpinesPerSiteInMso")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "SiteId")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SiteVersion")
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

type NullableNiatelemetryMsoSiteDetails struct {
	value *NiatelemetryMsoSiteDetails
	isSet bool
}

func (v NullableNiatelemetryMsoSiteDetails) Get() *NiatelemetryMsoSiteDetails {
	return v.value
}

func (v *NullableNiatelemetryMsoSiteDetails) Set(val *NiatelemetryMsoSiteDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMsoSiteDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMsoSiteDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMsoSiteDetails(val *NiatelemetryMsoSiteDetails) *NullableNiatelemetryMsoSiteDetails {
	return &NullableNiatelemetryMsoSiteDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryMsoSiteDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMsoSiteDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
