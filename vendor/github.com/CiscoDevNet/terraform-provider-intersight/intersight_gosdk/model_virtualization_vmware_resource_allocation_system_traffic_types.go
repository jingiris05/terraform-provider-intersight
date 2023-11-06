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

// VirtualizationVmwareResourceAllocationSystemTrafficTypes Bandwidth Allocation for System Traffic.
type VirtualizationVmwareResourceAllocationSystemTrafficTypes struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The maximum allowed usage for a traffic class belonging to this resource pool per host physical NIC. The utilization of a traffic class will not exceed the specified limit even if there are available network resources. If this value is unset or set to -1 in an update operation, then there is no limit on the network resource usage (only bounded by available resource and shares). Units are in Mbits/sec.
	Limit *int64 `json:"Limit,omitempty"`
	// Amount of bandwidth resource that is guaranteed available to the host infrastructure traffic class. If the utilization is less than the reservation, the extra bandwidth is used for other host infrastructure traffic class types.
	Reservation *int64 `json:"Reservation,omitempty"`
	// Network share. The value is used as a relative weight in competing for shared bandwidth, in case of resource contention. * `low` - Share of the traffic in the overall flow through a physical adapter. * `high` - Share of the traffic in the overall flow through a physical adapter. * `normal` - Share of the traffic in the overall flow through a physical adapter. * `custom` - Share of the traffic in the overall flow through a physical adapter.
	Shares *string `json:"Shares,omitempty"`
	// The number of shares allocated. Used to determine resource allocation in case of resource contention. Shares value is only set if level is set to custom. If level is not set to custom, this value is ignored. Therefore, only shares with custom values can be compared.
	SharesValue *int32 `json:"SharesValue,omitempty"`
	// Key of the host infrastructure resource.
	TrafficType          *string `json:"TrafficType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareResourceAllocationSystemTrafficTypes VirtualizationVmwareResourceAllocationSystemTrafficTypes

// NewVirtualizationVmwareResourceAllocationSystemTrafficTypes instantiates a new VirtualizationVmwareResourceAllocationSystemTrafficTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareResourceAllocationSystemTrafficTypes(classId string, objectType string) *VirtualizationVmwareResourceAllocationSystemTrafficTypes {
	this := VirtualizationVmwareResourceAllocationSystemTrafficTypes{}
	this.ClassId = classId
	this.ObjectType = objectType
	var shares string = "low"
	this.Shares = &shares
	return &this
}

// NewVirtualizationVmwareResourceAllocationSystemTrafficTypesWithDefaults instantiates a new VirtualizationVmwareResourceAllocationSystemTrafficTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareResourceAllocationSystemTrafficTypesWithDefaults() *VirtualizationVmwareResourceAllocationSystemTrafficTypes {
	this := VirtualizationVmwareResourceAllocationSystemTrafficTypes{}
	var classId string = "virtualization.VmwareResourceAllocationSystemTrafficTypes"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareResourceAllocationSystemTrafficTypes"
	this.ObjectType = objectType
	var shares string = "low"
	this.Shares = &shares
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) SetLimit(v int64) {
	o.Limit = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetReservation() int64 {
	if o == nil || o.Reservation == nil {
		var ret int64
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetReservationOk() (*int64, bool) {
	if o == nil || o.Reservation == nil {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) HasReservation() bool {
	if o != nil && o.Reservation != nil {
		return true
	}

	return false
}

// SetReservation gets a reference to the given int64 and assigns it to the Reservation field.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) SetReservation(v int64) {
	o.Reservation = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetShares() string {
	if o == nil || o.Shares == nil {
		var ret string
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetSharesOk() (*string, bool) {
	if o == nil || o.Shares == nil {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) HasShares() bool {
	if o != nil && o.Shares != nil {
		return true
	}

	return false
}

// SetShares gets a reference to the given string and assigns it to the Shares field.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) SetShares(v string) {
	o.Shares = &v
}

// GetSharesValue returns the SharesValue field value if set, zero value otherwise.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetSharesValue() int32 {
	if o == nil || o.SharesValue == nil {
		var ret int32
		return ret
	}
	return *o.SharesValue
}

// GetSharesValueOk returns a tuple with the SharesValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetSharesValueOk() (*int32, bool) {
	if o == nil || o.SharesValue == nil {
		return nil, false
	}
	return o.SharesValue, true
}

// HasSharesValue returns a boolean if a field has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) HasSharesValue() bool {
	if o != nil && o.SharesValue != nil {
		return true
	}

	return false
}

// SetSharesValue gets a reference to the given int32 and assigns it to the SharesValue field.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) SetSharesValue(v int32) {
	o.SharesValue = &v
}

// GetTrafficType returns the TrafficType field value if set, zero value otherwise.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetTrafficType() string {
	if o == nil || o.TrafficType == nil {
		var ret string
		return ret
	}
	return *o.TrafficType
}

// GetTrafficTypeOk returns a tuple with the TrafficType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) GetTrafficTypeOk() (*string, bool) {
	if o == nil || o.TrafficType == nil {
		return nil, false
	}
	return o.TrafficType, true
}

// HasTrafficType returns a boolean if a field has been set.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) HasTrafficType() bool {
	if o != nil && o.TrafficType != nil {
		return true
	}

	return false
}

// SetTrafficType gets a reference to the given string and assigns it to the TrafficType field.
func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) SetTrafficType(v string) {
	o.TrafficType = &v
}

func (o VirtualizationVmwareResourceAllocationSystemTrafficTypes) MarshalJSON() ([]byte, error) {
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
	if o.Limit != nil {
		toSerialize["Limit"] = o.Limit
	}
	if o.Reservation != nil {
		toSerialize["Reservation"] = o.Reservation
	}
	if o.Shares != nil {
		toSerialize["Shares"] = o.Shares
	}
	if o.SharesValue != nil {
		toSerialize["SharesValue"] = o.SharesValue
	}
	if o.TrafficType != nil {
		toSerialize["TrafficType"] = o.TrafficType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareResourceAllocationSystemTrafficTypes) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The maximum allowed usage for a traffic class belonging to this resource pool per host physical NIC. The utilization of a traffic class will not exceed the specified limit even if there are available network resources. If this value is unset or set to -1 in an update operation, then there is no limit on the network resource usage (only bounded by available resource and shares). Units are in Mbits/sec.
		Limit *int64 `json:"Limit,omitempty"`
		// Amount of bandwidth resource that is guaranteed available to the host infrastructure traffic class. If the utilization is less than the reservation, the extra bandwidth is used for other host infrastructure traffic class types.
		Reservation *int64 `json:"Reservation,omitempty"`
		// Network share. The value is used as a relative weight in competing for shared bandwidth, in case of resource contention. * `low` - Share of the traffic in the overall flow through a physical adapter. * `high` - Share of the traffic in the overall flow through a physical adapter. * `normal` - Share of the traffic in the overall flow through a physical adapter. * `custom` - Share of the traffic in the overall flow through a physical adapter.
		Shares *string `json:"Shares,omitempty"`
		// The number of shares allocated. Used to determine resource allocation in case of resource contention. Shares value is only set if level is set to custom. If level is not set to custom, this value is ignored. Therefore, only shares with custom values can be compared.
		SharesValue *int32 `json:"SharesValue,omitempty"`
		// Key of the host infrastructure resource.
		TrafficType *string `json:"TrafficType,omitempty"`
	}

	varVirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct := VirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareResourceAllocationSystemTrafficTypes := _VirtualizationVmwareResourceAllocationSystemTrafficTypes{}
		varVirtualizationVmwareResourceAllocationSystemTrafficTypes.ClassId = varVirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareResourceAllocationSystemTrafficTypes.ObjectType = varVirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareResourceAllocationSystemTrafficTypes.Limit = varVirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct.Limit
		varVirtualizationVmwareResourceAllocationSystemTrafficTypes.Reservation = varVirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct.Reservation
		varVirtualizationVmwareResourceAllocationSystemTrafficTypes.Shares = varVirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct.Shares
		varVirtualizationVmwareResourceAllocationSystemTrafficTypes.SharesValue = varVirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct.SharesValue
		varVirtualizationVmwareResourceAllocationSystemTrafficTypes.TrafficType = varVirtualizationVmwareResourceAllocationSystemTrafficTypesWithoutEmbeddedStruct.TrafficType
		*o = VirtualizationVmwareResourceAllocationSystemTrafficTypes(varVirtualizationVmwareResourceAllocationSystemTrafficTypes)
	} else {
		return err
	}

	varVirtualizationVmwareResourceAllocationSystemTrafficTypes := _VirtualizationVmwareResourceAllocationSystemTrafficTypes{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareResourceAllocationSystemTrafficTypes)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVmwareResourceAllocationSystemTrafficTypes.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Limit")
		delete(additionalProperties, "Reservation")
		delete(additionalProperties, "Shares")
		delete(additionalProperties, "SharesValue")
		delete(additionalProperties, "TrafficType")

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

type NullableVirtualizationVmwareResourceAllocationSystemTrafficTypes struct {
	value *VirtualizationVmwareResourceAllocationSystemTrafficTypes
	isSet bool
}

func (v NullableVirtualizationVmwareResourceAllocationSystemTrafficTypes) Get() *VirtualizationVmwareResourceAllocationSystemTrafficTypes {
	return v.value
}

func (v *NullableVirtualizationVmwareResourceAllocationSystemTrafficTypes) Set(val *VirtualizationVmwareResourceAllocationSystemTrafficTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareResourceAllocationSystemTrafficTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareResourceAllocationSystemTrafficTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareResourceAllocationSystemTrafficTypes(val *VirtualizationVmwareResourceAllocationSystemTrafficTypes) *NullableVirtualizationVmwareResourceAllocationSystemTrafficTypes {
	return &NullableVirtualizationVmwareResourceAllocationSystemTrafficTypes{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareResourceAllocationSystemTrafficTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareResourceAllocationSystemTrafficTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
