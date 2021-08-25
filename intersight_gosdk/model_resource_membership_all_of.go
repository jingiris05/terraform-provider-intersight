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
)

// ResourceMembershipAllOf Definition of the list of properties defined in 'resource.Membership', excluding properties defined in parent classes.
type ResourceMembershipAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                      `json:"ObjectType"`
	GroupPermissionRoles []IamGroupPermissionToRoles `json:"GroupPermissionRoles,omitempty"`
	// Name of the Service owning the resource.
	TargetApp            *string                               `json:"TargetApp,omitempty"`
	Holder               *ResourceMembershipHolderRelationship `json:"Holder,omitempty"`
	Resource             *MoBaseMoRelationship                 `json:"Resource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceMembershipAllOf ResourceMembershipAllOf

// NewResourceMembershipAllOf instantiates a new ResourceMembershipAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceMembershipAllOf(classId string, objectType string) *ResourceMembershipAllOf {
	this := ResourceMembershipAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourceMembershipAllOfWithDefaults instantiates a new ResourceMembershipAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceMembershipAllOfWithDefaults() *ResourceMembershipAllOf {
	this := ResourceMembershipAllOf{}
	var classId string = "resource.Membership"
	this.ClassId = classId
	var objectType string = "resource.Membership"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceMembershipAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceMembershipAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceMembershipAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourceMembershipAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceMembershipAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceMembershipAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGroupPermissionRoles returns the GroupPermissionRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceMembershipAllOf) GetGroupPermissionRoles() []IamGroupPermissionToRoles {
	if o == nil {
		var ret []IamGroupPermissionToRoles
		return ret
	}
	return o.GroupPermissionRoles
}

// GetGroupPermissionRolesOk returns a tuple with the GroupPermissionRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceMembershipAllOf) GetGroupPermissionRolesOk() (*[]IamGroupPermissionToRoles, bool) {
	if o == nil || o.GroupPermissionRoles == nil {
		return nil, false
	}
	return &o.GroupPermissionRoles, true
}

// HasGroupPermissionRoles returns a boolean if a field has been set.
func (o *ResourceMembershipAllOf) HasGroupPermissionRoles() bool {
	if o != nil && o.GroupPermissionRoles != nil {
		return true
	}

	return false
}

// SetGroupPermissionRoles gets a reference to the given []IamGroupPermissionToRoles and assigns it to the GroupPermissionRoles field.
func (o *ResourceMembershipAllOf) SetGroupPermissionRoles(v []IamGroupPermissionToRoles) {
	o.GroupPermissionRoles = v
}

// GetTargetApp returns the TargetApp field value if set, zero value otherwise.
func (o *ResourceMembershipAllOf) GetTargetApp() string {
	if o == nil || o.TargetApp == nil {
		var ret string
		return ret
	}
	return *o.TargetApp
}

// GetTargetAppOk returns a tuple with the TargetApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceMembershipAllOf) GetTargetAppOk() (*string, bool) {
	if o == nil || o.TargetApp == nil {
		return nil, false
	}
	return o.TargetApp, true
}

// HasTargetApp returns a boolean if a field has been set.
func (o *ResourceMembershipAllOf) HasTargetApp() bool {
	if o != nil && o.TargetApp != nil {
		return true
	}

	return false
}

// SetTargetApp gets a reference to the given string and assigns it to the TargetApp field.
func (o *ResourceMembershipAllOf) SetTargetApp(v string) {
	o.TargetApp = &v
}

// GetHolder returns the Holder field value if set, zero value otherwise.
func (o *ResourceMembershipAllOf) GetHolder() ResourceMembershipHolderRelationship {
	if o == nil || o.Holder == nil {
		var ret ResourceMembershipHolderRelationship
		return ret
	}
	return *o.Holder
}

// GetHolderOk returns a tuple with the Holder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceMembershipAllOf) GetHolderOk() (*ResourceMembershipHolderRelationship, bool) {
	if o == nil || o.Holder == nil {
		return nil, false
	}
	return o.Holder, true
}

// HasHolder returns a boolean if a field has been set.
func (o *ResourceMembershipAllOf) HasHolder() bool {
	if o != nil && o.Holder != nil {
		return true
	}

	return false
}

// SetHolder gets a reference to the given ResourceMembershipHolderRelationship and assigns it to the Holder field.
func (o *ResourceMembershipAllOf) SetHolder(v ResourceMembershipHolderRelationship) {
	o.Holder = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourceMembershipAllOf) GetResource() MoBaseMoRelationship {
	if o == nil || o.Resource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceMembershipAllOf) GetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourceMembershipAllOf) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given MoBaseMoRelationship and assigns it to the Resource field.
func (o *ResourceMembershipAllOf) SetResource(v MoBaseMoRelationship) {
	o.Resource = &v
}

func (o ResourceMembershipAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.GroupPermissionRoles != nil {
		toSerialize["GroupPermissionRoles"] = o.GroupPermissionRoles
	}
	if o.TargetApp != nil {
		toSerialize["TargetApp"] = o.TargetApp
	}
	if o.Holder != nil {
		toSerialize["Holder"] = o.Holder
	}
	if o.Resource != nil {
		toSerialize["Resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceMembershipAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourceMembershipAllOf := _ResourceMembershipAllOf{}

	if err = json.Unmarshal(bytes, &varResourceMembershipAllOf); err == nil {
		*o = ResourceMembershipAllOf(varResourceMembershipAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "GroupPermissionRoles")
		delete(additionalProperties, "TargetApp")
		delete(additionalProperties, "Holder")
		delete(additionalProperties, "Resource")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceMembershipAllOf struct {
	value *ResourceMembershipAllOf
	isSet bool
}

func (v NullableResourceMembershipAllOf) Get() *ResourceMembershipAllOf {
	return v.value
}

func (v *NullableResourceMembershipAllOf) Set(val *ResourceMembershipAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceMembershipAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceMembershipAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceMembershipAllOf(val *ResourceMembershipAllOf) *NullableResourceMembershipAllOf {
	return &NullableResourceMembershipAllOf{value: val, isSet: true}
}

func (v NullableResourceMembershipAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceMembershipAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
