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
	"reflect"
	"strings"
)

// BiosBootDevice Actual boot devices of the system as enumerated by BIOS.
type BiosBootDevice struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the Configured Boot Device.
	DeviceName *string `json:"DeviceName,omitempty"`
	// Type of the Configured Boot Device.
	DeviceType           *string                              `json:"DeviceType,omitempty"`
	BiosSystemBootOrder  *BiosSystemBootOrderRelationship     `json:"BiosSystemBootOrder,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BiosBootDevice BiosBootDevice

// NewBiosBootDevice instantiates a new BiosBootDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosBootDevice(classId string, objectType string) *BiosBootDevice {
	this := BiosBootDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBiosBootDeviceWithDefaults instantiates a new BiosBootDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosBootDeviceWithDefaults() *BiosBootDevice {
	this := BiosBootDevice{}
	var classId string = "bios.BootDevice"
	this.ClassId = classId
	var objectType string = "bios.BootDevice"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BiosBootDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BiosBootDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BiosBootDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BiosBootDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BiosBootDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BiosBootDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *BiosBootDevice) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootDevice) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *BiosBootDevice) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *BiosBootDevice) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *BiosBootDevice) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootDevice) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *BiosBootDevice) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *BiosBootDevice) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetBiosSystemBootOrder returns the BiosSystemBootOrder field value if set, zero value otherwise.
func (o *BiosBootDevice) GetBiosSystemBootOrder() BiosSystemBootOrderRelationship {
	if o == nil || o.BiosSystemBootOrder == nil {
		var ret BiosSystemBootOrderRelationship
		return ret
	}
	return *o.BiosSystemBootOrder
}

// GetBiosSystemBootOrderOk returns a tuple with the BiosSystemBootOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootDevice) GetBiosSystemBootOrderOk() (*BiosSystemBootOrderRelationship, bool) {
	if o == nil || o.BiosSystemBootOrder == nil {
		return nil, false
	}
	return o.BiosSystemBootOrder, true
}

// HasBiosSystemBootOrder returns a boolean if a field has been set.
func (o *BiosBootDevice) HasBiosSystemBootOrder() bool {
	if o != nil && o.BiosSystemBootOrder != nil {
		return true
	}

	return false
}

// SetBiosSystemBootOrder gets a reference to the given BiosSystemBootOrderRelationship and assigns it to the BiosSystemBootOrder field.
func (o *BiosBootDevice) SetBiosSystemBootOrder(v BiosSystemBootOrderRelationship) {
	o.BiosSystemBootOrder = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BiosBootDevice) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosBootDevice) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosBootDevice) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosBootDevice) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BiosBootDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.DeviceType != nil {
		toSerialize["DeviceType"] = o.DeviceType
	}
	if o.BiosSystemBootOrder != nil {
		toSerialize["BiosSystemBootOrder"] = o.BiosSystemBootOrder
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BiosBootDevice) UnmarshalJSON(bytes []byte) (err error) {
	type BiosBootDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the Configured Boot Device.
		DeviceName *string `json:"DeviceName,omitempty"`
		// Type of the Configured Boot Device.
		DeviceType          *string                              `json:"DeviceType,omitempty"`
		BiosSystemBootOrder *BiosSystemBootOrderRelationship     `json:"BiosSystemBootOrder,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varBiosBootDeviceWithoutEmbeddedStruct := BiosBootDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBiosBootDeviceWithoutEmbeddedStruct)
	if err == nil {
		varBiosBootDevice := _BiosBootDevice{}
		varBiosBootDevice.ClassId = varBiosBootDeviceWithoutEmbeddedStruct.ClassId
		varBiosBootDevice.ObjectType = varBiosBootDeviceWithoutEmbeddedStruct.ObjectType
		varBiosBootDevice.DeviceName = varBiosBootDeviceWithoutEmbeddedStruct.DeviceName
		varBiosBootDevice.DeviceType = varBiosBootDeviceWithoutEmbeddedStruct.DeviceType
		varBiosBootDevice.BiosSystemBootOrder = varBiosBootDeviceWithoutEmbeddedStruct.BiosSystemBootOrder
		varBiosBootDevice.RegisteredDevice = varBiosBootDeviceWithoutEmbeddedStruct.RegisteredDevice
		*o = BiosBootDevice(varBiosBootDevice)
	} else {
		return err
	}

	varBiosBootDevice := _BiosBootDevice{}

	err = json.Unmarshal(bytes, &varBiosBootDevice)
	if err == nil {
		o.MoBaseMo = varBiosBootDevice.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceName")
		delete(additionalProperties, "DeviceType")
		delete(additionalProperties, "BiosSystemBootOrder")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableBiosBootDevice struct {
	value *BiosBootDevice
	isSet bool
}

func (v NullableBiosBootDevice) Get() *BiosBootDevice {
	return v.value
}

func (v *NullableBiosBootDevice) Set(val *BiosBootDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosBootDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosBootDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosBootDevice(val *BiosBootDevice) *NullableBiosBootDevice {
	return &NullableBiosBootDevice{value: val, isSet: true}
}

func (v NullableBiosBootDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosBootDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
