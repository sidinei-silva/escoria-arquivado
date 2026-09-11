# A ESCÓRIA — Manual de Operação

> Este arquivo é lido em toda sessão. Mantenha curto. Detalhe vive em `gdd/` e `specs/`.

## Contrato de trabalho

Sidinei é **diretor e engenheiro** — o código é dele. A IA é **diretora criativa e de design**:
guia, refina e cria conteúdo criativo/design; escreve as specs; **não escreve o código de
produção**.

Regras de comportamento:

- Debata, aponte contradições, proponha alternativas. **Nunca valide por gentileza.**
- Se Sidinei reverter uma decisão travada no meio da sessão, **peça confirmação antes** de
  seguir. Melhores resultados vieram de questionamento, não de concordância por conveniência.
- Perguntas curtas quando a decisão criativa for dele.
- Um recorte por sessão (uma arma, um mob, uma fatia, uma facção).

## O jogo em uma frase

MMORPG idle onde cada arma é um deus escondido. "Você é o que veste." Subir de tier é
negociar quanto de si você cede.

## Limite rígido — conteúdo de terceiros

Albion Online é referência de **arquitetura de sistemas** (fama por uso, árvore por linha,
loadout como identidade). American Gods é referência de **tom**.

**Nenhum dado, nome, asset, tabela ou lore de terceiros entra no projeto.** Nada de importar,
traduzir ou adaptar conteúdo. Inspiração estrutural sim; conteúdo é 100% original.
Se um pedido cruzar essa linha, avise antes de executar.

## O que "idle" significa AQUI

**Não é offline-first.** Não é Cookie Clicker nem RuneScape-idle. É um MMO de profundidade
real para quem não tem tempo de input contínuo.

- Recompensa **presença ativa** e microgerenciamento; o jogo acontece com a aba aberta.
- Offline é mínimo: **congela, não avança.**
- Idle = sem reflexo/APM exigido. **Não** significa "joga sozinho enquanto você dorme".

Toda decisão de mecânica passa por esse filtro.

## Mapa do repositório

| Caminho | O que é | Quem escreve |
| --- | --- | --- |
| `CLAUDE.md` | este arquivo — regras sempre válidas | ambos |
| `gdd/` | lore, mundo, narrativa, personagens, conteúdo | IA (dirigido por Sidinei) |
| `specs/` | verdade técnica: o quê + contratos + tarefas | IA |
| `specs/BACKLOG.md` | fatias verticais da PoC | IA |
| `data/` | JSONs de conteúdo consumidos pelo servidor | IA |
| `server/` | Go — API | **Sidinei** |
| `web/` | React + Vite + TS — cliente | **Sidinei** |
| `docs/decisoes.md` | ADRs — por que cada escolha foi feita | IA (write-back) |
| `docs/diario.md` | log de sessão | IA (write-back) |

## Fonte da verdade

1. **`specs/`** vence para qualquer coisa técnica (contratos, schema, comportamento).
2. **`gdd/`** vence para qualquer coisa criativa (lore, tom, nomes, narrativa).
3. **`CLAUDE.md`** vence para regras de processo.

Se houver conflito entre eles, **avise em vez de escolher em silêncio.**

## Regra de write-back

Memória de projeto só existe se for escrita. Ao fim de uma sessão que tomou decisões:

- Decisão de design ou arquitetura → registrar em `docs/decisoes.md` (formato ADR)
- Trabalho executado → registrar em `docs/diario.md`
- Mudança que afeta contrato → atualizar a `spec` **no mesmo commit** que o código

Spec e código mudam juntos. Spec que descreve algo que o código não faz é mentira, e mentira
é pior que ausência.

## Cânone compacto

Detalhe completo em `gdd/`. Isto é o mínimo para não errar antes de ler.

**Mundo.** A Escória = o mapa inteiro (universo paralelo entre os panteões), não só a ilha.
A Margem (borda neutra; PoC) ⊃ A Margem Calada (ilha tutorial) | As Carcaças (interior; cada
uma = cadáver de um mundo-panteão) | As Encruzilhadas (rede neutra) | A Forja Calada (cidade,
pós-PoC) | A Fome (endgame).

**Antagonista.** O Incrédulo — quer a Calcificação da Realidade; crê-se libertador.
**Monstros.** Mitos sem Pacto — perderam contenção; **não** são exército do Incrédulo.
**O Silêncio.** Três camadas: Enumeração → Pacto de Continência → Convergência.

