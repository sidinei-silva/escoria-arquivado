---
tags: [processo]
---

# A Receita

> Como se trabalha neste projeto. Se você voltar depois de um mês sem tocar,
> leia só isto.
>
> Regras de uso de IA: [[../AGENTS|AGENTS.md]].

## O princípio

Uma fatia entrega algo que o jogador **vê ou faz**. Ela não fecha até que fluxo,
regra, teste, código e lore estejam juntos no mesmo commit.

Nada avança sozinho. O GDD não pode crescer à frente do código, e o código não
pode crescer sem regra escrita. Foi assim que dois projetos morreram antes.

## As seis peças

| Peça | O que é | Vive onde | Muda? |
|---|---|---|---|
| **Fluxo** | o jogo em ordem, do ponto de vista do jogador | `gdd/README.md`, `gdd/fluxos/` | sim — é a porta de entrada |
| **Regra** | o que o jogo faz. Numerada, testável | `gdd/sistemas/` | sim — vira regra nova |
| **Prosa** | por quê, tom, ficção. Não testável | `gdd/sistemas/`, `gdd/livro-de-imersao.md` | quando der vontade |
| **Spec** | como esta entrega acontece | `specs/` | **morre no fim** |
| **Teste** | a regra, executável | junto do código | não desincroniza |
| **ADR** | por que decidimos assim | `docs/adr/` | **nunca** — só nova |

Mais uma, que não é do jogo: **Estudo** — conceito e trade-off para eu decidir
depois, em `docs/estudos/`. Envelhece sem problema. **Não é decisão** e nunca
vira ADR por conveniência.

Só a spec é efêmera. Só a ADR é imutável.

## Por que o fluxo vem primeiro

Um GDD organizado por sistema é um índice: ótimo para procurar, péssimo para
lembrar. Se `gdd/` só tiver `sistemas/`, abrir a pasta não vai te dizer que jogo
é esse — foi exatamente o que aconteceu com o Eras do Brasil, e é por isso que
"não conheço mais o meu próprio projeto" é um risco real.

O `gdd/README.md` abre pelos 8 passos, em ordem, e cada passo aponta para os
sistemas que aparecem nele. Você lê o jogo acontecendo e desce para o detalhe.

**Todo arquivo de sistema começa com a linha de volta:** em que passo aparece,
com quais sistemas conversa. Navegação nas duas direções. É isso que faz o grafo
do Obsidian valer alguma coisa.

---

## Passo 1 — Localizar no fluxo

**Antes de tudo:** qual passo dos 8 esta fatia atende? Se não couber em nenhum,
ou é escopo novo — e aí vira decisão consciente — ou é `IDEIAS.md`.

Se o fluxo daquele passo não estiver escrito com detalhe suficiente, escreva
antes, em `gdd/fluxos/`. À mão, informal, do ponto de vista do jogador. É de lá
que as regras saem.

## Passo 2 — Escrever a regra no GDD

**Quando:** antes de qualquer código. Sem regra escrita, não há o que construir.

Abra o arquivo do sistema em `gdd/sistemas/`. Se não existir, crie — e é aqui
que a **migração preguiçosa do Notion** acontece: desce só o sistema desta
fatia, não o resto.

Separe prosa de regra fisicamente.

```markdown
# Têmpera

> Aparece no passo 5 (apresentada) e no passo 7 (nomeada).
> Conversa com [[loadout]] e [[fama]].

A arma esquenta com uso. Trocar é largar o metal antes da hora — não destrói
nada, mas apaga o forno. É o que faz presença valer mais que calendário.

## Regras

- **R-TEM-01** — sobe apenas com ação ativa da arma equipada
- **R-TEM-02** — congela offline: não sobe nem decai
- **R-TEM-03** — reseta ao valor base quando a arma equipada muda
- **R-TEM-04** — multiplica ganho de Fama; não altera dano
```

**O teste que separa regra de prosa:**

> *"O Guia parou perto do T8 por medo"* → prosa. Não testável, e não deve ser.
> *"Têmpera reseta na troca"* → regra. Numerada, testável.

