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

// checks if the ResourceMembership type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceMembership{}

// ResourceMembership ResourceMembership represents a resource's associated groups, organizations and the permissions associated to this resource via organizations.
type ResourceMembership struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                      `json:"ObjectType"`
	GroupPermissionRoles []IamGroupPermissionToRoles `json:"GroupPermissionRoles,omitempty"`
	// Set Reevaluate to true to reevaluate the membership of a resource.
	Reevaluate *bool `json:"Reevaluate,omitempty"`
	// Name of the Service owning the resource.
	TargetApp *string                                      `json:"TargetApp,omitempty"`
	Holder    NullableResourceMembershipHolderRelationship `json:"Holder,omitempty"`
	Resource  NullableMoBaseMoRelationship                 `json:"Resource,omitempty"`
	// An array of relationships to moBaseMo resources.
	ResourceAncestors    []MoBaseMoRelationship `json:"ResourceAncestors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceMembership ResourceMembership

// NewResourceMembership instantiates a new ResourceMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceMembership(classId string, objectType string) *ResourceMembership {
	this := ResourceMembership{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourceMembershipWithDefaults instantiates a new ResourceMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceMembershipWithDefaults() *ResourceMembership {
	this := ResourceMembership{}
	var classId string = "resource.Membership"
	this.ClassId = classId
	var objectType string = "resource.Membership"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceMembership) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceMembership) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceMembership) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "resource.Membership" of the ClassId field.
func (o *ResourceMembership) GetDefaultClassId() interface{} {
	return "resource.Membership"
}

// GetObjectType returns the ObjectType field value
func (o *ResourceMembership) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceMembership) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceMembership) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "resource.Membership" of the ObjectType field.
func (o *ResourceMembership) GetDefaultObjectType() interface{} {
	return "resource.Membership"
}

// GetGroupPermissionRoles returns the GroupPermissionRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceMembership) GetGroupPermissionRoles() []IamGroupPermissionToRoles {
	if o == nil {
		var ret []IamGroupPermissionToRoles
		return ret
	}
	return o.GroupPermissionRoles
}

// GetGroupPermissionRolesOk returns a tuple with the GroupPermissionRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceMembership) GetGroupPermissionRolesOk() ([]IamGroupPermissionToRoles, bool) {
	if o == nil || IsNil(o.GroupPermissionRoles) {
		return nil, false
	}
	return o.GroupPermissionRoles, true
}

// HasGroupPermissionRoles returns a boolean if a field has been set.
func (o *ResourceMembership) HasGroupPermissionRoles() bool {
	if o != nil && !IsNil(o.GroupPermissionRoles) {
		return true
	}

	return false
}

// SetGroupPermissionRoles gets a reference to the given []IamGroupPermissionToRoles and assigns it to the GroupPermissionRoles field.
func (o *ResourceMembership) SetGroupPermissionRoles(v []IamGroupPermissionToRoles) {
	o.GroupPermissionRoles = v
}

// GetReevaluate returns the Reevaluate field value if set, zero value otherwise.
func (o *ResourceMembership) GetReevaluate() bool {
	if o == nil || IsNil(o.Reevaluate) {
		var ret bool
		return ret
	}
	return *o.Reevaluate
}

// GetReevaluateOk returns a tuple with the Reevaluate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceMembership) GetReevaluateOk() (*bool, bool) {
	if o == nil || IsNil(o.Reevaluate) {
		return nil, false
	}
	return o.Reevaluate, true
}

// HasReevaluate returns a boolean if a field has been set.
func (o *ResourceMembership) HasReevaluate() bool {
	if o != nil && !IsNil(o.Reevaluate) {
		return true
	}

	return false
}

// SetReevaluate gets a reference to the given bool and assigns it to the Reevaluate field.
func (o *ResourceMembership) SetReevaluate(v bool) {
	o.Reevaluate = &v
}

// GetTargetApp returns the TargetApp field value if set, zero value otherwise.
func (o *ResourceMembership) GetTargetApp() string {
	if o == nil || IsNil(o.TargetApp) {
		var ret string
		return ret
	}
	return *o.TargetApp
}

// GetTargetAppOk returns a tuple with the TargetApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceMembership) GetTargetAppOk() (*string, bool) {
	if o == nil || IsNil(o.TargetApp) {
		return nil, false
	}
	return o.TargetApp, true
}

// HasTargetApp returns a boolean if a field has been set.
func (o *ResourceMembership) HasTargetApp() bool {
	if o != nil && !IsNil(o.TargetApp) {
		return true
	}

	return false
}

// SetTargetApp gets a reference to the given string and assigns it to the TargetApp field.
func (o *ResourceMembership) SetTargetApp(v string) {
	o.TargetApp = &v
}

// GetHolder returns the Holder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceMembership) GetHolder() ResourceMembershipHolderRelationship {
	if o == nil || IsNil(o.Holder.Get()) {
		var ret ResourceMembershipHolderRelationship
		return ret
	}
	return *o.Holder.Get()
}

// GetHolderOk returns a tuple with the Holder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceMembership) GetHolderOk() (*ResourceMembershipHolderRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Holder.Get(), o.Holder.IsSet()
}

// HasHolder returns a boolean if a field has been set.
func (o *ResourceMembership) HasHolder() bool {
	if o != nil && o.Holder.IsSet() {
		return true
	}

	return false
}

// SetHolder gets a reference to the given NullableResourceMembershipHolderRelationship and assigns it to the Holder field.
func (o *ResourceMembership) SetHolder(v ResourceMembershipHolderRelationship) {
	o.Holder.Set(&v)
}

// SetHolderNil sets the value for Holder to be an explicit nil
func (o *ResourceMembership) SetHolderNil() {
	o.Holder.Set(nil)
}

// UnsetHolder ensures that no value is present for Holder, not even an explicit nil
func (o *ResourceMembership) UnsetHolder() {
	o.Holder.Unset()
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceMembership) GetResource() MoBaseMoRelationship {
	if o == nil || IsNil(o.Resource.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceMembership) GetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// HasResource returns a boolean if a field has been set.
func (o *ResourceMembership) HasResource() bool {
	if o != nil && o.Resource.IsSet() {
		return true
	}

	return false
}

// SetResource gets a reference to the given NullableMoBaseMoRelationship and assigns it to the Resource field.
func (o *ResourceMembership) SetResource(v MoBaseMoRelationship) {
	o.Resource.Set(&v)
}

// SetResourceNil sets the value for Resource to be an explicit nil
func (o *ResourceMembership) SetResourceNil() {
	o.Resource.Set(nil)
}

// UnsetResource ensures that no value is present for Resource, not even an explicit nil
func (o *ResourceMembership) UnsetResource() {
	o.Resource.Unset()
}

// GetResourceAncestors returns the ResourceAncestors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceMembership) GetResourceAncestors() []MoBaseMoRelationship {
	if o == nil {
		var ret []MoBaseMoRelationship
		return ret
	}
	return o.ResourceAncestors
}

// GetResourceAncestorsOk returns a tuple with the ResourceAncestors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceMembership) GetResourceAncestorsOk() ([]MoBaseMoRelationship, bool) {
	if o == nil || IsNil(o.ResourceAncestors) {
		return nil, false
	}
	return o.ResourceAncestors, true
}

// HasResourceAncestors returns a boolean if a field has been set.
func (o *ResourceMembership) HasResourceAncestors() bool {
	if o != nil && !IsNil(o.ResourceAncestors) {
		return true
	}

	return false
}

// SetResourceAncestors gets a reference to the given []MoBaseMoRelationship and assigns it to the ResourceAncestors field.
func (o *ResourceMembership) SetResourceAncestors(v []MoBaseMoRelationship) {
	o.ResourceAncestors = v
}

func (o ResourceMembership) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceMembership) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.GroupPermissionRoles != nil {
		toSerialize["GroupPermissionRoles"] = o.GroupPermissionRoles
	}
	if !IsNil(o.Reevaluate) {
		toSerialize["Reevaluate"] = o.Reevaluate
	}
	if !IsNil(o.TargetApp) {
		toSerialize["TargetApp"] = o.TargetApp
	}
	if o.Holder.IsSet() {
		toSerialize["Holder"] = o.Holder.Get()
	}
	if o.Resource.IsSet() {
		toSerialize["Resource"] = o.Resource.Get()
	}
	if o.ResourceAncestors != nil {
		toSerialize["ResourceAncestors"] = o.ResourceAncestors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceMembership) UnmarshalJSON(data []byte) (err error) {
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
	type ResourceMembershipWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                      `json:"ObjectType"`
		GroupPermissionRoles []IamGroupPermissionToRoles `json:"GroupPermissionRoles,omitempty"`
		// Set Reevaluate to true to reevaluate the membership of a resource.
		Reevaluate *bool `json:"Reevaluate,omitempty"`
		// Name of the Service owning the resource.
		TargetApp *string                                      `json:"TargetApp,omitempty"`
		Holder    NullableResourceMembershipHolderRelationship `json:"Holder,omitempty"`
		Resource  NullableMoBaseMoRelationship                 `json:"Resource,omitempty"`
		// An array of relationships to moBaseMo resources.
		ResourceAncestors []MoBaseMoRelationship `json:"ResourceAncestors,omitempty"`
	}

	varResourceMembershipWithoutEmbeddedStruct := ResourceMembershipWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varResourceMembershipWithoutEmbeddedStruct)
	if err == nil {
		varResourceMembership := _ResourceMembership{}
		varResourceMembership.ClassId = varResourceMembershipWithoutEmbeddedStruct.ClassId
		varResourceMembership.ObjectType = varResourceMembershipWithoutEmbeddedStruct.ObjectType
		varResourceMembership.GroupPermissionRoles = varResourceMembershipWithoutEmbeddedStruct.GroupPermissionRoles
		varResourceMembership.Reevaluate = varResourceMembershipWithoutEmbeddedStruct.Reevaluate
		varResourceMembership.TargetApp = varResourceMembershipWithoutEmbeddedStruct.TargetApp
		varResourceMembership.Holder = varResourceMembershipWithoutEmbeddedStruct.Holder
		varResourceMembership.Resource = varResourceMembershipWithoutEmbeddedStruct.Resource
		varResourceMembership.ResourceAncestors = varResourceMembershipWithoutEmbeddedStruct.ResourceAncestors
		*o = ResourceMembership(varResourceMembership)
	} else {
		return err
	}

	varResourceMembership := _ResourceMembership{}

	err = json.Unmarshal(data, &varResourceMembership)
	if err == nil {
		o.MoBaseMo = varResourceMembership.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "GroupPermissionRoles")
		delete(additionalProperties, "Reevaluate")
		delete(additionalProperties, "TargetApp")
		delete(additionalProperties, "Holder")
		delete(additionalProperties, "Resource")
		delete(additionalProperties, "ResourceAncestors")

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

type NullableResourceMembership struct {
	value *ResourceMembership
	isSet bool
}

func (v NullableResourceMembership) Get() *ResourceMembership {
	return v.value
}

func (v *NullableResourceMembership) Set(val *ResourceMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceMembership(val *ResourceMembership) *NullableResourceMembership {
	return &NullableResourceMembership{value: val, isSet: true}
}

func (v NullableResourceMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
