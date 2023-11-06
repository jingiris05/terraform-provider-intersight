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

// PolicyScheduledAction ScheduledAction can be used to schedule actions after current workflow in an asynchronous manner.
type PolicyScheduledAction struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the action to be performed on the profile.
	Action *string `json:"Action,omitempty"`
	// ProceedOnReboot can be used to acknowledge server reboot while triggering deploy/activate.
	ProceedOnReboot      *bool `json:"ProceedOnReboot,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyScheduledAction PolicyScheduledAction

// NewPolicyScheduledAction instantiates a new PolicyScheduledAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyScheduledAction(classId string, objectType string) *PolicyScheduledAction {
	this := PolicyScheduledAction{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyScheduledActionWithDefaults instantiates a new PolicyScheduledAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyScheduledActionWithDefaults() *PolicyScheduledAction {
	this := PolicyScheduledAction{}
	var classId string = "policy.ScheduledAction"
	this.ClassId = classId
	var objectType string = "policy.ScheduledAction"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyScheduledAction) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyScheduledAction) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyScheduledAction) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyScheduledAction) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyScheduledAction) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyScheduledAction) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PolicyScheduledAction) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyScheduledAction) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PolicyScheduledAction) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PolicyScheduledAction) SetAction(v string) {
	o.Action = &v
}

// GetProceedOnReboot returns the ProceedOnReboot field value if set, zero value otherwise.
func (o *PolicyScheduledAction) GetProceedOnReboot() bool {
	if o == nil || o.ProceedOnReboot == nil {
		var ret bool
		return ret
	}
	return *o.ProceedOnReboot
}

// GetProceedOnRebootOk returns a tuple with the ProceedOnReboot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyScheduledAction) GetProceedOnRebootOk() (*bool, bool) {
	if o == nil || o.ProceedOnReboot == nil {
		return nil, false
	}
	return o.ProceedOnReboot, true
}

// HasProceedOnReboot returns a boolean if a field has been set.
func (o *PolicyScheduledAction) HasProceedOnReboot() bool {
	if o != nil && o.ProceedOnReboot != nil {
		return true
	}

	return false
}

// SetProceedOnReboot gets a reference to the given bool and assigns it to the ProceedOnReboot field.
func (o *PolicyScheduledAction) SetProceedOnReboot(v bool) {
	o.ProceedOnReboot = &v
}

func (o PolicyScheduledAction) MarshalJSON() ([]byte, error) {
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
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.ProceedOnReboot != nil {
		toSerialize["ProceedOnReboot"] = o.ProceedOnReboot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyScheduledAction) UnmarshalJSON(bytes []byte) (err error) {
	type PolicyScheduledActionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the action to be performed on the profile.
		Action *string `json:"Action,omitempty"`
		// ProceedOnReboot can be used to acknowledge server reboot while triggering deploy/activate.
		ProceedOnReboot *bool `json:"ProceedOnReboot,omitempty"`
	}

	varPolicyScheduledActionWithoutEmbeddedStruct := PolicyScheduledActionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPolicyScheduledActionWithoutEmbeddedStruct)
	if err == nil {
		varPolicyScheduledAction := _PolicyScheduledAction{}
		varPolicyScheduledAction.ClassId = varPolicyScheduledActionWithoutEmbeddedStruct.ClassId
		varPolicyScheduledAction.ObjectType = varPolicyScheduledActionWithoutEmbeddedStruct.ObjectType
		varPolicyScheduledAction.Action = varPolicyScheduledActionWithoutEmbeddedStruct.Action
		varPolicyScheduledAction.ProceedOnReboot = varPolicyScheduledActionWithoutEmbeddedStruct.ProceedOnReboot
		*o = PolicyScheduledAction(varPolicyScheduledAction)
	} else {
		return err
	}

	varPolicyScheduledAction := _PolicyScheduledAction{}

	err = json.Unmarshal(bytes, &varPolicyScheduledAction)
	if err == nil {
		o.MoBaseComplexType = varPolicyScheduledAction.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "ProceedOnReboot")

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

type NullablePolicyScheduledAction struct {
	value *PolicyScheduledAction
	isSet bool
}

func (v NullablePolicyScheduledAction) Get() *PolicyScheduledAction {
	return v.value
}

func (v *NullablePolicyScheduledAction) Set(val *PolicyScheduledAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyScheduledAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyScheduledAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyScheduledAction(val *PolicyScheduledAction) *NullablePolicyScheduledAction {
	return &NullablePolicyScheduledAction{value: val, isSet: true}
}

func (v NullablePolicyScheduledAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyScheduledAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
