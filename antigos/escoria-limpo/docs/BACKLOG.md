# A Escória — Backlog da PoC

> Organizado por **fatias verticais** seguindo o fluxo do jogador, não por camada.
> Cada fatia termina em algo que o jogador **vê ou faz**. As fatias mapeiam os 8 passos
> do tutorial, que vivem em [gdd/README.md](../gdd/README.md).
>
> **Títulos descrevem capacidade, não implementação** (ADR-008). Nomes de rota, tabela e
> função são decisão do implementador. A exceção são os arquivos de `data/`, cujos nomes são
> acordados porque a IA os produz e o código os consome.
>
> Destino: GitHub Issues. Cada `- [ ]` vira uma issue; cada fatia, um milestone.
>
> **Estado em 31/08/2026:** nada de código existe. De `data/`, só `tutorial/dialogs.json`.
> De `gdd/sistemas/`, só `jogador.md` e `zonas.md`. As ADRs citadas abaixo (ADR-004, ADR-008)
> foram decisões tomadas fora do repo e ainda não têm arquivo — reescrever quando reaparecerem.

---

## Fatia 0 — Fundação

**Entrega:** servidor sobe, carrega o conteúdo, responde a uma verificação de saúde. Nada
visível ao jogador.
**Por que existe:** é a espinha mínima. Sem ela nenhuma fatia vertical fecha.
**Mantenha curta.** Se crescer, está virando arquitetura-primeiro.

### Backend

- [ ] Inicializar o projeto Go e a estrutura de pastas
- [ ] Carregar o conteúdo JSON no boot e expor tipado em memória
- [ ] Verificação de saúde: confirma que o servidor subiu e o conteúdo carregou

### Frontend

- [ ] Inicializar o projeto do cliente

### Dados

- [ ] `zones.json` — A Ressaca, A Bigorna, O Verde Surdo, A Costela, A Encruzilhada
- [ ] `resources/nodes.json` — nós de recurso por zona
- [ ] `mobs/ressaca.json`, `mobs/verde-surdo.json`, `mobs/costela.json`
- [ ] `weapons.json` e `tools.json` (T1–T2)
- [ ] `armor.json` (T1–T2)
- [ ] `crafting/recipes.json` (T1 e T2)
- [ ] `skills/abilities.json` — habilidades por arma e ataque automático
- [ ] `progression/litania.json` — Linha do Catador, Guerreiro, Caçador, Conjurador, Coleta, Forja
- [ ] `progression/fame_per_action.json` e `tutorial/steps.json`
- [ ] `tutorial/dialogs.json` — **PRONTO** (entregue 21/07/2026)

---

## Fatia 1 — Personagem (Passo 1: Chegada)

**Entrega:** o jogador cria um personagem e vê A Ressaca pela primeira vez.
**Spec:** `specs/001-personagem/`

### Backend

- [ ] Persistir o jogador: quem é, onde está, em que passo do tutorial
- [ ] Criar jogador e abrir sessão
- [ ] Devolver o estado do jogador — a rota de reidratação
- [ ] Devolver os dados de uma zona, incluindo as ações disponíveis

### Frontend

- [ ] Store do jogador: estado mínimo e reidratação no boot
- [ ] Tela de criação de personagem
- [ ] Ligar a criação à tela de zona
- [ ] Tela de zona: nome e lista de ações vinda do servidor
- [ ] Reidratação: recarregar mantém o jogador na zona correta

### Narrativa

- [ ] Exibir a fala de chegada do Guia, sem repetir ao recarregar

---

## Fatia 2 — Primeira arma (Passo 2)

**Entrega:** o jogador recebe uma arma, equipa, e vê o loadout mudar.

### Backend

- [ ] Persistir o equipamento do jogador
- [ ] Equipar um item, limpando as habilidades da peça trocada

### Frontend

- [ ] Tela de loadout: slots de equipamento e escolha de item por slot

### Narrativa

- [ ] Exibir os diálogos do Passo 2 — chegada, reação e espera

---

## Fatia 3 — Primeiro combate (Passo 3)

**Entrega:** o jogador mata o primeiro Mito sem Pacto, ganha Fama e sente a Têmpera pela primeira vez.
**Nota:** a Têmpera foi antecipada da Fatia 8 para cá. É a mecânica que define o projeto — construir sete fatias de idle genérico antes de tocá-la é risco de perder o gosto no meio.

### Backend

- [ ] Combate T1: ciclo fixo com loot e Fama, sem fila de habilidades ainda
- [ ] Iniciar uma ação — o servidor é a autoridade do tempo
- [ ] Consultar a ação em andamento e receber o resultado quando o ciclo fecha
- [ ] Cancelar uma ação em andamento
- [ ] Creditar Fama: Linha do Catador sempre, mais o nó específico da ação
- [ ] Têmpera simplificada: sobe a cada ciclo concluído, congela offline, reseta na troca de arma. Sem fila de habilidades, sem T2 — só a curva e o castigo

### Frontend

- [ ] Componente de ação em andamento, com progresso e tempo restante
- [ ] Ciclo de consulta automática: ao fechar um ciclo, reinicia sozinho
- [ ] Mostrar a Fama ganha a cada ciclo e em qual nó ela caiu

### Narrativa

- [ ] Exibir os diálogos do Passo 3

---

## Fatia 4 — Coleta (Passo 4)

**Entrega:** o jogador coleta recursos em loop e vê o inventário encher.

