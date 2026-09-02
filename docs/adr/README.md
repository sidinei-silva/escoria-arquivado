# adr/

**As decisões tomadas. Um arquivo só: `decisoes.md`, append-only.**

ADR responde *por que foi assim*, meses depois, quando você não lembra e está
prestes a redecidir a mesma coisa.

## Quando escrever

Ao fechar uma fatia em que você tomou uma decisão que:

- é difícil de reverter, **ou**
- você vai esquecer o motivo, **ou**
- alguém (inclusive você) discordaria sem contexto

Escolher nome de variável não é ADR. Escolher SQLite em vez de Postgres é.

## Quando NÃO escrever

Antes de decidir. ADR registra decisão tomada, não hipótese. Material para
decidir é estudo, e estudo não mora no repo.

## A regra que faz ela funcionar

**ADR nunca é editada.** Mudou de ideia? ADR nova com `Status: substitui
ADR-007`. A antiga fica, com `Status: substituída por ADR-015`.

O valor está no histórico. Editar apaga o histórico.

## Numeração

Sequencial, no arquivo único, em ordem. Próxima: verifique a última.