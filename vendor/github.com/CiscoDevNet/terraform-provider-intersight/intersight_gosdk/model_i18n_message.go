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

// I18nMessage An i18n message that can be translated in multiple languages to support internationalization. An i18n message includes a unique message identifier, a text format string, a list of message parameters that can be used  to substitute parameters, and a translated string based on a user's locale.
type I18nMessage struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The default (en-US) localized message. Default localized message will be stored and directly retrieved when the user's locale setting is en-US.
	Message *string `json:"Message,omitempty"`
	// The unique message identitifer used to lookup text templates in a multi-language message catalog.
	MessageId            *string            `json:"MessageId,omitempty"`
	MessageParams        []I18nMessageParam `json:"MessageParams,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _I18nMessage I18nMessage

// NewI18nMessage instantiates a new I18nMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewI18nMessage(classId string, objectType string) *I18nMessage {
	this := I18nMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewI18nMessageWithDefaults instantiates a new I18nMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewI18nMessageWithDefaults() *I18nMessage {
	this := I18nMessage{}
	var classId string = "i18n.Message"
	this.ClassId = classId
	var objectType string = "i18n.Message"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *I18nMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *I18nMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *I18nMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *I18nMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *I18nMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *I18nMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *I18nMessage) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *I18nMessage) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *I18nMessage) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *I18nMessage) SetMessage(v string) {
	o.Message = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *I18nMessage) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *I18nMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *I18nMessage) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *I18nMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageParams returns the MessageParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *I18nMessage) GetMessageParams() []I18nMessageParam {
	if o == nil {
		var ret []I18nMessageParam
		return ret
	}
	return o.MessageParams
}

// GetMessageParamsOk returns a tuple with the MessageParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *I18nMessage) GetMessageParamsOk() ([]I18nMessageParam, bool) {
	if o == nil || o.MessageParams == nil {
		return nil, false
	}
	return o.MessageParams, true
}

// HasMessageParams returns a boolean if a field has been set.
func (o *I18nMessage) HasMessageParams() bool {
	if o != nil && o.MessageParams != nil {
		return true
	}

	return false
}

// SetMessageParams gets a reference to the given []I18nMessageParam and assigns it to the MessageParams field.
func (o *I18nMessage) SetMessageParams(v []I18nMessageParam) {
	o.MessageParams = v
}

func (o I18nMessage) MarshalJSON() ([]byte, error) {
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
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.MessageId != nil {
		toSerialize["MessageId"] = o.MessageId
	}
	if o.MessageParams != nil {
		toSerialize["MessageParams"] = o.MessageParams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *I18nMessage) UnmarshalJSON(bytes []byte) (err error) {
	type I18nMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The default (en-US) localized message. Default localized message will be stored and directly retrieved when the user's locale setting is en-US.
		Message *string `json:"Message,omitempty"`
		// The unique message identitifer used to lookup text templates in a multi-language message catalog.
		MessageId     *string            `json:"MessageId,omitempty"`
		MessageParams []I18nMessageParam `json:"MessageParams,omitempty"`
	}

	varI18nMessageWithoutEmbeddedStruct := I18nMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varI18nMessageWithoutEmbeddedStruct)
	if err == nil {
		varI18nMessage := _I18nMessage{}
		varI18nMessage.ClassId = varI18nMessageWithoutEmbeddedStruct.ClassId
		varI18nMessage.ObjectType = varI18nMessageWithoutEmbeddedStruct.ObjectType
		varI18nMessage.Message = varI18nMessageWithoutEmbeddedStruct.Message
		varI18nMessage.MessageId = varI18nMessageWithoutEmbeddedStruct.MessageId
		varI18nMessage.MessageParams = varI18nMessageWithoutEmbeddedStruct.MessageParams
		*o = I18nMessage(varI18nMessage)
	} else {
		return err
	}

	varI18nMessage := _I18nMessage{}

	err = json.Unmarshal(bytes, &varI18nMessage)
	if err == nil {
		o.MoBaseComplexType = varI18nMessage.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "MessageId")
		delete(additionalProperties, "MessageParams")

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

type NullableI18nMessage struct {
	value *I18nMessage
	isSet bool
}

func (v NullableI18nMessage) Get() *I18nMessage {
	return v.value
}

func (v *NullableI18nMessage) Set(val *I18nMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableI18nMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableI18nMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableI18nMessage(val *I18nMessage) *NullableI18nMessage {
	return &NullableI18nMessage{value: val, isSet: true}
}

func (v NullableI18nMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableI18nMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
