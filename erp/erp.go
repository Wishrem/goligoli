package erp

import (
	"encoding/json"
	"errors"

	"google.golang.org/grpc/status"
)

type ErrResp struct {
	c int64
	m string
}

func _new(code int64) *ErrResp {
	msg := getMessage(code)
	if msg == "Unknown" {
		return &ErrResp{c: INTERNAL_ERROR, m: msg}
	}
	return &ErrResp{c: code, m: msg}
}

func New(code int64, msg string) *ErrResp {
	if getMessage(code) == "Unknown" {
		code = INTERNAL_ERROR
	}
	return &ErrResp{c: code, m: msg}
}

func Covert(err error) *ErrResp {
	erp := new(ErrResp)
	if ok := errors.As(err, &erp); ok {
		return erp
	}

	st := status.Convert(err)
	erp.c = INTERNAL_ERROR
	erp.m = st.Message()
	return erp
}

/*
*

	json
*/
type resp struct {
	Detail *detail `json:"error"`
}

type detail struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func (erp *ErrResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(&resp{Detail: &detail{Code: erp.c, Msg: erp.m}})
}

func (erp *ErrResp) UnmarshalJSON(data []byte) error {
	return errors.New("ErrResp's unmarshalJSON func doesn't implement")
}

/*
*

	error
*/
func (erp *ErrResp) Error() string {
	return erp.m
}
