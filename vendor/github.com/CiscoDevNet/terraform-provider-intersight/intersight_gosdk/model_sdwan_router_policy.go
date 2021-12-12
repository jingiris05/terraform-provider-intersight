/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// SdwanRouterPolicy A policy specifying SD-WAN router details.
type SdwanRouterPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Scale of the SD-WAN router virtual machine deployment. * `Typical` - Typical deployment configuration with 4 vCPUs and 4GB RAM. * `Minimal` - Minimal deployment configuration with 2 vCPUs and 4GB RAM.
	DeploymentSize *string `json:"DeploymentSize,omitempty"`
	// Number of WAN connections across the SD-WAN site.
	WanCount *int64 `json:"WanCount,omitempty"`
	// Defines if the WAN networks are singly or dually terminated. Dually terminated WANs are configured on all the SD-WAN routers. Singly terminated WANs are configured only on one of the SD-WAN routers. * `Single` - Singly terminated WANs ar evenly distributed across SD-WAN router nodes. A given WAN connection is available only on one of the router nodes. * `Dual` - Dually terminated WANs are configured on all the SD-WAN routers. A given WAN connection is available on multiple router nodes.
	WanTerminationType *string                               `json:"WanTerminationType,omitempty"`
	Organization       *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to sdwanProfile resources.
	Profiles             []SdwanProfileRelationship                 `json:"Profiles,omitempty"`
	SolutionImage        *SoftwareSolutionDistributableRelationship `json:"SolutionImage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdwanRouterPolicy SdwanRouterPolicy

// NewSdwanRouterPolicy instantiates a new SdwanRouterPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdwanRouterPolicy(classId string, objectType string) *SdwanRouterPolicy {
	this := SdwanRouterPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var deploymentSize string = "Typical"
	this.DeploymentSize = &deploymentSize
	var wanCount int64 = 2
	this.WanCount = &wanCount
	var wanTerminationType string = "Single"
	this.WanTerminationType = &wanTerminationType
	return &this
}

// NewSdwanRouterPolicyWithDefaults instantiates a new SdwanRouterPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdwanRouterPolicyWithDefaults() *SdwanRouterPolicy {
	this := SdwanRouterPolicy{}
	var classId string = "sdwan.RouterPolicy"
	this.ClassId = classId
	var objectType string = "sdwan.RouterPolicy"
	this.ObjectType = objectType
	var deploymentSize string = "Typical"
	this.DeploymentSize = &deploymentSize
	var wanCount int64 = 2
	this.WanCount = &wanCount
	var wanTerminationType string = "Single"
	this.WanTerminationType = &wanTerminationType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SdwanRouterPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SdwanRouterPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SdwanRouterPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SdwanRouterPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SdwanRouterPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SdwanRouterPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeploymentSize returns the DeploymentSize field value if set, zero value otherwise.
func (o *SdwanRouterPolicy) GetDeploymentSize() string {
	if o == nil || o.DeploymentSize == nil {
		var ret string
		return ret
	}
	return *o.DeploymentSize
}

// GetDeploymentSizeOk returns a tuple with the DeploymentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterPolicy) GetDeploymentSizeOk() (*string, bool) {
	if o == nil || o.DeploymentSize == nil {
		return nil, false
	}
	return o.DeploymentSize, true
}

// HasDeploymentSize returns a boolean if a field has been set.
func (o *SdwanRouterPolicy) HasDeploymentSize() bool {
	if o != nil && o.DeploymentSize != nil {
		return true
	}

	return false
}

// SetDeploymentSize gets a reference to the given string and assigns it to the DeploymentSize field.
func (o *SdwanRouterPolicy) SetDeploymentSize(v string) {
	o.DeploymentSize = &v
}

// GetWanCount returns the WanCount field value if set, zero value otherwise.
func (o *SdwanRouterPolicy) GetWanCount() int64 {
	if o == nil || o.WanCount == nil {
		var ret int64
		return ret
	}
	return *o.WanCount
}

// GetWanCountOk returns a tuple with the WanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterPolicy) GetWanCountOk() (*int64, bool) {
	if o == nil || o.WanCount == nil {
		return nil, false
	}
	return o.WanCount, true
}

// HasWanCount returns a boolean if a field has been set.
func (o *SdwanRouterPolicy) HasWanCount() bool {
	if o != nil && o.WanCount != nil {
		return true
	}

	return false
}

// SetWanCount gets a reference to the given int64 and assigns it to the WanCount field.
func (o *SdwanRouterPolicy) SetWanCount(v int64) {
	o.WanCount = &v
}

// GetWanTerminationType returns the WanTerminationType field value if set, zero value otherwise.
func (o *SdwanRouterPolicy) GetWanTerminationType() string {
	if o == nil || o.WanTerminationType == nil {
		var ret string
		return ret
	}
	return *o.WanTerminationType
}

// GetWanTerminationTypeOk returns a tuple with the WanTerminationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterPolicy) GetWanTerminationTypeOk() (*string, bool) {
	if o == nil || o.WanTerminationType == nil {
		return nil, false
	}
	return o.WanTerminationType, true
}

// HasWanTerminationType returns a boolean if a field has been set.
func (o *SdwanRouterPolicy) HasWanTerminationType() bool {
	if o != nil && o.WanTerminationType != nil {
		return true
	}

	return false
}

// SetWanTerminationType gets a reference to the given string and assigns it to the WanTerminationType field.
func (o *SdwanRouterPolicy) SetWanTerminationType(v string) {
	o.WanTerminationType = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SdwanRouterPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SdwanRouterPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SdwanRouterPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SdwanRouterPolicy) GetProfiles() []SdwanProfileRelationship {
	if o == nil {
		var ret []SdwanProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SdwanRouterPolicy) GetProfilesOk() (*[]SdwanProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SdwanRouterPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []SdwanProfileRelationship and assigns it to the Profiles field.
func (o *SdwanRouterPolicy) SetProfiles(v []SdwanProfileRelationship) {
	o.Profiles = v
}

// GetSolutionImage returns the SolutionImage field value if set, zero value otherwise.
func (o *SdwanRouterPolicy) GetSolutionImage() SoftwareSolutionDistributableRelationship {
	if o == nil || o.SolutionImage == nil {
		var ret SoftwareSolutionDistributableRelationship
		return ret
	}
	return *o.SolutionImage
}

// GetSolutionImageOk returns a tuple with the SolutionImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterPolicy) GetSolutionImageOk() (*SoftwareSolutionDistributableRelationship, bool) {
	if o == nil || o.SolutionImage == nil {
		return nil, false
	}
	return o.SolutionImage, true
}

// HasSolutionImage returns a boolean if a field has been set.
func (o *SdwanRouterPolicy) HasSolutionImage() bool {
	if o != nil && o.SolutionImage != nil {
		return true
	}

	return false
}

// SetSolutionImage gets a reference to the given SoftwareSolutionDistributableRelationship and assigns it to the SolutionImage field.
func (o *SdwanRouterPolicy) SetSolutionImage(v SoftwareSolutionDistributableRelationship) {
	o.SolutionImage = &v
}

func (o SdwanRouterPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeploymentSize != nil {
		toSerialize["DeploymentSize"] = o.DeploymentSize
	}
	if o.WanCount != nil {
		toSerialize["WanCount"] = o.WanCount
	}
	if o.WanTerminationType != nil {
		toSerialize["WanTerminationType"] = o.WanTerminationType
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	if o.SolutionImage != nil {
		toSerialize["SolutionImage"] = o.SolutionImage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SdwanRouterPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type SdwanRouterPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Scale of the SD-WAN router virtual machine deployment. * `Typical` - Typical deployment configuration with 4 vCPUs and 4GB RAM. * `Minimal` - Minimal deployment configuration with 2 vCPUs and 4GB RAM.
		DeploymentSize *string `json:"DeploymentSize,omitempty"`
		// Number of WAN connections across the SD-WAN site.
		WanCount *int64 `json:"WanCount,omitempty"`
		// Defines if the WAN networks are singly or dually terminated. Dually terminated WANs are configured on all the SD-WAN routers. Singly terminated WANs are configured only on one of the SD-WAN routers. * `Single` - Singly terminated WANs ar evenly distributed across SD-WAN router nodes. A given WAN connection is available only on one of the router nodes. * `Dual` - Dually terminated WANs are configured on all the SD-WAN routers. A given WAN connection is available on multiple router nodes.
		WanTerminationType *string                               `json:"WanTerminationType,omitempty"`
		Organization       *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to sdwanProfile resources.
		Profiles      []SdwanProfileRelationship                 `json:"Profiles,omitempty"`
		SolutionImage *SoftwareSolutionDistributableRelationship `json:"SolutionImage,omitempty"`
	}

	varSdwanRouterPolicyWithoutEmbeddedStruct := SdwanRouterPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSdwanRouterPolicyWithoutEmbeddedStruct)
	if err == nil {
		varSdwanRouterPolicy := _SdwanRouterPolicy{}
		varSdwanRouterPolicy.ClassId = varSdwanRouterPolicyWithoutEmbeddedStruct.ClassId
		varSdwanRouterPolicy.ObjectType = varSdwanRouterPolicyWithoutEmbeddedStruct.ObjectType
		varSdwanRouterPolicy.DeploymentSize = varSdwanRouterPolicyWithoutEmbeddedStruct.DeploymentSize
		varSdwanRouterPolicy.WanCount = varSdwanRouterPolicyWithoutEmbeddedStruct.WanCount
		varSdwanRouterPolicy.WanTerminationType = varSdwanRouterPolicyWithoutEmbeddedStruct.WanTerminationType
		varSdwanRouterPolicy.Organization = varSdwanRouterPolicyWithoutEmbeddedStruct.Organization
		varSdwanRouterPolicy.Profiles = varSdwanRouterPolicyWithoutEmbeddedStruct.Profiles
		varSdwanRouterPolicy.SolutionImage = varSdwanRouterPolicyWithoutEmbeddedStruct.SolutionImage
		*o = SdwanRouterPolicy(varSdwanRouterPolicy)
	} else {
		return err
	}

	varSdwanRouterPolicy := _SdwanRouterPolicy{}

	err = json.Unmarshal(bytes, &varSdwanRouterPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varSdwanRouterPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeploymentSize")
		delete(additionalProperties, "WanCount")
		delete(additionalProperties, "WanTerminationType")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		delete(additionalProperties, "SolutionImage")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableSdwanRouterPolicy struct {
	value *SdwanRouterPolicy
	isSet bool
}

func (v NullableSdwanRouterPolicy) Get() *SdwanRouterPolicy {
	return v.value
}

func (v *NullableSdwanRouterPolicy) Set(val *SdwanRouterPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSdwanRouterPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSdwanRouterPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdwanRouterPolicy(val *SdwanRouterPolicy) *NullableSdwanRouterPolicy {
	return &NullableSdwanRouterPolicy{value: val, isSet: true}
}

func (v NullableSdwanRouterPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdwanRouterPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
