package configs

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	// Teste com arquivo .env inexistente
	err := LoadEnv()
	if err == nil {
		t.Error("Esperava erro ao carregar arquivo .env inexistente")
	}

	// Criar arquivo .env temporário para teste
	content := []byte(`
AD_SERVER=ldap.exemplo.com
AD_PORT=389
AD_DOMAIN=exemplo.com
AD_USERNAME=admin
AD_PASSWORD=senha123
AD_BASE_DN=dc=exemplo,dc=com
API_URL=https://api-gtw.smarketsolutions.com.br/v1
`)
	err = os.WriteFile(".env", content, 0644)
	if err != nil {
		t.Fatalf("Erro ao criar arquivo .env de teste: %v", err)
	}
	defer os.Remove(".env")

	// Teste com arquivo .env válido
	err = LoadEnv()
	if err != nil {
		t.Errorf("Não esperava erro ao carregar arquivo .env válido: %v", err)
	}
}

func TestGetADConfig(t *testing.T) {
	// Configurar variáveis de ambiente para teste
	os.Setenv("AD_SERVER", "ldap.exemplo.com")
	os.Setenv("AD_PORT", "389")
	os.Setenv("AD_DOMAIN", "exemplo.com")
	os.Setenv("AD_USERNAME", "admin")
	os.Setenv("AD_PASSWORD", "senha123")
	os.Setenv("AD_BASE_DN", "dc=exemplo,dc=com")
	os.Setenv("API_URL", "https://api-gtw.smarketsolutions.com.br/v1")

	config, err := GetADConfig()
	if err != nil {
		t.Errorf("Não esperava erro ao obter configurações: %v", err)
	}

	// Verificar valores carregados
	if config.Server != "ldap.exemplo.com" {
		t.Errorf("Server incorreto, obtido: %s, esperado: %s", config.Server, "ldap.exemplo.com")
	}
	if config.Port != 389 {
		t.Errorf("Port incorreta, obtido: %d, esperado: %d", config.Port, 389)
	}

	// Teste com porta inválida
	os.Setenv("AD_PORT", "porta_invalida")
	_, err = GetADConfig()
	if err == nil {
		t.Error("Esperava erro ao converter porta inválida")
	}
}
