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

// PartnerintegrationDeviceConnectorAllOf Definition of the list of properties defined in 'partnerintegration.DeviceConnector', excluding properties defined in parent classes.
type PartnerintegrationDeviceConnectorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action to be performed on the device connector recipe. * `None` - Default Value of the action, i.e. do nothing. * `Build` - Build the device connector service image. * `Deploy` - Deploy the device connector service on the appliance. * `Upload` - Upload a file to the Partner Integration Appliance bucket.
	Action *string `json:"Action,omitempty"`
	// Time when build was triggered.
	BuildStartTime *string `json:"BuildStartTime,omitempty"`
	// Status of build for device connector recipe. * `None` - Default value of the status. i.e. done nothing. * `BackendInProgress` - The backend build is in progress. * `BackendFailed` - The backend build has failed. * `DockerInProgress` - The docker build is in progress. * `DockerFailed` - The docker build has failed. * `Completed` - The operation completed successfully.
	BuildStatus *string `json:"BuildStatus,omitempty"`
	// Name of the docker image that is built.
	ImageName *string `json:"ImageName,omitempty"`
	// Name of the device connector recipe.
	Name *string `json:"Name,omitempty"`
	// Name of the bucket to pick up the file from.
	SrcBucket *string `json:"SrcBucket,omitempty"`
	// Name of source file to upload.
	SrcFileName *string `json:"SrcFileName,omitempty"`
	// An array of relationships to partnerintegrationDcLogs resources.
	Logs                 []PartnerintegrationDcLogsRelationship `json:"Logs,omitempty"`
	Organization         *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerintegrationDeviceConnectorAllOf PartnerintegrationDeviceConnectorAllOf

// NewPartnerintegrationDeviceConnectorAllOf instantiates a new PartnerintegrationDeviceConnectorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerintegrationDeviceConnectorAllOf(classId string, objectType string) *PartnerintegrationDeviceConnectorAllOf {
	this := PartnerintegrationDeviceConnectorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// NewPartnerintegrationDeviceConnectorAllOfWithDefaults instantiates a new PartnerintegrationDeviceConnectorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerintegrationDeviceConnectorAllOfWithDefaults() *PartnerintegrationDeviceConnectorAllOf {
	this := PartnerintegrationDeviceConnectorAllOf{}
	var classId string = "partnerintegration.DeviceConnector"
	this.ClassId = classId
	var objectType string = "partnerintegration.DeviceConnector"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *PartnerintegrationDeviceConnectorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PartnerintegrationDeviceConnectorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PartnerintegrationDeviceConnectorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PartnerintegrationDeviceConnectorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PartnerintegrationDeviceConnectorAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PartnerintegrationDeviceConnectorAllOf) SetAction(v string) {
	o.Action = &v
}

// GetBuildStartTime returns the BuildStartTime field value if set, zero value otherwise.
func (o *PartnerintegrationDeviceConnectorAllOf) GetBuildStartTime() string {
	if o == nil || o.BuildStartTime == nil {
		var ret string
		return ret
	}
	return *o.BuildStartTime
}

// GetBuildStartTimeOk returns a tuple with the BuildStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) GetBuildStartTimeOk() (*string, bool) {
	if o == nil || o.BuildStartTime == nil {
		return nil, false
	}
	return o.BuildStartTime, true
}

// HasBuildStartTime returns a boolean if a field has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) HasBuildStartTime() bool {
	if o != nil && o.BuildStartTime != nil {
		return true
	}

	return false
}

// SetBuildStartTime gets a reference to the given string and assigns it to the BuildStartTime field.
func (o *PartnerintegrationDeviceConnectorAllOf) SetBuildStartTime(v string) {
	o.BuildStartTime = &v
}

// GetBuildStatus returns the BuildStatus field value if set, zero value otherwise.
func (o *PartnerintegrationDeviceConnectorAllOf) GetBuildStatus() string {
	if o == nil || o.BuildStatus == nil {
		var ret string
		return ret
	}
	return *o.BuildStatus
}

// GetBuildStatusOk returns a tuple with the BuildStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) GetBuildStatusOk() (*string, bool) {
	if o == nil || o.BuildStatus == nil {
		return nil, false
	}
	return o.BuildStatus, true
}

// HasBuildStatus returns a boolean if a field has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) HasBuildStatus() bool {
	if o != nil && o.BuildStatus != nil {
		return true
	}

	return false
}

// SetBuildStatus gets a reference to the given string and assigns it to the BuildStatus field.
func (o *PartnerintegrationDeviceConnectorAllOf) SetBuildStatus(v string) {
	o.BuildStatus = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *PartnerintegrationDeviceConnectorAllOf) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *PartnerintegrationDeviceConnectorAllOf) SetImageName(v string) {
	o.ImageName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartnerintegrationDeviceConnectorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartnerintegrationDeviceConnectorAllOf) SetName(v string) {
	o.Name = &v
}

