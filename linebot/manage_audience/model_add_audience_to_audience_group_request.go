/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package manage_audience

// AddAudienceToAudienceGroupRequest
// Add user IDs or Identifiers for Advertisers (IFAs) to an audience for uploading user IDs (by JSON)
// https://developers.line.biz/en/reference/messaging-api/#update-upload-audience-group
type AddAudienceToAudienceGroupRequest struct {

	/**
	 * The audience ID.
	 */
	AudienceGroupId int64 `json:"audienceGroupId"`

	/**
	 * The audience&#39;s name.
	 */
	UploadDescription string `json:"uploadDescription,omitempty"`

	/**
	 * An array of up to 10,000 user IDs or IFAs.
	 */
	Audiences []Audience `json:"audiences,omitempty"`
}