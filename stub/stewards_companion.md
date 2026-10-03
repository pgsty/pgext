## Usage

Sources:

- [packs/companion/extension/README.md](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/packs/companion/extension/README.md)
- [packs/companion/extension/stewards_companion.control](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/packs/companion/extension/stewards_companion.control)
- [packs/companion/extension/stewards_companion--0.3.0.sql](https://github.com/cpuchip/pg-ai-stewards/blob/c78212cf4dbca4a711cbc6416e06dad25d576b1d/packs/companion/extension/stewards_companion--0.3.0.sql)

`stewards_companion` 0.3.0 is the SQL companion pack for `pg_ai_stewards`. It adds reminders, approval helpers and runtime tool registration in the `companion` and `forge` schemas. It inherits the core PostgreSQL 18 deployment boundary.

### Core workflow

```sql
CREATE EXTENSION pg_ai_stewards CASCADE;
CREATE EXTENSION stewards_companion;
SELECT * FROM companion.reminders LIMIT 5;
SELECT * FROM forge.forged_tools LIMIT 5;
```

The pack control and versioned SQL files must be installed separately; the ordinary core image build does not automatically include them. Creation requires elevated extension-install privileges. The pack has no native library or preload entry of its own.

### Tools and removal

Reminder helpers create, list, cancel and claim reminders; a separate client delivers them. Forge registration applies approved SQL tools inside the surrounding core workflow. Restrict those privileges and review the exact function definitions before approval. Version 0.3.0 resets loop counters when unsticking work items.

Before removal, call `companion.companion_uninstall()` to deactivate the pack’s tool definitions and shrink its write allowlist, then use `DROP EXTENSION stewards_companion`. A normal drop refuses while runtime-forged functions remain; `CASCADE` also destroys those functions. Pipeline and history rows remain, but the extension-owned reminder table is dropped. Its configuration-dump registration carries reminder contents through backup/restore, not through dropping the extension.
