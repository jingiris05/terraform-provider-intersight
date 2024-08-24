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

// checks if the PartnerintegrationMetricsModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnerintegrationMetricsModel{}

// PartnerintegrationMetricsModel Transformation model in metrics format.
type PartnerintegrationMetricsModel struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Transformation attributes model in yaml format.
	Attributes interface{} `json:"Attributes,omitempty"`
	// Transformation druid instrument exporters model in yaml format.
	DruidInstrumentExporters interface{} `json:"DruidInstrumentExporters,omitempty"`
	// Transformation instruments model in yaml format.
	Instruments interface{} `json:"Instruments,omitempty"`
	// Transformation meter providers model in yaml format.
	MeterProviders interface{} `json:"MeterProviders,omitempty"`
	// Transformation metrics model in yaml format.
	Metrics              interface{} `json:"Metrics,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerintegrationMetricsModel PartnerintegrationMetricsModel

// NewPartnerintegrationMetricsModel instantiates a new PartnerintegrationMetricsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerintegrationMetricsModel(classId string, objectType string) *PartnerintegrationMetricsModel {
	this := PartnerintegrationMetricsModel{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPartnerintegrationMetricsModelWithDefaults instantiates a new PartnerintegrationMetricsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerintegrationMetricsModelWithDefaults() *PartnerintegrationMetricsModel {
	this := PartnerintegrationMetricsModel{}
	var classId string = "partnerintegration.MetricsModel"
	this.ClassId = classId
	var objectType string = "partnerintegration.MetricsModel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PartnerintegrationMetricsModel) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationMetricsModel) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PartnerintegrationMetricsModel) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "partnerintegration.MetricsModel" of the ClassId field.
func (o *PartnerintegrationMetricsModel) GetDefaultClassId() interface{} {
	return "partnerintegration.MetricsModel"
}

// GetObjectType returns the ObjectType field value
func (o *PartnerintegrationMetricsModel) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationMetricsModel) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PartnerintegrationMetricsModel) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "partnerintegration.MetricsModel" of the ObjectType field.
func (o *PartnerintegrationMetricsModel) GetDefaultObjectType() interface{} {
	return "partnerintegration.MetricsModel"
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationMetricsModel) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationMetricsModel) GetAttributesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PartnerintegrationMetricsModel) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *PartnerintegrationMetricsModel) SetAttributes(v interface{}) {
	o.Attributes = v
}

// GetDruidInstrumentExporters returns the DruidInstrumentExporters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationMetricsModel) GetDruidInstrumentExporters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DruidInstrumentExporters
}

// GetDruidInstrumentExportersOk returns a tuple with the DruidInstrumentExporters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationMetricsModel) GetDruidInstrumentExportersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DruidInstrumentExporters) {
		return nil, false
	}
	return &o.DruidInstrumentExporters, true
}

// HasDruidInstrumentExporters returns a boolean if a field has been set.
func (o *PartnerintegrationMetricsModel) HasDruidInstrumentExporters() bool {
	if o != nil && !IsNil(o.DruidInstrumentExporters) {
		return true
	}

	return false
}

// SetDruidInstrumentExporters gets a reference to the given interface{} and assigns it to the DruidInstrumentExporters field.
func (o *PartnerintegrationMetricsModel) SetDruidInstrumentExporters(v interface{}) {
	o.DruidInstrumentExporters = v
}

// GetInstruments returns the Instruments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationMetricsModel) GetInstruments() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Instruments
}

// GetInstrumentsOk returns a tuple with the Instruments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationMetricsModel) GetInstrumentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Instruments) {
		return nil, false
	}
	return &o.Instruments, true
}

// HasInstruments returns a boolean if a field has been set.
func (o *PartnerintegrationMetricsModel) HasInstruments() bool {
	if o != nil && !IsNil(o.Instruments) {
		return true
	}

	return false
}

// SetInstruments gets a reference to the given interface{} and assigns it to the Instruments field.
func (o *PartnerintegrationMetricsModel) SetInstruments(v interface{}) {
	o.Instruments = v
}

// GetMeterProviders returns the MeterProviders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationMetricsModel) GetMeterProviders() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MeterProviders
}

// GetMeterProvidersOk returns a tuple with the MeterProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationMetricsModel) GetMeterProvidersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MeterProviders) {
		return nil, false
	}
	return &o.MeterProviders, true
}

// HasMeterProviders returns a boolean if a field has been set.
func (o *PartnerintegrationMetricsModel) HasMeterProviders() bool {
	if o != nil && !IsNil(o.MeterProviders) {
		return true
	}

	return false
}

// SetMeterProviders gets a reference to the given interface{} and assigns it to the MeterProviders field.
func (o *PartnerintegrationMetricsModel) SetMeterProviders(v interface{}) {
	o.MeterProviders = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationMetricsModel) GetMetrics() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationMetricsModel) GetMetricsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *PartnerintegrationMetricsModel) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given interface{} and assigns it to the Metrics field.
func (o *PartnerintegrationMetricsModel) SetMetrics(v interface{}) {
	o.Metrics = v
}

func (o PartnerintegrationMetricsModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnerintegrationMetricsModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Attributes != nil {
		toSerialize["Attributes"] = o.Attributes
	}
	if o.DruidInstrumentExporters != nil {
		toSerialize["DruidInstrumentExporters"] = o.DruidInstrumentExporters
	}
	if o.Instruments != nil {
		toSerialize["Instruments"] = o.Instruments
	}
	if o.MeterProviders != nil {
		toSerialize["MeterProviders"] = o.MeterProviders
	}
	if o.Metrics != nil {
		toSerialize["Metrics"] = o.Metrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PartnerintegrationMetricsModel) UnmarshalJSON(data []byte) (err error) {
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
	type PartnerintegrationMetricsModelWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Transformation attributes model in yaml format.
		Attributes interface{} `json:"Attributes,omitempty"`
		// Transformation druid instrument exporters model in yaml format.
		DruidInstrumentExporters interface{} `json:"DruidInstrumentExporters,omitempty"`
		// Transformation instruments model in yaml format.
		Instruments interface{} `json:"Instruments,omitempty"`
		// Transformation meter providers model in yaml format.
		MeterProviders interface{} `json:"MeterProviders,omitempty"`
		// Transformation metrics model in yaml format.
		Metrics interface{} `json:"Metrics,omitempty"`
	}

	varPartnerintegrationMetricsModelWithoutEmbeddedStruct := PartnerintegrationMetricsModelWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPartnerintegrationMetricsModelWithoutEmbeddedStruct)
	if err == nil {
		varPartnerintegrationMetricsModel := _PartnerintegrationMetricsModel{}
		varPartnerintegrationMetricsModel.ClassId = varPartnerintegrationMetricsModelWithoutEmbeddedStruct.ClassId
		varPartnerintegrationMetricsModel.ObjectType = varPartnerintegrationMetricsModelWithoutEmbeddedStruct.ObjectType
		varPartnerintegrationMetricsModel.Attributes = varPartnerintegrationMetricsModelWithoutEmbeddedStruct.Attributes
		varPartnerintegrationMetricsModel.DruidInstrumentExporters = varPartnerintegrationMetricsModelWithoutEmbeddedStruct.DruidInstrumentExporters
		varPartnerintegrationMetricsModel.Instruments = varPartnerintegrationMetricsModelWithoutEmbeddedStruct.Instruments
		varPartnerintegrationMetricsModel.MeterProviders = varPartnerintegrationMetricsModelWithoutEmbeddedStruct.MeterProviders
		varPartnerintegrationMetricsModel.Metrics = varPartnerintegrationMetricsModelWithoutEmbeddedStruct.Metrics
		*o = PartnerintegrationMetricsModel(varPartnerintegrationMetricsModel)
	} else {
		return err
	}

	varPartnerintegrationMetricsModel := _PartnerintegrationMetricsModel{}

	err = json.Unmarshal(data, &varPartnerintegrationMetricsModel)
	if err == nil {
		o.MoBaseComplexType = varPartnerintegrationMetricsModel.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Attributes")
		delete(additionalProperties, "DruidInstrumentExporters")
		delete(additionalProperties, "Instruments")
		delete(additionalProperties, "MeterProviders")
		delete(additionalProperties, "Metrics")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullablePartnerintegrationMetricsModel struct {
	value *PartnerintegrationMetricsModel
	isSet bool
}

func (v NullablePartnerintegrationMetricsModel) Get() *PartnerintegrationMetricsModel {
	return v.value
}

func (v *NullablePartnerintegrationMetricsModel) Set(val *PartnerintegrationMetricsModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationMetricsModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationMetricsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationMetricsModel(val *PartnerintegrationMetricsModel) *NullablePartnerintegrationMetricsModel {
	return &NullablePartnerintegrationMetricsModel{value: val, isSet: true}
}

func (v NullablePartnerintegrationMetricsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationMetricsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
