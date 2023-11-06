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

// CertificatemanagementImcAllOf Definition of the list of properties defined in 'certificatemanagement.Imc', excluding properties defined in parent classes.
type CertificatemanagementImcAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Certificate Type for the certificate management. * `None` - Set certificate on the selected end point . * `KMIPClient` - Set KMIP certificate on the selected end point.
	CertType *string `json:"CertType,omitempty"`
	// Indicates whether the value of the 'privatekey' property has been set.
	IsPrivatekeySet *bool `json:"IsPrivatekeySet,omitempty"`
	// Private Key which is used to validate the certificate.
	Privatekey           *string `json:"Privatekey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificatemanagementImcAllOf CertificatemanagementImcAllOf

// NewCertificatemanagementImcAllOf instantiates a new CertificatemanagementImcAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificatemanagementImcAllOf(classId string, objectType string) *CertificatemanagementImcAllOf {
	this := CertificatemanagementImcAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var certType string = "None"
	this.CertType = &certType
	return &this
}

// NewCertificatemanagementImcAllOfWithDefaults instantiates a new CertificatemanagementImcAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatemanagementImcAllOfWithDefaults() *CertificatemanagementImcAllOf {
	this := CertificatemanagementImcAllOf{}
	var classId string = "certificatemanagement.Imc"
	this.ClassId = classId
	var objectType string = "certificatemanagement.Imc"
	this.ObjectType = objectType
	var certType string = "None"
	this.CertType = &certType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CertificatemanagementImcAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CertificatemanagementImcAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CertificatemanagementImcAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CertificatemanagementImcAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CertificatemanagementImcAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CertificatemanagementImcAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *CertificatemanagementImcAllOf) GetCertType() string {
	if o == nil || o.CertType == nil {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementImcAllOf) GetCertTypeOk() (*string, bool) {
	if o == nil || o.CertType == nil {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *CertificatemanagementImcAllOf) HasCertType() bool {
	if o != nil && o.CertType != nil {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *CertificatemanagementImcAllOf) SetCertType(v string) {
	o.CertType = &v
}

// GetIsPrivatekeySet returns the IsPrivatekeySet field value if set, zero value otherwise.
func (o *CertificatemanagementImcAllOf) GetIsPrivatekeySet() bool {
	if o == nil || o.IsPrivatekeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivatekeySet
}

// GetIsPrivatekeySetOk returns a tuple with the IsPrivatekeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementImcAllOf) GetIsPrivatekeySetOk() (*bool, bool) {
	if o == nil || o.IsPrivatekeySet == nil {
		return nil, false
	}
	return o.IsPrivatekeySet, true
}

// HasIsPrivatekeySet returns a boolean if a field has been set.
func (o *CertificatemanagementImcAllOf) HasIsPrivatekeySet() bool {
	if o != nil && o.IsPrivatekeySet != nil {
		return true
	}

	return false
}

// SetIsPrivatekeySet gets a reference to the given bool and assigns it to the IsPrivatekeySet field.
func (o *CertificatemanagementImcAllOf) SetIsPrivatekeySet(v bool) {
	o.IsPrivatekeySet = &v
}

// GetPrivatekey returns the Privatekey field value if set, zero value otherwise.
func (o *CertificatemanagementImcAllOf) GetPrivatekey() string {
	if o == nil || o.Privatekey == nil {
		var ret string
		return ret
	}
	return *o.Privatekey
}

// GetPrivatekeyOk returns a tuple with the Privatekey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementImcAllOf) GetPrivatekeyOk() (*string, bool) {
	if o == nil || o.Privatekey == nil {
		return nil, false
	}
	return o.Privatekey, true
}

// HasPrivatekey returns a boolean if a field has been set.
func (o *CertificatemanagementImcAllOf) HasPrivatekey() bool {
	if o != nil && o.Privatekey != nil {
		return true
	}

	return false
}

// SetPrivatekey gets a reference to the given string and assigns it to the Privatekey field.
func (o *CertificatemanagementImcAllOf) SetPrivatekey(v string) {
	o.Privatekey = &v
}

func (o CertificatemanagementImcAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CertType != nil {
		toSerialize["CertType"] = o.CertType
	}
	if o.IsPrivatekeySet != nil {
		toSerialize["IsPrivatekeySet"] = o.IsPrivatekeySet
	}
	if o.Privatekey != nil {
		toSerialize["Privatekey"] = o.Privatekey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CertificatemanagementImcAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCertificatemanagementImcAllOf := _CertificatemanagementImcAllOf{}

	if err = json.Unmarshal(bytes, &varCertificatemanagementImcAllOf); err == nil {
		*o = CertificatemanagementImcAllOf(varCertificatemanagementImcAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CertType")
		delete(additionalProperties, "IsPrivatekeySet")
		delete(additionalProperties, "Privatekey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificatemanagementImcAllOf struct {
	value *CertificatemanagementImcAllOf
	isSet bool
}

func (v NullableCertificatemanagementImcAllOf) Get() *CertificatemanagementImcAllOf {
	return v.value
}

func (v *NullableCertificatemanagementImcAllOf) Set(val *CertificatemanagementImcAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificatemanagementImcAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificatemanagementImcAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificatemanagementImcAllOf(val *CertificatemanagementImcAllOf) *NullableCertificatemanagementImcAllOf {
	return &NullableCertificatemanagementImcAllOf{value: val, isSet: true}
}

func (v NullableCertificatemanagementImcAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificatemanagementImcAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
