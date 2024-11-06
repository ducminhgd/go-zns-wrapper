// Package pkce implements the PKCE (Proof Key for Code Exchange) protocol
// according to RFC-7636. https://datatracker.ietf.org/doc/html/rfc7636
//
// The main function is GetCodeChallenge, which generates a code challenge
// based on the given code verifier. The code challenge is the SHA256 hash
// of the code verifier, base64 encoded with URL-safe characters.
package pkce

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

// GetCodeChallenge generates a code challenge based on the given code verifier.
// The code challenge is the SHA256 hash of the code verifier, base64 encoded
// without padding. See https://tools.ietf.org/html/rfc7636#section-4.2
func GetCodeChallenge(codeVerifier string) string {
	asciiBytes := []byte(codeVerifier)
	hash := sha256.Sum256(asciiBytes)
	base64Encoded := Base64UrlEncode(hash)
	return base64Encoded
}

// Base64UrlEncode encodes a 256-bit hash as a URL-safe base64-encoded string,
// according to the rules in https://tools.ietf.org/html/rfc4648#section-5
// (with padding removed, and "+" and "/" characters replaced with "-" and "_"
// respectively).
func Base64UrlEncode(hash [32]byte) string {
	result := base64.RawURLEncoding.EncodeToString(hash[:])
	result = strings.TrimRight(result, "=")
	result = strings.ReplaceAll(result, "+", "-")
	result = strings.ReplaceAll(result, "/", "_")
	return result
}
