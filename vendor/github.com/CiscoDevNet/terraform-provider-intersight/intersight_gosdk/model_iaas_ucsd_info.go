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

// checks if the IaasUcsdInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IaasUcsdInfo{}

// IaasUcsdInfo UCS Director accounts managed by Intersight.
type IaasUcsdInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Moid of the UCS Director device connector's asset.DeviceRegistration.
	DeviceId *string `json:"DeviceId,omitempty"`
	// Unique ID of UCS Director getting registerd with Intersight.
	Guid *string `json:"Guid,omitempty"`
	// The UCS Director hostname for management.
	HostName *string `json:"HostName,omitempty"`
	// The UCS Director IP address for management.
	Ip *string `json:"Ip,omitempty"`
	// Last successful backup created for this UCS Director appliance if backup is configured.
	LastBackup *time.Time `json:"LastBackup,omitempty"`
	// NodeType specifies if UCS Director is deployed in Stand-alone or Multi Node.
	NodeType *string `json:"NodeType,omitempty"`
	// The UCS Director product name.
	ProductName *string `json:"ProductName,omitempty"`
	// The UCS Director product vendor.
	ProductVendor *string `json:"ProductVendor,omitempty"`
	// The UCS Director product/platform version.
	ProductVersion *string `json:"ProductVersion,omitempty"`
	// The UCS Director status. Possible values are Active, Inactive, Unknown.
	Status *string `json:"Status,omitempty"`
	// An array of relationships to iaasConnectorPack resources.
	ConnectorPack []IaasConnectorPackRelationship `json:"ConnectorPack,omitempty"`
	// An array of relationships to iaasCustomTaskInfo resources.
	CustomTaskInfo []IaasCustomTaskInfoRelationship `json:"CustomTaskInfo,omitempty"`
	// An array of relationships to iaasDeviceStatus resources.
	DeviceStatus []IaasDeviceStatusRelationship      `json:"DeviceStatus,omitempty"`
	LicenseInfo  NullableIaasLicenseInfoRelationship `json:"LicenseInfo,omitempty"`
	// An array of relationships to iaasMostRunTasks resources.
	MostRunTasks     []IaasMostRunTasksRelationship              `json:"MostRunTasks,omitempty"`
	RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to iaasSystemTaskInfo resources.
	SystemTaskInfo       []IaasSystemTaskInfoRelationship         `json:"SystemTaskInfo,omitempty"`
	UcsdManagedInfra     NullableIaasUcsdManagedInfraRelationship `json:"UcsdManagedInfra,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasUcsdInfo IaasUcsdInfo

// NewIaasUcsdInfo instantiates a new IaasUcsdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasUcsdInfo(classId string, objectType string) *IaasUcsdInfo {
	this := IaasUcsdInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasUcsdInfoWithDefaults instantiates a new IaasUcsdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasUcsdInfoWithDefaults() *IaasUcsdInfo {
	this := IaasUcsdInfo{}
	var classId string = "iaas.UcsdInfo"
	this.ClassId = classId
	var objectType string = "iaas.UcsdInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasUcsdInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasUcsdInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iaas.UcsdInfo" of the ClassId field.
func (o *IaasUcsdInfo) GetDefaultClassId() interface{} {
	return "iaas.UcsdInfo"
}

// GetObjectType returns the ObjectType field value
func (o *IaasUcsdInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasUcsdInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iaas.UcsdInfo" of the ObjectType field.
func (o *IaasUcsdInfo) GetDefaultObjectType() interface{} {
	return "iaas.UcsdInfo"
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *IaasUcsdInfo) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *IaasUcsdInfo) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasUcsdInfo) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *IaasUcsdInfo) SetGuid(v string) {
	o.Guid = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *IaasUcsdInfo) GetHostName() string {
	if o == nil || IsNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *IaasUcsdInfo) SetHostName(v string) {
	o.HostName = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *IaasUcsdInfo) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *IaasUcsdInfo) SetIp(v string) {
	o.Ip = &v
}

// GetLastBackup returns the LastBackup field value if set, zero value otherwise.
func (o *IaasUcsdInfo) GetLastBackup() time.Time {
	if o == nil || IsNil(o.LastBackup) {
		var ret time.Time
		return ret
	}
	return *o.LastBackup
}

// GetLastBackupOk returns a tuple with the LastBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetLastBackupOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastBackup) {
		return nil, false
	}
	return o.LastBackup, true
}

// HasLastBackup returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasLastBackup() bool {
	if o != nil && !IsNil(o.LastBackup) {
		return true
	}

	return false
}

// SetLastBackup gets a reference to the given time.Time and assigns it to the LastBackup field.
func (o *IaasUcsdInfo) SetLastBackup(v time.Time) {
	o.LastBackup = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *IaasUcsdInfo) GetNodeType() string {
	if o == nil || IsNil(o.NodeType) {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetNodeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NodeType) {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasNodeType() bool {
	if o != nil && !IsNil(o.NodeType) {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *IaasUcsdInfo) SetNodeType(v string) {
	o.NodeType = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *IaasUcsdInfo) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *IaasUcsdInfo) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductVendor returns the ProductVendor field value if set, zero value otherwise.
func (o *IaasUcsdInfo) GetProductVendor() string {
	if o == nil || IsNil(o.ProductVendor) {
		var ret string
		return ret
	}
	return *o.ProductVendor
}

// GetProductVendorOk returns a tuple with the ProductVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetProductVendorOk() (*string, bool) {
	if o == nil || IsNil(o.ProductVendor) {
		return nil, false
	}
	return o.ProductVendor, true
}

// HasProductVendor returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasProductVendor() bool {
	if o != nil && !IsNil(o.ProductVendor) {
		return true
	}

	return false
}

// SetProductVendor gets a reference to the given string and assigns it to the ProductVendor field.
func (o *IaasUcsdInfo) SetProductVendor(v string) {
	o.ProductVendor = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *IaasUcsdInfo) GetProductVersion() string {
	if o == nil || IsNil(o.ProductVersion) {
		var ret string
		return ret
	}
	return *o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetProductVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductVersion) {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasProductVersion() bool {
	if o != nil && !IsNil(o.ProductVersion) {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.
func (o *IaasUcsdInfo) SetProductVersion(v string) {
	o.ProductVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IaasUcsdInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasUcsdInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IaasUcsdInfo) SetStatus(v string) {
	o.Status = &v
}

// GetConnectorPack returns the ConnectorPack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfo) GetConnectorPack() []IaasConnectorPackRelationship {
	if o == nil {
		var ret []IaasConnectorPackRelationship
		return ret
	}
	return o.ConnectorPack
}

// GetConnectorPackOk returns a tuple with the ConnectorPack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfo) GetConnectorPackOk() ([]IaasConnectorPackRelationship, bool) {
	if o == nil || IsNil(o.ConnectorPack) {
		return nil, false
	}
	return o.ConnectorPack, true
}

// HasConnectorPack returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasConnectorPack() bool {
	if o != nil && !IsNil(o.ConnectorPack) {
		return true
	}

	return false
}

// SetConnectorPack gets a reference to the given []IaasConnectorPackRelationship and assigns it to the ConnectorPack field.
func (o *IaasUcsdInfo) SetConnectorPack(v []IaasConnectorPackRelationship) {
	o.ConnectorPack = v
}

// GetCustomTaskInfo returns the CustomTaskInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfo) GetCustomTaskInfo() []IaasCustomTaskInfoRelationship {
	if o == nil {
		var ret []IaasCustomTaskInfoRelationship
		return ret
	}
	return o.CustomTaskInfo
}

// GetCustomTaskInfoOk returns a tuple with the CustomTaskInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfo) GetCustomTaskInfoOk() ([]IaasCustomTaskInfoRelationship, bool) {
	if o == nil || IsNil(o.CustomTaskInfo) {
		return nil, false
	}
	return o.CustomTaskInfo, true
}

// HasCustomTaskInfo returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasCustomTaskInfo() bool {
	if o != nil && !IsNil(o.CustomTaskInfo) {
		return true
	}

	return false
}

// SetCustomTaskInfo gets a reference to the given []IaasCustomTaskInfoRelationship and assigns it to the CustomTaskInfo field.
func (o *IaasUcsdInfo) SetCustomTaskInfo(v []IaasCustomTaskInfoRelationship) {
	o.CustomTaskInfo = v
}

// GetDeviceStatus returns the DeviceStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfo) GetDeviceStatus() []IaasDeviceStatusRelationship {
	if o == nil {
		var ret []IaasDeviceStatusRelationship
		return ret
	}
	return o.DeviceStatus
}

// GetDeviceStatusOk returns a tuple with the DeviceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfo) GetDeviceStatusOk() ([]IaasDeviceStatusRelationship, bool) {
	if o == nil || IsNil(o.DeviceStatus) {
		return nil, false
	}
	return o.DeviceStatus, true
}

// HasDeviceStatus returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasDeviceStatus() bool {
	if o != nil && !IsNil(o.DeviceStatus) {
		return true
	}

	return false
}

// SetDeviceStatus gets a reference to the given []IaasDeviceStatusRelationship and assigns it to the DeviceStatus field.
func (o *IaasUcsdInfo) SetDeviceStatus(v []IaasDeviceStatusRelationship) {
	o.DeviceStatus = v
}

// GetLicenseInfo returns the LicenseInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfo) GetLicenseInfo() IaasLicenseInfoRelationship {
	if o == nil || IsNil(o.LicenseInfo.Get()) {
		var ret IaasLicenseInfoRelationship
		return ret
	}
	return *o.LicenseInfo.Get()
}

// GetLicenseInfoOk returns a tuple with the LicenseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfo) GetLicenseInfoOk() (*IaasLicenseInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenseInfo.Get(), o.LicenseInfo.IsSet()
}

// HasLicenseInfo returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasLicenseInfo() bool {
	if o != nil && o.LicenseInfo.IsSet() {
		return true
	}

	return false
}

// SetLicenseInfo gets a reference to the given NullableIaasLicenseInfoRelationship and assigns it to the LicenseInfo field.
func (o *IaasUcsdInfo) SetLicenseInfo(v IaasLicenseInfoRelationship) {
	o.LicenseInfo.Set(&v)
}

// SetLicenseInfoNil sets the value for LicenseInfo to be an explicit nil
func (o *IaasUcsdInfo) SetLicenseInfoNil() {
	o.LicenseInfo.Set(nil)
}

// UnsetLicenseInfo ensures that no value is present for LicenseInfo, not even an explicit nil
func (o *IaasUcsdInfo) UnsetLicenseInfo() {
	o.LicenseInfo.Unset()
}

// GetMostRunTasks returns the MostRunTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfo) GetMostRunTasks() []IaasMostRunTasksRelationship {
	if o == nil {
		var ret []IaasMostRunTasksRelationship
		return ret
	}
	return o.MostRunTasks
}

// GetMostRunTasksOk returns a tuple with the MostRunTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfo) GetMostRunTasksOk() ([]IaasMostRunTasksRelationship, bool) {
	if o == nil || IsNil(o.MostRunTasks) {
		return nil, false
	}
	return o.MostRunTasks, true
}

// HasMostRunTasks returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasMostRunTasks() bool {
	if o != nil && !IsNil(o.MostRunTasks) {
		return true
	}

	return false
}

// SetMostRunTasks gets a reference to the given []IaasMostRunTasksRelationship and assigns it to the MostRunTasks field.
func (o *IaasUcsdInfo) SetMostRunTasks(v []IaasMostRunTasksRelationship) {
	o.MostRunTasks = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *IaasUcsdInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *IaasUcsdInfo) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *IaasUcsdInfo) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetSystemTaskInfo returns the SystemTaskInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfo) GetSystemTaskInfo() []IaasSystemTaskInfoRelationship {
	if o == nil {
		var ret []IaasSystemTaskInfoRelationship
		return ret
	}
	return o.SystemTaskInfo
}

// GetSystemTaskInfoOk returns a tuple with the SystemTaskInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfo) GetSystemTaskInfoOk() ([]IaasSystemTaskInfoRelationship, bool) {
	if o == nil || IsNil(o.SystemTaskInfo) {
		return nil, false
	}
	return o.SystemTaskInfo, true
}

// HasSystemTaskInfo returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasSystemTaskInfo() bool {
	if o != nil && !IsNil(o.SystemTaskInfo) {
		return true
	}

	return false
}

// SetSystemTaskInfo gets a reference to the given []IaasSystemTaskInfoRelationship and assigns it to the SystemTaskInfo field.
func (o *IaasUcsdInfo) SetSystemTaskInfo(v []IaasSystemTaskInfoRelationship) {
	o.SystemTaskInfo = v
}

// GetUcsdManagedInfra returns the UcsdManagedInfra field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasUcsdInfo) GetUcsdManagedInfra() IaasUcsdManagedInfraRelationship {
	if o == nil || IsNil(o.UcsdManagedInfra.Get()) {
		var ret IaasUcsdManagedInfraRelationship
		return ret
	}
	return *o.UcsdManagedInfra.Get()
}

// GetUcsdManagedInfraOk returns a tuple with the UcsdManagedInfra field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasUcsdInfo) GetUcsdManagedInfraOk() (*IaasUcsdManagedInfraRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.UcsdManagedInfra.Get(), o.UcsdManagedInfra.IsSet()
}

// HasUcsdManagedInfra returns a boolean if a field has been set.
func (o *IaasUcsdInfo) HasUcsdManagedInfra() bool {
	if o != nil && o.UcsdManagedInfra.IsSet() {
		return true
	}

	return false
}

// SetUcsdManagedInfra gets a reference to the given NullableIaasUcsdManagedInfraRelationship and assigns it to the UcsdManagedInfra field.
func (o *IaasUcsdInfo) SetUcsdManagedInfra(v IaasUcsdManagedInfraRelationship) {
	o.UcsdManagedInfra.Set(&v)
}

// SetUcsdManagedInfraNil sets the value for UcsdManagedInfra to be an explicit nil
func (o *IaasUcsdInfo) SetUcsdManagedInfraNil() {
	o.UcsdManagedInfra.Set(nil)
}

// UnsetUcsdManagedInfra ensures that no value is present for UcsdManagedInfra, not even an explicit nil
func (o *IaasUcsdInfo) UnsetUcsdManagedInfra() {
	o.UcsdManagedInfra.Unset()
}

func (o IaasUcsdInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IaasUcsdInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DeviceId) {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if !IsNil(o.Guid) {
		toSerialize["Guid"] = o.Guid
	}
	if !IsNil(o.HostName) {
		toSerialize["HostName"] = o.HostName
	}
	if !IsNil(o.Ip) {
		toSerialize["Ip"] = o.Ip
	}
	if !IsNil(o.LastBackup) {
		toSerialize["LastBackup"] = o.LastBackup
	}
	if !IsNil(o.NodeType) {
		toSerialize["NodeType"] = o.NodeType
	}
	if !IsNil(o.ProductName) {
		toSerialize["ProductName"] = o.ProductName
	}
	if !IsNil(o.ProductVendor) {
		toSerialize["ProductVendor"] = o.ProductVendor
	}
	if !IsNil(o.ProductVersion) {
		toSerialize["ProductVersion"] = o.ProductVersion
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.ConnectorPack != nil {
		toSerialize["ConnectorPack"] = o.ConnectorPack
	}
	if o.CustomTaskInfo != nil {
		toSerialize["CustomTaskInfo"] = o.CustomTaskInfo
	}
	if o.DeviceStatus != nil {
		toSerialize["DeviceStatus"] = o.DeviceStatus
	}
	if o.LicenseInfo.IsSet() {
		toSerialize["LicenseInfo"] = o.LicenseInfo.Get()
	}
	if o.MostRunTasks != nil {
		toSerialize["MostRunTasks"] = o.MostRunTasks
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.SystemTaskInfo != nil {
		toSerialize["SystemTaskInfo"] = o.SystemTaskInfo
	}
	if o.UcsdManagedInfra.IsSet() {
		toSerialize["UcsdManagedInfra"] = o.UcsdManagedInfra.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IaasUcsdInfo) UnmarshalJSON(data []byte) (err error) {
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
	type IaasUcsdInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Moid of the UCS Director device connector's asset.DeviceRegistration.
		DeviceId *string `json:"DeviceId,omitempty"`
		// Unique ID of UCS Director getting registerd with Intersight.
		Guid *string `json:"Guid,omitempty"`
		// The UCS Director hostname for management.
		HostName *string `json:"HostName,omitempty"`
		// The UCS Director IP address for management.
		Ip *string `json:"Ip,omitempty"`
		// Last successful backup created for this UCS Director appliance if backup is configured.
		LastBackup *time.Time `json:"LastBackup,omitempty"`
		// NodeType specifies if UCS Director is deployed in Stand-alone or Multi Node.
		NodeType *string `json:"NodeType,omitempty"`
		// The UCS Director product name.
		ProductName *string `json:"ProductName,omitempty"`
		// The UCS Director product vendor.
		ProductVendor *string `json:"ProductVendor,omitempty"`
		// The UCS Director product/platform version.
		ProductVersion *string `json:"ProductVersion,omitempty"`
		// The UCS Director status. Possible values are Active, Inactive, Unknown.
		Status *string `json:"Status,omitempty"`
		// An array of relationships to iaasConnectorPack resources.
		ConnectorPack []IaasConnectorPackRelationship `json:"ConnectorPack,omitempty"`
		// An array of relationships to iaasCustomTaskInfo resources.
		CustomTaskInfo []IaasCustomTaskInfoRelationship `json:"CustomTaskInfo,omitempty"`
		// An array of relationships to iaasDeviceStatus resources.
		DeviceStatus []IaasDeviceStatusRelationship      `json:"DeviceStatus,omitempty"`
		LicenseInfo  NullableIaasLicenseInfoRelationship `json:"LicenseInfo,omitempty"`
		// An array of relationships to iaasMostRunTasks resources.
		MostRunTasks     []IaasMostRunTasksRelationship              `json:"MostRunTasks,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to iaasSystemTaskInfo resources.
		SystemTaskInfo   []IaasSystemTaskInfoRelationship         `json:"SystemTaskInfo,omitempty"`
		UcsdManagedInfra NullableIaasUcsdManagedInfraRelationship `json:"UcsdManagedInfra,omitempty"`
	}

	varIaasUcsdInfoWithoutEmbeddedStruct := IaasUcsdInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIaasUcsdInfoWithoutEmbeddedStruct)
	if err == nil {
		varIaasUcsdInfo := _IaasUcsdInfo{}
		varIaasUcsdInfo.ClassId = varIaasUcsdInfoWithoutEmbeddedStruct.ClassId
		varIaasUcsdInfo.ObjectType = varIaasUcsdInfoWithoutEmbeddedStruct.ObjectType
		varIaasUcsdInfo.DeviceId = varIaasUcsdInfoWithoutEmbeddedStruct.DeviceId
		varIaasUcsdInfo.Guid = varIaasUcsdInfoWithoutEmbeddedStruct.Guid
		varIaasUcsdInfo.HostName = varIaasUcsdInfoWithoutEmbeddedStruct.HostName
		varIaasUcsdInfo.Ip = varIaasUcsdInfoWithoutEmbeddedStruct.Ip
		varIaasUcsdInfo.LastBackup = varIaasUcsdInfoWithoutEmbeddedStruct.LastBackup
		varIaasUcsdInfo.NodeType = varIaasUcsdInfoWithoutEmbeddedStruct.NodeType
		varIaasUcsdInfo.ProductName = varIaasUcsdInfoWithoutEmbeddedStruct.ProductName
		varIaasUcsdInfo.ProductVendor = varIaasUcsdInfoWithoutEmbeddedStruct.ProductVendor
		varIaasUcsdInfo.ProductVersion = varIaasUcsdInfoWithoutEmbeddedStruct.ProductVersion
		varIaasUcsdInfo.Status = varIaasUcsdInfoWithoutEmbeddedStruct.Status
		varIaasUcsdInfo.ConnectorPack = varIaasUcsdInfoWithoutEmbeddedStruct.ConnectorPack
		varIaasUcsdInfo.CustomTaskInfo = varIaasUcsdInfoWithoutEmbeddedStruct.CustomTaskInfo
		varIaasUcsdInfo.DeviceStatus = varIaasUcsdInfoWithoutEmbeddedStruct.DeviceStatus
		varIaasUcsdInfo.LicenseInfo = varIaasUcsdInfoWithoutEmbeddedStruct.LicenseInfo
		varIaasUcsdInfo.MostRunTasks = varIaasUcsdInfoWithoutEmbeddedStruct.MostRunTasks
		varIaasUcsdInfo.RegisteredDevice = varIaasUcsdInfoWithoutEmbeddedStruct.RegisteredDevice
		varIaasUcsdInfo.SystemTaskInfo = varIaasUcsdInfoWithoutEmbeddedStruct.SystemTaskInfo
		varIaasUcsdInfo.UcsdManagedInfra = varIaasUcsdInfoWithoutEmbeddedStruct.UcsdManagedInfra
		*o = IaasUcsdInfo(varIaasUcsdInfo)
	} else {
		return err
	}

	varIaasUcsdInfo := _IaasUcsdInfo{}

	err = json.Unmarshal(data, &varIaasUcsdInfo)
	if err == nil {
		o.MoBaseMo = varIaasUcsdInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "Guid")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "LastBackup")
		delete(additionalProperties, "NodeType")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "ProductVendor")
		delete(additionalProperties, "ProductVersion")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "ConnectorPack")
		delete(additionalProperties, "CustomTaskInfo")
		delete(additionalProperties, "DeviceStatus")
		delete(additionalProperties, "LicenseInfo")
		delete(additionalProperties, "MostRunTasks")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "SystemTaskInfo")
		delete(additionalProperties, "UcsdManagedInfra")

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

type NullableIaasUcsdInfo struct {
	value *IaasUcsdInfo
	isSet bool
}

func (v NullableIaasUcsdInfo) Get() *IaasUcsdInfo {
	return v.value
}

func (v *NullableIaasUcsdInfo) Set(val *IaasUcsdInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasUcsdInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasUcsdInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasUcsdInfo(val *IaasUcsdInfo) *NullableIaasUcsdInfo {
	return &NullableIaasUcsdInfo{value: val, isSet: true}
}

func (v NullableIaasUcsdInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasUcsdInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
