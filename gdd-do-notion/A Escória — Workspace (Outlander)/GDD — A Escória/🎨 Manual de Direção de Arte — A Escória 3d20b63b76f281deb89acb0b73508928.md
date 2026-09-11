# 🎨 Manual de Direção de Arte — A Escória

# Objetivo

Este documento define a linguagem visual de **A Escória** para manter personagens, equipamentos, monstros, cenários, efeitos e interface coerentes durante a produção. O jogo é um **MMORPG idle (gear-based)**; a identidade vem do equipamento, a zona determina a ação e o jogador observa o personagem executar a atividade. fileciteturn11file0

## 1. Estilo gráfico

**Descrição curta:** Stylized 2D cartoon game art, chibi proportions, bold black outlines, flat colors, minimal shading.

Em português: **arte 2D cartoon estilizada, proporções chibi, contornos pretos fortes, cores chapadas e sombreamento mínimo**.

### Princípios

- Visual desenhado e estilizado; não realista.
- Formas grandes, simples e legíveis.
- Contorno escuro consistente.
- Paleta limitada por cena.
- Sombra em poucos volumes; preferir formas chapadas a pintura.
- Silhueta vem antes do detalhe.
- Personagens e equipamentos são mais limpos que o cenário.
- O cenário é simples o suficiente para ser reproduzido e reutilizado.
- A animação é curta e expressiva: o jogo é idle, mas o mundo não parece parado.

## 2. Inspirações e o que absorvemos

**Albion Online:** clareza de RPG, personagem/equipamento como identidade e progressão por uso. O GDD usa Albion como referência mecânica, mas o vocabulário interno é **A Litania**, não Destiny Board. fileciteturn27file0

**Femur Idle / Ygg Idle / outros idles animados:** sensação de mundo vivo mesmo quando a interação é mínima; o personagem executa ações e o jogador observa.

**Epic Idle Journey:** leitura imediata, personagens e objetos simplificados e apresentação casual de RPG.

As referências devem orientar a linguagem visual, não ser copiadas literalmente.

## 3. Personagens

### Proporção

- Cabeça grande, sem exagero; referência aproximada de 2,5–3 cabeças de altura.
- Corpo curto.
- Braços e pernas simples.
- Mãos e pés simplificados.
- Posturas fáceis de ler.
- A proporção pode variar levemente quando necessário para preservar silhueta, leitura e encaixe dos equipamentos.

### Rosto

- Poucos elementos.
- Olhos/visor como detalhe principal.
- Evitar anatomia e expressões faciais detalhadas.

### Silhueta

O personagem deve ser reconhecível em tamanho pequeno. A ordem de prioridade é:

**silhueta → cor → equipamento → detalhe**.

### Catador

O **Catador** começa visualmente simples e vulnerável. No tutorial inicial, ele aparece sem equipamento e vai adquirindo identidade conforme recebe e veste a primeira arma. O protagonista é um catador de sucata que é puxado para A Escória após despertar uma arma. fileciteturn12file4

## 4. Equipamento

Equipamento é parte da identidade visual do personagem. O loadout deve ser visível sempre que possível, porque o GDD trata equipamento como identidade. fileciteturn17file0

### Regras

- Cada peça precisa ter silhueta clara.
- Preferir uma forma básica que possa ser redesenhada muitas vezes.
- Evoluções de T1/T2/T3 devem alterar silhueta e alguns detalhes, não exigir um desenho totalmente novo.
- Armas recebem mais atenção visual que pequenos acessórios.
- Slots visuais principais: cabeça, torso, pernas/pés, mão principal e mão secundária; acessórios/capas podem existir como camadas adicionais.
- Cada peça deve usar anchors/pontos de encaixe consistentes para permitir troca de equipamento sem reposicionar manualmente o personagem.
- A ordem de layers deve ser previsível e documentada por slot, preservando mãos, armas, corpo e efeitos na frente/trás corretos.
- Escala, pivô e orientação das peças devem permanecer padronizados entre personagens e tiers.

### Têmpera

A **Têmpera** é o diferencial mecânico da arma equipada e deve ter uma representação visual simples, ligada à sensação de calor/esfriamento. O estado visual deve estar associado à **arma equipada e à sua progressão**, e não funcionar como um efeito decorativo ou elemento de UI genérico. fileciteturn27file0

## 5. Monstros e NPCs

### Monstros

- Silhueta diferente do jogador.
- Formas simples.
- Poucas cores.
- Um elemento visual marcante por tipo.
- Variações por tier devem partir da mesma família de desenho.

### Guia

O Guia deve ter uma silhueta imediatamente reconhecível: figura simples, roupa/capuz e cajado. Ele é um personagem recorrente do tutorial e recebe os recém-chegados. fileciteturn12file4

## 6. Cenários

A regra é **parecer cheio sem precisar desenhar muito**.

Um cenário deve ser construído com poucos assets reutilizáveis:

- 2–3 pedras.
- 1–2 madeiras/estruturas.
- 1 caixa/barril.
- 1 elemento exclusivo da zona.
- chão simples.

### Identidade por zona

**A Ressaca:** terra/areia, água, madeira, pedra e resíduos da maré.

