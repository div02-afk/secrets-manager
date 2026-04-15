package main

import (
	"log"
	"net"

	proto "github.com/div02-afk/secrets-manager/gen/kms"
	"github.com/div02-afk/secrets-manager/pkg/kms"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	//TODO: add encryption
	grpcServer := grpc.NewServer()

	KMSImpl := kms.KMS{}

	proto.RegisterKMSServer(grpcServer, kms.KMSService{
		KMS: &KMSImpl,
	})
	log.Println("KMS server running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
