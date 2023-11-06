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
	"time"
)

// RecommendationExpansionRequest Entity representing hardware expansion request like name, requestTime, etc.
type RecommendationExpansionRequest struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The name of the expansion specification.
	Name *string `json:"Name,omitempty"`
	// The time when the expansion was requested.
	RequestTime          *time.Time `json:"RequestTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationExpansionRequest RecommendationExpansionRequest

// NewRecommendationExpansionRequest instantiates a new RecommendationExpansionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationExpansionRequest(classId string, objectType string) *RecommendationExpansionRequest {
	this := RecommendationExpansionRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecommendationExpansionRequestWithDefaults instantiates a new RecommendationExpansionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationExpansionRequestWithDefaults() *RecommendationExpansionRequest {
	this := RecommendationExpansionRequest{}
	var classId string = "recommendation.HardwareExpansionRequest"
	this.ClassId = classId
	var objectType string = "recommendation.HardwareExpansionRequest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationExpansionRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationExpansionRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationExpansionRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationExpansionRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationExpansionRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationExpansionRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RecommendationExpansionRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationExpansionRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RecommendationExpansionRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RecommendationExpansionRequest) SetName(v string) {
	o.Name = &v
}

// GetRequestTime returns the RequestTime field value if set, zero value otherwise.
func (o *RecommendationExpansionRequest) GetRequestTime() time.Time {
	if o == nil || o.RequestTime == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestTime
}

// GetRequestTimeOk returns a tuple with the RequestTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationExpansionRequest) GetRequestTimeOk() (*time.Time, bool) {
	if o == nil || o.RequestTime == nil {
		return nil, false
	}
	return o.RequestTime, true
}

// HasRequestTime returns a boolean if a field has been set.
func (o *RecommendationExpansionRequest) HasRequestTime() bool {
	if o != nil && o.RequestTime != nil {
		return true
	}

	return false
}

// SetRequestTime gets a reference to the given time.Time and assigns it to the RequestTime field.
func (o *RecommendationExpansionRequest) SetRequestTime(v time.Time) {
	o.RequestTime = &v
}

func (o RecommendationExpansionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RequestTime != nil {
		toSerialize["RequestTime"] = o.RequestTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationExpansionRequest) UnmarshalJSON(bytes []byte) (err error) {
	type RecommendationExpansionRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The name of the expansion specification.
		Name *string `json:"Name,omitempty"`
		// The time when the expansion was requested.
		RequestTime *time.Time `json:"RequestTime,omitempty"`
	}

	varRecommendationExpansionRequestWithoutEmbeddedStruct := RecommendationExpansionRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecommendationExpansionRequestWithoutEmbeddedStruct)
	if err == nil {
		varRecommendationExpansionRequest := _RecommendationExpansionRequest{}
		varRecommendationExpansionRequest.ClassId = varRecommendationExpansionRequestWithoutEmbeddedStruct.ClassId
		varRecommendationExpansionRequest.ObjectType = varRecommendationExpansionRequestWithoutEmbeddedStruct.ObjectType
		varRecommendationExpansionRequest.Name = varRecommendationExpansionRequestWithoutEmbeddedStruct.Name
		varRecommendationExpansionRequest.RequestTime = varRecommendationExpansionRequestWithoutEmbeddedStruct.RequestTime
		*o = RecommendationExpansionRequest(varRecommendationExpansionRequest)
	} else {
		return err
	}

	varRecommendationExpansionRequest := _RecommendationExpansionRequest{}

	err = json.Unmarshal(bytes, &varRecommendationExpansionRequest)
	if err == nil {
		o.MoBaseMo = varRecommendationExpansionRequest.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RequestTime")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableRecommendationExpansionRequest struct {
	value *RecommendationExpansionRequest
	isSet bool
}

func (v NullableRecommendationExpansionRequest) Get() *RecommendationExpansionRequest {
	return v.value
}

func (v *NullableRecommendationExpansionRequest) Set(val *RecommendationExpansionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationExpansionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationExpansionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationExpansionRequest(val *RecommendationExpansionRequest) *NullableRecommendationExpansionRequest {
	return &NullableRecommendationExpansionRequest{value: val, isSet: true}
}

func (v NullableRecommendationExpansionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationExpansionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
