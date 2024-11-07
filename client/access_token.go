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
	"strings"
)

type AccessToken struct {
	AccessToken  string `json:"access_token" mapstructure:"access_token"`
	RefreshToken string `json:"refresh_token" mapstructure:"refresh_token"`
	ExpiresIn    int    `json:"expires_in" mapstructure:"expires_in"`
}

type AccessTokenRequest struct {
	Code         string `json:"code" mapstructure:"code"`
	RefreshToken string `json:"refresh_token" mapstructure:"refresh_token"`
}

type ErrorResp struct {
	Error            int    `json:"error" mapstructure:"error"`
	ErrorDescription string `json:"error_description" mapstructure:"error_description"`
}

// RequestAccessToken exchanges an authorization code for an access token.
// It sends a POST request to the Zalo API using the provided context and request data.
// The request includes the authorization code, app ID, and code verifier.
// On success, it returns the access token, refresh token, and expiration time.
// If an error occurs during the request or response processing, it returns the error.
func (z *ZaloClient) RequestAccessToken(ctx context.Context, request AccessTokenRequest) (AccessToken, error) {
	var token AccessToken

	// Set up the form data
	formData := url.Values{}
	formData.Set("code", request.Code)
	formData.Set("app_id", z.appID)
	formData.Set("grant_type", "authorization_code")
	formData.Set("code_verifier", z.codeVerifier)

	req, err := http.NewRequest("POST", ENDPOINT_GET_ACCESS_TOKEN, strings.NewReader(formData.Encode()))
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error creating request:", slog.Any("err", err))
		return token, err
	}
	// Set headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("secret_key", z.secretKey)
	resp, err := z.GetHTTPClient().Do(req)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error sending request:", slog.Any("err", err))
		return token, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error reading response:", slog.Any("err", err))
		return token, err
	}

	var respBody map[string]string
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		var errResp ErrorResp
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			z.GetLogger().ErrorContext(ctx, "Error unmarshalling response:", slog.Any("err", err))
			return token, err
		}
		err = fmt.Errorf("code: %d, description: %s", errResp.Error, errResp.ErrorDescription)
		z.GetLogger().ErrorContext(ctx, "Error:", slog.Any("err", err))
		return token, err
	}

	token.AccessToken = respBody["access_token"]
	token.RefreshToken = respBody["refresh_token"]
	token.ExpiresIn, _ = strconv.Atoi(respBody["expires_in"])

	return token, nil
}

func (z *ZaloClient) RefreshAccessToken(ctx context.Context, request AccessTokenRequest) (AccessToken, error) {
	var token AccessToken

	// Set up the form data
	formData := url.Values{}
	formData.Set("refresh_token", request.RefreshToken)
	formData.Set("app_id", z.appID)
	formData.Set("grant_type", "refresh_token")

	req, err := http.NewRequest("POST", ENDPOINT_GET_ACCESS_TOKEN, strings.NewReader(formData.Encode()))
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error creating request:", slog.Any("err", err))
		return token, err
	}
	// Set headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("secret_key", z.secretKey)
	resp, err := z.GetHTTPClient().Do(req)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error sending request:", slog.Any("err", err))
		return token, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error reading response:", slog.Any("err", err))
		return token, err
	}

	var respBody map[string]string
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		var errResp ErrorResp
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			z.GetLogger().ErrorContext(ctx, "Error unmarshalling response:", slog.Any("err", err))
			return token, err
		}
		err = fmt.Errorf("code: %d, description: %s", errResp.Error, errResp.ErrorDescription)
		z.GetLogger().ErrorContext(ctx, "Error:", slog.Any("err", err))
		return token, err
	}

	token.AccessToken = respBody["access_token"]
	token.RefreshToken = respBody["refresh_token"]
	token.ExpiresIn, _ = strconv.Atoi(respBody["expires_in"])

	return token, nil
}
