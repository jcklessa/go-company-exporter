# Go Company Exporter

Uma aplicação simples em Go que:

- Lê códigos de um arquivo `.txt`
- Consulta uma tabela `company` no PostgreSQL
- Gera um arquivo `.xlsx` com os dados encontrados
- Cria uma segunda aba no Excel com os códigos não encontrados

## 🛠 Pré-requisitos

- Go 1.18+
- PostgreSQL rodando e acessível
- Arquivo `codes_all.txt` com os códigos (um por linha)
- Arquivo `.env` com as credenciais do banco

## 📁 Estrutura do Projeto

```
go-company-exporter/
│
├── cmd/
│   └── main.go             # Ponto de entrada
│
├── config/
│   └── db.go               # Configuração do banco
│
├── internal/
│   ├── models/
│   │   └── company.go      # Structs e tipos
│   ├── services/
│   │   └── company.go      # Funções de serviço (fetch, processamento)
│   └── utils/
│       ├── excel.go        # Geração do Excel
│       └── file.go         # Leitura do .txt
│
├── .env                    # Variáveis de ambiente (não versionado)
├── .gitignore
└── README.md
```

## ⚙️ Variáveis de Ambiente

Crie um arquivo `.env` na raiz com o seguinte conteúdo:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=apprest
```

> ⚠️ **Nunca versionar este arquivo.** O `.env` já está no `.gitignore`.

## 🚀 Como executar

```bash
go run cmd/main.go
```

## 📦 Saída

Será gerado um arquivo `companies.xlsx` com duas abas:

- **Resultados** → Dados encontrados na base
- **NaoEncontrados** → Códigos que não foram localizados

## 🧪 Testando

Você pode adicionar testes unitários em `internal/services` ou `internal/utils` para validar a lógica.

## 📄 Licença

MIT. Sinta-se livre para modificar e adaptar este projeto.
