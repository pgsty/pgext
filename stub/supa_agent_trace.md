## Usage

Sources:

- [Official upstream README](https://github.com/jessevent/supa-agent/blob/e2a69b5c551d4925358252724788887bcf1cb862/supabase/extensions/supa_agent_trace/README.md)
- [Official extension control file (supa_agent_trace.control)](https://github.com/jessevent/supa-agent/blob/e2a69b5c551d4925358252724788887bcf1cb862/supabase/extensions/supa_agent_trace/supa_agent_trace.control)
- [Official extension SQL (supa_agent_trace--0.1.0.sql)](https://github.com/jessevent/supa-agent/blob/e2a69b5c551d4925358252724788887bcf1cb862/supabase/extensions/supa_agent_trace/supa_agent_trace--0.1.0.sql)

`supa_agent_trace` — Both the extension and the DevTool authenticate with the Supabase **Management API** (account-level OAuth) — neither holds a project GoTrue session. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION supa_agent_trace;

select dbdev.install('jessevent@supa_agent_trace');
create extension "jessevent@supa_agent_trace";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `public.agent_trace_broadcast()` is an extension function and returns `trigger`.
- `public.agent_trace_prune(retention interval default interval '7 days')` is an extension function and returns `bigint`.
- `public.agent_trace_topic(uid uuid)` is an extension function and returns `text`.
- `public.agent_trace_events` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- Install the confirmed extension dependencies first: `pgcrypto`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
