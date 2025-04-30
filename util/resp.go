package util

import (
	"github.com/lee31802/comment_lib/ginerrors"
	"github.com/lee31802/comment_lib/gweb"
	"github.com/lee31802/gotemplate/errno"
	"net/http"
)

var nilData = NilData{}

type NilData struct {
}

func ErrTranferWithResponse(err error) gweb.Response {
	return WithError(err)
}

func WithCodeResp(resp interface{}, err ginerrors.Error) gweb.Response {
	return gweb.JSONResponse(http.StatusOK, err, resp)
}

func WithError(err error) gweb.Response {
	return gweb.JSONResponse(http.StatusOK, errno.GinError(err), nil)
}

func WithSuccess(resp interface{}) gweb.Response {
	return gweb.JSONResponse(http.StatusOK, ginerrors.Success, resp)
}
