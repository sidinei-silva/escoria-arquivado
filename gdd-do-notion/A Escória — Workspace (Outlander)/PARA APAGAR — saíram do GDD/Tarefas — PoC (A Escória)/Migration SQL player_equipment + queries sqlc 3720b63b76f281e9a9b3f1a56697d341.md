# Migration SQL player_equipment + queries sqlc

Camada: Backend
Notas: Criar migrations/003_create_player_equipment.up.sql (slots: weapon, head, chest, boots, bag) e migrations/004_create_skill_queue.up.sql. Escrever internal/equipment/query.sql: GetEquipment, EquipItem, UnequipSlot, GetSkillQueue, SaveSkillQueue. Handler em internal/equipment/handler.go, lógica em internal/equipment/service.go.
Status: Backlog