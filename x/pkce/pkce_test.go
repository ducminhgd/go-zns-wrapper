package pkce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCodeChallenge(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU"}, // this is empty string
		{"single character", "a", "ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs"},
		{"multi-character string", "hello", "LPJNul-wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ"},
		{"non-ASCII characters", "héllo", "PEhZHY0JikU49eAT389AbpSOrE0yd7EL9hTildYGgXk"},
		{"long string", "abcdefghijklmnopqrstuvwxyz", "ccSA35PWri8e-tFEfGbJUl4xYhjPUfyNntgy8trxi3M"},
	}

	for _, tt := range tests {
		actual := GetCodeChallenge(tt.input)
		assert.Equal(t, tt.expected, actual, "GetCodeChallenge(%q) = %q, want %q", tt.input, actual, tt.expected)
	}
}
