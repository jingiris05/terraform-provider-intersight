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
	"reflect"
	"strings"
)

// WorkflowCustomDataTypeDefinition Captures a customized data type definition that can be used for task or workflow input/output. This can be reused across multiple tasks and workflow definitions.
type WorkflowCustomDataTypeDefinition struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When true this data type definition is a collection of type definitions to represent composite data like JSON.
	CompositeType *bool `json:"CompositeType,omitempty"`
	// A human-friendly description of this custom data type indicating it's domain and usage.
	Description *string `json:"Description,omitempty"`
	// A user friendly short name to identify the custom data type definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), or an underscore (_) and must be at least 2 characters.
	Label *string `json:"Label,omitempty"`
	// The name of custom data type definition. The valid name can contain lower case and upper case alphabetic characters, digits and special characters '-' and '_'.
	Name                 *string                                       `json:"Name,omitempty"`
	ParameterSet         []WorkflowParameterSet                        `json:"ParameterSet,omitempty"`
	Properties           NullableWorkflowCustomDataTypeProperties      `json:"Properties,omitempty"`
	TypeDefinition       []WorkflowBaseDataType                        `json:"TypeDefinition,omitempty"`
	Catalog              *WorkflowCatalogRelationship                  `json:"Catalog,omitempty"`
	ClonedFrom           *WorkflowCustomDataTypeDefinitionRelationship `json:"ClonedFrom,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCustomDataTypeDefinition WorkflowCustomDataTypeDefinition

// NewWorkflowCustomDataTypeDefinition instantiates a new WorkflowCustomDataTypeDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCustomDataTypeDefinition(classId string, objectType string) *WorkflowCustomDataTypeDefinition {
	this := WorkflowCustomDataTypeDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var compositeType bool = false
	this.CompositeType = &compositeType
	return &this
}

// NewWorkflowCustomDataTypeDefinitionWithDefaults instantiates a new WorkflowCustomDataTypeDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCustomDataTypeDefinitionWithDefaults() *WorkflowCustomDataTypeDefinition {
	this := WorkflowCustomDataTypeDefinition{}
	var classId string = "workflow.CustomDataTypeDefinition"
	this.ClassId = classId
	var objectType string = "workflow.CustomDataTypeDefinition"
	this.ObjectType = objectType
	var compositeType bool = false
	this.CompositeType = &compositeType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCustomDataTypeDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCustomDataTypeDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCustomDataTypeDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCustomDataTypeDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCompositeType returns the CompositeType field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinition) GetCompositeType() bool {
	if o == nil || o.CompositeType == nil {
		var ret bool
		return ret
	}
	return *o.CompositeType
}

// GetCompositeTypeOk returns a tuple with the CompositeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinition) GetCompositeTypeOk() (*bool, bool) {
	if o == nil || o.CompositeType == nil {
		return nil, false
	}
	return o.CompositeType, true
}

// HasCompositeType returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinition) HasCompositeType() bool {
	if o != nil && o.CompositeType != nil {
		return true
	}

	return false
}

// SetCompositeType gets a reference to the given bool and assigns it to the CompositeType field.
func (o *WorkflowCustomDataTypeDefinition) SetCompositeType(v bool) {
	o.CompositeType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowCustomDataTypeDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinition) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinition) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinition) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowCustomDataTypeDefinition) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowCustomDataTypeDefinition) SetName(v string) {
	o.Name = &v
}

// GetParameterSet returns the ParameterSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCustomDataTypeDefinition) GetParameterSet() []WorkflowParameterSet {
	if o == nil {
		var ret []WorkflowParameterSet
		return ret
	}
	return o.ParameterSet
}

// GetParameterSetOk returns a tuple with the ParameterSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCustomDataTypeDefinition) GetParameterSetOk() (*[]WorkflowParameterSet, bool) {
	if o == nil || o.ParameterSet == nil {
		return nil, false
	}
	return &o.ParameterSet, true
}

// HasParameterSet returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinition) HasParameterSet() bool {
	if o != nil && o.ParameterSet != nil {
		return true
	}

	return false
}

// SetParameterSet gets a reference to the given []WorkflowParameterSet and assigns it to the ParameterSet field.
func (o *WorkflowCustomDataTypeDefinition) SetParameterSet(v []WorkflowParameterSet) {
	o.ParameterSet = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCustomDataTypeDefinition) GetProperties() WorkflowCustomDataTypeProperties {
	if o == nil || o.Properties.Get() == nil {
		var ret WorkflowCustomDataTypeProperties
		return ret
	}
	return *o.Properties.Get()
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCustomDataTypeDefinition) GetPropertiesOk() (*WorkflowCustomDataTypeProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Get(), o.Properties.IsSet()
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinition) HasProperties() bool {
	if o != nil && o.Properties.IsSet() {
		return true
	}

	return false
}

// SetProperties gets a reference to the given NullableWorkflowCustomDataTypeProperties and assigns it to the Properties field.
func (o *WorkflowCustomDataTypeDefinition) SetProperties(v WorkflowCustomDataTypeProperties) {
	o.Properties.Set(&v)
}

// SetPropertiesNil sets the value for Properties to be an explicit nil
func (o *WorkflowCustomDataTypeDefinition) SetPropertiesNil() {
	o.Properties.Set(nil)
}

// UnsetProperties ensures that no value is present for Properties, not even an explicit nil
func (o *WorkflowCustomDataTypeDefinition) UnsetProperties() {
	o.Properties.Unset()
}

// GetTypeDefinition returns the TypeDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCustomDataTypeDefinition) GetTypeDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.TypeDefinition
}

// GetTypeDefinitionOk returns a tuple with the TypeDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCustomDataTypeDefinition) GetTypeDefinitionOk() (*[]WorkflowBaseDataType, bool) {
	if o == nil || o.TypeDefinition == nil {
		return nil, false
	}
	return &o.TypeDefinition, true
}

// HasTypeDefinition returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinition) HasTypeDefinition() bool {
	if o != nil && o.TypeDefinition != nil {
		return true
	}

	return false
}

// SetTypeDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the TypeDefinition field.
func (o *WorkflowCustomDataTypeDefinition) SetTypeDefinition(v []WorkflowBaseDataType) {
	o.TypeDefinition = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinition) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinition) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinition) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given WorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowCustomDataTypeDefinition) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog = &v
}

// GetClonedFrom returns the ClonedFrom field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypeDefinition) GetClonedFrom() WorkflowCustomDataTypeDefinitionRelationship {
	if o == nil || o.ClonedFrom == nil {
		var ret WorkflowCustomDataTypeDefinitionRelationship
		return ret
	}
	return *o.ClonedFrom
}

// GetClonedFromOk returns a tuple with the ClonedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypeDefinition) GetClonedFromOk() (*WorkflowCustomDataTypeDefinitionRelationship, bool) {
	if o == nil || o.ClonedFrom == nil {
		return nil, false
	}
	return o.ClonedFrom, true
}

// HasClonedFrom returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypeDefinition) HasClonedFrom() bool {
	if o != nil && o.ClonedFrom != nil {
		return true
	}

	return false
}

// SetClonedFrom gets a reference to the given WorkflowCustomDataTypeDefinitionRelationship and assigns it to the ClonedFrom field.
func (o *WorkflowCustomDataTypeDefinition) SetClonedFrom(v WorkflowCustomDataTypeDefinitionRelationship) {
	o.ClonedFrom = &v
}

func (o WorkflowCustomDataTypeDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CompositeType != nil {
		toSerialize["CompositeType"] = o.CompositeType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ParameterSet != nil {
		toSerialize["ParameterSet"] = o.ParameterSet
	}
	if o.Properties.IsSet() {
		toSerialize["Properties"] = o.Properties.Get()
	}
	if o.TypeDefinition != nil {
		toSerialize["TypeDefinition"] = o.TypeDefinition
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.ClonedFrom != nil {
		toSerialize["ClonedFrom"] = o.ClonedFrom
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowCustomDataTypeDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// When true this data type definition is a collection of type definitions to represent composite data like JSON.
		CompositeType *bool `json:"CompositeType,omitempty"`
		// A human-friendly description of this custom data type indicating it's domain and usage.
		Description *string `json:"Description,omitempty"`
		// A user friendly short name to identify the custom data type definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), or an underscore (_) and must be at least 2 characters.
		Label *string `json:"Label,omitempty"`
		// The name of custom data type definition. The valid name can contain lower case and upper case alphabetic characters, digits and special characters '-' and '_'.
		Name           *string                                       `json:"Name,omitempty"`
		ParameterSet   []WorkflowParameterSet                        `json:"ParameterSet,omitempty"`
		Properties     NullableWorkflowCustomDataTypeProperties      `json:"Properties,omitempty"`
		TypeDefinition []WorkflowBaseDataType                        `json:"TypeDefinition,omitempty"`
		Catalog        *WorkflowCatalogRelationship                  `json:"Catalog,omitempty"`
		ClonedFrom     *WorkflowCustomDataTypeDefinitionRelationship `json:"ClonedFrom,omitempty"`
	}

	varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct := WorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowCustomDataTypeDefinition := _WorkflowCustomDataTypeDefinition{}
		varWorkflowCustomDataTypeDefinition.ClassId = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.ClassId
		varWorkflowCustomDataTypeDefinition.ObjectType = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.ObjectType
		varWorkflowCustomDataTypeDefinition.CompositeType = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.CompositeType
		varWorkflowCustomDataTypeDefinition.Description = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.Description
		varWorkflowCustomDataTypeDefinition.Label = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.Label
		varWorkflowCustomDataTypeDefinition.Name = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.Name
		varWorkflowCustomDataTypeDefinition.ParameterSet = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.ParameterSet
		varWorkflowCustomDataTypeDefinition.Properties = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.Properties
		varWorkflowCustomDataTypeDefinition.TypeDefinition = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.TypeDefinition
		varWorkflowCustomDataTypeDefinition.Catalog = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.Catalog
		varWorkflowCustomDataTypeDefinition.ClonedFrom = varWorkflowCustomDataTypeDefinitionWithoutEmbeddedStruct.ClonedFrom
		*o = WorkflowCustomDataTypeDefinition(varWorkflowCustomDataTypeDefinition)
	} else {
		return err
	}

	varWorkflowCustomDataTypeDefinition := _WorkflowCustomDataTypeDefinition{}

	err = json.Unmarshal(bytes, &varWorkflowCustomDataTypeDefinition)
	if err == nil {
		o.MoBaseMo = varWorkflowCustomDataTypeDefinition.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CompositeType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ParameterSet")
		delete(additionalProperties, "Properties")
		delete(additionalProperties, "TypeDefinition")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "ClonedFrom")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableWorkflowCustomDataTypeDefinition struct {
	value *WorkflowCustomDataTypeDefinition
	isSet bool
}

func (v NullableWorkflowCustomDataTypeDefinition) Get() *WorkflowCustomDataTypeDefinition {
	return v.value
}

func (v *NullableWorkflowCustomDataTypeDefinition) Set(val *WorkflowCustomDataTypeDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCustomDataTypeDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCustomDataTypeDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCustomDataTypeDefinition(val *WorkflowCustomDataTypeDefinition) *NullableWorkflowCustomDataTypeDefinition {
	return &NullableWorkflowCustomDataTypeDefinition{value: val, isSet: true}
}

func (v NullableWorkflowCustomDataTypeDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCustomDataTypeDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
