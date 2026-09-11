---
tags: [estudo, arquitetura, concorrencia, go]
fatia: 3, 8
---

# 02 — Ownership e concorrência

> O documento mais importante da pasta. É onde os dois estudos originais
> divergiam, e é a decisão que mais custa reverter.

## A pergunta errada e a certa

**Errada:** "mutex, channel ou actor?"
**Certa:** "quem é dono deste estado?"

Mutex, channel e actor não são três sabores da mesma coisa. São respostas a
perguntas diferentes:

| Ferramenta | Responde |
|---|---|
| **mutex** | várias goroutines tocam o mesmo dado; como não se atropelam? |
| **channel** | uma goroutine precisa mandar mensagem para outra; como? |
| **actor** | quem é dono é uma goroutine só; como falam entre si? |

Escolher por "ser idiomático" é escolher a resposta antes de saber a pergunta. As
perguntas do [[../AGENTS|AGENTS.md]] existem para isso.

## As duas estratégias

Todo estado compartilhado cai numa destas duas. Não existe terceira.

### A — estado compartilhado, protegido

```
goroutine A ─┐
goroutine B ─┼──→ 🔒 GameState
goroutine C ─┘
```

Várias goroutines acessam o mesmo estado. Mutex garante uma por vez.

Vantagem: paralelismo real entre cores.
Custo: **toda** goroutine precisa lembrar de travar. Esquecer uma vez é race
condition silenciosa. E quando uma operação toca duas entidades, aparece ordem de
lock — e com ela, deadlock.

### B — estado possuído por uma goroutine

```
goroutine A ─┐
goroutine B ─┼──→ commands ──→ Game Loop ──→ GameState
goroutine C ─┘                  (dono)
```

Ninguém compartilha estado. Quem quer mudar algo **manda mensagem**. Uma goroutine
só aplica.

Vantagem: race é estruturalmente impossível no estado autoritativo. Não tem lock
a esquecer.
Custo: serializa. Um comando lento segura a fila.

## A decisão: B, single-owner

E o motivo não é "é mais simples para a PoC" — é mais simples, mas isso sozinho
não decidiria nada, porque a PoC não é o destino.

O motivo real:

> **A Escória particiona por zona. E single-owner é o caso N=1 dessa partição.**

Olha o que o GDD pede no estado final:

- o jogador está sempre em **exatamente uma** zona
- a zona define quais ações existem nela
- gank procura alvo na **mesma zona ou adjacente**
- viagem é sair de uma zona e entrar em outra, com trânsito

Isso é uma descrição de partição. O jogo já está dividido por zona no design; a
arquitetura só precisa não brigar com isso.

O destino natural, no multiplayer, é **cada zona ser dona do seu próprio estado**:
processa os ticks dos jogadores que estão nela, resolve o matchmaking de gank
local, e conversa com zonas vizinhas por mensagem. Travel vira handoff entre
donos.

E aqui está o ponto que decide:

**Um único loop dono de tudo é literalmente esse desenho com uma zona.** Crescer
significa trocar o roteamento — de "tudo vai para o loop" para "cada comando vai
para o dono da zona do jogador". A lógica de jogo não muda: validar, transicionar,
emitir continua idêntico.

A estratégia A **não** cresce para lá. Ela cresce para lock ordering — e o
primeiro lugar onde isso dói é exatamente o gank, que toca dois jogadores e uma
ou duas zonas ao mesmo tempo. Ou seja: o mutex fica confortável durante toda a
PoC e vira problema exatamente quando o recurso mais característico do jogo
chegar.

Escolher B agora é comprar a opção de crescer sem reescrever.

## O contra-argumento honesto

Single-owner serializa tudo. Isso não é gratuito.

Mas faz as contas do caso concreto: por tick, por jogador, o trabalho é comparar
`time.Now()` com `EndsAt`, e quando fecha, sortear loot e somar Fama. São
dezenas de nanossegundos. Milhares de jogadores num loop de 100ms cabem
folgadamente em um core.

