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

// checks if the VirtualizationBaseVirtualMachineSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationBaseVirtualMachineSnapshot{}

// VirtualizationBaseVirtualMachineSnapshot A snapshot preserves the state and data of a virtual machine at a specific point in time.
type VirtualizationBaseVirtualMachineSnapshot struct {
	VirtualizationBaseSourceDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The internally generated identity of the snapshot. This entity is not manipulated by users. It aids in uniquely identifying the snapshot object. For VMware, this is a MOR (managed object reference).
	Identity *string `json:"Identity,omitempty"`
	// User name provided to identify the snapshot.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseVirtualMachineSnapshot VirtualizationBaseVirtualMachineSnapshot

// NewVirtualizationBaseVirtualMachineSnapshot instantiates a new VirtualizationBaseVirtualMachineSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVirtualMachineSnapshot(classId string, objectType string) *VirtualizationBaseVirtualMachineSnapshot {
	this := VirtualizationBaseVirtualMachineSnapshot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationBaseVirtualMachineSnapshotWithDefaults instantiates a new VirtualizationBaseVirtualMachineSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVirtualMachineSnapshotWithDefaults() *VirtualizationBaseVirtualMachineSnapshot {
	this := VirtualizationBaseVirtualMachineSnapshot{}
	var classId string = "virtualization.VmwareVirtualMachineSnapshot"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVirtualMachineSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseVirtualMachineSnapshot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineSnapshot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseVirtualMachineSnapshot) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareVirtualMachineSnapshot" of the ClassId field.
func (o *VirtualizationBaseVirtualMachineSnapshot) GetDefaultClassId() interface{} {
	return "virtualization.VmwareVirtualMachineSnapshot"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseVirtualMachineSnapshot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineSnapshot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseVirtualMachineSnapshot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareVirtualMachineSnapshot" of the ObjectType field.
func (o *VirtualizationBaseVirtualMachineSnapshot) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareVirtualMachineSnapshot"
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineSnapshot) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineSnapshot) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineSnapshot) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseVirtualMachineSnapshot) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachineSnapshot) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachineSnapshot) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachineSnapshot) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseVirtualMachineSnapshot) SetName(v string) {
	o.Name = &v
}

func (o VirtualizationBaseVirtualMachineSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationBaseVirtualMachineSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseSourceDevice, errVirtualizationBaseSourceDevice := json.Marshal(o.VirtualizationBaseSourceDevice)
	if errVirtualizationBaseSourceDevice != nil {
		return map[string]interface{}{}, errVirtualizationBaseSourceDevice
	}
	errVirtualizationBaseSourceDevice = json.Unmarshal([]byte(serializedVirtualizationBaseSourceDevice), &toSerialize)
	if errVirtualizationBaseSourceDevice != nil {
		return map[string]interface{}{}, errVirtualizationBaseSourceDevice
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Identity) {
		toSerialize["Identity"] = o.Identity
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationBaseVirtualMachineSnapshot) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationBaseVirtualMachineSnapshotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The internally generated identity of the snapshot. This entity is not manipulated by users. It aids in uniquely identifying the snapshot object. For VMware, this is a MOR (managed object reference).
		Identity *string `json:"Identity,omitempty"`
		// User name provided to identify the snapshot.
		Name *string `json:"Name,omitempty"`
	}

	varVirtualizationBaseVirtualMachineSnapshotWithoutEmbeddedStruct := VirtualizationBaseVirtualMachineSnapshotWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationBaseVirtualMachineSnapshotWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBaseVirtualMachineSnapshot := _VirtualizationBaseVirtualMachineSnapshot{}
		varVirtualizationBaseVirtualMachineSnapshot.ClassId = varVirtualizationBaseVirtualMachineSnapshotWithoutEmbeddedStruct.ClassId
		varVirtualizationBaseVirtualMachineSnapshot.ObjectType = varVirtualizationBaseVirtualMachineSnapshotWithoutEmbeddedStruct.ObjectType
		varVirtualizationBaseVirtualMachineSnapshot.Identity = varVirtualizationBaseVirtualMachineSnapshotWithoutEmbeddedStruct.Identity
		varVirtualizationBaseVirtualMachineSnapshot.Name = varVirtualizationBaseVirtualMachineSnapshotWithoutEmbeddedStruct.Name
		*o = VirtualizationBaseVirtualMachineSnapshot(varVirtualizationBaseVirtualMachineSnapshot)
	} else {
		return err
	}

	varVirtualizationBaseVirtualMachineSnapshot := _VirtualizationBaseVirtualMachineSnapshot{}

	err = json.Unmarshal(data, &varVirtualizationBaseVirtualMachineSnapshot)
	if err == nil {
		o.VirtualizationBaseSourceDevice = varVirtualizationBaseVirtualMachineSnapshot.VirtualizationBaseSourceDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")

		// remove fields from embedded structs
		reflectVirtualizationBaseSourceDevice := reflect.ValueOf(o.VirtualizationBaseSourceDevice)
		for i := 0; i < reflectVirtualizationBaseSourceDevice.Type().NumField(); i++ {
			t := reflectVirtualizationBaseSourceDevice.Type().Field(i)

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

type NullableVirtualizationBaseVirtualMachineSnapshot struct {
	value *VirtualizationBaseVirtualMachineSnapshot
	isSet bool
}

func (v NullableVirtualizationBaseVirtualMachineSnapshot) Get() *VirtualizationBaseVirtualMachineSnapshot {
	return v.value
}

func (v *NullableVirtualizationBaseVirtualMachineSnapshot) Set(val *VirtualizationBaseVirtualMachineSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVirtualMachineSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVirtualMachineSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVirtualMachineSnapshot(val *VirtualizationBaseVirtualMachineSnapshot) *NullableVirtualizationBaseVirtualMachineSnapshot {
	return &NullableVirtualizationBaseVirtualMachineSnapshot{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVirtualMachineSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVirtualMachineSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
