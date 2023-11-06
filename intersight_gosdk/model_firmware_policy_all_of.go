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

// FirmwarePolicyAllOf Definition of the list of properties defined in 'firmware.Policy', excluding properties defined in parent classes.
type FirmwarePolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                       `json:"ObjectType"`
	ExcludeComponentList []string                     `json:"ExcludeComponentList,omitempty"`
	ModelBundleCombo     []FirmwareModelBundleVersion `json:"ModelBundleCombo,omitempty"`
	// The target platform on which the policy to be applied. Either standalone or connected. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	TargetPlatform *string                               `json:"TargetPlatform,omitempty"`
	Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwarePolicyAllOf FirmwarePolicyAllOf

// NewFirmwarePolicyAllOf instantiates a new FirmwarePolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwarePolicyAllOf(classId string, objectType string) *FirmwarePolicyAllOf {
	this := FirmwarePolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewFirmwarePolicyAllOfWithDefaults instantiates a new FirmwarePolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwarePolicyAllOfWithDefaults() *FirmwarePolicyAllOf {
	this := FirmwarePolicyAllOf{}
	var classId string = "firmware.Policy"
	this.ClassId = classId
	var objectType string = "firmware.Policy"
	this.ObjectType = objectType
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwarePolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwarePolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwarePolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwarePolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwarePolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwarePolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExcludeComponentList returns the ExcludeComponentList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwarePolicyAllOf) GetExcludeComponentList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeComponentList
}

// GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwarePolicyAllOf) GetExcludeComponentListOk() ([]string, bool) {
	if o == nil || o.ExcludeComponentList == nil {
		return nil, false
	}
	return o.ExcludeComponentList, true
}

// HasExcludeComponentList returns a boolean if a field has been set.
func (o *FirmwarePolicyAllOf) HasExcludeComponentList() bool {
	if o != nil && o.ExcludeComponentList != nil {
		return true
	}

	return false
}

// SetExcludeComponentList gets a reference to the given []string and assigns it to the ExcludeComponentList field.
func (o *FirmwarePolicyAllOf) SetExcludeComponentList(v []string) {
	o.ExcludeComponentList = v
}

// GetModelBundleCombo returns the ModelBundleCombo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwarePolicyAllOf) GetModelBundleCombo() []FirmwareModelBundleVersion {
	if o == nil {
		var ret []FirmwareModelBundleVersion
		return ret
	}
	return o.ModelBundleCombo
}

// GetModelBundleComboOk returns a tuple with the ModelBundleCombo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwarePolicyAllOf) GetModelBundleComboOk() ([]FirmwareModelBundleVersion, bool) {
	if o == nil || o.ModelBundleCombo == nil {
		return nil, false
	}
	return o.ModelBundleCombo, true
}

// HasModelBundleCombo returns a boolean if a field has been set.
func (o *FirmwarePolicyAllOf) HasModelBundleCombo() bool {
	if o != nil && o.ModelBundleCombo != nil {
		return true
	}

	return false
}

// SetModelBundleCombo gets a reference to the given []FirmwareModelBundleVersion and assigns it to the ModelBundleCombo field.
func (o *FirmwarePolicyAllOf) SetModelBundleCombo(v []FirmwareModelBundleVersion) {
	o.ModelBundleCombo = v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *FirmwarePolicyAllOf) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwarePolicyAllOf) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *FirmwarePolicyAllOf) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *FirmwarePolicyAllOf) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FirmwarePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwarePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FirmwarePolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FirmwarePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwarePolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwarePolicyAllOf) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FirmwarePolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *FirmwarePolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o FirmwarePolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExcludeComponentList != nil {
		toSerialize["ExcludeComponentList"] = o.ExcludeComponentList
	}
	if o.ModelBundleCombo != nil {
		toSerialize["ModelBundleCombo"] = o.ModelBundleCombo
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwarePolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwarePolicyAllOf := _FirmwarePolicyAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwarePolicyAllOf); err == nil {
		*o = FirmwarePolicyAllOf(varFirmwarePolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExcludeComponentList")
		delete(additionalProperties, "ModelBundleCombo")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwarePolicyAllOf struct {
	value *FirmwarePolicyAllOf
	isSet bool
}

func (v NullableFirmwarePolicyAllOf) Get() *FirmwarePolicyAllOf {
	return v.value
}

func (v *NullableFirmwarePolicyAllOf) Set(val *FirmwarePolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwarePolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwarePolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwarePolicyAllOf(val *FirmwarePolicyAllOf) *NullableFirmwarePolicyAllOf {
	return &NullableFirmwarePolicyAllOf{value: val, isSet: true}
}

func (v NullableFirmwarePolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwarePolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
