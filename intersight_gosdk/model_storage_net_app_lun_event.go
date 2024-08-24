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

// checks if the StorageNetAppLunEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppLunEvent{}

// StorageNetAppLunEvent An event where the impacted resource type is a lun.
type StorageNetAppLunEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	Lun                  NullableStorageNetAppLunRelationship `json:"Lun,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppLunEvent StorageNetAppLunEvent

// NewStorageNetAppLunEvent instantiates a new StorageNetAppLunEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppLunEvent(classId string, objectType string) *StorageNetAppLunEvent {
	this := StorageNetAppLunEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppLunEventWithDefaults instantiates a new StorageNetAppLunEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppLunEventWithDefaults() *StorageNetAppLunEvent {
	this := StorageNetAppLunEvent{}
	var classId string = "storage.NetAppLunEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppLunEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppLunEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppLunEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppLunEvent" of the ClassId field.
func (o *StorageNetAppLunEvent) GetDefaultClassId() interface{} {
	return "storage.NetAppLunEvent"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppLunEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppLunEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppLunEvent" of the ObjectType field.
func (o *StorageNetAppLunEvent) GetDefaultObjectType() interface{} {
	return "storage.NetAppLunEvent"
}

// GetLun returns the Lun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppLunEvent) GetLun() StorageNetAppLunRelationship {
	if o == nil || IsNil(o.Lun.Get()) {
		var ret StorageNetAppLunRelationship
		return ret
	}
	return *o.Lun.Get()
}

// GetLunOk returns a tuple with the Lun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppLunEvent) GetLunOk() (*StorageNetAppLunRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lun.Get(), o.Lun.IsSet()
}

// HasLun returns a boolean if a field has been set.
func (o *StorageNetAppLunEvent) HasLun() bool {
	if o != nil && o.Lun.IsSet() {
		return true
	}

	return false
}

// SetLun gets a reference to the given NullableStorageNetAppLunRelationship and assigns it to the Lun field.
func (o *StorageNetAppLunEvent) SetLun(v StorageNetAppLunRelationship) {
	o.Lun.Set(&v)
}

// SetLunNil sets the value for Lun to be an explicit nil
func (o *StorageNetAppLunEvent) SetLunNil() {
	o.Lun.Set(nil)
}

// UnsetLun ensures that no value is present for Lun, not even an explicit nil
func (o *StorageNetAppLunEvent) UnsetLun() {
	o.Lun.Unset()
}

func (o StorageNetAppLunEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppLunEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageNetAppBaseEvent, errStorageNetAppBaseEvent := json.Marshal(o.StorageNetAppBaseEvent)
	if errStorageNetAppBaseEvent != nil {
		return map[string]interface{}{}, errStorageNetAppBaseEvent
	}
	errStorageNetAppBaseEvent = json.Unmarshal([]byte(serializedStorageNetAppBaseEvent), &toSerialize)
	if errStorageNetAppBaseEvent != nil {
		return map[string]interface{}{}, errStorageNetAppBaseEvent
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Lun.IsSet() {
		toSerialize["Lun"] = o.Lun.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppLunEvent) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppLunEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                               `json:"ObjectType"`
		Lun        NullableStorageNetAppLunRelationship `json:"Lun,omitempty"`
	}

	varStorageNetAppLunEventWithoutEmbeddedStruct := StorageNetAppLunEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppLunEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppLunEvent := _StorageNetAppLunEvent{}
		varStorageNetAppLunEvent.ClassId = varStorageNetAppLunEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppLunEvent.ObjectType = varStorageNetAppLunEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppLunEvent.Lun = varStorageNetAppLunEventWithoutEmbeddedStruct.Lun
		*o = StorageNetAppLunEvent(varStorageNetAppLunEvent)
	} else {
		return err
	}

	varStorageNetAppLunEvent := _StorageNetAppLunEvent{}

	err = json.Unmarshal(data, &varStorageNetAppLunEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppLunEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Lun")

		// remove fields from embedded structs
		reflectStorageNetAppBaseEvent := reflect.ValueOf(o.StorageNetAppBaseEvent)
		for i := 0; i < reflectStorageNetAppBaseEvent.Type().NumField(); i++ {
			t := reflectStorageNetAppBaseEvent.Type().Field(i)

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

type NullableStorageNetAppLunEvent struct {
	value *StorageNetAppLunEvent
	isSet bool
}

func (v NullableStorageNetAppLunEvent) Get() *StorageNetAppLunEvent {
	return v.value
}

func (v *NullableStorageNetAppLunEvent) Set(val *StorageNetAppLunEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppLunEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppLunEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppLunEvent(val *StorageNetAppLunEvent) *NullableStorageNetAppLunEvent {
	return &NullableStorageNetAppLunEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppLunEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppLunEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
