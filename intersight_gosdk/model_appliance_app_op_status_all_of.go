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

// ApplianceAppOpStatusAllOf Definition of the list of properties defined in 'appliance.AppOpStatus', excluding properties defined in parent classes.
type ApplianceAppOpStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string               `json:"ObjectType"`
	ApiStatuses []ApplianceApiStatus `json:"ApiStatuses,omitempty"`
	// Unique label to identify the application.
	AppLabel *string `json:"AppLabel,omitempty"`
	// Operational status of the application. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - The status of the appliance node is unknown. * `Operational` - The appliance node is operational. * `Impaired` - The appliance node is impaired. * `AttentionNeeded` - The appliance node needs attention. * `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster. * `OutOfService` - The user has taken this node (part of a cluster) to out of service. * `ReadyForReplacement` - The cluster node is ready to be replaced. * `ReplacementInProgress` - The cluster node replacement is in progress. * `ReplacementFailed` - There was a failure during the cluster node replacement.
	OperationalStatus *string `json:"OperationalStatus,omitempty"`
	// Number of replicas ready.  The number of instances of the application currently ready to perform its intended functions.
	ReadyCount *int64 `json:"ReadyCount,omitempty"`
	// Number of replicas provisioned. The number of instances of the application provisioned to run on the Intersight appliance.
	ReplicaCount *int64 `json:"ReplicaCount,omitempty"`
	// Number of instance restarts in the last hour.
	RestartCount1Hour *int64 `json:"RestartCount1Hour,omitempty"`
	// Number of instance restarts in the last 24 hours.
	RestartCount24Hours *int64 `json:"RestartCount24Hours,omitempty"`
	// Number of instance restarts in the last 5 minutes.
	RestartCount5Mins *int64 `json:"RestartCount5Mins,omitempty"`
	// Total number of restarts since last deployment.
	RestartCountTotal *int64 `json:"RestartCountTotal,omitempty"`
	// Number of replicas running. The number of instances of the application currently running.
	RunningCount         *int64                               `json:"RunningCount,omitempty"`
	GroupOpStatus        *ApplianceGroupOpStatusRelationship  `json:"GroupOpStatus,omitempty"`
	SystemOpStatus       *ApplianceSystemOpStatusRelationship `json:"SystemOpStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceAppOpStatusAllOf ApplianceAppOpStatusAllOf

// NewApplianceAppOpStatusAllOf instantiates a new ApplianceAppOpStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAppOpStatusAllOf(classId string, objectType string) *ApplianceAppOpStatusAllOf {
	this := ApplianceAppOpStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceAppOpStatusAllOfWithDefaults instantiates a new ApplianceAppOpStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAppOpStatusAllOfWithDefaults() *ApplianceAppOpStatusAllOf {
	this := ApplianceAppOpStatusAllOf{}
	var classId string = "appliance.AppOpStatus"
	this.ClassId = classId
	var objectType string = "appliance.AppOpStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceAppOpStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceAppOpStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceAppOpStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceAppOpStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApiStatuses returns the ApiStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceAppOpStatusAllOf) GetApiStatuses() []ApplianceApiStatus {
	if o == nil {
		var ret []ApplianceApiStatus
		return ret
	}
	return o.ApiStatuses
}

// GetApiStatusesOk returns a tuple with the ApiStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceAppOpStatusAllOf) GetApiStatusesOk() ([]ApplianceApiStatus, bool) {
	if o == nil || o.ApiStatuses == nil {
		return nil, false
	}
	return o.ApiStatuses, true
}

// HasApiStatuses returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasApiStatuses() bool {
	if o != nil && o.ApiStatuses != nil {
		return true
	}

	return false
}

// SetApiStatuses gets a reference to the given []ApplianceApiStatus and assigns it to the ApiStatuses field.
func (o *ApplianceAppOpStatusAllOf) SetApiStatuses(v []ApplianceApiStatus) {
	o.ApiStatuses = v
}

// GetAppLabel returns the AppLabel field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetAppLabel() string {
	if o == nil || o.AppLabel == nil {
		var ret string
		return ret
	}
	return *o.AppLabel
}

// GetAppLabelOk returns a tuple with the AppLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetAppLabelOk() (*string, bool) {
	if o == nil || o.AppLabel == nil {
		return nil, false
	}
	return o.AppLabel, true
}

// HasAppLabel returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasAppLabel() bool {
	if o != nil && o.AppLabel != nil {
		return true
	}

	return false
}

// SetAppLabel gets a reference to the given string and assigns it to the AppLabel field.
func (o *ApplianceAppOpStatusAllOf) SetAppLabel(v string) {
	o.AppLabel = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceAppOpStatusAllOf) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetReadyCount returns the ReadyCount field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetReadyCount() int64 {
	if o == nil || o.ReadyCount == nil {
		var ret int64
		return ret
	}
	return *o.ReadyCount
}

// GetReadyCountOk returns a tuple with the ReadyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetReadyCountOk() (*int64, bool) {
	if o == nil || o.ReadyCount == nil {
		return nil, false
	}
	return o.ReadyCount, true
}

// HasReadyCount returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasReadyCount() bool {
	if o != nil && o.ReadyCount != nil {
		return true
	}

	return false
}

// SetReadyCount gets a reference to the given int64 and assigns it to the ReadyCount field.
func (o *ApplianceAppOpStatusAllOf) SetReadyCount(v int64) {
	o.ReadyCount = &v
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetReplicaCount() int64 {
	if o == nil || o.ReplicaCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetReplicaCountOk() (*int64, bool) {
	if o == nil || o.ReplicaCount == nil {
		return nil, false
	}
	return o.ReplicaCount, true
}

// HasReplicaCount returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasReplicaCount() bool {
	if o != nil && o.ReplicaCount != nil {
		return true
	}

	return false
}

// SetReplicaCount gets a reference to the given int64 and assigns it to the ReplicaCount field.
func (o *ApplianceAppOpStatusAllOf) SetReplicaCount(v int64) {
	o.ReplicaCount = &v
}

// GetRestartCount1Hour returns the RestartCount1Hour field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetRestartCount1Hour() int64 {
	if o == nil || o.RestartCount1Hour == nil {
		var ret int64
		return ret
	}
	return *o.RestartCount1Hour
}

// GetRestartCount1HourOk returns a tuple with the RestartCount1Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetRestartCount1HourOk() (*int64, bool) {
	if o == nil || o.RestartCount1Hour == nil {
		return nil, false
	}
	return o.RestartCount1Hour, true
}

// HasRestartCount1Hour returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasRestartCount1Hour() bool {
	if o != nil && o.RestartCount1Hour != nil {
		return true
	}

	return false
}

// SetRestartCount1Hour gets a reference to the given int64 and assigns it to the RestartCount1Hour field.
func (o *ApplianceAppOpStatusAllOf) SetRestartCount1Hour(v int64) {
	o.RestartCount1Hour = &v
}

// GetRestartCount24Hours returns the RestartCount24Hours field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetRestartCount24Hours() int64 {
	if o == nil || o.RestartCount24Hours == nil {
		var ret int64
		return ret
	}
	return *o.RestartCount24Hours
}

// GetRestartCount24HoursOk returns a tuple with the RestartCount24Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetRestartCount24HoursOk() (*int64, bool) {
	if o == nil || o.RestartCount24Hours == nil {
		return nil, false
	}
	return o.RestartCount24Hours, true
}

// HasRestartCount24Hours returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasRestartCount24Hours() bool {
	if o != nil && o.RestartCount24Hours != nil {
		return true
	}

	return false
}

// SetRestartCount24Hours gets a reference to the given int64 and assigns it to the RestartCount24Hours field.
func (o *ApplianceAppOpStatusAllOf) SetRestartCount24Hours(v int64) {
	o.RestartCount24Hours = &v
}

// GetRestartCount5Mins returns the RestartCount5Mins field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetRestartCount5Mins() int64 {
	if o == nil || o.RestartCount5Mins == nil {
		var ret int64
		return ret
	}
	return *o.RestartCount5Mins
}

// GetRestartCount5MinsOk returns a tuple with the RestartCount5Mins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetRestartCount5MinsOk() (*int64, bool) {
	if o == nil || o.RestartCount5Mins == nil {
		return nil, false
	}
	return o.RestartCount5Mins, true
}

// HasRestartCount5Mins returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasRestartCount5Mins() bool {
	if o != nil && o.RestartCount5Mins != nil {
		return true
	}

	return false
}

// SetRestartCount5Mins gets a reference to the given int64 and assigns it to the RestartCount5Mins field.
func (o *ApplianceAppOpStatusAllOf) SetRestartCount5Mins(v int64) {
	o.RestartCount5Mins = &v
}

// GetRestartCountTotal returns the RestartCountTotal field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetRestartCountTotal() int64 {
	if o == nil || o.RestartCountTotal == nil {
		var ret int64
		return ret
	}
	return *o.RestartCountTotal
}

// GetRestartCountTotalOk returns a tuple with the RestartCountTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetRestartCountTotalOk() (*int64, bool) {
	if o == nil || o.RestartCountTotal == nil {
		return nil, false
	}
	return o.RestartCountTotal, true
}

// HasRestartCountTotal returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasRestartCountTotal() bool {
	if o != nil && o.RestartCountTotal != nil {
		return true
	}

	return false
}

// SetRestartCountTotal gets a reference to the given int64 and assigns it to the RestartCountTotal field.
func (o *ApplianceAppOpStatusAllOf) SetRestartCountTotal(v int64) {
	o.RestartCountTotal = &v
}

// GetRunningCount returns the RunningCount field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetRunningCount() int64 {
	if o == nil || o.RunningCount == nil {
		var ret int64
		return ret
	}
	return *o.RunningCount
}

// GetRunningCountOk returns a tuple with the RunningCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetRunningCountOk() (*int64, bool) {
	if o == nil || o.RunningCount == nil {
		return nil, false
	}
	return o.RunningCount, true
}

// HasRunningCount returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasRunningCount() bool {
	if o != nil && o.RunningCount != nil {
		return true
	}

	return false
}

// SetRunningCount gets a reference to the given int64 and assigns it to the RunningCount field.
func (o *ApplianceAppOpStatusAllOf) SetRunningCount(v int64) {
	o.RunningCount = &v
}

// GetGroupOpStatus returns the GroupOpStatus field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetGroupOpStatus() ApplianceGroupOpStatusRelationship {
	if o == nil || o.GroupOpStatus == nil {
		var ret ApplianceGroupOpStatusRelationship
		return ret
	}
	return *o.GroupOpStatus
}

// GetGroupOpStatusOk returns a tuple with the GroupOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetGroupOpStatusOk() (*ApplianceGroupOpStatusRelationship, bool) {
	if o == nil || o.GroupOpStatus == nil {
		return nil, false
	}
	return o.GroupOpStatus, true
}

// HasGroupOpStatus returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasGroupOpStatus() bool {
	if o != nil && o.GroupOpStatus != nil {
		return true
	}

	return false
}

// SetGroupOpStatus gets a reference to the given ApplianceGroupOpStatusRelationship and assigns it to the GroupOpStatus field.
func (o *ApplianceAppOpStatusAllOf) SetGroupOpStatus(v ApplianceGroupOpStatusRelationship) {
	o.GroupOpStatus = &v
}

// GetSystemOpStatus returns the SystemOpStatus field value if set, zero value otherwise.
func (o *ApplianceAppOpStatusAllOf) GetSystemOpStatus() ApplianceSystemOpStatusRelationship {
	if o == nil || o.SystemOpStatus == nil {
		var ret ApplianceSystemOpStatusRelationship
		return ret
	}
	return *o.SystemOpStatus
}

// GetSystemOpStatusOk returns a tuple with the SystemOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAppOpStatusAllOf) GetSystemOpStatusOk() (*ApplianceSystemOpStatusRelationship, bool) {
	if o == nil || o.SystemOpStatus == nil {
		return nil, false
	}
	return o.SystemOpStatus, true
}

// HasSystemOpStatus returns a boolean if a field has been set.
func (o *ApplianceAppOpStatusAllOf) HasSystemOpStatus() bool {
	if o != nil && o.SystemOpStatus != nil {
		return true
	}

	return false
}

// SetSystemOpStatus gets a reference to the given ApplianceSystemOpStatusRelationship and assigns it to the SystemOpStatus field.
func (o *ApplianceAppOpStatusAllOf) SetSystemOpStatus(v ApplianceSystemOpStatusRelationship) {
	o.SystemOpStatus = &v
}

func (o ApplianceAppOpStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ApiStatuses != nil {
		toSerialize["ApiStatuses"] = o.ApiStatuses
	}
	if o.AppLabel != nil {
		toSerialize["AppLabel"] = o.AppLabel
	}
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if o.ReadyCount != nil {
		toSerialize["ReadyCount"] = o.ReadyCount
	}
	if o.ReplicaCount != nil {
		toSerialize["ReplicaCount"] = o.ReplicaCount
	}
	if o.RestartCount1Hour != nil {
		toSerialize["RestartCount1Hour"] = o.RestartCount1Hour
	}
	if o.RestartCount24Hours != nil {
		toSerialize["RestartCount24Hours"] = o.RestartCount24Hours
	}
	if o.RestartCount5Mins != nil {
		toSerialize["RestartCount5Mins"] = o.RestartCount5Mins
	}
	if o.RestartCountTotal != nil {
		toSerialize["RestartCountTotal"] = o.RestartCountTotal
	}
	if o.RunningCount != nil {
		toSerialize["RunningCount"] = o.RunningCount
	}
	if o.GroupOpStatus != nil {
		toSerialize["GroupOpStatus"] = o.GroupOpStatus
	}
	if o.SystemOpStatus != nil {
		toSerialize["SystemOpStatus"] = o.SystemOpStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceAppOpStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceAppOpStatusAllOf := _ApplianceAppOpStatusAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceAppOpStatusAllOf); err == nil {
		*o = ApplianceAppOpStatusAllOf(varApplianceAppOpStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiStatuses")
		delete(additionalProperties, "AppLabel")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "ReadyCount")
		delete(additionalProperties, "ReplicaCount")
		delete(additionalProperties, "RestartCount1Hour")
		delete(additionalProperties, "RestartCount24Hours")
		delete(additionalProperties, "RestartCount5Mins")
		delete(additionalProperties, "RestartCountTotal")
		delete(additionalProperties, "RunningCount")
		delete(additionalProperties, "GroupOpStatus")
		delete(additionalProperties, "SystemOpStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceAppOpStatusAllOf struct {
	value *ApplianceAppOpStatusAllOf
	isSet bool
}

func (v NullableApplianceAppOpStatusAllOf) Get() *ApplianceAppOpStatusAllOf {
	return v.value
}

func (v *NullableApplianceAppOpStatusAllOf) Set(val *ApplianceAppOpStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAppOpStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAppOpStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAppOpStatusAllOf(val *ApplianceAppOpStatusAllOf) *NullableApplianceAppOpStatusAllOf {
	return &NullableApplianceAppOpStatusAllOf{value: val, isSet: true}
}

func (v NullableApplianceAppOpStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAppOpStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
