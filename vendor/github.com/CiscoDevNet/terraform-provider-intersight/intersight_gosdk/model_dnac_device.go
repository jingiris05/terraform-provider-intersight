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

// checks if the DnacDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnacDevice{}

// DnacDevice Collection of network devices.
type DnacDevice struct {
	DnacInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identity of the network device.
	DeviceId *string `json:"DeviceId,omitempty"`
	// Host name of the network device.
	HostName *string `json:"HostName,omitempty"`
	// IP address of the network device.
	ManagementIpAddress  *string `json:"ManagementIpAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DnacDevice DnacDevice

// NewDnacDevice instantiates a new DnacDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnacDevice(classId string, objectType string) *DnacDevice {
	this := DnacDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewDnacDeviceWithDefaults instantiates a new DnacDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnacDeviceWithDefaults() *DnacDevice {
	this := DnacDevice{}
	var classId string = "dnac.Device"
	this.ClassId = classId
	var objectType string = "dnac.Device"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *DnacDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *DnacDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *DnacDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "dnac.Device" of the ClassId field.
func (o *DnacDevice) GetDefaultClassId() interface{} {
	return "dnac.Device"
}

// GetObjectType returns the ObjectType field value
func (o *DnacDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DnacDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DnacDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "dnac.Device" of the ObjectType field.
func (o *DnacDevice) GetDefaultObjectType() interface{} {
	return "dnac.Device"
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DnacDevice) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DnacDevice) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DnacDevice) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *DnacDevice) GetHostName() string {
	if o == nil || IsNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacDevice) GetHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *DnacDevice) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *DnacDevice) SetHostName(v string) {
	o.HostName = &v
}

// GetManagementIpAddress returns the ManagementIpAddress field value if set, zero value otherwise.
func (o *DnacDevice) GetManagementIpAddress() string {
	if o == nil || IsNil(o.ManagementIpAddress) {
		var ret string
		return ret
	}
	return *o.ManagementIpAddress
}

// GetManagementIpAddressOk returns a tuple with the ManagementIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacDevice) GetManagementIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementIpAddress) {
		return nil, false
	}
	return o.ManagementIpAddress, true
}

// HasManagementIpAddress returns a boolean if a field has been set.
func (o *DnacDevice) HasManagementIpAddress() bool {
	if o != nil && !IsNil(o.ManagementIpAddress) {
		return true
	}

	return false
}

// SetManagementIpAddress gets a reference to the given string and assigns it to the ManagementIpAddress field.
func (o *DnacDevice) SetManagementIpAddress(v string) {
	o.ManagementIpAddress = &v
}

func (o DnacDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnacDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDnacInventoryEntity, errDnacInventoryEntity := json.Marshal(o.DnacInventoryEntity)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	errDnacInventoryEntity = json.Unmarshal([]byte(serializedDnacInventoryEntity), &toSerialize)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
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
	if !IsNil(o.HostName) {
		toSerialize["HostName"] = o.HostName
	}
	if !IsNil(o.ManagementIpAddress) {
		toSerialize["ManagementIpAddress"] = o.ManagementIpAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DnacDevice) UnmarshalJSON(data []byte) (err error) {
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
	type DnacDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Unique identity of the network device.
		DeviceId *string `json:"DeviceId,omitempty"`
		// Host name of the network device.
		HostName *string `json:"HostName,omitempty"`
		// IP address of the network device.
		ManagementIpAddress *string `json:"ManagementIpAddress,omitempty"`
	}

	varDnacDeviceWithoutEmbeddedStruct := DnacDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDnacDeviceWithoutEmbeddedStruct)
	if err == nil {
		varDnacDevice := _DnacDevice{}
		varDnacDevice.ClassId = varDnacDeviceWithoutEmbeddedStruct.ClassId
		varDnacDevice.ObjectType = varDnacDeviceWithoutEmbeddedStruct.ObjectType
		varDnacDevice.DeviceId = varDnacDeviceWithoutEmbeddedStruct.DeviceId
		varDnacDevice.HostName = varDnacDeviceWithoutEmbeddedStruct.HostName
		varDnacDevice.ManagementIpAddress = varDnacDeviceWithoutEmbeddedStruct.ManagementIpAddress
		*o = DnacDevice(varDnacDevice)
	} else {
		return err
	}

	varDnacDevice := _DnacDevice{}

	err = json.Unmarshal(data, &varDnacDevice)
	if err == nil {
		o.DnacInventoryEntity = varDnacDevice.DnacInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "ManagementIpAddress")

		// remove fields from embedded structs
		reflectDnacInventoryEntity := reflect.ValueOf(o.DnacInventoryEntity)
		for i := 0; i < reflectDnacInventoryEntity.Type().NumField(); i++ {
			t := reflectDnacInventoryEntity.Type().Field(i)

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

type NullableDnacDevice struct {
	value *DnacDevice
	isSet bool
}

func (v NullableDnacDevice) Get() *DnacDevice {
	return v.value
}

func (v *NullableDnacDevice) Set(val *DnacDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableDnacDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDnacDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnacDevice(val *DnacDevice) *NullableDnacDevice {
	return &NullableDnacDevice{value: val, isSet: true}
}

func (v NullableDnacDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnacDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