// GetSrcBucket returns the SrcBucket field value if set, zero value otherwise.
func (o *PartnerintegrationDeviceConnectorAllOf) GetSrcBucket() string {
	if o == nil || o.SrcBucket == nil {
		var ret string
		return ret
	}
	return *o.SrcBucket
}

// GetSrcBucketOk returns a tuple with the SrcBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) GetSrcBucketOk() (*string, bool) {
	if o == nil || o.SrcBucket == nil {
		return nil, false
	}
	return o.SrcBucket, true
}

// HasSrcBucket returns a boolean if a field has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) HasSrcBucket() bool {
	if o != nil && o.SrcBucket != nil {
		return true
	}

	return false
}

// SetSrcBucket gets a reference to the given string and assigns it to the SrcBucket field.
func (o *PartnerintegrationDeviceConnectorAllOf) SetSrcBucket(v string) {
	o.SrcBucket = &v
}

// GetSrcFileName returns the SrcFileName field value if set, zero value otherwise.
func (o *PartnerintegrationDeviceConnectorAllOf) GetSrcFileName() string {
	if o == nil || o.SrcFileName == nil {
		var ret string
		return ret
	}
	return *o.SrcFileName
}

// GetSrcFileNameOk returns a tuple with the SrcFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) GetSrcFileNameOk() (*string, bool) {
	if o == nil || o.SrcFileName == nil {
		return nil, false
	}
	return o.SrcFileName, true
}

// HasSrcFileName returns a boolean if a field has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) HasSrcFileName() bool {
	if o != nil && o.SrcFileName != nil {
		return true
	}

	return false
}

// SetSrcFileName gets a reference to the given string and assigns it to the SrcFileName field.
func (o *PartnerintegrationDeviceConnectorAllOf) SetSrcFileName(v string) {
	o.SrcFileName = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationDeviceConnectorAllOf) GetLogs() []PartnerintegrationDcLogsRelationship {
	if o == nil {
		var ret []PartnerintegrationDcLogsRelationship
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationDeviceConnectorAllOf) GetLogsOk() ([]PartnerintegrationDcLogsRelationship, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []PartnerintegrationDcLogsRelationship and assigns it to the Logs field.
func (o *PartnerintegrationDeviceConnectorAllOf) SetLogs(v []PartnerintegrationDcLogsRelationship) {
	o.Logs = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PartnerintegrationDeviceConnectorAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PartnerintegrationDeviceConnectorAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *PartnerintegrationDeviceConnectorAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o PartnerintegrationDeviceConnectorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.BuildStartTime != nil {
		toSerialize["BuildStartTime"] = o.BuildStartTime
	}
	if o.BuildStatus != nil {
		toSerialize["BuildStatus"] = o.BuildStatus
	}
	if o.ImageName != nil {
		toSerialize["ImageName"] = o.ImageName
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SrcBucket != nil {
		toSerialize["SrcBucket"] = o.SrcBucket
	}
	if o.SrcFileName != nil {
		toSerialize["SrcFileName"] = o.SrcFileName
	}
	if o.Logs != nil {
		toSerialize["Logs"] = o.Logs
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerintegrationDeviceConnectorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerintegrationDeviceConnectorAllOf := _PartnerintegrationDeviceConnectorAllOf{}

	if err = json.Unmarshal(bytes, &varPartnerintegrationDeviceConnectorAllOf); err == nil {
		*o = PartnerintegrationDeviceConnectorAllOf(varPartnerintegrationDeviceConnectorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "BuildStartTime")
		delete(additionalProperties, "BuildStatus")
		delete(additionalProperties, "ImageName")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SrcBucket")
		delete(additionalProperties, "SrcFileName")
		delete(additionalProperties, "Logs")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerintegrationDeviceConnectorAllOf struct {
	value *PartnerintegrationDeviceConnectorAllOf
	isSet bool
}

func (v NullablePartnerintegrationDeviceConnectorAllOf) Get() *PartnerintegrationDeviceConnectorAllOf {
	return v.value
}

func (v *NullablePartnerintegrationDeviceConnectorAllOf) Set(val *PartnerintegrationDeviceConnectorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationDeviceConnectorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationDeviceConnectorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationDeviceConnectorAllOf(val *PartnerintegrationDeviceConnectorAllOf) *NullablePartnerintegrationDeviceConnectorAllOf {
	return &NullablePartnerintegrationDeviceConnectorAllOf{value: val, isSet: true}
}

func (v NullablePartnerintegrationDeviceConnectorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationDeviceConnectorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
