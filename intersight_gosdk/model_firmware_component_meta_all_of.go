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

// FirmwareComponentMetaAllOf Definition of the list of properties defined in 'firmware.ComponentMeta', excluding properties defined in parent classes.
type FirmwareComponentMetaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the component in the compressed HSU bundle.
	ComponentLabel *string `json:"ComponentLabel,omitempty"`
	// The type of component image within the distributable. * `ALL` - This represents all the components. * `ALL,HDD` - This represents all the components plus the HDDs. * `Drive-U.2` - This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category). * `Storage` - This represents the storage controller components. * `None` - This represents none of the components. * `NXOS` - This represents NXOS components. * `IOM` - This represents IOM components. * `PSU` - This represents PSU components. * `CIMC` - This represents CIMC components. * `BIOS` - This represents BIOS components. * `PCIE` - This represents PCIE components. * `Drive` - This represents Drive components. * `DIMM` - This represents DIMM components. * `BoardController` - This represents Board Controller components. * `StorageController` - This represents Storage Controller components. * `Storage-Sasexpander` - This represents Storage Sas-Expander components. * `Storage-U.2` - This represents U2 Storage Controller components. * `HBA` - This represents HBA components. * `GPU` - This represents GPU components. * `SasExpander` - This represents SasExpander components. * `MSwitch` - This represents mSwitch components. * `CMC` - This represents CMC components.
	ComponentType *string `json:"ComponentType,omitempty"`
	// This shows the description of component image within the distributable.
	Description *string `json:"Description,omitempty"`
	// The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle. * `None` - Indicates that the component did not receive a disruption request. * `HostReboot` - Indicates that the component received a host reboot request. * `EndpointReboot` - Indicates that the component received an end point reboot request. * `ManualPowerCycle` - Indicates that the component received a manual power cycle request. * `AutomaticPowerCycle` - Indicates that the component received an automatic power cycle request.
	Disruption *string `json:"Disruption,omitempty"`
	// This shows the path of component image within the distributable.
	ImagePath *string `json:"ImagePath,omitempty"`
	// If set, the component can be updated through out-of-band management, else, is updated through host service utility boot.
	IsOobSupported *bool `json:"IsOobSupported,omitempty"`
	// The model of the component image in the distributable.
	Model            *string  `json:"Model,omitempty"`
	OobManageability []string `json:"OobManageability,omitempty"`
	// The image version of components packaged in the distributable.
	PackedVersion *string `json:"PackedVersion,omitempty"`
	// The redfish target for each component.
	RedfishUrl *string `json:"RedfishUrl,omitempty"`
	// The version of component image in the distributable.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareComponentMetaAllOf FirmwareComponentMetaAllOf

// NewFirmwareComponentMetaAllOf instantiates a new FirmwareComponentMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareComponentMetaAllOf(classId string, objectType string) *FirmwareComponentMetaAllOf {
	this := FirmwareComponentMetaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var componentType string = "ALL"
	this.ComponentType = &componentType
	var disruption string = "None"
	this.Disruption = &disruption
	return &this
}

// NewFirmwareComponentMetaAllOfWithDefaults instantiates a new FirmwareComponentMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareComponentMetaAllOfWithDefaults() *FirmwareComponentMetaAllOf {
	this := FirmwareComponentMetaAllOf{}
	var classId string = "firmware.ComponentMeta"
	this.ClassId = classId
	var objectType string = "firmware.ComponentMeta"
	this.ObjectType = objectType
	var componentType string = "ALL"
	this.ComponentType = &componentType
	var disruption string = "None"
	this.Disruption = &disruption
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareComponentMetaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareComponentMetaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareComponentMetaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareComponentMetaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComponentLabel returns the ComponentLabel field value if set, zero value otherwise.
func (o *FirmwareComponentMetaAllOf) GetComponentLabel() string {
	if o == nil || o.ComponentLabel == nil {
		var ret string
		return ret
	}
	return *o.ComponentLabel
}

// GetComponentLabelOk returns a tuple with the ComponentLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetComponentLabelOk() (*string, bool) {
	if o == nil || o.ComponentLabel == nil {
		return nil, false
	}
	return o.ComponentLabel, true
}

