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

// checks if the DnacFabricSite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnacFabricSite{}

// DnacFabricSite Details for the Fabric Site.
type DnacFabricSite struct {
	DnacInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Authentication profile name.
	AuthenticationProfileName *string `json:"AuthenticationProfileName,omitempty"`
	// UUID for the Fabric Site.
	FabricSiteId *string `json:"FabricSiteId,omitempty"`
	// Fabric site name hierarchy.
	FabricSiteNameHierarchy *string `json:"FabricSiteNameHierarchy,omitempty"`
	// Pub sub for the fabric site.
	IsPubSubEnabled *string `json:"IsPubSubEnabled,omitempty"`
	// Site id for the fabric site.
	SiteId               *string `json:"SiteId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DnacFabricSite DnacFabricSite

// NewDnacFabricSite instantiates a new DnacFabricSite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnacFabricSite(classId string, objectType string) *DnacFabricSite {
	this := DnacFabricSite{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewDnacFabricSiteWithDefaults instantiates a new DnacFabricSite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnacFabricSiteWithDefaults() *DnacFabricSite {
	this := DnacFabricSite{}
	var classId string = "dnac.FabricSite"
	this.ClassId = classId
	var objectType string = "dnac.FabricSite"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *DnacFabricSite) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *DnacFabricSite) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *DnacFabricSite) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "dnac.FabricSite" of the ClassId field.
func (o *DnacFabricSite) GetDefaultClassId() interface{} {
	return "dnac.FabricSite"
}

// GetObjectType returns the ObjectType field value
func (o *DnacFabricSite) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DnacFabricSite) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DnacFabricSite) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "dnac.FabricSite" of the ObjectType field.
func (o *DnacFabricSite) GetDefaultObjectType() interface{} {
	return "dnac.FabricSite"
}

// GetAuthenticationProfileName returns the AuthenticationProfileName field value if set, zero value otherwise.
func (o *DnacFabricSite) GetAuthenticationProfileName() string {
	if o == nil || IsNil(o.AuthenticationProfileName) {
		var ret string
		return ret
	}
	return *o.AuthenticationProfileName
}

// GetAuthenticationProfileNameOk returns a tuple with the AuthenticationProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacFabricSite) GetAuthenticationProfileNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationProfileName) {
		return nil, false
	}
	return o.AuthenticationProfileName, true
}

// HasAuthenticationProfileName returns a boolean if a field has been set.
func (o *DnacFabricSite) HasAuthenticationProfileName() bool {
	if o != nil && !IsNil(o.AuthenticationProfileName) {
		return true
	}

	return false
}

// SetAuthenticationProfileName gets a reference to the given string and assigns it to the AuthenticationProfileName field.
func (o *DnacFabricSite) SetAuthenticationProfileName(v string) {
	o.AuthenticationProfileName = &v
}

// GetFabricSiteId returns the FabricSiteId field value if set, zero value otherwise.
func (o *DnacFabricSite) GetFabricSiteId() string {
	if o == nil || IsNil(o.FabricSiteId) {
		var ret string
		return ret
	}
	return *o.FabricSiteId
}

// GetFabricSiteIdOk returns a tuple with the FabricSiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacFabricSite) GetFabricSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.FabricSiteId) {
		return nil, false
	}
	return o.FabricSiteId, true
}

// HasFabricSiteId returns a boolean if a field has been set.
func (o *DnacFabricSite) HasFabricSiteId() bool {
	if o != nil && !IsNil(o.FabricSiteId) {
		return true
	}

	return false
}

// SetFabricSiteId gets a reference to the given string and assigns it to the FabricSiteId field.
func (o *DnacFabricSite) SetFabricSiteId(v string) {
	o.FabricSiteId = &v
}

// GetFabricSiteNameHierarchy returns the FabricSiteNameHierarchy field value if set, zero value otherwise.
func (o *DnacFabricSite) GetFabricSiteNameHierarchy() string {
	if o == nil || IsNil(o.FabricSiteNameHierarchy) {
		var ret string
		return ret
	}
	return *o.FabricSiteNameHierarchy
}

// GetFabricSiteNameHierarchyOk returns a tuple with the FabricSiteNameHierarchy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacFabricSite) GetFabricSiteNameHierarchyOk() (*string, bool) {
	if o == nil || IsNil(o.FabricSiteNameHierarchy) {
		return nil, false
	}
	return o.FabricSiteNameHierarchy, true
}

// HasFabricSiteNameHierarchy returns a boolean if a field has been set.
func (o *DnacFabricSite) HasFabricSiteNameHierarchy() bool {
	if o != nil && !IsNil(o.FabricSiteNameHierarchy) {
		return true
	}

	return false
}

// SetFabricSiteNameHierarchy gets a reference to the given string and assigns it to the FabricSiteNameHierarchy field.
func (o *DnacFabricSite) SetFabricSiteNameHierarchy(v string) {
	o.FabricSiteNameHierarchy = &v
}

// GetIsPubSubEnabled returns the IsPubSubEnabled field value if set, zero value otherwise.
func (o *DnacFabricSite) GetIsPubSubEnabled() string {
	if o == nil || IsNil(o.IsPubSubEnabled) {
		var ret string
		return ret
	}
	return *o.IsPubSubEnabled
}

// GetIsPubSubEnabledOk returns a tuple with the IsPubSubEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacFabricSite) GetIsPubSubEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.IsPubSubEnabled) {
		return nil, false
	}
	return o.IsPubSubEnabled, true
}

// HasIsPubSubEnabled returns a boolean if a field has been set.
func (o *DnacFabricSite) HasIsPubSubEnabled() bool {
	if o != nil && !IsNil(o.IsPubSubEnabled) {
		return true
	}

	return false
}

// SetIsPubSubEnabled gets a reference to the given string and assigns it to the IsPubSubEnabled field.
func (o *DnacFabricSite) SetIsPubSubEnabled(v string) {
	o.IsPubSubEnabled = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *DnacFabricSite) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnacFabricSite) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *DnacFabricSite) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *DnacFabricSite) SetSiteId(v string) {
	o.SiteId = &v
}

func (o DnacFabricSite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnacFabricSite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDnacInventoryEntity, errDnacInventoryEntity := json.Marshal(o.DnacInventoryEntity)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	errDnacInventoryEntity = json.Unmarshal([]byte(serializedDnacInventoryEntity), &toSerialize)
	if errDnacInventoryEntity != nil {
		return map[string]interface{}{}, errDnacInventoryEntity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AuthenticationProfileName) {
		toSerialize["AuthenticationProfileName"] = o.AuthenticationProfileName
	}
	if !IsNil(o.FabricSiteId) {
		toSerialize["FabricSiteId"] = o.FabricSiteId
	}
	if !IsNil(o.FabricSiteNameHierarchy) {
		toSerialize["FabricSiteNameHierarchy"] = o.FabricSiteNameHierarchy
	}
	if !IsNil(o.IsPubSubEnabled) {
		toSerialize["IsPubSubEnabled"] = o.IsPubSubEnabled
	}
	if !IsNil(o.SiteId) {
		toSerialize["SiteId"] = o.SiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DnacFabricSite) UnmarshalJSON(data []byte) (err error) {
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
	type DnacFabricSiteWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Authentication profile name.
		AuthenticationProfileName *string `json:"AuthenticationProfileName,omitempty"`
		// UUID for the Fabric Site.
		FabricSiteId *string `json:"FabricSiteId,omitempty"`
		// Fabric site name hierarchy.
		FabricSiteNameHierarchy *string `json:"FabricSiteNameHierarchy,omitempty"`
		// Pub sub for the fabric site.
		IsPubSubEnabled *string `json:"IsPubSubEnabled,omitempty"`
		// Site id for the fabric site.
		SiteId *string `json:"SiteId,omitempty"`
	}

	varDnacFabricSiteWithoutEmbeddedStruct := DnacFabricSiteWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDnacFabricSiteWithoutEmbeddedStruct)
	if err == nil {
		varDnacFabricSite := _DnacFabricSite{}
		varDnacFabricSite.ClassId = varDnacFabricSiteWithoutEmbeddedStruct.ClassId
		varDnacFabricSite.ObjectType = varDnacFabricSiteWithoutEmbeddedStruct.ObjectType
		varDnacFabricSite.AuthenticationProfileName = varDnacFabricSiteWithoutEmbeddedStruct.AuthenticationProfileName
		varDnacFabricSite.FabricSiteId = varDnacFabricSiteWithoutEmbeddedStruct.FabricSiteId
		varDnacFabricSite.FabricSiteNameHierarchy = varDnacFabricSiteWithoutEmbeddedStruct.FabricSiteNameHierarchy
		varDnacFabricSite.IsPubSubEnabled = varDnacFabricSiteWithoutEmbeddedStruct.IsPubSubEnabled
		varDnacFabricSite.SiteId = varDnacFabricSiteWithoutEmbeddedStruct.SiteId
		*o = DnacFabricSite(varDnacFabricSite)
	} else {
		return err
	}

	varDnacFabricSite := _DnacFabricSite{}

	err = json.Unmarshal(data, &varDnacFabricSite)
	if err == nil {
		o.DnacInventoryEntity = varDnacFabricSite.DnacInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthenticationProfileName")
		delete(additionalProperties, "FabricSiteId")
		delete(additionalProperties, "FabricSiteNameHierarchy")
		delete(additionalProperties, "IsPubSubEnabled")
		delete(additionalProperties, "SiteId")

		// remove fields from embedded structs
		reflectDnacInventoryEntity := reflect.ValueOf(o.DnacInventoryEntity)
		for i := 0; i < reflectDnacInventoryEntity.Type().NumField(); i++ {
			t := reflectDnacInventoryEntity.Type().Field(i)

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

type NullableDnacFabricSite struct {
	value *DnacFabricSite
	isSet bool
}

func (v NullableDnacFabricSite) Get() *DnacFabricSite {
	return v.value
}

func (v *NullableDnacFabricSite) Set(val *DnacFabricSite) {
	v.value = val
	v.isSet = true
}

func (v NullableDnacFabricSite) IsSet() bool {
	return v.isSet
}

func (v *NullableDnacFabricSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnacFabricSite(val *DnacFabricSite) *NullableDnacFabricSite {
	return &NullableDnacFabricSite{value: val, isSet: true}
}

func (v NullableDnacFabricSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnacFabricSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
