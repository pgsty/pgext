## Usage

Sources:

- [Official upstream README](https://github.com/arssher/hodgepodge/blob/31e1a36a2e27d7402b4f15570ec1a3e1bbb53cef/readme.md)
- [Official extension control file (shardman.control)](https://github.com/arssher/hodgepodge/blob/31e1a36a2e27d7402b4f15570ec1a3e1bbb53cef/ext/shardman.control)
- [Official extension SQL (shardman--0.0.1.sql)](https://github.com/arssher/hodgepodge/blob/31e1a36a2e27d7402b4f15570ec1a3e1bbb53cef/ext/shardman--0.0.1.sql)

`shardman` — Experimental sharding, rebalance, and cross-node lock-graph prototype. Use it for the corresponding analytical or storage workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION shardman;

create table pt (id serial, payload real) partition by hash(id);
select shardman.hash_shard_table('pt', 10)
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `bcst_all_sql(cmd text)` is an extension function and returns `void`.
- `bcst_sql(cmd text)` is an extension function and returns `void`.
- `deny_access()` is an extension function and returns `trigger`.
- `drop_repslot(slot_name text, with_force bool default true)` is an extension function and returns `void`.
- `eliminate_sub(subname name)` is an extension function and returns `void`.
- `ex_sql(rgid int, cmd text)` is an extension function and returns `void`.
- `hash_shard_table(relid regclass, nparts int, colocate_with regclass = null)` is an extension function and returns `void`.
- `is_subscription_ready(sname text)` is an extension function.
- `part_moved(relid regclass, pnum int, src_rgid int, dst_rgid int)` is an extension function and returns `void`.
- `postgres_fdw_handler()` is an extension function and returns `fdw_handler`.
- `postgres_fdw_validator(text[], oid)` is an extension function and returns `void`.
- `rebalance_cleanup()` is an extension function and returns `void`.
- `rebalance_cleanup_local()` is an extension function and returns `void`.
- `restore_foreign_tables()` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
