package erp

func (erp *ErrResp) HttpCode() int {
	switch erp.c {
	case VIDEO_NOT_FOUND:
		return 404
	case INTERNAL_ERROR:
		return 500
	case BAD_REQUEST:
		return 400
	case UNAUTHORIZED:
		return 401
	default:
		return 500
	}
}
