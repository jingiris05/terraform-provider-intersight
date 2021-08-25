/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-24T21:18:00Z.
 *
 * API version: 1.0.9-4404
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IamFeatureDefinition Feature parameters specified for the account.
type IamFeatureDefinition struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The beta feature that will be enabled for specific account. * `IWO` - Intersight Workflow Optimizer. * `Hitachi` - Support to claim Hitachi Storage arrays using the Intersight Orchestrator framework. * `Kubernetes` - Enables ability to create and manage Kubernetes clusters. * `NetAppIO` - Support to claim NetApp Storage arrays as IO targets. * `IvsPublicCloud` - Enables virtualization service for public clouds. * `TerraformCloud` - Enables an ability to create Terraform Cloud. * `WashingtonEFT` - Support for EFT customers to use Washington firmware images for upgrades.
	Feature              *string `json:"Feature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamFeatureDefinition IamFeatureDefinition

// NewIamFeatureDefinition instantiates a new IamFeatureDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamFeatureDefinition(classId string, objectType string) *IamFeatureDefinition {
	this := IamFeatureDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var feature string = "IWO"
	this.Feature = &feature
	return &this
}

// NewIamFeatureDefinitionWithDefaults instantiates a new IamFeatureDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamFeatureDefinitionWithDefaults() *IamFeatureDefinition {
	this := IamFeatureDefinition{}
	var classId string = "iam.FeatureDefinition"
	this.ClassId = classId
	var objectType string = "iam.FeatureDefinition"
	this.ObjectType = objectType
	var feature string = "IWO"
	this.Feature = &feature
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamFeatureDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamFeatureDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamFeatureDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamFeatureDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamFeatureDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamFeatureDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *IamFeatureDefinition) GetFeature() string {
	if o == nil || o.Feature == nil {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamFeatureDefinition) GetFeatureOk() (*string, bool) {
	if o == nil || o.Feature == nil {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *IamFeatureDefinition) HasFeature() bool {
	if o != nil && o.Feature != nil {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *IamFeatureDefinition) SetFeature(v string) {
	o.Feature = &v
}

func (o IamFeatureDefinition) MarshalJSON() ([]byte, error) {
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
	if o.Feature != nil {
		toSerialize["Feature"] = o.Feature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamFeatureDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type IamFeatureDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The beta feature that will be enabled for specific account. * `IWO` - Intersight Workflow Optimizer. * `Hitachi` - Support to claim Hitachi Storage arrays using the Intersight Orchestrator framework. * `Kubernetes` - Enables ability to create and manage Kubernetes clusters. * `NetAppIO` - Support to claim NetApp Storage arrays as IO targets. * `IvsPublicCloud` - Enables virtualization service for public clouds. * `TerraformCloud` - Enables an ability to create Terraform Cloud. * `WashingtonEFT` - Support for EFT customers to use Washington firmware images for upgrades.
		Feature *string `json:"Feature,omitempty"`
	}

	varIamFeatureDefinitionWithoutEmbeddedStruct := IamFeatureDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamFeatureDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varIamFeatureDefinition := _IamFeatureDefinition{}
		varIamFeatureDefinition.ClassId = varIamFeatureDefinitionWithoutEmbeddedStruct.ClassId
		varIamFeatureDefinition.ObjectType = varIamFeatureDefinitionWithoutEmbeddedStruct.ObjectType
		varIamFeatureDefinition.Feature = varIamFeatureDefinitionWithoutEmbeddedStruct.Feature
		*o = IamFeatureDefinition(varIamFeatureDefinition)
	} else {
		return err
	}

	varIamFeatureDefinition := _IamFeatureDefinition{}

	err = json.Unmarshal(bytes, &varIamFeatureDefinition)
	if err == nil {
		o.MoBaseComplexType = varIamFeatureDefinition.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Feature")

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

type NullableIamFeatureDefinition struct {
	value *IamFeatureDefinition
	isSet bool
}

func (v NullableIamFeatureDefinition) Get() *IamFeatureDefinition {
	return v.value
}

func (v *NullableIamFeatureDefinition) Set(val *IamFeatureDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableIamFeatureDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableIamFeatureDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamFeatureDefinition(val *IamFeatureDefinition) *NullableIamFeatureDefinition {
	return &NullableIamFeatureDefinition{value: val, isSet: true}
}

func (v NullableIamFeatureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamFeatureDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
