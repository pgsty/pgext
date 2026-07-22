## Usage

Sources:

- [Official README for version 0.1.0](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/README.md)
- [Official extension SQL](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/sql/pgclaw--0.1.0.sql)
- [Official background-worker implementation](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/src/worker.rs)

`pgclaw` stores an LLM-agent specification in the `claw` data type and asynchronously processes inserted or updated rows with a PostgreSQL background worker. It is useful for experimental agent-driven workflows in which a model writes values back to a row; it also supports reusable agents, channel sessions, and optional Claude Code workspaces.

### Configuration and Core Workflow

The background worker requires preload-time configuration and a restart:

```conf
shared_preload_libraries = 'pgclaw'
pgclaw.database = 'appdb'
pgclaw.api_provider = 'anthropic'
pgclaw.api_key = 'replace-with-a-secret'
```

Then create the extension, add a primary-key table with a `claw` column, and register it:

```sql
CREATE EXTENSION pgclaw;

CREATE TABLE tickets (
    id bigserial PRIMARY KEY,
    title text,
    priority text,
    agent claw DEFAULT claw(
        'Set priority to low, medium, high, or critical from the ticket data.'
    )
);

SELECT claw_watch('tickets'::regclass);
INSERT INTO tickets(title) VALUES ('Production login returns HTTP 500');

SELECT id, priority FROM tickets;
SELECT id, error, created_at, done_at FROM claw.queue ORDER BY id DESC;
```

The trigger writes row JSON to `claw.queue`; the worker calls the configured provider, parses its response, and updates the source row. `claw_unwatch(regclass)` removes the trigger. `claw_model()`, `claw_prompt()`, and `claw_agent_id()` inspect stored values. Reusable definitions live in `claw.agents`; sessions, messages, inbox/outbox, channel bindings, heartbeats, and cron definitions live in the corresponding `claw` tables.

### Operational and Security Boundaries

`claw_watch()` requires a primary key and at least one `claw` column. Processing is asynchronous and can fail; monitor `claw.queue.error` and design idempotent retries. Row contents are sent to an external LLM endpoint, so minimize exposed fields and obtain appropriate data-handling approval. The API key is a superuser-only GUC but still resides in server configuration/process memory.

Agents with a `workspace` can create directories and invoke Claude Code, which can read/write files and run tools under the PostgreSQL service account. Treat that mode as remote code execution and isolate it from database hosts carrying important data. Version 0.1.0 is an early design; test provider responses, permissions, trigger recursion protection, rate limits, transaction visibility, and crash recovery before relying on it.
