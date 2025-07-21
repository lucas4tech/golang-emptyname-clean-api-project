## 🏠 Visão Geral

Este projeto é uma API moderna e eficiente para gerenciar usuários, produtos e pedidos, seguindo princípios de Clean Architecture, Domain-Driven Design (DDD), boas práticas de Go e desenvolvido com TDD (Test Driven Development) 🧪.

---

## 🛠️ Decisões Técnicas

### 1. 🏗️ Arquitetura e Organização
- **Domain-Driven Design (DDD):**
  - Entidades, agregados e value objects modelam o domínio de forma explícita.
  - Serviços de domínio encapsulam regras de negócio complexas.
- **Clean Architecture:**
  - Separação clara entre domínio, usecases, infraestrutura e interfaces (API).
  - Facilita testes, manutenção e evolução do sistema.
- **Pacotes:**
  - `internal/domain/entity`: Entidades do domínio (User, Product, OrderItem).
  - `internal/domain/aggregate`: Agregados (Order).
  - `internal/domain/value_object`: Tipos de valor imutáveis (UUID, Money, Email).
  - `internal/usecase`: Casos de uso (application service layer).
  - `internal/infra/persistence/gorm`: Models e repositórios usando GORM.
  - `internal/infra/api/graphql`: Schema e resolvers GraphQL.

### 2. 💾 Persistência e Banco de Dados
- **GORM + SQLite:**
  - GORM facilita o mapeamento objeto-relacional e preload de relacionamentos.
  - SQLite é usado para simplicidade local, mas a arquitetura permite trocar facilmente por outros bancos.
- **Relacionamentos explícitos:**
  - Uso de `Preload` e queries manuais para garantir que os itens e produtos de cada pedido sejam sempre retornados corretamente.

### 3. 🔗 GraphQL
- **graphql-go:**
  - Schema e resolvers definidos explicitamente para garantir flexibilidade e performance.
  - Resolvers suportam múltiplos tipos de resposta (ex: ListOrderResponse e CreateOrderResponse) para queries e mutations.
- **Retorno de dados aninhados:**
  - Queries retornam dados completos e aninhados (usuário, itens, produto de cada item, etc).

### 4. 🧪 Testes (TDD)
- **Testes unitários com testify:**
  - Cobrem entidades, agregados e value objects.
  - Garantem regras de negócio, validações e operações matemáticas.
- **Cobertura de testes:**
  - Comando `make coverage` e integração com GitHub Actions.
- **Desenvolvimento orientado a testes (TDD):**
  - As regras de negócio e entidades foram implementadas sempre acompanhadas de testes automatizados.

### 5. 🐳 Docker e Deploy
- **Dockerfile multi-stage:**
  - Imagem enxuta e segura, build separado do runtime.
- **docker-compose:**
  - Facilita rodar a aplicação localmente, com persistência do banco de dados.

### 6. ⚙️ CI/CD
- **GitHub Actions:**
  - Workflow automatizado para build, testes e cobertura a cada push/pull request.

---

## ▶️ Como rodar localmente

1. **Build e subir com Docker Compose:**
   ```bash
   docker-compose up --build
   ```
2. **Rodar testes:**
   ```bash
   make test
   ```
3. **Cobertura de testes:**
   ```bash
   make coverage
   ```
4. **Acessar a API:**
   - [http://localhost:8080](http://localhost:8080)

---

## 💡 Observações
- O projeto é facilmente adaptável para outros bancos relacionais.
- O uso de value objects e entidades imutáveis garante integridade dos dados.
- O GraphQL permite consultas flexíveis e aninhadas, ideais para frontend moderno.
- Desenvolvido com TDD para máxima confiabilidade e facilidade de refatoração.

---

## 🤝 Contribuição
Pull requests são bem-vindos! Siga o padrão dos testes e mantenha a cobertura alta. 💚

