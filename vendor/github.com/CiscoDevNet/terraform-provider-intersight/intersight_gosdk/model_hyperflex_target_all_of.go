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

// HyperflexTargetAllOf Definition of the list of properties defined in 'hyperflex.Target', excluding properties defined in parent classes.
type HyperflexTargetAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Auth method of the target inventory. * `NOT_APPLICABLE` - Authorization method of the HyperFlex iSCSI target is not applicable. * `CHAP` - Authorization method of the HyperFlex iSCSI target is CHAP. * `NONE` - Authorization method of the HyperFlex iSCSI target is none.
	AuthMethod          *string  `json:"AuthMethod,omitempty"`
	InitiatorGroupUuids []string `json:"InitiatorGroupUuids,omitempty"`
	// Source of the target inventory. * `NOT_APPLICABLE` - The source of the HyperFlex inventory is not applicable. * `ONLINE` - The source of the HyperFlex inventory is online. * `OFFLINE` - The source of the HyperFlex inventory is offline.
	InventorySource *string `json:"InventorySource,omitempty"`
	// The iSCSI qualified name (IQN) of target.
	Iqn      *string  `json:"Iqn,omitempty"`
	LunUuids []string `json:"LunUuids,omitempty"`
	// Number of active initiators in the initiator group.
	NumActiveInitiators *int64 `json:"NumActiveInitiators,omitempty"`
	// UUID of the HyperFlex iSCSI target.
	Uuid *string `json:"Uuid,omitempty"`
	// Version of the Initiator Group.
	Version *int64                        `json:"Version,omitempty"`
	Cluster *HyperflexClusterRelationship `json:"Cluster,omitempty"`
	// An array of relationships to hyperflexInitiatorGroup resources.
	InitiatorGroups []HyperflexInitiatorGroupRelationship `json:"InitiatorGroups,omitempty"`
	// An array of relationships to hyperflexLun resources.
	Luns                 []HyperflexLunRelationship `json:"Luns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexTargetAllOf HyperflexTargetAllOf

// NewHyperflexTargetAllOf instantiates a new HyperflexTargetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexTargetAllOf(classId string, objectType string) *HyperflexTargetAllOf {
	this := HyperflexTargetAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexTargetAllOfWithDefaults instantiates a new HyperflexTargetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexTargetAllOfWithDefaults() *HyperflexTargetAllOf {
	this := HyperflexTargetAllOf{}
	var classId string = "hyperflex.Target"
	this.ClassId = classId
	var objectType string = "hyperflex.Target"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexTargetAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexTargetAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexTargetAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexTargetAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexTargetAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexTargetAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthMethod returns the AuthMethod field value if set, zero value otherwise.
func (o *HyperflexTargetAllOf) GetAuthMethod() string {
	if o == nil || o.AuthMethod == nil {
		var ret string
		return ret
	}
	return *o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexTargetAllOf) GetAuthMethodOk() (*string, bool) {
	if o == nil || o.AuthMethod == nil {
		return nil, false
	}
	return o.AuthMethod, true
}

// HasAuthMethod returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasAuthMethod() bool {
	if o != nil && o.AuthMethod != nil {
		return true
	}

	return false
}

// SetAuthMethod gets a reference to the given string and assigns it to the AuthMethod field.
func (o *HyperflexTargetAllOf) SetAuthMethod(v string) {
	o.AuthMethod = &v
}

// GetInitiatorGroupUuids returns the InitiatorGroupUuids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexTargetAllOf) GetInitiatorGroupUuids() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.InitiatorGroupUuids
}

// GetInitiatorGroupUuidsOk returns a tuple with the InitiatorGroupUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexTargetAllOf) GetInitiatorGroupUuidsOk() ([]string, bool) {
	if o == nil || o.InitiatorGroupUuids == nil {
		return nil, false
	}
	return o.InitiatorGroupUuids, true
}

// HasInitiatorGroupUuids returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasInitiatorGroupUuids() bool {
	if o != nil && o.InitiatorGroupUuids != nil {
		return true
	}

	return false
}

// SetInitiatorGroupUuids gets a reference to the given []string and assigns it to the InitiatorGroupUuids field.
func (o *HyperflexTargetAllOf) SetInitiatorGroupUuids(v []string) {
	o.InitiatorGroupUuids = v
}

// GetInventorySource returns the InventorySource field value if set, zero value otherwise.
func (o *HyperflexTargetAllOf) GetInventorySource() string {
	if o == nil || o.InventorySource == nil {
		var ret string
		return ret
	}
	return *o.InventorySource
}

// GetInventorySourceOk returns a tuple with the InventorySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexTargetAllOf) GetInventorySourceOk() (*string, bool) {
	if o == nil || o.InventorySource == nil {
		return nil, false
	}
	return o.InventorySource, true
}

// HasInventorySource returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasInventorySource() bool {
	if o != nil && o.InventorySource != nil {
		return true
	}

	return false
}

// SetInventorySource gets a reference to the given string and assigns it to the InventorySource field.
func (o *HyperflexTargetAllOf) SetInventorySource(v string) {
	o.InventorySource = &v
}

// GetIqn returns the Iqn field value if set, zero value otherwise.
func (o *HyperflexTargetAllOf) GetIqn() string {
	if o == nil || o.Iqn == nil {
		var ret string
		return ret
	}
	return *o.Iqn
}

// GetIqnOk returns a tuple with the Iqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexTargetAllOf) GetIqnOk() (*string, bool) {
	if o == nil || o.Iqn == nil {
		return nil, false
	}
	return o.Iqn, true
}

// HasIqn returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasIqn() bool {
	if o != nil && o.Iqn != nil {
		return true
	}

	return false
}

// SetIqn gets a reference to the given string and assigns it to the Iqn field.
func (o *HyperflexTargetAllOf) SetIqn(v string) {
	o.Iqn = &v
}

// GetLunUuids returns the LunUuids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexTargetAllOf) GetLunUuids() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.LunUuids
}

// GetLunUuidsOk returns a tuple with the LunUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexTargetAllOf) GetLunUuidsOk() ([]string, bool) {
	if o == nil || o.LunUuids == nil {
		return nil, false
	}
	return o.LunUuids, true
}

// HasLunUuids returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasLunUuids() bool {
	if o != nil && o.LunUuids != nil {
		return true
	}

	return false
}

// SetLunUuids gets a reference to the given []string and assigns it to the LunUuids field.
func (o *HyperflexTargetAllOf) SetLunUuids(v []string) {
	o.LunUuids = v
}

// GetNumActiveInitiators returns the NumActiveInitiators field value if set, zero value otherwise.
func (o *HyperflexTargetAllOf) GetNumActiveInitiators() int64 {
	if o == nil || o.NumActiveInitiators == nil {
		var ret int64
		return ret
	}
	return *o.NumActiveInitiators
}

// GetNumActiveInitiatorsOk returns a tuple with the NumActiveInitiators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexTargetAllOf) GetNumActiveInitiatorsOk() (*int64, bool) {
	if o == nil || o.NumActiveInitiators == nil {
		return nil, false
	}
	return o.NumActiveInitiators, true
}

// HasNumActiveInitiators returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasNumActiveInitiators() bool {
	if o != nil && o.NumActiveInitiators != nil {
		return true
	}

	return false
}

// SetNumActiveInitiators gets a reference to the given int64 and assigns it to the NumActiveInitiators field.
func (o *HyperflexTargetAllOf) SetNumActiveInitiators(v int64) {
	o.NumActiveInitiators = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexTargetAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexTargetAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexTargetAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexTargetAllOf) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexTargetAllOf) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *HyperflexTargetAllOf) SetVersion(v int64) {
	o.Version = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexTargetAllOf) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexTargetAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexTargetAllOf) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetInitiatorGroups returns the InitiatorGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexTargetAllOf) GetInitiatorGroups() []HyperflexInitiatorGroupRelationship {
	if o == nil {
		var ret []HyperflexInitiatorGroupRelationship
		return ret
	}
	return o.InitiatorGroups
}

// GetInitiatorGroupsOk returns a tuple with the InitiatorGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexTargetAllOf) GetInitiatorGroupsOk() ([]HyperflexInitiatorGroupRelationship, bool) {
	if o == nil || o.InitiatorGroups == nil {
		return nil, false
	}
	return o.InitiatorGroups, true
}

// HasInitiatorGroups returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasInitiatorGroups() bool {
	if o != nil && o.InitiatorGroups != nil {
		return true
	}

	return false
}

// SetInitiatorGroups gets a reference to the given []HyperflexInitiatorGroupRelationship and assigns it to the InitiatorGroups field.
func (o *HyperflexTargetAllOf) SetInitiatorGroups(v []HyperflexInitiatorGroupRelationship) {
	o.InitiatorGroups = v
}

// GetLuns returns the Luns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexTargetAllOf) GetLuns() []HyperflexLunRelationship {
	if o == nil {
		var ret []HyperflexLunRelationship
		return ret
	}
	return o.Luns
}

// GetLunsOk returns a tuple with the Luns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexTargetAllOf) GetLunsOk() ([]HyperflexLunRelationship, bool) {
	if o == nil || o.Luns == nil {
		return nil, false
	}
	return o.Luns, true
}

// HasLuns returns a boolean if a field has been set.
func (o *HyperflexTargetAllOf) HasLuns() bool {
	if o != nil && o.Luns != nil {
		return true
	}

	return false
}

// SetLuns gets a reference to the given []HyperflexLunRelationship and assigns it to the Luns field.
func (o *HyperflexTargetAllOf) SetLuns(v []HyperflexLunRelationship) {
	o.Luns = v
}

func (o HyperflexTargetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AuthMethod != nil {
		toSerialize["AuthMethod"] = o.AuthMethod
	}
	if o.InitiatorGroupUuids != nil {
		toSerialize["InitiatorGroupUuids"] = o.InitiatorGroupUuids
	}
	if o.InventorySource != nil {
		toSerialize["InventorySource"] = o.InventorySource
	}
	if o.Iqn != nil {
		toSerialize["Iqn"] = o.Iqn
	}
	if o.LunUuids != nil {
		toSerialize["LunUuids"] = o.LunUuids
	}
	if o.NumActiveInitiators != nil {
		toSerialize["NumActiveInitiators"] = o.NumActiveInitiators
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.InitiatorGroups != nil {
		toSerialize["InitiatorGroups"] = o.InitiatorGroups
	}
	if o.Luns != nil {
		toSerialize["Luns"] = o.Luns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexTargetAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexTargetAllOf := _HyperflexTargetAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexTargetAllOf); err == nil {
		*o = HyperflexTargetAllOf(varHyperflexTargetAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthMethod")
		delete(additionalProperties, "InitiatorGroupUuids")
		delete(additionalProperties, "InventorySource")
		delete(additionalProperties, "Iqn")
		delete(additionalProperties, "LunUuids")
		delete(additionalProperties, "NumActiveInitiators")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "InitiatorGroups")
		delete(additionalProperties, "Luns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexTargetAllOf struct {
	value *HyperflexTargetAllOf
	isSet bool
}

func (v NullableHyperflexTargetAllOf) Get() *HyperflexTargetAllOf {
	return v.value
}

func (v *NullableHyperflexTargetAllOf) Set(val *HyperflexTargetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexTargetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexTargetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexTargetAllOf(val *HyperflexTargetAllOf) *NullableHyperflexTargetAllOf {
	return &NullableHyperflexTargetAllOf{value: val, isSet: true}
}

func (v NullableHyperflexTargetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexTargetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
