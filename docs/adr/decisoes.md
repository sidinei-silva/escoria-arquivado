# Decisões (ADRs)

> Por que cada escolha foi feita. Uma entrada por decisão relevante de design ou arquitetura.
> Ordem cronológica inversa não é necessária — numeração sequencial basta.
>
> Formato: **contexto** (o que forçou a decisão) → **decisão** → **consequência**.

---

## ADR-001 — A Litania tem três camadas, e a Têmpera é o diferencial

**Data:** 21/07/2026

**Contexto.** O GDD tinha duas afirmações incompatíveis: o modelo de Fama seguia Albion (fama
por arma, nunca reseta), mas a Narrativa Jogada dizia que trocar de arma "zera a Fama". Uma
das duas estava errada.

**Decisão.** Três camadas independentes:

1. **Linha do Catador** — espinha; ganha de qualquer ação; nunca reseta.
2. **Fama por arma** — só com a arma equipada; nunca reseta.
3. **Têmpera** — multiplicador efêmero da arma equipada; sobe com uso ativo da mesma arma;
   **congela offline**; **reseta só na troca de arma**.

A frase "trocar zera a Fama" era vestígio de versão antiga e foi removida do GDD.

**Consequência.** Trocar de arma passa a ser **custo de oportunidade, não punição** — nada
acumulado se perde, mas abandona-se um multiplicador maduro. Isso preserva o modelo Albion e
o Pilar 2 ao mesmo tempo. A Têmpera vira o diferencial mecânico do projeto: recompensa
presença ativa sem punir ausência.

**Escopo.** PoC = Têmpera por arma. Pós-PoC = por loadout. Armadura existe na PoC mas ainda
não carrega Têmpera.

---

## ADR-002 — "Idle" aqui não é offline-first

**Data:** 21/07/2026

**Contexto.** A definição de mercado de idle game (Cookie Clicker, RuneScape-idle) é
offline-first: o jogo avança sozinho e o jogador volta para coletar. Uma proposta de mecânica
foi rejeitada por assumir essa definição.

**Decisão.** A Escória é um MMO de profundidade real para quem não tem tempo de input
contínuo. Recompensa **presença ativa** e microgerenciamento. Offline **congela, não avança**.
Idle significa "sem reflexo/APM exigido", não "joga sozinho enquanto você dorme".

**Consequência.** É regra de arquitetura, não só de design (ver `specs/000-constituicao.md`,
princípio 5). Habilita mecânicas como a Têmpera, que seriam punitivas num idle offline-first.
Toda decisão de mecânica passa por esse filtro.

---

## ADR-003 — Nenhum conteúdo de terceiros no projeto

**Data:** 21/07/2026

**Contexto.** O backlog tinha tarefas para importar dados do Albion (`ao-bin-dumps`), e o GDD
estava povoado de nomenclatura traduzida: recursos, mobs, estações de craft, nós de progressão.
Pior: a Especificação da PoC **afirmava** que nenhum nome de terceiros entrava no projeto,
enquanto a mesma página listava dezenas deles.

**Decisão.** Nenhum dado, nome, asset, tabela ou lore de terceiros entra. Nada de importar,
traduzir ou adaptar. Albion permanece como referência de **arquitetura de sistemas**
(fama por uso, árvore por linha, loadout como identidade) — nunca de conteúdo.

**Consequência.** As tarefas de importação foram reescritas para produção original. Todo nome
derivado virou placeholder `PH_*` rastreável (ver `gdd/60-producao/placeholders.md`). Números
de balanceamento permanecem como baseline a re-tunar — números não carregam identidade.

---

## ADR-005 — Monorepo com specs, GDD e código juntos

**Data:** 21/07/2026

**Contexto.** Em projeto anterior (Eras do Brasil), GDD e código viviam em repos separados.
Resultado observado: o GDD deixava de ser consultado durante o desenvolvimento. Distância
gerou negligência.

**Decisão.** Um repositório contendo `gdd/`, `data/`, `backend/`, `web/`, `docs/`.

