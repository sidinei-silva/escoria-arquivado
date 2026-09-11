# Estudos — Arquitetura do servidor

> **Isto não é decisão. É estudo.**
>
> Nada aqui é ADR. ADR é decisão tomada, datada e imutável (`docs/RECEITA.md`).
> Estes documentos são o material que me deixa **tomar** a decisão quando a fatia
> chegar — com o trade-off na mesa em vez de copiado.
>
> Quando eu decidir de fato, escrevo a ADR com as minhas palavras, em três
> parágrafos. O estudo é insumo da ADR, nunca a ADR.

## Como usar

**Não leia isto de ponta a ponta.** São documentos de consulta. Ler tudo antes de
começar é estudar em vez de codar — o padrão que estou tentando quebrar.

Leia o estudo da fatia, quando a fatia chegar.

## O código nestes documentos

Os trechos de código são **ilustração de forma**, não implementação. São curtos
de propósito, omitem tratamento de erro e casos de borda, e nenhum deles compila
sozinho.

Não copiar. Se eu me pegar colando daqui, o documento falhou no propósito dele —
e eu quebrei o [AGENTS.md](../../AGENTS.md).

## Mapa: qual estudo em qual fatia

| Fatia | O que entrega | Estudo |
|---|---|---|
| 0 — Fundação | servidor sobe, carrega conteúdo | [01-modelo-mental](01-modelo-mental.md) |
| 1 — Personagem | criar personagem, ver A Ressaca | [01-modelo-mental](01-modelo-mental.md), [04-rede-http-e-websocket](04-rede-http-e-websocket.md) |
| 2 — Primeira arma | equipar, loadout muda | — |
| 3 — Primeiro combate | matar o primeiro Mito, Fama, Têmpera simples | [02-ownership-e-concorrencia](02-ownership-e-concorrencia.md), [03-tempo-e-atividades](03-tempo-e-atividades.md) |
| 4 — Coleta | coletar em loop, inventário, banco entra | [03-tempo-e-atividades](03-tempo-e-atividades.md), [05-persistencia](05-persistencia.md) |
| 5 — Bigorna | viajar, refinar, craftar | [03-tempo-e-atividades](03-tempo-e-atividades.md) |
| 6 — Outro portador | encontro com o NPC-Portador | — |
| 7 — Zonas T2 | viajar e agir em tier 2 | [06-zonas-presenca-e-pvp](06-zonas-presenca-e-pvp.md) |
| 8 — Segunda arma | combate T2, fila de habilidades, Têmpera completa | [02-ownership-e-concorrencia](02-ownership-e-concorrencia.md) |
| 9 — Litania e Saída | ver a Litania, atravessar A Encruzilhada | [05-persistencia](05-persistencia.md) |
| pós-PoC | multiplayer, gank, PvP | [06-zonas-presenca-e-pvp](06-zonas-presenca-e-pvp.md) |

Fatias 0 a 3 rodam com estado em memória e JSON no disco. O banco entra na
Fatia 4, com o inventário — a primeira coisa que dói perder.

## ADRs que vou precisar escrever

Não escrever agora. Cada uma nasce na fatia que a exige, decidida por mim:

- **ADR-001 — ownership do estado autoritativo** → Fatia 3
- **ADR-002 — o servidor é dono do tempo das atividades** → Fatia 3
- **ADR-003 — quando persistir** → Fatia 4
- **ADR-004 — HTTP para pré-jogo, WebSocket para gameplay** → Fatia 3 ou 4
- **ADR-005 — partição por zona** → pós-PoC, só quando o multiplayer chegar

## Os documentos

1. [01-modelo-mental](01-modelo-mental.md) — command, state, event. O que o servidor é.
2. [02-ownership-e-concorrencia](02-ownership-e-concorrencia.md) — quem é dono do estado. Mutex, channel, actor.
3. [03-tempo-e-atividades](03-tempo-e-atividades.md) — atividade não é instantânea. Game loop e tick.
4. [04-rede-http-e-websocket](04-rede-http-e-websocket.md) — qual porta serve o quê, e por quê.
5. [05-persistencia](05-persistencia.md) — o banco não dirige o game loop.
6. [06-zonas-presenca-e-pvp](06-zonas-presenca-e-pvp.md) — o desenho que o jogo final quer.

## Procedência

Destilado de dois estudos com ChatGPT (26/08/2026), reorganizado e com a decisão
de arquitetura revista. Os chats originais ficam arquivados fora do repo. O que
sobreviveu é o que passa no teste do [AGENTS.md](../../AGENTS.md): consigo explicar
amanhã, sem olhar.
