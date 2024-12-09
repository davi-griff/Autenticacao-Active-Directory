package authentication

import (
	"auth-ad/src/internal/interfaces"
	"auth-ad/src/internal/models"
	"time"
)

type Authentication struct {
	adService  interfaces.IActiveDirectoryService
	apiService interfaces.IApiService
}

// NewAuthentication cria uma nova instância de Authentication.
// Parâmetros:
// - adService: serviço de autenticação.
// - apiService: serviço de API.
// Retorna: uma nova instância de Authentication.
func NewAuthentication(adService interfaces.IActiveDirectoryService, apiService interfaces.IApiService) *Authentication {
	return &Authentication{
		adService:  adService,
		apiService: apiService,
	}
}

// Start inicia o processo de autenticação.
// Retorna: um erro caso ocorra algum problema durante a execução.
func (a *Authentication) Start() error {
	for {
		requests, err := a.apiService.GetRequest()
		if err != nil {
			return err
		}

		for _, request := range requests {

			authenticated, err := a.adService.Authenticate(request.Username, request.Password)
			if err != nil {
				return err
			}

			if !authenticated {
				continue
			}

			user, err := a.adService.GetUser(request.Username)
			if err != nil {
				return err
			}

			response := models.AuthResponse{
				RequestID: request.RequestID,
				Success:   authenticated,
				UserData: models.UserData{
					Username: user.Username,
					Email:    user.Email,
					Groups:   user.Groups,
				},
			}

			err = a.apiService.SendResponse(request.RequestID, response)
			if err != nil {
				return err
			}

			a.adService.Unbind()
		}

		time.Sleep(1 * time.Second)
	}
}
