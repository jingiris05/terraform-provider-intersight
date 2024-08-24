/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18012
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the HclExemptedCatalog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HclExemptedCatalog{}

// HclExemptedCatalog Collection used to store exempted products (ie. adapters, storage controllers, etc). These products should be ignored for HCL validation purposes.
type HclExemptedCatalog struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Reason for the exemption.
	Comments *string `json:"Comments,omitempty"`
	// A unique descriptive name of the exemption.
	Name *string `json:"Name,omitempty"`
	// Vendor of the Operating System.
	OsVendor *string `json:"OsVendor,omitempty"`
	// Version of the Operating system.
	OsVersion *string `json:"OsVersion,omitempty"`
	// It indicates the personality of the sever.
	Personality *string `json:"Personality,omitempty"`
	// Name of the processor supported for the server.
	ProcessorName *string  `json:"ProcessorName,omitempty"`
	ProductModels []string `json:"ProductModels,omitempty"`
	// Type of the product/adapter say GPU for graphic cards. * `` - Default type of the Product. * `Adapter` - Represents network adapter cards. * `StorageController` - Represents storage controllers. * `GPU` - Represents graphics cards.
	ProductType *string `json:"ProductType,omitempty"`
	// Three part ID representing the server model as returned by UCSM/CIMC XML APIs.
	ServerPid *string `json:"ServerPid,omitempty"`
	// Version of the UCS software.
	UcsVersion *string `json:"UcsVersion,omitempty"`
	// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
	VersionType          *string `json:"VersionType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclExemptedCatalog HclExemptedCatalog

// NewHclExemptedCatalog instantiates a new HclExemptedCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclExemptedCatalog(classId string, objectType string) *HclExemptedCatalog {
	this := HclExemptedCatalog{}
	this.ClassId = classId
	this.ObjectType = objectType
	var productType string = ""
	this.ProductType = &productType
	return &this
}

// NewHclExemptedCatalogWithDefaults instantiates a new HclExemptedCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclExemptedCatalogWithDefaults() *HclExemptedCatalog {
	this := HclExemptedCatalog{}
	var classId string = "hcl.ExemptedCatalog"
	this.ClassId = classId
	var objectType string = "hcl.ExemptedCatalog"
	this.ObjectType = objectType
	var productType string = ""
	this.ProductType = &productType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclExemptedCatalog) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclExemptedCatalog) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hcl.ExemptedCatalog" of the ClassId field.
func (o *HclExemptedCatalog) GetDefaultClassId() interface{} {
	return "hcl.ExemptedCatalog"
}

// GetObjectType returns the ObjectType field value
func (o *HclExemptedCatalog) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclExemptedCatalog) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hcl.ExemptedCatalog" of the ObjectType field.
func (o *HclExemptedCatalog) GetDefaultObjectType() interface{} {
	return "hcl.ExemptedCatalog"
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *HclExemptedCatalog) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *HclExemptedCatalog) SetComments(v string) {
	o.Comments = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HclExemptedCatalog) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HclExemptedCatalog) SetName(v string) {
	o.Name = &v
}

// GetOsVendor returns the OsVendor field value if set, zero value otherwise.
func (o *HclExemptedCatalog) GetOsVendor() string {
	if o == nil || IsNil(o.OsVendor) {
		var ret string
		return ret
	}
	return *o.OsVendor
}

// GetOsVendorOk returns a tuple with the OsVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetOsVendorOk() (*string, bool) {
	if o == nil || IsNil(o.OsVendor) {
		return nil, false
	}
	return o.OsVendor, true
}

// HasOsVendor returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasOsVendor() bool {
	if o != nil && !IsNil(o.OsVendor) {
		return true
	}

	return false
}

// SetOsVendor gets a reference to the given string and assigns it to the OsVendor field.
func (o *HclExemptedCatalog) SetOsVendor(v string) {
	o.OsVendor = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *HclExemptedCatalog) GetOsVersion() string {
	if o == nil || IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetOsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *HclExemptedCatalog) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetPersonality returns the Personality field value if set, zero value otherwise.
func (o *HclExemptedCatalog) GetPersonality() string {
	if o == nil || IsNil(o.Personality) {
		var ret string
		return ret
	}
	return *o.Personality
}

// GetPersonalityOk returns a tuple with the Personality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetPersonalityOk() (*string, bool) {
	if o == nil || IsNil(o.Personality) {
		return nil, false
	}
	return o.Personality, true
}

// HasPersonality returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasPersonality() bool {
	if o != nil && !IsNil(o.Personality) {
		return true
	}

	return false
}

// SetPersonality gets a reference to the given string and assigns it to the Personality field.
func (o *HclExemptedCatalog) SetPersonality(v string) {
	o.Personality = &v
}

// GetProcessorName returns the ProcessorName field value if set, zero value otherwise.
func (o *HclExemptedCatalog) GetProcessorName() string {
	if o == nil || IsNil(o.ProcessorName) {
		var ret string
		return ret
	}
	return *o.ProcessorName
}

// GetProcessorNameOk returns a tuple with the ProcessorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetProcessorNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorName) {
		return nil, false
	}
	return o.ProcessorName, true
}

// HasProcessorName returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasProcessorName() bool {
	if o != nil && !IsNil(o.ProcessorName) {
		return true
	}

	return false
}

// SetProcessorName gets a reference to the given string and assigns it to the ProcessorName field.
func (o *HclExemptedCatalog) SetProcessorName(v string) {
	o.ProcessorName = &v
}

// GetProductModels returns the ProductModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclExemptedCatalog) GetProductModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProductModels
}

// GetProductModelsOk returns a tuple with the ProductModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclExemptedCatalog) GetProductModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductModels) {
		return nil, false
	}
	return o.ProductModels, true
}

// HasProductModels returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasProductModels() bool {
	if o != nil && !IsNil(o.ProductModels) {
		return true
	}

	return false
}

// SetProductModels gets a reference to the given []string and assigns it to the ProductModels field.
func (o *HclExemptedCatalog) SetProductModels(v []string) {
	o.ProductModels = v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *HclExemptedCatalog) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *HclExemptedCatalog) SetProductType(v string) {
	o.ProductType = &v
}

// GetServerPid returns the ServerPid field value if set, zero value otherwise.
func (o *HclExemptedCatalog) GetServerPid() string {
	if o == nil || IsNil(o.ServerPid) {
		var ret string
		return ret
	}
	return *o.ServerPid
}

// GetServerPidOk returns a tuple with the ServerPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetServerPidOk() (*string, bool) {
	if o == nil || IsNil(o.ServerPid) {
		return nil, false
	}
	return o.ServerPid, true
}

// HasServerPid returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasServerPid() bool {
	if o != nil && !IsNil(o.ServerPid) {
		return true
	}

	return false
}

// SetServerPid gets a reference to the given string and assigns it to the ServerPid field.
func (o *HclExemptedCatalog) SetServerPid(v string) {
	o.ServerPid = &v
}

// GetUcsVersion returns the UcsVersion field value if set, zero value otherwise.
func (o *HclExemptedCatalog) GetUcsVersion() string {
	if o == nil || IsNil(o.UcsVersion) {
		var ret string
		return ret
	}
	return *o.UcsVersion
}

// GetUcsVersionOk returns a tuple with the UcsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetUcsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.UcsVersion) {
		return nil, false
	}
	return o.UcsVersion, true
}

// HasUcsVersion returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasUcsVersion() bool {
	if o != nil && !IsNil(o.UcsVersion) {
		return true
	}

	return false
}

// SetUcsVersion gets a reference to the given string and assigns it to the UcsVersion field.
func (o *HclExemptedCatalog) SetUcsVersion(v string) {
	o.UcsVersion = &v
}

// GetVersionType returns the VersionType field value if set, zero value otherwise.
func (o *HclExemptedCatalog) GetVersionType() string {
	if o == nil || IsNil(o.VersionType) {
		var ret string
		return ret
	}
	return *o.VersionType
}

// GetVersionTypeOk returns a tuple with the VersionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalog) GetVersionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VersionType) {
		return nil, false
	}
	return o.VersionType, true
}

// HasVersionType returns a boolean if a field has been set.
func (o *HclExemptedCatalog) HasVersionType() bool {
	if o != nil && !IsNil(o.VersionType) {
		return true
	}

	return false
}

// SetVersionType gets a reference to the given string and assigns it to the VersionType field.
func (o *HclExemptedCatalog) SetVersionType(v string) {
	o.VersionType = &v
}

func (o HclExemptedCatalog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HclExemptedCatalog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Comments) {
		toSerialize["Comments"] = o.Comments
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.OsVendor) {
		toSerialize["OsVendor"] = o.OsVendor
	}
	if !IsNil(o.OsVersion) {
		toSerialize["OsVersion"] = o.OsVersion
	}
	if !IsNil(o.Personality) {
		toSerialize["Personality"] = o.Personality
	}
	if !IsNil(o.ProcessorName) {
		toSerialize["ProcessorName"] = o.ProcessorName
	}
	if o.ProductModels != nil {
		toSerialize["ProductModels"] = o.ProductModels
	}
	if !IsNil(o.ProductType) {
		toSerialize["ProductType"] = o.ProductType
	}
	if !IsNil(o.ServerPid) {
		toSerialize["ServerPid"] = o.ServerPid
	}
	if !IsNil(o.UcsVersion) {
		toSerialize["UcsVersion"] = o.UcsVersion
	}
	if !IsNil(o.VersionType) {
		toSerialize["VersionType"] = o.VersionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HclExemptedCatalog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type HclExemptedCatalogWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Reason for the exemption.
		Comments *string `json:"Comments,omitempty"`
		// A unique descriptive name of the exemption.
		Name *string `json:"Name,omitempty"`
		// Vendor of the Operating System.
		OsVendor *string `json:"OsVendor,omitempty"`
		// Version of the Operating system.
		OsVersion *string `json:"OsVersion,omitempty"`
		// It indicates the personality of the sever.
		Personality *string `json:"Personality,omitempty"`
		// Name of the processor supported for the server.
		ProcessorName *string  `json:"ProcessorName,omitempty"`
		ProductModels []string `json:"ProductModels,omitempty"`
		// Type of the product/adapter say GPU for graphic cards. * `` - Default type of the Product. * `Adapter` - Represents network adapter cards. * `StorageController` - Represents storage controllers. * `GPU` - Represents graphics cards.
		ProductType *string `json:"ProductType,omitempty"`
		// Three part ID representing the server model as returned by UCSM/CIMC XML APIs.
		ServerPid *string `json:"ServerPid,omitempty"`
		// Version of the UCS software.
		UcsVersion *string `json:"UcsVersion,omitempty"`
		// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
		VersionType *string `json:"VersionType,omitempty"`
	}

	varHclExemptedCatalogWithoutEmbeddedStruct := HclExemptedCatalogWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHclExemptedCatalogWithoutEmbeddedStruct)
	if err == nil {
		varHclExemptedCatalog := _HclExemptedCatalog{}
		varHclExemptedCatalog.ClassId = varHclExemptedCatalogWithoutEmbeddedStruct.ClassId
		varHclExemptedCatalog.ObjectType = varHclExemptedCatalogWithoutEmbeddedStruct.ObjectType
		varHclExemptedCatalog.Comments = varHclExemptedCatalogWithoutEmbeddedStruct.Comments
		varHclExemptedCatalog.Name = varHclExemptedCatalogWithoutEmbeddedStruct.Name
		varHclExemptedCatalog.OsVendor = varHclExemptedCatalogWithoutEmbeddedStruct.OsVendor
		varHclExemptedCatalog.OsVersion = varHclExemptedCatalogWithoutEmbeddedStruct.OsVersion
		varHclExemptedCatalog.Personality = varHclExemptedCatalogWithoutEmbeddedStruct.Personality
		varHclExemptedCatalog.ProcessorName = varHclExemptedCatalogWithoutEmbeddedStruct.ProcessorName
		varHclExemptedCatalog.ProductModels = varHclExemptedCatalogWithoutEmbeddedStruct.ProductModels
		varHclExemptedCatalog.ProductType = varHclExemptedCatalogWithoutEmbeddedStruct.ProductType
		varHclExemptedCatalog.ServerPid = varHclExemptedCatalogWithoutEmbeddedStruct.ServerPid
		varHclExemptedCatalog.UcsVersion = varHclExemptedCatalogWithoutEmbeddedStruct.UcsVersion
		varHclExemptedCatalog.VersionType = varHclExemptedCatalogWithoutEmbeddedStruct.VersionType
		*o = HclExemptedCatalog(varHclExemptedCatalog)
	} else {
		return err
	}

	varHclExemptedCatalog := _HclExemptedCatalog{}

	err = json.Unmarshal(data, &varHclExemptedCatalog)
	if err == nil {
		o.MoBaseMo = varHclExemptedCatalog.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Comments")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OsVendor")
		delete(additionalProperties, "OsVersion")
		delete(additionalProperties, "Personality")
		delete(additionalProperties, "ProcessorName")
		delete(additionalProperties, "ProductModels")
		delete(additionalProperties, "ProductType")
		delete(additionalProperties, "ServerPid")
		delete(additionalProperties, "UcsVersion")
		delete(additionalProperties, "VersionType")

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

type NullableHclExemptedCatalog struct {
	value *HclExemptedCatalog
	isSet bool
}

func (v NullableHclExemptedCatalog) Get() *HclExemptedCatalog {
	return v.value
}

func (v *NullableHclExemptedCatalog) Set(val *HclExemptedCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableHclExemptedCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableHclExemptedCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclExemptedCatalog(val *HclExemptedCatalog) *NullableHclExemptedCatalog {
	return &NullableHclExemptedCatalog{value: val, isSet: true}
}

func (v NullableHclExemptedCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclExemptedCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
