# Intenção de Layout

> 🖼️ **Intenção de interface, não especificação.** Registrada em 25/08/2026. Descreve o que a tela precisa **comunicar** e por quê. Componentes, telas e implementação vivem no repositório.
> 

> 
> 

> **Nada disto é da PoC.** Na PoC o quadro central pode ser uma imagem estática ou nem existir.
> 

## O problema com o padrão do gênero

A maioria dos jogos idle é uma pilha de menus. Funciona, mas custa toda a atmosfera — e n'A Escória a atmosfera **é** o produto. Um jogo sobre deuses escondidos em metal não pode parecer uma planilha.

## A decisão — layout ancorado com quadro de observação

Tela dividida em painéis fixos, com um **quadro no centro** mostrando a cena da zona e o personagem executando a ação atual.

Ancorado, não flutuante: painéis ocupam lugar próprio e nada esconde nada. Referência de época são os MMOs de navegador e os RPGs online dos anos 2000 — moldura visível, densidade de informação alta, sem minimalismo.

Distribuição pretendida: personagem e loadout de um lado, ação em andamento e informação da zona do outro, inventário e ações disponíveis embaixo, quadro no meio.

## Por que o quadro existe — o argumento temático

Não é enfeite. É o **Pilar 2** renderizado.

O jogo diz que o portador não executa o golpe — ele sobrevive ao fato de o golpe passar por ele. Uma janela onde você **assiste** seu personagem agir sem controlá-lo não é limitação técnica disfarçada de escolha: é a ficção. Em outros idles a ausência de controle é defeito aceito; aqui é o tema.

E resolve outros dois pilares de graça:

- **"Você é o que veste"** só é abstrato porque loadout é lista. Se o personagem aparece vestindo, identidade vira visível — a diferença entre saber que você tem uma espada e ver que você é o sujeito da espada.
- **"Zona define a ação"** deixa de ser regra de menu e vira lugar. A Ressaca com maré e sucata é outra coisa que uma entrada de lista escrita "A Ressaca".

Ataca também um buraco de mercado: a maioria dos "idle MMOs" são clickers de um jogador só, fantasiados. Ver outros portadores na zona é o que faz parecer MMO.

## O quadro pode ser pequeno

Em MMO tradicional o mundo precisa ocupar a tela porque o jogador **anda** nele — precisa ver o inimigo chegando, mirar, desviar. Aqui não há movimentação nem mira. Você não joga dentro do quadro; você olha para ele.

Então ele pode ser um retângulo modesto e ainda funcionar — e isso derruba o custo de arte, que é o único custo real desta ideia.

## Regra que sustenta o desligar

O quadro é opcional: quem quer deixar o jogo aberto em segundo plano ou minimizado pode desligá-lo para poupar recurso.

Para isso funcionar, **toda informação mora nos painéis, nunca no quadro.** Ele ilustra; não informa. Se algum dado só existir ali, desligar vira desvantagem competitiva e a opção deixa de ser livre.

Com os painéis ancorados, desligar o quadro também não deixa buraco: os painéis se expandem e ocupam o espaço.

## Questões em aberto

**Celular.** Jogo idle vive em tela pequena, e layout ancorado de três lados não cabe. Vai precisar de outro arranjo — provavelmente abas no lugar de painéis, com o quadro no topo. Não resolvido.

**Custo de arte.** Sprite por peça de equipamento visível, animação por ação, cenário por zona. É o gargalo real. Versão barata que captura quase todo o efeito: **cena estática por zona com silhueta do personagem**, sem animação. Cinco imagens cobrem a PoC inteira; animação entra por cima depois, sem refazer nada.

**Tom das imagens de referência.** Modelos de imagem puxam para MMO alegre por padrão — céu azul, verde vivo, vila povoada. É o oposto d'A Escória. Qualquer referência visual precisa empurrar com força para ferrugem, sal e cinza.

---

## Prompt para gerar imagem de conceito

Usado em 25/08/2026 com bom resultado. Guardado para reuso.

```
Mockup de interface de um MMORPG idle para PC, em pixel art, estilo old-school
de MMO de navegador dos anos 2000. Layout ancorado, não flutuante.

ESTRUTURA DA TELA:
- Centro: um quadro retangular médio (não ocupa a tela toda) mostrando a cena do
  mundo. É uma janela de observação, não uma área jogável — o personagem age
  sozinho, o jogador apenas assiste.
- Painel esquerdo: retrato do personagem, barras de vida, e os slots de
  equipamento vestido (arma, peito, elmo, botas).
- Painel direito: ação em andamento com barra de progresso, e informações da
  zona atual.
- Painel inferior: grade de inventário e a lista de ações disponíveis na zona.
- Todos os painéis são molduras de metal escuro gasto, com rebites, como chapa
  de ferro velha aparafusada. Nada dourado, nada ornamentado.

A CENA NO QUADRO CENTRAL:
Uma praia de pedra e sal ao amanhecer. Maré cinzenta trazendo destroços. Na
areia, entulho de civilizações que nunca se conheceram, misturado: um pilar de
pedra clara caído ao lado de um tronco retorcido, quilha de barco enterrada,
metal oxidado por toda parte.

Duas figuras pequenas: um homem magro em roupas pobres de catador de sucata,
sem armadura, segurando uma lâmina simples e escura — e a alguns passos, um
velho de pé, imóvel, observando sem ajudar. Ao fundo, muito distante, silhuetas
de ruínas.

PALETA E CLIMA:
Ferrugem, cinza-chumbo, sal esbranquiçado, verde-musgo apagado, marrom
oxidado. Céu com uma cor levemente errada — nem azul, nem cinza comum.
Iluminação baixa e difusa. Atmosfera de ferro-velho, abandono e sobrevivência.
Melancólico e gasto, não heroico.

EVITAR:
Céu azul claro, verde vivo, vila alegre e povoada, cores saturadas, brilho
dourado, magia colorida, ícones cintilantes, estética de fantasia medieval
otimista, logotipos ou marcas, texto legível em outro idioma.
```

**Ao reusar:** variar o tamanho do quadro central entre as tentativas (metade da tela, um terço) — a proporção certa é difícil de decidir no abstrato e óbvia de reconhecer na imagem. Traduzir para inglês costuma melhorar o resultado.

**Aviso.** Modelos de imagem pegam vocabulário do projeto e chutam o significado — uma geração de teste devolveu "Escória" como recurso coletado, Têmpera como botão que se compra, e inventou moedas que não existem. Bonito e canonicamente errado. Trate resultado de imagem como referência de **tom**, nunca de conteúdo.