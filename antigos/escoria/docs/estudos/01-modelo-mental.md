---
tags: [estudo, arquitetura, fundamentos]
fatia: 0, 1
---

# 01 — Modelo mental

## O erro que este documento evita

Vindo de Node e Express, o instinto é este:

```
request → controller → service → banco → response
```

Isso funciona para CRUD. **Não funciona para A Escória**, e a razão é uma só:

> Num idle, o tempo é o sistema.

Coletar leva 12 segundos. Combate roda em ciclos. Viagem tem trânsito. Têmpera
sobe com uso ativo e congela offline. Nada disso é "o cliente pediu, o servidor
respondeu". O servidor precisa continuar existindo e agindo **entre** os pedidos.

Se eu construir CRUD e depois tentar encaixar tempo, eu reescrevo o servidor. Por
isso este documento vem antes de qualquer código.

## As três palavras

Todo o resto se apoia nelas.

### Command — intenção do jogador

O cliente **nunca** afirma um resultado. Ele expressa vontade.

Errado: `POST /inventory { add: "minério", qtd: 3 }`
Certo: `IniciarColeta{ Jogador, Nó }` — e o servidor decide se pode, quanto
demora, e o que sai.

Isso não é purismo. É o que o GDD exige: se o cliente pudesse afirmar resultado,
Fama, Lastro e Têmpera seriam trivialmente falsificáveis. E quando o gank chegar,
"eu escapei" viraria uma afirmação do cliente.

### State — a verdade

Existe um estado autoritativo, no servidor, e ele é a única verdade. O cliente
tem uma **cópia atrasada** para desenhar tela.

### Event — o que aconteceu

Resultado de um comando aplicado. Vai para o cliente para ele atualizar a cópia
dele.

Evento é fato consumado. `ColetaConcluída{ Recurso, Qtd, FamaGanha, Nó }`.

## O ciclo

```
Command → Validar → Transicionar State → Emitir Events
```

Sempre nessa ordem, sem atalho. Validar antes de mudar. Mudar antes de emitir.
Se a validação falha, o estado não muda e o evento é de erro.

O que isso me dá de graça:

- **Testabilidade** — comando entra, estado e eventos saem. Sem HTTP, sem banco.
  É a `RECEITA` funcionando: `R-COL-01` vira um teste que aplica um comando e
  verifica o estado.
- **Reprodutibilidade** — um bug vira uma sequência de comandos.
- **Uma porta só** — HTTP e WebSocket viram tradutores para comando. Não duas
  lógicas.

## O desenho

```
                    CLIENT
                       │
                       ▼
                    NETWORK           traduz protocolo em Command
                  HTTP / WS
                       │
                       ▼
                   COMMANDS
                       │
                       ▼
                  GAME CORE           valida, aplica, emite
                       │
              ┌────────┴────────┐
              ▼                 ▼
         GAME STATE          EVENTS
              │                 │
              ▼                 ▼
         PostgreSQL          CLIENT
```

Ler de cima para baixo: **intenção desce, verdade sobe**.

Duas leituras importantes desse desenho:

**Network não conhece regra de jogo.** Ela sabe traduzir JSON em `Command` e
`Event` em JSON. Se eu trocar WebSocket por outra coisa, o game core não muda.

**PostgreSQL fica embaixo, não no meio.** O banco é consequência do estado, não
dono dele. Ver [[05-persistencia]].

## Os tipos, em forma

> Ilustração de forma. Não copiar. Ver o aviso no [[README]].

```go
// O que o jogador quer
type Command interface{ isCommand() }

type IniciarColeta struct {
    JogadorID string
    NóID      string
}

// O que aconteceu
type Event interface{ isEvent() }

type ColetaConcluída struct {
    JogadorID string
    Recurso   string
    Qtd       int
    Fama      int
}
```

Por que interface vazia com método privado? É o jeito Go de fazer união de tipos:
só quem está neste pacote pode implementar `Command`. `switch cmd := c.(type)`
resolve o despacho.

Existem alternativas — struct com campo `Tipo`, ou generics. Vale eu comparar
quando chegar na Fatia 3, e escolher. Não aceitar esta forma só porque está
escrita aqui.

## Conteúdo estático vs estado vivo

Distinção que evita muito problema:

| | Conteúdo | Estado |
|---|---|---|
| O que é | zonas, receitas, mobs, armas, nós da Litania | jogadores, inventários, atividades em curso, Fama |
| De onde vem | JSON em `data/` | comandos aplicados |
| Muda em runtime | **não** | sim, o tempo todo |
| Concorrência | nenhuma — imutável após boot | ver [[02-ownership-e-concorrencia]] |
| Persistência | é o repositório | é o Postgres |

Conteúdo carrega uma vez no boot, vira struct tipada em memória, e **nunca muda**.
Sem mutex, sem lock, sem cache. Muitas goroutines lendo algo que ninguém escreve é
seguro em Go.

Isso é metade do problema de concorrência resolvido de graça — desde que eu
respeite a imutabilidade. Se um dia eu escrever num mapa de conteúdo, quebrei a
garantia.

Fatia 0 é exatamente isto: subir, carregar `data/`, expor tipado, confirmar que
carregou.

## Simplificação didática

O que este documento está deixando de fora de propósito:

- validação e erro de verdade (aqui é `isCommand()` e pronto)
- versionamento de eventos, que importa quando o cliente for publicado
- o custo de `interface{}` e type switch com muitos tipos de comando

Ver [[README]] para o que é estudo e o que vira ADR.
