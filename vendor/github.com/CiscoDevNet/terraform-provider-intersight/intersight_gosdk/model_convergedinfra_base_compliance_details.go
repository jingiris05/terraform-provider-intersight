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

// ConvergedinfraBaseComplianceDetails Abstract type for all compliance details that may be part of a converged infrastructure pod.
type ConvergedinfraBaseComplianceDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The name of the component for which the compliance is evaluated.
	Name *string `json:"Name,omitempty"`
	// Reason for the status, when the status is Incomplete or NotApproved. Reason should help to identify the component that is not compliant. * `NotEvaluated` - The validation for the HCL or Interop status is yet to be performed. * `Missing-Os-Info` - This means the Interop status for the sever cannot be computed because we have missing OS information. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Incompatible-Storage-Os` - The validation failed because the given storage name and version combination is not found in Interop matrix. * `Incompatible-Os-Info` - The validation failed because the given OS name and version combination is not found in HCL. * `Incompatible-Processor` - The validation failed because the given processor is not found for the given server PID in HCL. * `Incompatible-Server-Platform` - The validation failed because the given server platform is not found in the Interop matrix. * `Incompatible-Server-Model` - The validation failed because the given server model is not found in the Interop matrix. * `Incompatible-Adapter-Model` - The validation failed because the given adapter model is not found in the Interop matrix. * `Incompatible-Switch-Model` - The validation failed because the given switch model is not found in the Interop matrix. * `Incompatible-Server-Firmware` - The validation failed because the given server firmware version is not found in HCL. * `Incompatible-Switch-Firmware` - The validation failed because the given switch firmware version is not found in Interop matrix. * `Incompatible-Firmware` - The validation failed because the given adapter firmware version is not found in either HCL or Interop matrix. * `Incompatible-Driver` - The validation failed because the given driver version is not found in either HCL or Interop matrix. * `Incompatible-Firmware-Driver` - The validation failed because the given adapter firmware and driver version is not found in either HCL or Interop matrix. * `Missing-Os-Driver-Info` - The validation failed because the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Missing-Os-Or-Driver-Info` - This means the Interop status for the switch or storage array cannot be computed because we have either missing Host OS information or missing required driver information. Either install ucstools vib or use power shell scripts to tag proper OS information or install required drivers. * `Incompatible-Components` - The validation failed for the given server because one or more of its components failed validation. * `Compatible` - This means the HCL status and Interop status for the component have passed all validations for all of its related components.
	Reason *string `json:"Reason,omitempty"`
	// Compliance status for the component. * `NotEvaluated` - The interoperability compliance for the component has not be checked. * `Approved` - The component is valid as per the interoperability compliance check. * `NotApproved` - The component is not valid as per the interoperability compliance check. * `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraBaseComplianceDetails ConvergedinfraBaseComplianceDetails

// NewConvergedinfraBaseComplianceDetails instantiates a new ConvergedinfraBaseComplianceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraBaseComplianceDetails(classId string, objectType string) *ConvergedinfraBaseComplianceDetails {
	this := ConvergedinfraBaseComplianceDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraBaseComplianceDetailsWithDefaults instantiates a new ConvergedinfraBaseComplianceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraBaseComplianceDetailsWithDefaults() *ConvergedinfraBaseComplianceDetails {
	this := ConvergedinfraBaseComplianceDetails{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraBaseComplianceDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBaseComplianceDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraBaseComplianceDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraBaseComplianceDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBaseComplianceDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraBaseComplianceDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConvergedinfraBaseComplianceDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBaseComplianceDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConvergedinfraBaseComplianceDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConvergedinfraBaseComplianceDetails) SetName(v string) {
	o.Name = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ConvergedinfraBaseComplianceDetails) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBaseComplianceDetails) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ConvergedinfraBaseComplianceDetails) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ConvergedinfraBaseComplianceDetails) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConvergedinfraBaseComplianceDetails) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBaseComplianceDetails) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConvergedinfraBaseComplianceDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConvergedinfraBaseComplianceDetails) SetStatus(v string) {
	o.Status = &v
}

