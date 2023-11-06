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

// TelemetryDruidBaseRequest The base structure for all Druid requests.
type TelemetryDruidBaseRequest struct {
	QueryType            string `json:"queryType"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidBaseRequest TelemetryDruidBaseRequest

// NewTelemetryDruidBaseRequest instantiates a new TelemetryDruidBaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidBaseRequest(queryType string) *TelemetryDruidBaseRequest {
	this := TelemetryDruidBaseRequest{}
	this.QueryType = queryType
	return &this
}

// NewTelemetryDruidBaseRequestWithDefaults instantiates a new TelemetryDruidBaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidBaseRequestWithDefaults() *TelemetryDruidBaseRequest {
	this := TelemetryDruidBaseRequest{}
	return &this
}

// GetQueryType returns the QueryType field value
func (o *TelemetryDruidBaseRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidBaseRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *TelemetryDruidBaseRequest) SetQueryType(v string) {
	o.QueryType = v
}

func (o TelemetryDruidBaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["queryType"] = o.QueryType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidBaseRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidBaseRequest := _TelemetryDruidBaseRequest{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidBaseRequest); err == nil {
		*o = TelemetryDruidBaseRequest(varTelemetryDruidBaseRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "queryType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidBaseRequest struct {
	value *TelemetryDruidBaseRequest
	isSet bool
}

func (v NullableTelemetryDruidBaseRequest) Get() *TelemetryDruidBaseRequest {
	return v.value
}

func (v *NullableTelemetryDruidBaseRequest) Set(val *TelemetryDruidBaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidBaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidBaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidBaseRequest(val *TelemetryDruidBaseRequest) *NullableTelemetryDruidBaseRequest {
	return &NullableTelemetryDruidBaseRequest{value: val, isSet: true}
}

func (v NullableTelemetryDruidBaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidBaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
