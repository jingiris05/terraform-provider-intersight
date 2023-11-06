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

// AssetDeploymentDevice Contains information about Cisco devices associated with consumption-based subscriptions. In addition to device installation status, information about returns and replacements is also recorded here. We listen to messages sent by Cisco Install Base and create/update an instance of this object.
type AssetDeploymentDevice struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                                 `json:"ObjectType"`
	AlarmInfo  NullableAssetDeploymentDeviceAlarmInfo `json:"AlarmInfo,omitempty"`
	// Unique identifier of the Cisco device.
	DeviceId          *string                                  `json:"DeviceId,omitempty"`
	DeviceInformation NullableAssetDeploymentDeviceInformation `json:"DeviceInformation,omitempty"`
	// Product identifier for the specified Cisco device. It is used to distinguish between HyperFlex and UCS devices.
	DevicePid        *string                       `json:"DevicePid,omitempty"`
	DeviceStatistics NullableAssetDeviceStatistics `json:"DeviceStatistics,omitempty"`
	// Product Subgroup type helps to determine if device subgroup within Product type has to be billed using consumption metering. example \"N9300 Series\" in Product type \"SWITCH\".
	ProductSubgroup *string `json:"ProductSubgroup,omitempty"`
	// Product type helps to determine if device has to be billed using consumption metering. example \"SERVER\".
	ProductType *string `json:"ProductType,omitempty"`
	// String reference to the device connector.
	RegisteredDeviceMoid *string             `json:"RegisteredDeviceMoid,omitempty"`
	UnitOfMeasure        []AssetMeteringType `json:"UnitOfMeasure,omitempty"`
	// Virtualization platform is used to identify the hypervisor type. example \"ESXi\".
	VirtualizationPlatform *string `json:"VirtualizationPlatform,omitempty"`
	// Workload/Usecase running on the device.
	Workload                  *string                                     `json:"Workload,omitempty"`
	Deployment                *AssetDeploymentRelationship                `json:"Deployment,omitempty"`
	DeviceContractInformation *AssetDeviceContractInformationRelationship `json:"DeviceContractInformation,omitempty"`
	RegisteredDevice          *AssetDeviceRegistrationRelationship        `json:"RegisteredDevice,omitempty"`
	Subscription              *AssetSubscriptionRelationship              `json:"Subscription,omitempty"`
	SubscriptionAccount       *AssetSubscriptionAccountRelationship       `json:"SubscriptionAccount,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _AssetDeploymentDevice AssetDeploymentDevice

// NewAssetDeploymentDevice instantiates a new AssetDeploymentDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeploymentDevice(classId string, objectType string) *AssetDeploymentDevice {
	this := AssetDeploymentDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeploymentDeviceWithDefaults instantiates a new AssetDeploymentDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeploymentDeviceWithDefaults() *AssetDeploymentDevice {
	this := AssetDeploymentDevice{}
	var classId string = "asset.DeploymentDevice"
	this.ClassId = classId
	var objectType string = "asset.DeploymentDevice"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeploymentDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeploymentDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeploymentDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeploymentDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlarmInfo returns the AlarmInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeploymentDevice) GetAlarmInfo() AssetDeploymentDeviceAlarmInfo {
	if o == nil || o.AlarmInfo.Get() == nil {
		var ret AssetDeploymentDeviceAlarmInfo
		return ret
	}
	return *o.AlarmInfo.Get()
}

// GetAlarmInfoOk returns a tuple with the AlarmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeploymentDevice) GetAlarmInfoOk() (*AssetDeploymentDeviceAlarmInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlarmInfo.Get(), o.AlarmInfo.IsSet()
}

// HasAlarmInfo returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasAlarmInfo() bool {
	if o != nil && o.AlarmInfo.IsSet() {
		return true
	}

	return false
}

// SetAlarmInfo gets a reference to the given NullableAssetDeploymentDeviceAlarmInfo and assigns it to the AlarmInfo field.
func (o *AssetDeploymentDevice) SetAlarmInfo(v AssetDeploymentDeviceAlarmInfo) {
	o.AlarmInfo.Set(&v)
}

// SetAlarmInfoNil sets the value for AlarmInfo to be an explicit nil
func (o *AssetDeploymentDevice) SetAlarmInfoNil() {
	o.AlarmInfo.Set(nil)
}

// UnsetAlarmInfo ensures that no value is present for AlarmInfo, not even an explicit nil
func (o *AssetDeploymentDevice) UnsetAlarmInfo() {
	o.AlarmInfo.Unset()
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *AssetDeploymentDevice) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceInformation returns the DeviceInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeploymentDevice) GetDeviceInformation() AssetDeploymentDeviceInformation {
	if o == nil || o.DeviceInformation.Get() == nil {
		var ret AssetDeploymentDeviceInformation
		return ret
	}
	return *o.DeviceInformation.Get()
}

// GetDeviceInformationOk returns a tuple with the DeviceInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeploymentDevice) GetDeviceInformationOk() (*AssetDeploymentDeviceInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceInformation.Get(), o.DeviceInformation.IsSet()
}

// HasDeviceInformation returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasDeviceInformation() bool {
	if o != nil && o.DeviceInformation.IsSet() {
		return true
	}

	return false
}

// SetDeviceInformation gets a reference to the given NullableAssetDeploymentDeviceInformation and assigns it to the DeviceInformation field.
func (o *AssetDeploymentDevice) SetDeviceInformation(v AssetDeploymentDeviceInformation) {
	o.DeviceInformation.Set(&v)
}

// SetDeviceInformationNil sets the value for DeviceInformation to be an explicit nil
func (o *AssetDeploymentDevice) SetDeviceInformationNil() {
	o.DeviceInformation.Set(nil)
}

// UnsetDeviceInformation ensures that no value is present for DeviceInformation, not even an explicit nil
func (o *AssetDeploymentDevice) UnsetDeviceInformation() {
	o.DeviceInformation.Unset()
}

// GetDevicePid returns the DevicePid field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetDevicePid() string {
	if o == nil || o.DevicePid == nil {
		var ret string
		return ret
	}
	return *o.DevicePid
}

// GetDevicePidOk returns a tuple with the DevicePid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetDevicePidOk() (*string, bool) {
	if o == nil || o.DevicePid == nil {
		return nil, false
	}
	return o.DevicePid, true
}

// HasDevicePid returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasDevicePid() bool {
	if o != nil && o.DevicePid != nil {
		return true
	}

	return false
}

// SetDevicePid gets a reference to the given string and assigns it to the DevicePid field.
func (o *AssetDeploymentDevice) SetDevicePid(v string) {
	o.DevicePid = &v
}

// GetDeviceStatistics returns the DeviceStatistics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeploymentDevice) GetDeviceStatistics() AssetDeviceStatistics {
	if o == nil || o.DeviceStatistics.Get() == nil {
		var ret AssetDeviceStatistics
		return ret
	}
	return *o.DeviceStatistics.Get()
}

// GetDeviceStatisticsOk returns a tuple with the DeviceStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeploymentDevice) GetDeviceStatisticsOk() (*AssetDeviceStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceStatistics.Get(), o.DeviceStatistics.IsSet()
}

// HasDeviceStatistics returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasDeviceStatistics() bool {
	if o != nil && o.DeviceStatistics.IsSet() {
		return true
	}

	return false
}

// SetDeviceStatistics gets a reference to the given NullableAssetDeviceStatistics and assigns it to the DeviceStatistics field.
func (o *AssetDeploymentDevice) SetDeviceStatistics(v AssetDeviceStatistics) {
	o.DeviceStatistics.Set(&v)
}

// SetDeviceStatisticsNil sets the value for DeviceStatistics to be an explicit nil
func (o *AssetDeploymentDevice) SetDeviceStatisticsNil() {
	o.DeviceStatistics.Set(nil)
}

// UnsetDeviceStatistics ensures that no value is present for DeviceStatistics, not even an explicit nil
func (o *AssetDeploymentDevice) UnsetDeviceStatistics() {
	o.DeviceStatistics.Unset()
}

// GetProductSubgroup returns the ProductSubgroup field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetProductSubgroup() string {
	if o == nil || o.ProductSubgroup == nil {
		var ret string
		return ret
	}
	return *o.ProductSubgroup
}

// GetProductSubgroupOk returns a tuple with the ProductSubgroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetProductSubgroupOk() (*string, bool) {
	if o == nil || o.ProductSubgroup == nil {
		return nil, false
	}
	return o.ProductSubgroup, true
}

// HasProductSubgroup returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasProductSubgroup() bool {
	if o != nil && o.ProductSubgroup != nil {
		return true
	}

	return false
}

// SetProductSubgroup gets a reference to the given string and assigns it to the ProductSubgroup field.
func (o *AssetDeploymentDevice) SetProductSubgroup(v string) {
	o.ProductSubgroup = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetProductType() string {
	if o == nil || o.ProductType == nil {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetProductTypeOk() (*string, bool) {
	if o == nil || o.ProductType == nil {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasProductType() bool {
	if o != nil && o.ProductType != nil {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *AssetDeploymentDevice) SetProductType(v string) {
	o.ProductType = &v
}

// GetRegisteredDeviceMoid returns the RegisteredDeviceMoid field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetRegisteredDeviceMoid() string {
	if o == nil || o.RegisteredDeviceMoid == nil {
		var ret string
		return ret
	}
	return *o.RegisteredDeviceMoid
}

// GetRegisteredDeviceMoidOk returns a tuple with the RegisteredDeviceMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetRegisteredDeviceMoidOk() (*string, bool) {
	if o == nil || o.RegisteredDeviceMoid == nil {
		return nil, false
	}
	return o.RegisteredDeviceMoid, true
}

// HasRegisteredDeviceMoid returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasRegisteredDeviceMoid() bool {
	if o != nil && o.RegisteredDeviceMoid != nil {
		return true
	}

	return false
}

// SetRegisteredDeviceMoid gets a reference to the given string and assigns it to the RegisteredDeviceMoid field.
func (o *AssetDeploymentDevice) SetRegisteredDeviceMoid(v string) {
	o.RegisteredDeviceMoid = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeploymentDevice) GetUnitOfMeasure() []AssetMeteringType {
	if o == nil {
		var ret []AssetMeteringType
		return ret
	}
	return o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeploymentDevice) GetUnitOfMeasureOk() ([]AssetMeteringType, bool) {
	if o == nil || o.UnitOfMeasure == nil {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasUnitOfMeasure() bool {
	if o != nil && o.UnitOfMeasure != nil {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given []AssetMeteringType and assigns it to the UnitOfMeasure field.
func (o *AssetDeploymentDevice) SetUnitOfMeasure(v []AssetMeteringType) {
	o.UnitOfMeasure = v
}

// GetVirtualizationPlatform returns the VirtualizationPlatform field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetVirtualizationPlatform() string {
	if o == nil || o.VirtualizationPlatform == nil {
		var ret string
		return ret
	}
	return *o.VirtualizationPlatform
}

// GetVirtualizationPlatformOk returns a tuple with the VirtualizationPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetVirtualizationPlatformOk() (*string, bool) {
	if o == nil || o.VirtualizationPlatform == nil {
		return nil, false
	}
	return o.VirtualizationPlatform, true
}

// HasVirtualizationPlatform returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasVirtualizationPlatform() bool {
	if o != nil && o.VirtualizationPlatform != nil {
		return true
	}

	return false
}

// SetVirtualizationPlatform gets a reference to the given string and assigns it to the VirtualizationPlatform field.
func (o *AssetDeploymentDevice) SetVirtualizationPlatform(v string) {
	o.VirtualizationPlatform = &v
}

// GetWorkload returns the Workload field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetWorkload() string {
	if o == nil || o.Workload == nil {
		var ret string
		return ret
	}
	return *o.Workload
}

// GetWorkloadOk returns a tuple with the Workload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetWorkloadOk() (*string, bool) {
	if o == nil || o.Workload == nil {
		return nil, false
	}
	return o.Workload, true
}

// HasWorkload returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasWorkload() bool {
	if o != nil && o.Workload != nil {
		return true
	}

	return false
}

// SetWorkload gets a reference to the given string and assigns it to the Workload field.
func (o *AssetDeploymentDevice) SetWorkload(v string) {
	o.Workload = &v
}

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetDeployment() AssetDeploymentRelationship {
	if o == nil || o.Deployment == nil {
		var ret AssetDeploymentRelationship
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetDeploymentOk() (*AssetDeploymentRelationship, bool) {
	if o == nil || o.Deployment == nil {
		return nil, false
	}
	return o.Deployment, true
}

// HasDeployment returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasDeployment() bool {
	if o != nil && o.Deployment != nil {
		return true
	}

	return false
}

// SetDeployment gets a reference to the given AssetDeploymentRelationship and assigns it to the Deployment field.
func (o *AssetDeploymentDevice) SetDeployment(v AssetDeploymentRelationship) {
	o.Deployment = &v
}

// GetDeviceContractInformation returns the DeviceContractInformation field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetDeviceContractInformation() AssetDeviceContractInformationRelationship {
	if o == nil || o.DeviceContractInformation == nil {
		var ret AssetDeviceContractInformationRelationship
		return ret
	}
	return *o.DeviceContractInformation
}

// GetDeviceContractInformationOk returns a tuple with the DeviceContractInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetDeviceContractInformationOk() (*AssetDeviceContractInformationRelationship, bool) {
	if o == nil || o.DeviceContractInformation == nil {
		return nil, false
	}
	return o.DeviceContractInformation, true
}

// HasDeviceContractInformation returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasDeviceContractInformation() bool {
	if o != nil && o.DeviceContractInformation != nil {
		return true
	}

	return false
}

// SetDeviceContractInformation gets a reference to the given AssetDeviceContractInformationRelationship and assigns it to the DeviceContractInformation field.
func (o *AssetDeploymentDevice) SetDeviceContractInformation(v AssetDeviceContractInformationRelationship) {
	o.DeviceContractInformation = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AssetDeploymentDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetSubscription() AssetSubscriptionRelationship {
	if o == nil || o.Subscription == nil {
		var ret AssetSubscriptionRelationship
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetSubscriptionOk() (*AssetSubscriptionRelationship, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given AssetSubscriptionRelationship and assigns it to the Subscription field.
func (o *AssetDeploymentDevice) SetSubscription(v AssetSubscriptionRelationship) {
	o.Subscription = &v
}

// GetSubscriptionAccount returns the SubscriptionAccount field value if set, zero value otherwise.
func (o *AssetDeploymentDevice) GetSubscriptionAccount() AssetSubscriptionAccountRelationship {
	if o == nil || o.SubscriptionAccount == nil {
		var ret AssetSubscriptionAccountRelationship
		return ret
	}
	return *o.SubscriptionAccount
}

// GetSubscriptionAccountOk returns a tuple with the SubscriptionAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeploymentDevice) GetSubscriptionAccountOk() (*AssetSubscriptionAccountRelationship, bool) {
	if o == nil || o.SubscriptionAccount == nil {
		return nil, false
	}
	return o.SubscriptionAccount, true
}

// HasSubscriptionAccount returns a boolean if a field has been set.
func (o *AssetDeploymentDevice) HasSubscriptionAccount() bool {
	if o != nil && o.SubscriptionAccount != nil {
		return true
	}

	return false
}

// SetSubscriptionAccount gets a reference to the given AssetSubscriptionAccountRelationship and assigns it to the SubscriptionAccount field.
func (o *AssetDeploymentDevice) SetSubscriptionAccount(v AssetSubscriptionAccountRelationship) {
	o.SubscriptionAccount = &v
}

func (o AssetDeploymentDevice) MarshalJSON() ([]byte, error) {
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
	if o.AlarmInfo.IsSet() {
		toSerialize["AlarmInfo"] = o.AlarmInfo.Get()
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.DeviceInformation.IsSet() {
		toSerialize["DeviceInformation"] = o.DeviceInformation.Get()
	}
	if o.DevicePid != nil {
		toSerialize["DevicePid"] = o.DevicePid
	}
	if o.DeviceStatistics.IsSet() {
		toSerialize["DeviceStatistics"] = o.DeviceStatistics.Get()
	}
	if o.ProductSubgroup != nil {
		toSerialize["ProductSubgroup"] = o.ProductSubgroup
	}
	if o.ProductType != nil {
		toSerialize["ProductType"] = o.ProductType
	}
	if o.RegisteredDeviceMoid != nil {
		toSerialize["RegisteredDeviceMoid"] = o.RegisteredDeviceMoid
	}
	if o.UnitOfMeasure != nil {
		toSerialize["UnitOfMeasure"] = o.UnitOfMeasure
	}
	if o.VirtualizationPlatform != nil {
		toSerialize["VirtualizationPlatform"] = o.VirtualizationPlatform
	}
	if o.Workload != nil {
		toSerialize["Workload"] = o.Workload
	}
	if o.Deployment != nil {
		toSerialize["Deployment"] = o.Deployment
	}
	if o.DeviceContractInformation != nil {
		toSerialize["DeviceContractInformation"] = o.DeviceContractInformation
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Subscription != nil {
		toSerialize["Subscription"] = o.Subscription
	}
	if o.SubscriptionAccount != nil {
		toSerialize["SubscriptionAccount"] = o.SubscriptionAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetDeploymentDevice) UnmarshalJSON(bytes []byte) (err error) {
	type AssetDeploymentDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                 `json:"ObjectType"`
		AlarmInfo  NullableAssetDeploymentDeviceAlarmInfo `json:"AlarmInfo,omitempty"`
		// Unique identifier of the Cisco device.
		DeviceId          *string                                  `json:"DeviceId,omitempty"`
		DeviceInformation NullableAssetDeploymentDeviceInformation `json:"DeviceInformation,omitempty"`
		// Product identifier for the specified Cisco device. It is used to distinguish between HyperFlex and UCS devices.
		DevicePid        *string                       `json:"DevicePid,omitempty"`
		DeviceStatistics NullableAssetDeviceStatistics `json:"DeviceStatistics,omitempty"`
		// Product Subgroup type helps to determine if device subgroup within Product type has to be billed using consumption metering. example \"N9300 Series\" in Product type \"SWITCH\".
		ProductSubgroup *string `json:"ProductSubgroup,omitempty"`
		// Product type helps to determine if device has to be billed using consumption metering. example \"SERVER\".
		ProductType *string `json:"ProductType,omitempty"`
		// String reference to the device connector.
		RegisteredDeviceMoid *string             `json:"RegisteredDeviceMoid,omitempty"`
		UnitOfMeasure        []AssetMeteringType `json:"UnitOfMeasure,omitempty"`
		// Virtualization platform is used to identify the hypervisor type. example \"ESXi\".
		VirtualizationPlatform *string `json:"VirtualizationPlatform,omitempty"`
		// Workload/Usecase running on the device.
		Workload                  *string                                     `json:"Workload,omitempty"`
		Deployment                *AssetDeploymentRelationship                `json:"Deployment,omitempty"`
		DeviceContractInformation *AssetDeviceContractInformationRelationship `json:"DeviceContractInformation,omitempty"`
		RegisteredDevice          *AssetDeviceRegistrationRelationship        `json:"RegisteredDevice,omitempty"`
		Subscription              *AssetSubscriptionRelationship              `json:"Subscription,omitempty"`
		SubscriptionAccount       *AssetSubscriptionAccountRelationship       `json:"SubscriptionAccount,omitempty"`
	}

	varAssetDeploymentDeviceWithoutEmbeddedStruct := AssetDeploymentDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetDeploymentDeviceWithoutEmbeddedStruct)
	if err == nil {
		varAssetDeploymentDevice := _AssetDeploymentDevice{}
		varAssetDeploymentDevice.ClassId = varAssetDeploymentDeviceWithoutEmbeddedStruct.ClassId
		varAssetDeploymentDevice.ObjectType = varAssetDeploymentDeviceWithoutEmbeddedStruct.ObjectType
		varAssetDeploymentDevice.AlarmInfo = varAssetDeploymentDeviceWithoutEmbeddedStruct.AlarmInfo
		varAssetDeploymentDevice.DeviceId = varAssetDeploymentDeviceWithoutEmbeddedStruct.DeviceId
		varAssetDeploymentDevice.DeviceInformation = varAssetDeploymentDeviceWithoutEmbeddedStruct.DeviceInformation
		varAssetDeploymentDevice.DevicePid = varAssetDeploymentDeviceWithoutEmbeddedStruct.DevicePid
		varAssetDeploymentDevice.DeviceStatistics = varAssetDeploymentDeviceWithoutEmbeddedStruct.DeviceStatistics
		varAssetDeploymentDevice.ProductSubgroup = varAssetDeploymentDeviceWithoutEmbeddedStruct.ProductSubgroup
		varAssetDeploymentDevice.ProductType = varAssetDeploymentDeviceWithoutEmbeddedStruct.ProductType
		varAssetDeploymentDevice.RegisteredDeviceMoid = varAssetDeploymentDeviceWithoutEmbeddedStruct.RegisteredDeviceMoid
		varAssetDeploymentDevice.UnitOfMeasure = varAssetDeploymentDeviceWithoutEmbeddedStruct.UnitOfMeasure
		varAssetDeploymentDevice.VirtualizationPlatform = varAssetDeploymentDeviceWithoutEmbeddedStruct.VirtualizationPlatform
		varAssetDeploymentDevice.Workload = varAssetDeploymentDeviceWithoutEmbeddedStruct.Workload
		varAssetDeploymentDevice.Deployment = varAssetDeploymentDeviceWithoutEmbeddedStruct.Deployment
		varAssetDeploymentDevice.DeviceContractInformation = varAssetDeploymentDeviceWithoutEmbeddedStruct.DeviceContractInformation
		varAssetDeploymentDevice.RegisteredDevice = varAssetDeploymentDeviceWithoutEmbeddedStruct.RegisteredDevice
		varAssetDeploymentDevice.Subscription = varAssetDeploymentDeviceWithoutEmbeddedStruct.Subscription
		varAssetDeploymentDevice.SubscriptionAccount = varAssetDeploymentDeviceWithoutEmbeddedStruct.SubscriptionAccount
		*o = AssetDeploymentDevice(varAssetDeploymentDevice)
	} else {
		return err
	}

	varAssetDeploymentDevice := _AssetDeploymentDevice{}

	err = json.Unmarshal(bytes, &varAssetDeploymentDevice)
	if err == nil {
		o.MoBaseMo = varAssetDeploymentDevice.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlarmInfo")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "DeviceInformation")
		delete(additionalProperties, "DevicePid")
		delete(additionalProperties, "DeviceStatistics")
		delete(additionalProperties, "ProductSubgroup")
		delete(additionalProperties, "ProductType")
		delete(additionalProperties, "RegisteredDeviceMoid")
		delete(additionalProperties, "UnitOfMeasure")
		delete(additionalProperties, "VirtualizationPlatform")
		delete(additionalProperties, "Workload")
		delete(additionalProperties, "Deployment")
		delete(additionalProperties, "DeviceContractInformation")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Subscription")
		delete(additionalProperties, "SubscriptionAccount")

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

type NullableAssetDeploymentDevice struct {
	value *AssetDeploymentDevice
	isSet bool
}

func (v NullableAssetDeploymentDevice) Get() *AssetDeploymentDevice {
	return v.value
}

func (v *NullableAssetDeploymentDevice) Set(val *AssetDeploymentDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeploymentDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeploymentDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeploymentDevice(val *AssetDeploymentDevice) *NullableAssetDeploymentDevice {
	return &NullableAssetDeploymentDevice{value: val, isSet: true}
}

func (v NullableAssetDeploymentDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeploymentDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
