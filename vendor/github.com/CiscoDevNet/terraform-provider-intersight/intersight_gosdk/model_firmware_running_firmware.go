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

// checks if the FirmwareRunningFirmware type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareRunningFirmware{}

// FirmwareRunningFirmware Running Firmware on an endpoint.
type FirmwareRunningFirmware struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Kind of the firmware - boot-booloader/system/kernel.
	Component *string `json:"Component,omitempty"`
	// Bundle version which the firmware belongs to.
	PackageVersion *string `json:"PackageVersion,omitempty"`
	// The type of the firmware.
	Type *string `json:"Type,omitempty"`
	// The version of the firmware.
	Version              *string                                  `json:"Version,omitempty"`
	BiosUnit             NullableBiosUnitRelationship             `json:"BiosUnit,omitempty"`
	GraphicsCard         NullableGraphicsCardRelationship         `json:"GraphicsCard,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship  `json:"InventoryDeviceInfo,omitempty"`
	ManagementController NullableManagementControllerRelationship `json:"ManagementController,omitempty"`
	// An array of relationships to networkElement resources.
	NetworkElements            []NetworkElementRelationship                   `json:"NetworkElements,omitempty"`
	PciSwitch                  NullablePciSwitchRelationship                  `json:"PciSwitch,omitempty"`
	RegisteredDevice           NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	StorageController          NullableStorageControllerRelationship          `json:"StorageController,omitempty"`
	StorageFlexFlashController NullableStorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	StoragePhysicalDisk        NullableStoragePhysicalDiskRelationship        `json:"StoragePhysicalDisk,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _FirmwareRunningFirmware FirmwareRunningFirmware

// NewFirmwareRunningFirmware instantiates a new FirmwareRunningFirmware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareRunningFirmware(classId string, objectType string) *FirmwareRunningFirmware {
	this := FirmwareRunningFirmware{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareRunningFirmwareWithDefaults instantiates a new FirmwareRunningFirmware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareRunningFirmwareWithDefaults() *FirmwareRunningFirmware {
	this := FirmwareRunningFirmware{}
	var classId string = "firmware.RunningFirmware"
	this.ClassId = classId
	var objectType string = "firmware.RunningFirmware"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareRunningFirmware) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareRunningFirmware) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "firmware.RunningFirmware" of the ClassId field.
func (o *FirmwareRunningFirmware) GetDefaultClassId() interface{} {
	return "firmware.RunningFirmware"
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareRunningFirmware) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareRunningFirmware) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "firmware.RunningFirmware" of the ObjectType field.
func (o *FirmwareRunningFirmware) GetDefaultObjectType() interface{} {
	return "firmware.RunningFirmware"
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *FirmwareRunningFirmware) SetComponent(v string) {
	o.Component = &v
}

// GetPackageVersion returns the PackageVersion field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetPackageVersion() string {
	if o == nil || IsNil(o.PackageVersion) {
		var ret string
		return ret
	}
	return *o.PackageVersion
}

// GetPackageVersionOk returns a tuple with the PackageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetPackageVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PackageVersion) {
		return nil, false
	}
	return o.PackageVersion, true
}

// HasPackageVersion returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasPackageVersion() bool {
	if o != nil && !IsNil(o.PackageVersion) {
		return true
	}

	return false
}

// SetPackageVersion gets a reference to the given string and assigns it to the PackageVersion field.
func (o *FirmwareRunningFirmware) SetPackageVersion(v string) {
	o.PackageVersion = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FirmwareRunningFirmware) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FirmwareRunningFirmware) SetVersion(v string) {
	o.Version = &v
}

// GetBiosUnit returns the BiosUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetBiosUnit() BiosUnitRelationship {
	if o == nil || IsNil(o.BiosUnit.Get()) {
		var ret BiosUnitRelationship
		return ret
	}
	return *o.BiosUnit.Get()
}

// GetBiosUnitOk returns a tuple with the BiosUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetBiosUnitOk() (*BiosUnitRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.BiosUnit.Get(), o.BiosUnit.IsSet()
}

