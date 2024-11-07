package client

import (
	"log/slog"
	"net/http"

	"github.com/ducminhgd/zalo-go-sdk/x/pkce"
)

type ZaloClient struct {
	httpClient *http.Client
	logger     *slog.Logger

	AppID         string `json:"app_id" mapstructure:"app_id"`
	SecretKey     string `json:"secret_key" mapstructure:"secret_key"`
	CodeVerifier  string `json:"code_verifier" mapstructure:"code_verifier"`
	CodeChallenge string `json:"code_challenge" mapstructure:"code_challenge"`

	token AccessToken `json:"token" mapstructure:"token"`
}

func NewZaloClient(appID, secretKey, codeVerifier string) *ZaloClient {
	return &ZaloClient{
		AppID:         appID,
		SecretKey:     secretKey,
		CodeVerifier:  codeVerifier,
		CodeChallenge: pkce.GetCodeChallenge(codeVerifier),
	}
}

func (z *ZaloClient) UseHTTPClient(client *http.Client) {
	z.httpClient = client
}

func (z *ZaloClient) GetHTTPClient() *http.Client {
	if z.httpClient == nil {
		z.httpClient = &http.Client{}
	}
	return z.httpClient
}

func (z *ZaloClient) UseLogger(logger *slog.Logger) {
	z.logger = logger
}

func (z *ZaloClient) GetLogger() *slog.Logger {
	if z.logger == nil {
		z.logger = slog.Default()
	}
	return z.logger
}

func (z *ZaloClient) SetAccessToken(token AccessToken) {
	z.token = token
}
