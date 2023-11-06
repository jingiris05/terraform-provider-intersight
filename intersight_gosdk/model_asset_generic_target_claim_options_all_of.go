/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-14237
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// AssetGenericTargetClaimOptionsAllOf Definition of the list of properties defined in 'asset.GenericTargetClaimOptions', excluding properties defined in parent classes.
type AssetGenericTargetClaimOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The generic device connector will consume these options.
	TargetOptions        interface{} `json:"TargetOptions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetGenericTargetClaimOptionsAllOf AssetGenericTargetClaimOptionsAllOf

// NewAssetGenericTargetClaimOptionsAllOf instantiates a new AssetGenericTargetClaimOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetGenericTargetClaimOptionsAllOf(classId string, objectType string) *AssetGenericTargetClaimOptionsAllOf {
	this := AssetGenericTargetClaimOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetGenericTargetClaimOptionsAllOfWithDefaults instantiates a new AssetGenericTargetClaimOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetGenericTargetClaimOptionsAllOfWithDefaults() *AssetGenericTargetClaimOptionsAllOf {
	this := AssetGenericTargetClaimOptionsAllOf{}
	var classId string = "asset.GenericTargetClaimOptions"
	this.ClassId = classId
	var objectType string = "asset.GenericTargetClaimOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetGenericTargetClaimOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetGenericTargetClaimOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetGenericTargetClaimOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetGenericTargetClaimOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetGenericTargetClaimOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetGenericTargetClaimOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetOptions returns the TargetOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetGenericTargetClaimOptionsAllOf) GetTargetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TargetOptions
}

// GetTargetOptionsOk returns a tuple with the TargetOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetGenericTargetClaimOptionsAllOf) GetTargetOptionsOk() (*interface{}, bool) {
	if o == nil || o.TargetOptions == nil {
		return nil, false
	}
	return &o.TargetOptions, true
}

// HasTargetOptions returns a boolean if a field has been set.
func (o *AssetGenericTargetClaimOptionsAllOf) HasTargetOptions() bool {
	if o != nil && o.TargetOptions != nil {
		return true
	}

	return false
}

// SetTargetOptions gets a reference to the given interface{} and assigns it to the TargetOptions field.
func (o *AssetGenericTargetClaimOptionsAllOf) SetTargetOptions(v interface{}) {
	o.TargetOptions = v
}

func (o AssetGenericTargetClaimOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetOptions != nil {
		toSerialize["TargetOptions"] = o.TargetOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetGenericTargetClaimOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetGenericTargetClaimOptionsAllOf := _AssetGenericTargetClaimOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetGenericTargetClaimOptionsAllOf); err == nil {
		*o = AssetGenericTargetClaimOptionsAllOf(varAssetGenericTargetClaimOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetOptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetGenericTargetClaimOptionsAllOf struct {
	value *AssetGenericTargetClaimOptionsAllOf
	isSet bool
}

func (v NullableAssetGenericTargetClaimOptionsAllOf) Get() *AssetGenericTargetClaimOptionsAllOf {
	return v.value
}

func (v *NullableAssetGenericTargetClaimOptionsAllOf) Set(val *AssetGenericTargetClaimOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetGenericTargetClaimOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetGenericTargetClaimOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetGenericTargetClaimOptionsAllOf(val *AssetGenericTargetClaimOptionsAllOf) *NullableAssetGenericTargetClaimOptionsAllOf {
	return &NullableAssetGenericTargetClaimOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetGenericTargetClaimOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetGenericTargetClaimOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
