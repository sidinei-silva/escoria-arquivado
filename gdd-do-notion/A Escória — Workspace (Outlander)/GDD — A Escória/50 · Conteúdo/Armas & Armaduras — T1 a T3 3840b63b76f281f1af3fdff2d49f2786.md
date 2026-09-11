# Armas & Armaduras — T1 a T3

Roster de T1 a T3, no modelo do Albion: **três arquétipos, uma oficina cada, variedade que só abre no T3.** Nomes próprios são Fase 2 — aqui está a estrutura e a contagem.

> 📐 **Regra de leitura.** Tier não é uma lista diferente de armas. A mesma linha atravessa os tiers em material melhor. O que cresce com o tier é **quantas árvores existem**. Ver *Tiers & Manifestação*.
> 

## As três linhas e suas oficinas

| Linha | Função | Oficina | Ruído divino (T2–T3) |
| --- | --- | --- | --- |
| **Guerreiro** (corte) | dano consistente, frontline | A Bigorna | urgência; sangue velho; cinza e calor; disciplina |
| **Caçador** (distância) | rápido, risco maior | *(PH)* | paranoia; predação; frio; respiração curta |
| **Conjurador** (controle) | controle de espaço, lento | *(PH)* | gravidade; obstinação; silêncio sufocante |

## T1 — sucata sem dono

**Não craftável.** Vem do Guia (Passo 2) ou de achado. Sem linha, sem árvore, sem variação, sem deus.

- 1 arma genérica — *Lâmina Crua*
- 1 peça de corpo genérica — *Trapo Cru*

## T2 — uma por linha, zero variação

Primeira forja de verdade (Passo 5). É aqui que um deus repara no player pela primeira vez.

| Linha | Arma | Set de armadura |
| --- | --- | --- |
| Guerreiro | Espada Batida | Casco (placa) |
| Caçador | Arco Batido | Andarilho (couro) |
| Conjurador | Cajado de Fogo Batido | Ledor (pano) |

**Total T2:** 3 armas + 3 sets × 3 slots = **12 itens**.

## T3 — as árvores abrem

Três árvores por linha, uma arma-base cada, mais um off-hand por linha. **A armadura não ramifica** — continua um set por linha, em material melhor. Essa é a lição de escopo do Albion: dobre a variedade de armas e segure a de armaduras.

| Linha | Mão principal (3 árvores) | Off-hand |
| --- | --- | --- |
| Guerreiro | Espada · Machado · Maça | Escudo |
| Caçador | Arco · Lança · Adaga | Tocha |
| Conjurador | Cajado de Fogo · Cajado de Gelo · Cajado Sagrado | Tomo |

**Total T3:** 9 armas principais + 3 off-hand + 3 sets × 3 slots = **21 itens**.

Três árvores por linha em vez das seis ou sete do Albion. É o mínimo para uma árvore parecer árvore, e já são 9 armas com Q/W/E a desenhar, ou seja **27 skills**. As árvores restantes entram como conteúdo pós-PoC, uma por vez.

## Fora do escopo (T4+)

Variações dentro da mesma árvore (tipo Claymore e Dual Swords saindo de Espada) e as **armas de panteão** — o equivalente às artifact do Albion. É no T4 que "cada deus é uma arma" fica literal.

## Nomenclatura

Cru (T1) → Batida (T2) → Temperada (T3) → **nome próprio** (T4+). Ver *Tiers & Manifestação*.

## Template de ficha (Fase 2)

```jsx
id: <slot>_<linha>_<arvore>_t<tier>   // ex: weapon_guerreiro_espada_t3
slot: weapon | offhand | cabeca | torso | botas
linha: guerreiro | cacador | conjurador
arvore: ""           // vazio em T1-T2; obrigatório do T3
tier: 1..8
nome: ""             // sem nome próprio antes do T4
ruido_divino: ""     // textura sensorial; vazio em T1
skills: [Q, W, E]    // armas; armadura contribui 1 skill
```

## A fazer

- [ ]  Nomes das 3 oficinas (falta caçador e conjurador).
- [ ]  Nomes dos 3 sets — *Casco*, *Andarilho* e *Ledor* são placeholders.
- [ ]  Ruído divino por linha em T2 e por árvore em T3.
- [ ]  Mapear Q/W/E das 9 armas T3 + 1 skill por peça de armadura.