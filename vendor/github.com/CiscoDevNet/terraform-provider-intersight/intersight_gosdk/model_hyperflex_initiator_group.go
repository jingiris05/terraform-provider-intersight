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

// checks if the HyperflexInitiatorGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexInitiatorGroup{}

// HyperflexInitiatorGroup A HyperFlex iSCSI initiator group entity. Contains detailed information about the initaitor group which includes a list of iSCSI initiators and iSCSI target objects.
type HyperflexInitiatorGroup struct {
	StorageBaseHostGroup
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Count of initiators in the iSCSI initiator group.
	InitiatorCount *int64                           `json:"InitiatorCount,omitempty"`
	Initiators     []StorageHyperFlexIscsiInitiator `json:"Initiators,omitempty"`
	// The source of the iSCSI initiator group inventory. * `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable. * `ONLINE` - The source of the HyperFlex inventory is online. * `OFFLINE` - The source of the HyperFlex inventory is offline.
	InventorySource *string  `json:"InventorySource,omitempty"`
	TargetUuids     []string `json:"TargetUuids,omitempty"`
	// UUID of the HyperFlex iSCSI initiator group.
	Uuid *string `json:"Uuid,omitempty"`
	// Version of the iSCSI initiator group.
	Version *int64                               `json:"Version,omitempty"`
	Cluster NullableHyperflexClusterRelationship `json:"Cluster,omitempty"`
	// An array of relationships to hyperflexTarget resources.
	Targets              []HyperflexTargetRelationship `json:"Targets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexInitiatorGroup HyperflexInitiatorGroup

// NewHyperflexInitiatorGroup instantiates a new HyperflexInitiatorGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexInitiatorGroup(classId string, objectType string) *HyperflexInitiatorGroup {
	this := HyperflexInitiatorGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexInitiatorGroupWithDefaults instantiates a new HyperflexInitiatorGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexInitiatorGroupWithDefaults() *HyperflexInitiatorGroup {
	this := HyperflexInitiatorGroup{}
	var classId string = "hyperflex.InitiatorGroup"
	this.ClassId = classId
	var objectType string = "hyperflex.InitiatorGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexInitiatorGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexInitiatorGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexInitiatorGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.InitiatorGroup" of the ClassId field.
func (o *HyperflexInitiatorGroup) GetDefaultClassId() interface{} {
	return "hyperflex.InitiatorGroup"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexInitiatorGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexInitiatorGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexInitiatorGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.InitiatorGroup" of the ObjectType field.
func (o *HyperflexInitiatorGroup) GetDefaultObjectType() interface{} {
	return "hyperflex.InitiatorGroup"
}

// GetInitiatorCount returns the InitiatorCount field value if set, zero value otherwise.
func (o *HyperflexInitiatorGroup) GetInitiatorCount() int64 {
	if o == nil || IsNil(o.InitiatorCount) {
		var ret int64
		return ret
	}
	return *o.InitiatorCount
}

// GetInitiatorCountOk returns a tuple with the InitiatorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexInitiatorGroup) GetInitiatorCountOk() (*int64, bool) {
	if o == nil || IsNil(o.InitiatorCount) {
		return nil, false
	}
	return o.InitiatorCount, true
}

// HasInitiatorCount returns a boolean if a field has been set.
func (o *HyperflexInitiatorGroup) HasInitiatorCount() bool {
	if o != nil && !IsNil(o.InitiatorCount) {
		return true
	}

	return false
}

// SetInitiatorCount gets a reference to the given int64 and assigns it to the InitiatorCount field.
func (o *HyperflexInitiatorGroup) SetInitiatorCount(v int64) {
	o.InitiatorCount = &v
}

// GetInitiators returns the Initiators field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexInitiatorGroup) GetInitiators() []StorageHyperFlexIscsiInitiator {
	if o == nil {
		var ret []StorageHyperFlexIscsiInitiator
		return ret
	}
	return o.Initiators
}

// GetInitiatorsOk returns a tuple with the Initiators field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexInitiatorGroup) GetInitiatorsOk() ([]StorageHyperFlexIscsiInitiator, bool) {
	if o == nil || IsNil(o.Initiators) {
		return nil, false
	}
	return o.Initiators, true
}

// HasInitiators returns a boolean if a field has been set.
func (o *HyperflexInitiatorGroup) HasInitiators() bool {
	if o != nil && !IsNil(o.Initiators) {
		return true
	}

	return false
}

// SetInitiators gets a reference to the given []StorageHyperFlexIscsiInitiator and assigns it to the Initiators field.
func (o *HyperflexInitiatorGroup) SetInitiators(v []StorageHyperFlexIscsiInitiator) {
	o.Initiators = v
}

// GetInventorySource returns the InventorySource field value if set, zero value otherwise.
func (o *HyperflexInitiatorGroup) GetInventorySource() string {
	if o == nil || IsNil(o.InventorySource) {
		var ret string
		return ret
	}
	return *o.InventorySource
}

// GetInventorySourceOk returns a tuple with the InventorySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexInitiatorGroup) GetInventorySourceOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySource) {
		return nil, false
	}
	return o.InventorySource, true
}

// HasInventorySource returns a boolean if a field has been set.
func (o *HyperflexInitiatorGroup) HasInventorySource() bool {
	if o != nil && !IsNil(o.InventorySource) {
		return true
	}

	return false
}

// SetInventorySource gets a reference to the given string and assigns it to the InventorySource field.
func (o *HyperflexInitiatorGroup) SetInventorySource(v string) {
	o.InventorySource = &v
}

// GetTargetUuids returns the TargetUuids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexInitiatorGroup) GetTargetUuids() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TargetUuids
}

// GetTargetUuidsOk returns a tuple with the TargetUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexInitiatorGroup) GetTargetUuidsOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetUuids) {
		return nil, false
	}
	return o.TargetUuids, true
}

// HasTargetUuids returns a boolean if a field has been set.
func (o *HyperflexInitiatorGroup) HasTargetUuids() bool {
	if o != nil && !IsNil(o.TargetUuids) {
		return true
	}

	return false
}

// SetTargetUuids gets a reference to the given []string and assigns it to the TargetUuids field.
func (o *HyperflexInitiatorGroup) SetTargetUuids(v []string) {
	o.TargetUuids = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexInitiatorGroup) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexInitiatorGroup) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexInitiatorGroup) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexInitiatorGroup) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexInitiatorGroup) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexInitiatorGroup) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexInitiatorGroup) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *HyperflexInitiatorGroup) SetVersion(v int64) {
	o.Version = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexInitiatorGroup) GetCluster() HyperflexClusterRelationship {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexInitiatorGroup) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexInitiatorGroup) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableHyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexInitiatorGroup) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *HyperflexInitiatorGroup) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *HyperflexInitiatorGroup) UnsetCluster() {
	o.Cluster.Unset()
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexInitiatorGroup) GetTargets() []HyperflexTargetRelationship {
	if o == nil {
		var ret []HyperflexTargetRelationship
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexInitiatorGroup) GetTargetsOk() ([]HyperflexTargetRelationship, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *HyperflexInitiatorGroup) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []HyperflexTargetRelationship and assigns it to the Targets field.
func (o *HyperflexInitiatorGroup) SetTargets(v []HyperflexTargetRelationship) {
	o.Targets = v
}

func (o HyperflexInitiatorGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexInitiatorGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseHostGroup, errStorageBaseHostGroup := json.Marshal(o.StorageBaseHostGroup)
	if errStorageBaseHostGroup != nil {
		return map[string]interface{}{}, errStorageBaseHostGroup
	}
	errStorageBaseHostGroup = json.Unmarshal([]byte(serializedStorageBaseHostGroup), &toSerialize)
	if errStorageBaseHostGroup != nil {
		return map[string]interface{}{}, errStorageBaseHostGroup
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.InitiatorCount) {
		toSerialize["InitiatorCount"] = o.InitiatorCount
	}
	if o.Initiators != nil {
		toSerialize["Initiators"] = o.Initiators
	}
	if !IsNil(o.InventorySource) {
		toSerialize["InventorySource"] = o.InventorySource
	}
	if o.TargetUuids != nil {
		toSerialize["TargetUuids"] = o.TargetUuids
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if o.Cluster.IsSet() {
		toSerialize["Cluster"] = o.Cluster.Get()
	}
	if o.Targets != nil {
		toSerialize["Targets"] = o.Targets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexInitiatorGroup) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexInitiatorGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Count of initiators in the iSCSI initiator group.
		InitiatorCount *int64                           `json:"InitiatorCount,omitempty"`
		Initiators     []StorageHyperFlexIscsiInitiator `json:"Initiators,omitempty"`
		// The source of the iSCSI initiator group inventory. * `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable. * `ONLINE` - The source of the HyperFlex inventory is online. * `OFFLINE` - The source of the HyperFlex inventory is offline.
		InventorySource *string  `json:"InventorySource,omitempty"`
		TargetUuids     []string `json:"TargetUuids,omitempty"`
		// UUID of the HyperFlex iSCSI initiator group.
		Uuid *string `json:"Uuid,omitempty"`
		// Version of the iSCSI initiator group.
		Version *int64                               `json:"Version,omitempty"`
		Cluster NullableHyperflexClusterRelationship `json:"Cluster,omitempty"`
		// An array of relationships to hyperflexTarget resources.
		Targets []HyperflexTargetRelationship `json:"Targets,omitempty"`
	}

	varHyperflexInitiatorGroupWithoutEmbeddedStruct := HyperflexInitiatorGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexInitiatorGroupWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexInitiatorGroup := _HyperflexInitiatorGroup{}
		varHyperflexInitiatorGroup.ClassId = varHyperflexInitiatorGroupWithoutEmbeddedStruct.ClassId
		varHyperflexInitiatorGroup.ObjectType = varHyperflexInitiatorGroupWithoutEmbeddedStruct.ObjectType
		varHyperflexInitiatorGroup.InitiatorCount = varHyperflexInitiatorGroupWithoutEmbeddedStruct.InitiatorCount
		varHyperflexInitiatorGroup.Initiators = varHyperflexInitiatorGroupWithoutEmbeddedStruct.Initiators
		varHyperflexInitiatorGroup.InventorySource = varHyperflexInitiatorGroupWithoutEmbeddedStruct.InventorySource
		varHyperflexInitiatorGroup.TargetUuids = varHyperflexInitiatorGroupWithoutEmbeddedStruct.TargetUuids
		varHyperflexInitiatorGroup.Uuid = varHyperflexInitiatorGroupWithoutEmbeddedStruct.Uuid
		varHyperflexInitiatorGroup.Version = varHyperflexInitiatorGroupWithoutEmbeddedStruct.Version
		varHyperflexInitiatorGroup.Cluster = varHyperflexInitiatorGroupWithoutEmbeddedStruct.Cluster
		varHyperflexInitiatorGroup.Targets = varHyperflexInitiatorGroupWithoutEmbeddedStruct.Targets
		*o = HyperflexInitiatorGroup(varHyperflexInitiatorGroup)
	} else {
		return err
	}

	varHyperflexInitiatorGroup := _HyperflexInitiatorGroup{}

	err = json.Unmarshal(data, &varHyperflexInitiatorGroup)
	if err == nil {
		o.StorageBaseHostGroup = varHyperflexInitiatorGroup.StorageBaseHostGroup
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InitiatorCount")
		delete(additionalProperties, "Initiators")
		delete(additionalProperties, "InventorySource")
		delete(additionalProperties, "TargetUuids")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Targets")

		// remove fields from embedded structs
		reflectStorageBaseHostGroup := reflect.ValueOf(o.StorageBaseHostGroup)
		for i := 0; i < reflectStorageBaseHostGroup.Type().NumField(); i++ {
			t := reflectStorageBaseHostGroup.Type().Field(i)

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

type NullableHyperflexInitiatorGroup struct {
	value *HyperflexInitiatorGroup
	isSet bool
}

func (v NullableHyperflexInitiatorGroup) Get() *HyperflexInitiatorGroup {
	return v.value
}

func (v *NullableHyperflexInitiatorGroup) Set(val *HyperflexInitiatorGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexInitiatorGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexInitiatorGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexInitiatorGroup(val *HyperflexInitiatorGroup) *NullableHyperflexInitiatorGroup {
	return &NullableHyperflexInitiatorGroup{value: val, isSet: true}
}

func (v NullableHyperflexInitiatorGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexInitiatorGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
