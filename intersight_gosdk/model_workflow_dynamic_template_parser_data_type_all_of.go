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
)

// WorkflowDynamicTemplateParserDataTypeAllOf Definition of the list of properties defined in 'workflow.DynamicTemplateParserDataType', excluding properties defined in parent classes.
type WorkflowDynamicTemplateParserDataTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When set to true, the template is marked as secure and the content is encrypted and stored.
	IsTemplateSecure *bool `json:"IsTemplateSecure,omitempty"`
	// Template type decides on the API to be used to fetch the placeholders present inside the template. * `OsInstall` - This refers to the OS configuration template MO.
	TemplateType         *string `json:"TemplateType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowDynamicTemplateParserDataTypeAllOf WorkflowDynamicTemplateParserDataTypeAllOf

// NewWorkflowDynamicTemplateParserDataTypeAllOf instantiates a new WorkflowDynamicTemplateParserDataTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDynamicTemplateParserDataTypeAllOf(classId string, objectType string) *WorkflowDynamicTemplateParserDataTypeAllOf {
	this := WorkflowDynamicTemplateParserDataTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var templateType string = "OsInstall"
	this.TemplateType = &templateType
	return &this
}

// NewWorkflowDynamicTemplateParserDataTypeAllOfWithDefaults instantiates a new WorkflowDynamicTemplateParserDataTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDynamicTemplateParserDataTypeAllOfWithDefaults() *WorkflowDynamicTemplateParserDataTypeAllOf {
	this := WorkflowDynamicTemplateParserDataTypeAllOf{}
	var classId string = "workflow.DynamicTemplateParserDataType"
	this.ClassId = classId
	var objectType string = "workflow.DynamicTemplateParserDataType"
	this.ObjectType = objectType
	var templateType string = "OsInstall"
	this.TemplateType = &templateType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsTemplateSecure returns the IsTemplateSecure field value if set, zero value otherwise.
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetIsTemplateSecure() bool {
	if o == nil || o.IsTemplateSecure == nil {
		var ret bool
		return ret
	}
	return *o.IsTemplateSecure
}

// GetIsTemplateSecureOk returns a tuple with the IsTemplateSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetIsTemplateSecureOk() (*bool, bool) {
	if o == nil || o.IsTemplateSecure == nil {
		return nil, false
	}
	return o.IsTemplateSecure, true
}

// HasIsTemplateSecure returns a boolean if a field has been set.
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) HasIsTemplateSecure() bool {
	if o != nil && o.IsTemplateSecure != nil {
		return true
	}

	return false
}

// SetIsTemplateSecure gets a reference to the given bool and assigns it to the IsTemplateSecure field.
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) SetIsTemplateSecure(v bool) {
	o.IsTemplateSecure = &v
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetTemplateType() string {
	if o == nil || o.TemplateType == nil {
		var ret string
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) GetTemplateTypeOk() (*string, bool) {
	if o == nil || o.TemplateType == nil {
		return nil, false
	}
	return o.TemplateType, true
}

// HasTemplateType returns a boolean if a field has been set.
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) HasTemplateType() bool {
	if o != nil && o.TemplateType != nil {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given string and assigns it to the TemplateType field.
func (o *WorkflowDynamicTemplateParserDataTypeAllOf) SetTemplateType(v string) {
	o.TemplateType = &v
}

func (o WorkflowDynamicTemplateParserDataTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsTemplateSecure != nil {
		toSerialize["IsTemplateSecure"] = o.IsTemplateSecure
	}
	if o.TemplateType != nil {
		toSerialize["TemplateType"] = o.TemplateType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowDynamicTemplateParserDataTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowDynamicTemplateParserDataTypeAllOf := _WorkflowDynamicTemplateParserDataTypeAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowDynamicTemplateParserDataTypeAllOf); err == nil {
		*o = WorkflowDynamicTemplateParserDataTypeAllOf(varWorkflowDynamicTemplateParserDataTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsTemplateSecure")
		delete(additionalProperties, "TemplateType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowDynamicTemplateParserDataTypeAllOf struct {
	value *WorkflowDynamicTemplateParserDataTypeAllOf
	isSet bool
}

func (v NullableWorkflowDynamicTemplateParserDataTypeAllOf) Get() *WorkflowDynamicTemplateParserDataTypeAllOf {
	return v.value
}

func (v *NullableWorkflowDynamicTemplateParserDataTypeAllOf) Set(val *WorkflowDynamicTemplateParserDataTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDynamicTemplateParserDataTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDynamicTemplateParserDataTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDynamicTemplateParserDataTypeAllOf(val *WorkflowDynamicTemplateParserDataTypeAllOf) *NullableWorkflowDynamicTemplateParserDataTypeAllOf {
	return &NullableWorkflowDynamicTemplateParserDataTypeAllOf{value: val, isSet: true}
}

func (v NullableWorkflowDynamicTemplateParserDataTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDynamicTemplateParserDataTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
