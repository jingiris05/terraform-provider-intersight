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

// CloudTfcWorkspace Terraform workspace which represents running infrastructure managed by Terraform.
type CloudTfcWorkspace struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The agent pool associated with this workspace.
	AgentPoolId *string `json:"AgentPoolId,omitempty"`
	// Whether or not Terraform Cloud should automatically apply a successful Terraform plan.
	ApplyMethod *bool `json:"ApplyMethod,omitempty"`
	// Indicates where the Terraform cloud should execute the runs.
	ExecutionMode *string `json:"ExecutionMode,omitempty"`
	// The unique id for this workspace.
	Identity *string `json:"Identity,omitempty"`
	// The status of the last executed run in this workspace.
	LastRunStatus *string `json:"LastRunStatus,omitempty"`
	// The name of the workspace.
	Name                 *string                           `json:"Name,omitempty"`
	WorkspaceVariables   []CloudTfcWorkspaceVariables      `json:"WorkspaceVariables,omitempty"`
	Organization         *CloudTfcOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudTfcWorkspace CloudTfcWorkspace

// NewCloudTfcWorkspace instantiates a new CloudTfcWorkspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudTfcWorkspace(classId string, objectType string) *CloudTfcWorkspace {
	this := CloudTfcWorkspace{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudTfcWorkspaceWithDefaults instantiates a new CloudTfcWorkspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudTfcWorkspaceWithDefaults() *CloudTfcWorkspace {
	this := CloudTfcWorkspace{}
	var classId string = "cloud.TfcWorkspace"
	this.ClassId = classId
	var objectType string = "cloud.TfcWorkspace"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudTfcWorkspace) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspace) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudTfcWorkspace) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudTfcWorkspace) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspace) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudTfcWorkspace) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAgentPoolId returns the AgentPoolId field value if set, zero value otherwise.
func (o *CloudTfcWorkspace) GetAgentPoolId() string {
	if o == nil || o.AgentPoolId == nil {
		var ret string
		return ret
	}
	return *o.AgentPoolId
}

// GetAgentPoolIdOk returns a tuple with the AgentPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspace) GetAgentPoolIdOk() (*string, bool) {
	if o == nil || o.AgentPoolId == nil {
		return nil, false
	}
	return o.AgentPoolId, true
}

// HasAgentPoolId returns a boolean if a field has been set.
func (o *CloudTfcWorkspace) HasAgentPoolId() bool {
	if o != nil && o.AgentPoolId != nil {
		return true
	}

	return false
}

// SetAgentPoolId gets a reference to the given string and assigns it to the AgentPoolId field.
func (o *CloudTfcWorkspace) SetAgentPoolId(v string) {
	o.AgentPoolId = &v
}

// GetApplyMethod returns the ApplyMethod field value if set, zero value otherwise.
func (o *CloudTfcWorkspace) GetApplyMethod() bool {
	if o == nil || o.ApplyMethod == nil {
		var ret bool
		return ret
	}
	return *o.ApplyMethod
}

// GetApplyMethodOk returns a tuple with the ApplyMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspace) GetApplyMethodOk() (*bool, bool) {
	if o == nil || o.ApplyMethod == nil {
		return nil, false
	}
	return o.ApplyMethod, true
}

// HasApplyMethod returns a boolean if a field has been set.
func (o *CloudTfcWorkspace) HasApplyMethod() bool {
	if o != nil && o.ApplyMethod != nil {
		return true
	}

	return false
}

// SetApplyMethod gets a reference to the given bool and assigns it to the ApplyMethod field.
func (o *CloudTfcWorkspace) SetApplyMethod(v bool) {
	o.ApplyMethod = &v
}

// GetExecutionMode returns the ExecutionMode field value if set, zero value otherwise.
func (o *CloudTfcWorkspace) GetExecutionMode() string {
	if o == nil || o.ExecutionMode == nil {
		var ret string
		return ret
	}
	return *o.ExecutionMode
}

