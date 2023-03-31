package handler

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

func toHttpCode(c codes.Code) int {
	switch c {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusForbidden
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusRequestTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusAlreadyReported
	case codes.PermissionDenied:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusNotFound
	case codes.FailedPrecondition:
		return http.StatusInternalServerError
	case codes.Aborted:
		return http.StatusForbidden
	case codes.OutOfRange:
		return http.StatusInternalServerError
	case codes.Unimplemented:
		return http.StatusServiceUnavailable
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusNotFound
	case codes.Unauthenticated:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
