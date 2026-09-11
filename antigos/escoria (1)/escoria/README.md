# A Escória

MMORPG idle onde cada arma é um deus escondido.
"Você é o que veste." Subir de tier é negociar quanto de si você cede.

## Estrutura

| Caminho | O que é |
| --- | --- |
| `CLAUDE.md` | manual de operação — regras do projeto, lido em toda sessão de IA |
| `gdd/` | game design document: lore, mundo, narrativa, personagens |
| `specs/` | verdade técnica: o quê, contratos, tarefas |
| `data/` | conteúdo de jogo em JSON, consumido pelo servidor |
| `server/` | API em Go |
| `web/` | cliente React + Vite + TypeScript |
| `docs/` | decisões (ADRs) e diário de sessões |

## Por onde começar

1. `CLAUDE.md` — as regras
2. `specs/000-constituicao.md` — invariantes técnicos
3. `specs/BACKLOG.md` — as fatias em ordem
4. `specs/001-personagem/` — a fatia atual

## Stack

Go + chi + SQLite (modernc, sem CGO) + sqlc + golang-migrate ·
React + Vite + TypeScript + Zustand · conteúdo em JSON estático.
