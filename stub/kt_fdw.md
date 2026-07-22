## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/README.md)
- [Version 0.0.1 SQL objects](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/sql/kt_fdw.sql)
- [FDW implementation](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/src/kt_fdw.c)
- [Optional Kyoto Tycoon transaction procedures](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/transactions.lua)

`kt_fdw` maps a Kyoto Tycoon key-value database to a writable PostgreSQL foreign table. It supports reads plus insert, update, and delete, but the upstream repository is archived and the code targets old PostgreSQL and Kyoto Tycoon APIs. Rebuild and integration-test it for the exact versions in an isolated environment before considering use.

### Core Workflow

Define a server endpoint and a foreign table with exactly two columns: the key first and the value second.

```sql
CREATE EXTENSION kt_fdw;

CREATE SERVER kt_cache
  FOREIGN DATA WRAPPER kt_fdw
  OPTIONS (host '127.0.0.1', port '1978', timeout '5');

CREATE USER MAPPING FOR CURRENT_USER SERVER kt_cache;

CREATE FOREIGN TABLE cache_entry (
  key text,
  value text
) SERVER kt_cache;

SELECT value FROM cache_entry WHERE key = 'session:42';
INSERT INTO cache_entry VALUES ('session:42', '{"state":"ready"}');
UPDATE cache_entry SET value = '{"state":"done"}' WHERE key = 'session:42';
DELETE FROM cache_entry WHERE key = 'session:42';
```

The only reviewed server options are `host` (default `127.0.0.1`), `port` (default `1978`), and `timeout` (default `-1`). There are no extension options for authentication, TLS, or a remote database number, so keep the endpoint on a trusted private network.

### Query and Write Semantics

A simple equality predicate on the first column can become one remote key lookup. Other scans walk the entire remote cursor and leave filtering to PostgreSQL. Both columns must be non-NULL. Insert uses Kyoto Tycoon's add operation and fails if the key already exists; update can change only the value, not the key.

Transactional writes require compiling the extension with its transaction support, building Kyoto Tycoon with Lua, and loading the supplied `transactions.lua` procedures on the server. Even then, savepoints and prepared transactions are unsupported. Without that setup, remote writes do not follow PostgreSQL transaction rollback. Treat cross-system failures and retries as application-visible consistency events.
