## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/postgrespro/aqo/blob/4a7fcd7291a8c4209c9102d1736f9f754d311cf2/README.md)
- [Extension control file](https://github.com/postgrespro/aqo/blob/4a7fcd7291a8c4209c9102d1736f9f754d311cf2/aqo.control)

`aqo` augments PostgreSQL cardinality estimation with models learned from prior executions of normalized query shapes. It targets expensive queries whose bad row estimates lead to bad plans; it is experimental optimizer research, not a guaranteed accelerator.

The reviewed upstream requires a PostgreSQL-version-specific core patch and a complete server rebuild, plus cluster-wide preload:

```conf
shared_preload_libraries = 'aqo'
```

Restart PostgreSQL, create the extension, and train only a chosen query inside a controlled transaction:

```sql
CREATE EXTENSION aqo;
BEGIN;
SET aqo.mode = 'learn';
EXPLAIN ANALYZE SELECT * FROM fact JOIN dim USING (dim_id) WHERE dim.kind = 'x';
SET aqo.mode = 'controlled';
COMMIT;
```

`controlled` is the upstream production default for ignoring unknown query shapes. `learn` records them; `intelligent` auto-tunes; `forced` shares a common model; `disabled` bypasses AQO. Inspect `aqo_queries` and `aqo_query_texts`, and use `aqo.show_details` or `aqo.show_hash` only for diagnosis.

Training executes the real query and can add CPU, memory, storage, and latency. Plans can regress as data distributions change. AQO does not support temporary objects, and replicas cannot collect training statistics. Pin the exact patched server branch, benchmark representative parameters, retain a fast disable path, and test failover, upgrade, model cleanup, and backup/restore before cluster-wide use.
