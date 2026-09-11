---
tags: [estudo, arquitetura, gameloop]
fatia: 3, 4, 5
---

# 03 — Tempo e atividades

> Se [[02-ownership-e-concorrencia]] é a decisão mais cara de reverter, esta é a
> mais característica do gênero. Idle **é** um jogo sobre tempo.

## Ação não é instantânea

Quase nada em A Escória acontece no momento do pedido:

| Ação | Duração |
|---|---|
| Coletar | ciclo fixo, repete |
| Combate automático | ciclos, com fila de habilidades e cooldowns |
| Refino / craft | tempo por receita |
| Viagem | trânsito entre zonas |

Isso muda a forma da API. Não existe `POST /coletar` que devolve o minério.
Existe:

```
IniciarColeta       → aceita, ou recusa com motivo
(tempo passa)
ColetaConcluída     → evento com recurso, quantidade, Fama e o nó que recebeu
```

E o `BACKLOG` já diz isso na Fatia 3: *"Iniciar uma ação — o servidor é a
autoridade do tempo"*.

## Duas formas de representar tempo

Essa é a escolha real, e ela tem consequência grande.

### A — contador que decrementa

```go
type Activity struct {
    Restante time.Duration
}
// a cada tick: a.Restante -= delta
```

Parece natural vindo de game loop clássico. **Não use.** Três problemas:

- depende de o loop rodar sem falhar; um GC pause ou um tick perdido e o jogo
  atrasa
- não sobrevive a restart: o valor salvo é relativo a um "agora" que passou
- não dá para responder "quanto falta?" sem o loop ter rodado

### B — marcos absolutos

```go
type Activity struct {
    Tipo      string
    AlvoID    string
    IniciadaEm time.Time
    TerminaEm  time.Time
}
```

O loop não decrementa nada. Ele **verifica**: `if now.After(a.TerminaEm)`.

Por que é melhor:

- **tick perdido não atrasa o jogo.** Se o loop travar 2 segundos, na volta ele
  fecha tudo que venceu. O tempo é o relógio, não o loop.
- **sobrevive a restart de graça.** `TerminaEm` salvo no banco continua
  significando a mesma coisa depois de reiniciar.
- **"quanto falta" é subtração**, sem estado auxiliar.
- **o tick vira decisão de latência, não de correção.** 100ms ou 1s muda só a
  precisão percebida, nunca o resultado.

Essa última linha é a que importa mais: com A, o intervalo do tick é
correção. Com B, é qualidade de serviço.

## O loop

```go
func (e *Engine) advance(now time.Time) {
    for _, p := range e.state.Jogadores {
        if p.Atividade == nil {
            continue
        }
        if now.Before(p.Atividade.TerminaEm) {
            continue
        }
        e.concluirAtividade(p, now)
    }
}
```

Duas observações:

**Varrer todos os jogadores todo tick não escala** — mas escala muito mais do que
parece, porque a comparação é barata. Quando não escalar, a solução é fila de
prioridade por `TerminaEm`, e aí só olho o topo. Otimização para quando o número
existir, não agora.

**`now` entra como parâmetro, não como `time.Now()` dentro.** Isso deixa o loop
testável: injeto um instante e verifico a transição, sem `sleep` no teste. Vale
para toda a lógica que depende de tempo.

## Ciclo que reinicia

Coleta e combate não são um evento único — são ciclo. O `BACKLOG` da Fatia 3 diz:
*"ao fechar um ciclo, reinicia sozinho"*.

Ao concluir:

1. aplicar o resultado (recurso, loot, Fama no nó certo)
2. emitir o evento
3. verificar se ainda pode continuar — inventário cheio? nó esgotou? jogador
   mandou parar?
4. se pode, agendar o próximo ciclo com novos `IniciadaEm`/`TerminaEm`

O passo 3 é onde as regras do GDD entram, e é onde vão nascer regras que ninguém
escreveu ainda — *"e se o inventário encher no meio do ciclo?"*. Isso é regra
descoberta: volta para o GDD com ID novo, como manda a `RECEITA`.

## Têmpera é uma regra sobre tempo

Vale isolar porque é a mecânica mais própria do jogo e depende inteiramente
disto:

- sobe **apenas com ação ativa** da arma equipada → incrementa na conclusão do
  ciclo, não no tick
- **congela offline** → não é função de tempo decorrido; é função de ciclos
  concluídos
- **reseta na troca de arma** → é estado da arma equipada, não do jogador
- **multiplica ganho de Fama, não dano**

"Congela offline" é o motivo de Têmpera não poder ser calculada como
`agora - últimaAtualização`. Se um dia eu me pegar querendo escrever isso, é
sinal de que a implementação divergiu da regra.

E é isso que faz o jogo ser sobre presença, não sobre calendário — a promessa do
GDD. A arquitetura precisa expressar isso, não contornar.

## Cliente não conta tempo

O cliente **desenha** uma barra de progresso interpolando entre `IniciadaEm` e
`TerminaEm`. Ele não decide quando termina.

Consequência: `TerminaEm` vai no evento. O cliente anima localmente e confirma
quando o evento de conclusão chega. Se divergir, o servidor ganha, sempre.

Isso já é o desenho certo para o gank futuro, onde "eu escapei" nunca pode ser
afirmação do cliente.

## Simplificação didática

- relógio: `time.Now()` monotônico vs wall clock importa se o servidor mudar de
  hora; ignorar por ora
- fila de prioridade em vez de varredura, quando o número exigir
- offline: o GDD diz que a recompensa offline é mínima, mas *o que exatamente*
  acontece com uma atividade em curso quando o jogador cai ainda não está
  decidido — regra a descobrir na Fatia 3
- fila de habilidades com cooldowns individuais (Fatia 8) é mais complexa que
  ciclo fixo e merece uma segunda passada neste documento

## O que vira ADR-002

Atividade é representada por marcos absolutos, não por contador. O servidor é a
autoridade do tempo. O tick é decisão de latência, não de correção.
