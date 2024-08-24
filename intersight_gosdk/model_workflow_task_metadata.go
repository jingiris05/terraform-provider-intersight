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

// checks if the WorkflowTaskMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowTaskMetadata{}

// WorkflowTaskMetadata Task details is a collection of properties that are common across all the versions of a task.
type WorkflowTaskMetadata struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The task metadata description to describe what this task will do when executed.
	Description *string `json:"Description,omitempty"`
	// A user friendly short name to identify the task metadata.
	Label *string `json:"Label,omitempty" validate:"regexp=^[a-zA-Z0-9]{1}[\\\\sa-zA-Z0-9_'.:-]{0,92}$"`
	// The name of the task metadata. The name should follow this convention <Verb or Action><Category><Vendor><Product><Noun or object> Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \"Generic\" if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume.
	Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9]{1}[a-zA-Z0-9_.:-]{0,63}$"`
	// An array of relationships to iamRole resources.
	AssociatedRoles      []IamRoleRelationship               `json:"AssociatedRoles,omitempty"`
	Catalog              NullableWorkflowCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskMetadata WorkflowTaskMetadata

// NewWorkflowTaskMetadata instantiates a new WorkflowTaskMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskMetadata(classId string, objectType string) *WorkflowTaskMetadata {
	this := WorkflowTaskMetadata{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTaskMetadataWithDefaults instantiates a new WorkflowTaskMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskMetadataWithDefaults() *WorkflowTaskMetadata {
	this := WorkflowTaskMetadata{}
	var classId string = "workflow.TaskMetadata"
	this.ClassId = classId
	var objectType string = "workflow.TaskMetadata"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTaskMetadata) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskMetadata) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTaskMetadata) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.TaskMetadata" of the ClassId field.
func (o *WorkflowTaskMetadata) GetDefaultClassId() interface{} {
	return "workflow.TaskMetadata"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTaskMetadata) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskMetadata) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTaskMetadata) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.TaskMetadata" of the ObjectType field.
func (o *WorkflowTaskMetadata) GetDefaultObjectType() interface{} {
	return "workflow.TaskMetadata"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowTaskMetadata) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowTaskMetadata) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowTaskMetadata) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowTaskMetadata) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskMetadata) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowTaskMetadata) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowTaskMetadata) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowTaskMetadata) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskMetadata) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowTaskMetadata) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowTaskMetadata) SetName(v string) {
	o.Name = &v
}

// GetAssociatedRoles returns the AssociatedRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskMetadata) GetAssociatedRoles() []IamRoleRelationship {
	if o == nil {
		var ret []IamRoleRelationship
		return ret
	}
	return o.AssociatedRoles
}

// GetAssociatedRolesOk returns a tuple with the AssociatedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskMetadata) GetAssociatedRolesOk() ([]IamRoleRelationship, bool) {
	if o == nil || IsNil(o.AssociatedRoles) {
		return nil, false
	}
	return o.AssociatedRoles, true
}

// HasAssociatedRoles returns a boolean if a field has been set.
func (o *WorkflowTaskMetadata) HasAssociatedRoles() bool {
	if o != nil && !IsNil(o.AssociatedRoles) {
		return true
	}

	return false
}

// SetAssociatedRoles gets a reference to the given []IamRoleRelationship and assigns it to the AssociatedRoles field.
func (o *WorkflowTaskMetadata) SetAssociatedRoles(v []IamRoleRelationship) {
	o.AssociatedRoles = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTaskMetadata) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || IsNil(o.Catalog.Get()) {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog.Get()
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTaskMetadata) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Catalog.Get(), o.Catalog.IsSet()
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowTaskMetadata) HasCatalog() bool {
	if o != nil && o.Catalog.IsSet() {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given NullableWorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowTaskMetadata) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog.Set(&v)
}

// SetCatalogNil sets the value for Catalog to be an explicit nil
func (o *WorkflowTaskMetadata) SetCatalogNil() {
	o.Catalog.Set(nil)
}

// UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
func (o *WorkflowTaskMetadata) UnsetCatalog() {
	o.Catalog.Unset()
}

func (o WorkflowTaskMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowTaskMetadata) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Label) {
		toSerialize["Label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.AssociatedRoles != nil {
		toSerialize["AssociatedRoles"] = o.AssociatedRoles
	}
	if o.Catalog.IsSet() {
		toSerialize["Catalog"] = o.Catalog.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowTaskMetadata) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowTaskMetadataWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The task metadata description to describe what this task will do when executed.
		Description *string `json:"Description,omitempty"`
		// A user friendly short name to identify the task metadata.
		Label *string `json:"Label,omitempty" validate:"regexp=^[a-zA-Z0-9]{1}[\\\\sa-zA-Z0-9_'.:-]{0,92}$"`
		// The name of the task metadata. The name should follow this convention <Verb or Action><Category><Vendor><Product><Noun or object> Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \"Generic\" if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume.
		Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9]{1}[a-zA-Z0-9_.:-]{0,63}$"`
		// An array of relationships to iamRole resources.
		AssociatedRoles []IamRoleRelationship               `json:"AssociatedRoles,omitempty"`
		Catalog         NullableWorkflowCatalogRelationship `json:"Catalog,omitempty"`
	}

	varWorkflowTaskMetadataWithoutEmbeddedStruct := WorkflowTaskMetadataWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowTaskMetadataWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowTaskMetadata := _WorkflowTaskMetadata{}
		varWorkflowTaskMetadata.ClassId = varWorkflowTaskMetadataWithoutEmbeddedStruct.ClassId
		varWorkflowTaskMetadata.ObjectType = varWorkflowTaskMetadataWithoutEmbeddedStruct.ObjectType
		varWorkflowTaskMetadata.Description = varWorkflowTaskMetadataWithoutEmbeddedStruct.Description
		varWorkflowTaskMetadata.Label = varWorkflowTaskMetadataWithoutEmbeddedStruct.Label
		varWorkflowTaskMetadata.Name = varWorkflowTaskMetadataWithoutEmbeddedStruct.Name
		varWorkflowTaskMetadata.AssociatedRoles = varWorkflowTaskMetadataWithoutEmbeddedStruct.AssociatedRoles
		varWorkflowTaskMetadata.Catalog = varWorkflowTaskMetadataWithoutEmbeddedStruct.Catalog
		*o = WorkflowTaskMetadata(varWorkflowTaskMetadata)
	} else {
		return err
	}

	varWorkflowTaskMetadata := _WorkflowTaskMetadata{}

	err = json.Unmarshal(data, &varWorkflowTaskMetadata)
	if err == nil {
		o.MoBaseMo = varWorkflowTaskMetadata.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "AssociatedRoles")
		delete(additionalProperties, "Catalog")

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

type NullableWorkflowTaskMetadata struct {
	value *WorkflowTaskMetadata
	isSet bool
}

func (v NullableWorkflowTaskMetadata) Get() *WorkflowTaskMetadata {
	return v.value
}

func (v *NullableWorkflowTaskMetadata) Set(val *WorkflowTaskMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskMetadata(val *WorkflowTaskMetadata) *NullableWorkflowTaskMetadata {
	return &NullableWorkflowTaskMetadata{value: val, isSet: true}
}

func (v NullableWorkflowTaskMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
