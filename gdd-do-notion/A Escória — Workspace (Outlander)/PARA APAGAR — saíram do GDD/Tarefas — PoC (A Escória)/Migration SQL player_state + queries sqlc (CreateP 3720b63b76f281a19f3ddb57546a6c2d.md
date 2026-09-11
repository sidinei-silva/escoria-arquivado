# Migration SQL player_state + queries sqlc (CreatePlayer, GetPlayer, UpdateZone...)

Camada: Backend
Fase: Fase 1 — Personagem
Notas: Criar migrations/001_create_player_state.up.sql. Escrever internal/player/query.sql com: CreatePlayer, GetPlayerByID, UpdatePlayerZone, UpdateCurrentAction, SetTutorialCompleted. Rodar sqlc generate → código vai para internal/db/. Handler em internal/player/handler.go, lógica em internal/player/service.go.
Status: Backlog