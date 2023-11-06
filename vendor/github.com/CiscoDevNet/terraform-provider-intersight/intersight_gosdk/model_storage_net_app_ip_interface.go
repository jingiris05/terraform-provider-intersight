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
	"reflect"
	"strings"
)

// StorageNetAppIpInterface NetApp IP interface is a logical interface.
type StorageNetAppIpInterface struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP interface is enabled or not.
	Enabled *string `json:"Enabled,omitempty"`
	// Name of home node of IP interface.
	HomeNode *string `json:"HomeNode,omitempty"`
	// Name of home port of IP interface.
	HomePort *string `json:"HomePort,omitempty"`
	// The IP address of interface.
	IpAddress *string `json:"IpAddress,omitempty"`
	// IP address family of interface. * `IPv4` - IP address family type is IPv4. * `IPv6` - IP address family type is IP6.
	IpFamily *string `json:"IpFamily,omitempty"`
	// The name of the IP interface.
	Name *string `json:"Name,omitempty"`
	// The netmask of the interface.
	Netmask *string `json:"Netmask,omitempty"`
	// Service policy name of IP interface.
	ServicePolicyName *string `json:"ServicePolicyName,omitempty"`
	// Service policy UUID of IP interface.
	ServicePolicyUuid *string  `json:"ServicePolicyUuid,omitempty"`
	Services          []string `json:"Services,omitempty"`
	// The state of the IP interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
	State *string `json:"State,omitempty"`
	// Uuid of NetApp IP Interface.
	Uuid            *string                        `json:"Uuid,omitempty"`
	ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppIpInterfaceEvent resources.
	Events               []StorageNetAppIpInterfaceEventRelationship `json:"Events,omitempty"`
	NetAppEthernetPort   *StorageNetAppEthernetPortRelationship      `json:"NetAppEthernetPort,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship         `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppIpInterface StorageNetAppIpInterface

// NewStorageNetAppIpInterface instantiates a new StorageNetAppIpInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppIpInterface(classId string, objectType string) *StorageNetAppIpInterface {
	this := StorageNetAppIpInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppIpInterfaceWithDefaults instantiates a new StorageNetAppIpInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppIpInterfaceWithDefaults() *StorageNetAppIpInterface {
	this := StorageNetAppIpInterface{}
	var classId string = "storage.NetAppIpInterface"
	this.ClassId = classId
	var objectType string = "storage.NetAppIpInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppIpInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppIpInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppIpInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppIpInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetEnabled() string {
	if o == nil || o.Enabled == nil {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetEnabledOk() (*string, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppIpInterface) SetEnabled(v string) {
	o.Enabled = &v
}

// GetHomeNode returns the HomeNode field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetHomeNode() string {
	if o == nil || o.HomeNode == nil {
		var ret string
		return ret
	}
	return *o.HomeNode
}

// GetHomeNodeOk returns a tuple with the HomeNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetHomeNodeOk() (*string, bool) {
	if o == nil || o.HomeNode == nil {
		return nil, false
	}
	return o.HomeNode, true
}

// HasHomeNode returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasHomeNode() bool {
	if o != nil && o.HomeNode != nil {
		return true
	}

	return false
}

// SetHomeNode gets a reference to the given string and assigns it to the HomeNode field.
func (o *StorageNetAppIpInterface) SetHomeNode(v string) {
	o.HomeNode = &v
}

// GetHomePort returns the HomePort field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetHomePort() string {
	if o == nil || o.HomePort == nil {
		var ret string
		return ret
	}
	return *o.HomePort
}

// GetHomePortOk returns a tuple with the HomePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetHomePortOk() (*string, bool) {
	if o == nil || o.HomePort == nil {
		return nil, false
	}
	return o.HomePort, true
}

// HasHomePort returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasHomePort() bool {
	if o != nil && o.HomePort != nil {
		return true
	}

	return false
}

// SetHomePort gets a reference to the given string and assigns it to the HomePort field.
func (o *StorageNetAppIpInterface) SetHomePort(v string) {
	o.HomePort = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *StorageNetAppIpInterface) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIpFamily returns the IpFamily field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetIpFamily() string {
	if o == nil || o.IpFamily == nil {
		var ret string
		return ret
	}
	return *o.IpFamily
}

// GetIpFamilyOk returns a tuple with the IpFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetIpFamilyOk() (*string, bool) {
	if o == nil || o.IpFamily == nil {
		return nil, false
	}
	return o.IpFamily, true
}

