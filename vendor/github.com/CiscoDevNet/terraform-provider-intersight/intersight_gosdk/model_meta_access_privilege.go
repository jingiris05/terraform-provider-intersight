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

// MetaAccessPrivilege The required access privilege for a given Managed Object and CRUD operation.
type MetaAccessPrivilege struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The type of CRUD operation (create, read, update, delete) for which an access privilege is required. * `Update` - The 'update' operation/state. * `Create` - The 'create' operation/state. * `Read` - The 'read' operation/state. * `Delete` - The 'delete' operation/state.
	Method *string `json:"Method,omitempty"`
	// The name of the privilege which is required to invoke the specified CRUD method.
	Privilege            *string `json:"Privilege,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetaAccessPrivilege MetaAccessPrivilege

// NewMetaAccessPrivilege instantiates a new MetaAccessPrivilege object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaAccessPrivilege(classId string, objectType string) *MetaAccessPrivilege {
	this := MetaAccessPrivilege{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMetaAccessPrivilegeWithDefaults instantiates a new MetaAccessPrivilege object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaAccessPrivilegeWithDefaults() *MetaAccessPrivilege {
	this := MetaAccessPrivilege{}
	var classId string = "meta.AccessPrivilege"
	this.ClassId = classId
	var objectType string = "meta.AccessPrivilege"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetaAccessPrivilege) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetaAccessPrivilege) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetaAccessPrivilege) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MetaAccessPrivilege) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetaAccessPrivilege) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetaAccessPrivilege) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *MetaAccessPrivilege) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAccessPrivilege) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *MetaAccessPrivilege) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *MetaAccessPrivilege) SetMethod(v string) {
	o.Method = &v
}

// GetPrivilege returns the Privilege field value if set, zero value otherwise.
func (o *MetaAccessPrivilege) GetPrivilege() string {
	if o == nil || o.Privilege == nil {
		var ret string
		return ret
	}
	return *o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAccessPrivilege) GetPrivilegeOk() (*string, bool) {
	if o == nil || o.Privilege == nil {
		return nil, false
	}
	return o.Privilege, true
}

// HasPrivilege returns a boolean if a field has been set.
func (o *MetaAccessPrivilege) HasPrivilege() bool {
	if o != nil && o.Privilege != nil {
		return true
	}

	return false
}

// SetPrivilege gets a reference to the given string and assigns it to the Privilege field.
func (o *MetaAccessPrivilege) SetPrivilege(v string) {
	o.Privilege = &v
}

func (o MetaAccessPrivilege) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.Privilege != nil {
		toSerialize["Privilege"] = o.Privilege
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MetaAccessPrivilege) UnmarshalJSON(bytes []byte) (err error) {
	type MetaAccessPrivilegeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The type of CRUD operation (create, read, update, delete) for which an access privilege is required. * `Update` - The 'update' operation/state. * `Create` - The 'create' operation/state. * `Read` - The 'read' operation/state. * `Delete` - The 'delete' operation/state.
		Method *string `json:"Method,omitempty"`
		// The name of the privilege which is required to invoke the specified CRUD method.
		Privilege *string `json:"Privilege,omitempty"`
	}

	varMetaAccessPrivilegeWithoutEmbeddedStruct := MetaAccessPrivilegeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMetaAccessPrivilegeWithoutEmbeddedStruct)
	if err == nil {
		varMetaAccessPrivilege := _MetaAccessPrivilege{}
		varMetaAccessPrivilege.ClassId = varMetaAccessPrivilegeWithoutEmbeddedStruct.ClassId
		varMetaAccessPrivilege.ObjectType = varMetaAccessPrivilegeWithoutEmbeddedStruct.ObjectType
		varMetaAccessPrivilege.Method = varMetaAccessPrivilegeWithoutEmbeddedStruct.Method
		varMetaAccessPrivilege.Privilege = varMetaAccessPrivilegeWithoutEmbeddedStruct.Privilege
		*o = MetaAccessPrivilege(varMetaAccessPrivilege)
	} else {
		return err
	}

	varMetaAccessPrivilege := _MetaAccessPrivilege{}

	err = json.Unmarshal(bytes, &varMetaAccessPrivilege)
	if err == nil {
		o.MoBaseComplexType = varMetaAccessPrivilege.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "Privilege")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableMetaAccessPrivilege struct {
	value *MetaAccessPrivilege
	isSet bool
}

func (v NullableMetaAccessPrivilege) Get() *MetaAccessPrivilege {
	return v.value
}

func (v *NullableMetaAccessPrivilege) Set(val *MetaAccessPrivilege) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaAccessPrivilege) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaAccessPrivilege) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaAccessPrivilege(val *MetaAccessPrivilege) *NullableMetaAccessPrivilege {
	return &NullableMetaAccessPrivilege{value: val, isSet: true}
}

func (v NullableMetaAccessPrivilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaAccessPrivilege) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
