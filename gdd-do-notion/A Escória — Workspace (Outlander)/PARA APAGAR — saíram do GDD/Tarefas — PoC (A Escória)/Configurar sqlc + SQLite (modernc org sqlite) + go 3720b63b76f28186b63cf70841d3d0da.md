# Configurar sqlc + SQLite (modernc.org/sqlite) + golang-migrate

Camada: Backend
Fase: Fase 0 — Setup
Notas: Criar sqlc.yaml apontando queries para cada internal/*/query.sql e output para internal/db/. Configurar golang-migrate com pasta migrations/. O código em internal/db/ é gerado — nunca editar manualmente. Rodar sqlc generate após cada query.sql novo.
Status: Backlog