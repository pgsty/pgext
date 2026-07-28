## 用法

来源：

- [Official database.dev 包页面](https://database.dev/dventimi/pg_partition_magician)

`dventimi@pg_partition_magician` — Pure-SQL 在线 RANGE 分区管理器（时间 / ID / UUIDv7），由 pg_cron 驱动。请使用它来对应相应的分析或存储工作流。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "dventimi@pg_partition_magician";
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `pgpm._adopt` 是一个扩展函数。
- `pgpm._create_partition` 是一个扩展函数。
- `pgpm._decode` 是一个扩展函数。
- `pgpm._encode` 是一个扩展函数。
- `pgpm._frontier_native` 是一个扩展函数。
- `pgpm._grid_floor` 是一个扩展函数。
- `pgpm._grid_next` 是一个扩展函数。
- `pgpm._native_gt` 是一个扩展函数。
- `pgpm._native_type` 是一个扩展函数。
- `pgpm._part_name` 是一个扩展函数。
- `pgpm._ts_to_uuid` 是一个扩展函数。
- `pgpm._uuid_to_ts` 是一个扩展函数。
- `pgpm.adopt(p_parent regclass, p_control name, p_interval interval, p_premake int default 4, p_retention interval default null, p_keep_default boolean default true, p_drain_batch int default 5000, p_anchor timestamptz default '2000-01-01 00:00:00+00', p_paused boolean def…)` 是一个扩展函数并返回 `regclass`。
- `pgpm.adopt_by_id(p_parent regclass, p_control name, p_step bigint, p_premake int default 4, p_retention bigint default null, p_keep_default boolean default true, p_drain_batch int default 5000, p_anchor bigint default 0, p_paused boolean default true, p_incoming_fks text defau…)` 是一个扩展函数并返回 `regclass`。

### 要求与注意事项

- 该目录记录版本为 `0.1.0`。
- 先安装确认的扩展依赖项：`pg_cron`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 身份之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
