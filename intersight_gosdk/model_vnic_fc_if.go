/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18012
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the VnicFcIf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicFcIf{}

// VnicFcIf Virtual Fibre Channel Interface.
type VnicFcIf struct {
	VnicBaseFcIf
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the virtual fibre channel interface.
	Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-\\\\._:]+$"`
	// The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The order should start from zero with no overlaps. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1340, VIC 1380 and VIC 1385 which have two.
	Order          *int64                        `json:"Order,omitempty"`
	OverriddenList []string                      `json:"OverriddenList,omitempty"`
	Placement      NullableVnicPlacementSettings `json:"Placement,omitempty"`
	// The WWPN address must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx.
	StaticWwpnAddress  *string                 `json:"StaticWwpnAddress,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
	TemplateActions    []MotemplateActionEntry `json:"TemplateActions,omitempty"`
	TemplateSyncErrors []MotemplateSyncError   `json:"TemplateSyncErrors,omitempty"`
	// The sync status of the current MO wrt the attached Template MO. * `None` - The Enum value represents that the object is not attached to any template. * `OK` - The Enum value represents that the object values are in sync with attached template. * `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template. * `InProgress` - The Enum value represents that the object sync with the attached template is in progress. * `OutOfSync` - The Enum value represents that the object values are not in sync with attached template.
	TemplateSyncStatus *string `json:"TemplateSyncStatus,omitempty"`
	// This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA.
	VifId *int64 `json:"VifId,omitempty"`
	// The WWPN address that is assigned to the vHBA based on the wwn pool that has been assigned to the SAN Connectivity Policy.
	Wwpn *string `json:"Wwpn,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
	// Type of allocation selected to assign a WWPN address to the vhba. * `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.
	WwpnAddressType       *string                                         `json:"WwpnAddressType,omitempty"`
	Profile               NullablePolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
	SanConnectivityPolicy NullableVnicSanConnectivityPolicyRelationship   `json:"SanConnectivityPolicy,omitempty"`
	ScpVhba               NullableVnicFcIfRelationship                    `json:"ScpVhba,omitempty"`
	// An array of relationships to vnicFcIf resources.
	SpVhbas              []VnicFcIfRelationship               `json:"SpVhbas,omitempty"`
	SrcTemplate          NullableVnicVhbaTemplateRelationship `json:"SrcTemplate,omitempty"`
	WwpnLease            NullableFcpoolLeaseRelationship      `json:"WwpnLease,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcIf VnicFcIf

// NewVnicFcIf instantiates a new VnicFcIf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcIf(classId string, objectType string) *VnicFcIf {
	this := VnicFcIf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "fc-initiator"
	this.Type = &type_
	var wwpnAddressType string = "POOL"
	this.WwpnAddressType = &wwpnAddressType
	return &this
}

// NewVnicFcIfWithDefaults instantiates a new VnicFcIf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcIfWithDefaults() *VnicFcIf {
	this := VnicFcIf{}
	var classId string = "vnic.FcIf"
	this.ClassId = classId
	var objectType string = "vnic.FcIf"
	this.ObjectType = objectType
	var wwpnAddressType string = "POOL"
	this.WwpnAddressType = &wwpnAddressType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcIf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcIf) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.FcIf" of the ClassId field.
func (o *VnicFcIf) GetDefaultClassId() interface{} {
	return "vnic.FcIf"
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcIf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcIf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.FcIf" of the ObjectType field.
func (o *VnicFcIf) GetDefaultObjectType() interface{} {
	return "vnic.FcIf"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicFcIf) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicFcIf) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicFcIf) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VnicFcIf) GetOrder() int64 {
	if o == nil || IsNil(o.Order) {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VnicFcIf) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *VnicFcIf) SetOrder(v int64) {
	o.Order = &v
}

// GetOverriddenList returns the OverriddenList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetOverriddenList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OverriddenList
}

// GetOverriddenListOk returns a tuple with the OverriddenList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetOverriddenListOk() ([]string, bool) {
	if o == nil || IsNil(o.OverriddenList) {
		return nil, false
	}
	return o.OverriddenList, true
}

// HasOverriddenList returns a boolean if a field has been set.
func (o *VnicFcIf) HasOverriddenList() bool {
	if o != nil && !IsNil(o.OverriddenList) {
		return true
	}

	return false
}

// SetOverriddenList gets a reference to the given []string and assigns it to the OverriddenList field.
func (o *VnicFcIf) SetOverriddenList(v []string) {
	o.OverriddenList = v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetPlacement() VnicPlacementSettings {
	if o == nil || IsNil(o.Placement.Get()) {
		var ret VnicPlacementSettings
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetPlacementOk() (*VnicPlacementSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *VnicFcIf) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableVnicPlacementSettings and assigns it to the Placement field.
func (o *VnicFcIf) SetPlacement(v VnicPlacementSettings) {
	o.Placement.Set(&v)
}

// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *VnicFcIf) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *VnicFcIf) UnsetPlacement() {
	o.Placement.Unset()
}

// GetStaticWwpnAddress returns the StaticWwpnAddress field value if set, zero value otherwise.
func (o *VnicFcIf) GetStaticWwpnAddress() string {
	if o == nil || IsNil(o.StaticWwpnAddress) {
		var ret string
		return ret
	}
	return *o.StaticWwpnAddress
}

// GetStaticWwpnAddressOk returns a tuple with the StaticWwpnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetStaticWwpnAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StaticWwpnAddress) {
		return nil, false
	}
	return o.StaticWwpnAddress, true
}

// HasStaticWwpnAddress returns a boolean if a field has been set.
func (o *VnicFcIf) HasStaticWwpnAddress() bool {
	if o != nil && !IsNil(o.StaticWwpnAddress) {
		return true
	}

	return false
}

// SetStaticWwpnAddress gets a reference to the given string and assigns it to the StaticWwpnAddress field.
func (o *VnicFcIf) SetStaticWwpnAddress(v string) {
	o.StaticWwpnAddress = &v
}

// GetTemplateActions returns the TemplateActions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetTemplateActions() []MotemplateActionEntry {
	if o == nil {
		var ret []MotemplateActionEntry
		return ret
	}
	return o.TemplateActions
}

// GetTemplateActionsOk returns a tuple with the TemplateActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetTemplateActionsOk() ([]MotemplateActionEntry, bool) {
	if o == nil || IsNil(o.TemplateActions) {
		return nil, false
	}
	return o.TemplateActions, true
}

// HasTemplateActions returns a boolean if a field has been set.
func (o *VnicFcIf) HasTemplateActions() bool {
	if o != nil && !IsNil(o.TemplateActions) {
		return true
	}

	return false
}

// SetTemplateActions gets a reference to the given []MotemplateActionEntry and assigns it to the TemplateActions field.
func (o *VnicFcIf) SetTemplateActions(v []MotemplateActionEntry) {
	o.TemplateActions = v
}

// GetTemplateSyncErrors returns the TemplateSyncErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetTemplateSyncErrors() []MotemplateSyncError {
	if o == nil {
		var ret []MotemplateSyncError
		return ret
	}
	return o.TemplateSyncErrors
}

// GetTemplateSyncErrorsOk returns a tuple with the TemplateSyncErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetTemplateSyncErrorsOk() ([]MotemplateSyncError, bool) {
	if o == nil || IsNil(o.TemplateSyncErrors) {
		return nil, false
	}
	return o.TemplateSyncErrors, true
}

// HasTemplateSyncErrors returns a boolean if a field has been set.
func (o *VnicFcIf) HasTemplateSyncErrors() bool {
	if o != nil && !IsNil(o.TemplateSyncErrors) {
		return true
	}

	return false
}

// SetTemplateSyncErrors gets a reference to the given []MotemplateSyncError and assigns it to the TemplateSyncErrors field.
func (o *VnicFcIf) SetTemplateSyncErrors(v []MotemplateSyncError) {
	o.TemplateSyncErrors = v
}

// GetTemplateSyncStatus returns the TemplateSyncStatus field value if set, zero value otherwise.
func (o *VnicFcIf) GetTemplateSyncStatus() string {
	if o == nil || IsNil(o.TemplateSyncStatus) {
		var ret string
		return ret
	}
	return *o.TemplateSyncStatus
}

// GetTemplateSyncStatusOk returns a tuple with the TemplateSyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetTemplateSyncStatusOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateSyncStatus) {
		return nil, false
	}
	return o.TemplateSyncStatus, true
}

// HasTemplateSyncStatus returns a boolean if a field has been set.
func (o *VnicFcIf) HasTemplateSyncStatus() bool {
	if o != nil && !IsNil(o.TemplateSyncStatus) {
		return true
	}

	return false
}

// SetTemplateSyncStatus gets a reference to the given string and assigns it to the TemplateSyncStatus field.
func (o *VnicFcIf) SetTemplateSyncStatus(v string) {
	o.TemplateSyncStatus = &v
}

// GetVifId returns the VifId field value if set, zero value otherwise.
func (o *VnicFcIf) GetVifId() int64 {
	if o == nil || IsNil(o.VifId) {
		var ret int64
		return ret
	}
	return *o.VifId
}

// GetVifIdOk returns a tuple with the VifId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetVifIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VifId) {
		return nil, false
	}
	return o.VifId, true
}

// HasVifId returns a boolean if a field has been set.
func (o *VnicFcIf) HasVifId() bool {
	if o != nil && !IsNil(o.VifId) {
		return true
	}

	return false
}

// SetVifId gets a reference to the given int64 and assigns it to the VifId field.
func (o *VnicFcIf) SetVifId(v int64) {
	o.VifId = &v
}

// GetWwpn returns the Wwpn field value if set, zero value otherwise.
func (o *VnicFcIf) GetWwpn() string {
	if o == nil || IsNil(o.Wwpn) {
		var ret string
		return ret
	}
	return *o.Wwpn
}

// GetWwpnOk returns a tuple with the Wwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetWwpnOk() (*string, bool) {
	if o == nil || IsNil(o.Wwpn) {
		return nil, false
	}
	return o.Wwpn, true
}

// HasWwpn returns a boolean if a field has been set.
func (o *VnicFcIf) HasWwpn() bool {
	if o != nil && !IsNil(o.Wwpn) {
		return true
	}

	return false
}

// SetWwpn gets a reference to the given string and assigns it to the Wwpn field.
func (o *VnicFcIf) SetWwpn(v string) {
	o.Wwpn = &v
}

// GetWwpnAddressType returns the WwpnAddressType field value if set, zero value otherwise.
func (o *VnicFcIf) GetWwpnAddressType() string {
	if o == nil || IsNil(o.WwpnAddressType) {
		var ret string
		return ret
	}
	return *o.WwpnAddressType
}

// GetWwpnAddressTypeOk returns a tuple with the WwpnAddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcIf) GetWwpnAddressTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WwpnAddressType) {
		return nil, false
	}
	return o.WwpnAddressType, true
}

// HasWwpnAddressType returns a boolean if a field has been set.
func (o *VnicFcIf) HasWwpnAddressType() bool {
	if o != nil && !IsNil(o.WwpnAddressType) {
		return true
	}

	return false
}

// SetWwpnAddressType gets a reference to the given string and assigns it to the WwpnAddressType field.
func (o *VnicFcIf) SetWwpnAddressType(v string) {
	o.WwpnAddressType = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetProfile() PolicyAbstractConfigProfileRelationship {
	if o == nil || IsNil(o.Profile.Get()) {
		var ret PolicyAbstractConfigProfileRelationship
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// HasProfile returns a boolean if a field has been set.
func (o *VnicFcIf) HasProfile() bool {
	if o != nil && o.Profile.IsSet() {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NullablePolicyAbstractConfigProfileRelationship and assigns it to the Profile field.
func (o *VnicFcIf) SetProfile(v PolicyAbstractConfigProfileRelationship) {
	o.Profile.Set(&v)
}

// SetProfileNil sets the value for Profile to be an explicit nil
func (o *VnicFcIf) SetProfileNil() {
	o.Profile.Set(nil)
}

// UnsetProfile ensures that no value is present for Profile, not even an explicit nil
func (o *VnicFcIf) UnsetProfile() {
	o.Profile.Unset()
}

// GetSanConnectivityPolicy returns the SanConnectivityPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetSanConnectivityPolicy() VnicSanConnectivityPolicyRelationship {
	if o == nil || IsNil(o.SanConnectivityPolicy.Get()) {
		var ret VnicSanConnectivityPolicyRelationship
		return ret
	}
	return *o.SanConnectivityPolicy.Get()
}

// GetSanConnectivityPolicyOk returns a tuple with the SanConnectivityPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetSanConnectivityPolicyOk() (*VnicSanConnectivityPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SanConnectivityPolicy.Get(), o.SanConnectivityPolicy.IsSet()
}

// HasSanConnectivityPolicy returns a boolean if a field has been set.
func (o *VnicFcIf) HasSanConnectivityPolicy() bool {
	if o != nil && o.SanConnectivityPolicy.IsSet() {
		return true
	}

	return false
}

// SetSanConnectivityPolicy gets a reference to the given NullableVnicSanConnectivityPolicyRelationship and assigns it to the SanConnectivityPolicy field.
func (o *VnicFcIf) SetSanConnectivityPolicy(v VnicSanConnectivityPolicyRelationship) {
	o.SanConnectivityPolicy.Set(&v)
}

// SetSanConnectivityPolicyNil sets the value for SanConnectivityPolicy to be an explicit nil
func (o *VnicFcIf) SetSanConnectivityPolicyNil() {
	o.SanConnectivityPolicy.Set(nil)
}

// UnsetSanConnectivityPolicy ensures that no value is present for SanConnectivityPolicy, not even an explicit nil
func (o *VnicFcIf) UnsetSanConnectivityPolicy() {
	o.SanConnectivityPolicy.Unset()
}

// GetScpVhba returns the ScpVhba field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetScpVhba() VnicFcIfRelationship {
	if o == nil || IsNil(o.ScpVhba.Get()) {
		var ret VnicFcIfRelationship
		return ret
	}
	return *o.ScpVhba.Get()
}

// GetScpVhbaOk returns a tuple with the ScpVhba field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetScpVhbaOk() (*VnicFcIfRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScpVhba.Get(), o.ScpVhba.IsSet()
}

// HasScpVhba returns a boolean if a field has been set.
func (o *VnicFcIf) HasScpVhba() bool {
	if o != nil && o.ScpVhba.IsSet() {
		return true
	}

	return false
}

// SetScpVhba gets a reference to the given NullableVnicFcIfRelationship and assigns it to the ScpVhba field.
func (o *VnicFcIf) SetScpVhba(v VnicFcIfRelationship) {
	o.ScpVhba.Set(&v)
}

// SetScpVhbaNil sets the value for ScpVhba to be an explicit nil
func (o *VnicFcIf) SetScpVhbaNil() {
	o.ScpVhba.Set(nil)
}

// UnsetScpVhba ensures that no value is present for ScpVhba, not even an explicit nil
func (o *VnicFcIf) UnsetScpVhba() {
	o.ScpVhba.Unset()
}

// GetSpVhbas returns the SpVhbas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetSpVhbas() []VnicFcIfRelationship {
	if o == nil {
		var ret []VnicFcIfRelationship
		return ret
	}
	return o.SpVhbas
}

// GetSpVhbasOk returns a tuple with the SpVhbas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetSpVhbasOk() ([]VnicFcIfRelationship, bool) {
	if o == nil || IsNil(o.SpVhbas) {
		return nil, false
	}
	return o.SpVhbas, true
}

// HasSpVhbas returns a boolean if a field has been set.
func (o *VnicFcIf) HasSpVhbas() bool {
	if o != nil && !IsNil(o.SpVhbas) {
		return true
	}

	return false
}

// SetSpVhbas gets a reference to the given []VnicFcIfRelationship and assigns it to the SpVhbas field.
func (o *VnicFcIf) SetSpVhbas(v []VnicFcIfRelationship) {
	o.SpVhbas = v
}

// GetSrcTemplate returns the SrcTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetSrcTemplate() VnicVhbaTemplateRelationship {
	if o == nil || IsNil(o.SrcTemplate.Get()) {
		var ret VnicVhbaTemplateRelationship
		return ret
	}
	return *o.SrcTemplate.Get()
}

// GetSrcTemplateOk returns a tuple with the SrcTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetSrcTemplateOk() (*VnicVhbaTemplateRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SrcTemplate.Get(), o.SrcTemplate.IsSet()
}

// HasSrcTemplate returns a boolean if a field has been set.
func (o *VnicFcIf) HasSrcTemplate() bool {
	if o != nil && o.SrcTemplate.IsSet() {
		return true
	}

	return false
}

// SetSrcTemplate gets a reference to the given NullableVnicVhbaTemplateRelationship and assigns it to the SrcTemplate field.
func (o *VnicFcIf) SetSrcTemplate(v VnicVhbaTemplateRelationship) {
	o.SrcTemplate.Set(&v)
}

// SetSrcTemplateNil sets the value for SrcTemplate to be an explicit nil
func (o *VnicFcIf) SetSrcTemplateNil() {
	o.SrcTemplate.Set(nil)
}

// UnsetSrcTemplate ensures that no value is present for SrcTemplate, not even an explicit nil
func (o *VnicFcIf) UnsetSrcTemplate() {
	o.SrcTemplate.Unset()
}

// GetWwpnLease returns the WwpnLease field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcIf) GetWwpnLease() FcpoolLeaseRelationship {
	if o == nil || IsNil(o.WwpnLease.Get()) {
		var ret FcpoolLeaseRelationship
		return ret
	}
	return *o.WwpnLease.Get()
}

// GetWwpnLeaseOk returns a tuple with the WwpnLease field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcIf) GetWwpnLeaseOk() (*FcpoolLeaseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.WwpnLease.Get(), o.WwpnLease.IsSet()
}

// HasWwpnLease returns a boolean if a field has been set.
func (o *VnicFcIf) HasWwpnLease() bool {
	if o != nil && o.WwpnLease.IsSet() {
		return true
	}

	return false
}

// SetWwpnLease gets a reference to the given NullableFcpoolLeaseRelationship and assigns it to the WwpnLease field.
func (o *VnicFcIf) SetWwpnLease(v FcpoolLeaseRelationship) {
	o.WwpnLease.Set(&v)
}

// SetWwpnLeaseNil sets the value for WwpnLease to be an explicit nil
func (o *VnicFcIf) SetWwpnLeaseNil() {
	o.WwpnLease.Set(nil)
}

// UnsetWwpnLease ensures that no value is present for WwpnLease, not even an explicit nil
func (o *VnicFcIf) UnsetWwpnLease() {
	o.WwpnLease.Unset()
}

func (o VnicFcIf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicFcIf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVnicBaseFcIf, errVnicBaseFcIf := json.Marshal(o.VnicBaseFcIf)
	if errVnicBaseFcIf != nil {
		return map[string]interface{}{}, errVnicBaseFcIf
	}
	errVnicBaseFcIf = json.Unmarshal([]byte(serializedVnicBaseFcIf), &toSerialize)
	if errVnicBaseFcIf != nil {
		return map[string]interface{}{}, errVnicBaseFcIf
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Order) {
		toSerialize["Order"] = o.Order
	}
	if o.OverriddenList != nil {
		toSerialize["OverriddenList"] = o.OverriddenList
	}
	if o.Placement.IsSet() {
		toSerialize["Placement"] = o.Placement.Get()
	}
	if !IsNil(o.StaticWwpnAddress) {
		toSerialize["StaticWwpnAddress"] = o.StaticWwpnAddress
	}
	if o.TemplateActions != nil {
		toSerialize["TemplateActions"] = o.TemplateActions
	}
	if o.TemplateSyncErrors != nil {
		toSerialize["TemplateSyncErrors"] = o.TemplateSyncErrors
	}
	if !IsNil(o.TemplateSyncStatus) {
		toSerialize["TemplateSyncStatus"] = o.TemplateSyncStatus
	}
	if !IsNil(o.VifId) {
		toSerialize["VifId"] = o.VifId
	}
	if !IsNil(o.Wwpn) {
		toSerialize["Wwpn"] = o.Wwpn
	}
	if !IsNil(o.WwpnAddressType) {
		toSerialize["WwpnAddressType"] = o.WwpnAddressType
	}
	if o.Profile.IsSet() {
		toSerialize["Profile"] = o.Profile.Get()
	}
	if o.SanConnectivityPolicy.IsSet() {
		toSerialize["SanConnectivityPolicy"] = o.SanConnectivityPolicy.Get()
	}
	if o.ScpVhba.IsSet() {
		toSerialize["ScpVhba"] = o.ScpVhba.Get()
	}
	if o.SpVhbas != nil {
		toSerialize["SpVhbas"] = o.SpVhbas
	}
	if o.SrcTemplate.IsSet() {
		toSerialize["SrcTemplate"] = o.SrcTemplate.Get()
	}
	if o.WwpnLease.IsSet() {
		toSerialize["WwpnLease"] = o.WwpnLease.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicFcIf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type VnicFcIfWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the virtual fibre channel interface.
		Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-\\\\._:]+$"`
		// The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The order should start from zero with no overlaps. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1340, VIC 1380 and VIC 1385 which have two.
		Order          *int64                        `json:"Order,omitempty"`
		OverriddenList []string                      `json:"OverriddenList,omitempty"`
		Placement      NullableVnicPlacementSettings `json:"Placement,omitempty"`
		// The WWPN address must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx.
		StaticWwpnAddress  *string                 `json:"StaticWwpnAddress,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
		TemplateActions    []MotemplateActionEntry `json:"TemplateActions,omitempty"`
		TemplateSyncErrors []MotemplateSyncError   `json:"TemplateSyncErrors,omitempty"`
		// The sync status of the current MO wrt the attached Template MO. * `None` - The Enum value represents that the object is not attached to any template. * `OK` - The Enum value represents that the object values are in sync with attached template. * `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template. * `InProgress` - The Enum value represents that the object sync with the attached template is in progress. * `OutOfSync` - The Enum value represents that the object values are not in sync with attached template.
		TemplateSyncStatus *string `json:"TemplateSyncStatus,omitempty"`
		// This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA.
		VifId *int64 `json:"VifId,omitempty"`
		// The WWPN address that is assigned to the vHBA based on the wwn pool that has been assigned to the SAN Connectivity Policy.
		Wwpn *string `json:"Wwpn,omitempty" validate:"regexp=^$|((^20|5[0-9a-fA-F]{1}):([0-9a-fA-F]{2}:){6}([0-9a-fA-F]{2})$)"`
		// Type of allocation selected to assign a WWPN address to the vhba. * `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.
		WwpnAddressType       *string                                         `json:"WwpnAddressType,omitempty"`
		Profile               NullablePolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
		SanConnectivityPolicy NullableVnicSanConnectivityPolicyRelationship   `json:"SanConnectivityPolicy,omitempty"`
		ScpVhba               NullableVnicFcIfRelationship                    `json:"ScpVhba,omitempty"`
		// An array of relationships to vnicFcIf resources.
		SpVhbas     []VnicFcIfRelationship               `json:"SpVhbas,omitempty"`
		SrcTemplate NullableVnicVhbaTemplateRelationship `json:"SrcTemplate,omitempty"`
		WwpnLease   NullableFcpoolLeaseRelationship      `json:"WwpnLease,omitempty"`
	}

	varVnicFcIfWithoutEmbeddedStruct := VnicFcIfWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicFcIfWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcIf := _VnicFcIf{}
		varVnicFcIf.ClassId = varVnicFcIfWithoutEmbeddedStruct.ClassId
		varVnicFcIf.ObjectType = varVnicFcIfWithoutEmbeddedStruct.ObjectType
		varVnicFcIf.Name = varVnicFcIfWithoutEmbeddedStruct.Name
		varVnicFcIf.Order = varVnicFcIfWithoutEmbeddedStruct.Order
		varVnicFcIf.OverriddenList = varVnicFcIfWithoutEmbeddedStruct.OverriddenList
		varVnicFcIf.Placement = varVnicFcIfWithoutEmbeddedStruct.Placement
		varVnicFcIf.StaticWwpnAddress = varVnicFcIfWithoutEmbeddedStruct.StaticWwpnAddress
		varVnicFcIf.TemplateActions = varVnicFcIfWithoutEmbeddedStruct.TemplateActions
		varVnicFcIf.TemplateSyncErrors = varVnicFcIfWithoutEmbeddedStruct.TemplateSyncErrors
		varVnicFcIf.TemplateSyncStatus = varVnicFcIfWithoutEmbeddedStruct.TemplateSyncStatus
		varVnicFcIf.VifId = varVnicFcIfWithoutEmbeddedStruct.VifId
		varVnicFcIf.Wwpn = varVnicFcIfWithoutEmbeddedStruct.Wwpn
		varVnicFcIf.WwpnAddressType = varVnicFcIfWithoutEmbeddedStruct.WwpnAddressType
		varVnicFcIf.Profile = varVnicFcIfWithoutEmbeddedStruct.Profile
		varVnicFcIf.SanConnectivityPolicy = varVnicFcIfWithoutEmbeddedStruct.SanConnectivityPolicy
		varVnicFcIf.ScpVhba = varVnicFcIfWithoutEmbeddedStruct.ScpVhba
		varVnicFcIf.SpVhbas = varVnicFcIfWithoutEmbeddedStruct.SpVhbas
		varVnicFcIf.SrcTemplate = varVnicFcIfWithoutEmbeddedStruct.SrcTemplate
		varVnicFcIf.WwpnLease = varVnicFcIfWithoutEmbeddedStruct.WwpnLease
		*o = VnicFcIf(varVnicFcIf)
	} else {
		return err
	}

	varVnicFcIf := _VnicFcIf{}

	err = json.Unmarshal(data, &varVnicFcIf)
	if err == nil {
		o.VnicBaseFcIf = varVnicFcIf.VnicBaseFcIf
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Order")
		delete(additionalProperties, "OverriddenList")
		delete(additionalProperties, "Placement")
		delete(additionalProperties, "StaticWwpnAddress")
		delete(additionalProperties, "TemplateActions")
		delete(additionalProperties, "TemplateSyncErrors")
		delete(additionalProperties, "TemplateSyncStatus")
		delete(additionalProperties, "VifId")
		delete(additionalProperties, "Wwpn")
		delete(additionalProperties, "WwpnAddressType")
		delete(additionalProperties, "Profile")
		delete(additionalProperties, "SanConnectivityPolicy")
		delete(additionalProperties, "ScpVhba")
		delete(additionalProperties, "SpVhbas")
		delete(additionalProperties, "SrcTemplate")
		delete(additionalProperties, "WwpnLease")

		// remove fields from embedded structs
		reflectVnicBaseFcIf := reflect.ValueOf(o.VnicBaseFcIf)
		for i := 0; i < reflectVnicBaseFcIf.Type().NumField(); i++ {
			t := reflectVnicBaseFcIf.Type().Field(i)

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

type NullableVnicFcIf struct {
	value *VnicFcIf
	isSet bool
}

func (v NullableVnicFcIf) Get() *VnicFcIf {
	return v.value
}

func (v *NullableVnicFcIf) Set(val *VnicFcIf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcIf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcIf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcIf(val *VnicFcIf) *NullableVnicFcIf {
	return &NullableVnicFcIf{value: val, isSet: true}
}

func (v NullableVnicFcIf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcIf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
