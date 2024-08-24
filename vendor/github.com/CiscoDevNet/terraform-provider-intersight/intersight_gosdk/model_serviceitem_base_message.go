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
	"time"
)

// checks if the ServiceitemBaseMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceitemBaseMessage{}

// ServiceitemBaseMessage BaseMessage holds the common set of properties that are logged across CatalogServiceRequest, and ServiceItemActionInstance. The messages can be of different severity levels and hence messages are created for successful or unsuccessful executions.
type ServiceitemBaseMessage struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The type of action instance operation, such as Validate, Start, Retry, RetryFailed, Cancel, Stop etc. * `None` - No action is set, this is the default value for action field. * `Validate` - Validate the action instance inputs and run the validation workflows. * `Start` - Start a new execution of the action instance. * `Rerun` - Rerun the service item action instance from the beginning. * `Retry` - Retry the workflow that has failed from the failure point. * `Cancel` - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * `Stop` - Stop the action instance which is in progress and didn't complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped.
	ActionOperation *string `json:"ActionOperation,omitempty"`
	// The timestamp when the message was created.
	CreateTime *time.Time `json:"CreateTime,omitempty"`
	// An i18n message which can be localized and conveys status on the action.
	Message *string `json:"Message,omitempty"`
	// The severity of the message such as error, warning, info etc. * `Info` - The enum represents the log level to be used to convey info message. * `Warning` - The enum represents the log level to be used to convey warning message. * `Debug` - The enum represents the log level to be used to convey debug message. * `Error` - The enum represents the log level to be used to convey error message.
	Severity             *string `json:"Severity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceitemBaseMessage ServiceitemBaseMessage

// NewServiceitemBaseMessage instantiates a new ServiceitemBaseMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceitemBaseMessage(classId string, objectType string) *ServiceitemBaseMessage {
	this := ServiceitemBaseMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServiceitemBaseMessageWithDefaults instantiates a new ServiceitemBaseMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceitemBaseMessageWithDefaults() *ServiceitemBaseMessage {
	this := ServiceitemBaseMessage{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServiceitemBaseMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServiceitemBaseMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServiceitemBaseMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServiceitemBaseMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServiceitemBaseMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServiceitemBaseMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActionOperation returns the ActionOperation field value if set, zero value otherwise.
func (o *ServiceitemBaseMessage) GetActionOperation() string {
	if o == nil || IsNil(o.ActionOperation) {
		var ret string
		return ret
	}
	return *o.ActionOperation
}

// GetActionOperationOk returns a tuple with the ActionOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceitemBaseMessage) GetActionOperationOk() (*string, bool) {
	if o == nil || IsNil(o.ActionOperation) {
		return nil, false
	}
	return o.ActionOperation, true
}

// HasActionOperation returns a boolean if a field has been set.
func (o *ServiceitemBaseMessage) HasActionOperation() bool {
	if o != nil && !IsNil(o.ActionOperation) {
		return true
	}

	return false
}

// SetActionOperation gets a reference to the given string and assigns it to the ActionOperation field.
func (o *ServiceitemBaseMessage) SetActionOperation(v string) {
	o.ActionOperation = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ServiceitemBaseMessage) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceitemBaseMessage) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ServiceitemBaseMessage) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *ServiceitemBaseMessage) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ServiceitemBaseMessage) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceitemBaseMessage) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ServiceitemBaseMessage) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ServiceitemBaseMessage) SetMessage(v string) {
	o.Message = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ServiceitemBaseMessage) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceitemBaseMessage) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ServiceitemBaseMessage) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ServiceitemBaseMessage) SetSeverity(v string) {
	o.Severity = &v
}

func (o ServiceitemBaseMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceitemBaseMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ActionOperation) {
		toSerialize["ActionOperation"] = o.ActionOperation
	}
	if !IsNil(o.CreateTime) {
		toSerialize["CreateTime"] = o.CreateTime
	}
	if !IsNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !IsNil(o.Severity) {
		toSerialize["Severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceitemBaseMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type ServiceitemBaseMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The type of action instance operation, such as Validate, Start, Retry, RetryFailed, Cancel, Stop etc. * `None` - No action is set, this is the default value for action field. * `Validate` - Validate the action instance inputs and run the validation workflows. * `Start` - Start a new execution of the action instance. * `Rerun` - Rerun the service item action instance from the beginning. * `Retry` - Retry the workflow that has failed from the failure point. * `Cancel` - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * `Stop` - Stop the action instance which is in progress and didn't complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped.
		ActionOperation *string `json:"ActionOperation,omitempty"`
		// The timestamp when the message was created.
		CreateTime *time.Time `json:"CreateTime,omitempty"`
		// An i18n message which can be localized and conveys status on the action.
		Message *string `json:"Message,omitempty"`
		// The severity of the message such as error, warning, info etc. * `Info` - The enum represents the log level to be used to convey info message. * `Warning` - The enum represents the log level to be used to convey warning message. * `Debug` - The enum represents the log level to be used to convey debug message. * `Error` - The enum represents the log level to be used to convey error message.
		Severity *string `json:"Severity,omitempty"`
	}

	varServiceitemBaseMessageWithoutEmbeddedStruct := ServiceitemBaseMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varServiceitemBaseMessageWithoutEmbeddedStruct)
	if err == nil {
		varServiceitemBaseMessage := _ServiceitemBaseMessage{}
		varServiceitemBaseMessage.ClassId = varServiceitemBaseMessageWithoutEmbeddedStruct.ClassId
		varServiceitemBaseMessage.ObjectType = varServiceitemBaseMessageWithoutEmbeddedStruct.ObjectType
		varServiceitemBaseMessage.ActionOperation = varServiceitemBaseMessageWithoutEmbeddedStruct.ActionOperation
		varServiceitemBaseMessage.CreateTime = varServiceitemBaseMessageWithoutEmbeddedStruct.CreateTime
		varServiceitemBaseMessage.Message = varServiceitemBaseMessageWithoutEmbeddedStruct.Message
		varServiceitemBaseMessage.Severity = varServiceitemBaseMessageWithoutEmbeddedStruct.Severity
		*o = ServiceitemBaseMessage(varServiceitemBaseMessage)
	} else {
		return err
	}

	varServiceitemBaseMessage := _ServiceitemBaseMessage{}

	err = json.Unmarshal(data, &varServiceitemBaseMessage)
	if err == nil {
		o.MoBaseComplexType = varServiceitemBaseMessage.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActionOperation")
		delete(additionalProperties, "CreateTime")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Severity")

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

type NullableServiceitemBaseMessage struct {
	value *ServiceitemBaseMessage
	isSet bool
}

func (v NullableServiceitemBaseMessage) Get() *ServiceitemBaseMessage {
	return v.value
}

func (v *NullableServiceitemBaseMessage) Set(val *ServiceitemBaseMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceitemBaseMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceitemBaseMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceitemBaseMessage(val *ServiceitemBaseMessage) *NullableServiceitemBaseMessage {
	return &NullableServiceitemBaseMessage{value: val, isSet: true}
}

func (v NullableServiceitemBaseMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceitemBaseMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
