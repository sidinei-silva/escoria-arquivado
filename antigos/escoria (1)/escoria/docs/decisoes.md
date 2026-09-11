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

## ADR-004 — Identificadores de código seguem o lore em português

**Data:** 21/07/2026

**Contexto.** Com a renomeação de "Destiny Board" para "A Litania", três referências ficaram
em código: `progression/destiny_board.json`, `GET /destiny-board`, e o pacote
`internal/destiny/`. Manter código em inglês com lore em português é prática comum e legítima.

**Decisão.** Código segue o lore: `internal/litania/`, `GET /litania`, `litania.json`,
`player_litania_progress`. Módulo Go: `escoria`.

**Consequência.** Um vocabulário só no projeto inteiro. Custo: identificadores em português
com risco de acento — mitigado usando apenas ASCII em nomes de arquivo e símbolo.

---

## ADR-005 — Monorepo com specs, GDD e código juntos

**Data:** 21/07/2026

**Contexto.** Em projeto anterior (Eras do Brasil), GDD e código viviam em repos separados.
Resultado observado: o GDD deixava de ser consultado durante o desenvolvimento. Distância
gerou negligência.

**Decisão.** Um repositório contendo `gdd/`, `specs/`, `data/`, `server/`, `web/`, `docs/`.
O vault do Obsidian é a raiz do repo.

**Consequência.** Spec e código mudam no mesmo commit — regra cardinal do SDD, impossível com
repos separados. Contexto de IA é controlado apontando caminhos, não por fronteira de repo.
Se o GDD crescer demais, dá para separar depois; o inverso é mais caro.

---

## ADR-006 — Specs param no contrato

**Data:** 21/07/2026

**Contexto.** No SDD com agente escrevendo código, specs são exaustivas — não deixam decisão
em aberto. Aqui o implementador é humano, e um dos objetivos declarados do projeto é treinar
engenharia.

**Decisão.** Specs contêm: o quê, por quê, critérios de aceite, contratos de dados e API,
ligações com o GDD, escopo negativo. **Não** contêm: estrutura interna de funções, algoritmos,
tratamento de erro, onde quebrar arquivo.

**Consequência.** O implementador projeta, não transcreve. Efeito colateral valioso: specs no
nível de contrato quase não driftam, enquanto specs no nível de implementação viram mentira no
primeiro dia em que o código diverge.

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

## ADR-008 — Specs são direção de arte, não engenharia

**Data:** 21/07/2026

**Contexto.** O ADR-006 estabeleceu specs no nível de contrato, mas a primeira spec piloto
ainda continha schema SQL, shapes de endpoint em JSON e tipos TypeScript. Isso é código —
e código escrito pela IA remove exatamente a parte que Sidinei quer exercitar.

**Decisão.** A IA atua como **diretora de arte**: descreve o que precisa existir, por quê, e
o que precisa ser verdade no fim. Sidinei atua como **engenheiro-chefe**: projeta todas as
estruturas. Specs descrevem em prosa e **fazem perguntas** em vez de dar respostas.

A IA não escreve schema, shape de endpoint, tipo de cliente, assinatura de função nem
algoritmo — nem como exemplo.

**Exceção única.** A forma dos arquivos em `data/`. A IA os produz nos chats de conteúdo e o
código os consome; forma não acordada significa arquivo que não carrega. Sidinei define a
forma, a IA segue, e fica registrada na `dados.md` da fatia.

**Consequência.** `contracts.md` virou `dados.md` e mudou de natureza: de especificação de
shapes para descrição de necessidades + perguntas em aberto. Critérios de aceite ganham peso —
com as estruturas em aberto, eles passam a ser o único contrato verificável.

---

## ADR-009 — Migração do GDD via exportação nativa, não transcrição

**Data:** 21/07/2026

**Contexto.** O GDD tem cerca de 25 páginas no Notion. A opção óbvia era a IA buscar cada uma e
reescrever em Markdown.

**Decisão.** Usar a exportação nativa do Notion (Markdown & CSV, com subpáginas) e limpar com
script. A IA não transcreve conteúdo que já existe em formato exportável.

**Consequência.** Mais fiel (tabelas, blocos de código e mermaid preservados), sem risco de erro
de transcrição, e muito mais barato. O valor da IA fica onde ela é insubstituível: definir a
estrutura de destino, escrever o script de limpeza e corrigir drift depois que o conteúdo
aterrissar.

**Efeito colateral.** O Notion vira arquivo morto depois da migração. Manter os dois vivos
repetiria o erro dos placeholders — duas fontes da mesma verdade divergem.

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
transporte. O `scripts/limpar-export-notion.py` permanece no repo como ferramenta auxiliar.

**Nota.** O ADR-009 não estava errado nos fatos — estava errado no critério. Otimizou custo de
transporte quando o gargalo real era qualidade de conteúdo.