**A Bigorna:** fogo, metal, madeira e estrutura de forja.

**O Verde Surdo:** vegetação, madeira e materiais naturais.

**A Costela:** ossos/costelas, cobre e algodão.

O GDD define A Ressaca como zona T1 de **coleta e combate**; A Bigorna como hub de **forja, refino, equipar e viajar**; O Verde Surdo e A Costela como zonas T2 com combinações próprias de coleta, combate e viagem. **Regra de interface:** a zona deve comunicar claramente quais ações estão disponíveis naquele local; a ação selecionada fica visualmente destacada e sua execução/progresso deve permanecer legível. fileciteturn25file0

## 7. Perspectiva

**Isométrica simplificada / 3/4 estilizada.**

Não é necessário perseguir isometria matemática perfeita. A prioridade é:

- sensação de profundidade;
- leitura de personagens e equipamentos;
- espaço suficiente para animação;
- produção simples.

## 8. Cores

### Mundo

Base preferencial:

- marrom-terra;
- ferrugem;
- cinza-chumbo;
- carvão;
- bege/sal;
- verde-musgo;
- azul dessaturado para água.

### Acentos

Usar acentos de maneira controlada, principalmente para:

- seleção;
- progresso;
- alerta;
- fogo/calor;
- raridade;
- recompensa.

### Regra prática

Objeto comum: **2–3 cores**.

Personagem: **3–6 cores principais**.

Evitar arco-íris e gradientes complexos em personagens.

### Leitura de recursos e progressão

Recursos, Lastro e Fama devem possuir tratamentos visuais distintos e consistentes.

- **Recursos:** associados visualmente ao material obtido (madeira, pedra, minério etc.).
- **Lastro:** associado a ganho/uso econômico e progressão prática do personagem.
- **Fama:** associado à progressão e reconhecimento, com destaque visual próprio.

A linguagem visual deve permitir identificar rapidamente qual tipo de recompensa foi obtida, sem depender apenas do texto.

## 9. Sombreamento e textura

### Sombreamento

Preferir **base + sombra + pequeno brilho**.

Não usar pintura realista, múltiplos degradês ou textura fotográfica.

### Textura

- Cenário pode ter desgaste, ferrugem, rachaduras e sujeira.
- Personagens devem permanecer relativamente limpos e legíveis.
- Textura deve ser sugerida por poucas formas.

## 10. Animação

A Escória é idle, mas as ações são observáveis. O GDD propõe um quadro central em que o personagem age sozinho; o quadro ilustra a ação enquanto as informações importantes continuam nos painéis. fileciteturn17file0

### Animações mínimas

**Idle:** respiração, balanço pequeno, mudança de postura.

**Combate:** preparar → atacar → impacto → recuperar.

**Coleta:** aproximar → trabalhar → pausa → recompensa.

**Forja:** levantar ferramenta → bater → faísca → pausa.

**Viagem:** caminhar → transição.

A meta é produzir poucas poses reutilizáveis, não animações longas.

## 11. Efeitos

Efeitos devem complementar a leitura sem encobrir o personagem.

- Impacto simples.
- Faíscas pequenas.
- Brilho curto.
- Dano flutuante.
- Indicador de progresso.
- Pequenos sinais de calor para Têmpera.

## 12. Interface

A intenção de layout é uma tela ancorada: quadro de observação no centro, painéis ao redor e informação importante sempre fora do quadro. fileciteturn17file0

### Componentes

- painel do personagem;
- equipamento;
- inventário;
- A Litania;
- ação em andamento;
- ações disponíveis na zona;
- log de eventos;
- mapa/minimapa;
- navegação;
- configurações.

### Linguagem

- retângulos arredondados;
- contorno escuro;
- fundo escuro;
- ícones simples;
- títulos curtos;
- números destacados.

Evitar UI dourada e ornamental demais.

## 13. Texto e tipografia

A voz visual e textual deve combinar: o GDD define registro **oral, gasto e pragmático**, com campo semântico de metalurgia. fileciteturn13file0

### Tipografia

- sem serifa;
- grossa e legível;
- formas arredondadas;
- tamanhos grandes para títulos e valores;
- contraste forte.

### Texto diegético

Preferir frases curtas e concretas.

Evitar jargão técnico de interface dentro do mundo.

## 14. Tutorial — princípio visual

O tutorial da PoC tem **8 passos** na Margem Calada e é a espinha dorsal desta primeira experiência. fileciteturn16file0

A curva visual deve ser:

**descartável → equipado → assustado → curioso → comprometido → reconhecido → em dúvida → responsável**. fileciteturn16file0

Cada passo deve introduzir poucos elementos novos e apontar claramente para a próxima ação.

---

# 15. Telas e elementos do tutorial

## Tela 01 — Chegada: A Ressaca

**Narrativa:** o Catador chega à ilha tutorial; o Guia o recebe.

**Elementos:**

- cena da A Ressaca;
- Catador sem equipamento;
- Guia;
- marcador visual do Guia;
- nome da zona;
- painel básico do personagem;
- diálogo;
- objetivo tutorial: falar com o Guia;
- indicação visual do clique/interação.

A PoC começa pela chegada em **A Ressaca** e o primeiro beat narrativo é “Mais um.” fileciteturn16file0

