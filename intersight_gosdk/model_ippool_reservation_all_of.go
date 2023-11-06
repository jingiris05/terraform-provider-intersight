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
)

// IppoolReservationAllOf Definition of the list of properties defined in 'ippool.Reservation', excluding properties defined in parent classes.
type IppoolReservationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP identity to be reserved.
	Identity *string `json:"Identity,omitempty"`
	// Type of the IP address that needs to be reserved. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
	IpType *string `json:"IpType,omitempty"`
	// IPv4 Address to be reserved.
	IpV4Address *string `json:"IpV4Address,omitempty"`
	// IPv6 Address to be reserved.
	IpV6Address *string `json:"IpV6Address,omitempty"`
	// The moid of the Virtual Routing and Forwarding MO.
	VrfMoid              *string                               `json:"VrfMoid,omitempty"`
	BlockHead            *IppoolShadowBlockRelationship        `json:"BlockHead,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Pool                 *IppoolPoolRelationship               `json:"Pool,omitempty"`
	PoolMember           *IppoolPoolMemberRelationship         `json:"PoolMember,omitempty"`
	ShadowPool           *IppoolShadowPoolRelationship         `json:"ShadowPool,omitempty"`
	Universe             *IppoolUniverseRelationship           `json:"Universe,omitempty"`
	Vrf                  *VrfVrfRelationship                   `json:"Vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolReservationAllOf IppoolReservationAllOf

// NewIppoolReservationAllOf instantiates a new IppoolReservationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolReservationAllOf(classId string, objectType string) *IppoolReservationAllOf {
	this := IppoolReservationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// NewIppoolReservationAllOfWithDefaults instantiates a new IppoolReservationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolReservationAllOfWithDefaults() *IppoolReservationAllOf {
	this := IppoolReservationAllOf{}
	var classId string = "ippool.Reservation"
	this.ClassId = classId
	var objectType string = "ippool.Reservation"
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolReservationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolReservationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IppoolReservationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolReservationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *IppoolReservationAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetIpType() string {
	if o == nil || o.IpType == nil {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetIpTypeOk() (*string, bool) {
	if o == nil || o.IpType == nil {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasIpType() bool {
	if o != nil && o.IpType != nil {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *IppoolReservationAllOf) SetIpType(v string) {
	o.IpType = &v
}

// GetIpV4Address returns the IpV4Address field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetIpV4Address() string {
	if o == nil || o.IpV4Address == nil {
		var ret string
		return ret
	}
	return *o.IpV4Address
}

// GetIpV4AddressOk returns a tuple with the IpV4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetIpV4AddressOk() (*string, bool) {
	if o == nil || o.IpV4Address == nil {
		return nil, false
	}
	return o.IpV4Address, true
}

// HasIpV4Address returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasIpV4Address() bool {
	if o != nil && o.IpV4Address != nil {
		return true
	}

	return false
}

// SetIpV4Address gets a reference to the given string and assigns it to the IpV4Address field.
func (o *IppoolReservationAllOf) SetIpV4Address(v string) {
	o.IpV4Address = &v
}

// GetIpV6Address returns the IpV6Address field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetIpV6Address() string {
	if o == nil || o.IpV6Address == nil {
		var ret string
		return ret
	}
	return *o.IpV6Address
}

// GetIpV6AddressOk returns a tuple with the IpV6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetIpV6AddressOk() (*string, bool) {
	if o == nil || o.IpV6Address == nil {
		return nil, false
	}
	return o.IpV6Address, true
}

// HasIpV6Address returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasIpV6Address() bool {
	if o != nil && o.IpV6Address != nil {
		return true
	}

	return false
}

// SetIpV6Address gets a reference to the given string and assigns it to the IpV6Address field.
func (o *IppoolReservationAllOf) SetIpV6Address(v string) {
	o.IpV6Address = &v
}

// GetVrfMoid returns the VrfMoid field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetVrfMoid() string {
	if o == nil || o.VrfMoid == nil {
		var ret string
		return ret
	}
	return *o.VrfMoid
}

// GetVrfMoidOk returns a tuple with the VrfMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetVrfMoidOk() (*string, bool) {
	if o == nil || o.VrfMoid == nil {
		return nil, false
	}
	return o.VrfMoid, true
}

// HasVrfMoid returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasVrfMoid() bool {
	if o != nil && o.VrfMoid != nil {
		return true
	}

	return false
}

// SetVrfMoid gets a reference to the given string and assigns it to the VrfMoid field.
func (o *IppoolReservationAllOf) SetVrfMoid(v string) {
	o.VrfMoid = &v
}