// HasBiosUnit returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasBiosUnit() bool {
	if o != nil && o.BiosUnit.IsSet() {
		return true
	}

	return false
}

// SetBiosUnit gets a reference to the given NullableBiosUnitRelationship and assigns it to the BiosUnit field.
func (o *FirmwareRunningFirmware) SetBiosUnit(v BiosUnitRelationship) {
	o.BiosUnit.Set(&v)
}

// SetBiosUnitNil sets the value for BiosUnit to be an explicit nil
func (o *FirmwareRunningFirmware) SetBiosUnitNil() {
	o.BiosUnit.Set(nil)
}

// UnsetBiosUnit ensures that no value is present for BiosUnit, not even an explicit nil
func (o *FirmwareRunningFirmware) UnsetBiosUnit() {
	o.BiosUnit.Unset()
}

// GetGraphicsCard returns the GraphicsCard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetGraphicsCard() GraphicsCardRelationship {
	if o == nil || IsNil(o.GraphicsCard.Get()) {
		var ret GraphicsCardRelationship
		return ret
	}
	return *o.GraphicsCard.Get()
}

// GetGraphicsCardOk returns a tuple with the GraphicsCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetGraphicsCardOk() (*GraphicsCardRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.GraphicsCard.Get(), o.GraphicsCard.IsSet()
}

// HasGraphicsCard returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasGraphicsCard() bool {
	if o != nil && o.GraphicsCard.IsSet() {
		return true
	}

	return false
}

// SetGraphicsCard gets a reference to the given NullableGraphicsCardRelationship and assigns it to the GraphicsCard field.
func (o *FirmwareRunningFirmware) SetGraphicsCard(v GraphicsCardRelationship) {
	o.GraphicsCard.Set(&v)
}

// SetGraphicsCardNil sets the value for GraphicsCard to be an explicit nil
func (o *FirmwareRunningFirmware) SetGraphicsCardNil() {
	o.GraphicsCard.Set(nil)
}

// UnsetGraphicsCard ensures that no value is present for GraphicsCard, not even an explicit nil
func (o *FirmwareRunningFirmware) UnsetGraphicsCard() {
	o.GraphicsCard.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *FirmwareRunningFirmware) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *FirmwareRunningFirmware) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *FirmwareRunningFirmware) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetManagementController returns the ManagementController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetManagementController() ManagementControllerRelationship {
	if o == nil || IsNil(o.ManagementController.Get()) {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.ManagementController.Get()
}

// GetManagementControllerOk returns a tuple with the ManagementController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetManagementControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagementController.Get(), o.ManagementController.IsSet()
}

// HasManagementController returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasManagementController() bool {
	if o != nil && o.ManagementController.IsSet() {
		return true
	}

	return false
}

// SetManagementController gets a reference to the given NullableManagementControllerRelationship and assigns it to the ManagementController field.
func (o *FirmwareRunningFirmware) SetManagementController(v ManagementControllerRelationship) {
	o.ManagementController.Set(&v)
}

// SetManagementControllerNil sets the value for ManagementController to be an explicit nil
func (o *FirmwareRunningFirmware) SetManagementControllerNil() {
	o.ManagementController.Set(nil)
}

// UnsetManagementController ensures that no value is present for ManagementController, not even an explicit nil
func (o *FirmwareRunningFirmware) UnsetManagementController() {
	o.ManagementController.Unset()
}

// GetNetworkElements returns the NetworkElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetNetworkElements() []NetworkElementRelationship {
	if o == nil {
		var ret []NetworkElementRelationship
		return ret
	}
	return o.NetworkElements
}

// GetNetworkElementsOk returns a tuple with the NetworkElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetNetworkElementsOk() ([]NetworkElementRelationship, bool) {
	if o == nil || IsNil(o.NetworkElements) {
		return nil, false
	}
	return o.NetworkElements, true
}

