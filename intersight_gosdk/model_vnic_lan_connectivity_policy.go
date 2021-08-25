/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-24T21:18:00Z.
 *
 * API version: 1.0.9-4404
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VnicLanConnectivityPolicy A LAN Connectivity Policy determines the network resources and the connections between the server and the LAN on the network. This policy uses Consistent Device Naming to configure the vNIC. You can configure a usNIC or VMQ connection for the vNIC to improve network performance.
type VnicLanConnectivityPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enabling AzureStack-Host QoS on an adapter allows the user to carve out traffic classes for RDMA traffic which ensures that a desired portion of the bandwidth is allocated to it.
	AzureQosEnabled *bool `json:"AzureQosEnabled,omitempty"`
	// Allocation Type of iSCSI Qualified Name - Static/Pool/None. * `None` - Type indicates that there is no IQN associated to an interface. * `Static` - Type represents that static IQN is associated to an interface. * `Pool` - Type indicates that IQN value is sourced from an associated pool.
	IqnAllocationType *string `json:"IqnAllocationType,omitempty"`
	// The mode used for placement of vNICs on network adapters. It can either be Auto or Custom. * `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system.
	PlacementMode *string `json:"PlacementMode,omitempty"`
	// User provided static iSCSI Qualified Name (IQN) for use as initiator identifiers by iSCSI vNICs in a Fabric Interconnect domain.
	StaticIqnName *string `json:"StaticIqnName,omitempty"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	TargetPlatform *string `json:"TargetPlatform,omitempty"`
	// An array of relationships to vnicEthIf resources.
	EthIfs       []VnicEthIfRelationship               `json:"EthIfs,omitempty"`
	IqnPool      *IqnpoolPoolRelationship              `json:"IqnPool,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicLanConnectivityPolicy VnicLanConnectivityPolicy

// NewVnicLanConnectivityPolicy instantiates a new VnicLanConnectivityPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicLanConnectivityPolicy(classId string, objectType string) *VnicLanConnectivityPolicy {
	this := VnicLanConnectivityPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var azureQosEnabled bool = false
	this.AzureQosEnabled = &azureQosEnabled
	var iqnAllocationType string = "None"
	this.IqnAllocationType = &iqnAllocationType
	var placementMode string = "custom"
	this.PlacementMode = &placementMode
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewVnicLanConnectivityPolicyWithDefaults instantiates a new VnicLanConnectivityPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicLanConnectivityPolicyWithDefaults() *VnicLanConnectivityPolicy {
	this := VnicLanConnectivityPolicy{}
	var classId string = "vnic.LanConnectivityPolicy"
	this.ClassId = classId
	var objectType string = "vnic.LanConnectivityPolicy"
	this.ObjectType = objectType
	var azureQosEnabled bool = false
	this.AzureQosEnabled = &azureQosEnabled
	var iqnAllocationType string = "None"
	this.IqnAllocationType = &iqnAllocationType
	var placementMode string = "custom"
	this.PlacementMode = &placementMode
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicLanConnectivityPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicLanConnectivityPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicLanConnectivityPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicLanConnectivityPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAzureQosEnabled returns the AzureQosEnabled field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicy) GetAzureQosEnabled() bool {
	if o == nil || o.AzureQosEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AzureQosEnabled
}

// GetAzureQosEnabledOk returns a tuple with the AzureQosEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetAzureQosEnabledOk() (*bool, bool) {
	if o == nil || o.AzureQosEnabled == nil {
		return nil, false
	}
	return o.AzureQosEnabled, true
}

// HasAzureQosEnabled returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasAzureQosEnabled() bool {
	if o != nil && o.AzureQosEnabled != nil {
		return true
	}

	return false
}

// SetAzureQosEnabled gets a reference to the given bool and assigns it to the AzureQosEnabled field.
func (o *VnicLanConnectivityPolicy) SetAzureQosEnabled(v bool) {
	o.AzureQosEnabled = &v
}

// GetIqnAllocationType returns the IqnAllocationType field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicy) GetIqnAllocationType() string {
	if o == nil || o.IqnAllocationType == nil {
		var ret string
		return ret
	}
	return *o.IqnAllocationType
}

