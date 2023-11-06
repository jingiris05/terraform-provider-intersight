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
	"reflect"
	"strings"
)

// AssetNoAuthenticationCredential Used against a target which does not require any credential.
type AssetNoAuthenticationCredential struct {
	AssetCredential
	AdditionalProperties map[string]interface{}
}

type _AssetNoAuthenticationCredential AssetNoAuthenticationCredential

// NewAssetNoAuthenticationCredential instantiates a new AssetNoAuthenticationCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetNoAuthenticationCredential(classId string, objectType string) *AssetNoAuthenticationCredential {
	this := AssetNoAuthenticationCredential{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetNoAuthenticationCredentialWithDefaults instantiates a new AssetNoAuthenticationCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetNoAuthenticationCredentialWithDefaults() *AssetNoAuthenticationCredential {
	this := AssetNoAuthenticationCredential{}
	return &this
}

func (o AssetNoAuthenticationCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetCredential, errAssetCredential := json.Marshal(o.AssetCredential)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	errAssetCredential = json.Unmarshal([]byte(serializedAssetCredential), &toSerialize)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetNoAuthenticationCredential) UnmarshalJSON(bytes []byte) (err error) {
	type AssetNoAuthenticationCredentialWithoutEmbeddedStruct struct {
	}

	varAssetNoAuthenticationCredentialWithoutEmbeddedStruct := AssetNoAuthenticationCredentialWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetNoAuthenticationCredentialWithoutEmbeddedStruct)
	if err == nil {
		varAssetNoAuthenticationCredential := _AssetNoAuthenticationCredential{}
		*o = AssetNoAuthenticationCredential(varAssetNoAuthenticationCredential)
	} else {
		return err
	}

	varAssetNoAuthenticationCredential := _AssetNoAuthenticationCredential{}

	err = json.Unmarshal(bytes, &varAssetNoAuthenticationCredential)
	if err == nil {
		o.AssetCredential = varAssetNoAuthenticationCredential.AssetCredential
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectAssetCredential := reflect.ValueOf(o.AssetCredential)
		for i := 0; i < reflectAssetCredential.Type().NumField(); i++ {
			t := reflectAssetCredential.Type().Field(i)

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

type NullableAssetNoAuthenticationCredential struct {
	value *AssetNoAuthenticationCredential
	isSet bool
}

func (v NullableAssetNoAuthenticationCredential) Get() *AssetNoAuthenticationCredential {
	return v.value
}

func (v *NullableAssetNoAuthenticationCredential) Set(val *AssetNoAuthenticationCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetNoAuthenticationCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetNoAuthenticationCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetNoAuthenticationCredential(val *AssetNoAuthenticationCredential) *NullableAssetNoAuthenticationCredential {
	return &NullableAssetNoAuthenticationCredential{value: val, isSet: true}
}

func (v NullableAssetNoAuthenticationCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetNoAuthenticationCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
