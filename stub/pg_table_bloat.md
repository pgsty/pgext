## Usage

Sources:

- [Official upstream README](https://github.com/lucasfrederico/pg_table_bloat/blob/686385d88dfc8318aa57351ae260c87399f70868/README.md)
- [Official extension control file (pg_table_bloat.control)](https://github.com/lucasfrederico/pg_table_bloat/blob/686385d88dfc8318aa57351ae260c87399f70868/pg_table_bloat.control)
- [Official implementation source](https://github.com/lucasfrederico/pg_table_bloat/blob/686385d88dfc8318aa57351ae260c87399f70868/src/lib.rs)

`pg_table_bloat` — > A PostgreSQL extension that estimates table bloat from pg_class and pg_stats > — no pgstattuple required, no per-page lock, microseconds on tables of any size. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_table_bloat;

CREATE TABLE users (
    id bigserial PRIMARY KEY,
    email varchar(255) UNIQUE NOT NULL,
    payload jsonb
);
INSERT INTO users (email, payload)
SELECT 'user'||g||'@example.com', jsonb_build_object('data', repeat('x', 500))
FROM generate_series(1, 50000) g;
DELETE FROM users WHERE id > 500;   -- delete 99% of rows
ANALYZE users;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
