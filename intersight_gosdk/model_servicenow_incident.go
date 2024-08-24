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

// checks if the ServicenowIncident type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicenowIncident{}

// ServicenowIncident A Incident on ServiceNow.
type ServicenowIncident struct {
	ServicenowInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The approver property of Incident.
	Approval *string `json:"Approval,omitempty"`
	// Category property for Incident.
	Category *string `json:"Category,omitempty"`
	// Comments property on Incident.
	Comments *string `json:"Comments,omitempty"`
	// Creator property of Incident.
	CreatedBy *string `json:"CreatedBy,omitempty"`
	// Incident create date property.
	CreatedOn *string `json:"CreatedOn,omitempty"`
	// Description for Incident.
	Description *string `json:"Description,omitempty"`
	// Due date property for Incident.
	DueDate *string `json:"DueDate,omitempty"`
	// Expected start date for Incident.
	ExpectedStart *string `json:"ExpectedStart,omitempty"`
	// Impact property for Incident.
	Impact *string `json:"Impact,omitempty"`
	// State property of the Incident.
	IncidentState *string `json:"IncidentState,omitempty"`
	// Assigned to value for Incident.
	OpenedBy *string `json:"OpenedBy,omitempty"`
	// Priority property for Incident.
	Priority *string `json:"Priority,omitempty"`
	// The risk property for Incident.
	Risk *string `json:"Risk,omitempty"`
	// Severity property of the Incident.
	Severity *string `json:"Severity,omitempty"`
	// Short Description for Incident.
	ShortDescription *string `json:"ShortDescription,omitempty"`
	// System Id property for Incident.
	SysId *string `json:"SysId,omitempty"`
	// Task Effective Number for Incident.
	TaskEffectiveNumber *string `json:"TaskEffectiveNumber,omitempty"`
	// Last update by on Incident.
	UpdatedBy *string `json:"UpdatedBy,omitempty"`
	// Urgency property of the Incident.
	Urgency *string `json:"Urgency,omitempty"`
	// Work end date for Incident.
	WorkEnd *string `json:"WorkEnd,omitempty"`
	// Work start date for Incident.
	WorkStart            *string `json:"WorkStart,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicenowIncident ServicenowIncident

// NewServicenowIncident instantiates a new ServicenowIncident object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicenowIncident(classId string, objectType string) *ServicenowIncident {
	this := ServicenowIncident{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServicenowIncidentWithDefaults instantiates a new ServicenowIncident object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicenowIncidentWithDefaults() *ServicenowIncident {
	this := ServicenowIncident{}
	var classId string = "servicenow.Incident"
	this.ClassId = classId
	var objectType string = "servicenow.Incident"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServicenowIncident) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServicenowIncident) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "servicenow.Incident" of the ClassId field.
func (o *ServicenowIncident) GetDefaultClassId() interface{} {
	return "servicenow.Incident"
}

// GetObjectType returns the ObjectType field value
func (o *ServicenowIncident) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServicenowIncident) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "servicenow.Incident" of the ObjectType field.
func (o *ServicenowIncident) GetDefaultObjectType() interface{} {
	return "servicenow.Incident"
}

// GetApproval returns the Approval field value if set, zero value otherwise.
func (o *ServicenowIncident) GetApproval() string {
	if o == nil || IsNil(o.Approval) {
		var ret string
		return ret
	}
	return *o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetApprovalOk() (*string, bool) {
	if o == nil || IsNil(o.Approval) {
		return nil, false
	}
	return o.Approval, true
}

// HasApproval returns a boolean if a field has been set.
func (o *ServicenowIncident) HasApproval() bool {
	if o != nil && !IsNil(o.Approval) {
		return true
	}

	return false
}

// SetApproval gets a reference to the given string and assigns it to the Approval field.
func (o *ServicenowIncident) SetApproval(v string) {
	o.Approval = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ServicenowIncident) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ServicenowIncident) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ServicenowIncident) SetCategory(v string) {
	o.Category = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ServicenowIncident) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ServicenowIncident) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ServicenowIncident) SetComments(v string) {
	o.Comments = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ServicenowIncident) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ServicenowIncident) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ServicenowIncident) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *ServicenowIncident) GetCreatedOn() string {
	if o == nil || IsNil(o.CreatedOn) {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetCreatedOnOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *ServicenowIncident) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *ServicenowIncident) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServicenowIncident) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServicenowIncident) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServicenowIncident) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *ServicenowIncident) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *ServicenowIncident) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *ServicenowIncident) SetDueDate(v string) {
	o.DueDate = &v
}

// GetExpectedStart returns the ExpectedStart field value if set, zero value otherwise.
func (o *ServicenowIncident) GetExpectedStart() string {
	if o == nil || IsNil(o.ExpectedStart) {
		var ret string
		return ret
	}
	return *o.ExpectedStart
}

// GetExpectedStartOk returns a tuple with the ExpectedStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetExpectedStartOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedStart) {
		return nil, false
	}
	return o.ExpectedStart, true
}

// HasExpectedStart returns a boolean if a field has been set.
func (o *ServicenowIncident) HasExpectedStart() bool {
	if o != nil && !IsNil(o.ExpectedStart) {
		return true
	}

	return false
}

// SetExpectedStart gets a reference to the given string and assigns it to the ExpectedStart field.
func (o *ServicenowIncident) SetExpectedStart(v string) {
	o.ExpectedStart = &v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *ServicenowIncident) GetImpact() string {
	if o == nil || IsNil(o.Impact) {
		var ret string
		return ret
	}
	return *o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetImpactOk() (*string, bool) {
	if o == nil || IsNil(o.Impact) {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *ServicenowIncident) HasImpact() bool {
	if o != nil && !IsNil(o.Impact) {
		return true
	}

	return false
}

// SetImpact gets a reference to the given string and assigns it to the Impact field.
func (o *ServicenowIncident) SetImpact(v string) {
	o.Impact = &v
}

// GetIncidentState returns the IncidentState field value if set, zero value otherwise.
func (o *ServicenowIncident) GetIncidentState() string {
	if o == nil || IsNil(o.IncidentState) {
		var ret string
		return ret
	}
	return *o.IncidentState
}

// GetIncidentStateOk returns a tuple with the IncidentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetIncidentStateOk() (*string, bool) {
	if o == nil || IsNil(o.IncidentState) {
		return nil, false
	}
	return o.IncidentState, true
}

// HasIncidentState returns a boolean if a field has been set.
func (o *ServicenowIncident) HasIncidentState() bool {
	if o != nil && !IsNil(o.IncidentState) {
		return true
	}

	return false
}

// SetIncidentState gets a reference to the given string and assigns it to the IncidentState field.
func (o *ServicenowIncident) SetIncidentState(v string) {
	o.IncidentState = &v
}

// GetOpenedBy returns the OpenedBy field value if set, zero value otherwise.
func (o *ServicenowIncident) GetOpenedBy() string {
	if o == nil || IsNil(o.OpenedBy) {
		var ret string
		return ret
	}
	return *o.OpenedBy
}

// GetOpenedByOk returns a tuple with the OpenedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetOpenedByOk() (*string, bool) {
	if o == nil || IsNil(o.OpenedBy) {
		return nil, false
	}
	return o.OpenedBy, true
}

// HasOpenedBy returns a boolean if a field has been set.
func (o *ServicenowIncident) HasOpenedBy() bool {
	if o != nil && !IsNil(o.OpenedBy) {
		return true
	}

	return false
}

// SetOpenedBy gets a reference to the given string and assigns it to the OpenedBy field.
func (o *ServicenowIncident) SetOpenedBy(v string) {
	o.OpenedBy = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ServicenowIncident) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ServicenowIncident) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *ServicenowIncident) SetPriority(v string) {
	o.Priority = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *ServicenowIncident) GetRisk() string {
	if o == nil || IsNil(o.Risk) {
		var ret string
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetRiskOk() (*string, bool) {
	if o == nil || IsNil(o.Risk) {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *ServicenowIncident) HasRisk() bool {
	if o != nil && !IsNil(o.Risk) {
		return true
	}

	return false
}

// SetRisk gets a reference to the given string and assigns it to the Risk field.
func (o *ServicenowIncident) SetRisk(v string) {
	o.Risk = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ServicenowIncident) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ServicenowIncident) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ServicenowIncident) SetSeverity(v string) {
	o.Severity = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *ServicenowIncident) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *ServicenowIncident) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *ServicenowIncident) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetSysId returns the SysId field value if set, zero value otherwise.
func (o *ServicenowIncident) GetSysId() string {
	if o == nil || IsNil(o.SysId) {
		var ret string
		return ret
	}
	return *o.SysId
}

// GetSysIdOk returns a tuple with the SysId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetSysIdOk() (*string, bool) {
	if o == nil || IsNil(o.SysId) {
		return nil, false
	}
	return o.SysId, true
}

// HasSysId returns a boolean if a field has been set.
func (o *ServicenowIncident) HasSysId() bool {
	if o != nil && !IsNil(o.SysId) {
		return true
	}

	return false
}

// SetSysId gets a reference to the given string and assigns it to the SysId field.
func (o *ServicenowIncident) SetSysId(v string) {
	o.SysId = &v
}

// GetTaskEffectiveNumber returns the TaskEffectiveNumber field value if set, zero value otherwise.
func (o *ServicenowIncident) GetTaskEffectiveNumber() string {
	if o == nil || IsNil(o.TaskEffectiveNumber) {
		var ret string
		return ret
	}
	return *o.TaskEffectiveNumber
}

// GetTaskEffectiveNumberOk returns a tuple with the TaskEffectiveNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetTaskEffectiveNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TaskEffectiveNumber) {
		return nil, false
	}
	return o.TaskEffectiveNumber, true
}

// HasTaskEffectiveNumber returns a boolean if a field has been set.
func (o *ServicenowIncident) HasTaskEffectiveNumber() bool {
	if o != nil && !IsNil(o.TaskEffectiveNumber) {
		return true
	}

	return false
}

// SetTaskEffectiveNumber gets a reference to the given string and assigns it to the TaskEffectiveNumber field.
func (o *ServicenowIncident) SetTaskEffectiveNumber(v string) {
	o.TaskEffectiveNumber = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ServicenowIncident) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ServicenowIncident) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *ServicenowIncident) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUrgency returns the Urgency field value if set, zero value otherwise.
func (o *ServicenowIncident) GetUrgency() string {
	if o == nil || IsNil(o.Urgency) {
		var ret string
		return ret
	}
	return *o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetUrgencyOk() (*string, bool) {
	if o == nil || IsNil(o.Urgency) {
		return nil, false
	}
	return o.Urgency, true
}

// HasUrgency returns a boolean if a field has been set.
func (o *ServicenowIncident) HasUrgency() bool {
	if o != nil && !IsNil(o.Urgency) {
		return true
	}

	return false
}

// SetUrgency gets a reference to the given string and assigns it to the Urgency field.
func (o *ServicenowIncident) SetUrgency(v string) {
	o.Urgency = &v
}

// GetWorkEnd returns the WorkEnd field value if set, zero value otherwise.
func (o *ServicenowIncident) GetWorkEnd() string {
	if o == nil || IsNil(o.WorkEnd) {
		var ret string
		return ret
	}
	return *o.WorkEnd
}

// GetWorkEndOk returns a tuple with the WorkEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetWorkEndOk() (*string, bool) {
	if o == nil || IsNil(o.WorkEnd) {
		return nil, false
	}
	return o.WorkEnd, true
}

// HasWorkEnd returns a boolean if a field has been set.
func (o *ServicenowIncident) HasWorkEnd() bool {
	if o != nil && !IsNil(o.WorkEnd) {
		return true
	}

	return false
}

// SetWorkEnd gets a reference to the given string and assigns it to the WorkEnd field.
func (o *ServicenowIncident) SetWorkEnd(v string) {
	o.WorkEnd = &v
}

// GetWorkStart returns the WorkStart field value if set, zero value otherwise.
func (o *ServicenowIncident) GetWorkStart() string {
	if o == nil || IsNil(o.WorkStart) {
		var ret string
		return ret
	}
	return *o.WorkStart
}

// GetWorkStartOk returns a tuple with the WorkStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowIncident) GetWorkStartOk() (*string, bool) {
	if o == nil || IsNil(o.WorkStart) {
		return nil, false
	}
	return o.WorkStart, true
}

// HasWorkStart returns a boolean if a field has been set.
func (o *ServicenowIncident) HasWorkStart() bool {
	if o != nil && !IsNil(o.WorkStart) {
		return true
	}

	return false
}

// SetWorkStart gets a reference to the given string and assigns it to the WorkStart field.
func (o *ServicenowIncident) SetWorkStart(v string) {
	o.WorkStart = &v
}

func (o ServicenowIncident) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicenowIncident) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedServicenowInventoryEntity, errServicenowInventoryEntity := json.Marshal(o.ServicenowInventoryEntity)
	if errServicenowInventoryEntity != nil {
		return map[string]interface{}{}, errServicenowInventoryEntity
	}
	errServicenowInventoryEntity = json.Unmarshal([]byte(serializedServicenowInventoryEntity), &toSerialize)
	if errServicenowInventoryEntity != nil {
		return map[string]interface{}{}, errServicenowInventoryEntity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Approval) {
		toSerialize["Approval"] = o.Approval
	}
	if !IsNil(o.Category) {
		toSerialize["Category"] = o.Category
	}
	if !IsNil(o.Comments) {
		toSerialize["Comments"] = o.Comments
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["CreatedBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["CreatedOn"] = o.CreatedOn
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.DueDate) {
		toSerialize["DueDate"] = o.DueDate
	}
	if !IsNil(o.ExpectedStart) {
		toSerialize["ExpectedStart"] = o.ExpectedStart
	}
	if !IsNil(o.Impact) {
		toSerialize["Impact"] = o.Impact
	}
	if !IsNil(o.IncidentState) {
		toSerialize["IncidentState"] = o.IncidentState
	}
	if !IsNil(o.OpenedBy) {
		toSerialize["OpenedBy"] = o.OpenedBy
	}
	if !IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	if !IsNil(o.Risk) {
		toSerialize["Risk"] = o.Risk
	}
	if !IsNil(o.Severity) {
		toSerialize["Severity"] = o.Severity
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["ShortDescription"] = o.ShortDescription
	}
	if !IsNil(o.SysId) {
		toSerialize["SysId"] = o.SysId
	}
	if !IsNil(o.TaskEffectiveNumber) {
		toSerialize["TaskEffectiveNumber"] = o.TaskEffectiveNumber
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["UpdatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.Urgency) {
		toSerialize["Urgency"] = o.Urgency
	}
	if !IsNil(o.WorkEnd) {
		toSerialize["WorkEnd"] = o.WorkEnd
	}
	if !IsNil(o.WorkStart) {
		toSerialize["WorkStart"] = o.WorkStart
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicenowIncident) UnmarshalJSON(data []byte) (err error) {
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
	type ServicenowIncidentWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The approver property of Incident.
		Approval *string `json:"Approval,omitempty"`
		// Category property for Incident.
		Category *string `json:"Category,omitempty"`
		// Comments property on Incident.
		Comments *string `json:"Comments,omitempty"`
		// Creator property of Incident.
		CreatedBy *string `json:"CreatedBy,omitempty"`
		// Incident create date property.
		CreatedOn *string `json:"CreatedOn,omitempty"`
		// Description for Incident.
		Description *string `json:"Description,omitempty"`
		// Due date property for Incident.
		DueDate *string `json:"DueDate,omitempty"`
		// Expected start date for Incident.
		ExpectedStart *string `json:"ExpectedStart,omitempty"`
		// Impact property for Incident.
		Impact *string `json:"Impact,omitempty"`
		// State property of the Incident.
		IncidentState *string `json:"IncidentState,omitempty"`
		// Assigned to value for Incident.
		OpenedBy *string `json:"OpenedBy,omitempty"`
		// Priority property for Incident.
		Priority *string `json:"Priority,omitempty"`
		// The risk property for Incident.
		Risk *string `json:"Risk,omitempty"`
		// Severity property of the Incident.
		Severity *string `json:"Severity,omitempty"`
		// Short Description for Incident.
		ShortDescription *string `json:"ShortDescription,omitempty"`
		// System Id property for Incident.
		SysId *string `json:"SysId,omitempty"`
		// Task Effective Number for Incident.
		TaskEffectiveNumber *string `json:"TaskEffectiveNumber,omitempty"`
		// Last update by on Incident.
		UpdatedBy *string `json:"UpdatedBy,omitempty"`
		// Urgency property of the Incident.
		Urgency *string `json:"Urgency,omitempty"`
		// Work end date for Incident.
		WorkEnd *string `json:"WorkEnd,omitempty"`
		// Work start date for Incident.
		WorkStart *string `json:"WorkStart,omitempty"`
	}

	varServicenowIncidentWithoutEmbeddedStruct := ServicenowIncidentWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varServicenowIncidentWithoutEmbeddedStruct)
	if err == nil {
		varServicenowIncident := _ServicenowIncident{}
		varServicenowIncident.ClassId = varServicenowIncidentWithoutEmbeddedStruct.ClassId
		varServicenowIncident.ObjectType = varServicenowIncidentWithoutEmbeddedStruct.ObjectType
		varServicenowIncident.Approval = varServicenowIncidentWithoutEmbeddedStruct.Approval
		varServicenowIncident.Category = varServicenowIncidentWithoutEmbeddedStruct.Category
		varServicenowIncident.Comments = varServicenowIncidentWithoutEmbeddedStruct.Comments
		varServicenowIncident.CreatedBy = varServicenowIncidentWithoutEmbeddedStruct.CreatedBy
		varServicenowIncident.CreatedOn = varServicenowIncidentWithoutEmbeddedStruct.CreatedOn
		varServicenowIncident.Description = varServicenowIncidentWithoutEmbeddedStruct.Description
		varServicenowIncident.DueDate = varServicenowIncidentWithoutEmbeddedStruct.DueDate
		varServicenowIncident.ExpectedStart = varServicenowIncidentWithoutEmbeddedStruct.ExpectedStart
		varServicenowIncident.Impact = varServicenowIncidentWithoutEmbeddedStruct.Impact
		varServicenowIncident.IncidentState = varServicenowIncidentWithoutEmbeddedStruct.IncidentState
		varServicenowIncident.OpenedBy = varServicenowIncidentWithoutEmbeddedStruct.OpenedBy
		varServicenowIncident.Priority = varServicenowIncidentWithoutEmbeddedStruct.Priority
		varServicenowIncident.Risk = varServicenowIncidentWithoutEmbeddedStruct.Risk
		varServicenowIncident.Severity = varServicenowIncidentWithoutEmbeddedStruct.Severity
		varServicenowIncident.ShortDescription = varServicenowIncidentWithoutEmbeddedStruct.ShortDescription
		varServicenowIncident.SysId = varServicenowIncidentWithoutEmbeddedStruct.SysId
		varServicenowIncident.TaskEffectiveNumber = varServicenowIncidentWithoutEmbeddedStruct.TaskEffectiveNumber
		varServicenowIncident.UpdatedBy = varServicenowIncidentWithoutEmbeddedStruct.UpdatedBy
		varServicenowIncident.Urgency = varServicenowIncidentWithoutEmbeddedStruct.Urgency
		varServicenowIncident.WorkEnd = varServicenowIncidentWithoutEmbeddedStruct.WorkEnd
		varServicenowIncident.WorkStart = varServicenowIncidentWithoutEmbeddedStruct.WorkStart
		*o = ServicenowIncident(varServicenowIncident)
	} else {
		return err
	}

	varServicenowIncident := _ServicenowIncident{}

	err = json.Unmarshal(data, &varServicenowIncident)
	if err == nil {
		o.ServicenowInventoryEntity = varServicenowIncident.ServicenowInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Approval")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "Comments")
		delete(additionalProperties, "CreatedBy")
		delete(additionalProperties, "CreatedOn")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "DueDate")
		delete(additionalProperties, "ExpectedStart")
		delete(additionalProperties, "Impact")
		delete(additionalProperties, "IncidentState")
		delete(additionalProperties, "OpenedBy")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "Risk")
		delete(additionalProperties, "Severity")
		delete(additionalProperties, "ShortDescription")
		delete(additionalProperties, "SysId")
		delete(additionalProperties, "TaskEffectiveNumber")
		delete(additionalProperties, "UpdatedBy")
		delete(additionalProperties, "Urgency")
		delete(additionalProperties, "WorkEnd")
		delete(additionalProperties, "WorkStart")

		// remove fields from embedded structs
		reflectServicenowInventoryEntity := reflect.ValueOf(o.ServicenowInventoryEntity)
		for i := 0; i < reflectServicenowInventoryEntity.Type().NumField(); i++ {
			t := reflectServicenowInventoryEntity.Type().Field(i)

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

type NullableServicenowIncident struct {
	value *ServicenowIncident
	isSet bool
}

func (v NullableServicenowIncident) Get() *ServicenowIncident {
	return v.value
}

func (v *NullableServicenowIncident) Set(val *ServicenowIncident) {
	v.value = val
	v.isSet = true
}

func (v NullableServicenowIncident) IsSet() bool {
	return v.isSet
}

func (v *NullableServicenowIncident) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicenowIncident(val *ServicenowIncident) *NullableServicenowIncident {
	return &NullableServicenowIncident{value: val, isSet: true}
}

func (v NullableServicenowIncident) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicenowIncident) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
