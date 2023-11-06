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

// TelemetryDruidExpressionPostAggregatorAllOf struct for TelemetryDruidExpressionPostAggregatorAllOf
type TelemetryDruidExpressionPostAggregatorAllOf struct {
	// Output name for the return value of the expression.
	Name *string `json:"name,omitempty"`
	// The Druid expression.
	Expression *string `json:"expression,omitempty"`
	// Expression post-aggregators may specify an ordering, which defines the order of resulting values when sorting results. If no ordering (or null) is specified, the default floating point ordering is used. numericFirst ordering always returns finite values first, followed by NaN, and infinite values last.
	Ordering             *string `json:"ordering,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidExpressionPostAggregatorAllOf TelemetryDruidExpressionPostAggregatorAllOf

// NewTelemetryDruidExpressionPostAggregatorAllOf instantiates a new TelemetryDruidExpressionPostAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidExpressionPostAggregatorAllOf() *TelemetryDruidExpressionPostAggregatorAllOf {
	this := TelemetryDruidExpressionPostAggregatorAllOf{}
	return &this
}

// NewTelemetryDruidExpressionPostAggregatorAllOfWithDefaults instantiates a new TelemetryDruidExpressionPostAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidExpressionPostAggregatorAllOfWithDefaults() *TelemetryDruidExpressionPostAggregatorAllOf {
	this := TelemetryDruidExpressionPostAggregatorAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) SetExpression(v string) {
	o.Expression = &v
}

// GetOrdering returns the Ordering field value if set, zero value otherwise.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetOrdering() string {
	if o == nil || o.Ordering == nil {
		var ret string
		return ret
	}
	return *o.Ordering
}

// GetOrderingOk returns a tuple with the Ordering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) GetOrderingOk() (*string, bool) {
	if o == nil || o.Ordering == nil {
		return nil, false
	}
	return o.Ordering, true
}

// HasOrdering returns a boolean if a field has been set.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) HasOrdering() bool {
	if o != nil && o.Ordering != nil {
		return true
	}

	return false
}

// SetOrdering gets a reference to the given string and assigns it to the Ordering field.
func (o *TelemetryDruidExpressionPostAggregatorAllOf) SetOrdering(v string) {
	o.Ordering = &v
}

func (o TelemetryDruidExpressionPostAggregatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Ordering != nil {
		toSerialize["ordering"] = o.Ordering
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidExpressionPostAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidExpressionPostAggregatorAllOf := _TelemetryDruidExpressionPostAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidExpressionPostAggregatorAllOf); err == nil {
		*o = TelemetryDruidExpressionPostAggregatorAllOf(varTelemetryDruidExpressionPostAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "expression")
		delete(additionalProperties, "ordering")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidExpressionPostAggregatorAllOf struct {
	value *TelemetryDruidExpressionPostAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidExpressionPostAggregatorAllOf) Get() *TelemetryDruidExpressionPostAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidExpressionPostAggregatorAllOf) Set(val *TelemetryDruidExpressionPostAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidExpressionPostAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidExpressionPostAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidExpressionPostAggregatorAllOf(val *TelemetryDruidExpressionPostAggregatorAllOf) *NullableTelemetryDruidExpressionPostAggregatorAllOf {
	return &NullableTelemetryDruidExpressionPostAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidExpressionPostAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidExpressionPostAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
