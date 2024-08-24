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

// checks if the StorageNetAppEthernetPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppEthernetPort{}

// StorageNetAppEthernetPort Ethernet port is a port on a node in a storage array.
type StorageNetAppEthernetPort struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the broadcast domain, scoped to its IPspace.
	BroadcastDomainName *string `json:"BroadcastDomainName,omitempty"`
	// Status of port to determine if its enabled or not.
	Enabled *string `json:"Enabled,omitempty"`
	// MAC address of the port available in storage array.
	MacAddress *string `json:"MacAddress,omitempty" validate:"regexp=^$|^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"`
	// Maximum transmission unit of the physical port available in storage array.
	Mtu *int64 `json:"Mtu,omitempty"`
	// Name of the port available in storage array.
	Name                   *string                               `json:"Name,omitempty"`
	NetAppEthernetPortLag  NullableStorageNetAppEthernetPortLag  `json:"NetAppEthernetPortLag,omitempty"`
	NetAppEthernetPortVlan NullableStorageNetAppEthernetPortVlan `json:"NetAppEthernetPortVlan,omitempty"`
	// The node name for the port.
	NodeName *string `json:"NodeName,omitempty"`
	// State of the port available in storage array. * `Down` - An inactive port is listed as Down. * `Up` - An active port is listed as Up. * `Degraded` - An active port that is Up but unhealthy.
	PortState *string `json:"PortState,omitempty"`
	// Operational speed of port measured.
	Speed *int64 `json:"Speed,omitempty"`
	// State of the port available in storage array. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
	// Deprecated
	State *string `json:"State,omitempty"`
	// Type of the port available in storage array. * `LAG` - Storage port of type lag. * `physical` - LIFs can be configured directly on physical ports. * `VLAN` - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port.
	Type *string `json:"Type,omitempty"`
	// Universally unique identifier of the physical port.
	Uuid            *string                               `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	ArrayController NullableStorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppEthernetPortEvent resources.
	Events               []StorageNetAppEthernetPortEventRelationship `json:"Events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppEthernetPort StorageNetAppEthernetPort

// NewStorageNetAppEthernetPort instantiates a new StorageNetAppEthernetPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppEthernetPort(classId string, objectType string) *StorageNetAppEthernetPort {
	this := StorageNetAppEthernetPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppEthernetPortWithDefaults instantiates a new StorageNetAppEthernetPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppEthernetPortWithDefaults() *StorageNetAppEthernetPort {
	this := StorageNetAppEthernetPort{}
	var classId string = "storage.NetAppEthernetPort"
	this.ClassId = classId
	var objectType string = "storage.NetAppEthernetPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppEthernetPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppEthernetPort) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppEthernetPort" of the ClassId field.
func (o *StorageNetAppEthernetPort) GetDefaultClassId() interface{} {
	return "storage.NetAppEthernetPort"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppEthernetPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppEthernetPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppEthernetPort" of the ObjectType field.
func (o *StorageNetAppEthernetPort) GetDefaultObjectType() interface{} {
	return "storage.NetAppEthernetPort"
}

// GetBroadcastDomainName returns the BroadcastDomainName field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetBroadcastDomainName() string {
	if o == nil || IsNil(o.BroadcastDomainName) {
		var ret string
		return ret
	}
	return *o.BroadcastDomainName
}

// GetBroadcastDomainNameOk returns a tuple with the BroadcastDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetBroadcastDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.BroadcastDomainName) {
		return nil, false
	}
	return o.BroadcastDomainName, true
}

// HasBroadcastDomainName returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasBroadcastDomainName() bool {
	if o != nil && !IsNil(o.BroadcastDomainName) {
		return true
	}

	return false
}

