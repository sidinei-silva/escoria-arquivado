# Criar GameDataLoader: carregar todos os JSONs no boot e expor structs tipadas em memória

Camada: Backend
Fase: Fase 0 — Setup
Notas: Criar internal/gamedata/loader.go (os.ReadFile + json.Unmarshal para cada arquivo de /data/), internal/gamedata/types.go (structs Zone, Item, Mob, Recipe, Ability, LitaniaNode), internal/gamedata/store.go (GetZone, GetItem, GetRecipe, GetAbilities, etc). Chamar gamedata.Load() no main.go antes de iniciar o servidor. Não é domínio DDD — é infraestrutura de conteúdo.
Status: Backlog