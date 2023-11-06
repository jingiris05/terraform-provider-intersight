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

// NiatelemetryNexusCloudSiteAllOf Definition of the list of properties defined in 'niatelemetry.NexusCloudSite', excluding properties defined in parent classes.
type NiatelemetryNexusCloudSiteAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Advisories setting status, based on license type.
	Advisories *bool `json:"Advisories,omitempty"`
	// Anomalies setting status, based on license type.
	Anomalies *bool `json:"Anomalies,omitempty"`
	// Capacity utilization setting status, based on license type.
	CapacityUtilization *bool `json:"CapacityUtilization,omitempty"`
	// Returns the license type of the device.
	LicenseType *string `json:"LicenseType,omitempty"`
	// Returns the name of the Nexus Cloud site.
	Name *string `json:"Name,omitempty"`
	// Returns the type of the Nexus Cloud site.
	SiteType *string `json:"SiteType,omitempty"`
	// Software management setting status, based on license type.
	SoftwareManagement *bool `json:"SoftwareManagement,omitempty"`
	// Returns the uuid of the Nexus Cloud site.
	Uuid                 *string                                    `json:"Uuid,omitempty"`
	NexusCloudAccount    *NiatelemetryNexusCloudAccountRelationship `json:"NexusCloudAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNexusCloudSiteAllOf NiatelemetryNexusCloudSiteAllOf

// NewNiatelemetryNexusCloudSiteAllOf instantiates a new NiatelemetryNexusCloudSiteAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNexusCloudSiteAllOf(classId string, objectType string) *NiatelemetryNexusCloudSiteAllOf {
	this := NiatelemetryNexusCloudSiteAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNexusCloudSiteAllOfWithDefaults instantiates a new NiatelemetryNexusCloudSiteAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNexusCloudSiteAllOfWithDefaults() *NiatelemetryNexusCloudSiteAllOf {
	this := NiatelemetryNexusCloudSiteAllOf{}
	var classId string = "niatelemetry.NexusCloudSite"
	this.ClassId = classId
	var objectType string = "niatelemetry.NexusCloudSite"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNexusCloudSiteAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNexusCloudSiteAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNexusCloudSiteAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNexusCloudSiteAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdvisories returns the Advisories field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSiteAllOf) GetAdvisories() bool {
	if o == nil || o.Advisories == nil {
		var ret bool
		return ret
	}
	return *o.Advisories
}

// GetAdvisoriesOk returns a tuple with the Advisories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetAdvisoriesOk() (*bool, bool) {
	if o == nil || o.Advisories == nil {
		return nil, false
	}
	return o.Advisories, true
}

// HasAdvisories returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) HasAdvisories() bool {
	if o != nil && o.Advisories != nil {
		return true
	}

	return false
}

// SetAdvisories gets a reference to the given bool and assigns it to the Advisories field.
func (o *NiatelemetryNexusCloudSiteAllOf) SetAdvisories(v bool) {
	o.Advisories = &v
}

// GetAnomalies returns the Anomalies field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSiteAllOf) GetAnomalies() bool {
	if o == nil || o.Anomalies == nil {
		var ret bool
		return ret
	}
	return *o.Anomalies
}

// GetAnomaliesOk returns a tuple with the Anomalies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetAnomaliesOk() (*bool, bool) {
	if o == nil || o.Anomalies == nil {
		return nil, false
	}
	return o.Anomalies, true
}

// HasAnomalies returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) HasAnomalies() bool {
	if o != nil && o.Anomalies != nil {
		return true
	}

	return false
}

// SetAnomalies gets a reference to the given bool and assigns it to the Anomalies field.
func (o *NiatelemetryNexusCloudSiteAllOf) SetAnomalies(v bool) {
	o.Anomalies = &v
}

// GetCapacityUtilization returns the CapacityUtilization field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSiteAllOf) GetCapacityUtilization() bool {
	if o == nil || o.CapacityUtilization == nil {
		var ret bool
		return ret
	}
	return *o.CapacityUtilization
}

// GetCapacityUtilizationOk returns a tuple with the CapacityUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetCapacityUtilizationOk() (*bool, bool) {
	if o == nil || o.CapacityUtilization == nil {
		return nil, false
	}
	return o.CapacityUtilization, true
}

// HasCapacityUtilization returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) HasCapacityUtilization() bool {
	if o != nil && o.CapacityUtilization != nil {
		return true
	}

	return false
}

// SetCapacityUtilization gets a reference to the given bool and assigns it to the CapacityUtilization field.
func (o *NiatelemetryNexusCloudSiteAllOf) SetCapacityUtilization(v bool) {
	o.CapacityUtilization = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSiteAllOf) GetLicenseType() string {
	if o == nil || o.LicenseType == nil {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetLicenseTypeOk() (*string, bool) {
	if o == nil || o.LicenseType == nil {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) HasLicenseType() bool {
	if o != nil && o.LicenseType != nil {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *NiatelemetryNexusCloudSiteAllOf) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSiteAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryNexusCloudSiteAllOf) SetName(v string) {
	o.Name = &v
}

// GetSiteType returns the SiteType field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSiteAllOf) GetSiteType() string {
	if o == nil || o.SiteType == nil {
		var ret string
		return ret
	}
	return *o.SiteType
}

// GetSiteTypeOk returns a tuple with the SiteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetSiteTypeOk() (*string, bool) {
	if o == nil || o.SiteType == nil {
		return nil, false
	}
	return o.SiteType, true
}

// HasSiteType returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) HasSiteType() bool {
	if o != nil && o.SiteType != nil {
		return true
	}

	return false
}

// SetSiteType gets a reference to the given string and assigns it to the SiteType field.
func (o *NiatelemetryNexusCloudSiteAllOf) SetSiteType(v string) {
	o.SiteType = &v
}

// GetSoftwareManagement returns the SoftwareManagement field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSiteAllOf) GetSoftwareManagement() bool {
	if o == nil || o.SoftwareManagement == nil {
		var ret bool
		return ret
	}
	return *o.SoftwareManagement
}

// GetSoftwareManagementOk returns a tuple with the SoftwareManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetSoftwareManagementOk() (*bool, bool) {
	if o == nil || o.SoftwareManagement == nil {
		return nil, false
	}
	return o.SoftwareManagement, true
}

// HasSoftwareManagement returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) HasSoftwareManagement() bool {
	if o != nil && o.SoftwareManagement != nil {
		return true
	}

	return false
}

// SetSoftwareManagement gets a reference to the given bool and assigns it to the SoftwareManagement field.
func (o *NiatelemetryNexusCloudSiteAllOf) SetSoftwareManagement(v bool) {
	o.SoftwareManagement = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSiteAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *NiatelemetryNexusCloudSiteAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetNexusCloudAccount returns the NexusCloudAccount field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudSiteAllOf) GetNexusCloudAccount() NiatelemetryNexusCloudAccountRelationship {
	if o == nil || o.NexusCloudAccount == nil {
		var ret NiatelemetryNexusCloudAccountRelationship
		return ret
	}
	return *o.NexusCloudAccount
}

// GetNexusCloudAccountOk returns a tuple with the NexusCloudAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) GetNexusCloudAccountOk() (*NiatelemetryNexusCloudAccountRelationship, bool) {
	if o == nil || o.NexusCloudAccount == nil {
		return nil, false
	}
	return o.NexusCloudAccount, true
}

// HasNexusCloudAccount returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudSiteAllOf) HasNexusCloudAccount() bool {
	if o != nil && o.NexusCloudAccount != nil {
		return true
	}

	return false
}

// SetNexusCloudAccount gets a reference to the given NiatelemetryNexusCloudAccountRelationship and assigns it to the NexusCloudAccount field.
func (o *NiatelemetryNexusCloudSiteAllOf) SetNexusCloudAccount(v NiatelemetryNexusCloudAccountRelationship) {
	o.NexusCloudAccount = &v
}

func (o NiatelemetryNexusCloudSiteAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Advisories != nil {
		toSerialize["Advisories"] = o.Advisories
	}
	if o.Anomalies != nil {
		toSerialize["Anomalies"] = o.Anomalies
	}
	if o.CapacityUtilization != nil {
		toSerialize["CapacityUtilization"] = o.CapacityUtilization
	}
	if o.LicenseType != nil {
		toSerialize["LicenseType"] = o.LicenseType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SiteType != nil {
		toSerialize["SiteType"] = o.SiteType
	}
	if o.SoftwareManagement != nil {
		toSerialize["SoftwareManagement"] = o.SoftwareManagement
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.NexusCloudAccount != nil {
		toSerialize["NexusCloudAccount"] = o.NexusCloudAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNexusCloudSiteAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNexusCloudSiteAllOf := _NiatelemetryNexusCloudSiteAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNexusCloudSiteAllOf); err == nil {
		*o = NiatelemetryNexusCloudSiteAllOf(varNiatelemetryNexusCloudSiteAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Advisories")
		delete(additionalProperties, "Anomalies")
		delete(additionalProperties, "CapacityUtilization")
		delete(additionalProperties, "LicenseType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SiteType")
		delete(additionalProperties, "SoftwareManagement")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "NexusCloudAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNexusCloudSiteAllOf struct {
	value *NiatelemetryNexusCloudSiteAllOf
	isSet bool
}

func (v NullableNiatelemetryNexusCloudSiteAllOf) Get() *NiatelemetryNexusCloudSiteAllOf {
	return v.value
}

func (v *NullableNiatelemetryNexusCloudSiteAllOf) Set(val *NiatelemetryNexusCloudSiteAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusCloudSiteAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusCloudSiteAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusCloudSiteAllOf(val *NiatelemetryNexusCloudSiteAllOf) *NullableNiatelemetryNexusCloudSiteAllOf {
	return &NullableNiatelemetryNexusCloudSiteAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusCloudSiteAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusCloudSiteAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