// SetBroadcastDomainName gets a reference to the given string and assigns it to the BroadcastDomainName field.
func (o *StorageNetAppEthernetPort) SetBroadcastDomainName(v string) {
	o.BroadcastDomainName = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetEnabled() string {
	if o == nil || IsNil(o.Enabled) {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppEthernetPort) SetEnabled(v string) {
	o.Enabled = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *StorageNetAppEthernetPort) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetMtu() int64 {
	if o == nil || IsNil(o.Mtu) {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetMtuOk() (*int64, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *StorageNetAppEthernetPort) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppEthernetPort) SetName(v string) {
	o.Name = &v
}

// GetNetAppEthernetPortLag returns the NetAppEthernetPortLag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppEthernetPort) GetNetAppEthernetPortLag() StorageNetAppEthernetPortLag {
	if o == nil || IsNil(o.NetAppEthernetPortLag.Get()) {
		var ret StorageNetAppEthernetPortLag
		return ret
	}
	return *o.NetAppEthernetPortLag.Get()
}

// GetNetAppEthernetPortLagOk returns a tuple with the NetAppEthernetPortLag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppEthernetPort) GetNetAppEthernetPortLagOk() (*StorageNetAppEthernetPortLag, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetAppEthernetPortLag.Get(), o.NetAppEthernetPortLag.IsSet()
}

// HasNetAppEthernetPortLag returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasNetAppEthernetPortLag() bool {
	if o != nil && o.NetAppEthernetPortLag.IsSet() {
		return true
	}

	return false
}

// SetNetAppEthernetPortLag gets a reference to the given NullableStorageNetAppEthernetPortLag and assigns it to the NetAppEthernetPortLag field.
func (o *StorageNetAppEthernetPort) SetNetAppEthernetPortLag(v StorageNetAppEthernetPortLag) {
	o.NetAppEthernetPortLag.Set(&v)
}

// SetNetAppEthernetPortLagNil sets the value for NetAppEthernetPortLag to be an explicit nil
func (o *StorageNetAppEthernetPort) SetNetAppEthernetPortLagNil() {
	o.NetAppEthernetPortLag.Set(nil)
}

// UnsetNetAppEthernetPortLag ensures that no value is present for NetAppEthernetPortLag, not even an explicit nil
func (o *StorageNetAppEthernetPort) UnsetNetAppEthernetPortLag() {
	o.NetAppEthernetPortLag.Unset()
}

// GetNetAppEthernetPortVlan returns the NetAppEthernetPortVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppEthernetPort) GetNetAppEthernetPortVlan() StorageNetAppEthernetPortVlan {
	if o == nil || IsNil(o.NetAppEthernetPortVlan.Get()) {
		var ret StorageNetAppEthernetPortVlan
		return ret
	}
	return *o.NetAppEthernetPortVlan.Get()
}

// GetNetAppEthernetPortVlanOk returns a tuple with the NetAppEthernetPortVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppEthernetPort) GetNetAppEthernetPortVlanOk() (*StorageNetAppEthernetPortVlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetAppEthernetPortVlan.Get(), o.NetAppEthernetPortVlan.IsSet()
}

// HasNetAppEthernetPortVlan returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasNetAppEthernetPortVlan() bool {
	if o != nil && o.NetAppEthernetPortVlan.IsSet() {
		return true
	}

	return false
}

// SetNetAppEthernetPortVlan gets a reference to the given NullableStorageNetAppEthernetPortVlan and assigns it to the NetAppEthernetPortVlan field.
func (o *StorageNetAppEthernetPort) SetNetAppEthernetPortVlan(v StorageNetAppEthernetPortVlan) {
	o.NetAppEthernetPortVlan.Set(&v)
}

// SetNetAppEthernetPortVlanNil sets the value for NetAppEthernetPortVlan to be an explicit nil
func (o *StorageNetAppEthernetPort) SetNetAppEthernetPortVlanNil() {
	o.NetAppEthernetPortVlan.Set(nil)
}

