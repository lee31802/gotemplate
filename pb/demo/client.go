package pb

import "context"

type Client interface {
	GetDriverTags(ctx context.Context, in *GetDriverTagsRequest, opts ...callopt.Option) (*GetDriverTagsResponse, error)
}

func (client *foodyRatingClient) GetDriverTags(ctx context.Context, in *GetDriverTagsRequest, opts ...callopt.Option) (*GetDriverTagsResponse, error) {
	if greySwitch(ctx, client.namespace, "get_driver_tags") { // grey related
		callOpts, _ := ctx.Value(callOptsCtxKey{}).([]microclient.CallOption)
		return client.microkitClient.GetDriverTags(ctx, in, callOpts...)
	}

	out := new(GetDriverTagsResponse)

	err := client.c.Invoke(ctx, client.namespace+"."+"get_driver_tags", in, out, opts...)
	if err != nil {
		return nil, err
	}

	return out, nil
}