### Backend

- [ ] Configurar banco e migrations (adiado da Fatia 0 — ver docs/estudos/05-persistencia)
- [ ] Persistir o inventário do jogador
- [ ] Coleta: ciclo fixo que entrega recurso e Fama, e reinicia
- [ ] Consultar o inventário com quantidades e peso
- [ ] Validar slots e peso a cada ciclo

### Frontend

- [ ] Painel de inventário: slots, peso e itens

### Narrativa

- [ ] Exibir os diálogos do Passo 4

---

## Fatia 5 — A Bigorna: viagem e forja (Passo 5)

**Entrega:** o jogador viaja até A Bigorna, refina e crafta a primeira ferramenta.

### Backend

- [ ] Viajar entre zonas, com tempo de trânsito
- [ ] Refinar recursos nas estações de refino
- [ ] Craftar: validar a estação, consumir recursos, gerar o item e creditar Fama

### Frontend

- [ ] Integrar a viagem: tempo de trânsito e troca de zona ao chegar
- [ ] Tela de forja: receitas disponíveis na estação atual
- [ ] Retorno visível do craft: o que foi consumido e o que saiu

### Narrativa

- [ ] Exibir os diálogos do Passo 5, incluindo a apresentação da Têmpera

---

## Fatia 6 — Outro portador (Passo 6)

**Entrega:** o jogador cruza com o NPC-Portador; as armas se reconhecem.

### Backend

- [ ] Verificar e avançar os passos do tutorial
- [ ] Disparar o encontro com o NPC-Portador quando as condições baterem

### Frontend

- [ ] Indicador de progresso do tutorial: passo atual e o que falta

### Narrativa

- [ ] Exibir a fala do NPC-Portador e a reação do Guia

---

## Fatia 7 — Zonas T2 (O Verde Surdo, A Costela)

**Entrega:** o jogador viaja para as zonas T2 e coleta e luta em tier 2.

### Backend

- [ ] Creditar Fama de coleta T2 no nó certo (`PH_NODE_COLETA_MADEIRA`, `_MINERIO`, `_COURO`)

### Dados

- [ ] Recursos e mobs de O Verde Surdo e A Costela em `zones.json`

---

## Fatia 8 — Segunda arma e combate T2 (Passo 7)

**Entrega:** o jogador descobre a segunda arma, troca, e sente a Têmpera esfriar.

### Backend

- [ ] Combate T2: fila de habilidades com cooldowns individuais
- [ ] Salvar a ordem da fila de habilidades
- [ ] Identificar a peça que mais contribuiu com dano e creditar a Fama no nó dela
- [ ] Têmpera completa: multiplicador aplicado sobre o ganho de Fama, integrado à fila de habilidades e ao T2

### Frontend

- [ ] Tela de fila de habilidades: ver e reordenar
- [ ] Indicador de Têmpera, em metáfora de calor e sem jargão de sistema

### Narrativa

- [ ] Exibir os diálogos do Passo 7 — único momento em que o Guia diz "têmpera"

---

## Fatia 9 — A Litania e a Saída (Passo 8)

**Entrega:** o jogador vê A Litania inteira e atravessa A Encruzilhada.

### Backend

- [ ] Persistir o progresso do jogador na Litania
- [ ] Consultar A Litania: a árvore e o progresso do jogador
- [ ] Marcar o tutorial como concluído e liberar a saída da ilha

### Frontend

- [ ] Tela da Litania: nós, progresso e desbloqueios
- [ ] Tela de escolha do destino de saída (ver questão aberta 1)
- [ ] Tela de conclusão do tutorial

### Narrativa

- [ ] Exibir os diálogos do Passo 8 e o fechamento da chave temática

---

## Correções de nomenclatura aplicadas

| Antes (drift)                                                    | Depois (cânone)                                  |
| ---------------------------------------------------------------- | ------------------------------------------------ |
| O Farol                                                          | A Ressaca                                        |
| A Cova / A Cova Hub                                              | A Bigorna                                        |
| Floresta Esquecida                                               | O Verde Surdo                                    |
| Forte da Montanha                                                | A Costela                                        |
| nó Adventurer                                                    | Linha do Catador                                 |
| Woodcutter, Miner, Skinner                                       | `PH_NODE_COLETA_MADEIRA` / `_MINERIO` / `_COURO` |
| "o board"                                                        | A Litania                                        |
| `mobs/forest.json`, `mobs/mountain.json`                         | `mobs/verde-surdo.json`, `mobs/costela.json`     |
| `go mod init albion-idle`                                        | `go mod init escoria`                            |
| `internal/destiny/`, `/destiny-board`, `player_destiny_progress` | `litania` (ADR-004)                              |

## Questões abertas

1. **Destino de saída da ilha.** A tarefa original previa "5 cidades iniciais". No cânone a
   saída é **A Encruzilhada**, e a cidade neutra (A Forja Calada) é pós-PoC. As 5 opções podem
   mapear as 5 Carcaças (Pálida, Solar, Reverenciada, Trovejada, Tropical) — falta decidir.
2. **Slots de habilidade Q/W/E.** Convenção genérica de MMO ou trocar por outra coisa?
   Decisão natural na Etapa 7 da migração, junto com a página de Skills.
3. **Armadura no tutorial.** Existe na PoC e não carrega Têmpera. Confirmar se o jogador chega
   a equipar armadura durante os 8 passos, ou se ela só existe como dado.
