/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowTemplateEvaluationAllOf Definition of the list of properties defined in 'workflow.TemplateEvaluation', excluding properties defined in parent classes.
type WorkflowTemplateEvaluationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType             string                       `json:"ObjectType"`
	ExpectedOutputDataType NullableWorkflowBaseDataType `json:"ExpectedOutputDataType,omitempty"`
	// The output generated by the template execution.
	Output interface{} `json:"Output,omitempty"`
	// The generated template based on the stages provided in the request body.
	PreviewTemplate *string                       `json:"PreviewTemplate,omitempty"`
	Stages          []TemplateTransformationStage `json:"Stages,omitempty"`
	// Values to be fed to the template for execution.
	Values               interface{} `json:"Values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTemplateEvaluationAllOf WorkflowTemplateEvaluationAllOf

// NewWorkflowTemplateEvaluationAllOf instantiates a new WorkflowTemplateEvaluationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTemplateEvaluationAllOf(classId string, objectType string) *WorkflowTemplateEvaluationAllOf {
	this := WorkflowTemplateEvaluationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTemplateEvaluationAllOfWithDefaults instantiates a new WorkflowTemplateEvaluationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTemplateEvaluationAllOfWithDefaults() *WorkflowTemplateEvaluationAllOf {
	this := WorkflowTemplateEvaluationAllOf{}
	var classId string = "workflow.TemplateEvaluation"
	this.ClassId = classId
	var objectType string = "workflow.TemplateEvaluation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTemplateEvaluationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateEvaluationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTemplateEvaluationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTemplateEvaluationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateEvaluationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTemplateEvaluationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExpectedOutputDataType returns the ExpectedOutputDataType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateEvaluationAllOf) GetExpectedOutputDataType() WorkflowBaseDataType {
	if o == nil || o.ExpectedOutputDataType.Get() == nil {
		var ret WorkflowBaseDataType
		return ret
	}
	return *o.ExpectedOutputDataType.Get()
}

// GetExpectedOutputDataTypeOk returns a tuple with the ExpectedOutputDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateEvaluationAllOf) GetExpectedOutputDataTypeOk() (*WorkflowBaseDataType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedOutputDataType.Get(), o.ExpectedOutputDataType.IsSet()
}

// HasExpectedOutputDataType returns a boolean if a field has been set.
func (o *WorkflowTemplateEvaluationAllOf) HasExpectedOutputDataType() bool {
	if o != nil && o.ExpectedOutputDataType.IsSet() {
		return true
	}

	return false
}

// SetExpectedOutputDataType gets a reference to the given NullableWorkflowBaseDataType and assigns it to the ExpectedOutputDataType field.
func (o *WorkflowTemplateEvaluationAllOf) SetExpectedOutputDataType(v WorkflowBaseDataType) {
	o.ExpectedOutputDataType.Set(&v)
}

// SetExpectedOutputDataTypeNil sets the value for ExpectedOutputDataType to be an explicit nil
func (o *WorkflowTemplateEvaluationAllOf) SetExpectedOutputDataTypeNil() {
	o.ExpectedOutputDataType.Set(nil)
}

// UnsetExpectedOutputDataType ensures that no value is present for ExpectedOutputDataType, not even an explicit nil
func (o *WorkflowTemplateEvaluationAllOf) UnsetExpectedOutputDataType() {
	o.ExpectedOutputDataType.Unset()
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateEvaluationAllOf) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateEvaluationAllOf) GetOutputOk() (*interface{}, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowTemplateEvaluationAllOf) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given interface{} and assigns it to the Output field.
func (o *WorkflowTemplateEvaluationAllOf) SetOutput(v interface{}) {
	o.Output = v
}

// GetPreviewTemplate returns the PreviewTemplate field value if set, zero value otherwise.
func (o *WorkflowTemplateEvaluationAllOf) GetPreviewTemplate() string {
	if o == nil || o.PreviewTemplate == nil {
		var ret string
		return ret
	}
	return *o.PreviewTemplate
}

// GetPreviewTemplateOk returns a tuple with the PreviewTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateEvaluationAllOf) GetPreviewTemplateOk() (*string, bool) {
	if o == nil || o.PreviewTemplate == nil {
		return nil, false
	}
	return o.PreviewTemplate, true
}

// HasPreviewTemplate returns a boolean if a field has been set.
func (o *WorkflowTemplateEvaluationAllOf) HasPreviewTemplate() bool {
	if o != nil && o.PreviewTemplate != nil {
		return true
	}

	return false
}

// SetPreviewTemplate gets a reference to the given string and assigns it to the PreviewTemplate field.
func (o *WorkflowTemplateEvaluationAllOf) SetPreviewTemplate(v string) {
	o.PreviewTemplate = &v
}

// GetStages returns the Stages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateEvaluationAllOf) GetStages() []TemplateTransformationStage {
	if o == nil {
		var ret []TemplateTransformationStage
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateEvaluationAllOf) GetStagesOk() (*[]TemplateTransformationStage, bool) {
	if o == nil || o.Stages == nil {
		return nil, false
	}
	return &o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *WorkflowTemplateEvaluationAllOf) HasStages() bool {
	if o != nil && o.Stages != nil {
		return true
	}

	return false
}

// SetStages gets a reference to the given []TemplateTransformationStage and assigns it to the Stages field.
func (o *WorkflowTemplateEvaluationAllOf) SetStages(v []TemplateTransformationStage) {
	o.Stages = v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateEvaluationAllOf) GetValues() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateEvaluationAllOf) GetValuesOk() (*interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *WorkflowTemplateEvaluationAllOf) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given interface{} and assigns it to the Values field.
func (o *WorkflowTemplateEvaluationAllOf) SetValues(v interface{}) {
	o.Values = v
}

func (o WorkflowTemplateEvaluationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExpectedOutputDataType.IsSet() {
		toSerialize["ExpectedOutputDataType"] = o.ExpectedOutputDataType.Get()
	}
	if o.Output != nil {
		toSerialize["Output"] = o.Output
	}
	if o.PreviewTemplate != nil {
		toSerialize["PreviewTemplate"] = o.PreviewTemplate
	}
	if o.Stages != nil {
		toSerialize["Stages"] = o.Stages
	}
	if o.Values != nil {
		toSerialize["Values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTemplateEvaluationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTemplateEvaluationAllOf := _WorkflowTemplateEvaluationAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTemplateEvaluationAllOf); err == nil {
		*o = WorkflowTemplateEvaluationAllOf(varWorkflowTemplateEvaluationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExpectedOutputDataType")
		delete(additionalProperties, "Output")
		delete(additionalProperties, "PreviewTemplate")
		delete(additionalProperties, "Stages")
		delete(additionalProperties, "Values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTemplateEvaluationAllOf struct {
	value *WorkflowTemplateEvaluationAllOf
	isSet bool
}

func (v NullableWorkflowTemplateEvaluationAllOf) Get() *WorkflowTemplateEvaluationAllOf {
	return v.value
}

func (v *NullableWorkflowTemplateEvaluationAllOf) Set(val *WorkflowTemplateEvaluationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTemplateEvaluationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTemplateEvaluationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTemplateEvaluationAllOf(val *WorkflowTemplateEvaluationAllOf) *NullableWorkflowTemplateEvaluationAllOf {
	return &NullableWorkflowTemplateEvaluationAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTemplateEvaluationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTemplateEvaluationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