// GetBlockHead returns the BlockHead field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetBlockHead() IppoolShadowBlockRelationship {
	if o == nil || o.BlockHead == nil {
		var ret IppoolShadowBlockRelationship
		return ret
	}
	return *o.BlockHead
}

// GetBlockHeadOk returns a tuple with the BlockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetBlockHeadOk() (*IppoolShadowBlockRelationship, bool) {
	if o == nil || o.BlockHead == nil {
		return nil, false
	}
	return o.BlockHead, true
}

// HasBlockHead returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasBlockHead() bool {
	if o != nil && o.BlockHead != nil {
		return true
	}

	return false
}

// SetBlockHead gets a reference to the given IppoolShadowBlockRelationship and assigns it to the BlockHead field.
func (o *IppoolReservationAllOf) SetBlockHead(v IppoolShadowBlockRelationship) {
	o.BlockHead = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IppoolReservationAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetPool() IppoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolPoolRelationship and assigns it to the Pool field.
func (o *IppoolReservationAllOf) SetPool(v IppoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetPoolMember() IppoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret IppoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetPoolMemberOk() (*IppoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given IppoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *IppoolReservationAllOf) SetPoolMember(v IppoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetShadowPool returns the ShadowPool field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetShadowPool() IppoolShadowPoolRelationship {
	if o == nil || o.ShadowPool == nil {
		var ret IppoolShadowPoolRelationship
		return ret
	}
	return *o.ShadowPool
}

// GetShadowPoolOk returns a tuple with the ShadowPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetShadowPoolOk() (*IppoolShadowPoolRelationship, bool) {
	if o == nil || o.ShadowPool == nil {
		return nil, false
	}
	return o.ShadowPool, true
}

// HasShadowPool returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasShadowPool() bool {
	if o != nil && o.ShadowPool != nil {
		return true
	}

	return false
}

// SetShadowPool gets a reference to the given IppoolShadowPoolRelationship and assigns it to the ShadowPool field.
func (o *IppoolReservationAllOf) SetShadowPool(v IppoolShadowPoolRelationship) {
	o.ShadowPool = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetUniverse() IppoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret IppoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetUniverseOk() (*IppoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given IppoolUniverseRelationship and assigns it to the Universe field.
func (o *IppoolReservationAllOf) SetUniverse(v IppoolUniverseRelationship) {
	o.Universe = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *IppoolReservationAllOf) GetVrf() VrfVrfRelationship {
	if o == nil || o.Vrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolReservationAllOf) GetVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *IppoolReservationAllOf) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given VrfVrfRelationship and assigns it to the Vrf field.
func (o *IppoolReservationAllOf) SetVrf(v VrfVrfRelationship) {
	o.Vrf = &v
}

func (o IppoolReservationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.IpType != nil {
		toSerialize["IpType"] = o.IpType
	}
	if o.IpV4Address != nil {
		toSerialize["IpV4Address"] = o.IpV4Address
	}
	if o.IpV6Address != nil {
		toSerialize["IpV6Address"] = o.IpV6Address
	}
	if o.VrfMoid != nil {
		toSerialize["VrfMoid"] = o.VrfMoid
	}
	if o.BlockHead != nil {
		toSerialize["BlockHead"] = o.BlockHead
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.ShadowPool != nil {
		toSerialize["ShadowPool"] = o.ShadowPool
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}
	if o.Vrf != nil {
		toSerialize["Vrf"] = o.Vrf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolReservationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIppoolReservationAllOf := _IppoolReservationAllOf{}

	if err = json.Unmarshal(bytes, &varIppoolReservationAllOf); err == nil {
		*o = IppoolReservationAllOf(varIppoolReservationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "IpType")
		delete(additionalProperties, "IpV4Address")
		delete(additionalProperties, "IpV6Address")
		delete(additionalProperties, "VrfMoid")
		delete(additionalProperties, "BlockHead")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "ShadowPool")
		delete(additionalProperties, "Universe")
		delete(additionalProperties, "Vrf")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIppoolReservationAllOf struct {
	value *IppoolReservationAllOf
	isSet bool
}

func (v NullableIppoolReservationAllOf) Get() *IppoolReservationAllOf {
	return v.value
}

func (v *NullableIppoolReservationAllOf) Set(val *IppoolReservationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolReservationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolReservationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolReservationAllOf(val *IppoolReservationAllOf) *NullableIppoolReservationAllOf {
	return &NullableIppoolReservationAllOf{value: val, isSet: true}
}

func (v NullableIppoolReservationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolReservationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
