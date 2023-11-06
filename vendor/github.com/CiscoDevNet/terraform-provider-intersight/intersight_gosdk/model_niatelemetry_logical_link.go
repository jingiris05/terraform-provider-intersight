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

// NiatelemetryLogicalLink Stores logical links per fabric.
type NiatelemetryLogicalLink struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return value of dbId attribute.
	DbId *int64 `json:"DbId,omitempty"`
	// Return value of isPresent attribute.
	IsPresent *bool `json:"IsPresent,omitempty"`
	// Return value of linkAddr1 attribute.
	LinkAddr1 *string `json:"LinkAddr1,omitempty"`
	// Return value of linkAddr2 attribute.
	LinkAddr2 *string `json:"LinkAddr2,omitempty"`
	// Return value of linkState attribute.
	LinkState *string `json:"LinkState,omitempty"`
	// Return value of linkType attribute.
	LinkType *string `json:"LinkType,omitempty"`
	// Return value of uptime attribute.
	Uptime               *string `json:"Uptime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryLogicalLink NiatelemetryLogicalLink

// NewNiatelemetryLogicalLink instantiates a new NiatelemetryLogicalLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryLogicalLink(classId string, objectType string) *NiatelemetryLogicalLink {
	this := NiatelemetryLogicalLink{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryLogicalLinkWithDefaults instantiates a new NiatelemetryLogicalLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryLogicalLinkWithDefaults() *NiatelemetryLogicalLink {
	this := NiatelemetryLogicalLink{}
	var classId string = "niatelemetry.LogicalLink"
	this.ClassId = classId
	var objectType string = "niatelemetry.LogicalLink"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryLogicalLink) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLink) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryLogicalLink) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryLogicalLink) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLink) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryLogicalLink) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDbId returns the DbId field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLink) GetDbId() int64 {
	if o == nil || o.DbId == nil {
		var ret int64
		return ret
	}
	return *o.DbId
}

// GetDbIdOk returns a tuple with the DbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLink) GetDbIdOk() (*int64, bool) {
	if o == nil || o.DbId == nil {
		return nil, false
	}
	return o.DbId, true
}

// HasDbId returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLink) HasDbId() bool {
	if o != nil && o.DbId != nil {
		return true
	}

	return false
}

// SetDbId gets a reference to the given int64 and assigns it to the DbId field.
func (o *NiatelemetryLogicalLink) SetDbId(v int64) {
	o.DbId = &v
}

// GetIsPresent returns the IsPresent field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLink) GetIsPresent() bool {
	if o == nil || o.IsPresent == nil {
		var ret bool
		return ret
	}
	return *o.IsPresent
}

// GetIsPresentOk returns a tuple with the IsPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLink) GetIsPresentOk() (*bool, bool) {
	if o == nil || o.IsPresent == nil {
		return nil, false
	}
	return o.IsPresent, true
}

// HasIsPresent returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLink) HasIsPresent() bool {
	if o != nil && o.IsPresent != nil {
		return true
	}

	return false
}

// SetIsPresent gets a reference to the given bool and assigns it to the IsPresent field.
func (o *NiatelemetryLogicalLink) SetIsPresent(v bool) {
	o.IsPresent = &v
}

// GetLinkAddr1 returns the LinkAddr1 field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLink) GetLinkAddr1() string {
	if o == nil || o.LinkAddr1 == nil {
		var ret string
		return ret
	}
	return *o.LinkAddr1
}

// GetLinkAddr1Ok returns a tuple with the LinkAddr1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLink) GetLinkAddr1Ok() (*string, bool) {
	if o == nil || o.LinkAddr1 == nil {
		return nil, false
	}
	return o.LinkAddr1, true
}

// HasLinkAddr1 returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLink) HasLinkAddr1() bool {
	if o != nil && o.LinkAddr1 != nil {
		return true
	}

	return false
}

// SetLinkAddr1 gets a reference to the given string and assigns it to the LinkAddr1 field.
func (o *NiatelemetryLogicalLink) SetLinkAddr1(v string) {
	o.LinkAddr1 = &v
}

// GetLinkAddr2 returns the LinkAddr2 field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLink) GetLinkAddr2() string {
	if o == nil || o.LinkAddr2 == nil {
		var ret string
		return ret
	}
	return *o.LinkAddr2
}

// GetLinkAddr2Ok returns a tuple with the LinkAddr2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLink) GetLinkAddr2Ok() (*string, bool) {
	if o == nil || o.LinkAddr2 == nil {
		return nil, false
	}
	return o.LinkAddr2, true
}

// HasLinkAddr2 returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLink) HasLinkAddr2() bool {
	if o != nil && o.LinkAddr2 != nil {
		return true
	}

	return false
}

// SetLinkAddr2 gets a reference to the given string and assigns it to the LinkAddr2 field.
func (o *NiatelemetryLogicalLink) SetLinkAddr2(v string) {
	o.LinkAddr2 = &v
}

// GetLinkState returns the LinkState field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLink) GetLinkState() string {
	if o == nil || o.LinkState == nil {
		var ret string
		return ret
	}
	return *o.LinkState
}

// GetLinkStateOk returns a tuple with the LinkState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLink) GetLinkStateOk() (*string, bool) {
	if o == nil || o.LinkState == nil {
		return nil, false
	}
	return o.LinkState, true
}

// HasLinkState returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLink) HasLinkState() bool {
	if o != nil && o.LinkState != nil {
		return true
	}

	return false
}

// SetLinkState gets a reference to the given string and assigns it to the LinkState field.
func (o *NiatelemetryLogicalLink) SetLinkState(v string) {
	o.LinkState = &v
}

// GetLinkType returns the LinkType field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLink) GetLinkType() string {
	if o == nil || o.LinkType == nil {
		var ret string
		return ret
	}
	return *o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLink) GetLinkTypeOk() (*string, bool) {
	if o == nil || o.LinkType == nil {
		return nil, false
	}
	return o.LinkType, true
}

// HasLinkType returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLink) HasLinkType() bool {
	if o != nil && o.LinkType != nil {
		return true
	}

	return false
}

// SetLinkType gets a reference to the given string and assigns it to the LinkType field.
func (o *NiatelemetryLogicalLink) SetLinkType(v string) {
	o.LinkType = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLink) GetUptime() string {
	if o == nil || o.Uptime == nil {
		var ret string
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLink) GetUptimeOk() (*string, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLink) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given string and assigns it to the Uptime field.
func (o *NiatelemetryLogicalLink) SetUptime(v string) {
	o.Uptime = &v
}

func (o NiatelemetryLogicalLink) MarshalJSON() ([]byte, error) {
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
	if o.DbId != nil {
		toSerialize["DbId"] = o.DbId
	}
	if o.IsPresent != nil {
		toSerialize["IsPresent"] = o.IsPresent
	}
	if o.LinkAddr1 != nil {
		toSerialize["LinkAddr1"] = o.LinkAddr1
	}
	if o.LinkAddr2 != nil {
		toSerialize["LinkAddr2"] = o.LinkAddr2
	}
	if o.LinkState != nil {
		toSerialize["LinkState"] = o.LinkState
	}
	if o.LinkType != nil {
		toSerialize["LinkType"] = o.LinkType
	}
	if o.Uptime != nil {
		toSerialize["Uptime"] = o.Uptime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryLogicalLink) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryLogicalLinkWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Return value of dbId attribute.
		DbId *int64 `json:"DbId,omitempty"`
		// Return value of isPresent attribute.
		IsPresent *bool `json:"IsPresent,omitempty"`
		// Return value of linkAddr1 attribute.
		LinkAddr1 *string `json:"LinkAddr1,omitempty"`
		// Return value of linkAddr2 attribute.
		LinkAddr2 *string `json:"LinkAddr2,omitempty"`
		// Return value of linkState attribute.
		LinkState *string `json:"LinkState,omitempty"`
		// Return value of linkType attribute.
		LinkType *string `json:"LinkType,omitempty"`
		// Return value of uptime attribute.
		Uptime *string `json:"Uptime,omitempty"`
	}

	varNiatelemetryLogicalLinkWithoutEmbeddedStruct := NiatelemetryLogicalLinkWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryLogicalLinkWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryLogicalLink := _NiatelemetryLogicalLink{}
		varNiatelemetryLogicalLink.ClassId = varNiatelemetryLogicalLinkWithoutEmbeddedStruct.ClassId
		varNiatelemetryLogicalLink.ObjectType = varNiatelemetryLogicalLinkWithoutEmbeddedStruct.ObjectType
		varNiatelemetryLogicalLink.DbId = varNiatelemetryLogicalLinkWithoutEmbeddedStruct.DbId
		varNiatelemetryLogicalLink.IsPresent = varNiatelemetryLogicalLinkWithoutEmbeddedStruct.IsPresent
		varNiatelemetryLogicalLink.LinkAddr1 = varNiatelemetryLogicalLinkWithoutEmbeddedStruct.LinkAddr1
		varNiatelemetryLogicalLink.LinkAddr2 = varNiatelemetryLogicalLinkWithoutEmbeddedStruct.LinkAddr2
		varNiatelemetryLogicalLink.LinkState = varNiatelemetryLogicalLinkWithoutEmbeddedStruct.LinkState
		varNiatelemetryLogicalLink.LinkType = varNiatelemetryLogicalLinkWithoutEmbeddedStruct.LinkType
		varNiatelemetryLogicalLink.Uptime = varNiatelemetryLogicalLinkWithoutEmbeddedStruct.Uptime
		*o = NiatelemetryLogicalLink(varNiatelemetryLogicalLink)
	} else {
		return err
	}

	varNiatelemetryLogicalLink := _NiatelemetryLogicalLink{}

	err = json.Unmarshal(bytes, &varNiatelemetryLogicalLink)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryLogicalLink.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DbId")
		delete(additionalProperties, "IsPresent")
		delete(additionalProperties, "LinkAddr1")
		delete(additionalProperties, "LinkAddr2")
		delete(additionalProperties, "LinkState")
		delete(additionalProperties, "LinkType")
		delete(additionalProperties, "Uptime")

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

type NullableNiatelemetryLogicalLink struct {
	value *NiatelemetryLogicalLink
	isSet bool
}

func (v NullableNiatelemetryLogicalLink) Get() *NiatelemetryLogicalLink {
	return v.value
}

func (v *NullableNiatelemetryLogicalLink) Set(val *NiatelemetryLogicalLink) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryLogicalLink) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryLogicalLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryLogicalLink(val *NiatelemetryLogicalLink) *NullableNiatelemetryLogicalLink {
	return &NullableNiatelemetryLogicalLink{value: val, isSet: true}
}

func (v NullableNiatelemetryLogicalLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryLogicalLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
