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

// ServicerequestMessage The instance status message can be used at CatalogServiceRequest level.
type ServicerequestMessage struct {
	ServiceitemBaseMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The service item in which operation is perfomed.
	ServiceItemName      *string `json:"ServiceItemName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicerequestMessage ServicerequestMessage

// NewServicerequestMessage instantiates a new ServicerequestMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicerequestMessage(classId string, objectType string) *ServicerequestMessage {
	this := ServicerequestMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServicerequestMessageWithDefaults instantiates a new ServicerequestMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicerequestMessageWithDefaults() *ServicerequestMessage {
	this := ServicerequestMessage{}
	var classId string = "servicerequest.Message"
	this.ClassId = classId
	var objectType string = "servicerequest.Message"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServicerequestMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServicerequestMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServicerequestMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServicerequestMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServicerequestMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServicerequestMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetServiceItemName returns the ServiceItemName field value if set, zero value otherwise.
func (o *ServicerequestMessage) GetServiceItemName() string {
	if o == nil || o.ServiceItemName == nil {
		var ret string
		return ret
	}
	return *o.ServiceItemName
}

// GetServiceItemNameOk returns a tuple with the ServiceItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicerequestMessage) GetServiceItemNameOk() (*string, bool) {
	if o == nil || o.ServiceItemName == nil {
		return nil, false
	}
	return o.ServiceItemName, true
}

// HasServiceItemName returns a boolean if a field has been set.
func (o *ServicerequestMessage) HasServiceItemName() bool {
	if o != nil && o.ServiceItemName != nil {
		return true
	}

	return false
}

// SetServiceItemName gets a reference to the given string and assigns it to the ServiceItemName field.
func (o *ServicerequestMessage) SetServiceItemName(v string) {
	o.ServiceItemName = &v
}

func (o ServicerequestMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedServiceitemBaseMessage, errServiceitemBaseMessage := json.Marshal(o.ServiceitemBaseMessage)
	if errServiceitemBaseMessage != nil {
		return []byte{}, errServiceitemBaseMessage
	}
	errServiceitemBaseMessage = json.Unmarshal([]byte(serializedServiceitemBaseMessage), &toSerialize)
	if errServiceitemBaseMessage != nil {
		return []byte{}, errServiceitemBaseMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ServiceItemName != nil {
		toSerialize["ServiceItemName"] = o.ServiceItemName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicerequestMessage) UnmarshalJSON(bytes []byte) (err error) {
	type ServicerequestMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The service item in which operation is perfomed.
		ServiceItemName *string `json:"ServiceItemName,omitempty"`
	}

	varServicerequestMessageWithoutEmbeddedStruct := ServicerequestMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varServicerequestMessageWithoutEmbeddedStruct)
	if err == nil {
		varServicerequestMessage := _ServicerequestMessage{}
		varServicerequestMessage.ClassId = varServicerequestMessageWithoutEmbeddedStruct.ClassId
		varServicerequestMessage.ObjectType = varServicerequestMessageWithoutEmbeddedStruct.ObjectType
		varServicerequestMessage.ServiceItemName = varServicerequestMessageWithoutEmbeddedStruct.ServiceItemName
		*o = ServicerequestMessage(varServicerequestMessage)
	} else {
		return err
	}

	varServicerequestMessage := _ServicerequestMessage{}

	err = json.Unmarshal(bytes, &varServicerequestMessage)
	if err == nil {
		o.ServiceitemBaseMessage = varServicerequestMessage.ServiceitemBaseMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ServiceItemName")

		// remove fields from embedded structs
		reflectServiceitemBaseMessage := reflect.ValueOf(o.ServiceitemBaseMessage)
		for i := 0; i < reflectServiceitemBaseMessage.Type().NumField(); i++ {
			t := reflectServiceitemBaseMessage.Type().Field(i)

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

type NullableServicerequestMessage struct {
	value *ServicerequestMessage
	isSet bool
}

func (v NullableServicerequestMessage) Get() *ServicerequestMessage {
	return v.value
}

func (v *NullableServicerequestMessage) Set(val *ServicerequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableServicerequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableServicerequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicerequestMessage(val *ServicerequestMessage) *NullableServicerequestMessage {
	return &NullableServicerequestMessage{value: val, isSet: true}
}

func (v NullableServicerequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicerequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
