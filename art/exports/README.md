# Art Exports

Arquivos artísticos exportados e prontos para serem utilizados pelo **Godot**.

Este diretório contém apenas versões finais de assets que o jogo pode consumir.

## Estrutura

```text
exports/
├── characters/
├── equipment/
├── environment/
├── animations/
├── vfx/
└── ui/
```

## `characters/`

Sprites finais dos personagens.

Exemplos:

- corpo base;
- variações;
- NPCs;
- monstros;
- sprites direcionais.

## `equipment/`

Sprites finais dos equipamentos.

```text
equipment/
├── helmets/
├── armor/
├── boots/
├── main_hand/
├── off_hand/
└── accessories/
```

Os exports devem preservar os padrões definidos no `source/`.

## `environment/`

Assets finais utilizados no cenário.

Inclui:

- terrenos;
- objetos;
- estruturas;
- elementos das zonas;
- recursos de coleta;
- objetos decorativos reutilizáveis.

## `animations/`

Arquivos finais utilizados pelo sistema de animação do Godot.

```text
animations/
├── player/
├── enemies/
├── npcs/
└── ui/
```

Podem incluir:

- spritesheets;
- sequências de frames;
- sprites individuais;
- outros formatos necessários pelo Godot.

Exemplo:

```text
animations/
└── player/
    ├── idle.png
    ├── walk.png
    ├── attack.png
    ├── gather.png
    ├── hit.png
    └── death.png
```

## `vfx/`

Efeitos visuais exportados.

Exemplos:

- impacto;
- dano;
- coleta;
- fogo;
- Têmpera;
- habilidades;
- recompensas.

## `ui/`

Elementos finais da interface.

Exemplos:

- botões;
- ícones;
- painéis;
- modais;
- HUD;
- barras;
- elementos de menus.

## Formatos

O formato padrão para sprites e elementos 2D é:

```text
.png
```

Sempre que possível:

- preservar transparência;
- manter resolução definida pelo guia de escala;
- evitar compressão que prejudique a arte;
- manter nomes consistentes.

## Regras

- Não editar arquivos diretamente neste diretório.
- Alterações devem ser feitas em `art/source/`.
- Depois da alteração, gerar um novo export.
- Não guardar arquivos temporários ou testes aqui.
- Cada arquivo deve representar um asset utilizável pelo jogo.

## Pipeline

```text
art/source/
     ↓
Krita
     ↓
Exportação
     ↓
art/exports/
     ↓
Godot
```

`exports/` é a **versão pronta para o jogo**.
