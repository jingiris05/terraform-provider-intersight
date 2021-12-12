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
	"reflect"
	"strings"
)

// ConnectorWinrmRequest A WinRM message sent from a cloud service to be proxied through a device connector. Encapsulates the details needed to execute a powershell script against an Intersight claimed, WinRM enabled windows target.
type ConnectorWinrmRequest struct {
	ConnectorBaseMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Target endpoint Moid which is used to fetch the previously persisted Target information in Intersight.
	AssetTargetMoid *string `json:"AssetTargetMoid,omitempty"`
	// Contains the file content to be copied or the script to be executed.
	Body *string `json:"Body,omitempty"`
	// The absolute filename to which body is to be written to. This field is ignored if the opType is ExecuteScript.
	Filename *string `json:"Filename,omitempty"`
	// The type of operation to be performed on the endpoint. File copy and script execution are the possible options. Only script execution is supported for now. * `ExecuteScript` - Executes the given script on the target windows machine. * `CopyToFile` - The plugin copies the body of the incoming message to the given location.
	OpType *string `json:"OpType,omitempty"`
	// A unique id to track long running script executions. Every execution request is identified by a unique session id. Scripts that have longer execution times can be tracked to completion by using their unique session id.
	SessionId *string `json:"SessionId,omitempty"`
	// The time within which the script execution must be completed. If the script execution exceeds the given time limit an appropriate response is sent back to the calling service.
	Timeout              *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorWinrmRequest ConnectorWinrmRequest

// NewConnectorWinrmRequest instantiates a new ConnectorWinrmRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorWinrmRequest(classId string, objectType string) *ConnectorWinrmRequest {
	this := ConnectorWinrmRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	var opType string = "ExecuteScript"
	this.OpType = &opType
	return &this
}

// NewConnectorWinrmRequestWithDefaults instantiates a new ConnectorWinrmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorWinrmRequestWithDefaults() *ConnectorWinrmRequest {
	this := ConnectorWinrmRequest{}
	var classId string = "connector.WinrmRequest"
	this.ClassId = classId
	var objectType string = "connector.WinrmRequest"
	this.ObjectType = objectType
	var opType string = "ExecuteScript"
	this.OpType = &opType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorWinrmRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorWinrmRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorWinrmRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorWinrmRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorWinrmRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorWinrmRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssetTargetMoid returns the AssetTargetMoid field value if set, zero value otherwise.
func (o *ConnectorWinrmRequest) GetAssetTargetMoid() string {
	if o == nil || o.AssetTargetMoid == nil {
		var ret string
		return ret
	}
	return *o.AssetTargetMoid
}

// GetAssetTargetMoidOk returns a tuple with the AssetTargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorWinrmRequest) GetAssetTargetMoidOk() (*string, bool) {
	if o == nil || o.AssetTargetMoid == nil {
		return nil, false
	}
	return o.AssetTargetMoid, true
}

// HasAssetTargetMoid returns a boolean if a field has been set.
func (o *ConnectorWinrmRequest) HasAssetTargetMoid() bool {
	if o != nil && o.AssetTargetMoid != nil {
		return true
	}

	return false
}

// SetAssetTargetMoid gets a reference to the given string and assigns it to the AssetTargetMoid field.
func (o *ConnectorWinrmRequest) SetAssetTargetMoid(v string) {
	o.AssetTargetMoid = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ConnectorWinrmRequest) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorWinrmRequest) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ConnectorWinrmRequest) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ConnectorWinrmRequest) SetBody(v string) {
	o.Body = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ConnectorWinrmRequest) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorWinrmRequest) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ConnectorWinrmRequest) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ConnectorWinrmRequest) SetFilename(v string) {
	o.Filename = &v
}

// GetOpType returns the OpType field value if set, zero value otherwise.
func (o *ConnectorWinrmRequest) GetOpType() string {
	if o == nil || o.OpType == nil {
		var ret string
		return ret
	}
	return *o.OpType
}

// GetOpTypeOk returns a tuple with the OpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorWinrmRequest) GetOpTypeOk() (*string, bool) {
	if o == nil || o.OpType == nil {
		return nil, false
	}
	return o.OpType, true
}

// HasOpType returns a boolean if a field has been set.
func (o *ConnectorWinrmRequest) HasOpType() bool {
	if o != nil && o.OpType != nil {
		return true
	}

	return false
}

