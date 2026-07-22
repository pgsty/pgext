## Usage

Sources:

- [Official README](https://github.com/asotolongo/rep_fdw/blob/990fb7b36c5334db9dabf2ccfd1d8502bd057cad/README)
- [Official extension SQL](https://github.com/asotolongo/rep_fdw/blob/990fb7b36c5334db9dabf2ccfd1d8502bd057cad/rep_fdw--0.1.0.sql)
- [Official extension control file](https://github.com/asotolongo/rep_fdw/blob/990fb7b36c5334db9dabf2ccfd1d8502bd057cad/rep_fdw.control)

`rep_fdw` version `0.1.0` is an experimental trigger-based row copier built on `postgres_fdw`. It synchronously mirrors INSERT, UPDATE, and DELETE operations to a generated foreign table, but it has no durable queue, conflict protocol, schema synchronization, or robust SQL quoting and should not be used as a production replication system.

### Core Workflow

The extension creates objects in the `rep_fdw` schema and requires `postgres_fdw`:

```sql
CREATE EXTENSION postgres_fdw;
CREATE EXTENSION rep_fdw;

SELECT rep_fdw.create_server('remote_db', '127.0.0.1', '5433', 'app');
SELECT rep_fdw.create_usermap('replica_user', 'secret', 'remote_db');
SELECT rep_fdw.create_f_table('public', 'orders', 'remote_db');
SELECT rep_fdw.generar_trigger('public.orders');
```

`create_f_table(text, text, text)` introspects a local table and creates a foreign table named with the `f_` prefix. `generar_trigger(text)` installs an AFTER-row trigger whose `tr_ftable()` handler writes the old or new row to the remote table during local DML.

### Correctness and Security Boundary

`create_usermap(text, text, text)` creates a user mapping for PUBLIC and places the supplied password in PostgreSQL catalog options. This is much broader than a per-role mapping; restrict catalog visibility, ownership, and function execution, and prefer manually managed least-privilege mappings.

The trigger constructs identifiers, values, and predicates through string concatenation. Quotes, NULL values, unusual types, or hostile input can break statements or alter their meaning. UPDATE and DELETE match all old column values rather than a declared primary key, so duplicates and concurrent changes are ambiguous.

Remote writes are synchronous: an FDW error can abort local DML, while network partitions and transaction recovery still require explicit testing. Use `rep_fdw` only with disposable data to study the approach. Use maintained logical replication or a queue-based design when correctness, replay, monitoring, or recovery matters.
