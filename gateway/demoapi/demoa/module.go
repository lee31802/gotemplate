package demoa

import (
	"context"
	"github.com/lee31802/comment_lib/gweb"
	"github.com/lee31802/gotemplate/gateway/demoapi/client"
	"github.com/lee31802/gotemplate/pb/demoa"
	"github.com/lee31802/gotemplate/util"
)

type DemoModule struct {
	client demoa.DemoClient
}

func NewDemoModule() *DemoModule {
	return &DemoModule{}
}

func (m *DemoModule) Init(r gweb.Router) {
	m.client = client.DemoAClient()

	group := r.Group("api/buyer/rating")
	{
		group.GET("/tags/driver", m.GetDriverTags)
	}
}

func (m *DemoModule) GetDriverTags(ctx context.Context) gweb.Response {
	rsp, err := m.client.GetDriverTags(ctx, &demoa.GetDriverTagsRequest{})
	if err != nil {
		return util.WithError(err)
	}
	return util.WithSuccess(rsp)
}
