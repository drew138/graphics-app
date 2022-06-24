package helpers

import (
	"net/url"
	"os"
	"strconv"
)

func FormatPaginatedURL(model string, from, to int) string {
	scheme := os.Getenv("SCHEME")
	host := os.Getenv("HOST")
	url := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   model,
	}
	query := url.Query()
	query.Set("from", strconv.Itoa(from))
	query.Set("to", strconv.Itoa(to))
	url.RawQuery = query.Encode()
	return url.String()
}
