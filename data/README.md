# data/

**As instâncias do jogo. Carregado no boot, imutável em runtime.**

Zonas, mobs, armas, receitas, nós da Litania, diálogos, objetivos.

## O invariante (ADR-011)

**Conteúdo é dado, não código.** Adicionar uma zona, um mob ou um objetivo
nunca exige recompilar lógica. Se você precisou mexer no Go para adicionar
conteúdo, o conteúdo virou código e o invariante quebrou.

## Regras de forma

- IDs são **strings estáveis** definidas aqui, nunca inteiros de banco
- Nada muda em runtime — carrega no boot e é servido em memória
- Sem mutex: muitas goroutines lendo o que ninguém escreve é seguro
- Referência quebrada **falha no boot**, com mensagem clara. Um typo em
  `receitas.json` não pode virar `nil pointer` três dias depois

## Sem gabarito

Cada arquivo tem forma própria. A forma nasce no fluxo que precisa dele e é
decidida junto com as regras do sistema correspondente.

O que vale para todos: um `id` string, e todo campo que aponta para outro
arquivo usa o `id` de lá.