func (o ConvergedinfraBaseComplianceDetails) MarshalJSON() ([]byte, error) {
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
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraBaseComplianceDetails) UnmarshalJSON(bytes []byte) (err error) {
	type ConvergedinfraBaseComplianceDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The name of the component for which the compliance is evaluated.
		Name *string `json:"Name,omitempty"`
		// Reason for the status, when the status is Incomplete or NotApproved. Reason should help to identify the component that is not compliant. * `NotEvaluated` - The validation for the HCL or Interop status is yet to be performed. * `Missing-Os-Info` - This means the Interop status for the sever cannot be computed because we have missing OS information. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Incompatible-Storage-Os` - The validation failed because the given storage name and version combination is not found in Interop matrix. * `Incompatible-Os-Info` - The validation failed because the given OS name and version combination is not found in HCL. * `Incompatible-Processor` - The validation failed because the given processor is not found for the given server PID in HCL. * `Incompatible-Server-Platform` - The validation failed because the given server platform is not found in the Interop matrix. * `Incompatible-Server-Model` - The validation failed because the given server model is not found in the Interop matrix. * `Incompatible-Adapter-Model` - The validation failed because the given adapter model is not found in the Interop matrix. * `Incompatible-Switch-Model` - The validation failed because the given switch model is not found in the Interop matrix. * `Incompatible-Server-Firmware` - The validation failed because the given server firmware version is not found in HCL. * `Incompatible-Switch-Firmware` - The validation failed because the given switch firmware version is not found in Interop matrix. * `Incompatible-Firmware` - The validation failed because the given adapter firmware version is not found in either HCL or Interop matrix. * `Incompatible-Driver` - The validation failed because the given driver version is not found in either HCL or Interop matrix. * `Incompatible-Firmware-Driver` - The validation failed because the given adapter firmware and driver version is not found in either HCL or Interop matrix. * `Missing-Os-Driver-Info` - The validation failed because the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Missing-Os-Or-Driver-Info` - This means the Interop status for the switch or storage array cannot be computed because we have either missing Host OS information or missing required driver information. Either install ucstools vib or use power shell scripts to tag proper OS information or install required drivers. * `Incompatible-Components` - The validation failed for the given server because one or more of its components failed validation. * `Compatible` - This means the HCL status and Interop status for the component have passed all validations for all of its related components.
		Reason *string `json:"Reason,omitempty"`
		// Compliance status for the component. * `NotEvaluated` - The interoperability compliance for the component has not be checked. * `Approved` - The component is valid as per the interoperability compliance check. * `NotApproved` - The component is not valid as per the interoperability compliance check. * `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data.
		Status *string `json:"Status,omitempty"`
	}

	varConvergedinfraBaseComplianceDetailsWithoutEmbeddedStruct := ConvergedinfraBaseComplianceDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConvergedinfraBaseComplianceDetailsWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraBaseComplianceDetails := _ConvergedinfraBaseComplianceDetails{}
		varConvergedinfraBaseComplianceDetails.ClassId = varConvergedinfraBaseComplianceDetailsWithoutEmbeddedStruct.ClassId
		varConvergedinfraBaseComplianceDetails.ObjectType = varConvergedinfraBaseComplianceDetailsWithoutEmbeddedStruct.ObjectType
		varConvergedinfraBaseComplianceDetails.Name = varConvergedinfraBaseComplianceDetailsWithoutEmbeddedStruct.Name
		varConvergedinfraBaseComplianceDetails.Reason = varConvergedinfraBaseComplianceDetailsWithoutEmbeddedStruct.Reason
		varConvergedinfraBaseComplianceDetails.Status = varConvergedinfraBaseComplianceDetailsWithoutEmbeddedStruct.Status
		*o = ConvergedinfraBaseComplianceDetails(varConvergedinfraBaseComplianceDetails)
	} else {
		return err
	}

	varConvergedinfraBaseComplianceDetails := _ConvergedinfraBaseComplianceDetails{}

	err = json.Unmarshal(bytes, &varConvergedinfraBaseComplianceDetails)
	if err == nil {
		o.MoBaseMo = varConvergedinfraBaseComplianceDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Status")

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

type NullableConvergedinfraBaseComplianceDetails struct {
	value *ConvergedinfraBaseComplianceDetails
	isSet bool
}

func (v NullableConvergedinfraBaseComplianceDetails) Get() *ConvergedinfraBaseComplianceDetails {
	return v.value
}

func (v *NullableConvergedinfraBaseComplianceDetails) Set(val *ConvergedinfraBaseComplianceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraBaseComplianceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraBaseComplianceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraBaseComplianceDetails(val *ConvergedinfraBaseComplianceDetails) *NullableConvergedinfraBaseComplianceDetails {
	return &NullableConvergedinfraBaseComplianceDetails{value: val, isSet: true}
}

func (v NullableConvergedinfraBaseComplianceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraBaseComplianceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
