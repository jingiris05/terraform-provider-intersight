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

// checks if the IssueOdataCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueOdataCondition{}

// IssueOdataCondition OdataCondition defines, via an Odata predicate, a set of criteria under which an issue exists.
type IssueOdataCondition struct {
	IssueCondition
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string           `json:"ObjectType"`
	DeviceTags []IssueDeviceTag `json:"DeviceTags,omitempty"`
	// The Intersight object type on which the condition is to be applied. The object type may be either a concrete type such as compute.RackUnit or a shared parent type such as compute.Physical.
	Motype *string `json:"Motype,omitempty"`
	// An Odata query string interpreted to be the value portion of a $filter expression. For example, for the function query $filter=Serial eq 'ABC', the 'query' field should be assigned the string Serial eq 'ABC.
	Query                *string `json:"Query,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IssueOdataCondition IssueOdataCondition

// NewIssueOdataCondition instantiates a new IssueOdataCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueOdataCondition(classId string, objectType string) *IssueOdataCondition {
	this := IssueOdataCondition{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIssueOdataConditionWithDefaults instantiates a new IssueOdataCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueOdataConditionWithDefaults() *IssueOdataCondition {
	this := IssueOdataCondition{}
	var classId string = "issue.OdataCondition"
	this.ClassId = classId
	var objectType string = "issue.OdataCondition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IssueOdataCondition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IssueOdataCondition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IssueOdataCondition) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "issue.OdataCondition" of the ClassId field.
func (o *IssueOdataCondition) GetDefaultClassId() interface{} {
	return "issue.OdataCondition"
}

// GetObjectType returns the ObjectType field value
func (o *IssueOdataCondition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IssueOdataCondition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IssueOdataCondition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "issue.OdataCondition" of the ObjectType field.
func (o *IssueOdataCondition) GetDefaultObjectType() interface{} {
	return "issue.OdataCondition"
}

// GetDeviceTags returns the DeviceTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueOdataCondition) GetDeviceTags() []IssueDeviceTag {
	if o == nil {
		var ret []IssueDeviceTag
		return ret
	}
	return o.DeviceTags
}

// GetDeviceTagsOk returns a tuple with the DeviceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueOdataCondition) GetDeviceTagsOk() ([]IssueDeviceTag, bool) {
	if o == nil || IsNil(o.DeviceTags) {
		return nil, false
	}
	return o.DeviceTags, true
}

// HasDeviceTags returns a boolean if a field has been set.
func (o *IssueOdataCondition) HasDeviceTags() bool {
	if o != nil && !IsNil(o.DeviceTags) {
		return true
	}

	return false
}

// SetDeviceTags gets a reference to the given []IssueDeviceTag and assigns it to the DeviceTags field.
func (o *IssueOdataCondition) SetDeviceTags(v []IssueDeviceTag) {
	o.DeviceTags = v
}

// GetMotype returns the Motype field value if set, zero value otherwise.
func (o *IssueOdataCondition) GetMotype() string {
	if o == nil || IsNil(o.Motype) {
		var ret string
		return ret
	}
	return *o.Motype
}

// GetMotypeOk returns a tuple with the Motype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueOdataCondition) GetMotypeOk() (*string, bool) {
	if o == nil || IsNil(o.Motype) {
		return nil, false
	}
	return o.Motype, true
}

// HasMotype returns a boolean if a field has been set.
func (o *IssueOdataCondition) HasMotype() bool {
	if o != nil && !IsNil(o.Motype) {
		return true
	}

	return false
}

// SetMotype gets a reference to the given string and assigns it to the Motype field.
func (o *IssueOdataCondition) SetMotype(v string) {
	o.Motype = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *IssueOdataCondition) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueOdataCondition) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *IssueOdataCondition) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *IssueOdataCondition) SetQuery(v string) {
	o.Query = &v
}

func (o IssueOdataCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueOdataCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedIssueCondition, errIssueCondition := json.Marshal(o.IssueCondition)
	if errIssueCondition != nil {
		return map[string]interface{}{}, errIssueCondition
	}
	errIssueCondition = json.Unmarshal([]byte(serializedIssueCondition), &toSerialize)
	if errIssueCondition != nil {
		return map[string]interface{}{}, errIssueCondition
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.DeviceTags != nil {
		toSerialize["DeviceTags"] = o.DeviceTags
	}
	if !IsNil(o.Motype) {
		toSerialize["Motype"] = o.Motype
	}
	if !IsNil(o.Query) {
		toSerialize["Query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IssueOdataCondition) UnmarshalJSON(data []byte) (err error) {
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
	type IssueOdataConditionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string           `json:"ObjectType"`
		DeviceTags []IssueDeviceTag `json:"DeviceTags,omitempty"`
		// The Intersight object type on which the condition is to be applied. The object type may be either a concrete type such as compute.RackUnit or a shared parent type such as compute.Physical.
		Motype *string `json:"Motype,omitempty"`
		// An Odata query string interpreted to be the value portion of a $filter expression. For example, for the function query $filter=Serial eq 'ABC', the 'query' field should be assigned the string Serial eq 'ABC.
		Query *string `json:"Query,omitempty"`
	}

	varIssueOdataConditionWithoutEmbeddedStruct := IssueOdataConditionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIssueOdataConditionWithoutEmbeddedStruct)
	if err == nil {
		varIssueOdataCondition := _IssueOdataCondition{}
		varIssueOdataCondition.ClassId = varIssueOdataConditionWithoutEmbeddedStruct.ClassId
		varIssueOdataCondition.ObjectType = varIssueOdataConditionWithoutEmbeddedStruct.ObjectType
		varIssueOdataCondition.DeviceTags = varIssueOdataConditionWithoutEmbeddedStruct.DeviceTags
		varIssueOdataCondition.Motype = varIssueOdataConditionWithoutEmbeddedStruct.Motype
		varIssueOdataCondition.Query = varIssueOdataConditionWithoutEmbeddedStruct.Query
		*o = IssueOdataCondition(varIssueOdataCondition)
	} else {
		return err
	}

	varIssueOdataCondition := _IssueOdataCondition{}

	err = json.Unmarshal(data, &varIssueOdataCondition)
	if err == nil {
		o.IssueCondition = varIssueOdataCondition.IssueCondition
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceTags")
		delete(additionalProperties, "Motype")
		delete(additionalProperties, "Query")

		// remove fields from embedded structs
		reflectIssueCondition := reflect.ValueOf(o.IssueCondition)
		for i := 0; i < reflectIssueCondition.Type().NumField(); i++ {
			t := reflectIssueCondition.Type().Field(i)

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

type NullableIssueOdataCondition struct {
	value *IssueOdataCondition
	isSet bool
}

func (v NullableIssueOdataCondition) Get() *IssueOdataCondition {
	return v.value
}

func (v *NullableIssueOdataCondition) Set(val *IssueOdataCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueOdataCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueOdataCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueOdataCondition(val *IssueOdataCondition) *NullableIssueOdataCondition {
	return &NullableIssueOdataCondition{value: val, isSet: true}
}

func (v NullableIssueOdataCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueOdataCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