## Tela 02 — Conversa com o Guia

**Elementos:**

- retrato do Guia;
- nome do Guia;
- caixa de diálogo;
- botão/ação continuar;
- cenário parcialmente visível;
- objetivo atual.

Fala-âncora do Passo 1: **“Mais um. Anda. A maré não traz quem ela não pode usar.”** fileciteturn20file0

## Tela 03 — Primeira arma

**Narrativa:** o Guia entrega a primeira arma T1.

**Elementos:**

- arma T1;
- janela de entrega/recompensa;
- inventário;
- destaque do slot de arma;
- diálogo;
- indicador de nova ação.

Fala-âncora: **“Pega. Pra você não morrer na primeira noite.”** fileciteturn20file0

## Tela 04 — Equipar a primeira arma

**Elementos:**

- personagem;
- arma na mão;
- slot de equipamento;
- destaque do botão/slot;
- painel de atributos;
- visual antes/depois;
- orientação para equipar.

## Tela 05 — Primeiro combate

**Narrativa:** o jogador inicia o combate e percebe que o corpo age sozinho.

**Elementos:**

- Catador equipado;
- um inimigo T1 simples;
- barras de vida;
- ação “Matar”;
- fila de habilidades;
- progresso/tempo da ação;
- dano flutuante;
- animação de ataque;
- reação do inimigo;
- Guia observando;
- mensagem tutorial.

O combate da PoC é automatizado com fila; o jogador monta a ação e assiste. fileciteturn15file1

Fala-âncora: **“Vai. Deixa o braço fazer o que ele lembra — você só assina depois.”** fileciteturn21file0

## Tela 06 — Resultado do combate

**Elementos:**

- inimigo derrotado;
- recompensa;
- Fama;
- Lastro/drops quando aplicável;
- log de evento;
- atualização da progressão;
- fala do Guia.

Após o primeiro combate, o Guia comenta: **“Engraçado, né. A lasca lembra antes de você.”** fileciteturn21file0

## Tela 07 — Primeira coleta

**Narrativa:** A Escória fornece materiais estranhos ao mundo exterior.

**Elementos:**

- A Ressaca;
- escolha entre **Coletar** e **Matar**;
- ação Coletar selecionada;
- recurso da zona;
- Catador executando a coleta;
- progresso;
- tempo;
- recompensa de recurso;
- Fama da coleta;
- log.

A Ressaca é definida como zona T1 de **Tronco/Pedra Bruta**, com ações **Coletar** e **Matar**. fileciteturn25file0

## Tela 08 — Coleta em andamento

Esta é a manifestação mais clara do “idle com animação”.

**Elementos:**

- personagem em loop de coleta;
- barra/progresso;
- contador/tempo;
- recurso ganho;
- Fama ganha;
- ação atual destacada;
- alternativa de trocar para outra ação da zona;
- log de eventos.

O core loop é **zona → ação → recursos/Lastro/Fama**. fileciteturn8file0

## Tela 09 — A Bigorna

**Narrativa:** o tutorial leva o jogador à estação de forja.

**Elementos:**

- zona/hub A Bigorna;
- estações;
- inventário;
- materiais;
- receitas;
- custo em Lastro;
- resultado;
- ações Craftar, Refinar, Equipar e Viajar.

## Tela 10 — Primeira forja

**Elementos:**

- item de entrada;
- materiais necessários;
- custo em Lastro;
- equipamento produzido;
- novo tier;
- botão Forjar;
- resultado;
- animação/efeito de forja.

O loop da sessão transforma recursos e drops em equipamento melhor e tier maior. fileciteturn8file0

## Tela 11 — Têmpera

**Elementos:**

- arma equipada;
- estado quente/esfriando;
- ganho multiplicado pela Têmpera;
- indicação visual simples;
- informação sobre troca de arma.

A Têmpera sobe com uso ativo da mesma arma, congela offline e reseta na troca; a Fama acumulada não é perdida. fileciteturn27file0

## Tela 12 — Outro Portador

**Elementos:**

- segundo Portador;
- personagem do jogador;
- armas/equipamentos visíveis;
- diálogo;
- reação ao loadout;
- marcador do NPC.

Fala-âncora: **“Sua arma tem gosto. Cuidado quando ela começar a ter voz.”** fileciteturn20file0

## Tela 13 — Segunda arma escondida

**Elementos:**

- segunda arma;
- comparação com a arma atual;
- Fama de cada linha;
- Têmpera atual;
- opção de equipar;
- indicação da troca.

A primeira arma esfria, mas a Fama dela permanece; a escolha começa a definir a identidade do jogador. fileciteturn16file0

## Tela 14 — Troca de arma

**Elementos:**

- arma anterior;
- nova arma;
- Têmpera antes/depois;
- confirmação de equipar;
- mudança visual do personagem;
- atualização dos atributos e ações.

## Tela 15 — Saída / A Encruzilhada

**Elementos:**

- caminho de saída;
- destino A Encruzilhada;
- resumo do tutorial;
- ação Viajar;
- última fala do Guia;
- transição de zona.

O tutorial fecha em **A Encruzilhada**. fileciteturn16file0

