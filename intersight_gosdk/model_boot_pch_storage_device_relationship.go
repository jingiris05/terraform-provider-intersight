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
	"fmt"
)

// BootPchStorageDeviceRelationship - A relationship to the 'boot.PchStorageDevice' resource, or the expanded 'boot.PchStorageDevice' resource, or the 'null' value.
type BootPchStorageDeviceRelationship struct {
	BootPchStorageDevice *BootPchStorageDevice
	MoMoRef              *MoMoRef
}

// BootPchStorageDeviceAsBootPchStorageDeviceRelationship is a convenience function that returns BootPchStorageDevice wrapped in BootPchStorageDeviceRelationship
func BootPchStorageDeviceAsBootPchStorageDeviceRelationship(v *BootPchStorageDevice) BootPchStorageDeviceRelationship {
	return BootPchStorageDeviceRelationship{BootPchStorageDevice: v}
}

// MoMoRefAsBootPchStorageDeviceRelationship is a convenience function that returns MoMoRef wrapped in BootPchStorageDeviceRelationship
func MoMoRefAsBootPchStorageDeviceRelationship(v *MoMoRef) BootPchStorageDeviceRelationship {
	return BootPchStorageDeviceRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BootPchStorageDeviceRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'boot.PchStorageDevice'
	if jsonDict["ClassId"] == "boot.PchStorageDevice" {
		// try to unmarshal JSON data into BootPchStorageDevice
		err = json.Unmarshal(data, &dst.BootPchStorageDevice)
		if err == nil {
			return nil // data stored in dst.BootPchStorageDevice, return on the first match
		} else {
			dst.BootPchStorageDevice = nil
			return fmt.Errorf("Failed to unmarshal BootPchStorageDeviceRelationship as BootPchStorageDevice: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal BootPchStorageDeviceRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BootPchStorageDeviceRelationship) MarshalJSON() ([]byte, error) {
	if src.BootPchStorageDevice != nil {
		return json.Marshal(&src.BootPchStorageDevice)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BootPchStorageDeviceRelationship) GetActualInstance() interface{} {
	if obj.BootPchStorageDevice != nil {
		return obj.BootPchStorageDevice
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableBootPchStorageDeviceRelationship struct {
	value *BootPchStorageDeviceRelationship
	isSet bool
}

func (v NullableBootPchStorageDeviceRelationship) Get() *BootPchStorageDeviceRelationship {
	return v.value
}

func (v *NullableBootPchStorageDeviceRelationship) Set(val *BootPchStorageDeviceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableBootPchStorageDeviceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableBootPchStorageDeviceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootPchStorageDeviceRelationship(val *BootPchStorageDeviceRelationship) *NullableBootPchStorageDeviceRelationship {
	return &NullableBootPchStorageDeviceRelationship{value: val, isSet: true}
}

func (v NullableBootPchStorageDeviceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootPchStorageDeviceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
