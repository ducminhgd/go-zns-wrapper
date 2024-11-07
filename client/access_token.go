package client

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type AccessTokenRequest struct {
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}

// GetAccessToken exchanges an authorization code for an access token.
// It sends a POST request to the Zalo API using the provided context and request data.
// The request includes the authorization code, app ID, and code verifier.
// On success, it returns the access token, refresh token, and expiration time.
// If an error occurs during the request or response processing, it returns the error.
func (z *ZaloClient) GetAccessToken(ctx context.Context, request AccessTokenRequest) (AccessToken, error) {
	var response AccessToken

	// Set up the form data
	formData := url.Values{}
	formData.Set("code", request.Code)
	formData.Set("app_id", z.AppID)
	formData.Set("grant_type", "authorization_code")
	formData.Set("code_verifier", z.CodeVerifier)

	req, err := http.NewRequest("POST", ENDPOINT_GET_ACCESS_TOKEN, strings.NewReader(formData.Encode()))
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error creating request:", slog.Any("err", err))
		return response, err
	}
	// Set headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("secret_key", z.SecretKey)
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
	z.SetAccessToken(response)

	return response, nil
}

func (z *ZaloClient) RefreshAccessToken(ctx context.Context, request AccessTokenRequest) (AccessToken, error) {
	var response AccessToken

	// Set up the form data
	formData := url.Values{}
	formData.Set("refresh_token", request.RefreshToken)
	formData.Set("app_id", z.AppID)
	formData.Set("grant_type", "refresh_token")

	req, err := http.NewRequest("POST", ENDPOINT_GET_ACCESS_TOKEN, strings.NewReader(formData.Encode()))
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error creating request:", slog.Any("err", err))
		return response, err
	}
	// Set headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("secret_key", z.SecretKey)
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
	z.SetAccessToken(response)

	return response, nil
}
