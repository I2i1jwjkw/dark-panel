# Dark Panel

A Persian-first VPN management panel scaffold with a Vue 3 frontend, Go API, PostgreSQL, Caddy, and an Xray configuration placeholder.

پنل مدیریت VPN با رابط فارسی، فرانت‌اند Vue، بک‌اند Go، پایگاه داده PostgreSQL، پراکسی Caddy و پوشه پیکربندی Xray.

## Project structure

```text
dark-panel/
├── .env.example
├── .gitignore
├── docker-compose.yml
├── README.md
├── backend/
│   ├── Dockerfile
│   ├── go.mod
│   ├── cmd/main.go
│   └── internal/
│       ├── config/config.go
│       ├── database/db.go
│       ├── models/user.go
│       ├── handlers/auth.go
│       ├── handlers/user.go
│       ├── handlers/sub.go
│       └── i18n/locales/
│           ├── active.fa.toml
│           └── active.en.toml
├── frontend/
│   ├── Dockerfile
│   ├── nginx.conf
│   ├── package.json
│   ├── vite.config.ts
│   ├── tailwind.config.js
│   ├── index.html
│   └── src/
│       ├── main.ts
│       ├── App.vue
│       ├── assets/main.css
│       ├── router/index.ts
│       ├── i18n/index.ts
│       ├── locales/fa.json
│       ├── locales/en.json
│       ├── components/
│       │   ├── AnimatedBackground.vue
│       │   ├── GlassCard.vue
│       │   └── LangSwitcher.vue
│       └── views/
│           ├── LoginView.vue
│           └── DashboardView.vue
├── caddy/Caddyfile
└── xray/config.json
```

## Run with Docker Compose

1. Install Docker Engine and the Docker Compose plugin on a Linux VPS.
2. Clone this repository and enter its directory:
   ```bash
   git clone https://github.com/I2i1jwjkw/dark-panel.git
   cd dark-panel
   ```
3. Create your environment file:
   ```bash
   cp .env.example .env
   ```
4. Edit `.env`. Set a long random `JWT_SECRET`, a strong `ADMIN_PASSWORD`, a unique database password in `DATABASE_URL` and `POSTGRES_PASSWORD`, and your public hostname in `PUBLIC_URL` and `PUBLIC_HOST`.
5. Set DNS for your hostname to the VPS, then start:
   ```bash
   docker compose up -d --build
   docker compose logs -f
   ```
6. Open `https://your-hostname`.

Caddy obtains TLS certificates automatically when DNS and ports 80/443 are correctly configured. Do not expose the API or database directly to the public internet.

## Xray status

The `xray/config.json` file is an empty-inbound starter configuration. It does not accept VPN clients yet. Configure a secure inbound, client credentials, TLS, firewall, and Xray service before use. The current compose stack does not provision or supervise Xray automatically.

## Important limitations

This repository is an early scaffold, **not production-ready**. The API currently has a health endpoint and starter handlers; frontend authentication/session handling and protected API access are not complete. Subscription creation is not connected to Xray and does not issue a working subscription. Traffic accounting, expiry enforcement, client provisioning, password rotation, rate limiting, CSRF protections, and audit logs still need implementation and testing.

Never commit `.env`, real client UUIDs, private keys, or production Xray configuration. Change all example credentials before deployment.

## Development

Frontend:
```bash
cd frontend
npm install
npm run dev
```

Backend:
```bash
cd backend
go run ./cmd
```
The backend requires a reachable PostgreSQL instance and the environment variables above.
