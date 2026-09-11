# Mecânicas — Overview

As mecânicas são **âncoras inegociáveis** (têm spec). A narrativa se adapta a elas.

```mermaid
flowchart TD
    ZONA["Zona define a ação"] --> ACAO{Ação}
    ACAO -->|coletar| COL["recursos + Fama"]
    ACAO -->|lutar| CMB["drops + Lastro + Fama"]
    ACAO -->|forjar| CRA["equipamento (consome Lastro)"]
    COL & CMB --> CRA --> GEAR["Loadout"]
    GEAR --> TIER["Tier ↑ = deus desperta"]
    CMB --> FAMA["Fama"] --> DB["Destiny Board"] --> TIER --> CMB
```

## Duas moedas, dois eixos (não embolar)

```mermaid
flowchart LR
    FAMA["Fama (lembrança)"] --> DB["Destiny Board / desbloqueios"]
    LASTRO["Lastro (resíduo)"] --> USO["reparo / craft / viagem / ancorar"]
```

Matar um mob com espada T2 rende **as duas coisas, separadas**: Fama de espada T2 **e** Lastro.

## Vínculo mecânica/narrativa

- Tier ↑ = o deus toma mais de você (Pilar 2).
- Fama = lembrança (mantém o deus vivo).
- Lastro = realidade calcificada (miniatura do mundo do Incrédulo).