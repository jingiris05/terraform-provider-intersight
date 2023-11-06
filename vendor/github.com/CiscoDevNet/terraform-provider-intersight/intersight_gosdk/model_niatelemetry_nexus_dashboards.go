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

// NiatelemetryNexusDashboards Object is available for Nexus Dashboard devices.
type NiatelemetryNexusDashboards struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Nexus Dashboard can onboard multiple APIC clusters/sites.
	ClusterName *string `json:"ClusterName,omitempty"`
	// UUID of the Nexus Dashboard cluster.
	ClusterUuid *string `json:"ClusterUuid,omitempty"`
	// Dn of the objects present for Nexus Dashboard devices.
	Dn *string `json:"Dn,omitempty"`
	// Health of Nexus Dashboard cluster.
	IsClusterHealthy *string `json:"IsClusterHealthy,omitempty"`
	// Number of nodes in Nexus Dashboard cluster.
	NdClusterSize *int64              `json:"NdClusterSize,omitempty"`
	NdSites       []NiatelemetrySites `json:"NdSites,omitempty"`
	// Node type in Nexus Dashboard cluster.
	NdType *string `json:"NdType,omitempty"`
	// Version running on Nexus Dashboard.
	NdVersion *string `json:"NdVersion,omitempty"`
	// Number of applications installed in the Nexus Dashboard.
	NumberOfApps *int64 `json:"NumberOfApps,omitempty"`
	// Number of total insight groups in ND.
	NumberOfInsightGroups *int64 `json:"NumberOfInsightGroups,omitempty"`
	// Number of total NIR dashboards in ND.
	NumberOfNirDashboards *int64 `json:"NumberOfNirDashboards,omitempty"`
	// Number of total schemas in Multi-Site Orchestrator.
	NumberOfSchemasInMso *int64 `json:"NumberOfSchemasInMso,omitempty"`
	// Number of sites in Multi-Site Orchestrator.
	NumberOfSitesInMso *int64 `json:"NumberOfSitesInMso,omitempty"`
	// Number of sites serviced by ND.
	NumberOfSitesServiced *int64 `json:"NumberOfSitesServiced,omitempty"`
	// Number of total tenants in Multi-Site Orchestrator.
	NumberOfTenantsInMso *int64 `json:"NumberOfTenantsInMso,omitempty"`
	// Number of sites with vxLan type fabric in Multi-Site Orchestrator.
	NumberOfVxlanFabricSitesInMso *int64 `json:"NumberOfVxlanFabricSitesInMso,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Type of site added to Multi-Site Orchestrator.
	TypeOfSiteInMso      *string                              `json:"TypeOfSiteInMso,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNexusDashboards NiatelemetryNexusDashboards

// NewNiatelemetryNexusDashboards instantiates a new NiatelemetryNexusDashboards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNexusDashboards(classId string, objectType string) *NiatelemetryNexusDashboards {
	this := NiatelemetryNexusDashboards{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNexusDashboardsWithDefaults instantiates a new NiatelemetryNexusDashboards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNexusDashboardsWithDefaults() *NiatelemetryNexusDashboards {
	this := NiatelemetryNexusDashboards{}
	var classId string = "niatelemetry.NexusDashboards"
	this.ClassId = classId
	var objectType string = "niatelemetry.NexusDashboards"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNexusDashboards) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNexusDashboards) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNexusDashboards) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNexusDashboards) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *NiatelemetryNexusDashboards) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *NiatelemetryNexusDashboards) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryNexusDashboards) SetDn(v string) {
	o.Dn = &v
}

// GetIsClusterHealthy returns the IsClusterHealthy field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetIsClusterHealthy() string {
	if o == nil || o.IsClusterHealthy == nil {
		var ret string
		return ret
	}
	return *o.IsClusterHealthy
}

// GetIsClusterHealthyOk returns a tuple with the IsClusterHealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetIsClusterHealthyOk() (*string, bool) {
	if o == nil || o.IsClusterHealthy == nil {
		return nil, false
	}
	return o.IsClusterHealthy, true
}

