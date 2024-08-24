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

// checks if the FabricVlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricVlan{}

// FabricVlan Configuration object for Virtual LAN.
type FabricVlan struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable to automatically allow this VLAN on all uplinks. Disable must be specified for Disjoint Layer 2 VLAN configuration. Default VLAN-1 cannot be configured as Disjoint Layer 2 VLAN.
	AutoAllowOnUplinks *bool `json:"AutoAllowOnUplinks,omitempty"`
	// Used to define whether this VLAN is to be classified as 'native' for traffic in this FI.
	IsNative *bool `json:"IsNative,omitempty"`
	// The 'name' used to identify this VLAN.
	Name *string `json:"Name,omitempty"`
	// The Primary VLAN ID of the VLAN, if the sharing type of the VLAN is Isolated or Community.
	PrimaryVlanId *int64 `json:"PrimaryVlanId,omitempty"`
	// The sharing type of this VLAN. * `None` - This represents a regular VLAN. * `Primary` - This represents a primary VLAN. * `Isolated` - This represents an isolated VLAN. * `Community` - This represents a community VLAN.
	SharingType *string `json:"SharingType,omitempty"`
	// The identifier for this Virtual LAN.
	VlanId               *int64                                     `json:"VlanId,omitempty"`
	EthNetworkPolicy     NullableFabricEthNetworkPolicyRelationship `json:"EthNetworkPolicy,omitempty"`
	MulticastPolicy      NullableFabricMulticastPolicyRelationship  `json:"MulticastPolicy,omitempty"`
	VlanSet              NullableFabricVlanSetRelationship          `json:"VlanSet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricVlan FabricVlan

// NewFabricVlan instantiates a new FabricVlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricVlan(classId string, objectType string) *FabricVlan {
	this := FabricVlan{}
	this.ClassId = classId
	this.ObjectType = objectType
	var autoAllowOnUplinks bool = true
	this.AutoAllowOnUplinks = &autoAllowOnUplinks
	var primaryVlanId int64 = 0
	this.PrimaryVlanId = &primaryVlanId
	var sharingType string = "None"
	this.SharingType = &sharingType
	return &this
}

// NewFabricVlanWithDefaults instantiates a new FabricVlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricVlanWithDefaults() *FabricVlan {
	this := FabricVlan{}
	var classId string = "fabric.Vlan"
	this.ClassId = classId
	var objectType string = "fabric.Vlan"
	this.ObjectType = objectType
	var autoAllowOnUplinks bool = true
	this.AutoAllowOnUplinks = &autoAllowOnUplinks
	var primaryVlanId int64 = 0
	this.PrimaryVlanId = &primaryVlanId
	var sharingType string = "None"
	this.SharingType = &sharingType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricVlan) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricVlan) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.Vlan" of the ClassId field.
func (o *FabricVlan) GetDefaultClassId() interface{} {
	return "fabric.Vlan"
}

// GetObjectType returns the ObjectType field value
func (o *FabricVlan) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricVlan) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.Vlan" of the ObjectType field.
func (o *FabricVlan) GetDefaultObjectType() interface{} {
	return "fabric.Vlan"
}

// GetAutoAllowOnUplinks returns the AutoAllowOnUplinks field value if set, zero value otherwise.
func (o *FabricVlan) GetAutoAllowOnUplinks() bool {
	if o == nil || IsNil(o.AutoAllowOnUplinks) {
		var ret bool
		return ret
	}
	return *o.AutoAllowOnUplinks
}

// GetAutoAllowOnUplinksOk returns a tuple with the AutoAllowOnUplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetAutoAllowOnUplinksOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoAllowOnUplinks) {
		return nil, false
	}
	return o.AutoAllowOnUplinks, true
}

// HasAutoAllowOnUplinks returns a boolean if a field has been set.
func (o *FabricVlan) HasAutoAllowOnUplinks() bool {
	if o != nil && !IsNil(o.AutoAllowOnUplinks) {
		return true
	}

	return false
}

// SetAutoAllowOnUplinks gets a reference to the given bool and assigns it to the AutoAllowOnUplinks field.
func (o *FabricVlan) SetAutoAllowOnUplinks(v bool) {
	o.AutoAllowOnUplinks = &v
}

// GetIsNative returns the IsNative field value if set, zero value otherwise.
func (o *FabricVlan) GetIsNative() bool {
	if o == nil || IsNil(o.IsNative) {
		var ret bool
		return ret
	}
	return *o.IsNative
}

// GetIsNativeOk returns a tuple with the IsNative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetIsNativeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNative) {
		return nil, false
	}
	return o.IsNative, true
}

// HasIsNative returns a boolean if a field has been set.
func (o *FabricVlan) HasIsNative() bool {
	if o != nil && !IsNil(o.IsNative) {
		return true
	}

	return false
}

// SetIsNative gets a reference to the given bool and assigns it to the IsNative field.
func (o *FabricVlan) SetIsNative(v bool) {
	o.IsNative = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricVlan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricVlan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricVlan) SetName(v string) {
	o.Name = &v
}

// GetPrimaryVlanId returns the PrimaryVlanId field value if set, zero value otherwise.
func (o *FabricVlan) GetPrimaryVlanId() int64 {
	if o == nil || IsNil(o.PrimaryVlanId) {
		var ret int64
		return ret
	}
	return *o.PrimaryVlanId
}

// GetPrimaryVlanIdOk returns a tuple with the PrimaryVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetPrimaryVlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PrimaryVlanId) {
		return nil, false
	}
	return o.PrimaryVlanId, true
}

// HasPrimaryVlanId returns a boolean if a field has been set.
func (o *FabricVlan) HasPrimaryVlanId() bool {
	if o != nil && !IsNil(o.PrimaryVlanId) {
		return true
	}

	return false
}

// SetPrimaryVlanId gets a reference to the given int64 and assigns it to the PrimaryVlanId field.
func (o *FabricVlan) SetPrimaryVlanId(v int64) {
	o.PrimaryVlanId = &v
}

// GetSharingType returns the SharingType field value if set, zero value otherwise.
func (o *FabricVlan) GetSharingType() string {
	if o == nil || IsNil(o.SharingType) {
		var ret string
		return ret
	}
	return *o.SharingType
}

// GetSharingTypeOk returns a tuple with the SharingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetSharingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SharingType) {
		return nil, false
	}
	return o.SharingType, true
}

// HasSharingType returns a boolean if a field has been set.
func (o *FabricVlan) HasSharingType() bool {
	if o != nil && !IsNil(o.SharingType) {
		return true
	}

	return false
}

// SetSharingType gets a reference to the given string and assigns it to the SharingType field.
func (o *FabricVlan) SetSharingType(v string) {
	o.SharingType = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *FabricVlan) GetVlanId() int64 {
	if o == nil || IsNil(o.VlanId) {
		var ret int64
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlan) GetVlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *FabricVlan) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int64 and assigns it to the VlanId field.
func (o *FabricVlan) SetVlanId(v int64) {
	o.VlanId = &v
}

// GetEthNetworkPolicy returns the EthNetworkPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricVlan) GetEthNetworkPolicy() FabricEthNetworkPolicyRelationship {
	if o == nil || IsNil(o.EthNetworkPolicy.Get()) {
		var ret FabricEthNetworkPolicyRelationship
		return ret
	}
	return *o.EthNetworkPolicy.Get()
}

// GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricVlan) GetEthNetworkPolicyOk() (*FabricEthNetworkPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.EthNetworkPolicy.Get(), o.EthNetworkPolicy.IsSet()
}

// HasEthNetworkPolicy returns a boolean if a field has been set.
func (o *FabricVlan) HasEthNetworkPolicy() bool {
	if o != nil && o.EthNetworkPolicy.IsSet() {
		return true
	}

	return false
}

// SetEthNetworkPolicy gets a reference to the given NullableFabricEthNetworkPolicyRelationship and assigns it to the EthNetworkPolicy field.
func (o *FabricVlan) SetEthNetworkPolicy(v FabricEthNetworkPolicyRelationship) {
	o.EthNetworkPolicy.Set(&v)
}

// SetEthNetworkPolicyNil sets the value for EthNetworkPolicy to be an explicit nil
func (o *FabricVlan) SetEthNetworkPolicyNil() {
	o.EthNetworkPolicy.Set(nil)
}

// UnsetEthNetworkPolicy ensures that no value is present for EthNetworkPolicy, not even an explicit nil
func (o *FabricVlan) UnsetEthNetworkPolicy() {
	o.EthNetworkPolicy.Unset()
}

// GetMulticastPolicy returns the MulticastPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricVlan) GetMulticastPolicy() FabricMulticastPolicyRelationship {
	if o == nil || IsNil(o.MulticastPolicy.Get()) {
		var ret FabricMulticastPolicyRelationship
		return ret
	}
	return *o.MulticastPolicy.Get()
}

// GetMulticastPolicyOk returns a tuple with the MulticastPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricVlan) GetMulticastPolicyOk() (*FabricMulticastPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.MulticastPolicy.Get(), o.MulticastPolicy.IsSet()
}

// HasMulticastPolicy returns a boolean if a field has been set.
func (o *FabricVlan) HasMulticastPolicy() bool {
	if o != nil && o.MulticastPolicy.IsSet() {
		return true
	}

	return false
}

// SetMulticastPolicy gets a reference to the given NullableFabricMulticastPolicyRelationship and assigns it to the MulticastPolicy field.
func (o *FabricVlan) SetMulticastPolicy(v FabricMulticastPolicyRelationship) {
	o.MulticastPolicy.Set(&v)
}

// SetMulticastPolicyNil sets the value for MulticastPolicy to be an explicit nil
func (o *FabricVlan) SetMulticastPolicyNil() {
	o.MulticastPolicy.Set(nil)
}

// UnsetMulticastPolicy ensures that no value is present for MulticastPolicy, not even an explicit nil
func (o *FabricVlan) UnsetMulticastPolicy() {
	o.MulticastPolicy.Unset()
}

// GetVlanSet returns the VlanSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricVlan) GetVlanSet() FabricVlanSetRelationship {
	if o == nil || IsNil(o.VlanSet.Get()) {
		var ret FabricVlanSetRelationship
		return ret
	}
	return *o.VlanSet.Get()
}

// GetVlanSetOk returns a tuple with the VlanSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricVlan) GetVlanSetOk() (*FabricVlanSetRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanSet.Get(), o.VlanSet.IsSet()
}

// HasVlanSet returns a boolean if a field has been set.
func (o *FabricVlan) HasVlanSet() bool {
	if o != nil && o.VlanSet.IsSet() {
		return true
	}

	return false
}

// SetVlanSet gets a reference to the given NullableFabricVlanSetRelationship and assigns it to the VlanSet field.
func (o *FabricVlan) SetVlanSet(v FabricVlanSetRelationship) {
	o.VlanSet.Set(&v)
}

// SetVlanSetNil sets the value for VlanSet to be an explicit nil
func (o *FabricVlan) SetVlanSetNil() {
	o.VlanSet.Set(nil)
}

// UnsetVlanSet ensures that no value is present for VlanSet, not even an explicit nil
func (o *FabricVlan) UnsetVlanSet() {
	o.VlanSet.Unset()
}

func (o FabricVlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricVlan) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AutoAllowOnUplinks) {
		toSerialize["AutoAllowOnUplinks"] = o.AutoAllowOnUplinks
	}
	if !IsNil(o.IsNative) {
		toSerialize["IsNative"] = o.IsNative
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PrimaryVlanId) {
		toSerialize["PrimaryVlanId"] = o.PrimaryVlanId
	}
	if !IsNil(o.SharingType) {
		toSerialize["SharingType"] = o.SharingType
	}
	if !IsNil(o.VlanId) {
		toSerialize["VlanId"] = o.VlanId
	}
	if o.EthNetworkPolicy.IsSet() {
		toSerialize["EthNetworkPolicy"] = o.EthNetworkPolicy.Get()
	}
	if o.MulticastPolicy.IsSet() {
		toSerialize["MulticastPolicy"] = o.MulticastPolicy.Get()
	}
	if o.VlanSet.IsSet() {
		toSerialize["VlanSet"] = o.VlanSet.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricVlan) UnmarshalJSON(data []byte) (err error) {
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
	type FabricVlanWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enable to automatically allow this VLAN on all uplinks. Disable must be specified for Disjoint Layer 2 VLAN configuration. Default VLAN-1 cannot be configured as Disjoint Layer 2 VLAN.
		AutoAllowOnUplinks *bool `json:"AutoAllowOnUplinks,omitempty"`
		// Used to define whether this VLAN is to be classified as 'native' for traffic in this FI.
		IsNative *bool `json:"IsNative,omitempty"`
		// The 'name' used to identify this VLAN.
		Name *string `json:"Name,omitempty"`
		// The Primary VLAN ID of the VLAN, if the sharing type of the VLAN is Isolated or Community.
		PrimaryVlanId *int64 `json:"PrimaryVlanId,omitempty"`
		// The sharing type of this VLAN. * `None` - This represents a regular VLAN. * `Primary` - This represents a primary VLAN. * `Isolated` - This represents an isolated VLAN. * `Community` - This represents a community VLAN.
		SharingType *string `json:"SharingType,omitempty"`
		// The identifier for this Virtual LAN.
		VlanId           *int64                                     `json:"VlanId,omitempty"`
		EthNetworkPolicy NullableFabricEthNetworkPolicyRelationship `json:"EthNetworkPolicy,omitempty"`
		MulticastPolicy  NullableFabricMulticastPolicyRelationship  `json:"MulticastPolicy,omitempty"`
		VlanSet          NullableFabricVlanSetRelationship          `json:"VlanSet,omitempty"`
	}

	varFabricVlanWithoutEmbeddedStruct := FabricVlanWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricVlanWithoutEmbeddedStruct)
	if err == nil {
		varFabricVlan := _FabricVlan{}
		varFabricVlan.ClassId = varFabricVlanWithoutEmbeddedStruct.ClassId
		varFabricVlan.ObjectType = varFabricVlanWithoutEmbeddedStruct.ObjectType
		varFabricVlan.AutoAllowOnUplinks = varFabricVlanWithoutEmbeddedStruct.AutoAllowOnUplinks
		varFabricVlan.IsNative = varFabricVlanWithoutEmbeddedStruct.IsNative
		varFabricVlan.Name = varFabricVlanWithoutEmbeddedStruct.Name
		varFabricVlan.PrimaryVlanId = varFabricVlanWithoutEmbeddedStruct.PrimaryVlanId
		varFabricVlan.SharingType = varFabricVlanWithoutEmbeddedStruct.SharingType
		varFabricVlan.VlanId = varFabricVlanWithoutEmbeddedStruct.VlanId
		varFabricVlan.EthNetworkPolicy = varFabricVlanWithoutEmbeddedStruct.EthNetworkPolicy
		varFabricVlan.MulticastPolicy = varFabricVlanWithoutEmbeddedStruct.MulticastPolicy
		varFabricVlan.VlanSet = varFabricVlanWithoutEmbeddedStruct.VlanSet
		*o = FabricVlan(varFabricVlan)
	} else {
		return err
	}

	varFabricVlan := _FabricVlan{}

	err = json.Unmarshal(data, &varFabricVlan)
	if err == nil {
		o.MoBaseMo = varFabricVlan.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoAllowOnUplinks")
		delete(additionalProperties, "IsNative")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PrimaryVlanId")
		delete(additionalProperties, "SharingType")
		delete(additionalProperties, "VlanId")
		delete(additionalProperties, "EthNetworkPolicy")
		delete(additionalProperties, "MulticastPolicy")
		delete(additionalProperties, "VlanSet")

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

type NullableFabricVlan struct {
	value *FabricVlan
	isSet bool
}

func (v NullableFabricVlan) Get() *FabricVlan {
	return v.value
}

func (v *NullableFabricVlan) Set(val *FabricVlan) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricVlan(val *FabricVlan) *NullableFabricVlan {
	return &NullableFabricVlan{value: val, isSet: true}
}

func (v NullableFabricVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
