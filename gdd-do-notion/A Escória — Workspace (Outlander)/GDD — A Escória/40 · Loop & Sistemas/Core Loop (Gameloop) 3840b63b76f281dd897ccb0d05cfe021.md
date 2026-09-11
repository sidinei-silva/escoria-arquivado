# Core Loop (Gameloop)

Três loops aninhados: **momento → sessão → meta**.

## Loop de momento

```mermaid
flowchart LR
    Z["Estou numa zona"] --> A["Ação (a zona define)"] --> C["Ciclo roda no servidor"] --> O["Saída: recursos + Lastro + Fama"] --> Z
```

## Loop de sessão (o que o tutorial valida)

```mermaid
flowchart TD
    COL["Coletar"] --> MAT["Recursos"]
    CMB["Lutar"] --> DROP["Lastro + drops"]
    MAT & DROP --> FORJA["Forjar / Refinar"]
    FORJA --> GEAR["Equipamento melhor / tier maior"]
    GEAR --> EQUIP["Equipar + fila"] --> CMB
    CMB --> FAMA["Fama"] --> DB["A Litania"]
    DB -->|desbloqueia| COL
    DB --> GEAR
```

> *zona define ação → ação gera recurso/fama → fama progride A Litania → progressão desbloqueia novo conteúdo.*
> 

## Loop de meta

```mermaid
flowchart LR
    SESSAO["Loop de sessão"] -->|Fama| T4["T4: cristaliza um deus"] -->|vínculo| T8["T8: pessoa-mito"] -->|rebirth| NOVA["Nova trajetória"] --> SESSAO
    T4 -.descobre.-> GUERRA["Guerra dos panteões"]
```

## Gatilhos de engajamento

1. Curiosidade da voz (o tier seguinte muda o que a arma diz).
2. Dilema de identidade (comprometer com uma arma e aprofundar vs. espalhar-se entre várias; ao trocar, a arma esfria — a Têmpera reseta — mas a Fama dela permanece). O custo real de ceder só cobra no T4.
3. Mistério descoberto (facções reagem ao loadout).