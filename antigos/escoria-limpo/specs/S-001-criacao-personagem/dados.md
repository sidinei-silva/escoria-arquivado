# 001 — Dados e Fronteiras

> **O que precisa existir e por quê.** Como estruturar é decisão do engenheiro.
>
> Este documento faz perguntas em vez de dar respostas. As perguntas marcadas com **▸** são
> as que a implementação precisa responder. Se você responder diferente do que eu imaginaria,
> ótimo — desde que os critérios de aceite da `spec.md` continuem verdadeiros.

---

## O que precisa ser persistido

O jogo precisa lembrar de **três coisas** sobre um jogador, e só três nesta fatia:

**Quem ele é.** Um nome que ele escolheu, entre 3 e 20 caracteres. E alguma forma de o
servidor reconhecê-lo de novo quando ele voltar.

**Onde ele está.** Todo jogador está sempre em exatamente uma zona. Nunca em nenhuma, nunca em
duas. Ao ser criado, ele está em A Ressaca — isso não é configurável.

**Em que ponto do tutorial ele está.** Um marcador do passo atual (1 a 8). Nesta fatia ele
sempre nasce em 1 e nunca avança. Existe agora por dois motivos: sustenta o critério de aceite
5 (não repetir a fala do Guia ao recarregar) e evita uma migration extra na Fatia 6.

### Perguntas para você decidir

**▸ Como o servidor reconhece um jogador que volta?** A PoC não tem autenticação real. Você
precisa de algo que sobreviva a um refresh da página. Cookie de sessão? Token no
localStorage? Id na URL? Cada opção tem consequência diferente quando a autenticação real
chegar — mas nenhuma delas deve fazer as outras rotas mudarem de forma.

**▸ A zona do jogador é uma referência a quê?** Zonas são conteúdo, vivem em JSON, não em
tabela (ver `000-constituicao.md`, princípio 1). Então o banco guarda uma string que aponta
para um id do JSON. Isso significa que o banco **não consegue** garantir que a zona existe.
Onde essa validação acontece? O que acontece se alguém editar o JSON e remover uma zona onde
há jogadores parados?

**▸ Onde a regra "3 a 20 caracteres" vive?** No banco, na camada de API, no cliente, ou em
mais de um? Cada resposta tem um custo distinto. Não existe resposta única correta — mas
existe resposta inconsistente.

---

## O que a API precisa oferecer

Três capacidades. Os nomes de rota abaixo são sugestão; o que importa é a capacidade existir.

**Criar um jogador e abrir sessão.** Recebe um nome. Se o nome for inválido, recusa de forma
que o cliente consiga mostrar mensagem útil — não um 500 genérico. Se for válido, cria o
jogador em A Ressaca no passo 1 e devolve o suficiente para o cliente se montar.

**Devolver o estado atual do jogador.** Esta é a rota de reidratação: o cliente chama ao
abrir e monta a tela a partir dela. É a rota mais importante da fatia, porque toda fatia
seguinte vai pendurar campo nela.

**Devolver os dados de uma zona, incluindo as ações disponíveis.** Lê do conteúdo carregado
em memória, não do banco. Zona inexistente responde "não encontrado", não erro de servidor.

### Perguntas para você decidir

**▸ A rota de estado devolve campos que ainda não existem?** Nas Fatias 3 e 4 o jogador vai
ter ação em andamento e inventário. Você pode: (a) já devolver esses campos vazios agora, e o
cliente nasce preparado; (b) adicionar depois, e mexer no cliente de novo. A opção (a) custa
quase nada agora e evita retrabalho — mas é uma aposta de que você sabe a forma final, e você
não sabe. **Decida ciente do trade-off.**

**▸ A tela de zona pede a zona separadamente, ou o estado do jogador já traz tudo?** Duas
chamadas é mais simples de cachear e mais fácil de raciocinar. Uma chamada é menos ida-e-volta
e menos estado dessincronizado. A escolha aqui define o ritmo de todas as telas seguintes.

**▸ Como você identifica erros?** Precisa de algo que o cliente consiga distinguir
programaticamente — "nome inválido" é diferente de "sem sessão" e o cliente reage diferente.
Código de erro no corpo? Só o status HTTP? A convenção que você escolher aqui vale para o
projeto inteiro; anote em `000-constituicao.md`.

---

## O que o cliente precisa guardar

Um único store com o estado do jogador. Nesta fatia ele guarda pouco: quem é, onde está, em
que passo do tutorial. Mas ele é o esqueleto que todas as fatias seguintes engordam.

O cliente também precisa saber **se já mostrou a fala de chegada** — sem isso, recarregar a
página repete o diálogo e quebra o critério 5.

### Perguntas para você decidir

**▸ "Já mostrei a fala" mora no cliente ou no servidor?** Se morar no cliente, limpar o
navegador faz o Guia te receber de novo — o que talvez seja aceitável, talvez não. Se morar no
servidor, é estado de verdade e precisa de coluna. O `tutorial_passo` já existe; ele resolve?

**▸ O store espelha a forma da API ou tem forma própria?** Espelhar reduz tradução e bugs de
mapeamento. Ter forma própria protege o cliente de mudanças na API. Para um projeto solo com
uma API só, espelhar costuma ganhar — mas é decisão sua.

---

## Fronteira com o conteúdo — aqui a forma importa

**Esta é a única seção deste documento que é exata**, e a razão é específica: os arquivos de
`data/` são produzidos nos chats de conteúdo e consumidos pelo seu código. Se a forma não
estiver combinada, você recebe JSON que seu loader não lê.

Você define a forma. Eu sigo. O que está abaixo é **proposta minha, sujeita à sua revisão** —
mude à vontade e eu me adapto.

### `data/zones.json`

Precisa conter, para cada zona: um **id estável** (string, sem acento, usado pelo banco e pelas
rotas), um **nome de exibição**, o **tier**, e a **lista de ações disponíveis**.

Cada ação precisa de: um **id estável**, um **tipo** (que o servidor usa para decidir qual
lógica roda — coleta, combate, craft, refino, viagem) e um **rótulo diegético** (o texto que o
jogador lê, na voz do jogo — "Juntar lascas", não "Coletar madeira").

A separação entre `tipo` e `rotulo` é deliberada: o tipo é mecânico e estável, o rótulo é
criativo e vai mudar várias vezes sem quebrar nada.

**▸ Decisão sua:** as ações vivem dentro da zona (como acima) ou num arquivo separado que
referencia zonas? Dentro é mais simples de ler. Fora evita duplicação quando a mesma ação
aparecer em várias zonas — o que vai acontecer a partir da Fatia 7.

### `data/tutorial/dialogs.json`

**Já existe e está fechado.** Forma documentada no bloco `meta` do próprio arquivo.

Nesta fatia você consome um bloco: `dialog_passo1_abertura`, com gatilho
`player_desembarca_primeira_vez`.

Dois avisos sobre esse arquivo:

- O campo `subtexto` é **nota de direção interna**. Nunca exibido ao jogador. É contexto para
  quem for dirigir voz ou tomar decisão de UI.
- O bloco `dialog_passo1_reacao` tem gatilho `player_conclui_primeira_navegacao`. Não há
  navegação nesta fatia — ele fica para a Fatia 5.

---

## O que não decidir agora

Estas perguntas vão aparecer e **não** precisam de resposta nesta fatia. Anotadas para você
não gastar energia nelas:

- Como paginar inventário (Fatia 4)
- Como o servidor valida que uma ação é permitida na zona (Fatia 3)
- Onde a Têmpera é calculada (Fatia 8)
- Como autenticação real substitui a sessão mock (pós-PoC)
