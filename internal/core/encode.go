package core

import (
	"net/url"
)

func UrlEncode(value string) string {
	encodeOutput := url.QueryEscape(value)
	return encodeOutput
}
