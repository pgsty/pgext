


## Usage

Sources: [official README](https://github.com/supabase/pg_net), [v0.20.3 release notes](https://github.com/supabase/pg_net/releases/tag/v0.20.3), [local package metadata](../db/extension.csv).

`pg_net` queues asynchronous HTTP and HTTPS requests from SQL. It creates the `net` schema, stores pending work in `net.http_request_queue`, and stores responses in `net._http_response`. A background worker uses `libcurl` to process queued requests.

The extension requires `shared_preload_libraries = 'pg_net'` and `libcurl >= 7.83`.

### GET Request

```sql
CREATE EXTENSION pg_net;

SELECT net.http_get(
  'https://postman-echo.com/get',
  params := '{"foo": "bar"}'::jsonb,
  headers := '{"API-KEY": "<key>"}'::jsonb,
  timeout_milliseconds := 1000
) AS request_id;
```

`net.http_get(url, params, headers, timeout_milliseconds)` returns a `bigint` request id.

### POST Request

```sql
SELECT net.http_post(
  'https://postman-echo.com/post',
  body := '{"key": "value"}'::jsonb,
  headers := '{"Content-Type": "application/json"}'::jsonb,
  timeout_milliseconds := 1000
) AS request_id;
```

Send one table row as JSON:

```sql
WITH selected_row AS (
  SELECT * FROM my_table LIMIT 1
)
SELECT net.http_post(
  'https://api.example.com/data',
  to_jsonb(selected_row.*)
) AS request_id
FROM selected_row;
```

### DELETE Request

```sql
SELECT net.http_delete(
  'https://api.example.com/resource/42',
  timeout_milliseconds := 2000
) AS request_id;
```

`net.http_delete(url, params, headers, timeout_milliseconds)` is `SECURITY DEFINER` and returns a `bigint` request id.

### Checking Responses

```sql
SELECT id, status_code, content, error_msg, created
FROM net._http_response
ORDER BY created DESC;
```

Failed requests can be identified from `status_code` and `error_msg`. The response table does not preserve every original request argument, so store request metadata separately if you need retry workflows.

### Configuration

```sql
SHOW pg_net.batch_size;
SHOW pg_net.ttl;
SHOW pg_net.database_name;
SHOW pg_net.username;
```

- `pg_net.batch_size`, default `200`, limits how many queued requests the worker processes per cycle.
- `pg_net.ttl`, default `6 hours`, controls response retention.
- `pg_net.database_name`, default `postgres`, selects the database where the worker runs.
- `pg_net.username`, default NULL, selects the worker connection user; NULL uses the bootstrap user.

Settings can be changed in `postgresql.conf` or with `ALTER SYSTEM`:

```sql
ALTER SYSTEM SET pg_net.ttl TO '1 hour';
ALTER SYSTEM SET pg_net.batch_size TO 500;
SELECT pg_reload_conf();
```

Changing `pg_net.database_name` or `pg_net.username` requires restarting the worker:

```sql
SELECT net.worker_restart();
```

### Caveats

- Pigsty metadata carries `pg_net` 0.20.3 for PostgreSQL 14-18, but local package notes say 0.20.3 is available only on `d12`, `d13`, `el10`, `u24`, and `u26`; `el8`, `el9`, and `u22` remain on 0.9.2 because of older `libcurl`.
- Upstream documents PostgreSQL 12+ compatibility, but this catalog row is packaged for PostgreSQL 14-18.
- `pg_net` supports only one database per cluster through `pg_net.database_name`.
- v0.20.3 is a worker/maintenance release: it flushes pgstat counters for autovacuum visibility and reports worker activity to `pg_stat_activity`; no new SQL request API was documented.
