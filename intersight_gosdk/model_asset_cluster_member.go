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
	"reflect"
	"strings"
)

// AssetClusterMember A node within a cluster of device connectors. A Device Registration may contain multiple ClusterMembers with each holding the connection details of the device connector as well as the nodes current leadership within the cluster.
type AssetClusterMember struct {
	AssetDeviceConnection
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The current leadershipstate of this member. Updated by the device connector on failover or leadership change. If a member is elected as Primary within the cluster this connection will be the same as the DeviceRegistration connection. E.g a message addressed to the DeviceRegistration will be forwarded to the ClusterMember connection. * `Unknown` - The node is unable to complete election or determine the current state. If the device has been registered before and the node has access to the current credentials it will establish a connection to Intersight with limited capabilities that can be used to debug the HA failure from Intersight. * `Primary` - The node has been elected as the primary and will establish a connection to the Intersight service and accept all message types enabled for a primary node. There can only be one primary in a given cluster, while the underlying platform may be active-active only one connector will assume the primary role. * `Secondary` - The node has been elected as a secondary node in the cluster. The device connector will establish a connection to the Intersight service with limited capabilities. e.g. file upload will be enabled, but requests to the underlying platform management will be disabled.
	Leadership *string `json:"Leadership,omitempty"`
	// Devices lock their leadership on failure to heartbeat with peers. Value acts as a third party tie breaker in election between nodes. Intersight enforces that only one cluster member exists with this value set to true.
	LockedLeader *bool `json:"LockedLeader,omitempty"`
	// The unique identity of the member within the cluster. The identity is retrieved from the platform and reported by the device connector at connection time.
	MemberIdentity *string `json:"MemberIdentity,omitempty"`
	// The member idenity of the cluster member through which this device is connected if applicable.
	ParentClusterMemberIdentity *string                              `json:"ParentClusterMemberIdentity,omitempty"`
	Sudi                        NullableAssetSudiInfo                `json:"Sudi,omitempty"`
	Device                      *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _AssetClusterMember AssetClusterMember

// NewAssetClusterMember instantiates a new AssetClusterMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetClusterMember(classId string, objectType string) *AssetClusterMember {
	this := AssetClusterMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	var leadership string = "Unknown"
	this.Leadership = &leadership
	return &this
}

// NewAssetClusterMemberWithDefaults instantiates a new AssetClusterMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetClusterMemberWithDefaults() *AssetClusterMember {
	this := AssetClusterMember{}
	var classId string = "asset.ClusterMember"
	this.ClassId = classId
	var objectType string = "asset.ClusterMember"
	this.ObjectType = objectType
	var leadership string = "Unknown"
	this.Leadership = &leadership
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetClusterMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetClusterMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetClusterMember) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetClusterMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetClusterMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetClusterMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLeadership returns the Leadership field value if set, zero value otherwise.
func (o *AssetClusterMember) GetLeadership() string {
	if o == nil || o.Leadership == nil {
		var ret string
		return ret
	}
	return *o.Leadership
}

// GetLeadershipOk returns a tuple with the Leadership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetClusterMember) GetLeadershipOk() (*string, bool) {
	if o == nil || o.Leadership == nil {
		return nil, false
	}
	return o.Leadership, true
}

// HasLeadership returns a boolean if a field has been set.
func (o *AssetClusterMember) HasLeadership() bool {
	if o != nil && o.Leadership != nil {
		return true
	}

	return false
}

// SetLeadership gets a reference to the given string and assigns it to the Leadership field.
func (o *AssetClusterMember) SetLeadership(v string) {
	o.Leadership = &v
}

// GetLockedLeader returns the LockedLeader field value if set, zero value otherwise.
func (o *AssetClusterMember) GetLockedLeader() bool {
	if o == nil || o.LockedLeader == nil {
		var ret bool
		return ret
	}
	return *o.LockedLeader
}

// GetLockedLeaderOk returns a tuple with the LockedLeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetClusterMember) GetLockedLeaderOk() (*bool, bool) {
	if o == nil || o.LockedLeader == nil {
		return nil, false
	}
	return o.LockedLeader, true
}

// HasLockedLeader returns a boolean if a field has been set.
func (o *AssetClusterMember) HasLockedLeader() bool {
	if o != nil && o.LockedLeader != nil {
		return true
	}

	return false
}

// SetLockedLeader gets a reference to the given bool and assigns it to the LockedLeader field.
func (o *AssetClusterMember) SetLockedLeader(v bool) {
	o.LockedLeader = &v
}

