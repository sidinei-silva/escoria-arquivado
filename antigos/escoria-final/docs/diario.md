# Diário de sessões

> Log do que foi feito, sessão a sessão. Complementa `decisoes.md`: aqui vai o **trabalho**,
> lá vai o **porquê**.

---

## 2026-07-21 — Diálogos do tutorial, Têmpera, e reorganização do projeto

**Entregue**

- `data/tutorial/dialogs.json` — 22 blocos de diálogo (8 passos) + 11 idle barks por loadout.
  Preserva as 7 falas-âncora canônicas. Chave temática "você só assina depois" recorre nos
  passos 2 e 3 e fecha no 8 em tempo passado.
- Modelo de progressão da Litania documentado com as três camadas (ADR-001).
- `specs/BACKLOG.md` — 58 tarefas reorganizadas em 10 fatias verticais (ADR-007).
- Esqueleto do monorepo: `CLAUDE.md`, `specs/000-constituicao.md`, `specs/001-personagem/`.

**Corrigido**

- Migração completa "Destiny Board" → "A Litania" em 8 páginas do GDD, incluindo
  identificadores de código (ADR-004).
- Passo 7 reformulado: deixou de ser "trocar zera a Fama / sacrifício" e virou descoberta de
  identidade. A regra antiga estava replicada em três páginas.
- Dependência de dados do Albion eliminada do backlog; nomenclatura derivada neutralizada em
  placeholders `PH_*` (ADR-003).
- Drift de nomes de zona no backlog: O Farol → A Ressaca, A Cova → A Bigorna,
  Floresta Esquecida → O Verde Surdo, Forte da Montanha → A Costela.

**Decidido**

ADR-001 a ADR-008. O ADR-008 reajustou o nível das specs no fim da sessão:
de contrato para direção de arte — a IA parou de escrever qualquer código, inclusive de
exemplo.

**Fase B — migração do GDD (revertida para manual)**

Primeiro decidida por exportação nativa (ADR-009). Sidinei reverteu com argumento melhor: o
export carrega o drift junto, em silêncio, e a migração é a única chance de revisar o conteúdo
antes que ele vire fonte da verdade. Duas seções nunca haviam sido revisadas. Ver ADR-010.

Plano de 8 etapas criado. `scripts/limpar-export-notion.py` fica no repo como ferramenta
auxiliar, mas não é mais o caminho principal.

**Etapa 1 concluída**

- Árvore real levantada: **45 páginas**, não as ~25 estimadas.
- `gdd/MIGRACAO.md` — checklist completo por etapa, com marcação do que já foi revisado,
  do que nunca foi, e do que se sobrepõe a `specs/` e `docs/`.
- `specs/BACKLOG.md` — 71 títulos reescritos em prosa conforme ADR-008. Nenhum título nomeia
  rota, tabela ou função. Única exceção mantida: nomes de arquivo em `data/`.

**Pendente**

- Etapas 2 a 8 da migração.
- `00 · Overview`, `20 · Mundo` e parte de `40 · Loop` nunca revisados — espere drift.
- Cinco páginas de `60 · Produção` se sobrepõem a `specs/` e `docs/`; decidir na Etapa 8.
- Três questões abertas no backlog: as 5 opções de destino de saída, os slots Q/W/E,
  e se a armadura entra no tutorial.
- Nomear os placeholders `PH_MOB_*` (4 Mitos sem Pacto da ilha).
