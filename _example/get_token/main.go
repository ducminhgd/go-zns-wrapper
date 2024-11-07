package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ducminhgd/zalo-go-sdk/client"
)

func main() {
	appID := os.Getenv("ZALO_APP_ID")
	secretKey := os.Getenv("ZALO_SECRET_KEY")
	codeVerifier := os.Getenv("ZALO_CODE_VERIFIER")
	znsCode := os.Getenv("ZNS_CODE")

	zc := client.NewZaloClient(appID, secretKey, codeVerifier)
	token, err := zc.RequestAccessToken(context.Background(), client.AccessTokenRequest{
		Code: znsCode,
	})
	if err != nil {
		panic(err)
	}
	zc.SetAccessToken(token)
	fmt.Println(token)

	token, err = zc.RefreshAccessToken(context.Background(), client.AccessTokenRequest{
		RefreshToken: token.RefreshToken,
	})
	if err != nil {
		panic(err)
	}
	zc.SetAccessToken(token)
	fmt.Println(token)
}
