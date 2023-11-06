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

// WorkflowOperationTypeDecommission Operation details for decommission actions.
type WorkflowOperationTypeDecommission struct {
	WorkflowBaseOperation
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string   `json:"ObjectType"`
	ServiceItemInstance  *MoMoRef `json:"ServiceItemInstance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowOperationTypeDecommission WorkflowOperationTypeDecommission

// NewWorkflowOperationTypeDecommission instantiates a new WorkflowOperationTypeDecommission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowOperationTypeDecommission(classId string, objectType string) *WorkflowOperationTypeDecommission {
	this := WorkflowOperationTypeDecommission{}
	this.ClassId = classId
	this.ObjectType = objectType
	var operationType string = "PostDeployment"
	this.OperationType = &operationType
	return &this
}

// NewWorkflowOperationTypeDecommissionWithDefaults instantiates a new WorkflowOperationTypeDecommission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowOperationTypeDecommissionWithDefaults() *WorkflowOperationTypeDecommission {
	this := WorkflowOperationTypeDecommission{}
	var classId string = "workflow.OperationTypeDecommission"
	this.ClassId = classId
	var objectType string = "workflow.OperationTypeDecommission"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowOperationTypeDecommission) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowOperationTypeDecommission) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowOperationTypeDecommission) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowOperationTypeDecommission) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowOperationTypeDecommission) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowOperationTypeDecommission) SetObjectType(v string) {
	o.ObjectType = v
}

// GetServiceItemInstance returns the ServiceItemInstance field value if set, zero value otherwise.
func (o *WorkflowOperationTypeDecommission) GetServiceItemInstance() MoMoRef {
	if o == nil || o.ServiceItemInstance == nil {
		var ret MoMoRef
		return ret
	}
	return *o.ServiceItemInstance
}

// GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowOperationTypeDecommission) GetServiceItemInstanceOk() (*MoMoRef, bool) {
	if o == nil || o.ServiceItemInstance == nil {
		return nil, false
	}
	return o.ServiceItemInstance, true
}

// HasServiceItemInstance returns a boolean if a field has been set.
func (o *WorkflowOperationTypeDecommission) HasServiceItemInstance() bool {
	if o != nil && o.ServiceItemInstance != nil {
		return true
	}

	return false
}

// SetServiceItemInstance gets a reference to the given MoMoRef and assigns it to the ServiceItemInstance field.
func (o *WorkflowOperationTypeDecommission) SetServiceItemInstance(v MoMoRef) {
	o.ServiceItemInstance = &v
}

func (o WorkflowOperationTypeDecommission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowBaseOperation, errWorkflowBaseOperation := json.Marshal(o.WorkflowBaseOperation)
	if errWorkflowBaseOperation != nil {
		return []byte{}, errWorkflowBaseOperation
	}
	errWorkflowBaseOperation = json.Unmarshal([]byte(serializedWorkflowBaseOperation), &toSerialize)
	if errWorkflowBaseOperation != nil {
		return []byte{}, errWorkflowBaseOperation
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ServiceItemInstance != nil {
		toSerialize["ServiceItemInstance"] = o.ServiceItemInstance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowOperationTypeDecommission) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowOperationTypeDecommissionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string   `json:"ObjectType"`
		ServiceItemInstance *MoMoRef `json:"ServiceItemInstance,omitempty"`
	}

	varWorkflowOperationTypeDecommissionWithoutEmbeddedStruct := WorkflowOperationTypeDecommissionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowOperationTypeDecommissionWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowOperationTypeDecommission := _WorkflowOperationTypeDecommission{}
		varWorkflowOperationTypeDecommission.ClassId = varWorkflowOperationTypeDecommissionWithoutEmbeddedStruct.ClassId
		varWorkflowOperationTypeDecommission.ObjectType = varWorkflowOperationTypeDecommissionWithoutEmbeddedStruct.ObjectType
		varWorkflowOperationTypeDecommission.ServiceItemInstance = varWorkflowOperationTypeDecommissionWithoutEmbeddedStruct.ServiceItemInstance
		*o = WorkflowOperationTypeDecommission(varWorkflowOperationTypeDecommission)
	} else {
		return err
	}

	varWorkflowOperationTypeDecommission := _WorkflowOperationTypeDecommission{}

	err = json.Unmarshal(bytes, &varWorkflowOperationTypeDecommission)
	if err == nil {
		o.WorkflowBaseOperation = varWorkflowOperationTypeDecommission.WorkflowBaseOperation
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ServiceItemInstance")

		// remove fields from embedded structs
		reflectWorkflowBaseOperation := reflect.ValueOf(o.WorkflowBaseOperation)
		for i := 0; i < reflectWorkflowBaseOperation.Type().NumField(); i++ {
			t := reflectWorkflowBaseOperation.Type().Field(i)

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

type NullableWorkflowOperationTypeDecommission struct {
	value *WorkflowOperationTypeDecommission
	isSet bool
}

func (v NullableWorkflowOperationTypeDecommission) Get() *WorkflowOperationTypeDecommission {
	return v.value
}

func (v *NullableWorkflowOperationTypeDecommission) Set(val *WorkflowOperationTypeDecommission) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowOperationTypeDecommission) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowOperationTypeDecommission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowOperationTypeDecommission(val *WorkflowOperationTypeDecommission) *NullableWorkflowOperationTypeDecommission {
	return &NullableWorkflowOperationTypeDecommission{value: val, isSet: true}
}

func (v NullableWorkflowOperationTypeDecommission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowOperationTypeDecommission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
