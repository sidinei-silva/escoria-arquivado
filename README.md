---
tags: [índice]
---

# A Escória

MMORPG idle. Codinome `Outlander`. Servidor em Go, autoritativo.

> **Você é o que veste.** A build vem do conjunto de armas e armaduras
> equipadas. Tudo que você faz rende XP da atividade e Fama no centro. A arma
> que você usa por muito tempo retribui — e cobra quando você a troca.
>
> Um MMORPG full loot com a tensão inteira, jogado no tempo que a vida hoje
> permite.

## Por onde começar

| Quero…                            | Vá para                                           |
| --------------------------------- | ------------------------------------------------- |
| entender o jogo                   | [gdd/fluxos/](gdd/fluxos/) — o jogo em ordem      |
| consultar uma regra               | [gdd/sistemas/](gdd/sistemas/) — regras numeradas |
| saber como se trabalha aqui       | [docs/RECEITA.md](docs/RECEITA.md)                |
| saber por que algo é assim        | [docs/adr/](docs/adr/) — decisões, imutáveis      |
| ver o que falta                   | [docs/BACKLOG.md](docs/BACKLOG.md)                |
| guardar uma ideia que não é agora | [docs/ideias/](docs/ideias/)                      |
| saber a regra de IA               | [AGENTS.md](AGENTS.md)                            |

## Como o repositório está organizado

Separado por **tempo de vida**, não por assunto:

| Pasta               | O que é                                   | O que a muda                     |
| ------------------- | ----------------------------------------- | -------------------------------- |
| `gdd/fluxos/`       | o jogo em ordem, na visão do jogador      | cresce, uma fatia por vez        |
| `gdd/sistemas/`     | as regras numeradas                       | cresce por descoberta            |
| `data/`             | as instâncias: zonas, diálogos, objetivos | cresce                           |
| `docs/adr/`         | as decisões tomadas                       | **nunca** — só é substituída     |
| `docs/ideias/`      | o que talvez nunca aconteça               | é promovida ou apagada           |
| `specs/` ou issues  | a entrega em curso                        | **morre quando a entrega fecha** |
| `cmd/`, `internal/` | o servidor                                | sempre                           |
| `web/`              | o cliente                                 | sempre                           |

Só uma coisa morre: **spec**. Só uma é imutável: **ADR**. O resto cresce.

Cada pasta tem um `README.md` dizendo o que entra nela, e um `_GABARITO.md`
para copiar.

## A regra que sustenta tudo

**A regra mora em um lugar só: `gdd/sistemas/`.**

Fluxo cita `R-COL-01`. Spec cita `R-COL-01`. Teste nomeia `R-COL-01`. Nenhum
deles repete o texto nem o número. Se a duração da coleta aparecer em dois
arquivos, um dos dois vai ficar velho — e você não vai saber qual.

## Decisões que já valem

- **Estado autoritativo com dono único.** Uma goroutine é dona; comandos chegam
  por channel; o loop só toca memória.
- **O servidor é dono do tempo.** Atividade é marco absoluto (`inicio`, `fim`),
  nunca contador que decrementa. Tick perdido não atrasa o jogo, e o estado
  sobrevive a restart.
- **Offline congela.** Nenhuma mecânica avança com o jogador desconectado.
- **Conteúdo é dado, não código.** Adicionar zona, mob ou objetivo nunca exige
  recompilar.
- **Sem WebSocket na PoC.** Polling. Simples, suficiente, adiável.
- **A posição é do cliente.** O servidor descreve combate abstrato; a encenação
  é decoração. O mesmo servidor alimentaria um cliente de texto.

O porquê de cada uma está em [docs/adr/](docs/adr/). Uma spec nunca contradiz
uma ADR.

## Como uma fatia acontece

    fluxo → sistema → conteúdo → código+teste → lore → commit

Uma fatia entrega algo que o jogador **vê ou faz**, e não fecha até que as cinco
peças estejam no mesmo commit. Detalhe em [docs/RECEITA.md](docs/RECEITA.md).

## Estado

Nenhuma linha de código escrita.

O GDD completo está no Notion, **congelado** como referência. Migração
preguiçosa: um sistema desce para `gdd/sistemas/` quando a fatia que precisa
dele chegar.

**Próxima entrega — Fatia 0:** o servidor sobe, carrega `data/`, valida
referências e falha no boot com mensagem clara se alguma apontar para lugar
nenhum. Sem banco, sem rede, sem cliente.

## Sobre IA neste projeto

O código deste jogo é escrito à mão. IA explica conceito, compara abordagens,
revisa o que eu escrevi e ajuda a depurar — não gera arquivo, não gera função
para colar, não escreve lore.

O teste, antes de aceitar qualquer coisa: **eu consigo explicar isso amanhã, sem
olhar?**

Regras completas em [AGENTS.md](AGENTS.md).