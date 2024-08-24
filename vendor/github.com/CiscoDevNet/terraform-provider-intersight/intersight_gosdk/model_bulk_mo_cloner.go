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

// checks if the BulkMoCloner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkMoCloner{}

// BulkMoCloner The MO Cloner interface facilitates making n number of shallow copies of any resource instance which supports the CREATE operation. For deep copy and reference clone support, refer to the MoDeepCloner API. The MO to be cloned must be specified as an MoRef object in the \"Sources\". The \"Targets\" array must contain n JSON documents each compliant to RFC 7386.  For each target MO to be created, you can specify the following - - new values for the identity properties, if applicable - new values for specific properties or references of the source MO which need to be overridden in the cloned object. Currently this API is used to perform template derive operations for Server Profile Templates, vNIC Templates and vHBA Templates.
type BulkMoCloner struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string           `json:"ObjectType"`
	ExcludeProperties []string         `json:"ExcludeProperties,omitempty"`
	Responses         []BulkRestResult `json:"Responses,omitempty"`
	Sources           []MoBaseMo       `json:"Sources,omitempty"`
	Targets           []MoBaseMo       `json:"Targets,omitempty"`
	// A user-friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_).
	WorkflowNameSuffix   *string                                      `json:"WorkflowNameSuffix,omitempty" validate:"regexp=^$|^[a-zA-Z0-9]{1}[\\\\sa-zA-Z0-9_.\\\\,\\/:-]{0,63}$"`
	AsyncResult          NullableBulkResultRelationship               `json:"AsyncResult,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkMoCloner BulkMoCloner

// NewBulkMoCloner instantiates a new BulkMoCloner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkMoCloner(classId string, objectType string) *BulkMoCloner {
	this := BulkMoCloner{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBulkMoClonerWithDefaults instantiates a new BulkMoCloner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkMoClonerWithDefaults() *BulkMoCloner {
	this := BulkMoCloner{}
	var classId string = "bulk.MoCloner"
	this.ClassId = classId
	var objectType string = "bulk.MoCloner"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkMoCloner) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkMoCloner) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkMoCloner) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "bulk.MoCloner" of the ClassId field.
func (o *BulkMoCloner) GetDefaultClassId() interface{} {
	return "bulk.MoCloner"
}

// GetObjectType returns the ObjectType field value
func (o *BulkMoCloner) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkMoCloner) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkMoCloner) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "bulk.MoCloner" of the ObjectType field.
func (o *BulkMoCloner) GetDefaultObjectType() interface{} {
	return "bulk.MoCloner"
}

// GetExcludeProperties returns the ExcludeProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoCloner) GetExcludeProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeProperties
}

// GetExcludePropertiesOk returns a tuple with the ExcludeProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoCloner) GetExcludePropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeProperties) {
		return nil, false
	}
	return o.ExcludeProperties, true
}

// HasExcludeProperties returns a boolean if a field has been set.
func (o *BulkMoCloner) HasExcludeProperties() bool {
	if o != nil && !IsNil(o.ExcludeProperties) {
		return true
	}

	return false
}

// SetExcludeProperties gets a reference to the given []string and assigns it to the ExcludeProperties field.
func (o *BulkMoCloner) SetExcludeProperties(v []string) {
	o.ExcludeProperties = v
}

// GetResponses returns the Responses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoCloner) GetResponses() []BulkRestResult {
	if o == nil {
		var ret []BulkRestResult
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoCloner) GetResponsesOk() ([]BulkRestResult, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *BulkMoCloner) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []BulkRestResult and assigns it to the Responses field.
func (o *BulkMoCloner) SetResponses(v []BulkRestResult) {
	o.Responses = v
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoCloner) GetSources() []MoBaseMo {
	if o == nil {
		var ret []MoBaseMo
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoCloner) GetSourcesOk() ([]MoBaseMo, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *BulkMoCloner) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []MoBaseMo and assigns it to the Sources field.
func (o *BulkMoCloner) SetSources(v []MoBaseMo) {
	o.Sources = v
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoCloner) GetTargets() []MoBaseMo {
	if o == nil {
		var ret []MoBaseMo
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoCloner) GetTargetsOk() ([]MoBaseMo, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *BulkMoCloner) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []MoBaseMo and assigns it to the Targets field.
func (o *BulkMoCloner) SetTargets(v []MoBaseMo) {
	o.Targets = v
}

// GetWorkflowNameSuffix returns the WorkflowNameSuffix field value if set, zero value otherwise.
func (o *BulkMoCloner) GetWorkflowNameSuffix() string {
	if o == nil || IsNil(o.WorkflowNameSuffix) {
		var ret string
		return ret
	}
	return *o.WorkflowNameSuffix
}

// GetWorkflowNameSuffixOk returns a tuple with the WorkflowNameSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoCloner) GetWorkflowNameSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowNameSuffix) {
		return nil, false
	}
	return o.WorkflowNameSuffix, true
}

// HasWorkflowNameSuffix returns a boolean if a field has been set.
func (o *BulkMoCloner) HasWorkflowNameSuffix() bool {
	if o != nil && !IsNil(o.WorkflowNameSuffix) {
		return true
	}

	return false
}

// SetWorkflowNameSuffix gets a reference to the given string and assigns it to the WorkflowNameSuffix field.
func (o *BulkMoCloner) SetWorkflowNameSuffix(v string) {
	o.WorkflowNameSuffix = &v
}

// GetAsyncResult returns the AsyncResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoCloner) GetAsyncResult() BulkResultRelationship {
	if o == nil || IsNil(o.AsyncResult.Get()) {
		var ret BulkResultRelationship
		return ret
	}
	return *o.AsyncResult.Get()
}

// GetAsyncResultOk returns a tuple with the AsyncResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoCloner) GetAsyncResultOk() (*BulkResultRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AsyncResult.Get(), o.AsyncResult.IsSet()
}

// HasAsyncResult returns a boolean if a field has been set.
func (o *BulkMoCloner) HasAsyncResult() bool {
	if o != nil && o.AsyncResult.IsSet() {
		return true
	}

	return false
}

// SetAsyncResult gets a reference to the given NullableBulkResultRelationship and assigns it to the AsyncResult field.
func (o *BulkMoCloner) SetAsyncResult(v BulkResultRelationship) {
	o.AsyncResult.Set(&v)
}

// SetAsyncResultNil sets the value for AsyncResult to be an explicit nil
func (o *BulkMoCloner) SetAsyncResultNil() {
	o.AsyncResult.Set(nil)
}

// UnsetAsyncResult ensures that no value is present for AsyncResult, not even an explicit nil
func (o *BulkMoCloner) UnsetAsyncResult() {
	o.AsyncResult.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoCloner) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoCloner) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkMoCloner) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkMoCloner) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *BulkMoCloner) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *BulkMoCloner) UnsetOrganization() {
	o.Organization.Unset()
}

func (o BulkMoCloner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkMoCloner) ToMap() (map[string]interface{}, error) {
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
	if o.ExcludeProperties != nil {
		toSerialize["ExcludeProperties"] = o.ExcludeProperties
	}
	if o.Responses != nil {
		toSerialize["Responses"] = o.Responses
	}
	if o.Sources != nil {
		toSerialize["Sources"] = o.Sources
	}
	if o.Targets != nil {
		toSerialize["Targets"] = o.Targets
	}
	if !IsNil(o.WorkflowNameSuffix) {
		toSerialize["WorkflowNameSuffix"] = o.WorkflowNameSuffix
	}
	if o.AsyncResult.IsSet() {
		toSerialize["AsyncResult"] = o.AsyncResult.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkMoCloner) UnmarshalJSON(data []byte) (err error) {
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
	type BulkMoClonerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string           `json:"ObjectType"`
		ExcludeProperties []string         `json:"ExcludeProperties,omitempty"`
		Responses         []BulkRestResult `json:"Responses,omitempty"`
		Sources           []MoBaseMo       `json:"Sources,omitempty"`
		Targets           []MoBaseMo       `json:"Targets,omitempty"`
		// A user-friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_).
		WorkflowNameSuffix *string                                      `json:"WorkflowNameSuffix,omitempty" validate:"regexp=^$|^[a-zA-Z0-9]{1}[\\\\sa-zA-Z0-9_.\\\\,\\/:-]{0,63}$"`
		AsyncResult        NullableBulkResultRelationship               `json:"AsyncResult,omitempty"`
		Organization       NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varBulkMoClonerWithoutEmbeddedStruct := BulkMoClonerWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBulkMoClonerWithoutEmbeddedStruct)
	if err == nil {
		varBulkMoCloner := _BulkMoCloner{}
		varBulkMoCloner.ClassId = varBulkMoClonerWithoutEmbeddedStruct.ClassId
		varBulkMoCloner.ObjectType = varBulkMoClonerWithoutEmbeddedStruct.ObjectType
		varBulkMoCloner.ExcludeProperties = varBulkMoClonerWithoutEmbeddedStruct.ExcludeProperties
		varBulkMoCloner.Responses = varBulkMoClonerWithoutEmbeddedStruct.Responses
		varBulkMoCloner.Sources = varBulkMoClonerWithoutEmbeddedStruct.Sources
		varBulkMoCloner.Targets = varBulkMoClonerWithoutEmbeddedStruct.Targets
		varBulkMoCloner.WorkflowNameSuffix = varBulkMoClonerWithoutEmbeddedStruct.WorkflowNameSuffix
		varBulkMoCloner.AsyncResult = varBulkMoClonerWithoutEmbeddedStruct.AsyncResult
		varBulkMoCloner.Organization = varBulkMoClonerWithoutEmbeddedStruct.Organization
		*o = BulkMoCloner(varBulkMoCloner)
	} else {
		return err
	}

	varBulkMoCloner := _BulkMoCloner{}

	err = json.Unmarshal(data, &varBulkMoCloner)
	if err == nil {
		o.MoBaseMo = varBulkMoCloner.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExcludeProperties")
		delete(additionalProperties, "Responses")
		delete(additionalProperties, "Sources")
		delete(additionalProperties, "Targets")
		delete(additionalProperties, "WorkflowNameSuffix")
		delete(additionalProperties, "AsyncResult")
		delete(additionalProperties, "Organization")

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

type NullableBulkMoCloner struct {
	value *BulkMoCloner
	isSet bool
}

func (v NullableBulkMoCloner) Get() *BulkMoCloner {
	return v.value
}

func (v *NullableBulkMoCloner) Set(val *BulkMoCloner) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkMoCloner) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkMoCloner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkMoCloner(val *BulkMoCloner) *NullableBulkMoCloner {
	return &NullableBulkMoCloner{value: val, isSet: true}
}

func (v NullableBulkMoCloner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkMoCloner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