// GetMemberIdentity returns the MemberIdentity field value if set, zero value otherwise.
func (o *AssetClusterMember) GetMemberIdentity() string {
	if o == nil || o.MemberIdentity == nil {
		var ret string
		return ret
	}
	return *o.MemberIdentity
}

// GetMemberIdentityOk returns a tuple with the MemberIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetClusterMember) GetMemberIdentityOk() (*string, bool) {
	if o == nil || o.MemberIdentity == nil {
		return nil, false
	}
	return o.MemberIdentity, true
}

// HasMemberIdentity returns a boolean if a field has been set.
func (o *AssetClusterMember) HasMemberIdentity() bool {
	if o != nil && o.MemberIdentity != nil {
		return true
	}

	return false
}

// SetMemberIdentity gets a reference to the given string and assigns it to the MemberIdentity field.
func (o *AssetClusterMember) SetMemberIdentity(v string) {
	o.MemberIdentity = &v
}

// GetParentClusterMemberIdentity returns the ParentClusterMemberIdentity field value if set, zero value otherwise.
func (o *AssetClusterMember) GetParentClusterMemberIdentity() string {
	if o == nil || o.ParentClusterMemberIdentity == nil {
		var ret string
		return ret
	}
	return *o.ParentClusterMemberIdentity
}

// GetParentClusterMemberIdentityOk returns a tuple with the ParentClusterMemberIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetClusterMember) GetParentClusterMemberIdentityOk() (*string, bool) {
	if o == nil || o.ParentClusterMemberIdentity == nil {
		return nil, false
	}
	return o.ParentClusterMemberIdentity, true
}

// HasParentClusterMemberIdentity returns a boolean if a field has been set.
func (o *AssetClusterMember) HasParentClusterMemberIdentity() bool {
	if o != nil && o.ParentClusterMemberIdentity != nil {
		return true
	}

	return false
}

// SetParentClusterMemberIdentity gets a reference to the given string and assigns it to the ParentClusterMemberIdentity field.
func (o *AssetClusterMember) SetParentClusterMemberIdentity(v string) {
	o.ParentClusterMemberIdentity = &v
}

// GetSudi returns the Sudi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetClusterMember) GetSudi() AssetSudiInfo {
	if o == nil || o.Sudi.Get() == nil {
		var ret AssetSudiInfo
		return ret
	}
	return *o.Sudi.Get()
}

// GetSudiOk returns a tuple with the Sudi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetClusterMember) GetSudiOk() (*AssetSudiInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sudi.Get(), o.Sudi.IsSet()
}

// HasSudi returns a boolean if a field has been set.
func (o *AssetClusterMember) HasSudi() bool {
	if o != nil && o.Sudi.IsSet() {
		return true
	}

	return false
}

// SetSudi gets a reference to the given NullableAssetSudiInfo and assigns it to the Sudi field.
func (o *AssetClusterMember) SetSudi(v AssetSudiInfo) {
	o.Sudi.Set(&v)
}

// SetSudiNil sets the value for Sudi to be an explicit nil
func (o *AssetClusterMember) SetSudiNil() {
	o.Sudi.Set(nil)
}

// UnsetSudi ensures that no value is present for Sudi, not even an explicit nil
func (o *AssetClusterMember) UnsetSudi() {
	o.Sudi.Unset()
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *AssetClusterMember) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetClusterMember) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *AssetClusterMember) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *AssetClusterMember) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

