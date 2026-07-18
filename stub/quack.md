## Usage

Sources:

- [Version 0.0.1 README](https://github.com/hydradatabase/pg_quack/blob/5a49c79fd363bb27dcf6205fb3dd5d788758382b/README.md)
- [Extension SQL](https://github.com/hydradatabase/pg_quack/blob/5a49c79fd363bb27dcf6205fb3dd5d788758382b/quack--0.0.1.sql)
- [Extension control file](https://github.com/hydradatabase/pg_quack/blob/5a49c79fd363bb27dcf6205fb3dd5d788758382b/quack.control)
- [Current main-branch notice](https://github.com/hydradatabase/pg_quack/blob/6213d66399c2b0fb8efd6640cfff2f8c6976a1e5/README.md)

`quack` version `0.0.1` is a proof-of-concept table access method that embeds DuckDB in PostgreSQL. Ensure the directory selected by `quack.data_dir` exists and is writable by the PostgreSQL operating-system account, then create and query a table using the access method:

```sql
CREATE EXTENSION quack;

CREATE TABLE quack_events (
  id bigint,
  payload text
) USING quack;

INSERT INTO quack_events VALUES (1, 'first event');
SELECT * FROM quack_events;
```

### Proof-of-concept limits

The published branch supports only PostgreSQL 14 and 15, basic data types, and `COPY`, `INSERT`, and `SELECT`. Only one connection can execute `INSERT` or `SELECT` against a quack table at a time, and queries cannot combine quack tables with heap or other storage methods. The unarchived main branch only announces a future version and directs users back to this proof of concept. Use it for controlled evaluation, not as a general production table engine.
