# fluxos/

**O jogo acontecendo, em ordem, na visão do jogador.**

Um arquivo por trecho do fluxo, numerado. Lidos em sequência, contam o jogo
inteiro. É por aqui que se entra no GDD.

## O que entra

O que o jogador faz, o que o servidor faz em resposta, quais regras valem, que
conteúdo é usado, e o que ainda não foi decidido.

## O que NÃO entra

**Regra escrita por extenso.** O fluxo cita `R-COL-01`; nunca repete o texto
dela nem o número que ela define. Se a duração da coleta aparecer aqui, existem
duas fontes de verdade e uma vai ficar velha.

O fluxo diz *"o servidor decide quanto demora (R-COL-02)"*.
O sistema diz *quanto* demora.

## Um arquivo por trecho, não por ação

Coletar madeira e coletar pedra usam as mesmas regras e o mesmo sistema — é um
arquivo. Separe só quando as regras divergirem de verdade.

## Como escrever

1. Copie `_GABARITO.md`
2. Numere: `NN-nome-curto.md`
3. Preencha "o que o jogador faz" primeiro, em passos numerados
4. Depois "o que o servidor faz" — é aqui que o fluxo vira arquitetura
5. Toda pergunta que aparecer vai para "a decidir". **Não resolva na hora.**
6. As regras citadas que ainda não existem: crie em `../sistemas/` agora

## Escreva um por vez

O da fatia atual. Não os seis.