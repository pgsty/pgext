## Usage

Sources:

- [Version 14.0 manual](https://github.com/ossc-db/pg_dbms_stats/blob/REL14_0/doc/pg_dbms_stats-en.md)
- [Version 14.0 object reference](https://github.com/ossc-db/pg_dbms_stats/blob/REL14_0/doc/objects-en.md)
- [Extension control file](https://github.com/ossc-db/pg_dbms_stats/blob/REL14_0/pg_dbms_stats.control)

`pg_dbms_stats` 14.0 stores, restores, and locks planner statistics so a PostgreSQL 14 optimizer can continue seeing a chosen statistics snapshot. Use it to reduce plan changes caused by `ANALYZE` or distribution drift; it does not pin a plan or freeze table data.

### Enable the Extension

Install and register `pg_dbms_stats` in every database where the library will intercept planner catalog reads. Loading the library in a database without the `dbms_stats` schema can make ordinary statements fail, so complete per-database registration before enabling cluster-wide preload.

```sql
CREATE EXTENSION pg_dbms_stats;

SHOW pg_dbms_stats.use_locked_stats;
```

For automatic loading, add it to `shared_preload_libraries` and restart PostgreSQL.

```conf
shared_preload_libraries = 'pg_dbms_stats'
```

For controlled session testing, the upstream manual also permits `LOAD 'pg_dbms_stats'` after the extension exists in that database.

### Backup, Restore, and Lock

```sql
SELECT dbms_stats.backup_database_stats('before scheduled analyze');

ANALYZE;

SELECT dbms_stats.restore_database_stats(now() - interval '1 day');

SELECT *
FROM dbms_stats.backup_history
ORDER BY time DESC;

SET pg_dbms_stats.use_locked_stats = off;
EXPLAIN SELECT * FROM orders WHERE customer_id = 42;
SET pg_dbms_stats.use_locked_stats = on;
```

Restore also locks the restored statistics. Database-, schema-, table-, and column-level variants are available for backup, restore, lock, unlock, and import. Target objects need existing PostgreSQL statistics before they can be meaningfully backed up or locked.

### Important Objects

- `dbms_stats.backup_history` records backup sets.
- `dbms_stats.relation_stats_effective` and `dbms_stats.column_stats_effective` show the values currently presented to the planner.
- `dbms_stats.stats` summarizes locked statistics.
- `dbms_stats.clean_up_stats()` removes locked rows left behind after tables or columns are dropped.
- Export scripts and import functions move statistics through server-side files; imports require an absolute path readable by the PostgreSQL server account.

### Operational Boundaries

Locked statistics do not guarantee an unchanged plan: indexes, schema, planner GUCs, PostgreSQL code, parameter values, and table size still affect planning. Monitor query plans and unlock quickly if the fixed snapshot becomes harmful.

Dropped objects do not automatically remove their locked rows, so schedule `dbms_stats.clean_up_stats()`. Statistics containing polymorphic `anyarray` values need the manual's special dump/restore sequence. Version 14.0 is the PostgreSQL 14 branch; use the matching upstream branch and regression-test upgrades rather than loading it on another major version.
