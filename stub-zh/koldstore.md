## 用法

来源：

- [官方上游 README](https://github.com/kalamdb/koldstore/blob/2a50565936bc97377a439dcbb4cc2ae5c07db1a5/README.md)
- [官方扩展控制文件 (koldstore.control)](https://github.com/kalamdb/koldstore/blob/2a50565936bc97377a439dcbb4cc2ae5c07db1a5/crates/pg_koldstore/koldstore.control)
- [官方扩展 SQL (koldstore--0.1.0.sql)](https://github.com/kalamdb/koldstore/blob/2a50565936bc97377a439dcbb4cc2ae5c07db1a5/crates/pg_koldstore/sql/koldstore--0.1.0.sql)

`koldstore` — PostgreSQL 分层存储，将历史数据移至 Parquet，同时保持原始表完全可查询，并支持更新和删除操作。使用它来实现相应的分析或存储工作流。上游明确表示该项目尚未准备好用于生产环境。

### 核心工作流

```sql
CREATE EXTENSION koldstore;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `koldstore.internal_apply_flush_row_counts(p_table_oid oid, p_mirror_pruned bigint, p_hot_pruned bigint, p_cold_rows_added bigint)` 是一个扩展函数，返回 `void`。
- `koldstore.internal_bump_row_counts(p_table_oid oid, p_hot_delta bigint, p_mirror_delta bigint)` 是一个扩展函数，返回 `void`。
- `koldstore.internal_ensure_manifest_row(p_table_oid oid)` 是一个扩展函数，返回 `void`。
- `koldstore.internal_refresh_row_counts(p_table_oid oid, p_hot_rows bigint, p_mirror_rows bigint)` 是一个扩展函数，返回 `void`。
- `koldstore.change_event` 是一个扩展定义的类型。
- `koldstore.dml_result` 是一个扩展定义的类型。
- `koldstore.managed_table_info` 是一个扩展定义的类型。
- `koldstore.async_mirror_state` 是一个由扩展安装或管理的表。
- `koldstore.cold_segment_stats` 是一个由扩展安装或管理的表。
- `koldstore.cold_segments` 是一个由扩展安装或管理的表。
- `koldstore.jobs` 是一个由扩展安装或管理的表。
- `koldstore.manifest` 是一个由扩展安装或管理的表。
- `koldstore.schemas` 是一个由扩展安装或管理的表。
- `koldstore.storage` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 该目录记录版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 上游明确表示该项目尚未准备好用于生产环境。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
