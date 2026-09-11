# GDD — A Escória

Documento de design vivo. Lore, mundo, narrativa, personagens, conteúdo.

**Regra de precedência:** para qualquer coisa criativa, este diretório vence. Para qualquer
coisa técnica, `specs/` vence. Para processo, `CLAUDE.md`. Se houver conflito, avise em vez de
escolher em silêncio.

## Estrutura

```
gdd/
├── canone.md              Consolidado — Cânone v1.0 (a ata de decisões que fundou o GDD)
├── 00-overview/           High Concept, Pilares, Glossário
├── 10-narrativa/          História e como ela é contada
├── 20-mundo/             A Escória, facções, bestiário, geografia
├── 30-mecanicas/          Gear, tiers, combate, A Litania, Lastro, crafting
├── 40-loop/               Core loop e metagame
├── 50-conteudo/           Os criativos: armas, mobs, diálogos, zonas
└── 60-producao/           Escopo, stack, roadmap, decisões, placeholders
```

## Migração do Notion — como fazer

O conteúdo vive no Notion (workspace "A Escória") e precisa vir para cá. **Use a exportação
nativa**, não transcreva à mão.

### 1. Exportar

No Notion, na página **GDD — A Escória**:

`···` (canto superior direito) → **Export** → formato **Markdown & CSV** → marcar
**Include subpages** → **Create folders for subpages** → Export.

Vai baixar um `.zip`.

### 2. Limpar

O export vem com UUID no nome de todo arquivo e pasta, e os links internos apontam para esses
nomes. O script resolve:

```bash
unzip ~/Downloads/Export-*.zip -d /tmp/gdd-export
python3 scripts/limpar-export-notion.py /tmp/gdd-export gdd/
```

Ele remove os UUIDs, transforma os nomes em slugs sem acento, reescreve os links internos e
converte para `[[wikilinks]]` do Obsidian se você pedir com `--wikilinks`.

### 3. Conferir

Depois de importar, rode uma varredura pelos termos que **não** deveriam mais existir:

```bash
grep -rin "destiny board\|adventurer\|woodcutter\|skinner\|o farol\|a cova\|floresta esquecida" gdd/
```

Se aparecer algo, é drift que escapou — corrija ou traga para uma sessão de revisão.

### 4. Abrir no Obsidian

Aponte o vault para a **raiz do repositório**, não para `gdd/`. Assim `specs/` e `docs/` também
ficam navegáveis e linkáveis, que é o ponto do monorepo.

## O que já está correto no Notion

Estas páginas foram revisadas e reescritas em 21/07/2026. Podem ser exportadas como estão:

| Página | O que mudou |
| --- | --- |
| A Litania & Fama | reescrita completa: 3 camadas + Têmpera + migração de nomenclatura |
| Tom & Estilo | léxico com Litania, "Destiny Board" proibido, nota dos dois sentidos de têmpera |
| Narrativa Jogada — PoC | Passo 7 reformulado; Têmpera no Passo 5 |
| Core Loop (Gameloop) | mesma correção do Passo 7, replicada aqui |
| Consolidado — Cânone v1.0 | gancho A11 corrigido |
| Escopo da PoC | 3 menções migradas |
| Diálogos do Guia | `dialogs.json` completo inserido |
| Mobs — Mitos sem Pacto | slots virando `PH_MOB_*` |
| Conteúdo — Overview | princípios de nomeação reescritos |
| Especificação da PoC | placeholders + Caminho 1 + contradição de IP corrigida |
| Placeholders de Conteúdo | página nova, índice dos 26 slots `PH_*` |

## O que ainda tem pendência conhecida

- **20 · Mundo** e **00 · Overview** não foram revisados nesta sessão. Podem conter drift de
  nomenclatura (nomes de zona antigos, "Destiny Board", nomes derivados de Albion).
- **40 · Loop & Sistemas** — só o Core Loop foi revisado; as demais páginas da seção não.
- Três questões abertas listadas em `specs/BACKLOG.md`.

## Depois de migrar

Quando o conteúdo estiver aqui, o Notion vira arquivo morto. Não mantenha os dois vivos — é
a mesma regra dos placeholders: os dois nunca convivem. Escolha um e aposente o outro.
