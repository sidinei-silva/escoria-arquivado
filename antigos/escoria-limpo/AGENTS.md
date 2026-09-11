# AGENTS.md

> Regras de uso de IA em **A Escória**. Valem para qualquer ferramenta: Copilot,
> ChatGPT, Claude, o que vier depois. Se uma ferramenta não lê este arquivo, a
> regra continua valendo — ela é minha, não dela.
>
> Ferramentas que leem `CLAUDE.md` em vez de `AGENTS.md`: criar um `CLAUDE.md`
> com uma linha só — `Ver AGENTS.md.`

## Por que este arquivo existe

Este é um projeto pessoal. O objetivo dele **não é** entregar rápido. O objetivo
é escrever código que é meu, aprender Go de verdade e voltar a gostar de
programar.

Já perdi dois projetos por esquecer isso. No Eras do Brasil, agentes escreveram
parte significativa do servidor, eu fiquei sem entender as decisões, e apaguei o
repositório. Duas vezes. Este arquivo existe para isso não acontecer uma
terceira.

Velocidade aqui não é qualidade. Entender é.

## O teste

Antes de aceitar qualquer coisa vinda de uma IA:

> **Eu consigo explicar isso amanhã, sem olhar?**

Se não, não entra. Não importa se funciona, se é idiomático, ou se economizaria
duas horas.

## Permitido

- Explicar um conceito que eu não conheço
- Comparar duas ou mais abordagens, com os problemas que cada uma resolve
- Pesquisar (documentação, prática comum, erro obscuro)
- **Code review de código que eu escrevi**
- Ajudar a depurar: ler stack trace, sugerir onde investigar, explicar o erro
- Revisar texto que eu escrevi — pontuação, clareza, concordância
- Traduzir

## Proibido

- Gerar arquivo
- Gerar função para eu colar
- Modo agente, modo plan que produz a implementação, edição multi-arquivo
- Gerar prosa de lore ou de GDD que vá para `gdd/`
- Escrever regra numerada (`R-XXX-NN`) por mim
- Abrir spec ou ADR por mim

## Autocomplete

Permitido: completar no padrão que **eu** já estabeleci no arquivo. Fechar uma
função cuja forma eu defini duas funções acima não é a IA decidindo nada — é
economia de digitação sobre decisão minha.

**Exceção, e essa importa:** na primeira vez que eu encostar num conceito de Go
que ainda não é meu, **desligo inline suggestions naquela sessão** e escrevo
errado sozinho. Depois que o padrão for meu, ligo de volta.

A lista de "ainda não é meu" hoje:

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

Se cheguei no 5, escrevo depois o que eu não sabia. Isso é candidato a post no
devlog.

## Perguntas que eu faço antes de escrever concorrência

Sempre. Não é cerimônia — é o único jeito de a decisão ser minha.

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

Ver [02-ownership-e-concorrencia](docs/estudos/02-ownership-e-concorrencia.md).

## Ideias vindas de outros jogos

Vão para [IDEIAS](IDEIAS.md). **Não** para o GDD, **não** para o backlog.

Foi assim que o Eras do Brasil morreu: sistema legal de outro jogo entrava
direto no escopo. O GDD cresce por descoberta durante a construção, nunca por
antecipação — ver `docs/RECEITA.md`.

Se depois da PoC a ideia ainda importar, ela era real.

## GDD e lore

Regra numerada eu escrevo. Prosa de lore eu escrevo — **depois** do código da
fatia funcionar, nunca antes — é o passo 6.3 da [receita](docs/RECEITA.md). IA pode revisar
pontuação e clareza depois de pronta.

A lore descrever algo que já roda é mais fácil e mais gostoso do que inventar no
vazio. E é o que faz o jogo virar meu aos poucos, sem reset.

## No trabalho é diferente

Na empresa: plan mode para descobrir o fluxo do sistema, e eu escrevo o código.
Lá existe pressão por velocidade e a linha vai ceder às vezes.

**Aqui não cede.** É a única que só depende de mim.
