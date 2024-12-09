package main

import (
	"auth-ad/src/internal/authentication"
	"auth-ad/src/internal/repositories/microsoftActiveDirectory"
	"auth-ad/src/internal/repositories/smarketAPIGateway"
	"auth-ad/src/internal/services/apiService"
	"auth-ad/src/internal/services/authService"
	"auth-ad/src/pkg/configs"
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func main() {

	configs.LoadEnv()

	adConfig, err := configs.GetADConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar as configurações: %v", err)
	}

	ldapConn, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%d", adConfig.Server, adConfig.Port))
	if err != nil {
		log.Fatalf("Erro ao conectar ao AD: %v", err)
	}

	adRepository, err := microsoftActiveDirectory.NewADRepository(adConfig, ldapConn)
	if err != nil {
		log.Fatalf("Erro ao criar o repositório: %v", err)
	}

	apiRepository := smarketAPIGateway.NewSmarketGateway("1234567890", adConfig.ApiUrl)

	apiService := apiService.NewApiService(apiRepository)
	authService := authService.NewAuthService(adRepository)

	authentication := authentication.NewAuthentication(authService, apiService)

	err = authentication.Start()
	if err != nil {
		authService.Close()
		log.Fatalf("Erro ao iniciar a autenticação: %v", err)
	}
}
