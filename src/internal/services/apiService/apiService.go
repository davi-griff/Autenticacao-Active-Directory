package apiService

import (
	"auth-ad/src/internal/interfaces"
	"auth-ad/src/internal/models"
)

type ApiService struct {
	apiRepository interfaces.IApiRepository
}

// NewApiService cria uma nova instância de ApiService.
// Parâmetros:
// - apiRepository: uma implementação da interface IApiRepository.
// Retorno:
// - *ApiService: uma nova instância de ApiService.
func NewApiService(apiRepository interfaces.IApiRepository) *ApiService {
	return &ApiService{
		apiRepository: apiRepository,
	}
}

// GetRequest obtém uma lista de AuthRequest.
// Retorno:
// - []models.AuthRequest: uma lista de solicitações de autenticação.
// - error: um erro, se ocorrer.
func (s *ApiService) GetRequest() ([]models.AuthRequest, error) {
	return s.apiRepository.GetRequest()
}

// SendResponse envia uma resposta de autenticação.
// Parâmetros:
// - requestID: o ID da solicitação.
// - response: a resposta de autenticação a ser enviada.
// Retorno:
// - error: um erro, se ocorrer.
func (s *ApiService) SendResponse(requestID string, response models.AuthResponse) error {
	return s.apiRepository.SendResponse(requestID, response)
}
