package buy

import (
	"context"
	"github.com/lee31802/comment_lib/gweb"
	"github.com/lee31802/comment_lib/logkit"
	"github.com/lee31802/gotemplate/gateway/demoapi/client"
	"github.com/lee31802/gotemplate/pb/buy"
	"github.com/lee31802/gotemplate/util"
	"google.golang.org/protobuf/proto"
)

type BuyModule struct {
	client buy.Client
}

func NewDemoModule() *BuyModule {
	return &BuyModule{}
}

func (m *BuyModule) Init(r gweb.Router) {
	m.client = client.BuyClient()
	group := r.Group("api/buyer/order")
	{
		group.POST("/create", m.CreateOrder)
	}
}

func (m *BuyModule) CreateOrder(ctx context.Context, req *CreateOrderRequest) gweb.Response {
	rsp, err := m.client.CreateOrder(ctx, &buy.CreateOrderRequest{
		UserId:    proto.Int64(req.UserId),
		ProductId: proto.Int64(req.ProductId),
	})
	if err != nil {
		logkit.FromContext(ctx).Error("CreateOrder API", logkit.Any("req", req), logkit.Err(err))
		return util.WithError(err)
	}
	resp := &CreateOrderResponse{
		OrderId: rsp.OrderId,
	}
	return util.WithSuccess(resp)
}
