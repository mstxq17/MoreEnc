package core

import (
	"encoding/base64"
	"net/url"
)

func UrlEncode(value string) string {
	encodeOutput := url.QueryEscape(value)
	return encodeOutput
}

func B64Encode(value string) string {
	encodeOutput := base64.StdEncoding.EncodeToString([]byte(value))
	return encodeOutput
}
