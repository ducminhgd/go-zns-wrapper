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

	response, err := zc.GetZnsTemplateDetail(ctx, os.Getenv("ZNS_TEMPLATE_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("RESPONSE:\n%+v\n", response)
}