// GetExecutionModeOk returns a tuple with the ExecutionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspace) GetExecutionModeOk() (*string, bool) {
	if o == nil || o.ExecutionMode == nil {
		return nil, false
	}
	return o.ExecutionMode, true
}

// HasExecutionMode returns a boolean if a field has been set.
func (o *CloudTfcWorkspace) HasExecutionMode() bool {
	if o != nil && o.ExecutionMode != nil {
		return true
	}

	return false
}

// SetExecutionMode gets a reference to the given string and assigns it to the ExecutionMode field.
func (o *CloudTfcWorkspace) SetExecutionMode(v string) {
	o.ExecutionMode = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudTfcWorkspace) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspace) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudTfcWorkspace) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudTfcWorkspace) SetIdentity(v string) {
	o.Identity = &v
}

// GetLastRunStatus returns the LastRunStatus field value if set, zero value otherwise.
func (o *CloudTfcWorkspace) GetLastRunStatus() string {
	if o == nil || o.LastRunStatus == nil {
		var ret string
		return ret
	}
	return *o.LastRunStatus
}

// GetLastRunStatusOk returns a tuple with the LastRunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspace) GetLastRunStatusOk() (*string, bool) {
	if o == nil || o.LastRunStatus == nil {
		return nil, false
	}
	return o.LastRunStatus, true
}

// HasLastRunStatus returns a boolean if a field has been set.
func (o *CloudTfcWorkspace) HasLastRunStatus() bool {
	if o != nil && o.LastRunStatus != nil {
		return true
	}

	return false
}

// SetLastRunStatus gets a reference to the given string and assigns it to the LastRunStatus field.
func (o *CloudTfcWorkspace) SetLastRunStatus(v string) {
	o.LastRunStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudTfcWorkspace) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspace) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudTfcWorkspace) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudTfcWorkspace) SetName(v string) {
	o.Name = &v
}

// GetWorkspaceVariables returns the WorkspaceVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudTfcWorkspace) GetWorkspaceVariables() []CloudTfcWorkspaceVariables {
	if o == nil {
		var ret []CloudTfcWorkspaceVariables
		return ret
	}
	return o.WorkspaceVariables
}

// GetWorkspaceVariablesOk returns a tuple with the WorkspaceVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudTfcWorkspace) GetWorkspaceVariablesOk() ([]CloudTfcWorkspaceVariables, bool) {
	if o == nil || o.WorkspaceVariables == nil {
		return nil, false
	}
	return o.WorkspaceVariables, true
}

// HasWorkspaceVariables returns a boolean if a field has been set.
func (o *CloudTfcWorkspace) HasWorkspaceVariables() bool {
	if o != nil && o.WorkspaceVariables != nil {
		return true
	}

	return false
}

// SetWorkspaceVariables gets a reference to the given []CloudTfcWorkspaceVariables and assigns it to the WorkspaceVariables field.
func (o *CloudTfcWorkspace) SetWorkspaceVariables(v []CloudTfcWorkspaceVariables) {
	o.WorkspaceVariables = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CloudTfcWorkspace) GetOrganization() CloudTfcOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret CloudTfcOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspace) GetOrganizationOk() (*CloudTfcOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CloudTfcWorkspace) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given CloudTfcOrganizationRelationship and assigns it to the Organization field.
func (o *CloudTfcWorkspace) SetOrganization(v CloudTfcOrganizationRelationship) {
	o.Organization = &v
}

