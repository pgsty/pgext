## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pmpp/)
- [Official README](https://github.com/coreyhuinker/pmpp/blob/45afa3434d6de14006ce43e43412f23ad4242fd7/README.md)
- [Official function documentation](https://github.com/coreyhuinker/pmpp/blob/45afa3434d6de14006ce43e43412f23ad4242fd7/doc/pmpp.md)
- [Official extension control file](https://github.com/coreyhuinker/pmpp/blob/45afa3434d6de14006ce43e43412f23ad4242fd7/pmpp.control)

`pmpp` 1.2.3, Poor Man’s Parallel Processing, runs manually partitioned SQL statements over multiple asynchronous libpq connections and streams compatible rows back into one PostgreSQL query. Use it when you can split work into independent statements yourself; it is not an optimizer-driven parallel-query engine or a distributed transaction manager.

### Core Workflow

`pmpp.distribute` needs a local composite type that describes every returned row. A temporary table can supply that type without storing results. Pass an explicit worker count when the target server does not also have `pmpp` installed.

```sql
CREATE EXTENSION pmpp;

CREATE TEMPORARY TABLE pmpp_result(value integer);

SELECT *
FROM pmpp.distribute(
    NULL::pmpp_result,
    'dbname=' || current_database(),
    ARRAY['SELECT 1', 'SELECT 2'],
    1.0,
    2
);
```

Installation creates the role `pmpp`. Grant that role only to users who are allowed to open the configured remote connections and execute the submitted SQL.

### Main Interfaces

- `pmpp.distribute(row_type, connection, sql_list, cpu_multiplier, num_workers, setup_commands, result_format)` handles one target and returns `SETOF` the supplied row type.
- `pmpp.distribute` overloads accept `query_manifest[]`, `json`, or `jsonb` for multiple targets, with typed and `record` return variants.
- `pmpp.meta` runs commands such as DDL in parallel and reports each command’s result.
- `pmpp.broadcast` runs one row-returning query across a list of connections.
- `pmpp.num_cpus()` reports the server CPU count used by automatic worker sizing.

The connection argument may be a libpq connection string or a foreign-server name. A `postgres_fdw` server plus user mapping keeps credentials out of each function call, although the credentials still exist in PostgreSQL metadata.

### Concurrency, Transactions, and Privileges

Each worker is a separate database session. It uses the remote user’s privileges and cannot see uncommitted changes in the caller’s transaction. Returned row order follows worker completion and is not deterministic unless the outer query sorts it.

If one subquery fails, `pmpp` attempts to cancel the others, but separate sessions do not provide atomic distributed commit or rollback. Do not use independent DDL/DML tasks as though they were one transaction; design idempotency and recovery explicitly. `setup_commands` and `pmpp.meta` can execute arbitrary SQL on worker sessions, so restrict the `pmpp` role, server mappings, and accepted query text.

Worker counts translate directly into additional connections and can overload the local or remote server. Set `num_workers` deliberately, account for connection limits and per-query resource use, and benchmark the complete workload. The upstream project targets PostgreSQL 9.4-era features and was last updated in 2018, so verify build and runtime behavior on the exact target release.
