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
)

// AssetOauthBearerTokenCredentialAllOf Definition of the list of properties defined in 'asset.OauthBearerTokenCredential', excluding properties defined in parent classes.
type AssetOauthBearerTokenCredentialAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'token' property has been set.
	IsTokenSet *bool `json:"IsTokenSet,omitempty"`
	// Scope type for the crendetial i.e. User, Organization, Team.
	ScopeType *string `json:"ScopeType,omitempty"`
	// Scope value for the credential i.e. username, organization name or team name.
	ScopeValue *string `json:"ScopeValue,omitempty"`
	// The token used to authenticate with a managed target.
	Token                *string `json:"Token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetOauthBearerTokenCredentialAllOf AssetOauthBearerTokenCredentialAllOf

// NewAssetOauthBearerTokenCredentialAllOf instantiates a new AssetOauthBearerTokenCredentialAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOauthBearerTokenCredentialAllOf(classId string, objectType string) *AssetOauthBearerTokenCredentialAllOf {
	this := AssetOauthBearerTokenCredentialAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isTokenSet bool = false
	this.IsTokenSet = &isTokenSet
	return &this
}

// NewAssetOauthBearerTokenCredentialAllOfWithDefaults instantiates a new AssetOauthBearerTokenCredentialAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOauthBearerTokenCredentialAllOfWithDefaults() *AssetOauthBearerTokenCredentialAllOf {
	this := AssetOauthBearerTokenCredentialAllOf{}
	var classId string = "asset.OauthBearerTokenCredential"
	this.ClassId = classId
	var objectType string = "asset.OauthBearerTokenCredential"
	this.ObjectType = objectType
	var isTokenSet bool = false
	this.IsTokenSet = &isTokenSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetOauthBearerTokenCredentialAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredentialAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetOauthBearerTokenCredentialAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetOauthBearerTokenCredentialAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredentialAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetOauthBearerTokenCredentialAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsTokenSet returns the IsTokenSet field value if set, zero value otherwise.
func (o *AssetOauthBearerTokenCredentialAllOf) GetIsTokenSet() bool {
	if o == nil || o.IsTokenSet == nil {
		var ret bool
		return ret
	}
	return *o.IsTokenSet
}

// GetIsTokenSetOk returns a tuple with the IsTokenSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredentialAllOf) GetIsTokenSetOk() (*bool, bool) {
	if o == nil || o.IsTokenSet == nil {
		return nil, false
	}
	return o.IsTokenSet, true
}

// HasIsTokenSet returns a boolean if a field has been set.
func (o *AssetOauthBearerTokenCredentialAllOf) HasIsTokenSet() bool {
	if o != nil && o.IsTokenSet != nil {
		return true
	}

	return false
}

// SetIsTokenSet gets a reference to the given bool and assigns it to the IsTokenSet field.
func (o *AssetOauthBearerTokenCredentialAllOf) SetIsTokenSet(v bool) {
	o.IsTokenSet = &v
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise.
func (o *AssetOauthBearerTokenCredentialAllOf) GetScopeType() string {
	if o == nil || o.ScopeType == nil {
		var ret string
		return ret
	}
	return *o.ScopeType
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredentialAllOf) GetScopeTypeOk() (*string, bool) {
	if o == nil || o.ScopeType == nil {
		return nil, false
	}
	return o.ScopeType, true
}

// HasScopeType returns a boolean if a field has been set.
func (o *AssetOauthBearerTokenCredentialAllOf) HasScopeType() bool {
	if o != nil && o.ScopeType != nil {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given string and assigns it to the ScopeType field.
func (o *AssetOauthBearerTokenCredentialAllOf) SetScopeType(v string) {
	o.ScopeType = &v
}

// GetScopeValue returns the ScopeValue field value if set, zero value otherwise.
func (o *AssetOauthBearerTokenCredentialAllOf) GetScopeValue() string {
	if o == nil || o.ScopeValue == nil {
		var ret string
		return ret
	}
	return *o.ScopeValue
}

// GetScopeValueOk returns a tuple with the ScopeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredentialAllOf) GetScopeValueOk() (*string, bool) {
	if o == nil || o.ScopeValue == nil {
		return nil, false
	}
	return o.ScopeValue, true
}

// HasScopeValue returns a boolean if a field has been set.
func (o *AssetOauthBearerTokenCredentialAllOf) HasScopeValue() bool {
	if o != nil && o.ScopeValue != nil {
		return true
	}

	return false
}

// SetScopeValue gets a reference to the given string and assigns it to the ScopeValue field.
func (o *AssetOauthBearerTokenCredentialAllOf) SetScopeValue(v string) {
	o.ScopeValue = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AssetOauthBearerTokenCredentialAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthBearerTokenCredentialAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AssetOauthBearerTokenCredentialAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AssetOauthBearerTokenCredentialAllOf) SetToken(v string) {
	o.Token = &v
}

func (o AssetOauthBearerTokenCredentialAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsTokenSet != nil {
		toSerialize["IsTokenSet"] = o.IsTokenSet
	}
	if o.ScopeType != nil {
		toSerialize["ScopeType"] = o.ScopeType
	}
	if o.ScopeValue != nil {
		toSerialize["ScopeValue"] = o.ScopeValue
	}
	if o.Token != nil {
		toSerialize["Token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetOauthBearerTokenCredentialAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetOauthBearerTokenCredentialAllOf := _AssetOauthBearerTokenCredentialAllOf{}

	if err = json.Unmarshal(bytes, &varAssetOauthBearerTokenCredentialAllOf); err == nil {
		*o = AssetOauthBearerTokenCredentialAllOf(varAssetOauthBearerTokenCredentialAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsTokenSet")
		delete(additionalProperties, "ScopeType")
		delete(additionalProperties, "ScopeValue")
		delete(additionalProperties, "Token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetOauthBearerTokenCredentialAllOf struct {
	value *AssetOauthBearerTokenCredentialAllOf
	isSet bool
}

func (v NullableAssetOauthBearerTokenCredentialAllOf) Get() *AssetOauthBearerTokenCredentialAllOf {
	return v.value
}

func (v *NullableAssetOauthBearerTokenCredentialAllOf) Set(val *AssetOauthBearerTokenCredentialAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOauthBearerTokenCredentialAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOauthBearerTokenCredentialAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOauthBearerTokenCredentialAllOf(val *AssetOauthBearerTokenCredentialAllOf) *NullableAssetOauthBearerTokenCredentialAllOf {
	return &NullableAssetOauthBearerTokenCredentialAllOf{value: val, isSet: true}
}

func (v NullableAssetOauthBearerTokenCredentialAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOauthBearerTokenCredentialAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
