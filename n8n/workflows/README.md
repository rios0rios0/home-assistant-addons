# n8n workflows

Importable workflow definitions for the n8n add-on. Kept in Git so the
automations survive an n8n database wipe. Credentials are **not** stored
here -- those stay in n8n's encrypted credential store.

## `gg-incident.json` -- GitGuardian -> Telegram

Receives a GitGuardian webhook payload, formats it as a Markdown message,
and sends it to a Telegram chat via a bot.

### Import

1. Open n8n (Home Assistant ingress: `Settings > Add-ons > n8n > Open Web UI`).
2. `Workflows` > menu > `Import from File` -> pick `gg-incident.json`.
3. Set up the Telegram node and fill in the two workflow placeholders:
   - **Telegram credential** -- `Credentials` > `New` > `Telegram API`, paste
     the bot token from `@BotFather`, then assign that credential to the
     `Send to Telegram` node.
   - **`chatId`** -- open
     `https://api.telegram.org/bot<TOKEN>/getUpdates` in a browser after
     sending `/start` to your bot, copy `result[0].message.chat.id`, and
     paste it into the `Send to Telegram` node parameters.
4. `Execute Workflow` and then trigger `Webhook` (use `Listen for test event`
   and POST a sample payload) to verify.
5. Toggle the workflow **Active**.
6. Copy the `Production URL` of the webhook and paste it into
   `dashboard.gitguardian.com > Settings > Alerting > Add Custom Webhook`.
   **Note:** GitGuardian must be able to reach this URL from outside your
   Home Assistant instance. If the generated `Production URL` is not publicly
   reachable, set the add-on `webhook_url` option to your public n8n base URL
   (this sets `WEBHOOK_URL` in the container) or otherwise expose n8n
   externally.

### Hardening

The webhook node is unauthenticated by default, which is fine on a trusted
LAN but risky once the `Production URL` is exposed externally. Before
publishing the URL, open the `Webhook` node and set `Authentication` to
one of:

- **Header Auth** / **Basic Auth** -- generate a long random secret, add it
  to the Telegram-facing node, and configure GitGuardian to send the
  matching header (see
  `dashboard.gitguardian.com > Settings > Alerting > Add Custom Webhook >
  Custom Headers`).
- **JWT Auth** -- if you already manage JWTs for your n8n instance.

Anyone who discovers an unauthenticated `Production URL` can trigger
Telegram messages through your bot, so do not skip this step for
internet-facing deployments.

### Payload assumptions

The `Format message` node reads these fields from the GitGuardian payload
(defensive -- accepts either `{ body: { incident: ... } }` or a flat
`incident` object):

| Field | Source |
|-------|--------|
| `secret_type` | `incident.secret_type` or `incident.detector.display_name` |
| `severity` | `incident.severity` |
| `status` | `incident.status` (drives the emoji/title) |
| `source` | `incident.source.full_name` or `incident.source.name` |
| `url` | `incident.incident_url` or `incident.gitguardian_url` |

Adjust the `jsCode` if GitGuardian changes its payload shape.
