#!/usr/bin/env bash
#
# Semeia o GitHub com o backlog da PoC.
#
# Cria: labels de camada e fatia, 10 milestones (uma por fatia), 71 issues.
# Opcionalmente cria um Project v2 e adiciona as issues nele.
#
# Pré-requisitos:
#   - gh CLI instalado           https://cli.github.com
#   - autenticado                gh auth login
#   - escopos necessários        gh auth refresh -s project,repo
#   - repositório já existente   (ver --criar-repo abaixo)
#
# Uso:
#   ./scripts/seed-github.sh --repo SEU_USUARIO/escoria
#   ./scripts/seed-github.sh --repo SEU_USUARIO/escoria --com-project
#   ./scripts/seed-github.sh --repo SEU_USUARIO/escoria --dry-run
#
set -euo pipefail

REPO=""
DRY_RUN=0
COM_PROJECT=0
AQUI="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --repo)        REPO="$2"; shift 2 ;;
    --dry-run)     DRY_RUN=1; shift ;;
    --com-project) COM_PROJECT=1; shift ;;
    -h|--help)     sed -n '2,20p' "$0"; exit 0 ;;
    *) echo "argumento desconhecido: $1" >&2; exit 1 ;;
  esac
done

if [[ -z "$REPO" ]]; then
  echo "erro: informe --repo USUARIO/REPOSITORIO" >&2
  exit 1
fi

run() {
  if [[ $DRY_RUN -eq 1 ]]; then
    printf '  [dry-run] %s\n' "$*"
  else
    "$@"
  fi
}

command -v gh >/dev/null || { echo "erro: gh CLI não encontrado" >&2; exit 1; }
gh auth status >/dev/null 2>&1 || { echo "erro: rode 'gh auth login' primeiro" >&2; exit 1; }

echo "==> Repositório: $REPO"
[[ $DRY_RUN -eq 1 ]] && echo "==> MODO DRY-RUN — nada será criado"

# ---------------------------------------------------------------- labels

echo
echo "==> Labels"

criar_label() {
  local nome="$1" cor="$2" desc="$3"
  echo "  $nome"
  run gh label create "$nome" --repo "$REPO" --color "$cor" --description "$desc" --force
}

criar_label "camada:backend"   "1D76DB" "API em Go"
criar_label "camada:frontend"  "5319E7" "Cliente React"
criar_label "camada:dados"     "0E8A16" "JSON de conteúdo em data/"
criar_label "camada:narrativa" "D93F0B" "Diálogo, texto diegético"

for i in $(seq 0 9); do
  criar_label "fatia:$i" "FBCA04" "Fatia vertical $i"
done

# ------------------------------------------------------------ milestones

echo
echo "==> Milestones"

# pula o cabeçalho; separador é TAB
tail -n +2 "$AQUI/milestones.tsv" | while IFS=$'\t' read -r titulo descricao; do
  [[ -z "$titulo" ]] && continue
  echo "  $titulo"
  if [[ $DRY_RUN -eq 1 ]]; then
    printf '  [dry-run] gh api repos/%s/milestones -f title=%q\n' "$REPO" "$titulo"
  else
    gh api "repos/$REPO/milestones" \
      -f title="$titulo" \
      -f description="Entrega: $descricao" \
      --silent 2>/dev/null || echo "     (já existe, pulando)"
  fi
done

# ---------------------------------------------------------------- issues

echo
echo "==> Issues"

tail -n +2 "$AQUI/backlog.tsv" | while IFS=$'\t' read -r fatia milestone entrega camada titulo; do
  [[ -z "$titulo" ]] && continue

  corpo="**$milestone**

Entrega da fatia: $entrega

---

Contexto e critérios de aceite: \`specs/BACKLOG.md\` e a spec da fatia em \`specs/\`.

Antes de fechar: confira se a spec continua verdadeira. Se o código divergiu, atualize a spec
no mesmo commit (ver \`specs/000-constituicao.md\`, Definição de pronto)."

  echo "  [$fatia/$camada] $titulo"
  if [[ $DRY_RUN -eq 1 ]]; then
    continue
  fi

  gh issue create \
    --repo "$REPO" \
    --title "$titulo" \
    --body "$corpo" \
    --label "camada:$camada" \
    --label "fatia:$fatia" \
    --milestone "$milestone" \
    >/dev/null

  # respiro para não estourar rate limit secundário
  sleep 1
done

# --------------------------------------------------------------- project

if [[ $COM_PROJECT -eq 1 ]]; then
  echo
  echo "==> Project"
  OWNER="${REPO%%/*}"

  if [[ $DRY_RUN -eq 1 ]]; then
    echo "  [dry-run] gh project create --owner $OWNER --title 'A Escória — PoC'"
  else
    NUM=$(gh project create --owner "$OWNER" --title "A Escória — PoC" --format json | jq -r '.number')
    echo "  Project #$NUM criado"
    echo "  Adicionando issues..."
    gh issue list --repo "$REPO" --limit 200 --json url --jq '.[].url' | while read -r url; do
      gh project item-add "$NUM" --owner "$OWNER" --url "$url" >/dev/null
      sleep 1
    done
    echo "  pronto"
  fi
fi

echo
echo "==> Concluído."
echo
echo "Próximos passos manuais (mais fáceis pela interface web):"
echo "  1. No Project, agrupe por Milestone para ver as fatias como colunas"
echo "  2. Crie uma view 'Fatia atual' filtrando por milestone"
echo "  3. Ative o workflow de auto-add para novas issues entrarem sozinhas"
