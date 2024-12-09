# Changelog
Todas as alterações notáveis neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/),
e este projeto adere ao [Semantic Versioning](https://semver.org/lang/pt-BR/).

## [0.1.0] - 2024-12-09

### Adicionado
- Integração inicial com Microsoft Active Directory para autenticação de usuários
- API REST para integração com outros sistemas
- Funcionalidade de autenticação de usuários contra AD
- Recuperação de informações de usuários do AD
- Busca de usuários por grupo no AD
- Sistema de configuração via variáveis de ambiente
- Suporte a arquivo .env para configuração local
- Testes unitários para os principais componentes
- Estrutura inicial do projeto seguindo princípios SOLID
- Documentação básica no README
- Configuração de debug para VSCode

### Segurança
- Implementação de autenticação segura via LDAP
- Gerenciamento seguro de credenciais via variáveis de ambiente

[0.1.0]: https://github.com/seu-usuario/auth-ad/releases/tag/v0.1.0 