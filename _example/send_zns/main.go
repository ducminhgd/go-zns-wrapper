package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ducminhgd/zalo-go-sdk/client"
)

func main() {
	ctx := context.Background()
	zc := client.NewZaloClient(
		os.Getenv("ZALO_APP_ID"),
		os.Getenv("ZALO_SECRET_KEY"),
		os.Getenv("ZALO_CODE_VERIFIER"),
	)

	fmt.Printf("CODE CHALLENGE: %s\n", zc.GetCodeChallenge())

	zc.SetAccessToken(client.AccessToken{
		AccessToken:  os.Getenv("ZALO_ACCESS_TOKEN"),
		RefreshToken: os.Getenv("ZALO_REFRESH_TOKEN"),
		ExpiresIn:    90_000,
	})

	response, err := zc.SendZnsMessage(ctx, client.ZnsSendMsgRequest{
		Phone:      os.Getenv("ZNS_RECEIVER_PHONE"),
		TemplateID: os.Getenv("ZNS_TEMPLATE_ID"),
		TemplateData: map[string]string{
			"cost": "10,000",
			"note": "10,000 đồng",
		},
		TrackingID: "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("RESPONSE:\n%+v\n", response)
}
