package apihandler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/div02-afk/secrets-manager/pkg/auth"
	"github.com/div02-afk/secrets-manager/pkg/secret"
	"github.com/div02-afk/secrets-manager/pkg/storage"
)

type HttpApiHandler struct {
	secretService *secret.SecretService
	auth          auth.Auth
}

func CreateHttpApiHandler(s *secret.SecretService, a auth.Auth) *HttpApiHandler {
	return &HttpApiHandler{
		secretService: s,
		auth:          a,
	}
}

func (h *HttpApiHandler) AddSecret(w http.ResponseWriter, r *http.Request) {
	var req storage.SecretDTO
	log.Println(r.Method, " ", r.RequestURI)
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	apiKey := r.Header.Get("x-api-key")
	if apiKey == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Invalid Request: ", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	authResult, err := h.auth.Validate(req.TenantID, apiKey)

	if err != nil || !authResult {
		log.Println("auth failed with: ",err,authResult)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	version, err := h.secretService.Add(req.TenantID, req.SecretKey, req.SecretValue)
	if err != nil {
		log.Println("Add Failed with error: ", err)
		http.Error(w, "New Secret Add failed", http.StatusBadRequest)
	}
	resp := map[string][]byte{
		"version": {byte(version)},
	}
	w.Write((resp["version"]))

}
func (h *HttpApiHandler) GetSecretValue(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, " ", r.RequestURI)
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	apiKey := r.Header.Get("x-api-key")
	if apiKey == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !r.URL.Query().Has("id") || !r.URL.Query().Has("secret") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tenetIDString := r.URL.Query().Get("id")
	secret := r.URL.Query().Get("secret")
	versionString := r.URL.Query().Get("version")

	tenetID, err := strconv.ParseInt(tenetIDString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authResult, err := h.auth.Validate(tenetID, apiKey)

	if err != nil || !authResult {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var version *int
	if versionString != "" {
		v, err := strconv.Atoi(versionString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		version = &v
	}

	result, err := h.secretService.Get(tenetID, secret, version)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(result)

}