// UnsetNetAppEthernetPortVlan ensures that no value is present for NetAppEthernetPortVlan, not even an explicit nil
func (o *StorageNetAppEthernetPort) UnsetNetAppEthernetPortVlan() {
	o.NetAppEthernetPortVlan.Unset()
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetNodeName() string {
	if o == nil || IsNil(o.NodeName) {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetNodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.NodeName) {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasNodeName() bool {
	if o != nil && !IsNil(o.NodeName) {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *StorageNetAppEthernetPort) SetNodeName(v string) {
	o.NodeName = &v
}

// GetPortState returns the PortState field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetPortState() string {
	if o == nil || IsNil(o.PortState) {
		var ret string
		return ret
	}
	return *o.PortState
}

// GetPortStateOk returns a tuple with the PortState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetPortStateOk() (*string, bool) {
	if o == nil || IsNil(o.PortState) {
		return nil, false
	}
	return o.PortState, true
}

// HasPortState returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasPortState() bool {
	if o != nil && !IsNil(o.PortState) {
		return true
	}

	return false
}

// SetPortState gets a reference to the given string and assigns it to the PortState field.
func (o *StorageNetAppEthernetPort) SetPortState(v string) {
	o.PortState = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetSpeed() int64 {
	if o == nil || IsNil(o.Speed) {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetSpeedOk() (*int64, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *StorageNetAppEthernetPort) SetSpeed(v int64) {
	o.Speed = &v
}

// GetState returns the State field value if set, zero value otherwise.
// Deprecated
func (o *StorageNetAppEthernetPort) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *StorageNetAppEthernetPort) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
// Deprecated
func (o *StorageNetAppEthernetPort) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageNetAppEthernetPort) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppEthernetPort) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppEthernetPort) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || IsNil(o.ArrayController.Get()) {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController.Get()
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppEthernetPort) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArrayController.Get(), o.ArrayController.IsSet()
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasArrayController() bool {
	if o != nil && o.ArrayController.IsSet() {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given NullableStorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppEthernetPort) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController.Set(&v)
}

// SetArrayControllerNil sets the value for ArrayController to be an explicit nil
func (o *StorageNetAppEthernetPort) SetArrayControllerNil() {
	o.ArrayController.Set(nil)
}

