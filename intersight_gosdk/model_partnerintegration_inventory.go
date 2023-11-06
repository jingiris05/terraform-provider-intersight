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

// PartnerintegrationInventory Inventory Collection object that acts as an aggregator object for the underlying model and ETL objects.
type PartnerintegrationInventory struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action to be performed on the inventory collection. * `None` - Default Value of the action, i.e. do nothing. * `Build` - Build the inventory service image. * `Deploy` - Deploy the inventory service on the appliance.
	Action *string `json:"Action,omitempty"`
	// Addtional flags to control build action.
	BuildFlags *string `json:"BuildFlags,omitempty"`
	// Time when build was triggered.
	BuildStartTime *string `json:"BuildStartTime,omitempty"`
	// Status of build for inventory collection. * `None` - Default value of the status. i.e. done nothing. * `BackendInProgress` - The backend build is in progress. * `BackendFailed` - The backend build has failed. * `DockerInProgress` - The docker build is in progress. * `DockerFailed` - The docker build has failed. * `UiInProgress` - The UI build is in progress. * `UiFailed` - The inventory UI build has failed. * `ApidocsInProgress` - The apidocs build is in progress. * `ApidocsFailed` - The apidocs build has failed. * `Completed` - The operation completed successfully.
	BuildStatus *string `json:"BuildStatus,omitempty"`
	// Time when deploy was triggered.
	DeployStartTime *string `json:"DeployStartTime,omitempty"`
	// Status of deployment of the inventory microservice. * `None` - Default value of the status. i.e. done nothing. * `Completed` - The operation completed successfully. * `Failed` - The deploy operation failed.
	DeployStatus *string `json:"DeployStatus,omitempty"`
	// Name of the docker image that is built.
	ImageName *string `json:"ImageName,omitempty"`
	// Name of the inventory collection.
	Name *string `json:"Name,omitempty"`
	// Link to the generated v3 python SDK.
	PythonSdkUrl *string `json:"PythonSdkUrl,omitempty"`
	// An array of relationships to partnerintegrationDocIssues resources.
	DocIssues []PartnerintegrationDocIssuesRelationship `json:"DocIssues,omitempty"`
	// An array of relationships to partnerintegrationEtl resources.
	Etls []PartnerintegrationEtlRelationship `json:"Etls,omitempty"`
	// An array of relationships to partnerintegrationLogs resources.
	Logs []PartnerintegrationLogsRelationship `json:"Logs,omitempty"`
	// An array of relationships to partnerintegrationModel resources.
	Models               []PartnerintegrationModelRelationship `json:"Models,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerintegrationInventory PartnerintegrationInventory

// NewPartnerintegrationInventory instantiates a new PartnerintegrationInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerintegrationInventory(classId string, objectType string) *PartnerintegrationInventory {
	this := PartnerintegrationInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// NewPartnerintegrationInventoryWithDefaults instantiates a new PartnerintegrationInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerintegrationInventoryWithDefaults() *PartnerintegrationInventory {
	this := PartnerintegrationInventory{}
	var classId string = "partnerintegration.Inventory"
	this.ClassId = classId
	var objectType string = "partnerintegration.Inventory"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *PartnerintegrationInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PartnerintegrationInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PartnerintegrationInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PartnerintegrationInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PartnerintegrationInventory) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PartnerintegrationInventory) SetAction(v string) {
	o.Action = &v
}

// GetBuildFlags returns the BuildFlags field value if set, zero value otherwise.
func (o *PartnerintegrationInventory) GetBuildFlags() string {
	if o == nil || o.BuildFlags == nil {
		var ret string
		return ret
	}
	return *o.BuildFlags
}

// GetBuildFlagsOk returns a tuple with the BuildFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetBuildFlagsOk() (*string, bool) {
	if o == nil || o.BuildFlags == nil {
		return nil, false
	}
	return o.BuildFlags, true
}

// HasBuildFlags returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasBuildFlags() bool {
	if o != nil && o.BuildFlags != nil {
		return true
	}

	return false
}

// SetBuildFlags gets a reference to the given string and assigns it to the BuildFlags field.
func (o *PartnerintegrationInventory) SetBuildFlags(v string) {
	o.BuildFlags = &v
}

// GetBuildStartTime returns the BuildStartTime field value if set, zero value otherwise.
func (o *PartnerintegrationInventory) GetBuildStartTime() string {
	if o == nil || o.BuildStartTime == nil {
		var ret string
		return ret
	}
	return *o.BuildStartTime
}

// GetBuildStartTimeOk returns a tuple with the BuildStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetBuildStartTimeOk() (*string, bool) {
	if o == nil || o.BuildStartTime == nil {
		return nil, false
	}
	return o.BuildStartTime, true
}

// HasBuildStartTime returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasBuildStartTime() bool {
	if o != nil && o.BuildStartTime != nil {
		return true
	}

	return false
}

// SetBuildStartTime gets a reference to the given string and assigns it to the BuildStartTime field.
func (o *PartnerintegrationInventory) SetBuildStartTime(v string) {
	o.BuildStartTime = &v
}

// GetBuildStatus returns the BuildStatus field value if set, zero value otherwise.
func (o *PartnerintegrationInventory) GetBuildStatus() string {
	if o == nil || o.BuildStatus == nil {
		var ret string
		return ret
	}
	return *o.BuildStatus
}

// GetBuildStatusOk returns a tuple with the BuildStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetBuildStatusOk() (*string, bool) {
	if o == nil || o.BuildStatus == nil {
		return nil, false
	}
	return o.BuildStatus, true
}

// HasBuildStatus returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasBuildStatus() bool {
	if o != nil && o.BuildStatus != nil {
		return true
	}

	return false
}

// SetBuildStatus gets a reference to the given string and assigns it to the BuildStatus field.
func (o *PartnerintegrationInventory) SetBuildStatus(v string) {
	o.BuildStatus = &v
}

// GetDeployStartTime returns the DeployStartTime field value if set, zero value otherwise.
func (o *PartnerintegrationInventory) GetDeployStartTime() string {
	if o == nil || o.DeployStartTime == nil {
		var ret string
		return ret
	}
	return *o.DeployStartTime
}

// GetDeployStartTimeOk returns a tuple with the DeployStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetDeployStartTimeOk() (*string, bool) {
	if o == nil || o.DeployStartTime == nil {
		return nil, false
	}
	return o.DeployStartTime, true
}

// HasDeployStartTime returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasDeployStartTime() bool {
	if o != nil && o.DeployStartTime != nil {
		return true
	}

	return false
}

// SetDeployStartTime gets a reference to the given string and assigns it to the DeployStartTime field.
func (o *PartnerintegrationInventory) SetDeployStartTime(v string) {
	o.DeployStartTime = &v
}

// GetDeployStatus returns the DeployStatus field value if set, zero value otherwise.
func (o *PartnerintegrationInventory) GetDeployStatus() string {
	if o == nil || o.DeployStatus == nil {
		var ret string
		return ret
	}
	return *o.DeployStatus
}

// GetDeployStatusOk returns a tuple with the DeployStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetDeployStatusOk() (*string, bool) {
	if o == nil || o.DeployStatus == nil {
		return nil, false
	}
	return o.DeployStatus, true
}

// HasDeployStatus returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasDeployStatus() bool {
	if o != nil && o.DeployStatus != nil {
		return true
	}

	return false
}

// SetDeployStatus gets a reference to the given string and assigns it to the DeployStatus field.
func (o *PartnerintegrationInventory) SetDeployStatus(v string) {
	o.DeployStatus = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *PartnerintegrationInventory) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *PartnerintegrationInventory) SetImageName(v string) {
	o.ImageName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartnerintegrationInventory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartnerintegrationInventory) SetName(v string) {
	o.Name = &v
}

// GetPythonSdkUrl returns the PythonSdkUrl field value if set, zero value otherwise.
func (o *PartnerintegrationInventory) GetPythonSdkUrl() string {
	if o == nil || o.PythonSdkUrl == nil {
		var ret string
		return ret
	}
	return *o.PythonSdkUrl
}

// GetPythonSdkUrlOk returns a tuple with the PythonSdkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetPythonSdkUrlOk() (*string, bool) {
	if o == nil || o.PythonSdkUrl == nil {
		return nil, false
	}
	return o.PythonSdkUrl, true
}

// HasPythonSdkUrl returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasPythonSdkUrl() bool {
	if o != nil && o.PythonSdkUrl != nil {
		return true
	}

	return false
}

// SetPythonSdkUrl gets a reference to the given string and assigns it to the PythonSdkUrl field.
func (o *PartnerintegrationInventory) SetPythonSdkUrl(v string) {
	o.PythonSdkUrl = &v
}

// GetDocIssues returns the DocIssues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationInventory) GetDocIssues() []PartnerintegrationDocIssuesRelationship {
	if o == nil {
		var ret []PartnerintegrationDocIssuesRelationship
		return ret
	}
	return o.DocIssues
}

// GetDocIssuesOk returns a tuple with the DocIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationInventory) GetDocIssuesOk() ([]PartnerintegrationDocIssuesRelationship, bool) {
	if o == nil || o.DocIssues == nil {
		return nil, false
	}
	return o.DocIssues, true
}

// HasDocIssues returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasDocIssues() bool {
	if o != nil && o.DocIssues != nil {
		return true
	}

	return false
}

// SetDocIssues gets a reference to the given []PartnerintegrationDocIssuesRelationship and assigns it to the DocIssues field.
func (o *PartnerintegrationInventory) SetDocIssues(v []PartnerintegrationDocIssuesRelationship) {
	o.DocIssues = v
}

// GetEtls returns the Etls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationInventory) GetEtls() []PartnerintegrationEtlRelationship {
	if o == nil {
		var ret []PartnerintegrationEtlRelationship
		return ret
	}
	return o.Etls
}

// GetEtlsOk returns a tuple with the Etls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationInventory) GetEtlsOk() ([]PartnerintegrationEtlRelationship, bool) {
	if o == nil || o.Etls == nil {
		return nil, false
	}
	return o.Etls, true
}

// HasEtls returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasEtls() bool {
	if o != nil && o.Etls != nil {
		return true
	}

	return false
}

// SetEtls gets a reference to the given []PartnerintegrationEtlRelationship and assigns it to the Etls field.
func (o *PartnerintegrationInventory) SetEtls(v []PartnerintegrationEtlRelationship) {
	o.Etls = v
}

// GetLogs returns the Logs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationInventory) GetLogs() []PartnerintegrationLogsRelationship {
	if o == nil {
		var ret []PartnerintegrationLogsRelationship
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationInventory) GetLogsOk() ([]PartnerintegrationLogsRelationship, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []PartnerintegrationLogsRelationship and assigns it to the Logs field.
func (o *PartnerintegrationInventory) SetLogs(v []PartnerintegrationLogsRelationship) {
	o.Logs = v
}

// GetModels returns the Models field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationInventory) GetModels() []PartnerintegrationModelRelationship {
	if o == nil {
		var ret []PartnerintegrationModelRelationship
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationInventory) GetModelsOk() ([]PartnerintegrationModelRelationship, bool) {
	if o == nil || o.Models == nil {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasModels() bool {
	if o != nil && o.Models != nil {
		return true
	}

	return false
}

// SetModels gets a reference to the given []PartnerintegrationModelRelationship and assigns it to the Models field.
func (o *PartnerintegrationInventory) SetModels(v []PartnerintegrationModelRelationship) {
	o.Models = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PartnerintegrationInventory) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationInventory) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PartnerintegrationInventory) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *PartnerintegrationInventory) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o PartnerintegrationInventory) MarshalJSON() ([]byte, error) {
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
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.BuildFlags != nil {
		toSerialize["BuildFlags"] = o.BuildFlags
	}
	if o.BuildStartTime != nil {
		toSerialize["BuildStartTime"] = o.BuildStartTime
	}
	if o.BuildStatus != nil {
		toSerialize["BuildStatus"] = o.BuildStatus
	}
	if o.DeployStartTime != nil {
		toSerialize["DeployStartTime"] = o.DeployStartTime
	}
	if o.DeployStatus != nil {
		toSerialize["DeployStatus"] = o.DeployStatus
	}
	if o.ImageName != nil {
		toSerialize["ImageName"] = o.ImageName
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PythonSdkUrl != nil {
		toSerialize["PythonSdkUrl"] = o.PythonSdkUrl
	}
	if o.DocIssues != nil {
		toSerialize["DocIssues"] = o.DocIssues
	}
	if o.Etls != nil {
		toSerialize["Etls"] = o.Etls
	}
	if o.Logs != nil {
		toSerialize["Logs"] = o.Logs
	}
	if o.Models != nil {
		toSerialize["Models"] = o.Models
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerintegrationInventory) UnmarshalJSON(bytes []byte) (err error) {
	type PartnerintegrationInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Action to be performed on the inventory collection. * `None` - Default Value of the action, i.e. do nothing. * `Build` - Build the inventory service image. * `Deploy` - Deploy the inventory service on the appliance.
		Action *string `json:"Action,omitempty"`
		// Addtional flags to control build action.
		BuildFlags *string `json:"BuildFlags,omitempty"`
		// Time when build was triggered.
		BuildStartTime *string `json:"BuildStartTime,omitempty"`
		// Status of build for inventory collection. * `None` - Default value of the status. i.e. done nothing. * `BackendInProgress` - The backend build is in progress. * `BackendFailed` - The backend build has failed. * `DockerInProgress` - The docker build is in progress. * `DockerFailed` - The docker build has failed. * `UiInProgress` - The UI build is in progress. * `UiFailed` - The inventory UI build has failed. * `ApidocsInProgress` - The apidocs build is in progress. * `ApidocsFailed` - The apidocs build has failed. * `Completed` - The operation completed successfully.
		BuildStatus *string `json:"BuildStatus,omitempty"`
		// Time when deploy was triggered.
		DeployStartTime *string `json:"DeployStartTime,omitempty"`
		// Status of deployment of the inventory microservice. * `None` - Default value of the status. i.e. done nothing. * `Completed` - The operation completed successfully. * `Failed` - The deploy operation failed.
		DeployStatus *string `json:"DeployStatus,omitempty"`
		// Name of the docker image that is built.
		ImageName *string `json:"ImageName,omitempty"`
		// Name of the inventory collection.
		Name *string `json:"Name,omitempty"`
		// Link to the generated v3 python SDK.
		PythonSdkUrl *string `json:"PythonSdkUrl,omitempty"`
		// An array of relationships to partnerintegrationDocIssues resources.
		DocIssues []PartnerintegrationDocIssuesRelationship `json:"DocIssues,omitempty"`
		// An array of relationships to partnerintegrationEtl resources.
		Etls []PartnerintegrationEtlRelationship `json:"Etls,omitempty"`
		// An array of relationships to partnerintegrationLogs resources.
		Logs []PartnerintegrationLogsRelationship `json:"Logs,omitempty"`
		// An array of relationships to partnerintegrationModel resources.
		Models       []PartnerintegrationModelRelationship `json:"Models,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varPartnerintegrationInventoryWithoutEmbeddedStruct := PartnerintegrationInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPartnerintegrationInventoryWithoutEmbeddedStruct)
	if err == nil {
		varPartnerintegrationInventory := _PartnerintegrationInventory{}
		varPartnerintegrationInventory.ClassId = varPartnerintegrationInventoryWithoutEmbeddedStruct.ClassId
		varPartnerintegrationInventory.ObjectType = varPartnerintegrationInventoryWithoutEmbeddedStruct.ObjectType
		varPartnerintegrationInventory.Action = varPartnerintegrationInventoryWithoutEmbeddedStruct.Action
		varPartnerintegrationInventory.BuildFlags = varPartnerintegrationInventoryWithoutEmbeddedStruct.BuildFlags
		varPartnerintegrationInventory.BuildStartTime = varPartnerintegrationInventoryWithoutEmbeddedStruct.BuildStartTime
		varPartnerintegrationInventory.BuildStatus = varPartnerintegrationInventoryWithoutEmbeddedStruct.BuildStatus
		varPartnerintegrationInventory.DeployStartTime = varPartnerintegrationInventoryWithoutEmbeddedStruct.DeployStartTime
		varPartnerintegrationInventory.DeployStatus = varPartnerintegrationInventoryWithoutEmbeddedStruct.DeployStatus
		varPartnerintegrationInventory.ImageName = varPartnerintegrationInventoryWithoutEmbeddedStruct.ImageName
		varPartnerintegrationInventory.Name = varPartnerintegrationInventoryWithoutEmbeddedStruct.Name
		varPartnerintegrationInventory.PythonSdkUrl = varPartnerintegrationInventoryWithoutEmbeddedStruct.PythonSdkUrl
		varPartnerintegrationInventory.DocIssues = varPartnerintegrationInventoryWithoutEmbeddedStruct.DocIssues
		varPartnerintegrationInventory.Etls = varPartnerintegrationInventoryWithoutEmbeddedStruct.Etls
		varPartnerintegrationInventory.Logs = varPartnerintegrationInventoryWithoutEmbeddedStruct.Logs
		varPartnerintegrationInventory.Models = varPartnerintegrationInventoryWithoutEmbeddedStruct.Models
		varPartnerintegrationInventory.Organization = varPartnerintegrationInventoryWithoutEmbeddedStruct.Organization
		*o = PartnerintegrationInventory(varPartnerintegrationInventory)
	} else {
		return err
	}

	varPartnerintegrationInventory := _PartnerintegrationInventory{}

	err = json.Unmarshal(bytes, &varPartnerintegrationInventory)
	if err == nil {
		o.MoBaseMo = varPartnerintegrationInventory.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "BuildFlags")
		delete(additionalProperties, "BuildStartTime")
		delete(additionalProperties, "BuildStatus")
		delete(additionalProperties, "DeployStartTime")
		delete(additionalProperties, "DeployStatus")
		delete(additionalProperties, "ImageName")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PythonSdkUrl")
		delete(additionalProperties, "DocIssues")
		delete(additionalProperties, "Etls")
		delete(additionalProperties, "Logs")
		delete(additionalProperties, "Models")
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

type NullablePartnerintegrationInventory struct {
	value *PartnerintegrationInventory
	isSet bool
}

func (v NullablePartnerintegrationInventory) Get() *PartnerintegrationInventory {
	return v.value
}

func (v *NullablePartnerintegrationInventory) Set(val *PartnerintegrationInventory) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationInventory) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationInventory(val *PartnerintegrationInventory) *NullablePartnerintegrationInventory {
	return &NullablePartnerintegrationInventory{value: val, isSet: true}
}

func (v NullablePartnerintegrationInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
