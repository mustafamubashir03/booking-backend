package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseURL string, pathPrefix string) http.HandlerFunc {
	targetURL, err := url.Parse(targetBaseURL)

	fmt.Println("Target URL:", targetBaseURL)
	fmt.Println("Path Prefix:", pathPrefix)

	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			WriteJsonErrorResponse(
				w,
				http.StatusInternalServerError,
				"Failed to parse target URL",
				err,
			)
		}
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(req *httputil.ProxyRequest) {

			// Save the original incoming path
			incomingPath := req.In.URL.Path

			// Set target URL
			req.SetURL(targetURL)

			// Remove the gateway prefix from the incoming path
			remainingPath := strings.TrimPrefix(incomingPath, pathPrefix)

			// If only "/" remains, use target's existing path
			if remainingPath == "" || remainingPath == "/" {
				req.Out.URL.Path = targetURL.Path
			} else {
				req.Out.URL.Path = strings.TrimRight(targetURL.Path, "/") + remainingPath
			}

			req.SetXForwarded()
			req.Out.Host = targetURL.Host

			if userId, ok := req.In.Context().Value("userId").(string); ok {
				req.Out.Header.Set("X-User-ID", userId)
			}
		},
	}

	return proxy.ServeHTTP
}
