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

// checks if the AssetDeviceInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetDeviceInformation{}

// AssetDeviceInformation Type for saving the device information reported by Cisco Install Base. It is used in asset.SubscriptionDeviceInformation object to save the device information.
type AssetDeviceInformation struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Application name reported by Cisco Install Base.
	ApplicationName    *string                          `json:"ApplicationName,omitempty"`
	DeviceTransactions []AssetDeviceTransaction         `json:"DeviceTransactions,omitempty"`
	EndCustomer        NullableAssetCustomerInformation `json:"EndCustomer,omitempty"`
	// Instance number of the device. example \"917280220\".
	InstanceId *string `json:"InstanceId,omitempty"`
	// Item type flag. example ATO, Child, Standalone. ATO - refers to Cisco Block based major device. Child - refers to Child device part of an ATO block. Standalone - refers to a device that is managed and configured as an individual entity with limited capacity.
	ItemType *string `json:"ItemType,omitempty"`
	// Identifier for consumption based pricing. MLB refers to MultiLine Bundle.
	MlbOfferType *string `json:"MlbOfferType,omitempty"`
	// Identifier corresponding to MLB Product Name. MLB refers to MultiLine Bundle.
	MlbProductId *int64 `json:"MlbProductId,omitempty"`
	// Product Name for the device. It is used to determine if the server is to be billed as a UCS compute device or a storage cluster. MLB refers to MultiLine Bundle.
	MlbProductName *string `json:"MlbProductName,omitempty"`
	// Unique identifier of old Cisco device. It is the device ID of old Cisco device, which got replaced by the new device.
	OldDeviceId *string `json:"OldDeviceId,omitempty"`
	// Description of status of old Cisco device, which got replaced by the new device.
	OldDeviceStatusDescription *string `json:"OldDeviceStatusDescription,omitempty"`
	// Status ID of old Cisco device, which got replaced by the new device. * `0` - A default value to catch cases where device status is not correctly detected. * `10000` - Device is installed successfully. * `1010041` - Device is currently in Return Material Authorization process. * `10005` - Device is replaced successfully with another device. * `10007` - Device is returned succcessfuly. * `10009` - Device is terminated at customer's end.
	OldDeviceStatusId *int32 `json:"OldDeviceStatusId,omitempty"`
	// Instance number of the old device, which got replaced by the new device.
	OldInstanceId *string `json:"OldInstanceId,omitempty"`
	// Product Family is the field used to identify the hypervisor type. example \"ESXi\".
	ProductFamily *string `json:"ProductFamily,omitempty"`
	// Product type helps to determine if device has to be billed using consumption metering. example \"SERVER\".
	ProductType *string `json:"ProductType,omitempty"`
	// Unit of Measure is flag used to identify the type of metric being pushed. example - \"STORAGE\" for hardware metrics , \"VM\" - for hypervisor metrics. * `None` - A default value to catch cases where unit of measure is not correctly detected. * `STORAGE` - The metric type of the device is a storage metric. * `NODE` - The metric type of the device is a hardware metric. * `VM` - The metric type of the device is a hypervisor metric.
	UnitOfMeasure        *string `json:"UnitOfMeasure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDeviceInformation AssetDeviceInformation

// NewAssetDeviceInformation instantiates a new AssetDeviceInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceInformation(classId string, objectType string) *AssetDeviceInformation {
	this := AssetDeviceInformation{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeviceInformationWithDefaults instantiates a new AssetDeviceInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceInformationWithDefaults() *AssetDeviceInformation {
	this := AssetDeviceInformation{}
	var classId string = "asset.DeviceInformation"
	this.ClassId = classId
	var objectType string = "asset.DeviceInformation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceInformation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceInformation) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.DeviceInformation" of the ClassId field.
func (o *AssetDeviceInformation) GetDefaultClassId() interface{} {
	return "asset.DeviceInformation"
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceInformation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceInformation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.DeviceInformation" of the ObjectType field.
func (o *AssetDeviceInformation) GetDefaultObjectType() interface{} {
	return "asset.DeviceInformation"
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetApplicationName() string {
	if o == nil || IsNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasApplicationName() bool {
	if o != nil && !IsNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *AssetDeviceInformation) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetDeviceTransactions returns the DeviceTransactions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceInformation) GetDeviceTransactions() []AssetDeviceTransaction {
	if o == nil {
		var ret []AssetDeviceTransaction
		return ret
	}
	return o.DeviceTransactions
}

// GetDeviceTransactionsOk returns a tuple with the DeviceTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceInformation) GetDeviceTransactionsOk() ([]AssetDeviceTransaction, bool) {
	if o == nil || IsNil(o.DeviceTransactions) {
		return nil, false
	}
	return o.DeviceTransactions, true
}

// HasDeviceTransactions returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasDeviceTransactions() bool {
	if o != nil && !IsNil(o.DeviceTransactions) {
		return true
	}

	return false
}

// SetDeviceTransactions gets a reference to the given []AssetDeviceTransaction and assigns it to the DeviceTransactions field.
func (o *AssetDeviceInformation) SetDeviceTransactions(v []AssetDeviceTransaction) {
	o.DeviceTransactions = v
}

// GetEndCustomer returns the EndCustomer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceInformation) GetEndCustomer() AssetCustomerInformation {
	if o == nil || IsNil(o.EndCustomer.Get()) {
		var ret AssetCustomerInformation
		return ret
	}
	return *o.EndCustomer.Get()
}

// GetEndCustomerOk returns a tuple with the EndCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceInformation) GetEndCustomerOk() (*AssetCustomerInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndCustomer.Get(), o.EndCustomer.IsSet()
}

// HasEndCustomer returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasEndCustomer() bool {
	if o != nil && o.EndCustomer.IsSet() {
		return true
	}

	return false
}

// SetEndCustomer gets a reference to the given NullableAssetCustomerInformation and assigns it to the EndCustomer field.
func (o *AssetDeviceInformation) SetEndCustomer(v AssetCustomerInformation) {
	o.EndCustomer.Set(&v)
}

// SetEndCustomerNil sets the value for EndCustomer to be an explicit nil
func (o *AssetDeviceInformation) SetEndCustomerNil() {
	o.EndCustomer.Set(nil)
}

// UnsetEndCustomer ensures that no value is present for EndCustomer, not even an explicit nil
func (o *AssetDeviceInformation) UnsetEndCustomer() {
	o.EndCustomer.Unset()
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *AssetDeviceInformation) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetItemType() string {
	if o == nil || IsNil(o.ItemType) {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetItemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ItemType) {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasItemType() bool {
	if o != nil && !IsNil(o.ItemType) {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *AssetDeviceInformation) SetItemType(v string) {
	o.ItemType = &v
}

// GetMlbOfferType returns the MlbOfferType field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetMlbOfferType() string {
	if o == nil || IsNil(o.MlbOfferType) {
		var ret string
		return ret
	}
	return *o.MlbOfferType
}

// GetMlbOfferTypeOk returns a tuple with the MlbOfferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetMlbOfferTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MlbOfferType) {
		return nil, false
	}
	return o.MlbOfferType, true
}

// HasMlbOfferType returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasMlbOfferType() bool {
	if o != nil && !IsNil(o.MlbOfferType) {
		return true
	}

	return false
}

// SetMlbOfferType gets a reference to the given string and assigns it to the MlbOfferType field.
func (o *AssetDeviceInformation) SetMlbOfferType(v string) {
	o.MlbOfferType = &v
}

// GetMlbProductId returns the MlbProductId field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetMlbProductId() int64 {
	if o == nil || IsNil(o.MlbProductId) {
		var ret int64
		return ret
	}
	return *o.MlbProductId
}

// GetMlbProductIdOk returns a tuple with the MlbProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetMlbProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MlbProductId) {
		return nil, false
	}
	return o.MlbProductId, true
}

// HasMlbProductId returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasMlbProductId() bool {
	if o != nil && !IsNil(o.MlbProductId) {
		return true
	}

	return false
}

// SetMlbProductId gets a reference to the given int64 and assigns it to the MlbProductId field.
func (o *AssetDeviceInformation) SetMlbProductId(v int64) {
	o.MlbProductId = &v
}

// GetMlbProductName returns the MlbProductName field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetMlbProductName() string {
	if o == nil || IsNil(o.MlbProductName) {
		var ret string
		return ret
	}
	return *o.MlbProductName
}

// GetMlbProductNameOk returns a tuple with the MlbProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetMlbProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.MlbProductName) {
		return nil, false
	}
	return o.MlbProductName, true
}

// HasMlbProductName returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasMlbProductName() bool {
	if o != nil && !IsNil(o.MlbProductName) {
		return true
	}

	return false
}

// SetMlbProductName gets a reference to the given string and assigns it to the MlbProductName field.
func (o *AssetDeviceInformation) SetMlbProductName(v string) {
	o.MlbProductName = &v
}

// GetOldDeviceId returns the OldDeviceId field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetOldDeviceId() string {
	if o == nil || IsNil(o.OldDeviceId) {
		var ret string
		return ret
	}
	return *o.OldDeviceId
}

// GetOldDeviceIdOk returns a tuple with the OldDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetOldDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.OldDeviceId) {
		return nil, false
	}
	return o.OldDeviceId, true
}

// HasOldDeviceId returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasOldDeviceId() bool {
	if o != nil && !IsNil(o.OldDeviceId) {
		return true
	}

	return false
}

// SetOldDeviceId gets a reference to the given string and assigns it to the OldDeviceId field.
func (o *AssetDeviceInformation) SetOldDeviceId(v string) {
	o.OldDeviceId = &v
}

// GetOldDeviceStatusDescription returns the OldDeviceStatusDescription field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetOldDeviceStatusDescription() string {
	if o == nil || IsNil(o.OldDeviceStatusDescription) {
		var ret string
		return ret
	}
	return *o.OldDeviceStatusDescription
}

// GetOldDeviceStatusDescriptionOk returns a tuple with the OldDeviceStatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetOldDeviceStatusDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.OldDeviceStatusDescription) {
		return nil, false
	}
	return o.OldDeviceStatusDescription, true
}

// HasOldDeviceStatusDescription returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasOldDeviceStatusDescription() bool {
	if o != nil && !IsNil(o.OldDeviceStatusDescription) {
		return true
	}

	return false
}

// SetOldDeviceStatusDescription gets a reference to the given string and assigns it to the OldDeviceStatusDescription field.
func (o *AssetDeviceInformation) SetOldDeviceStatusDescription(v string) {
	o.OldDeviceStatusDescription = &v
}

// GetOldDeviceStatusId returns the OldDeviceStatusId field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetOldDeviceStatusId() int32 {
	if o == nil || IsNil(o.OldDeviceStatusId) {
		var ret int32
		return ret
	}
	return *o.OldDeviceStatusId
}

// GetOldDeviceStatusIdOk returns a tuple with the OldDeviceStatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetOldDeviceStatusIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OldDeviceStatusId) {
		return nil, false
	}
	return o.OldDeviceStatusId, true
}

// HasOldDeviceStatusId returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasOldDeviceStatusId() bool {
	if o != nil && !IsNil(o.OldDeviceStatusId) {
		return true
	}

	return false
}

// SetOldDeviceStatusId gets a reference to the given int32 and assigns it to the OldDeviceStatusId field.
func (o *AssetDeviceInformation) SetOldDeviceStatusId(v int32) {
	o.OldDeviceStatusId = &v
}

// GetOldInstanceId returns the OldInstanceId field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetOldInstanceId() string {
	if o == nil || IsNil(o.OldInstanceId) {
		var ret string
		return ret
	}
	return *o.OldInstanceId
}

// GetOldInstanceIdOk returns a tuple with the OldInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetOldInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.OldInstanceId) {
		return nil, false
	}
	return o.OldInstanceId, true
}

// HasOldInstanceId returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasOldInstanceId() bool {
	if o != nil && !IsNil(o.OldInstanceId) {
		return true
	}

	return false
}

// SetOldInstanceId gets a reference to the given string and assigns it to the OldInstanceId field.
func (o *AssetDeviceInformation) SetOldInstanceId(v string) {
	o.OldInstanceId = &v
}

// GetProductFamily returns the ProductFamily field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetProductFamily() string {
	if o == nil || IsNil(o.ProductFamily) {
		var ret string
		return ret
	}
	return *o.ProductFamily
}

// GetProductFamilyOk returns a tuple with the ProductFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetProductFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.ProductFamily) {
		return nil, false
	}
	return o.ProductFamily, true
}

// HasProductFamily returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasProductFamily() bool {
	if o != nil && !IsNil(o.ProductFamily) {
		return true
	}

	return false
}

// SetProductFamily gets a reference to the given string and assigns it to the ProductFamily field.
func (o *AssetDeviceInformation) SetProductFamily(v string) {
	o.ProductFamily = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *AssetDeviceInformation) SetProductType(v string) {
	o.ProductType = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *AssetDeviceInformation) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceInformation) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *AssetDeviceInformation) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *AssetDeviceInformation) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

func (o AssetDeviceInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetDeviceInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ApplicationName) {
		toSerialize["ApplicationName"] = o.ApplicationName
	}
	if o.DeviceTransactions != nil {
		toSerialize["DeviceTransactions"] = o.DeviceTransactions
	}
	if o.EndCustomer.IsSet() {
		toSerialize["EndCustomer"] = o.EndCustomer.Get()
	}
	if !IsNil(o.InstanceId) {
		toSerialize["InstanceId"] = o.InstanceId
	}
	if !IsNil(o.ItemType) {
		toSerialize["ItemType"] = o.ItemType
	}
	if !IsNil(o.MlbOfferType) {
		toSerialize["MlbOfferType"] = o.MlbOfferType
	}
	if !IsNil(o.MlbProductId) {
		toSerialize["MlbProductId"] = o.MlbProductId
	}
	if !IsNil(o.MlbProductName) {
		toSerialize["MlbProductName"] = o.MlbProductName
	}
	if !IsNil(o.OldDeviceId) {
		toSerialize["OldDeviceId"] = o.OldDeviceId
	}
	if !IsNil(o.OldDeviceStatusDescription) {
		toSerialize["OldDeviceStatusDescription"] = o.OldDeviceStatusDescription
	}
	if !IsNil(o.OldDeviceStatusId) {
		toSerialize["OldDeviceStatusId"] = o.OldDeviceStatusId
	}
	if !IsNil(o.OldInstanceId) {
		toSerialize["OldInstanceId"] = o.OldInstanceId
	}
	if !IsNil(o.ProductFamily) {
		toSerialize["ProductFamily"] = o.ProductFamily
	}
	if !IsNil(o.ProductType) {
		toSerialize["ProductType"] = o.ProductType
	}
	if !IsNil(o.UnitOfMeasure) {
		toSerialize["UnitOfMeasure"] = o.UnitOfMeasure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetDeviceInformation) UnmarshalJSON(data []byte) (err error) {
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
	type AssetDeviceInformationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Application name reported by Cisco Install Base.
		ApplicationName    *string                          `json:"ApplicationName,omitempty"`
		DeviceTransactions []AssetDeviceTransaction         `json:"DeviceTransactions,omitempty"`
		EndCustomer        NullableAssetCustomerInformation `json:"EndCustomer,omitempty"`
		// Instance number of the device. example \"917280220\".
		InstanceId *string `json:"InstanceId,omitempty"`
		// Item type flag. example ATO, Child, Standalone. ATO - refers to Cisco Block based major device. Child - refers to Child device part of an ATO block. Standalone - refers to a device that is managed and configured as an individual entity with limited capacity.
		ItemType *string `json:"ItemType,omitempty"`
		// Identifier for consumption based pricing. MLB refers to MultiLine Bundle.
		MlbOfferType *string `json:"MlbOfferType,omitempty"`
		// Identifier corresponding to MLB Product Name. MLB refers to MultiLine Bundle.
		MlbProductId *int64 `json:"MlbProductId,omitempty"`
		// Product Name for the device. It is used to determine if the server is to be billed as a UCS compute device or a storage cluster. MLB refers to MultiLine Bundle.
		MlbProductName *string `json:"MlbProductName,omitempty"`
		// Unique identifier of old Cisco device. It is the device ID of old Cisco device, which got replaced by the new device.
		OldDeviceId *string `json:"OldDeviceId,omitempty"`
		// Description of status of old Cisco device, which got replaced by the new device.
		OldDeviceStatusDescription *string `json:"OldDeviceStatusDescription,omitempty"`
		// Status ID of old Cisco device, which got replaced by the new device. * `0` - A default value to catch cases where device status is not correctly detected. * `10000` - Device is installed successfully. * `1010041` - Device is currently in Return Material Authorization process. * `10005` - Device is replaced successfully with another device. * `10007` - Device is returned succcessfuly. * `10009` - Device is terminated at customer's end.
		OldDeviceStatusId *int32 `json:"OldDeviceStatusId,omitempty"`
		// Instance number of the old device, which got replaced by the new device.
		OldInstanceId *string `json:"OldInstanceId,omitempty"`
		// Product Family is the field used to identify the hypervisor type. example \"ESXi\".
		ProductFamily *string `json:"ProductFamily,omitempty"`
		// Product type helps to determine if device has to be billed using consumption metering. example \"SERVER\".
		ProductType *string `json:"ProductType,omitempty"`
		// Unit of Measure is flag used to identify the type of metric being pushed. example - \"STORAGE\" for hardware metrics , \"VM\" - for hypervisor metrics. * `None` - A default value to catch cases where unit of measure is not correctly detected. * `STORAGE` - The metric type of the device is a storage metric. * `NODE` - The metric type of the device is a hardware metric. * `VM` - The metric type of the device is a hypervisor metric.
		UnitOfMeasure *string `json:"UnitOfMeasure,omitempty"`
	}

	varAssetDeviceInformationWithoutEmbeddedStruct := AssetDeviceInformationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetDeviceInformationWithoutEmbeddedStruct)
	if err == nil {
		varAssetDeviceInformation := _AssetDeviceInformation{}
		varAssetDeviceInformation.ClassId = varAssetDeviceInformationWithoutEmbeddedStruct.ClassId
		varAssetDeviceInformation.ObjectType = varAssetDeviceInformationWithoutEmbeddedStruct.ObjectType
		varAssetDeviceInformation.ApplicationName = varAssetDeviceInformationWithoutEmbeddedStruct.ApplicationName
		varAssetDeviceInformation.DeviceTransactions = varAssetDeviceInformationWithoutEmbeddedStruct.DeviceTransactions
		varAssetDeviceInformation.EndCustomer = varAssetDeviceInformationWithoutEmbeddedStruct.EndCustomer
		varAssetDeviceInformation.InstanceId = varAssetDeviceInformationWithoutEmbeddedStruct.InstanceId
		varAssetDeviceInformation.ItemType = varAssetDeviceInformationWithoutEmbeddedStruct.ItemType
		varAssetDeviceInformation.MlbOfferType = varAssetDeviceInformationWithoutEmbeddedStruct.MlbOfferType
		varAssetDeviceInformation.MlbProductId = varAssetDeviceInformationWithoutEmbeddedStruct.MlbProductId
		varAssetDeviceInformation.MlbProductName = varAssetDeviceInformationWithoutEmbeddedStruct.MlbProductName
		varAssetDeviceInformation.OldDeviceId = varAssetDeviceInformationWithoutEmbeddedStruct.OldDeviceId
		varAssetDeviceInformation.OldDeviceStatusDescription = varAssetDeviceInformationWithoutEmbeddedStruct.OldDeviceStatusDescription
		varAssetDeviceInformation.OldDeviceStatusId = varAssetDeviceInformationWithoutEmbeddedStruct.OldDeviceStatusId
		varAssetDeviceInformation.OldInstanceId = varAssetDeviceInformationWithoutEmbeddedStruct.OldInstanceId
		varAssetDeviceInformation.ProductFamily = varAssetDeviceInformationWithoutEmbeddedStruct.ProductFamily
		varAssetDeviceInformation.ProductType = varAssetDeviceInformationWithoutEmbeddedStruct.ProductType
		varAssetDeviceInformation.UnitOfMeasure = varAssetDeviceInformationWithoutEmbeddedStruct.UnitOfMeasure
		*o = AssetDeviceInformation(varAssetDeviceInformation)
	} else {
		return err
	}

	varAssetDeviceInformation := _AssetDeviceInformation{}

	err = json.Unmarshal(data, &varAssetDeviceInformation)
	if err == nil {
		o.MoBaseComplexType = varAssetDeviceInformation.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApplicationName")
		delete(additionalProperties, "DeviceTransactions")
		delete(additionalProperties, "EndCustomer")
		delete(additionalProperties, "InstanceId")
		delete(additionalProperties, "ItemType")
		delete(additionalProperties, "MlbOfferType")
		delete(additionalProperties, "MlbProductId")
		delete(additionalProperties, "MlbProductName")
		delete(additionalProperties, "OldDeviceId")
		delete(additionalProperties, "OldDeviceStatusDescription")
		delete(additionalProperties, "OldDeviceStatusId")
		delete(additionalProperties, "OldInstanceId")
		delete(additionalProperties, "ProductFamily")
		delete(additionalProperties, "ProductType")
		delete(additionalProperties, "UnitOfMeasure")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableAssetDeviceInformation struct {
	value *AssetDeviceInformation
	isSet bool
}

func (v NullableAssetDeviceInformation) Get() *AssetDeviceInformation {
	return v.value
}

func (v *NullableAssetDeviceInformation) Set(val *AssetDeviceInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceInformation(val *AssetDeviceInformation) *NullableAssetDeviceInformation {
	return &NullableAssetDeviceInformation{value: val, isSet: true}
}

func (v NullableAssetDeviceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
