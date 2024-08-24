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

// checks if the PciNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PciNode{}

// PciNode External pci nodes connected to a server.
type PciNode struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The id of the chassis that the pcie node is currently located in.
	ChassisId  *string  `json:"ChassisId,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	// Operational state of the pcie node.
	OperState *string `json:"OperState,omitempty"`
	// The slot number in the chassis that the pcie node is currently located in.
	SlotId       *string                          `json:"SlotId,omitempty"`
	ComputeBlade NullableComputeBladeRelationship `json:"ComputeBlade,omitempty"`
	// An array of relationships to graphicsCard resources.
	GraphicsCards        []GraphicsCardRelationship                  `json:"GraphicsCards,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PciNode PciNode

// NewPciNode instantiates a new PciNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciNode(classId string, objectType string) *PciNode {
	this := PciNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPciNodeWithDefaults instantiates a new PciNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciNodeWithDefaults() *PciNode {
	this := PciNode{}
	var classId string = "pci.Node"
	this.ClassId = classId
	var objectType string = "pci.Node"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PciNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PciNode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PciNode) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "pci.Node" of the ClassId field.
func (o *PciNode) GetDefaultClassId() interface{} {
	return "pci.Node"
}

// GetObjectType returns the ObjectType field value
func (o *PciNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PciNode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PciNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "pci.Node" of the ObjectType field.
func (o *PciNode) GetDefaultObjectType() interface{} {
	return "pci.Node"
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *PciNode) GetChassisId() string {
	if o == nil || IsNil(o.ChassisId) {
		var ret string
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciNode) GetChassisIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChassisId) {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *PciNode) HasChassisId() bool {
	if o != nil && !IsNil(o.ChassisId) {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given string and assigns it to the ChassisId field.
func (o *PciNode) SetChassisId(v string) {
	o.ChassisId = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciNode) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciNode) GetOperReasonOk() ([]string, bool) {
	if o == nil || IsNil(o.OperReason) {
		return nil, false
	}
	return o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *PciNode) HasOperReason() bool {
	if o != nil && !IsNil(o.OperReason) {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *PciNode) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *PciNode) GetOperState() string {
	if o == nil || IsNil(o.OperState) {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciNode) GetOperStateOk() (*string, bool) {
	if o == nil || IsNil(o.OperState) {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *PciNode) HasOperState() bool {
	if o != nil && !IsNil(o.OperState) {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *PciNode) SetOperState(v string) {
	o.OperState = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *PciNode) GetSlotId() string {
	if o == nil || IsNil(o.SlotId) {
		var ret string
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciNode) GetSlotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SlotId) {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *PciNode) HasSlotId() bool {
	if o != nil && !IsNil(o.SlotId) {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given string and assigns it to the SlotId field.
func (o *PciNode) SetSlotId(v string) {
	o.SlotId = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciNode) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || IsNil(o.ComputeBlade.Get()) {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade.Get()
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciNode) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBlade.Get(), o.ComputeBlade.IsSet()
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *PciNode) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade.IsSet() {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given NullableComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *PciNode) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade.Set(&v)
}

// SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil
func (o *PciNode) SetComputeBladeNil() {
	o.ComputeBlade.Set(nil)
}

// UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
func (o *PciNode) UnsetComputeBlade() {
	o.ComputeBlade.Unset()
}

// GetGraphicsCards returns the GraphicsCards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciNode) GetGraphicsCards() []GraphicsCardRelationship {
	if o == nil {
		var ret []GraphicsCardRelationship
		return ret
	}
	return o.GraphicsCards
}

// GetGraphicsCardsOk returns a tuple with the GraphicsCards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciNode) GetGraphicsCardsOk() ([]GraphicsCardRelationship, bool) {
	if o == nil || IsNil(o.GraphicsCards) {
		return nil, false
	}
	return o.GraphicsCards, true
}

// HasGraphicsCards returns a boolean if a field has been set.
func (o *PciNode) HasGraphicsCards() bool {
	if o != nil && !IsNil(o.GraphicsCards) {
		return true
	}

	return false
}

// SetGraphicsCards gets a reference to the given []GraphicsCardRelationship and assigns it to the GraphicsCards field.
func (o *PciNode) SetGraphicsCards(v []GraphicsCardRelationship) {
	o.GraphicsCards = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciNode) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciNode) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *PciNode) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *PciNode) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *PciNode) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *PciNode) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciNode) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciNode) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PciNode) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PciNode) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *PciNode) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *PciNode) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o PciNode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PciNode) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ChassisId) {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if !IsNil(o.OperState) {
		toSerialize["OperState"] = o.OperState
	}
	if !IsNil(o.SlotId) {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.ComputeBlade.IsSet() {
		toSerialize["ComputeBlade"] = o.ComputeBlade.Get()
	}
	if o.GraphicsCards != nil {
		toSerialize["GraphicsCards"] = o.GraphicsCards
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PciNode) UnmarshalJSON(data []byte) (err error) {
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
	type PciNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The id of the chassis that the pcie node is currently located in.
		ChassisId  *string  `json:"ChassisId,omitempty"`
		OperReason []string `json:"OperReason,omitempty"`
		// Operational state of the pcie node.
		OperState *string `json:"OperState,omitempty"`
		// The slot number in the chassis that the pcie node is currently located in.
		SlotId       *string                          `json:"SlotId,omitempty"`
		ComputeBlade NullableComputeBladeRelationship `json:"ComputeBlade,omitempty"`
		// An array of relationships to graphicsCard resources.
		GraphicsCards       []GraphicsCardRelationship                  `json:"GraphicsCards,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varPciNodeWithoutEmbeddedStruct := PciNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPciNodeWithoutEmbeddedStruct)
	if err == nil {
		varPciNode := _PciNode{}
		varPciNode.ClassId = varPciNodeWithoutEmbeddedStruct.ClassId
		varPciNode.ObjectType = varPciNodeWithoutEmbeddedStruct.ObjectType
		varPciNode.ChassisId = varPciNodeWithoutEmbeddedStruct.ChassisId
		varPciNode.OperReason = varPciNodeWithoutEmbeddedStruct.OperReason
		varPciNode.OperState = varPciNodeWithoutEmbeddedStruct.OperState
		varPciNode.SlotId = varPciNodeWithoutEmbeddedStruct.SlotId
		varPciNode.ComputeBlade = varPciNodeWithoutEmbeddedStruct.ComputeBlade
		varPciNode.GraphicsCards = varPciNodeWithoutEmbeddedStruct.GraphicsCards
		varPciNode.InventoryDeviceInfo = varPciNodeWithoutEmbeddedStruct.InventoryDeviceInfo
		varPciNode.RegisteredDevice = varPciNodeWithoutEmbeddedStruct.RegisteredDevice
		*o = PciNode(varPciNode)
	} else {
		return err
	}

	varPciNode := _PciNode{}

	err = json.Unmarshal(data, &varPciNode)
	if err == nil {
		o.EquipmentBase = varPciNode.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "GraphicsCards")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

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

type NullablePciNode struct {
	value *PciNode
	isSet bool
}

func (v NullablePciNode) Get() *PciNode {
	return v.value
}

func (v *NullablePciNode) Set(val *PciNode) {
	v.value = val
	v.isSet = true
}

func (v NullablePciNode) IsSet() bool {
	return v.isSet
}

func (v *NullablePciNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciNode(val *PciNode) *NullablePciNode {
	return &NullablePciNode{value: val, isSet: true}
}

func (v NullablePciNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
