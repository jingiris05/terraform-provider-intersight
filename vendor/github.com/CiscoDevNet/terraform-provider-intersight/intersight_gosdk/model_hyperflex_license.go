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

// HyperflexLicense Feature Identifier
type HyperflexLicense struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Is the cluster complaint with the license entitlements?
	ComplianceState *string `json:"ComplianceState,omitempty"`
	// Out of compliance date of the cluster
	GetOutOfComplianceStartAt *string `json:"GetOutOfComplianceStartAt,omitempty"`
	// Is the cluster in evalution period?
	InEvaluation         *bool                                            `json:"InEvaluation,omitempty"`
	LicenseAuthorization NullableHyperflexHxLicenseAuthorizationDetailsDt `json:"LicenseAuthorization,omitempty"`
	LicenseRegistration  NullableHyperflexHxRegistrationDetailsDt         `json:"LicenseRegistration,omitempty"`
	// The type of license applied on the cluster
	LicenseType *string `json:"LicenseType,omitempty"`
	// Is reservation enabled for the cluster?
	PlrEnabled *bool `json:"PlrEnabled,omitempty"`
	// Is Smart Licensing Enabled for this cluster?
	SmartLicensingEnabled *bool                                `json:"SmartLicensingEnabled,omitempty"`
	Cluster               *HyperflexClusterRelationship        `json:"Cluster,omitempty"`
	RegisteredDevice      *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _HyperflexLicense HyperflexLicense

// NewHyperflexLicense instantiates a new HyperflexLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexLicense(classId string, objectType string) *HyperflexLicense {
	this := HyperflexLicense{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexLicenseWithDefaults instantiates a new HyperflexLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexLicenseWithDefaults() *HyperflexLicense {
	this := HyperflexLicense{}
	var classId string = "hyperflex.License"
	this.ClassId = classId
	var objectType string = "hyperflex.License"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexLicense) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexLicense) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexLicense) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexLicense) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexLicense) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexLicense) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComplianceState returns the ComplianceState field value if set, zero value otherwise.
func (o *HyperflexLicense) GetComplianceState() string {
	if o == nil || o.ComplianceState == nil {
		var ret string
		return ret
	}
	return *o.ComplianceState
}

// GetComplianceStateOk returns a tuple with the ComplianceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLicense) GetComplianceStateOk() (*string, bool) {
	if o == nil || o.ComplianceState == nil {
		return nil, false
	}
	return o.ComplianceState, true
}

// HasComplianceState returns a boolean if a field has been set.
func (o *HyperflexLicense) HasComplianceState() bool {
	if o != nil && o.ComplianceState != nil {
		return true
	}

	return false
}

// SetComplianceState gets a reference to the given string and assigns it to the ComplianceState field.
func (o *HyperflexLicense) SetComplianceState(v string) {
	o.ComplianceState = &v
}

// GetGetOutOfComplianceStartAt returns the GetOutOfComplianceStartAt field value if set, zero value otherwise.
func (o *HyperflexLicense) GetGetOutOfComplianceStartAt() string {
	if o == nil || o.GetOutOfComplianceStartAt == nil {
		var ret string
		return ret
	}
	return *o.GetOutOfComplianceStartAt
}

// GetGetOutOfComplianceStartAtOk returns a tuple with the GetOutOfComplianceStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLicense) GetGetOutOfComplianceStartAtOk() (*string, bool) {
	if o == nil || o.GetOutOfComplianceStartAt == nil {
		return nil, false
	}
	return o.GetOutOfComplianceStartAt, true
}

// HasGetOutOfComplianceStartAt returns a boolean if a field has been set.
func (o *HyperflexLicense) HasGetOutOfComplianceStartAt() bool {
	if o != nil && o.GetOutOfComplianceStartAt != nil {
		return true
	}

	return false
}

// SetGetOutOfComplianceStartAt gets a reference to the given string and assigns it to the GetOutOfComplianceStartAt field.
func (o *HyperflexLicense) SetGetOutOfComplianceStartAt(v string) {
	o.GetOutOfComplianceStartAt = &v
}

