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

// ComputePersistentMemoryOperationAllOf Definition of the list of properties defined in 'compute.PersistentMemoryOperation', excluding properties defined in parent classes.
type ComputePersistentMemoryOperationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Administrative actions that can be performed on the Persistent Memory Modules. * `None` - No action on the selected Persistent Memory Modules. * `SecureErase` - Secure Erase action on the selected Persistent Memory Modules. * `Unlock` - Unlock action on the selected Persistent Memory Modules.
	AdminAction *string `json:"AdminAction,omitempty"`
	// Indicates whether the value of the 'securePassphrase' property has been set.
	IsSecurePassphraseSet *bool                           `json:"IsSecurePassphraseSet,omitempty"`
	Modules               []ComputePersistentMemoryModule `json:"Modules,omitempty"`
	// Secure passphrase of the Persistent Memory Modules of the server.
	SecurePassphrase     *string `json:"SecurePassphrase,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputePersistentMemoryOperationAllOf ComputePersistentMemoryOperationAllOf

// NewComputePersistentMemoryOperationAllOf instantiates a new ComputePersistentMemoryOperationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputePersistentMemoryOperationAllOf(classId string, objectType string) *ComputePersistentMemoryOperationAllOf {
	this := ComputePersistentMemoryOperationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	var isSecurePassphraseSet bool = false
	this.IsSecurePassphraseSet = &isSecurePassphraseSet
	return &this
}

// NewComputePersistentMemoryOperationAllOfWithDefaults instantiates a new ComputePersistentMemoryOperationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputePersistentMemoryOperationAllOfWithDefaults() *ComputePersistentMemoryOperationAllOf {
	this := ComputePersistentMemoryOperationAllOf{}
	var classId string = "compute.PersistentMemoryOperation"
	this.ClassId = classId
	var objectType string = "compute.PersistentMemoryOperation"
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	var isSecurePassphraseSet bool = false
	this.IsSecurePassphraseSet = &isSecurePassphraseSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputePersistentMemoryOperationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryOperationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputePersistentMemoryOperationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputePersistentMemoryOperationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryOperationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputePersistentMemoryOperationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminAction returns the AdminAction field value if set, zero value otherwise.
func (o *ComputePersistentMemoryOperationAllOf) GetAdminAction() string {
	if o == nil || o.AdminAction == nil {
		var ret string
		return ret
	}
	return *o.AdminAction
}

// GetAdminActionOk returns a tuple with the AdminAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryOperationAllOf) GetAdminActionOk() (*string, bool) {
	if o == nil || o.AdminAction == nil {
		return nil, false
	}
	return o.AdminAction, true
}

// HasAdminAction returns a boolean if a field has been set.
func (o *ComputePersistentMemoryOperationAllOf) HasAdminAction() bool {
	if o != nil && o.AdminAction != nil {
		return true
	}

	return false
}

// SetAdminAction gets a reference to the given string and assigns it to the AdminAction field.
func (o *ComputePersistentMemoryOperationAllOf) SetAdminAction(v string) {
	o.AdminAction = &v
}

// GetIsSecurePassphraseSet returns the IsSecurePassphraseSet field value if set, zero value otherwise.
func (o *ComputePersistentMemoryOperationAllOf) GetIsSecurePassphraseSet() bool {
	if o == nil || o.IsSecurePassphraseSet == nil {
		var ret bool
		return ret
	}
	return *o.IsSecurePassphraseSet
}

// GetIsSecurePassphraseSetOk returns a tuple with the IsSecurePassphraseSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryOperationAllOf) GetIsSecurePassphraseSetOk() (*bool, bool) {
	if o == nil || o.IsSecurePassphraseSet == nil {
		return nil, false
	}
	return o.IsSecurePassphraseSet, true
}

// HasIsSecurePassphraseSet returns a boolean if a field has been set.
func (o *ComputePersistentMemoryOperationAllOf) HasIsSecurePassphraseSet() bool {
	if o != nil && o.IsSecurePassphraseSet != nil {
		return true
	}

	return false
}

// SetIsSecurePassphraseSet gets a reference to the given bool and assigns it to the IsSecurePassphraseSet field.
func (o *ComputePersistentMemoryOperationAllOf) SetIsSecurePassphraseSet(v bool) {
	o.IsSecurePassphraseSet = &v
}

// GetModules returns the Modules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputePersistentMemoryOperationAllOf) GetModules() []ComputePersistentMemoryModule {
	if o == nil {
		var ret []ComputePersistentMemoryModule
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputePersistentMemoryOperationAllOf) GetModulesOk() (*[]ComputePersistentMemoryModule, bool) {
	if o == nil || o.Modules == nil {
		return nil, false
	}
	return &o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *ComputePersistentMemoryOperationAllOf) HasModules() bool {
	if o != nil && o.Modules != nil {
		return true
	}

	return false
}

// SetModules gets a reference to the given []ComputePersistentMemoryModule and assigns it to the Modules field.
func (o *ComputePersistentMemoryOperationAllOf) SetModules(v []ComputePersistentMemoryModule) {
	o.Modules = v
}

// GetSecurePassphrase returns the SecurePassphrase field value if set, zero value otherwise.
func (o *ComputePersistentMemoryOperationAllOf) GetSecurePassphrase() string {
	if o == nil || o.SecurePassphrase == nil {
		var ret string
		return ret
	}
	return *o.SecurePassphrase
}

// GetSecurePassphraseOk returns a tuple with the SecurePassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryOperationAllOf) GetSecurePassphraseOk() (*string, bool) {
	if o == nil || o.SecurePassphrase == nil {
		return nil, false
	}
	return o.SecurePassphrase, true
}

// HasSecurePassphrase returns a boolean if a field has been set.
func (o *ComputePersistentMemoryOperationAllOf) HasSecurePassphrase() bool {
	if o != nil && o.SecurePassphrase != nil {
		return true
	}

	return false
}

// SetSecurePassphrase gets a reference to the given string and assigns it to the SecurePassphrase field.
func (o *ComputePersistentMemoryOperationAllOf) SetSecurePassphrase(v string) {
	o.SecurePassphrase = &v
}

func (o ComputePersistentMemoryOperationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminAction != nil {
		toSerialize["AdminAction"] = o.AdminAction
	}
	if o.IsSecurePassphraseSet != nil {
		toSerialize["IsSecurePassphraseSet"] = o.IsSecurePassphraseSet
	}
	if o.Modules != nil {
		toSerialize["Modules"] = o.Modules
	}
	if o.SecurePassphrase != nil {
		toSerialize["SecurePassphrase"] = o.SecurePassphrase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputePersistentMemoryOperationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputePersistentMemoryOperationAllOf := _ComputePersistentMemoryOperationAllOf{}

	if err = json.Unmarshal(bytes, &varComputePersistentMemoryOperationAllOf); err == nil {
		*o = ComputePersistentMemoryOperationAllOf(varComputePersistentMemoryOperationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminAction")
		delete(additionalProperties, "IsSecurePassphraseSet")
		delete(additionalProperties, "Modules")
		delete(additionalProperties, "SecurePassphrase")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputePersistentMemoryOperationAllOf struct {
	value *ComputePersistentMemoryOperationAllOf
	isSet bool
}

func (v NullableComputePersistentMemoryOperationAllOf) Get() *ComputePersistentMemoryOperationAllOf {
	return v.value
}

func (v *NullableComputePersistentMemoryOperationAllOf) Set(val *ComputePersistentMemoryOperationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputePersistentMemoryOperationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputePersistentMemoryOperationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputePersistentMemoryOperationAllOf(val *ComputePersistentMemoryOperationAllOf) *NullableComputePersistentMemoryOperationAllOf {
	return &NullableComputePersistentMemoryOperationAllOf{value: val, isSet: true}
}

func (v NullableComputePersistentMemoryOperationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputePersistentMemoryOperationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
