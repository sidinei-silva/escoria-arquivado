# Inicializar projeto Go: criar estrutura de pastas (Standard Layout) + instalar dependências

Camada: Backend
Fase: Fase 0 — Setup
Notas: go mod init escoria. Criar estrutura: cmd/api/main.go, internal/gamedata/, internal/player/, internal/inventory/, internal/action/, internal/equipment/, internal/litania/, internal/db/ (gerado pelo sqlc), data/, migrations/. Instalar dependências: go get http://github.com/go-chi/chi/v5 http://modernc.org/sqlite http://github.com/golang-migrate/migrate/v4 http://github.com/sqlc-dev/sqlc.
Status: Backlog