// GetInEvaluation returns the InEvaluation field value if set, zero value otherwise.
func (o *HyperflexLicense) GetInEvaluation() bool {
	if o == nil || o.InEvaluation == nil {
		var ret bool
		return ret
	}
	return *o.InEvaluation
}

// GetInEvaluationOk returns a tuple with the InEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLicense) GetInEvaluationOk() (*bool, bool) {
	if o == nil || o.InEvaluation == nil {
		return nil, false
	}
	return o.InEvaluation, true
}

// HasInEvaluation returns a boolean if a field has been set.
func (o *HyperflexLicense) HasInEvaluation() bool {
	if o != nil && o.InEvaluation != nil {
		return true
	}

	return false
}

// SetInEvaluation gets a reference to the given bool and assigns it to the InEvaluation field.
func (o *HyperflexLicense) SetInEvaluation(v bool) {
	o.InEvaluation = &v
}

// GetLicenseAuthorization returns the LicenseAuthorization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexLicense) GetLicenseAuthorization() HyperflexHxLicenseAuthorizationDetailsDt {
	if o == nil || o.LicenseAuthorization.Get() == nil {
		var ret HyperflexHxLicenseAuthorizationDetailsDt
		return ret
	}
	return *o.LicenseAuthorization.Get()
}

// GetLicenseAuthorizationOk returns a tuple with the LicenseAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexLicense) GetLicenseAuthorizationOk() (*HyperflexHxLicenseAuthorizationDetailsDt, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenseAuthorization.Get(), o.LicenseAuthorization.IsSet()
}

// HasLicenseAuthorization returns a boolean if a field has been set.
func (o *HyperflexLicense) HasLicenseAuthorization() bool {
	if o != nil && o.LicenseAuthorization.IsSet() {
		return true
	}

	return false
}

// SetLicenseAuthorization gets a reference to the given NullableHyperflexHxLicenseAuthorizationDetailsDt and assigns it to the LicenseAuthorization field.
func (o *HyperflexLicense) SetLicenseAuthorization(v HyperflexHxLicenseAuthorizationDetailsDt) {
	o.LicenseAuthorization.Set(&v)
}

// SetLicenseAuthorizationNil sets the value for LicenseAuthorization to be an explicit nil
func (o *HyperflexLicense) SetLicenseAuthorizationNil() {
	o.LicenseAuthorization.Set(nil)
}

// UnsetLicenseAuthorization ensures that no value is present for LicenseAuthorization, not even an explicit nil
func (o *HyperflexLicense) UnsetLicenseAuthorization() {
	o.LicenseAuthorization.Unset()
}

// GetLicenseRegistration returns the LicenseRegistration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexLicense) GetLicenseRegistration() HyperflexHxRegistrationDetailsDt {
	if o == nil || o.LicenseRegistration.Get() == nil {
		var ret HyperflexHxRegistrationDetailsDt
		return ret
	}
	return *o.LicenseRegistration.Get()
}

// GetLicenseRegistrationOk returns a tuple with the LicenseRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexLicense) GetLicenseRegistrationOk() (*HyperflexHxRegistrationDetailsDt, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenseRegistration.Get(), o.LicenseRegistration.IsSet()
}

// HasLicenseRegistration returns a boolean if a field has been set.
func (o *HyperflexLicense) HasLicenseRegistration() bool {
	if o != nil && o.LicenseRegistration.IsSet() {
		return true
	}

	return false
}

// SetLicenseRegistration gets a reference to the given NullableHyperflexHxRegistrationDetailsDt and assigns it to the LicenseRegistration field.
func (o *HyperflexLicense) SetLicenseRegistration(v HyperflexHxRegistrationDetailsDt) {
	o.LicenseRegistration.Set(&v)
}

