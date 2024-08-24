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

// checks if the ServerDisruption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerDisruption{}

// ServerDisruption The server disruption details are captured here.
type ServerDisruption struct {
	PolicyAbstractDisruption
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                       `json:"ObjectType"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerDisruption ServerDisruption

// NewServerDisruption instantiates a new ServerDisruption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerDisruption(classId string, objectType string) *ServerDisruption {
	this := ServerDisruption{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServerDisruptionWithDefaults instantiates a new ServerDisruption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerDisruptionWithDefaults() *ServerDisruption {
	this := ServerDisruption{}
	var classId string = "server.Disruption"
	this.ClassId = classId
	var objectType string = "server.Disruption"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerDisruption) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerDisruption) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerDisruption) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "server.Disruption" of the ClassId field.
func (o *ServerDisruption) GetDefaultClassId() interface{} {
	return "server.Disruption"
}

// GetObjectType returns the ObjectType field value
func (o *ServerDisruption) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerDisruption) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerDisruption) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "server.Disruption" of the ObjectType field.
func (o *ServerDisruption) GetDefaultObjectType() interface{} {
	return "server.Disruption"
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerDisruption) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerDisruption) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *ServerDisruption) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ServerDisruption) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *ServerDisruption) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *ServerDisruption) UnsetOrganization() {
	o.Organization.Unset()
}

func (o ServerDisruption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerDisruption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractDisruption, errPolicyAbstractDisruption := json.Marshal(o.PolicyAbstractDisruption)
	if errPolicyAbstractDisruption != nil {
		return map[string]interface{}{}, errPolicyAbstractDisruption
	}
	errPolicyAbstractDisruption = json.Unmarshal([]byte(serializedPolicyAbstractDisruption), &toSerialize)
	if errPolicyAbstractDisruption != nil {
		return map[string]interface{}{}, errPolicyAbstractDisruption
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerDisruption) UnmarshalJSON(data []byte) (err error) {
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
	type ServerDisruptionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                       `json:"ObjectType"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varServerDisruptionWithoutEmbeddedStruct := ServerDisruptionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varServerDisruptionWithoutEmbeddedStruct)
	if err == nil {
		varServerDisruption := _ServerDisruption{}
		varServerDisruption.ClassId = varServerDisruptionWithoutEmbeddedStruct.ClassId
		varServerDisruption.ObjectType = varServerDisruptionWithoutEmbeddedStruct.ObjectType
		varServerDisruption.Organization = varServerDisruptionWithoutEmbeddedStruct.Organization
		*o = ServerDisruption(varServerDisruption)
	} else {
		return err
	}

	varServerDisruption := _ServerDisruption{}

	err = json.Unmarshal(data, &varServerDisruption)
	if err == nil {
		o.PolicyAbstractDisruption = varServerDisruption.PolicyAbstractDisruption
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractDisruption := reflect.ValueOf(o.PolicyAbstractDisruption)
		for i := 0; i < reflectPolicyAbstractDisruption.Type().NumField(); i++ {
			t := reflectPolicyAbstractDisruption.Type().Field(i)

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

type NullableServerDisruption struct {
	value *ServerDisruption
	isSet bool
}

func (v NullableServerDisruption) Get() *ServerDisruption {
	return v.value
}

func (v *NullableServerDisruption) Set(val *ServerDisruption) {
	v.value = val
	v.isSet = true
}

func (v NullableServerDisruption) IsSet() bool {
	return v.isSet
}

func (v *NullableServerDisruption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerDisruption(val *ServerDisruption) *NullableServerDisruption {
	return &NullableServerDisruption{value: val, isSet: true}
}

func (v NullableServerDisruption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerDisruption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
