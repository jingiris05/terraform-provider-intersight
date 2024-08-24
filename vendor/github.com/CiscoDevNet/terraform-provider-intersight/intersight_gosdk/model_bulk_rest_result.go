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

// checks if the BulkRestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkRestResult{}

// BulkRestResult The result of an individual REST API action.
type BulkRestResult struct {
	BulkApiResult
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string    `json:"ObjectType"`
	Body       *MoBaseMo `json:"Body,omitempty"`
	// The response string for an individual REST API action.
	BodyString           *string `json:"BodyString,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkRestResult BulkRestResult

// NewBulkRestResult instantiates a new BulkRestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkRestResult(classId string, objectType string) *BulkRestResult {
	this := BulkRestResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBulkRestResultWithDefaults instantiates a new BulkRestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkRestResultWithDefaults() *BulkRestResult {
	this := BulkRestResult{}
	var classId string = "bulk.RestResult"
	this.ClassId = classId
	var objectType string = "bulk.RestResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkRestResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkRestResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkRestResult) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "bulk.RestResult" of the ClassId field.
func (o *BulkRestResult) GetDefaultClassId() interface{} {
	return "bulk.RestResult"
}

// GetObjectType returns the ObjectType field value
func (o *BulkRestResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkRestResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkRestResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "bulk.RestResult" of the ObjectType field.
func (o *BulkRestResult) GetDefaultObjectType() interface{} {
	return "bulk.RestResult"
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BulkRestResult) GetBody() MoBaseMo {
	if o == nil || IsNil(o.Body) {
		var ret MoBaseMo
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRestResult) GetBodyOk() (*MoBaseMo, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BulkRestResult) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given MoBaseMo and assigns it to the Body field.
func (o *BulkRestResult) SetBody(v MoBaseMo) {
	o.Body = &v
}

// GetBodyString returns the BodyString field value if set, zero value otherwise.
func (o *BulkRestResult) GetBodyString() string {
	if o == nil || IsNil(o.BodyString) {
		var ret string
		return ret
	}
	return *o.BodyString
}

// GetBodyStringOk returns a tuple with the BodyString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkRestResult) GetBodyStringOk() (*string, bool) {
	if o == nil || IsNil(o.BodyString) {
		return nil, false
	}
	return o.BodyString, true
}

// HasBodyString returns a boolean if a field has been set.
func (o *BulkRestResult) HasBodyString() bool {
	if o != nil && !IsNil(o.BodyString) {
		return true
	}

	return false
}

// SetBodyString gets a reference to the given string and assigns it to the BodyString field.
func (o *BulkRestResult) SetBodyString(v string) {
	o.BodyString = &v
}

func (o BulkRestResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkRestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBulkApiResult, errBulkApiResult := json.Marshal(o.BulkApiResult)
	if errBulkApiResult != nil {
		return map[string]interface{}{}, errBulkApiResult
	}
	errBulkApiResult = json.Unmarshal([]byte(serializedBulkApiResult), &toSerialize)
	if errBulkApiResult != nil {
		return map[string]interface{}{}, errBulkApiResult
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Body) {
		toSerialize["Body"] = o.Body
	}
	if !IsNil(o.BodyString) {
		toSerialize["BodyString"] = o.BodyString
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkRestResult) UnmarshalJSON(data []byte) (err error) {
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
	type BulkRestResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string    `json:"ObjectType"`
		Body       *MoBaseMo `json:"Body,omitempty"`
		// The response string for an individual REST API action.
		BodyString *string `json:"BodyString,omitempty"`
	}

	varBulkRestResultWithoutEmbeddedStruct := BulkRestResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBulkRestResultWithoutEmbeddedStruct)
	if err == nil {
		varBulkRestResult := _BulkRestResult{}
		varBulkRestResult.ClassId = varBulkRestResultWithoutEmbeddedStruct.ClassId
		varBulkRestResult.ObjectType = varBulkRestResultWithoutEmbeddedStruct.ObjectType
		varBulkRestResult.Body = varBulkRestResultWithoutEmbeddedStruct.Body
		varBulkRestResult.BodyString = varBulkRestResultWithoutEmbeddedStruct.BodyString
		*o = BulkRestResult(varBulkRestResult)
	} else {
		return err
	}

	varBulkRestResult := _BulkRestResult{}

	err = json.Unmarshal(data, &varBulkRestResult)
	if err == nil {
		o.BulkApiResult = varBulkRestResult.BulkApiResult
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "BodyString")

		// remove fields from embedded structs
		reflectBulkApiResult := reflect.ValueOf(o.BulkApiResult)
		for i := 0; i < reflectBulkApiResult.Type().NumField(); i++ {
			t := reflectBulkApiResult.Type().Field(i)

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

type NullableBulkRestResult struct {
	value *BulkRestResult
	isSet bool
}

func (v NullableBulkRestResult) Get() *BulkRestResult {
	return v.value
}

func (v *NullableBulkRestResult) Set(val *BulkRestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkRestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkRestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkRestResult(val *BulkRestResult) *NullableBulkRestResult {
	return &NullableBulkRestResult{value: val, isSet: true}
}

func (v NullableBulkRestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkRestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
