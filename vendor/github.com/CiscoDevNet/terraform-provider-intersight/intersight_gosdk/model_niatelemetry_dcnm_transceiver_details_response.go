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

// NiatelemetryDcnmTransceiverDetailsResponse - The response body of a HTTP GET request for the 'niatelemetry.DcnmTransceiverDetails' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niatelemetry.DcnmTransceiverDetails' resources.
type NiatelemetryDcnmTransceiverDetailsResponse struct {
	MoAggregateTransform                   *MoAggregateTransform
	MoDocumentCount                        *MoDocumentCount
	MoTagSummary                           *MoTagSummary
	NiatelemetryDcnmTransceiverDetailsList *NiatelemetryDcnmTransceiverDetailsList
}

// MoAggregateTransformAsNiatelemetryDcnmTransceiverDetailsResponse is a convenience function that returns MoAggregateTransform wrapped in NiatelemetryDcnmTransceiverDetailsResponse
func MoAggregateTransformAsNiatelemetryDcnmTransceiverDetailsResponse(v *MoAggregateTransform) NiatelemetryDcnmTransceiverDetailsResponse {
	return NiatelemetryDcnmTransceiverDetailsResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiatelemetryDcnmTransceiverDetailsResponse is a convenience function that returns MoDocumentCount wrapped in NiatelemetryDcnmTransceiverDetailsResponse
func MoDocumentCountAsNiatelemetryDcnmTransceiverDetailsResponse(v *MoDocumentCount) NiatelemetryDcnmTransceiverDetailsResponse {
	return NiatelemetryDcnmTransceiverDetailsResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiatelemetryDcnmTransceiverDetailsResponse is a convenience function that returns MoTagSummary wrapped in NiatelemetryDcnmTransceiverDetailsResponse
func MoTagSummaryAsNiatelemetryDcnmTransceiverDetailsResponse(v *MoTagSummary) NiatelemetryDcnmTransceiverDetailsResponse {
	return NiatelemetryDcnmTransceiverDetailsResponse{
		MoTagSummary: v,
	}
}

// NiatelemetryDcnmTransceiverDetailsListAsNiatelemetryDcnmTransceiverDetailsResponse is a convenience function that returns NiatelemetryDcnmTransceiverDetailsList wrapped in NiatelemetryDcnmTransceiverDetailsResponse
func NiatelemetryDcnmTransceiverDetailsListAsNiatelemetryDcnmTransceiverDetailsResponse(v *NiatelemetryDcnmTransceiverDetailsList) NiatelemetryDcnmTransceiverDetailsResponse {
	return NiatelemetryDcnmTransceiverDetailsResponse{
		NiatelemetryDcnmTransceiverDetailsList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryDcnmTransceiverDetailsResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("Failed to unmarshal NiatelemetryDcnmTransceiverDetailsResponse as MoAggregateTransform: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal NiatelemetryDcnmTransceiverDetailsResponse as MoDocumentCount: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal NiatelemetryDcnmTransceiverDetailsResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.DcnmTransceiverDetails.List'
	if jsonDict["ObjectType"] == "niatelemetry.DcnmTransceiverDetails.List" {
		// try to unmarshal JSON data into NiatelemetryDcnmTransceiverDetailsList
		err = json.Unmarshal(data, &dst.NiatelemetryDcnmTransceiverDetailsList)
		if err == nil {
			return nil // data stored in dst.NiatelemetryDcnmTransceiverDetailsList, return on the first match
		} else {
			dst.NiatelemetryDcnmTransceiverDetailsList = nil
			return fmt.Errorf("Failed to unmarshal NiatelemetryDcnmTransceiverDetailsResponse as NiatelemetryDcnmTransceiverDetailsList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryDcnmTransceiverDetailsResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiatelemetryDcnmTransceiverDetailsList != nil {
		return json.Marshal(&src.NiatelemetryDcnmTransceiverDetailsList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryDcnmTransceiverDetailsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
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

	if obj.NiatelemetryDcnmTransceiverDetailsList != nil {
		return obj.NiatelemetryDcnmTransceiverDetailsList
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryDcnmTransceiverDetailsResponse struct {
	value *NiatelemetryDcnmTransceiverDetailsResponse
	isSet bool
}

func (v NullableNiatelemetryDcnmTransceiverDetailsResponse) Get() *NiatelemetryDcnmTransceiverDetailsResponse {
	return v.value
}

func (v *NullableNiatelemetryDcnmTransceiverDetailsResponse) Set(val *NiatelemetryDcnmTransceiverDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryDcnmTransceiverDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryDcnmTransceiverDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryDcnmTransceiverDetailsResponse(val *NiatelemetryDcnmTransceiverDetailsResponse) *NullableNiatelemetryDcnmTransceiverDetailsResponse {
	return &NullableNiatelemetryDcnmTransceiverDetailsResponse{value: val, isSet: true}
}

func (v NullableNiatelemetryDcnmTransceiverDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryDcnmTransceiverDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