// HasIsClusterHealthy returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasIsClusterHealthy() bool {
	if o != nil && o.IsClusterHealthy != nil {
		return true
	}

	return false
}

// SetIsClusterHealthy gets a reference to the given string and assigns it to the IsClusterHealthy field.
func (o *NiatelemetryNexusDashboards) SetIsClusterHealthy(v string) {
	o.IsClusterHealthy = &v
}

// GetNdClusterSize returns the NdClusterSize field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNdClusterSize() int64 {
	if o == nil || o.NdClusterSize == nil {
		var ret int64
		return ret
	}
	return *o.NdClusterSize
}

// GetNdClusterSizeOk returns a tuple with the NdClusterSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNdClusterSizeOk() (*int64, bool) {
	if o == nil || o.NdClusterSize == nil {
		return nil, false
	}
	return o.NdClusterSize, true
}

// HasNdClusterSize returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNdClusterSize() bool {
	if o != nil && o.NdClusterSize != nil {
		return true
	}

	return false
}

// SetNdClusterSize gets a reference to the given int64 and assigns it to the NdClusterSize field.
func (o *NiatelemetryNexusDashboards) SetNdClusterSize(v int64) {
	o.NdClusterSize = &v
}

// GetNdSites returns the NdSites field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryNexusDashboards) GetNdSites() []NiatelemetrySites {
	if o == nil {
		var ret []NiatelemetrySites
		return ret
	}
	return o.NdSites
}

// GetNdSitesOk returns a tuple with the NdSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryNexusDashboards) GetNdSitesOk() ([]NiatelemetrySites, bool) {
	if o == nil || o.NdSites == nil {
		return nil, false
	}
	return o.NdSites, true
}

// HasNdSites returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNdSites() bool {
	if o != nil && o.NdSites != nil {
		return true
	}

	return false
}

// SetNdSites gets a reference to the given []NiatelemetrySites and assigns it to the NdSites field.
func (o *NiatelemetryNexusDashboards) SetNdSites(v []NiatelemetrySites) {
	o.NdSites = v
}

// GetNdType returns the NdType field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNdType() string {
	if o == nil || o.NdType == nil {
		var ret string
		return ret
	}
	return *o.NdType
}

// GetNdTypeOk returns a tuple with the NdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNdTypeOk() (*string, bool) {
	if o == nil || o.NdType == nil {
		return nil, false
	}
	return o.NdType, true
}

// HasNdType returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNdType() bool {
	if o != nil && o.NdType != nil {
		return true
	}

	return false
}

// SetNdType gets a reference to the given string and assigns it to the NdType field.
func (o *NiatelemetryNexusDashboards) SetNdType(v string) {
	o.NdType = &v
}

// GetNdVersion returns the NdVersion field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNdVersion() string {
	if o == nil || o.NdVersion == nil {
		var ret string
		return ret
	}
	return *o.NdVersion
}

// GetNdVersionOk returns a tuple with the NdVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNdVersionOk() (*string, bool) {
	if o == nil || o.NdVersion == nil {
		return nil, false
	}
	return o.NdVersion, true
}

// HasNdVersion returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNdVersion() bool {
	if o != nil && o.NdVersion != nil {
		return true
	}

	return false
}

// SetNdVersion gets a reference to the given string and assigns it to the NdVersion field.
func (o *NiatelemetryNexusDashboards) SetNdVersion(v string) {
	o.NdVersion = &v
}

// GetNumberOfApps returns the NumberOfApps field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNumberOfApps() int64 {
	if o == nil || o.NumberOfApps == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfApps
}

// GetNumberOfAppsOk returns a tuple with the NumberOfApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNumberOfAppsOk() (*int64, bool) {
	if o == nil || o.NumberOfApps == nil {
		return nil, false
	}
	return o.NumberOfApps, true
}

// HasNumberOfApps returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNumberOfApps() bool {
	if o != nil && o.NumberOfApps != nil {
		return true
	}

	return false
}

