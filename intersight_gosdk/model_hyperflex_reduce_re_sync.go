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

// checks if the HyperflexReduceReSync type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexReduceReSync{}

// HyperflexReduceReSync The execution status of reduce re-sync and stale mirror cleanup for the HyperFlex cluster.
type HyperflexReduceReSync struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The HyperFlex reduce re-sync script execution completion status.
	CompletionStatus *bool `json:"CompletionStatus,omitempty"`
	// The execution output of the script.
	ExecutionOutput      *string                              `json:"ExecutionOutput,omitempty"`
	Cluster              NullableHyperflexClusterRelationship `json:"Cluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexReduceReSync HyperflexReduceReSync

// NewHyperflexReduceReSync instantiates a new HyperflexReduceReSync object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexReduceReSync(classId string, objectType string) *HyperflexReduceReSync {
	this := HyperflexReduceReSync{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexReduceReSyncWithDefaults instantiates a new HyperflexReduceReSync object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexReduceReSyncWithDefaults() *HyperflexReduceReSync {
	this := HyperflexReduceReSync{}
	var classId string = "hyperflex.ReduceReSync"
	this.ClassId = classId
	var objectType string = "hyperflex.ReduceReSync"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexReduceReSync) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexReduceReSync) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexReduceReSync) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ReduceReSync" of the ClassId field.
func (o *HyperflexReduceReSync) GetDefaultClassId() interface{} {
	return "hyperflex.ReduceReSync"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexReduceReSync) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexReduceReSync) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexReduceReSync) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ReduceReSync" of the ObjectType field.
func (o *HyperflexReduceReSync) GetDefaultObjectType() interface{} {
	return "hyperflex.ReduceReSync"
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *HyperflexReduceReSync) GetCompletionStatus() bool {
	if o == nil || IsNil(o.CompletionStatus) {
		var ret bool
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReduceReSync) GetCompletionStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.CompletionStatus) {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *HyperflexReduceReSync) HasCompletionStatus() bool {
	if o != nil && !IsNil(o.CompletionStatus) {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given bool and assigns it to the CompletionStatus field.
func (o *HyperflexReduceReSync) SetCompletionStatus(v bool) {
	o.CompletionStatus = &v
}

// GetExecutionOutput returns the ExecutionOutput field value if set, zero value otherwise.
func (o *HyperflexReduceReSync) GetExecutionOutput() string {
	if o == nil || IsNil(o.ExecutionOutput) {
		var ret string
		return ret
	}
	return *o.ExecutionOutput
}

// GetExecutionOutputOk returns a tuple with the ExecutionOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReduceReSync) GetExecutionOutputOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionOutput) {
		return nil, false
	}
	return o.ExecutionOutput, true
}

// HasExecutionOutput returns a boolean if a field has been set.
func (o *HyperflexReduceReSync) HasExecutionOutput() bool {
	if o != nil && !IsNil(o.ExecutionOutput) {
		return true
	}

	return false
}

// SetExecutionOutput gets a reference to the given string and assigns it to the ExecutionOutput field.
func (o *HyperflexReduceReSync) SetExecutionOutput(v string) {
	o.ExecutionOutput = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReduceReSync) GetCluster() HyperflexClusterRelationship {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReduceReSync) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexReduceReSync) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableHyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexReduceReSync) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *HyperflexReduceReSync) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *HyperflexReduceReSync) UnsetCluster() {
	o.Cluster.Unset()
}

func (o HyperflexReduceReSync) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexReduceReSync) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CompletionStatus) {
		toSerialize["CompletionStatus"] = o.CompletionStatus
	}
	if !IsNil(o.ExecutionOutput) {
		toSerialize["ExecutionOutput"] = o.ExecutionOutput
	}
	if o.Cluster.IsSet() {
		toSerialize["Cluster"] = o.Cluster.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexReduceReSync) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexReduceReSyncWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The HyperFlex reduce re-sync script execution completion status.
		CompletionStatus *bool `json:"CompletionStatus,omitempty"`
		// The execution output of the script.
		ExecutionOutput *string                              `json:"ExecutionOutput,omitempty"`
		Cluster         NullableHyperflexClusterRelationship `json:"Cluster,omitempty"`
	}

	varHyperflexReduceReSyncWithoutEmbeddedStruct := HyperflexReduceReSyncWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexReduceReSyncWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexReduceReSync := _HyperflexReduceReSync{}
		varHyperflexReduceReSync.ClassId = varHyperflexReduceReSyncWithoutEmbeddedStruct.ClassId
		varHyperflexReduceReSync.ObjectType = varHyperflexReduceReSyncWithoutEmbeddedStruct.ObjectType
		varHyperflexReduceReSync.CompletionStatus = varHyperflexReduceReSyncWithoutEmbeddedStruct.CompletionStatus
		varHyperflexReduceReSync.ExecutionOutput = varHyperflexReduceReSyncWithoutEmbeddedStruct.ExecutionOutput
		varHyperflexReduceReSync.Cluster = varHyperflexReduceReSyncWithoutEmbeddedStruct.Cluster
		*o = HyperflexReduceReSync(varHyperflexReduceReSync)
	} else {
		return err
	}

	varHyperflexReduceReSync := _HyperflexReduceReSync{}

	err = json.Unmarshal(data, &varHyperflexReduceReSync)
	if err == nil {
		o.MoBaseMo = varHyperflexReduceReSync.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CompletionStatus")
		delete(additionalProperties, "ExecutionOutput")
		delete(additionalProperties, "Cluster")

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

type NullableHyperflexReduceReSync struct {
	value *HyperflexReduceReSync
	isSet bool
}

func (v NullableHyperflexReduceReSync) Get() *HyperflexReduceReSync {
	return v.value
}

func (v *NullableHyperflexReduceReSync) Set(val *HyperflexReduceReSync) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexReduceReSync) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexReduceReSync) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexReduceReSync(val *HyperflexReduceReSync) *NullableHyperflexReduceReSync {
	return &NullableHyperflexReduceReSync{value: val, isSet: true}
}

func (v NullableHyperflexReduceReSync) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexReduceReSync) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
