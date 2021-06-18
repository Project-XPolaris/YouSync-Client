package client

import (
	"google.golang.org/grpc"
	"yousyncclient/pb"
)

var DefaultSyncClient = SyncClient{
	Address:           "localhost:50051",
}

type SyncClient struct {
	Address           string
	Client            pb.FileSyncClient
}

func (c *SyncClient) Run() error {
	conn, err := grpc.Dial(
		c.Address, grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return err
	}
	c.Client = pb.NewFileSyncClient(conn)
	return nil
}
