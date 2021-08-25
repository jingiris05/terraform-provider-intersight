/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-24T21:18:00Z.
 *
 * API version: 1.0.9-4404
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// MacpoolPoolMemberAllOf Definition of the list of properties defined in 'macpool.PoolMember', excluding properties defined in parent classes.
type MacpoolPoolMemberAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// MAC Address of this pool member.
	MacAddress           *string                     `json:"MacAddress,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship       `json:"AssignedToEntity,omitempty"`
	BlockHead            *MacpoolIdBlockRelationship `json:"BlockHead,omitempty"`
	Peer                 *MacpoolLeaseRelationship   `json:"Peer,omitempty"`
	Pool                 *MacpoolPoolRelationship    `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MacpoolPoolMemberAllOf MacpoolPoolMemberAllOf

// NewMacpoolPoolMemberAllOf instantiates a new MacpoolPoolMemberAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacpoolPoolMemberAllOf(classId string, objectType string) *MacpoolPoolMemberAllOf {
	this := MacpoolPoolMemberAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMacpoolPoolMemberAllOfWithDefaults instantiates a new MacpoolPoolMemberAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacpoolPoolMemberAllOfWithDefaults() *MacpoolPoolMemberAllOf {
	this := MacpoolPoolMemberAllOf{}
	var classId string = "macpool.PoolMember"
	this.ClassId = classId
	var objectType string = "macpool.PoolMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MacpoolPoolMemberAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MacpoolPoolMemberAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MacpoolPoolMemberAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MacpoolPoolMemberAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MacpoolPoolMemberAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MacpoolPoolMemberAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *MacpoolPoolMemberAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolPoolMemberAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *MacpoolPoolMemberAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *MacpoolPoolMemberAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *MacpoolPoolMemberAllOf) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolPoolMemberAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *MacpoolPoolMemberAllOf) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *MacpoolPoolMemberAllOf) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetBlockHead returns the BlockHead field value if set, zero value otherwise.
func (o *MacpoolPoolMemberAllOf) GetBlockHead() MacpoolIdBlockRelationship {
	if o == nil || o.BlockHead == nil {
		var ret MacpoolIdBlockRelationship
		return ret
	}
	return *o.BlockHead
}

// GetBlockHeadOk returns a tuple with the BlockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolPoolMemberAllOf) GetBlockHeadOk() (*MacpoolIdBlockRelationship, bool) {
	if o == nil || o.BlockHead == nil {
		return nil, false
	}
	return o.BlockHead, true
}

// HasBlockHead returns a boolean if a field has been set.
func (o *MacpoolPoolMemberAllOf) HasBlockHead() bool {
	if o != nil && o.BlockHead != nil {
		return true
	}

	return false
}

// SetBlockHead gets a reference to the given MacpoolIdBlockRelationship and assigns it to the BlockHead field.
func (o *MacpoolPoolMemberAllOf) SetBlockHead(v MacpoolIdBlockRelationship) {
	o.BlockHead = &v
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *MacpoolPoolMemberAllOf) GetPeer() MacpoolLeaseRelationship {
	if o == nil || o.Peer == nil {
		var ret MacpoolLeaseRelationship
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolPoolMemberAllOf) GetPeerOk() (*MacpoolLeaseRelationship, bool) {
	if o == nil || o.Peer == nil {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *MacpoolPoolMemberAllOf) HasPeer() bool {
	if o != nil && o.Peer != nil {
		return true
	}

	return false
}

// SetPeer gets a reference to the given MacpoolLeaseRelationship and assigns it to the Peer field.
func (o *MacpoolPoolMemberAllOf) SetPeer(v MacpoolLeaseRelationship) {
	o.Peer = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *MacpoolPoolMemberAllOf) GetPool() MacpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret MacpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolPoolMemberAllOf) GetPoolOk() (*MacpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *MacpoolPoolMemberAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given MacpoolPoolRelationship and assigns it to the Pool field.
func (o *MacpoolPoolMemberAllOf) SetPool(v MacpoolPoolRelationship) {
	o.Pool = &v
}

func (o MacpoolPoolMemberAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.BlockHead != nil {
		toSerialize["BlockHead"] = o.BlockHead
	}
	if o.Peer != nil {
		toSerialize["Peer"] = o.Peer
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MacpoolPoolMemberAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMacpoolPoolMemberAllOf := _MacpoolPoolMemberAllOf{}

	if err = json.Unmarshal(bytes, &varMacpoolPoolMemberAllOf); err == nil {
		*o = MacpoolPoolMemberAllOf(varMacpoolPoolMemberAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "BlockHead")
		delete(additionalProperties, "Peer")
		delete(additionalProperties, "Pool")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMacpoolPoolMemberAllOf struct {
	value *MacpoolPoolMemberAllOf
	isSet bool
}

func (v NullableMacpoolPoolMemberAllOf) Get() *MacpoolPoolMemberAllOf {
	return v.value
}

func (v *NullableMacpoolPoolMemberAllOf) Set(val *MacpoolPoolMemberAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolPoolMemberAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolPoolMemberAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolPoolMemberAllOf(val *MacpoolPoolMemberAllOf) *NullableMacpoolPoolMemberAllOf {
	return &NullableMacpoolPoolMemberAllOf{value: val, isSet: true}
}

func (v NullableMacpoolPoolMemberAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolPoolMemberAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
