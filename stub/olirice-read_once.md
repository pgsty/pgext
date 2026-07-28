## Usage

Sources:

- [Official database.dev package page](https://database.dev/olirice/read_once)

`olirice-read_once` — Send messages that can only be read once. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "olirice-read_once";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `read_message(id uuid)` is an extension function and returns `text`.
- `send_message(contents text)` is an extension function and returns `uuid`.
- `read_once` is a schema created by the extension.

### Requirements and Caveats

- The catalog records version `0.3.2`.
- Install the confirmed extension dependencies first: `pg_cron`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
