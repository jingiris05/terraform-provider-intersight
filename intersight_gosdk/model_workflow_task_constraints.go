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

// checks if the WorkflowTaskConstraints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowTaskConstraints{}

// WorkflowTaskConstraints Captures the constraints for a task. Currently, it hold only targetDataType constraints such as Vendor, ObjectType that are properties of a target device.
type WorkflowTaskConstraints struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// List of property constraints that helps to narrow down task implementations based on target device input.
	TargetDataType       interface{} `json:"TargetDataType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskConstraints WorkflowTaskConstraints

// NewWorkflowTaskConstraints instantiates a new WorkflowTaskConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskConstraints(classId string, objectType string) *WorkflowTaskConstraints {
	this := WorkflowTaskConstraints{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTaskConstraintsWithDefaults instantiates a new WorkflowTaskConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskConstraintsWithDefaults() *WorkflowTaskConstraints {
	this := WorkflowTaskConstraints{}
	var classId string = "workflow.TaskConstraints"
	this.ClassId = classId
	var objectType string = "workflow.TaskConstraints"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTaskConstraints) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskConstraints) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTaskConstraints) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.TaskConstraints" of the ClassId field.
func (o *WorkflowTaskConstraints) GetDefaultClassId() interface{} {
	return "workflow.TaskConstraints"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTaskConstraints) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskConstraints) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTaskConstraints) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.TaskConstraints" of the ObjectType field.
func (o *WorkflowTaskConstraints) GetDefaultObjectType() interface{} {
	return "workflow.TaskConstraints"
}

// GetTargetDataType returns the TargetDataType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskConstraints) GetTargetDataType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TargetDataType
}

// GetTargetDataTypeOk returns a tuple with the TargetDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskConstraints) GetTargetDataTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TargetDataType) {
		return nil, false
	}
	return &o.TargetDataType, true
}

// HasTargetDataType returns a boolean if a field has been set.
func (o *WorkflowTaskConstraints) HasTargetDataType() bool {
	if o != nil && !IsNil(o.TargetDataType) {
		return true
	}

	return false
}

// SetTargetDataType gets a reference to the given interface{} and assigns it to the TargetDataType field.
func (o *WorkflowTaskConstraints) SetTargetDataType(v interface{}) {
	o.TargetDataType = v
}

func (o WorkflowTaskConstraints) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowTaskConstraints) ToMap() (map[string]interface{}, error) {
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
	if o.TargetDataType != nil {
		toSerialize["TargetDataType"] = o.TargetDataType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowTaskConstraints) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowTaskConstraintsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// List of property constraints that helps to narrow down task implementations based on target device input.
		TargetDataType interface{} `json:"TargetDataType,omitempty"`
	}

	varWorkflowTaskConstraintsWithoutEmbeddedStruct := WorkflowTaskConstraintsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowTaskConstraintsWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowTaskConstraints := _WorkflowTaskConstraints{}
		varWorkflowTaskConstraints.ClassId = varWorkflowTaskConstraintsWithoutEmbeddedStruct.ClassId
		varWorkflowTaskConstraints.ObjectType = varWorkflowTaskConstraintsWithoutEmbeddedStruct.ObjectType
		varWorkflowTaskConstraints.TargetDataType = varWorkflowTaskConstraintsWithoutEmbeddedStruct.TargetDataType
		*o = WorkflowTaskConstraints(varWorkflowTaskConstraints)
	} else {
		return err
	}

	varWorkflowTaskConstraints := _WorkflowTaskConstraints{}

	err = json.Unmarshal(data, &varWorkflowTaskConstraints)
	if err == nil {
		o.MoBaseComplexType = varWorkflowTaskConstraints.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetDataType")

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

type NullableWorkflowTaskConstraints struct {
	value *WorkflowTaskConstraints
	isSet bool
}

func (v NullableWorkflowTaskConstraints) Get() *WorkflowTaskConstraints {
	return v.value
}

func (v *NullableWorkflowTaskConstraints) Set(val *WorkflowTaskConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskConstraints(val *WorkflowTaskConstraints) *NullableWorkflowTaskConstraints {
	return &NullableWorkflowTaskConstraints{value: val, isSet: true}
}

func (v NullableWorkflowTaskConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
