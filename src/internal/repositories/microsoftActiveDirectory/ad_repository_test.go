package microsoftActiveDirectory

import (
	"auth-ad/src/pkg/configs"
	"testing"

	"github.com/go-ldap/ldap/v3"
	"github.com/stretchr/testify/assert"
)

// MockLDAPConn é uma estrutura de mock para a conexão LDAP
type MockLDAPConn struct {
	BindFunc   func(username, password string) error
	SearchFunc func(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error)
	CloseFunc  func() error
	UnbindFunc func() error
}

func (m *MockLDAPConn) Bind(username, password string) error {
	return m.BindFunc(username, password)
}

func (m *MockLDAPConn) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	return m.SearchFunc(searchRequest)
}

func (m *MockLDAPConn) Close() error {
	return m.CloseFunc()
}

func (m *MockLDAPConn) Unbind() error {
	return m.UnbindFunc()
}

func TestADRepository_Authenticate(t *testing.T) {
	mockConn := &MockLDAPConn{
		BindFunc: func(username, password string) error {
			if username == "validUser@domain.com" && password == "validPassword" {
				return nil
			}
			return ldap.NewError(ldap.LDAPResultInvalidCredentials, nil)
		},
	}

	repo := &ADRepository{conn: mockConn, config: &configs.ADConfig{Domain: "domain.com"}}

	success, err := repo.Authenticate("validUser", "validPassword")
	assert.NoError(t, err)
	assert.True(t, success)

	success, err = repo.Authenticate("invalidUser", "invalidPassword")
	assert.Error(t, err)
	assert.False(t, success)
}

func TestADRepository_GetUser(t *testing.T) {
	mockConn := &MockLDAPConn{
		SearchFunc: func(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
			entry := ldap.NewEntry("cn=Test User,dc=example,dc=com", map[string][]string{
				"uid":               {"testuser"},
				"cn":                {"Test User"},
				"mail":              {"test@example.com"},
				"sAMAccountName":    {"testuser"},
				"userPrincipalName": {"testuser@example.com"},
			})
			return &ldap.SearchResult{Entries: []*ldap.Entry{entry}}, nil
		},
	}

	repo := &ADRepository{conn: mockConn, config: &configs.ADConfig{BaseDN: "dc=example,dc=com"}}

	user, err := repo.GetUser("testuser")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.SAMAccountName)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestADRepository_GetUsers(t *testing.T) {
	mockConn := &MockLDAPConn{
		SearchFunc: func(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
			if searchRequest.Filter == "(&(objectClass=group)(cn=TestGroup))" {
				entry := ldap.NewEntry("cn=TestGroup,dc=example,dc=com", nil)
				return &ldap.SearchResult{Entries: []*ldap.Entry{entry}}, nil
			}

			entries := []*ldap.Entry{
				ldap.NewEntry("cn=User1,dc=example,dc=com", map[string][]string{
					"uid":               {"user1"},
					"cn":                {"User One"},
					"mail":              {"user1@example.com"},
					"sAMAccountName":    {"user1"},
					"userPrincipalName": {"user1@example.com"},
				}),
				ldap.NewEntry("cn=User2,dc=example,dc=com", map[string][]string{
					"uid":               {"user2"},
					"cn":                {"User Two"},
					"mail":              {"user2@example.com"},
					"sAMAccountName":    {"user2"},
					"userPrincipalName": {"user2@example.com"},
				}),
			}
			return &ldap.SearchResult{Entries: entries}, nil
		},
	}

	repo := &ADRepository{conn: mockConn, config: &configs.ADConfig{BaseDN: "dc=example,dc=com"}}

	users, err := repo.GetUsers("TestGroup")
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "user1", users[0].SAMAccountName)
	assert.Equal(t, "user2", users[1].SAMAccountName)
}

func TestADRepository_Close(t *testing.T) {
	mockConn := &MockLDAPConn{
		CloseFunc: func() error {
			return nil
		},
	}

	repo := &ADRepository{conn: mockConn}

	err := repo.Close()
	assert.NoError(t, err)
}

func TestADRepository_Bind(t *testing.T) {
	mockConn := &MockLDAPConn{
		BindFunc: func(username, password string) error {
			if username == "validUser@domain.com" && password == "validPassword" {
				return nil
			}
			return ldap.NewError(ldap.LDAPResultInvalidCredentials, nil)
		},
	}

	repo := &ADRepository{conn: mockConn, config: &configs.ADConfig{Domain: "domain.com"}}

	err := repo.Bind("validUser", "validPassword")
	assert.NoError(t, err)

	err = repo.Bind("invalidUser", "invalidPassword")
	assert.Error(t, err)
}

func TestADRepository_Unbind(t *testing.T) {
	mockConn := &MockLDAPConn{
		UnbindFunc: func() error {
			return nil
		},
	}

	repo := &ADRepository{conn: mockConn}

	err := repo.Unbind()
	assert.NoError(t, err)
}
