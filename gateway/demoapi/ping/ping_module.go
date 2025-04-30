package ping

import (
	"github.com/lee31802/comment_lib/gweb"
	"github.com/lee31802/gotemplate/util"
	//"git.garena.com/shopee/foody/service/buyerapi/resp"
)

type PingModule struct{}

func (m *PingModule) Init(r gweb.Router) {
	group := r.Group("api/buyer")
	{
		group.GET("/ping", m.Get)
	}
}

func (m *PingModule) Get() gweb.Response {
	return util.WithSuccess(PingModule{})
}
