package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ducminhgd/zalo-go-sdk/client"
)

func main() {
	zc := client.NewZaloClient(
		os.Getenv("ZALO_APP_ID"),
		os.Getenv("ZALO_SECRET_KEY"),
		os.Getenv("ZALO_CODE_VERIFIER"),
	)
	token, err := zc.RequestAccessToken(context.Background(), client.AccessTokenRequest{
		Code: os.Getenv("ZNS_CODE"),
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
