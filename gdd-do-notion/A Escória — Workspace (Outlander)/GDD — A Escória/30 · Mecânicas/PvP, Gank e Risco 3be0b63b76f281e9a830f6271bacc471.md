# PvP, Gank e Risco

> ⚔️ **Decisões de PvP, gank e risco.** Registradas em 15/08/2026, recuperadas de conversas de 27 e 31 de maio de 2026 que nunca foram documentadas.
> 

> 
> 

> **Nada disto está na PoC.** A PoC é o tutorial, em zona segura, sem PvP. Esta página existe para que estas decisões não se percam de novo — e para que a arquitetura do servidor não feche portas que elas vão precisar.
> 

## Por que PvP em jogo idle é um problema

Full loot funciona no Albion porque o jogador **escolheu** entrar na zona perigosa, está presente, pode lutar, fugir e decidir. A perda é consequência das escolhas dele.

Em um jogo idle, se o jogador é morto enquanto está no trabalho, ele perdeu tudo sem tomar nenhuma decisão. Isso não é tensão — é uma máquina caça-níqueis que rouba, e historicamente causa abandono em jogos idle e mobile.

As decisões abaixo existem para resolver essa tensão sem abrir mão do full loot.

## Decisão 1 — Fila de prioridade de alvos

O gank é uma **ação de zona**: o jogador seleciona "gankar" e entra em um matchmaking que procura alvos em ordem de consentimento decrescente.

| Prioridade | Alvo | Consentimento | Escopo |
| --- | --- | --- | --- |
| 1 | quem também está gankando | total — os dois escolheram | zona atual + adjacentes |
| 2 | quem está viajando | implícito — trânsito é o momento exposto | zona do gank; adjacentes contam como 2 ganks |
| 3 | quem está em qualquer atividade | presente, mas não escolheu | zona atual |
| 4 | quem está ocioso (online) | nenhum | zona atual |

**O princípio:** o PvP só encontra quem não quer depois de esgotar quem quer. Cada nível só é alcançado após um tempo de busca sem resultado no anterior.

## Decisão 2 — Offline é intocável

**Só quem está online pode ser gankado.** Jogador desconectado nunca entra no matchmaking, em nenhuma prioridade.

Isso é coerente com a definição de idle do projeto: offline congela, não avança. **Estar online é o que expõe e o que recompensa** — a Têmpera só sobe com presença ativa, e a presença ativa é o que te coloca na fila de alvos. Risco e recompensa moram na mesma escolha.

## Decisão 3 — Postura, não loadout separado

Quando o gank dispara, o jogador escolhe uma **postura**: enfrentar ou fugir.

**Um loadout só.** Não existe kit de PvP separado nem kit de corrida. O que decide se a fuga dá certo é se as habilidades que o jogador **já tem equipadas** servem para fugir.

Três razões:

- **Pilar 1.** Loadouts trocados na hora desmontam "você é o que veste". Com guarda-roupa, você não é o que veste — você tem opções.
- **Têmpera.** Ela é por arma equipada e reseta na troca. Trocar de loadout ao ser atacado zeraria a Têmpera de quem foi atacado — punindo a vítima.
- **Dominância.** Se todos podem ter kit de fuga, todos têm. A escolha evapora e o kit vira imposto.

**A consequência desejada:** quem montou loadout de mobilidade foge bem e luta mal. Quem montou de dano luta bem e foge mal. A decisão continua sendo o que você vestiu ao sair — antes de saber o que ia encontrar.

## Decisão 4 — Caça de gato e rato

Quando um lado quer fugir e o outro quer caçar, há uma resolução baseada em atributos dos dois jogadores e nas habilidades equipadas.

- **Fugitivo vence** → escapa, sem combate
- **Caçador vence** → começa o combate

**Em aberto.** Se a resolução for puramente comparação de atributos, não é jogo — é uma rolagem de dados disfarçada. Quem está sendo gankado está pedindo agência. Vale o fugitivo ter ao menos **uma decisão real** na perseguição: direção, gastar um recurso, tentar uma vez ou aguentar. Não precisa ser complexo — precisa ter um momento em que ele escolheu algo.

Decidir quando o sistema for construído.

## Decisão 5 — Faseamento

| Fase | O que entra |
| --- | --- |
| PoC | **Nada.** Tutorial em zona segura, sem PvP |
| V1 | Zonas de risco baixo e médio. PvP sem perda, ou com penalidade pequena |
| V2 | Zona de risco alto. **Full loot entra aqui**, não antes |
| V3+ | Zona de risco máximo — só com base de jogadores real |

O raciocínio: validar que o PvP simples engaja **antes** de introduzir perda total. Full loot é a decisão mais delicada do projeto e não deve ser a primeira testada.

## Restrição conhecida — massa crítica

Dois sistemas do projeto dependem de população e vão falhar com poucos jogadores:

- **Matchmaking de gank** em mapa vazio não encontra ninguém
- **Economia feita por jogador** não se forma sem volume de trocas

Saída registrada para validação: não são necessários milhares de jogadores para validar o conceito. Vinte pessoas de uma comunidade pequena bastam, e o mercado pode nascer com oferta semeada artificialmente.

## O que isto exige da arquitetura desde já

Nada precisa ser construído agora, mas duas coisas não devem ser fechadas de um jeito que impeça estas decisões depois:

- **Estado do jogador precisa ser observável pelo servidor** — ocioso, viajando, em atividade, offline. A fila de prioridade depende inteiramente disso.
- **O servidor é a autoridade do tempo e do estado.** Um cliente que informa o próprio estado permitiria mentir sobre estar offline para escapar da fila.

## Origem

Decisões 1, 3 e a base da 5 vieram das conversas de 27 e 31 de maio de 2026, recuperadas de um export de histórico. Decisões 2 e 4 foram fechadas em 15 de agosto de 2026.

A lição registrada junto: **decisão tomada em conversa não sobrevive.** Estas ficaram três meses fora de qualquer documento e quase foram redecididas do zero.

## Nota — Temporadas (pós-lançamento, não decidido)

Ideia levantada em 15/08/2026 a partir do formato de Leagues do RuneScape. Não é decisão, é

anotação para não perder.

O formato: mundo paralelo, personagem novo descartável, taxas muito aceleradas, restrição de

região que força builds diferentes, duração de semanas, sem economia (todos autossuficientes).

Três coisas encaixariam bem n'A Escória:

- **Restrição por Carcaça** — "esta temporada só a Pálida e a Solar estão abertas" força linhas
    
    de arma diferentes e recontextualiza o mesmo conteúdo. As cinco Carcaças já existem.
    
- **Taxa acelerada** resolve a fraqueza estrutural do idle: progressão custa tempo de calendário
    
    e não há como comprimir. É a única forma de compressão que o gênero permite.
    
- **Resolve o buraco veterano/novato** — em MMO idle o veterano fica sem o que fazer e o novato
    
    nunca alcança.
    

**O problema:** Leagues remove a economia (todos jogam isolados, sem troca), porque economia leva

meses para se formar e a temporada dura semanas. Isso arrancaria justamente a alma d'A Escória —

economia feita por jogador e full loot com valor real.

**Conclusão provisória:** só faz sentido depois que existir jogo base suficiente para valer a

pena rejogar. Temporada é variação sazonal de um jogo que já existe; sem o jogo, não há o que

recontextualizar.