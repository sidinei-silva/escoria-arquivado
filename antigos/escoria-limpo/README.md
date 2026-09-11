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
| entender o jogo | [gdd/README.md](gdd/README.md) — os 8 passos, em ordem |
| saber como se trabalha aqui | [docs/RECEITA.md](docs/RECEITA.md) |
| ver o que falta construir | [docs/BACKLOG.md](docs/BACKLOG.md) |
| decidir algo de arquitetura | [docs/estudos/](docs/estudos/README.md) |
| saber a regra de IA | [AGENTS.md](AGENTS.md) |
| jogar fora uma ideia boa | [IDEIAS.md](IDEIAS.md) |

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
internal/   Go — domínios do jogo
cmd/        Go — binário do servidor
web/        o cliente (React + Vite)
```

## Decisões que já valem

- **Estado autoritativo com dono único.** Uma goroutine é dona; comandos chegam
  por channel; o loop só toca memória. Cresce para partição por zona sem
  reescrever a lógica de jogo. → `docs/estudos/02`
- **O servidor é dono do tempo.** Atividade é marco absoluto, não contador. →
  `docs/estudos/03`
- **HTTP para pré-jogo, WebSocket para gameplay.** → `docs/estudos/04`
- **Sem event sourcing.** → `docs/estudos/05`

Stack e invariantes estão na **ADR-011**. Os estudos discutem alternativas —
a ADR vence.

## Próxima entrega

**Fatia 0** — servidor sobe, carrega `data/`, valida referências. Sem spec.

Depois, `specs/S-001-criacao-personagem/` — escrita em julho, revisar antes de
implementar. Regras `R-JOG-01`, `R-JOG-02`, `R-JOG-03`, todas já existentes.