---

## ADR-007 — Backlog organizado em fatias verticais pelo fluxo do jogador

**Data:** 21/07/2026

**Contexto.** O backlog tinha 58 tarefas organizadas por fase e camada, com 8 tarefas órfãs e
uma fase inchada (19 tarefas contra 2–4 nas demais). Havia também a preocupação de que ordenar
por jornada do usuário contrariasse boa arquitetura.

**Decisão.** Dez fatias verticais mapeando os 8 passos do tutorial. Cada fatia entrega algo
visível ou jogável. Fatia 0 (fundação) é a única exceção e deve ser mínima.

**Consequência.** Ordenar por jornada **é** boa prática (walking skeleton, vertical slice
architecture) — arquitetura-primeiro adia aprendizado e validação. Para um projeto solo,
motivação é restrição de engenharia legítima: projeto solo morre de abandono mais do que de
arquitetura ruim.

---

## ADR-010 — Migração manual, revertendo o ADR-009

**Data:** 21/07/2026

**Contexto.** O ADR-009 decidiu migrar o GDD por exportação nativa do Notion, por ser mais fiel
e mais barato. Sidinei reverteu com um argumento que o ADR-009 não considerou: **a exportação
transporta o drift junto, em silêncio.** Duas seções inteiras (`00 · Overview` e `20 · Mundo`)
nunca haviam sido revisadas, e o levantamento mostrou 45 páginas em vez das ~25 estimadas.

**Decisão.** Migração manual, página a página, em 8 etapas. A cada página, o conteúdo é
classificado em três baldes: **drift mecânico** (nome já decidido — corrigido em silêncio e
listado), **lacuna** (informação faltando ou contraditória — pergunta), **conteúdo novo** (algo
que o GDD nunca decidiu — pergunta). A IA não resolve cânone sozinha.

**Consequência.** Mais caro e mais lento, mas a migração vira revisão editorial em vez de
transporte.

---

## ADR-012 — Código em inglês, conteúdo de jogo em português

**Data:** 02/09/2026
**Status:** substitui parcialmente o ADR-004 e o invariante 3 do ADR-011

**Contexto.** O ADR-004 decidiu que identificadores de código seguiriam o lore em
português (`internal/litania/`, `GET /litania`). Ao escrever a Fatia 0, o código
nasceu em inglês: `internal/game/`, `internal/gamedata/`, `ZoneID`, `ActionType`,
e as chaves dos JSONs viraram `adjacent_zones`, `requires_objective`. O
invariante 3 do ADR-011 também previa domínios em português.

A prática divergiu da decisão sem ninguém perceber, e é melhor decidir do que
descobrir a divergência mais tarde.

**Decisão.** A fronteira passa a ser **o que o compilador vê versus o que o
jogador vê**, não o idioma do projeto.

Em inglês:
- pacotes, tipos, funções, variáveis e constantes Go
- chaves de JSON em `data/`
- valores que viram constante e entram em `switch`: `gather`, `chain`, `daily`,
  `entered_zone`
- IDs genéricos de conteúdo: `wood`, `rough_stone`, `tut_04_gathering`

Em português:
- IDs de entidade nomeada do lore: `a_ressaca`, `guia`, `portador_npc`
- todo texto que o jogador lê
- `subtext` e qualquer nota de direção
- toda a documentação: GDD, fluxos, sistemas, ADRs, mensagens de commit

Mensagens de erro do servidor são em português, porque hoje o único leitor sou eu.

**Consequência.** Um nome próprio do mundo mantém o nome em qualquer camada:
`a_ressaca` no JSON, no banco e na URL. O que o código enumera não depende de
acento nem de tradução.

Custo: dois vocabulários no mesmo arquivo. `zones[ZoneID("a_ressaca")]` mistura
os dois de propósito, e isso é o desenho, não descuido.

O ADR-004 continua valendo para o **conteúdo** — nada de "Destiny Board". O que
muda é só a camada de código.