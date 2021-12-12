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
)

// IaasConnectorPackAllOf Definition of the list of properties defined in 'iaas.ConnectorPack', excluding properties defined in parent classes.
type IaasConnectorPackAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Complete version of the connector pack including build number.
	CompleteVersion *string  `json:"CompleteVersion,omitempty"`
	DependencyNames []string `json:"DependencyNames,omitempty"`
	// Version of the connector pack that is last downloaded successfully to UCSD.
	DownloadedVersion *string `json:"DownloadedVersion,omitempty"`
	// Name of the connector pack running on the UCSD.
	Name *string `json:"Name,omitempty"`
	// State of the connector pack whether it is enabled or disabled.
	State *string `json:"State,omitempty"`
	// Version of the connector pack.
	Version              *string                   `json:"Version,omitempty"`
	Guid                 *IaasUcsdInfoRelationship `json:"Guid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasConnectorPackAllOf IaasConnectorPackAllOf

// NewIaasConnectorPackAllOf instantiates a new IaasConnectorPackAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasConnectorPackAllOf(classId string, objectType string) *IaasConnectorPackAllOf {
	this := IaasConnectorPackAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasConnectorPackAllOfWithDefaults instantiates a new IaasConnectorPackAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasConnectorPackAllOfWithDefaults() *IaasConnectorPackAllOf {
	this := IaasConnectorPackAllOf{}
	var classId string = "iaas.ConnectorPack"
	this.ClassId = classId
	var objectType string = "iaas.ConnectorPack"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasConnectorPackAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasConnectorPackAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasConnectorPackAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasConnectorPackAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasConnectorPackAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasConnectorPackAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCompleteVersion returns the CompleteVersion field value if set, zero value otherwise.
func (o *IaasConnectorPackAllOf) GetCompleteVersion() string {
	if o == nil || o.CompleteVersion == nil {
		var ret string
		return ret
	}
	return *o.CompleteVersion
}

// GetCompleteVersionOk returns a tuple with the CompleteVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPackAllOf) GetCompleteVersionOk() (*string, bool) {
	if o == nil || o.CompleteVersion == nil {
		return nil, false
	}
	return o.CompleteVersion, true
}

// HasCompleteVersion returns a boolean if a field has been set.
func (o *IaasConnectorPackAllOf) HasCompleteVersion() bool {
	if o != nil && o.CompleteVersion != nil {
		return true
	}

	return false
}

// SetCompleteVersion gets a reference to the given string and assigns it to the CompleteVersion field.
func (o *IaasConnectorPackAllOf) SetCompleteVersion(v string) {
	o.CompleteVersion = &v
}

// GetDependencyNames returns the DependencyNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasConnectorPackAllOf) GetDependencyNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DependencyNames
}

// GetDependencyNamesOk returns a tuple with the DependencyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasConnectorPackAllOf) GetDependencyNamesOk() (*[]string, bool) {
	if o == nil || o.DependencyNames == nil {
		return nil, false
	}
	return &o.DependencyNames, true
}

// HasDependencyNames returns a boolean if a field has been set.
func (o *IaasConnectorPackAllOf) HasDependencyNames() bool {
	if o != nil && o.DependencyNames != nil {
		return true
	}

	return false
}

// SetDependencyNames gets a reference to the given []string and assigns it to the DependencyNames field.
func (o *IaasConnectorPackAllOf) SetDependencyNames(v []string) {
	o.DependencyNames = v
}

// GetDownloadedVersion returns the DownloadedVersion field value if set, zero value otherwise.
func (o *IaasConnectorPackAllOf) GetDownloadedVersion() string {
	if o == nil || o.DownloadedVersion == nil {
		var ret string
		return ret
	}
	return *o.DownloadedVersion
}

// GetDownloadedVersionOk returns a tuple with the DownloadedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPackAllOf) GetDownloadedVersionOk() (*string, bool) {
	if o == nil || o.DownloadedVersion == nil {
		return nil, false
	}
	return o.DownloadedVersion, true
}

// HasDownloadedVersion returns a boolean if a field has been set.
func (o *IaasConnectorPackAllOf) HasDownloadedVersion() bool {
	if o != nil && o.DownloadedVersion != nil {
		return true
	}

	return false
}

// SetDownloadedVersion gets a reference to the given string and assigns it to the DownloadedVersion field.
func (o *IaasConnectorPackAllOf) SetDownloadedVersion(v string) {
	o.DownloadedVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IaasConnectorPackAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPackAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IaasConnectorPackAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IaasConnectorPackAllOf) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IaasConnectorPackAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPackAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IaasConnectorPackAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IaasConnectorPackAllOf) SetState(v string) {
	o.State = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IaasConnectorPackAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPackAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IaasConnectorPackAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *IaasConnectorPackAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasConnectorPackAllOf) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || o.Guid == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPackAllOf) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasConnectorPackAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given IaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasConnectorPackAllOf) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid = &v
}

func (o IaasConnectorPackAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CompleteVersion != nil {
		toSerialize["CompleteVersion"] = o.CompleteVersion
	}
	if o.DependencyNames != nil {
		toSerialize["DependencyNames"] = o.DependencyNames
	}
	if o.DownloadedVersion != nil {
		toSerialize["DownloadedVersion"] = o.DownloadedVersion
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasConnectorPackAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIaasConnectorPackAllOf := _IaasConnectorPackAllOf{}

	if err = json.Unmarshal(bytes, &varIaasConnectorPackAllOf); err == nil {
		*o = IaasConnectorPackAllOf(varIaasConnectorPackAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CompleteVersion")
		delete(additionalProperties, "DependencyNames")
		delete(additionalProperties, "DownloadedVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Guid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIaasConnectorPackAllOf struct {
	value *IaasConnectorPackAllOf
	isSet bool
}

func (v NullableIaasConnectorPackAllOf) Get() *IaasConnectorPackAllOf {
	return v.value
}

func (v *NullableIaasConnectorPackAllOf) Set(val *IaasConnectorPackAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasConnectorPackAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasConnectorPackAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasConnectorPackAllOf(val *IaasConnectorPackAllOf) *NullableIaasConnectorPackAllOf {
	return &NullableIaasConnectorPackAllOf{value: val, isSet: true}
}

func (v NullableIaasConnectorPackAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasConnectorPackAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