// SetNumberOfApps gets a reference to the given int64 and assigns it to the NumberOfApps field.
func (o *NiatelemetryNexusDashboards) SetNumberOfApps(v int64) {
	o.NumberOfApps = &v
}

// GetNumberOfInsightGroups returns the NumberOfInsightGroups field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNumberOfInsightGroups() int64 {
	if o == nil || o.NumberOfInsightGroups == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfInsightGroups
}

// GetNumberOfInsightGroupsOk returns a tuple with the NumberOfInsightGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNumberOfInsightGroupsOk() (*int64, bool) {
	if o == nil || o.NumberOfInsightGroups == nil {
		return nil, false
	}
	return o.NumberOfInsightGroups, true
}

// HasNumberOfInsightGroups returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNumberOfInsightGroups() bool {
	if o != nil && o.NumberOfInsightGroups != nil {
		return true
	}

	return false
}

// SetNumberOfInsightGroups gets a reference to the given int64 and assigns it to the NumberOfInsightGroups field.
func (o *NiatelemetryNexusDashboards) SetNumberOfInsightGroups(v int64) {
	o.NumberOfInsightGroups = &v
}

// GetNumberOfNirDashboards returns the NumberOfNirDashboards field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNumberOfNirDashboards() int64 {
	if o == nil || o.NumberOfNirDashboards == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfNirDashboards
}

// GetNumberOfNirDashboardsOk returns a tuple with the NumberOfNirDashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNumberOfNirDashboardsOk() (*int64, bool) {
	if o == nil || o.NumberOfNirDashboards == nil {
		return nil, false
	}
	return o.NumberOfNirDashboards, true
}

// HasNumberOfNirDashboards returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNumberOfNirDashboards() bool {
	if o != nil && o.NumberOfNirDashboards != nil {
		return true
	}

	return false
}

// SetNumberOfNirDashboards gets a reference to the given int64 and assigns it to the NumberOfNirDashboards field.
func (o *NiatelemetryNexusDashboards) SetNumberOfNirDashboards(v int64) {
	o.NumberOfNirDashboards = &v
}

// GetNumberOfSchemasInMso returns the NumberOfSchemasInMso field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNumberOfSchemasInMso() int64 {
	if o == nil || o.NumberOfSchemasInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSchemasInMso
}

// GetNumberOfSchemasInMsoOk returns a tuple with the NumberOfSchemasInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNumberOfSchemasInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfSchemasInMso == nil {
		return nil, false
	}
	return o.NumberOfSchemasInMso, true
}

// HasNumberOfSchemasInMso returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNumberOfSchemasInMso() bool {
	if o != nil && o.NumberOfSchemasInMso != nil {
		return true
	}

	return false
}

// SetNumberOfSchemasInMso gets a reference to the given int64 and assigns it to the NumberOfSchemasInMso field.
func (o *NiatelemetryNexusDashboards) SetNumberOfSchemasInMso(v int64) {
	o.NumberOfSchemasInMso = &v
}

// GetNumberOfSitesInMso returns the NumberOfSitesInMso field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNumberOfSitesInMso() int64 {
	if o == nil || o.NumberOfSitesInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSitesInMso
}

// GetNumberOfSitesInMsoOk returns a tuple with the NumberOfSitesInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNumberOfSitesInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfSitesInMso == nil {
		return nil, false
	}
	return o.NumberOfSitesInMso, true
}

// HasNumberOfSitesInMso returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNumberOfSitesInMso() bool {
	if o != nil && o.NumberOfSitesInMso != nil {
		return true
	}

	return false
}

// SetNumberOfSitesInMso gets a reference to the given int64 and assigns it to the NumberOfSitesInMso field.
func (o *NiatelemetryNexusDashboards) SetNumberOfSitesInMso(v int64) {
	o.NumberOfSitesInMso = &v
}

// GetNumberOfSitesServiced returns the NumberOfSitesServiced field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNumberOfSitesServiced() int64 {
	if o == nil || o.NumberOfSitesServiced == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSitesServiced
}

