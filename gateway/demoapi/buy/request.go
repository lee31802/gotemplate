package buy

import (
	"github.com/lee31802/comment_lib/gerrors"
	"github.com/lee31802/comment_lib/gweb"
	"github.com/lee31802/comment_lib/logkit"
)

type CreateOrderRequest struct {
	gweb.Request
	UserId    int64 `json:"user_id"`
	ProductId int64 `json:"product_id"`
}

func (req *CreateOrderRequest) Validate() gerrors.Error {
	if req.UserId <= 0 || req.ProductId <= 0 {
		logkit.Error("CreateOrderRequest req ParamsInvalid", logkit.Reflect("req", req))
		return gerrors.ErrorParamsInvalid
	}
	return nil
}
