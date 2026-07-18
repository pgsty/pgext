## Usage

Sources:

- [Official IMCS user guide](https://pgxn.org/dist/imcs/user_guide.html)
- [Pinned extension control file](https://github.com/knizhnik/imcs/blob/940964df4f3d03ea67ebe868691a17b41c89218b/imcs.control)
- [Official PGXN distribution page](https://pgxn.org/dist/imcs/)

`imcs` adds an in-memory columnar store for analytic time-series workloads. Ordinary tables remain the durable source; selected columns are loaded into shared memory, and generated time-series accessors apply vectorized operations, compression, and optional parallel execution.

Load the module at server start and restart PostgreSQL after changing the setting:

```conf
shared_preload_libraries = '$libdir/imcs'
```

For a table named `Quote` with timestamp key `Day` and series identifier `Symbol`, the guide demonstrates the following workflow:

```sql
CREATE EXTENSION imcs;
SELECT cs_create('Quote', 'Day', 'Symbol');
SELECT Quote_load();
SELECT cs_max(Close) FROM Quote_get('IBM');
```

`cs_create()` generates load, append, delete, trigger, and accessor objects for the selected table. Data loaded from a table is sorted by timestamp; incremental or trigger-driven inserts must remain timestamp-ascending or an out-of-order error is raised. The store cannot represent SQL `NULL` directly, and its supported element types are primarily fixed-size values; review dictionary mapping before using variable-length strings.

Shared-memory contents must be reloaded after restart unless the documented autoload mode is used. Size `imcs.shmem_size` and related limits before loading data, and test the first-access latency of autoload. The pinned control reports version `1.1`, while the latest listed PGXN distribution is `0.1.7`; treat those as separate version lines and verify the exact source, license, server compatibility, and upgrade path.