// HasNetworkElements returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasNetworkElements() bool {
	if o != nil && !IsNil(o.NetworkElements) {
		return true
	}

	return false
}

// SetNetworkElements gets a reference to the given []NetworkElementRelationship and assigns it to the NetworkElements field.
func (o *FirmwareRunningFirmware) SetNetworkElements(v []NetworkElementRelationship) {
	o.NetworkElements = v
}

// GetPciSwitch returns the PciSwitch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetPciSwitch() PciSwitchRelationship {
	if o == nil || IsNil(o.PciSwitch.Get()) {
		var ret PciSwitchRelationship
		return ret
	}
	return *o.PciSwitch.Get()
}

// GetPciSwitchOk returns a tuple with the PciSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetPciSwitchOk() (*PciSwitchRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PciSwitch.Get(), o.PciSwitch.IsSet()
}

// HasPciSwitch returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasPciSwitch() bool {
	if o != nil && o.PciSwitch.IsSet() {
		return true
	}

	return false
}

// SetPciSwitch gets a reference to the given NullablePciSwitchRelationship and assigns it to the PciSwitch field.
func (o *FirmwareRunningFirmware) SetPciSwitch(v PciSwitchRelationship) {
	o.PciSwitch.Set(&v)
}

// SetPciSwitchNil sets the value for PciSwitch to be an explicit nil
func (o *FirmwareRunningFirmware) SetPciSwitchNil() {
	o.PciSwitch.Set(nil)
}

// UnsetPciSwitch ensures that no value is present for PciSwitch, not even an explicit nil
func (o *FirmwareRunningFirmware) UnsetPciSwitch() {
	o.PciSwitch.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *FirmwareRunningFirmware) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *FirmwareRunningFirmware) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *FirmwareRunningFirmware) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetStorageController returns the StorageController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetStorageController() StorageControllerRelationship {
	if o == nil || IsNil(o.StorageController.Get()) {
		var ret StorageControllerRelationship
		return ret
	}
	return *o.StorageController.Get()
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetStorageControllerOk() (*StorageControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageController.Get(), o.StorageController.IsSet()
}

// HasStorageController returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasStorageController() bool {
	if o != nil && o.StorageController.IsSet() {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given NullableStorageControllerRelationship and assigns it to the StorageController field.
func (o *FirmwareRunningFirmware) SetStorageController(v StorageControllerRelationship) {
	o.StorageController.Set(&v)
}

// SetStorageControllerNil sets the value for StorageController to be an explicit nil
func (o *FirmwareRunningFirmware) SetStorageControllerNil() {
	o.StorageController.Set(nil)
}

// UnsetStorageController ensures that no value is present for StorageController, not even an explicit nil
func (o *FirmwareRunningFirmware) UnsetStorageController() {
	o.StorageController.Unset()
}

// GetStorageFlexFlashController returns the StorageFlexFlashController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship {
	if o == nil || IsNil(o.StorageFlexFlashController.Get()) {
		var ret StorageFlexFlashControllerRelationship
		return ret
	}
	return *o.StorageFlexFlashController.Get()
}

// GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageFlexFlashController.Get(), o.StorageFlexFlashController.IsSet()
}

// HasStorageFlexFlashController returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasStorageFlexFlashController() bool {
	if o != nil && o.StorageFlexFlashController.IsSet() {
		return true
	}

	return false
}

// SetStorageFlexFlashController gets a reference to the given NullableStorageFlexFlashControllerRelationship and assigns it to the StorageFlexFlashController field.
func (o *FirmwareRunningFirmware) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship) {
	o.StorageFlexFlashController.Set(&v)
}

// SetStorageFlexFlashControllerNil sets the value for StorageFlexFlashController to be an explicit nil
func (o *FirmwareRunningFirmware) SetStorageFlexFlashControllerNil() {
	o.StorageFlexFlashController.Set(nil)
}

// UnsetStorageFlexFlashController ensures that no value is present for StorageFlexFlashController, not even an explicit nil
func (o *FirmwareRunningFirmware) UnsetStorageFlexFlashController() {
	o.StorageFlexFlashController.Unset()
}

