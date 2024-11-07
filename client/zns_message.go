package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type ZnsSendMsgRequest struct {
	Phone        string            `json:"phone"`
	TemplateID   string            `json:"template_id"`
	TemplateData map[string]string `json:"template_data"` // This base on each template
	TrackingID   string            `json:"tracking_id"`
}

type ZnsSendMsgResponseData struct {
	MsgID       string    `json:"msg_id"`
	SentTime    time.Time `json:"sent_time"`
	SendingMode string    `json:"sending_mode"`
	Quota       struct {
		DailyQuota     string `json:"dailyQuota"`
		RemainingQuota string `json:"remainingQuota"`
	} `json:"quota"`
}

type ZnsSendMsgReponse struct {
	Error   int                    `json:"error"`
	Message string                 `json:"message"`
	Data    ZnsSendMsgResponseData `json:"data"`
}

func (z *ZaloClient) SendZnsMessage(ctx context.Context, request ZnsSendMsgRequest) (ZnsSendMsgReponse, error) {
	var response ZnsSendMsgReponse

	// Set up the request body as a JSON object
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		z.GetLogger().ErrorContext(ctx, "Error encoding request as JSON:", slog.Any("err", err))
		return response, err
	}

	req, err := http.NewRequest("POST", ENDPOINT_MESSAGE_SEND, bytes.NewReader(jsonBytes))
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
		z.GetLogger().ErrorContext(ctx, "Error unmarshaling response:", slog.Any("err", err))
		return response, err
	}
	return response, nil
}
