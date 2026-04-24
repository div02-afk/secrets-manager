package main

import (
	"log"
	"net/http"

	apihandler "github.com/div02-afk/secrets-manager/pkg/api-handler"
	"github.com/div02-afk/secrets-manager/pkg/auth"
	"github.com/div02-afk/secrets-manager/pkg/encryption"
	"github.com/div02-afk/secrets-manager/pkg/kms"
	"github.com/div02-afk/secrets-manager/pkg/secret"
	"github.com/div02-afk/secrets-manager/pkg/storage"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	KMSClient, err := kms.NewClient("localhost:50051")
	if err != nil {
		log.Fatal("KMS gRPC client startup failed with: ", err)
	}
	
	storageImpl := storage.NewPostgresStorage()
	secretService := secret.NewSecretService(KMSClient, storageImpl, &encryption.AESProvider{})
	authProvider := auth.CreateAPIAuthProvider(storageImpl)

	apiHandler := apihandler.CreateHttpApiHandler(secretService,authProvider)

	http.HandleFunc("/add", apiHandler.AddSecret)
	http.HandleFunc("/get",apiHandler.GetSecretValue)

	log.Println("HTTP server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