## 16. Estados recorrentes da tela principal

Em vez de criar uma tela separada para cada ação, a maior parte do jogo deve acontecer em uma mesma tela de zona.

### Estrutura

- quadro central de observação;
- painel do personagem;
- painel da ação atual;
- painel das ações da zona;
- inventário/hotbar;
- log;
- mapa.

### Estados principais

**Sem ação:** personagem em idle.

**Coletando:** personagem executa animação de coleta.

**Matando:** personagem executa combate automático.

**Forjando:** personagem/estação mostra animação de forja.

**Viajar:** animação/transição.

A intenção de layout é que o quadro seja opcional e ilustrativo; as informações competitivamente relevantes devem existir nos painéis. fileciteturn17file0

---

# 17. Checklist de produção de arte

Antes de considerar um asset pronto, verificar:

- A silhueta é reconhecível?
- O desenho ainda funciona pequeno?
- Há poucos tons?
- O contorno está consistente?
- A peça combina com o restante do jogo?
- É fácil redesenhar ou variar este asset?
- A animação pode ser feita com poucas poses?
- O objeto está comunicando algo do mundo ou da mecânica?

# 18. Regra de ouro

> **A Escória deve ser simples de desenhar, fácil de animar e difícil de confundir com outro jogo.**
> 

A complexidade deve vir da **combinação de personagens, equipamentos, animações, zonas e interface**, não de cada desenho individual.

## Referências do GDD

- **GDD — A Escória:** visão geral, nomenclatura e estrutura do projeto. fileciteturn11file0
- **Narrativa Jogada — PoC (8 passos):** sequência e intenção do tutorial. fileciteturn16file0
- **Intenção de Layout:** composição da tela e papel do quadro central. fileciteturn17file0
- **Tom & Estilo:** voz, léxico e comunicação por tier. fileciteturn13file0
- **Geografia & Zonas:** zonas e ações disponíveis. fileciteturn25file0
- **Core Loop:** relações entre zona, ação, recompensa e progressão. fileciteturn8file0
- **A Litania & Fama:** Fama e Têmpera. fileciteturn27file0

# 25. Catálogo de telas, estados e elementos — referência visual

> **Decisão de câmera:** A Escória usa **isométrico simplificado / 3/4 estilizado** como perspectiva visual padrão. A prioridade é sensação de profundidade, leitura do personagem/equipamento e produção simples; não é necessário perseguir isometria matemática perfeita.
> 

Esta seção transforma o GDD e o manual de arte em uma especificação prática para mockups. **“Tela” não significa necessariamente uma cena separada:** em um MMORPG idle, várias funções devem continuar sendo estados, painéis ou janelas da tela principal da zona.

## 25.1 Fluxo de entrada

### Tela 01 — Criar conta

- Logo A Escória
- E-mail
- Senha
- Confirmar senha
- Criar conta
- Já tenho uma conta → Entrar
- Mensagens de validação
- Fundo ilustrado simples do mundo

### Tela 02 — Login

- Logo
- E-mail
- Senha
- Entrar no mundo
- Esqueci minha senha
- Criar conta
- Mensagens de erro/validação

### Tela 03 — Seleção de personagem

- Personagens existentes
- Personagem selecionado em destaque
- Nome
- Nível
- Equipamento principal
- Última região/zona
- Criar personagem
- Excluir personagem
- Entrar no mundo
- Estado vazio quando não houver personagem

### Tela 04 — Criação de personagem

- Nome do personagem
- Aparência simples
- Prévia grande do personagem
- Criar personagem
- Voltar
- A criação não deve virar uma escolha tradicional de classe: a identidade do jogo vem principalmente do equipamento. As linhas de arma da PoC são Guerreiro, Caçador e Conjurador.

## 25.2 Tela principal do jogo — conceito-base

Esta é a **principal referência de layout** para o jogo completo. O cenário isométrico fica no centro e os painéis permanecem ancorados ao redor dele.

```
┌──────────────────────────────────────────────┐
│ FAMA     LASTRO      ZONA           MENU     │
├──────────┬───────────────────────┬───────────┤
│          │                       │           │
│ PERSONA- │                       │  AÇÃO     │
│ GEM      │     CENA DA ZONA      │           │
│          │       ISOMÉTRICA      │ PROGRESSO │
│ EQUIPA-  │   personagem animado  │           │
│ MENTO    │   executando ação     │  EVENTOS  │
│          │                       │           │
├──────────┴───────────────────────┴───────────┤
│ INVENTÁRIO              │ AÇÕES DA ZONA      │
└─────────────────────────┴────────────────────┘
```

### Elementos persistentes/prioritários

- Cenário central isométrico
- Personagem do jogador
- NPCs e monstros quando aplicável
- Objetos/recursos da zona
- Animação idle e animação da ação
- Painel do personagem
- Equipamento visível
- Ação atual e progresso
- Ações disponíveis na zona
- Inventário/resumo
- Fama e Lastro
- Log de eventos
- Navegação/mapa quando aplicável

### Regra de leitura

**Personagem → ação → zona → consequências → informação secundária.**

O quadro central é observacional: ele mostra o personagem vivendo a ação, mas toda informação importante deve continuar existindo nos painéis.

