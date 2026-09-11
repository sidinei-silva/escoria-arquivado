---
tags: [indice]
---

# A Escória

MMORPG idle. Codinome `Outlander`. Servidor em Go, autoritativo.

> **Você é o que veste.** A build vem do conjunto de armas e armaduras
> equipadas. Tudo que você faz rende XP da atividade e Fama no centro. A arma que
> você usa por muito tempo retribui — e cobra quando você a troca.
>
> Um MMORPG full loot com a tensão inteira, jogado no tempo que a vida hoje
> permite.

## Por onde começar

| Quero… | Vá para |
|---|---|
| entender o jogo | [[gdd/README|gdd/README.md]] — os 8 passos, em ordem |
| saber como se trabalha aqui | [[docs/RECEITA|docs/RECEITA.md]] |
| ver o que falta construir | [[docs/BACKLOG|docs/BACKLOG.md]] |
| decidir algo de arquitetura | [[docs/estudos/README|docs/estudos/]] |
| saber a regra de IA | [[AGENTS|AGENTS.md]] |
| jogar fora uma ideia boa | [[IDEIAS|IDEIAS.md]] |

## Estado — 31/08/2026

Nenhuma linha de código escrita. O que existe:

- 8 regras numeradas: `R-JOG-01..05`, `R-ZON-01..03`
- 22 diálogos do tutorial, 8 passos, em `data/tutorial/dialogs.json`
- o Livro de Imersão
- o processo, o backlog e os estudos de arquitetura

O GDD completo está no [Notion](https://app.notion.com/p/3720b63b76f281bc89c4d679f9d8cd1f),
**congelado**. Migração preguiçosa: um sistema desce para `gdd/sistemas/` quando
a fatia dele chegar.

## Estrutura

```
AGENTS.md   CLAUDE.md   IDEIAS.md

gdd/        o jogo — fluxo, sistemas, regras, prosa
specs/      a entrega em curso. Uma por vez. Morre no fim.
docs/       processo, backlog, ADRs, estudos
data/       as instâncias: zonas, mobs, armas, diálogos
server/     Go — o servidor autoritativo
web/        o cliente
```

## Decisões que já valem

- **Estado autoritativo com dono único.** Uma goroutine é dona; comandos chegam
  por channel; o loop só toca memória. Cresce para partição por zona sem
  reescrever a lógica de jogo. → `docs/estudos/02`
- **O servidor é dono do tempo.** Atividade é marco absoluto, não contador. →
  `docs/estudos/03`
- **HTTP para pré-jogo, WebSocket para gameplay.** → `docs/estudos/04`
- **Sem event sourcing.** → `docs/estudos/05`

Nenhuma delas é ADR ainda. Viram ADR quando a fatia que as exige chegar, escritas
por mim.

## Próxima entrega

`S-001 · Criação de personagem` — Fatia 1, passo 1.
Regras: `R-JOG-01`, `R-JOG-02`, `R-JOG-03`. As três já existem.
