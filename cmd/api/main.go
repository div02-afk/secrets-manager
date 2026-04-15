package main

import (
	"encoding/json"
	"log"
	"net/http"

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
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		var req storage.SecretDTO
		log.Println(r.Method, " ", r.RequestURI)
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println("Invalid Request: ", err)
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		version, err := secretService.Add(req.TenantID, req.SecretKey, req.SecretValue)
		if err != nil {
			log.Println("Add Failed with error: ", err)
			http.Error(w, "New Secret Add failed", http.StatusBadRequest)
		}
		resp := map[string][]byte{
			"version": {byte(version)},
		}
		w.Write((resp["version"]))

	})

	log.Println("HTTP server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