## 25.3 Estados da tela principal

### Estado Idle

O personagem respira, balança levemente, muda de postura e permanece no mundo.

### Estado de Combate

- Matar
- Alvo atual
- Progresso/tempo
- Fila de habilidades quando aplicável
- Vida do jogador/inimigo
- Dano flutuante
- Animação de ataque/reação
- Pequenos efeitos

### Estado de Coleta

- Coletar
- Recurso da zona
- Personagem em loop de coleta
- Progresso
- Tempo
- Recurso ganho
- Fama ganha
- Log

### Estado de Forja

- Forjar/Craftar
- Refinar
- Material
- Lastro
- Resultado
- Animação de forja

### Estado de Viagem

- Destino
- Caminho/conexão
- Ação de viajar
- Transição visual

### Regra

**Não criar telas independentes de combate, coleta ou viagem sem necessidade.** Essas funções são estados da tela principal.

## 25.4 Ações por zona — regra visual e funcional

A zona define quais ações aparecem.

| Zona | Tier | Ações | Recursos principais |
| --- | --- | --- | --- |
| A Ressaca | T1 | Coletar, Matar | Tronco, Pedra Bruta |
| A Bigorna | Hub | Craftar, Refinar, Equipar, Viajar | materiais/Lastro |
| O Verde Surdo | T2 | Coletar, Matar, Viajar | Bétula, Couro |
| A Costela | T2 | Coletar, Matar, Viajar | Cobre, Algodão |

## 25.5 Tutorial — catálogo de mockups

A PoC tem 8 passos. O fluxo narrativo está fechado; o catálogo abaixo detalha os estados visuais que precisam ser produzidos.

### Mockup 01 — Chegada: A Ressaca

- Cena isométrica
- Catador sem equipamento
- Guia
- Marcador no Guia
- Nome da zona
- HUD básico
- Objetivo tutorial
- Indicação de interação

Fala-âncora: **“Mais um. Anda. A maré não traz quem ela não pode usar.”**

### Mockup 02 — Conversa com o Guia

- Retrato/nome do Guia
- Caixa de diálogo
- Botão continuar
- Cenário parcialmente visível
- Objetivo atual

### Mockup 03 — Recebimento da primeira arma

- Primeira arma T1
- Janela de entrega/recompensa
- Inventário
- Destaque do slot de arma
- Tooltip curto

Fala-âncora: **“Pega. Pra você não morrer na primeira noite.”**

### Mockup 04 — Equipar a primeira arma

- Personagem
- Arma na mão
- Slot de arma
- Destaque do slot/botão
- Atributos necessários
- Visual antes/depois

### Mockup 05 — Primeiro combate

- Catador equipado
- Um Mito sem Pacto T1 simples
- Barras de vida
- Ação Matar
- Alvo atual
- Progresso/tempo
- Fila de habilidades, quando aplicável
- Dano flutuante
- Animação de ataque
- Guia observando

Mensagem visual do tutorial: **“Matar — ação automática.”**

Fala-âncora: **“Vai. Deixa o braço fazer o que ele lembra — você só assina depois.”**

### Mockup 06 — Resultado do combate

- Inimigo derrotado
- Recompensa
- Fama
- Lastro/drops quando aplicável
- Log de evento
- Atualização de progressão

Fala-âncora: **“Engraçado, né. A lasca lembra antes de você.”**

### Mockup 07 — Primeira coleta

- Ações COLETAR / MATAR
- COLETAR selecionado
- Recursos da zona
- Personagem executando coleta
- Progresso
- Tempo
- Recurso ganho
- Fama ganha
- Log

### Mockup 08 — Coleta em andamento

- Loop visual da coleta
- Ação atual destacada
- Barra/progresso
- Contador/tempo
- Recurso ganho
- Fama ganha
- Ações alternativas disponíveis

Fala-âncora: **“Isso aqui dá o que não existe em lugar nenhum. Não pergunta por quê.”**

### Mockup 09 — A Bigorna

- Forja/estação
- Recursos disponíveis
- Receitas
- Craftar/Forjar
- Refinar
- Equipar
- Viajar
- Equipamento atual

### Mockup 10 — Primeira forja T1 → T2

- Arma atual
- Materiais
- Lastro
- Resultado
- Novo tier
- Ação de forjar

### Mockup 11 — Têmpera

A Têmpera é uma informação da arma/equipamento, não precisa de uma tela exclusiva.

- Nome da arma
- Tier
- Fama da arma
- Têmpera atual
- Multiplicador
- Indicação visual de calor/uso contínuo
- Informação de resfriamento/reset ao trocar de arma

### Mockup 12 — Outro Portador

- Seu personagem
- Outro Portador
- Equipamentos visíveis
- Nome
- Arma/linha quando necessário
- Conversa

Fala-âncora: **“Sua arma tem gosto. Cuidado quando ela começar a ter voz.”**

Observação: na PoC isso é um NPC/encontro narrativo; multiplayer/PvP ficam fora do primeiro corte.

### Mockup 13 — Segunda arma

- Segunda arma
- Arma atual
- Comparação
- Fama de cada arma
- Têmpera atual
- Ação de equipar

