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

// checks if the HyperflexClusterHealthCheckExecutionSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexClusterHealthCheckExecutionSnapshot{}

// HyperflexClusterHealthCheckExecutionSnapshot Health check execution snapshot of the HyperFlex cluster.
type HyperflexClusterHealthCheckExecutionSnapshot struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The execution context of the HyperFlex health checks. * `UNKNOWN` - The current context of HyperFlex health check execution is unknown. * `WORKFLOW` - The HyperFlex health check execution is initiated through an orchestration workflow. * `SCHEDULED` - The HyperFlex health check execution is through a scheduled run.
	ExecutionContext *string `json:"ExecutionContext,omitempty"`
	// Timestamp of the last health check execution on the HyperFlex cluster.
	Timestamp            *time.Time                                  `json:"Timestamp,omitempty"`
	HxCluster            NullableHyperflexClusterRelationship        `json:"HxCluster,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	Workflow             NullableWorkflowWorkflowInfoRelationship    `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterHealthCheckExecutionSnapshot HyperflexClusterHealthCheckExecutionSnapshot

// NewHyperflexClusterHealthCheckExecutionSnapshot instantiates a new HyperflexClusterHealthCheckExecutionSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterHealthCheckExecutionSnapshot(classId string, objectType string) *HyperflexClusterHealthCheckExecutionSnapshot {
	this := HyperflexClusterHealthCheckExecutionSnapshot{}
	this.ClassId = classId
	this.ObjectType = objectType
	var executionContext string = "UNKNOWN"
	this.ExecutionContext = &executionContext
	return &this
}

// NewHyperflexClusterHealthCheckExecutionSnapshotWithDefaults instantiates a new HyperflexClusterHealthCheckExecutionSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterHealthCheckExecutionSnapshotWithDefaults() *HyperflexClusterHealthCheckExecutionSnapshot {
	this := HyperflexClusterHealthCheckExecutionSnapshot{}
	var classId string = "hyperflex.ClusterHealthCheckExecutionSnapshot"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterHealthCheckExecutionSnapshot"
	this.ObjectType = objectType
	var executionContext string = "UNKNOWN"
	this.ExecutionContext = &executionContext
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ClusterHealthCheckExecutionSnapshot" of the ClassId field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetDefaultClassId() interface{} {
	return "hyperflex.ClusterHealthCheckExecutionSnapshot"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ClusterHealthCheckExecutionSnapshot" of the ObjectType field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetDefaultObjectType() interface{} {
	return "hyperflex.ClusterHealthCheckExecutionSnapshot"
}

// GetExecutionContext returns the ExecutionContext field value if set, zero value otherwise.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetExecutionContext() string {
	if o == nil || IsNil(o.ExecutionContext) {
		var ret string
		return ret
	}
	return *o.ExecutionContext
}

// GetExecutionContextOk returns a tuple with the ExecutionContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetExecutionContextOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionContext) {
		return nil, false
	}
	return o.ExecutionContext, true
}

// HasExecutionContext returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasExecutionContext() bool {
	if o != nil && !IsNil(o.ExecutionContext) {
		return true
	}

	return false
}

// SetExecutionContext gets a reference to the given string and assigns it to the ExecutionContext field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetExecutionContext(v string) {
	o.ExecutionContext = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetHxCluster returns the HxCluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetHxCluster() HyperflexClusterRelationship {
	if o == nil || IsNil(o.HxCluster.Get()) {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.HxCluster.Get()
}

// GetHxClusterOk returns a tuple with the HxCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetHxClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.HxCluster.Get(), o.HxCluster.IsSet()
}

// HasHxCluster returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasHxCluster() bool {
	if o != nil && o.HxCluster.IsSet() {
		return true
	}

	return false
}

// SetHxCluster gets a reference to the given NullableHyperflexClusterRelationship and assigns it to the HxCluster field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetHxCluster(v HyperflexClusterRelationship) {
	o.HxCluster.Set(&v)
}

// SetHxClusterNil sets the value for HxCluster to be an explicit nil
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetHxClusterNil() {
	o.HxCluster.Set(nil)
}

