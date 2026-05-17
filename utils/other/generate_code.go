package other

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

const codeLength = 6

// GenerateShortCode creates a random short code if customCode is empty
// If customCode is provided, it validates and returns it (after cleaning)
func GenerateShortCode(customCode string) (string, error) {
	// If user provided a custom code
	if strings.TrimSpace(customCode) != "" {
		// Remove spaces and convert to lowercase
		cleanCode := strings.ToLower(strings.TrimSpace(customCode))

		return cleanCode, nil
	}

	// Generate random code
	return generateRandomCode()
}

func generateRandomCode() (string, error) {
	bytes := make([]byte, codeLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	// Base64 encoding gives alphanumeric + some symbols
	// Replace symbols with letters and take first codeLength chars
	code := base64.URLEncoding.EncodeToString(bytes)
	code = strings.ReplaceAll(code, "-", "a")
	code = strings.ReplaceAll(code, "_", "b")
	return strings.ToLower(code[:codeLength]), nil
}
