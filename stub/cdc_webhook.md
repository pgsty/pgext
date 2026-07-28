## Usage

Sources:

- [Official upstream README](https://github.com/hasura/postgres-cdc-plugin/blob/e08af17fe6f2a71fb1f181d1647e9bcff27b7f58/README.md)
- [Official extension control file (cdc_webhook.control)](https://github.com/hasura/postgres-cdc-plugin/blob/e08af17fe6f2a71fb1f181d1647e9bcff27b7f58/cdc_webhook.control)
- [Official extension SQL (cdc_webhook--1.0.sql)](https://github.com/hasura/postgres-cdc-plugin/blob/e08af17fe6f2a71fb1f181d1647e9bcff27b7f58/cdc_webhook--1.0.sql)

`cdc_webhook` — A PostgreSQL extension that enables Change Data Capture (CDC) by sending webhook notifications for database changes. This extension is written in C and SQL, allowing real-time monitoring of INSERT, UPDATE, and DELETE operations on specified tables. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION cdc_webhook;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `call_webhook(payload JSONB, webhook_url TEXT, headers JSONB, timeout INT, cancel_on_failure BOOLEAN, retry_number INT, retry_interval INT, retry_backoff TEXT)` is an extension function and returns `void`.
- `cdc_webhook.credentials_audit_trigger()` is an extension function and returns `TRIGGER`.
- `cdc_webhook.event_log_audit_trigger()` is an extension function and returns `TRIGGER`.
- `create_event_trigger` is an extension function.
- `cdc_webhook.credentials` is a table installed or managed by the extension.
- `cdc_webhook.event_log` is a table installed or managed by the extension.
- `cdc_webhook` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
