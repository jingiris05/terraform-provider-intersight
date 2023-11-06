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

// MoTagSummaryAllOf struct for MoTagSummaryAllOf
type MoTagSummaryAllOf struct {
	// A discriminator value to disambiguate the schema of a HTTP GET response body.
	ObjectType           *string           `json:"ObjectType,omitempty"`
	Results              []MoTagKeySummary `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MoTagSummaryAllOf MoTagSummaryAllOf

// NewMoTagSummaryAllOf instantiates a new MoTagSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoTagSummaryAllOf() *MoTagSummaryAllOf {
	this := MoTagSummaryAllOf{}
	return &this
}

// NewMoTagSummaryAllOfWithDefaults instantiates a new MoTagSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoTagSummaryAllOfWithDefaults() *MoTagSummaryAllOf {
	this := MoTagSummaryAllOf{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *MoTagSummaryAllOf) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoTagSummaryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *MoTagSummaryAllOf) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *MoTagSummaryAllOf) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoTagSummaryAllOf) GetResults() []MoTagKeySummary {
	if o == nil {
		var ret []MoTagKeySummary
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoTagSummaryAllOf) GetResultsOk() ([]MoTagKeySummary, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MoTagSummaryAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MoTagKeySummary and assigns it to the Results field.
func (o *MoTagSummaryAllOf) SetResults(v []MoTagKeySummary) {
	o.Results = v
}

func (o MoTagSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MoTagSummaryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMoTagSummaryAllOf := _MoTagSummaryAllOf{}

	if err = json.Unmarshal(bytes, &varMoTagSummaryAllOf); err == nil {
		*o = MoTagSummaryAllOf(varMoTagSummaryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMoTagSummaryAllOf struct {
	value *MoTagSummaryAllOf
	isSet bool
}

func (v NullableMoTagSummaryAllOf) Get() *MoTagSummaryAllOf {
	return v.value
}

func (v *NullableMoTagSummaryAllOf) Set(val *MoTagSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMoTagSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMoTagSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoTagSummaryAllOf(val *MoTagSummaryAllOf) *NullableMoTagSummaryAllOf {
	return &NullableMoTagSummaryAllOf{value: val, isSet: true}
}

func (v NullableMoTagSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoTagSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
