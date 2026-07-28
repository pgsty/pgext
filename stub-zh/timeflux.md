## 用法

来源：

- [官方上游 README](https://github.com/aintliy/timefluxplus/blob/6cb43ce01a256bf22c5b3fb9ed0014cbf3084ef1/README.md)
- [官方扩展控制文件 (timeflux.control)](https://github.com/aintliy/timefluxplus/blob/6cb43ce01a256bf22c5b3fb9ed0014cbf3084ef1/timeflux.control)
- [官方扩展 SQL (timeflux--1.0.sql)](https://github.com/aintliy/timefluxplus/blob/6cb43ce01a256bf22c5b3fb9ed0014cbf3084ef1/timeflux--1.0.sql)

`timeflux` — TimeFlux+: 高级时间序列存储引擎。用于相应的调度、时间或时间序列工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION timeflux;

INSERT → Hook 拦截 → Buffer 排序 → WAL → 异步 Flush → 分区表
                                              ↓
                              RANGE+HASH 两层分区 (热/温/冷 三级生命周期)
                                              ↓
                              MOT 内存热层 ←→ 行存温层 ←→ 列存冷层
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `tfp_mot.action_log(since_us bigint DEFAULT 0)` 是一个扩展函数，返回 `TABLE`。
- `tfp_mot.chunk_mirror_status(ht_oid regclass)` 是一个扩展函数，返回 `TABLE`。
- `tfp_mot.coordinator_status()` 是一个扩展函数，返回 `text`。
- `tfp_mot.ensure_chunk_mirrors(source regclass)` 是一个扩展函数，返回 `integer`。
- `tfp_mot.get_flush_apply_batch_rows()` 是一个扩展函数，返回 `int`。
- `tfp_mot.guc_show()` 是一个扩展函数，返回 `TABLE`。
- `tfp_mot.install_mirror_write_trigger(src regclass)` 是一个扩展函数，返回 `boolean`。
- `tfp_mot.mirror_apply_stats()` 是一个扩展函数，返回 `TABLE`。
- `tfp_mot.mirror_list()` 是一个扩展函数，返回 `TABLE`。
- `tfp_mot.mirror_write()` 是一个扩展函数，返回 `trigger`。
- `tfp_mot.refresh_cache()` 是一个扩展函数，返回 `boolean`。
- `tfp_mot.register_table(source regclass)` 是一个扩展函数，返回 `boolean`。
- `tfp_mot.remap_source_oid(old_source oid, new_source oid)` 是一个扩展函数，返回 `void`。
- `tfp_mot.set_flush_apply_batch_rows(new_value int)` 是一个扩展函数，返回 `int`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