// UnsetHxCluster ensures that no value is present for HxCluster, not even an explicit nil
func (o *HyperflexClusterHealthCheckExecutionSnapshot) UnsetHxCluster() {
	o.HxCluster.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *HyperflexClusterHealthCheckExecutionSnapshot) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || IsNil(o.Workflow.Get()) {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow.Get()
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workflow.Get(), o.Workflow.IsSet()
}

// HasWorkflow returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasWorkflow() bool {
	if o != nil && o.Workflow.IsSet() {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given NullableWorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow.Set(&v)
}

// SetWorkflowNil sets the value for Workflow to be an explicit nil
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetWorkflowNil() {
	o.Workflow.Set(nil)
}

// UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
func (o *HyperflexClusterHealthCheckExecutionSnapshot) UnsetWorkflow() {
	o.Workflow.Unset()
}

func (o HyperflexClusterHealthCheckExecutionSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexClusterHealthCheckExecutionSnapshot) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ExecutionContext) {
		toSerialize["ExecutionContext"] = o.ExecutionContext
	}
	if !IsNil(o.Timestamp) {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.HxCluster.IsSet() {
		toSerialize["HxCluster"] = o.HxCluster.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.Workflow.IsSet() {
		toSerialize["Workflow"] = o.Workflow.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexClusterHealthCheckExecutionSnapshot) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The execution context of the HyperFlex health checks. * `UNKNOWN` - The current context of HyperFlex health check execution is unknown. * `WORKFLOW` - The HyperFlex health check execution is initiated through an orchestration workflow. * `SCHEDULED` - The HyperFlex health check execution is through a scheduled run.
		ExecutionContext *string `json:"ExecutionContext,omitempty"`
		// Timestamp of the last health check execution on the HyperFlex cluster.
		Timestamp        *time.Time                                  `json:"Timestamp,omitempty"`
		HxCluster        NullableHyperflexClusterRelationship        `json:"HxCluster,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		Workflow         NullableWorkflowWorkflowInfoRelationship    `json:"Workflow,omitempty"`
	}

	varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct := HyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexClusterHealthCheckExecutionSnapshot := _HyperflexClusterHealthCheckExecutionSnapshot{}
		varHyperflexClusterHealthCheckExecutionSnapshot.ClassId = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.ClassId
		varHyperflexClusterHealthCheckExecutionSnapshot.ObjectType = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.ObjectType
		varHyperflexClusterHealthCheckExecutionSnapshot.ExecutionContext = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.ExecutionContext
		varHyperflexClusterHealthCheckExecutionSnapshot.Timestamp = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.Timestamp
		varHyperflexClusterHealthCheckExecutionSnapshot.HxCluster = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.HxCluster
		varHyperflexClusterHealthCheckExecutionSnapshot.RegisteredDevice = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.RegisteredDevice
		varHyperflexClusterHealthCheckExecutionSnapshot.Workflow = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.Workflow
		*o = HyperflexClusterHealthCheckExecutionSnapshot(varHyperflexClusterHealthCheckExecutionSnapshot)
	} else {
		return err
	}

	varHyperflexClusterHealthCheckExecutionSnapshot := _HyperflexClusterHealthCheckExecutionSnapshot{}

	err = json.Unmarshal(data, &varHyperflexClusterHealthCheckExecutionSnapshot)
	if err == nil {
		o.MoBaseMo = varHyperflexClusterHealthCheckExecutionSnapshot.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExecutionContext")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "HxCluster")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Workflow")

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

type NullableHyperflexClusterHealthCheckExecutionSnapshot struct {
	value *HyperflexClusterHealthCheckExecutionSnapshot
	isSet bool
}

func (v NullableHyperflexClusterHealthCheckExecutionSnapshot) Get() *HyperflexClusterHealthCheckExecutionSnapshot {
	return v.value
}

func (v *NullableHyperflexClusterHealthCheckExecutionSnapshot) Set(val *HyperflexClusterHealthCheckExecutionSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterHealthCheckExecutionSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterHealthCheckExecutionSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterHealthCheckExecutionSnapshot(val *HyperflexClusterHealthCheckExecutionSnapshot) *NullableHyperflexClusterHealthCheckExecutionSnapshot {
	return &NullableHyperflexClusterHealthCheckExecutionSnapshot{value: val, isSet: true}
}

func (v NullableHyperflexClusterHealthCheckExecutionSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterHealthCheckExecutionSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
