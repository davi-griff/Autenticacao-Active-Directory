package smarketAPIGateway

import (
	"auth-ad/src/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequest(t *testing.T) {
	// Configurar servidor mock
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verificar método e path
		if r.Method != "GET" || r.URL.Path != "/v1/auth" {
			t.Errorf("Esperado GET /v1/auth, recebido %s %s", r.Method, r.URL.Path)
		}

		// Verificar headers
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Error("Token de autorização inválido")
		}

		// Responder com dados mock
		mockResponse := []models.AuthRequest{
			{
				RequestID: "123",
				Username:  "test-user",
			},
		}
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Criar gateway com URL do servidor mock
	gateway := &SmarketGateway{
		httpClient: server.Client(),
		baseUrl:    server.URL + "/v1",
		token:      "test-token",
	}

	// Executar teste
	requests, err := gateway.GetRequest()
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	// Verificar resultado
	if len(requests) != 1 {
		t.Errorf("Esperado 1 request, recebido %d", len(requests))
	}
	if requests[0].RequestID != "123" {
		t.Errorf("Esperado ID 123, recebido %s", requests[0].RequestID)
	}
}

func TestSendResponse(t *testing.T) {
	// Configurar servidor mock
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verificar método e path
		if r.Method != "POST" || r.URL.Path != "/v1/auth/123" {
			t.Errorf("Esperado POST /v1/auth/123, recebido %s %s", r.Method, r.URL.Path)
		}

		// Verificar headers
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Error("Token de autorização inválido")
		}

		// Verificar body
		var response models.AuthResponse
		if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
			t.Fatalf("Erro ao decodificar body: %v", err)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Criar gateway com URL do servidor mock
	gateway := &SmarketGateway{
		httpClient: server.Client(),
		baseUrl:    server.URL + "/v1",
		token:      "test-token",
	}

	// Executar teste
	err := gateway.SendResponse("123", models.AuthResponse{
		Success: true,
	})

	// Verificar resultado
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}
}

func TestSendResponse_Error(t *testing.T) {
	// Configurar servidor mock com erro
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("erro de teste"))
	}))
	defer server.Close()

	// Criar gateway com URL do servidor mock
	gateway := &SmarketGateway{
		httpClient: server.Client(),
		baseUrl:    server.URL + "/v1",
		token:      "test-token",
	}

	// Executar teste
	err := gateway.SendResponse("123", models.AuthResponse{
		Success: true,
	})

	// Verificar se o erro foi retornado
	if err == nil {
		t.Fatal("Esperado erro, recebido nil")
	}
}
