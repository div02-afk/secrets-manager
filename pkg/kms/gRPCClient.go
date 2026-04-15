package kms

import (
	"context"
	"time"

	pb "github.com/div02-afk/secrets-manager/gen/kms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCKMSClient struct {
	client  pb.KMSClient
	timeout time.Duration
}

func NewGRPCKMSClient(client pb.KMSClient) *GRPCKMSClient {
	return &GRPCKMSClient{
		client:  client,
		timeout: 5 * time.Second,
	}
}

func NewClient(addr string) (*GRPCKMSClient, error) {
    conn, err := grpc.NewClient(
        addr,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
        return nil, err
    }

    client := pb.NewKMSClient(conn)
    return NewGRPCKMSClient(client), nil
}

func (c *GRPCKMSClient) Encrypt(dek []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	res, err := c.client.Encrypt(ctx, &pb.EncryptRequest{
		Dek: dek,
	})
	if err != nil {
		return nil, err
	}

	return res.EncryptedDek, nil
}

func (c *GRPCKMSClient) Decrypt(encryptedDEK []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	res, err := c.client.Decrypt(ctx, &pb.DecryptRequest{
		EncryptedDek: encryptedDEK,
	})
	if err != nil {
		return nil, err
	}

	return res.Dek, nil
}