// HasComponentLabel returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasComponentLabel() bool {
	if o != nil && o.ComponentLabel != nil {
		return true
	}

	return false
}

// SetComponentLabel gets a reference to the given string and assigns it to the ComponentLabel field.
func (o *FirmwareComponentMetaAllOf) SetComponentLabel(v string) {
	o.ComponentLabel = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *FirmwareComponentMetaAllOf) GetComponentType() string {
	if o == nil || o.ComponentType == nil {
		var ret string
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetComponentTypeOk() (*string, bool) {
	if o == nil || o.ComponentType == nil {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasComponentType() bool {
	if o != nil && o.ComponentType != nil {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given string and assigns it to the ComponentType field.
func (o *FirmwareComponentMetaAllOf) SetComponentType(v string) {
	o.ComponentType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FirmwareComponentMetaAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FirmwareComponentMetaAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDisruption returns the Disruption field value if set, zero value otherwise.
func (o *FirmwareComponentMetaAllOf) GetDisruption() string {
	if o == nil || o.Disruption == nil {
		var ret string
		return ret
	}
	return *o.Disruption
}

// GetDisruptionOk returns a tuple with the Disruption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetDisruptionOk() (*string, bool) {
	if o == nil || o.Disruption == nil {
		return nil, false
	}
	return o.Disruption, true
}

// HasDisruption returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasDisruption() bool {
	if o != nil && o.Disruption != nil {
		return true
	}

	return false
}

// SetDisruption gets a reference to the given string and assigns it to the Disruption field.
func (o *FirmwareComponentMetaAllOf) SetDisruption(v string) {
	o.Disruption = &v
}

// GetImagePath returns the ImagePath field value if set, zero value otherwise.
func (o *FirmwareComponentMetaAllOf) GetImagePath() string {
	if o == nil || o.ImagePath == nil {
		var ret string
		return ret
	}
	return *o.ImagePath
}

// GetImagePathOk returns a tuple with the ImagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetImagePathOk() (*string, bool) {
	if o == nil || o.ImagePath == nil {
		return nil, false
	}
	return o.ImagePath, true
}

// HasImagePath returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasImagePath() bool {
	if o != nil && o.ImagePath != nil {
		return true
	}

	return false
}

// SetImagePath gets a reference to the given string and assigns it to the ImagePath field.
func (o *FirmwareComponentMetaAllOf) SetImagePath(v string) {
	o.ImagePath = &v
}

// GetIsOobSupported returns the IsOobSupported field value if set, zero value otherwise.
func (o *FirmwareComponentMetaAllOf) GetIsOobSupported() bool {
	if o == nil || o.IsOobSupported == nil {
		var ret bool
		return ret
	}
	return *o.IsOobSupported
}

// GetIsOobSupportedOk returns a tuple with the IsOobSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetIsOobSupportedOk() (*bool, bool) {
	if o == nil || o.IsOobSupported == nil {
		return nil, false
	}
	return o.IsOobSupported, true
}

// HasIsOobSupported returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasIsOobSupported() bool {
	if o != nil && o.IsOobSupported != nil {
		return true
	}

	return false
}

// SetIsOobSupported gets a reference to the given bool and assigns it to the IsOobSupported field.
func (o *FirmwareComponentMetaAllOf) SetIsOobSupported(v bool) {
	o.IsOobSupported = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *FirmwareComponentMetaAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *FirmwareComponentMetaAllOf) SetModel(v string) {
	o.Model = &v
}

// GetOobManageability returns the OobManageability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareComponentMetaAllOf) GetOobManageability() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OobManageability
}

// GetOobManageabilityOk returns a tuple with the OobManageability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareComponentMetaAllOf) GetOobManageabilityOk() ([]string, bool) {
	if o == nil || o.OobManageability == nil {
		return nil, false
	}
	return o.OobManageability, true
}

// HasOobManageability returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasOobManageability() bool {
	if o != nil && o.OobManageability != nil {
		return true
	}

	return false
}

