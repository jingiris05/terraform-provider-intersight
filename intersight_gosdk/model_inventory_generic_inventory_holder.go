/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-24T21:18:00Z.
 *
 * API version: 1.0.9-4404
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// InventoryGenericInventoryHolder A container class for generic inventory.
type InventoryGenericInventoryHolder struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The endpoint represented by this holder.
	Endpoint        *string                      `json:"Endpoint,omitempty"`
	ComputeBlade    *ComputeBladeRelationship    `json:"ComputeBlade,omitempty"`
	ComputeRackUnit *ComputeRackUnitRelationship `json:"ComputeRackUnit,omitempty"`
	// An array of relationships to inventoryGenericInventory resources.
	GenericInventory     []InventoryGenericInventoryRelationship `json:"GenericInventory,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryGenericInventoryHolder InventoryGenericInventoryHolder

// NewInventoryGenericInventoryHolder instantiates a new InventoryGenericInventoryHolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryGenericInventoryHolder(classId string, objectType string) *InventoryGenericInventoryHolder {
	this := InventoryGenericInventoryHolder{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryGenericInventoryHolderWithDefaults instantiates a new InventoryGenericInventoryHolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryGenericInventoryHolderWithDefaults() *InventoryGenericInventoryHolder {
	this := InventoryGenericInventoryHolder{}
	var classId string = "inventory.GenericInventoryHolder"
	this.ClassId = classId
	var objectType string = "inventory.GenericInventoryHolder"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryGenericInventoryHolder) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryGenericInventoryHolder) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryGenericInventoryHolder) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryGenericInventoryHolder) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolder) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *InventoryGenericInventoryHolder) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolder) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *InventoryGenericInventoryHolder) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolder) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *InventoryGenericInventoryHolder) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetGenericInventory returns the GenericInventory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryGenericInventoryHolder) GetGenericInventory() []InventoryGenericInventoryRelationship {
	if o == nil {
		var ret []InventoryGenericInventoryRelationship
		return ret
	}
	return o.GenericInventory
}

// GetGenericInventoryOk returns a tuple with the GenericInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryGenericInventoryHolder) GetGenericInventoryOk() (*[]InventoryGenericInventoryRelationship, bool) {
	if o == nil || o.GenericInventory == nil {
		return nil, false
	}
	return &o.GenericInventory, true
}

// HasGenericInventory returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasGenericInventory() bool {
	if o != nil && o.GenericInventory != nil {
		return true
	}

	return false
}

// SetGenericInventory gets a reference to the given []InventoryGenericInventoryRelationship and assigns it to the GenericInventory field.
func (o *InventoryGenericInventoryHolder) SetGenericInventory(v []InventoryGenericInventoryRelationship) {
	o.GenericInventory = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolder) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *InventoryGenericInventoryHolder) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolder) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *InventoryGenericInventoryHolder) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o InventoryGenericInventoryHolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Endpoint != nil {
		toSerialize["Endpoint"] = o.Endpoint
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.GenericInventory != nil {
		toSerialize["GenericInventory"] = o.GenericInventory
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

func (o *InventoryGenericInventoryHolder) UnmarshalJSON(bytes []byte) (err error) {
	type InventoryGenericInventoryHolderWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The endpoint represented by this holder.
		Endpoint        *string                      `json:"Endpoint,omitempty"`
		ComputeBlade    *ComputeBladeRelationship    `json:"ComputeBlade,omitempty"`
		ComputeRackUnit *ComputeRackUnitRelationship `json:"ComputeRackUnit,omitempty"`
		// An array of relationships to inventoryGenericInventory resources.
		GenericInventory    []InventoryGenericInventoryRelationship `json:"GenericInventory,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	}

	varInventoryGenericInventoryHolderWithoutEmbeddedStruct := InventoryGenericInventoryHolderWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInventoryGenericInventoryHolderWithoutEmbeddedStruct)
	if err == nil {
		varInventoryGenericInventoryHolder := _InventoryGenericInventoryHolder{}
		varInventoryGenericInventoryHolder.ClassId = varInventoryGenericInventoryHolderWithoutEmbeddedStruct.ClassId
		varInventoryGenericInventoryHolder.ObjectType = varInventoryGenericInventoryHolderWithoutEmbeddedStruct.ObjectType
		varInventoryGenericInventoryHolder.Endpoint = varInventoryGenericInventoryHolderWithoutEmbeddedStruct.Endpoint
		varInventoryGenericInventoryHolder.ComputeBlade = varInventoryGenericInventoryHolderWithoutEmbeddedStruct.ComputeBlade
		varInventoryGenericInventoryHolder.ComputeRackUnit = varInventoryGenericInventoryHolderWithoutEmbeddedStruct.ComputeRackUnit
		varInventoryGenericInventoryHolder.GenericInventory = varInventoryGenericInventoryHolderWithoutEmbeddedStruct.GenericInventory
		varInventoryGenericInventoryHolder.InventoryDeviceInfo = varInventoryGenericInventoryHolderWithoutEmbeddedStruct.InventoryDeviceInfo
		varInventoryGenericInventoryHolder.RegisteredDevice = varInventoryGenericInventoryHolderWithoutEmbeddedStruct.RegisteredDevice
		*o = InventoryGenericInventoryHolder(varInventoryGenericInventoryHolder)
	} else {
		return err
	}

	varInventoryGenericInventoryHolder := _InventoryGenericInventoryHolder{}

	err = json.Unmarshal(bytes, &varInventoryGenericInventoryHolder)
	if err == nil {
		o.InventoryBase = varInventoryGenericInventoryHolder.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Endpoint")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "GenericInventory")
		delete(additionalProperties, "InventoryDeviceInfo")
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

type NullableInventoryGenericInventoryHolder struct {
	value *InventoryGenericInventoryHolder
	isSet bool
}

func (v NullableInventoryGenericInventoryHolder) Get() *InventoryGenericInventoryHolder {
	return v.value
}

func (v *NullableInventoryGenericInventoryHolder) Set(val *InventoryGenericInventoryHolder) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryGenericInventoryHolder) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryGenericInventoryHolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryGenericInventoryHolder(val *InventoryGenericInventoryHolder) *NullableInventoryGenericInventoryHolder {
	return &NullableInventoryGenericInventoryHolder{value: val, isSet: true}
}

func (v NullableInventoryGenericInventoryHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryGenericInventoryHolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
