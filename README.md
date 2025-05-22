
📘 Exercício: Expansão do CRUD de Livros com Categoria
🎯 Objetivo

Adicionar ao sistema uma funcionalidade de categorias de livros, incluindo:

    Novo model: Category

    Associação entre livros e categorias (relacionamento)

    Novas rotas para CRUD de categorias

    Possibilidade de listar livros por categoria

🧱 1. Criar o Model Category

    Campos:

        ID (uint, chave primária)

        Name (string, obrigatório, único)

        Relacionamento Books com []Book

    📝 Dica: use a tag gorm:"foreignKey:CategoryID" no model Book

🧰 2. Ajustar o Model Book

    Adicionar campo CategoryID (uint)

    Adicionar campo Category com gorm:"foreignKey:CategoryID"

🔧 3. Atualizar o banco de dados

    Fazer AutoMigrate com o novo model

    Se quiser, criar categorias manualmente via psql ou POST /categories

🔁 4. Criar os arquivos MVC para Categoria

    models/category.go

    repositories/category_repository.go

    services/category_service.go

    controllers/category_controller.go

Funções mínimas:

    GetCategories

    GetCategoryByID

    CreateCategory

    UpdateCategory

    DeleteCategory

🌐 5. Atualizar o roteamento

    Adicionar grupo de rotas /categories com os métodos:

        GET /categories

        GET /categories/:id

        POST /categories

        PUT /categories/:id

        DELETE /categories/:id

🔍 6. Endpoint extra (desafio)

Criar endpoint:

GET /categories/:id/books

    Retorna todos os livros pertencentes a uma determinada categoria.

✅ Critérios de Aceitação

    O código deve compilar e rodar sem erros

    As rotas de categoria devem funcionar via Insomnia/Postman

    Ao criar um livro, deve ser possível associá-lo a uma categoria

    Deve ser possível listar livros filtrando por categoria

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
