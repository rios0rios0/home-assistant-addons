# Home Assistant Add-on Ideas

Future add-on ideas for this repository. These are self-hosted tools that would complement Home Assistant but don't have widely available community add-ons yet.

---

## AI & Machine Learning

### Open WebUI
- **Project**: [github.com/open-webui/open-webui](https://github.com/open-webui/open-webui)
- **Description**: Feature-rich web UI for interacting with LLMs. Pairs naturally with the ollama-container add-on as a chat frontend.
- **Tech Stack**: Python (FastAPI), SvelteKit
- **Default Port**: 3000
- **HA Integration**: Provides a polished chat interface for local LLMs; can be exposed via ingress alongside Ollama
- **Complexity**: Medium

### LocalAI
- **Project**: [github.com/mudler/LocalAI](https://github.com/mudler/LocalAI)
- **Description**: Drop-in OpenAI API replacement that runs LLMs, image generation, audio transcription, and embeddings locally.
- **Tech Stack**: Go, C++
- **Default Port**: 8080
- **HA Integration**: OpenAI-compatible API enables direct use with HA's OpenAI conversation integration for automations and voice assistants
- **Complexity**: High (GPU support, large models)

### Whisper.cpp Server
- **Project**: [github.com/ggerganov/whisper.cpp](https://github.com/ggerganov/whisper.cpp)
- **Description**: High-performance local speech-to-text server using OpenAI's Whisper models compiled to C++.
- **Tech Stack**: C++
- **Default Port**: 8080
- **HA Integration**: Wyoming protocol compatible; plugs directly into HA's voice assistant pipeline for local STT
- **Complexity**: Medium

### Piper TTS
- **Project**: [github.com/rhasspy/piper](https://github.com/rhasspy/piper)
- **Description**: Fast, local neural text-to-speech engine with support for many languages and voices.
- **Tech Stack**: C++, Python
- **Default Port**: 10200 (Wyoming)
- **HA Integration**: Wyoming protocol compatible; local TTS for HA voice assistants without cloud dependency
- **Complexity**: Low

---

## Automation & Workflows

### n8n
- **Project**: [github.com/n8n-io/n8n](https://github.com/n8n-io/n8n)
- **Description**: Workflow automation platform with 400+ integrations, AI capabilities, and a visual flow editor.
- **Tech Stack**: TypeScript, Node.js
- **Default Port**: 5678
- **HA Integration**: Bridges HA with external services (CRM, email, databases, APIs) for complex multi-step workflows beyond HA's native automations
- **Complexity**: Medium

### Huginn
- **Project**: [github.com/huginn/huginn](https://github.com/huginn/huginn)
- **Description**: System for building agents that perform automated tasks online — monitoring websites, scraping data, sending notifications.
- **Tech Stack**: Ruby on Rails
- **Default Port**: 3000
- **HA Integration**: Create intelligent agents that feed data into HA entities or trigger automations based on external events
- **Complexity**: High

---

## Monitoring & Infrastructure

### Uptime Kuma
- **Project**: [github.com/louislam/uptime-kuma](https://github.com/louislam/uptime-kuma)
- **Description**: Self-hosted monitoring tool for HTTP(s), TCP, DNS, and other protocols with a clean status page.
- **Tech Stack**: Node.js, Vue.js
- **Default Port**: 3001
- **HA Integration**: Monitor HA instance health, network devices, and IoT endpoints; send alerts via HA notification services
- **Complexity**: Low

### Changedetection.io
- **Project**: [github.com/dgtlmoon/changedetection.io](https://github.com/dgtlmoon/changedetection.io)
- **Description**: Website change detection and notification service. Monitors web pages for content changes.
- **Tech Stack**: Python (Flask)
- **Default Port**: 5000
- **HA Integration**: Trigger HA automations when monitored websites change (price drops, availability, news updates)
- **Complexity**: Medium

### Ntfy
- **Project**: [github.com/binwiederhier/ntfy](https://github.com/binwiederhier/ntfy)
- **Description**: Simple HTTP-based push notification server. Send notifications to phones/desktops via PUT/POST.
- **Tech Stack**: Go
- **Default Port**: 8080
- **HA Integration**: Lightweight self-hosted push notification backend for HA automations; no vendor lock-in
- **Complexity**: Low

---

## Household & Organization

### Homebox
- **Project**: [github.com/sysadminsmedia/homebox](https://github.com/sysadminsmedia/homebox)
- **Description**: Home inventory and asset management system. Track items, locations, warranties, and maintenance schedules.
- **Tech Stack**: Go, Vue.js
- **Default Port**: 7745
- **HA Integration**: Track home assets and link them to HA areas/devices; trigger maintenance reminders via automations
- **Complexity**: Low

### Mealie
- **Project**: [github.com/mealie-recipes/mealie](https://github.com/mealie-recipes/mealie)
- **Description**: Recipe manager, meal planner, and shopping list with auto-import from recipe websites.
- **Tech Stack**: Python (FastAPI), Vue.js
- **Default Port**: 9925
- **HA Integration**: Display meal plans on HA dashboards; integrate shopping lists with smart displays
- **Complexity**: Medium

### Grocy
- **Project**: [github.com/grocy/grocy](https://github.com/grocy/grocy)
- **Description**: ERP-style household management for groceries, chores, and batteries with barcode scanning.
- **Tech Stack**: PHP
- **Default Port**: 9283
- **HA Integration**: Track pantry inventory with barcode scanning; automate shopping list notifications when stock is low
- **Complexity**: Medium

---

## Developer Tools

### Stirling PDF
- **Project**: [github.com/Stirling-Tools/Stirling-PDF](https://github.com/Stirling-Tools/Stirling-PDF)
- **Description**: Self-hosted PDF manipulation toolkit — merge, split, convert, compress, OCR, and more.
- **Tech Stack**: Java (Spring Boot)
- **Default Port**: 8080
- **HA Integration**: Process documents locally; useful for scanning/OCR workflows triggered by HA automations
- **Complexity**: Low

### IT Tools
- **Project**: [github.com/CorentinTh/it-tools](https://github.com/CorentinTh/it-tools)
- **Description**: Collection of 80+ developer and sysadmin utilities (encoders, converters, generators, network tools).
- **Tech Stack**: Vue.js
- **Default Port**: 8080
- **HA Integration**: Handy utility collection accessible from the HA dashboard for quick network and dev tasks
- **Complexity**: Low

### Linkwarden
- **Project**: [github.com/linkwarden/linkwarden](https://github.com/linkwarden/linkwarden)
- **Description**: Self-hosted bookmark and link manager with archiving, tagging, and collaboration features.
- **Tech Stack**: TypeScript, Next.js, PostgreSQL
- **Default Port**: 3000
- **HA Integration**: Bookmark HA community resources, automation snippets, and documentation in one place
- **Complexity**: Medium (requires PostgreSQL)

---

## Priority Recommendations

| Priority | Add-on | Rationale |
|----------|--------|-----------|
| High | Open WebUI | Natural companion to ollama-container; high user demand |
| High | Ntfy | Lightweight, fills a real gap for self-hosted notifications |
| High | Uptime Kuma | Universal need; low complexity |
| Medium | n8n | Powerful automation complement to HA |
| Medium | Stirling PDF | Very popular self-hosted tool; easy to package |
| Medium | Changedetection.io | Unique capability not native to HA |
| Low | LocalAI | Overlaps with Ollama; higher resource requirements |
| Low | Homebox | Niche use case but growing community |
| Low | Linkwarden | Less HA-specific; more of a general self-hosted tool |
