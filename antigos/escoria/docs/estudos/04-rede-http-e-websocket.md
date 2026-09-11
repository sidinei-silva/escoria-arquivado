---
tags: [estudo, arquitetura, rede, go]
fatia: 1, 3
---

# 04 — Rede: HTTP e WebSocket

## Vindo do Express

O modelo que eu já tenho na cabeça:

```js
const app = express()
app.get('/rota', (req, res) => res.json({...}))
```

Em Go a peça equivalente é a `net/http` da biblioteca padrão. Não precisa de
framework:

```go
mux := http.NewServeMux()
mux.HandleFunc("POST /personagens", h.criarPersonagem)
http.ListenAndServe(":8080", mux)
```

Diferenças que importam:

- **`ResponseWriter` é um stream, não um objeto de resposta.** Escreveu, foi.
  Chamar `WriteHeader` duas vezes é bug. Não existe "montar a resposta e mandar
  no final" a não ser que eu faça isso à mão.
- **Cada request roda na própria goroutine.** O `net/http` faz isso sozinho. Ou
  seja: **todo handler já é código concorrente**, mesmo que não pareça. Se dois
  handlers tocarem o mesmo mapa, é race. Esse é o motivo de
  [[02-ownership-e-concorrencia]] existir.
- **Middleware é `func(http.Handler) http.Handler`**, não `(req, res, next)`.
  Composição por wrapping.
- **`context.Context` viaja no request** e cancela quando o cliente desiste.

## A divisão de trabalho

Não é "HTTP para umas coisas, WS para outras" por gosto. A pergunta é: **isso é
pedido-resposta ou é fluxo contínuo?**

### HTTP — pré-jogo e pedido-resposta

- criar conta, login
- criar personagem, listar personagens
- entrar no servidor
- consultas pontuais que não precisam de push

Característica comum: **nada disso toca o estado autoritativo do jogo em curso**.
São operações sobre conta e persona, contra o banco, antes de o jogador existir
dentro do game core.

Isso é uma fronteira limpa e vale explicitá-la, porque simplifica muito o começo:
a Fatia 1 pode ser HTTP puro contra o banco, sem game loop nenhum. O loop entra
quando o jogador **entra** no servidor.

### WebSocket — gameplay

- comandos de ação: iniciar coleta, equipar, viajar, cancelar
- eventos do servidor: ciclo concluído, Fama ganha, Têmpera mudou, encontro

Característica comum: **o servidor precisa falar sem ser perguntado**. Um ciclo
de coleta fecha 12 segundos depois do pedido. Com HTTP puro, a única saída é o
cliente ficar perguntando.

## Por que não polling em REST

Funciona? Funciona. Mas:

- **latência é o intervalo do poll.** Poll de 2s significa que "a coleta
  terminou" chega até 2 segundos atrasado, e o jogo parece travado.
- **custo cresce com jogadores ociosos.** Idle tem muita gente com a janela
  aberta sem fazer nada. Polling faz custo proporcional a presença, não a
  atividade — exatamente o oposto do que eu quero.
- **quando o gank chegar, não tem escolha.** Um encontro é iniciado pelo
  servidor. O cliente não sabe que precisa perguntar.

Não é sobre elegância. É sobre o servidor ser quem inicia a conversa.

## A camada de rede não conhece regra

Isso vale para as duas portas:

```
JSON → Command → [game core] → Event → JSON
```

A rede traduz. Não valida regra de jogo, não decide resultado, não conhece Fama.
Validação de formato — campo faltando, ID malformado — é dela. Validação de
regra — "pode coletar nesta zona?" — é do core.

O teste do `RECEITA` se aplica: se eu trocasse WebSocket por outra coisa, isso
continuaria valendo? Sim → é regra, fica no core. Não → é encanamento, fica na
rede.

## Uma armadilha de WebSocket em Go

A que mais aparece na prática, vale saber antes:

**Uma conexão WebSocket não suporta duas goroutines escrevendo ao mesmo tempo.**

E o desenho natural leva direto a isso: uma goroutine lendo comandos do cliente,
outra empurrando eventos do servidor. Se as duas escreverem, corrompe o frame.

O padrão comum: **uma goroutine de escrita por conexão**, com um channel de saída.
Quem quer mandar evento manda para o channel; só ela escreve no socket.

Repare que é o mesmo princípio de [[02-ownership-e-concorrencia]] em escala
menor: a conexão tem dono.

E aí aparecem as perguntas que eu tenho que responder:

- o que acontece se o cliente parar de ler e o channel encher?
- o jogador desconectou ou só está lento?
- na desconexão, a atividade em curso continua?

Nenhuma dessas é técnica pura — a terceira é regra de jogo, e vai para o GDD.

## Ordem de construção

Segue as fatias:

1. **Fatia 1** — HTTP puro. Conta, login, personagem. Sem game loop.
2. **Fatia 3** — entra o loop e o WebSocket, porque a coleta precisa de push.
3. Daí em diante, comando novo é caso novo no despacho, não código de rede novo.

Se eu me pegar escrevendo muito código de rede depois da Fatia 3, alguma regra
vazou para a camada errada.

## Simplificação didática

- autenticação de verdade (token, expiração, refresh) não está resolvida aqui
- reconexão e reidratação de estado: o `BACKLOG` pede reidratação na Fatia 1 por
  HTTP; com WS isso vira "o que mando no connect?"
- heartbeat/ping-pong e detecção de conexão morta
- rate limiting por conexão

## O que vira ADR-004

HTTP para pré-jogo e pedido-resposta; WebSocket para gameplay. A camada de rede
traduz protocolo em comando e não conhece regra de jogo. Uma goroutine de escrita
por conexão.
