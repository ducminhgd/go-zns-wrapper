package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
)

const (
	ZNS_TPL_STATUS_UNKNOWN_CODE        = 0 // ZNS not defined this
	ZNS_TPL_STATUS_ENABLED_CODE        = 1
	ZNS_TPL_STATUS_PENDING_REVIEW_CODE = 2
	ZNS_TPL_STATUS_REJECTED_CODE       = 3
	ZNS_TPL_STATUS_DISABLED_CODE       = 4

	ZNS_TPL_STATUS_UNKNOWN_NAME        = "UNKNOWN" // ZNS not defined this
	ZNS_TPL_STATUS_ENABLED_NAME        = "ENABLE"
	ZNS_TPL_STATUS_PENDING_REVIEW_NAME = "PENDING_REVIEW"
	ZNS_TPL_STATUS_REJECTED_NAME       = "REJECT"
	ZNS_TPL_STATUS_DISABLED_NAME       = "ENABLE"
)

type ZnsTplListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Status int `json:"status"`
}

type ZnsTplListRecord struct {
	TemplateID      int    `json:"templateId"`
	TemplateName    int    `json:"templateName"`
	CreatedTime     int    `json:"createdTime"`
	Status          string `json:"status"`
	TemplateQuality string `json:"templateQuality"`
}

type ZnsTplListMetadata struct {
	Total int `json:"total"`
}

type ZnsTplListResponse struct {
	Error    int                `json:"error"`
	Message  string             `json:"message"`
	Data     []ZnsTplListRecord `json:"data"`
	Metadata ZnsTplListMetadata `json:"metadata"`
}

// GetZnsTemplateList gets a list of ZNS templates.
//
// It sends a GET request to the Zalo API using the provided context and request data.
// The request includes the offset, limit, and status as query string parameters.
// On success, it returns the list response.
// If an error occurs during the request or response processing, it returns the error.
func (z *ZaloClient) GetZnsTemplateList(ctx context.Context, request ZnsTplListRequest) (ZnsTplListResponse, error) {
	var response ZnsTplListResponse

	// Set up the query string parameters
	query := url.Values{}
	if request.Offset < 0 {
		request.Offset = 0
	}
	query.Set("offset", strconv.Itoa(request.Offset))

	if request.Limit < 0 || request.Limit > 100 {
		request.Limit = 100
	}
	query.Set("limit", strconv.Itoa(request.Limit))

	if request.Status != ZNS_TPL_STATUS_UNKNOWN_CODE {
		query.Set("status", strconv.Itoa(request.Status))
	}

	// Create the request URL with the query string parameters
	reqUrl := fmt.Sprintf("%s?%s", ENDPOINT_TEMPLATE_LIST, query.Encode())

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error creating request:", slog.Any("err", err))
		return response, err
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", z.token.AccessToken)
	resp, err := z.GetHTTPClient().Do(req)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error sending request:", slog.Any("err", err))
		return response, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error reading response:", slog.Any("err", err))
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error unmarshalling response:", slog.Any("err", err))
		return response, err
	}

	return response, nil
}

type ZnsTplDetailParam struct {
	Name       string `json:"name"`
	Require    bool   `json:"require"`
	Type       string `json:"type"`
	MaxLength  int    `json:"maxLength"`
	MinLength  int    `json:"minLength"`
	AcceptNull bool   `json:"acceptNull"`
}

type ZnsTplDetailData struct {
	TemplateID      int                 `json:"templateId"`
	TemplateName    string              `json:"templateName"`
	Status          string              `json:"status"`
	ListParams      []ZnsTplDetailParam `json:"listParams"`
	Timeout         int                 `json:"timeout"`
	PreviewURL      string              `json:"previewUrl"`
	TemplateQuality string              `json:"templateQuality"`
	TemplateTag     string              `json:"templateTag"`
	Price           string              `json:"price"`
}

type ZnsTplDetailResponse struct {
	Error   int              `json:"error"`
	Message string           `json:"message"`
	Data    ZnsTplDetailData `json:"data"`
}

// GetZnsTemplateDetail gets a template detail by template ID.
//
// It sends a GET request to the Zalo API using the provided context and template ID.
// The request includes the template ID as a query string parameter.
// On success, it returns the template detail response.
// If an error occurs during the request or response processing, it returns the error.
func (z *ZaloClient) GetZnsTemplateDetail(ctx context.Context, templateID int) (ZnsTplDetailResponse, error) {
	var response ZnsTplDetailResponse

	// Set up the query string parameters
	query := url.Values{}
	query.Set("template_id", strconv.Itoa(templateID))

	// Create the request URL with the query string parameters
	reqUrl := fmt.Sprintf("%s?%s", ENDPOINT_TEMPLATE_DETAIL, query.Encode())

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error creating request:", slog.Any("err", err))
		return response, err
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", z.token.AccessToken)
	resp, err := z.GetHTTPClient().Do(req)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error sending request:", slog.Any("err", err))
		return response, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error reading response:", slog.Any("err", err))
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error unmarshalling response:", slog.Any("err", err))
		return response, err
	}

	return response, nil
}
