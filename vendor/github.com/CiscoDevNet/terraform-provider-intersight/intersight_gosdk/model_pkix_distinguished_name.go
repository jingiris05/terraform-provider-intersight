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

// PkixDistinguishedName The identifier for the owner of an X.509 certificate and the authority that issued the certificate.
type PkixDistinguishedName struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A required component that identifies a person or an object.
	CommonName           *string  `json:"CommonName,omitempty"`
	Country              []string `json:"Country,omitempty"`
	Locality             []string `json:"Locality,omitempty"`
	Organization         []string `json:"Organization,omitempty"`
	OrganizationalUnit   []string `json:"OrganizationalUnit,omitempty"`
	State                []string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PkixDistinguishedName PkixDistinguishedName

// NewPkixDistinguishedName instantiates a new PkixDistinguishedName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixDistinguishedName(classId string, objectType string) *PkixDistinguishedName {
	this := PkixDistinguishedName{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPkixDistinguishedNameWithDefaults instantiates a new PkixDistinguishedName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixDistinguishedNameWithDefaults() *PkixDistinguishedName {
	this := PkixDistinguishedName{}
	var classId string = "pkix.DistinguishedName"
	this.ClassId = classId
	var objectType string = "pkix.DistinguishedName"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PkixDistinguishedName) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PkixDistinguishedName) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PkixDistinguishedName) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PkixDistinguishedName) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PkixDistinguishedName) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PkixDistinguishedName) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *PkixDistinguishedName) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixDistinguishedName) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *PkixDistinguishedName) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedName) GetCountry() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedName) GetCountryOk() (*[]string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return &o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given []string and assigns it to the Country field.
func (o *PkixDistinguishedName) SetCountry(v []string) {
	o.Country = v
}

// GetLocality returns the Locality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedName) GetLocality() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedName) GetLocalityOk() (*[]string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return &o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given []string and assigns it to the Locality field.
func (o *PkixDistinguishedName) SetLocality(v []string) {
	o.Locality = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedName) GetOrganization() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedName) GetOrganizationOk() (*[]string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return &o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given []string and assigns it to the Organization field.
func (o *PkixDistinguishedName) SetOrganization(v []string) {
	o.Organization = v
}

// GetOrganizationalUnit returns the OrganizationalUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedName) GetOrganizationalUnit() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OrganizationalUnit
}

// GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedName) GetOrganizationalUnitOk() (*[]string, bool) {
	if o == nil || o.OrganizationalUnit == nil {
		return nil, false
	}
	return &o.OrganizationalUnit, true
}

// HasOrganizationalUnit returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasOrganizationalUnit() bool {
	if o != nil && o.OrganizationalUnit != nil {
		return true
	}

	return false
}

// SetOrganizationalUnit gets a reference to the given []string and assigns it to the OrganizationalUnit field.
func (o *PkixDistinguishedName) SetOrganizationalUnit(v []string) {
	o.OrganizationalUnit = v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PkixDistinguishedName) GetState() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PkixDistinguishedName) GetStateOk() (*[]string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PkixDistinguishedName) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given []string and assigns it to the State field.
func (o *PkixDistinguishedName) SetState(v []string) {
	o.State = v
}

func (o PkixDistinguishedName) MarshalJSON() ([]byte, error) {
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
	if o.CommonName != nil {
		toSerialize["CommonName"] = o.CommonName
	}
	if o.Country != nil {
		toSerialize["Country"] = o.Country
	}
	if o.Locality != nil {
		toSerialize["Locality"] = o.Locality
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.OrganizationalUnit != nil {
		toSerialize["OrganizationalUnit"] = o.OrganizationalUnit
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PkixDistinguishedName) UnmarshalJSON(bytes []byte) (err error) {
	type PkixDistinguishedNameWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A required component that identifies a person or an object.
		CommonName         *string  `json:"CommonName,omitempty"`
		Country            []string `json:"Country,omitempty"`
		Locality           []string `json:"Locality,omitempty"`
		Organization       []string `json:"Organization,omitempty"`
		OrganizationalUnit []string `json:"OrganizationalUnit,omitempty"`
		State              []string `json:"State,omitempty"`
	}

	varPkixDistinguishedNameWithoutEmbeddedStruct := PkixDistinguishedNameWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPkixDistinguishedNameWithoutEmbeddedStruct)
	if err == nil {
		varPkixDistinguishedName := _PkixDistinguishedName{}
		varPkixDistinguishedName.ClassId = varPkixDistinguishedNameWithoutEmbeddedStruct.ClassId
		varPkixDistinguishedName.ObjectType = varPkixDistinguishedNameWithoutEmbeddedStruct.ObjectType
		varPkixDistinguishedName.CommonName = varPkixDistinguishedNameWithoutEmbeddedStruct.CommonName
		varPkixDistinguishedName.Country = varPkixDistinguishedNameWithoutEmbeddedStruct.Country
		varPkixDistinguishedName.Locality = varPkixDistinguishedNameWithoutEmbeddedStruct.Locality
		varPkixDistinguishedName.Organization = varPkixDistinguishedNameWithoutEmbeddedStruct.Organization
		varPkixDistinguishedName.OrganizationalUnit = varPkixDistinguishedNameWithoutEmbeddedStruct.OrganizationalUnit
		varPkixDistinguishedName.State = varPkixDistinguishedNameWithoutEmbeddedStruct.State
		*o = PkixDistinguishedName(varPkixDistinguishedName)
	} else {
		return err
	}

	varPkixDistinguishedName := _PkixDistinguishedName{}

	err = json.Unmarshal(bytes, &varPkixDistinguishedName)
	if err == nil {
		o.MoBaseComplexType = varPkixDistinguishedName.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommonName")
		delete(additionalProperties, "Country")
		delete(additionalProperties, "Locality")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "OrganizationalUnit")
		delete(additionalProperties, "State")

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

type NullablePkixDistinguishedName struct {
	value *PkixDistinguishedName
	isSet bool
}

func (v NullablePkixDistinguishedName) Get() *PkixDistinguishedName {
	return v.value
}

func (v *NullablePkixDistinguishedName) Set(val *PkixDistinguishedName) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixDistinguishedName) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixDistinguishedName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixDistinguishedName(val *PkixDistinguishedName) *NullablePkixDistinguishedName {
	return &NullablePkixDistinguishedName{value: val, isSet: true}
}

func (v NullablePkixDistinguishedName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixDistinguishedName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
