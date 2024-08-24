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

// checks if the CondAlarmAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CondAlarmAction{}

// CondAlarmAction A rule for constructing an alarm for a detected issue.
type CondAlarmAction struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                `json:"ObjectType"`
	Message              NullableIssueMessage  `json:"Message,omitempty"`
	Spec                 NullableCondAlarmSpec `json:"Spec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CondAlarmAction CondAlarmAction

// NewCondAlarmAction instantiates a new CondAlarmAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondAlarmAction(classId string, objectType string) *CondAlarmAction {
	this := CondAlarmAction{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCondAlarmActionWithDefaults instantiates a new CondAlarmAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondAlarmActionWithDefaults() *CondAlarmAction {
	this := CondAlarmAction{}
	var classId string = "cond.AlarmAction"
	this.ClassId = classId
	var objectType string = "cond.AlarmAction"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CondAlarmAction) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CondAlarmAction) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CondAlarmAction) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "cond.AlarmAction" of the ClassId field.
func (o *CondAlarmAction) GetDefaultClassId() interface{} {
	return "cond.AlarmAction"
}

// GetObjectType returns the ObjectType field value
func (o *CondAlarmAction) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CondAlarmAction) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CondAlarmAction) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "cond.AlarmAction" of the ObjectType field.
func (o *CondAlarmAction) GetDefaultObjectType() interface{} {
	return "cond.AlarmAction"
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondAlarmAction) GetMessage() IssueMessage {
	if o == nil || IsNil(o.Message.Get()) {
		var ret IssueMessage
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondAlarmAction) GetMessageOk() (*IssueMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *CondAlarmAction) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableIssueMessage and assigns it to the Message field.
func (o *CondAlarmAction) SetMessage(v IssueMessage) {
	o.Message.Set(&v)
}

// SetMessageNil sets the value for Message to be an explicit nil
func (o *CondAlarmAction) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *CondAlarmAction) UnsetMessage() {
	o.Message.Unset()
}

// GetSpec returns the Spec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondAlarmAction) GetSpec() CondAlarmSpec {
	if o == nil || IsNil(o.Spec.Get()) {
		var ret CondAlarmSpec
		return ret
	}
	return *o.Spec.Get()
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondAlarmAction) GetSpecOk() (*CondAlarmSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Spec.Get(), o.Spec.IsSet()
}

// HasSpec returns a boolean if a field has been set.
func (o *CondAlarmAction) HasSpec() bool {
	if o != nil && o.Spec.IsSet() {
		return true
	}

	return false
}

// SetSpec gets a reference to the given NullableCondAlarmSpec and assigns it to the Spec field.
func (o *CondAlarmAction) SetSpec(v CondAlarmSpec) {
	o.Spec.Set(&v)
}

// SetSpecNil sets the value for Spec to be an explicit nil
func (o *CondAlarmAction) SetSpecNil() {
	o.Spec.Set(nil)
}

// UnsetSpec ensures that no value is present for Spec, not even an explicit nil
func (o *CondAlarmAction) UnsetSpec() {
	o.Spec.Unset()
}

func (o CondAlarmAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CondAlarmAction) ToMap() (map[string]interface{}, error) {
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
	if o.Message.IsSet() {
		toSerialize["Message"] = o.Message.Get()
	}
	if o.Spec.IsSet() {
		toSerialize["Spec"] = o.Spec.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CondAlarmAction) UnmarshalJSON(data []byte) (err error) {
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
	type CondAlarmActionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                `json:"ObjectType"`
		Message    NullableIssueMessage  `json:"Message,omitempty"`
		Spec       NullableCondAlarmSpec `json:"Spec,omitempty"`
	}

	varCondAlarmActionWithoutEmbeddedStruct := CondAlarmActionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCondAlarmActionWithoutEmbeddedStruct)
	if err == nil {
		varCondAlarmAction := _CondAlarmAction{}
		varCondAlarmAction.ClassId = varCondAlarmActionWithoutEmbeddedStruct.ClassId
		varCondAlarmAction.ObjectType = varCondAlarmActionWithoutEmbeddedStruct.ObjectType
		varCondAlarmAction.Message = varCondAlarmActionWithoutEmbeddedStruct.Message
		varCondAlarmAction.Spec = varCondAlarmActionWithoutEmbeddedStruct.Spec
		*o = CondAlarmAction(varCondAlarmAction)
	} else {
		return err
	}

	varCondAlarmAction := _CondAlarmAction{}

	err = json.Unmarshal(data, &varCondAlarmAction)
	if err == nil {
		o.MoBaseComplexType = varCondAlarmAction.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Spec")

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

type NullableCondAlarmAction struct {
	value *CondAlarmAction
	isSet bool
}

func (v NullableCondAlarmAction) Get() *CondAlarmAction {
	return v.value
}

func (v *NullableCondAlarmAction) Set(val *CondAlarmAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCondAlarmAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCondAlarmAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondAlarmAction(val *CondAlarmAction) *NullableCondAlarmAction {
	return &NullableCondAlarmAction{value: val, isSet: true}
}

func (v NullableCondAlarmAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondAlarmAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
