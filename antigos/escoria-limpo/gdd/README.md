# GDD — A Escória

> **Esta página é o jogo acontecendo, em ordem.** Não é índice.
>
> Se você quer consultar um sistema específico, ele está em `sistemas/` — mas
> chegue lá **por aqui**, pelo passo em que ele aparece. É assim que o jogo fica
> na cabeça.

## Onde o design vive

| Onde | O quê |
|---|---|
| **Este repositório** | fluxos, regras numeradas, prosa dos sistemas que já existem |
| **Notion** ([workspace](https://app.notion.com/p/3720b63b76f281bc89c4d679f9d8cd1f)) | GDD completo — Overview, Narrativa, Mundo, Mecânicas, Loop, Conteúdo, Cânone v1.0 |

O Notion está **congelado** como referência histórica. Nada novo é escrito lá.

**Migração preguiçosa:** um sistema só desce do Notion para `sistemas/` quando a
fatia que precisa dele chegar. Migrar tudo de uma vez seria semanas sem código —
e o GDD voltaria a crescer mais rápido que o jogo, que é como o Eras do Brasil
morreu.

### Estado do Notion — o que confiar

São 45 páginas. Ao migrar um sistema, saiba o que esperar:

- **`00 · Overview`, `20 · Mundo` e parte de `40 · Loop` nunca foram revisados.**
  Espere drift de nomenclatura. Confira nomes contra o Cânone antes de trazer.
- Nomes de zona já corrigidos em 21/07/2026: O Farol → **A Ressaca**,
  A Cova → **A Bigorna**, Floresta Esquecida → **O Verde Surdo**,
  Forte da Montanha → **A Costela**. Se aparecer nome antigo, é página velha.
- "Destiny Board" foi migrado para **A Litania** em 8 páginas (ADR-004).
  Ocorrência remanescente é drift.
- Conteúdo derivado de Albion foi neutralizado em placeholders `PH_*` (ADR-003).
  Os 4 Mitos sem Pacto da ilha ainda são `PH_MOB_*` e precisam de nome.
- Cinco páginas de `60 · Produção` se sobrepõem a `docs/` e `specs/`.
  Provavelmente não migram.

## O fluxo do jogador — os 8 passos

A PoC inteira acontece na **Margem Calada**, a ilha inicial, sem outros
jogadores. Cada passo abaixo é uma fatia do `docs/BACKLOG.md`.

### 1 · Chegada — A Ressaca

O jogador cria o personagem e é cuspido na praia. Um velho o espera sem oferecer
ajuda para levantar. Ninguém o esperava, nenhuma profecia foi escrita.

→ [jogador](sistemas/jogador.md) · [zonas](sistemas/zonas.md)
→ diálogos: `dialog_passo1_abertura`, `dialog_passo1_reacao`
→ detalhe: [01-inicio](fluxos/01-inicio.md)

### 2 · A primeira arma

O jogador recebe uma arma, equipa, e vê o loadout mudar. A arma ainda é ruído
difuso — não tem deus decidido lá dentro.

→ `sistemas/loadout` *(a migrar do Notion)*
→ diálogos: `dialog_passo2_abertura`, `_reacao`, `_espera`

### 3 · O corpo age sozinho

Primeiro combate. O servidor conduz o ciclo; o jogador assiste e ganha Fama. É
aqui que o idle se explica sem ninguém explicar.

→ `sistemas/combate`, `sistemas/fama` *(a migrar)*
→ diálogos: `dialog_passo3_abertura`, `_reacao`, `_espera`

### 4 · A Escória dá

Coleta em loop, inventário enchendo. O lugar é generoso antes de cobrar.

→ [zonas](sistemas/zonas.md) · `sistemas/inventario` *(a migrar)*
→ diálogos: `dialog_passo4_abertura`, `_reacao`, `_espera`

### 5 · A forja — A Bigorna

Primeira viagem entre zonas, com trânsito. Refino e craft. O Guia apresenta a
Têmpera **sem usar a palavra**, só em metáfora de forja.

→ `sistemas/craft`, `sistemas/lastro`, `sistemas/viagem` *(a migrar)*
→ diálogos: `dialog_passo5_abertura`, `_reacao`, `_espera`

### 6 · O outro portador

Encontro com o NPC-Portador. **As armas se reconhecem** antes das pessoas. É o
primeiro sinal de que existe algo vivo no metal.

→ diálogos: `dialog_passo6_abertura`, `dialog_passo6_npc`, `dialog_passo6_reacao`

### 7 · A segunda arma

O jogador troca de arma e sente a Têmpera esfriar. **Único momento em que o Guia
diz a palavra "têmpera".** Combate T2 com fila de habilidades.

→ `sistemas/tempera` *(a migrar — a mecânica que define o projeto)*
→ diálogos: `dialog_passo7_abertura`, `_reacao`, `_espera`

### 8 · A saída — A Encruzilhada

A Litania inteira à vista. O jogador atravessa e sai da ilha. Fecha a chave
temática: **"você só assina depois"**.

→ `sistemas/litania` *(a migrar)*
→ diálogos: `dialog_passo8_abertura`, `dialog_passo8_reacao`

## Critério de PoC validada

1. O jogador completa os 8 passos sem travar
2. O loop fecha: zona → ação → recurso e Fama → Litania → novo conteúdo
3. Narrativa mínima presente
4. **Lastro e Fama percebidos como coisas diferentes**

## Sistemas que já existem aqui

- [jogador](sistemas/jogador.md) — `R-JOG-01` a `R-JOG-05`
- [zonas](sistemas/zonas.md) — `R-ZON-01` a `R-ZON-03`

Todo o resto ainda está no Notion. Isso é esperado.

## Prosa longa

- [livro-de-imersao](livro-de-imersao.md) — o mundo em prosa contínua, 7 livros. Não é cânone
  confirmado; é referência. Se a lore for reescrita, ela vira insumo, não trava.
