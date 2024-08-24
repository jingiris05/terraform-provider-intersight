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

// checks if the ApicVrfs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApicVrfs{}

// ApicVrfs VRFs configured on cisco APIC.
type ApicVrfs struct {
	ApicInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Distinguished Name generated from URL Parameters.
	Dn *string `json:"Dn,omitempty"`
	// VRF Namegenerated from URL Parameters.
	Name                 *string                        `json:"Name,omitempty"`
	Tenant               NullableApicTenantRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApicVrfs ApicVrfs

// NewApicVrfs instantiates a new ApicVrfs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApicVrfs(classId string, objectType string) *ApicVrfs {
	this := ApicVrfs{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApicVrfsWithDefaults instantiates a new ApicVrfs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApicVrfsWithDefaults() *ApicVrfs {
	this := ApicVrfs{}
	var classId string = "apic.Vrfs"
	this.ClassId = classId
	var objectType string = "apic.Vrfs"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApicVrfs) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApicVrfs) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApicVrfs) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "apic.Vrfs" of the ClassId field.
func (o *ApicVrfs) GetDefaultClassId() interface{} {
	return "apic.Vrfs"
}

// GetObjectType returns the ObjectType field value
func (o *ApicVrfs) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApicVrfs) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApicVrfs) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "apic.Vrfs" of the ObjectType field.
func (o *ApicVrfs) GetDefaultObjectType() interface{} {
	return "apic.Vrfs"
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *ApicVrfs) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicVrfs) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *ApicVrfs) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *ApicVrfs) SetDn(v string) {
	o.Dn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApicVrfs) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApicVrfs) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApicVrfs) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApicVrfs) SetName(v string) {
	o.Name = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApicVrfs) GetTenant() ApicTenantRelationship {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret ApicTenantRelationship
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApicVrfs) GetTenantOk() (*ApicTenantRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *ApicVrfs) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableApicTenantRelationship and assigns it to the Tenant field.
func (o *ApicVrfs) SetTenant(v ApicTenantRelationship) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *ApicVrfs) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *ApicVrfs) UnsetTenant() {
	o.Tenant.Unset()
}

func (o ApicVrfs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApicVrfs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedApicInventoryEntity, errApicInventoryEntity := json.Marshal(o.ApicInventoryEntity)
	if errApicInventoryEntity != nil {
		return map[string]interface{}{}, errApicInventoryEntity
	}
	errApicInventoryEntity = json.Unmarshal([]byte(serializedApicInventoryEntity), &toSerialize)
	if errApicInventoryEntity != nil {
		return map[string]interface{}{}, errApicInventoryEntity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Tenant.IsSet() {
		toSerialize["Tenant"] = o.Tenant.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApicVrfs) UnmarshalJSON(data []byte) (err error) {
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
	type ApicVrfsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Distinguished Name generated from URL Parameters.
		Dn *string `json:"Dn,omitempty"`
		// VRF Namegenerated from URL Parameters.
		Name   *string                        `json:"Name,omitempty"`
		Tenant NullableApicTenantRelationship `json:"Tenant,omitempty"`
	}

	varApicVrfsWithoutEmbeddedStruct := ApicVrfsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApicVrfsWithoutEmbeddedStruct)
	if err == nil {
		varApicVrfs := _ApicVrfs{}
		varApicVrfs.ClassId = varApicVrfsWithoutEmbeddedStruct.ClassId
		varApicVrfs.ObjectType = varApicVrfsWithoutEmbeddedStruct.ObjectType
		varApicVrfs.Dn = varApicVrfsWithoutEmbeddedStruct.Dn
		varApicVrfs.Name = varApicVrfsWithoutEmbeddedStruct.Name
		varApicVrfs.Tenant = varApicVrfsWithoutEmbeddedStruct.Tenant
		*o = ApicVrfs(varApicVrfs)
	} else {
		return err
	}

	varApicVrfs := _ApicVrfs{}

	err = json.Unmarshal(data, &varApicVrfs)
	if err == nil {
		o.ApicInventoryEntity = varApicVrfs.ApicInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Tenant")

		// remove fields from embedded structs
		reflectApicInventoryEntity := reflect.ValueOf(o.ApicInventoryEntity)
		for i := 0; i < reflectApicInventoryEntity.Type().NumField(); i++ {
			t := reflectApicInventoryEntity.Type().Field(i)

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

type NullableApicVrfs struct {
	value *ApicVrfs
	isSet bool
}

func (v NullableApicVrfs) Get() *ApicVrfs {
	return v.value
}

func (v *NullableApicVrfs) Set(val *ApicVrfs) {
	v.value = val
	v.isSet = true
}

func (v NullableApicVrfs) IsSet() bool {
	return v.isSet
}

func (v *NullableApicVrfs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApicVrfs(val *ApicVrfs) *NullableApicVrfs {
	return &NullableApicVrfs{value: val, isSet: true}
}

func (v NullableApicVrfs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApicVrfs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
