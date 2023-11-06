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
	"time"
)

// ApplianceClusterInstallBaseAllOf Definition of the list of properties defined in 'appliance.ClusterInstallBase', excluding properties defined in parent classes.
type ApplianceClusterInstallBaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType      string                               `json:"ObjectType"`
	CompletedPhases []ApplianceClusterInstallPhase       `json:"CompletedPhases,omitempty"`
	CurrentPhase    NullableApplianceClusterInstallPhase `json:"CurrentPhase,omitempty"`
	// Elapsed time in seconds during the software install.
	ElapsedTime *int64 `json:"ElapsedTime,omitempty"`
	// End date of the software install.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Error code for Intersight Appliance's software install. In case of failure - this code will help decide if software install can be retried.
	ErrorCode *int64                `json:"ErrorCode,omitempty"`
	Messages  []string              `json:"Messages,omitempty"`
	NodeInfo  []ApplianceNodeIpInfo `json:"NodeInfo,omitempty"`
	// Round robin DNS address, which should be able to resolve the hostnames of all the nodes in the cluster.
	RemoteDns *string `json:"RemoteDns,omitempty"`
	// Start date of the software install. UI can modify startTime to re-schedule an install.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Status of the Intersight Appliance's software install. * `NotReady` - Cluster is not ready. Install cannot be triggered. * `Ready` - Cluster is ready. Install can be triggered. * `InProgress` - Install is currently in progress. * `Success` - Install was run and succeeded. * `Fail` - Install was run and failed.
	Status *string `json:"Status,omitempty"`
	// Total number of nodes in the system.
	TotalNodes *int64 `json:"TotalNodes,omitempty"`
	// TotalPhase represents the total number of the install phases for one install.
	TotalPhases          *int64                      `json:"TotalPhases,omitempty"`
	Vip                  NullableApplianceNodeIpInfo `json:"Vip,omitempty"`
	Account              *IamAccountRelationship     `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceClusterInstallBaseAllOf ApplianceClusterInstallBaseAllOf

// NewApplianceClusterInstallBaseAllOf instantiates a new ApplianceClusterInstallBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceClusterInstallBaseAllOf(classId string, objectType string) *ApplianceClusterInstallBaseAllOf {
	this := ApplianceClusterInstallBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceClusterInstallBaseAllOfWithDefaults instantiates a new ApplianceClusterInstallBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceClusterInstallBaseAllOfWithDefaults() *ApplianceClusterInstallBaseAllOf {
	this := ApplianceClusterInstallBaseAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceClusterInstallBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceClusterInstallBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceClusterInstallBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceClusterInstallBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCompletedPhases returns the CompletedPhases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceClusterInstallBaseAllOf) GetCompletedPhases() []ApplianceClusterInstallPhase {
	if o == nil {
		var ret []ApplianceClusterInstallPhase
		return ret
	}
	return o.CompletedPhases
}

// GetCompletedPhasesOk returns a tuple with the CompletedPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceClusterInstallBaseAllOf) GetCompletedPhasesOk() ([]ApplianceClusterInstallPhase, bool) {
	if o == nil || o.CompletedPhases == nil {
		return nil, false
	}
	return o.CompletedPhases, true
}

// HasCompletedPhases returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasCompletedPhases() bool {
	if o != nil && o.CompletedPhases != nil {
		return true
	}

	return false
}

// SetCompletedPhases gets a reference to the given []ApplianceClusterInstallPhase and assigns it to the CompletedPhases field.
func (o *ApplianceClusterInstallBaseAllOf) SetCompletedPhases(v []ApplianceClusterInstallPhase) {
	o.CompletedPhases = v
}

// GetCurrentPhase returns the CurrentPhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceClusterInstallBaseAllOf) GetCurrentPhase() ApplianceClusterInstallPhase {
	if o == nil || o.CurrentPhase.Get() == nil {
		var ret ApplianceClusterInstallPhase
		return ret
	}
	return *o.CurrentPhase.Get()
}

// GetCurrentPhaseOk returns a tuple with the CurrentPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceClusterInstallBaseAllOf) GetCurrentPhaseOk() (*ApplianceClusterInstallPhase, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPhase.Get(), o.CurrentPhase.IsSet()
}

// HasCurrentPhase returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasCurrentPhase() bool {
	if o != nil && o.CurrentPhase.IsSet() {
		return true
	}

	return false
}

// SetCurrentPhase gets a reference to the given NullableApplianceClusterInstallPhase and assigns it to the CurrentPhase field.
func (o *ApplianceClusterInstallBaseAllOf) SetCurrentPhase(v ApplianceClusterInstallPhase) {
	o.CurrentPhase.Set(&v)
}

// SetCurrentPhaseNil sets the value for CurrentPhase to be an explicit nil
func (o *ApplianceClusterInstallBaseAllOf) SetCurrentPhaseNil() {
	o.CurrentPhase.Set(nil)
}

// UnsetCurrentPhase ensures that no value is present for CurrentPhase, not even an explicit nil
func (o *ApplianceClusterInstallBaseAllOf) UnsetCurrentPhase() {
	o.CurrentPhase.Unset()
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *ApplianceClusterInstallBaseAllOf) GetElapsedTime() int64 {
	if o == nil || o.ElapsedTime == nil {
		var ret int64
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetElapsedTimeOk() (*int64, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int64 and assigns it to the ElapsedTime field.
func (o *ApplianceClusterInstallBaseAllOf) SetElapsedTime(v int64) {
	o.ElapsedTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplianceClusterInstallBaseAllOf) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplianceClusterInstallBaseAllOf) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ApplianceClusterInstallBaseAllOf) GetErrorCode() int64 {
	if o == nil || o.ErrorCode == nil {
		var ret int64
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetErrorCodeOk() (*int64, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int64 and assigns it to the ErrorCode field.
func (o *ApplianceClusterInstallBaseAllOf) SetErrorCode(v int64) {
	o.ErrorCode = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceClusterInstallBaseAllOf) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceClusterInstallBaseAllOf) GetMessagesOk() ([]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *ApplianceClusterInstallBaseAllOf) SetMessages(v []string) {
	o.Messages = v
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceClusterInstallBaseAllOf) GetNodeInfo() []ApplianceNodeIpInfo {
	if o == nil {
		var ret []ApplianceNodeIpInfo
		return ret
	}
	return o.NodeInfo
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceClusterInstallBaseAllOf) GetNodeInfoOk() ([]ApplianceNodeIpInfo, bool) {
	if o == nil || o.NodeInfo == nil {
		return nil, false
	}
	return o.NodeInfo, true
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasNodeInfo() bool {
	if o != nil && o.NodeInfo != nil {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given []ApplianceNodeIpInfo and assigns it to the NodeInfo field.
func (o *ApplianceClusterInstallBaseAllOf) SetNodeInfo(v []ApplianceNodeIpInfo) {
	o.NodeInfo = v
}

// GetRemoteDns returns the RemoteDns field value if set, zero value otherwise.
func (o *ApplianceClusterInstallBaseAllOf) GetRemoteDns() string {
	if o == nil || o.RemoteDns == nil {
		var ret string
		return ret
	}
	return *o.RemoteDns
}

// GetRemoteDnsOk returns a tuple with the RemoteDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetRemoteDnsOk() (*string, bool) {
	if o == nil || o.RemoteDns == nil {
		return nil, false
	}
	return o.RemoteDns, true
}

// HasRemoteDns returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasRemoteDns() bool {
	if o != nil && o.RemoteDns != nil {
		return true
	}

	return false
}

// SetRemoteDns gets a reference to the given string and assigns it to the RemoteDns field.
func (o *ApplianceClusterInstallBaseAllOf) SetRemoteDns(v string) {
	o.RemoteDns = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplianceClusterInstallBaseAllOf) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplianceClusterInstallBaseAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceClusterInstallBaseAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceClusterInstallBaseAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTotalNodes returns the TotalNodes field value if set, zero value otherwise.
func (o *ApplianceClusterInstallBaseAllOf) GetTotalNodes() int64 {
	if o == nil || o.TotalNodes == nil {
		var ret int64
		return ret
	}
	return *o.TotalNodes
}

// GetTotalNodesOk returns a tuple with the TotalNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetTotalNodesOk() (*int64, bool) {
	if o == nil || o.TotalNodes == nil {
		return nil, false
	}
	return o.TotalNodes, true
}

// HasTotalNodes returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasTotalNodes() bool {
	if o != nil && o.TotalNodes != nil {
		return true
	}

	return false
}

// SetTotalNodes gets a reference to the given int64 and assigns it to the TotalNodes field.
func (o *ApplianceClusterInstallBaseAllOf) SetTotalNodes(v int64) {
	o.TotalNodes = &v
}

// GetTotalPhases returns the TotalPhases field value if set, zero value otherwise.
func (o *ApplianceClusterInstallBaseAllOf) GetTotalPhases() int64 {
	if o == nil || o.TotalPhases == nil {
		var ret int64
		return ret
	}
	return *o.TotalPhases
}

// GetTotalPhasesOk returns a tuple with the TotalPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetTotalPhasesOk() (*int64, bool) {
	if o == nil || o.TotalPhases == nil {
		return nil, false
	}
	return o.TotalPhases, true
}

// HasTotalPhases returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasTotalPhases() bool {
	if o != nil && o.TotalPhases != nil {
		return true
	}

	return false
}

// SetTotalPhases gets a reference to the given int64 and assigns it to the TotalPhases field.
func (o *ApplianceClusterInstallBaseAllOf) SetTotalPhases(v int64) {
	o.TotalPhases = &v
}

// GetVip returns the Vip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceClusterInstallBaseAllOf) GetVip() ApplianceNodeIpInfo {
	if o == nil || o.Vip.Get() == nil {
		var ret ApplianceNodeIpInfo
		return ret
	}
	return *o.Vip.Get()
}

// GetVipOk returns a tuple with the Vip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceClusterInstallBaseAllOf) GetVipOk() (*ApplianceNodeIpInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vip.Get(), o.Vip.IsSet()
}

// HasVip returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasVip() bool {
	if o != nil && o.Vip.IsSet() {
		return true
	}

	return false
}

// SetVip gets a reference to the given NullableApplianceNodeIpInfo and assigns it to the Vip field.
func (o *ApplianceClusterInstallBaseAllOf) SetVip(v ApplianceNodeIpInfo) {
	o.Vip.Set(&v)
}

// SetVipNil sets the value for Vip to be an explicit nil
func (o *ApplianceClusterInstallBaseAllOf) SetVipNil() {
	o.Vip.Set(nil)
}

// UnsetVip ensures that no value is present for Vip, not even an explicit nil
func (o *ApplianceClusterInstallBaseAllOf) UnsetVip() {
	o.Vip.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceClusterInstallBaseAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterInstallBaseAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceClusterInstallBaseAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceClusterInstallBaseAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceClusterInstallBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CompletedPhases != nil {
		toSerialize["CompletedPhases"] = o.CompletedPhases
	}
	if o.CurrentPhase.IsSet() {
		toSerialize["CurrentPhase"] = o.CurrentPhase.Get()
	}
	if o.ElapsedTime != nil {
		toSerialize["ElapsedTime"] = o.ElapsedTime
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.NodeInfo != nil {
		toSerialize["NodeInfo"] = o.NodeInfo
	}
	if o.RemoteDns != nil {
		toSerialize["RemoteDns"] = o.RemoteDns
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TotalNodes != nil {
		toSerialize["TotalNodes"] = o.TotalNodes
	}
	if o.TotalPhases != nil {
		toSerialize["TotalPhases"] = o.TotalPhases
	}
	if o.Vip.IsSet() {
		toSerialize["Vip"] = o.Vip.Get()
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceClusterInstallBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceClusterInstallBaseAllOf := _ApplianceClusterInstallBaseAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceClusterInstallBaseAllOf); err == nil {
		*o = ApplianceClusterInstallBaseAllOf(varApplianceClusterInstallBaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CompletedPhases")
		delete(additionalProperties, "CurrentPhase")
		delete(additionalProperties, "ElapsedTime")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "ErrorCode")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "NodeInfo")
		delete(additionalProperties, "RemoteDns")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TotalNodes")
		delete(additionalProperties, "TotalPhases")
		delete(additionalProperties, "Vip")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceClusterInstallBaseAllOf struct {
	value *ApplianceClusterInstallBaseAllOf
	isSet bool
}

func (v NullableApplianceClusterInstallBaseAllOf) Get() *ApplianceClusterInstallBaseAllOf {
	return v.value
}

func (v *NullableApplianceClusterInstallBaseAllOf) Set(val *ApplianceClusterInstallBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceClusterInstallBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceClusterInstallBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceClusterInstallBaseAllOf(val *ApplianceClusterInstallBaseAllOf) *NullableApplianceClusterInstallBaseAllOf {
	return &NullableApplianceClusterInstallBaseAllOf{value: val, isSet: true}
}

func (v NullableApplianceClusterInstallBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceClusterInstallBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