// GetNumberOfSitesServicedOk returns a tuple with the NumberOfSitesServiced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNumberOfSitesServicedOk() (*int64, bool) {
	if o == nil || o.NumberOfSitesServiced == nil {
		return nil, false
	}
	return o.NumberOfSitesServiced, true
}

// HasNumberOfSitesServiced returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNumberOfSitesServiced() bool {
	if o != nil && o.NumberOfSitesServiced != nil {
		return true
	}

	return false
}

// SetNumberOfSitesServiced gets a reference to the given int64 and assigns it to the NumberOfSitesServiced field.
func (o *NiatelemetryNexusDashboards) SetNumberOfSitesServiced(v int64) {
	o.NumberOfSitesServiced = &v
}

// GetNumberOfTenantsInMso returns the NumberOfTenantsInMso field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNumberOfTenantsInMso() int64 {
	if o == nil || o.NumberOfTenantsInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfTenantsInMso
}

// GetNumberOfTenantsInMsoOk returns a tuple with the NumberOfTenantsInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNumberOfTenantsInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfTenantsInMso == nil {
		return nil, false
	}
	return o.NumberOfTenantsInMso, true
}

// HasNumberOfTenantsInMso returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNumberOfTenantsInMso() bool {
	if o != nil && o.NumberOfTenantsInMso != nil {
		return true
	}

	return false
}

// SetNumberOfTenantsInMso gets a reference to the given int64 and assigns it to the NumberOfTenantsInMso field.
func (o *NiatelemetryNexusDashboards) SetNumberOfTenantsInMso(v int64) {
	o.NumberOfTenantsInMso = &v
}

// GetNumberOfVxlanFabricSitesInMso returns the NumberOfVxlanFabricSitesInMso field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetNumberOfVxlanFabricSitesInMso() int64 {
	if o == nil || o.NumberOfVxlanFabricSitesInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfVxlanFabricSitesInMso
}

// GetNumberOfVxlanFabricSitesInMsoOk returns a tuple with the NumberOfVxlanFabricSitesInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetNumberOfVxlanFabricSitesInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfVxlanFabricSitesInMso == nil {
		return nil, false
	}
	return o.NumberOfVxlanFabricSitesInMso, true
}

// HasNumberOfVxlanFabricSitesInMso returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasNumberOfVxlanFabricSitesInMso() bool {
	if o != nil && o.NumberOfVxlanFabricSitesInMso != nil {
		return true
	}

	return false
}

// SetNumberOfVxlanFabricSitesInMso gets a reference to the given int64 and assigns it to the NumberOfVxlanFabricSitesInMso field.
func (o *NiatelemetryNexusDashboards) SetNumberOfVxlanFabricSitesInMso(v int64) {
	o.NumberOfVxlanFabricSitesInMso = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryNexusDashboards) SetRecordType(v string) {
	o.RecordType = &v
}

// GetTypeOfSiteInMso returns the TypeOfSiteInMso field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetTypeOfSiteInMso() string {
	if o == nil || o.TypeOfSiteInMso == nil {
		var ret string
		return ret
	}
	return *o.TypeOfSiteInMso
}

// GetTypeOfSiteInMsoOk returns a tuple with the TypeOfSiteInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetTypeOfSiteInMsoOk() (*string, bool) {
	if o == nil || o.TypeOfSiteInMso == nil {
		return nil, false
	}
	return o.TypeOfSiteInMso, true
}

// HasTypeOfSiteInMso returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasTypeOfSiteInMso() bool {
	if o != nil && o.TypeOfSiteInMso != nil {
		return true
	}

	return false
}

