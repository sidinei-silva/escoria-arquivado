# Migration SQL player_litania_progress + queries sqlc

Camada: Backend
Notas: Criar migrations/005_create_player_litania_progress.up.sql. Escrever internal/litania/query.sql: GetNodeProgress, UpsertFame, GetAllProgressByPlayer, UnlockNode. Handler em internal/litania/handler.go, lógica em internal/litania/service.go (inclui regras de distribuição de fama).
Status: Backlog