**Cuidado com o falso terceiro tipo.** "Nome entre 3 e 20 caracteres" não é regra
de jogo — é validação de software. Não ganha ID, não vai pro GDD. Vive na spec.

Pergunta que decide: *se eu trocasse Go por Rust e web por mobile, isso
continuaria valendo?* Sim → regra. Não → encanamento, fica na spec.

## Passo 3 — Escrever a spec da entrega

**Quando:** quando as regras que ela implementa já existem.

**Tamanho:** algo que você consegue **terminar e demonstrar**. "Criação de
personagem" é uma spec. "Fatia 1" não é.

Um arquivo em `specs/`, quatro seções. As regras são **citadas, nunca
reescritas**.

```markdown
# S-001 · Criação de personagem

## Regras que esta entrega implementa
- R-JOG-01 — o jogador tem uma aparência escolhida por ele
- R-JOG-02 — o jogador tem um nome escolhido por ele
- R-JOG-03 — o jogador novo nasce em A Ressaca

## Comportamento de software
Não é regra de jogo; é o que o GDD nunca vai dizer.

- Nome entre 3 e 20 caracteres, com erro visível
- Recarregar a página mantém o jogador logado e na zona correta
- Nome vazio não cria jogador

## Como demonstro que funcionou
- Informo um nome, confirmo, e vejo a tela de A Ressaca
- Recarrego a página e continuo lá
- Tento confirmar vazio e vejo o erro

## Fora de escopo
- Executar ações na zona — S-002
- Inventário, equipamento, Fama
```

**Escopo negativo não é opcional.** É ele que impede a entrega de inchar.

## Passo 4 — Escrever os testes, falhando

Cada regra citada vira ao menos um teste que **nomeia o ID**.

```
Test_R_JOG_03_JogadorNovoNasceEmARessaca

  // DADO    nenhum jogador existente
  // QUANDO  crio um jogador chamado "Sidinei"
  // ENTÃO   a zona atual dele é A Ressaca
```

O comportamento de software também vira teste, mas **sem ID** — não é regra do
jogo.

Rode. Tem que falhar. Se passar sem código, o teste está errado.

## Passo 5 — Escrever o código

Até os testes passarem. Nada além disso.

Se precisar de uma decisão que a spec não responde — como estruturar, onde
validar, que biblioteca — **decida você**. É por isso que a spec para no
comportamento.

Se a decisão for **arquitetural e difícil de reverter**, anote para o Passo 6. Se
for de concorrência, `docs/estudos/` tem o material — e as perguntas do
[[../AGENTS|AGENTS.md]] valem antes de escrever.

`go test -race`, sempre.

## Passo 6 — Fechar

Cinco coisas, nesta ordem:

**1. Regra descoberta volta pro GDD.** Construindo, você esbarra em algo que
ninguém decidiu: *"e se o inventário encher no meio do ciclo?"*. É regra nova. ID
novo, teste novo. **Não pode morrer com a spec.**

É assim que o GDD cresce: por descoberta, nunca por antecipação.

**2. Decisão relevante vira ADR.**

```markdown
# ADR-001 — Estado autoritativo com dono único

**Data:** 2026-09-XX
**Status:** aceita

## Contexto
Várias goroutines poderiam tocar o estado do jogo.

## Decisão
Uma goroutine é dona do estado. Comandos chegam por channel.

## Consequência
Race é estruturalmente impossível no estado de jogo. O loop só toca memória;
I/O sai por channel. Custo: serializa. Compatível com partição por zona no
futuro — ver docs/estudos/02.

## Regras afetadas
nenhuma
```

**ADR nunca é editada.** Mudou de ideia? ADR nova, com `Status: substitui
ADR-001`. A antiga fica, com `Status: substituída por ADR-009`. O histórico é o
valor.

**3. A lore da fatia, escrita por mim.** Prosa em `gdd/`, à mão, sem IA gerando.
Vem **por último** de propósito: descrever algo que já roda é mais fácil e mais
gostoso que inventar no vazio, e assim a lore nunca vira gargalo da entrega.

