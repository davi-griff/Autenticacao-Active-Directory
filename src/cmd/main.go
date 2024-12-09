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

	apiRepository := smarketAPIGateway.NewSmarketGateway("1234567890")

	apiService := apiService.NewApiService(apiRepository)
	authService := authService.NewAuthService(adRepository)

	authentication := authentication.NewAuthentication(authService, apiService)

	err = authentication.Start()
	if err != nil {
		authService.Close()
		log.Fatalf("Erro ao iniciar a autenticação: %v", err)
	}

	// success, err := authService.Authenticate("smarket.ad", "cHL_bm9N@10")
	// if err != nil {
	// 	log.Fatalf("Erro ao autenticar: %v", err)
	// }

	// fmt.Println(success)

	// user, err := authService.GetUser("smarket.ad")
	// if err != nil {
	// 	log.Fatalf("Erro ao buscar usuário: %v", err)
	// }

	// users, err := authService.GetUsers("TI - Colaboradores")
	// if err != nil {
	// 	log.Fatalf("Erro ao buscar usuários: %v", err)
	// }

	// // Configuração
	// config := ADConfig{
	// 	// Server:   "dtc02srvad01.koch.intranet",
	// 	// Server:   "172.19.20.1",
	// 	Server:   "localhost",
	// 	Port:     389,
	// 	Domain:   "koch.intranet",
	// 	Username: "smarket.ad",
	// 	Password: "cHL_bm9N@10",
	// 	BaseDN:   "DC=koch,DC=intranet", // Construído a partir do domínio
	// }
}
