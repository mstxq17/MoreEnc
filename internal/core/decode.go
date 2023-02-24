package core

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
)

func UrlDecode(value string) string {
	decodeOutput, err := url.QueryUnescape(value)
	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, "Error decoding input", err)
		os.Exit(1)
	}
	return decodeOutput
}

func B64Decode(value string) string {
	decodeOutput, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, "Error base64 decoding input")
	}
	return string(decodeOutput)
}
