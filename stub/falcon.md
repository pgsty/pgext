## Usage

Sources:

- [Official upstream README](https://github.com/falcon-infra/falconfs/blob/0e4f460c2600352529bb9a320fb667121f277658/README.md)
- [Official extension control file (falcon.control)](https://github.com/falcon-infra/falconfs/blob/0e4f460c2600352529bb9a320fb667121f277658/falcon/falcon.control)
- [Official extension SQL (falcon--1.0.sql)](https://github.com/falcon-infra/falconfs/blob/0e4f460c2600352529bb9a320fb667121f277658/falcon/falcon--1.0.sql)

`falcon` — FalconFS is a high-performance distributed file system (DFS) designed for AI workloads. Use it for the corresponding analytical or storage workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION falcon;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_catalog.falcon_acquire_hash_lock(IN path cstring, IN parentId bigint, IN lockmode bigint)` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_build_shard_table(shard_count int)` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_clear_all_data_func()` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_clear_cached_relation_oid_func()` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_clear_user_data_func()` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_create_distributed_data_table()` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_create_distributed_data_table_by_range_point(range_point int)` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_create_kvmeta_table()` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_create_slice_table()` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_delete_foreign_server(server_id int)` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_drop_distributed_data_table_by_range_point(range_point int)` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_foreign_server_test(mode cstring)` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_insert_foreign_server(server_id int, server_name cstring, host cstring, port int, is_local bool, user_name cstring)` is an extension function and returns `INTEGER`.
- `pg_catalog.falcon_meta_call_by_serialized_data(type int, count int, param bytea)` is an extension function and returns `bytea`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
