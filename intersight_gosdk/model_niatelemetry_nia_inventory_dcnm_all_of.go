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
)

// NiatelemetryNiaInventoryDcnmAllOf Definition of the list of properties defined in 'niatelemetry.NiaInventoryDcnm', excluding properties defined in parent classes.
type NiatelemetryNiaInventoryDcnmAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the value of the dev Field.
	Dev *bool `json:"Dev,omitempty"`
	// Number of EPLD images uploaded to DCNM.
	EpldImageCount *int64 `json:"EpldImageCount,omitempty"`
	// Returns the value of the haEnabled field.
	HaEnabled *bool `json:"HaEnabled,omitempty"`
	// Returns the value of the haReplicationStatus field.
	HaReplicationStatus *string `json:"HaReplicationStatus,omitempty"`
	// Returns the value of the install field.
	Install *string `json:"Install,omitempty"`
	// Returns true if ISN is configured.
	IsIsnConfigured *bool `json:"IsIsnConfigured,omitempty"`
	// Returns the value of the isMediaController field.
	IsMediaController *bool `json:"IsMediaController,omitempty"`
	// Returns true if the Smart license is enabled and is in use.
	IsSmartLicenseEnabled *bool `json:"IsSmartLicenseEnabled,omitempty"`
	// Returns total number of fabrics in DCNM set-up.
	NumFabrics *int64 `json:"NumFabrics,omitempty"`
	// Returns the number of fabrics in msd.
	NumFabricsInMsd *int64 `json:"NumFabricsInMsd,omitempty"`
	// Returns the number of fabrics that have ingress replication type.
	NumIngressReplicationFabrics *int64 `json:"NumIngressReplicationFabrics,omitempty"`
	// Returns the number of local users other than admin user.
	NumLocalUsers *int64 `json:"NumLocalUsers,omitempty"`
	// Returns the number of MSD fabrics.
	NumMsd *int64 `json:"NumMsd,omitempty"`
	// Returns the number of svi interfaces configured for VRF vlans.
	NumSviVrfCount *int64 `json:"NumSviVrfCount,omitempty"`
	// Returns the number of links where TRM is enabled.
	NumTrmEnabledCount *int64 `json:"NumTrmEnabledCount,omitempty"`
	// Number of users who have upgrade privileges excluding the admin.
	NumUpgUsers *int64 `json:"NumUpgUsers,omitempty"`
	// Number of NXOS images uploaded to DCNM.
	NxosImageCount *int64 `json:"NxosImageCount,omitempty"`
	// Serial number of device being inventoried. The serial number is unique per device.
	Serial *string `json:"Serial,omitempty"`
	// Name of fabric domain of the controller.
	SiteName *string `json:"SiteName,omitempty"`
	// Returns the number of underlay peering active links.
	UnderlayPeeringActiveLinksCount *int64 `json:"UnderlayPeeringActiveLinksCount,omitempty"`
	// Number of upgrade jobs configured on DCNM.
	UpgJobCount *int64 `json:"UpgJobCount,omitempty"`
	// Upgrade status of jobs created on DCNM.
	UpgStatus *string `json:"UpgStatus,omitempty"`
	// Returns the value of the version field.
	Version              *string                              `json:"Version,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNiaInventoryDcnmAllOf NiatelemetryNiaInventoryDcnmAllOf

// NewNiatelemetryNiaInventoryDcnmAllOf instantiates a new NiatelemetryNiaInventoryDcnmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNiaInventoryDcnmAllOf(classId string, objectType string) *NiatelemetryNiaInventoryDcnmAllOf {
	this := NiatelemetryNiaInventoryDcnmAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNiaInventoryDcnmAllOfWithDefaults instantiates a new NiatelemetryNiaInventoryDcnmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNiaInventoryDcnmAllOfWithDefaults() *NiatelemetryNiaInventoryDcnmAllOf {
	this := NiatelemetryNiaInventoryDcnmAllOf{}
	var classId string = "niatelemetry.NiaInventoryDcnm"
	this.ClassId = classId
	var objectType string = "niatelemetry.NiaInventoryDcnm"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDev returns the Dev field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetDev() bool {
	if o == nil || o.Dev == nil {
		var ret bool
		return ret
	}
	return *o.Dev
}

// GetDevOk returns a tuple with the Dev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetDevOk() (*bool, bool) {
	if o == nil || o.Dev == nil {
		return nil, false
	}
	return o.Dev, true
}

// HasDev returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasDev() bool {
	if o != nil && o.Dev != nil {
		return true
	}

	return false
}

// SetDev gets a reference to the given bool and assigns it to the Dev field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetDev(v bool) {
	o.Dev = &v
}

// GetEpldImageCount returns the EpldImageCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetEpldImageCount() int64 {
	if o == nil || o.EpldImageCount == nil {
		var ret int64
		return ret
	}
	return *o.EpldImageCount
}

// GetEpldImageCountOk returns a tuple with the EpldImageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetEpldImageCountOk() (*int64, bool) {
	if o == nil || o.EpldImageCount == nil {
		return nil, false
	}
	return o.EpldImageCount, true
}

// HasEpldImageCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasEpldImageCount() bool {
	if o != nil && o.EpldImageCount != nil {
		return true
	}

	return false
}

// SetEpldImageCount gets a reference to the given int64 and assigns it to the EpldImageCount field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetEpldImageCount(v int64) {
	o.EpldImageCount = &v
}

// GetHaEnabled returns the HaEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaEnabled() bool {
	if o == nil || o.HaEnabled == nil {
		var ret bool
		return ret
	}
	return *o.HaEnabled
}

// GetHaEnabledOk returns a tuple with the HaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaEnabledOk() (*bool, bool) {
	if o == nil || o.HaEnabled == nil {
		return nil, false
	}
	return o.HaEnabled, true
}

// HasHaEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasHaEnabled() bool {
	if o != nil && o.HaEnabled != nil {
		return true
	}

	return false
}

// SetHaEnabled gets a reference to the given bool and assigns it to the HaEnabled field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetHaEnabled(v bool) {
	o.HaEnabled = &v
}

// GetHaReplicationStatus returns the HaReplicationStatus field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaReplicationStatus() string {
	if o == nil || o.HaReplicationStatus == nil {
		var ret string
		return ret
	}
	return *o.HaReplicationStatus
}

// GetHaReplicationStatusOk returns a tuple with the HaReplicationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaReplicationStatusOk() (*string, bool) {
	if o == nil || o.HaReplicationStatus == nil {
		return nil, false
	}
	return o.HaReplicationStatus, true
}

// HasHaReplicationStatus returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasHaReplicationStatus() bool {
	if o != nil && o.HaReplicationStatus != nil {
		return true
	}

	return false
}

// SetHaReplicationStatus gets a reference to the given string and assigns it to the HaReplicationStatus field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetHaReplicationStatus(v string) {
	o.HaReplicationStatus = &v
}

// GetInstall returns the Install field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetInstall() string {
	if o == nil || o.Install == nil {
		var ret string
		return ret
	}
	return *o.Install
}

// GetInstallOk returns a tuple with the Install field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetInstallOk() (*string, bool) {
	if o == nil || o.Install == nil {
		return nil, false
	}
	return o.Install, true
}

// HasInstall returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasInstall() bool {
	if o != nil && o.Install != nil {
		return true
	}

	return false
}

// SetInstall gets a reference to the given string and assigns it to the Install field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetInstall(v string) {
	o.Install = &v
}

// GetIsIsnConfigured returns the IsIsnConfigured field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsIsnConfigured() bool {
	if o == nil || o.IsIsnConfigured == nil {
		var ret bool
		return ret
	}
	return *o.IsIsnConfigured
}

// GetIsIsnConfiguredOk returns a tuple with the IsIsnConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsIsnConfiguredOk() (*bool, bool) {
	if o == nil || o.IsIsnConfigured == nil {
		return nil, false
	}
	return o.IsIsnConfigured, true
}

// HasIsIsnConfigured returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasIsIsnConfigured() bool {
	if o != nil && o.IsIsnConfigured != nil {
		return true
	}

	return false
}

// SetIsIsnConfigured gets a reference to the given bool and assigns it to the IsIsnConfigured field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetIsIsnConfigured(v bool) {
	o.IsIsnConfigured = &v
}

// GetIsMediaController returns the IsMediaController field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsMediaController() bool {
	if o == nil || o.IsMediaController == nil {
		var ret bool
		return ret
	}
	return *o.IsMediaController
}

// GetIsMediaControllerOk returns a tuple with the IsMediaController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsMediaControllerOk() (*bool, bool) {
	if o == nil || o.IsMediaController == nil {
		return nil, false
	}
	return o.IsMediaController, true
}

// HasIsMediaController returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasIsMediaController() bool {
	if o != nil && o.IsMediaController != nil {
		return true
	}

	return false
}

// SetIsMediaController gets a reference to the given bool and assigns it to the IsMediaController field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetIsMediaController(v bool) {
	o.IsMediaController = &v
}

// GetIsSmartLicenseEnabled returns the IsSmartLicenseEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsSmartLicenseEnabled() bool {
	if o == nil || o.IsSmartLicenseEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsSmartLicenseEnabled
}

// GetIsSmartLicenseEnabledOk returns a tuple with the IsSmartLicenseEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsSmartLicenseEnabledOk() (*bool, bool) {
	if o == nil || o.IsSmartLicenseEnabled == nil {
		return nil, false
	}
	return o.IsSmartLicenseEnabled, true
}

// HasIsSmartLicenseEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasIsSmartLicenseEnabled() bool {
	if o != nil && o.IsSmartLicenseEnabled != nil {
		return true
	}

	return false
}

// SetIsSmartLicenseEnabled gets a reference to the given bool and assigns it to the IsSmartLicenseEnabled field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetIsSmartLicenseEnabled(v bool) {
	o.IsSmartLicenseEnabled = &v
}

// GetNumFabrics returns the NumFabrics field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabrics() int64 {
	if o == nil || o.NumFabrics == nil {
		var ret int64
		return ret
	}
	return *o.NumFabrics
}

// GetNumFabricsOk returns a tuple with the NumFabrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabricsOk() (*int64, bool) {
	if o == nil || o.NumFabrics == nil {
		return nil, false
	}
	return o.NumFabrics, true
}

// HasNumFabrics returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumFabrics() bool {
	if o != nil && o.NumFabrics != nil {
		return true
	}

	return false
}

// SetNumFabrics gets a reference to the given int64 and assigns it to the NumFabrics field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumFabrics(v int64) {
	o.NumFabrics = &v
}

// GetNumFabricsInMsd returns the NumFabricsInMsd field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabricsInMsd() int64 {
	if o == nil || o.NumFabricsInMsd == nil {
		var ret int64
		return ret
	}
	return *o.NumFabricsInMsd
}

// GetNumFabricsInMsdOk returns a tuple with the NumFabricsInMsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabricsInMsdOk() (*int64, bool) {
	if o == nil || o.NumFabricsInMsd == nil {
		return nil, false
	}
	return o.NumFabricsInMsd, true
}

// HasNumFabricsInMsd returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumFabricsInMsd() bool {
	if o != nil && o.NumFabricsInMsd != nil {
		return true
	}

	return false
}

// SetNumFabricsInMsd gets a reference to the given int64 and assigns it to the NumFabricsInMsd field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumFabricsInMsd(v int64) {
	o.NumFabricsInMsd = &v
}

// GetNumIngressReplicationFabrics returns the NumIngressReplicationFabrics field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumIngressReplicationFabrics() int64 {
	if o == nil || o.NumIngressReplicationFabrics == nil {
		var ret int64
		return ret
	}
	return *o.NumIngressReplicationFabrics
}

// GetNumIngressReplicationFabricsOk returns a tuple with the NumIngressReplicationFabrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumIngressReplicationFabricsOk() (*int64, bool) {
	if o == nil || o.NumIngressReplicationFabrics == nil {
		return nil, false
	}
	return o.NumIngressReplicationFabrics, true
}

// HasNumIngressReplicationFabrics returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumIngressReplicationFabrics() bool {
	if o != nil && o.NumIngressReplicationFabrics != nil {
		return true
	}

	return false
}

// SetNumIngressReplicationFabrics gets a reference to the given int64 and assigns it to the NumIngressReplicationFabrics field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumIngressReplicationFabrics(v int64) {
	o.NumIngressReplicationFabrics = &v
}

// GetNumLocalUsers returns the NumLocalUsers field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumLocalUsers() int64 {
	if o == nil || o.NumLocalUsers == nil {
		var ret int64
		return ret
	}
	return *o.NumLocalUsers
}

// GetNumLocalUsersOk returns a tuple with the NumLocalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumLocalUsersOk() (*int64, bool) {
	if o == nil || o.NumLocalUsers == nil {
		return nil, false
	}
	return o.NumLocalUsers, true
}

// HasNumLocalUsers returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumLocalUsers() bool {
	if o != nil && o.NumLocalUsers != nil {
		return true
	}

	return false
}

// SetNumLocalUsers gets a reference to the given int64 and assigns it to the NumLocalUsers field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumLocalUsers(v int64) {
	o.NumLocalUsers = &v
}

// GetNumMsd returns the NumMsd field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumMsd() int64 {
	if o == nil || o.NumMsd == nil {
		var ret int64
		return ret
	}
	return *o.NumMsd
}

// GetNumMsdOk returns a tuple with the NumMsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumMsdOk() (*int64, bool) {
	if o == nil || o.NumMsd == nil {
		return nil, false
	}
	return o.NumMsd, true
}

// HasNumMsd returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumMsd() bool {
	if o != nil && o.NumMsd != nil {
		return true
	}

	return false
}

// SetNumMsd gets a reference to the given int64 and assigns it to the NumMsd field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumMsd(v int64) {
	o.NumMsd = &v
}

// GetNumSviVrfCount returns the NumSviVrfCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumSviVrfCount() int64 {
	if o == nil || o.NumSviVrfCount == nil {
		var ret int64
		return ret
	}
	return *o.NumSviVrfCount
}

// GetNumSviVrfCountOk returns a tuple with the NumSviVrfCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumSviVrfCountOk() (*int64, bool) {
	if o == nil || o.NumSviVrfCount == nil {
		return nil, false
	}
	return o.NumSviVrfCount, true
}

// HasNumSviVrfCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumSviVrfCount() bool {
	if o != nil && o.NumSviVrfCount != nil {
		return true
	}

	return false
}

// SetNumSviVrfCount gets a reference to the given int64 and assigns it to the NumSviVrfCount field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumSviVrfCount(v int64) {
	o.NumSviVrfCount = &v
}

// GetNumTrmEnabledCount returns the NumTrmEnabledCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumTrmEnabledCount() int64 {
	if o == nil || o.NumTrmEnabledCount == nil {
		var ret int64
		return ret
	}
	return *o.NumTrmEnabledCount
}

// GetNumTrmEnabledCountOk returns a tuple with the NumTrmEnabledCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumTrmEnabledCountOk() (*int64, bool) {
	if o == nil || o.NumTrmEnabledCount == nil {
		return nil, false
	}
	return o.NumTrmEnabledCount, true
}

// HasNumTrmEnabledCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumTrmEnabledCount() bool {
	if o != nil && o.NumTrmEnabledCount != nil {
		return true
	}

	return false
}

// SetNumTrmEnabledCount gets a reference to the given int64 and assigns it to the NumTrmEnabledCount field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumTrmEnabledCount(v int64) {
	o.NumTrmEnabledCount = &v
}

// GetNumUpgUsers returns the NumUpgUsers field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumUpgUsers() int64 {
	if o == nil || o.NumUpgUsers == nil {
		var ret int64
		return ret
	}
	return *o.NumUpgUsers
}

// GetNumUpgUsersOk returns a tuple with the NumUpgUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumUpgUsersOk() (*int64, bool) {
	if o == nil || o.NumUpgUsers == nil {
		return nil, false
	}
	return o.NumUpgUsers, true
}

// HasNumUpgUsers returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumUpgUsers() bool {
	if o != nil && o.NumUpgUsers != nil {
		return true
	}

	return false
}

// SetNumUpgUsers gets a reference to the given int64 and assigns it to the NumUpgUsers field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumUpgUsers(v int64) {
	o.NumUpgUsers = &v
}

// GetNxosImageCount returns the NxosImageCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNxosImageCount() int64 {
	if o == nil || o.NxosImageCount == nil {
		var ret int64
		return ret
	}
	return *o.NxosImageCount
}

// GetNxosImageCountOk returns a tuple with the NxosImageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNxosImageCountOk() (*int64, bool) {
	if o == nil || o.NxosImageCount == nil {
		return nil, false
	}
	return o.NxosImageCount, true
}

// HasNxosImageCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNxosImageCount() bool {
	if o != nil && o.NxosImageCount != nil {
		return true
	}

	return false
}

// SetNxosImageCount gets a reference to the given int64 and assigns it to the NxosImageCount field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNxosImageCount(v int64) {
	o.NxosImageCount = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetUnderlayPeeringActiveLinksCount returns the UnderlayPeeringActiveLinksCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUnderlayPeeringActiveLinksCount() int64 {
	if o == nil || o.UnderlayPeeringActiveLinksCount == nil {
		var ret int64
		return ret
	}
	return *o.UnderlayPeeringActiveLinksCount
}

// GetUnderlayPeeringActiveLinksCountOk returns a tuple with the UnderlayPeeringActiveLinksCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUnderlayPeeringActiveLinksCountOk() (*int64, bool) {
	if o == nil || o.UnderlayPeeringActiveLinksCount == nil {
		return nil, false
	}
	return o.UnderlayPeeringActiveLinksCount, true
}

// HasUnderlayPeeringActiveLinksCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasUnderlayPeeringActiveLinksCount() bool {
	if o != nil && o.UnderlayPeeringActiveLinksCount != nil {
		return true
	}

	return false
}

// SetUnderlayPeeringActiveLinksCount gets a reference to the given int64 and assigns it to the UnderlayPeeringActiveLinksCount field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetUnderlayPeeringActiveLinksCount(v int64) {
	o.UnderlayPeeringActiveLinksCount = &v
}

// GetUpgJobCount returns the UpgJobCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUpgJobCount() int64 {
	if o == nil || o.UpgJobCount == nil {
		var ret int64
		return ret
	}
	return *o.UpgJobCount
}

// GetUpgJobCountOk returns a tuple with the UpgJobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUpgJobCountOk() (*int64, bool) {
	if o == nil || o.UpgJobCount == nil {
		return nil, false
	}
	return o.UpgJobCount, true
}

// HasUpgJobCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasUpgJobCount() bool {
	if o != nil && o.UpgJobCount != nil {
		return true
	}

	return false
}

// SetUpgJobCount gets a reference to the given int64 and assigns it to the UpgJobCount field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetUpgJobCount(v int64) {
	o.UpgJobCount = &v
}

// GetUpgStatus returns the UpgStatus field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUpgStatus() string {
	if o == nil || o.UpgStatus == nil {
		var ret string
		return ret
	}
	return *o.UpgStatus
}

// GetUpgStatusOk returns a tuple with the UpgStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetUpgStatusOk() (*string, bool) {
	if o == nil || o.UpgStatus == nil {
		return nil, false
	}
	return o.UpgStatus, true
}

// HasUpgStatus returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasUpgStatus() bool {
	if o != nil && o.UpgStatus != nil {
		return true
	}

	return false
}

// SetUpgStatus gets a reference to the given string and assigns it to the UpgStatus field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetUpgStatus(v string) {
	o.UpgStatus = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNiaInventoryDcnmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dev != nil {
		toSerialize["Dev"] = o.Dev
	}
	if o.EpldImageCount != nil {
		toSerialize["EpldImageCount"] = o.EpldImageCount
	}
	if o.HaEnabled != nil {
		toSerialize["HaEnabled"] = o.HaEnabled
	}
	if o.HaReplicationStatus != nil {
		toSerialize["HaReplicationStatus"] = o.HaReplicationStatus
	}
	if o.Install != nil {
		toSerialize["Install"] = o.Install
	}
	if o.IsIsnConfigured != nil {
		toSerialize["IsIsnConfigured"] = o.IsIsnConfigured
	}
	if o.IsMediaController != nil {
		toSerialize["IsMediaController"] = o.IsMediaController
	}
	if o.IsSmartLicenseEnabled != nil {
		toSerialize["IsSmartLicenseEnabled"] = o.IsSmartLicenseEnabled
	}
	if o.NumFabrics != nil {
		toSerialize["NumFabrics"] = o.NumFabrics
	}
	if o.NumFabricsInMsd != nil {
		toSerialize["NumFabricsInMsd"] = o.NumFabricsInMsd
	}
	if o.NumIngressReplicationFabrics != nil {
		toSerialize["NumIngressReplicationFabrics"] = o.NumIngressReplicationFabrics
	}
	if o.NumLocalUsers != nil {
		toSerialize["NumLocalUsers"] = o.NumLocalUsers
	}
	if o.NumMsd != nil {
		toSerialize["NumMsd"] = o.NumMsd
	}
	if o.NumSviVrfCount != nil {
		toSerialize["NumSviVrfCount"] = o.NumSviVrfCount
	}
	if o.NumTrmEnabledCount != nil {
		toSerialize["NumTrmEnabledCount"] = o.NumTrmEnabledCount
	}
	if o.NumUpgUsers != nil {
		toSerialize["NumUpgUsers"] = o.NumUpgUsers
	}
	if o.NxosImageCount != nil {
		toSerialize["NxosImageCount"] = o.NxosImageCount
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.UnderlayPeeringActiveLinksCount != nil {
		toSerialize["UnderlayPeeringActiveLinksCount"] = o.UnderlayPeeringActiveLinksCount
	}
	if o.UpgJobCount != nil {
		toSerialize["UpgJobCount"] = o.UpgJobCount
	}
	if o.UpgStatus != nil {
		toSerialize["UpgStatus"] = o.UpgStatus
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNiaInventoryDcnmAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNiaInventoryDcnmAllOf := _NiatelemetryNiaInventoryDcnmAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNiaInventoryDcnmAllOf); err == nil {
		*o = NiatelemetryNiaInventoryDcnmAllOf(varNiatelemetryNiaInventoryDcnmAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dev")
		delete(additionalProperties, "EpldImageCount")
		delete(additionalProperties, "HaEnabled")
		delete(additionalProperties, "HaReplicationStatus")
		delete(additionalProperties, "Install")
		delete(additionalProperties, "IsIsnConfigured")
		delete(additionalProperties, "IsMediaController")
		delete(additionalProperties, "IsSmartLicenseEnabled")
		delete(additionalProperties, "NumFabrics")
		delete(additionalProperties, "NumFabricsInMsd")
		delete(additionalProperties, "NumIngressReplicationFabrics")
		delete(additionalProperties, "NumLocalUsers")
		delete(additionalProperties, "NumMsd")
		delete(additionalProperties, "NumSviVrfCount")
		delete(additionalProperties, "NumTrmEnabledCount")
		delete(additionalProperties, "NumUpgUsers")
		delete(additionalProperties, "NxosImageCount")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "UnderlayPeeringActiveLinksCount")
		delete(additionalProperties, "UpgJobCount")
		delete(additionalProperties, "UpgStatus")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNiaInventoryDcnmAllOf struct {
	value *NiatelemetryNiaInventoryDcnmAllOf
	isSet bool
}

func (v NullableNiatelemetryNiaInventoryDcnmAllOf) Get() *NiatelemetryNiaInventoryDcnmAllOf {
	return v.value
}

func (v *NullableNiatelemetryNiaInventoryDcnmAllOf) Set(val *NiatelemetryNiaInventoryDcnmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiaInventoryDcnmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiaInventoryDcnmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiaInventoryDcnmAllOf(val *NiatelemetryNiaInventoryDcnmAllOf) *NullableNiatelemetryNiaInventoryDcnmAllOf {
	return &NullableNiatelemetryNiaInventoryDcnmAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNiaInventoryDcnmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiaInventoryDcnmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
