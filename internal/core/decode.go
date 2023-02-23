package core

import (
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