// SetLicenseRegistrationNil sets the value for LicenseRegistration to be an explicit nil
func (o *HyperflexLicense) SetLicenseRegistrationNil() {
	o.LicenseRegistration.Set(nil)
}

// UnsetLicenseRegistration ensures that no value is present for LicenseRegistration, not even an explicit nil
func (o *HyperflexLicense) UnsetLicenseRegistration() {
	o.LicenseRegistration.Unset()
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *HyperflexLicense) GetLicenseType() string {
	if o == nil || o.LicenseType == nil {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLicense) GetLicenseTypeOk() (*string, bool) {
	if o == nil || o.LicenseType == nil {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *HyperflexLicense) HasLicenseType() bool {
	if o != nil && o.LicenseType != nil {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *HyperflexLicense) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetPlrEnabled returns the PlrEnabled field value if set, zero value otherwise.
func (o *HyperflexLicense) GetPlrEnabled() bool {
	if o == nil || o.PlrEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PlrEnabled
}

// GetPlrEnabledOk returns a tuple with the PlrEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLicense) GetPlrEnabledOk() (*bool, bool) {
	if o == nil || o.PlrEnabled == nil {
		return nil, false
	}
	return o.PlrEnabled, true
}

// HasPlrEnabled returns a boolean if a field has been set.
func (o *HyperflexLicense) HasPlrEnabled() bool {
	if o != nil && o.PlrEnabled != nil {
		return true
	}

	return false
}

// SetPlrEnabled gets a reference to the given bool and assigns it to the PlrEnabled field.
func (o *HyperflexLicense) SetPlrEnabled(v bool) {
	o.PlrEnabled = &v
}

// GetSmartLicensingEnabled returns the SmartLicensingEnabled field value if set, zero value otherwise.
func (o *HyperflexLicense) GetSmartLicensingEnabled() bool {
	if o == nil || o.SmartLicensingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SmartLicensingEnabled
}

// GetSmartLicensingEnabledOk returns a tuple with the SmartLicensingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLicense) GetSmartLicensingEnabledOk() (*bool, bool) {
	if o == nil || o.SmartLicensingEnabled == nil {
		return nil, false
	}
	return o.SmartLicensingEnabled, true
}

// HasSmartLicensingEnabled returns a boolean if a field has been set.
func (o *HyperflexLicense) HasSmartLicensingEnabled() bool {
	if o != nil && o.SmartLicensingEnabled != nil {
		return true
	}

	return false
}

// SetSmartLicensingEnabled gets a reference to the given bool and assigns it to the SmartLicensingEnabled field.
func (o *HyperflexLicense) SetSmartLicensingEnabled(v bool) {
	o.SmartLicensingEnabled = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexLicense) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLicense) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexLicense) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexLicense) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *HyperflexLicense) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexLicense) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexLicense) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexLicense) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o HyperflexLicense) MarshalJSON() ([]byte, error) {
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
	if o.ComplianceState != nil {
		toSerialize["ComplianceState"] = o.ComplianceState
	}
	if o.GetOutOfComplianceStartAt != nil {
		toSerialize["GetOutOfComplianceStartAt"] = o.GetOutOfComplianceStartAt
	}
	if o.InEvaluation != nil {
		toSerialize["InEvaluation"] = o.InEvaluation
	}
	if o.LicenseAuthorization.IsSet() {
		toSerialize["LicenseAuthorization"] = o.LicenseAuthorization.Get()
	}
	if o.LicenseRegistration.IsSet() {
		toSerialize["LicenseRegistration"] = o.LicenseRegistration.Get()
	}
	if o.LicenseType != nil {
		toSerialize["LicenseType"] = o.LicenseType
	}
	if o.PlrEnabled != nil {
		toSerialize["PlrEnabled"] = o.PlrEnabled
	}
	if o.SmartLicensingEnabled != nil {
		toSerialize["SmartLicensingEnabled"] = o.SmartLicensingEnabled
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexLicense) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexLicenseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Is the cluster complaint with the license entitlements?
		ComplianceState *string `json:"ComplianceState,omitempty"`
		// Out of compliance date of the cluster
		GetOutOfComplianceStartAt *string `json:"GetOutOfComplianceStartAt,omitempty"`
		// Is the cluster in evalution period?
		InEvaluation         *bool                                            `json:"InEvaluation,omitempty"`
		LicenseAuthorization NullableHyperflexHxLicenseAuthorizationDetailsDt `json:"LicenseAuthorization,omitempty"`
		LicenseRegistration  NullableHyperflexHxRegistrationDetailsDt         `json:"LicenseRegistration,omitempty"`
		// The type of license applied on the cluster
		LicenseType *string `json:"LicenseType,omitempty"`
		// Is reservation enabled for the cluster?
		PlrEnabled *bool `json:"PlrEnabled,omitempty"`
		// Is Smart Licensing Enabled for this cluster?
		SmartLicensingEnabled *bool                                `json:"SmartLicensingEnabled,omitempty"`
		Cluster               *HyperflexClusterRelationship        `json:"Cluster,omitempty"`
		RegisteredDevice      *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varHyperflexLicenseWithoutEmbeddedStruct := HyperflexLicenseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexLicenseWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexLicense := _HyperflexLicense{}
		varHyperflexLicense.ClassId = varHyperflexLicenseWithoutEmbeddedStruct.ClassId
		varHyperflexLicense.ObjectType = varHyperflexLicenseWithoutEmbeddedStruct.ObjectType
		varHyperflexLicense.ComplianceState = varHyperflexLicenseWithoutEmbeddedStruct.ComplianceState
		varHyperflexLicense.GetOutOfComplianceStartAt = varHyperflexLicenseWithoutEmbeddedStruct.GetOutOfComplianceStartAt
		varHyperflexLicense.InEvaluation = varHyperflexLicenseWithoutEmbeddedStruct.InEvaluation
		varHyperflexLicense.LicenseAuthorization = varHyperflexLicenseWithoutEmbeddedStruct.LicenseAuthorization
		varHyperflexLicense.LicenseRegistration = varHyperflexLicenseWithoutEmbeddedStruct.LicenseRegistration
		varHyperflexLicense.LicenseType = varHyperflexLicenseWithoutEmbeddedStruct.LicenseType
		varHyperflexLicense.PlrEnabled = varHyperflexLicenseWithoutEmbeddedStruct.PlrEnabled
		varHyperflexLicense.SmartLicensingEnabled = varHyperflexLicenseWithoutEmbeddedStruct.SmartLicensingEnabled
		varHyperflexLicense.Cluster = varHyperflexLicenseWithoutEmbeddedStruct.Cluster
		varHyperflexLicense.RegisteredDevice = varHyperflexLicenseWithoutEmbeddedStruct.RegisteredDevice
		*o = HyperflexLicense(varHyperflexLicense)
	} else {
		return err
	}

	varHyperflexLicense := _HyperflexLicense{}

	err = json.Unmarshal(bytes, &varHyperflexLicense)
	if err == nil {
		o.MoBaseMo = varHyperflexLicense.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ComplianceState")
		delete(additionalProperties, "GetOutOfComplianceStartAt")
		delete(additionalProperties, "InEvaluation")
		delete(additionalProperties, "LicenseAuthorization")
		delete(additionalProperties, "LicenseRegistration")
		delete(additionalProperties, "LicenseType")
		delete(additionalProperties, "PlrEnabled")
		delete(additionalProperties, "SmartLicensingEnabled")
		delete(additionalProperties, "Cluster")
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

type NullableHyperflexLicense struct {
	value *HyperflexLicense
	isSet bool
}

func (v NullableHyperflexLicense) Get() *HyperflexLicense {
	return v.value
}

func (v *NullableHyperflexLicense) Set(val *HyperflexLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexLicense(val *HyperflexLicense) *NullableHyperflexLicense {
	return &NullableHyperflexLicense{value: val, isSet: true}
}

func (v NullableHyperflexLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