// SetOobManageability gets a reference to the given []string and assigns it to the OobManageability field.
func (o *FirmwareComponentMetaAllOf) SetOobManageability(v []string) {
	o.OobManageability = v
}

// GetPackedVersion returns the PackedVersion field value if set, zero value otherwise.
func (o *FirmwareComponentMetaAllOf) GetPackedVersion() string {
	if o == nil || o.PackedVersion == nil {
		var ret string
		return ret
	}
	return *o.PackedVersion
}

// GetPackedVersionOk returns a tuple with the PackedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetPackedVersionOk() (*string, bool) {
	if o == nil || o.PackedVersion == nil {
		return nil, false
	}
	return o.PackedVersion, true
}

// HasPackedVersion returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasPackedVersion() bool {
	if o != nil && o.PackedVersion != nil {
		return true
	}

	return false
}

// SetPackedVersion gets a reference to the given string and assigns it to the PackedVersion field.
func (o *FirmwareComponentMetaAllOf) SetPackedVersion(v string) {
	o.PackedVersion = &v
}

// GetRedfishUrl returns the RedfishUrl field value if set, zero value otherwise.
func (o *FirmwareComponentMetaAllOf) GetRedfishUrl() string {
	if o == nil || o.RedfishUrl == nil {
		var ret string
		return ret
	}
	return *o.RedfishUrl
}

// GetRedfishUrlOk returns a tuple with the RedfishUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetRedfishUrlOk() (*string, bool) {
	if o == nil || o.RedfishUrl == nil {
		return nil, false
	}
	return o.RedfishUrl, true
}

// HasRedfishUrl returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasRedfishUrl() bool {
	if o != nil && o.RedfishUrl != nil {
		return true
	}

	return false
}

// SetRedfishUrl gets a reference to the given string and assigns it to the RedfishUrl field.
func (o *FirmwareComponentMetaAllOf) SetRedfishUrl(v string) {
	o.RedfishUrl = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *FirmwareComponentMetaAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareComponentMetaAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *FirmwareComponentMetaAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *FirmwareComponentMetaAllOf) SetVendor(v string) {
	o.Vendor = &v
}

func (o FirmwareComponentMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ComponentLabel != nil {
		toSerialize["ComponentLabel"] = o.ComponentLabel
	}
	if o.ComponentType != nil {
		toSerialize["ComponentType"] = o.ComponentType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Disruption != nil {
		toSerialize["Disruption"] = o.Disruption
	}
	if o.ImagePath != nil {
		toSerialize["ImagePath"] = o.ImagePath
	}
	if o.IsOobSupported != nil {
		toSerialize["IsOobSupported"] = o.IsOobSupported
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.OobManageability != nil {
		toSerialize["OobManageability"] = o.OobManageability
	}
	if o.PackedVersion != nil {
		toSerialize["PackedVersion"] = o.PackedVersion
	}
	if o.RedfishUrl != nil {
		toSerialize["RedfishUrl"] = o.RedfishUrl
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareComponentMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareComponentMetaAllOf := _FirmwareComponentMetaAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareComponentMetaAllOf); err == nil {
		*o = FirmwareComponentMetaAllOf(varFirmwareComponentMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ComponentLabel")
		delete(additionalProperties, "ComponentType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Disruption")
		delete(additionalProperties, "ImagePath")
		delete(additionalProperties, "IsOobSupported")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "OobManageability")
		delete(additionalProperties, "PackedVersion")
		delete(additionalProperties, "RedfishUrl")
		delete(additionalProperties, "Vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareComponentMetaAllOf struct {
	value *FirmwareComponentMetaAllOf
	isSet bool
}

func (v NullableFirmwareComponentMetaAllOf) Get() *FirmwareComponentMetaAllOf {
	return v.value
}

func (v *NullableFirmwareComponentMetaAllOf) Set(val *FirmwareComponentMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareComponentMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareComponentMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareComponentMetaAllOf(val *FirmwareComponentMetaAllOf) *NullableFirmwareComponentMetaAllOf {
	return &NullableFirmwareComponentMetaAllOf{value: val, isSet: true}
}

func (v NullableFirmwareComponentMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareComponentMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
