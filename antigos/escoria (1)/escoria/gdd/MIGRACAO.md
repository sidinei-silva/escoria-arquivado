# Migração do Notion → `gdd/`

> Checklist vivo. Marque conforme cada página for reescrita e revisada.
> Reescrita **manual**, página a página — o objetivo é revisar o conteúdo antes que ele vire
> fonte da verdade, não só transportá-lo. Ver ADR-009.

**Total: 45 páginas.** 11 já foram revisadas na sessão de 21/07/2026 (marcadas ✅ revisada).

## Legenda

- `[ ]` pendente · `[x]` migrada e revisada
- ✅ **revisada** — conteúdo já corrigido no Notion em 21/07/2026, deve vir limpo
- ⚠️ **nunca revisada** — espere drift de nomenclatura
- 🗑️ **candidata a descarte** — avaliar se migra ou morre
- 🔀 **sobreposição** — conteúdo agora duplicado em `specs/` ou `docs/`

---

## Etapa 1 — Inventário e backlog

- [x] Levantar a árvore completa (45 páginas)
- [x] Reescrever os 67 títulos de tarefa conforme ADR-008

## Etapa 2 — Cânone e Overview · 4 páginas

| | Página | Destino | Nota |
| --- | --- | --- | --- |
| [ ] | Consolidado — Cânone v1.0 | `gdd/canone.md` | ✅ revisada (gancho A11) |
| [ ] | High Concept | `gdd/00-overview/high-concept.md` | ⚠️ |
| [ ] | Design Pillars | `gdd/00-overview/pilares.md` | ⚠️ |
| [ ] | Glossário | `gdd/00-overview/glossario.md` | ⚠️ crítico — é onde o vocabulário vive |

## Etapa 3 — Narrativa I · 3 páginas

| | Página | Destino | Nota |
| --- | --- | --- | --- |
| [ ] | Narrativa — Overview | `gdd/10-narrativa/overview.md` | ⚠️ |
| [ ] | Narrativa Jogada — PoC (8 passos) | `gdd/10-narrativa/narrativa-jogada-poc.md` | ✅ revisada (Passo 5 e 7) |
| [ ] | Personagens | `gdd/10-narrativa/personagens.md` | parcialmente revisada |

## Etapa 4 — Narrativa II · 4 páginas

| | Página | Destino | Nota |
| --- | --- | --- | --- |
| [ ] | Prévio — Ambientação do Player | `gdd/10-narrativa/previo-ambientacao.md` | ⚠️ |
| [ ] | Cosmologia | `gdd/10-narrativa/cosmologia.md` | ⚠️ |
| [ ] | Ganchos de Arco Longo | `gdd/10-narrativa/ganchos.md` | parcialmente revisada |
| [ ] | Tom & Estilo | `gdd/10-narrativa/tom-e-estilo.md` | ✅ revisada (léxico + Têmpera) |

## Etapa 5 — Mundo · 5 páginas

⚠️ **Seção inteira nunca revisada.** É de onde saem os nomes de zona que vão para `zones.json`.

| | Página | Destino | Nota |
| --- | --- | --- | --- |
| [ ] | Mundo — Overview | `gdd/20-mundo/overview.md` | ⚠️ |
| [ ] | A Escória | `gdd/20-mundo/a-escoria.md` | ⚠️ |
| [ ] | Facções & Panteões | `gdd/20-mundo/faccoes-panteoes.md` | ⚠️ |
| [ ] | Bestiário — Mitos sem Pacto | `gdd/20-mundo/bestiario.md` | ⚠️ pode duplicar `50 · Mobs` |
| [ ] | Geografia & Zonas | `gdd/20-mundo/geografia-zonas.md` | ⚠️ crítico — nomes de zona |

## Etapa 6 — Mecânicas e Loop · 9 páginas

