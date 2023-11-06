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

// ComputeBladeIdentity Identity object that uniquely represents a blade server object under a DR.
type ComputeBladeIdentity struct {
	EquipmentPhysicalIdentity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Chassis Identifier of a blade server.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// The id of the chassis that the blade is currently located in.
	CurrentChassisId *int64 `json:"CurrentChassisId,omitempty"`
	// The slot number in the chassis that the blade is currently located in.
	CurrentSlotId *int64 `json:"CurrentSlotId,omitempty"`
	// Describes whether the running CIMC version supports Intersight managed mode. * `Unknown` - The running firmware version is unknown. * `Supported` - The running firmware version is known and supports IMM mode. * `NotSupported` - The running firmware version is known and does not support IMM mode.
	FirmwareSupportability *string `json:"FirmwareSupportability,omitempty"`
	// Chassis slot number of the manager compute server.
	ManagerSlotId *int64 `json:"ManagerSlotId,omitempty"`
	// The presence state of the blade server. * `Unknown` - The default presence state. * `Equipped` - The server is equipped in the slot. * `EquippedMismatch` - The slot is equipped, but there is another server currently inventoried in the slot. * `Missing` - The server is not present in the given slot.
	Presence *string `json:"Presence,omitempty"`
	// Chassis slot number of a blade server.
	SlotId *int64 `json:"SlotId,omitempty"`
	// An array of relationships to computeBladeIdentity resources.
	ManagedNodes         []ComputeBladeIdentityRelationship `json:"ManagedNodes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeBladeIdentity ComputeBladeIdentity

// NewComputeBladeIdentity instantiates a new ComputeBladeIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeBladeIdentity(classId string, objectType string) *ComputeBladeIdentity {
	this := ComputeBladeIdentity{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	return &this
}

// NewComputeBladeIdentityWithDefaults instantiates a new ComputeBladeIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeBladeIdentityWithDefaults() *ComputeBladeIdentity {
	this := ComputeBladeIdentity{}
	var classId string = "compute.BladeIdentity"
	this.ClassId = classId
	var objectType string = "compute.BladeIdentity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeBladeIdentity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeBladeIdentity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeBladeIdentity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeBladeIdentity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetChassisId() int64 {
	if o == nil || o.ChassisId == nil {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetChassisIdOk() (*int64, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *ComputeBladeIdentity) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetCurrentChassisId returns the CurrentChassisId field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetCurrentChassisId() int64 {
	if o == nil || o.CurrentChassisId == nil {
		var ret int64
		return ret
	}
	return *o.CurrentChassisId
}

// GetCurrentChassisIdOk returns a tuple with the CurrentChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetCurrentChassisIdOk() (*int64, bool) {
	if o == nil || o.CurrentChassisId == nil {
		return nil, false
	}
	return o.CurrentChassisId, true
}

// HasCurrentChassisId returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasCurrentChassisId() bool {
	if o != nil && o.CurrentChassisId != nil {
		return true
	}

	return false
}

// SetCurrentChassisId gets a reference to the given int64 and assigns it to the CurrentChassisId field.
func (o *ComputeBladeIdentity) SetCurrentChassisId(v int64) {
	o.CurrentChassisId = &v
}

// GetCurrentSlotId returns the CurrentSlotId field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetCurrentSlotId() int64 {
	if o == nil || o.CurrentSlotId == nil {
		var ret int64
		return ret
	}
	return *o.CurrentSlotId
}

// GetCurrentSlotIdOk returns a tuple with the CurrentSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetCurrentSlotIdOk() (*int64, bool) {
	if o == nil || o.CurrentSlotId == nil {
		return nil, false
	}
	return o.CurrentSlotId, true
}

// HasCurrentSlotId returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasCurrentSlotId() bool {
	if o != nil && o.CurrentSlotId != nil {
		return true
	}

	return false
}

// SetCurrentSlotId gets a reference to the given int64 and assigns it to the CurrentSlotId field.
func (o *ComputeBladeIdentity) SetCurrentSlotId(v int64) {
	o.CurrentSlotId = &v
}

// GetFirmwareSupportability returns the FirmwareSupportability field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetFirmwareSupportability() string {
	if o == nil || o.FirmwareSupportability == nil {
		var ret string
		return ret
	}
	return *o.FirmwareSupportability
}

// GetFirmwareSupportabilityOk returns a tuple with the FirmwareSupportability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetFirmwareSupportabilityOk() (*string, bool) {
	if o == nil || o.FirmwareSupportability == nil {
		return nil, false
	}
	return o.FirmwareSupportability, true
}

// HasFirmwareSupportability returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasFirmwareSupportability() bool {
	if o != nil && o.FirmwareSupportability != nil {
		return true
	}

	return false
}

// SetFirmwareSupportability gets a reference to the given string and assigns it to the FirmwareSupportability field.
func (o *ComputeBladeIdentity) SetFirmwareSupportability(v string) {
	o.FirmwareSupportability = &v
}

// GetManagerSlotId returns the ManagerSlotId field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetManagerSlotId() int64 {
	if o == nil || o.ManagerSlotId == nil {
		var ret int64
		return ret
	}
	return *o.ManagerSlotId
}

// GetManagerSlotIdOk returns a tuple with the ManagerSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetManagerSlotIdOk() (*int64, bool) {
	if o == nil || o.ManagerSlotId == nil {
		return nil, false
	}
	return o.ManagerSlotId, true
}

// HasManagerSlotId returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasManagerSlotId() bool {
	if o != nil && o.ManagerSlotId != nil {
		return true
	}

	return false
}

// SetManagerSlotId gets a reference to the given int64 and assigns it to the ManagerSlotId field.
func (o *ComputeBladeIdentity) SetManagerSlotId(v int64) {
	o.ManagerSlotId = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *ComputeBladeIdentity) SetPresence(v string) {
	o.Presence = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *ComputeBladeIdentity) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeIdentity) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *ComputeBladeIdentity) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetManagedNodes returns the ManagedNodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeIdentity) GetManagedNodes() []ComputeBladeIdentityRelationship {
	if o == nil {
		var ret []ComputeBladeIdentityRelationship
		return ret
	}
	return o.ManagedNodes
}

// GetManagedNodesOk returns a tuple with the ManagedNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeIdentity) GetManagedNodesOk() ([]ComputeBladeIdentityRelationship, bool) {
	if o == nil || o.ManagedNodes == nil {
		return nil, false
	}
	return o.ManagedNodes, true
}

// HasManagedNodes returns a boolean if a field has been set.
func (o *ComputeBladeIdentity) HasManagedNodes() bool {
	if o != nil && o.ManagedNodes != nil {
		return true
	}

	return false
}

// SetManagedNodes gets a reference to the given []ComputeBladeIdentityRelationship and assigns it to the ManagedNodes field.
func (o *ComputeBladeIdentity) SetManagedNodes(v []ComputeBladeIdentityRelationship) {
	o.ManagedNodes = v
}

func (o ComputeBladeIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentPhysicalIdentity, errEquipmentPhysicalIdentity := json.Marshal(o.EquipmentPhysicalIdentity)
	if errEquipmentPhysicalIdentity != nil {
		return []byte{}, errEquipmentPhysicalIdentity
	}
	errEquipmentPhysicalIdentity = json.Unmarshal([]byte(serializedEquipmentPhysicalIdentity), &toSerialize)
	if errEquipmentPhysicalIdentity != nil {
		return []byte{}, errEquipmentPhysicalIdentity
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.CurrentChassisId != nil {
		toSerialize["CurrentChassisId"] = o.CurrentChassisId
	}
	if o.CurrentSlotId != nil {
		toSerialize["CurrentSlotId"] = o.CurrentSlotId
	}
	if o.FirmwareSupportability != nil {
		toSerialize["FirmwareSupportability"] = o.FirmwareSupportability
	}
	if o.ManagerSlotId != nil {
		toSerialize["ManagerSlotId"] = o.ManagerSlotId
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.ManagedNodes != nil {
		toSerialize["ManagedNodes"] = o.ManagedNodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeBladeIdentity) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeBladeIdentityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Chassis Identifier of a blade server.
		ChassisId *int64 `json:"ChassisId,omitempty"`
		// The id of the chassis that the blade is currently located in.
		CurrentChassisId *int64 `json:"CurrentChassisId,omitempty"`
		// The slot number in the chassis that the blade is currently located in.
		CurrentSlotId *int64 `json:"CurrentSlotId,omitempty"`
		// Describes whether the running CIMC version supports Intersight managed mode. * `Unknown` - The running firmware version is unknown. * `Supported` - The running firmware version is known and supports IMM mode. * `NotSupported` - The running firmware version is known and does not support IMM mode.
		FirmwareSupportability *string `json:"FirmwareSupportability,omitempty"`
		// Chassis slot number of the manager compute server.
		ManagerSlotId *int64 `json:"ManagerSlotId,omitempty"`
		// The presence state of the blade server. * `Unknown` - The default presence state. * `Equipped` - The server is equipped in the slot. * `EquippedMismatch` - The slot is equipped, but there is another server currently inventoried in the slot. * `Missing` - The server is not present in the given slot.
		Presence *string `json:"Presence,omitempty"`
		// Chassis slot number of a blade server.
		SlotId *int64 `json:"SlotId,omitempty"`
		// An array of relationships to computeBladeIdentity resources.
		ManagedNodes []ComputeBladeIdentityRelationship `json:"ManagedNodes,omitempty"`
	}

	varComputeBladeIdentityWithoutEmbeddedStruct := ComputeBladeIdentityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeBladeIdentityWithoutEmbeddedStruct)
	if err == nil {
		varComputeBladeIdentity := _ComputeBladeIdentity{}
		varComputeBladeIdentity.ClassId = varComputeBladeIdentityWithoutEmbeddedStruct.ClassId
		varComputeBladeIdentity.ObjectType = varComputeBladeIdentityWithoutEmbeddedStruct.ObjectType
		varComputeBladeIdentity.ChassisId = varComputeBladeIdentityWithoutEmbeddedStruct.ChassisId
		varComputeBladeIdentity.CurrentChassisId = varComputeBladeIdentityWithoutEmbeddedStruct.CurrentChassisId
		varComputeBladeIdentity.CurrentSlotId = varComputeBladeIdentityWithoutEmbeddedStruct.CurrentSlotId
		varComputeBladeIdentity.FirmwareSupportability = varComputeBladeIdentityWithoutEmbeddedStruct.FirmwareSupportability
		varComputeBladeIdentity.ManagerSlotId = varComputeBladeIdentityWithoutEmbeddedStruct.ManagerSlotId
		varComputeBladeIdentity.Presence = varComputeBladeIdentityWithoutEmbeddedStruct.Presence
		varComputeBladeIdentity.SlotId = varComputeBladeIdentityWithoutEmbeddedStruct.SlotId
		varComputeBladeIdentity.ManagedNodes = varComputeBladeIdentityWithoutEmbeddedStruct.ManagedNodes
		*o = ComputeBladeIdentity(varComputeBladeIdentity)
	} else {
		return err
	}

	varComputeBladeIdentity := _ComputeBladeIdentity{}

	err = json.Unmarshal(bytes, &varComputeBladeIdentity)
	if err == nil {
		o.EquipmentPhysicalIdentity = varComputeBladeIdentity.EquipmentPhysicalIdentity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "CurrentChassisId")
		delete(additionalProperties, "CurrentSlotId")
		delete(additionalProperties, "FirmwareSupportability")
		delete(additionalProperties, "ManagerSlotId")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "ManagedNodes")

		// remove fields from embedded structs
		reflectEquipmentPhysicalIdentity := reflect.ValueOf(o.EquipmentPhysicalIdentity)
		for i := 0; i < reflectEquipmentPhysicalIdentity.Type().NumField(); i++ {
			t := reflectEquipmentPhysicalIdentity.Type().Field(i)

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

type NullableComputeBladeIdentity struct {
	value *ComputeBladeIdentity
	isSet bool
}

func (v NullableComputeBladeIdentity) Get() *ComputeBladeIdentity {
	return v.value
}

func (v *NullableComputeBladeIdentity) Set(val *ComputeBladeIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeBladeIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeBladeIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeBladeIdentity(val *ComputeBladeIdentity) *NullableComputeBladeIdentity {
	return &NullableComputeBladeIdentity{value: val, isSet: true}
}

func (v NullableComputeBladeIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeBladeIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
