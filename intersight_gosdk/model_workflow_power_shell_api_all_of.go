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

// WorkflowPowerShellApiAllOf Definition of the list of properties defined in 'workflow.PowerShellApi', excluding properties defined in parent classes.
type WorkflowPowerShellApiAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The response of a PowerShell script is an object, since PowerShell is an Object based language. Each object can contain multiple objects as properties, each of which in turn can contain multiple objects and so on and so forth. The depth field specifies how many levels of contained objects are included in the JSON representation.
	Depth *int64 `json:"Depth,omitempty"`
	// The timeout in seconds for the execution of the script against the given endpoint.
	OperationTimeout *string `json:"OperationTimeout,omitempty"`
	// The grammar specification to parse the response and extract the required values.
	PowerShellResponseSpec interface{} `json:"PowerShellResponseSpec,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _WorkflowPowerShellApiAllOf WorkflowPowerShellApiAllOf

// NewWorkflowPowerShellApiAllOf instantiates a new WorkflowPowerShellApiAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPowerShellApiAllOf(classId string, objectType string) *WorkflowPowerShellApiAllOf {
	this := WorkflowPowerShellApiAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowPowerShellApiAllOfWithDefaults instantiates a new WorkflowPowerShellApiAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPowerShellApiAllOfWithDefaults() *WorkflowPowerShellApiAllOf {
	this := WorkflowPowerShellApiAllOf{}
	var classId string = "workflow.PowerShellApi"
	this.ClassId = classId
	var objectType string = "workflow.PowerShellApi"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowPowerShellApiAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellApiAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowPowerShellApiAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowPowerShellApiAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellApiAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowPowerShellApiAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *WorkflowPowerShellApiAllOf) GetDepth() int64 {
	if o == nil || o.Depth == nil {
		var ret int64
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellApiAllOf) GetDepthOk() (*int64, bool) {
	if o == nil || o.Depth == nil {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *WorkflowPowerShellApiAllOf) HasDepth() bool {
	if o != nil && o.Depth != nil {
		return true
	}

	return false
}

// SetDepth gets a reference to the given int64 and assigns it to the Depth field.
func (o *WorkflowPowerShellApiAllOf) SetDepth(v int64) {
	o.Depth = &v
}

// GetOperationTimeout returns the OperationTimeout field value if set, zero value otherwise.
func (o *WorkflowPowerShellApiAllOf) GetOperationTimeout() string {
	if o == nil || o.OperationTimeout == nil {
		var ret string
		return ret
	}
	return *o.OperationTimeout
}

// GetOperationTimeoutOk returns a tuple with the OperationTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellApiAllOf) GetOperationTimeoutOk() (*string, bool) {
	if o == nil || o.OperationTimeout == nil {
		return nil, false
	}
	return o.OperationTimeout, true
}

// HasOperationTimeout returns a boolean if a field has been set.
func (o *WorkflowPowerShellApiAllOf) HasOperationTimeout() bool {
	if o != nil && o.OperationTimeout != nil {
		return true
	}

	return false
}

// SetOperationTimeout gets a reference to the given string and assigns it to the OperationTimeout field.
func (o *WorkflowPowerShellApiAllOf) SetOperationTimeout(v string) {
	o.OperationTimeout = &v
}

// GetPowerShellResponseSpec returns the PowerShellResponseSpec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPowerShellApiAllOf) GetPowerShellResponseSpec() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PowerShellResponseSpec
}

// GetPowerShellResponseSpecOk returns a tuple with the PowerShellResponseSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPowerShellApiAllOf) GetPowerShellResponseSpecOk() (*interface{}, bool) {
	if o == nil || o.PowerShellResponseSpec == nil {
		return nil, false
	}
	return &o.PowerShellResponseSpec, true
}

// HasPowerShellResponseSpec returns a boolean if a field has been set.
func (o *WorkflowPowerShellApiAllOf) HasPowerShellResponseSpec() bool {
	if o != nil && o.PowerShellResponseSpec != nil {
		return true
	}

	return false
}

// SetPowerShellResponseSpec gets a reference to the given interface{} and assigns it to the PowerShellResponseSpec field.
func (o *WorkflowPowerShellApiAllOf) SetPowerShellResponseSpec(v interface{}) {
	o.PowerShellResponseSpec = v
}

func (o WorkflowPowerShellApiAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Depth != nil {
		toSerialize["Depth"] = o.Depth
	}
	if o.OperationTimeout != nil {
		toSerialize["OperationTimeout"] = o.OperationTimeout
	}
	if o.PowerShellResponseSpec != nil {
		toSerialize["PowerShellResponseSpec"] = o.PowerShellResponseSpec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPowerShellApiAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowPowerShellApiAllOf := _WorkflowPowerShellApiAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowPowerShellApiAllOf); err == nil {
		*o = WorkflowPowerShellApiAllOf(varWorkflowPowerShellApiAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Depth")
		delete(additionalProperties, "OperationTimeout")
		delete(additionalProperties, "PowerShellResponseSpec")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowPowerShellApiAllOf struct {
	value *WorkflowPowerShellApiAllOf
	isSet bool
}

func (v NullableWorkflowPowerShellApiAllOf) Get() *WorkflowPowerShellApiAllOf {
	return v.value
}

func (v *NullableWorkflowPowerShellApiAllOf) Set(val *WorkflowPowerShellApiAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPowerShellApiAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPowerShellApiAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPowerShellApiAllOf(val *WorkflowPowerShellApiAllOf) *NullableWorkflowPowerShellApiAllOf {
	return &NullableWorkflowPowerShellApiAllOf{value: val, isSet: true}
}

func (v NullableWorkflowPowerShellApiAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPowerShellApiAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
