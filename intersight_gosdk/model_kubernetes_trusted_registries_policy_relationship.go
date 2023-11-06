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
	"fmt"
)

// KubernetesTrustedRegistriesPolicyRelationship - A relationship to the 'kubernetes.TrustedRegistriesPolicy' resource, or the expanded 'kubernetes.TrustedRegistriesPolicy' resource, or the 'null' value.
type KubernetesTrustedRegistriesPolicyRelationship struct {
	KubernetesTrustedRegistriesPolicy *KubernetesTrustedRegistriesPolicy
	MoMoRef                           *MoMoRef
}

// KubernetesTrustedRegistriesPolicyAsKubernetesTrustedRegistriesPolicyRelationship is a convenience function that returns KubernetesTrustedRegistriesPolicy wrapped in KubernetesTrustedRegistriesPolicyRelationship
func KubernetesTrustedRegistriesPolicyAsKubernetesTrustedRegistriesPolicyRelationship(v *KubernetesTrustedRegistriesPolicy) KubernetesTrustedRegistriesPolicyRelationship {
	return KubernetesTrustedRegistriesPolicyRelationship{
		KubernetesTrustedRegistriesPolicy: v,
	}
}

// MoMoRefAsKubernetesTrustedRegistriesPolicyRelationship is a convenience function that returns MoMoRef wrapped in KubernetesTrustedRegistriesPolicyRelationship
func MoMoRefAsKubernetesTrustedRegistriesPolicyRelationship(v *MoMoRef) KubernetesTrustedRegistriesPolicyRelationship {
	return KubernetesTrustedRegistriesPolicyRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *KubernetesTrustedRegistriesPolicyRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'kubernetes.TrustedRegistriesPolicy'
	if jsonDict["ClassId"] == "kubernetes.TrustedRegistriesPolicy" {
		// try to unmarshal JSON data into KubernetesTrustedRegistriesPolicy
		err = json.Unmarshal(data, &dst.KubernetesTrustedRegistriesPolicy)
		if err == nil {
			return nil // data stored in dst.KubernetesTrustedRegistriesPolicy, return on the first match
		} else {
			dst.KubernetesTrustedRegistriesPolicy = nil
			return fmt.Errorf("Failed to unmarshal KubernetesTrustedRegistriesPolicyRelationship as KubernetesTrustedRegistriesPolicy: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal KubernetesTrustedRegistriesPolicyRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src KubernetesTrustedRegistriesPolicyRelationship) MarshalJSON() ([]byte, error) {
	if src.KubernetesTrustedRegistriesPolicy != nil {
		return json.Marshal(&src.KubernetesTrustedRegistriesPolicy)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *KubernetesTrustedRegistriesPolicyRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.KubernetesTrustedRegistriesPolicy != nil {
		return obj.KubernetesTrustedRegistriesPolicy
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableKubernetesTrustedRegistriesPolicyRelationship struct {
	value *KubernetesTrustedRegistriesPolicyRelationship
	isSet bool
}

func (v NullableKubernetesTrustedRegistriesPolicyRelationship) Get() *KubernetesTrustedRegistriesPolicyRelationship {
	return v.value
}

func (v *NullableKubernetesTrustedRegistriesPolicyRelationship) Set(val *KubernetesTrustedRegistriesPolicyRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesTrustedRegistriesPolicyRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesTrustedRegistriesPolicyRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesTrustedRegistriesPolicyRelationship(val *KubernetesTrustedRegistriesPolicyRelationship) *NullableKubernetesTrustedRegistriesPolicyRelationship {
	return &NullableKubernetesTrustedRegistriesPolicyRelationship{value: val, isSet: true}
}

func (v NullableKubernetesTrustedRegistriesPolicyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesTrustedRegistriesPolicyRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
