                 Browser / phone
                       │  HTTPS :443  (only open port)
                       ▼
              ┌─────────────────┐
              │  proxy (Caddy)  │  TLS, routing
              └───┬─────────┬───┘
         /api/*   │         │   everything else
                  ▼         ▼
        ┌──────────────┐  ┌───────────────────────────┐
        │ backend (Gin)│  │ frontend                  │
        │    :3000     │  │ dev: Vite :5173           │
        └──────┬───────┘  │ prod: static files served │
               │          │       by Caddy itself     │
               ▼          └───────────────────────────┘
        ┌──────────────┐
        │ db (Postgres)│  only the backend can reach it
        │    :5432     │
        └──────────────┘
