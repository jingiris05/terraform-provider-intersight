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
)

// checks if the TelemetryDruidSegmentMetadataResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidSegmentMetadataResult{}

// TelemetryDruidSegmentMetadataResult Metadata about a Druid segment
type TelemetryDruidSegmentMetadataResult struct {
	// An identifier for the metadata
	Id *string `json:"id,omitempty"`
	// String representation of the interval queried
	Intervals *string `json:"intervals,omitempty"`
	// A map of columns and their properties
	Columns map[string]interface{} `json:"columns,omitempty"`
	// A map of metrics and their properties
	Aggregators map[string]interface{} `json:"aggregators,omitempty"`
	// query granularity of data stored in segments, may be 'null'
	QueryGranularity map[string]interface{} `json:"queryGranularity,omitempty"`
	// estimated total segment byte size as if stored as text
	Size *int32 `json:"size,omitempty"`
	// number of rows stored in segment
	NumRows              *int32 `json:"numRows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidSegmentMetadataResult TelemetryDruidSegmentMetadataResult

// NewTelemetryDruidSegmentMetadataResult instantiates a new TelemetryDruidSegmentMetadataResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidSegmentMetadataResult() *TelemetryDruidSegmentMetadataResult {
	this := TelemetryDruidSegmentMetadataResult{}
	return &this
}

// NewTelemetryDruidSegmentMetadataResultWithDefaults instantiates a new TelemetryDruidSegmentMetadataResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidSegmentMetadataResultWithDefaults() *TelemetryDruidSegmentMetadataResult {
	this := TelemetryDruidSegmentMetadataResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TelemetryDruidSegmentMetadataResult) SetId(v string) {
	o.Id = &v
}

// GetIntervals returns the Intervals field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataResult) GetIntervals() string {
	if o == nil || IsNil(o.Intervals) {
		var ret string
		return ret
	}
	return *o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataResult) GetIntervalsOk() (*string, bool) {
	if o == nil || IsNil(o.Intervals) {
		return nil, false
	}
	return o.Intervals, true
}

// HasIntervals returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataResult) HasIntervals() bool {
	if o != nil && !IsNil(o.Intervals) {
		return true
	}

	return false
}

// SetIntervals gets a reference to the given string and assigns it to the Intervals field.
func (o *TelemetryDruidSegmentMetadataResult) SetIntervals(v string) {
	o.Intervals = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataResult) GetColumns() map[string]interface{} {
	if o == nil || IsNil(o.Columns) {
		var ret map[string]interface{}
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataResult) GetColumnsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Columns) {
		return map[string]interface{}{}, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataResult) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given map[string]interface{} and assigns it to the Columns field.
func (o *TelemetryDruidSegmentMetadataResult) SetColumns(v map[string]interface{}) {
	o.Columns = v
}

// GetAggregators returns the Aggregators field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataResult) GetAggregators() map[string]interface{} {
	if o == nil || IsNil(o.Aggregators) {
		var ret map[string]interface{}
		return ret
	}
	return o.Aggregators
}

// GetAggregatorsOk returns a tuple with the Aggregators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataResult) GetAggregatorsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Aggregators) {
		return map[string]interface{}{}, false
	}
	return o.Aggregators, true
}

// HasAggregators returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataResult) HasAggregators() bool {
	if o != nil && !IsNil(o.Aggregators) {
		return true
	}

	return false
}

// SetAggregators gets a reference to the given map[string]interface{} and assigns it to the Aggregators field.
func (o *TelemetryDruidSegmentMetadataResult) SetAggregators(v map[string]interface{}) {
	o.Aggregators = v
}

// GetQueryGranularity returns the QueryGranularity field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataResult) GetQueryGranularity() map[string]interface{} {
	if o == nil || IsNil(o.QueryGranularity) {
		var ret map[string]interface{}
		return ret
	}
	return o.QueryGranularity
}

// GetQueryGranularityOk returns a tuple with the QueryGranularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataResult) GetQueryGranularityOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.QueryGranularity) {
		return map[string]interface{}{}, false
	}
	return o.QueryGranularity, true
}

// HasQueryGranularity returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataResult) HasQueryGranularity() bool {
	if o != nil && !IsNil(o.QueryGranularity) {
		return true
	}

	return false
}

// SetQueryGranularity gets a reference to the given map[string]interface{} and assigns it to the QueryGranularity field.
func (o *TelemetryDruidSegmentMetadataResult) SetQueryGranularity(v map[string]interface{}) {
	o.QueryGranularity = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataResult) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataResult) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataResult) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *TelemetryDruidSegmentMetadataResult) SetSize(v int32) {
	o.Size = &v
}

// GetNumRows returns the NumRows field value if set, zero value otherwise.
func (o *TelemetryDruidSegmentMetadataResult) GetNumRows() int32 {
	if o == nil || IsNil(o.NumRows) {
		var ret int32
		return ret
	}
	return *o.NumRows
}

// GetNumRowsOk returns a tuple with the NumRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidSegmentMetadataResult) GetNumRowsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumRows) {
		return nil, false
	}
	return o.NumRows, true
}

// HasNumRows returns a boolean if a field has been set.
func (o *TelemetryDruidSegmentMetadataResult) HasNumRows() bool {
	if o != nil && !IsNil(o.NumRows) {
		return true
	}

	return false
}

// SetNumRows gets a reference to the given int32 and assigns it to the NumRows field.
func (o *TelemetryDruidSegmentMetadataResult) SetNumRows(v int32) {
	o.NumRows = &v
}

func (o TelemetryDruidSegmentMetadataResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidSegmentMetadataResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Intervals) {
		toSerialize["intervals"] = o.Intervals
	}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	if !IsNil(o.Aggregators) {
		toSerialize["aggregators"] = o.Aggregators
	}
	if !IsNil(o.QueryGranularity) {
		toSerialize["queryGranularity"] = o.QueryGranularity
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.NumRows) {
		toSerialize["numRows"] = o.NumRows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidSegmentMetadataResult) UnmarshalJSON(data []byte) (err error) {
	varTelemetryDruidSegmentMetadataResult := _TelemetryDruidSegmentMetadataResult{}

	err = json.Unmarshal(data, &varTelemetryDruidSegmentMetadataResult)

	if err != nil {
		return err
	}

	*o = TelemetryDruidSegmentMetadataResult(varTelemetryDruidSegmentMetadataResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "intervals")
		delete(additionalProperties, "columns")
		delete(additionalProperties, "aggregators")
		delete(additionalProperties, "queryGranularity")
		delete(additionalProperties, "size")
		delete(additionalProperties, "numRows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidSegmentMetadataResult struct {
	value *TelemetryDruidSegmentMetadataResult
	isSet bool
}

func (v NullableTelemetryDruidSegmentMetadataResult) Get() *TelemetryDruidSegmentMetadataResult {
	return v.value
}

func (v *NullableTelemetryDruidSegmentMetadataResult) Set(val *TelemetryDruidSegmentMetadataResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidSegmentMetadataResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidSegmentMetadataResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidSegmentMetadataResult(val *TelemetryDruidSegmentMetadataResult) *NullableTelemetryDruidSegmentMetadataResult {
	return &NullableTelemetryDruidSegmentMetadataResult{value: val, isSet: true}
}

func (v NullableTelemetryDruidSegmentMetadataResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidSegmentMetadataResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
