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

// NiatelemetryInsightGroupDetails Insight group details in ND.
type NiatelemetryInsightGroupDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Alert rules count of the Insight group.
	AlertRulesCount *int64 `json:"AlertRulesCount,omitempty"`
	// Analysis setting status of the Insight group.
	AnalysisSettingsStatus *string `json:"AnalysisSettingsStatus,omitempty"`
	// Bug scan settings status of the Insight group.
	BugScanSettingsStatus *string `json:"BugScanSettingsStatus,omitempty"`
	// Delta analysis job count of the Insight group.
	DeltaAnalysisJobCount *int64 `json:"DeltaAnalysisJobCount,omitempty"`
	// Email settings count of the Insight group.
	EmailSettingsCount *int64 `json:"EmailSettingsCount,omitempty"`
	// Flow setting count of the Insight group.
	FlowSettingsCount *int64 `json:"FlowSettingsCount,omitempty"`
	// Flow setting status of the Insight group.
	FlowSettingsStatus *string `json:"FlowSettingsStatus,omitempty"`
	// Name of the Insight group.
	GroupName    *string             `json:"GroupName,omitempty"`
	InsightSites []NiatelemetrySites `json:"InsightSites,omitempty"`
	// Kafka settings count of the Insight group.
	KafkaSettingsCount *int64 `json:"KafkaSettingsCount,omitempty"`
	// Microburst setting status of the Insight group.
	MicroBurstSettingsStatus *string `json:"MicroBurstSettingsStatus,omitempty"`
	// Prechange analysis count of the Insight group.
	PrechangeAnalysisCount *int64 `json:"PrechangeAnalysisCount,omitempty"`
	// TAC collection config count of the Insight group.
	TacCollectionConfigCount *int64                               `json:"TacCollectionConfigCount,omitempty"`
	RegisteredDevice         *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _NiatelemetryInsightGroupDetails NiatelemetryInsightGroupDetails

// NewNiatelemetryInsightGroupDetails instantiates a new NiatelemetryInsightGroupDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryInsightGroupDetails(classId string, objectType string) *NiatelemetryInsightGroupDetails {
	this := NiatelemetryInsightGroupDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryInsightGroupDetailsWithDefaults instantiates a new NiatelemetryInsightGroupDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryInsightGroupDetailsWithDefaults() *NiatelemetryInsightGroupDetails {
	this := NiatelemetryInsightGroupDetails{}
	var classId string = "niatelemetry.InsightGroupDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.InsightGroupDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryInsightGroupDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryInsightGroupDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryInsightGroupDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryInsightGroupDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlertRulesCount returns the AlertRulesCount field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetAlertRulesCount() int64 {
	if o == nil || o.AlertRulesCount == nil {
		var ret int64
		return ret
	}
	return *o.AlertRulesCount
}

// GetAlertRulesCountOk returns a tuple with the AlertRulesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetAlertRulesCountOk() (*int64, bool) {
	if o == nil || o.AlertRulesCount == nil {
		return nil, false
	}
	return o.AlertRulesCount, true
}

// HasAlertRulesCount returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasAlertRulesCount() bool {
	if o != nil && o.AlertRulesCount != nil {
		return true
	}

	return false
}

// SetAlertRulesCount gets a reference to the given int64 and assigns it to the AlertRulesCount field.
func (o *NiatelemetryInsightGroupDetails) SetAlertRulesCount(v int64) {
	o.AlertRulesCount = &v
}

// GetAnalysisSettingsStatus returns the AnalysisSettingsStatus field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetAnalysisSettingsStatus() string {
	if o == nil || o.AnalysisSettingsStatus == nil {
		var ret string
		return ret
	}
	return *o.AnalysisSettingsStatus
}

// GetAnalysisSettingsStatusOk returns a tuple with the AnalysisSettingsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetAnalysisSettingsStatusOk() (*string, bool) {
	if o == nil || o.AnalysisSettingsStatus == nil {
		return nil, false
	}
	return o.AnalysisSettingsStatus, true
}

// HasAnalysisSettingsStatus returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasAnalysisSettingsStatus() bool {
	if o != nil && o.AnalysisSettingsStatus != nil {
		return true
	}

	return false
}

