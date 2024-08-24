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

// checks if the HyperflexHealth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexHealth{}

// HyperflexHealth The data health status and ability of the HyperFlex storage cluster to tolerate failures.
type HyperflexHealth struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The status of the HyperFlex cluster's connection to the Intersight arbitration service. The arbitration service state is only applicable to 2-node edge clusters. * `NOT_AVAILABLE` - The cluster does not require a connection to the arbitration service. * `UNKNOWN` - The cluster's connection state to the arbitration service cannot be determined. * `ONLINE` - The cluster is connected to the arbitration service. * `OFFLINE` - The cluster is disconnected from the arbitration service.
	ArbitrationServiceState *string `json:"ArbitrationServiceState,omitempty"`
	// The HyperFlex cluster's compliance to the configured replication factor. It indicates that the compliance has degraded if the number of copies of data is reduced. * `UNKNOWN` - The replication compliance of the HyperFlex cluster is not known. * `COMPLIANT` - The HyperFlex cluster is compliant with the replication policy. All data on the cluster is replicated according to the configured replication factor. * `NON_COMPLIANT` - The HyperFlex cluster is not compliant with the replication policy. Some data on the cluster is not replicated in accordance with the configured replication factor.
	DataReplicationCompliance *string                             `json:"DataReplicationCompliance,omitempty"`
	ResiliencyDetails         NullableHyperflexHxResiliencyInfoDt `json:"ResiliencyDetails,omitempty"`
	// The operational status of the HyperFlex cluster. * `UNKNOWN` - The operational status of the cluster cannot be determined. * `ONLINE` - The HyperFlex cluster is online and is performing IO operations. * `OFFLINE` - The HyperFlex cluster is offline and is not ready to perform IO operations. * `ENOSPACE` - The HyperFlex cluster is out of available storage capacity and cannot perform write transactions. * `READONLY` - The HyperFlex cluster is not accepting write transactions, but can still display static cluster information.
	State *string `json:"State,omitempty"`
	// The unique identifier for the cluster.
	Uuid *string `json:"Uuid,omitempty"`
	// The health status of the HyperFlex cluster's zookeeper ensemble. * `NOT_AVAILABLE` - The operational status of the ZK ensemble is not provided by the HyperFlex cluster. * `UNKNOWN` - The operational status of the ZK ensemble cannot be determined. * `ONLINE` - The ZK ensemble is online and operational. * `OFFLINE` - The ZK ensemble is offline and not operational.
	ZkHealth             *string                              `json:"ZkHealth,omitempty"`
	ZoneResiliencyList   []HyperflexHxZoneResiliencyInfoDt    `json:"ZoneResiliencyList,omitempty"`
	Cluster              NullableHyperflexClusterRelationship `json:"Cluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHealth HyperflexHealth

// NewHyperflexHealth instantiates a new HyperflexHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHealth(classId string, objectType string) *HyperflexHealth {
	this := HyperflexHealth{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHealthWithDefaults instantiates a new HyperflexHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHealthWithDefaults() *HyperflexHealth {
	this := HyperflexHealth{}
	var classId string = "hyperflex.Health"
	this.ClassId = classId
	var objectType string = "hyperflex.Health"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHealth) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHealth) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.Health" of the ClassId field.
func (o *HyperflexHealth) GetDefaultClassId() interface{} {
	return "hyperflex.Health"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHealth) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHealth) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.Health" of the ObjectType field.
func (o *HyperflexHealth) GetDefaultObjectType() interface{} {
	return "hyperflex.Health"
}

// GetArbitrationServiceState returns the ArbitrationServiceState field value if set, zero value otherwise.
func (o *HyperflexHealth) GetArbitrationServiceState() string {
	if o == nil || IsNil(o.ArbitrationServiceState) {
		var ret string
		return ret
	}
	return *o.ArbitrationServiceState
}

// GetArbitrationServiceStateOk returns a tuple with the ArbitrationServiceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetArbitrationServiceStateOk() (*string, bool) {
	if o == nil || IsNil(o.ArbitrationServiceState) {
		return nil, false
	}
	return o.ArbitrationServiceState, true
}

// HasArbitrationServiceState returns a boolean if a field has been set.
func (o *HyperflexHealth) HasArbitrationServiceState() bool {
	if o != nil && !IsNil(o.ArbitrationServiceState) {
		return true
	}

	return false
}

// SetArbitrationServiceState gets a reference to the given string and assigns it to the ArbitrationServiceState field.
func (o *HyperflexHealth) SetArbitrationServiceState(v string) {
	o.ArbitrationServiceState = &v
}

// GetDataReplicationCompliance returns the DataReplicationCompliance field value if set, zero value otherwise.
func (o *HyperflexHealth) GetDataReplicationCompliance() string {
	if o == nil || IsNil(o.DataReplicationCompliance) {
		var ret string
		return ret
	}
	return *o.DataReplicationCompliance
}

// GetDataReplicationComplianceOk returns a tuple with the DataReplicationCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetDataReplicationComplianceOk() (*string, bool) {
	if o == nil || IsNil(o.DataReplicationCompliance) {
		return nil, false
	}
	return o.DataReplicationCompliance, true
}

// HasDataReplicationCompliance returns a boolean if a field has been set.
func (o *HyperflexHealth) HasDataReplicationCompliance() bool {
	if o != nil && !IsNil(o.DataReplicationCompliance) {
		return true
	}

	return false
}

// SetDataReplicationCompliance gets a reference to the given string and assigns it to the DataReplicationCompliance field.
func (o *HyperflexHealth) SetDataReplicationCompliance(v string) {
	o.DataReplicationCompliance = &v
}

// GetResiliencyDetails returns the ResiliencyDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHealth) GetResiliencyDetails() HyperflexHxResiliencyInfoDt {
	if o == nil || IsNil(o.ResiliencyDetails.Get()) {
		var ret HyperflexHxResiliencyInfoDt
		return ret
	}
	return *o.ResiliencyDetails.Get()
}

// GetResiliencyDetailsOk returns a tuple with the ResiliencyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHealth) GetResiliencyDetailsOk() (*HyperflexHxResiliencyInfoDt, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResiliencyDetails.Get(), o.ResiliencyDetails.IsSet()
}

// HasResiliencyDetails returns a boolean if a field has been set.
func (o *HyperflexHealth) HasResiliencyDetails() bool {
	if o != nil && o.ResiliencyDetails.IsSet() {
		return true
	}

	return false
}

// SetResiliencyDetails gets a reference to the given NullableHyperflexHxResiliencyInfoDt and assigns it to the ResiliencyDetails field.
func (o *HyperflexHealth) SetResiliencyDetails(v HyperflexHxResiliencyInfoDt) {
	o.ResiliencyDetails.Set(&v)
}

// SetResiliencyDetailsNil sets the value for ResiliencyDetails to be an explicit nil
func (o *HyperflexHealth) SetResiliencyDetailsNil() {
	o.ResiliencyDetails.Set(nil)
}

// UnsetResiliencyDetails ensures that no value is present for ResiliencyDetails, not even an explicit nil
func (o *HyperflexHealth) UnsetResiliencyDetails() {
	o.ResiliencyDetails.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *HyperflexHealth) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *HyperflexHealth) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *HyperflexHealth) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexHealth) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexHealth) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexHealth) SetUuid(v string) {
	o.Uuid = &v
}

// GetZkHealth returns the ZkHealth field value if set, zero value otherwise.
func (o *HyperflexHealth) GetZkHealth() string {
	if o == nil || IsNil(o.ZkHealth) {
		var ret string
		return ret
	}
	return *o.ZkHealth
}

// GetZkHealthOk returns a tuple with the ZkHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealth) GetZkHealthOk() (*string, bool) {
	if o == nil || IsNil(o.ZkHealth) {
		return nil, false
	}
	return o.ZkHealth, true
}

// HasZkHealth returns a boolean if a field has been set.
func (o *HyperflexHealth) HasZkHealth() bool {
	if o != nil && !IsNil(o.ZkHealth) {
		return true
	}

	return false
}

// SetZkHealth gets a reference to the given string and assigns it to the ZkHealth field.
func (o *HyperflexHealth) SetZkHealth(v string) {
	o.ZkHealth = &v
}

// GetZoneResiliencyList returns the ZoneResiliencyList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHealth) GetZoneResiliencyList() []HyperflexHxZoneResiliencyInfoDt {
	if o == nil {
		var ret []HyperflexHxZoneResiliencyInfoDt
		return ret
	}
	return o.ZoneResiliencyList
}

// GetZoneResiliencyListOk returns a tuple with the ZoneResiliencyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHealth) GetZoneResiliencyListOk() ([]HyperflexHxZoneResiliencyInfoDt, bool) {
	if o == nil || IsNil(o.ZoneResiliencyList) {
		return nil, false
	}
	return o.ZoneResiliencyList, true
}

// HasZoneResiliencyList returns a boolean if a field has been set.
func (o *HyperflexHealth) HasZoneResiliencyList() bool {
	if o != nil && !IsNil(o.ZoneResiliencyList) {
		return true
	}

	return false
}

// SetZoneResiliencyList gets a reference to the given []HyperflexHxZoneResiliencyInfoDt and assigns it to the ZoneResiliencyList field.
func (o *HyperflexHealth) SetZoneResiliencyList(v []HyperflexHxZoneResiliencyInfoDt) {
	o.ZoneResiliencyList = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHealth) GetCluster() HyperflexClusterRelationship {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHealth) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHealth) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableHyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHealth) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *HyperflexHealth) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *HyperflexHealth) UnsetCluster() {
	o.Cluster.Unset()
}

func (o HyperflexHealth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexHealth) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ArbitrationServiceState) {
		toSerialize["ArbitrationServiceState"] = o.ArbitrationServiceState
	}
	if !IsNil(o.DataReplicationCompliance) {
		toSerialize["DataReplicationCompliance"] = o.DataReplicationCompliance
	}
	if o.ResiliencyDetails.IsSet() {
		toSerialize["ResiliencyDetails"] = o.ResiliencyDetails.Get()
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if !IsNil(o.ZkHealth) {
		toSerialize["ZkHealth"] = o.ZkHealth
	}
	if o.ZoneResiliencyList != nil {
		toSerialize["ZoneResiliencyList"] = o.ZoneResiliencyList
	}
	if o.Cluster.IsSet() {
		toSerialize["Cluster"] = o.Cluster.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexHealth) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexHealthWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The status of the HyperFlex cluster's connection to the Intersight arbitration service. The arbitration service state is only applicable to 2-node edge clusters. * `NOT_AVAILABLE` - The cluster does not require a connection to the arbitration service. * `UNKNOWN` - The cluster's connection state to the arbitration service cannot be determined. * `ONLINE` - The cluster is connected to the arbitration service. * `OFFLINE` - The cluster is disconnected from the arbitration service.
		ArbitrationServiceState *string `json:"ArbitrationServiceState,omitempty"`
		// The HyperFlex cluster's compliance to the configured replication factor. It indicates that the compliance has degraded if the number of copies of data is reduced. * `UNKNOWN` - The replication compliance of the HyperFlex cluster is not known. * `COMPLIANT` - The HyperFlex cluster is compliant with the replication policy. All data on the cluster is replicated according to the configured replication factor. * `NON_COMPLIANT` - The HyperFlex cluster is not compliant with the replication policy. Some data on the cluster is not replicated in accordance with the configured replication factor.
		DataReplicationCompliance *string                             `json:"DataReplicationCompliance,omitempty"`
		ResiliencyDetails         NullableHyperflexHxResiliencyInfoDt `json:"ResiliencyDetails,omitempty"`
		// The operational status of the HyperFlex cluster. * `UNKNOWN` - The operational status of the cluster cannot be determined. * `ONLINE` - The HyperFlex cluster is online and is performing IO operations. * `OFFLINE` - The HyperFlex cluster is offline and is not ready to perform IO operations. * `ENOSPACE` - The HyperFlex cluster is out of available storage capacity and cannot perform write transactions. * `READONLY` - The HyperFlex cluster is not accepting write transactions, but can still display static cluster information.
		State *string `json:"State,omitempty"`
		// The unique identifier for the cluster.
		Uuid *string `json:"Uuid,omitempty"`
		// The health status of the HyperFlex cluster's zookeeper ensemble. * `NOT_AVAILABLE` - The operational status of the ZK ensemble is not provided by the HyperFlex cluster. * `UNKNOWN` - The operational status of the ZK ensemble cannot be determined. * `ONLINE` - The ZK ensemble is online and operational. * `OFFLINE` - The ZK ensemble is offline and not operational.
		ZkHealth           *string                              `json:"ZkHealth,omitempty"`
		ZoneResiliencyList []HyperflexHxZoneResiliencyInfoDt    `json:"ZoneResiliencyList,omitempty"`
		Cluster            NullableHyperflexClusterRelationship `json:"Cluster,omitempty"`
	}

	varHyperflexHealthWithoutEmbeddedStruct := HyperflexHealthWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexHealthWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHealth := _HyperflexHealth{}
		varHyperflexHealth.ClassId = varHyperflexHealthWithoutEmbeddedStruct.ClassId
		varHyperflexHealth.ObjectType = varHyperflexHealthWithoutEmbeddedStruct.ObjectType
		varHyperflexHealth.ArbitrationServiceState = varHyperflexHealthWithoutEmbeddedStruct.ArbitrationServiceState
		varHyperflexHealth.DataReplicationCompliance = varHyperflexHealthWithoutEmbeddedStruct.DataReplicationCompliance
		varHyperflexHealth.ResiliencyDetails = varHyperflexHealthWithoutEmbeddedStruct.ResiliencyDetails
		varHyperflexHealth.State = varHyperflexHealthWithoutEmbeddedStruct.State
		varHyperflexHealth.Uuid = varHyperflexHealthWithoutEmbeddedStruct.Uuid
		varHyperflexHealth.ZkHealth = varHyperflexHealthWithoutEmbeddedStruct.ZkHealth
		varHyperflexHealth.ZoneResiliencyList = varHyperflexHealthWithoutEmbeddedStruct.ZoneResiliencyList
		varHyperflexHealth.Cluster = varHyperflexHealthWithoutEmbeddedStruct.Cluster
		*o = HyperflexHealth(varHyperflexHealth)
	} else {
		return err
	}

	varHyperflexHealth := _HyperflexHealth{}

	err = json.Unmarshal(data, &varHyperflexHealth)
	if err == nil {
		o.MoBaseMo = varHyperflexHealth.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArbitrationServiceState")
		delete(additionalProperties, "DataReplicationCompliance")
		delete(additionalProperties, "ResiliencyDetails")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ZkHealth")
		delete(additionalProperties, "ZoneResiliencyList")
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

type NullableHyperflexHealth struct {
	value *HyperflexHealth
	isSet bool
}

func (v NullableHyperflexHealth) Get() *HyperflexHealth {
	return v.value
}

func (v *NullableHyperflexHealth) Set(val *HyperflexHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHealth(val *HyperflexHealth) *NullableHyperflexHealth {
	return &NullableHyperflexHealth{value: val, isSet: true}
}

func (v NullableHyperflexHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
