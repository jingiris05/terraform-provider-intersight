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
)

// IamEndPointUserRoleInventoryRelationship - A relationship to the 'iam.EndPointUserRoleInventory' resource, or the expanded 'iam.EndPointUserRoleInventory' resource, or the 'null' value.
type IamEndPointUserRoleInventoryRelationship struct {
	IamEndPointUserRoleInventory *IamEndPointUserRoleInventory
	MoMoRef                      *MoMoRef
}

// IamEndPointUserRoleInventoryAsIamEndPointUserRoleInventoryRelationship is a convenience function that returns IamEndPointUserRoleInventory wrapped in IamEndPointUserRoleInventoryRelationship
func IamEndPointUserRoleInventoryAsIamEndPointUserRoleInventoryRelationship(v *IamEndPointUserRoleInventory) IamEndPointUserRoleInventoryRelationship {
	return IamEndPointUserRoleInventoryRelationship{
		IamEndPointUserRoleInventory: v,
	}
}

// MoMoRefAsIamEndPointUserRoleInventoryRelationship is a convenience function that returns MoMoRef wrapped in IamEndPointUserRoleInventoryRelationship
func MoMoRefAsIamEndPointUserRoleInventoryRelationship(v *MoMoRef) IamEndPointUserRoleInventoryRelationship {
	return IamEndPointUserRoleInventoryRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IamEndPointUserRoleInventoryRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'iam.EndPointUserRoleInventory'
	if jsonDict["ClassId"] == "iam.EndPointUserRoleInventory" {
		// try to unmarshal JSON data into IamEndPointUserRoleInventory
		err = json.Unmarshal(data, &dst.IamEndPointUserRoleInventory)
		if err == nil {
			return nil // data stored in dst.IamEndPointUserRoleInventory, return on the first match
		} else {
			dst.IamEndPointUserRoleInventory = nil
			return fmt.Errorf("failed to unmarshal IamEndPointUserRoleInventoryRelationship as IamEndPointUserRoleInventory: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal IamEndPointUserRoleInventoryRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IamEndPointUserRoleInventoryRelationship) MarshalJSON() ([]byte, error) {
	if src.IamEndPointUserRoleInventory != nil {
		return json.Marshal(&src.IamEndPointUserRoleInventory)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IamEndPointUserRoleInventoryRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IamEndPointUserRoleInventory != nil {
		return obj.IamEndPointUserRoleInventory
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableIamEndPointUserRoleInventoryRelationship struct {
	value *IamEndPointUserRoleInventoryRelationship
	isSet bool
}

func (v NullableIamEndPointUserRoleInventoryRelationship) Get() *IamEndPointUserRoleInventoryRelationship {
	return v.value
}

func (v *NullableIamEndPointUserRoleInventoryRelationship) Set(val *IamEndPointUserRoleInventoryRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserRoleInventoryRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserRoleInventoryRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserRoleInventoryRelationship(val *IamEndPointUserRoleInventoryRelationship) *NullableIamEndPointUserRoleInventoryRelationship {
	return &NullableIamEndPointUserRoleInventoryRelationship{value: val, isSet: true}
}

func (v NullableIamEndPointUserRoleInventoryRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserRoleInventoryRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
