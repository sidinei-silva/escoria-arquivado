# ideias/

**Quarentena.** Sistema legal visto em outro jogo cai aqui, nunca direto no GDD
nem no backlog.

## Por que existe

Foi carregar sistema de outros jogos por antecipação que matou o Eras do Brasil:
virou coop, voltou pra multiplayer, quase virou dungeon crawler, absorvia
qualquer mecânica boa que aparecesse. Projeto sem critério de recusa não termina.

Esta pasta é o critério de recusa. A ideia não é rejeitada — é **adiada com
endereço**. Isso basta pra ela parar de pressionar o escopo.

## A regra

**Nada aqui está aceito.** Um arquivo em `ideias/` não é compromisso, não é
roadmap, não entra em estimativa. É uma anotação com data.

**A data é o dado mais importante.** Ideia que ainda incomoda seis meses depois
é diferente de ideia que empolgou numa noite. O nome do arquivo carrega isso.

## Como uma ideia sai daqui

Quando **as três** forem verdade:

1. A PoC está fechada, ou a ideia é claramente independente dela
2. Você consegue dizer que problema do jogador ela resolve — não "seria legal"
3. Você consegue dizer o que ela **custa**: que sistema encosta, que regra muda,
   o que ela torna mais difícil depois

Aí ela vira fluxo em `gdd/fluxos/` ou item em `docs/BACKLOG.md`, e o arquivo é
apagado. Ideia promovida não fica nos dois lugares.

## Como uma ideia morre

Você relê e não sente mais nada. Apague. O `git log` guarda.

Não existe pasta de "ideias arquivadas" — isso seria quarentena da quarentena.

## O que NÃO entra aqui

- **Bug ou tarefa** → backlog
- **Regra que o jogo já tem** → `gdd/sistemas/`
- **Decisão já tomada** → `docs/adr/decisoes.md`
- **Dúvida de uma fatia em curso** → "A decidir" no arquivo de fluxo

Isso aqui é só o que o jogo **não faz e talvez nunca faça**.

## Formato

Um arquivo por ideia: `AAAA-MM-DD-nome-curto.md`. Copie `_GABARITO.md`.

Curto. Se você precisar de mais de uma tela, já está projetando — e projetar o
que não foi aceito é o que inflou o Eras.