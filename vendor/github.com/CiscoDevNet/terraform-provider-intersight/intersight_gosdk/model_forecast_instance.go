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
	"time"
)

// checks if the ForecastInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForecastInstance{}

// ForecastInstance Entity representing forecast result for instance of managed object, ie, data source.
type ForecastInstance struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action to be triggered on forecast instance. Default value is None. * `None` - The Enum value None represents that no action is triggered on the forecast Instance managed object. * `Evaluate` - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered.
	Action   *string   `json:"Action,omitempty"`
	AltModel []float32 `json:"AltModel,omitempty"`
	// The time interval (in days) for the data to be used for computing forecast model.
	DataInterval *int64 `json:"DataInterval,omitempty"`
	// The start date from when the data should be used for computing forecast model.
	DataStartDate *time.Time `json:"DataStartDate,omitempty"`
	// The Moid of the Intersight managed device instance for which regression model is derived.
	DeviceId *string `json:"DeviceId,omitempty"`
	// The number of days remaining before the device reaches its full functional capacity.
	FullCapDays *int64 `json:"FullCapDays,omitempty"`
	// The time when the forecast model was last updated.
	LastModelUpdateTime *time.Time `json:"LastModelUpdateTime,omitempty"`
	// The name of the metric for which regression model is generated.
	MetricName *string `json:"MetricName,omitempty"`
	// The minimum number of days the HyperFlex cluster should be up for computing forecast.
	MinDaysForForecast *int64                `json:"MinDaysForForecast,omitempty"`
	Model              NullableForecastModel `json:"Model,omitempty"`
	// The number of days remaining before the device reaches the specified threshold for the metric as defined in definition.
	ThresholdDays        *int64                                      `json:"ThresholdDays,omitempty"`
	ForecastDef          NullableForecastDefinitionRelationship      `json:"ForecastDef,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ForecastInstance ForecastInstance

// NewForecastInstance instantiates a new ForecastInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecastInstance(classId string, objectType string) *ForecastInstance {
	this := ForecastInstance{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var dataInterval int64 = 180
	this.DataInterval = &dataInterval
	return &this
}

// NewForecastInstanceWithDefaults instantiates a new ForecastInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastInstanceWithDefaults() *ForecastInstance {
	this := ForecastInstance{}
	var classId string = "forecast.Instance"
	this.ClassId = classId
	var objectType string = "forecast.Instance"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var dataInterval int64 = 180
	this.DataInterval = &dataInterval
	return &this
}

// GetClassId returns the ClassId field value
func (o *ForecastInstance) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ForecastInstance) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "forecast.Instance" of the ClassId field.
func (o *ForecastInstance) GetDefaultClassId() interface{} {
	return "forecast.Instance"
}

// GetObjectType returns the ObjectType field value
func (o *ForecastInstance) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ForecastInstance) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "forecast.Instance" of the ObjectType field.
func (o *ForecastInstance) GetDefaultObjectType() interface{} {
	return "forecast.Instance"
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ForecastInstance) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ForecastInstance) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ForecastInstance) SetAction(v string) {
	o.Action = &v
}

// GetAltModel returns the AltModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastInstance) GetAltModel() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}
	return o.AltModel
}

// GetAltModelOk returns a tuple with the AltModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastInstance) GetAltModelOk() ([]float32, bool) {
	if o == nil || IsNil(o.AltModel) {
		return nil, false
	}
	return o.AltModel, true
}

// HasAltModel returns a boolean if a field has been set.
func (o *ForecastInstance) HasAltModel() bool {
	if o != nil && !IsNil(o.AltModel) {
		return true
	}

	return false
}

// SetAltModel gets a reference to the given []float32 and assigns it to the AltModel field.
func (o *ForecastInstance) SetAltModel(v []float32) {
	o.AltModel = v
}

// GetDataInterval returns the DataInterval field value if set, zero value otherwise.
func (o *ForecastInstance) GetDataInterval() int64 {
	if o == nil || IsNil(o.DataInterval) {
		var ret int64
		return ret
	}
	return *o.DataInterval
}

// GetDataIntervalOk returns a tuple with the DataInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetDataIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.DataInterval) {
		return nil, false
	}
	return o.DataInterval, true
}

// HasDataInterval returns a boolean if a field has been set.
func (o *ForecastInstance) HasDataInterval() bool {
	if o != nil && !IsNil(o.DataInterval) {
		return true
	}

	return false
}

// SetDataInterval gets a reference to the given int64 and assigns it to the DataInterval field.
func (o *ForecastInstance) SetDataInterval(v int64) {
	o.DataInterval = &v
}

// GetDataStartDate returns the DataStartDate field value if set, zero value otherwise.
func (o *ForecastInstance) GetDataStartDate() time.Time {
	if o == nil || IsNil(o.DataStartDate) {
		var ret time.Time
		return ret
	}
	return *o.DataStartDate
}

// GetDataStartDateOk returns a tuple with the DataStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetDataStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DataStartDate) {
		return nil, false
	}
	return o.DataStartDate, true
}

// HasDataStartDate returns a boolean if a field has been set.
func (o *ForecastInstance) HasDataStartDate() bool {
	if o != nil && !IsNil(o.DataStartDate) {
		return true
	}

	return false
}

// SetDataStartDate gets a reference to the given time.Time and assigns it to the DataStartDate field.
func (o *ForecastInstance) SetDataStartDate(v time.Time) {
	o.DataStartDate = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ForecastInstance) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ForecastInstance) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ForecastInstance) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetFullCapDays returns the FullCapDays field value if set, zero value otherwise.
func (o *ForecastInstance) GetFullCapDays() int64 {
	if o == nil || IsNil(o.FullCapDays) {
		var ret int64
		return ret
	}
	return *o.FullCapDays
}

// GetFullCapDaysOk returns a tuple with the FullCapDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetFullCapDaysOk() (*int64, bool) {
	if o == nil || IsNil(o.FullCapDays) {
		return nil, false
	}
	return o.FullCapDays, true
}

// HasFullCapDays returns a boolean if a field has been set.
func (o *ForecastInstance) HasFullCapDays() bool {
	if o != nil && !IsNil(o.FullCapDays) {
		return true
	}

	return false
}

// SetFullCapDays gets a reference to the given int64 and assigns it to the FullCapDays field.
func (o *ForecastInstance) SetFullCapDays(v int64) {
	o.FullCapDays = &v
}

// GetLastModelUpdateTime returns the LastModelUpdateTime field value if set, zero value otherwise.
func (o *ForecastInstance) GetLastModelUpdateTime() time.Time {
	if o == nil || IsNil(o.LastModelUpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModelUpdateTime
}

// GetLastModelUpdateTimeOk returns a tuple with the LastModelUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetLastModelUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModelUpdateTime) {
		return nil, false
	}
	return o.LastModelUpdateTime, true
}

// HasLastModelUpdateTime returns a boolean if a field has been set.
func (o *ForecastInstance) HasLastModelUpdateTime() bool {
	if o != nil && !IsNil(o.LastModelUpdateTime) {
		return true
	}

	return false
}

// SetLastModelUpdateTime gets a reference to the given time.Time and assigns it to the LastModelUpdateTime field.
func (o *ForecastInstance) SetLastModelUpdateTime(v time.Time) {
	o.LastModelUpdateTime = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *ForecastInstance) GetMetricName() string {
	if o == nil || IsNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricName) {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *ForecastInstance) HasMetricName() bool {
	if o != nil && !IsNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *ForecastInstance) SetMetricName(v string) {
	o.MetricName = &v
}

// GetMinDaysForForecast returns the MinDaysForForecast field value if set, zero value otherwise.
func (o *ForecastInstance) GetMinDaysForForecast() int64 {
	if o == nil || IsNil(o.MinDaysForForecast) {
		var ret int64
		return ret
	}
	return *o.MinDaysForForecast
}

// GetMinDaysForForecastOk returns a tuple with the MinDaysForForecast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetMinDaysForForecastOk() (*int64, bool) {
	if o == nil || IsNil(o.MinDaysForForecast) {
		return nil, false
	}
	return o.MinDaysForForecast, true
}

// HasMinDaysForForecast returns a boolean if a field has been set.
func (o *ForecastInstance) HasMinDaysForForecast() bool {
	if o != nil && !IsNil(o.MinDaysForForecast) {
		return true
	}

	return false
}

// SetMinDaysForForecast gets a reference to the given int64 and assigns it to the MinDaysForForecast field.
func (o *ForecastInstance) SetMinDaysForForecast(v int64) {
	o.MinDaysForForecast = &v
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastInstance) GetModel() ForecastModel {
	if o == nil || IsNil(o.Model.Get()) {
		var ret ForecastModel
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastInstance) GetModelOk() (*ForecastModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *ForecastInstance) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableForecastModel and assigns it to the Model field.
func (o *ForecastInstance) SetModel(v ForecastModel) {
	o.Model.Set(&v)
}

// SetModelNil sets the value for Model to be an explicit nil
func (o *ForecastInstance) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *ForecastInstance) UnsetModel() {
	o.Model.Unset()
}

// GetThresholdDays returns the ThresholdDays field value if set, zero value otherwise.
func (o *ForecastInstance) GetThresholdDays() int64 {
	if o == nil || IsNil(o.ThresholdDays) {
		var ret int64
		return ret
	}
	return *o.ThresholdDays
}

// GetThresholdDaysOk returns a tuple with the ThresholdDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstance) GetThresholdDaysOk() (*int64, bool) {
	if o == nil || IsNil(o.ThresholdDays) {
		return nil, false
	}
	return o.ThresholdDays, true
}

// HasThresholdDays returns a boolean if a field has been set.
func (o *ForecastInstance) HasThresholdDays() bool {
	if o != nil && !IsNil(o.ThresholdDays) {
		return true
	}

	return false
}

// SetThresholdDays gets a reference to the given int64 and assigns it to the ThresholdDays field.
func (o *ForecastInstance) SetThresholdDays(v int64) {
	o.ThresholdDays = &v
}

// GetForecastDef returns the ForecastDef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastInstance) GetForecastDef() ForecastDefinitionRelationship {
	if o == nil || IsNil(o.ForecastDef.Get()) {
		var ret ForecastDefinitionRelationship
		return ret
	}
	return *o.ForecastDef.Get()
}

// GetForecastDefOk returns a tuple with the ForecastDef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastInstance) GetForecastDefOk() (*ForecastDefinitionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForecastDef.Get(), o.ForecastDef.IsSet()
}

// HasForecastDef returns a boolean if a field has been set.
func (o *ForecastInstance) HasForecastDef() bool {
	if o != nil && o.ForecastDef.IsSet() {
		return true
	}

	return false
}

// SetForecastDef gets a reference to the given NullableForecastDefinitionRelationship and assigns it to the ForecastDef field.
func (o *ForecastInstance) SetForecastDef(v ForecastDefinitionRelationship) {
	o.ForecastDef.Set(&v)
}

// SetForecastDefNil sets the value for ForecastDef to be an explicit nil
func (o *ForecastInstance) SetForecastDefNil() {
	o.ForecastDef.Set(nil)
}

// UnsetForecastDef ensures that no value is present for ForecastDef, not even an explicit nil
func (o *ForecastInstance) UnsetForecastDef() {
	o.ForecastDef.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastInstance) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastInstance) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ForecastInstance) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ForecastInstance) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *ForecastInstance) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *ForecastInstance) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o ForecastInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForecastInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Action) {
		toSerialize["Action"] = o.Action
	}
	if o.AltModel != nil {
		toSerialize["AltModel"] = o.AltModel
	}
	if !IsNil(o.DataInterval) {
		toSerialize["DataInterval"] = o.DataInterval
	}
	if !IsNil(o.DataStartDate) {
		toSerialize["DataStartDate"] = o.DataStartDate
	}
	if !IsNil(o.DeviceId) {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if !IsNil(o.FullCapDays) {
		toSerialize["FullCapDays"] = o.FullCapDays
	}
	if !IsNil(o.LastModelUpdateTime) {
		toSerialize["LastModelUpdateTime"] = o.LastModelUpdateTime
	}
	if !IsNil(o.MetricName) {
		toSerialize["MetricName"] = o.MetricName
	}
	if !IsNil(o.MinDaysForForecast) {
		toSerialize["MinDaysForForecast"] = o.MinDaysForForecast
	}
	if o.Model.IsSet() {
		toSerialize["Model"] = o.Model.Get()
	}
	if !IsNil(o.ThresholdDays) {
		toSerialize["ThresholdDays"] = o.ThresholdDays
	}
	if o.ForecastDef.IsSet() {
		toSerialize["ForecastDef"] = o.ForecastDef.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ForecastInstance) UnmarshalJSON(data []byte) (err error) {
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
	type ForecastInstanceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Action to be triggered on forecast instance. Default value is None. * `None` - The Enum value None represents that no action is triggered on the forecast Instance managed object. * `Evaluate` - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered.
		Action   *string   `json:"Action,omitempty"`
		AltModel []float32 `json:"AltModel,omitempty"`
		// The time interval (in days) for the data to be used for computing forecast model.
		DataInterval *int64 `json:"DataInterval,omitempty"`
		// The start date from when the data should be used for computing forecast model.
		DataStartDate *time.Time `json:"DataStartDate,omitempty"`
		// The Moid of the Intersight managed device instance for which regression model is derived.
		DeviceId *string `json:"DeviceId,omitempty"`
		// The number of days remaining before the device reaches its full functional capacity.
		FullCapDays *int64 `json:"FullCapDays,omitempty"`
		// The time when the forecast model was last updated.
		LastModelUpdateTime *time.Time `json:"LastModelUpdateTime,omitempty"`
		// The name of the metric for which regression model is generated.
		MetricName *string `json:"MetricName,omitempty"`
		// The minimum number of days the HyperFlex cluster should be up for computing forecast.
		MinDaysForForecast *int64                `json:"MinDaysForForecast,omitempty"`
		Model              NullableForecastModel `json:"Model,omitempty"`
		// The number of days remaining before the device reaches the specified threshold for the metric as defined in definition.
		ThresholdDays    *int64                                      `json:"ThresholdDays,omitempty"`
		ForecastDef      NullableForecastDefinitionRelationship      `json:"ForecastDef,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varForecastInstanceWithoutEmbeddedStruct := ForecastInstanceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varForecastInstanceWithoutEmbeddedStruct)
	if err == nil {
		varForecastInstance := _ForecastInstance{}
		varForecastInstance.ClassId = varForecastInstanceWithoutEmbeddedStruct.ClassId
		varForecastInstance.ObjectType = varForecastInstanceWithoutEmbeddedStruct.ObjectType
		varForecastInstance.Action = varForecastInstanceWithoutEmbeddedStruct.Action
		varForecastInstance.AltModel = varForecastInstanceWithoutEmbeddedStruct.AltModel
		varForecastInstance.DataInterval = varForecastInstanceWithoutEmbeddedStruct.DataInterval
		varForecastInstance.DataStartDate = varForecastInstanceWithoutEmbeddedStruct.DataStartDate
		varForecastInstance.DeviceId = varForecastInstanceWithoutEmbeddedStruct.DeviceId
		varForecastInstance.FullCapDays = varForecastInstanceWithoutEmbeddedStruct.FullCapDays
		varForecastInstance.LastModelUpdateTime = varForecastInstanceWithoutEmbeddedStruct.LastModelUpdateTime
		varForecastInstance.MetricName = varForecastInstanceWithoutEmbeddedStruct.MetricName
		varForecastInstance.MinDaysForForecast = varForecastInstanceWithoutEmbeddedStruct.MinDaysForForecast
		varForecastInstance.Model = varForecastInstanceWithoutEmbeddedStruct.Model
		varForecastInstance.ThresholdDays = varForecastInstanceWithoutEmbeddedStruct.ThresholdDays
		varForecastInstance.ForecastDef = varForecastInstanceWithoutEmbeddedStruct.ForecastDef
		varForecastInstance.RegisteredDevice = varForecastInstanceWithoutEmbeddedStruct.RegisteredDevice
		*o = ForecastInstance(varForecastInstance)
	} else {
		return err
	}

	varForecastInstance := _ForecastInstance{}

	err = json.Unmarshal(data, &varForecastInstance)
	if err == nil {
		o.MoBaseMo = varForecastInstance.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "AltModel")
		delete(additionalProperties, "DataInterval")
		delete(additionalProperties, "DataStartDate")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "FullCapDays")
		delete(additionalProperties, "LastModelUpdateTime")
		delete(additionalProperties, "MetricName")
		delete(additionalProperties, "MinDaysForForecast")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "ThresholdDays")
		delete(additionalProperties, "ForecastDef")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableForecastInstance struct {
	value *ForecastInstance
	isSet bool
}

func (v NullableForecastInstance) Get() *ForecastInstance {
	return v.value
}

func (v *NullableForecastInstance) Set(val *ForecastInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastInstance(val *ForecastInstance) *NullableForecastInstance {
	return &NullableForecastInstance{value: val, isSet: true}
}

func (v NullableForecastInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForecastInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
