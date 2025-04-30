package handler

import (
	"context"
	"github.com/lee31802/gotemplate/pb/demoa"
)

type DemoHandler struct {
}

func (s *DemoHandler) GetDriverTags(ctx context.Context, req *demoa.GetDriverTagsRequest, rsp *demoa.GetDriverTagsResponse) error {
	return nil
}
