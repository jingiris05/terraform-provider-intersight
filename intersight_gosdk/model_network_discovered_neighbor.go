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

// checks if the NetworkDiscoveredNeighbor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDiscoveredNeighbor{}

// NetworkDiscoveredNeighbor L2 neighbor (LLDP and CDP) available on the switch.
type NetworkDiscoveredNeighbor struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Device identifier of the switch.
	DeviceId *string `json:"DeviceId,omitempty"`
	// The management address of the switch.
	ManagementAddress        *string  `json:"ManagementAddress,omitempty"`
	NeighborDeviceCapability []string `json:"NeighborDeviceCapability,omitempty"`
	// Device identifier of the neighboring device.
	NeighborDeviceId *string `json:"NeighborDeviceId,omitempty"`
	// Neighboring device interface.
	NeighborInterface *string `json:"NeighborInterface,omitempty"`
	// PortID of the neighbor device configured.
	NeighborPortId *string `json:"NeighborPortId,omitempty"`
	// Name of the local interface.
	SwitchPortId         *string                            `json:"SwitchPortId,omitempty"`
	CdpNeighbor          NullableNetworkElementRelationship `json:"CdpNeighbor,omitempty"`
	LldpNeighbor         NullableNetworkElementRelationship `json:"LldpNeighbor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkDiscoveredNeighbor NetworkDiscoveredNeighbor

// NewNetworkDiscoveredNeighbor instantiates a new NetworkDiscoveredNeighbor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDiscoveredNeighbor(classId string, objectType string) *NetworkDiscoveredNeighbor {
	this := NetworkDiscoveredNeighbor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkDiscoveredNeighborWithDefaults instantiates a new NetworkDiscoveredNeighbor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDiscoveredNeighborWithDefaults() *NetworkDiscoveredNeighbor {
	this := NetworkDiscoveredNeighbor{}
	var classId string = "network.DiscoveredNeighbor"
	this.ClassId = classId
	var objectType string = "network.DiscoveredNeighbor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkDiscoveredNeighbor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveredNeighbor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkDiscoveredNeighbor) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "network.DiscoveredNeighbor" of the ClassId field.
func (o *NetworkDiscoveredNeighbor) GetDefaultClassId() interface{} {
	return "network.DiscoveredNeighbor"
}

// GetObjectType returns the ObjectType field value
func (o *NetworkDiscoveredNeighbor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveredNeighbor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkDiscoveredNeighbor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "network.DiscoveredNeighbor" of the ObjectType field.
func (o *NetworkDiscoveredNeighbor) GetDefaultObjectType() interface{} {
	return "network.DiscoveredNeighbor"
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *NetworkDiscoveredNeighbor) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveredNeighbor) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *NetworkDiscoveredNeighbor) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *NetworkDiscoveredNeighbor) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *NetworkDiscoveredNeighbor) GetManagementAddress() string {
	if o == nil || IsNil(o.ManagementAddress) {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveredNeighbor) GetManagementAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementAddress) {
		return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *NetworkDiscoveredNeighbor) HasManagementAddress() bool {
	if o != nil && !IsNil(o.ManagementAddress) {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *NetworkDiscoveredNeighbor) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetNeighborDeviceCapability returns the NeighborDeviceCapability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkDiscoveredNeighbor) GetNeighborDeviceCapability() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NeighborDeviceCapability
}

// GetNeighborDeviceCapabilityOk returns a tuple with the NeighborDeviceCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkDiscoveredNeighbor) GetNeighborDeviceCapabilityOk() ([]string, bool) {
	if o == nil || IsNil(o.NeighborDeviceCapability) {
		return nil, false
	}
	return o.NeighborDeviceCapability, true
}

// HasNeighborDeviceCapability returns a boolean if a field has been set.
func (o *NetworkDiscoveredNeighbor) HasNeighborDeviceCapability() bool {
	if o != nil && !IsNil(o.NeighborDeviceCapability) {
		return true
	}

	return false
}

// SetNeighborDeviceCapability gets a reference to the given []string and assigns it to the NeighborDeviceCapability field.
func (o *NetworkDiscoveredNeighbor) SetNeighborDeviceCapability(v []string) {
	o.NeighborDeviceCapability = v
}

// GetNeighborDeviceId returns the NeighborDeviceId field value if set, zero value otherwise.
func (o *NetworkDiscoveredNeighbor) GetNeighborDeviceId() string {
	if o == nil || IsNil(o.NeighborDeviceId) {
		var ret string
		return ret
	}
	return *o.NeighborDeviceId
}

// GetNeighborDeviceIdOk returns a tuple with the NeighborDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveredNeighbor) GetNeighborDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.NeighborDeviceId) {
		return nil, false
	}
	return o.NeighborDeviceId, true
}

// HasNeighborDeviceId returns a boolean if a field has been set.
func (o *NetworkDiscoveredNeighbor) HasNeighborDeviceId() bool {
	if o != nil && !IsNil(o.NeighborDeviceId) {
		return true
	}

	return false
}

// SetNeighborDeviceId gets a reference to the given string and assigns it to the NeighborDeviceId field.
func (o *NetworkDiscoveredNeighbor) SetNeighborDeviceId(v string) {
	o.NeighborDeviceId = &v
}

// GetNeighborInterface returns the NeighborInterface field value if set, zero value otherwise.
func (o *NetworkDiscoveredNeighbor) GetNeighborInterface() string {
	if o == nil || IsNil(o.NeighborInterface) {
		var ret string
		return ret
	}
	return *o.NeighborInterface
}

// GetNeighborInterfaceOk returns a tuple with the NeighborInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveredNeighbor) GetNeighborInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.NeighborInterface) {
		return nil, false
	}
	return o.NeighborInterface, true
}

// HasNeighborInterface returns a boolean if a field has been set.
func (o *NetworkDiscoveredNeighbor) HasNeighborInterface() bool {
	if o != nil && !IsNil(o.NeighborInterface) {
		return true
	}

	return false
}

// SetNeighborInterface gets a reference to the given string and assigns it to the NeighborInterface field.
func (o *NetworkDiscoveredNeighbor) SetNeighborInterface(v string) {
	o.NeighborInterface = &v
}

// GetNeighborPortId returns the NeighborPortId field value if set, zero value otherwise.
func (o *NetworkDiscoveredNeighbor) GetNeighborPortId() string {
	if o == nil || IsNil(o.NeighborPortId) {
		var ret string
		return ret
	}
	return *o.NeighborPortId
}

// GetNeighborPortIdOk returns a tuple with the NeighborPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveredNeighbor) GetNeighborPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.NeighborPortId) {
		return nil, false
	}
	return o.NeighborPortId, true
}

// HasNeighborPortId returns a boolean if a field has been set.
func (o *NetworkDiscoveredNeighbor) HasNeighborPortId() bool {
	if o != nil && !IsNil(o.NeighborPortId) {
		return true
	}

	return false
}

// SetNeighborPortId gets a reference to the given string and assigns it to the NeighborPortId field.
func (o *NetworkDiscoveredNeighbor) SetNeighborPortId(v string) {
	o.NeighborPortId = &v
}

// GetSwitchPortId returns the SwitchPortId field value if set, zero value otherwise.
func (o *NetworkDiscoveredNeighbor) GetSwitchPortId() string {
	if o == nil || IsNil(o.SwitchPortId) {
		var ret string
		return ret
	}
	return *o.SwitchPortId
}

// GetSwitchPortIdOk returns a tuple with the SwitchPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveredNeighbor) GetSwitchPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchPortId) {
		return nil, false
	}
	return o.SwitchPortId, true
}

// HasSwitchPortId returns a boolean if a field has been set.
func (o *NetworkDiscoveredNeighbor) HasSwitchPortId() bool {
	if o != nil && !IsNil(o.SwitchPortId) {
		return true
	}

	return false
}

// SetSwitchPortId gets a reference to the given string and assigns it to the SwitchPortId field.
func (o *NetworkDiscoveredNeighbor) SetSwitchPortId(v string) {
	o.SwitchPortId = &v
}

// GetCdpNeighbor returns the CdpNeighbor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkDiscoveredNeighbor) GetCdpNeighbor() NetworkElementRelationship {
	if o == nil || IsNil(o.CdpNeighbor.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.CdpNeighbor.Get()
}

// GetCdpNeighborOk returns a tuple with the CdpNeighbor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkDiscoveredNeighbor) GetCdpNeighborOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.CdpNeighbor.Get(), o.CdpNeighbor.IsSet()
}

// HasCdpNeighbor returns a boolean if a field has been set.
func (o *NetworkDiscoveredNeighbor) HasCdpNeighbor() bool {
	if o != nil && o.CdpNeighbor.IsSet() {
		return true
	}

	return false
}

// SetCdpNeighbor gets a reference to the given NullableNetworkElementRelationship and assigns it to the CdpNeighbor field.
func (o *NetworkDiscoveredNeighbor) SetCdpNeighbor(v NetworkElementRelationship) {
	o.CdpNeighbor.Set(&v)
}

// SetCdpNeighborNil sets the value for CdpNeighbor to be an explicit nil
func (o *NetworkDiscoveredNeighbor) SetCdpNeighborNil() {
	o.CdpNeighbor.Set(nil)
}

// UnsetCdpNeighbor ensures that no value is present for CdpNeighbor, not even an explicit nil
func (o *NetworkDiscoveredNeighbor) UnsetCdpNeighbor() {
	o.CdpNeighbor.Unset()
}

// GetLldpNeighbor returns the LldpNeighbor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkDiscoveredNeighbor) GetLldpNeighbor() NetworkElementRelationship {
	if o == nil || IsNil(o.LldpNeighbor.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.LldpNeighbor.Get()
}

// GetLldpNeighborOk returns a tuple with the LldpNeighbor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkDiscoveredNeighbor) GetLldpNeighborOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.LldpNeighbor.Get(), o.LldpNeighbor.IsSet()
}

// HasLldpNeighbor returns a boolean if a field has been set.
func (o *NetworkDiscoveredNeighbor) HasLldpNeighbor() bool {
	if o != nil && o.LldpNeighbor.IsSet() {
		return true
	}

	return false
}

// SetLldpNeighbor gets a reference to the given NullableNetworkElementRelationship and assigns it to the LldpNeighbor field.
func (o *NetworkDiscoveredNeighbor) SetLldpNeighbor(v NetworkElementRelationship) {
	o.LldpNeighbor.Set(&v)
}

// SetLldpNeighborNil sets the value for LldpNeighbor to be an explicit nil
func (o *NetworkDiscoveredNeighbor) SetLldpNeighborNil() {
	o.LldpNeighbor.Set(nil)
}

// UnsetLldpNeighbor ensures that no value is present for LldpNeighbor, not even an explicit nil
func (o *NetworkDiscoveredNeighbor) UnsetLldpNeighbor() {
	o.LldpNeighbor.Unset()
}

func (o NetworkDiscoveredNeighbor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDiscoveredNeighbor) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DeviceId) {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if !IsNil(o.ManagementAddress) {
		toSerialize["ManagementAddress"] = o.ManagementAddress
	}
	if o.NeighborDeviceCapability != nil {
		toSerialize["NeighborDeviceCapability"] = o.NeighborDeviceCapability
	}
	if !IsNil(o.NeighborDeviceId) {
		toSerialize["NeighborDeviceId"] = o.NeighborDeviceId
	}
	if !IsNil(o.NeighborInterface) {
		toSerialize["NeighborInterface"] = o.NeighborInterface
	}
	if !IsNil(o.NeighborPortId) {
		toSerialize["NeighborPortId"] = o.NeighborPortId
	}
	if !IsNil(o.SwitchPortId) {
		toSerialize["SwitchPortId"] = o.SwitchPortId
	}
	if o.CdpNeighbor.IsSet() {
		toSerialize["CdpNeighbor"] = o.CdpNeighbor.Get()
	}
	if o.LldpNeighbor.IsSet() {
		toSerialize["LldpNeighbor"] = o.LldpNeighbor.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkDiscoveredNeighbor) UnmarshalJSON(data []byte) (err error) {
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
	type NetworkDiscoveredNeighborWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Device identifier of the switch.
		DeviceId *string `json:"DeviceId,omitempty"`
		// The management address of the switch.
		ManagementAddress        *string  `json:"ManagementAddress,omitempty"`
		NeighborDeviceCapability []string `json:"NeighborDeviceCapability,omitempty"`
		// Device identifier of the neighboring device.
		NeighborDeviceId *string `json:"NeighborDeviceId,omitempty"`
		// Neighboring device interface.
		NeighborInterface *string `json:"NeighborInterface,omitempty"`
		// PortID of the neighbor device configured.
		NeighborPortId *string `json:"NeighborPortId,omitempty"`
		// Name of the local interface.
		SwitchPortId *string                            `json:"SwitchPortId,omitempty"`
		CdpNeighbor  NullableNetworkElementRelationship `json:"CdpNeighbor,omitempty"`
		LldpNeighbor NullableNetworkElementRelationship `json:"LldpNeighbor,omitempty"`
	}

	varNetworkDiscoveredNeighborWithoutEmbeddedStruct := NetworkDiscoveredNeighborWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNetworkDiscoveredNeighborWithoutEmbeddedStruct)
	if err == nil {
		varNetworkDiscoveredNeighbor := _NetworkDiscoveredNeighbor{}
		varNetworkDiscoveredNeighbor.ClassId = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.ClassId
		varNetworkDiscoveredNeighbor.ObjectType = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.ObjectType
		varNetworkDiscoveredNeighbor.DeviceId = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.DeviceId
		varNetworkDiscoveredNeighbor.ManagementAddress = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.ManagementAddress
		varNetworkDiscoveredNeighbor.NeighborDeviceCapability = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.NeighborDeviceCapability
		varNetworkDiscoveredNeighbor.NeighborDeviceId = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.NeighborDeviceId
		varNetworkDiscoveredNeighbor.NeighborInterface = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.NeighborInterface
		varNetworkDiscoveredNeighbor.NeighborPortId = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.NeighborPortId
		varNetworkDiscoveredNeighbor.SwitchPortId = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.SwitchPortId
		varNetworkDiscoveredNeighbor.CdpNeighbor = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.CdpNeighbor
		varNetworkDiscoveredNeighbor.LldpNeighbor = varNetworkDiscoveredNeighborWithoutEmbeddedStruct.LldpNeighbor
		*o = NetworkDiscoveredNeighbor(varNetworkDiscoveredNeighbor)
	} else {
		return err
	}

	varNetworkDiscoveredNeighbor := _NetworkDiscoveredNeighbor{}

	err = json.Unmarshal(data, &varNetworkDiscoveredNeighbor)
	if err == nil {
		o.InventoryBase = varNetworkDiscoveredNeighbor.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "ManagementAddress")
		delete(additionalProperties, "NeighborDeviceCapability")
		delete(additionalProperties, "NeighborDeviceId")
		delete(additionalProperties, "NeighborInterface")
		delete(additionalProperties, "NeighborPortId")
		delete(additionalProperties, "SwitchPortId")
		delete(additionalProperties, "CdpNeighbor")
		delete(additionalProperties, "LldpNeighbor")

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

type NullableNetworkDiscoveredNeighbor struct {
	value *NetworkDiscoveredNeighbor
	isSet bool
}

func (v NullableNetworkDiscoveredNeighbor) Get() *NetworkDiscoveredNeighbor {
	return v.value
}

func (v *NullableNetworkDiscoveredNeighbor) Set(val *NetworkDiscoveredNeighbor) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDiscoveredNeighbor) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDiscoveredNeighbor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDiscoveredNeighbor(val *NetworkDiscoveredNeighbor) *NullableNetworkDiscoveredNeighbor {
	return &NullableNetworkDiscoveredNeighbor{value: val, isSet: true}
}

func (v NullableNetworkDiscoveredNeighbor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDiscoveredNeighbor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
