## Usage

Sources:

- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_vectorize/alohadb_vectorize--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vectorize/alohadb_vectorize--1.0.sql)
- [contrib/alohadb_vectorize/alohadb_vectorize.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vectorize/alohadb_vectorize.c)
- [contrib/alohadb_vectorize/alohadb_vectorize.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_vectorize/alohadb_vectorize.control)

`alohadb_vectorize` 1.0 wraps analytical SQL with local memory, parallel-planning and JIT settings in AlohaDB. Its implementation uses PostgreSQL SPI and does not introduce a new vectorized execution engine.

### Core Workflow

```sql
CREATE EXTENSION alohadb_vectorize;
BEGIN;
SELECT * FROM vectorize_status();
SELECT * FROM vectorize_query('SELECT count(*) FROM pg_class') AS t(n bigint);
SELECT * FROM vectorize_explain('SELECT count(*) FROM pg_class');
ROLLBACK;
```

### Operational Boundaries

`vectorize_query` returns a record set whose column definitions the caller must supply. `vectorize_explain` runs EXPLAIN ANALYZE, so it executes the supplied query. `vectorize_status` shows relevant settings and `vectorize_benchmark` repeats a query and returns timings.

The wrappers apply `SET LOCAL` settings, including increased work memory and parallel-worker settings, for the transaction. Run experiments in a bounded transaction and consider aggregate memory use under concurrent queries. Do not pass untrusted SQL to these functions. Installation requires a superuser; no preload or restart is declared. The reviewed distribution belongs to the PostgreSQL 18-based AlohaDB fork; portable stock-PostgreSQL package support has not been established.
