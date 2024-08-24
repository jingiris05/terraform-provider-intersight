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

// checks if the IqnpoolLease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IqnpoolLease{}

// IqnpoolLease Lease represents a single IQN address that is part of the universe, allocated either from a pool or through static assignment.
type IqnpoolLease struct {
	PoolAbstractLease
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IQN address allocated for pool-based allocation. It is constructed as <prefix>:<suffix>:<number>.
	IqnAddress *string `json:"IqnAddress,omitempty" validate:"regexp=^$|^(?:iqn\\\\.[0-9]{4}-[0-9]{2}(?:\\\\.[A-Za-z](?:[A-Za-z0-9\\\\-]*[A-Za-z0-9])?)+(?::.*)?|eui\\\\.[0-9A-Fa-f]{16})"`
	// Number of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
	IqnNumber *int64 `json:"IqnNumber,omitempty"`
	// Prefix of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
	IqnPrefix *string `json:"IqnPrefix,omitempty"`
	// Suffix of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
	IqnSuffix            *string                               `json:"IqnSuffix,omitempty"`
	Reservation          *IqnpoolReservationReference          `json:"Reservation,omitempty"`
	AssignedToEntity     NullableMoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
	Pool                 NullableIqnpoolPoolRelationship       `json:"Pool,omitempty"`
	PoolMember           NullableIqnpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
	Universe             NullableIqnpoolUniverseRelationship   `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IqnpoolLease IqnpoolLease

// NewIqnpoolLease instantiates a new IqnpoolLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqnpoolLease(classId string, objectType string) *IqnpoolLease {
	this := IqnpoolLease{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	var hasDuplicate bool = false
	this.HasDuplicate = &hasDuplicate
	return &this
}

// NewIqnpoolLeaseWithDefaults instantiates a new IqnpoolLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqnpoolLeaseWithDefaults() *IqnpoolLease {
	this := IqnpoolLease{}
	var classId string = "iqnpool.Lease"
	this.ClassId = classId
	var objectType string = "iqnpool.Lease"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IqnpoolLease) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IqnpoolLease) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IqnpoolLease) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iqnpool.Lease" of the ClassId field.
func (o *IqnpoolLease) GetDefaultClassId() interface{} {
	return "iqnpool.Lease"
}

// GetObjectType returns the ObjectType field value
func (o *IqnpoolLease) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IqnpoolLease) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IqnpoolLease) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iqnpool.Lease" of the ObjectType field.
func (o *IqnpoolLease) GetDefaultObjectType() interface{} {
	return "iqnpool.Lease"
}

// GetIqnAddress returns the IqnAddress field value if set, zero value otherwise.
func (o *IqnpoolLease) GetIqnAddress() string {
	if o == nil || IsNil(o.IqnAddress) {
		var ret string
		return ret
	}
	return *o.IqnAddress
}

// GetIqnAddressOk returns a tuple with the IqnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolLease) GetIqnAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IqnAddress) {
		return nil, false
	}
	return o.IqnAddress, true
}

// HasIqnAddress returns a boolean if a field has been set.
func (o *IqnpoolLease) HasIqnAddress() bool {
	if o != nil && !IsNil(o.IqnAddress) {
		return true
	}

	return false
}

// SetIqnAddress gets a reference to the given string and assigns it to the IqnAddress field.
func (o *IqnpoolLease) SetIqnAddress(v string) {
	o.IqnAddress = &v
}

// GetIqnNumber returns the IqnNumber field value if set, zero value otherwise.
func (o *IqnpoolLease) GetIqnNumber() int64 {
	if o == nil || IsNil(o.IqnNumber) {
		var ret int64
		return ret
	}
	return *o.IqnNumber
}

// GetIqnNumberOk returns a tuple with the IqnNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolLease) GetIqnNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.IqnNumber) {
		return nil, false
	}
	return o.IqnNumber, true
}

// HasIqnNumber returns a boolean if a field has been set.
func (o *IqnpoolLease) HasIqnNumber() bool {
	if o != nil && !IsNil(o.IqnNumber) {
		return true
	}

	return false
}

// SetIqnNumber gets a reference to the given int64 and assigns it to the IqnNumber field.
func (o *IqnpoolLease) SetIqnNumber(v int64) {
	o.IqnNumber = &v
}

// GetIqnPrefix returns the IqnPrefix field value if set, zero value otherwise.
func (o *IqnpoolLease) GetIqnPrefix() string {
	if o == nil || IsNil(o.IqnPrefix) {
		var ret string
		return ret
	}
	return *o.IqnPrefix
}

// GetIqnPrefixOk returns a tuple with the IqnPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolLease) GetIqnPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.IqnPrefix) {
		return nil, false
	}
	return o.IqnPrefix, true
}

// HasIqnPrefix returns a boolean if a field has been set.
func (o *IqnpoolLease) HasIqnPrefix() bool {
	if o != nil && !IsNil(o.IqnPrefix) {
		return true
	}

	return false
}

// SetIqnPrefix gets a reference to the given string and assigns it to the IqnPrefix field.
func (o *IqnpoolLease) SetIqnPrefix(v string) {
	o.IqnPrefix = &v
}

// GetIqnSuffix returns the IqnSuffix field value if set, zero value otherwise.
func (o *IqnpoolLease) GetIqnSuffix() string {
	if o == nil || IsNil(o.IqnSuffix) {
		var ret string
		return ret
	}
	return *o.IqnSuffix
}

// GetIqnSuffixOk returns a tuple with the IqnSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolLease) GetIqnSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.IqnSuffix) {
		return nil, false
	}
	return o.IqnSuffix, true
}

// HasIqnSuffix returns a boolean if a field has been set.
func (o *IqnpoolLease) HasIqnSuffix() bool {
	if o != nil && !IsNil(o.IqnSuffix) {
		return true
	}

	return false
}

// SetIqnSuffix gets a reference to the given string and assigns it to the IqnSuffix field.
func (o *IqnpoolLease) SetIqnSuffix(v string) {
	o.IqnSuffix = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *IqnpoolLease) GetReservation() IqnpoolReservationReference {
	if o == nil || IsNil(o.Reservation) {
		var ret IqnpoolReservationReference
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolLease) GetReservationOk() (*IqnpoolReservationReference, bool) {
	if o == nil || IsNil(o.Reservation) {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *IqnpoolLease) HasReservation() bool {
	if o != nil && !IsNil(o.Reservation) {
		return true
	}

	return false
}

// SetReservation gets a reference to the given IqnpoolReservationReference and assigns it to the Reservation field.
func (o *IqnpoolLease) SetReservation(v IqnpoolReservationReference) {
	o.Reservation = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || IsNil(o.AssignedToEntity.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity.Get()
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedToEntity.Get(), o.AssignedToEntity.IsSet()
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *IqnpoolLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity.IsSet() {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given NullableMoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *IqnpoolLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity.Set(&v)
}

// SetAssignedToEntityNil sets the value for AssignedToEntity to be an explicit nil
func (o *IqnpoolLease) SetAssignedToEntityNil() {
	o.AssignedToEntity.Set(nil)
}

// UnsetAssignedToEntity ensures that no value is present for AssignedToEntity, not even an explicit nil
func (o *IqnpoolLease) UnsetAssignedToEntity() {
	o.AssignedToEntity.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolLease) GetPool() IqnpoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret IqnpoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolLease) GetPoolOk() (*IqnpoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *IqnpoolLease) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableIqnpoolPoolRelationship and assigns it to the Pool field.
func (o *IqnpoolLease) SetPool(v IqnpoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *IqnpoolLease) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *IqnpoolLease) UnsetPool() {
	o.Pool.Unset()
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolLease) GetPoolMember() IqnpoolPoolMemberRelationship {
	if o == nil || IsNil(o.PoolMember.Get()) {
		var ret IqnpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember.Get()
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolLease) GetPoolMemberOk() (*IqnpoolPoolMemberRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolMember.Get(), o.PoolMember.IsSet()
}

// HasPoolMember returns a boolean if a field has been set.
func (o *IqnpoolLease) HasPoolMember() bool {
	if o != nil && o.PoolMember.IsSet() {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given NullableIqnpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *IqnpoolLease) SetPoolMember(v IqnpoolPoolMemberRelationship) {
	o.PoolMember.Set(&v)
}

// SetPoolMemberNil sets the value for PoolMember to be an explicit nil
func (o *IqnpoolLease) SetPoolMemberNil() {
	o.PoolMember.Set(nil)
}

// UnsetPoolMember ensures that no value is present for PoolMember, not even an explicit nil
func (o *IqnpoolLease) UnsetPoolMember() {
	o.PoolMember.Unset()
}

// GetUniverse returns the Universe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IqnpoolLease) GetUniverse() IqnpoolUniverseRelationship {
	if o == nil || IsNil(o.Universe.Get()) {
		var ret IqnpoolUniverseRelationship
		return ret
	}
	return *o.Universe.Get()
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IqnpoolLease) GetUniverseOk() (*IqnpoolUniverseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Universe.Get(), o.Universe.IsSet()
}

// HasUniverse returns a boolean if a field has been set.
func (o *IqnpoolLease) HasUniverse() bool {
	if o != nil && o.Universe.IsSet() {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given NullableIqnpoolUniverseRelationship and assigns it to the Universe field.
func (o *IqnpoolLease) SetUniverse(v IqnpoolUniverseRelationship) {
	o.Universe.Set(&v)
}

// SetUniverseNil sets the value for Universe to be an explicit nil
func (o *IqnpoolLease) SetUniverseNil() {
	o.Universe.Set(nil)
}

// UnsetUniverse ensures that no value is present for Universe, not even an explicit nil
func (o *IqnpoolLease) UnsetUniverse() {
	o.Universe.Unset()
}

func (o IqnpoolLease) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IqnpoolLease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractLease, errPoolAbstractLease := json.Marshal(o.PoolAbstractLease)
	if errPoolAbstractLease != nil {
		return map[string]interface{}{}, errPoolAbstractLease
	}
	errPoolAbstractLease = json.Unmarshal([]byte(serializedPoolAbstractLease), &toSerialize)
	if errPoolAbstractLease != nil {
		return map[string]interface{}{}, errPoolAbstractLease
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.IqnAddress) {
		toSerialize["IqnAddress"] = o.IqnAddress
	}
	if !IsNil(o.IqnNumber) {
		toSerialize["IqnNumber"] = o.IqnNumber
	}
	if !IsNil(o.IqnPrefix) {
		toSerialize["IqnPrefix"] = o.IqnPrefix
	}
	if !IsNil(o.IqnSuffix) {
		toSerialize["IqnSuffix"] = o.IqnSuffix
	}
	if !IsNil(o.Reservation) {
		toSerialize["Reservation"] = o.Reservation
	}
	if o.AssignedToEntity.IsSet() {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity.Get()
	}
	if o.Pool.IsSet() {
		toSerialize["Pool"] = o.Pool.Get()
	}
	if o.PoolMember.IsSet() {
		toSerialize["PoolMember"] = o.PoolMember.Get()
	}
	if o.Universe.IsSet() {
		toSerialize["Universe"] = o.Universe.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IqnpoolLease) UnmarshalJSON(data []byte) (err error) {
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
	type IqnpoolLeaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IQN address allocated for pool-based allocation. It is constructed as <prefix>:<suffix>:<number>.
		IqnAddress *string `json:"IqnAddress,omitempty" validate:"regexp=^$|^(?:iqn\\\\.[0-9]{4}-[0-9]{2}(?:\\\\.[A-Za-z](?:[A-Za-z0-9\\\\-]*[A-Za-z0-9])?)+(?::.*)?|eui\\\\.[0-9A-Fa-f]{16})"`
		// Number of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
		IqnNumber *int64 `json:"IqnNumber,omitempty"`
		// Prefix of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
		IqnPrefix *string `json:"IqnPrefix,omitempty"`
		// Suffix of the IQN address. IQN Address is constructed as <prefix>:<suffix>:<number>.
		IqnSuffix        *string                               `json:"IqnSuffix,omitempty"`
		Reservation      *IqnpoolReservationReference          `json:"Reservation,omitempty"`
		AssignedToEntity NullableMoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
		Pool             NullableIqnpoolPoolRelationship       `json:"Pool,omitempty"`
		PoolMember       NullableIqnpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
		Universe         NullableIqnpoolUniverseRelationship   `json:"Universe,omitempty"`
	}

	varIqnpoolLeaseWithoutEmbeddedStruct := IqnpoolLeaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIqnpoolLeaseWithoutEmbeddedStruct)
	if err == nil {
		varIqnpoolLease := _IqnpoolLease{}
		varIqnpoolLease.ClassId = varIqnpoolLeaseWithoutEmbeddedStruct.ClassId
		varIqnpoolLease.ObjectType = varIqnpoolLeaseWithoutEmbeddedStruct.ObjectType
		varIqnpoolLease.IqnAddress = varIqnpoolLeaseWithoutEmbeddedStruct.IqnAddress
		varIqnpoolLease.IqnNumber = varIqnpoolLeaseWithoutEmbeddedStruct.IqnNumber
		varIqnpoolLease.IqnPrefix = varIqnpoolLeaseWithoutEmbeddedStruct.IqnPrefix
		varIqnpoolLease.IqnSuffix = varIqnpoolLeaseWithoutEmbeddedStruct.IqnSuffix
		varIqnpoolLease.Reservation = varIqnpoolLeaseWithoutEmbeddedStruct.Reservation
		varIqnpoolLease.AssignedToEntity = varIqnpoolLeaseWithoutEmbeddedStruct.AssignedToEntity
		varIqnpoolLease.Pool = varIqnpoolLeaseWithoutEmbeddedStruct.Pool
		varIqnpoolLease.PoolMember = varIqnpoolLeaseWithoutEmbeddedStruct.PoolMember
		varIqnpoolLease.Universe = varIqnpoolLeaseWithoutEmbeddedStruct.Universe
		*o = IqnpoolLease(varIqnpoolLease)
	} else {
		return err
	}

	varIqnpoolLease := _IqnpoolLease{}

	err = json.Unmarshal(data, &varIqnpoolLease)
	if err == nil {
		o.PoolAbstractLease = varIqnpoolLease.PoolAbstractLease
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IqnAddress")
		delete(additionalProperties, "IqnNumber")
		delete(additionalProperties, "IqnPrefix")
		delete(additionalProperties, "IqnSuffix")
		delete(additionalProperties, "Reservation")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")

		// remove fields from embedded structs
		reflectPoolAbstractLease := reflect.ValueOf(o.PoolAbstractLease)
		for i := 0; i < reflectPoolAbstractLease.Type().NumField(); i++ {
			t := reflectPoolAbstractLease.Type().Field(i)

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

type NullableIqnpoolLease struct {
	value *IqnpoolLease
	isSet bool
}

func (v NullableIqnpoolLease) Get() *IqnpoolLease {
	return v.value
}

func (v *NullableIqnpoolLease) Set(val *IqnpoolLease) {
	v.value = val
	v.isSet = true
}

func (v NullableIqnpoolLease) IsSet() bool {
	return v.isSet
}

func (v *NullableIqnpoolLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqnpoolLease(val *IqnpoolLease) *NullableIqnpoolLease {
	return &NullableIqnpoolLease{value: val, isSet: true}
}

func (v NullableIqnpoolLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqnpoolLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
