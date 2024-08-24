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

// checks if the AssetPrivateKeyCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetPrivateKeyCredential{}

// AssetPrivateKeyCredential A credential which performs authentication based on a private key for targets like SSH Endpoint. PrivateKeyCredential used to establish password-less SSH connection to the endpoint.
type AssetPrivateKeyCredential struct {
	AssetCredential
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'passphrase' property has been set.
	IsPassphraseSet *bool `json:"IsPassphraseSet,omitempty"`
	// Indicates whether the value of the 'privateKey' property has been set.
	IsPrivateKeySet *bool `json:"IsPrivateKeySet,omitempty"`
	// The passphrase associated with the private key - Optional.
	Passphrase *string `json:"Passphrase,omitempty"`
	// The private key used to authenticate with a managed target. The corresponding public key needs to be added in the auth list of the remote endpoint.
	PrivateKey *string `json:"PrivateKey,omitempty"`
	// The username used to authenticate with a managed target.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetPrivateKeyCredential AssetPrivateKeyCredential

// NewAssetPrivateKeyCredential instantiates a new AssetPrivateKeyCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetPrivateKeyCredential(classId string, objectType string) *AssetPrivateKeyCredential {
	this := AssetPrivateKeyCredential{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetPrivateKeyCredentialWithDefaults instantiates a new AssetPrivateKeyCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetPrivateKeyCredentialWithDefaults() *AssetPrivateKeyCredential {
	this := AssetPrivateKeyCredential{}
	var classId string = "asset.PrivateKeyCredential"
	this.ClassId = classId
	var objectType string = "asset.PrivateKeyCredential"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetPrivateKeyCredential) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetPrivateKeyCredential) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetPrivateKeyCredential) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.PrivateKeyCredential" of the ClassId field.
func (o *AssetPrivateKeyCredential) GetDefaultClassId() interface{} {
	return "asset.PrivateKeyCredential"
}

// GetObjectType returns the ObjectType field value
func (o *AssetPrivateKeyCredential) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetPrivateKeyCredential) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetPrivateKeyCredential) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.PrivateKeyCredential" of the ObjectType field.
func (o *AssetPrivateKeyCredential) GetDefaultObjectType() interface{} {
	return "asset.PrivateKeyCredential"
}

// GetIsPassphraseSet returns the IsPassphraseSet field value if set, zero value otherwise.
func (o *AssetPrivateKeyCredential) GetIsPassphraseSet() bool {
	if o == nil || IsNil(o.IsPassphraseSet) {
		var ret bool
		return ret
	}
	return *o.IsPassphraseSet
}

// GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetPrivateKeyCredential) GetIsPassphraseSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPassphraseSet) {
		return nil, false
	}
	return o.IsPassphraseSet, true
}

// HasIsPassphraseSet returns a boolean if a field has been set.
func (o *AssetPrivateKeyCredential) HasIsPassphraseSet() bool {
	if o != nil && !IsNil(o.IsPassphraseSet) {
		return true
	}

	return false
}

// SetIsPassphraseSet gets a reference to the given bool and assigns it to the IsPassphraseSet field.
func (o *AssetPrivateKeyCredential) SetIsPassphraseSet(v bool) {
	o.IsPassphraseSet = &v
}

// GetIsPrivateKeySet returns the IsPrivateKeySet field value if set, zero value otherwise.
func (o *AssetPrivateKeyCredential) GetIsPrivateKeySet() bool {
	if o == nil || IsNil(o.IsPrivateKeySet) {
		var ret bool
		return ret
	}
	return *o.IsPrivateKeySet
}

// GetIsPrivateKeySetOk returns a tuple with the IsPrivateKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetPrivateKeyCredential) GetIsPrivateKeySetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivateKeySet) {
		return nil, false
	}
	return o.IsPrivateKeySet, true
}

// HasIsPrivateKeySet returns a boolean if a field has been set.
func (o *AssetPrivateKeyCredential) HasIsPrivateKeySet() bool {
	if o != nil && !IsNil(o.IsPrivateKeySet) {
		return true
	}

	return false
}

// SetIsPrivateKeySet gets a reference to the given bool and assigns it to the IsPrivateKeySet field.
func (o *AssetPrivateKeyCredential) SetIsPrivateKeySet(v bool) {
	o.IsPrivateKeySet = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *AssetPrivateKeyCredential) GetPassphrase() string {
	if o == nil || IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetPrivateKeyCredential) GetPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *AssetPrivateKeyCredential) HasPassphrase() bool {
	if o != nil && !IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *AssetPrivateKeyCredential) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *AssetPrivateKeyCredential) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetPrivateKeyCredential) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *AssetPrivateKeyCredential) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *AssetPrivateKeyCredential) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AssetPrivateKeyCredential) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetPrivateKeyCredential) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AssetPrivateKeyCredential) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AssetPrivateKeyCredential) SetUsername(v string) {
	o.Username = &v
}

