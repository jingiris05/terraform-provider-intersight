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

// checks if the AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions{}

// AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions Captures configuration specific to the Microsoft Azure Service Principal target for the Workload Optimizer service.
type AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Azure has different endpoints for managing Germany subscriptions. Azure cloud type helps in differentiating between regular subscriptions and Germany subscriptions to manage the Azure services for workload optimization. Documentation for germany cloud [link](https://docs.microsoft.com/en-us/azure/germany/germany-manage-subscriptions). * `Global` - Global cloud type for Azure subscription. * `Germany` - Germany cloud type for Azure subscription.
	AzureCloudType *string `json:"AzureCloudType,omitempty"`
	// ID of the tenant used while authenticating the managed target.
	TenantId             *string `json:"TenantId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions

// NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions instantiates a new AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions(classId string, objectType string) *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions {
	this := AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	var azureCloudType string = "Global"
	this.AzureCloudType = &azureCloudType
	return &this
}

// NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsWithDefaults() *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions {
	this := AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions{}
	var classId string = "asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions"
	this.ObjectType = objectType
	var azureCloudType string = "Global"
	this.AzureCloudType = &azureCloudType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions" of the ClassId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) GetDefaultClassId() interface{} {
	return "asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions"
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions" of the ObjectType field.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) GetDefaultObjectType() interface{} {
	return "asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions"
}

// GetAzureCloudType returns the AzureCloudType field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) GetAzureCloudType() string {
	if o == nil || IsNil(o.AzureCloudType) {
		var ret string
		return ret
	}
	return *o.AzureCloudType
}

// GetAzureCloudTypeOk returns a tuple with the AzureCloudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) GetAzureCloudTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AzureCloudType) {
		return nil, false
	}
	return o.AzureCloudType, true
}

// HasAzureCloudType returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) HasAzureCloudType() bool {
	if o != nil && !IsNil(o.AzureCloudType) {
		return true
	}

	return false
}

// SetAzureCloudType gets a reference to the given string and assigns it to the AzureCloudType field.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) SetAzureCloudType(v string) {
	o.AzureCloudType = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) SetTenantId(v string) {
	o.TenantId = &v
}

func (o AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return map[string]interface{}{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return map[string]interface{}{}, errAssetServiceOptions
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AzureCloudType) {
		toSerialize["AzureCloudType"] = o.AzureCloudType
	}
	if !IsNil(o.TenantId) {
		toSerialize["TenantId"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) UnmarshalJSON(data []byte) (err error) {
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
	type AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Azure has different endpoints for managing Germany subscriptions. Azure cloud type helps in differentiating between regular subscriptions and Germany subscriptions to manage the Azure services for workload optimization. Documentation for germany cloud [link](https://docs.microsoft.com/en-us/azure/germany/germany-manage-subscriptions). * `Global` - Global cloud type for Azure subscription. * `Germany` - Germany cloud type for Azure subscription.
		AzureCloudType *string `json:"AzureCloudType,omitempty"`
		// ID of the tenant used while authenticating the managed target.
		TenantId *string `json:"TenantId,omitempty"`
	}

	varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsWithoutEmbeddedStruct := AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions := _AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions{}
		varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions.ClassId = varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsWithoutEmbeddedStruct.ClassId
		varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions.ObjectType = varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsWithoutEmbeddedStruct.ObjectType
		varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions.AzureCloudType = varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsWithoutEmbeddedStruct.AzureCloudType
		varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions.TenantId = varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsWithoutEmbeddedStruct.TenantId
		*o = AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions(varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions)
	} else {
		return err
	}

	varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions := _AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions{}

	err = json.Unmarshal(data, &varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AzureCloudType")
		delete(additionalProperties, "TenantId")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions struct {
	value *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions
	isSet bool
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) Get() *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) Set(val *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions(val *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) *NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions {
	return &NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
