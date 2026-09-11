# 001 — Personagem e Chegada

**Fatia 1** · Passo 1 do tutorial (*A Chegada*) · Status: pronta para implementar

## Entrega

O jogador cria um personagem, chega em **A Ressaca**, ouve a primeira fala do Guia e vê as
ações disponíveis da zona. Ele ainda não consegue executar nenhuma ação — apenas vê que
existem.

Ao terminar esta fatia, existe um jogo que se pode abrir.

## Por que esta fatia vem primeiro

É o primeiro contato do jogador com a Escória. Estabelece três coisas que todas as fatias
seguintes reusam: **identidade** (existe um jogador com estado persistido), **lugar** (o
jogador está sempre em uma zona) e **zona define a ação** (Pilar 3 — a tela de zona é a
interface principal do jogo, não um menu global).

O jogador não recebe arma aqui. Receber e equipar a primeira arma é a Fatia 2. Esta fatia
termina com ele parado na praia, olhando as opções — que é exatamente a sensação de chegada
que a narrativa pede.

## Escopo

### Dentro

- Criar personagem informando um nome
- Persistir estado do jogador (id, nome, zona atual)
- Colocar o jogador em A Ressaca na criação
- Listar as ações disponíveis da zona atual
- Exibir a fala de abertura do Guia (Passo 1)

### Fora

- Executar qualquer ação (coleta, combate, viagem) → Fatias 3, 4, 5
- Inventário → Fatia 4
- Equipamento e loadout → Fatia 2
- Autenticação real (senha, token, registro) — a PoC usa sessão mock
- Fama, Litania, Têmpera → Fatias 3 e 9
- Bark de espera do Passo 1 (não existe — ver `dados.md`)

## Critérios de aceite

Verificáveis à mão. Cada um deve funcionar de ponta a ponta antes da fatia fechar.

1. **Criação.** Com o cliente aberto na tela inicial, informar um nome e confirmar cria um
   jogador e leva à tela de zona. Recarregar a página mantém o jogador (a sessão sobrevive).

2. **Nome obrigatório.** Confirmar com o campo vazio não cria jogador e mostra erro. Nomes
   aceitos: 3 a 20 caracteres.

3. **Zona inicial.** Um jogador recém-criado está em **A Ressaca**. Nunca em outra zona,
   nunca sem zona.

4. **Tela de zona.** A tela mostra o nome da zona e a lista de ações disponíveis nela, vindas
   do servidor — não fixas no cliente. As ações aparecem desabilitadas ou inertes; clicar não
   precisa fazer nada ainda.

5. **Fala do Guia.** Na primeira vez que o jogador chega em A Ressaca, a fala de abertura do
   Passo 1 é exibida. Recarregar a página **não** a exibe de novo.

6. **Estado consultável.** `GET /player/state` devolve o estado atual do jogador e é a fonte
   que o cliente usa para se reidratar ao abrir.

7. **Zona desconhecida.** Pedir uma zona que não existe em `data/zones.json` responde 404,
   não 500.

## Ligações com o GDD

| O quê | Onde |
| --- | --- |
| Narrativa do Passo 1 | `gdd/10-narrativa/narrativa-jogada-poc.md` |
| Falas do Guia (Passo 1) | `data/tutorial/dialogs.json` → `dialog_passo1_abertura` |
| Ficha do Guia | `gdd/10-narrativa/personagens.md` |
| A Ressaca | `gdd/20-mundo/zonas-poc.md` |
| Pilar 3 (zona define a ação) | `gdd/30-mecanicas/pilares.md` |

## Notas de direção

**A tela de zona é o jogo.** Não é um menu de navegação — é onde o jogador vive. Vale investir
nela agora, porque todas as fatias seguintes penduram coisas aqui (ação em andamento,
inventário, loadout, progresso do tutorial). Deixe espaço.

**A fala do Guia não é um popup de tutorial.** É diálogo diegético. Evite linguagem de
onboarding ("Bem-vindo!", "Clique aqui para começar"). O Guia não está te ensinando a jogar —
ele está te recebendo com má vontade profissional.

**Mostrar ações que ainda não funcionam é intencional.** O jogador vê o que a zona oferece
antes de poder fazer. Isso comunica o Pilar 3 desde o primeiro segundo e dá continuidade
visível quando as fatias seguintes ligarem cada ação.

## Questões abertas

Nenhuma bloqueante. Duas para decidir na implementação:

- A fala do Guia aparece sobreposta à tela de zona ou numa tela de chegada anterior?
- O nome do jogador é exibido em algum lugar da tela de zona nesta fatia, ou só depois?
