package smarketAPIGateway

import (
	"auth-ad/src/internal/interfaces"
	"auth-ad/src/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SmarketGateway implementa a interface IApiRepository para comunicação com a API Smarket
type SmarketGateway struct {
	httpClient *http.Client
	baseUrl    string
	token      string
}

// NewSmarketGateway cria uma nova instância de SmarketGateway
// Parâmetros:
//   - token: Token de autenticação para a API
//
// Retorna:
//   - interfaces.IApiRepository: Interface implementada pelo gateway
func NewSmarketGateway(token string, baseUrl string) interfaces.IApiRepository {
	return &SmarketGateway{
		httpClient: http.DefaultClient,
		baseUrl:    baseUrl,
		token:      token,
	}
}

// GetRequest busca as requisições de autenticação pendentes
// Retorna:
//   - []models.AuthRequest: Lista de requisições de autenticação
//   - error: Erro em caso de falha na requisição
func (s *SmarketGateway) GetRequest() ([]models.AuthRequest, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth", s.baseUrl), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var responseJsons []models.AuthRequest

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &responseJsons); err != nil {
		return nil, err
	}

	return responseJsons, nil
}

// SendResponse envia uma resposta de autenticação para uma requisição específica
// Parâmetros:
//   - requestId: ID da requisição a ser respondida
//   - response: Dados da resposta de autenticação
//
// Retorna:
//   - error: Erro em caso de falha no envio da resposta
func (s *SmarketGateway) SendResponse(requestId string, response models.AuthResponse) error {

	responseBody, err := json.Marshal(response)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/%s", s.baseUrl, requestId), bytes.NewBuffer(responseBody))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	resp, err := s.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("failed to send response: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}
