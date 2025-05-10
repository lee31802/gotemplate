package buy

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type Client interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
}

type BuyService struct {
	name string
	cli  zrpc.Client
}

func GetServiceName() string {
	return "buy"
}

func NewClient(cli zrpc.Client) Client {
	return &BuyService{
		name: GetServiceName(),
		cli:  cli,
	}
}

func (m *BuyService) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	client := NewBuyClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}

//import (
//	"context"
//	microkit1 "github.com/lee31802/comment_lib/gmicrokit"
//	"github.com/micro/go-micro"
//	"github.com/micro/go-micro/client"
//	"github.com/micro/go-micro/server"
//)
//
//const (
//	serviceName = "buy"
//)
//
//var (
//	service micro.Service
//	cli     DemoClient
//)
//
//type DemoClient interface {
//	GetDriverTags(ctx context.Context, in *GetDriverTagsRequest) (*GetDriverTagsResponse, error)
//}
//
//type DemoAService struct {
//	name           string
//	microkitClient client.Client
//}
//
//type callOptsCtxKey struct{}
//
//func (c *DemoAService) GetDriverTags(ctx context.Context, in *GetDriverTagsRequest) (*GetDriverTagsResponse, error) {
//	callOpts, _ := ctx.Value(callOptsCtxKey{}).([]client.CallOption)
//	req := c.microkitClient.NewRequest(c.name, "Demo.GetDriverTags", in)
//	out := new(GetDriverTagsResponse)
//	err := c.microkitClient.Call(ctx, req, out, callOpts...)
//	if err != nil {
//		return nil, err
//	}
//	return out, nil
//}
//func NewService() micro.Service {
//	service = micro.NewService(microkit1.DefaultOptions...)
//	return service
//}
//
//func NewClient() DemoClient {
//	cli = newDemoAService(serviceName, microkit1.DefaultClient)
//	return cli
//}
//
//func newDemoAService(name string, c client.Client) DemoClient {
//	if c == nil {
//		c = client.NewClient()
//	}
//	if len(name) == 0 {
//		name = serviceName
//	}
//	return &DemoAService{
//		microkitClient: c,
//		name:           name,
//	}
//}
//
//type DemoHandler interface {
//	GetDriverTags(context.Context, *GetDriverTagsRequest, *GetDriverTagsResponse) error
//}
//
//// 可以约束他必须实现了所有的方法
//func RegisterDemoHandler(hdlr DemoHandler, opts ...server.HandlerOption) error {
//	return registerFoodyRatingHandler(service.Server(), hdlr, opts...)
//}
//
//type demoHandler struct {
//	DemoHandler
//}
//
//func registerFoodyRatingHandler(s server.Server, hdlr DemoHandler, opts ...server.HandlerOption) error {
//	type foodyRating interface {
//		GetDriverTags(ctx context.Context, in *GetDriverTagsRequest, out *GetDriverTagsResponse) error
//	}
//	type FoodyRating struct {
//		DemoHandler
//	}
//	h := &demoHandler{hdlr}
//	return s.Handle(s.NewHandler(&FoodyRating{h}, opts...))
//}