// UnsetArrayController ensures that no value is present for ArrayController, not even an explicit nil
func (o *StorageNetAppEthernetPort) UnsetArrayController() {
	o.ArrayController.Unset()
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppEthernetPort) GetEvents() []StorageNetAppEthernetPortEventRelationship {
	if o == nil {
		var ret []StorageNetAppEthernetPortEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppEthernetPort) GetEventsOk() ([]StorageNetAppEthernetPortEventRelationship, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppEthernetPortEventRelationship and assigns it to the Events field.
func (o *StorageNetAppEthernetPort) SetEvents(v []StorageNetAppEthernetPortEventRelationship) {
	o.Events = v
}

func (o StorageNetAppEthernetPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppEthernetPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.BroadcastDomainName) {
		toSerialize["BroadcastDomainName"] = o.BroadcastDomainName
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.MacAddress) {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if !IsNil(o.Mtu) {
		toSerialize["Mtu"] = o.Mtu
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.NetAppEthernetPortLag.IsSet() {
		toSerialize["NetAppEthernetPortLag"] = o.NetAppEthernetPortLag.Get()
	}
	if o.NetAppEthernetPortVlan.IsSet() {
		toSerialize["NetAppEthernetPortVlan"] = o.NetAppEthernetPortVlan.Get()
	}
	if !IsNil(o.NodeName) {
		toSerialize["NodeName"] = o.NodeName
	}
	if !IsNil(o.PortState) {
		toSerialize["PortState"] = o.PortState
	}
	if !IsNil(o.Speed) {
		toSerialize["Speed"] = o.Speed
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ArrayController.IsSet() {
		toSerialize["ArrayController"] = o.ArrayController.Get()
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppEthernetPort) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppEthernetPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the broadcast domain, scoped to its IPspace.
		BroadcastDomainName *string `json:"BroadcastDomainName,omitempty"`
		// Status of port to determine if its enabled or not.
		Enabled *string `json:"Enabled,omitempty"`
		// MAC address of the port available in storage array.
		MacAddress *string `json:"MacAddress,omitempty" validate:"regexp=^$|^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"`
		// Maximum transmission unit of the physical port available in storage array.
		Mtu *int64 `json:"Mtu,omitempty"`
		// Name of the port available in storage array.
		Name                   *string                               `json:"Name,omitempty"`
		NetAppEthernetPortLag  NullableStorageNetAppEthernetPortLag  `json:"NetAppEthernetPortLag,omitempty"`
		NetAppEthernetPortVlan NullableStorageNetAppEthernetPortVlan `json:"NetAppEthernetPortVlan,omitempty"`
		// The node name for the port.
		NodeName *string `json:"NodeName,omitempty"`
		// State of the port available in storage array. * `Down` - An inactive port is listed as Down. * `Up` - An active port is listed as Up. * `Degraded` - An active port that is Up but unhealthy.
		PortState *string `json:"PortState,omitempty"`
		// Operational speed of port measured.
		Speed *int64 `json:"Speed,omitempty"`
		// State of the port available in storage array. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
		// Deprecated
		State *string `json:"State,omitempty"`
		// Type of the port available in storage array. * `LAG` - Storage port of type lag. * `physical` - LIFs can be configured directly on physical ports. * `VLAN` - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port.
		Type *string `json:"Type,omitempty"`
		// Universally unique identifier of the physical port.
		Uuid            *string                               `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		ArrayController NullableStorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
		// An array of relationships to storageNetAppEthernetPortEvent resources.
		Events []StorageNetAppEthernetPortEventRelationship `json:"Events,omitempty"`
	}

	varStorageNetAppEthernetPortWithoutEmbeddedStruct := StorageNetAppEthernetPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppEthernetPortWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppEthernetPort := _StorageNetAppEthernetPort{}
		varStorageNetAppEthernetPort.ClassId = varStorageNetAppEthernetPortWithoutEmbeddedStruct.ClassId
		varStorageNetAppEthernetPort.ObjectType = varStorageNetAppEthernetPortWithoutEmbeddedStruct.ObjectType
		varStorageNetAppEthernetPort.BroadcastDomainName = varStorageNetAppEthernetPortWithoutEmbeddedStruct.BroadcastDomainName
		varStorageNetAppEthernetPort.Enabled = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Enabled
		varStorageNetAppEthernetPort.MacAddress = varStorageNetAppEthernetPortWithoutEmbeddedStruct.MacAddress
		varStorageNetAppEthernetPort.Mtu = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Mtu
		varStorageNetAppEthernetPort.Name = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Name
		varStorageNetAppEthernetPort.NetAppEthernetPortLag = varStorageNetAppEthernetPortWithoutEmbeddedStruct.NetAppEthernetPortLag
		varStorageNetAppEthernetPort.NetAppEthernetPortVlan = varStorageNetAppEthernetPortWithoutEmbeddedStruct.NetAppEthernetPortVlan
		varStorageNetAppEthernetPort.NodeName = varStorageNetAppEthernetPortWithoutEmbeddedStruct.NodeName
		varStorageNetAppEthernetPort.PortState = varStorageNetAppEthernetPortWithoutEmbeddedStruct.PortState
		varStorageNetAppEthernetPort.Speed = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Speed
		varStorageNetAppEthernetPort.State = varStorageNetAppEthernetPortWithoutEmbeddedStruct.State
		varStorageNetAppEthernetPort.Type = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Type
		varStorageNetAppEthernetPort.Uuid = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Uuid
		varStorageNetAppEthernetPort.ArrayController = varStorageNetAppEthernetPortWithoutEmbeddedStruct.ArrayController
		varStorageNetAppEthernetPort.Events = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Events
		*o = StorageNetAppEthernetPort(varStorageNetAppEthernetPort)
	} else {
		return err
	}

	varStorageNetAppEthernetPort := _StorageNetAppEthernetPort{}

	err = json.Unmarshal(data, &varStorageNetAppEthernetPort)
	if err == nil {
		o.MoBaseMo = varStorageNetAppEthernetPort.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BroadcastDomainName")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NetAppEthernetPortLag")
		delete(additionalProperties, "NetAppEthernetPortVlan")
		delete(additionalProperties, "NodeName")
		delete(additionalProperties, "PortState")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")

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

type NullableStorageNetAppEthernetPort struct {
	value *StorageNetAppEthernetPort
	isSet bool
}

func (v NullableStorageNetAppEthernetPort) Get() *StorageNetAppEthernetPort {
	return v.value
}

func (v *NullableStorageNetAppEthernetPort) Set(val *StorageNetAppEthernetPort) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppEthernetPort) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppEthernetPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppEthernetPort(val *StorageNetAppEthernetPort) *NullableStorageNetAppEthernetPort {
	return &NullableStorageNetAppEthernetPort{value: val, isSet: true}
}

func (v NullableStorageNetAppEthernetPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppEthernetPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
