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

// checks if the StorageFlexFlashVirtualDrive type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageFlexFlashVirtualDrive{}

// StorageFlexFlashVirtualDrive Virtual Drive repersenting a SD Card.
type StorageFlexFlashVirtualDrive struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The drive scope of the flex flash virtual drive.
	DriveScope *string `json:"DriveScope,omitempty"`
	// Status of virtual drive on the flex controller.
	DriveStatus *string `json:"DriveStatus,omitempty"`
	// The partition Id of the flex flash virtual Drive.
	PartitionId *string `json:"PartitionId,omitempty"`
	// The resident image on the flex flash virtual Drive.
	ResidentImage *string `json:"ResidentImage,omitempty"`
	// Size of virtual drive on the flex controller.
	Size *string `json:"Size,omitempty"`
	// Virtual drive on the flex flash controller.
	VirtualDrive               *string                                        `json:"VirtualDrive,omitempty"`
	InventoryDeviceInfo        NullableInventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice           NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	StorageFlexFlashController NullableStorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _StorageFlexFlashVirtualDrive StorageFlexFlashVirtualDrive

// NewStorageFlexFlashVirtualDrive instantiates a new StorageFlexFlashVirtualDrive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexFlashVirtualDrive(classId string, objectType string) *StorageFlexFlashVirtualDrive {
	this := StorageFlexFlashVirtualDrive{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexFlashVirtualDriveWithDefaults instantiates a new StorageFlexFlashVirtualDrive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexFlashVirtualDriveWithDefaults() *StorageFlexFlashVirtualDrive {
	this := StorageFlexFlashVirtualDrive{}
	var classId string = "storage.FlexFlashVirtualDrive"
	this.ClassId = classId
	var objectType string = "storage.FlexFlashVirtualDrive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexFlashVirtualDrive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexFlashVirtualDrive) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.FlexFlashVirtualDrive" of the ClassId field.
func (o *StorageFlexFlashVirtualDrive) GetDefaultClassId() interface{} {
	return "storage.FlexFlashVirtualDrive"
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexFlashVirtualDrive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexFlashVirtualDrive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.FlexFlashVirtualDrive" of the ObjectType field.
func (o *StorageFlexFlashVirtualDrive) GetDefaultObjectType() interface{} {
	return "storage.FlexFlashVirtualDrive"
}

// GetDriveScope returns the DriveScope field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetDriveScope() string {
	if o == nil || IsNil(o.DriveScope) {
		var ret string
		return ret
	}
	return *o.DriveScope
}

// GetDriveScopeOk returns a tuple with the DriveScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetDriveScopeOk() (*string, bool) {
	if o == nil || IsNil(o.DriveScope) {
		return nil, false
	}
	return o.DriveScope, true
}

// HasDriveScope returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasDriveScope() bool {
	if o != nil && !IsNil(o.DriveScope) {
		return true
	}

	return false
}

// SetDriveScope gets a reference to the given string and assigns it to the DriveScope field.
func (o *StorageFlexFlashVirtualDrive) SetDriveScope(v string) {
	o.DriveScope = &v
}

// GetDriveStatus returns the DriveStatus field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetDriveStatus() string {
	if o == nil || IsNil(o.DriveStatus) {
		var ret string
		return ret
	}
	return *o.DriveStatus
}

// GetDriveStatusOk returns a tuple with the DriveStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetDriveStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DriveStatus) {
		return nil, false
	}
	return o.DriveStatus, true
}

// HasDriveStatus returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasDriveStatus() bool {
	if o != nil && !IsNil(o.DriveStatus) {
		return true
	}

	return false
}

// SetDriveStatus gets a reference to the given string and assigns it to the DriveStatus field.
func (o *StorageFlexFlashVirtualDrive) SetDriveStatus(v string) {
	o.DriveStatus = &v
}

// GetPartitionId returns the PartitionId field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetPartitionId() string {
	if o == nil || IsNil(o.PartitionId) {
		var ret string
		return ret
	}
	return *o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetPartitionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartitionId) {
		return nil, false
	}
	return o.PartitionId, true
}

// HasPartitionId returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasPartitionId() bool {
	if o != nil && !IsNil(o.PartitionId) {
		return true
	}

	return false
}

// SetPartitionId gets a reference to the given string and assigns it to the PartitionId field.
func (o *StorageFlexFlashVirtualDrive) SetPartitionId(v string) {
	o.PartitionId = &v
}

// GetResidentImage returns the ResidentImage field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetResidentImage() string {
	if o == nil || IsNil(o.ResidentImage) {
		var ret string
		return ret
	}
	return *o.ResidentImage
}

// GetResidentImageOk returns a tuple with the ResidentImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetResidentImageOk() (*string, bool) {
	if o == nil || IsNil(o.ResidentImage) {
		return nil, false
	}
	return o.ResidentImage, true
}

// HasResidentImage returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasResidentImage() bool {
	if o != nil && !IsNil(o.ResidentImage) {
		return true
	}

	return false
}

// SetResidentImage gets a reference to the given string and assigns it to the ResidentImage field.
func (o *StorageFlexFlashVirtualDrive) SetResidentImage(v string) {
	o.ResidentImage = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StorageFlexFlashVirtualDrive) SetSize(v string) {
	o.Size = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetVirtualDrive() string {
	if o == nil || IsNil(o.VirtualDrive) {
		var ret string
		return ret
	}
	return *o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetVirtualDriveOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualDrive) {
		return nil, false
	}
	return o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasVirtualDrive() bool {
	if o != nil && !IsNil(o.VirtualDrive) {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given string and assigns it to the VirtualDrive field.
func (o *StorageFlexFlashVirtualDrive) SetVirtualDrive(v string) {
	o.VirtualDrive = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashVirtualDrive) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashVirtualDrive) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexFlashVirtualDrive) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *StorageFlexFlashVirtualDrive) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *StorageFlexFlashVirtualDrive) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashVirtualDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashVirtualDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexFlashVirtualDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageFlexFlashVirtualDrive) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageFlexFlashVirtualDrive) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetStorageFlexFlashController returns the StorageFlexFlashController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashVirtualDrive) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship {
	if o == nil || IsNil(o.StorageFlexFlashController.Get()) {
		var ret StorageFlexFlashControllerRelationship
		return ret
	}
	return *o.StorageFlexFlashController.Get()
}

// GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashVirtualDrive) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageFlexFlashController.Get(), o.StorageFlexFlashController.IsSet()
}

// HasStorageFlexFlashController returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasStorageFlexFlashController() bool {
	if o != nil && o.StorageFlexFlashController.IsSet() {
		return true
	}

	return false
}

// SetStorageFlexFlashController gets a reference to the given NullableStorageFlexFlashControllerRelationship and assigns it to the StorageFlexFlashController field.
func (o *StorageFlexFlashVirtualDrive) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship) {
	o.StorageFlexFlashController.Set(&v)
}

// SetStorageFlexFlashControllerNil sets the value for StorageFlexFlashController to be an explicit nil
func (o *StorageFlexFlashVirtualDrive) SetStorageFlexFlashControllerNil() {
	o.StorageFlexFlashController.Set(nil)
}

// UnsetStorageFlexFlashController ensures that no value is present for StorageFlexFlashController, not even an explicit nil
func (o *StorageFlexFlashVirtualDrive) UnsetStorageFlexFlashController() {
	o.StorageFlexFlashController.Unset()
}

func (o StorageFlexFlashVirtualDrive) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageFlexFlashVirtualDrive) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DriveScope) {
		toSerialize["DriveScope"] = o.DriveScope
	}
	if !IsNil(o.DriveStatus) {
		toSerialize["DriveStatus"] = o.DriveStatus
	}
	if !IsNil(o.PartitionId) {
		toSerialize["PartitionId"] = o.PartitionId
	}
	if !IsNil(o.ResidentImage) {
		toSerialize["ResidentImage"] = o.ResidentImage
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	if !IsNil(o.VirtualDrive) {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.StorageFlexFlashController.IsSet() {
		toSerialize["StorageFlexFlashController"] = o.StorageFlexFlashController.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageFlexFlashVirtualDrive) UnmarshalJSON(data []byte) (err error) {
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
	type StorageFlexFlashVirtualDriveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The drive scope of the flex flash virtual drive.
		DriveScope *string `json:"DriveScope,omitempty"`
		// Status of virtual drive on the flex controller.
		DriveStatus *string `json:"DriveStatus,omitempty"`
		// The partition Id of the flex flash virtual Drive.
		PartitionId *string `json:"PartitionId,omitempty"`
		// The resident image on the flex flash virtual Drive.
		ResidentImage *string `json:"ResidentImage,omitempty"`
		// Size of virtual drive on the flex controller.
		Size *string `json:"Size,omitempty"`
		// Virtual drive on the flex flash controller.
		VirtualDrive               *string                                        `json:"VirtualDrive,omitempty"`
		InventoryDeviceInfo        NullableInventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice           NullableAssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
		StorageFlexFlashController NullableStorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	}

	varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct := StorageFlexFlashVirtualDriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct)
	if err == nil {
		varStorageFlexFlashVirtualDrive := _StorageFlexFlashVirtualDrive{}
		varStorageFlexFlashVirtualDrive.ClassId = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.ClassId
		varStorageFlexFlashVirtualDrive.ObjectType = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.ObjectType
		varStorageFlexFlashVirtualDrive.DriveScope = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.DriveScope
		varStorageFlexFlashVirtualDrive.DriveStatus = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.DriveStatus
		varStorageFlexFlashVirtualDrive.PartitionId = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.PartitionId
		varStorageFlexFlashVirtualDrive.ResidentImage = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.ResidentImage
		varStorageFlexFlashVirtualDrive.Size = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.Size
		varStorageFlexFlashVirtualDrive.VirtualDrive = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.VirtualDrive
		varStorageFlexFlashVirtualDrive.InventoryDeviceInfo = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageFlexFlashVirtualDrive.RegisteredDevice = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.RegisteredDevice
		varStorageFlexFlashVirtualDrive.StorageFlexFlashController = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.StorageFlexFlashController
		*o = StorageFlexFlashVirtualDrive(varStorageFlexFlashVirtualDrive)
	} else {
		return err
	}

	varStorageFlexFlashVirtualDrive := _StorageFlexFlashVirtualDrive{}

	err = json.Unmarshal(data, &varStorageFlexFlashVirtualDrive)
	if err == nil {
		o.EquipmentBase = varStorageFlexFlashVirtualDrive.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriveScope")
		delete(additionalProperties, "DriveStatus")
		delete(additionalProperties, "PartitionId")
		delete(additionalProperties, "ResidentImage")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "VirtualDrive")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageFlexFlashController")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableStorageFlexFlashVirtualDrive struct {
	value *StorageFlexFlashVirtualDrive
	isSet bool
}

func (v NullableStorageFlexFlashVirtualDrive) Get() *StorageFlexFlashVirtualDrive {
	return v.value
}

func (v *NullableStorageFlexFlashVirtualDrive) Set(val *StorageFlexFlashVirtualDrive) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexFlashVirtualDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexFlashVirtualDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexFlashVirtualDrive(val *StorageFlexFlashVirtualDrive) *NullableStorageFlexFlashVirtualDrive {
	return &NullableStorageFlexFlashVirtualDrive{value: val, isSet: true}
}

func (v NullableStorageFlexFlashVirtualDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexFlashVirtualDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
