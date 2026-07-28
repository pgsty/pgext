## 用法

来源：

- [官方上游 README](https://github.com/ai-blaise/citus/blob/b2bbe2e0d29ec70eb32e7e44c7450d03aaa52659/companion/README.md)
- [官方扩展控制文件 (ai_blaise_citus.control)](https://github.com/ai-blaise/citus/blob/b2bbe2e0d29ec70eb32e7e44c7450d03aaa52659/companion/ai_blaise_citus.control)
- [官方实现源代码](https://github.com/ai-blaise/citus/blob/b2bbe2e0d29ec70eb32e7e44c7450d03aaa52659/companion/src/lib.rs)

`ai_blaise_citus` — Rust pgrx 同伴扩展，用于协调 Citus、TimescaleDB、捆绑扩展和侧车。当应用程序需要此特定数据库功能时，请使用它。上游明确表示该项目尚未准备好生产使用。

### 核心工作流

```sql
CREATE EXTENSION ai_blaise_citus;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `add_compression_policy_distributed` 是一个扩展函数。
- `add_continuous_aggregate_distributed` 是一个扩展函数。
- `add_reorder_policy_distributed` 是一个扩展函数。
- `add_retention_policy_distributed` 是一个扩展函数。
- `companion_feature_status()` 是一个扩展函数。
- `distribute_hypertable` 是一个扩展函数。
- `time_range_shard_pruner` 是一个扩展函数。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 上游明确表示该项目尚未准备好生产使用。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
