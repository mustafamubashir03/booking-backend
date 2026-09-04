package utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseURL string, pathPrefix string) http.HandlerFunc {
	targetURL, err := url.Parse(targetBaseURL)
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
			incomingPath := req.In.URL.Path
			req.SetURL(targetURL)
			remainingPath := strings.TrimPrefix(incomingPath, pathPrefix)
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
