package utils

import (
	"net"
	"net/http"
)

func ExtractIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	return host
}
