
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION "Snehil_Shah@pg_dispatch";
> SELECT pgdispatch.fire('SELECT pg_sleep(40);');
> SELECT pgdispatch.snooze('SELECT pg_sleep(20);', '20 seconds');
> ```
>
> Sources: [README](https://github.com/Snehil-Shah/pg_dispatch), [database.dev page](https://database.dev/Snehil_Shah/pg_dispatch)

`pg_dispatch` is an asynchronous SQL dispatcher for PostgreSQL. It is designed as a TLE-compatible alternative to `pg_later`, built on top of `pg_cron`, so it can be used in environments such as Supabase and AWS RDS.

## Prerequisites

The upstream README lists:

- PostgreSQL 13 or newer
- `pg_cron` 1.5.0 or newer
- `pgcrypto`

## Installation

The documented TLE installation path is:

```sql
SELECT dbdev.install(Snehil_Shah@pg_dispatch);
CREATE EXTENSION "Snehil_Shah@pg_dispatch";
```

The README warns that the extension installs into the `pgdispatch` schema and can collide with an existing schema of the same name.

## Functions

### `pgdispatch.fire(command text)`

Dispatch an SQL command for asynchronous execution:

```sql
SELECT pgdispatch.fire('SELECT pg_sleep(40);');
```

### `pgdispatch.snooze(command text, delay interval)`

Dispatch a delayed SQL command:

```sql
SELECT pgdispatch.snooze('SELECT pg_sleep(20);', '20 seconds');
```

The README notes that the delay is scheduled asynchronously and does not block the caller's main transaction.

## Use Cases

The project positions itself for database-native async side effects, especially in PL/pgSQL or trigger-based workflows. Its example use case is moving expensive notification or analytics work out of an `AFTER INSERT` trigger so the main RPC or application request can return sooner.
