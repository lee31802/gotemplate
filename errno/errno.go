package errno

import (
	"errors"
	"fmt"
	"github.com/lee31802/comment_lib/gerrors"
)

type ErrNo gerrors.GinError

func (e *ErrNo) Error() string {
	return fmt.Sprintf("error: code = %d desc = %s", e.GetCode(), e.GetMsg())
}

func (e *ErrNo) GetCode() int32 {
	return e.Code
}

func (e *ErrNo) GetMsg() string {
	return e.Msg
}

func GinError(err error) gerrors.Error {
	// 处理业务错误
	var e gerrors.Error
	ok := errors.As(err, &e)
	if ok {
		return gerrors.New(e.GetCode(), e.GetMsg())
	}
	// 处理一些框架的错误
	e = As(err)
	return gerrors.New(e.GetCode(), e.GetMsg())
}

func As(err error) gerrors.Error {
	if err == nil {
		return gerrors.Success
	}
	//
	//// 处理框架错误
	//var e *gmerrors.Error
	//if errors.As(err, &e) {
	//	return &ErrNo{
	//		Code: e.Code,
	//		Msg:  e.Error(),
	//	}
	//}

	return gerrors.ErrorUnKnown
}
