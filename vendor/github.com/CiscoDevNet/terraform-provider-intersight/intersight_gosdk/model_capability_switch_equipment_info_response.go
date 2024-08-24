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

// CapabilitySwitchEquipmentInfoResponse - The response body of a HTTP GET request for the 'capability.SwitchEquipmentInfo' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'capability.SwitchEquipmentInfo' resources.
type CapabilitySwitchEquipmentInfoResponse struct {
	CapabilitySwitchEquipmentInfoList *CapabilitySwitchEquipmentInfoList
	MoAggregateTransform              *MoAggregateTransform
	MoDocumentCount                   *MoDocumentCount
	MoTagSummary                      *MoTagSummary
}

// CapabilitySwitchEquipmentInfoListAsCapabilitySwitchEquipmentInfoResponse is a convenience function that returns CapabilitySwitchEquipmentInfoList wrapped in CapabilitySwitchEquipmentInfoResponse
func CapabilitySwitchEquipmentInfoListAsCapabilitySwitchEquipmentInfoResponse(v *CapabilitySwitchEquipmentInfoList) CapabilitySwitchEquipmentInfoResponse {
	return CapabilitySwitchEquipmentInfoResponse{
		CapabilitySwitchEquipmentInfoList: v,
	}
}

// MoAggregateTransformAsCapabilitySwitchEquipmentInfoResponse is a convenience function that returns MoAggregateTransform wrapped in CapabilitySwitchEquipmentInfoResponse
func MoAggregateTransformAsCapabilitySwitchEquipmentInfoResponse(v *MoAggregateTransform) CapabilitySwitchEquipmentInfoResponse {
	return CapabilitySwitchEquipmentInfoResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsCapabilitySwitchEquipmentInfoResponse is a convenience function that returns MoDocumentCount wrapped in CapabilitySwitchEquipmentInfoResponse
func MoDocumentCountAsCapabilitySwitchEquipmentInfoResponse(v *MoDocumentCount) CapabilitySwitchEquipmentInfoResponse {
	return CapabilitySwitchEquipmentInfoResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsCapabilitySwitchEquipmentInfoResponse is a convenience function that returns MoTagSummary wrapped in CapabilitySwitchEquipmentInfoResponse
func MoTagSummaryAsCapabilitySwitchEquipmentInfoResponse(v *MoTagSummary) CapabilitySwitchEquipmentInfoResponse {
	return CapabilitySwitchEquipmentInfoResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CapabilitySwitchEquipmentInfoResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'capability.SwitchEquipmentInfo.List'
	if jsonDict["ObjectType"] == "capability.SwitchEquipmentInfo.List" {
		// try to unmarshal JSON data into CapabilitySwitchEquipmentInfoList
		err = json.Unmarshal(data, &dst.CapabilitySwitchEquipmentInfoList)
		if err == nil {
			return nil // data stored in dst.CapabilitySwitchEquipmentInfoList, return on the first match
		} else {
			dst.CapabilitySwitchEquipmentInfoList = nil
			return fmt.Errorf("failed to unmarshal CapabilitySwitchEquipmentInfoResponse as CapabilitySwitchEquipmentInfoList: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("failed to unmarshal CapabilitySwitchEquipmentInfoResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("failed to unmarshal CapabilitySwitchEquipmentInfoResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("failed to unmarshal CapabilitySwitchEquipmentInfoResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CapabilitySwitchEquipmentInfoResponse) MarshalJSON() ([]byte, error) {
	if src.CapabilitySwitchEquipmentInfoList != nil {
		return json.Marshal(&src.CapabilitySwitchEquipmentInfoList)
	}

	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CapabilitySwitchEquipmentInfoResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CapabilitySwitchEquipmentInfoList != nil {
		return obj.CapabilitySwitchEquipmentInfoList
	}

	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	// all schemas are nil
	return nil
}

type NullableCapabilitySwitchEquipmentInfoResponse struct {
	value *CapabilitySwitchEquipmentInfoResponse
	isSet bool
}

func (v NullableCapabilitySwitchEquipmentInfoResponse) Get() *CapabilitySwitchEquipmentInfoResponse {
	return v.value
}

func (v *NullableCapabilitySwitchEquipmentInfoResponse) Set(val *CapabilitySwitchEquipmentInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchEquipmentInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchEquipmentInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchEquipmentInfoResponse(val *CapabilitySwitchEquipmentInfoResponse) *NullableCapabilitySwitchEquipmentInfoResponse {
	return &NullableCapabilitySwitchEquipmentInfoResponse{value: val, isSet: true}
}

func (v NullableCapabilitySwitchEquipmentInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchEquipmentInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
