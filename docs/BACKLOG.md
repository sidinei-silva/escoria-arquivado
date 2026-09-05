# Backlog — PoC da Margem Calada

> Cada fatia entrega algo que o jogador **vê ou faz**, e mapeia um dos 8 passos
> do tutorial. Cada `- [ ]` vira uma issue; cada fatia, um milestone.
>
> Ordem de trabalho dentro da fatia: fluxo → sistema → conteúdo → código+teste →
> lore → commit. Ver [RECEITA.md](RECEITA.md).
>
> **Estado:** nenhuma linha de código escrita.

## PoC validada quando

1. O jogador completa os 8 passos sem travar
2. O loop fecha: zona → ação → recurso e Fama → Litania → novo conteúdo
3. Narrativa mínima presente
4. **Lastro e Fama percebidos como coisas diferentes**

## Fora da PoC

Multiplayer, PvP, gank, mercado, dungeons, guilda, zonas vivas, Journal,
diárias, temporadas, autenticação real, engine gráfica.

Ideia que aparecer no caminho vai para [ideias/](ideias/).

---

## Fatia 0 — Fundação

**Entrega:** o servidor sobe, carrega o conteúdo e falha claro se algo estiver
errado. Ninguém joga nada.

- [x] `go mod init escoria`, esqueleto `cmd/` + `internal/`
- [x] Carregador de `data/` — lê, tipa e expõe. Sem regra de negócio
- [x] `data/zonas.json` com A Ressaca
- [x] Validar referências no boot: objetivo aponta para zona que existe? diálogo
      aponta para objetivo que existe? recurso citado existe?
- [x] Falha no boot com mensagem legível quando uma referência quebra
- [x] `/health`

**Sem** banco, **sem** rede de jogo, **sem** cliente. Mantenha curta — se
crescer, virou arquitetura-primeiro.

---

## Fatia 1 - Criação da conta e login
**Entrega:** o jogador cria uma conta, faz login e vê a tela de criação de
personagem. Não consegue fazer nada ainda — só vê que existe.

---

## Fatia 2 — Chegada · Passo 1

**Entrega:** o jogador cria um personagem, nasce em A Ressaca, é recebido pelo
Guia e vê o primeiro objetivo. Não consegue fazer nada ainda — só vê que existe.

Fluxo: `gdd/fluxos/00-pre-jogo.md`, `gdd/fluxos/01-chegada.md`

### Servidor
- [ ] Criar personagem: nome, aparência, nasce em A Ressaca (`R-JOG-01..03`)
- [ ] Estado do jogador em memória, persistido em JSON no disco
- [ ] Sistema de objetivos, versão mínima: um ativo por vez, fecha por condição
- [ ] Condição `alcancar_local`
- [ ] Motor de diálogo: gatilho + `requer_objetivo` + `uma_vez`
- [ ] Reidratar: recarregar a página devolve o jogador onde estava

### Cliente
- [ ] Tela de criação de personagem
- [ ] Tela de zona: nome, descrição, ações disponíveis (todas desabilitadas)
- [ ] Caixa de diálogo do Guia

### Narrativa
- [ ] `dialog_passo1_abertura` e `_reacao` tocam no momento certo
- [ ] Lore do passo 1, escrita à mão, depois do código passar

### A decidir
- Conta e login entram agora? Se sim, `gdd/sistemas/conta.md` com `R-CTA-*`
  antes do código: quantos personagens por conta, nome único onde, apagar
  existe, tutorial é por conta ou por personagem
- Aparência: quais opções na PoC?

---

## Fatia 3 — A primeira arma · Passo 2

**Entrega:** o jogador recebe uma arma, equipa, e vê o loadout mudar.

Fluxo: `gdd/fluxos/02-primeira-arma.md`

### Servidor
- [ ] Sistema de equipamento: slots arma, peito, elmo, botas
- [ ] Equipar e desequipar
- [ ] Condição `equipar`
- [ ] `data/armas.json` — a arma inicial (Guerreiro T1)
- [ ] Evento de equipar carrega **o que desenhar**: id da peça, camada, tier —
      não só "ok". Sem isso, o cliente com sprites exige mudar o protocolo depois

### Cliente
- [ ] Painel de personagem com os quatro slots nomeados
- [ ] Loadout visível

### Narrativa
- [ ] `dialog_passo2_*` (abertura, reação, espera)
- [ ] Lore do passo 2

### A decidir
- A arma vem do Guia ou é achada?
- Armadura entra no tutorial?

---

## Fatia 4 — O corpo age sozinho · Passo 3

**Entrega:** o jogador mata o primeiro Mito sem Pacto, ganha Fama, e sente a
Têmpera pela primeira vez.

> **Fatia mais importante da PoC.** É aqui que o idle se explica sem ninguém
> explicar, e onde a mecânica que define o projeto aparece. A Têmpera foi
> antecipada para cá: construir sete fatias de idle genérico antes de tocá-la é
> risco de perder o gosto no meio.

Fluxo: `gdd/fluxos/03-primeiro-combate.md`
Estudos: ownership e concorrência · tempo e atividades

### Servidor
- [ ] **Game loop com dono único.** Uma goroutine dona do estado; comandos por
      channel; o loop só toca memória
- [ ] **Atividade com marcos absolutos** (`inicio`, `fim`), nunca contador
- [ ] Ciclo de combate: inicia, fecha, reinicia sozinho
- [ ] Ids estáveis por alvo nos eventos — o cliente precisa saber qual sprite
      morreu
- [ ] Fama: Linha do Catador sempre, mais o nó específico da ação
- [ ] **Têmpera simplificada:** sobe a cada ciclo concluído, congela offline,
      reseta na troca de arma. Sem fila de habilidades, sem T2
