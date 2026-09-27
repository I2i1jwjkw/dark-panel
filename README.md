# 🦇 DARK PANEL

A Persian-first VPN management panel interface with a Go API, PostgreSQL, Vue 3, and Caddy.

پنل مدیریت با رابط کاربری شیشه‌ای، پشتیبانی رابط از پنج زبان و بک‌اند Go.

## Features in this repository

- Dark glass-style responsive interface
- Persian, English, Chinese, Russian, and Arabic language selector
- Login screen connected to `POST /api/auth/login`
- User list and create-user form connected to the API
- Docker Compose deployment scaffold and Caddy reverse proxy

## Run

1. Install Docker Engine and the Docker Compose plugin on a Linux VPS.
2. Clone the repository:
   ```bash
   git clone https://github.com/I2i1jwjkw/dark-panel.git
   cd dark-panel
   ```
3. Configure environment variables:
   ```bash
   cp .env.example .env
   nano .env
   ```
   Set a strong unique `JWT_SECRET`, `ADMIN_PASSWORD`, database password, and public hostname. Never use example credentials on a public server.
4. Point your domain's DNS to the VPS, allow inbound ports 80 and 443, then run:
   ```bash
   docker compose up -d --build
   docker compose logs -f
   ```
5. Open the hostname configured for Caddy.

## Important status

**This is an early-stage scaffold, not a production-ready VPN service.** The current Xray configuration is a placeholder and the compose stack does not provision working Xray clients. Subscription delivery, traffic accounting, expiry enforcement, and client provisioning are not implemented. The current API's user endpoints also need authentication middleware before public deployment.

The frontend is an interface and should not be interpreted as proof that the underlying VPN/subscription functions are operational. Do not expose this panel publicly until API authorization and deployment security have been reviewed.

Never commit `.env`, private keys, real client UUIDs, or production configuration.
