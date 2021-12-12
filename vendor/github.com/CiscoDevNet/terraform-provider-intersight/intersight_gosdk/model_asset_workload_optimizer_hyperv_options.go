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

// AssetWorkloadOptimizerHypervOptions Captures configuration specific to the Hyper-V target for the Workload Optimizer service.
type AssetWorkloadOptimizerHypervOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When enabled, all Hyper-V hosts in the same cluster of specified Hyper-V host target will be discovered Each server must still be configured to allow remote management (for example, configuring WinRM using a GPO).
	DiscoverHostCluster *bool `json:"DiscoverHostCluster,omitempty"`
	// Fully qualified domain name of the Hyper-V target. It is used to get the name of the Hyper-V cluster node and do Active Directory authentication process through Kerberos scheme.
	FullDomainName       *string `json:"FullDomainName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerHypervOptions AssetWorkloadOptimizerHypervOptions

// NewAssetWorkloadOptimizerHypervOptions instantiates a new AssetWorkloadOptimizerHypervOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerHypervOptions(classId string, objectType string) *AssetWorkloadOptimizerHypervOptions {
	this := AssetWorkloadOptimizerHypervOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerHypervOptionsWithDefaults instantiates a new AssetWorkloadOptimizerHypervOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerHypervOptionsWithDefaults() *AssetWorkloadOptimizerHypervOptions {
	this := AssetWorkloadOptimizerHypervOptions{}
	var classId string = "asset.WorkloadOptimizerHypervOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerHypervOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerHypervOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerHypervOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerHypervOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerHypervOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerHypervOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerHypervOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDiscoverHostCluster returns the DiscoverHostCluster field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerHypervOptions) GetDiscoverHostCluster() bool {
	if o == nil || o.DiscoverHostCluster == nil {
		var ret bool
		return ret
	}
	return *o.DiscoverHostCluster
}

// GetDiscoverHostClusterOk returns a tuple with the DiscoverHostCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerHypervOptions) GetDiscoverHostClusterOk() (*bool, bool) {
	if o == nil || o.DiscoverHostCluster == nil {
		return nil, false
	}
	return o.DiscoverHostCluster, true
}

// HasDiscoverHostCluster returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerHypervOptions) HasDiscoverHostCluster() bool {
	if o != nil && o.DiscoverHostCluster != nil {
		return true
	}

	return false
}

// SetDiscoverHostCluster gets a reference to the given bool and assigns it to the DiscoverHostCluster field.
func (o *AssetWorkloadOptimizerHypervOptions) SetDiscoverHostCluster(v bool) {
	o.DiscoverHostCluster = &v
}

// GetFullDomainName returns the FullDomainName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerHypervOptions) GetFullDomainName() string {
	if o == nil || o.FullDomainName == nil {
		var ret string
		return ret
	}
	return *o.FullDomainName
}

// GetFullDomainNameOk returns a tuple with the FullDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerHypervOptions) GetFullDomainNameOk() (*string, bool) {
	if o == nil || o.FullDomainName == nil {
		return nil, false
	}
	return o.FullDomainName, true
}

// HasFullDomainName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerHypervOptions) HasFullDomainName() bool {
	if o != nil && o.FullDomainName != nil {
		return true
	}

	return false
}

// SetFullDomainName gets a reference to the given string and assigns it to the FullDomainName field.
func (o *AssetWorkloadOptimizerHypervOptions) SetFullDomainName(v string) {
	o.FullDomainName = &v
}

func (o AssetWorkloadOptimizerHypervOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DiscoverHostCluster != nil {
		toSerialize["DiscoverHostCluster"] = o.DiscoverHostCluster
	}
	if o.FullDomainName != nil {
		toSerialize["FullDomainName"] = o.FullDomainName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerHypervOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetWorkloadOptimizerHypervOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// When enabled, all Hyper-V hosts in the same cluster of specified Hyper-V host target will be discovered Each server must still be configured to allow remote management (for example, configuring WinRM using a GPO).
		DiscoverHostCluster *bool `json:"DiscoverHostCluster,omitempty"`
		// Fully qualified domain name of the Hyper-V target. It is used to get the name of the Hyper-V cluster node and do Active Directory authentication process through Kerberos scheme.
		FullDomainName *string `json:"FullDomainName,omitempty"`
	}

	varAssetWorkloadOptimizerHypervOptionsWithoutEmbeddedStruct := AssetWorkloadOptimizerHypervOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerHypervOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetWorkloadOptimizerHypervOptions := _AssetWorkloadOptimizerHypervOptions{}
		varAssetWorkloadOptimizerHypervOptions.ClassId = varAssetWorkloadOptimizerHypervOptionsWithoutEmbeddedStruct.ClassId
		varAssetWorkloadOptimizerHypervOptions.ObjectType = varAssetWorkloadOptimizerHypervOptionsWithoutEmbeddedStruct.ObjectType
		varAssetWorkloadOptimizerHypervOptions.DiscoverHostCluster = varAssetWorkloadOptimizerHypervOptionsWithoutEmbeddedStruct.DiscoverHostCluster
		varAssetWorkloadOptimizerHypervOptions.FullDomainName = varAssetWorkloadOptimizerHypervOptionsWithoutEmbeddedStruct.FullDomainName
		*o = AssetWorkloadOptimizerHypervOptions(varAssetWorkloadOptimizerHypervOptions)
	} else {
		return err
	}

	varAssetWorkloadOptimizerHypervOptions := _AssetWorkloadOptimizerHypervOptions{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerHypervOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetWorkloadOptimizerHypervOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DiscoverHostCluster")
		delete(additionalProperties, "FullDomainName")

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

type NullableAssetWorkloadOptimizerHypervOptions struct {
	value *AssetWorkloadOptimizerHypervOptions
	isSet bool
}

func (v NullableAssetWorkloadOptimizerHypervOptions) Get() *AssetWorkloadOptimizerHypervOptions {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerHypervOptions) Set(val *AssetWorkloadOptimizerHypervOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerHypervOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerHypervOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerHypervOptions(val *AssetWorkloadOptimizerHypervOptions) *NullableAssetWorkloadOptimizerHypervOptions {
	return &NullableAssetWorkloadOptimizerHypervOptions{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerHypervOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerHypervOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
