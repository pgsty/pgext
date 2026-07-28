## 用法

来源：

- [官方上游 README](https://github.com/yash7312/dbis_assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/postgresql/contrib/temporal_rtree/README)
- [官方扩展控制文件 (temporal_rtree.control)](https://github.com/yash7312/dbis_assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/postgresql/contrib/temporal_rtree/temporal_rtree.control)
- [官方扩展 SQL (temporal_rtree--1.0.sql)](https://github.com/yash7312/dbis_assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/postgresql/contrib/temporal_rtree/temporal_rtree--1.0.sql)

`temporal_rtree` — 一个实现新的索引访问方法（temporal_rtree），优化用于时间范围索引的 PostgreSQL 扩展。使用它来进行相应的调度、时间或时间序列工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION temporal_rtree;

CREATE INDEX idx_temporal ON temporal_data
  USING temporal_rtree (temporalbox(attr, valid_period) temporal_cube_ops);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `temporal_rtree_handler(internal)` 是一个扩展函数，返回 `index_am_handler`。
- `temporal_rtree_hook_reset()` 是一个扩展函数，返回 `void`。
- `temporal_rtree_hook_stats()` 是一个扩展函数，返回 `TABLE`。
- `FAMILY` 是一个扩展定义的操作符。
- `temporal_cube_ops` 是一个扩展定义的操作符类。
- `temporal_rtree` 是一个扩展定义的访问方法。
- `temporal_tsrange_ops` 是一个扩展定义的操作符类。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
