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

// IamEndPointUserInventoryAllOf Definition of the list of properties defined in 'iam.EndPointUserInventory', excluding properties defined in parent classes.
type IamEndPointUserInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the user to be created on the endpoint. It can be any string that adheres to the following constraints. It can have alphanumeric characters, dots, underscores and hyphen. It cannot be more than 16 characters.
	Name *string `json:"Name,omitempty"`
	// UserId for the end point user.
	UserId *string `json:"UserId,omitempty"`
	// An array of relationships to iamEndPointUserRoleInventory resources.
	EndPointUserRole     []IamEndPointUserRoleInventoryRelationship `json:"EndPointUserRole,omitempty"`
	TargetMo             *MoBaseMoRelationship                      `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointUserInventoryAllOf IamEndPointUserInventoryAllOf

// NewIamEndPointUserInventoryAllOf instantiates a new IamEndPointUserInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointUserInventoryAllOf(classId string, objectType string) *IamEndPointUserInventoryAllOf {
	this := IamEndPointUserInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamEndPointUserInventoryAllOfWithDefaults instantiates a new IamEndPointUserInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointUserInventoryAllOfWithDefaults() *IamEndPointUserInventoryAllOf {
	this := IamEndPointUserInventoryAllOf{}
	var classId string = "iam.EndPointUserInventory"
	this.ClassId = classId
	var objectType string = "iam.EndPointUserInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamEndPointUserInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamEndPointUserInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamEndPointUserInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamEndPointUserInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamEndPointUserInventoryAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserInventoryAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamEndPointUserInventoryAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamEndPointUserInventoryAllOf) SetName(v string) {
	o.Name = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IamEndPointUserInventoryAllOf) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserInventoryAllOf) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IamEndPointUserInventoryAllOf) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IamEndPointUserInventoryAllOf) SetUserId(v string) {
	o.UserId = &v
}

// GetEndPointUserRole returns the EndPointUserRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserInventoryAllOf) GetEndPointUserRole() []IamEndPointUserRoleInventoryRelationship {
	if o == nil {
		var ret []IamEndPointUserRoleInventoryRelationship
		return ret
	}
	return o.EndPointUserRole
}

// GetEndPointUserRoleOk returns a tuple with the EndPointUserRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserInventoryAllOf) GetEndPointUserRoleOk() ([]IamEndPointUserRoleInventoryRelationship, bool) {
	if o == nil || o.EndPointUserRole == nil {
		return nil, false
	}
	return o.EndPointUserRole, true
}

// HasEndPointUserRole returns a boolean if a field has been set.
func (o *IamEndPointUserInventoryAllOf) HasEndPointUserRole() bool {
	if o != nil && o.EndPointUserRole != nil {
		return true
	}

	return false
}

// SetEndPointUserRole gets a reference to the given []IamEndPointUserRoleInventoryRelationship and assigns it to the EndPointUserRole field.
func (o *IamEndPointUserInventoryAllOf) SetEndPointUserRole(v []IamEndPointUserRoleInventoryRelationship) {
	o.EndPointUserRole = v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *IamEndPointUserInventoryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *IamEndPointUserInventoryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *IamEndPointUserInventoryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o IamEndPointUserInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.UserId != nil {
		toSerialize["UserId"] = o.UserId
	}
	if o.EndPointUserRole != nil {
		toSerialize["EndPointUserRole"] = o.EndPointUserRole
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamEndPointUserInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamEndPointUserInventoryAllOf := _IamEndPointUserInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varIamEndPointUserInventoryAllOf); err == nil {
		*o = IamEndPointUserInventoryAllOf(varIamEndPointUserInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "UserId")
		delete(additionalProperties, "EndPointUserRole")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamEndPointUserInventoryAllOf struct {
	value *IamEndPointUserInventoryAllOf
	isSet bool
}

func (v NullableIamEndPointUserInventoryAllOf) Get() *IamEndPointUserInventoryAllOf {
	return v.value
}

func (v *NullableIamEndPointUserInventoryAllOf) Set(val *IamEndPointUserInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserInventoryAllOf(val *IamEndPointUserInventoryAllOf) *NullableIamEndPointUserInventoryAllOf {
	return &NullableIamEndPointUserInventoryAllOf{value: val, isSet: true}
}

func (v NullableIamEndPointUserInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
