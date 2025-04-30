package errno

import (
	"errors"
	"fmt"
	"github.com/lee31802/comment_lib/ginerrors"
	gmerrors "github.com/micro/go-micro/errors"
)

type ErrNo ginerrors.GinError

func (e *ErrNo) Error() string {
	return fmt.Sprintf("error: code = %d desc = %s", e.GetCode(), e.GetMsg())
}

func (e *ErrNo) GetCode() int32 {
	return e.Code
}

func (e *ErrNo) GetMsg() string {
	return e.Msg
}

func GinError(err error) ginerrors.Error {
	// 处理业务错误
	var e ginerrors.Error
	ok := errors.As(err, &e)
	if ok {
		return ginerrors.New(e.GetCode(), e.GetMsg())
	}
	// 处理一些框架的错误
	e = As(err)
	return ginerrors.New(e.GetCode(), e.GetMsg())
}

func As(err error) ginerrors.Error {
	if err == nil {
		return ginerrors.Success
	}

	// 处理框架错误
	var e *gmerrors.Error
	if errors.As(err, &e) {
		return &ErrNo{
			Code: e.Code,
			Msg:  e.Error(),
		}
	}

	return ginerrors.ErrorUnKnown
}
