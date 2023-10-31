/**
 * Channel Access Token API
 * This document describes Channel Access Token API.
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
package channel_access_token

// ChannelAccessTokenKeyIdsResponse
// Channel access token key IDs
// https://developers.line.biz/en/reference/messaging-api/#get-all-valid-channel-access-token-key-ids-v2-1
type ChannelAccessTokenKeyIdsResponse struct {

	/**
	 * Array of channel access token key IDs. (Required)
	 */
	Kids []string `json:"kids"`
}