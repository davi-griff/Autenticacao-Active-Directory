package microsoftActiveDirectory

import (
	"auth-ad/src/internal/interfaces"
	"auth-ad/src/internal/models"
	"auth-ad/src/pkg/configs"
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

// ILDAPConnection define a interface para operações LDAP
type ILDAPConnection interface {
	Bind(username, password string) error
	Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error)
	Close() error
	Unbind() error
}

// ADRepository implementa a interface IActiveDirectoryInterface para interação com o Active Directory
type ADRepository struct {
	conn   ILDAPConnection
	config *configs.ADConfig
}

// NewADRepository cria uma nova instância do repositório AD que implementa IActiveDirectoryInterface
// Params:
//   - config: Configurações de conexão com o Active Directory
//
// Returns:
//   - interfaces.IActiveDirectoryRepository: Interface implementada do repositório
//   - error: Erro em caso de falha na conexão
func NewADRepository(config *configs.ADConfig, conn ILDAPConnection) (interfaces.IActiveDirectoryRepository, error) {
	return &ADRepository{conn: conn, config: config}, nil
}

// Close fecha a conexão com o Active Directory
// Returns:
//   - error: Erro em caso de falha ao fechar a conexão
func (r *ADRepository) Close() error {
	return r.conn.Close()
}

// Authenticate realiza a autenticação do usuário no Active Directory
// Params:
//   - username: Nome do usuário
//   - password: Senha do usuário
//
// Returns:
//   - bool: true se autenticação for bem sucedida
//   - error: Erro em caso de falha na autenticação
func (r *ADRepository) Authenticate(username, password string) (bool, error) {
	userDN := fmt.Sprintf("%s@%s", username, r.config.Domain)
	err := r.conn.Bind(userDN, password)
	if err != nil {
		return false, fmt.Errorf("erro na autenticação: %v", err)
	}

	return true, nil
}

// Bind realiza a vinculação da conexão com as credenciais do usuário
// Params:
//   - username: Nome do usuário
//   - password: Senha do usuário
//
// Returns:
//   - error: Erro em caso de falha no bind
func (r *ADRepository) Bind(username, password string) error {
	userDN := fmt.Sprintf("%s@%s", username, r.config.Domain)
	return r.conn.Bind(userDN, password)
}

// Unbind remove a vinculação atual da conexão
// Returns:
//   - error: Erro em caso de falha no unbind
func (r *ADRepository) Unbind() error {
	return r.conn.Unbind()
}

// GetUser busca informações de um usuário específico no Active Directory
// Params:
//   - username: Nome do usuário a ser buscado
//
// Returns:
//   - *models.ADUser: Dados do usuário encontrado
//   - error: Erro em caso de falha na busca
func (r *ADRepository) GetUser(username string) (*models.ADUser, error) {
	searchRequest := ldap.NewSearchRequest(
		r.config.BaseDN,        // BaseDN
		ldap.ScopeWholeSubtree, // Escopo
		ldap.NeverDerefAliases, // Dereferencing
		0,                      // Limite de tamanho (0 = sem limite)
		0,                      // Limite de tempo (0 = sem limite)
		false,                  // Somente tipos
		fmt.Sprintf("(&(objectClass=user)(sAMAccountName=%s))", username),                                                                // Filtro
		[]string{"cn", "mail", "sAMAccountName", "userPrincipalName", "distinguishedName", "department", "mailNickname", "title", "uid"}, // Atributos que queremos retornar
		nil,
	)

	result, err := r.conn.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário: %v", err)
	}

	if len(result.Entries) == 0 {
		return nil, fmt.Errorf("usuário não encontrado")
	}

	user := result.Entries[0]

	user.PrettyPrint(4)

	return &models.ADUser{
		UID:               user.GetAttributeValue("uid"),
		DN:                user.DN,
		CN:                user.GetAttributeValue("cn"),
		Email:             user.GetAttributeValue("mail"),
		SAMAccountName:    user.GetAttributeValue("sAMAccountName"),
		UserPrincipalName: user.GetAttributeValue("userPrincipalName"),
	}, nil
}

// GetUsers busca todos os usuários pertencentes a um grupo específico
// Params:
//   - group: Nome do grupo a ser consultado
//
// Returns:
//   - []*models.ADUser: Lista de usuários encontrados no grupo
//   - error: Erro em caso de falha na busca
func (r *ADRepository) GetUsers(group string) ([]*models.ADUser, error) {
	groupFilter := fmt.Sprintf("(&(objectClass=group)(cn=%s))", group)
	searchRequest := ldap.NewSearchRequest(
		r.config.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		groupFilter,
		[]string{"distinguishedName", "member"},
		nil,
	)

	result, err := r.conn.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %v", err)
	}

	if len(result.Entries) == 0 {
		return nil, fmt.Errorf("grupo não encontrado")
	}

	groupDN := result.Entries[0].DN

	userFilter := fmt.Sprintf("(&(objectClass=user)(objectCategory=person)(memberOf:1.2.840.113556.1.4.1941:=%s))", groupDN)
	userSearchRequest := ldap.NewSearchRequest(
		r.config.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		userFilter,
		[]string{"cn", "mail", "sAMAccountName", "userPrincipalName", "distinguishedName", "department", "mailNickname", "title", "uid"},
		nil,
	)

	userResult, err := r.conn.Search(userSearchRequest)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuários: %v", err)
	}

	users := make([]*models.ADUser, 0)
	for _, user := range userResult.Entries {
		users = append(users, &models.ADUser{
			UID:               user.GetAttributeValue("uid"),
			DN:                user.DN,
			CN:                user.GetAttributeValue("cn"),
			Email:             user.GetAttributeValue("mail"),
			SAMAccountName:    user.GetAttributeValue("sAMAccountName"),
			UserPrincipalName: user.GetAttributeValue("userPrincipalName"),
		})
	}

	return users, nil
}
