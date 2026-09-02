# AGENTS.md

> Regras de uso de IA em **A Escória**. Valem para qualquer ferramenta: Copilot,
> ChatGPT, Claude, o que vier depois. Se uma ferramenta não lê este arquivo, a
> regra continua valendo — ela é minha, não dela.

## Por que este arquivo existe

Este é um projeto pessoal. O objetivo **não é** entregar rápido. É escrever
código que é meu, aprender Go de verdade e voltar a gostar de programar.

Já perdi dois projetos por esquecer isso. No Eras do Brasil, agentes escreveram
parte do servidor, eu fiquei sem entender as decisões, e apaguei o repositório.
Duas vezes.

Velocidade aqui não é qualidade. Entender é.

## O teste

Antes de aceitar qualquer coisa vinda de uma IA:

> **Eu consigo explicar isso amanhã, sem olhar?**

Se não, não entra. Não importa se funciona, se é idiomático, ou se economizaria
duas horas.

## Permitido

- Explicar um conceito que eu não conheço
- Comparar duas ou mais abordagens, com o problema que cada uma resolve
- Pesquisar: documentação, prática comum, erro obscuro
- **Code review de código que eu escrevi**
- Ajudar a depurar: ler stack trace, sugerir onde investigar, explicar o erro
- Revisar texto que eu escrevi — pontuação, clareza, concordância
- Traduzir

## Proibido

- Gerar arquivo
- Gerar função para eu colar
- Modo agente, modo plan que produz a implementação, edição multi-arquivo
- Gerar prosa de lore ou de GDD
- Escrever regra numerada (`R-XXX-NN`) por mim
- Escrever fluxo, ADR ou spec por mim
- **Deixar material didático de IA morar no repo.** Explicação é conversa, não
  artefato versionado. O que sobreviver vira ADR com as minhas palavras.

## Autocomplete

Permitido: completar no padrão que **eu** já estabeleci no arquivo. Fechar uma
função cuja forma eu defini duas funções acima não é a IA decidindo nada — é
economia de digitação sobre decisão minha.

**Exceção:** na primeira vez que eu encostar num conceito de Go que ainda não é
meu, **desligo inline suggestions naquela sessão** e escrevo errado sozinho.
Depois que o padrão for meu, ligo de volta.

Lista do que ainda não é meu:

- goroutines e ownership de estado
- channels (buffered, select, fechamento)
- mutex e quando ele é a ferramenta certa
- `context` e cancelamento
- `defer` e ordem de execução
- tratamento de erro idiomático (wrapping, sentinel, `errors.Is/As`)
- receivers: valor vs ponteiro
- graceful shutdown

O autocomplete é mais confiante justamente onde eu sou mais fraco. O risco não é
código que eu não entendo — é **decisão que eu nunca precisei tomar**.

## Se eu estiver travado

Escalar aos poucos, nesta ordem. Não pular etapas:

1. Reformular o problema em uma frase
2. Perguntar qual conceito está faltando
3. Pedir um exemplo **em outro domínio**, não no meu código
4. Pedir para apontar a linha errada, sem a correção
5. Só então pedir a correção

Se cheguei no 5, escrevo depois o que eu não sabia. Vira post no devlog.

## Perguntas antes de escrever concorrência

Sempre. É o único jeito de a decisão ser minha.

- Quem é dono deste estado?
- Quem pode modificá-lo?
- Quem só precisa receber mensagem?
- Por que existe esta goroutine?
- Por que existe este mutex?
- Por que isso precisa ser um channel?
- O que pertence ao game loop?
- O que acontece numa race condition aqui?
- O que acontece quando o cliente desconecta?
- O que acontece no shutdown?

## Texto que o jogador lê

Nenhuma frase na tela pode ser gerada por IA — nem copy de interface, nem lore,
nem descrição de item.

E há um defeito específico a vigiar, que eu identifiquei jogando MU Idle: **texto
que sabe demais sobre a implementação.** Frases como *"eco do servidor"*,
*"a grade não antecipa o resultado"*, *"estado sincronizado"*. É o que acontece
quando quem escreve a copy tem a spec técnica em contexto e descreve o mecanismo
em vez do objetivo do jogador. Denuncia origem em IA e afasta quem lê.

O teste é o mesmo da receita: **se a frase muda quando eu troco WebSocket por
polling, ela é encanamento e não pode estar na tela.**

A arquitetura autoritativa está certa. O jogador nunca pode ficar sabendo dela
por escrito.

## Lore

Regra numerada eu escrevo. Prosa de lore eu escrevo — **depois** do código da
fatia funcionar, nunca antes. IA pode revisar pontuação e clareza depois de
pronta.

Pedir correção, não reescrita. Se ela reescrever o parágrafo inteiro e eu colar,
virou geração — e eu não vou reconhecer o texto, exatamente como aconteceu com o
código.

## Ideias

Sistema que a IA sugerir vai para `docs/ideias/`, como qualquer outra ideia.
Nunca direto para fluxo, sistema ou backlog.

## Mensagens de commit

Conventional Commits, descrição em pt-BR. Formato e exemplo em
[.github/commit-instructions.md](.github/commit-instructions.md).

Commit message é documentação que eu vou ler. É como eu vou descobrir, em março,
o que aconteceu numa noite de setembro.

O botão gera, **eu reescrevo**. O Copilot vê o diff — ele não sabe que
`R-COL-03` nasceu ali, nem que aquilo é a Fatia 4. Gerar e commitar sem ler
transforma o `git log` em ruído.

**Conventional Commits, com a descrição em pt-BR.**

Formato: `tipo(escopo): descrição`

- Tipos em inglês, que é o padrão: `feat`, `fix`, `docs`, `refactor`, `test`,
  `chore`
- **Descrição em português**, imperativo, minúscula, sem ponto final
  ✅ `feat(coleta): adiciona ciclo com nó de recurso`
  ❌ `feat(coleta): Adicionado o ciclo de coleta.`
- Escopos deste repo: `fluxo`, `sistema`, `data`, `adr`, `server`, `web`, `docs`

**O tipo é o que a fatia entrega**, não o arquivo que mais mudou. Fatia que
entrega coleta é `feat`, mesmo que 80% das linhas sejam markdown.

**A fatia inteira num commit só.** A receita manda fechar fluxo, regra, conteúdo,
código, teste e lore juntos. Então o assunto resume a entrega e o corpo lista as
peças: