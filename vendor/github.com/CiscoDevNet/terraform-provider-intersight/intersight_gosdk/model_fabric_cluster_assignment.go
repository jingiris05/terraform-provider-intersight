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

// checks if the FabricClusterAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricClusterAssignment{}

// FabricClusterAssignment Network Elements that are to be assigned to the SwitchProfiles in the SwitchClusterProfile Mo. This will be honored only in the case of profile creation through the clone/derive operations.
type FabricClusterAssignment struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType     string   `json:"ObjectType"`
	NetworkElement *MoMoRef `json:"NetworkElement,omitempty"`
	// Name of the source SwitchProfile or SwitchProfileTemplate whose clone has to be assigned to the network element mentioned in NetworkElement property under ClusterAssignments.
	SourceSwitchProfileOrTemplateName *string `json:"SourceSwitchProfileOrTemplateName,omitempty" validate:"regexp=^$|^[a-zA-Z0-9_.-]{1,64}$"`
	AdditionalProperties              map[string]interface{}
}

type _FabricClusterAssignment FabricClusterAssignment

// NewFabricClusterAssignment instantiates a new FabricClusterAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricClusterAssignment(classId string, objectType string) *FabricClusterAssignment {
	this := FabricClusterAssignment{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricClusterAssignmentWithDefaults instantiates a new FabricClusterAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricClusterAssignmentWithDefaults() *FabricClusterAssignment {
	this := FabricClusterAssignment{}
	var classId string = "fabric.ClusterAssignment"
	this.ClassId = classId
	var objectType string = "fabric.ClusterAssignment"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricClusterAssignment) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricClusterAssignment) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricClusterAssignment) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.ClusterAssignment" of the ClassId field.
func (o *FabricClusterAssignment) GetDefaultClassId() interface{} {
	return "fabric.ClusterAssignment"
}

// GetObjectType returns the ObjectType field value
func (o *FabricClusterAssignment) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricClusterAssignment) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricClusterAssignment) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.ClusterAssignment" of the ObjectType field.
func (o *FabricClusterAssignment) GetDefaultObjectType() interface{} {
	return "fabric.ClusterAssignment"
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *FabricClusterAssignment) GetNetworkElement() MoMoRef {
	if o == nil || IsNil(o.NetworkElement) {
		var ret MoMoRef
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricClusterAssignment) GetNetworkElementOk() (*MoMoRef, bool) {
	if o == nil || IsNil(o.NetworkElement) {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *FabricClusterAssignment) HasNetworkElement() bool {
	if o != nil && !IsNil(o.NetworkElement) {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given MoMoRef and assigns it to the NetworkElement field.
func (o *FabricClusterAssignment) SetNetworkElement(v MoMoRef) {
	o.NetworkElement = &v
}

// GetSourceSwitchProfileOrTemplateName returns the SourceSwitchProfileOrTemplateName field value if set, zero value otherwise.
func (o *FabricClusterAssignment) GetSourceSwitchProfileOrTemplateName() string {
	if o == nil || IsNil(o.SourceSwitchProfileOrTemplateName) {
		var ret string
		return ret
	}
	return *o.SourceSwitchProfileOrTemplateName
}

// GetSourceSwitchProfileOrTemplateNameOk returns a tuple with the SourceSwitchProfileOrTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricClusterAssignment) GetSourceSwitchProfileOrTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceSwitchProfileOrTemplateName) {
		return nil, false
	}
	return o.SourceSwitchProfileOrTemplateName, true
}

// HasSourceSwitchProfileOrTemplateName returns a boolean if a field has been set.
func (o *FabricClusterAssignment) HasSourceSwitchProfileOrTemplateName() bool {
	if o != nil && !IsNil(o.SourceSwitchProfileOrTemplateName) {
		return true
	}

	return false
}

// SetSourceSwitchProfileOrTemplateName gets a reference to the given string and assigns it to the SourceSwitchProfileOrTemplateName field.
func (o *FabricClusterAssignment) SetSourceSwitchProfileOrTemplateName(v string) {
	o.SourceSwitchProfileOrTemplateName = &v
}

func (o FabricClusterAssignment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricClusterAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.NetworkElement) {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if !IsNil(o.SourceSwitchProfileOrTemplateName) {
		toSerialize["SourceSwitchProfileOrTemplateName"] = o.SourceSwitchProfileOrTemplateName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricClusterAssignment) UnmarshalJSON(data []byte) (err error) {
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
	type FabricClusterAssignmentWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType     string   `json:"ObjectType"`
		NetworkElement *MoMoRef `json:"NetworkElement,omitempty"`
		// Name of the source SwitchProfile or SwitchProfileTemplate whose clone has to be assigned to the network element mentioned in NetworkElement property under ClusterAssignments.
		SourceSwitchProfileOrTemplateName *string `json:"SourceSwitchProfileOrTemplateName,omitempty" validate:"regexp=^$|^[a-zA-Z0-9_.-]{1,64}$"`
	}

	varFabricClusterAssignmentWithoutEmbeddedStruct := FabricClusterAssignmentWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricClusterAssignmentWithoutEmbeddedStruct)
	if err == nil {
		varFabricClusterAssignment := _FabricClusterAssignment{}
		varFabricClusterAssignment.ClassId = varFabricClusterAssignmentWithoutEmbeddedStruct.ClassId
		varFabricClusterAssignment.ObjectType = varFabricClusterAssignmentWithoutEmbeddedStruct.ObjectType
		varFabricClusterAssignment.NetworkElement = varFabricClusterAssignmentWithoutEmbeddedStruct.NetworkElement
		varFabricClusterAssignment.SourceSwitchProfileOrTemplateName = varFabricClusterAssignmentWithoutEmbeddedStruct.SourceSwitchProfileOrTemplateName
		*o = FabricClusterAssignment(varFabricClusterAssignment)
	} else {
		return err
	}

	varFabricClusterAssignment := _FabricClusterAssignment{}

	err = json.Unmarshal(data, &varFabricClusterAssignment)
	if err == nil {
		o.MoBaseComplexType = varFabricClusterAssignment.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "SourceSwitchProfileOrTemplateName")

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

type NullableFabricClusterAssignment struct {
	value *FabricClusterAssignment
	isSet bool
}

func (v NullableFabricClusterAssignment) Get() *FabricClusterAssignment {
	return v.value
}

func (v *NullableFabricClusterAssignment) Set(val *FabricClusterAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricClusterAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricClusterAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricClusterAssignment(val *FabricClusterAssignment) *NullableFabricClusterAssignment {
	return &NullableFabricClusterAssignment{value: val, isSet: true}
}

func (v NullableFabricClusterAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricClusterAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
