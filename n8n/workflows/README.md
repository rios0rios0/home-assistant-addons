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
3. Fill in the three placeholders:
   - **Telegram credential** -- `Credentials` > `New` > `Telegram API`, paste
     the bot token from `@BotFather`. Assign it to the `Send to Telegram`
     node.
   - **`chatId`** -- open
     `https://api.telegram.org/bot<TOKEN>/getUpdates` in a browser after
     sending `/start` to your bot, copy `result[0].message.chat.id`, and
     paste it into the `Send to Telegram` node parameters.
4. `Execute Workflow` and then trigger `Webhook` (use `Listen for test event`
   and POST a sample payload) to verify.
5. Toggle the workflow **Active**.
6. Copy the `Production URL` of the webhook and paste it into
   `dashboard.gitguardian.com > Settings > Alerting > Add Custom Webhook`.

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
