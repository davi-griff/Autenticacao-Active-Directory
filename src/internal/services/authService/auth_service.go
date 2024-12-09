package authService

import (
	"auth-ad/src/internal/interfaces"
	"auth-ad/src/internal/models"
)

// AuthService fornece métodos para autenticação e recuperação de dados de usuários.
type AuthService struct {
	adRepository interfaces.IActiveDirectoryRepository
}

// NewAuthService cria uma nova instância de AuthService.
//
// Parâmetros:
//   - adRepository: Interface para o repositório do Active Directory.
//
// Retorna:
//   - Ponteiro para uma nova instância de AuthService.
func NewAuthService(adRepository interfaces.IActiveDirectoryRepository) *AuthService {
	return &AuthService{adRepository: adRepository}
}

// Authenticate verifica as credenciais do usuário.
//
// Parâmetros:
//   - username: Nome de usuário para autenticação.
//   - password: Senha do usuário.
//
// Retorna:
//   - bool: Verdadeiro se a autenticação for bem-sucedida, falso caso contrário.
//   - error: Erro, se ocorrer.
func (s *AuthService) Authenticate(username, password string) (bool, error) {
	return s.adRepository.Authenticate(username, password)
}

// GetUser recupera os dados de um usuário pelo nome de usuário.
//
// Parâmetros:
//   - username: Nome de usuário para recuperar os dados.
//
// Retorna:
//   - *models.UserData: Dados do usuário.
//   - error: Erro, se ocorrer.
func (s *AuthService) GetUser(username string) (models.UserData, error) {
	user, err := s.adRepository.GetUser(username)
	if err != nil {
		return models.UserData{}, err
	}

	return models.UserData{
		Username: user.SAMAccountName,
		Email:    user.Email,
		Groups:   user.Groups,
	}, nil
}

// GetUsers recupera os dados de todos os usuários de um grupo.
//
// Parâmetros:
//   - group: Nome do grupo para recuperar os usuários.
//
// Retorna:
//   - []*models.UserData: Lista de dados dos usuários.
//   - error: Erro, se ocorrer.
func (s *AuthService) GetUsers(group string) ([]models.UserData, error) {
	users, err := s.adRepository.GetUsers(group)
	if err != nil {
		return nil, err
	}

	userData := make([]models.UserData, 0)
	for _, user := range users {
		userData = append(userData, models.UserData{
			Username: user.SAMAccountName,
			Email:    user.Email,
			Groups:   user.Groups,
		})
	}

	return userData, nil
}

// Unbind remove a vinculação atual da conexão
// Returns:
//   - error: Erro em caso de falha no unbind
func (s *AuthService) Unbind() error {
	return s.adRepository.Unbind()
}

// Close fecha a conexão com o Active Directory.
func (s *AuthService) Close() error {
	return s.adRepository.Close()
}
