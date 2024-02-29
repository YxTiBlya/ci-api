package executor

import (
	"context"

	"google.golang.org/grpc"
)

type clientOption = func(*Client)

func NewClient(address string, opts ...clientOption) *Client {
	c := &Client{
		address:     address,
		dialOptions: []grpc.DialOption{},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithClientDialOptions(dialOptions ...grpc.DialOption) clientOption {
	return func(c *Client) {
		c.dialOptions = append(c.dialOptions, dialOptions...)
	}
}

type Client struct {
	ExecutorAPIClient
	address     string
	conn        *grpc.ClientConn
	dialOptions []grpc.DialOption
}

func (c *Client) Start(ctx context.Context) error {
	var err error
	c.conn, err = grpc.DialContext(ctx, c.address, c.dialOptions...)
	if err != nil {
		return err
	}

	c.ExecutorAPIClient = NewExecutorAPIClient(c.conn)

	return nil
}

func (c *Client) Stop(ctx context.Context) error {
	return c.conn.Close()
}
