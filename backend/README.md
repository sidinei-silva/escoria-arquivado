# Go Server


## Arquitetura
```mermaid
flowchart TD
    CLIENT([CLIENT]) -->|json| NETWORK["NETWORK<br/>HTTP / WS"]
    NETWORK -->|Command| QUEUE["COMMAND QUEUE<br/>channel"]
    QUEUE --> LOOP["GAME LOOP"]
    LOOP --> Handle["Handle()"]
    LOOP --> Tick["Tick()"]
    Handle --> ENGINE["GAME ENGINE<br/>regras + transição"]
    Tick --> ENGINE
    ENGINE --> STATE["GAME STATE"]
    STATE --> PG[(PostgreSQL)]
    STATE --> EVENTS["Events"]
    EVENTS --> WS(["WebSocket"])
```

## Fluxo de inicialização

```mermaid
flowchart TD
    A["main()"] --> B[config]
    B --> C[Postgres]
    C --> D[load initial game data]
    D --> E[create GameState]
    E --> F[create Engine]
    F --> G[create command channel]
    G --> H[start game loop]
    H --> I[start HTTP]
    I --> J[start WebSocket]
```

*Visualmente*:
```mermaid
flowchart TD
    PG[(PostgreSQL)] -->|"Load()"| GS[GameState]
    GS -->|Engine| GL[Game Loop]
    GS -->|Engine| NET[Network]
    GL --> CMD[Commands]
    NET --> CMD
```

## Estrutura pensada inicialmente para o jogo
```
escoria/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   │
│   ├── account/
│   │   ├── account.go
│   │   ├── repository.go
│   │   └── service.go
│   │
│   ├── character/
│   │   ├── character.go
│   │   ├── repository.go
│   │   └── service.go
│   │
│   ├── world/
│   │   ├── world.go
│   │   └── service.go
│   │
│   ├── combat/
│   │   └── service.go
│   │
│   ├── gamedata/
│   │   └── ...
│   │
│   ├── persistence/
│   │   └── postgres/ 
│   │       ├── account_repository.go
│   │       └── character_repository.go
│   │
│   └── network/
│       │
│       ├── http/
│       │   ├── server.go
│       │   ├── account_handler.go
│       │   └── character_handler.go
│       │
│       └── ws/
│           ├── server.go
│           ├── connection.go
│           └── handler.go
│
└── data/
```

E a dependência ficaria assim:
```
                  ┌───────────────┐
                  │     main      │
                  └───────┬───────┘
                          │
                 composição da aplicação
                          │
         ┌────────────────┼────────────────┐
         │                │                │
         ▼                ▼                ▼
      Gamedata        Services         Network
                           │          ┌─────┴─────┐
                           │          │           │
                           │         HTTP         WS
                           │          │           │
                           └──────────┴───────────┘
                                      │
                                      ▼
                                  Repository
                                      │
                                      ▼
                                  PostgreSQL
``` 