O que **não** pode entrar no loop: I/O. Query no banco, escrita em socket,
chamada externa. Isso sai por channel e é feito por outra goroutine. Se eu puser
um `db.Query` dentro do loop, o jogo inteiro trava esperando o Postgres — e aí a
crítica ao single-owner vira verdadeira por culpa minha, não da arquitetura.

Regra de bolso: **o loop só toca memória.**

## Onde mutex ainda é a ferramenta certa

Não estou banindo mutex. Ele sai do estado de jogo e continua onde é a resposta
certa:

- **registro de conexões WebSocket** (`map[jogadorID]*conn`) — várias goroutines
  registram e removem; ninguém é dono natural; mutex é exatamente isso
- **métricas e contadores** — ou `sync/atomic`
- **cache do worker de persistência**, se existir

O que nunca leva mutex: conteúdo estático de `data/`, porque é imutável depois do
boot ([[01-modelo-mental]]).

## Actor model

É o destino, não o começo. "Uma goroutine por zona, conversando por channel" é
actor model — só que eu vou chegar lá por necessidade, com o desenho pronto, em
vez de adotar o nome antes de ter o problema.

Não adotar agora: com uma zona ativa, actor é cerimônia sem ganho.

Ver [[06-zonas-presenca-e-pvp]].

## A forma

> Ilustração de forma. Não copiar.

```go
type Engine struct {
    state    *GameState        // só o loop toca
    commands chan Command      // entrada
    events   chan Event        // saída
}

func (e *Engine) Run(ctx context.Context) {
    tick := time.NewTicker(100 * time.Millisecond)
    defer tick.Stop()

    for {
        select {
        case <-ctx.Done():
            return
        case cmd := <-e.commands:
            e.apply(cmd)
        case now := <-tick.C:
            e.advance(now)      // atividades que venceram
        }
    }
}
```

Quatro coisas para eu conseguir explicar antes de escrever isso:

1. **Por que `select` e não duas goroutines?** Porque é o `select` que garante o
   dono único. Duas goroutines tocando `e.state` seria a estratégia A disfarçada.
2. **Por que `ctx.Done()` primeiro?** Não é prioridade — `select` escolhe
   aleatoriamente entre casos prontos. É convenção de leitura. Shutdown de
   verdade precisa de mais que isso.
3. **O que acontece se `commands` encher?** Buffered ou não? Quem manda, bloqueia
   ou desiste? Essa decisão é minha e vira ADR.
4. **O tick de 100ms é chute.** Tem que sair de um requisito do jogo, não de um
   número redondo.

## Perguntas antes de escrever concorrência

Repetidas do [[../AGENTS|AGENTS.md]] porque é aqui que valem:

- Quem é dono deste estado?
- Quem pode modificá-lo?
- Quem só precisa receber mensagem?
- Por que existe esta goroutine?
- Por que existe este mutex?
- O que acontece numa race condition aqui?
- O que acontece quando o cliente desconecta?
- O que acontece no shutdown?

E a ferramenta que responde de verdade: **`go test -race`**. Rodar sempre. Race
detector acha o que revisão não acha.

## Simplificação didática

- shutdown de verdade é mais que `ctx.Done()`: drenar a fila, persistir, fechar
  conexões, esperar workers
- backpressure quando `commands` enche não está resolvido
- não existe prioridade entre comandos
- métricas de latência do loop (quanto tempo um tick leva) são o que vai me dizer
  se o single-owner está segurando — e isso não existe aqui

## O que vira ADR-001

Quando a Fatia 3 chegar, escrever com as minhas palavras: o estado autoritativo
tem dono único; comandos chegam por channel; mutex fica fora do estado de jogo;
o desenho é compatível com partição por zona no futuro. Contexto, decisão,
consequência. Três parágrafos.
