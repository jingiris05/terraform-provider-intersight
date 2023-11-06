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

// ApplianceFileSystemOpStatusAllOf Definition of the list of properties defined in 'appliance.FileSystemOpStatus', excluding properties defined in parent classes.
type ApplianceFileSystemOpStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Capacity of the file system in bytes.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Mount point of this file system.
	Mountpoint *string `json:"Mountpoint,omitempty"`
	// Operational status of the file system. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - The status of the appliance node is unknown. * `Operational` - The appliance node is operational. * `Impaired` - The appliance node is impaired. * `AttentionNeeded` - The appliance node needs attention. * `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * `OutOfService` - The user has taken this node (part of a cluster) to out of service. * `ReadyForReplacement` - The cluster node is ready to be replaced. * `ReplacementInProgress` - The cluster node replacement is in progress. * `ReplacementFailed` - There was a failure during the cluster node replacement.
	OperationalStatus *string `json:"OperationalStatus,omitempty"`
	// Percentage of the file system capacity currently in use.
	Usage                *float32                             `json:"Usage,omitempty"`
	NodeOpStatus         *ApplianceNodeOpStatusRelationship   `json:"NodeOpStatus,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceFileSystemOpStatusAllOf ApplianceFileSystemOpStatusAllOf

// NewApplianceFileSystemOpStatusAllOf instantiates a new ApplianceFileSystemOpStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceFileSystemOpStatusAllOf(classId string, objectType string) *ApplianceFileSystemOpStatusAllOf {
	this := ApplianceFileSystemOpStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceFileSystemOpStatusAllOfWithDefaults instantiates a new ApplianceFileSystemOpStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceFileSystemOpStatusAllOfWithDefaults() *ApplianceFileSystemOpStatusAllOf {
	this := ApplianceFileSystemOpStatusAllOf{}
	var classId string = "appliance.FileSystemOpStatus"
	this.ClassId = classId
	var objectType string = "appliance.FileSystemOpStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceFileSystemOpStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceFileSystemOpStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceFileSystemOpStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceFileSystemOpStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *ApplianceFileSystemOpStatusAllOf) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatusAllOf) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatusAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *ApplianceFileSystemOpStatusAllOf) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetMountpoint returns the Mountpoint field value if set, zero value otherwise.
func (o *ApplianceFileSystemOpStatusAllOf) GetMountpoint() string {
	if o == nil || o.Mountpoint == nil {
		var ret string
		return ret
	}
	return *o.Mountpoint
}

// GetMountpointOk returns a tuple with the Mountpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatusAllOf) GetMountpointOk() (*string, bool) {
	if o == nil || o.Mountpoint == nil {
		return nil, false
	}
	return o.Mountpoint, true
}

// HasMountpoint returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatusAllOf) HasMountpoint() bool {
	if o != nil && o.Mountpoint != nil {
		return true
	}

	return false
}

// SetMountpoint gets a reference to the given string and assigns it to the Mountpoint field.
func (o *ApplianceFileSystemOpStatusAllOf) SetMountpoint(v string) {
	o.Mountpoint = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceFileSystemOpStatusAllOf) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatusAllOf) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatusAllOf) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceFileSystemOpStatusAllOf) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ApplianceFileSystemOpStatusAllOf) GetUsage() float32 {
	if o == nil || o.Usage == nil {
		var ret float32
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatusAllOf) GetUsageOk() (*float32, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatusAllOf) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given float32 and assigns it to the Usage field.
func (o *ApplianceFileSystemOpStatusAllOf) SetUsage(v float32) {
	o.Usage = &v
}

// GetNodeOpStatus returns the NodeOpStatus field value if set, zero value otherwise.
func (o *ApplianceFileSystemOpStatusAllOf) GetNodeOpStatus() ApplianceNodeOpStatusRelationship {
	if o == nil || o.NodeOpStatus == nil {
		var ret ApplianceNodeOpStatusRelationship
		return ret
	}
	return *o.NodeOpStatus
}

// GetNodeOpStatusOk returns a tuple with the NodeOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatusAllOf) GetNodeOpStatusOk() (*ApplianceNodeOpStatusRelationship, bool) {
	if o == nil || o.NodeOpStatus == nil {
		return nil, false
	}
	return o.NodeOpStatus, true
}

// HasNodeOpStatus returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatusAllOf) HasNodeOpStatus() bool {
	if o != nil && o.NodeOpStatus != nil {
		return true
	}

	return false
}

// SetNodeOpStatus gets a reference to the given ApplianceNodeOpStatusRelationship and assigns it to the NodeOpStatus field.
func (o *ApplianceFileSystemOpStatusAllOf) SetNodeOpStatus(v ApplianceNodeOpStatusRelationship) {
	o.NodeOpStatus = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ApplianceFileSystemOpStatusAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemOpStatusAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ApplianceFileSystemOpStatusAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ApplianceFileSystemOpStatusAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o ApplianceFileSystemOpStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Mountpoint != nil {
		toSerialize["Mountpoint"] = o.Mountpoint
	}
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if o.Usage != nil {
		toSerialize["Usage"] = o.Usage
	}
	if o.NodeOpStatus != nil {
		toSerialize["NodeOpStatus"] = o.NodeOpStatus
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceFileSystemOpStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceFileSystemOpStatusAllOf := _ApplianceFileSystemOpStatusAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceFileSystemOpStatusAllOf); err == nil {
		*o = ApplianceFileSystemOpStatusAllOf(varApplianceFileSystemOpStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Mountpoint")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "NodeOpStatus")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceFileSystemOpStatusAllOf struct {
	value *ApplianceFileSystemOpStatusAllOf
	isSet bool
}

func (v NullableApplianceFileSystemOpStatusAllOf) Get() *ApplianceFileSystemOpStatusAllOf {
	return v.value
}

func (v *NullableApplianceFileSystemOpStatusAllOf) Set(val *ApplianceFileSystemOpStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceFileSystemOpStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceFileSystemOpStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceFileSystemOpStatusAllOf(val *ApplianceFileSystemOpStatusAllOf) *NullableApplianceFileSystemOpStatusAllOf {
	return &NullableApplianceFileSystemOpStatusAllOf{value: val, isSet: true}
}

func (v NullableApplianceFileSystemOpStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceFileSystemOpStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
