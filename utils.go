package fn

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func computedBodySignature(key, content []byte) []byte {
	computed := hmac.New(sha1.New, key)
	computed.Write(content)
	return []byte(computed.Sum(nil))
}

func checkContentSignature(content []byte, signature string, body []byte) bool {

	const signaturePrefix = "sha1="
	const signatureLength = 45 // len(SignaturePrefix) + len(hex(sha1))

	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))

	return hmac.Equal(computedBodySignature(content, body), actual)
}
