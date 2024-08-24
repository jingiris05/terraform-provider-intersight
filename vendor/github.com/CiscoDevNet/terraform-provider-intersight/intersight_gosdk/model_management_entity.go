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

// checks if the ManagementEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementEntity{}

// ManagementEntity Logical representation that captures the role of each Fabric Interconnect in UCS Manager.
type ManagementEntity struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Cluster link state between the Fabric Interconnects.
	ClusterLinkState *string `json:"ClusterLinkState,omitempty"`
	// Cluster readiness of the Fabric Interconnect.
	ClusterReadiness *string `json:"ClusterReadiness,omitempty"`
	// Cluster state of the Fabric Interconnect.
	ClusterState *string `json:"ClusterState,omitempty"`
	// Identity of the Fabric Interconnect - A/B.
	EntityId *string `json:"EntityId,omitempty"`
	// Role (Primary / Subordinate) of the Fabric Interconnect.
	Leadership           *string                                     `json:"Leadership,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement       NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagementEntity ManagementEntity

// NewManagementEntity instantiates a new ManagementEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementEntity(classId string, objectType string) *ManagementEntity {
	this := ManagementEntity{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewManagementEntityWithDefaults instantiates a new ManagementEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementEntityWithDefaults() *ManagementEntity {
	this := ManagementEntity{}
	var classId string = "management.Entity"
	this.ClassId = classId
	var objectType string = "management.Entity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ManagementEntity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ManagementEntity) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "management.Entity" of the ClassId field.
func (o *ManagementEntity) GetDefaultClassId() interface{} {
	return "management.Entity"
}

// GetObjectType returns the ObjectType field value
func (o *ManagementEntity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ManagementEntity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "management.Entity" of the ObjectType field.
func (o *ManagementEntity) GetDefaultObjectType() interface{} {
	return "management.Entity"
}

// GetClusterLinkState returns the ClusterLinkState field value if set, zero value otherwise.
func (o *ManagementEntity) GetClusterLinkState() string {
	if o == nil || IsNil(o.ClusterLinkState) {
		var ret string
		return ret
	}
	return *o.ClusterLinkState
}

// GetClusterLinkStateOk returns a tuple with the ClusterLinkState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetClusterLinkStateOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterLinkState) {
		return nil, false
	}
	return o.ClusterLinkState, true
}

// HasClusterLinkState returns a boolean if a field has been set.
func (o *ManagementEntity) HasClusterLinkState() bool {
	if o != nil && !IsNil(o.ClusterLinkState) {
		return true
	}

	return false
}

// SetClusterLinkState gets a reference to the given string and assigns it to the ClusterLinkState field.
func (o *ManagementEntity) SetClusterLinkState(v string) {
	o.ClusterLinkState = &v
}

// GetClusterReadiness returns the ClusterReadiness field value if set, zero value otherwise.
func (o *ManagementEntity) GetClusterReadiness() string {
	if o == nil || IsNil(o.ClusterReadiness) {
		var ret string
		return ret
	}
	return *o.ClusterReadiness
}

// GetClusterReadinessOk returns a tuple with the ClusterReadiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetClusterReadinessOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterReadiness) {
		return nil, false
	}
	return o.ClusterReadiness, true
}

// HasClusterReadiness returns a boolean if a field has been set.
func (o *ManagementEntity) HasClusterReadiness() bool {
	if o != nil && !IsNil(o.ClusterReadiness) {
		return true
	}

	return false
}

// SetClusterReadiness gets a reference to the given string and assigns it to the ClusterReadiness field.
func (o *ManagementEntity) SetClusterReadiness(v string) {
	o.ClusterReadiness = &v
}

// GetClusterState returns the ClusterState field value if set, zero value otherwise.
func (o *ManagementEntity) GetClusterState() string {
	if o == nil || IsNil(o.ClusterState) {
		var ret string
		return ret
	}
	return *o.ClusterState
}

// GetClusterStateOk returns a tuple with the ClusterState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetClusterStateOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterState) {
		return nil, false
	}
	return o.ClusterState, true
}

// HasClusterState returns a boolean if a field has been set.
func (o *ManagementEntity) HasClusterState() bool {
	if o != nil && !IsNil(o.ClusterState) {
		return true
	}

	return false
}

// SetClusterState gets a reference to the given string and assigns it to the ClusterState field.
func (o *ManagementEntity) SetClusterState(v string) {
	o.ClusterState = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *ManagementEntity) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *ManagementEntity) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *ManagementEntity) SetEntityId(v string) {
	o.EntityId = &v
}

// GetLeadership returns the Leadership field value if set, zero value otherwise.
func (o *ManagementEntity) GetLeadership() string {
	if o == nil || IsNil(o.Leadership) {
		var ret string
		return ret
	}
	return *o.Leadership
}

// GetLeadershipOk returns a tuple with the Leadership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementEntity) GetLeadershipOk() (*string, bool) {
	if o == nil || IsNil(o.Leadership) {
		return nil, false
	}
	return o.Leadership, true
}

// HasLeadership returns a boolean if a field has been set.
func (o *ManagementEntity) HasLeadership() bool {
	if o != nil && !IsNil(o.Leadership) {
		return true
	}

	return false
}

// SetLeadership gets a reference to the given string and assigns it to the Leadership field.
func (o *ManagementEntity) SetLeadership(v string) {
	o.Leadership = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagementEntity) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagementEntity) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ManagementEntity) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ManagementEntity) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *ManagementEntity) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *ManagementEntity) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagementEntity) GetNetworkElement() NetworkElementRelationship {
	if o == nil || IsNil(o.NetworkElement.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement.Get()
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagementEntity) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkElement.Get(), o.NetworkElement.IsSet()
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *ManagementEntity) HasNetworkElement() bool {
	if o != nil && o.NetworkElement.IsSet() {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NullableNetworkElementRelationship and assigns it to the NetworkElement field.
func (o *ManagementEntity) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement.Set(&v)
}

// SetNetworkElementNil sets the value for NetworkElement to be an explicit nil
func (o *ManagementEntity) SetNetworkElementNil() {
	o.NetworkElement.Set(nil)
}

// UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
func (o *ManagementEntity) UnsetNetworkElement() {
	o.NetworkElement.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagementEntity) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagementEntity) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ManagementEntity) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ManagementEntity) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *ManagementEntity) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *ManagementEntity) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o ManagementEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementEntity) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ClusterLinkState) {
		toSerialize["ClusterLinkState"] = o.ClusterLinkState
	}
	if !IsNil(o.ClusterReadiness) {
		toSerialize["ClusterReadiness"] = o.ClusterReadiness
	}
	if !IsNil(o.ClusterState) {
		toSerialize["ClusterState"] = o.ClusterState
	}
	if !IsNil(o.EntityId) {
		toSerialize["EntityId"] = o.EntityId
	}
	if !IsNil(o.Leadership) {
		toSerialize["Leadership"] = o.Leadership
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.NetworkElement.IsSet() {
		toSerialize["NetworkElement"] = o.NetworkElement.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagementEntity) UnmarshalJSON(data []byte) (err error) {
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
	type ManagementEntityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Cluster link state between the Fabric Interconnects.
		ClusterLinkState *string `json:"ClusterLinkState,omitempty"`
		// Cluster readiness of the Fabric Interconnect.
		ClusterReadiness *string `json:"ClusterReadiness,omitempty"`
		// Cluster state of the Fabric Interconnect.
		ClusterState *string `json:"ClusterState,omitempty"`
		// Identity of the Fabric Interconnect - A/B.
		EntityId *string `json:"EntityId,omitempty"`
		// Role (Primary / Subordinate) of the Fabric Interconnect.
		Leadership          *string                                     `json:"Leadership,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		NetworkElement      NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varManagementEntityWithoutEmbeddedStruct := ManagementEntityWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varManagementEntityWithoutEmbeddedStruct)
	if err == nil {
		varManagementEntity := _ManagementEntity{}
		varManagementEntity.ClassId = varManagementEntityWithoutEmbeddedStruct.ClassId
		varManagementEntity.ObjectType = varManagementEntityWithoutEmbeddedStruct.ObjectType
		varManagementEntity.ClusterLinkState = varManagementEntityWithoutEmbeddedStruct.ClusterLinkState
		varManagementEntity.ClusterReadiness = varManagementEntityWithoutEmbeddedStruct.ClusterReadiness
		varManagementEntity.ClusterState = varManagementEntityWithoutEmbeddedStruct.ClusterState
		varManagementEntity.EntityId = varManagementEntityWithoutEmbeddedStruct.EntityId
		varManagementEntity.Leadership = varManagementEntityWithoutEmbeddedStruct.Leadership
		varManagementEntity.InventoryDeviceInfo = varManagementEntityWithoutEmbeddedStruct.InventoryDeviceInfo
		varManagementEntity.NetworkElement = varManagementEntityWithoutEmbeddedStruct.NetworkElement
		varManagementEntity.RegisteredDevice = varManagementEntityWithoutEmbeddedStruct.RegisteredDevice
		*o = ManagementEntity(varManagementEntity)
	} else {
		return err
	}

	varManagementEntity := _ManagementEntity{}

	err = json.Unmarshal(data, &varManagementEntity)
	if err == nil {
		o.InventoryBase = varManagementEntity.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterLinkState")
		delete(additionalProperties, "ClusterReadiness")
		delete(additionalProperties, "ClusterState")
		delete(additionalProperties, "EntityId")
		delete(additionalProperties, "Leadership")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableManagementEntity struct {
	value *ManagementEntity
	isSet bool
}

func (v NullableManagementEntity) Get() *ManagementEntity {
	return v.value
}

func (v *NullableManagementEntity) Set(val *ManagementEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementEntity(val *ManagementEntity) *NullableManagementEntity {
	return &NullableManagementEntity{value: val, isSet: true}
}

func (v NullableManagementEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
