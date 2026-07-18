## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/pg_oltp_bench/blob/fbf92a58e12f6e155a4d2c09969063dc26e8a2a1/pg_oltp_bench.control)

`pg_oltp_bench` — Helper function and pgbench scripts for a sysbench-style OLTP benchmark.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_oltp_bench";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_oltp_bench';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
