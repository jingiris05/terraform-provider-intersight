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

// HyperflexExtFcStoragePolicyAllOf Definition of the list of properties defined in 'hyperflex.ExtFcStoragePolicy', excluding properties defined in parent classes.
type HyperflexExtFcStoragePolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables or disables external FC storage configuration.
	AdminState      *bool                            `json:"AdminState,omitempty"`
	ExtaTraffic     NullableHyperflexNamedVsan       `json:"ExtaTraffic,omitempty"`
	ExtbTraffic     NullableHyperflexNamedVsan       `json:"ExtbTraffic,omitempty"`
	WwxnPrefixRange NullableHyperflexWwxnPrefixRange `json:"WwxnPrefixRange,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexExtFcStoragePolicyAllOf HyperflexExtFcStoragePolicyAllOf

// NewHyperflexExtFcStoragePolicyAllOf instantiates a new HyperflexExtFcStoragePolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexExtFcStoragePolicyAllOf(classId string, objectType string) *HyperflexExtFcStoragePolicyAllOf {
	this := HyperflexExtFcStoragePolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexExtFcStoragePolicyAllOfWithDefaults instantiates a new HyperflexExtFcStoragePolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexExtFcStoragePolicyAllOfWithDefaults() *HyperflexExtFcStoragePolicyAllOf {
	this := HyperflexExtFcStoragePolicyAllOf{}
	var classId string = "hyperflex.ExtFcStoragePolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.ExtFcStoragePolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexExtFcStoragePolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexExtFcStoragePolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexExtFcStoragePolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexExtFcStoragePolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *HyperflexExtFcStoragePolicyAllOf) GetAdminState() bool {
	if o == nil || o.AdminState == nil {
		var ret bool
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) GetAdminStateOk() (*bool, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given bool and assigns it to the AdminState field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetAdminState(v bool) {
	o.AdminState = &v
}

// GetExtaTraffic returns the ExtaTraffic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexExtFcStoragePolicyAllOf) GetExtaTraffic() HyperflexNamedVsan {
	if o == nil || o.ExtaTraffic.Get() == nil {
		var ret HyperflexNamedVsan
		return ret
	}
	return *o.ExtaTraffic.Get()
}

// GetExtaTrafficOk returns a tuple with the ExtaTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexExtFcStoragePolicyAllOf) GetExtaTrafficOk() (*HyperflexNamedVsan, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtaTraffic.Get(), o.ExtaTraffic.IsSet()
}

// HasExtaTraffic returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasExtaTraffic() bool {
	if o != nil && o.ExtaTraffic.IsSet() {
		return true
	}

	return false
}

// SetExtaTraffic gets a reference to the given NullableHyperflexNamedVsan and assigns it to the ExtaTraffic field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetExtaTraffic(v HyperflexNamedVsan) {
	o.ExtaTraffic.Set(&v)
}

// SetExtaTrafficNil sets the value for ExtaTraffic to be an explicit nil
func (o *HyperflexExtFcStoragePolicyAllOf) SetExtaTrafficNil() {
	o.ExtaTraffic.Set(nil)
}

// UnsetExtaTraffic ensures that no value is present for ExtaTraffic, not even an explicit nil
func (o *HyperflexExtFcStoragePolicyAllOf) UnsetExtaTraffic() {
	o.ExtaTraffic.Unset()
}

// GetExtbTraffic returns the ExtbTraffic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexExtFcStoragePolicyAllOf) GetExtbTraffic() HyperflexNamedVsan {
	if o == nil || o.ExtbTraffic.Get() == nil {
		var ret HyperflexNamedVsan
		return ret
	}
	return *o.ExtbTraffic.Get()
}

// GetExtbTrafficOk returns a tuple with the ExtbTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexExtFcStoragePolicyAllOf) GetExtbTrafficOk() (*HyperflexNamedVsan, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtbTraffic.Get(), o.ExtbTraffic.IsSet()
}

// HasExtbTraffic returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasExtbTraffic() bool {
	if o != nil && o.ExtbTraffic.IsSet() {
		return true
	}

	return false
}

// SetExtbTraffic gets a reference to the given NullableHyperflexNamedVsan and assigns it to the ExtbTraffic field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetExtbTraffic(v HyperflexNamedVsan) {
	o.ExtbTraffic.Set(&v)
}

// SetExtbTrafficNil sets the value for ExtbTraffic to be an explicit nil
func (o *HyperflexExtFcStoragePolicyAllOf) SetExtbTrafficNil() {
	o.ExtbTraffic.Set(nil)
}

// UnsetExtbTraffic ensures that no value is present for ExtbTraffic, not even an explicit nil
func (o *HyperflexExtFcStoragePolicyAllOf) UnsetExtbTraffic() {
	o.ExtbTraffic.Unset()
}

// GetWwxnPrefixRange returns the WwxnPrefixRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexExtFcStoragePolicyAllOf) GetWwxnPrefixRange() HyperflexWwxnPrefixRange {
	if o == nil || o.WwxnPrefixRange.Get() == nil {
		var ret HyperflexWwxnPrefixRange
		return ret
	}
	return *o.WwxnPrefixRange.Get()
}

// GetWwxnPrefixRangeOk returns a tuple with the WwxnPrefixRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexExtFcStoragePolicyAllOf) GetWwxnPrefixRangeOk() (*HyperflexWwxnPrefixRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.WwxnPrefixRange.Get(), o.WwxnPrefixRange.IsSet()
}

// HasWwxnPrefixRange returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasWwxnPrefixRange() bool {
	if o != nil && o.WwxnPrefixRange.IsSet() {
		return true
	}

	return false
}

// SetWwxnPrefixRange gets a reference to the given NullableHyperflexWwxnPrefixRange and assigns it to the WwxnPrefixRange field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetWwxnPrefixRange(v HyperflexWwxnPrefixRange) {
	o.WwxnPrefixRange.Set(&v)
}

// SetWwxnPrefixRangeNil sets the value for WwxnPrefixRange to be an explicit nil
func (o *HyperflexExtFcStoragePolicyAllOf) SetWwxnPrefixRangeNil() {
	o.WwxnPrefixRange.Set(nil)
}

// UnsetWwxnPrefixRange ensures that no value is present for WwxnPrefixRange, not even an explicit nil
func (o *HyperflexExtFcStoragePolicyAllOf) UnsetWwxnPrefixRange() {
	o.WwxnPrefixRange.Unset()
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexExtFcStoragePolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexExtFcStoragePolicyAllOf) GetClusterProfilesOk() ([]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexExtFcStoragePolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexExtFcStoragePolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexExtFcStoragePolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexExtFcStoragePolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.ExtaTraffic.IsSet() {
		toSerialize["ExtaTraffic"] = o.ExtaTraffic.Get()
	}
	if o.ExtbTraffic.IsSet() {
		toSerialize["ExtbTraffic"] = o.ExtbTraffic.Get()
	}
	if o.WwxnPrefixRange.IsSet() {
		toSerialize["WwxnPrefixRange"] = o.WwxnPrefixRange.Get()
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexExtFcStoragePolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexExtFcStoragePolicyAllOf := _HyperflexExtFcStoragePolicyAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexExtFcStoragePolicyAllOf); err == nil {
		*o = HyperflexExtFcStoragePolicyAllOf(varHyperflexExtFcStoragePolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "ExtaTraffic")
		delete(additionalProperties, "ExtbTraffic")
		delete(additionalProperties, "WwxnPrefixRange")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexExtFcStoragePolicyAllOf struct {
	value *HyperflexExtFcStoragePolicyAllOf
	isSet bool
}

func (v NullableHyperflexExtFcStoragePolicyAllOf) Get() *HyperflexExtFcStoragePolicyAllOf {
	return v.value
}

func (v *NullableHyperflexExtFcStoragePolicyAllOf) Set(val *HyperflexExtFcStoragePolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexExtFcStoragePolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexExtFcStoragePolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexExtFcStoragePolicyAllOf(val *HyperflexExtFcStoragePolicyAllOf) *NullableHyperflexExtFcStoragePolicyAllOf {
	return &NullableHyperflexExtFcStoragePolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexExtFcStoragePolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexExtFcStoragePolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