// SetTypeOfSiteInMso gets a reference to the given string and assigns it to the TypeOfSiteInMso field.
func (o *NiatelemetryNexusDashboards) SetTypeOfSiteInMso(v string) {
	o.TypeOfSiteInMso = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboards) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboards) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboards) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNexusDashboards) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNexusDashboards) MarshalJSON() ([]byte, error) {
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
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.IsClusterHealthy != nil {
		toSerialize["IsClusterHealthy"] = o.IsClusterHealthy
	}
	if o.NdClusterSize != nil {
		toSerialize["NdClusterSize"] = o.NdClusterSize
	}
	if o.NdSites != nil {
		toSerialize["NdSites"] = o.NdSites
	}
	if o.NdType != nil {
		toSerialize["NdType"] = o.NdType
	}
	if o.NdVersion != nil {
		toSerialize["NdVersion"] = o.NdVersion
	}
	if o.NumberOfApps != nil {
		toSerialize["NumberOfApps"] = o.NumberOfApps
	}
	if o.NumberOfInsightGroups != nil {
		toSerialize["NumberOfInsightGroups"] = o.NumberOfInsightGroups
	}
	if o.NumberOfNirDashboards != nil {
		toSerialize["NumberOfNirDashboards"] = o.NumberOfNirDashboards
	}
	if o.NumberOfSchemasInMso != nil {
		toSerialize["NumberOfSchemasInMso"] = o.NumberOfSchemasInMso
	}
	if o.NumberOfSitesInMso != nil {
		toSerialize["NumberOfSitesInMso"] = o.NumberOfSitesInMso
	}
	if o.NumberOfSitesServiced != nil {
		toSerialize["NumberOfSitesServiced"] = o.NumberOfSitesServiced
	}
	if o.NumberOfTenantsInMso != nil {
		toSerialize["NumberOfTenantsInMso"] = o.NumberOfTenantsInMso
	}
	if o.NumberOfVxlanFabricSitesInMso != nil {
		toSerialize["NumberOfVxlanFabricSitesInMso"] = o.NumberOfVxlanFabricSitesInMso
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.TypeOfSiteInMso != nil {
		toSerialize["TypeOfSiteInMso"] = o.TypeOfSiteInMso
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNexusDashboards) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryNexusDashboardsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Nexus Dashboard can onboard multiple APIC clusters/sites.
		ClusterName *string `json:"ClusterName,omitempty"`
		// UUID of the Nexus Dashboard cluster.
		ClusterUuid *string `json:"ClusterUuid,omitempty"`
		// Dn of the objects present for Nexus Dashboard devices.
		Dn *string `json:"Dn,omitempty"`
		// Health of Nexus Dashboard cluster.
		IsClusterHealthy *string `json:"IsClusterHealthy,omitempty"`
		// Number of nodes in Nexus Dashboard cluster.
		NdClusterSize *int64              `json:"NdClusterSize,omitempty"`
		NdSites       []NiatelemetrySites `json:"NdSites,omitempty"`
		// Node type in Nexus Dashboard cluster.
		NdType *string `json:"NdType,omitempty"`
		// Version running on Nexus Dashboard.
		NdVersion *string `json:"NdVersion,omitempty"`
		// Number of applications installed in the Nexus Dashboard.
		NumberOfApps *int64 `json:"NumberOfApps,omitempty"`
		// Number of total insight groups in ND.
		NumberOfInsightGroups *int64 `json:"NumberOfInsightGroups,omitempty"`
		// Number of total NIR dashboards in ND.
		NumberOfNirDashboards *int64 `json:"NumberOfNirDashboards,omitempty"`
		// Number of total schemas in Multi-Site Orchestrator.
		NumberOfSchemasInMso *int64 `json:"NumberOfSchemasInMso,omitempty"`
		// Number of sites in Multi-Site Orchestrator.
		NumberOfSitesInMso *int64 `json:"NumberOfSitesInMso,omitempty"`
		// Number of sites serviced by ND.
		NumberOfSitesServiced *int64 `json:"NumberOfSitesServiced,omitempty"`
		// Number of total tenants in Multi-Site Orchestrator.
		NumberOfTenantsInMso *int64 `json:"NumberOfTenantsInMso,omitempty"`
		// Number of sites with vxLan type fabric in Multi-Site Orchestrator.
		NumberOfVxlanFabricSitesInMso *int64 `json:"NumberOfVxlanFabricSitesInMso,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Type of site added to Multi-Site Orchestrator.
		TypeOfSiteInMso  *string                              `json:"TypeOfSiteInMso,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryNexusDashboardsWithoutEmbeddedStruct := NiatelemetryNexusDashboardsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboardsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNexusDashboards := _NiatelemetryNexusDashboards{}
		varNiatelemetryNexusDashboards.ClassId = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.ClassId
		varNiatelemetryNexusDashboards.ObjectType = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNexusDashboards.ClusterName = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.ClusterName
		varNiatelemetryNexusDashboards.ClusterUuid = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.ClusterUuid
		varNiatelemetryNexusDashboards.Dn = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.Dn
		varNiatelemetryNexusDashboards.IsClusterHealthy = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.IsClusterHealthy
		varNiatelemetryNexusDashboards.NdClusterSize = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NdClusterSize
		varNiatelemetryNexusDashboards.NdSites = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NdSites
		varNiatelemetryNexusDashboards.NdType = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NdType
		varNiatelemetryNexusDashboards.NdVersion = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NdVersion
		varNiatelemetryNexusDashboards.NumberOfApps = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NumberOfApps
		varNiatelemetryNexusDashboards.NumberOfInsightGroups = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NumberOfInsightGroups
		varNiatelemetryNexusDashboards.NumberOfNirDashboards = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NumberOfNirDashboards
		varNiatelemetryNexusDashboards.NumberOfSchemasInMso = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NumberOfSchemasInMso
		varNiatelemetryNexusDashboards.NumberOfSitesInMso = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NumberOfSitesInMso
		varNiatelemetryNexusDashboards.NumberOfSitesServiced = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NumberOfSitesServiced
		varNiatelemetryNexusDashboards.NumberOfTenantsInMso = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NumberOfTenantsInMso
		varNiatelemetryNexusDashboards.NumberOfVxlanFabricSitesInMso = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.NumberOfVxlanFabricSitesInMso
		varNiatelemetryNexusDashboards.RecordType = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.RecordType
		varNiatelemetryNexusDashboards.TypeOfSiteInMso = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.TypeOfSiteInMso
		varNiatelemetryNexusDashboards.RegisteredDevice = varNiatelemetryNexusDashboardsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryNexusDashboards(varNiatelemetryNexusDashboards)
	} else {
		return err
	}

	varNiatelemetryNexusDashboards := _NiatelemetryNexusDashboards{}

	err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboards)
	if err == nil {
		o.MoBaseMo = varNiatelemetryNexusDashboards.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "IsClusterHealthy")
		delete(additionalProperties, "NdClusterSize")
		delete(additionalProperties, "NdSites")
		delete(additionalProperties, "NdType")
		delete(additionalProperties, "NdVersion")
		delete(additionalProperties, "NumberOfApps")
		delete(additionalProperties, "NumberOfInsightGroups")
		delete(additionalProperties, "NumberOfNirDashboards")
		delete(additionalProperties, "NumberOfSchemasInMso")
		delete(additionalProperties, "NumberOfSitesInMso")
		delete(additionalProperties, "NumberOfSitesServiced")
		delete(additionalProperties, "NumberOfTenantsInMso")
		delete(additionalProperties, "NumberOfVxlanFabricSitesInMso")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "TypeOfSiteInMso")
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

type NullableNiatelemetryNexusDashboards struct {
	value *NiatelemetryNexusDashboards
	isSet bool
}

func (v NullableNiatelemetryNexusDashboards) Get() *NiatelemetryNexusDashboards {
	return v.value
}

func (v *NullableNiatelemetryNexusDashboards) Set(val *NiatelemetryNexusDashboards) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusDashboards) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusDashboards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusDashboards(val *NiatelemetryNexusDashboards) *NullableNiatelemetryNexusDashboards {
	return &NullableNiatelemetryNexusDashboards{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusDashboards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusDashboards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
