# Xray server template

This folder contains a starter Xray-core configuration for a self-managed VPN server.
It is a template, not a live server. Replace every placeholder and configure TLS before exposing it.

## Requirements
- A Linux VPS with Docker and Docker Compose
- A domain name pointing to the VPS
- A valid TLS certificate and private key for that domain
- A unique UUID for each client

## Start
1. Copy `config.example.json` to `config.json`.
2. Replace `REPLACE_WITH_CLIENT_UUID`, `vpn.example.com`, and certificate paths.
3. Run `docker compose up -d` from this directory.
4. Check logs with `docker compose logs -f xray`.

Do not commit private keys, real UUIDs, or production configuration to a public repository.
The web dashboard currently does not call this service; subscription links generated in the UI are placeholders.
