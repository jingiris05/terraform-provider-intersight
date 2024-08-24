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

// checks if the IppoolIpV4Block type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IppoolIpV4Block{}

// IppoolIpV4Block A block of IPv4 addresses.
type IppoolIpV4Block struct {
	PoolAbstractBlockType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// First IPv4 address of the block.
	From       *string                  `json:"From,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	IpV4Config NullableIppoolIpV4Config `json:"IpV4Config,omitempty"`
	// Last IPv4 address of the block.
	To                   *string `json:"To,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	AdditionalProperties map[string]interface{}
}

type _IppoolIpV4Block IppoolIpV4Block

// NewIppoolIpV4Block instantiates a new IppoolIpV4Block object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolIpV4Block(classId string, objectType string) *IppoolIpV4Block {
	this := IppoolIpV4Block{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIppoolIpV4BlockWithDefaults instantiates a new IppoolIpV4Block object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolIpV4BlockWithDefaults() *IppoolIpV4Block {
	this := IppoolIpV4Block{}
	var classId string = "ippool.IpV4Block"
	this.ClassId = classId
	var objectType string = "ippool.IpV4Block"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolIpV4Block) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolIpV4Block) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolIpV4Block) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "ippool.IpV4Block" of the ClassId field.
func (o *IppoolIpV4Block) GetDefaultClassId() interface{} {
	return "ippool.IpV4Block"
}

// GetObjectType returns the ObjectType field value
func (o *IppoolIpV4Block) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolIpV4Block) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolIpV4Block) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "ippool.IpV4Block" of the ObjectType field.
func (o *IppoolIpV4Block) GetDefaultObjectType() interface{} {
	return "ippool.IpV4Block"
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *IppoolIpV4Block) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4Block) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *IppoolIpV4Block) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *IppoolIpV4Block) SetFrom(v string) {
	o.From = &v
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolIpV4Block) GetIpV4Config() IppoolIpV4Config {
	if o == nil || IsNil(o.IpV4Config.Get()) {
		var ret IppoolIpV4Config
		return ret
	}
	return *o.IpV4Config.Get()
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolIpV4Block) GetIpV4ConfigOk() (*IppoolIpV4Config, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpV4Config.Get(), o.IpV4Config.IsSet()
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *IppoolIpV4Block) HasIpV4Config() bool {
	if o != nil && o.IpV4Config.IsSet() {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given NullableIppoolIpV4Config and assigns it to the IpV4Config field.
func (o *IppoolIpV4Block) SetIpV4Config(v IppoolIpV4Config) {
	o.IpV4Config.Set(&v)
}

// SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil
func (o *IppoolIpV4Block) SetIpV4ConfigNil() {
	o.IpV4Config.Set(nil)
}

// UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil
func (o *IppoolIpV4Block) UnsetIpV4Config() {
	o.IpV4Config.Unset()
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *IppoolIpV4Block) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4Block) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *IppoolIpV4Block) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *IppoolIpV4Block) SetTo(v string) {
	o.To = &v
}

func (o IppoolIpV4Block) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IppoolIpV4Block) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractBlockType, errPoolAbstractBlockType := json.Marshal(o.PoolAbstractBlockType)
	if errPoolAbstractBlockType != nil {
		return map[string]interface{}{}, errPoolAbstractBlockType
	}
	errPoolAbstractBlockType = json.Unmarshal([]byte(serializedPoolAbstractBlockType), &toSerialize)
	if errPoolAbstractBlockType != nil {
		return map[string]interface{}{}, errPoolAbstractBlockType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.From) {
		toSerialize["From"] = o.From
	}
	if o.IpV4Config.IsSet() {
		toSerialize["IpV4Config"] = o.IpV4Config.Get()
	}
	if !IsNil(o.To) {
		toSerialize["To"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IppoolIpV4Block) UnmarshalJSON(data []byte) (err error) {
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
	type IppoolIpV4BlockWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// First IPv4 address of the block.
		From       *string                  `json:"From,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		IpV4Config NullableIppoolIpV4Config `json:"IpV4Config,omitempty"`
		// Last IPv4 address of the block.
		To *string `json:"To,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	}

	varIppoolIpV4BlockWithoutEmbeddedStruct := IppoolIpV4BlockWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIppoolIpV4BlockWithoutEmbeddedStruct)
	if err == nil {
		varIppoolIpV4Block := _IppoolIpV4Block{}
		varIppoolIpV4Block.ClassId = varIppoolIpV4BlockWithoutEmbeddedStruct.ClassId
		varIppoolIpV4Block.ObjectType = varIppoolIpV4BlockWithoutEmbeddedStruct.ObjectType
		varIppoolIpV4Block.From = varIppoolIpV4BlockWithoutEmbeddedStruct.From
		varIppoolIpV4Block.IpV4Config = varIppoolIpV4BlockWithoutEmbeddedStruct.IpV4Config
		varIppoolIpV4Block.To = varIppoolIpV4BlockWithoutEmbeddedStruct.To
		*o = IppoolIpV4Block(varIppoolIpV4Block)
	} else {
		return err
	}

	varIppoolIpV4Block := _IppoolIpV4Block{}

	err = json.Unmarshal(data, &varIppoolIpV4Block)
	if err == nil {
		o.PoolAbstractBlockType = varIppoolIpV4Block.PoolAbstractBlockType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "From")
		delete(additionalProperties, "IpV4Config")
		delete(additionalProperties, "To")

		// remove fields from embedded structs
		reflectPoolAbstractBlockType := reflect.ValueOf(o.PoolAbstractBlockType)
		for i := 0; i < reflectPoolAbstractBlockType.Type().NumField(); i++ {
			t := reflectPoolAbstractBlockType.Type().Field(i)

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

type NullableIppoolIpV4Block struct {
	value *IppoolIpV4Block
	isSet bool
}

func (v NullableIppoolIpV4Block) Get() *IppoolIpV4Block {
	return v.value
}

func (v *NullableIppoolIpV4Block) Set(val *IppoolIpV4Block) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolIpV4Block) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolIpV4Block) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolIpV4Block(val *IppoolIpV4Block) *NullableIppoolIpV4Block {
	return &NullableIppoolIpV4Block{value: val, isSet: true}
}

func (v NullableIppoolIpV4Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolIpV4Block) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
