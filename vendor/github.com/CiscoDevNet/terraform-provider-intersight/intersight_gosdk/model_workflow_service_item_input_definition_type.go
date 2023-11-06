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
	"reflect"
	"strings"
)

// WorkflowServiceItemInputDefinitionType Definition to capture the inputs for a service items included inside a catalog. It captures the input definition and the associated input parameter mappings for the service items.
type WorkflowServiceItemInputDefinitionType struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the service item action.
	ActionName      *string                `json:"ActionName,omitempty"`
	InputDefinition []WorkflowBaseDataType `json:"InputDefinition,omitempty"`
	// Capture the mapping of service item ActionDefinition inputDefinition to catalog service request.
	InputParameters interface{} `json:"InputParameters,omitempty"`
	// Type of action operation to be executed on the service item. * `PostDeployment` - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item. * `Deployment` - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment. * `Decommission` - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item.
	OperationType               *string  `json:"OperationType,omitempty"`
	ServiceItemActionDefinition *MoMoRef `json:"ServiceItemActionDefinition,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _WorkflowServiceItemInputDefinitionType WorkflowServiceItemInputDefinitionType

// NewWorkflowServiceItemInputDefinitionType instantiates a new WorkflowServiceItemInputDefinitionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemInputDefinitionType(classId string, objectType string) *WorkflowServiceItemInputDefinitionType {
	this := WorkflowServiceItemInputDefinitionType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowServiceItemInputDefinitionTypeWithDefaults instantiates a new WorkflowServiceItemInputDefinitionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemInputDefinitionTypeWithDefaults() *WorkflowServiceItemInputDefinitionType {
	this := WorkflowServiceItemInputDefinitionType{}
	var classId string = "workflow.ServiceItemInputDefinitionType"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemInputDefinitionType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemInputDefinitionType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInputDefinitionType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemInputDefinitionType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemInputDefinitionType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInputDefinitionType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemInputDefinitionType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActionName returns the ActionName field value if set, zero value otherwise.
func (o *WorkflowServiceItemInputDefinitionType) GetActionName() string {
	if o == nil || o.ActionName == nil {
		var ret string
		return ret
	}
	return *o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInputDefinitionType) GetActionNameOk() (*string, bool) {
	if o == nil || o.ActionName == nil {
		return nil, false
	}
	return o.ActionName, true
}

// HasActionName returns a boolean if a field has been set.
func (o *WorkflowServiceItemInputDefinitionType) HasActionName() bool {
	if o != nil && o.ActionName != nil {
		return true
	}

	return false
}

// SetActionName gets a reference to the given string and assigns it to the ActionName field.
func (o *WorkflowServiceItemInputDefinitionType) SetActionName(v string) {
	o.ActionName = &v
}

// GetInputDefinition returns the InputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemInputDefinitionType) GetInputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.InputDefinition
}

// GetInputDefinitionOk returns a tuple with the InputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemInputDefinitionType) GetInputDefinitionOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || o.InputDefinition == nil {
		return nil, false
	}
	return o.InputDefinition, true
}

// HasInputDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemInputDefinitionType) HasInputDefinition() bool {
	if o != nil && o.InputDefinition != nil {
		return true
	}

	return false
}

// SetInputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the InputDefinition field.
func (o *WorkflowServiceItemInputDefinitionType) SetInputDefinition(v []WorkflowBaseDataType) {
	o.InputDefinition = v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemInputDefinitionType) GetInputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemInputDefinitionType) GetInputParametersOk() (*interface{}, bool) {
	if o == nil || o.InputParameters == nil {
		return nil, false
	}
	return &o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *WorkflowServiceItemInputDefinitionType) HasInputParameters() bool {
	if o != nil && o.InputParameters != nil {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given interface{} and assigns it to the InputParameters field.
func (o *WorkflowServiceItemInputDefinitionType) SetInputParameters(v interface{}) {
	o.InputParameters = v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *WorkflowServiceItemInputDefinitionType) GetOperationType() string {
	if o == nil || o.OperationType == nil {
		var ret string
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInputDefinitionType) GetOperationTypeOk() (*string, bool) {
	if o == nil || o.OperationType == nil {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *WorkflowServiceItemInputDefinitionType) HasOperationType() bool {
	if o != nil && o.OperationType != nil {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given string and assigns it to the OperationType field.
func (o *WorkflowServiceItemInputDefinitionType) SetOperationType(v string) {
	o.OperationType = &v
}

// GetServiceItemActionDefinition returns the ServiceItemActionDefinition field value if set, zero value otherwise.
func (o *WorkflowServiceItemInputDefinitionType) GetServiceItemActionDefinition() MoMoRef {
	if o == nil || o.ServiceItemActionDefinition == nil {
		var ret MoMoRef
		return ret
	}
	return *o.ServiceItemActionDefinition
}

// GetServiceItemActionDefinitionOk returns a tuple with the ServiceItemActionDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInputDefinitionType) GetServiceItemActionDefinitionOk() (*MoMoRef, bool) {
	if o == nil || o.ServiceItemActionDefinition == nil {
		return nil, false
	}
	return o.ServiceItemActionDefinition, true
}

// HasServiceItemActionDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemInputDefinitionType) HasServiceItemActionDefinition() bool {
	if o != nil && o.ServiceItemActionDefinition != nil {
		return true
	}

	return false
}

// SetServiceItemActionDefinition gets a reference to the given MoMoRef and assigns it to the ServiceItemActionDefinition field.
func (o *WorkflowServiceItemInputDefinitionType) SetServiceItemActionDefinition(v MoMoRef) {
	o.ServiceItemActionDefinition = &v
}

func (o WorkflowServiceItemInputDefinitionType) MarshalJSON() ([]byte, error) {
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
	if o.ActionName != nil {
		toSerialize["ActionName"] = o.ActionName
	}
	if o.InputDefinition != nil {
		toSerialize["InputDefinition"] = o.InputDefinition
	}
	if o.InputParameters != nil {
		toSerialize["InputParameters"] = o.InputParameters
	}
	if o.OperationType != nil {
		toSerialize["OperationType"] = o.OperationType
	}
	if o.ServiceItemActionDefinition != nil {
		toSerialize["ServiceItemActionDefinition"] = o.ServiceItemActionDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemInputDefinitionType) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the service item action.
		ActionName      *string                `json:"ActionName,omitempty"`
		InputDefinition []WorkflowBaseDataType `json:"InputDefinition,omitempty"`
		// Capture the mapping of service item ActionDefinition inputDefinition to catalog service request.
		InputParameters interface{} `json:"InputParameters,omitempty"`
		// Type of action operation to be executed on the service item. * `PostDeployment` - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item. * `Deployment` - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment. * `Decommission` - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item.
		OperationType               *string  `json:"OperationType,omitempty"`
		ServiceItemActionDefinition *MoMoRef `json:"ServiceItemActionDefinition,omitempty"`
	}

	varWorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct := WorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowServiceItemInputDefinitionType := _WorkflowServiceItemInputDefinitionType{}
		varWorkflowServiceItemInputDefinitionType.ClassId = varWorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct.ClassId
		varWorkflowServiceItemInputDefinitionType.ObjectType = varWorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct.ObjectType
		varWorkflowServiceItemInputDefinitionType.ActionName = varWorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct.ActionName
		varWorkflowServiceItemInputDefinitionType.InputDefinition = varWorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct.InputDefinition
		varWorkflowServiceItemInputDefinitionType.InputParameters = varWorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct.InputParameters
		varWorkflowServiceItemInputDefinitionType.OperationType = varWorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct.OperationType
		varWorkflowServiceItemInputDefinitionType.ServiceItemActionDefinition = varWorkflowServiceItemInputDefinitionTypeWithoutEmbeddedStruct.ServiceItemActionDefinition
		*o = WorkflowServiceItemInputDefinitionType(varWorkflowServiceItemInputDefinitionType)
	} else {
		return err
	}

	varWorkflowServiceItemInputDefinitionType := _WorkflowServiceItemInputDefinitionType{}

	err = json.Unmarshal(bytes, &varWorkflowServiceItemInputDefinitionType)
	if err == nil {
		o.MoBaseComplexType = varWorkflowServiceItemInputDefinitionType.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActionName")
		delete(additionalProperties, "InputDefinition")
		delete(additionalProperties, "InputParameters")
		delete(additionalProperties, "OperationType")
		delete(additionalProperties, "ServiceItemActionDefinition")

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

type NullableWorkflowServiceItemInputDefinitionType struct {
	value *WorkflowServiceItemInputDefinitionType
	isSet bool
}

func (v NullableWorkflowServiceItemInputDefinitionType) Get() *WorkflowServiceItemInputDefinitionType {
	return v.value
}

func (v *NullableWorkflowServiceItemInputDefinitionType) Set(val *WorkflowServiceItemInputDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemInputDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemInputDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemInputDefinitionType(val *WorkflowServiceItemInputDefinitionType) *NullableWorkflowServiceItemInputDefinitionType {
	return &NullableWorkflowServiceItemInputDefinitionType{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemInputDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemInputDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