- [ ] Condição `abater`
- [ ] `data/mobs.json` — nomear os `PH_MOB_*` da ilha
- [ ] Polling: `GET /acao/atual` (ADR — sem WebSocket na PoC)

### Cliente
- [ ] Ação em andamento com destaque: progresso, tempo restante, alvo
- [ ] Log de eventos
- [ ] Têmpera visível **sem abrir nada**, junto da arma equipada

### Narrativa
- [ ] `dialog_passo3_*`
- [ ] Lore do passo 3

### A decidir
- O que acontece com a atividade em curso quando o jogador cai?
- Combate dá Fama já aqui; e coleta, dá?
- Tick: qual intervalo, e por qual requisito de jogo?

---

## Fatia 5 — A Escória dá · Passo 4

**Entrega:** o jogador coleta em loop no Verde Surdo e vê a mochila encher.

Fluxo: `gdd/fluxos/04-coleta.md`
Estudos: tempo e atividades · persistência

### Servidor
- [ ] **Banco entra aqui** (adiado da Fatia 0): SQLite, migrations, sqlc
- [ ] Migrar o estado do JSON para o banco
- [ ] Persistência fora do loop: snapshot periódico + escrita imediata no que
      dói perder
- [ ] Ciclo de coleta com nó de recurso
- [ ] Inventário com capacidade
- [ ] Parar a coleta manualmente
- [ ] Condições `possuir_recursos` e `acumular_recurso`
- [ ] `data/recursos.json`, nós em `zonas.json` — nomear os `PH_REC_*`

### Cliente
- [ ] Grade de inventário
- [ ] Ações da zona com duração visível

### Narrativa
- [ ] `dialog_passo4_*`
- [ ] Lore do passo 4

### A decidir
- E se a mochila encher no meio do ciclo?
- Parar cancela o ciclo em curso ou espera fechar?
- Quantas unidades o objetivo exige?

---

## Fatia 6 — A forja · Passo 5

**Entrega:** o jogador viaja até A Bigorna, refina e cria seu primeiro item. O
Guia apresenta a Têmpera **sem usar a palavra**.

Fluxo: `gdd/fluxos/05-forja.md`

### Servidor
- [ ] Viagem entre zonas, com trânsito
- [ ] Refino e craft
- [ ] Lastro creditado e gasto — **percebido como coisa diferente da Fama**
- [ ] Condição `craftar`
- [ ] `data/receitas.json`

### Cliente
- [ ] Tela de viagem
- [ ] Tela de forja
- [ ] Fama e Lastro sempre visíveis, separados

### Narrativa
- [ ] `dialog_passo5_*` — metáfora de forja, sem o termo técnico
- [ ] Lore do passo 5

### A decidir
- Craft contínuo existe?
- Viagem pode ser cancelada no meio?

---

## Fatia 7 — O outro portador · Passo 6

**Entrega:** encontro com o NPC-Portador n'A Costela. **As armas se reconhecem
antes das pessoas.**

Fluxo: `gdd/fluxos/06-portador.md`

### Servidor
- [ ] NPC posicionado em zona
- [ ] Condição `encontrar_npc`
- [ ] Diálogo com dois speakers alternando

### Cliente
- [ ] Diálogo com speaker identificado

### Narrativa
- [ ] `dialog_passo6_abertura`, `_npc`, `_reacao`
- [ ] Lore do passo 6

---

## Fatia 8 — Tier 2

**Entrega:** o jogador acessa zonas T2, coleta e luta em tier maior.

### Servidor
- [ ] Tier em zona, recurso, mob e item
- [ ] Requisito de tier para entrar e para equipar
- [ ] Conteúdo T2 em `data/`

### Cliente
- [ ] Tier visível na zona e no item

### A decidir
- O que destrava T2: Fama, nó da Litania, ou item equipado?

---

## Fatia 9 — A segunda arma · Passo 7

**Entrega:** o jogador troca de arma e **sente a Têmpera esfriar**. Único
momento do tutorial em que o Guia diz a palavra.

Fluxo: `gdd/fluxos/07-segunda-arma.md`

### Servidor
- [ ] Segunda linha de arma
- [ ] **Têmpera completa:** multiplicador sobre o ganho de Fama, por arma
- [ ] Têmpera por arma equipada, não por jogador
- [ ] Fila de habilidades com cooldowns
- [ ] Condição `trocar_equipado`

### Cliente
- [ ] Configurar a fila de habilidades
- [ ] Têmpera zerando de forma **visível** na troca

### Narrativa
- [ ] `dialog_passo7_*` — a palavra "têmpera" aparece na reação
- [ ] Lore do passo 7

### A decidir
- A Têmpera antiga volta se ele reequipar a primeira arma, ou perdeu?

---

## Fatia 10 — A Litania e a saída · Passo 8

**Entrega:** o jogador vê a Litania inteira e atravessa A Encruzilhada.
**Fecha a PoC.**

Fluxo: `gdd/fluxos/08-saida.md`

### Servidor
- [ ] Litania: Linha do Catador, linha por arma, progresso por nó
- [ ] Destravar por Fama acumulada no nó
- [ ] Fim do tutorial: cadeia com `proximo: null`

### Cliente
- [ ] Tela da Litania com os nós visíveis, destravados e não
- [ ] Tela de saída

### Narrativa
- [ ] `dialog_passo8_*` — fecha "você só assina depois", em tempo passado
- [ ] Lore do passo 8

### A decidir
- Quais opções de destino ao sair da ilha?
- A saída é irreversível?

---

## Depois da PoC

Vai para [ideias/](ideias/) ou vira backlog novo quando a hora chegar:
multiplayer e partição por zona, gank e PvP, mercado, Journal, diárias, cliente
em engine com sprites em camadas.