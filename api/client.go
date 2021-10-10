package api

import (
	"context"
	v1 "github.com/ffy/kratos-layout/api/manage/v1"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/zookeeper/registry"
)

// DiscoveryId for registry
const DiscoveryId = "github.com/ffy/kratos-layout"

// Client github.com/ffy/kratos-layout
type Client struct {
	ManageClient v1.ManageClient
}

// NewClient github.com/ffy/kratos-layout grpc client
func NewClient(r *registry.Registry) (client *Client, err error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+DiscoveryId),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		return
	}
	client = &Client{
		ManageClient: v1.NewManageClient(conn),
	}
	return
}
