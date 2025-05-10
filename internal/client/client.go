package client

import "net/http"

var RetryStatus = []int{
	// 4xx
	http.StatusBadRequest,
	http.StatusForbidden,
	http.StatusRequestTimeout,
	http.StatusConflict,
	http.StatusTooManyRequests,

	// 5xx
	http.StatusInternalServerError,
	http.StatusNotImplemented,
	http.StatusBadGateway,
	http.StatusServiceUnavailable,
	http.StatusGatewayTimeout,
	http.StatusHTTPVersionNotSupported,
	http.StatusVariantAlsoNegotiates,
	http.StatusInsufficientStorage,
	http.StatusLoopDetected,
	http.StatusNotExtended,
	http.StatusNetworkAuthenticationRequired,
}