| | Página | Destino | Nota |
| --- | --- | --- | --- |
| [ ] | Mecânicas — Overview | `gdd/30-mecanicas/overview.md` | ✅ revisada (diagrama) |
| [ ] | Gear = Identidade | `gdd/30-mecanicas/gear-identidade.md` | ⚠️ |
| [ ] | Tiers & Manifestação | `gdd/30-mecanicas/tiers.md` | ⚠️ |
| [ ] | Combate — Auto-battler | `gdd/30-mecanicas/combate.md` | ⚠️ |
| [ ] | A Litania & Fama | `gdd/30-mecanicas/a-litania-e-fama.md` | ✅ reescrita completa |
| [ ] | Economia — Lastro | `gdd/30-mecanicas/economia-lastro.md` | ⚠️ |
| [ ] | Crafting & Coleta | `gdd/30-mecanicas/crafting-coleta.md` | ⚠️ |
| [ ] | Core Loop (Gameloop) | `gdd/40-loop/core-loop.md` | ✅ revisada (Passo 7) |
| [ ] | Metagame | `gdd/40-loop/metagame.md` | ⚠️ toca T8 e rebirth |

## Etapa 7 — Conteúdo · 5 páginas

| | Página | Destino | Nota |
| --- | --- | --- | --- |
| [ ] | Conteúdo — Overview | `gdd/50-conteudo/overview.md` | ✅ revisada (nomeação) |
| [ ] | Armas da PoC — 3 Linhas | `gdd/50-conteudo/armas.md` | ⚠️ |
| [ ] | Skills | `gdd/50-conteudo/skills.md` | ⚠️ decide os slots Q/W/E |
| [ ] | Mobs — Mitos sem Pacto | `gdd/50-conteudo/mobs.md` | ✅ revisada (`PH_MOB_*`) |
| [ ] | Diálogos do Guia | `gdd/50-conteudo/dialogos-do-guia.md` | ✅ `dialogs.json` já em `data/` |

## Etapa 8 — Produção, índices e costura · 8 páginas

⚠️ **Atenção:** várias destas agora se sobrepõem ao que foi construído em `specs/` e `docs/`.
Migrar as duas versões criaria duas fontes da mesma verdade — o erro que a regra dos
placeholders existe para evitar.

| | Página | Destino | Nota |
| --- | --- | --- | --- |
| [ ] | Escopo da PoC | `gdd/60-producao/escopo-poc.md` | ✅ revisada |
| [ ] | Placeholders de Conteúdo | `gdd/60-producao/placeholders.md` | ✅ criada nesta sessão |
| [ ] | Tech Stack | — | 🔀 sobrepõe `specs/000-constituicao.md` |
| [ ] | Especificação da PoC | — | 🔀 sobrepõe `specs/` inteiro |
| [ ] | Decisões Fixas (ADR) | — | 🔀 sobrepõe `docs/decisoes.md` |
| [ ] | Roadmap & Fases | — | 🗑️ sobreposto por `specs/BACKLOG.md` |
| [ ] | Handoff — Produção com IA | — | 🗑️ plano de trabalho já executado e alterado |
| [ ] | GDD — A Escória (índice) | `gdd/README.md` | reescrever como índice do diretório |

Fecha com: índices de seção, varredura final de drift, repo empacotado.

---

## Varredura final de drift

Rodar ao fim da Etapa 8:

```bash
grep -rin "destiny board\|adventurer\|woodcutter\|skinner\|lumberjack\|miner\b" gdd/
grep -rin "o farol\|a cova\|floresta esquecida\|forte da montanha" gdd/
grep -rin "albion\|ao-bin\|premium\|zona azul\|focus" gdd/
grep -rn "PH_" gdd/ | wc -l    # deve bater com o índice de placeholders
```

## Páginas a apagar no Notion depois

Não tenho ferramenta de delete no Notion — as operações disponíveis são criar, ler, atualizar
e mover. Apague estas à mão quando a migração fechar:

- Handoff — Produção com IA (Fase 2+)
- Roadmap & Fases
- Tech Stack
- Decisões Fixas (ADR)
- Especificação da PoC

E, quando tudo estiver aqui, o workspace inteiro. Duas fontes vivas da mesma verdade divergem —
é a mesma regra dos placeholders.
