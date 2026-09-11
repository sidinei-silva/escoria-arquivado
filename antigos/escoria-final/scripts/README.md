# scripts/

## `seed-github.sh`

Cria no GitHub: labels, 10 milestones (uma por fatia) e 71 issues, a partir do
`specs/BACKLOG.md`.

### Antes de rodar

```bash
gh auth login
gh auth refresh -s project,repo     # só se for usar --com-project
gh repo create SEU_USUARIO/escoria --private --source=. --remote=origin --push
```

### Rodar

```bash
./scripts/seed-github.sh --repo SEU_USUARIO/escoria --dry-run   # confere primeiro
./scripts/seed-github.sh --repo SEU_USUARIO/escoria
./scripts/seed-github.sh --repo SEU_USUARIO/escoria --com-project
```

Leva alguns minutos: há uma pausa de 1s entre issues para não estourar o rate limit
secundário do GitHub.

### Regenerar os TSVs

`backlog.tsv` e `milestones.tsv` são derivados do `specs/BACKLOG.md`. Se você editar o
backlog, regenere antes de semear — ou edite os TSVs direto, se a mudança for pontual.

O gerador está no histórico da sessão de 21/07/2026; se precisar dele de novo, peça.

### O que fazer pela interface web depois

Coisas que a CLI faz mal e o navegador faz bem:

1. No Project, agrupar por Milestone → as fatias viram colunas
2. Criar uma view "Fatia atual" filtrada pelo milestone em que você está
3. Ativar o workflow de auto-add, para issues novas entrarem no Project sozinhas

---

## `limpar-export-notion.py`

Limpa uma exportação Markdown do Notion: remove UUID dos nomes de arquivo e pasta, slugifica
(sem acento, minúsculo, hifenizado) e reescreve os links internos.

```bash
unzip ~/Downloads/Export-*.zip -d /tmp/gdd-export
python3 scripts/limpar-export-notion.py /tmp/gdd-export gdd/ --dry-run
python3 scripts/limpar-export-notion.py /tmp/gdd-export gdd/
```

`--wikilinks` converte os links para `[[formato do Obsidian]]` em vez de caminho relativo.
Caminho relativo é mais portátil (funciona no GitHub); wikilink integra melhor com o grafo do
Obsidian. Escolha um e mantenha.

Passo a passo completo da migração em `gdd/README.md`.