### Mockup 14 — Troca de arma

Visualizar explicitamente:

- Arma anterior → esfria
- Fama → permanece
- Têmpera → reseta
- Nova arma → assume o uso ativo

### Mockup 15 — Saída / A Encruzilhada

- Caminho para fora
- Destino
- Ação Viajar
- Resumo curto do que foi aprendido
- Guia quando necessário
- A Encruzilhada como destino

## 25.6 Painéis e telas funcionais pós-tutorial

### Personagem

- atributos
- aparência
- equipamento
- progressão
- identidade visual

### Equipamento

- arma
- peças visuais do personagem
- comparação
- tiers
- Fama/Têmpera quando aplicável

### Inventário

- itens
- quantidade
- slots
- peso

### A Litania

- progresso
- linhas das armas
- desbloqueios
- Fama

### Crafting / Forja

- receitas
- materiais
- custo
- resultado

### Refino

- material bruto
- material refinado
- custos
- resultado

### Mapa

- zonas
- conexões
- localização atual
- destinos disponíveis

### Ação em andamento

Pode ser um painel da tela principal ou uma janela ampliada:

- ação
- alvo/recurso
- progresso
- tempo
- recompensa

### Eventos / Log

- ganhos
- derrotas
- coleta
- forja
- equipamento
- eventos narrativos

### Social / Portadores

Pós-PoC:

- outros Portadores
- encontros
- facções
- PvP
- risco

### Configurações

- opções gerais de interface e jogo

## 25.7 Elementos visuais que ainda precisam de especificação própria

O manual já define a linguagem, mas estes ativos ainda precisam de uma ficha visual/padrão de produção:

- Catador base em isométrico
- Guia em isométrico
- Mito sem Pacto T1
- Armas T1/T2/T3 das linhas Guerreiro, Caçador e Conjurador
- Peças de equipamento visíveis
- Kit de cenário da A Ressaca
- Kit de cenário da A Bigorna
- Kit de cenário do Verde Surdo
- Kit de cenário da Costela
- Recursos visuais: Tronco, Pedra Bruta, Bétula, Couro, Cobre, Algodão
- Forja/Bigorna
- Ícones
- Botões e estados de botão
- Slots de equipamento/inventário
- Barras de progresso/vida
- Janela de diálogo
- Indicadores de tutorial
- Efeitos de combate/coleta/forja

## 25.8 Regra de produção para iniciantes

A arte deve ser pensada para ser **repetida muitas vezes**. Um asset bom para A Escória é aquele que você consegue redesenhar 100 vezes mantendo a mesma linguagem.

Prioridade de produção:

**silhueta → cor → equipamento → detalhe**.

Regra prática:

- objeto comum: 2–3 cores;
- personagem: 3–6 cores principais;
- sombreamento: base + sombra + pequeno destaque;
- cenário: poucos elementos reutilizáveis;
- animações: poucas poses reutilizáveis;
- efeitos: simples e legíveis.

## 25.9 O que evitar

- Realismo
- Pintura digital detalhada
- Anatomia realista
- Texturas fotográficas
- Partículas excessivas
- Iluminação cinematográfica pesada
- Gradientes complexos em personagens
- Cenários lotados
- Linhas ultrafinas
- Pixel art detalhada
- Fantasia medieval excessivamente ornamentada
- UI dourada/luxuosa
- Efeitos neon

## 25.10 Fórmula visual consolidada

**CARTOON 2D + CHIBI + CONTORNO FORTE + CORES CHAPADAS + SOMBRAS SIMPLES + SILHUETAS LEGÍVEIS + EQUIPAMENTO VISÍVEL + ANIMAÇÕES CURTAS + CENÁRIOS SIMPLES + MUNDO SUJO = IDENTIDADE VISUAL DE A ESCÓRIA**

## 25.11 Referência de escopo

Os 31 itens de mockup não devem ser tratados como 31 telas isoladas. A direção recomendada é **poucas telas estruturais + muitos estados e painéis**, especialmente porque o jogo é idle e a tela da zona deve concentrar a maior parte da experiência.

A PoC continua limitada aos 8 passos na Margem Calada, sem multiplayer real, PvP ou gank. O objetivo é validar o loop **zona → ação → recurso/Fama → progressão**.

## Validação de conteúdo — Mensagens 1 e 2

Esta seção consolida e valida as duas mensagens de referência fornecidas pelo autor em 05/09/2026. O objetivo é registrar o conteúdo de design e direção de arte que deve permanecer como referência para os mockups e para a produção visual.

### Status de cobertura

**Mensagem 1 — Fluxo, telas, tutorial e UI:** o conteúdo conceitual está contemplado no manual, incluindo fluxo de entrada, tela principal da zona, ações por zona, passos 1–8 do tutorial, estados de combate/coleta/forja/Têmpera/segunda arma/saída, painéis pós-tutorial e a distinção entre telas reais, estados e painéis. Os elementos específicos de cada etapa também estão consolidados nas seções de telas do tutorial.

**Mensagem 2 — Manual Visual:** o conteúdo está contemplado nas seções de identidade visual, silhueta, personagens, equipamentos, cores, contorno, sombreamento, textura, cenários, perspectiva, animação, efeitos, interface, ícones, botões, texto diegético, monstros, zonas, prioridade visual, exclusões, regras positivas e princípio de produção.

