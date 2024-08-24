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

// checks if the WorkflowCustomDataTypeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowCustomDataTypeProperties{}

// WorkflowCustomDataTypeProperties Properties for a custom data type definition.
type WorkflowCustomDataTypeProperties struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When set to false custom data type is not cloneable. It is set to true only if data type is not internal and it is not using any internal custom data type.
	Cloneable *bool `json:"Cloneable,omitempty"`
	// When set to false the custom data type is owned by the system and used for internal services. Such custom data type cannot be directly used by external entities.
	ExternalMeta         *bool `json:"ExternalMeta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCustomDataTypeProperties WorkflowCustomDataTypeProperties

// NewWorkflowCustomDataTypeProperties instantiates a new WorkflowCustomDataTypeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCustomDataTypeProperties(classId string, objectType string) *WorkflowCustomDataTypeProperties {
	this := WorkflowCustomDataTypeProperties{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowCustomDataTypePropertiesWithDefaults instantiates a new WorkflowCustomDataTypeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCustomDataTypePropertiesWithDefaults() *WorkflowCustomDataTypeProperties {
	this := WorkflowCustomDataTypeProperties{}
	var classId string = "workflow.CustomDataTypeProperties"
	this.ClassId = classId
	var objectType string = "workflow.CustomDataTypeProperties"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCustomDataTypeProperties) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeProperties) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCustomDataTypeProperties) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.CustomDataTypeProperties" of the ClassId field.
func (o *WorkflowCustomDataTypeProperties) GetDefaultClassId() interface{} {
	return "workflow.CustomDataTypeProperties"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCustomDataTypeProperties) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeProperties) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCustomDataTypeProperties) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.CustomDataTypeProperties" of the ObjectType field.
func (o *WorkflowCustomDataTypeProperties) GetDefaultObjectType() interface{} {
	return "workflow.CustomDataTypeProperties"
}

// GetCloneable returns the Cloneable field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeProperties) GetCloneable() bool {
	if o == nil || IsNil(o.Cloneable) {
		var ret bool
		return ret
	}
	return *o.Cloneable
}

// GetCloneableOk returns a tuple with the Cloneable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeProperties) GetCloneableOk() (*bool, bool) {
	if o == nil || IsNil(o.Cloneable) {
		return nil, false
	}
	return o.Cloneable, true
}

// HasCloneable returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeProperties) HasCloneable() bool {
	if o != nil && !IsNil(o.Cloneable) {
		return true
	}

	return false
}

// SetCloneable gets a reference to the given bool and assigns it to the Cloneable field.
func (o *WorkflowCustomDataTypeProperties) SetCloneable(v bool) {
	o.Cloneable = &v
}

// GetExternalMeta returns the ExternalMeta field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeProperties) GetExternalMeta() bool {
	if o == nil || IsNil(o.ExternalMeta) {
		var ret bool
		return ret
	}
	return *o.ExternalMeta
}

// GetExternalMetaOk returns a tuple with the ExternalMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeProperties) GetExternalMetaOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalMeta) {
		return nil, false
	}
	return o.ExternalMeta, true
}

// HasExternalMeta returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeProperties) HasExternalMeta() bool {
	if o != nil && !IsNil(o.ExternalMeta) {
		return true
	}

	return false
}

// SetExternalMeta gets a reference to the given bool and assigns it to the ExternalMeta field.
func (o *WorkflowCustomDataTypeProperties) SetExternalMeta(v bool) {
	o.ExternalMeta = &v
}

func (o WorkflowCustomDataTypeProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowCustomDataTypeProperties) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Cloneable) {
		toSerialize["Cloneable"] = o.Cloneable
	}
	if !IsNil(o.ExternalMeta) {
		toSerialize["ExternalMeta"] = o.ExternalMeta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowCustomDataTypeProperties) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowCustomDataTypePropertiesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// When set to false custom data type is not cloneable. It is set to true only if data type is not internal and it is not using any internal custom data type.
		Cloneable *bool `json:"Cloneable,omitempty"`
		// When set to false the custom data type is owned by the system and used for internal services. Such custom data type cannot be directly used by external entities.
		ExternalMeta *bool `json:"ExternalMeta,omitempty"`
	}

	varWorkflowCustomDataTypePropertiesWithoutEmbeddedStruct := WorkflowCustomDataTypePropertiesWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowCustomDataTypePropertiesWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowCustomDataTypeProperties := _WorkflowCustomDataTypeProperties{}
		varWorkflowCustomDataTypeProperties.ClassId = varWorkflowCustomDataTypePropertiesWithoutEmbeddedStruct.ClassId
		varWorkflowCustomDataTypeProperties.ObjectType = varWorkflowCustomDataTypePropertiesWithoutEmbeddedStruct.ObjectType
		varWorkflowCustomDataTypeProperties.Cloneable = varWorkflowCustomDataTypePropertiesWithoutEmbeddedStruct.Cloneable
		varWorkflowCustomDataTypeProperties.ExternalMeta = varWorkflowCustomDataTypePropertiesWithoutEmbeddedStruct.ExternalMeta
		*o = WorkflowCustomDataTypeProperties(varWorkflowCustomDataTypeProperties)
	} else {
		return err
	}

	varWorkflowCustomDataTypeProperties := _WorkflowCustomDataTypeProperties{}

	err = json.Unmarshal(data, &varWorkflowCustomDataTypeProperties)
	if err == nil {
		o.MoBaseComplexType = varWorkflowCustomDataTypeProperties.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cloneable")
		delete(additionalProperties, "ExternalMeta")

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

type NullableWorkflowCustomDataTypeProperties struct {
	value *WorkflowCustomDataTypeProperties
	isSet bool
}

func (v NullableWorkflowCustomDataTypeProperties) Get() *WorkflowCustomDataTypeProperties {
	return v.value
}

func (v *NullableWorkflowCustomDataTypeProperties) Set(val *WorkflowCustomDataTypeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCustomDataTypeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCustomDataTypeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCustomDataTypeProperties(val *WorkflowCustomDataTypeProperties) *NullableWorkflowCustomDataTypeProperties {
	return &NullableWorkflowCustomDataTypeProperties{value: val, isSet: true}
}

func (v NullableWorkflowCustomDataTypeProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCustomDataTypeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
