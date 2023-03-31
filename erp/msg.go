package erp

func getMessage(code int64) string {
	switch code {
	case INTERNAL_ERROR:
		return "Internal Error"
	case BAD_REQUEST:
		return "Params Loss"
	case VIDEO_NOT_FOUND:
		return "Video Not Found"
	case UNAUTHORIZED:
		return "Failed to Authorize"
	case FORBIDDEN:
		return "Forbidden"
	default:
		return "Unknown"
	}
}