func (o AssetPrivateKeyCredential) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetPrivateKeyCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetCredential, errAssetCredential := json.Marshal(o.AssetCredential)
	if errAssetCredential != nil {
		return map[string]interface{}{}, errAssetCredential
	}
	errAssetCredential = json.Unmarshal([]byte(serializedAssetCredential), &toSerialize)
	if errAssetCredential != nil {
		return map[string]interface{}{}, errAssetCredential
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.IsPassphraseSet) {
		toSerialize["IsPassphraseSet"] = o.IsPassphraseSet
	}
	if !IsNil(o.IsPrivateKeySet) {
		toSerialize["IsPrivateKeySet"] = o.IsPrivateKeySet
	}
	if !IsNil(o.Passphrase) {
		toSerialize["Passphrase"] = o.Passphrase
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["PrivateKey"] = o.PrivateKey
	}
	if !IsNil(o.Username) {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetPrivateKeyCredential) UnmarshalJSON(data []byte) (err error) {
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
	type AssetPrivateKeyCredentialWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'passphrase' property has been set.
		IsPassphraseSet *bool `json:"IsPassphraseSet,omitempty"`
		// Indicates whether the value of the 'privateKey' property has been set.
		IsPrivateKeySet *bool `json:"IsPrivateKeySet,omitempty"`
		// The passphrase associated with the private key - Optional.
		Passphrase *string `json:"Passphrase,omitempty"`
		// The private key used to authenticate with a managed target. The corresponding public key needs to be added in the auth list of the remote endpoint.
		PrivateKey *string `json:"PrivateKey,omitempty"`
		// The username used to authenticate with a managed target.
		Username *string `json:"Username,omitempty"`
	}

	varAssetPrivateKeyCredentialWithoutEmbeddedStruct := AssetPrivateKeyCredentialWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetPrivateKeyCredentialWithoutEmbeddedStruct)
	if err == nil {
		varAssetPrivateKeyCredential := _AssetPrivateKeyCredential{}
		varAssetPrivateKeyCredential.ClassId = varAssetPrivateKeyCredentialWithoutEmbeddedStruct.ClassId
		varAssetPrivateKeyCredential.ObjectType = varAssetPrivateKeyCredentialWithoutEmbeddedStruct.ObjectType
		varAssetPrivateKeyCredential.IsPassphraseSet = varAssetPrivateKeyCredentialWithoutEmbeddedStruct.IsPassphraseSet
		varAssetPrivateKeyCredential.IsPrivateKeySet = varAssetPrivateKeyCredentialWithoutEmbeddedStruct.IsPrivateKeySet
		varAssetPrivateKeyCredential.Passphrase = varAssetPrivateKeyCredentialWithoutEmbeddedStruct.Passphrase
		varAssetPrivateKeyCredential.PrivateKey = varAssetPrivateKeyCredentialWithoutEmbeddedStruct.PrivateKey
		varAssetPrivateKeyCredential.Username = varAssetPrivateKeyCredentialWithoutEmbeddedStruct.Username
		*o = AssetPrivateKeyCredential(varAssetPrivateKeyCredential)
	} else {
		return err
	}

	varAssetPrivateKeyCredential := _AssetPrivateKeyCredential{}

	err = json.Unmarshal(data, &varAssetPrivateKeyCredential)
	if err == nil {
		o.AssetCredential = varAssetPrivateKeyCredential.AssetCredential
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPassphraseSet")
		delete(additionalProperties, "IsPrivateKeySet")
		delete(additionalProperties, "Passphrase")
		delete(additionalProperties, "PrivateKey")
		delete(additionalProperties, "Username")

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

type NullableAssetPrivateKeyCredential struct {
	value *AssetPrivateKeyCredential
	isSet bool
}

func (v NullableAssetPrivateKeyCredential) Get() *AssetPrivateKeyCredential {
	return v.value
}

func (v *NullableAssetPrivateKeyCredential) Set(val *AssetPrivateKeyCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetPrivateKeyCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetPrivateKeyCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetPrivateKeyCredential(val *AssetPrivateKeyCredential) *NullableAssetPrivateKeyCredential {
	return &NullableAssetPrivateKeyCredential{value: val, isSet: true}
}

func (v NullableAssetPrivateKeyCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetPrivateKeyCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
