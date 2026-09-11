# 06 — Zonas, presença e PvP futuro

> **Nada aqui é para construir agora.** A PoC não tem multiplayer nem gank.
>
> Este documento existe por um motivo só: garantir que as decisões das Fatias 0 a
> 9 não fechem a porta do jogo final. Ler para saber o que **não** fazer.

## A zona é a unidade natural de partição

Voltando ao que o GDD pede no estado final:

- o jogador está sempre em **exatamente uma** zona
- a zona define quais ações existem nela — coletar o quê, lutar contra quem, se
  pode gank
- gank procura alvo na **mesma zona ou adjacente**, com hierarquia de regras
- viagem tem trânsito e conecta zonas

Isso é uma partição pronta, escrita no design antes de existir arquitetura. A
maior parte da interação é **dentro** de uma zona; interação entre zonas é
limitada e explícita (viagem, gank adjacente).

Isso é sorte de design, e vale reconhecer: nem todo jogo tem uma fronteira tão
limpa. Um mundo aberto contínuo não teria.

## O desenho de destino

```
        ┌──────────────┐        ┌──────────────┐
        │  Zona: Verde │◄──────►│ Zona: Costela│
        │    Surdo     │        │              │
        │  (goroutine) │        │  (goroutine) │
        └──────────────┘        └──────────────┘
               ▲                       ▲
               │                       │
               └────── Roteador ───────┘
                          ▲
                          │
                      Commands
```

Cada zona é dona do estado dos jogadores que estão nela. Processa os ticks deles,
resolve o matchmaking de gank local. Zonas conversam por mensagem.

Isso é actor model — mas chego nele por necessidade, com o desenho já pronto, em
vez de adotar o nome antes de ter o problema.

**Viagem vira handoff:** zona de origem remove o jogador, manda mensagem para a
zona de destino, destino aceita. Enquanto isso, o jogador está "em trânsito" —
que é um estado real do GDD, não um limbo técnico. O design já resolveu o caso
difícil.

## Por que o single-owner de hoje cresce para isto

O ponto central de [02-ownership-e-concorrencia](02-ownership-e-concorrencia.md), repetido porque é aqui que ele
se paga:

> Um loop dono de tudo **é** este desenho com N=1.

Crescer muda o **roteamento** — de "todo comando vai para o loop" para "o comando
vai para o dono da zona do jogador". A lógica de jogo não muda: validar,
transicionar, emitir continua idêntico, só que dentro de outro dono.

Se eu tivesse escolhido mutex por entidade, o gank exigiria travar dois jogadores
e uma ou duas zonas juntos — ordem de lock, deadlock, e o problema aparecendo
exatamente na feature mais característica do jogo.

## O que o gank exige do servidor

Vale enumerar agora porque cada item é uma pressão sobre decisões da PoC:

**Presença é estado de servidor.** Quem está em qual zona, fazendo o quê, há
quanto tempo, ocioso ou ativo. O servidor já vai saber disso pelas atividades —
desde que atividade seja estado, e não algo que o cliente reporta.

**Matchmaking é iniciado pelo servidor.** Ninguém pede para ser gankado. Isso é o
argumento definitivo do WebSocket em [04-rede-http-e-websocket](04-rede-http-e-websocket.md): o servidor
inicia a conversa.

**Gato e rato é resolução autoritativa.** Fuga ou combate é decisão do servidor a
partir do estado dos dois. Se qualquer parte disso dependesse de afirmação do
cliente, seria falsificável — e é por isso que [01-modelo-mental](01-modelo-mental.md) insiste que
comando é intenção, nunca resultado.

**Dois jogadores mudam juntos.** A operação toca o estado de dois. Com dono único
por zona, isso é uma transição dentro de um dono. Entre zonas adjacentes, é
mensagem entre dois donos — e aí aparece um problema de verdade: os dois estados
não mudam no mesmo instante. Isso vai precisar de protocolo (reserva, confirma,
libera) e é onde vou ter que estudar de novo, com o problema na mão.

## Regras da PoC que protegem o futuro

O que fazer nas Fatias 0 a 9 para não fechar a porta:

1. **Nunca deixar o cliente afirmar resultado.** Sempre intenção.
2. **Zona como entidade de primeira classe**, com identidade e adjacência em
   `zones.json`, mesmo com uma zona ativa. Se zona virar string solta no
   jogador, a partição fica cara depois.
3. **Atividade é estado do servidor**, não do cliente. Presença sai de graça
   disso.
4. **Comando carrega quem e onde.** Se o comando já diz o jogador, o roteador
   futuro descobre a zona sem mudar a assinatura.
5. **Nada de estado global fora do dono.** Um contador compartilhado "só para
   métrica" vira o ponto de contenção que quebra a partição.

## O que explicitamente NÃO fazer agora

- goroutine por zona com uma zona ativa
- protocolo de mensagem entre zonas
- matchmaking, mesmo simplificado
- sharding, Redis, NATS, serviço separado
- abstrair "dono" atrás de interface antecipando N donos

O último merece ênfase. A tentação de "já deixar preparado" é como o Eras do
Brasil morreu — sistema entrando por antecipação. A porta fica aberta pelas cinco
regras acima, que custam quase nada. Abstração antecipada custa caro e quase
sempre erra a forma.

## Simplificação didática

- balanceamento entre zonas (uma zona lotada e outra vazia) não está pensado
- consistência entre zonas adjacentes num gank é o problema difícil de verdade
- reconexão durante um encontro
- anti-cheat além de "o servidor decide"

## O que vira ADR-005

Só quando o multiplayer chegar. Até lá, este documento é uma lista do que não
fazer.