func (o CloudTfcWorkspace) MarshalJSON() ([]byte, error) {
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
	if o.AgentPoolId != nil {
		toSerialize["AgentPoolId"] = o.AgentPoolId
	}
	if o.ApplyMethod != nil {
		toSerialize["ApplyMethod"] = o.ApplyMethod
	}
	if o.ExecutionMode != nil {
		toSerialize["ExecutionMode"] = o.ExecutionMode
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.LastRunStatus != nil {
		toSerialize["LastRunStatus"] = o.LastRunStatus
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.WorkspaceVariables != nil {
		toSerialize["WorkspaceVariables"] = o.WorkspaceVariables
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudTfcWorkspace) UnmarshalJSON(bytes []byte) (err error) {
	type CloudTfcWorkspaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The agent pool associated with this workspace.
		AgentPoolId *string `json:"AgentPoolId,omitempty"`
		// Whether or not Terraform Cloud should automatically apply a successful Terraform plan.
		ApplyMethod *bool `json:"ApplyMethod,omitempty"`
		// Indicates where the Terraform cloud should execute the runs.
		ExecutionMode *string `json:"ExecutionMode,omitempty"`
		// The unique id for this workspace.
		Identity *string `json:"Identity,omitempty"`
		// The status of the last executed run in this workspace.
		LastRunStatus *string `json:"LastRunStatus,omitempty"`
		// The name of the workspace.
		Name               *string                           `json:"Name,omitempty"`
		WorkspaceVariables []CloudTfcWorkspaceVariables      `json:"WorkspaceVariables,omitempty"`
		Organization       *CloudTfcOrganizationRelationship `json:"Organization,omitempty"`
	}

	varCloudTfcWorkspaceWithoutEmbeddedStruct := CloudTfcWorkspaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudTfcWorkspaceWithoutEmbeddedStruct)
	if err == nil {
		varCloudTfcWorkspace := _CloudTfcWorkspace{}
		varCloudTfcWorkspace.ClassId = varCloudTfcWorkspaceWithoutEmbeddedStruct.ClassId
		varCloudTfcWorkspace.ObjectType = varCloudTfcWorkspaceWithoutEmbeddedStruct.ObjectType
		varCloudTfcWorkspace.AgentPoolId = varCloudTfcWorkspaceWithoutEmbeddedStruct.AgentPoolId
		varCloudTfcWorkspace.ApplyMethod = varCloudTfcWorkspaceWithoutEmbeddedStruct.ApplyMethod
		varCloudTfcWorkspace.ExecutionMode = varCloudTfcWorkspaceWithoutEmbeddedStruct.ExecutionMode
		varCloudTfcWorkspace.Identity = varCloudTfcWorkspaceWithoutEmbeddedStruct.Identity
		varCloudTfcWorkspace.LastRunStatus = varCloudTfcWorkspaceWithoutEmbeddedStruct.LastRunStatus
		varCloudTfcWorkspace.Name = varCloudTfcWorkspaceWithoutEmbeddedStruct.Name
		varCloudTfcWorkspace.WorkspaceVariables = varCloudTfcWorkspaceWithoutEmbeddedStruct.WorkspaceVariables
		varCloudTfcWorkspace.Organization = varCloudTfcWorkspaceWithoutEmbeddedStruct.Organization
		*o = CloudTfcWorkspace(varCloudTfcWorkspace)
	} else {
		return err
	}

	varCloudTfcWorkspace := _CloudTfcWorkspace{}

	err = json.Unmarshal(bytes, &varCloudTfcWorkspace)
	if err == nil {
		o.MoBaseMo = varCloudTfcWorkspace.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AgentPoolId")
		delete(additionalProperties, "ApplyMethod")
		delete(additionalProperties, "ExecutionMode")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "LastRunStatus")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "WorkspaceVariables")
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

type NullableCloudTfcWorkspace struct {
	value *CloudTfcWorkspace
	isSet bool
}

func (v NullableCloudTfcWorkspace) Get() *CloudTfcWorkspace {
	return v.value
}

func (v *NullableCloudTfcWorkspace) Set(val *CloudTfcWorkspace) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudTfcWorkspace) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudTfcWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudTfcWorkspace(val *CloudTfcWorkspace) *NullableCloudTfcWorkspace {
	return &NullableCloudTfcWorkspace{value: val, isSet: true}
}

func (v NullableCloudTfcWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudTfcWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