// GetIqnAllocationTypeOk returns a tuple with the IqnAllocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetIqnAllocationTypeOk() (*string, bool) {
	if o == nil || o.IqnAllocationType == nil {
		return nil, false
	}
	return o.IqnAllocationType, true
}

// HasIqnAllocationType returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasIqnAllocationType() bool {
	if o != nil && o.IqnAllocationType != nil {
		return true
	}

	return false
}

// SetIqnAllocationType gets a reference to the given string and assigns it to the IqnAllocationType field.
func (o *VnicLanConnectivityPolicy) SetIqnAllocationType(v string) {
	o.IqnAllocationType = &v
}

// GetPlacementMode returns the PlacementMode field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicy) GetPlacementMode() string {
	if o == nil || o.PlacementMode == nil {
		var ret string
		return ret
	}
	return *o.PlacementMode
}

// GetPlacementModeOk returns a tuple with the PlacementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetPlacementModeOk() (*string, bool) {
	if o == nil || o.PlacementMode == nil {
		return nil, false
	}
	return o.PlacementMode, true
}

// HasPlacementMode returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasPlacementMode() bool {
	if o != nil && o.PlacementMode != nil {
		return true
	}

	return false
}

// SetPlacementMode gets a reference to the given string and assigns it to the PlacementMode field.
func (o *VnicLanConnectivityPolicy) SetPlacementMode(v string) {
	o.PlacementMode = &v
}

// GetStaticIqnName returns the StaticIqnName field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicy) GetStaticIqnName() string {
	if o == nil || o.StaticIqnName == nil {
		var ret string
		return ret
	}
	return *o.StaticIqnName
}

// GetStaticIqnNameOk returns a tuple with the StaticIqnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetStaticIqnNameOk() (*string, bool) {
	if o == nil || o.StaticIqnName == nil {
		return nil, false
	}
	return o.StaticIqnName, true
}

// HasStaticIqnName returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasStaticIqnName() bool {
	if o != nil && o.StaticIqnName != nil {
		return true
	}

	return false
}

// SetStaticIqnName gets a reference to the given string and assigns it to the StaticIqnName field.
func (o *VnicLanConnectivityPolicy) SetStaticIqnName(v string) {
	o.StaticIqnName = &v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicy) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *VnicLanConnectivityPolicy) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetEthIfs returns the EthIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicLanConnectivityPolicy) GetEthIfs() []VnicEthIfRelationship {
	if o == nil {
		var ret []VnicEthIfRelationship
		return ret
	}
	return o.EthIfs
}

// GetEthIfsOk returns a tuple with the EthIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicLanConnectivityPolicy) GetEthIfsOk() (*[]VnicEthIfRelationship, bool) {
	if o == nil || o.EthIfs == nil {
		return nil, false
	}
	return &o.EthIfs, true
}

// HasEthIfs returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasEthIfs() bool {
	if o != nil && o.EthIfs != nil {
		return true
	}

	return false
}

// SetEthIfs gets a reference to the given []VnicEthIfRelationship and assigns it to the EthIfs field.
func (o *VnicLanConnectivityPolicy) SetEthIfs(v []VnicEthIfRelationship) {
	o.EthIfs = v
}

// GetIqnPool returns the IqnPool field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicy) GetIqnPool() IqnpoolPoolRelationship {
	if o == nil || o.IqnPool == nil {
		var ret IqnpoolPoolRelationship
		return ret
	}
	return *o.IqnPool
}

// GetIqnPoolOk returns a tuple with the IqnPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetIqnPoolOk() (*IqnpoolPoolRelationship, bool) {
	if o == nil || o.IqnPool == nil {
		return nil, false
	}
	return o.IqnPool, true
}

// HasIqnPool returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasIqnPool() bool {
	if o != nil && o.IqnPool != nil {
		return true
	}

	return false
}

