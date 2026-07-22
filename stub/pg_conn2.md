## Usage

Sources:

- [Official README](https://github.com/serpent7776/pg_conn2/blob/296c2d7086dd0021a636c57964a454961cecfead/README.md)
- [Extension SQL API](https://github.com/serpent7776/pg_conn2/blob/296c2d7086dd0021a636c57964a454961cecfead/pg_conn2--0.1.sql)
- [libpq connection implementation](https://github.com/serpent7776/pg_conn2/blob/296c2d7086dd0021a636c57964a454961cecfead/pg_conn2.c)

`pg_conn2` 0.1 opens a second libpq connection from a PostgreSQL backend to another database on the same local server and executes SQL through an opaque `pg_conn2` handle. It is useful only for tightly controlled administrative experiments: it returns no query rows and its remote transaction is not atomic with the caller's local transaction.

### Core Workflow

The implementation hard-codes the Unix socket directory `/var/run/postgresql` and user `postgres`; the only supplied connection field is the database name:

```sql
CREATE EXTENSION pg_conn2;

SELECT pg_conn2exec(
  pg_conn2make('analytics'),
  'CREATE TABLE IF NOT EXISTS import_log(ts timestamptz DEFAULT now())'
);
```

For several statements that must succeed or fail together on the remote connection:

```sql
SELECT pg_conn2exec_many(
  pg_conn2make('analytics'),
  ARRAY[
    'INSERT INTO import_log DEFAULT VALUES',
    'ANALYZE import_log'
  ]
);
```

`pg_conn2exec_many` sends `BEGIN`, each array element, and `COMMIT` remotely. On any error it sends `ROLLBACK` and raises an error; its Boolean result is therefore `true` on success rather than a normally observable `false` failure status. `pg_conn2close(handle)` closes early; otherwise the connection closes at local top-level transaction commit or abort.

### Boundaries and Risks

`pg_conn2exec` accepts statements that return tuples but discards every result. Each statement normally autocommits on the remote connection. Even `pg_conn2exec_many` commits remotely before the surrounding local transaction finishes, so a later local rollback cannot undo remote work. Avoid self-referential locking patterns that can deadlock the two sessions.

Authentication follows local `pg_hba.conf` and the PostgreSQL server operating-system context. The database name is interpolated into libpq connection text without escaping; never pass untrusted input because it can inject additional connection parameters. The handle is transaction-scoped, cannot be meaningfully serialized, and must not be stored in a table. Production code should prefer a maintained FDW, `dblink`, or application-side connection with explicit credentials, result handling, and transaction semantics.
