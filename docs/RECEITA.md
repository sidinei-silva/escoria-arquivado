# A Receita

> Como se trabalha aqui. Se você voltar depois de um mês sem tocar, leia só isto.
>
> O que vai em cada pasta está no README de cada pasta.

## O princípio

Uma fatia entrega algo que o jogador **vê ou faz**. Ela não fecha até que fluxo,
regra, conteúdo, código e lore estejam no mesmo commit.

Nada avança sozinho. O GDD não cresce à frente do código, e o código não cresce
sem regra escrita. Foi assim que dois projetos morreram antes.

## A ordem de escrita

Sempre do jogador para a máquina, nunca ao contrário:

    fluxo → sistema → conteúdo → código+teste → lore

**Fluxo primeiro** porque regra sem contexto vira lista. Você não descobre que
precisa de `R-COL-04` até escrever "e se a mochila encher?" no fluxo.

**Lore por último** porque descrever algo que já roda é fácil, e inventar no
vazio não é. Escrita antes, ela vira gargalo da entrega.

**ADR a qualquer momento**, assim que a decisão for tomada.

---

## 1 · Localizar no fluxo

Qual trecho de `gdd/fluxos/` esta fatia atende? Se não couber em nenhum, ou é
escopo novo — decisão consciente — ou é `IDEIAS.md`.

Se aquele trecho ainda não estiver escrito, escreva agora. À mão, informal, na
visão do jogador. É de lá que as regras saem.

## 2 · Escrever a regra

Antes de qualquer código. Sem regra escrita, não há o que construir.

Abra o sistema em `gdd/sistemas/`. Se não existir, crie — é aqui que a
**migração preguiçosa** do Notion acontece: desce só o sistema desta fatia.

Separe prosa de regra fisicamente. O teste está no README de `sistemas/`.

## 3 · Definir o conteúdo

Se a fatia precisa de instância — zona, mob, receita, objetivo, diálogo — o JSON
em `data/` nasce agora, depois da regra que dá sentido a ele.

## 4 · Delimitar a entrega

Numa spec curta em `specs/`, ou na descrição da issue. Quatro blocos:

- **regras que implementa** — só os IDs, nunca reescritos
- **comportamento de software** — o que o GDD nunca vai dizer (validação de
  formato, recarregar a página, mensagem de erro). Sem ID; não é regra de jogo
- **como demonstro que funcionou**
- **o que fica de fora**

O último é a função inteira dela. `coleta.md` descreve o sistema completo; você
não entrega tudo numa noite.

## 5 · Testes falhando, depois código

Cada regra citada vira ao menos um teste que **nomeia o ID**:

    Test_R_COL_03_CicloConcluidoCreditaERecomeça

Rode. Tem que falhar. Se passar sem código, o teste está errado.

Depois o código, até passarem. Decisão que a spec não responde — estrutura,
onde validar, que biblioteca — é sua. `go test -race` sempre.

## 6 · Fechar

Cinco coisas, nesta ordem:

1. **Regra descoberta volta pro GDD.** Construindo, você esbarra em algo que
   ninguém decidiu. É regra nova, ID novo, teste novo. **Não pode morrer com a
   spec.** É assim que o GDD cresce: por descoberta, nunca por antecipação.
2. **Decisão relevante vira ADR** em `docs/adr/decisoes.md`.
3. **A lore da fatia, escrita por você**, à mão, depois do código passar.
4. **A spec morre** — `git mv` para `specs/arquivo/`, ou fecha a issue.
5. **Um commit com tudo.**

---

## Regras de bolso

**Regra sem teste é decoração.** Se não vale um teste, não vale um ID — é prosa.

**Não escreva spec para o óbvio.** Ajustar um número, adicionar item ao JSON:
vai direto ao código.

**Se escrever a spec demora mais que fazer a coisa, não escreva a spec.**

**Nunca reescreva uma regra fora de `sistemas/`.** Cite o ID. Duplicou,
dessincronizou.

**Se não há regra para citar, a entrega não está pronta para começar.**

**Uma spec por vez.** Duas abertas significa que nenhuma vai fechar.

**Ideia vinda de outro jogo vai para `IDEIAS.md`.** Nunca direto para fluxo,
sistema ou backlog.

**Escreva um fluxo por vez.** O da fatia atual. Se os seis nascerem hoje, o GDD
voltou a andar na frente do código.