// SetIqnPool gets a reference to the given IqnpoolPoolRelationship and assigns it to the IqnPool field.
func (o *VnicLanConnectivityPolicy) SetIqnPool(v IqnpoolPoolRelationship) {
	o.IqnPool = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicLanConnectivityPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicLanConnectivityPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicLanConnectivityPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *VnicLanConnectivityPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o VnicLanConnectivityPolicy) MarshalJSON() ([]byte, error) {
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
	if o.AzureQosEnabled != nil {
		toSerialize["AzureQosEnabled"] = o.AzureQosEnabled
	}
	if o.IqnAllocationType != nil {
		toSerialize["IqnAllocationType"] = o.IqnAllocationType
	}
	if o.PlacementMode != nil {
		toSerialize["PlacementMode"] = o.PlacementMode
	}
	if o.StaticIqnName != nil {
		toSerialize["StaticIqnName"] = o.StaticIqnName
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.EthIfs != nil {
		toSerialize["EthIfs"] = o.EthIfs
	}
	if o.IqnPool != nil {
		toSerialize["IqnPool"] = o.IqnPool
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicLanConnectivityPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type VnicLanConnectivityPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enabling AzureStack-Host QoS on an adapter allows the user to carve out traffic classes for RDMA traffic which ensures that a desired portion of the bandwidth is allocated to it.
		AzureQosEnabled *bool `json:"AzureQosEnabled,omitempty"`
		// Allocation Type of iSCSI Qualified Name - Static/Pool/None. * `None` - Type indicates that there is no IQN associated to an interface. * `Static` - Type represents that static IQN is associated to an interface. * `Pool` - Type indicates that IQN value is sourced from an associated pool.
		IqnAllocationType *string `json:"IqnAllocationType,omitempty"`
		// The mode used for placement of vNICs on network adapters. It can either be Auto or Custom. * `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system.
		PlacementMode *string `json:"PlacementMode,omitempty"`
		// User provided static iSCSI Qualified Name (IQN) for use as initiator identifiers by iSCSI vNICs in a Fabric Interconnect domain.
		StaticIqnName *string `json:"StaticIqnName,omitempty"`
		// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
		TargetPlatform *string `json:"TargetPlatform,omitempty"`
		// An array of relationships to vnicEthIf resources.
		EthIfs       []VnicEthIfRelationship               `json:"EthIfs,omitempty"`
		IqnPool      *IqnpoolPoolRelationship              `json:"IqnPool,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varVnicLanConnectivityPolicyWithoutEmbeddedStruct := VnicLanConnectivityPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicLanConnectivityPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVnicLanConnectivityPolicy := _VnicLanConnectivityPolicy{}
		varVnicLanConnectivityPolicy.ClassId = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.ClassId
		varVnicLanConnectivityPolicy.ObjectType = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.ObjectType
		varVnicLanConnectivityPolicy.AzureQosEnabled = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.AzureQosEnabled
		varVnicLanConnectivityPolicy.IqnAllocationType = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.IqnAllocationType
		varVnicLanConnectivityPolicy.PlacementMode = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.PlacementMode
		varVnicLanConnectivityPolicy.StaticIqnName = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.StaticIqnName
		varVnicLanConnectivityPolicy.TargetPlatform = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.TargetPlatform
		varVnicLanConnectivityPolicy.EthIfs = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.EthIfs
		varVnicLanConnectivityPolicy.IqnPool = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.IqnPool
		varVnicLanConnectivityPolicy.Organization = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.Organization
		varVnicLanConnectivityPolicy.Profiles = varVnicLanConnectivityPolicyWithoutEmbeddedStruct.Profiles
		*o = VnicLanConnectivityPolicy(varVnicLanConnectivityPolicy)
	} else {
		return err
	}

	varVnicLanConnectivityPolicy := _VnicLanConnectivityPolicy{}

	err = json.Unmarshal(bytes, &varVnicLanConnectivityPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVnicLanConnectivityPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AzureQosEnabled")
		delete(additionalProperties, "IqnAllocationType")
		delete(additionalProperties, "PlacementMode")
		delete(additionalProperties, "StaticIqnName")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "EthIfs")
		delete(additionalProperties, "IqnPool")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

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

type NullableVnicLanConnectivityPolicy struct {
	value *VnicLanConnectivityPolicy
	isSet bool
}

func (v NullableVnicLanConnectivityPolicy) Get() *VnicLanConnectivityPolicy {
	return v.value
}

func (v *NullableVnicLanConnectivityPolicy) Set(val *VnicLanConnectivityPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicLanConnectivityPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicLanConnectivityPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicLanConnectivityPolicy(val *VnicLanConnectivityPolicy) *NullableVnicLanConnectivityPolicy {
	return &NullableVnicLanConnectivityPolicy{value: val, isSet: true}
}

func (v NullableVnicLanConnectivityPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicLanConnectivityPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