// HasIpFamily returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasIpFamily() bool {
	if o != nil && o.IpFamily != nil {
		return true
	}

	return false
}

// SetIpFamily gets a reference to the given string and assigns it to the IpFamily field.
func (o *StorageNetAppIpInterface) SetIpFamily(v string) {
	o.IpFamily = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppIpInterface) SetName(v string) {
	o.Name = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *StorageNetAppIpInterface) SetNetmask(v string) {
	o.Netmask = &v
}

// GetServicePolicyName returns the ServicePolicyName field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetServicePolicyName() string {
	if o == nil || o.ServicePolicyName == nil {
		var ret string
		return ret
	}
	return *o.ServicePolicyName
}

// GetServicePolicyNameOk returns a tuple with the ServicePolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetServicePolicyNameOk() (*string, bool) {
	if o == nil || o.ServicePolicyName == nil {
		return nil, false
	}
	return o.ServicePolicyName, true
}

// HasServicePolicyName returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasServicePolicyName() bool {
	if o != nil && o.ServicePolicyName != nil {
		return true
	}

	return false
}

// SetServicePolicyName gets a reference to the given string and assigns it to the ServicePolicyName field.
func (o *StorageNetAppIpInterface) SetServicePolicyName(v string) {
	o.ServicePolicyName = &v
}

// GetServicePolicyUuid returns the ServicePolicyUuid field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetServicePolicyUuid() string {
	if o == nil || o.ServicePolicyUuid == nil {
		var ret string
		return ret
	}
	return *o.ServicePolicyUuid
}

// GetServicePolicyUuidOk returns a tuple with the ServicePolicyUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetServicePolicyUuidOk() (*string, bool) {
	if o == nil || o.ServicePolicyUuid == nil {
		return nil, false
	}
	return o.ServicePolicyUuid, true
}

// HasServicePolicyUuid returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasServicePolicyUuid() bool {
	if o != nil && o.ServicePolicyUuid != nil {
		return true
	}

	return false
}

// SetServicePolicyUuid gets a reference to the given string and assigns it to the ServicePolicyUuid field.
func (o *StorageNetAppIpInterface) SetServicePolicyUuid(v string) {
	o.ServicePolicyUuid = &v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppIpInterface) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppIpInterface) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *StorageNetAppIpInterface) SetServices(v []string) {
	o.Services = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppIpInterface) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppIpInterface) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppIpInterface) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppIpInterface) GetEvents() []StorageNetAppIpInterfaceEventRelationship {
	if o == nil {
		var ret []StorageNetAppIpInterfaceEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppIpInterface) GetEventsOk() ([]StorageNetAppIpInterfaceEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppIpInterfaceEventRelationship and assigns it to the Events field.
func (o *StorageNetAppIpInterface) SetEvents(v []StorageNetAppIpInterfaceEventRelationship) {
	o.Events = v
}

// GetNetAppEthernetPort returns the NetAppEthernetPort field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetNetAppEthernetPort() StorageNetAppEthernetPortRelationship {
	if o == nil || o.NetAppEthernetPort == nil {
		var ret StorageNetAppEthernetPortRelationship
		return ret
	}
	return *o.NetAppEthernetPort
}

// GetNetAppEthernetPortOk returns a tuple with the NetAppEthernetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetNetAppEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool) {
	if o == nil || o.NetAppEthernetPort == nil {
		return nil, false
	}
	return o.NetAppEthernetPort, true
}

// HasNetAppEthernetPort returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasNetAppEthernetPort() bool {
	if o != nil && o.NetAppEthernetPort != nil {
		return true
	}

	return false
}

