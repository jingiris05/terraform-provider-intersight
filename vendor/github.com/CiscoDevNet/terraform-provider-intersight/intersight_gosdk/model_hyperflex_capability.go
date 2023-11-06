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

// HyperflexCapability Capabilities and features supported in the HyperFlex cluster.
type HyperflexCapability struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies if encryption is supported in the HyperFlex cluster. The HyperFlex cluster supports self encrypting drives (SED).
	EncryptionSupported *bool `json:"EncryptionSupported,omitempty"`
	// Specifies if iSCSI is supported in the HyperFlex cluster. The HyperFlex cluster supports an iSCSI network, iSCSI initiator groups, targets, and logical unit number (LUN).
	IscsiSupported *bool `json:"IscsiSupported,omitempty"`
	// Specifies if replication is supported in the HyperFlex cluster.  The HyperFlex cluster supports 1:1 disaster recovery and N:1 backup.
	ReplicationSupported *bool `json:"ReplicationSupported,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexCapability HyperflexCapability

// NewHyperflexCapability instantiates a new HyperflexCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexCapability(classId string, objectType string) *HyperflexCapability {
	this := HyperflexCapability{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexCapabilityWithDefaults instantiates a new HyperflexCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexCapabilityWithDefaults() *HyperflexCapability {
	this := HyperflexCapability{}
	var classId string = "hyperflex.Capability"
	this.ClassId = classId
	var objectType string = "hyperflex.Capability"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexCapability) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexCapability) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexCapability) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexCapability) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexCapability) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexCapability) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEncryptionSupported returns the EncryptionSupported field value if set, zero value otherwise.
func (o *HyperflexCapability) GetEncryptionSupported() bool {
	if o == nil || o.EncryptionSupported == nil {
		var ret bool
		return ret
	}
	return *o.EncryptionSupported
}

// GetEncryptionSupportedOk returns a tuple with the EncryptionSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCapability) GetEncryptionSupportedOk() (*bool, bool) {
	if o == nil || o.EncryptionSupported == nil {
		return nil, false
	}
	return o.EncryptionSupported, true
}

// HasEncryptionSupported returns a boolean if a field has been set.
func (o *HyperflexCapability) HasEncryptionSupported() bool {
	if o != nil && o.EncryptionSupported != nil {
		return true
	}

	return false
}

// SetEncryptionSupported gets a reference to the given bool and assigns it to the EncryptionSupported field.
func (o *HyperflexCapability) SetEncryptionSupported(v bool) {
	o.EncryptionSupported = &v
}

// GetIscsiSupported returns the IscsiSupported field value if set, zero value otherwise.
func (o *HyperflexCapability) GetIscsiSupported() bool {
	if o == nil || o.IscsiSupported == nil {
		var ret bool
		return ret
	}
	return *o.IscsiSupported
}

// GetIscsiSupportedOk returns a tuple with the IscsiSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCapability) GetIscsiSupportedOk() (*bool, bool) {
	if o == nil || o.IscsiSupported == nil {
		return nil, false
	}
	return o.IscsiSupported, true
}

// HasIscsiSupported returns a boolean if a field has been set.
func (o *HyperflexCapability) HasIscsiSupported() bool {
	if o != nil && o.IscsiSupported != nil {
		return true
	}

	return false
}

// SetIscsiSupported gets a reference to the given bool and assigns it to the IscsiSupported field.
func (o *HyperflexCapability) SetIscsiSupported(v bool) {
	o.IscsiSupported = &v
}

// GetReplicationSupported returns the ReplicationSupported field value if set, zero value otherwise.
func (o *HyperflexCapability) GetReplicationSupported() bool {
	if o == nil || o.ReplicationSupported == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationSupported
}

// GetReplicationSupportedOk returns a tuple with the ReplicationSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexCapability) GetReplicationSupportedOk() (*bool, bool) {
	if o == nil || o.ReplicationSupported == nil {
		return nil, false
	}
	return o.ReplicationSupported, true
}

// HasReplicationSupported returns a boolean if a field has been set.
func (o *HyperflexCapability) HasReplicationSupported() bool {
	if o != nil && o.ReplicationSupported != nil {
		return true
	}

	return false
}

// SetReplicationSupported gets a reference to the given bool and assigns it to the ReplicationSupported field.
func (o *HyperflexCapability) SetReplicationSupported(v bool) {
	o.ReplicationSupported = &v
}

func (o HyperflexCapability) MarshalJSON() ([]byte, error) {
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
	if o.EncryptionSupported != nil {
		toSerialize["EncryptionSupported"] = o.EncryptionSupported
	}
	if o.IscsiSupported != nil {
		toSerialize["IscsiSupported"] = o.IscsiSupported
	}
	if o.ReplicationSupported != nil {
		toSerialize["ReplicationSupported"] = o.ReplicationSupported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexCapability) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexCapabilityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specifies if encryption is supported in the HyperFlex cluster. The HyperFlex cluster supports self encrypting drives (SED).
		EncryptionSupported *bool `json:"EncryptionSupported,omitempty"`
		// Specifies if iSCSI is supported in the HyperFlex cluster. The HyperFlex cluster supports an iSCSI network, iSCSI initiator groups, targets, and logical unit number (LUN).
		IscsiSupported *bool `json:"IscsiSupported,omitempty"`
		// Specifies if replication is supported in the HyperFlex cluster.  The HyperFlex cluster supports 1:1 disaster recovery and N:1 backup.
		ReplicationSupported *bool `json:"ReplicationSupported,omitempty"`
	}

	varHyperflexCapabilityWithoutEmbeddedStruct := HyperflexCapabilityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexCapabilityWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexCapability := _HyperflexCapability{}
		varHyperflexCapability.ClassId = varHyperflexCapabilityWithoutEmbeddedStruct.ClassId
		varHyperflexCapability.ObjectType = varHyperflexCapabilityWithoutEmbeddedStruct.ObjectType
		varHyperflexCapability.EncryptionSupported = varHyperflexCapabilityWithoutEmbeddedStruct.EncryptionSupported
		varHyperflexCapability.IscsiSupported = varHyperflexCapabilityWithoutEmbeddedStruct.IscsiSupported
		varHyperflexCapability.ReplicationSupported = varHyperflexCapabilityWithoutEmbeddedStruct.ReplicationSupported
		*o = HyperflexCapability(varHyperflexCapability)
	} else {
		return err
	}

	varHyperflexCapability := _HyperflexCapability{}

	err = json.Unmarshal(bytes, &varHyperflexCapability)
	if err == nil {
		o.MoBaseComplexType = varHyperflexCapability.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EncryptionSupported")
		delete(additionalProperties, "IscsiSupported")
		delete(additionalProperties, "ReplicationSupported")

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

type NullableHyperflexCapability struct {
	value *HyperflexCapability
	isSet bool
}

func (v NullableHyperflexCapability) Get() *HyperflexCapability {
	return v.value
}

func (v *NullableHyperflexCapability) Set(val *HyperflexCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexCapability(val *HyperflexCapability) *NullableHyperflexCapability {
	return &NullableHyperflexCapability{value: val, isSet: true}
}

func (v NullableHyperflexCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
