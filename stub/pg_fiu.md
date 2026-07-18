## Usage

Sources:

- [Upstream README](https://github.com/jasonmp85/pg_fiu/blob/f2e0afc6963e49f6fac31d4bb1d169f8a4eaf312/README.md)
- [SQL fault-point API](https://github.com/jasonmp85/pg_fiu/blob/f2e0afc6963e49f6fac31d4bb1d169f8a4eaf312/sql/pg_fiu.sql)
- [Hook and libfiu implementation](https://github.com/jasonmp85/pg_fiu/blob/f2e0afc6963e49f6fac31d4bb1d169f8a4eaf312/src/pg_fiu.c)
- [Build configuration](https://github.com/jasonmp85/pg_fiu/blob/f2e0afc6963e49f6fac31d4bb1d169f8a4eaf312/Makefile)

pg_fiu 0.0.1 is an unfinished 2018 fault-injection prototype built around libfiu and PostgreSQL parser, executor, utility, and transaction hooks. It belongs only on a disposable test cluster.

### Isolated setup

Preload the library and restart the disposable server:

```conf
shared_preload_libraries = 'pg_fiu'
```

Install into a dedicated schema and remove public function access before inspection:

```sql
CREATE SCHEMA fiu_lab;
CREATE EXTENSION pg_fiu WITH SCHEMA fiu_lab;
REVOKE EXECUTE ON ALL FUNCTIONS IN SCHEMA fiu_lab FROM PUBLIC;
SELECT * FROM fiu_lab.failure_points();
```

Do not enable an injection point until the exact linked binary and recovery procedure have been reviewed.

### Caveats

- Fault injection deliberately changes C return values and auxiliary pointers inside a backend. A wrong name, type, callback, stack target, or return value can corrupt memory, crash processes, or damage data.
- The hooks activate only when loaded through shared_preload_libraries, affecting every new backend. The next-transaction switch is superuser-only, but the SQL definition-creation functions are public unless revoked.
- SQL functions that mutate backend-global failure state are incorrectly declared immutable. Definitions persist for the life of the backend, and no SQL removal or clear function is provided.
- The target code must contain matching libfiu instrumentation or compatible stack symbols. Upstream provides no working usage example; its regression file calls a function that does not exist.
- The Makefile and one platform branch contain absolute developer paths. The 2018 hook signatures and backend internals require porting for current PostgreSQL majors.
- Never preload pg_fiu on production, a shared test service, or a cluster containing irreplaceable data.