### Decisão visual consolidada

**Perspectiva oficial:** isométrico estilizado / visão 3/4 simplificada. Não buscar isometria matemática perfeita. A prioridade é profundidade suficiente para leitura do personagem e do equipamento, espaço para animação e produção simples.

**Regra de produção:** desenhar pensando em reutilização e consistência. Preferir poucos assets, silhuetas fortes, evolução por variantes e animações curtas.

### Catálogo de telas e estados para mockups

#### Entrada

1. Criar conta — logo A Escória; e-mail; senha; confirmar senha; criar conta; link para entrar; validações; fundo ilustrado simples.
2. Login — logo; e-mail; senha; entrar no mundo; recuperação de senha; criar conta; erros/validações.
3. Seleção de personagem — personagens; personagem selecionado; nome; nível; equipamento principal; última região/zona; criar; excluir; entrar; estado vazio.
4. Criação de personagem — nome; aparência simples; prévia grande; criação; voltar. A escolha inicial de linha/arma permanece **condicional**, não tratada como classe tradicional enquanto o design não definir isso.

#### Tutorial / PoC

1. Chegada — A Ressaca: Catador sem equipamento, Guia, cenário, marcador, nome da zona, HUD básico, diálogo e objetivo.
2. Conversa com o Guia: retrato/nome, caixa de diálogo, continuar, cenário e objetivo.
3. Recebimento da primeira arma: recompensa, arma T1, inventário, destaque do slot, diálogo e indicação da próxima ação.
4. Equipar primeira arma: personagem, arma na mão, slot, atributos e visual antes/depois.
5. Primeiro combate: Catador equipado, Mito sem Pacto T1, vida, dano, ação Matar, alvo, progresso/tempo, fila de habilidades quando aplicável, animações, efeitos e Guia observando.
6. Resultado do combate: inimigo derrotado, recompensa, Fama, Lastro/drops quando aplicável, log, progressão e fala do Guia.
7. Primeira coleta: A Ressaca, ações Coletar/Matar, recurso da zona, personagem coletando, progresso, tempo, recurso recebido, Fama e log.
8. Coleta em andamento: loop de coleta, progresso, contador/tempo, recurso, Fama, ação destacada, troca de ação e log.
9. A Bigorna: recursos, receitas, equipamento, custo em Lastro, resultado e ações Forjar/Refinar/Equipar.
10. Primeira forja T1→T2: material, Lastro, arma, resultado e novo tier.
11. Têmpera: arma, tier, Fama da arma, Têmpera atual, multiplicador, sinal visual de calor e informação de resfriamento/reset na troca. A Têmpera não precisa ser uma tela independente.
12. Outro Portador: personagem do jogador, outro Portador, equipamentos visíveis, nome, arma, linha/facção quando aplicável e conversa. Na PoC, sem multiplayer real.
13. Segunda arma: arma atual, segunda arma, comparação, Fama, Têmpera e escolha.
14. Troca de arma: arma anterior esfria; Fama permanece; Têmpera reseta; confirmação visual.
15. Saída / A Encruzilhada: caminho, destino, viajar, resumo do aprendizado e próximo destino.

#### Jogo completo

1. Tela principal da zona — tela central do jogo; cena isométrica; personagem e animação; painel do personagem/equipamento; ação atual/progresso; log; inventário; ações da zona; navegação.
2. Personagem — atributos, aparência, equipamento e progresso.
3. Equipamento — arma, peito, elmo, botas, comparação e tiers.
4. Inventário — itens, quantidades, peso e slots.
5. A Litania — progresso do Catador, linhas de armas, desbloqueios e Fama.
6. Crafting — receitas, materiais e resultado.
7. Refino — materiais brutos/refinados e custos.
8. Mapa — zonas, conexões e localização atual.
9. Viagem — destino, transição e estado de deslocamento; preferencialmente como estado/painel da experiência de zona, não como tela independente obrigatória.
10. Eventos/Log — ganhos, derrotas, coleta, forja, troca de equipamento e eventos narrativos.
11. Social/Portadores — posteriormente: outros Portadores, encontros, facções, PvP e risco.
12. Configurações.

### Arquitetura de telas

Os 31 itens acima **não equivalem a 31 telas independentes**. Muitas funções são painéis ou estados da tela principal. Em especial, combate, coleta, viagem, Têmpera, recompensa e resultado devem ser tratados prioritariamente como estados/painéis da experiência da zona, evitando multiplicação desnecessária de cenas.

### Tela principal — composição de referência

```
┌──────────────────────────────────────────────┐
│ FAMA     LASTRO      ZONA           MENU     │
├──────────┬───────────────────────┬───────────┤
│          │                       │           │
│ PERSONA- │                       │  AÇÃO     │
│ GEM      │     CENA DA ZONA      │           │
│          │                       │  PROGRESSO│
│ EQUIPA-  │   personagem animado  │           │
│ MENTO    │   executando ação     │  EVENTOS  │
│          │                       │           │
├──────────┴───────────────────────┴───────────┤
│ INVENTÁRIO            │ AÇÕES DA ZONA         │
└────────────────────────┴─────────────────────┘
```