// GetStoragePhysicalDisk returns the StoragePhysicalDisk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetStoragePhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || IsNil(o.StoragePhysicalDisk.Get()) {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.StoragePhysicalDisk.Get()
}

// GetStoragePhysicalDiskOk returns a tuple with the StoragePhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetStoragePhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoragePhysicalDisk.Get(), o.StoragePhysicalDisk.IsSet()
}

// HasStoragePhysicalDisk returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasStoragePhysicalDisk() bool {
	if o != nil && o.StoragePhysicalDisk.IsSet() {
		return true
	}

	return false
}

// SetStoragePhysicalDisk gets a reference to the given NullableStoragePhysicalDiskRelationship and assigns it to the StoragePhysicalDisk field.
func (o *FirmwareRunningFirmware) SetStoragePhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.StoragePhysicalDisk.Set(&v)
}

// SetStoragePhysicalDiskNil sets the value for StoragePhysicalDisk to be an explicit nil
func (o *FirmwareRunningFirmware) SetStoragePhysicalDiskNil() {
	o.StoragePhysicalDisk.Set(nil)
}

// UnsetStoragePhysicalDisk ensures that no value is present for StoragePhysicalDisk, not even an explicit nil
func (o *FirmwareRunningFirmware) UnsetStoragePhysicalDisk() {
	o.StoragePhysicalDisk.Unset()
}

