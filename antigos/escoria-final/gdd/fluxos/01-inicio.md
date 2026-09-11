---
tags: [gdd, fluxo]
---

# Fluxo — Início

> Escrito à mão em 22/08/2026, antes de qualquer código. Era o `rascunho.md` da
> raiz. Promovido sem reescrever: é o passo a passo do jogador do jeito que eu
> pensei, e a linguagem informal é proposital.
>
> Cobre os **passos 1 a 3** de [[../README|os 8 passos]].

## Fluxo Start

1. **Criar o personagem**
   1. Selecionar o tipo de body do usuário, tendo 2 tipos no momento:
      `NORMAL_MALE` e `NORMAL_FEMALE`
   2. Selecionar o nick do personagem
2. **Personagem nasce na Ressaca**
3. **Jogador é recebido por um velho** que explica quem ele é e o que tem que
   fazer
   1. Diálogos do tutorial mandando coletar pedra e madeira para fazer espada
4. **Personagem coleta pedra**
   1. Jogador clica em opção de coletar pedra
   2. Inicia idle de coleta de pedra
   3. Jogador pode clicar para parar a coleta
5. **Personagem coleta madeira**
   1. Jogador clica em opção de coletar madeira
   2. Inicia idle de coletar madeira
   3. Jogador pode parar de coletar
6. **Criar uma espada e escudo**
   1. Jogador clica em criar espada
   2. Jogador seleciona quantas quer criar, ou criação contínua
7. **Aparece opção de lutar**
   1. Jogador clica em lutar
   2. Aparece idle de luta
   3. Jogador pode parar a luta

## Regras que este fluxo já produziu

- [[../sistemas/jogador|R-JOG-01, R-JOG-02]] — aparência e nome escolhidos pelo
  jogador (passo 1)
- [[../sistemas/jogador|R-JOG-03]] — nasce em A Ressaca (passo 2)
- [[../sistemas/jogador|R-JOG-04]] — recebido pelo Guia (passo 3)

## Divergências com o cânone, a resolver

Anotadas, não corrigidas — o fluxo é meu e o cânone pode ceder:

- Aqui a ordem é **coletar → craftar espada → lutar**. Nos 8 passos do cânone é
  **receber a arma (passo 2) → combate (passo 3) → coleta (passo 4)**.
- Aqui aparece **escudo**; o cânone da PoC não menciona escudo.
- "Criação contínua" no craft não existe como regra em lugar nenhum ainda.

Decidir quando a fatia chegar, não agora.
