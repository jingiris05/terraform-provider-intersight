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

// checks if the FirmwareFirmwareSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareFirmwareSummary{}

// FirmwareFirmwareSummary Update inventory that contains the details for the firmware running on each component in the compute.Physical object.
type FirmwareFirmwareSummary struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Version details at the bundle level for the each of server.
	BundleVersion         *string                             `json:"BundleVersion,omitempty"`
	ComponentsFwInventory []FirmwareFirmwareInventory         `json:"ComponentsFwInventory,omitempty"`
	Server                NullableComputePhysicalRelationship `json:"Server,omitempty"`
	TargetMo              NullableMoBaseMoRelationship        `json:"TargetMo,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _FirmwareFirmwareSummary FirmwareFirmwareSummary

// NewFirmwareFirmwareSummary instantiates a new FirmwareFirmwareSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareFirmwareSummary(classId string, objectType string) *FirmwareFirmwareSummary {
	this := FirmwareFirmwareSummary{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareFirmwareSummaryWithDefaults instantiates a new FirmwareFirmwareSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareFirmwareSummaryWithDefaults() *FirmwareFirmwareSummary {
	this := FirmwareFirmwareSummary{}
	var classId string = "firmware.FirmwareSummary"
	this.ClassId = classId
	var objectType string = "firmware.FirmwareSummary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareFirmwareSummary) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareSummary) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareFirmwareSummary) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "firmware.FirmwareSummary" of the ClassId field.
func (o *FirmwareFirmwareSummary) GetDefaultClassId() interface{} {
	return "firmware.FirmwareSummary"
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareFirmwareSummary) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareSummary) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareFirmwareSummary) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "firmware.FirmwareSummary" of the ObjectType field.
func (o *FirmwareFirmwareSummary) GetDefaultObjectType() interface{} {
	return "firmware.FirmwareSummary"
}

// GetBundleVersion returns the BundleVersion field value if set, zero value otherwise.
func (o *FirmwareFirmwareSummary) GetBundleVersion() string {
	if o == nil || IsNil(o.BundleVersion) {
		var ret string
		return ret
	}
	return *o.BundleVersion
}

// GetBundleVersionOk returns a tuple with the BundleVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareFirmwareSummary) GetBundleVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BundleVersion) {
		return nil, false
	}
	return o.BundleVersion, true
}

// HasBundleVersion returns a boolean if a field has been set.
func (o *FirmwareFirmwareSummary) HasBundleVersion() bool {
	if o != nil && !IsNil(o.BundleVersion) {
		return true
	}

	return false
}

// SetBundleVersion gets a reference to the given string and assigns it to the BundleVersion field.
func (o *FirmwareFirmwareSummary) SetBundleVersion(v string) {
	o.BundleVersion = &v
}

// GetComponentsFwInventory returns the ComponentsFwInventory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareFirmwareSummary) GetComponentsFwInventory() []FirmwareFirmwareInventory {
	if o == nil {
		var ret []FirmwareFirmwareInventory
		return ret
	}
	return o.ComponentsFwInventory
}

// GetComponentsFwInventoryOk returns a tuple with the ComponentsFwInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareFirmwareSummary) GetComponentsFwInventoryOk() ([]FirmwareFirmwareInventory, bool) {
	if o == nil || IsNil(o.ComponentsFwInventory) {
		return nil, false
	}
	return o.ComponentsFwInventory, true
}

// HasComponentsFwInventory returns a boolean if a field has been set.
func (o *FirmwareFirmwareSummary) HasComponentsFwInventory() bool {
	if o != nil && !IsNil(o.ComponentsFwInventory) {
		return true
	}

	return false
}

// SetComponentsFwInventory gets a reference to the given []FirmwareFirmwareInventory and assigns it to the ComponentsFwInventory field.
func (o *FirmwareFirmwareSummary) SetComponentsFwInventory(v []FirmwareFirmwareInventory) {
	o.ComponentsFwInventory = v
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareFirmwareSummary) GetServer() ComputePhysicalRelationship {
	if o == nil || IsNil(o.Server.Get()) {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareFirmwareSummary) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *FirmwareFirmwareSummary) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableComputePhysicalRelationship and assigns it to the Server field.
func (o *FirmwareFirmwareSummary) SetServer(v ComputePhysicalRelationship) {
	o.Server.Set(&v)
}

// SetServerNil sets the value for Server to be an explicit nil
func (o *FirmwareFirmwareSummary) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *FirmwareFirmwareSummary) UnsetServer() {
	o.Server.Unset()
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareFirmwareSummary) GetTargetMo() MoBaseMoRelationship {
	if o == nil || IsNil(o.TargetMo.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo.Get()
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareFirmwareSummary) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetMo.Get(), o.TargetMo.IsSet()
}

// HasTargetMo returns a boolean if a field has been set.
func (o *FirmwareFirmwareSummary) HasTargetMo() bool {
	if o != nil && o.TargetMo.IsSet() {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TargetMo field.
func (o *FirmwareFirmwareSummary) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo.Set(&v)
}

// SetTargetMoNil sets the value for TargetMo to be an explicit nil
func (o *FirmwareFirmwareSummary) SetTargetMoNil() {
	o.TargetMo.Set(nil)
}

// UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil
func (o *FirmwareFirmwareSummary) UnsetTargetMo() {
	o.TargetMo.Unset()
}

func (o FirmwareFirmwareSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareFirmwareSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.BundleVersion) {
		toSerialize["BundleVersion"] = o.BundleVersion
	}
	if o.ComponentsFwInventory != nil {
		toSerialize["ComponentsFwInventory"] = o.ComponentsFwInventory
	}
	if o.Server.IsSet() {
		toSerialize["Server"] = o.Server.Get()
	}
	if o.TargetMo.IsSet() {
		toSerialize["TargetMo"] = o.TargetMo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareFirmwareSummary) UnmarshalJSON(data []byte) (err error) {
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
	type FirmwareFirmwareSummaryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Version details at the bundle level for the each of server.
		BundleVersion         *string                             `json:"BundleVersion,omitempty"`
		ComponentsFwInventory []FirmwareFirmwareInventory         `json:"ComponentsFwInventory,omitempty"`
		Server                NullableComputePhysicalRelationship `json:"Server,omitempty"`
		TargetMo              NullableMoBaseMoRelationship        `json:"TargetMo,omitempty"`
	}

	varFirmwareFirmwareSummaryWithoutEmbeddedStruct := FirmwareFirmwareSummaryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFirmwareFirmwareSummaryWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareFirmwareSummary := _FirmwareFirmwareSummary{}
		varFirmwareFirmwareSummary.ClassId = varFirmwareFirmwareSummaryWithoutEmbeddedStruct.ClassId
		varFirmwareFirmwareSummary.ObjectType = varFirmwareFirmwareSummaryWithoutEmbeddedStruct.ObjectType
		varFirmwareFirmwareSummary.BundleVersion = varFirmwareFirmwareSummaryWithoutEmbeddedStruct.BundleVersion
		varFirmwareFirmwareSummary.ComponentsFwInventory = varFirmwareFirmwareSummaryWithoutEmbeddedStruct.ComponentsFwInventory
		varFirmwareFirmwareSummary.Server = varFirmwareFirmwareSummaryWithoutEmbeddedStruct.Server
		varFirmwareFirmwareSummary.TargetMo = varFirmwareFirmwareSummaryWithoutEmbeddedStruct.TargetMo
		*o = FirmwareFirmwareSummary(varFirmwareFirmwareSummary)
	} else {
		return err
	}

	varFirmwareFirmwareSummary := _FirmwareFirmwareSummary{}

	err = json.Unmarshal(data, &varFirmwareFirmwareSummary)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varFirmwareFirmwareSummary.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BundleVersion")
		delete(additionalProperties, "ComponentsFwInventory")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableFirmwareFirmwareSummary struct {
	value *FirmwareFirmwareSummary
	isSet bool
}

func (v NullableFirmwareFirmwareSummary) Get() *FirmwareFirmwareSummary {
	return v.value
}

func (v *NullableFirmwareFirmwareSummary) Set(val *FirmwareFirmwareSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareFirmwareSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareFirmwareSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareFirmwareSummary(val *FirmwareFirmwareSummary) *NullableFirmwareFirmwareSummary {
	return &NullableFirmwareFirmwareSummary{value: val, isSet: true}
}

func (v NullableFirmwareFirmwareSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareFirmwareSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
