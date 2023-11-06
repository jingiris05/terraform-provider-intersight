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
)

// LsServiceProfileAllOf Definition of the list of properties defined in 'ls.ServiceProfile', excluding properties defined in parent classes.
type LsServiceProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Assignment state of the service profile.
	AssignState *string `json:"AssignState,omitempty"`
	// Association state of the service profile.
	AssocState *string `json:"AssocState,omitempty"`
	// Server to which the UCS Manager service profile is associated to.
	AssociatedServer *string `json:"AssociatedServer,omitempty"`
	// Configuration state of the service profile.
	ConfigState *string `json:"ConfigState,omitempty"`
	// Name of the UCS Manager service profile.
	Name *string `json:"Name,omitempty"`
	// Operational state of the service profile.
	OperState            *string                              `json:"OperState,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LsServiceProfileAllOf LsServiceProfileAllOf

// NewLsServiceProfileAllOf instantiates a new LsServiceProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLsServiceProfileAllOf(classId string, objectType string) *LsServiceProfileAllOf {
	this := LsServiceProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewLsServiceProfileAllOfWithDefaults instantiates a new LsServiceProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLsServiceProfileAllOfWithDefaults() *LsServiceProfileAllOf {
	this := LsServiceProfileAllOf{}
	var classId string = "ls.ServiceProfile"
	this.ClassId = classId
	var objectType string = "ls.ServiceProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LsServiceProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LsServiceProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LsServiceProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LsServiceProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LsServiceProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LsServiceProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssignState returns the AssignState field value if set, zero value otherwise.
func (o *LsServiceProfileAllOf) GetAssignState() string {
	if o == nil || o.AssignState == nil {
		var ret string
		return ret
	}
	return *o.AssignState
}

// GetAssignStateOk returns a tuple with the AssignState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LsServiceProfileAllOf) GetAssignStateOk() (*string, bool) {
	if o == nil || o.AssignState == nil {
		return nil, false
	}
	return o.AssignState, true
}

// HasAssignState returns a boolean if a field has been set.
func (o *LsServiceProfileAllOf) HasAssignState() bool {
	if o != nil && o.AssignState != nil {
		return true
	}

	return false
}

// SetAssignState gets a reference to the given string and assigns it to the AssignState field.
func (o *LsServiceProfileAllOf) SetAssignState(v string) {
	o.AssignState = &v
}

// GetAssocState returns the AssocState field value if set, zero value otherwise.
func (o *LsServiceProfileAllOf) GetAssocState() string {
	if o == nil || o.AssocState == nil {
		var ret string
		return ret
	}
	return *o.AssocState
}

// GetAssocStateOk returns a tuple with the AssocState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LsServiceProfileAllOf) GetAssocStateOk() (*string, bool) {
	if o == nil || o.AssocState == nil {
		return nil, false
	}
	return o.AssocState, true
}

// HasAssocState returns a boolean if a field has been set.
func (o *LsServiceProfileAllOf) HasAssocState() bool {
	if o != nil && o.AssocState != nil {
		return true
	}

	return false
}

// SetAssocState gets a reference to the given string and assigns it to the AssocState field.
func (o *LsServiceProfileAllOf) SetAssocState(v string) {
	o.AssocState = &v
}

// GetAssociatedServer returns the AssociatedServer field value if set, zero value otherwise.
func (o *LsServiceProfileAllOf) GetAssociatedServer() string {
	if o == nil || o.AssociatedServer == nil {
		var ret string
		return ret
	}
	return *o.AssociatedServer
}

// GetAssociatedServerOk returns a tuple with the AssociatedServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LsServiceProfileAllOf) GetAssociatedServerOk() (*string, bool) {
	if o == nil || o.AssociatedServer == nil {
		return nil, false
	}
	return o.AssociatedServer, true
}

// HasAssociatedServer returns a boolean if a field has been set.
func (o *LsServiceProfileAllOf) HasAssociatedServer() bool {
	if o != nil && o.AssociatedServer != nil {
		return true
	}

	return false
}

// SetAssociatedServer gets a reference to the given string and assigns it to the AssociatedServer field.
func (o *LsServiceProfileAllOf) SetAssociatedServer(v string) {
	o.AssociatedServer = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *LsServiceProfileAllOf) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LsServiceProfileAllOf) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *LsServiceProfileAllOf) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *LsServiceProfileAllOf) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LsServiceProfileAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LsServiceProfileAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LsServiceProfileAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LsServiceProfileAllOf) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *LsServiceProfileAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LsServiceProfileAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *LsServiceProfileAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *LsServiceProfileAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *LsServiceProfileAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LsServiceProfileAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *LsServiceProfileAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *LsServiceProfileAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *LsServiceProfileAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LsServiceProfileAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *LsServiceProfileAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *LsServiceProfileAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o LsServiceProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AssignState != nil {
		toSerialize["AssignState"] = o.AssignState
	}
	if o.AssocState != nil {
		toSerialize["AssocState"] = o.AssocState
	}
	if o.AssociatedServer != nil {
		toSerialize["AssociatedServer"] = o.AssociatedServer
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LsServiceProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLsServiceProfileAllOf := _LsServiceProfileAllOf{}

	if err = json.Unmarshal(bytes, &varLsServiceProfileAllOf); err == nil {
		*o = LsServiceProfileAllOf(varLsServiceProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssignState")
		delete(additionalProperties, "AssocState")
		delete(additionalProperties, "AssociatedServer")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLsServiceProfileAllOf struct {
	value *LsServiceProfileAllOf
	isSet bool
}

func (v NullableLsServiceProfileAllOf) Get() *LsServiceProfileAllOf {
	return v.value
}

func (v *NullableLsServiceProfileAllOf) Set(val *LsServiceProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLsServiceProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLsServiceProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLsServiceProfileAllOf(val *LsServiceProfileAllOf) *NullableLsServiceProfileAllOf {
	return &NullableLsServiceProfileAllOf{value: val, isSet: true}
}

func (v NullableLsServiceProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLsServiceProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
