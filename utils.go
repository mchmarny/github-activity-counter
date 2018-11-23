package counter

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strings"
)

func getFileContent(path string) (conten []byte, err error) {

	jf, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jf.Close()
	data, _ := ioutil.ReadAll(jf)
	return data, nil
}

func makeNewSignature(key, content []byte) string {
	dst := make([]byte, 40)
	computed := hmac.New(sha1.New, key)
	computed.Write(content)
	hex.Encode(dst, computed.Sum(nil))
	return "sha1=" + string(dst)
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