// SetAnalysisSettingsStatus gets a reference to the given string and assigns it to the AnalysisSettingsStatus field.
func (o *NiatelemetryInsightGroupDetails) SetAnalysisSettingsStatus(v string) {
	o.AnalysisSettingsStatus = &v
}

// GetBugScanSettingsStatus returns the BugScanSettingsStatus field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetBugScanSettingsStatus() string {
	if o == nil || o.BugScanSettingsStatus == nil {
		var ret string
		return ret
	}
	return *o.BugScanSettingsStatus
}

// GetBugScanSettingsStatusOk returns a tuple with the BugScanSettingsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetBugScanSettingsStatusOk() (*string, bool) {
	if o == nil || o.BugScanSettingsStatus == nil {
		return nil, false
	}
	return o.BugScanSettingsStatus, true
}

// HasBugScanSettingsStatus returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasBugScanSettingsStatus() bool {
	if o != nil && o.BugScanSettingsStatus != nil {
		return true
	}

	return false
}

// SetBugScanSettingsStatus gets a reference to the given string and assigns it to the BugScanSettingsStatus field.
func (o *NiatelemetryInsightGroupDetails) SetBugScanSettingsStatus(v string) {
	o.BugScanSettingsStatus = &v
}

// GetDeltaAnalysisJobCount returns the DeltaAnalysisJobCount field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetDeltaAnalysisJobCount() int64 {
	if o == nil || o.DeltaAnalysisJobCount == nil {
		var ret int64
		return ret
	}
	return *o.DeltaAnalysisJobCount
}

// GetDeltaAnalysisJobCountOk returns a tuple with the DeltaAnalysisJobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetDeltaAnalysisJobCountOk() (*int64, bool) {
	if o == nil || o.DeltaAnalysisJobCount == nil {
		return nil, false
	}
	return o.DeltaAnalysisJobCount, true
}

// HasDeltaAnalysisJobCount returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasDeltaAnalysisJobCount() bool {
	if o != nil && o.DeltaAnalysisJobCount != nil {
		return true
	}

	return false
}

// SetDeltaAnalysisJobCount gets a reference to the given int64 and assigns it to the DeltaAnalysisJobCount field.
func (o *NiatelemetryInsightGroupDetails) SetDeltaAnalysisJobCount(v int64) {
	o.DeltaAnalysisJobCount = &v
}

// GetEmailSettingsCount returns the EmailSettingsCount field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetEmailSettingsCount() int64 {
	if o == nil || o.EmailSettingsCount == nil {
		var ret int64
		return ret
	}
	return *o.EmailSettingsCount
}

// GetEmailSettingsCountOk returns a tuple with the EmailSettingsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetEmailSettingsCountOk() (*int64, bool) {
	if o == nil || o.EmailSettingsCount == nil {
		return nil, false
	}
	return o.EmailSettingsCount, true
}

// HasEmailSettingsCount returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasEmailSettingsCount() bool {
	if o != nil && o.EmailSettingsCount != nil {
		return true
	}

	return false
}

// SetEmailSettingsCount gets a reference to the given int64 and assigns it to the EmailSettingsCount field.
func (o *NiatelemetryInsightGroupDetails) SetEmailSettingsCount(v int64) {
	o.EmailSettingsCount = &v
}

// GetFlowSettingsCount returns the FlowSettingsCount field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetFlowSettingsCount() int64 {
	if o == nil || o.FlowSettingsCount == nil {
		var ret int64
		return ret
	}
	return *o.FlowSettingsCount
}

// GetFlowSettingsCountOk returns a tuple with the FlowSettingsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetFlowSettingsCountOk() (*int64, bool) {
	if o == nil || o.FlowSettingsCount == nil {
		return nil, false
	}
	return o.FlowSettingsCount, true
}

// HasFlowSettingsCount returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasFlowSettingsCount() bool {
	if o != nil && o.FlowSettingsCount != nil {
		return true
	}

	return false
}

// SetFlowSettingsCount gets a reference to the given int64 and assigns it to the FlowSettingsCount field.
func (o *NiatelemetryInsightGroupDetails) SetFlowSettingsCount(v int64) {
	o.FlowSettingsCount = &v
}

// GetFlowSettingsStatus returns the FlowSettingsStatus field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetFlowSettingsStatus() string {
	if o == nil || o.FlowSettingsStatus == nil {
		var ret string
		return ret
	}
	return *o.FlowSettingsStatus
}

