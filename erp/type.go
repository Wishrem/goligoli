package erp

var (
	Internal = _new(INTERNAL_ERROR)

	BadRequest   = _new(BAD_REQUEST)
	Unauthorized = _new(UNAUTHORIZED)
	Forbidden    = _new(FORBIDDEN)

	VideoNotFound = _new(VIDEO_NOT_FOUND)
)
