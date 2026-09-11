# 001 — Tarefas

> Ordem sugerida: de baixo para cima (banco → API → cliente), fechando com o diálogo.
> Cada item vira uma issue.

## Pré-requisito

- [ ] Fatia 0 concluída (servidor sobe, carrega JSONs, responde `/health`)
- [ ] `data/zones.json` contém ao menos `a_ressaca` (forma acordada em `dados.md`)

## Backend

- [ ] **Persistência do jogador**
      Migration + queries para guardar quem é, onde está e em que passo do tutorial.
      Decisões em aberto na seção "O que precisa ser persistido" de `dados.md`.

- [ ] **Criar jogador e abrir sessão**
      Valida o nome, cria em A Ressaca no passo 1, devolve o suficiente para o cliente montar.
      Aceite 1, 2, 3.

- [ ] **Devolver estado do jogador**
      A rota de reidratação. Aceite 6.

- [ ] **Devolver dados de uma zona**
      Lê do conteúdo em memória, não do banco. Zona inexistente → "não encontrado".
      Aceite 4, 7.

## Frontend

- [ ] **Store do jogador (Zustand)**
      Estado mínimo + reidratação no boot. Ver "O que o cliente precisa guardar" em `dados.md`.

- [ ] **Tela: Criar Personagem**
      Campo nome + confirmar. Erro visível quando o nome é inválido. Aceite 2.

- [ ] **Integração criação → zona**
      `POST /auth/mock` → popular store → navegar para a tela de zona. Aceite 1.

- [ ] **Tela de Zona**
      Nome da zona + lista de ações vinda do servidor. Ações inertes nesta fatia.
      Deixar espaço no layout para: ação em andamento, inventário, loadout, progresso do
      tutorial. Aceite 4.

- [ ] **Reidratação**
      Recarregar a página mantém o jogador na tela de zona correta. Aceite 1.

## Narrativa

- [ ] **Exibir fala de abertura do Guia**
      Ler `dialog_passo1_abertura` de `data/tutorial/dialogs.json`. Exibir na primeira chegada.
      Não repetir ao recarregar — usar `tutorial_passo` para decidir. Aceite 5.
      Registro diegético, não onboarding (ver notas de direção na `spec.md`).

## Fechamento da fatia

- [ ] Percorrer os 7 critérios de aceite à mão
- [ ] Atualizar `spec.md` se algo divergiu na implementação
- [ ] Registrar decisões relevantes em `docs/decisoes.md`
- [ ] Registrar a sessão em `docs/diario.md`
