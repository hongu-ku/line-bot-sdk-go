/**
 * Mission Stickers API
 * This document describes LINE Mission Stickers API.
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

package shop

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

type ShopAPI struct {
	httpClient   *http.Client
	endpoint     *url.URL
	channelToken string
	ctx          context.Context
}

// ShopAPIOption type
type ShopAPIOption func(*ShopAPI) error

// New returns a new bot client instance.
func NewShopAPI(channelToken string, options ...ShopAPIOption) (*ShopAPI, error) {
	if channelToken == "" {
		return nil, errors.New("missing channel access token")
	}

	c := &ShopAPI{
		channelToken: channelToken,
		httpClient:   http.DefaultClient,
	}

	u, err := url.ParseRequestURI("https://api.line.me")
	if err != nil {
		return nil, err
	}
	c.endpoint = u

	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// WithContext method
func (call *ShopAPI) WithContext(ctx context.Context) *ShopAPI {
	call.ctx = ctx
	return call
}

func (client *ShopAPI) Do(req *http.Request) (*http.Response, error) {
	if client.channelToken != "" {
		req.Header.Set("Authorization", "Bearer "+client.channelToken)
	}
	req.Header.Set("User-Agent", "LINE-BotSDK-Go/"+linebot.GetVersion())
	if client.ctx != nil {
		req = req.WithContext(client.ctx)
	}
	return client.httpClient.Do(req)
}

func (client *ShopAPI) Url(endpointPath string) string {
	u := client.endpoint
	u.Path = path.Join(u.Path, endpointPath)
	return u.String()
}

// WithHTTPClient function
func WithHTTPClient(c *http.Client) ShopAPIOption {
	return func(client *ShopAPI) error {
		client.httpClient = c
		return nil
	}
}

// WithEndpointClient function
func WithEndpoint(endpoint string) ShopAPIOption {
	return func(client *ShopAPI) error {
		u, err := url.ParseRequestURI(endpoint)
		if err != nil {
			return err
		}
		client.endpoint = u
		return nil
	}
}

// MissionStickerV3
//
// Sends a mission sticker.
// Parameters:
//        missionStickerRequest

// https://developers.line.biz/en/reference/partner-docs/#send-mission-stickers-v3
func (client *ShopAPI) MissionStickerV3(

	missionStickerRequest *MissionStickerRequest,

) (struct{}, error) {
	_, body, error := client.MissionStickerV3WithHttpInfo(

		missionStickerRequest,
	)
	return body, error
}

// MissionStickerV3
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Sends a mission sticker.
// Parameters:
//        missionStickerRequest

// https://developers.line.biz/en/reference/partner-docs/#send-mission-stickers-v3
func (client *ShopAPI) MissionStickerV3WithHttpInfo(

	missionStickerRequest *MissionStickerRequest,

) (*http.Response, struct{}, error) {
	path := "/shop/v3/mission"

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(missionStickerRequest); err != nil {
		return nil, struct{}{}, err
	}
	req, err := http.NewRequest(http.MethodPost, client.Url(path), &buf)
	if err != nil {
		return nil, struct{}{}, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := client.Do(req)

	if err != nil {
		return res, struct{}{}, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, struct{}{}, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, struct{}{}, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	return res, struct{}{}, nil

}
