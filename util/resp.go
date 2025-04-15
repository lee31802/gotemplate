package util

import (
	"github.com/lee31802/comment_lib/ginerrors"
	"github.com/lee31802/comment_lib/ginserver"
	"github.com/lee31802/gotemplate/errno"
	"net/http"
)

var nilData = NilData{}

type NilData struct {
}

func ErrTranferWithResponse(err error) ginserver.Response {
	return WithError(err)
}

func WithCodeResp(resp interface{}, err ginerrors.Error) ginserver.Response {
	return ginserver.JSONResponse(http.StatusOK, err, resp)
}

func WithError(err error) ginserver.Response {
	return ginserver.JSONResponse(http.StatusOK, errno.GinError(err), nil)
}

func WithSuccess(resp interface{}) ginserver.Response {
	return ginserver.JSONResponse(http.StatusOK, ginerrors.Success, resp)
}
