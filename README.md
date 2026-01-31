# Telegram Bot Modular em Go

Este é um bot para Telegram desenvolvido em Go utilizando a biblioteca `gogram`, com uma arquitetura modular e suporte a configurações dinâmicas via YAML.

## 🚀 Funcionalidades

- **Arquitetura Modular**: Comandos e callbacks são organizados em módulos separados.
- **Configuração YAML**: Textos e botões são carregados de um arquivo `assets/messages.yml`.
- **Variáveis Dinâmicas**: Suporte a placeholders como `{firstName}`, `{userID}` e `{username}` nos textos e callbacks.
- **GORM & PostgreSQL**: Integração completa com banco de dados para persistência de usuários e configurações.
- **Modo Admin/Manutenção**: Painel administrativo para alternar o modo de manutenção do bot.
- **API/Mini App**: Servidor API básico integrado usando Gin.

## 🛠 Pré-requisitos

- Go 1.25+ (ou Go 1.24 com toolchain configurado)
- PostgreSQL

## 📦 Instalação e Execução

1. Clone o repositório.
2. Configure as variáveis de ambiente:
   ```bash
   export BOT_TOKEN="seu_token_aqui"
   export DATABASE_URL="host=localhost user=postgres password=postgres dbname=telegram_bot port=5432 sslmode=disable"
   ```
3. Instale as dependências:
   ```bash
   go mod tidy
   ```
4. Compile o projeto:
   ```bash
   go build -o bot cmd/main.go
   ```
5. Execute o bot:
   ```bash
   ./bot
   ```

## 📂 Estrutura do Projeto

- `cmd/`: Ponto de entrada da aplicação.
- `internal/bot/`: Core do bot e gerenciamento de handlers.
- `internal/config/`: Carregamento de mensagens e variáveis.
- `internal/database/`: Conexão e migrações do banco de dados.
- `internal/models/`: Modelos GORM.
- `internal/modules/`: Módulos de funcionalidades (Start, Admin, etc).
- `api/`: Servidor API para Mini Apps.
- `assets/`: Arquivos de configuração YAML.

## 📝 Exemplo de Configuração YAML (`assets/messages.yml`)

```yaml
messages:
  - name: start
    text: "<b>👋 Olá, {firstName}!</b>\n\n🤖 Eu estou aqui para automatizar..."
    buttons:
      - - text: "📋 Meus Dados"
          callback_data: "profile-info:{userID}"
        - text: "🆘 Como Usar"
          callback_data: "help"
```

## ⚖️ Licença

Este projeto está sob a licença MIT.