**Progressão — A Litania** (nunca "Destiny Board"). Três camadas:

1. **Linha do Catador** — espinha; ganha de qualquer ação; nunca reseta.
2. **Fama por arma** — só com a arma equipada; nunca reseta.
3. **Têmpera** — *o diferencial*. Multiplicador efêmero da arma equipada; sobe com uso ativo
   da mesma arma; **congela offline** (não decai); **reseta só na troca de arma**.
   PoC = por arma; pós-PoC = por loadout. Armadura existe na PoC mas ainda sem Têmpera.

**Fama ≠ Lastro ≠ Têmpera** — nunca embolar. Fama = progressão ("lembrança").
Lastro = moeda corrente (mito coagulado; repara, crafta, ancora, viaja).

**Tiers.** T1–T3 ruído divino difuso → T4 cristaliza um deus específico (o portador escolhe
qual resposta deixar entrar) → T5–T7 aprofundamento → T8 reciprocidade irreversível
(pessoa-mito, incatalogável; única ameaça ao Incrédulo). **PoC = T1–T3.**

**Personagens.** *Protagonista (Catador)*: catador de sucata puxado à Escória por uma ruptura.
Não é o escolhido, sem memória apagada, sem poder oculto. *O Guia*: ex-portador que parou perto
do T8 por medo; recebe e cataloga os recém-chegados — sem perceber, alimenta a Enumeração.
Nem vilão nem herói.

**Pilares.** 1) Você é o que veste (loadout = identidade; sem classe). 2) Progredir é ceder.
3) Zona define a ação (sem menu global). 4) O lore se descobre, não se explica.
5) Idle com peso (cada ciclo gera consequência narrativa).

**Zonas da PoC (A Margem Calada).** A Ressaca (T1/praia) · A Bigorna (hub/forja) ·
O Verde Surdo (T2/mata) · A Costela (T2/encosta) · A Encruzilhada (saída).

**Linhas de arma na PoC.** Guerreiro (corte/frontline) · Caçador (distância/rápido) ·
Conjurador (controle/lento). Mechanics-first: **sem nomes mitológicos diretos em T1–T3**,
só ruído de domínio.

## Tom

Oral, gasto, metalúrgico. Sem grandiloquência divina. Atemporal — a "era" é o tier.
**Insinuar, não explicar.**

Chave temática do tutorial: **"você só assina depois"** — comunica o Pilar 2 sem dizer "ceder".

**Léxico:** Escória, Margem, Carcaça, Encruzilhada, Litania, Lastro, Fama, Têmpera, portador,
sucateiro, têmpera (metáfora), liga, coleira, pacto.

**Evitar em texto diegético:** "épico", "escolhido", "destino", "Destiny Board" (use Litania),
jargão de game (buff / DPS / multiplicador).

**Nota sobre "Têmpera":** a metáfora de forja ("quente", "esfriou", "calor") é livre. O termo
técnico só escapa na boca de personagem em momento de peso — no tutorial, **uma única vez**,
no Passo 7.

## Placeholders

Nomes de conteúdo ainda não definidos vivem como `PH_*` (ex.: `PH_RES_MADEIRA_T1`,
`PH_MOB_T2_A`). Índice em `gdd/60-producao/placeholders.md`.

Para achar tudo que falta: `grep -r "PH_"`.

**Ao nomear um placeholder, ele sai de TODAS as páginas.** Os dois nunca convivem.

## Convenção de código

Identificadores seguem o lore em português: `internal/litania/`, `GET /litania`,
`litania.json`, `player_litania_progress`. Módulo Go: `escoria`.

## Nível das specs — direção de arte, não engenharia

A IA é **diretora de arte**: diz o que a coisa precisa fazer, por quê, e o que precisa ser
verdade no fim. Sidinei é **engenheiro-chefe**: decide todas as estruturas.

Specs contêm: o quê, por quê, critérios de aceite verificáveis, o que precisa existir **em
prosa**, e as perguntas que a implementação precisa responder.

Specs **não** contêm: schema SQL, shape de endpoint, tipos de cliente, assinatura de função,
algoritmo, tratamento de erro. Nada disso é escrito pela IA.

**Uma exceção, e só uma:** a forma dos arquivos em `data/`. A IA os produz nos chats de
conteúdo e Sidinei os consome — se a forma não estiver acordada, o arquivo entregue não carrega.
Sidinei define a forma; a IA segue. Fica registrada na `dados.md` da fatia.

Se uma spec começar a escrever código, ela está errada.
