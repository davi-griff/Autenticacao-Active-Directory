# Autenticação Active Directory

Serviço de autenticação que integra com Microsoft Active Directory e expõe uma API para validação de credenciais e obtenção de informações de usuários.

## 🚀 Tecnologias

- Go 1.23.2
- LDAP (Lightweight Directory Access Protocol)
- REST API

## 📚 Bibliotecas Principais

- `github.com/go-ldap/ldap/v3`: Cliente LDAP para Go
- `github.com/joho/godotenv`: Carregamento de variáveis de ambiente
- `github.com/stretchr/testify`: Framework de testes
- `github.com/google/uuid`: Geração de UUIDs
- `golang.org/x/crypto`: Funções criptográficas

## 🔧 Configuração

1. Clone o repositório:
```bash
git clone [url-do-repositorio]
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure o arquivo `.env` na raiz do projeto:
```env
AD_SERVER=seu-servidor-ad
AD_PORT=389
AD_DOMAIN=seu.dominio
AD_USERNAME=usuario-admin
AD_PASSWORD=senha-admin
AD_BASE_DN=DC=seu,DC=dominio
```

## ⚙️ Variáveis de Ambiente

| Variável | Descrição |
|----------|-----------|
| AD_SERVER | Endereço do servidor Active Directory |
| AD_PORT | Porta do servidor AD (geralmente 389) |
| AD_DOMAIN | Domínio do AD |
| AD_USERNAME | Usuário com permissões de administrador |
| AD_PASSWORD | Senha do usuário administrador |
| AD_BASE_DN | DN base para pesquisas LDAP |
| API_URL | URL da API de autenticação |

## 🚀 Executando o Projeto

### Via linha de comando:
```bash
go run src/cmd/main.go
```

### Via VSCode:
1. Abra o projeto no VSCode
2. Use a configuração de debug presente em `.vscode/launch.json`
3. Pressione F5 para iniciar em modo debug

## 🧪 Executando os Testes
```bash
go test ./...
```

## 🏗️ Estrutura do Projeto

```
src/
├── cmd/
│   └── main.go
├── internal/
│   ├── authentication/
│   ├── interfaces/
│   ├── models/
│   ├── repositories/
│   └── services/
└── pkg/
    └── configs/
```

## 🔍 Funcionalidades Principais

- Autenticação de usuários contra Active Directory
- Recuperação de informações de usuários
- Busca de usuários por grupo
- Interface REST para integração com outros sistemas

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## 🎯 Status do Projeto

Em desenvolvimento ativo.

---

⌨️ com ❤️ por Davi Araujo 😊