func (o FirmwareRunningFirmware) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareRunningFirmware) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Component) {
		toSerialize["Component"] = o.Component
	}
	if !IsNil(o.PackageVersion) {
		toSerialize["PackageVersion"] = o.PackageVersion
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if o.BiosUnit.IsSet() {
		toSerialize["BiosUnit"] = o.BiosUnit.Get()
	}
	if o.GraphicsCard.IsSet() {
		toSerialize["GraphicsCard"] = o.GraphicsCard.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.ManagementController.IsSet() {
		toSerialize["ManagementController"] = o.ManagementController.Get()
	}
	if o.NetworkElements != nil {
		toSerialize["NetworkElements"] = o.NetworkElements
	}
	if o.PciSwitch.IsSet() {
		toSerialize["PciSwitch"] = o.PciSwitch.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.StorageController.IsSet() {
		toSerialize["StorageController"] = o.StorageController.Get()
	}
	if o.StorageFlexFlashController.IsSet() {
		toSerialize["StorageFlexFlashController"] = o.StorageFlexFlashController.Get()
	}
	if o.StoragePhysicalDisk.IsSet() {
		toSerialize["StoragePhysicalDisk"] = o.StoragePhysicalDisk.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareRunningFirmware) UnmarshalJSON(data []byte) (err error) {
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
	type FirmwareRunningFirmwareWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Kind of the firmware - boot-booloader/system/kernel.
		Component *string `json:"Component,omitempty"`
		// Bundle version which the firmware belongs to.
		PackageVersion *string `json:"PackageVersion,omitempty"`
		// The type of the firmware.
		Type *string `json:"Type,omitempty"`
		// The version of the firmware.
		Version              *string                                  `json:"Version,omitempty"`
		BiosUnit             NullableBiosUnitRelationship             `json:"BiosUnit,omitempty"`
		GraphicsCard         NullableGraphicsCardRelationship         `json:"GraphicsCard,omitempty"`
		InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship  `json:"InventoryDeviceInfo,omitempty"`
		ManagementController NullableManagementControllerRelationship `json:"ManagementController,omitempty"`
		// An array of relationships to networkElement resources.
		NetworkElements            []NetworkElementRelationship                   `json:"NetworkElements,omitempty"`
		PciSwitch                  NullablePciSwitchRelationship                  `json:"PciSwitch,omitempty"`
		RegisteredDevice           NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
		StorageController          NullableStorageControllerRelationship          `json:"StorageController,omitempty"`
		StorageFlexFlashController NullableStorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
		StoragePhysicalDisk        NullableStoragePhysicalDiskRelationship        `json:"StoragePhysicalDisk,omitempty"`
	}

	varFirmwareRunningFirmwareWithoutEmbeddedStruct := FirmwareRunningFirmwareWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFirmwareRunningFirmwareWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareRunningFirmware := _FirmwareRunningFirmware{}
		varFirmwareRunningFirmware.ClassId = varFirmwareRunningFirmwareWithoutEmbeddedStruct.ClassId
		varFirmwareRunningFirmware.ObjectType = varFirmwareRunningFirmwareWithoutEmbeddedStruct.ObjectType
		varFirmwareRunningFirmware.Component = varFirmwareRunningFirmwareWithoutEmbeddedStruct.Component
		varFirmwareRunningFirmware.PackageVersion = varFirmwareRunningFirmwareWithoutEmbeddedStruct.PackageVersion
		varFirmwareRunningFirmware.Type = varFirmwareRunningFirmwareWithoutEmbeddedStruct.Type
		varFirmwareRunningFirmware.Version = varFirmwareRunningFirmwareWithoutEmbeddedStruct.Version
		varFirmwareRunningFirmware.BiosUnit = varFirmwareRunningFirmwareWithoutEmbeddedStruct.BiosUnit
		varFirmwareRunningFirmware.GraphicsCard = varFirmwareRunningFirmwareWithoutEmbeddedStruct.GraphicsCard
		varFirmwareRunningFirmware.InventoryDeviceInfo = varFirmwareRunningFirmwareWithoutEmbeddedStruct.InventoryDeviceInfo
		varFirmwareRunningFirmware.ManagementController = varFirmwareRunningFirmwareWithoutEmbeddedStruct.ManagementController
		varFirmwareRunningFirmware.NetworkElements = varFirmwareRunningFirmwareWithoutEmbeddedStruct.NetworkElements
		varFirmwareRunningFirmware.PciSwitch = varFirmwareRunningFirmwareWithoutEmbeddedStruct.PciSwitch
		varFirmwareRunningFirmware.RegisteredDevice = varFirmwareRunningFirmwareWithoutEmbeddedStruct.RegisteredDevice
		varFirmwareRunningFirmware.StorageController = varFirmwareRunningFirmwareWithoutEmbeddedStruct.StorageController
		varFirmwareRunningFirmware.StorageFlexFlashController = varFirmwareRunningFirmwareWithoutEmbeddedStruct.StorageFlexFlashController
		varFirmwareRunningFirmware.StoragePhysicalDisk = varFirmwareRunningFirmwareWithoutEmbeddedStruct.StoragePhysicalDisk
		*o = FirmwareRunningFirmware(varFirmwareRunningFirmware)
	} else {
		return err
	}

	varFirmwareRunningFirmware := _FirmwareRunningFirmware{}

	err = json.Unmarshal(data, &varFirmwareRunningFirmware)
	if err == nil {
		o.InventoryBase = varFirmwareRunningFirmware.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Component")
		delete(additionalProperties, "PackageVersion")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "BiosUnit")
		delete(additionalProperties, "GraphicsCard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "ManagementController")
		delete(additionalProperties, "NetworkElements")
		delete(additionalProperties, "PciSwitch")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageController")
		delete(additionalProperties, "StorageFlexFlashController")
		delete(additionalProperties, "StoragePhysicalDisk")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableFirmwareRunningFirmware struct {
	value *FirmwareRunningFirmware
	isSet bool
}

func (v NullableFirmwareRunningFirmware) Get() *FirmwareRunningFirmware {
	return v.value
}

func (v *NullableFirmwareRunningFirmware) Set(val *FirmwareRunningFirmware) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareRunningFirmware) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareRunningFirmware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareRunningFirmware(val *FirmwareRunningFirmware) *NullableFirmwareRunningFirmware {
	return &NullableFirmwareRunningFirmware{value: val, isSet: true}
}

func (v NullableFirmwareRunningFirmware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareRunningFirmware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
