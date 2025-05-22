✅ Visão geral do que vamos fazer:

    Estrutura básica da API Go com GIN + GORM + PostgreSQL

    Modelo Book

    Endpoint POST /books para adicionar livros

    Testar no Insomnia

📦 Estrutura básica esperada:

go-mvc-api/
├── main.go
├── config/
│   └── database.go
├── models/
│   └── book.go
├── controllers/
│   └── book_controller.go
├── routes/
│   └── routes.go

main.go

// main.go - ponto de entrada da aplicação
//
// Responsabilidades:
// - Inicializar o servidor HTTP com Gin
// - Configurar middlewares básicos (Logger, Recovery, CORS se necessário)
// - Estabelecer conexão com o banco via config
// - Registrar as rotas da API (centralizar endpoints)
// - Iniciar o servidor e escutar requisições na porta configurada
//
// Observações:
// - Mantenha o main.go limpo e enxuto, delegando lógica para outras camadas
// - Facilita testes e manutenção da aplicação

config/database.go

// config/database.go - configuração do banco de dados
//
// Responsabilidades:
// - Definir string de conexão segura (preferencialmente por variáveis ambiente)
// - Abrir conexão com PostgreSQL usando GORM
// - Gerenciar pool de conexões e configurações específicas do banco
// - Executar migrações automáticas para criar/alterar tabelas conforme os models
// - Exportar a instância DB para ser utilizada globalmente
//
// Boas práticas:
// - Usar variáveis de ambiente para credenciais (não deixar hardcoded)
// - Tratar erros de conexão para evitar crash inesperado
// - Configurar logs e níveis de debug para banco de dados

controllers/book_controller.go

// controllers/book_controller.go - controladores HTTP
//
// Responsabilidades:
// - Receber as requisições HTTP (body, params, headers)
// - Validar entrada e formatar saída (status HTTP, JSON)
// - Chamar a camada services para execução da lógica de negócio
// - Tratar erros para devolver respostas apropriadas (ex: 400, 404, 500)
// - Manter os handlers simples, focados em orquestração de dados
//
// Observações:
// - Evitar lógica complexa e acessos diretos ao banco aqui
// - Usar binding e validação integrada do Gin para simplificar

models/book.go

// models/book.go - definição dos modelos de dados
//
// Responsabilidades:
// - Representar a estrutura dos dados no banco de forma clara
// - Definir os campos e seus tipos compatíveis com o banco (int, string, etc.)
// - Usar tags para mapeamento ORM e serialização JSON
// - Facilitar manutenção e documentação da estrutura de dados
//
// Boas práticas:
// - Separar modelos por domínio (ex: book.go para livros, user.go para usuários)
// - Evitar lógica nos modelos; usar serviços para regras de negócio

repositories/book_repository.go

// repositories/book_repository.go - acesso a dados (DAO)
//
// Responsabilidades:
// - Implementar métodos para operações CRUD no banco com GORM
// - Isolar queries SQL e operações diretas do banco de outras camadas
// - Retornar resultados ou erros para camada de serviços
// - Permitir fácil troca da implementação do banco (ex: substituir PostgreSQL por outro)
// - Facilitar testes unitários com mocks para o banco
//
// Boas práticas:
// - Não conter lógica de negócio, só persistência e recuperação
// - Tratar erros específicos do banco e logar se necessário

routes/routes.go

// routes/routes.go - definição das rotas da aplicação
//
// Responsabilidades:
// - Centralizar todas as rotas da API REST (GET, POST, PUT, DELETE)
// - Associar cada rota ao respectivo controller handler
// - Facilitar a manutenção e expansão dos endpoints
// - Permitir configuração de middlewares por rota (ex: autenticação, CORS)
//
// Boas práticas:
// - Agrupar rotas relacionadas por grupo ou prefixo (ex: /books)
// - Manter o roteamento separado da lógica para melhor organização

services/book_service.go

// services/book_service.go - lógica de negócio e regras da aplicação
//
// Responsabilidades:
// - Implementar regras específicas do domínio (ex: validações complexas, cálculos)
// - Coordenar chamadas ao repositório para persistência e consultas
// - Transformar dados entre camadas (DTOs, formatações)
// - Garantir integridade e consistência antes de enviar dados ao controller
// - Facilitar testes unitários focados em lógica, isolando acesso ao banco
//
// Boas práticas:
// - Evitar chamadas HTTP ou lógica de apresentação aqui
// - Manter serviços pequenos e coesos, dividir em múltiplos serviços se necessário
