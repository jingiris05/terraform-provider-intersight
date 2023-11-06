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

// TemplateTransformationStageAllOf Definition of the list of properties defined in 'template.TransformationStage', excluding properties defined in parent classes.
type TemplateTransformationStageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The function to be executed.
	Function *string `json:"Function,omitempty"`
	// A collection of arguments for the function being executed.
	FunctionArguments interface{} `json:"FunctionArguments,omitempty"`
	// The unique name by which the output of this transformation stage can be accessed in further stages. Only alphanumeric characters are allowed.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateTransformationStageAllOf TemplateTransformationStageAllOf

// NewTemplateTransformationStageAllOf instantiates a new TemplateTransformationStageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateTransformationStageAllOf(classId string, objectType string) *TemplateTransformationStageAllOf {
	this := TemplateTransformationStageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTemplateTransformationStageAllOfWithDefaults instantiates a new TemplateTransformationStageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateTransformationStageAllOfWithDefaults() *TemplateTransformationStageAllOf {
	this := TemplateTransformationStageAllOf{}
	var classId string = "template.TransformationStage"
	this.ClassId = classId
	var objectType string = "template.TransformationStage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TemplateTransformationStageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TemplateTransformationStageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TemplateTransformationStageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TemplateTransformationStageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TemplateTransformationStageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TemplateTransformationStageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *TemplateTransformationStageAllOf) GetFunction() string {
	if o == nil || o.Function == nil {
		var ret string
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateTransformationStageAllOf) GetFunctionOk() (*string, bool) {
	if o == nil || o.Function == nil {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *TemplateTransformationStageAllOf) HasFunction() bool {
	if o != nil && o.Function != nil {
		return true
	}

	return false
}

// SetFunction gets a reference to the given string and assigns it to the Function field.
func (o *TemplateTransformationStageAllOf) SetFunction(v string) {
	o.Function = &v
}

// GetFunctionArguments returns the FunctionArguments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateTransformationStageAllOf) GetFunctionArguments() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FunctionArguments
}

// GetFunctionArgumentsOk returns a tuple with the FunctionArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateTransformationStageAllOf) GetFunctionArgumentsOk() (*interface{}, bool) {
	if o == nil || o.FunctionArguments == nil {
		return nil, false
	}
	return &o.FunctionArguments, true
}

// HasFunctionArguments returns a boolean if a field has been set.
func (o *TemplateTransformationStageAllOf) HasFunctionArguments() bool {
	if o != nil && o.FunctionArguments != nil {
		return true
	}

	return false
}

// SetFunctionArguments gets a reference to the given interface{} and assigns it to the FunctionArguments field.
func (o *TemplateTransformationStageAllOf) SetFunctionArguments(v interface{}) {
	o.FunctionArguments = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateTransformationStageAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateTransformationStageAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateTransformationStageAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateTransformationStageAllOf) SetName(v string) {
	o.Name = &v
}

func (o TemplateTransformationStageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Function != nil {
		toSerialize["Function"] = o.Function
	}
	if o.FunctionArguments != nil {
		toSerialize["FunctionArguments"] = o.FunctionArguments
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TemplateTransformationStageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTemplateTransformationStageAllOf := _TemplateTransformationStageAllOf{}

	if err = json.Unmarshal(bytes, &varTemplateTransformationStageAllOf); err == nil {
		*o = TemplateTransformationStageAllOf(varTemplateTransformationStageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Function")
		delete(additionalProperties, "FunctionArguments")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateTransformationStageAllOf struct {
	value *TemplateTransformationStageAllOf
	isSet bool
}

func (v NullableTemplateTransformationStageAllOf) Get() *TemplateTransformationStageAllOf {
	return v.value
}

func (v *NullableTemplateTransformationStageAllOf) Set(val *TemplateTransformationStageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateTransformationStageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateTransformationStageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateTransformationStageAllOf(val *TemplateTransformationStageAllOf) *NullableTemplateTransformationStageAllOf {
	return &NullableTemplateTransformationStageAllOf{value: val, isSet: true}
}

func (v NullableTemplateTransformationStageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateTransformationStageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