**4. A spec morre.** `git mv specs/S-001-*.md specs/arquivo/`
Não é mantida, não é sincronizada, não é lida de novo.

**5. Um commit com tudo.** Fluxo, regra, teste, código, ADR, lore. Junto.

---

## Onde cada coisa mora

```
AGENTS.md              regras de uso de IA
IDEIAS.md              ideia de outro jogo cai aqui, nunca no GDD
README.md              porta de entrada do repo

gdd/
├── README.md          O FLUXO — os 8 passos. Comece por aqui.
├── fluxos/            o jogo em ordem, detalhado por trecho
├── sistemas/          um arquivo por sistema COM REGRAS
│   ├── jogador.md         R-JOG-*
│   └── zonas.md           R-ZON-*
└── livro-de-imersao.md    prosa longa do mundo

specs/
├── S-001-*.md         ativa (uma por vez)
└── arquivo/           mortas, nunca mais lidas

docs/
├── RECEITA.md         este arquivo
├── BACKLOG.md         as 10 fatias
├── adr/               decisões tomadas, imutáveis
└── estudos/           conceito e trade-off. NÃO é decisão.

data/                  as instâncias: zonas, mobs, armas, diálogos
server/  web/          código e testes
```

**Sistema é por sistema, não por funcionalidade.** Algo com regras próprias que
outros sistemas consultam. O prefixo do ID espelha o arquivo.

**Arquivo de sistema é curto.** Linha de navegação, prosa em parágrafos, regras
em lista. Passou de duas telas, provavelmente são dois sistemas.

## Convenções

| | Formato | Exemplo |
|---|---|---|
| Regra | `R-<SISTEMA>-<NN>` | `R-TEM-03` |
| Spec | `S-<NNN>-<slug>` | `S-001-criacao-personagem` |
| ADR | `ADR-<NNN>-<slug>` | `ADR-001-dono-unico` |

Numeração sempre cresce, nunca é reaproveitada.

## Regras de bolso

**Regra sem teste é decoração.** Se não vale um teste, não vale um ID — é prosa.

**Não escreva spec para o óbvio.** Ajustar um número, adicionar item ao JSON:
vai direto ao código.

**Se escrever a spec demora mais que fazer a coisa, não escreva a spec.**

**Nunca reescreva uma regra na spec.** Cite o ID. Duplicou, dessincronizou.

**Se não há regra para citar, a entrega não está pronta para começar.**

**Uma spec por vez.** Duas abertas significa que nenhuma vai fechar.

**Ideia vinda de outro jogo vai para `IDEIAS.md`.** Nunca direto para o GDD nem
para o backlog. Se depois da PoC ainda importar, era real.

**Estudo nunca vira ADR sem eu decidir.** O documento é insumo; a decisão é
minha, escrita com as minhas palavras.

## O ciclo, inteiro

```
  FLUXO ──► GDD ──cita IDs──► SPEC ──vira──► TESTES ──► CÓDIGO
(8 passos)  (regras)         (efêmera)     (permanentes)  │
    ▲          ▲                 │                        │
    │          │                 ▼                        │
    │          │              arquivo/                    │
    │          │                                          │
    │          └──── regra descoberta ────────────────────┤
    │                                                     │
    └──────── lore da fatia ◄── escrita por mim ──────────┤
                                                          │
                                    ADR ◄── decisão ──────┘
                                 (em pedra)
```

**Morre:** a spec.
**Fica:** fluxo, regra, teste, código, ADR, lore.
**Sincronização manual:** nenhuma.

## Por onde começar

1. Ler `gdd/README.md` — os 8 passos.
2. Escrever `specs/S-001-criacao-personagem.md` citando `R-JOG-01`, `R-JOG-02`,
   `R-JOG-03`. As três já existem.
3. Escrever os testes. Rodar. Ver falhar.
4. Escrever o código até passarem.
5. Fechar: lore do passo 1, spec para `arquivo/`, ADR se houve decisão, um
   commit.

Terminou o ciclo uma vez. Agora você sabe o que a receita significa, e pode
discordar dela com fundamento.
