# 000 — Constituição Técnica

> Invariantes do projeto. Uma spec **nunca** contradiz este documento.
> Mudar algo aqui é uma decisão de arquitetura e exige um ADR em `docs/decisoes.md`.

## Stack

| Camada | Escolha | Não-negociável nesta PoC |
| --- | --- | --- |
| API | Go + chi | sim |
| Banco | SQLite (`modernc.org/sqlite`) | sim — sem CGO |
| Acesso a dados | sqlc (SQL escrito à mão, tipos gerados) | sim — **sem ORM** |
| Migrations | golang-migrate | sim |
| Cliente | React + Vite + TypeScript | sim |
| Estado (cliente) | Zustand | sim |
| Conteúdo de jogo | JSON estático em `data/` | sim |

Módulo Go: `escoria`.

## Princípios de arquitetura

**1. Conteúdo é dado, não código.**
Armas, mobs, zonas, receitas, nós da Litania, diálogos — tudo vive em JSON sob `data/`,
carregado no boot e servido em memória. Adicionar conteúdo **nunca** exige recompilar lógica.

**2. `gamedata` é infraestrutura, não domínio.**
O carregador de JSON não tem regra de negócio. Ele lê, tipa e expõe. Quem decide o que fazer
com o dado são os domínios.

**3. Domínios por substantivo do jogo, em português.**
`internal/player/`, `internal/inventory/`, `internal/action/`, `internal/equipment/`,
`internal/litania/`. Identificador de código segue o lore (ver `CLAUDE.md`).

**4. O servidor é a autoridade do tempo.**
Toda ação tem `started_at` + `duration_ms` **no servidor**. O cliente não decide quando um
ciclo termina — ele pergunta. Cliente adiantando relógio não deve conseguir nada.

**5. Offline congela.**
Nenhuma mecânica avança com o jogador desconectado. Quando ele volta, o estado é o que ele
deixou. Isso decorre da definição de idle do projeto (`CLAUDE.md`) e é **regra de arquitetura**,
não preferência de design.

**6. Sem WebSocket na PoC.**
Polling do cliente (`GET /actions/current`). Simples, suficiente, e adiável.

## Convenções de API

- REST sobre JSON. Rotas em português, seguindo o lore: `/litania`, `/viagem`, `/inventario`.
- Corpo e campos em **snake_case**.
- Erro: HTTP status adequado + corpo `{ "erro": "codigo_legivel", "mensagem": "..." }`.
- Toda rota que muda estado é `POST`. Toda leitura é `GET`.
- Identificação do jogador via sessão mock na PoC (`POST /auth/mock`). Autenticação real é
  pós-PoC e **não** deve influenciar o desenho das outras rotas.

## Convenções de dados

- Migrations numeradas e imutáveis depois de aplicadas: `NNN_descricao.up.sql` / `.down.sql`.
- Chaves primárias: `id` inteiro autoincremento, salvo quando houver razão explícita.
- Toda tabela de progresso do jogador começa com `player_`.
- Timestamps em UTC, ISO-8601.
- IDs de conteúdo (armas, mobs, zonas) são **strings estáveis** definidas no JSON, nunca
  inteiros de banco. Conteúdo é dado; dado tem nome.

## Definição de pronto

Uma tarefa está pronta quando:

1. O comportamento descrito nos **critérios de aceite** da spec acontece de ponta a ponta.
2. O jogador consegue ver ou fazer aquilo pela interface (fatias 1+; a Fatia 0 é exceção).
3. A spec continua verdadeira. Se o código divergiu, **a spec é atualizada no mesmo commit**.
4. Decisão relevante tomada no caminho virou linha em `docs/decisoes.md`.

Não há exigência de cobertura de testes nesta PoC. Há exigência de que o critério de aceite
seja verificável à mão.

## Fluxo de trabalho (SDD)

```
gdd/  ──────────►  specs/NNN-fatia/spec.md      (o quê + por quê + aceite)
                          │
                          ├─► dados.md          (o que precisa existir, em prosa)
                          │
                          └─► tasks.md          (quebra executável)
                                   │
                                   ▼
                          Sidinei implementa
                                   │
                                   ▼
                   docs/decisoes.md + docs/diario.md   (write-back)
```

A IA produz `spec.md`, `dados.md`, `tasks.md`. Sidinei produz o código — todo ele.

**O que uma spec contém:** o quê, por quê, critérios de aceite verificáveis, descrição em
prosa do que precisa existir, as perguntas que a implementação precisa responder, ligações
com o GDD, escopo negativo.

**O que uma spec NÃO contém:** schema SQL, shape de endpoint, tipos de cliente, assinatura de
função, algoritmo, tratamento de erro. A IA não escreve código — nem de exemplo.

**Exceção única:** a forma dos arquivos em `data/`, porque são produzidos pela IA e consumidos
pelo código. Sidinei define a forma; a IA segue.

## Ordem de execução

Fatias verticais seguindo o fluxo do jogador — ver `specs/BACKLOG.md`.

Cada fatia entrega algo **visível ou jogável**. A única exceção é a Fatia 0 (fundação), que
deve ser mantida no mínimo absoluto. Se uma fatia não tem entrega visível, ela não é fatia:
é arquitetura disfarçada, e provavelmente está fora de ordem.
