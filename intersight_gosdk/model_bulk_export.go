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

// checks if the BulkExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkExport{}

// BulkExport All export operations are captured as Export instances. Users shall use this Export mo to track the export operation progress.
type BulkExport struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action to be performed on the export operation. * `Start` - Starts the export operation. * `Cancel` - Cancels the export operation that is in progress.
	Action       *string  `json:"Action,omitempty"`
	ExcludePeers []string `json:"ExcludePeers,omitempty"`
	// Used to specify that none of the relationships should be exported.
	ExcludeRelations *bool `json:"ExcludeRelations,omitempty"`
	// Specifies whether tags must be exported and will be considered for all the items MOs.
	ExportTags      *bool            `json:"ExportTags,omitempty"`
	ExportedObjects []BulkSubRequest `json:"ExportedObjects,omitempty"`
	// Contains the list of import order.
	ImportOrder interface{} `json:"ImportOrder,omitempty"`
	// Indicates that exported references for objects which are organization owned should include the organization reference along with the other identity properties.
	IncludeOrgIdentity *bool     `json:"IncludeOrgIdentity,omitempty"`
	Items              []MoMoRef `json:"Items,omitempty"`
	// An identifier for the export instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_).
	Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9][a-zA-Z0-9_-]{1,92}$"`
	// The permission identifier which indicates the permission that current user has that will allow to start this export operation.
	PermissionId *string `json:"PermissionId,omitempty"`
	// Status of the export operation. * `` - The operation has not started. * `InProgress` - The operation is in progress. * `OrderInProgress` - The archive operation is in progress. * `Success` - The operation has succeeded. * `Failed` - The operation has failed. * `OperationTimedOut` - The operation has timed out. * `OperationCancelled` - The operation has been cancelled. * `CancelInProgress` - The operation is being cancelled.
	Status *string `json:"Status,omitempty"`
	// Status message associated with failures or progress indication.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// The user identifier which indicates the user that started this export operation.
	UserId *string `json:"UserId,omitempty"`
	// An array of relationships to bulkExportedItem resources.
	ExportedItems        []BulkExportedItemRelationship               `json:"ExportedItems,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkExport BulkExport

// NewBulkExport instantiates a new BulkExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkExport(classId string, objectType string) *BulkExport {
	this := BulkExport{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "Start"
	this.Action = &action
	var exportTags bool = true
	this.ExportTags = &exportTags
	return &this
}

// NewBulkExportWithDefaults instantiates a new BulkExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkExportWithDefaults() *BulkExport {
	this := BulkExport{}
	var classId string = "bulk.Export"
	this.ClassId = classId
	var objectType string = "bulk.Export"
	this.ObjectType = objectType
	var action string = "Start"
	this.Action = &action
	var exportTags bool = true
	this.ExportTags = &exportTags
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkExport) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkExport) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkExport) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "bulk.Export" of the ClassId field.
func (o *BulkExport) GetDefaultClassId() interface{} {
	return "bulk.Export"
}

// GetObjectType returns the ObjectType field value
func (o *BulkExport) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkExport) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkExport) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "bulk.Export" of the ObjectType field.
func (o *BulkExport) GetDefaultObjectType() interface{} {
	return "bulk.Export"
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BulkExport) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExport) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BulkExport) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BulkExport) SetAction(v string) {
	o.Action = &v
}

// GetExcludePeers returns the ExcludePeers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExport) GetExcludePeers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludePeers
}

// GetExcludePeersOk returns a tuple with the ExcludePeers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExport) GetExcludePeersOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludePeers) {
		return nil, false
	}
	return o.ExcludePeers, true
}

// HasExcludePeers returns a boolean if a field has been set.
func (o *BulkExport) HasExcludePeers() bool {
	if o != nil && !IsNil(o.ExcludePeers) {
		return true
	}

	return false
}

// SetExcludePeers gets a reference to the given []string and assigns it to the ExcludePeers field.
func (o *BulkExport) SetExcludePeers(v []string) {
	o.ExcludePeers = v
}

// GetExcludeRelations returns the ExcludeRelations field value if set, zero value otherwise.
func (o *BulkExport) GetExcludeRelations() bool {
	if o == nil || IsNil(o.ExcludeRelations) {
		var ret bool
		return ret
	}
	return *o.ExcludeRelations
}

// GetExcludeRelationsOk returns a tuple with the ExcludeRelations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExport) GetExcludeRelationsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeRelations) {
		return nil, false
	}
	return o.ExcludeRelations, true
}

// HasExcludeRelations returns a boolean if a field has been set.
func (o *BulkExport) HasExcludeRelations() bool {
	if o != nil && !IsNil(o.ExcludeRelations) {
		return true
	}

	return false
}

// SetExcludeRelations gets a reference to the given bool and assigns it to the ExcludeRelations field.
func (o *BulkExport) SetExcludeRelations(v bool) {
	o.ExcludeRelations = &v
}

// GetExportTags returns the ExportTags field value if set, zero value otherwise.
func (o *BulkExport) GetExportTags() bool {
	if o == nil || IsNil(o.ExportTags) {
		var ret bool
		return ret
	}
	return *o.ExportTags
}

// GetExportTagsOk returns a tuple with the ExportTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExport) GetExportTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExportTags) {
		return nil, false
	}
	return o.ExportTags, true
}

// HasExportTags returns a boolean if a field has been set.
func (o *BulkExport) HasExportTags() bool {
	if o != nil && !IsNil(o.ExportTags) {
		return true
	}

	return false
}

// SetExportTags gets a reference to the given bool and assigns it to the ExportTags field.
func (o *BulkExport) SetExportTags(v bool) {
	o.ExportTags = &v
}

// GetExportedObjects returns the ExportedObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExport) GetExportedObjects() []BulkSubRequest {
	if o == nil {
		var ret []BulkSubRequest
		return ret
	}
	return o.ExportedObjects
}

// GetExportedObjectsOk returns a tuple with the ExportedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExport) GetExportedObjectsOk() ([]BulkSubRequest, bool) {
	if o == nil || IsNil(o.ExportedObjects) {
		return nil, false
	}
	return o.ExportedObjects, true
}

// HasExportedObjects returns a boolean if a field has been set.
func (o *BulkExport) HasExportedObjects() bool {
	if o != nil && !IsNil(o.ExportedObjects) {
		return true
	}

	return false
}

// SetExportedObjects gets a reference to the given []BulkSubRequest and assigns it to the ExportedObjects field.
func (o *BulkExport) SetExportedObjects(v []BulkSubRequest) {
	o.ExportedObjects = v
}

// GetImportOrder returns the ImportOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExport) GetImportOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ImportOrder
}

// GetImportOrderOk returns a tuple with the ImportOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExport) GetImportOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ImportOrder) {
		return nil, false
	}
	return &o.ImportOrder, true
}

// HasImportOrder returns a boolean if a field has been set.
func (o *BulkExport) HasImportOrder() bool {
	if o != nil && !IsNil(o.ImportOrder) {
		return true
	}

	return false
}

// SetImportOrder gets a reference to the given interface{} and assigns it to the ImportOrder field.
func (o *BulkExport) SetImportOrder(v interface{}) {
	o.ImportOrder = v
}

// GetIncludeOrgIdentity returns the IncludeOrgIdentity field value if set, zero value otherwise.
func (o *BulkExport) GetIncludeOrgIdentity() bool {
	if o == nil || IsNil(o.IncludeOrgIdentity) {
		var ret bool
		return ret
	}
	return *o.IncludeOrgIdentity
}

// GetIncludeOrgIdentityOk returns a tuple with the IncludeOrgIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExport) GetIncludeOrgIdentityOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeOrgIdentity) {
		return nil, false
	}
	return o.IncludeOrgIdentity, true
}

// HasIncludeOrgIdentity returns a boolean if a field has been set.
func (o *BulkExport) HasIncludeOrgIdentity() bool {
	if o != nil && !IsNil(o.IncludeOrgIdentity) {
		return true
	}

	return false
}

// SetIncludeOrgIdentity gets a reference to the given bool and assigns it to the IncludeOrgIdentity field.
func (o *BulkExport) SetIncludeOrgIdentity(v bool) {
	o.IncludeOrgIdentity = &v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExport) GetItems() []MoMoRef {
	if o == nil {
		var ret []MoMoRef
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExport) GetItemsOk() ([]MoMoRef, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BulkExport) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MoMoRef and assigns it to the Items field.
func (o *BulkExport) SetItems(v []MoMoRef) {
	o.Items = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BulkExport) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExport) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BulkExport) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BulkExport) SetName(v string) {
	o.Name = &v
}

// GetPermissionId returns the PermissionId field value if set, zero value otherwise.
func (o *BulkExport) GetPermissionId() string {
	if o == nil || IsNil(o.PermissionId) {
		var ret string
		return ret
	}
	return *o.PermissionId
}

// GetPermissionIdOk returns a tuple with the PermissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExport) GetPermissionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionId) {
		return nil, false
	}
	return o.PermissionId, true
}

// HasPermissionId returns a boolean if a field has been set.
func (o *BulkExport) HasPermissionId() bool {
	if o != nil && !IsNil(o.PermissionId) {
		return true
	}

	return false
}

// SetPermissionId gets a reference to the given string and assigns it to the PermissionId field.
func (o *BulkExport) SetPermissionId(v string) {
	o.PermissionId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkExport) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExport) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkExport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BulkExport) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *BulkExport) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExport) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *BulkExport) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *BulkExport) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BulkExport) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkExport) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BulkExport) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BulkExport) SetUserId(v string) {
	o.UserId = &v
}

// GetExportedItems returns the ExportedItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExport) GetExportedItems() []BulkExportedItemRelationship {
	if o == nil {
		var ret []BulkExportedItemRelationship
		return ret
	}
	return o.ExportedItems
}

// GetExportedItemsOk returns a tuple with the ExportedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExport) GetExportedItemsOk() ([]BulkExportedItemRelationship, bool) {
	if o == nil || IsNil(o.ExportedItems) {
		return nil, false
	}
	return o.ExportedItems, true
}

// HasExportedItems returns a boolean if a field has been set.
func (o *BulkExport) HasExportedItems() bool {
	if o != nil && !IsNil(o.ExportedItems) {
		return true
	}

	return false
}

// SetExportedItems gets a reference to the given []BulkExportedItemRelationship and assigns it to the ExportedItems field.
func (o *BulkExport) SetExportedItems(v []BulkExportedItemRelationship) {
	o.ExportedItems = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkExport) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkExport) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkExport) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkExport) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *BulkExport) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *BulkExport) UnsetOrganization() {
	o.Organization.Unset()
}

func (o BulkExport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkExport) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Action) {
		toSerialize["Action"] = o.Action
	}
	if o.ExcludePeers != nil {
		toSerialize["ExcludePeers"] = o.ExcludePeers
	}
	if !IsNil(o.ExcludeRelations) {
		toSerialize["ExcludeRelations"] = o.ExcludeRelations
	}
	if !IsNil(o.ExportTags) {
		toSerialize["ExportTags"] = o.ExportTags
	}
	if o.ExportedObjects != nil {
		toSerialize["ExportedObjects"] = o.ExportedObjects
	}
	if o.ImportOrder != nil {
		toSerialize["ImportOrder"] = o.ImportOrder
	}
	if !IsNil(o.IncludeOrgIdentity) {
		toSerialize["IncludeOrgIdentity"] = o.IncludeOrgIdentity
	}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PermissionId) {
		toSerialize["PermissionId"] = o.PermissionId
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if !IsNil(o.UserId) {
		toSerialize["UserId"] = o.UserId
	}
	if o.ExportedItems != nil {
		toSerialize["ExportedItems"] = o.ExportedItems
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkExport) UnmarshalJSON(data []byte) (err error) {
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
	type BulkExportWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Action to be performed on the export operation. * `Start` - Starts the export operation. * `Cancel` - Cancels the export operation that is in progress.
		Action       *string  `json:"Action,omitempty"`
		ExcludePeers []string `json:"ExcludePeers,omitempty"`
		// Used to specify that none of the relationships should be exported.
		ExcludeRelations *bool `json:"ExcludeRelations,omitempty"`
		// Specifies whether tags must be exported and will be considered for all the items MOs.
		ExportTags      *bool            `json:"ExportTags,omitempty"`
		ExportedObjects []BulkSubRequest `json:"ExportedObjects,omitempty"`
		// Contains the list of import order.
		ImportOrder interface{} `json:"ImportOrder,omitempty"`
		// Indicates that exported references for objects which are organization owned should include the organization reference along with the other identity properties.
		IncludeOrgIdentity *bool     `json:"IncludeOrgIdentity,omitempty"`
		Items              []MoMoRef `json:"Items,omitempty"`
		// An identifier for the export instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_).
		Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9][a-zA-Z0-9_-]{1,92}$"`
		// The permission identifier which indicates the permission that current user has that will allow to start this export operation.
		PermissionId *string `json:"PermissionId,omitempty"`
		// Status of the export operation. * `` - The operation has not started. * `InProgress` - The operation is in progress. * `OrderInProgress` - The archive operation is in progress. * `Success` - The operation has succeeded. * `Failed` - The operation has failed. * `OperationTimedOut` - The operation has timed out. * `OperationCancelled` - The operation has been cancelled. * `CancelInProgress` - The operation is being cancelled.
		Status *string `json:"Status,omitempty"`
		// Status message associated with failures or progress indication.
		StatusMessage *string `json:"StatusMessage,omitempty"`
		// The user identifier which indicates the user that started this export operation.
		UserId *string `json:"UserId,omitempty"`
		// An array of relationships to bulkExportedItem resources.
		ExportedItems []BulkExportedItemRelationship               `json:"ExportedItems,omitempty"`
		Organization  NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varBulkExportWithoutEmbeddedStruct := BulkExportWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBulkExportWithoutEmbeddedStruct)
	if err == nil {
		varBulkExport := _BulkExport{}
		varBulkExport.ClassId = varBulkExportWithoutEmbeddedStruct.ClassId
		varBulkExport.ObjectType = varBulkExportWithoutEmbeddedStruct.ObjectType
		varBulkExport.Action = varBulkExportWithoutEmbeddedStruct.Action
		varBulkExport.ExcludePeers = varBulkExportWithoutEmbeddedStruct.ExcludePeers
		varBulkExport.ExcludeRelations = varBulkExportWithoutEmbeddedStruct.ExcludeRelations
		varBulkExport.ExportTags = varBulkExportWithoutEmbeddedStruct.ExportTags
		varBulkExport.ExportedObjects = varBulkExportWithoutEmbeddedStruct.ExportedObjects
		varBulkExport.ImportOrder = varBulkExportWithoutEmbeddedStruct.ImportOrder
		varBulkExport.IncludeOrgIdentity = varBulkExportWithoutEmbeddedStruct.IncludeOrgIdentity
		varBulkExport.Items = varBulkExportWithoutEmbeddedStruct.Items
		varBulkExport.Name = varBulkExportWithoutEmbeddedStruct.Name
		varBulkExport.PermissionId = varBulkExportWithoutEmbeddedStruct.PermissionId
		varBulkExport.Status = varBulkExportWithoutEmbeddedStruct.Status
		varBulkExport.StatusMessage = varBulkExportWithoutEmbeddedStruct.StatusMessage
		varBulkExport.UserId = varBulkExportWithoutEmbeddedStruct.UserId
		varBulkExport.ExportedItems = varBulkExportWithoutEmbeddedStruct.ExportedItems
		varBulkExport.Organization = varBulkExportWithoutEmbeddedStruct.Organization
		*o = BulkExport(varBulkExport)
	} else {
		return err
	}

	varBulkExport := _BulkExport{}

	err = json.Unmarshal(data, &varBulkExport)
	if err == nil {
		o.MoBaseMo = varBulkExport.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "ExcludePeers")
		delete(additionalProperties, "ExcludeRelations")
		delete(additionalProperties, "ExportTags")
		delete(additionalProperties, "ExportedObjects")
		delete(additionalProperties, "ImportOrder")
		delete(additionalProperties, "IncludeOrgIdentity")
		delete(additionalProperties, "Items")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PermissionId")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "UserId")
		delete(additionalProperties, "ExportedItems")
		delete(additionalProperties, "Organization")

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

type NullableBulkExport struct {
	value *BulkExport
	isSet bool
}

func (v NullableBulkExport) Get() *BulkExport {
	return v.value
}

func (v *NullableBulkExport) Set(val *BulkExport) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkExport) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkExport(val *BulkExport) *NullableBulkExport {
	return &NullableBulkExport{value: val, isSet: true}
}

func (v NullableBulkExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
