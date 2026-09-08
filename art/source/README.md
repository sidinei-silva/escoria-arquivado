# Art Source

Arquivos-fonte editáveis de toda a produção artística de **A Escória**.

Este diretório contém os arquivos originais usados para criar os assets do jogo. Eles devem permanecer editáveis e servir como fonte oficial para futuras alterações.

## Estrutura

```text
source/
├── characters/
├── equipment/
├── environment/
├── animations/
├── vfx/
└── ui/
```

## `characters/`

Arquivos-fonte dos personagens e suas partes visuais.

Exemplos:

```text
characters/
└── player/
    ├── base/
    ├── customization/
    └── variations/
```

Inclui:

- corpo base;
- cabeça;
- cabelo;
- braços;
- mãos;
- pernas;
- pés;
- variações visuais;
- NPCs;
- monstros, quando aplicável.

## `equipment/`

Arquivos-fonte dos equipamentos visuais.

```text
equipment/
├── helmets/
├── armor/
├── boots/
├── main_hand/
├── off_hand/
└── accessories/
```

Cada equipamento deve respeitar os mesmos padrões de:

- escala;
- pivô;
- orientação;
- anchors;
- layers;
- proporção.

## `environment/`

Arquivos-fonte dos elementos do mundo.

Exemplos:

- terrenos;
- pedras;
- árvores;
- madeira;
- estruturas;
- caixas;
- barris;
- objetos das zonas;
- elementos exclusivos de cada região.

Priorizar assets reutilizáveis.

## `animations/`

Arquivos-fonte das animações.

```text
animations/
├── player/
├── enemies/
├── npcs/
└── ui/
```

As animações devem utilizar os mesmos padrões de escala, pivô e posicionamento definidos pelos assets-base.

Exemplos:

```text
animations/
└── player/
    ├── idle.kra
    ├── walk.kra
    ├── attack.kra
    ├── gather.kra
    ├── hit.kra
    └── death.kra
```

Sempre que possível, priorizar animações curtas, reutilizáveis e adequadas ao caráter idle do jogo.

## `vfx/`

Arquivos-fonte dos efeitos visuais.

Exemplos:

- impactos;
- faíscas;
- dano;
- coleta;
- recompensas;
- fogo;
- Têmpera;
- efeitos de habilidades.

## `ui/`

Arquivos-fonte da interface.

Exemplos:

- botões;
- ícones;
- painéis;
- modais;
- barras;
- controles;
- HUD;
- elementos de menus.

## Formatos

Os arquivos editáveis devem permanecer em seus formatos originais.

Para ilustrações e sprites:

```text
.kra
```

Outros formatos editáveis podem ser utilizados quando fizerem sentido.

## Regras

- Nunca substituir o arquivo-fonte por um export.
- Não editar diretamente arquivos em `art/exports/`.
- Manter layers e grupos organizados.
- Preservar os guides e anchors necessários para assets modulares.
- Usar nomes claros e consistentes.
- Evitar arquivos duplicados sem necessidade.

## Pipeline

```text
SOURCE
  ↓
Edição no Krita
  ↓
Validação visual
  ↓
EXPORT
  ↓
Godot
```

`source/` é a **fonte oficial e editável** da arte.
