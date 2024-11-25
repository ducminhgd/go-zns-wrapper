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

	// CODE VERIFIER: ThisIsCodeVerifierToCreateCodeChallenge
	// CODE CHALLENGE: UMmgzWBy-6hAVyG8y-UD1tufv8rxiZP9AXJXrb8blYg
	fmt.Printf("Code Challenge: %+v\n", zc.GetCodeChallenge())

	token, err := zc.RequestAccessToken(context.Background(), client.AccessTokenRequest{
		Code: os.Getenv("ZNS_CODE"),
	})
	if err != nil {
		panic(err)
	}
	zc.SetAccessToken(token)
	fmt.Printf("First Access Token: %+v\n", token)

	token, err = zc.RefreshAccessToken(context.Background(), client.AccessTokenRequest{
		RefreshToken: token.RefreshToken,
	})
	if err != nil {
		panic(err)
	}
	zc.SetAccessToken(token)
	fmt.Printf("After Refresh Token: %+v\n", token)
}
