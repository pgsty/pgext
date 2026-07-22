## Usage

Sources:

- [Official upstream documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/using-the-index-adviser-extension-on-an-apsaradb-rds-for-postgresql-instance)
- [Official primary documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`index_adviser` 2.0 is an Alibaba Cloud RDS for PostgreSQL extension that estimates useful B-tree indexes from planned or executed queries. It is provider-managed software, and its recommendations are estimates that must be reviewed before executing the generated DDL.

### Core Workflow

Create the extension, then either load it for one session or configure `shared_preload_libraries` and restart the instance for all sessions.

```sql
CREATE EXTENSION index_adviser;
LOAD 'index_adviser';

EXPLAIN SELECT * FROM orders WHERE customer_id = 42;
SELECT * FROM index_advisory;
SELECT show_index_advisory(NULL);
```

Plain `EXPLAIN` records advice without running the query. Queries executed normally are analyzed as they run. `index_advisory` stores raw per-query estimates; `show_index_advisory(pid)` formats session recommendations as `CREATE INDEX` statements; `select_index_advisory` summarizes all recorded sessions.

### Requirements and Limits

The current vendor documentation requires RDS PostgreSQL minor engine version 20230830 or later when creating or recreating the extension. Do not use it in read-only transactions. It recommends only single-column or composite B-tree indexes, not GIN, GiST, or hash indexes. Validate workload representativeness, write amplification, storage, duplicate indexes, locking, and maintenance cost before applying any suggestion.