// GetFlowSettingsStatusOk returns a tuple with the FlowSettingsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetFlowSettingsStatusOk() (*string, bool) {
	if o == nil || o.FlowSettingsStatus == nil {
		return nil, false
	}
	return o.FlowSettingsStatus, true
}

// HasFlowSettingsStatus returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasFlowSettingsStatus() bool {
	if o != nil && o.FlowSettingsStatus != nil {
		return true
	}

	return false
}

// SetFlowSettingsStatus gets a reference to the given string and assigns it to the FlowSettingsStatus field.
func (o *NiatelemetryInsightGroupDetails) SetFlowSettingsStatus(v string) {
	o.FlowSettingsStatus = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *NiatelemetryInsightGroupDetails) SetGroupName(v string) {
	o.GroupName = &v
}

// GetInsightSites returns the InsightSites field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryInsightGroupDetails) GetInsightSites() []NiatelemetrySites {
	if o == nil {
		var ret []NiatelemetrySites
		return ret
	}
	return o.InsightSites
}

// GetInsightSitesOk returns a tuple with the InsightSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryInsightGroupDetails) GetInsightSitesOk() ([]NiatelemetrySites, bool) {
	if o == nil || o.InsightSites == nil {
		return nil, false
	}
	return o.InsightSites, true
}

// HasInsightSites returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasInsightSites() bool {
	if o != nil && o.InsightSites != nil {
		return true
	}

	return false
}

// SetInsightSites gets a reference to the given []NiatelemetrySites and assigns it to the InsightSites field.
func (o *NiatelemetryInsightGroupDetails) SetInsightSites(v []NiatelemetrySites) {
	o.InsightSites = v
}

// GetKafkaSettingsCount returns the KafkaSettingsCount field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetKafkaSettingsCount() int64 {
	if o == nil || o.KafkaSettingsCount == nil {
		var ret int64
		return ret
	}
	return *o.KafkaSettingsCount
}

// GetKafkaSettingsCountOk returns a tuple with the KafkaSettingsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetKafkaSettingsCountOk() (*int64, bool) {
	if o == nil || o.KafkaSettingsCount == nil {
		return nil, false
	}
	return o.KafkaSettingsCount, true
}

// HasKafkaSettingsCount returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasKafkaSettingsCount() bool {
	if o != nil && o.KafkaSettingsCount != nil {
		return true
	}

	return false
}

// SetKafkaSettingsCount gets a reference to the given int64 and assigns it to the KafkaSettingsCount field.
func (o *NiatelemetryInsightGroupDetails) SetKafkaSettingsCount(v int64) {
	o.KafkaSettingsCount = &v
}

// GetMicroBurstSettingsStatus returns the MicroBurstSettingsStatus field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetMicroBurstSettingsStatus() string {
	if o == nil || o.MicroBurstSettingsStatus == nil {
		var ret string
		return ret
	}
	return *o.MicroBurstSettingsStatus
}

// GetMicroBurstSettingsStatusOk returns a tuple with the MicroBurstSettingsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetMicroBurstSettingsStatusOk() (*string, bool) {
	if o == nil || o.MicroBurstSettingsStatus == nil {
		return nil, false
	}
	return o.MicroBurstSettingsStatus, true
}

// HasMicroBurstSettingsStatus returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasMicroBurstSettingsStatus() bool {
	if o != nil && o.MicroBurstSettingsStatus != nil {
		return true
	}

	return false
}

// SetMicroBurstSettingsStatus gets a reference to the given string and assigns it to the MicroBurstSettingsStatus field.
func (o *NiatelemetryInsightGroupDetails) SetMicroBurstSettingsStatus(v string) {
	o.MicroBurstSettingsStatus = &v
}

// GetPrechangeAnalysisCount returns the PrechangeAnalysisCount field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetPrechangeAnalysisCount() int64 {
	if o == nil || o.PrechangeAnalysisCount == nil {
		var ret int64
		return ret
	}
	return *o.PrechangeAnalysisCount
}

// GetPrechangeAnalysisCountOk returns a tuple with the PrechangeAnalysisCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetPrechangeAnalysisCountOk() (*int64, bool) {
	if o == nil || o.PrechangeAnalysisCount == nil {
		return nil, false
	}
	return o.PrechangeAnalysisCount, true
}