func (o AssetClusterMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetDeviceConnection, errAssetDeviceConnection := json.Marshal(o.AssetDeviceConnection)
	if errAssetDeviceConnection != nil {
		return []byte{}, errAssetDeviceConnection
	}
	errAssetDeviceConnection = json.Unmarshal([]byte(serializedAssetDeviceConnection), &toSerialize)
	if errAssetDeviceConnection != nil {
		return []byte{}, errAssetDeviceConnection
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Leadership != nil {
		toSerialize["Leadership"] = o.Leadership
	}
	if o.LockedLeader != nil {
		toSerialize["LockedLeader"] = o.LockedLeader
	}
	if o.MemberIdentity != nil {
		toSerialize["MemberIdentity"] = o.MemberIdentity
	}
	if o.ParentClusterMemberIdentity != nil {
		toSerialize["ParentClusterMemberIdentity"] = o.ParentClusterMemberIdentity
	}
	if o.Sudi.IsSet() {
		toSerialize["Sudi"] = o.Sudi.Get()
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetClusterMember) UnmarshalJSON(bytes []byte) (err error) {
	type AssetClusterMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The current leadershipstate of this member. Updated by the device connector on failover or leadership change. If a member is elected as Primary within the cluster this connection will be the same as the DeviceRegistration connection. E.g a message addressed to the DeviceRegistration will be forwarded to the ClusterMember connection. * `Unknown` - The node is unable to complete election or determine the current state. If the device has been registered before and the node has access to the current credentials it will establish a connection to Intersight with limited capabilities that can be used to debug the HA failure from Intersight. * `Primary` - The node has been elected as the primary and will establish a connection to the Intersight service and accept all message types enabled for a primary node. There can only be one primary in a given cluster, while the underlying platform may be active-active only one connector will assume the primary role. * `Secondary` - The node has been elected as a secondary node in the cluster. The device connector will establish a connection to the Intersight service with limited capabilities. e.g. file upload will be enabled, but requests to the underlying platform management will be disabled.
		Leadership *string `json:"Leadership,omitempty"`
		// Devices lock their leadership on failure to heartbeat with peers. Value acts as a third party tie breaker in election between nodes. Intersight enforces that only one cluster member exists with this value set to true.
		LockedLeader *bool `json:"LockedLeader,omitempty"`
		// The unique identity of the member within the cluster. The identity is retrieved from the platform and reported by the device connector at connection time.
		MemberIdentity *string `json:"MemberIdentity,omitempty"`
		// The member idenity of the cluster member through which this device is connected if applicable.
		ParentClusterMemberIdentity *string                              `json:"ParentClusterMemberIdentity,omitempty"`
		Sudi                        NullableAssetSudiInfo                `json:"Sudi,omitempty"`
		Device                      *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	}

	varAssetClusterMemberWithoutEmbeddedStruct := AssetClusterMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetClusterMemberWithoutEmbeddedStruct)
	if err == nil {
		varAssetClusterMember := _AssetClusterMember{}
		varAssetClusterMember.ClassId = varAssetClusterMemberWithoutEmbeddedStruct.ClassId
		varAssetClusterMember.ObjectType = varAssetClusterMemberWithoutEmbeddedStruct.ObjectType
		varAssetClusterMember.Leadership = varAssetClusterMemberWithoutEmbeddedStruct.Leadership
		varAssetClusterMember.LockedLeader = varAssetClusterMemberWithoutEmbeddedStruct.LockedLeader
		varAssetClusterMember.MemberIdentity = varAssetClusterMemberWithoutEmbeddedStruct.MemberIdentity
		varAssetClusterMember.ParentClusterMemberIdentity = varAssetClusterMemberWithoutEmbeddedStruct.ParentClusterMemberIdentity
		varAssetClusterMember.Sudi = varAssetClusterMemberWithoutEmbeddedStruct.Sudi
		varAssetClusterMember.Device = varAssetClusterMemberWithoutEmbeddedStruct.Device
		*o = AssetClusterMember(varAssetClusterMember)
	} else {
		return err
	}

	varAssetClusterMember := _AssetClusterMember{}

	err = json.Unmarshal(bytes, &varAssetClusterMember)
	if err == nil {
		o.AssetDeviceConnection = varAssetClusterMember.AssetDeviceConnection
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Leadership")
		delete(additionalProperties, "LockedLeader")
		delete(additionalProperties, "MemberIdentity")
		delete(additionalProperties, "ParentClusterMemberIdentity")
		delete(additionalProperties, "Sudi")
		delete(additionalProperties, "Device")

		// remove fields from embedded structs
		reflectAssetDeviceConnection := reflect.ValueOf(o.AssetDeviceConnection)
		for i := 0; i < reflectAssetDeviceConnection.Type().NumField(); i++ {
			t := reflectAssetDeviceConnection.Type().Field(i)

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

type NullableAssetClusterMember struct {
	value *AssetClusterMember
	isSet bool
}

func (v NullableAssetClusterMember) Get() *AssetClusterMember {
	return v.value
}

func (v *NullableAssetClusterMember) Set(val *AssetClusterMember) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetClusterMember) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetClusterMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetClusterMember(val *AssetClusterMember) *NullableAssetClusterMember {
	return &NullableAssetClusterMember{value: val, isSet: true}
}

func (v NullableAssetClusterMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetClusterMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