A cena central é isométrica e funciona como quadro de observação. Informações importantes não devem depender apenas da imagem central; elas precisam existir nos painéis.

### Ações por zona — referência da PoC

**A Ressaca:** Coletar, Matar.

**A Bigorna:** Craftar, Refinar, Equipar, Viajar.

**O Verde Surdo:** Coletar, Matar, Viajar.

**A Costela:** Coletar, Matar, Viajar.

### Movimento observável do idle

A experiência visual da zona deve permitir que o personagem faça pequenos deslocamentos e loops: **andar alguns passos → parar → executar a ação → repetir**. O objetivo é um mundo observável, não uma tela estática.

### Mensagens e falas de referência

**Passo 1:** “Mais um. Anda. A maré não traz quem ela não pode usar.”

**Passo 2:** “Pega. Pra você não morrer na primeira noite.”

**Passo 2 — frase adicional usada anteriormente:** “Não é prêmio. É ferramenta. Trata ela como sucata que ainda serve.” Essa frase fica registrada como texto de trabalho/referência, mas não substitui as falas canônicas já definidas no GDD.

**Passo 3:** “Vai. Deixa o braço fazer o que ele lembra — você só assina depois.”

**Após o primeiro combate:** “Engraçado, né. A lasca lembra antes de você.”

**Passo 4:** “Isso aqui dá o que não existe em lugar nenhum. Não pergunta por quê.”

**Passo 4 — formulação usada anteriormente:** “Guarda tudo. Aqui até chão firme custa.” Registrar como referência de texto, distinguindo-a das falas canônicas confirmadas do GDD quando necessário.

**Passo 6:** “Sua arma tem gosto. Cuidado quando ela começar a ter voz.”

**Passo 8:** “Lá fora, sabem o que você carrega. E o que você carrega escolhe com quem você se alinha.”

### Direção de arte — fórmula consolidada

**A ESCÓRIA = CARTOON 2D + CHIBI + CONTORNO FORTE + CORES CHAPADAS + SOMBRAS SIMPLES + SILHUETAS LEGÍVEIS + EQUIPAMENTO VISÍVEL + ANIMAÇÕES CURTAS + CENÁRIOS SIMPLES + MUNDO SUJO.**

### Regras visuais consolidadas

- Silhueta → cor → equipamento → detalhe.
- Personagem chibi, reconhecível em tamanho pequeno.
- Equipamento é identidade visual, não apenas número de atributo.
- Objetos devem ser simples o suficiente para serem redesenhados muitas vezes.
- Evolução T1/T2/T3 deve preferir variantes sobre redesenhos completos.
- Armas recebem mais atenção visual que pequenos acessórios.
- Mundo com marrom-terra, ferrugem, cinza-chumbo, carvão, bege/sal, verde-musgo e azul dessaturado; acentos controlados.
- Objeto comum: 2–3 cores; personagem: 3–6 cores principais.
- Contorno forte, consistente e simples.
- Sombreamento preferencial: base + sombra + pequeno brilho.
- Textura principalmente no ambiente; personagens relativamente limpos.
- Cenários compostos por poucos assets reutilizáveis; aparência de riqueza sem grande quantidade de desenho.
- Isométrico estilizado / 3/4 simplificado como padrão de câmera visual.
- Animações curtas e reutilizáveis; idle, combate, coleta, forja e viagem.
- Efeitos simples; não ocultar o personagem.
- UI com painéis ancorados, retângulos arredondados, contorno escuro, fundo escuro, ícones simples e números destacados.
- Ícones seguem a mesma linguagem de contorno e silhueta.
- Botões devem ter estados normal, selecionado e ativo, sem excesso de efeitos.
- Tipografia sem serifa, grossa, legível, arredondada e de alto contraste.
- Texto diegético curto, oral, gasto e pragmático; evitar jargão de interface dentro do mundo.
- Monstros simples, com silhueta distinta e elemento visual marcante; famílias evoluem por tier.
- Cada zona deve ser reconhecida por poucos elementos visuais fortes.
- Na tela principal: **personagem > ação > zona > UI secundária**.

### Evitar

Realismo; pintura digital detalhada; anatomia realista; textura fotográfica; excesso de partículas; iluminação cinematográfica pesada; gradientes complexos em personagens; cenários superlotados; linhas ultrafinas; pixel art detalhada; fantasia medieval extremamente ornamentada; UI dourada/luxuosa; efeitos neon; personagens extremamente musculosos; rostos realistas.

### Princípio de produção

> “Desenhe como se você tivesse que redesenhar o mesmo objeto 100 vezes.”
> 

Uma espada difícil de desenhar é um mau asset. Uma espada simples e memorável é um bom asset. O mesmo vale para monstros, cenários e personagens.

### Consequência para o desenvolvimento

O objetivo não é aprender a desenhar “bonito” antes de iniciar o jogo; é aprender a desenhar **consistentemente**. A produção deve começar por um pequeno conjunto de assets: Catador, espada, barril, pedra, Mito, fogueira e outros elementos básicos, submetidos às mesmas regras de proporção, contorno, cor e perspectiva.