// SetNetAppEthernetPort gets a reference to the given StorageNetAppEthernetPortRelationship and assigns it to the NetAppEthernetPort field.
func (o *StorageNetAppIpInterface) SetNetAppEthernetPort(v StorageNetAppEthernetPortRelationship) {
	o.NetAppEthernetPort = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppIpInterface) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterface) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppIpInterface) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppIpInterface) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppIpInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.HomeNode != nil {
		toSerialize["HomeNode"] = o.HomeNode
	}
	if o.HomePort != nil {
		toSerialize["HomePort"] = o.HomePort
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.IpFamily != nil {
		toSerialize["IpFamily"] = o.IpFamily
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Netmask != nil {
		toSerialize["Netmask"] = o.Netmask
	}
	if o.ServicePolicyName != nil {
		toSerialize["ServicePolicyName"] = o.ServicePolicyName
	}
	if o.ServicePolicyUuid != nil {
		toSerialize["ServicePolicyUuid"] = o.ServicePolicyUuid
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}
	if o.NetAppEthernetPort != nil {
		toSerialize["NetAppEthernetPort"] = o.NetAppEthernetPort
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppIpInterface) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppIpInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP interface is enabled or not.
		Enabled *string `json:"Enabled,omitempty"`
		// Name of home node of IP interface.
		HomeNode *string `json:"HomeNode,omitempty"`
		// Name of home port of IP interface.
		HomePort *string `json:"HomePort,omitempty"`
		// The IP address of interface.
		IpAddress *string `json:"IpAddress,omitempty"`
		// IP address family of interface. * `IPv4` - IP address family type is IPv4. * `IPv6` - IP address family type is IP6.
		IpFamily *string `json:"IpFamily,omitempty"`
		// The name of the IP interface.
		Name *string `json:"Name,omitempty"`
		// The netmask of the interface.
		Netmask *string `json:"Netmask,omitempty"`
		// Service policy name of IP interface.
		ServicePolicyName *string `json:"ServicePolicyName,omitempty"`
		// Service policy UUID of IP interface.
		ServicePolicyUuid *string  `json:"ServicePolicyUuid,omitempty"`
		Services          []string `json:"Services,omitempty"`
		// The state of the IP interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
		State *string `json:"State,omitempty"`
		// Uuid of NetApp IP Interface.
		Uuid            *string                        `json:"Uuid,omitempty"`
		ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
		// An array of relationships to storageNetAppIpInterfaceEvent resources.
		Events             []StorageNetAppIpInterfaceEventRelationship `json:"Events,omitempty"`
		NetAppEthernetPort *StorageNetAppEthernetPortRelationship      `json:"NetAppEthernetPort,omitempty"`
		Tenant             *StorageNetAppStorageVmRelationship         `json:"Tenant,omitempty"`
	}

	varStorageNetAppIpInterfaceWithoutEmbeddedStruct := StorageNetAppIpInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppIpInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppIpInterface := _StorageNetAppIpInterface{}
		varStorageNetAppIpInterface.ClassId = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.ClassId
		varStorageNetAppIpInterface.ObjectType = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.ObjectType
		varStorageNetAppIpInterface.Enabled = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Enabled
		varStorageNetAppIpInterface.HomeNode = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.HomeNode
		varStorageNetAppIpInterface.HomePort = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.HomePort
		varStorageNetAppIpInterface.IpAddress = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.IpAddress
		varStorageNetAppIpInterface.IpFamily = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.IpFamily
		varStorageNetAppIpInterface.Name = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Name
		varStorageNetAppIpInterface.Netmask = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Netmask
		varStorageNetAppIpInterface.ServicePolicyName = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.ServicePolicyName
		varStorageNetAppIpInterface.ServicePolicyUuid = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.ServicePolicyUuid
		varStorageNetAppIpInterface.Services = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Services
		varStorageNetAppIpInterface.State = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.State
		varStorageNetAppIpInterface.Uuid = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Uuid
		varStorageNetAppIpInterface.ArrayController = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.ArrayController
		varStorageNetAppIpInterface.Events = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Events
		varStorageNetAppIpInterface.NetAppEthernetPort = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.NetAppEthernetPort
		varStorageNetAppIpInterface.Tenant = varStorageNetAppIpInterfaceWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppIpInterface(varStorageNetAppIpInterface)
	} else {
		return err
	}

	varStorageNetAppIpInterface := _StorageNetAppIpInterface{}

	err = json.Unmarshal(bytes, &varStorageNetAppIpInterface)
	if err == nil {
		o.MoBaseMo = varStorageNetAppIpInterface.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "HomeNode")
		delete(additionalProperties, "HomePort")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "IpFamily")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Netmask")
		delete(additionalProperties, "ServicePolicyName")
		delete(additionalProperties, "ServicePolicyUuid")
		delete(additionalProperties, "Services")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "NetAppEthernetPort")
		delete(additionalProperties, "Tenant")

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

type NullableStorageNetAppIpInterface struct {
	value *StorageNetAppIpInterface
	isSet bool
}

func (v NullableStorageNetAppIpInterface) Get() *StorageNetAppIpInterface {
	return v.value
}

func (v *NullableStorageNetAppIpInterface) Set(val *StorageNetAppIpInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppIpInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppIpInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppIpInterface(val *StorageNetAppIpInterface) *NullableStorageNetAppIpInterface {
	return &NullableStorageNetAppIpInterface{value: val, isSet: true}
}

func (v NullableStorageNetAppIpInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppIpInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
