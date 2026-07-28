## Usage

Sources:

- [Official upstream README](https://github.com/aintliy/timefluxplus/blob/6cb43ce01a256bf22c5b3fb9ed0014cbf3084ef1/README.md)
- [Official extension control file (timeflux.control)](https://github.com/aintliy/timefluxplus/blob/6cb43ce01a256bf22c5b3fb9ed0014cbf3084ef1/timeflux.control)
- [Official extension SQL (timeflux--1.0.sql)](https://github.com/aintliy/timefluxplus/blob/6cb43ce01a256bf22c5b3fb9ed0014cbf3084ef1/timeflux--1.0.sql)

`timeflux` — TimeFlux+: Advanced Time-Series Storage Engine. Use it for the corresponding scheduling, temporal, or time-series workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION timeflux;

INSERT → Hook 拦截 → Buffer 排序 → WAL → 异步 Flush → 分区表
                                              ↓
                              RANGE+HASH 两层分区 (热/温/冷 三级生命周期)
                                              ↓
                              MOT 内存热层 ←→ 行存温层 ←→ 列存冷层
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `tfp_mot.action_log(since_us bigint DEFAULT 0)` is an extension function and returns `TABLE`.
- `tfp_mot.chunk_mirror_status(ht_oid regclass)` is an extension function and returns `TABLE`.
- `tfp_mot.coordinator_status()` is an extension function and returns `text`.
- `tfp_mot.ensure_chunk_mirrors(source regclass)` is an extension function and returns `integer`.
- `tfp_mot.get_flush_apply_batch_rows()` is an extension function and returns `int`.
- `tfp_mot.guc_show()` is an extension function and returns `TABLE`.
- `tfp_mot.install_mirror_write_trigger(src regclass)` is an extension function and returns `boolean`.
- `tfp_mot.mirror_apply_stats()` is an extension function and returns `TABLE`.
- `tfp_mot.mirror_list()` is an extension function and returns `TABLE`.
- `tfp_mot.mirror_write()` is an extension function and returns `trigger`.
- `tfp_mot.refresh_cache()` is an extension function and returns `boolean`.
- `tfp_mot.register_table(source regclass)` is an extension function and returns `boolean`.
- `tfp_mot.remap_source_oid(old_source oid, new_source oid)` is an extension function and returns `void`.
- `tfp_mot.set_flush_apply_batch_rows(new_value int)` is an extension function and returns `int`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
