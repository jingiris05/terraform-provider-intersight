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

// VirtualizationVmwareVmCpuShareInfoAllOf Definition of the list of properties defined in 'virtualization.VmwareVmCpuShareInfo', excluding properties defined in parent classes.
type VirtualizationVmwareVmCpuShareInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Upper limit on CPU allocation (MHz).
	CpuLimit *int64 `json:"CpuLimit,omitempty"`
	// Amount of CPU for virtualization overhead.
	CpuOverheadLimit *int64 `json:"CpuOverheadLimit,omitempty"`
	// Guaranteed minimum allocation of CPU resource (MHz).
	CpuReservation *int64 `json:"CpuReservation,omitempty"`
	// Shows the relative importance of a VM. There is no unit for this value. It is a relative measure based on the settings for other resource pools. For more information, see VMware documentation.
	CpuShares            *int64 `json:"CpuShares,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVmCpuShareInfoAllOf VirtualizationVmwareVmCpuShareInfoAllOf

// NewVirtualizationVmwareVmCpuShareInfoAllOf instantiates a new VirtualizationVmwareVmCpuShareInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVmCpuShareInfoAllOf(classId string, objectType string) *VirtualizationVmwareVmCpuShareInfoAllOf {
	this := VirtualizationVmwareVmCpuShareInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVmCpuShareInfoAllOfWithDefaults instantiates a new VirtualizationVmwareVmCpuShareInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVmCpuShareInfoAllOfWithDefaults() *VirtualizationVmwareVmCpuShareInfoAllOf {
	this := VirtualizationVmwareVmCpuShareInfoAllOf{}
	var classId string = "virtualization.VmwareVmCpuShareInfo"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVmCpuShareInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpuLimit returns the CpuLimit field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetCpuLimit() int64 {
	if o == nil || o.CpuLimit == nil {
		var ret int64
		return ret
	}
	return *o.CpuLimit
}

// GetCpuLimitOk returns a tuple with the CpuLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetCpuLimitOk() (*int64, bool) {
	if o == nil || o.CpuLimit == nil {
		return nil, false
	}
	return o.CpuLimit, true
}

// HasCpuLimit returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) HasCpuLimit() bool {
	if o != nil && o.CpuLimit != nil {
		return true
	}

	return false
}

// SetCpuLimit gets a reference to the given int64 and assigns it to the CpuLimit field.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) SetCpuLimit(v int64) {
	o.CpuLimit = &v
}

// GetCpuOverheadLimit returns the CpuOverheadLimit field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetCpuOverheadLimit() int64 {
	if o == nil || o.CpuOverheadLimit == nil {
		var ret int64
		return ret
	}
	return *o.CpuOverheadLimit
}

// GetCpuOverheadLimitOk returns a tuple with the CpuOverheadLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetCpuOverheadLimitOk() (*int64, bool) {
	if o == nil || o.CpuOverheadLimit == nil {
		return nil, false
	}
	return o.CpuOverheadLimit, true
}

// HasCpuOverheadLimit returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) HasCpuOverheadLimit() bool {
	if o != nil && o.CpuOverheadLimit != nil {
		return true
	}

	return false
}

// SetCpuOverheadLimit gets a reference to the given int64 and assigns it to the CpuOverheadLimit field.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) SetCpuOverheadLimit(v int64) {
	o.CpuOverheadLimit = &v
}

// GetCpuReservation returns the CpuReservation field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetCpuReservation() int64 {
	if o == nil || o.CpuReservation == nil {
		var ret int64
		return ret
	}
	return *o.CpuReservation
}

// GetCpuReservationOk returns a tuple with the CpuReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetCpuReservationOk() (*int64, bool) {
	if o == nil || o.CpuReservation == nil {
		return nil, false
	}
	return o.CpuReservation, true
}

// HasCpuReservation returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) HasCpuReservation() bool {
	if o != nil && o.CpuReservation != nil {
		return true
	}

	return false
}

// SetCpuReservation gets a reference to the given int64 and assigns it to the CpuReservation field.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) SetCpuReservation(v int64) {
	o.CpuReservation = &v
}

// GetCpuShares returns the CpuShares field value if set, zero value otherwise.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetCpuShares() int64 {
	if o == nil || o.CpuShares == nil {
		var ret int64
		return ret
	}
	return *o.CpuShares
}

// GetCpuSharesOk returns a tuple with the CpuShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) GetCpuSharesOk() (*int64, bool) {
	if o == nil || o.CpuShares == nil {
		return nil, false
	}
	return o.CpuShares, true
}

// HasCpuShares returns a boolean if a field has been set.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) HasCpuShares() bool {
	if o != nil && o.CpuShares != nil {
		return true
	}

	return false
}

// SetCpuShares gets a reference to the given int64 and assigns it to the CpuShares field.
func (o *VirtualizationVmwareVmCpuShareInfoAllOf) SetCpuShares(v int64) {
	o.CpuShares = &v
}

func (o VirtualizationVmwareVmCpuShareInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CpuLimit != nil {
		toSerialize["CpuLimit"] = o.CpuLimit
	}
	if o.CpuOverheadLimit != nil {
		toSerialize["CpuOverheadLimit"] = o.CpuOverheadLimit
	}
	if o.CpuReservation != nil {
		toSerialize["CpuReservation"] = o.CpuReservation
	}
	if o.CpuShares != nil {
		toSerialize["CpuShares"] = o.CpuShares
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVmCpuShareInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareVmCpuShareInfoAllOf := _VirtualizationVmwareVmCpuShareInfoAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareVmCpuShareInfoAllOf); err == nil {
		*o = VirtualizationVmwareVmCpuShareInfoAllOf(varVirtualizationVmwareVmCpuShareInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CpuLimit")
		delete(additionalProperties, "CpuOverheadLimit")
		delete(additionalProperties, "CpuReservation")
		delete(additionalProperties, "CpuShares")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareVmCpuShareInfoAllOf struct {
	value *VirtualizationVmwareVmCpuShareInfoAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareVmCpuShareInfoAllOf) Get() *VirtualizationVmwareVmCpuShareInfoAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareVmCpuShareInfoAllOf) Set(val *VirtualizationVmwareVmCpuShareInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVmCpuShareInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVmCpuShareInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVmCpuShareInfoAllOf(val *VirtualizationVmwareVmCpuShareInfoAllOf) *NullableVirtualizationVmwareVmCpuShareInfoAllOf {
	return &NullableVirtualizationVmwareVmCpuShareInfoAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVmCpuShareInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVmCpuShareInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
