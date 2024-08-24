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

// checks if the WorkflowServiceItemAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowServiceItemAttribute{}

// WorkflowServiceItemAttribute Service item attribute which represents all the artifacts created or related to this service item instance.
type WorkflowServiceItemAttribute struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Service item attribute for a service item instance and the format is specified by attribute definition of the service item definition.
	Attribute interface{} `json:"Attribute,omitempty"`
	// Datatype, if any, backing the service item attribute definition.
	Datatype *string `json:"Datatype,omitempty"`
	// Attribute name which is used in the attribute definition of the service item.
	Name *string `json:"Name,omitempty"`
	// Type of the service item attribute. * `None` - Default value if the service item attribute does not belong to any of the existing types. * `Configuration` - The service item attribute is a configuration from the designer or the end user. * `Inventory` - The service item attribute captures the inventory of the resource created by the service item deployment. * `Health` - The service item attribute describes the health of the resource created by the service item deployment. * `Output` - The service item attribute captures the artifact generated after performing an action on the service item.
	Type                  *string                                           `json:"Type,omitempty"`
	ServiceItemDefinition NullableWorkflowServiceItemDefinitionRelationship `json:"ServiceItemDefinition,omitempty"`
	ServiceItemInstance   NullableWorkflowServiceItemInstanceRelationship   `json:"ServiceItemInstance,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _WorkflowServiceItemAttribute WorkflowServiceItemAttribute

// NewWorkflowServiceItemAttribute instantiates a new WorkflowServiceItemAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemAttribute(classId string, objectType string) *WorkflowServiceItemAttribute {
	this := WorkflowServiceItemAttribute{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowServiceItemAttributeWithDefaults instantiates a new WorkflowServiceItemAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemAttributeWithDefaults() *WorkflowServiceItemAttribute {
	this := WorkflowServiceItemAttribute{}
	var classId string = "workflow.ServiceItemAttribute"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemAttribute"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemAttribute) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemAttribute) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemAttribute) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.ServiceItemAttribute" of the ClassId field.
func (o *WorkflowServiceItemAttribute) GetDefaultClassId() interface{} {
	return "workflow.ServiceItemAttribute"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemAttribute) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemAttribute) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemAttribute) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.ServiceItemAttribute" of the ObjectType field.
func (o *WorkflowServiceItemAttribute) GetDefaultObjectType() interface{} {
	return "workflow.ServiceItemAttribute"
}

// GetAttribute returns the Attribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemAttribute) GetAttribute() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemAttribute) GetAttributeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return &o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *WorkflowServiceItemAttribute) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given interface{} and assigns it to the Attribute field.
func (o *WorkflowServiceItemAttribute) SetAttribute(v interface{}) {
	o.Attribute = v
}

// GetDatatype returns the Datatype field value if set, zero value otherwise.
func (o *WorkflowServiceItemAttribute) GetDatatype() string {
	if o == nil || IsNil(o.Datatype) {
		var ret string
		return ret
	}
	return *o.Datatype
}

// GetDatatypeOk returns a tuple with the Datatype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemAttribute) GetDatatypeOk() (*string, bool) {
	if o == nil || IsNil(o.Datatype) {
		return nil, false
	}
	return o.Datatype, true
}

// HasDatatype returns a boolean if a field has been set.
func (o *WorkflowServiceItemAttribute) HasDatatype() bool {
	if o != nil && !IsNil(o.Datatype) {
		return true
	}

	return false
}

// SetDatatype gets a reference to the given string and assigns it to the Datatype field.
func (o *WorkflowServiceItemAttribute) SetDatatype(v string) {
	o.Datatype = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowServiceItemAttribute) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemAttribute) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowServiceItemAttribute) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowServiceItemAttribute) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowServiceItemAttribute) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemAttribute) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowServiceItemAttribute) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowServiceItemAttribute) SetType(v string) {
	o.Type = &v
}

// GetServiceItemDefinition returns the ServiceItemDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemAttribute) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship {
	if o == nil || IsNil(o.ServiceItemDefinition.Get()) {
		var ret WorkflowServiceItemDefinitionRelationship
		return ret
	}
	return *o.ServiceItemDefinition.Get()
}

// GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemAttribute) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceItemDefinition.Get(), o.ServiceItemDefinition.IsSet()
}

// HasServiceItemDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemAttribute) HasServiceItemDefinition() bool {
	if o != nil && o.ServiceItemDefinition.IsSet() {
		return true
	}

	return false
}

// SetServiceItemDefinition gets a reference to the given NullableWorkflowServiceItemDefinitionRelationship and assigns it to the ServiceItemDefinition field.
func (o *WorkflowServiceItemAttribute) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship) {
	o.ServiceItemDefinition.Set(&v)
}

// SetServiceItemDefinitionNil sets the value for ServiceItemDefinition to be an explicit nil
func (o *WorkflowServiceItemAttribute) SetServiceItemDefinitionNil() {
	o.ServiceItemDefinition.Set(nil)
}

// UnsetServiceItemDefinition ensures that no value is present for ServiceItemDefinition, not even an explicit nil
func (o *WorkflowServiceItemAttribute) UnsetServiceItemDefinition() {
	o.ServiceItemDefinition.Unset()
}

// GetServiceItemInstance returns the ServiceItemInstance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemAttribute) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship {
	if o == nil || IsNil(o.ServiceItemInstance.Get()) {
		var ret WorkflowServiceItemInstanceRelationship
		return ret
	}
	return *o.ServiceItemInstance.Get()
}

// GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemAttribute) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceItemInstance.Get(), o.ServiceItemInstance.IsSet()
}

// HasServiceItemInstance returns a boolean if a field has been set.
func (o *WorkflowServiceItemAttribute) HasServiceItemInstance() bool {
	if o != nil && o.ServiceItemInstance.IsSet() {
		return true
	}

	return false
}

// SetServiceItemInstance gets a reference to the given NullableWorkflowServiceItemInstanceRelationship and assigns it to the ServiceItemInstance field.
func (o *WorkflowServiceItemAttribute) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship) {
	o.ServiceItemInstance.Set(&v)
}

// SetServiceItemInstanceNil sets the value for ServiceItemInstance to be an explicit nil
func (o *WorkflowServiceItemAttribute) SetServiceItemInstanceNil() {
	o.ServiceItemInstance.Set(nil)
}

// UnsetServiceItemInstance ensures that no value is present for ServiceItemInstance, not even an explicit nil
func (o *WorkflowServiceItemAttribute) UnsetServiceItemInstance() {
	o.ServiceItemInstance.Unset()
}

func (o WorkflowServiceItemAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowServiceItemAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Attribute != nil {
		toSerialize["Attribute"] = o.Attribute
	}
	if !IsNil(o.Datatype) {
		toSerialize["Datatype"] = o.Datatype
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.ServiceItemDefinition.IsSet() {
		toSerialize["ServiceItemDefinition"] = o.ServiceItemDefinition.Get()
	}
	if o.ServiceItemInstance.IsSet() {
		toSerialize["ServiceItemInstance"] = o.ServiceItemInstance.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowServiceItemAttribute) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowServiceItemAttributeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Service item attribute for a service item instance and the format is specified by attribute definition of the service item definition.
		Attribute interface{} `json:"Attribute,omitempty"`
		// Datatype, if any, backing the service item attribute definition.
		Datatype *string `json:"Datatype,omitempty"`
		// Attribute name which is used in the attribute definition of the service item.
		Name *string `json:"Name,omitempty"`
		// Type of the service item attribute. * `None` - Default value if the service item attribute does not belong to any of the existing types. * `Configuration` - The service item attribute is a configuration from the designer or the end user. * `Inventory` - The service item attribute captures the inventory of the resource created by the service item deployment. * `Health` - The service item attribute describes the health of the resource created by the service item deployment. * `Output` - The service item attribute captures the artifact generated after performing an action on the service item.
		Type                  *string                                           `json:"Type,omitempty"`
		ServiceItemDefinition NullableWorkflowServiceItemDefinitionRelationship `json:"ServiceItemDefinition,omitempty"`
		ServiceItemInstance   NullableWorkflowServiceItemInstanceRelationship   `json:"ServiceItemInstance,omitempty"`
	}

	varWorkflowServiceItemAttributeWithoutEmbeddedStruct := WorkflowServiceItemAttributeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowServiceItemAttributeWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowServiceItemAttribute := _WorkflowServiceItemAttribute{}
		varWorkflowServiceItemAttribute.ClassId = varWorkflowServiceItemAttributeWithoutEmbeddedStruct.ClassId
		varWorkflowServiceItemAttribute.ObjectType = varWorkflowServiceItemAttributeWithoutEmbeddedStruct.ObjectType
		varWorkflowServiceItemAttribute.Attribute = varWorkflowServiceItemAttributeWithoutEmbeddedStruct.Attribute
		varWorkflowServiceItemAttribute.Datatype = varWorkflowServiceItemAttributeWithoutEmbeddedStruct.Datatype
		varWorkflowServiceItemAttribute.Name = varWorkflowServiceItemAttributeWithoutEmbeddedStruct.Name
		varWorkflowServiceItemAttribute.Type = varWorkflowServiceItemAttributeWithoutEmbeddedStruct.Type
		varWorkflowServiceItemAttribute.ServiceItemDefinition = varWorkflowServiceItemAttributeWithoutEmbeddedStruct.ServiceItemDefinition
		varWorkflowServiceItemAttribute.ServiceItemInstance = varWorkflowServiceItemAttributeWithoutEmbeddedStruct.ServiceItemInstance
		*o = WorkflowServiceItemAttribute(varWorkflowServiceItemAttribute)
	} else {
		return err
	}

	varWorkflowServiceItemAttribute := _WorkflowServiceItemAttribute{}

	err = json.Unmarshal(data, &varWorkflowServiceItemAttribute)
	if err == nil {
		o.MoBaseMo = varWorkflowServiceItemAttribute.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Attribute")
		delete(additionalProperties, "Datatype")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "ServiceItemDefinition")
		delete(additionalProperties, "ServiceItemInstance")

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

type NullableWorkflowServiceItemAttribute struct {
	value *WorkflowServiceItemAttribute
	isSet bool
}

func (v NullableWorkflowServiceItemAttribute) Get() *WorkflowServiceItemAttribute {
	return v.value
}

func (v *NullableWorkflowServiceItemAttribute) Set(val *WorkflowServiceItemAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemAttribute(val *WorkflowServiceItemAttribute) *NullableWorkflowServiceItemAttribute {
	return &NullableWorkflowServiceItemAttribute{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
