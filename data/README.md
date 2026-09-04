# data/

**As instâncias do jogo. Carregado no boot, imutável em runtime.**

    data/
    ├── zones.json          zonas da Margem Calada
    ├── objectives.json     tutorial, journal, daily
    └── dialogs.json        tutorial, idle_barks

## O invariante

**Conteúdo é dado, não código.** Adicionar uma zona, um mob ou um objetivo nunca
exige recompilar lógica. Se você precisou mexer no Go para adicionar conteúdo,
o conteúdo virou código e o invariante quebrou.

## Idioma

| O quê                             | Idioma     | Exemplo                                     |
| --------------------------------- | ---------- | ------------------------------------------- |
| Chaves                            | inglês     | `adjacent_zones`, `requires_objective`      |
| Valores que o código faz `switch` | inglês     | `gather`, `chain`, `daily`, `entered_zone`  |
| IDs de entidade do lore           | slug pt-BR | `a_ressaca`, `guia`, `portador_npc`         |
| IDs genéricos                     | inglês     | `wood`, `rough_stone`, `tut_04_gathering`   |
| Texto que o jogador lê            | pt-BR      | `"A maré não traz quem ela não pode usar."` |
| `subtext` (nota de direção)       | pt-BR      | não é exibido ao jogador                    |

A linha: **se vira constante em Go, é inglês. Se é nome próprio do mundo ou
texto de tela, é português.**

## Regras de forma

- IDs são **strings estáveis** definidas aqui, nunca inteiros de banco
- Nada muda em runtime — carrega no boot e é servido em memória
- Sem mutex: muitas goroutines lendo o que ninguém escreve é seguro
- **Referência quebrada falha no boot**, com mensagem clara. Um typo não pode
  virar `nil pointer` três dias depois

## Referências entre arquivos

O boot valida todas. Se uma quebrar, o servidor não sobe.

    zone.adjacent_zones           → zone.id
    zone.resources                → resource.id
    objective.zone                → zone.id
    objective.next                → objective.id
    objective.condition.zone      → zone.id
    objective.condition.resource  → resource.id
    dialog.requires_objective     → objective.id
    dialog.filters.zone           → zone.id

## Como o conteúdo cresce

**Mesmo tipo:** uma entrada nova no array. Não precisa de código.

**Grupo novo** dentro de `objectives.json` ou `dialogs.json` — `weekly`, por
exemplo: uma chave nova. Se o unmarshal for para um map por grupo em vez de
struct com campos fixos, também não precisa de código.

**Tipo novo de conteúdo** — recursos, mobs, receitas: um arquivo novo aqui, um
parser, e as referências dele na lista acima.

Nada nasce por antecipação. O arquivo aparece na fatia que precisa dele.

## Nota sobre subtext

`subtext` é nota de direção interna, nunca exibida ao jogador. Registra o que o
Guia omite.

É o que impede a lore de virar explicação de sistema, e é o motivo de um diálogo
poder ser reescrito depois sem perder a intenção original.