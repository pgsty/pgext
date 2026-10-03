## Usage

Sources:

- [README 0.1.0](https://github.com/snoutdata/snout-net/blob/6c309d0075e123ef88a55c521becd308046a49b0/README.md)
- [Control](https://github.com/snoutdata/snout-net/blob/6c309d0075e123ef88a55c521becd308046a49b0/snout_net.control)
- [SQL 0.1.0](https://github.com/snoutdata/snout-net/blob/6c309d0075e123ef88a55c521becd308046a49b0/sql/snout_net--0.1.0.sql)
- [Cargo](https://github.com/snoutdata/snout-net/blob/6c309d0075e123ef88a55c521becd308046a49b0/Cargo.toml)

`snout_net` 0.1.0 queues HTTP requests in PostgreSQL and sends them through a background worker after the submitting transaction commits. It targets PostgreSQL 17 and requires libcurl 7.85 or later. A superuser must preload the library, restart the server and install the extension in the configured database.

### Core Workflow

```conf
shared_preload_libraries = 'snout_net'
snout_net.database_name = 'postgres'
```

```sql
CREATE EXTENSION snout_net;
SELECT net.check_worker_is_up();

BEGIN;
SELECT net.http_post(
  url := 'https://example.com/hook',
  body := '{"event":"signup"}'::jsonb
) AS request_id \gset
COMMIT;

SELECT status_code, content, error_msg
FROM net._http_response WHERE id = :request_id;
```

The SQL example uses psql to retain the request ID. The response query can initially return no row; query again after the worker finishes. A rolled-back transaction sends nothing. Do not wait synchronously in the transaction that queued the request.

### Objects and Configuration

- `net.http_get`, `net.http_post`, `net.http_delete`: return a request ID; accept query parameters, headers and a timeout. The POST helper accepts JSON bodies.
- `net.http_request_queue` and `net._http_response`: pending requests and recorded responses.
- `net._http_collect_response`: collect a response; synchronous waiting belongs in a later transaction.
- `net.check_worker_is_up`, `net.wait_until_running`, `net.worker_restart`: worker health and restart helpers.
- `snout_net.database_name` selects the one database served; `snout_net.username` selects the worker role (default: bootstrap superuser).
- `snout_net.ttl` defaults to 6 hours; `snout_net.max_concurrent` to 200; `snout_net.max_timeout_ms` to 600000; `snout_net.max_response_bytes` to 64 MB.
- `snout_net.allowed_networks` permits explicitly listed internal CIDRs. Otherwise private, loopback, link-local and other non-public ranges are rejected, including after DNS resolution and redirects.
- `snout_net.worker_type` is set at server start; the other settings require a configuration reload. Library upgrades require a server restart.

### Delivery and Security Boundaries

Both queue and response tables are unlogged: they are not a durable job system across a database crash and are not replicated to physical standbys. A worker restarting mid-request may resend it, so receivers must tolerate duplicate effects. Requests to the same endpoint can arrive out of order. If recording a response fails because of a trigger or constraint, the worker logs a warning and does not resend that request.

The installation grants PUBLIC access to the schema, tables and sequences. Review those privileges, the worker role and outbound-network policy before exposing SQL access. The fixed `net` objects overlap with `pg_net`; this is not a side-by-side installation or an automatic migration procedure.
