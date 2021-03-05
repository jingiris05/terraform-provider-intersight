/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-24T06:47:07Z.
 *
 * API version: 1.0.9-3824
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// AssetVmHost Type for saving VM host details.
type AssetVmHost struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Reference to virtualization target account ID.
	AccountMoid *string `json:"AccountMoid,omitempty"`
	// The connection status of the host. 1 represents being connected and 0 represents being disconnected.
	Connected *int64 `json:"Connected,omitempty"`
	// Reference to virtualization target device ID.
	RegisteredDeviceMoid *string `json:"RegisteredDeviceMoid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetVmHost AssetVmHost

// NewAssetVmHost instantiates a new AssetVmHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetVmHost(classId string, objectType string) *AssetVmHost {
	this := AssetVmHost{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetVmHostWithDefaults instantiates a new AssetVmHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetVmHostWithDefaults() *AssetVmHost {
	this := AssetVmHost{}
	var classId string = "asset.VmHost"
	this.ClassId = classId
	var objectType string = "asset.VmHost"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetVmHost) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetVmHost) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetVmHost) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetVmHost) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccountMoid returns the AccountMoid field value if set, zero value otherwise.
func (o *AssetVmHost) GetAccountMoid() string {
	if o == nil || o.AccountMoid == nil {
		var ret string
		return ret
	}
	return *o.AccountMoid
}

// GetAccountMoidOk returns a tuple with the AccountMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetAccountMoidOk() (*string, bool) {
	if o == nil || o.AccountMoid == nil {
		return nil, false
	}
	return o.AccountMoid, true
}

// HasAccountMoid returns a boolean if a field has been set.
func (o *AssetVmHost) HasAccountMoid() bool {
	if o != nil && o.AccountMoid != nil {
		return true
	}

	return false
}

// SetAccountMoid gets a reference to the given string and assigns it to the AccountMoid field.
func (o *AssetVmHost) SetAccountMoid(v string) {
	o.AccountMoid = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *AssetVmHost) GetConnected() int64 {
	if o == nil || o.Connected == nil {
		var ret int64
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetConnectedOk() (*int64, bool) {
	if o == nil || o.Connected == nil {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *AssetVmHost) HasConnected() bool {
	if o != nil && o.Connected != nil {
		return true
	}

	return false
}

// SetConnected gets a reference to the given int64 and assigns it to the Connected field.
func (o *AssetVmHost) SetConnected(v int64) {
	o.Connected = &v
}

// GetRegisteredDeviceMoid returns the RegisteredDeviceMoid field value if set, zero value otherwise.
func (o *AssetVmHost) GetRegisteredDeviceMoid() string {
	if o == nil || o.RegisteredDeviceMoid == nil {
		var ret string
		return ret
	}
	return *o.RegisteredDeviceMoid
}

// GetRegisteredDeviceMoidOk returns a tuple with the RegisteredDeviceMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetVmHost) GetRegisteredDeviceMoidOk() (*string, bool) {
	if o == nil || o.RegisteredDeviceMoid == nil {
		return nil, false
	}
	return o.RegisteredDeviceMoid, true
}

// HasRegisteredDeviceMoid returns a boolean if a field has been set.
func (o *AssetVmHost) HasRegisteredDeviceMoid() bool {
	if o != nil && o.RegisteredDeviceMoid != nil {
		return true
	}

	return false
}

// SetRegisteredDeviceMoid gets a reference to the given string and assigns it to the RegisteredDeviceMoid field.
func (o *AssetVmHost) SetRegisteredDeviceMoid(v string) {
	o.RegisteredDeviceMoid = &v
}

func (o AssetVmHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccountMoid != nil {
		toSerialize["AccountMoid"] = o.AccountMoid
	}
	if o.Connected != nil {
		toSerialize["Connected"] = o.Connected
	}
	if o.RegisteredDeviceMoid != nil {
		toSerialize["RegisteredDeviceMoid"] = o.RegisteredDeviceMoid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetVmHost) UnmarshalJSON(bytes []byte) (err error) {
	type AssetVmHostWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Reference to virtualization target account ID.
		AccountMoid *string `json:"AccountMoid,omitempty"`
		// The connection status of the host. 1 represents being connected and 0 represents being disconnected.
		Connected *int64 `json:"Connected,omitempty"`
		// Reference to virtualization target device ID.
		RegisteredDeviceMoid *string `json:"RegisteredDeviceMoid,omitempty"`
	}

	varAssetVmHostWithoutEmbeddedStruct := AssetVmHostWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetVmHostWithoutEmbeddedStruct)
	if err == nil {
		varAssetVmHost := _AssetVmHost{}
		varAssetVmHost.ClassId = varAssetVmHostWithoutEmbeddedStruct.ClassId
		varAssetVmHost.ObjectType = varAssetVmHostWithoutEmbeddedStruct.ObjectType
		varAssetVmHost.AccountMoid = varAssetVmHostWithoutEmbeddedStruct.AccountMoid
		varAssetVmHost.Connected = varAssetVmHostWithoutEmbeddedStruct.Connected
		varAssetVmHost.RegisteredDeviceMoid = varAssetVmHostWithoutEmbeddedStruct.RegisteredDeviceMoid
		*o = AssetVmHost(varAssetVmHost)
	} else {
		return err
	}

	varAssetVmHost := _AssetVmHost{}

	err = json.Unmarshal(bytes, &varAssetVmHost)
	if err == nil {
		o.MoBaseComplexType = varAssetVmHost.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccountMoid")
		delete(additionalProperties, "Connected")
		delete(additionalProperties, "RegisteredDeviceMoid")

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

type NullableAssetVmHost struct {
	value *AssetVmHost
	isSet bool
}

func (v NullableAssetVmHost) Get() *AssetVmHost {
	return v.value
}

func (v *NullableAssetVmHost) Set(val *AssetVmHost) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetVmHost) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetVmHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetVmHost(val *AssetVmHost) *NullableAssetVmHost {
	return &NullableAssetVmHost{value: val, isSet: true}
}

func (v NullableAssetVmHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetVmHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}