/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-07-24T21:18:00Z.
 *
 * API version: 1.0.9-4404
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ConnectorHttpRequest A HTTP request sent by a cloud service to be proxied through a device connector.
type ConnectorHttpRequest struct {
	ConnectorBaseMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Target endpoint Moid which is used to fetch the previously persisted Target information in Intersight to create HTTP request along with any authentication info specifed.
	AssetTargetMoid *string `json:"AssetTargetMoid,omitempty"`
	// Contents of the request body to send for PUT/PATCH/POST requests.
	Body *string `json:"Body,omitempty"`
	// The timeout for establishing the TCP connection to the target host. If not set the request timeout value is used.
	DialTimeout *int64 `json:"DialTimeout,omitempty"`
	// The MO id of the asset.EndpointConnection this request is directed to. If set plugin will insert connection details into the request, including credentials if defined.
	EndpointMoid *string `json:"EndpointMoid,omitempty"`
	// Collection of key value pairs to set in the request header.
	Header interface{} `json:"Header,omitempty"`
	// The request is for an internal platform API that requires authentication to be inserted by the platform implementation.
	Internal *bool `json:"Internal,omitempty"`
	// Method specifies the HTTP method (GET, POST, PUT, etc.). For client requests an empty string means GET.
	Method *string `json:"Method,omitempty"`
	// The timeout for the HTTP request to complete, from connection establishment to response body read complete. If not set a default timeout of five minutes is used.
	Timeout              *int64               `json:"Timeout,omitempty"`
	Url                  NullableConnectorUrl `json:"Url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorHttpRequest ConnectorHttpRequest

// NewConnectorHttpRequest instantiates a new ConnectorHttpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorHttpRequest(classId string, objectType string) *ConnectorHttpRequest {
	this := ConnectorHttpRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorHttpRequestWithDefaults instantiates a new ConnectorHttpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorHttpRequestWithDefaults() *ConnectorHttpRequest {
	this := ConnectorHttpRequest{}
	var classId string = "connector.HttpRequest"
	this.ClassId = classId
	var objectType string = "connector.HttpRequest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorHttpRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorHttpRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorHttpRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorHttpRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssetTargetMoid returns the AssetTargetMoid field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetAssetTargetMoid() string {
	if o == nil || o.AssetTargetMoid == nil {
		var ret string
		return ret
	}
	return *o.AssetTargetMoid
}

// GetAssetTargetMoidOk returns a tuple with the AssetTargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetAssetTargetMoidOk() (*string, bool) {
	if o == nil || o.AssetTargetMoid == nil {
		return nil, false
	}
	return o.AssetTargetMoid, true
}

// HasAssetTargetMoid returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasAssetTargetMoid() bool {
	if o != nil && o.AssetTargetMoid != nil {
		return true
	}

	return false
}

// SetAssetTargetMoid gets a reference to the given string and assigns it to the AssetTargetMoid field.
func (o *ConnectorHttpRequest) SetAssetTargetMoid(v string) {
	o.AssetTargetMoid = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ConnectorHttpRequest) SetBody(v string) {
	o.Body = &v
}

// GetDialTimeout returns the DialTimeout field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetDialTimeout() int64 {
	if o == nil || o.DialTimeout == nil {
		var ret int64
		return ret
	}
	return *o.DialTimeout
}

// GetDialTimeoutOk returns a tuple with the DialTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetDialTimeoutOk() (*int64, bool) {
	if o == nil || o.DialTimeout == nil {
		return nil, false
	}
	return o.DialTimeout, true
}

// HasDialTimeout returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasDialTimeout() bool {
	if o != nil && o.DialTimeout != nil {
		return true
	}

	return false
}

// SetDialTimeout gets a reference to the given int64 and assigns it to the DialTimeout field.
func (o *ConnectorHttpRequest) SetDialTimeout(v int64) {
	o.DialTimeout = &v
}

// GetEndpointMoid returns the EndpointMoid field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetEndpointMoid() string {
	if o == nil || o.EndpointMoid == nil {
		var ret string
		return ret
	}
	return *o.EndpointMoid
}

// GetEndpointMoidOk returns a tuple with the EndpointMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetEndpointMoidOk() (*string, bool) {
	if o == nil || o.EndpointMoid == nil {
		return nil, false
	}
	return o.EndpointMoid, true
}

// HasEndpointMoid returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasEndpointMoid() bool {
	if o != nil && o.EndpointMoid != nil {
		return true
	}

	return false
}

// SetEndpointMoid gets a reference to the given string and assigns it to the EndpointMoid field.
func (o *ConnectorHttpRequest) SetEndpointMoid(v string) {
	o.EndpointMoid = &v
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorHttpRequest) GetHeader() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorHttpRequest) GetHeaderOk() (*interface{}, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return &o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given interface{} and assigns it to the Header field.
func (o *ConnectorHttpRequest) SetHeader(v interface{}) {
	o.Header = v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *ConnectorHttpRequest) SetInternal(v bool) {
	o.Internal = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *ConnectorHttpRequest) SetMethod(v string) {
	o.Method = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ConnectorHttpRequest) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorHttpRequest) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *ConnectorHttpRequest) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorHttpRequest) GetUrl() ConnectorUrl {
	if o == nil || o.Url.Get() == nil {
		var ret ConnectorUrl
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorHttpRequest) GetUrlOk() (*ConnectorUrl, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ConnectorHttpRequest) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableConnectorUrl and assigns it to the Url field.
func (o *ConnectorHttpRequest) SetUrl(v ConnectorUrl) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *ConnectorHttpRequest) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ConnectorHttpRequest) UnsetUrl() {
	o.Url.Unset()
}

func (o ConnectorHttpRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorBaseMessage, errConnectorBaseMessage := json.Marshal(o.ConnectorBaseMessage)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	errConnectorBaseMessage = json.Unmarshal([]byte(serializedConnectorBaseMessage), &toSerialize)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AssetTargetMoid != nil {
		toSerialize["AssetTargetMoid"] = o.AssetTargetMoid
	}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.DialTimeout != nil {
		toSerialize["DialTimeout"] = o.DialTimeout
	}
	if o.EndpointMoid != nil {
		toSerialize["EndpointMoid"] = o.EndpointMoid
	}
	if o.Header != nil {
		toSerialize["Header"] = o.Header
	}
	if o.Internal != nil {
		toSerialize["Internal"] = o.Internal
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}
	if o.Url.IsSet() {
		toSerialize["Url"] = o.Url.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorHttpRequest) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorHttpRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Target endpoint Moid which is used to fetch the previously persisted Target information in Intersight to create HTTP request along with any authentication info specifed.
		AssetTargetMoid *string `json:"AssetTargetMoid,omitempty"`
		// Contents of the request body to send for PUT/PATCH/POST requests.
		Body *string `json:"Body,omitempty"`
		// The timeout for establishing the TCP connection to the target host. If not set the request timeout value is used.
		DialTimeout *int64 `json:"DialTimeout,omitempty"`
		// The MO id of the asset.EndpointConnection this request is directed to. If set plugin will insert connection details into the request, including credentials if defined.
		EndpointMoid *string `json:"EndpointMoid,omitempty"`
		// Collection of key value pairs to set in the request header.
		Header interface{} `json:"Header,omitempty"`
		// The request is for an internal platform API that requires authentication to be inserted by the platform implementation.
		Internal *bool `json:"Internal,omitempty"`
		// Method specifies the HTTP method (GET, POST, PUT, etc.). For client requests an empty string means GET.
		Method *string `json:"Method,omitempty"`
		// The timeout for the HTTP request to complete, from connection establishment to response body read complete. If not set a default timeout of five minutes is used.
		Timeout *int64               `json:"Timeout,omitempty"`
		Url     NullableConnectorUrl `json:"Url,omitempty"`
	}

	varConnectorHttpRequestWithoutEmbeddedStruct := ConnectorHttpRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorHttpRequestWithoutEmbeddedStruct)
	if err == nil {
		varConnectorHttpRequest := _ConnectorHttpRequest{}
		varConnectorHttpRequest.ClassId = varConnectorHttpRequestWithoutEmbeddedStruct.ClassId
		varConnectorHttpRequest.ObjectType = varConnectorHttpRequestWithoutEmbeddedStruct.ObjectType
		varConnectorHttpRequest.AssetTargetMoid = varConnectorHttpRequestWithoutEmbeddedStruct.AssetTargetMoid
		varConnectorHttpRequest.Body = varConnectorHttpRequestWithoutEmbeddedStruct.Body
		varConnectorHttpRequest.DialTimeout = varConnectorHttpRequestWithoutEmbeddedStruct.DialTimeout
		varConnectorHttpRequest.EndpointMoid = varConnectorHttpRequestWithoutEmbeddedStruct.EndpointMoid
		varConnectorHttpRequest.Header = varConnectorHttpRequestWithoutEmbeddedStruct.Header
		varConnectorHttpRequest.Internal = varConnectorHttpRequestWithoutEmbeddedStruct.Internal
		varConnectorHttpRequest.Method = varConnectorHttpRequestWithoutEmbeddedStruct.Method
		varConnectorHttpRequest.Timeout = varConnectorHttpRequestWithoutEmbeddedStruct.Timeout
		varConnectorHttpRequest.Url = varConnectorHttpRequestWithoutEmbeddedStruct.Url
		*o = ConnectorHttpRequest(varConnectorHttpRequest)
	} else {
		return err
	}

	varConnectorHttpRequest := _ConnectorHttpRequest{}

	err = json.Unmarshal(bytes, &varConnectorHttpRequest)
	if err == nil {
		o.ConnectorBaseMessage = varConnectorHttpRequest.ConnectorBaseMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssetTargetMoid")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "DialTimeout")
		delete(additionalProperties, "EndpointMoid")
		delete(additionalProperties, "Header")
		delete(additionalProperties, "Internal")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "Timeout")
		delete(additionalProperties, "Url")

		// remove fields from embedded structs
		reflectConnectorBaseMessage := reflect.ValueOf(o.ConnectorBaseMessage)
		for i := 0; i < reflectConnectorBaseMessage.Type().NumField(); i++ {
			t := reflectConnectorBaseMessage.Type().Field(i)

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

type NullableConnectorHttpRequest struct {
	value *ConnectorHttpRequest
	isSet bool
}

func (v NullableConnectorHttpRequest) Get() *ConnectorHttpRequest {
	return v.value
}

func (v *NullableConnectorHttpRequest) Set(val *ConnectorHttpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorHttpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorHttpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorHttpRequest(val *ConnectorHttpRequest) *NullableConnectorHttpRequest {
	return &NullableConnectorHttpRequest{value: val, isSet: true}
}

func (v NullableConnectorHttpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorHttpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
