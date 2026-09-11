---
tags: [estudo, arquitetura, persistencia, postgres]
fatia: 4, 9
---

# 05 — Persistência

## O princípio

> **O banco não dirige o game loop.**

Vindo de web, o instinto é o oposto: o banco é a verdade, a request lê e escreve,
o processo é sem estado. Aqui não.

Aqui a verdade é o estado em memória, possuído pelo loop
([[02-ownership-e-concorrencia]]). O Postgres é **estado durável**: o que
sobrevive a restart. Consequência, não fonte.

O motivo é dureza, não gosto: se cada tick fizesse I/O, o tick passaria a durar o
que o banco demora. Uma query lenta viraria lag de jogo para todo mundo. É a
regra do documento 02 aplicada — **o loop só toca memória**.

## Quando persistir

Quatro estratégias, com o que cada uma custa:

| Estratégia | Perda máxima | Custo | Veredito |
|---|---|---|---|
| A cada tick | ~zero | inviável — I/O no loop | não |
| A cada evento relevante | ~zero | alto, mas fora do loop | talvez, seletivo |
| Periódico (snapshot) | o intervalo | baixo | **base** |
| Só no logout/shutdown | tudo, num crash | mínimo | não sozinho |

O desenho que faz sentido é **misto**:

- **snapshot periódico** como base (a cada N segundos, o que sujou desde o
  último)
- **escrita imediata** para o que não pode perder: conclusão de craft, avanço de
  nó da Litania, saída da ilha
- **flush no shutdown**, sempre

O que decide "não pode perder" é **quanto dói repetir**. Perder 30 segundos de
ciclo de coleta é irritante. Perder o craft de um item T2 é o jogador ir embora.
Essa é uma decisão de design, não técnica — e portanto vira regra no GDD, não só
ADR.

## Fora do loop, sempre

```
Game Loop  ──── snapshot ────→  channel  ────→  Persistence Worker  ──→  Postgres
(memória)                                        (goroutine própria)
```

O loop nunca espera o banco. Ele monta o que mudou, joga no channel, e segue.

E as perguntas que isso levanta — e que eu tenho que responder, não herdar:

- se o worker estiver atrasado e o channel encher, o loop bloqueia ou descarta?
- o snapshot é cópia ou referência? (referência = o worker lê enquanto o loop
  escreve = race — o `-race` pega isso)
- duas versões do mesmo jogador na fila: escreve as duas ou só a última?

A terceira tem resposta natural: para estado de jogador, **só a última importa**.
Um mapa `jogadorID → estado sujo`, drenado a cada intervalo, resolve.

## O que salvar

| Categoria | Vai para o banco? |
|---|---|
| Conta, credenciais | sim — é a base da conta |
| Personagem: nome, zona atual, passo do tutorial | sim |
| Inventário, equipamento | sim |
| Fama por nó, progresso da Litania | sim |
| Têmpera da arma equipada | sim — congela offline, então é estado |
| Atividade em curso (`TerminaEm`) | sim — marco absoluto sobrevive a restart |
| Conteúdo: zonas, receitas, mobs | **não** — vive em `data/`, é o repositório |
| Eventos emitidos | não — ver abaixo |
| Conexão, sessão em memória | não |

A linha da atividade em curso só funciona por causa da decisão de
[[03-tempo-e-atividades]]: marco absoluto é salvável, contador relativo não é.

## Persistência não é event sourcing

Tentação real: já que existem eventos, por que não gravar todos e reconstruir o
estado a partir deles?

Porque não é o mesmo problema. Event sourcing é a fonte da verdade ser o log, e
todo estado ser projeção. Isso traz versionamento de evento, replay, snapshots
para não replayar tudo, e migração de esquema de evento. É muito peso para um
projeto de uma pessoa com 6 a 8 horas por semana.

O evento aqui existe para **avisar o cliente**. Estado é estado, salvo como
estado.

Se um dia eu quiser log de auditoria — anti-cheat, investigar gank — isso é uma
tabela append-only ao lado, não uma mudança de arquitetura.

## Boot

```
1. carrega data/ (JSON) → conteúdo tipado, imutável
2. valida referências: toda receita aponta para recurso existente?
3. abre o Postgres, roda migrations
4. carrega estado durável do que precisa estar quente
5. sobe o game loop
6. só então abre HTTP e WebSocket
```

A ordem importa: aceitar conexão antes de o estado existir é aceitar comando que
não tem onde ser aplicado.

O passo 2 é barato e evita muita hora perdida. Um typo em `recipes.json` deve
falhar no boot com mensagem clara, não virar `nil pointer` no meio de um craft
três dias depois. Isso é a Fatia 0 e cabe nela.

## Não comece pelo banco

Recomendação prática, contra o que o `BACKLOG` sugere na Fatia 0: **adiar
Postgres**.

As Fatias 1 a 3 funcionam com estado em memória e JSON no disco. Isso corta
migrations, driver, conexão e schema das primeiras noites — que é justamente
quando a energia é mais escassa e o gosto de ver algo funcionando importa mais.

O banco entra na Fatia 4, com o inventário, que é a primeira coisa que dói
perder.

Contra-argumento honesto: adiar significa migrar depois, e migrar dá trabalho.
Mas se a fronteira estiver limpa — o loop fala com uma interface de repositório,
não com SQL — a troca é uma implementação nova, não uma reescrita. E desenhar
essa fronteira é aprendizado bom.

## Simplificação didática

- migrations: ferramenta e estratégia não escolhidas
- transação: um snapshot que toca inventário e Fama junto deveria ser atômico?
- pool de conexão e tuning ficam para quando houver número
- backup e recuperação não existem na PoC

## O que vira ADR-003

O banco é estado durável, não fonte da verdade. Escrita acontece fora do loop.
Snapshot periódico como base, escrita imediata para o que dói perder. Sem event
sourcing.
