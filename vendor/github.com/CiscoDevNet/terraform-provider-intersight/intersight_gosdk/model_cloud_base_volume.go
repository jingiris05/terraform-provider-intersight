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
	"reflect"
	"strings"
	"time"
)

// CloudBaseVolume A base Volume definition extended by cloud specific Volumes objects.
type CloudBaseVolume struct {
	VirtualizationBaseVirtualDisk
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType  string                   `json:"ObjectType"`
	BillingUnit NullableCloudBillingUnit `json:"BillingUnit,omitempty"`
	// Encryption state of this volume.For example ENCRYPTED, NOT ENCRYPTED, etc. * `Automatic` - Volume encryption state is automatic.Cloud provider takes the decision of when to encrypt the data. * `Encrypted` - Volume data is encrypted. Can be decrypted only by authorized users. * `Not_Encrypted` - Volume data is not encrypted.
	EncryptionState *string `json:"EncryptionState,omitempty"`
	// The internally generated identity of this VM. It aids in uniquely identifying the volume object.
	Identity            *string                         `json:"Identity,omitempty"`
	InstanceAttachments []CloudVolumeInstanceAttachment `json:"InstanceAttachments,omitempty"`
	IopsInfo            NullableCloudVolumeIopsInfo     `json:"IopsInfo,omitempty"`
	RegionInfo          NullableCloudCloudRegion        `json:"RegionInfo,omitempty"`
	// Source Image Id used for the volume.
	SourceImageId *string `json:"SourceImageId,omitempty"`
	// The volume state.For example AVAILABLE, IN_USE, DELETED, etc. * `UnRecognized` - Volume is in unrecognized state. * `Pending` - Volume is  in pending state, due to errors encountered during volume creation. * `Bound` - Volume is in bound state. * `Available` - Volume is in available state. * `Error` - Volume is in error state. * `Released` - Volume is in released state. * `in-use` - Volume is in in-use state. * `Creating` - Volume is in creating state. * `Deleting` - Volume is in deleting state. * `Deleted` - Volume is in deleted state.
	State *string `json:"State,omitempty"`
	// UUID (Universally Unique IDentifier) is a 128-bit value used for unique identification of Volume.
	Uuid *string `json:"Uuid,omitempty"`
	// Time when this volume is created.
	VolumeCreateTime     *time.Time                    `json:"VolumeCreateTime,omitempty"`
	VolumeTags           []CloudCloudTag               `json:"VolumeTags,omitempty"`
	VolumeType           NullableCloudVolumeType       `json:"VolumeType,omitempty"`
	ZoneInfo             NullableCloudAvailabilityZone `json:"ZoneInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudBaseVolume CloudBaseVolume

// NewCloudBaseVolume instantiates a new CloudBaseVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBaseVolume(classId string, objectType string) *CloudBaseVolume {
	this := CloudBaseVolume{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudBaseVolumeWithDefaults instantiates a new CloudBaseVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBaseVolumeWithDefaults() *CloudBaseVolume {
	this := CloudBaseVolume{}
	var classId string = "cloud.AwsVolume"
	this.ClassId = classId
	var objectType string = "cloud.AwsVolume"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudBaseVolume) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudBaseVolume) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudBaseVolume) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudBaseVolume) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudBaseVolume) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudBaseVolume) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBillingUnit returns the BillingUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVolume) GetBillingUnit() CloudBillingUnit {
	if o == nil || o.BillingUnit.Get() == nil {
		var ret CloudBillingUnit
		return ret
	}
	return *o.BillingUnit.Get()
}

// GetBillingUnitOk returns a tuple with the BillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVolume) GetBillingUnitOk() (*CloudBillingUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingUnit.Get(), o.BillingUnit.IsSet()
}

// HasBillingUnit returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasBillingUnit() bool {
	if o != nil && o.BillingUnit.IsSet() {
		return true
	}

	return false
}

// SetBillingUnit gets a reference to the given NullableCloudBillingUnit and assigns it to the BillingUnit field.
func (o *CloudBaseVolume) SetBillingUnit(v CloudBillingUnit) {
	o.BillingUnit.Set(&v)
}

// SetBillingUnitNil sets the value for BillingUnit to be an explicit nil
func (o *CloudBaseVolume) SetBillingUnitNil() {
	o.BillingUnit.Set(nil)
}

// UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
func (o *CloudBaseVolume) UnsetBillingUnit() {
	o.BillingUnit.Unset()
}

// GetEncryptionState returns the EncryptionState field value if set, zero value otherwise.
func (o *CloudBaseVolume) GetEncryptionState() string {
	if o == nil || o.EncryptionState == nil {
		var ret string
		return ret
	}
	return *o.EncryptionState
}

// GetEncryptionStateOk returns a tuple with the EncryptionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseVolume) GetEncryptionStateOk() (*string, bool) {
	if o == nil || o.EncryptionState == nil {
		return nil, false
	}
	return o.EncryptionState, true
}

// HasEncryptionState returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasEncryptionState() bool {
	if o != nil && o.EncryptionState != nil {
		return true
	}

	return false
}

// SetEncryptionState gets a reference to the given string and assigns it to the EncryptionState field.
func (o *CloudBaseVolume) SetEncryptionState(v string) {
	o.EncryptionState = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudBaseVolume) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseVolume) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudBaseVolume) SetIdentity(v string) {
	o.Identity = &v
}

// GetInstanceAttachments returns the InstanceAttachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVolume) GetInstanceAttachments() []CloudVolumeInstanceAttachment {
	if o == nil {
		var ret []CloudVolumeInstanceAttachment
		return ret
	}
	return o.InstanceAttachments
}

// GetInstanceAttachmentsOk returns a tuple with the InstanceAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVolume) GetInstanceAttachmentsOk() (*[]CloudVolumeInstanceAttachment, bool) {
	if o == nil || o.InstanceAttachments == nil {
		return nil, false
	}
	return &o.InstanceAttachments, true
}

// HasInstanceAttachments returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasInstanceAttachments() bool {
	if o != nil && o.InstanceAttachments != nil {
		return true
	}

	return false
}

// SetInstanceAttachments gets a reference to the given []CloudVolumeInstanceAttachment and assigns it to the InstanceAttachments field.
func (o *CloudBaseVolume) SetInstanceAttachments(v []CloudVolumeInstanceAttachment) {
	o.InstanceAttachments = v
}

// GetIopsInfo returns the IopsInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVolume) GetIopsInfo() CloudVolumeIopsInfo {
	if o == nil || o.IopsInfo.Get() == nil {
		var ret CloudVolumeIopsInfo
		return ret
	}
	return *o.IopsInfo.Get()
}

// GetIopsInfoOk returns a tuple with the IopsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVolume) GetIopsInfoOk() (*CloudVolumeIopsInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.IopsInfo.Get(), o.IopsInfo.IsSet()
}

// HasIopsInfo returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasIopsInfo() bool {
	if o != nil && o.IopsInfo.IsSet() {
		return true
	}

	return false
}

// SetIopsInfo gets a reference to the given NullableCloudVolumeIopsInfo and assigns it to the IopsInfo field.
func (o *CloudBaseVolume) SetIopsInfo(v CloudVolumeIopsInfo) {
	o.IopsInfo.Set(&v)
}

// SetIopsInfoNil sets the value for IopsInfo to be an explicit nil
func (o *CloudBaseVolume) SetIopsInfoNil() {
	o.IopsInfo.Set(nil)
}

// UnsetIopsInfo ensures that no value is present for IopsInfo, not even an explicit nil
func (o *CloudBaseVolume) UnsetIopsInfo() {
	o.IopsInfo.Unset()
}

// GetRegionInfo returns the RegionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVolume) GetRegionInfo() CloudCloudRegion {
	if o == nil || o.RegionInfo.Get() == nil {
		var ret CloudCloudRegion
		return ret
	}
	return *o.RegionInfo.Get()
}

// GetRegionInfoOk returns a tuple with the RegionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVolume) GetRegionInfoOk() (*CloudCloudRegion, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegionInfo.Get(), o.RegionInfo.IsSet()
}

// HasRegionInfo returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasRegionInfo() bool {
	if o != nil && o.RegionInfo.IsSet() {
		return true
	}

	return false
}

// SetRegionInfo gets a reference to the given NullableCloudCloudRegion and assigns it to the RegionInfo field.
func (o *CloudBaseVolume) SetRegionInfo(v CloudCloudRegion) {
	o.RegionInfo.Set(&v)
}

// SetRegionInfoNil sets the value for RegionInfo to be an explicit nil
func (o *CloudBaseVolume) SetRegionInfoNil() {
	o.RegionInfo.Set(nil)
}

// UnsetRegionInfo ensures that no value is present for RegionInfo, not even an explicit nil
func (o *CloudBaseVolume) UnsetRegionInfo() {
	o.RegionInfo.Unset()
}

// GetSourceImageId returns the SourceImageId field value if set, zero value otherwise.
func (o *CloudBaseVolume) GetSourceImageId() string {
	if o == nil || o.SourceImageId == nil {
		var ret string
		return ret
	}
	return *o.SourceImageId
}

// GetSourceImageIdOk returns a tuple with the SourceImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseVolume) GetSourceImageIdOk() (*string, bool) {
	if o == nil || o.SourceImageId == nil {
		return nil, false
	}
	return o.SourceImageId, true
}

// HasSourceImageId returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasSourceImageId() bool {
	if o != nil && o.SourceImageId != nil {
		return true
	}

	return false
}

// SetSourceImageId gets a reference to the given string and assigns it to the SourceImageId field.
func (o *CloudBaseVolume) SetSourceImageId(v string) {
	o.SourceImageId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CloudBaseVolume) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseVolume) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CloudBaseVolume) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CloudBaseVolume) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseVolume) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CloudBaseVolume) SetUuid(v string) {
	o.Uuid = &v
}

// GetVolumeCreateTime returns the VolumeCreateTime field value if set, zero value otherwise.
func (o *CloudBaseVolume) GetVolumeCreateTime() time.Time {
	if o == nil || o.VolumeCreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.VolumeCreateTime
}

// GetVolumeCreateTimeOk returns a tuple with the VolumeCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseVolume) GetVolumeCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.VolumeCreateTime == nil {
		return nil, false
	}
	return o.VolumeCreateTime, true
}

// HasVolumeCreateTime returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasVolumeCreateTime() bool {
	if o != nil && o.VolumeCreateTime != nil {
		return true
	}

	return false
}

// SetVolumeCreateTime gets a reference to the given time.Time and assigns it to the VolumeCreateTime field.
func (o *CloudBaseVolume) SetVolumeCreateTime(v time.Time) {
	o.VolumeCreateTime = &v
}

// GetVolumeTags returns the VolumeTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVolume) GetVolumeTags() []CloudCloudTag {
	if o == nil {
		var ret []CloudCloudTag
		return ret
	}
	return o.VolumeTags
}

// GetVolumeTagsOk returns a tuple with the VolumeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVolume) GetVolumeTagsOk() (*[]CloudCloudTag, bool) {
	if o == nil || o.VolumeTags == nil {
		return nil, false
	}
	return &o.VolumeTags, true
}

// HasVolumeTags returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasVolumeTags() bool {
	if o != nil && o.VolumeTags != nil {
		return true
	}

	return false
}

// SetVolumeTags gets a reference to the given []CloudCloudTag and assigns it to the VolumeTags field.
func (o *CloudBaseVolume) SetVolumeTags(v []CloudCloudTag) {
	o.VolumeTags = v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVolume) GetVolumeType() CloudVolumeType {
	if o == nil || o.VolumeType.Get() == nil {
		var ret CloudVolumeType
		return ret
	}
	return *o.VolumeType.Get()
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVolume) GetVolumeTypeOk() (*CloudVolumeType, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeType.Get(), o.VolumeType.IsSet()
}

// HasVolumeType returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasVolumeType() bool {
	if o != nil && o.VolumeType.IsSet() {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given NullableCloudVolumeType and assigns it to the VolumeType field.
func (o *CloudBaseVolume) SetVolumeType(v CloudVolumeType) {
	o.VolumeType.Set(&v)
}

// SetVolumeTypeNil sets the value for VolumeType to be an explicit nil
func (o *CloudBaseVolume) SetVolumeTypeNil() {
	o.VolumeType.Set(nil)
}

// UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
func (o *CloudBaseVolume) UnsetVolumeType() {
	o.VolumeType.Unset()
}

// GetZoneInfo returns the ZoneInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVolume) GetZoneInfo() CloudAvailabilityZone {
	if o == nil || o.ZoneInfo.Get() == nil {
		var ret CloudAvailabilityZone
		return ret
	}
	return *o.ZoneInfo.Get()
}

// GetZoneInfoOk returns a tuple with the ZoneInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVolume) GetZoneInfoOk() (*CloudAvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZoneInfo.Get(), o.ZoneInfo.IsSet()
}

// HasZoneInfo returns a boolean if a field has been set.
func (o *CloudBaseVolume) HasZoneInfo() bool {
	if o != nil && o.ZoneInfo.IsSet() {
		return true
	}

	return false
}

// SetZoneInfo gets a reference to the given NullableCloudAvailabilityZone and assigns it to the ZoneInfo field.
func (o *CloudBaseVolume) SetZoneInfo(v CloudAvailabilityZone) {
	o.ZoneInfo.Set(&v)
}

// SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil
func (o *CloudBaseVolume) SetZoneInfoNil() {
	o.ZoneInfo.Set(nil)
}

// UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil
func (o *CloudBaseVolume) UnsetZoneInfo() {
	o.ZoneInfo.Unset()
}

func (o CloudBaseVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualDisk, errVirtualizationBaseVirtualDisk := json.Marshal(o.VirtualizationBaseVirtualDisk)
	if errVirtualizationBaseVirtualDisk != nil {
		return []byte{}, errVirtualizationBaseVirtualDisk
	}
	errVirtualizationBaseVirtualDisk = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualDisk), &toSerialize)
	if errVirtualizationBaseVirtualDisk != nil {
		return []byte{}, errVirtualizationBaseVirtualDisk
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BillingUnit.IsSet() {
		toSerialize["BillingUnit"] = o.BillingUnit.Get()
	}
	if o.EncryptionState != nil {
		toSerialize["EncryptionState"] = o.EncryptionState
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.InstanceAttachments != nil {
		toSerialize["InstanceAttachments"] = o.InstanceAttachments
	}
	if o.IopsInfo.IsSet() {
		toSerialize["IopsInfo"] = o.IopsInfo.Get()
	}
	if o.RegionInfo.IsSet() {
		toSerialize["RegionInfo"] = o.RegionInfo.Get()
	}
	if o.SourceImageId != nil {
		toSerialize["SourceImageId"] = o.SourceImageId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.VolumeCreateTime != nil {
		toSerialize["VolumeCreateTime"] = o.VolumeCreateTime
	}
	if o.VolumeTags != nil {
		toSerialize["VolumeTags"] = o.VolumeTags
	}
	if o.VolumeType.IsSet() {
		toSerialize["VolumeType"] = o.VolumeType.Get()
	}
	if o.ZoneInfo.IsSet() {
		toSerialize["ZoneInfo"] = o.ZoneInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudBaseVolume) UnmarshalJSON(bytes []byte) (err error) {
	type CloudBaseVolumeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType  string                   `json:"ObjectType"`
		BillingUnit NullableCloudBillingUnit `json:"BillingUnit,omitempty"`
		// Encryption state of this volume.For example ENCRYPTED, NOT ENCRYPTED, etc. * `Automatic` - Volume encryption state is automatic.Cloud provider takes the decision of when to encrypt the data. * `Encrypted` - Volume data is encrypted. Can be decrypted only by authorized users. * `Not_Encrypted` - Volume data is not encrypted.
		EncryptionState *string `json:"EncryptionState,omitempty"`
		// The internally generated identity of this VM. It aids in uniquely identifying the volume object.
		Identity            *string                         `json:"Identity,omitempty"`
		InstanceAttachments []CloudVolumeInstanceAttachment `json:"InstanceAttachments,omitempty"`
		IopsInfo            NullableCloudVolumeIopsInfo     `json:"IopsInfo,omitempty"`
		RegionInfo          NullableCloudCloudRegion        `json:"RegionInfo,omitempty"`
		// Source Image Id used for the volume.
		SourceImageId *string `json:"SourceImageId,omitempty"`
		// The volume state.For example AVAILABLE, IN_USE, DELETED, etc. * `UnRecognized` - Volume is in unrecognized state. * `Pending` - Volume is  in pending state, due to errors encountered during volume creation. * `Bound` - Volume is in bound state. * `Available` - Volume is in available state. * `Error` - Volume is in error state. * `Released` - Volume is in released state. * `in-use` - Volume is in in-use state. * `Creating` - Volume is in creating state. * `Deleting` - Volume is in deleting state. * `Deleted` - Volume is in deleted state.
		State *string `json:"State,omitempty"`
		// UUID (Universally Unique IDentifier) is a 128-bit value used for unique identification of Volume.
		Uuid *string `json:"Uuid,omitempty"`
		// Time when this volume is created.
		VolumeCreateTime *time.Time                    `json:"VolumeCreateTime,omitempty"`
		VolumeTags       []CloudCloudTag               `json:"VolumeTags,omitempty"`
		VolumeType       NullableCloudVolumeType       `json:"VolumeType,omitempty"`
		ZoneInfo         NullableCloudAvailabilityZone `json:"ZoneInfo,omitempty"`
	}

	varCloudBaseVolumeWithoutEmbeddedStruct := CloudBaseVolumeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudBaseVolumeWithoutEmbeddedStruct)
	if err == nil {
		varCloudBaseVolume := _CloudBaseVolume{}
		varCloudBaseVolume.ClassId = varCloudBaseVolumeWithoutEmbeddedStruct.ClassId
		varCloudBaseVolume.ObjectType = varCloudBaseVolumeWithoutEmbeddedStruct.ObjectType
		varCloudBaseVolume.BillingUnit = varCloudBaseVolumeWithoutEmbeddedStruct.BillingUnit
		varCloudBaseVolume.EncryptionState = varCloudBaseVolumeWithoutEmbeddedStruct.EncryptionState
		varCloudBaseVolume.Identity = varCloudBaseVolumeWithoutEmbeddedStruct.Identity
		varCloudBaseVolume.InstanceAttachments = varCloudBaseVolumeWithoutEmbeddedStruct.InstanceAttachments
		varCloudBaseVolume.IopsInfo = varCloudBaseVolumeWithoutEmbeddedStruct.IopsInfo
		varCloudBaseVolume.RegionInfo = varCloudBaseVolumeWithoutEmbeddedStruct.RegionInfo
		varCloudBaseVolume.SourceImageId = varCloudBaseVolumeWithoutEmbeddedStruct.SourceImageId
		varCloudBaseVolume.State = varCloudBaseVolumeWithoutEmbeddedStruct.State
		varCloudBaseVolume.Uuid = varCloudBaseVolumeWithoutEmbeddedStruct.Uuid
		varCloudBaseVolume.VolumeCreateTime = varCloudBaseVolumeWithoutEmbeddedStruct.VolumeCreateTime
		varCloudBaseVolume.VolumeTags = varCloudBaseVolumeWithoutEmbeddedStruct.VolumeTags
		varCloudBaseVolume.VolumeType = varCloudBaseVolumeWithoutEmbeddedStruct.VolumeType
		varCloudBaseVolume.ZoneInfo = varCloudBaseVolumeWithoutEmbeddedStruct.ZoneInfo
		*o = CloudBaseVolume(varCloudBaseVolume)
	} else {
		return err
	}

	varCloudBaseVolume := _CloudBaseVolume{}

	err = json.Unmarshal(bytes, &varCloudBaseVolume)
	if err == nil {
		o.VirtualizationBaseVirtualDisk = varCloudBaseVolume.VirtualizationBaseVirtualDisk
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillingUnit")
		delete(additionalProperties, "EncryptionState")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "InstanceAttachments")
		delete(additionalProperties, "IopsInfo")
		delete(additionalProperties, "RegionInfo")
		delete(additionalProperties, "SourceImageId")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VolumeCreateTime")
		delete(additionalProperties, "VolumeTags")
		delete(additionalProperties, "VolumeType")
		delete(additionalProperties, "ZoneInfo")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualDisk := reflect.ValueOf(o.VirtualizationBaseVirtualDisk)
		for i := 0; i < reflectVirtualizationBaseVirtualDisk.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualDisk.Type().Field(i)

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

type NullableCloudBaseVolume struct {
	value *CloudBaseVolume
	isSet bool
}

func (v NullableCloudBaseVolume) Get() *CloudBaseVolume {
	return v.value
}

func (v *NullableCloudBaseVolume) Set(val *CloudBaseVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudBaseVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudBaseVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudBaseVolume(val *CloudBaseVolume) *NullableCloudBaseVolume {
	return &NullableCloudBaseVolume{value: val, isSet: true}
}

func (v NullableCloudBaseVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudBaseVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