// HasPrechangeAnalysisCount returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasPrechangeAnalysisCount() bool {
	if o != nil && o.PrechangeAnalysisCount != nil {
		return true
	}

	return false
}

// SetPrechangeAnalysisCount gets a reference to the given int64 and assigns it to the PrechangeAnalysisCount field.
func (o *NiatelemetryInsightGroupDetails) SetPrechangeAnalysisCount(v int64) {
	o.PrechangeAnalysisCount = &v
}

// GetTacCollectionConfigCount returns the TacCollectionConfigCount field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetTacCollectionConfigCount() int64 {
	if o == nil || o.TacCollectionConfigCount == nil {
		var ret int64
		return ret
	}
	return *o.TacCollectionConfigCount
}

// GetTacCollectionConfigCountOk returns a tuple with the TacCollectionConfigCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetTacCollectionConfigCountOk() (*int64, bool) {
	if o == nil || o.TacCollectionConfigCount == nil {
		return nil, false
	}
	return o.TacCollectionConfigCount, true
}

// HasTacCollectionConfigCount returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasTacCollectionConfigCount() bool {
	if o != nil && o.TacCollectionConfigCount != nil {
		return true
	}

	return false
}

// SetTacCollectionConfigCount gets a reference to the given int64 and assigns it to the TacCollectionConfigCount field.
func (o *NiatelemetryInsightGroupDetails) SetTacCollectionConfigCount(v int64) {
	o.TacCollectionConfigCount = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryInsightGroupDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryInsightGroupDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryInsightGroupDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryInsightGroupDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryInsightGroupDetails) MarshalJSON() ([]byte, error) {
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
	if o.AlertRulesCount != nil {
		toSerialize["AlertRulesCount"] = o.AlertRulesCount
	}
	if o.AnalysisSettingsStatus != nil {
		toSerialize["AnalysisSettingsStatus"] = o.AnalysisSettingsStatus
	}
	if o.BugScanSettingsStatus != nil {
		toSerialize["BugScanSettingsStatus"] = o.BugScanSettingsStatus
	}
	if o.DeltaAnalysisJobCount != nil {
		toSerialize["DeltaAnalysisJobCount"] = o.DeltaAnalysisJobCount
	}
	if o.EmailSettingsCount != nil {
		toSerialize["EmailSettingsCount"] = o.EmailSettingsCount
	}
	if o.FlowSettingsCount != nil {
		toSerialize["FlowSettingsCount"] = o.FlowSettingsCount
	}
	if o.FlowSettingsStatus != nil {
		toSerialize["FlowSettingsStatus"] = o.FlowSettingsStatus
	}
	if o.GroupName != nil {
		toSerialize["GroupName"] = o.GroupName
	}
	if o.InsightSites != nil {
		toSerialize["InsightSites"] = o.InsightSites
	}
	if o.KafkaSettingsCount != nil {
		toSerialize["KafkaSettingsCount"] = o.KafkaSettingsCount
	}
	if o.MicroBurstSettingsStatus != nil {
		toSerialize["MicroBurstSettingsStatus"] = o.MicroBurstSettingsStatus
	}
	if o.PrechangeAnalysisCount != nil {
		toSerialize["PrechangeAnalysisCount"] = o.PrechangeAnalysisCount
	}
	if o.TacCollectionConfigCount != nil {
		toSerialize["TacCollectionConfigCount"] = o.TacCollectionConfigCount
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryInsightGroupDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryInsightGroupDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Alert rules count of the Insight group.
		AlertRulesCount *int64 `json:"AlertRulesCount,omitempty"`
		// Analysis setting status of the Insight group.
		AnalysisSettingsStatus *string `json:"AnalysisSettingsStatus,omitempty"`
		// Bug scan settings status of the Insight group.
		BugScanSettingsStatus *string `json:"BugScanSettingsStatus,omitempty"`
		// Delta analysis job count of the Insight group.
		DeltaAnalysisJobCount *int64 `json:"DeltaAnalysisJobCount,omitempty"`
		// Email settings count of the Insight group.
		EmailSettingsCount *int64 `json:"EmailSettingsCount,omitempty"`
		// Flow setting count of the Insight group.
		FlowSettingsCount *int64 `json:"FlowSettingsCount,omitempty"`
		// Flow setting status of the Insight group.
		FlowSettingsStatus *string `json:"FlowSettingsStatus,omitempty"`
		// Name of the Insight group.
		GroupName    *string             `json:"GroupName,omitempty"`
		InsightSites []NiatelemetrySites `json:"InsightSites,omitempty"`
		// Kafka settings count of the Insight group.
		KafkaSettingsCount *int64 `json:"KafkaSettingsCount,omitempty"`
		// Microburst setting status of the Insight group.
		MicroBurstSettingsStatus *string `json:"MicroBurstSettingsStatus,omitempty"`
		// Prechange analysis count of the Insight group.
		PrechangeAnalysisCount *int64 `json:"PrechangeAnalysisCount,omitempty"`
		// TAC collection config count of the Insight group.
		TacCollectionConfigCount *int64                               `json:"TacCollectionConfigCount,omitempty"`
		RegisteredDevice         *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct := NiatelemetryInsightGroupDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryInsightGroupDetails := _NiatelemetryInsightGroupDetails{}
		varNiatelemetryInsightGroupDetails.ClassId = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryInsightGroupDetails.ObjectType = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryInsightGroupDetails.AlertRulesCount = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.AlertRulesCount
		varNiatelemetryInsightGroupDetails.AnalysisSettingsStatus = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.AnalysisSettingsStatus
		varNiatelemetryInsightGroupDetails.BugScanSettingsStatus = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.BugScanSettingsStatus
		varNiatelemetryInsightGroupDetails.DeltaAnalysisJobCount = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.DeltaAnalysisJobCount
		varNiatelemetryInsightGroupDetails.EmailSettingsCount = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.EmailSettingsCount
		varNiatelemetryInsightGroupDetails.FlowSettingsCount = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.FlowSettingsCount
		varNiatelemetryInsightGroupDetails.FlowSettingsStatus = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.FlowSettingsStatus
		varNiatelemetryInsightGroupDetails.GroupName = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.GroupName
		varNiatelemetryInsightGroupDetails.InsightSites = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.InsightSites
		varNiatelemetryInsightGroupDetails.KafkaSettingsCount = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.KafkaSettingsCount
		varNiatelemetryInsightGroupDetails.MicroBurstSettingsStatus = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.MicroBurstSettingsStatus
		varNiatelemetryInsightGroupDetails.PrechangeAnalysisCount = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.PrechangeAnalysisCount
		varNiatelemetryInsightGroupDetails.TacCollectionConfigCount = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.TacCollectionConfigCount
		varNiatelemetryInsightGroupDetails.RegisteredDevice = varNiatelemetryInsightGroupDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryInsightGroupDetails(varNiatelemetryInsightGroupDetails)
	} else {
		return err
	}

	varNiatelemetryInsightGroupDetails := _NiatelemetryInsightGroupDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryInsightGroupDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryInsightGroupDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlertRulesCount")
		delete(additionalProperties, "AnalysisSettingsStatus")
		delete(additionalProperties, "BugScanSettingsStatus")
		delete(additionalProperties, "DeltaAnalysisJobCount")
		delete(additionalProperties, "EmailSettingsCount")
		delete(additionalProperties, "FlowSettingsCount")
		delete(additionalProperties, "FlowSettingsStatus")
		delete(additionalProperties, "GroupName")
		delete(additionalProperties, "InsightSites")
		delete(additionalProperties, "KafkaSettingsCount")
		delete(additionalProperties, "MicroBurstSettingsStatus")
		delete(additionalProperties, "PrechangeAnalysisCount")
		delete(additionalProperties, "TacCollectionConfigCount")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryInsightGroupDetails struct {
	value *NiatelemetryInsightGroupDetails
	isSet bool
}

func (v NullableNiatelemetryInsightGroupDetails) Get() *NiatelemetryInsightGroupDetails {
	return v.value
}

func (v *NullableNiatelemetryInsightGroupDetails) Set(val *NiatelemetryInsightGroupDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryInsightGroupDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryInsightGroupDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryInsightGroupDetails(val *NiatelemetryInsightGroupDetails) *NullableNiatelemetryInsightGroupDetails {
	return &NullableNiatelemetryInsightGroupDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryInsightGroupDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryInsightGroupDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