// SetOpType gets a reference to the given string and assigns it to the OpType field.
func (o *ConnectorWinrmRequest) SetOpType(v string) {
	o.OpType = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *ConnectorWinrmRequest) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorWinrmRequest) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *ConnectorWinrmRequest) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *ConnectorWinrmRequest) SetSessionId(v string) {
	o.SessionId = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ConnectorWinrmRequest) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorWinrmRequest) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ConnectorWinrmRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *ConnectorWinrmRequest) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o ConnectorWinrmRequest) MarshalJSON() ([]byte, error) {
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
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.OpType != nil {
		toSerialize["OpType"] = o.OpType
	}
	if o.SessionId != nil {
		toSerialize["SessionId"] = o.SessionId
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorWinrmRequest) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorWinrmRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Target endpoint Moid which is used to fetch the previously persisted Target information in Intersight.
		AssetTargetMoid *string `json:"AssetTargetMoid,omitempty"`
		// Contains the file content to be copied or the script to be executed.
		Body *string `json:"Body,omitempty"`
		// The absolute filename to which body is to be written to. This field is ignored if the opType is ExecuteScript.
		Filename *string `json:"Filename,omitempty"`
		// The type of operation to be performed on the endpoint. File copy and script execution are the possible options. Only script execution is supported for now. * `ExecuteScript` - Executes the given script on the target windows machine. * `CopyToFile` - The plugin copies the body of the incoming message to the given location.
		OpType *string `json:"OpType,omitempty"`
		// A unique id to track long running script executions. Every execution request is identified by a unique session id. Scripts that have longer execution times can be tracked to completion by using their unique session id.
		SessionId *string `json:"SessionId,omitempty"`
		// The time within which the script execution must be completed. If the script execution exceeds the given time limit an appropriate response is sent back to the calling service.
		Timeout *int64 `json:"Timeout,omitempty"`
	}

	varConnectorWinrmRequestWithoutEmbeddedStruct := ConnectorWinrmRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorWinrmRequestWithoutEmbeddedStruct)
	if err == nil {
		varConnectorWinrmRequest := _ConnectorWinrmRequest{}
		varConnectorWinrmRequest.ClassId = varConnectorWinrmRequestWithoutEmbeddedStruct.ClassId
		varConnectorWinrmRequest.ObjectType = varConnectorWinrmRequestWithoutEmbeddedStruct.ObjectType
		varConnectorWinrmRequest.AssetTargetMoid = varConnectorWinrmRequestWithoutEmbeddedStruct.AssetTargetMoid
		varConnectorWinrmRequest.Body = varConnectorWinrmRequestWithoutEmbeddedStruct.Body
		varConnectorWinrmRequest.Filename = varConnectorWinrmRequestWithoutEmbeddedStruct.Filename
		varConnectorWinrmRequest.OpType = varConnectorWinrmRequestWithoutEmbeddedStruct.OpType
		varConnectorWinrmRequest.SessionId = varConnectorWinrmRequestWithoutEmbeddedStruct.SessionId
		varConnectorWinrmRequest.Timeout = varConnectorWinrmRequestWithoutEmbeddedStruct.Timeout
		*o = ConnectorWinrmRequest(varConnectorWinrmRequest)
	} else {
		return err
	}

	varConnectorWinrmRequest := _ConnectorWinrmRequest{}

	err = json.Unmarshal(bytes, &varConnectorWinrmRequest)
	if err == nil {
		o.ConnectorBaseMessage = varConnectorWinrmRequest.ConnectorBaseMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssetTargetMoid")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "OpType")
		delete(additionalProperties, "SessionId")
		delete(additionalProperties, "Timeout")

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

type NullableConnectorWinrmRequest struct {
	value *ConnectorWinrmRequest
	isSet bool
}

func (v NullableConnectorWinrmRequest) Get() *ConnectorWinrmRequest {
	return v.value
}

func (v *NullableConnectorWinrmRequest) Set(val *ConnectorWinrmRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorWinrmRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorWinrmRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorWinrmRequest(val *ConnectorWinrmRequest) *NullableConnectorWinrmRequest {
	return &NullableConnectorWinrmRequest{value: val, isSet: true}
}

func (v NullableConnectorWinrmRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorWinrmRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
