package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// ADConfig representa as configurações de conexão com o Active Directory
type ADConfig struct {
	Server   string // Endereço do servidor AD
	Port     int    // Porta de conexão
	Domain   string // Domínio do AD
	Username string // Nome de usuário para autenticação
	Password string // Senha para autenticação
	BaseDN   string // DN base para pesquisas
	ApiUrl   string // URL da API
}

// LoadEnv carrega as variáveis de ambiente do arquivo .env
// Retorna error em caso de falha ao carregar o arquivo
func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("erro ao carregar variáveis de ambiente: %v", err)
	}

	return nil
}

// GetADConfig recupera as configurações do Active Directory das variáveis de ambiente
// Retorna:
//   - *ADConfig: estrutura com as configurações carregadas
//   - error: erro em caso de falha ao carregar ou converter valores
func GetADConfig() (*ADConfig, error) {
	server := os.Getenv("AD_SERVER")
	port, err := strconv.Atoi(os.Getenv("AD_PORT"))
	if err != nil {
		return nil, fmt.Errorf("erro ao converter porta: %v", err)
	}
	domain := os.Getenv("AD_DOMAIN")
	username := os.Getenv("AD_USERNAME")
	password := os.Getenv("AD_PASSWORD")
	baseDN := os.Getenv("AD_BASE_DN")
	apiUrl := os.Getenv("API_URL")

	return &ADConfig{
		Server:   server,
		Port:     port,
		Domain:   domain,
		Username: username,
		Password: password,
		BaseDN:   baseDN,
		ApiUrl:   apiUrl,
	}, nil
}
