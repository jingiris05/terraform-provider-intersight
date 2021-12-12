/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4950
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// NiatelemetryBootflashDetailsAllOf Definition of the list of properties defined in 'niatelemetry.BootflashDetails', excluding properties defined in parent classes.
type NiatelemetryBootflashDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return firmware revision in boot flash details.
	FwRev *string `json:"FwRev,omitempty"`
	// Return model type in boot flash details.
	ModelType *string `json:"ModelType,omitempty"`
	// Return serial id in boot flash details.
	Serial               *string `json:"Serial,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryBootflashDetailsAllOf NiatelemetryBootflashDetailsAllOf

// NewNiatelemetryBootflashDetailsAllOf instantiates a new NiatelemetryBootflashDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryBootflashDetailsAllOf(classId string, objectType string) *NiatelemetryBootflashDetailsAllOf {
	this := NiatelemetryBootflashDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryBootflashDetailsAllOfWithDefaults instantiates a new NiatelemetryBootflashDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryBootflashDetailsAllOfWithDefaults() *NiatelemetryBootflashDetailsAllOf {
	this := NiatelemetryBootflashDetailsAllOf{}
	var classId string = "niatelemetry.BootflashDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.BootflashDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryBootflashDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryBootflashDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryBootflashDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryBootflashDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryBootflashDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryBootflashDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFwRev returns the FwRev field value if set, zero value otherwise.
func (o *NiatelemetryBootflashDetailsAllOf) GetFwRev() string {
	if o == nil || o.FwRev == nil {
		var ret string
		return ret
	}
	return *o.FwRev
}

// GetFwRevOk returns a tuple with the FwRev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryBootflashDetailsAllOf) GetFwRevOk() (*string, bool) {
	if o == nil || o.FwRev == nil {
		return nil, false
	}
	return o.FwRev, true
}

// HasFwRev returns a boolean if a field has been set.
func (o *NiatelemetryBootflashDetailsAllOf) HasFwRev() bool {
	if o != nil && o.FwRev != nil {
		return true
	}

	return false
}

// SetFwRev gets a reference to the given string and assigns it to the FwRev field.
func (o *NiatelemetryBootflashDetailsAllOf) SetFwRev(v string) {
	o.FwRev = &v
}

// GetModelType returns the ModelType field value if set, zero value otherwise.
func (o *NiatelemetryBootflashDetailsAllOf) GetModelType() string {
	if o == nil || o.ModelType == nil {
		var ret string
		return ret
	}
	return *o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryBootflashDetailsAllOf) GetModelTypeOk() (*string, bool) {
	if o == nil || o.ModelType == nil {
		return nil, false
	}
	return o.ModelType, true
}

// HasModelType returns a boolean if a field has been set.
func (o *NiatelemetryBootflashDetailsAllOf) HasModelType() bool {
	if o != nil && o.ModelType != nil {
		return true
	}

	return false
}

// SetModelType gets a reference to the given string and assigns it to the ModelType field.
func (o *NiatelemetryBootflashDetailsAllOf) SetModelType(v string) {
	o.ModelType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NiatelemetryBootflashDetailsAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryBootflashDetailsAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NiatelemetryBootflashDetailsAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NiatelemetryBootflashDetailsAllOf) SetSerial(v string) {
	o.Serial = &v
}

func (o NiatelemetryBootflashDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FwRev != nil {
		toSerialize["FwRev"] = o.FwRev
	}
	if o.ModelType != nil {
		toSerialize["ModelType"] = o.ModelType
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryBootflashDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryBootflashDetailsAllOf := _NiatelemetryBootflashDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryBootflashDetailsAllOf); err == nil {
		*o = NiatelemetryBootflashDetailsAllOf(varNiatelemetryBootflashDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FwRev")
		delete(additionalProperties, "ModelType")
		delete(additionalProperties, "Serial")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryBootflashDetailsAllOf struct {
	value *NiatelemetryBootflashDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryBootflashDetailsAllOf) Get() *NiatelemetryBootflashDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryBootflashDetailsAllOf) Set(val *NiatelemetryBootflashDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryBootflashDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryBootflashDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryBootflashDetailsAllOf(val *NiatelemetryBootflashDetailsAllOf) *NullableNiatelemetryBootflashDetailsAllOf {
	return &NullableNiatelemetryBootflashDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryBootflashDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryBootflashDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
