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

// checks if the VirtualizationVmwareVmMemoryShareInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationVmwareVmMemoryShareInfo{}

// VirtualizationVmwareVmMemoryShareInfo Information about the virtual machine's memory sharing and limits. For more information, see VMware documentation.
type VirtualizationVmwareVmMemoryShareInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Limit on the memory sharing imposed (in Mbytes).
	MemLimit *int64 `json:"MemLimit,omitempty"`
	// Limit on memory overhead imposed (in Mbytes).
	MemOverheadLimit *int64 `json:"MemOverheadLimit,omitempty"`
	// Similar to CPU reservations (Mbytes).
	MemReservation *int64 `json:"MemReservation,omitempty"`
	// Similar to CPU Shares but applicable to memory. There is no unit for this value. It is a relative measure based on the settings for other resource pools.
	MemShares            *int64 `json:"MemShares,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVmMemoryShareInfo VirtualizationVmwareVmMemoryShareInfo

// NewVirtualizationVmwareVmMemoryShareInfo instantiates a new VirtualizationVmwareVmMemoryShareInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVmMemoryShareInfo(classId string, objectType string) *VirtualizationVmwareVmMemoryShareInfo {
	this := VirtualizationVmwareVmMemoryShareInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVmMemoryShareInfoWithDefaults instantiates a new VirtualizationVmwareVmMemoryShareInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVmMemoryShareInfoWithDefaults() *VirtualizationVmwareVmMemoryShareInfo {
	this := VirtualizationVmwareVmMemoryShareInfo{}
	var classId string = "virtualization.VmwareVmMemoryShareInfo"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVmMemoryShareInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVmMemoryShareInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVmMemoryShareInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareVmMemoryShareInfo" of the ClassId field.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetDefaultClassId() interface{} {
	return "virtualization.VmwareVmMemoryShareInfo"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVmMemoryShareInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVmMemoryShareInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareVmMemoryShareInfo" of the ObjectType field.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareVmMemoryShareInfo"
}

// GetMemLimit returns the MemLimit field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemLimit() int64 {
	if o == nil || IsNil(o.MemLimit) {
		var ret int64
		return ret
	}
	return *o.MemLimit
}

// GetMemLimitOk returns a tuple with the MemLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.MemLimit) {
		return nil, false
	}
	return o.MemLimit, true
}

// HasMemLimit returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemLimit() bool {
	if o != nil && !IsNil(o.MemLimit) {
		return true
	}

	return false
}

// SetMemLimit gets a reference to the given int64 and assigns it to the MemLimit field.
func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemLimit(v int64) {
	o.MemLimit = &v
}

// GetMemOverheadLimit returns the MemOverheadLimit field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemOverheadLimit() int64 {
	if o == nil || IsNil(o.MemOverheadLimit) {
		var ret int64
		return ret
	}
	return *o.MemOverheadLimit
}

// GetMemOverheadLimitOk returns a tuple with the MemOverheadLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemOverheadLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.MemOverheadLimit) {
		return nil, false
	}
	return o.MemOverheadLimit, true
}

// HasMemOverheadLimit returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemOverheadLimit() bool {
	if o != nil && !IsNil(o.MemOverheadLimit) {
		return true
	}

	return false
}

// SetMemOverheadLimit gets a reference to the given int64 and assigns it to the MemOverheadLimit field.
func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemOverheadLimit(v int64) {
	o.MemOverheadLimit = &v
}

// GetMemReservation returns the MemReservation field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemReservation() int64 {
	if o == nil || IsNil(o.MemReservation) {
		var ret int64
		return ret
	}
	return *o.MemReservation
}

// GetMemReservationOk returns a tuple with the MemReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemReservationOk() (*int64, bool) {
	if o == nil || IsNil(o.MemReservation) {
		return nil, false
	}
	return o.MemReservation, true
}

// HasMemReservation returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemReservation() bool {
	if o != nil && !IsNil(o.MemReservation) {
		return true
	}

	return false
}

// SetMemReservation gets a reference to the given int64 and assigns it to the MemReservation field.
func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemReservation(v int64) {
	o.MemReservation = &v
}

// GetMemShares returns the MemShares field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemShares() int64 {
	if o == nil || IsNil(o.MemShares) {
		var ret int64
		return ret
	}
	return *o.MemShares
}

// GetMemSharesOk returns a tuple with the MemShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) GetMemSharesOk() (*int64, bool) {
	if o == nil || IsNil(o.MemShares) {
		return nil, false
	}
	return o.MemShares, true
}

// HasMemShares returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmMemoryShareInfo) HasMemShares() bool {
	if o != nil && !IsNil(o.MemShares) {
		return true
	}

	return false
}

// SetMemShares gets a reference to the given int64 and assigns it to the MemShares field.
func (o *VirtualizationVmwareVmMemoryShareInfo) SetMemShares(v int64) {
	o.MemShares = &v
}

func (o VirtualizationVmwareVmMemoryShareInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationVmwareVmMemoryShareInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.MemLimit) {
		toSerialize["MemLimit"] = o.MemLimit
	}
	if !IsNil(o.MemOverheadLimit) {
		toSerialize["MemOverheadLimit"] = o.MemOverheadLimit
	}
	if !IsNil(o.MemReservation) {
		toSerialize["MemReservation"] = o.MemReservation
	}
	if !IsNil(o.MemShares) {
		toSerialize["MemShares"] = o.MemShares
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationVmwareVmMemoryShareInfo) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Limit on the memory sharing imposed (in Mbytes).
		MemLimit *int64 `json:"MemLimit,omitempty"`
		// Limit on memory overhead imposed (in Mbytes).
		MemOverheadLimit *int64 `json:"MemOverheadLimit,omitempty"`
		// Similar to CPU reservations (Mbytes).
		MemReservation *int64 `json:"MemReservation,omitempty"`
		// Similar to CPU Shares but applicable to memory. There is no unit for this value. It is a relative measure based on the settings for other resource pools.
		MemShares *int64 `json:"MemShares,omitempty"`
	}

	varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct := VirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareVmMemoryShareInfo := _VirtualizationVmwareVmMemoryShareInfo{}
		varVirtualizationVmwareVmMemoryShareInfo.ClassId = varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareVmMemoryShareInfo.ObjectType = varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareVmMemoryShareInfo.MemLimit = varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct.MemLimit
		varVirtualizationVmwareVmMemoryShareInfo.MemOverheadLimit = varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct.MemOverheadLimit
		varVirtualizationVmwareVmMemoryShareInfo.MemReservation = varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct.MemReservation
		varVirtualizationVmwareVmMemoryShareInfo.MemShares = varVirtualizationVmwareVmMemoryShareInfoWithoutEmbeddedStruct.MemShares
		*o = VirtualizationVmwareVmMemoryShareInfo(varVirtualizationVmwareVmMemoryShareInfo)
	} else {
		return err
	}

	varVirtualizationVmwareVmMemoryShareInfo := _VirtualizationVmwareVmMemoryShareInfo{}

	err = json.Unmarshal(data, &varVirtualizationVmwareVmMemoryShareInfo)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVmwareVmMemoryShareInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MemLimit")
		delete(additionalProperties, "MemOverheadLimit")
		delete(additionalProperties, "MemReservation")
		delete(additionalProperties, "MemShares")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVirtualizationVmwareVmMemoryShareInfo struct {
	value *VirtualizationVmwareVmMemoryShareInfo
	isSet bool
}

func (v NullableVirtualizationVmwareVmMemoryShareInfo) Get() *VirtualizationVmwareVmMemoryShareInfo {
	return v.value
}

func (v *NullableVirtualizationVmwareVmMemoryShareInfo) Set(val *VirtualizationVmwareVmMemoryShareInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVmMemoryShareInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVmMemoryShareInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVmMemoryShareInfo(val *VirtualizationVmwareVmMemoryShareInfo) *NullableVirtualizationVmwareVmMemoryShareInfo {
	return &NullableVirtualizationVmwareVmMemoryShareInfo{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVmMemoryShareInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVmMemoryShareInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
