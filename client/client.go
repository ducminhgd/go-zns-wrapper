package client

import (
	"log/slog"
	"net/http"

	"github.com/ducminhgd/zalo-go-sdk/x/pkce"
)

type ZaloClient struct {
	httpClient    *http.Client
	logger        *slog.Logger
	token         AccessToken
	appID         string
	secretKey     string
	codeVerifier  string
	codeChallenge string
}

func NewZaloClient(appID, secretKey, codeVerifier string) *ZaloClient {
	return &ZaloClient{
		appID:         appID,
		secretKey:     secretKey,
		codeVerifier:  codeVerifier,
		codeChallenge: pkce.GetCodeChallenge(codeVerifier),
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

func (z *ZaloClient) GetAccessToken() AccessToken {
	return z.token
}
